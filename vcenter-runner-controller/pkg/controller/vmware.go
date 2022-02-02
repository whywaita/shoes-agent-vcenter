package controller

import (
	"context"
	"fmt"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/soap"
	"github.com/whywaita/shoes-agent-vcenter/vcenter-runner-controller/pkg/logger"
)

func (c *Controller) cloneVirtualMachine(ctx context.Context, hostID string, name string) (string, error) {
	logger.Logf(ctx, false, "will create virtual machine")

	u, err := soap.LinkedCloneVirtualMachine(ctx,
		c.govmomiClient,
		c.datacenter.Name,
		name,
		c.baseVMName,
		c.baseSnapshotName,
		c.datastoreName,
		&hostID,
	)
	if err != nil {
		return "", fmt.Errorf("failed to execute linked clone virtual machine: %w", err)
	}

	return u, nil
}

func (c *Controller) deleteVirtualMachine(ctx context.Context, vmID string) error {
	logger.Logf(ctx, false, "will delete virtual machine")

	if err := c.client.DeleteVM(ctx, vmID); err != nil {
		return fmt.Errorf("failed to execute delete virtual machine: %w", err)
	}
	return nil
}
