package vcenter

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/whywaita/myshoes/pkg/runner"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware"
	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/soap"
	"github.com/whywaita/shoes-agent/agent/pkg/server"
	pb "github.com/whywaita/shoes-agent/proto.go"
	"github.com/whywaita/shoes-agent/shoes-agent/pkg/backend"
	"github.com/whywaita/shoes-agent/shoes-agent/pkg/shoesprovider"

	"github.com/google/uuid"
	"github.com/vmware/govmomi"
	"google.golang.org/grpc"
)

const (
	EnvURL      = "SHOES_VCENTER_URL"
	EnvUser     = "SHOES_VCENTER_USERNAME"
	EnvPassword = "SHOES_VCENTER_PASSWORD"

	EnvDatacenterName     = "SHOES_VCENTER_DATACENTER_NAME"
	EnvDatastoreName      = "SHOES_VCENTER_DATASTORE_NAME"
	EnvSourceVMName       = "SHOES_VCENTER_SOURCE_VM_NAME"
	EnvSourceSnapshotName = "SHOES_VCENTER_SOURCE_SNAPSHOT_NAME"

	EnvNumberOfInstancePerHost = "SHOES_VCENTER_NUMBER_OF_INSTANCE_PER_HOST"
	EnvIgnoreVMNames           = "SHOES_VCENTER_IGNORE_VM_NAMES"
)

const (
	NoLimitNumberVirtualMachine = -1
)

type config struct {
	datacenterName     string
	datastoreName      string
	sourceVMName       string
	sourceSnapshotName string

	numberInstancePerHost int
	ignoreVMNames         []string
}

func load() (*url.URL, *config, error) {
	u, err := url.Parse(os.Getenv(EnvURL))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	u.User = url.UserPassword(os.Getenv(EnvUser), os.Getenv(EnvPassword))

	if os.Getenv(EnvDatacenterName) == "" {
		return nil, nil, fmt.Errorf("must set %s", EnvDatacenterName)
	}
	if os.Getenv(EnvDatastoreName) == "" {
		return nil, nil, fmt.Errorf("must set %s", EnvDatastoreName)
	}
	if os.Getenv(EnvSourceVMName) == "" {
		return nil, nil, fmt.Errorf("must set %s", EnvSourceVMName)
	}
	if os.Getenv(EnvSourceSnapshotName) == "" {
		return nil, nil, fmt.Errorf("must set %s", EnvSourceSnapshotName)
	}

	var numInstance int
	envNumInstance := os.Getenv(EnvNumberOfInstancePerHost)
	if envNumInstance == "" {
		numInstance = NoLimitNumberVirtualMachine
	} else {
		numInstance64, err := strconv.ParseUint(envNumInstance, 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse %s: %w", EnvNumberOfInstancePerHost, err)
		}
		numInstance = int(numInstance64)
	}

	ignoreVMName := os.Getenv(EnvIgnoreVMNames)
	ignoreVMNames := strings.Split(ignoreVMName, ",")

	return u, &config{
		datacenterName:     os.Getenv(EnvDatacenterName),
		datastoreName:      os.Getenv(EnvDatastoreName),
		sourceVMName:       os.Getenv(EnvSourceVMName),
		sourceSnapshotName: os.Getenv(EnvSourceSnapshotName),

		numberInstancePerHost: numInstance,
		ignoreVMNames:         ignoreVMNames,
	}, nil
}

func newClient(ctx context.Context, u url.URL) (*vmware.Client, *govmomi.Client, error) {
	password, ok := u.User.Password()
	if !ok {
		return nil, nil, fmt.Errorf("not set password")
	}
	client, err := vmware.NewClient(u.Host, u.User.Username(), password, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create a client: %w", err)
	}

	soapURL := u
	soapURL.Path = "/sdk"
	soapClient, err := govmomi.NewClient(ctx, &soapURL, true)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create a soap client: %w", err)
	}

	return client, soapClient, nil

}

// VCenter is a client of vCenter
type VCenter struct {
	Client     vmware.Client
	SOAPClient govmomi.Client

	Conf config
}

// New create a client
func New(ctx context.Context) (*VCenter, error) {
	rand.Seed(time.Now().UnixNano())

	u, c, err := load()
	if err != nil {
		return nil, fmt.Errorf("failed to load: %w", err)
	}

	client, soapClient, err := newClient(ctx, *u)
	if err != nil {
		return nil, fmt.Errorf("failed to create clients: %w", err)
	}

	return &VCenter{
		Client:     *client,
		SOAPClient: *soapClient,
		Conf:       *c,
	}, nil
}

// IsAgent check instance is agent
func IsAgent(instanceName string) bool {
	if _, err := runner.ToUUID(instanceName); err != nil {
		return false
	}
	return true
}

func (c *VCenter) isRunningAgent(ctx context.Context, vm string) (string, pb.Status, error) {
	cctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	vmInfo, err := c.Client.GetVMWithoutDevice(cctx, vm)
	if err != nil {
		return "", pb.Status_Unknown, fmt.Errorf("failed to get virtual machine: %w", err)
	}

	if vmInfo.PowerState != openapi.VCENTERVMPOWERSTATE_POWERED_ON {
		return "", pb.Status_Unknown, fmt.Errorf("%s is not powered on", vm)
	}

	identity, err := c.Client.GetVMGuestIdentity(cctx, vm)
	if err != nil {
		return "", pb.Status_Unknown, fmt.Errorf("failed to get virtual machine identity: %w", err)
	}

	ipAddress := identity.IpAddress
	if ipAddress == nil {
		return "", pb.Status_Unknown, fmt.Errorf("failed to get ip address: %w", err)
	}

	grpcHost := fmt.Sprintf("%s:%d", *ipAddress, server.AgentListenPort)

	grpcConn, err := grpc.DialContext(cctx, grpcHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", pb.Status_Unknown, fmt.Errorf("failed to dial agent server: %w", err)
	}
	client := pb.NewShoesAgentClient(grpcConn)
	resp, err := client.GetAgentStatus(cctx, &pb.GetAgentStatusRequest{})
	if err != nil {
		return "", pb.Status_Unknown, fmt.Errorf("failed to get agent status: %w", err)
	}

	return grpcHost, resp.GetStatus(), nil
}

func (c *VCenter) GetAgent(ctx context.Context, cloudID string) (*backend.Agent, error) {
	grpcHost, status, err := c.isRunningAgent(ctx, cloudID)
	if err != nil {
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}
	return &backend.Agent{
		CloudID:  cloudID,
		GRPCHost: grpcHost,
		Status:   status,
	}, nil
}

func (c *VCenter) ListAgent(ctx context.Context) ([]backend.Agent, error) {
	vms, err := c.Client.ListVM(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get list vm: %w", err)
	}

	var agents []backend.Agent
	for _, vm := range vms {
		if IsAgent(vm.Name) {
			grpcHost, status, err := c.isRunningAgent(ctx, vm.Vm)
			if err != nil {
				log.Printf("failed to check running agent: %+v\n", err)
				continue
			}

			a := backend.Agent{
				CloudID:  vm.Vm,
				GRPCHost: grpcHost,
				Status:   status,
			}

			agents = append(agents, a)
		}
	}

	return agents, nil
}

// CreateInstance create a VirtualMachine that running agent from snapshot
func (c *VCenter) CreateInstance(ctx context.Context, runnerName string) error {
	host, err := c.schedule(ctx)
	if err != nil {
		return fmt.Errorf("failed to schedule host: %w", err)
	}

	u := uuid.New()
	newVMName := runner.ToName(u.String())

	hostName := host.Name
	vmID, err := soap.LinkedCloneVirtualMachine(
		ctx,
		c.SOAPClient,
		c.Conf.datacenterName,
		newVMName,
		c.Conf.sourceVMName,
		c.Conf.sourceSnapshotName,
		c.Conf.datastoreName,
		&hostName)
	if err != nil {
		return fmt.Errorf("failed to linked clone virtual machine: %w", err)
	}

	log.Printf("creating virtual machine is completed! (VM ID: %s)\n", vmID)

	return nil
}

func (c *VCenter) schedule(ctx context.Context) (*openapi.VcenterHostSummary, error) {
	targetHost, err := c.getTargetHost(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get target hosts: %w", err)
	}

	return targetHost, nil
}

func (c *VCenter) getTargetHost(ctx context.Context) (*openapi.VcenterHostSummary, error) {
	dc, err := soap.GetDatacenter(ctx, c.SOAPClient, c.Conf.datacenterName)
	if err != nil {
		return nil, fmt.Errorf("failed to get datacenter: %w", err)
	}

	hosts, err := c.Client.GetHostsInDatacenter(ctx, dc.Reference().Value)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of host: %w", err)
	}

	rand.Shuffle(len(hosts), func(i, j int) { hosts[i], hosts[j] = hosts[j], hosts[i] })

	for _, host := range hosts {
		if err := c.canCreateHost(ctx, host.Host); err != nil {
			// can't create host, ignore
			continue
		}
		return &host, nil
	}
	return nil, fmt.Errorf("no host available: %w", shoesprovider.ErrAlreadyCreated)
}

func (c *VCenter) canCreateHost(ctx context.Context, hostID string) error {
	if c.Conf.numberInstancePerHost == NoLimitNumberVirtualMachine {
		return nil
	}

	vms, err := c.Client.GetVirtualMachinesInHost(ctx, hostID)
	if err != nil {
		return fmt.Errorf("failed to get list of virtual machine in host: %w", err)
	}
	ignoredVMs := ignoreSystemVirtualMachine(vms, c.Conf.ignoreVMNames)

	if len(ignoredVMs) >= c.Conf.numberInstancePerHost {
		return fmt.Errorf("already created virtual machine more than configured number (now: %d, c.numberInstancePerHost: %d)", len(ignoredVMs), c.Conf.numberInstancePerHost)
	}

	return nil
}

func ignoreSystemVirtualMachine(vms []openapi.VcenterVMSummary, ignoreVMNames []string) []openapi.VcenterVMSummary {
	var result []openapi.VcenterVMSummary

	for _, vm := range vms {
		if strings.EqualFold(vm.Name, "VMWare vCenter Server") {
			continue
		}

		if isSlice(vm.Name, ignoreVMNames) {
			continue
		}

		result = append(result, vm)
	}
	return result
}

func isSlice(value string, slice []string) bool {
	for _, s := range slice {
		if s == value {
			return true
		}
	}

	return false
}

// DeleteInstance delete a VirtualMachine
func (c *VCenter) DeleteInstance(ctx context.Context, cloudID string) error {
	// TODO: shutdown instance

	if err := c.Client.DeleteVM(ctx, cloudID); err != nil {
		return fmt.Errorf("failed to delete a virtual machine (VM ID: %s): %w", cloudID, err)
	}
	return nil
}
