package vmware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
)

func (c *Client) GetVMGuestIdentity(ctx context.Context, vm string) (*openapi.VcenterVmGuestIdentityInfo, error) {
	spath := fmt.Sprintf("/api/vcenter/vm/%s/guest/identity", vm)

	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}

	var info openapi.VcenterVmGuestIdentityInfo
	if err := c.request(req, &info); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}

	return &info, nil
}
