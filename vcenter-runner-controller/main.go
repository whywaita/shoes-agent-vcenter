package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/vmware/govmomi"

	"github.com/whywaita/shoes-agent-vcenter/vcenter-runner-controller/pkg/controller"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const (
	EnvURL                   = "SHOES_VCENTER_URL"
	EnvUser                  = "SHOES_VCENTER_USERNAME"
	EnvPassword              = "SHOES_VCENTER_PASSWORD"
	EnvDatacenterName        = "SHOES_VCENTER_DATACENTER_NAME"
	EnvNumberInstancePerHost = "SHOES_VCENTER_NUMBER_OF_INSTANCE_PER_HOST"
	EnvBaseVMName            = "SHOES_VCENTER_BASE_VM_NAME"
	EnvBaseSnapshotName      = "SHOES_VCENTER_BASE_SNAPSHOT_NAME"
	EnvDatastoreName         = "SHOES_VCENTER_DATASTORE_NAME"
	EnvIgnoreVMName          = "SHOES_VCENTER_IGNORE_VM_NAMES"
)

func load() (*url.URL, uint64, *controller.Config, []string, error) {
	u, err := url.Parse(os.Getenv(EnvURL))
	if err != nil {
		return nil, 0, nil, nil, fmt.Errorf("failed to parse URL: %w", err)
	}
	u.User = url.UserPassword(os.Getenv(EnvUser), os.Getenv(EnvPassword))

	envNumInstance := os.Getenv(EnvNumberInstancePerHost)
	numInstance, err := strconv.ParseUint(envNumInstance, 10, 10)
	if err != nil {
		return nil, 0, nil, nil, fmt.Errorf("failed to parse %s: %w", EnvNumberInstancePerHost, err)
	}

	dc := os.Getenv(EnvDatacenterName)
	baseVMName := os.Getenv(EnvBaseVMName)
	baseSnapshotName := os.Getenv(EnvBaseSnapshotName)
	datastoreName := os.Getenv(EnvDatastoreName)

	if dc == "" || baseVMName == "" || baseSnapshotName == "" || datastoreName == "" {
		return nil, 0, nil, nil, fmt.Errorf("must set %s and %s and %s and %s",
			EnvDatacenterName,
			EnvBaseVMName,
			EnvBaseSnapshotName,
			EnvDatastoreName,
		)
	}

	c := &controller.Config{
		DatacenterID:     dc,
		DatastoreName:    datastoreName,
		BaseVMName:       baseVMName,
		BaseSnapshotName: baseSnapshotName,
	}

	ignoreVMName := os.Getenv(EnvIgnoreVMName)
	ignoreVMNames := strings.Split(ignoreVMName, ",")

	return u, numInstance, c, ignoreVMNames, nil
}

func run() error {
	ctx := context.Background()
	u, numInstance, conf, ignoreVMNames, err := load()
	if err != nil {
		return fmt.Errorf("failed to load: %w", err)
	}

	soapURL := u
	soapURL.Path = "/sdk"
	c, err := govmomi.NewClient(ctx, soapURL, true)
	if err != nil {
		return fmt.Errorf("failed to create a soap client: %w", err)
	}

	con, err := controller.New(ctx, u, numInstance, *c, *conf, ignoreVMNames)
	if err != nil {
		return fmt.Errorf("failed to initalize controller: %w", err)
	}

	if err := con.Run(ctx); err != nil {
		return fmt.Errorf("failed to run controller: %w", err)
	}

	return nil
}
