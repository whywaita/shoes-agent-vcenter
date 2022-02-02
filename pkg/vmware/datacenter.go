package vmware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
)

func (c *Client) GetDatacenter(ctx context.Context, dcID string) (*openapi.VcenterDatacenterInfo, error) {
	spath := fmt.Sprintf("/api/vcenter/datacenter/%s", dcID)

	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}

	var dc openapi.VcenterDatacenterInfo
	if err := c.request(req, &dc); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}

	return &dc, nil
}
