package vmware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
)

func (c *Client) CloneVM(ctx context.Context, name, sourceVMID string, host *string, powerOn bool) (*VcenterVMCreateResult, error) {
	var p *openapi.VcenterVMClonePlacementSpec

	if host != nil {
		p = &openapi.VcenterVMClonePlacementSpec{
			Host: host,
		}
	}

	spec := openapi.VcenterVMCloneSpec{
		Name:      name,
		Source:    sourceVMID,
		Placement: p,
		PowerOn:   &powerOn,
	}
	jb, err := json.Marshal(spec)
	if err != nil {
		return nil, fmt.Errorf(errCreatePostValue+": %w", err)
	}

	return c.createVM(ctx, bytes.NewBuffer(jb), "clone")
}
