package vmware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// VcenterVMCreateResult struct for VcenterVMCreateResult
type VcenterVMCreateResult struct {
	VirtualMachineID string `json:"value"`
}

func (c *Client) createVM(ctx context.Context, body io.Reader, action string) (*VcenterVMCreateResult, error) {
	spath := "/api/vcenter/vm"

	req, err := c.newRequest(ctx, http.MethodPost, spath, body)
	if err != nil {
		return nil, fmt.Errorf(errHTTPCreateRequest+": %w", err)
	}
	if !strings.EqualFold(action, "") {
		q := req.URL.Query()
		q.Add("action", action)
		req.URL.RawQuery = q.Encode()
	}

	var result VcenterVMCreateResult
	if err := c.request(req, &result); err != nil {
		return nil, fmt.Errorf(errRequest+": %w", err)
	}
	return &result, nil
}
