package soap

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
)

// GetFinder get a finder
func GetFinder(ctx context.Context, client govmomi.Client, dcName string) (*find.Finder, *object.DatacenterFolders, error) {
	f := find.NewFinder(client.Client, true)

	dc, err := f.Datacenter(ctx, fmt.Sprintf("/%s", dcName))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get datacenter: %w", err)
	}
	f.SetDatacenter(dc)

	folders, err := dc.Folders(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get folders: %w", err)
	}

	return f, folders, nil
}

// GetDatacenter get a datacenter
func GetDatacenter(ctx context.Context, client govmomi.Client, dcName string) (*object.Datacenter, error) {
	f := find.NewFinder(client.Client, true)

	dc, err := f.Datacenter(ctx, fmt.Sprintf("/%s", dcName))
	if err != nil {
		return nil, fmt.Errorf("failed to get datacenter: %w", err)
	}
	return dc, nil
}

// GetVirtualMachine get a virtual machine
func GetVirtualMachine(ctx context.Context, client govmomi.Client, dcName, vmName string) (*object.VirtualMachine, error) {
	finder, folders, err := GetFinder(ctx, client, dcName)
	if err != nil {
		return nil, fmt.Errorf("failed to get finder: %w", err)
	}

	vm, err := finder.VirtualMachine(ctx, filepath.Join(folders.VmFolder.InventoryPath, vmName))
	if err != nil {
		return nil, fmt.Errorf("failed to get vm: %w", err)
	}

	return vm, nil
}

// GetDatastore get a datastore
func GetDatastore(ctx context.Context, client govmomi.Client, dcName, datastoreName string) (*object.Datastore, error) {
	finder, folders, err := GetFinder(ctx, client, dcName)
	if err != nil {
		return nil, fmt.Errorf("failed to get finder: %w", err)
	}

	datastore, err := finder.Datastore(ctx, filepath.Join(folders.DatastoreFolder.InventoryPath, datastoreName))
	if err != nil {
		return nil, fmt.Errorf("failed to get datastore: %w", err)
	}
	return datastore, nil
}

// GetHost get a host
func GetHost(ctx context.Context, client govmomi.Client, dcName, hostName string) (*object.HostSystem, error) {
	finder, _, err := GetFinder(ctx, client, dcName)
	if err != nil {
		return nil, fmt.Errorf("failed to get finder: %w", err)
	}

	//host, err := finder.HostSystem(ctx, filepath.Join(folders.HostFolder.InventoryPath, hostName))
	host, err := finder.HostSystem(ctx, fmt.Sprintf("*%s*", hostName))
	if err != nil {
		return nil, fmt.Errorf("failed to get host: %w", err)
	}

	return host, nil
}

// GetVirtualMachineCreateDate get information of virtual machine
func GetVirtualMachineCreateDate(ctx context.Context, client govmomi.Client, dcName, vmName string) (*time.Time, error) {
	vm, err := GetVirtualMachine(ctx, client, dcName, vmName)
	if err != nil {
		return nil, fmt.Errorf("failed to get virtual machine: %w", err)
	}
	vmRef := vm.Reference()

	pc := property.DefaultCollector(client.Client)
	var vmt mo.VirtualMachine
	if err := pc.RetrieveOne(ctx, vmRef, nil, &vmt); err != nil {
		return nil, fmt.Errorf("failed to retrieve virtual machine info: %w", err)
	}

	return vmt.Config.CreateDate, nil
}
