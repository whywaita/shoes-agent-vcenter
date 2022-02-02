package controller

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/vmware/govmomi"
	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware"
	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
	"github.com/whywaita/shoes-agent-vcenter/vcenter-runner-controller/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type Controller struct {
	client        vmware.Client
	govmomiClient govmomi.Client

	numberInstancePerHost int
	baseVMName            string
	baseSnapshotName      string
	datastoreName         string
	ignoreVMNames         []string

	datacenterID string
	datacenter   openapi.VcenterDatacenterInfo
}

type Config struct {
	DatacenterID     string
	DatastoreName    string
	BaseVMName       string
	BaseSnapshotName string
}

func New(ctx context.Context, u *url.URL, numInstance uint64, govmomiClient govmomi.Client, c Config, ignoreVMNames []string) (*Controller, error) {
	password, ok := u.User.Password()
	if !ok {
		return nil, fmt.Errorf("not set password")
	}
	apiClient, err := vmware.NewClient(u.Host, u.User.Username(), password, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a client: %w", err)
	}

	dc, err := apiClient.GetDatacenter(ctx, c.DatacenterID)
	if err != nil {
		return nil, fmt.Errorf("failed to get datacenter: %w", err)
	}

	return &Controller{
		client:                *apiClient,
		numberInstancePerHost: int(numInstance),
		baseVMName:            c.BaseVMName,
		baseSnapshotName:      c.BaseSnapshotName,
		datastoreName:         c.DatastoreName,
		ignoreVMNames:         ignoreVMNames,
		datacenterID:          c.DatacenterID,
		datacenter:            *dc,
		govmomiClient:         govmomiClient,
	}, nil
}

// Run is main loop for vcenter-runner-controller
func (c *Controller) Run(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.reconcile(ctx); err != nil {
				log.Printf("failed to reconcile: %+v", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// reconcile check and create and delete virtual machines
func (c *Controller) reconcile(ctx context.Context) error {
	logger.Logf(ctx, true, "start reconcile: %s", c.client.BaseURL)
	hosts, err := c.client.GetHostsInDatacenter(ctx, c.datacenterID)
	if err != nil {
		return fmt.Errorf("failed to get list of host: %w", err)
	}

	eg, cctx := errgroup.WithContext(ctx)
	for _, host := range hosts {
		eg.Go(func() error {
			cctx = logger.SetHostName(cctx, host.Name)
			if err := c.reconcileHost(cctx, host.Host); err != nil {
				return fmt.Errorf("failed to reconcile host: %w", err)
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to wait errgroup per host: %w", err)
	}

	return nil
}

func (c *Controller) reconcileHost(ctx context.Context, hostID string) error {
	logger.Logf(ctx, true, "start reconcile in host")
	vms, err := c.client.GetVirtualMachinesInHost(ctx, hostID)
	if err != nil {
		return fmt.Errorf("failed to get list of vm: %w", err)
	}
	ignoredVMs := ignoreSystemVirtualMachine(vms, c.ignoreVMNames)

	switch {
	case len(ignoredVMs) == c.numberInstancePerHost:
		// It's correct
		logger.Logf(ctx, true, "number of instance is correct, not doing")
		return nil
	case len(ignoredVMs) < c.numberInstancePerHost:
		// It's a less than numberInstancePerHost
		// Need to create a virtual machine
		logger.Logf(ctx, false, "need to create a virtual machine, will be creating (now: %d, c.numberInstancePerHost: %d)", len(ignoredVMs), c.numberInstancePerHost)
		if err := c.createInstanceIfNeed(ctx, hostID); err != nil {
			return fmt.Errorf("failed to create virtual machine: %w", err)
		}
	case len(ignoredVMs) > c.numberInstancePerHost:
		// It's a more than numberInstancePerHost
		// Need to delete a virtual machine
		logger.Logf(ctx, false, "need to delete a virtual machine, will be deleting (now: %d, c.numberInstancePerHost: %d)", len(ignoredVMs), c.numberInstancePerHost)
		if err := c.deleteOldVM(ctx, ignoredVMs); err != nil {
			return fmt.Errorf("failed to delete old virtual machine: %w", err)
		}
		return nil
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
