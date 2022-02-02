package vmware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
)

// VcenterVMInfo struct for VcenterVMInfo
type VcenterVMInfo struct {
	GuestOS openapi.VcenterVmGuestOS `json:"guest_OS"`
	// Virtual machine name.
	Name string `json:"name"`
	Identity *openapi.VcenterVmIdentityInfo `json:"identity,omitempty"`
	PowerState openapi.VcenterVmPowerState `json:"power_state"`
	// Indicates whether the virtual machine is frozen for instant clone, or not. This field is optional because it was added in a newer version than its parent node.
	InstantCloneFrozen *bool `json:"instant_clone_frozen,omitempty"`
	Hardware openapi.VcenterVmHardwareInfo `json:"hardware"`
	Boot openapi.VcenterVmHardwareBootInfo `json:"boot"`
	// Boot device configuration. If the list has no entries, a server-specific default boot sequence is used.
	BootDevices []openapi.VcenterVmHardwareBootDeviceEntry `json:"boot_devices"`
	Cpu openapi.VcenterVmHardwareCpuInfo `json:"cpu"`
	Memory openapi.VcenterVmHardwareMemoryInfo `json:"memory"`
	// List of disks. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk.
	Disks map[string]VcenterVMInfoDisks `json:"disks"`
	// List of Ethernet adapters. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Ethernet.
	Nics map[string]VcenterVMInfoNics `json:"nics"`
}

type VcenterVMInfoDisks struct {
		Backing struct {
			VmdkFile string `json:"vmdk_file"`
			Type     string `json:"type"`
		} `json:"backing"`
		Label string `json:"label"`
		Type  string `json:"type"`
		Sata  struct {
			Bus  int `json:"bus"`
			Unit int `json:"unit"`
		} `json:"sata"`
		Capacity int64 `json:"capacity"`
}

type VcenterVMInfoNics struct {
		StartConnected bool `json:"start_connected"`
		PciSlotNumber  int  `json:"pci_slot_number"`
		Backing        struct {
			NetworkName string `json:"network_name"`
			Type        string `json:"type"`
			Network     string `json:"network"`
		} `json:"backing"`
		MacAddress        string `json:"mac_address"`
		MacType           string `json:"mac_type"`
		AllowGuestControl bool   `json:"allow_guest_control"`
		WakeOnLanEnabled  bool   `json:"wake_on_lan_enabled"`
		Label             string `json:"label"`
		State             string `json:"state"`
		Type              string `json:"type"`
}

func (c *Client) GetVMWithoutDevice(ctx context.Context, vm string) (*VcenterVMInfo, error) {
	spath := fmt.Sprintf("/api/vcenter/vm/%s", vm)

	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}

	var info VcenterVMInfo
	if err := c.request(req, &info); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}

	return &info, nil
}


func (c *Client) GetVM(ctx context.Context, vm string) (*openapi.VcenterVMInfo, error) {
	spath := fmt.Sprintf("/api/vcenter/vm/%s", vm)

	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}

	var info openapi.VcenterVMInfo
	if err := c.request(req, &info); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}

	return &info, nil
}

func (c *Client) ListVM(ctx context.Context, query *SearchQuery) ([]openapi.VcenterVMSummary, error) {
	spath := "/api/vcenter/vm"

	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}
	req = AddSearchQuery(req, query)

	var summaries []openapi.VcenterVMSummary
	if err := c.request(req, &summaries); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}

	if len(summaries) == 0 {
		return nil, ErrNotFound
	}
	return summaries, nil
}

func (c *Client) DeleteVM(ctx context.Context, vm string) error {
	spath := fmt.Sprintf("api/vcenter/vm/%s", vm)

	req, err := c.newRequest(ctx, http.MethodDelete, spath, nil)
	if err != nil {
		return fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}

	var i interface{} // this endpoint return N/A
	if err = c.request(req, i); err != nil {
		return fmt.Errorf(errRequest+": %w", err)
	}
	return nil
}

func (c *Client) StopVM(ctx context.Context, vm string) error {
	spath := fmt.Sprintf("api/vcenter/vm/%s/power", vm)

	req, err := c.newRequest(ctx, http.MethodPost, spath, nil)
	if err != nil {
		return fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}
	q := req.URL.Query()
	q.Add("action", "stop")
	req.URL.RawQuery = q.Encode()

	var i interface{} // this endpoint return N/A
	if err = c.request(req, i); err != nil {
		return fmt.Errorf(errRequest+": %w", err)
	}
	return nil
}

// GetVirtualMachinesInHost get virtual machines in hostID
func (c *Client) GetVirtualMachinesInHost(ctx context.Context, hostID string) ([]openapi.VcenterVMSummary, error) {
	query := NewSearchQueryHosts([]string{hostID})

	summaries, err := c.ListVM(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute list of vm: %w", err)
	}
	return summaries, nil
}

// GetRunningVirtualMachinesInHost get virtual machines in hostID and running status
func (c *Client) GetRunningVirtualMachinesInHost(ctx context.Context, hostID string) ([]openapi.VcenterVMSummary, error) {
	r, err := c.GetVirtualMachinesInHost(ctx, hostID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute get list of virtual machine: %w", err)
	}

	var result []openapi.VcenterVMSummary
	for _, vm := range r {
		if vm.PowerState == openapi.VCENTERVMPOWERSTATE_POWERED_ON {
			result = append(result, vm)
		}
	}
	return result, nil
}
