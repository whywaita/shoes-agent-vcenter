package controller

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/whywaita/shoes-agent-vcenter/vcenter-runner-controller/pkg/logger"
)

func (c *Controller) createInstanceIfNeed(ctx context.Context, hostID string) error {
	isNeeded, reason, err := c.isNeededCreate(ctx, hostID)
	if err != nil {
		return fmt.Errorf("failed to check need: %w", err)
	}
	if !isNeeded {
		logger.Logf(ctx, false, "not need creating virtual machine (reason: %s)", reason)
		return nil
	}

	if err := c.createInstance(ctx, hostID); err != nil {
		return fmt.Errorf("failed to create virtual machine: %w", err)
	}
	return nil
}

// isNeededCreate check number of virtual machines in the host.
// It avoids creating many virtual machines when while booting.
func (c *Controller) isNeededCreate(ctx context.Context, hostID string) (bool, string, error) {
	vms, err := c.client.GetVirtualMachinesInHost(ctx, hostID)
	if err != nil {
		return false, "", fmt.Errorf("failed to get list of virtual machine in host: %w", err)
	}
	ignoredVMs := ignoreSystemVirtualMachine(vms, c.ignoreVMNames)

	if len(ignoredVMs) >= c.numberInstancePerHost {
		return false, fmt.Sprintf("already created virtual machine more than configured number, wait booting... (now: %d, c.numberInstancePerHost: %d)", len(ignoredVMs), c.numberInstancePerHost), nil
	}

	logger.Logf(ctx, false, "virtual machines: %+v", vms)

	return true, "", nil
}

func (c *Controller) createInstance(ctx context.Context, hostID string) error {
	u := uuid.New().String()

	vmID, err := c.cloneVirtualMachine(ctx, hostID, u)
	if err != nil {
		return fmt.Errorf("failed to clone virtual machine: %w", err)
	}

	logger.Logf(ctx, false, "[Virtual Machine ID: %s] creating virtual machine is completed!", vmID)
	return nil
}
