package controller

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/soap"
	"github.com/whywaita/shoes-agent-vcenter/vcenter-runner-controller/pkg/logger"
	"golang.org/x/sync/errgroup"
)

type vCenterSummary struct {
	openapi.VcenterVMSummary

	createDate time.Time
}

// deleteOldVM delete a virtual machine and leave numberInstancePerHost only
func (c *Controller) deleteOldVM(ctx context.Context, openAPISummaries []openapi.VcenterVMSummary) error {
	summaries, err := c.getVcenterSummaries(ctx, openAPISummaries)
	if err != nil {
		return fmt.Errorf("failed to get list of vCenter summaries: %w", err)
	}
	_, old := c.divideOldVirtualMachine(summaries)

	eg, cctx := errgroup.WithContext(ctx)
	for _, oldVM := range old {
		oldVM = oldVM
		eg.Go(func() error {
			cctx = logger.SetVirtualMachineName(cctx, oldVM.Name)

			if oldVM.PowerState == openapi.VCENTERVMPOWERSTATE_POWERED_ON {
				logger.Logf(ctx, false, "stopping virtual machine...")
				if err := c.client.StopVM(ctx, oldVM.Vm); err != nil {
					return fmt.Errorf("failed to stop virtual machine: %w", err)
				}
			}

			logger.Logf(ctx, false, "deleting virtual machine...")
			if err := c.deleteVirtualMachine(cctx, oldVM.Vm); err != nil {
				return fmt.Errorf("failed to delete virtual machine: %w", err)
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to wait errgroup per virtual machine: %w", err)
	}

	return nil
}

// divideOldVirtualMachine divide summaries to leave (only numberInstancePerHost) and old (not numberInstancePerHost)
func (c *Controller) divideOldVirtualMachine(summaries []vCenterSummary) ([]vCenterSummary, []vCenterSummary) {
	leave := summaries[0:c.numberInstancePerHost]
	old := summaries[c.numberInstancePerHost:]
	return leave, old
}

func (c *Controller) getVcenterSummaries(ctx context.Context, summaries []openapi.VcenterVMSummary) ([]vCenterSummary, error) {
	var result []vCenterSummary

	for _, vm := range summaries {
		t, err := soap.GetVirtualMachineCreateDate(ctx, c.govmomiClient, c.datacenter.Name, vm.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to get virtual machine info: %w", err)
		}
		result = append(result, vCenterSummary{
			VcenterVMSummary: vm,
			createDate:       *t,
		})
	}
	sort.SliceStable(result, func(i, j int) bool {
		// result[0] is the newest object
		// result[len(result)] is the oldest object
		return result[i].createDate.After(result[j].createDate)
	})

	return result, nil
}
