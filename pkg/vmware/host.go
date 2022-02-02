package vmware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
)

func (c *Client) ListHost(ctx context.Context, query *SearchQuery) ([]openapi.VcenterHostSummary, error) {
	spath := "/api/vcenter/host"

	req, err := c.newRequest(ctx, http.MethodGet, spath, nil)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}
	req = AddSearchQuery(req, query)

	var summaries []openapi.VcenterHostSummary
	if err := c.request(req, &summaries); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}

	if len(summaries) == 0 {
		return nil, ErrNotFound
	}
	return summaries, nil
}

// GetHostsInDatacenter get hosts in datacenterID
func (c *Client) GetHostsInDatacenter(ctx context.Context, datacenterID string) ([]openapi.VcenterHostSummary, error) {
	query := NewSearchQueryDatacenters([]string{datacenterID})

	summaries, err := c.ListHost(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute list of host: %w", err)
	}
	return summaries, nil
}
