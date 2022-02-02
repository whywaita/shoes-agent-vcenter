package soap

import (
	"context"
	"fmt"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

// LinkedCloneVirtualMachine create a virtual machine using linked clone (like a Copy on Write)
func LinkedCloneVirtualMachine(ctx context.Context, client govmomi.Client, dcName, newVMName, sourceVMName, sourceSnapshotName, datastoreName string, hostName *string) (string, error) {
	_, folders, err := GetFinder(ctx, client, dcName)
	if err != nil {
		return "", fmt.Errorf("failed to get finder: %w", err)
	}
	datastoreRef, sourceVM, err := getObjectsLinkedClone(ctx, client, dcName, datastoreName, sourceVMName)
	if err != nil {
		return "", fmt.Errorf("failed to get refercenses: %w", err)
	}

	snapshotRef, err := sourceVM.FindSnapshot(ctx, sourceSnapshotName)
	if err != nil {
		return "", fmt.Errorf("failed to find snapshot: %w", err)
	}

	hostRef, err := getHostReferences(ctx, client, dcName, hostName)
	if err != nil {
		return "", fmt.Errorf("failed to get host reference: %w", err)
	}

	vmFolderRef := folders.VmFolder.Reference()
	relocateSpec := types.VirtualMachineRelocateSpec{
		// DeviceChange: // Configure Ethernet card
		Host:         hostRef,
		Folder:       &vmFolderRef,
		Datastore:    datastoreRef,
		DiskMoveType: string(types.VirtualMachineRelocateDiskMoveOptionsCreateNewChildDiskBacking),
	}
	cloneSpec := types.VirtualMachineCloneSpec{
		PowerOn:  true,
		Template: false,
		Location: relocateSpec,
		Snapshot: snapshotRef,
	}

	task, err := sourceVM.Clone(ctx, folders.VmFolder, newVMName, cloneSpec)
	if err != nil {
		return "", fmt.Errorf("failed to create clone task: %w", err)
	}

	info, err := task.WaitForResult(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to wait task: %w", err)
	}
	newVM := object.NewVirtualMachine(client.Client, info.Result.(types.ManagedObjectReference))

	return newVM.Reference().Value, nil
}

func getObjectsLinkedClone(ctx context.Context, client govmomi.Client, dcName, datastoreName, sourceVMName string) (*types.ManagedObjectReference, *object.VirtualMachine, error) {
	datastore, err := GetDatastore(ctx, client, dcName, datastoreName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get datastore: %w", err)
	}
	dsRef := datastore.Reference()

	sourceVM, err := GetVirtualMachine(ctx, client, dcName, sourceVMName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get source virtual machine: %w", err)
	}

	return &dsRef, sourceVM, nil
}

func getHostReferences(ctx context.Context, client govmomi.Client, dcName string, hostName *string) (*types.ManagedObjectReference, error) {
	if hostName == nil {
		return nil, nil
	}

	host, err := GetHost(ctx, client, dcName, *hostName)
	if err != nil {
		return nil, fmt.Errorf("failed to get host: %w", err)
	}
	hostRef := host.Reference()
	return &hostRef, nil
}
