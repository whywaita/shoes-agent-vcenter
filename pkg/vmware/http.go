package vmware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/whywaita/shoes-agent-vcenter/pkg/vmware/openapi"
)

func decodeBody(resp *http.Response, out interface{}, logger *log.Logger) ([]byte, error) {
	defer resp.Body.Close()

	jb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(jb, &out); err != nil {
		return jb, fmt.Errorf("failed to unmarshal response json: %w", err)
	}
	return jb, nil
}

func (c *Client) request(req *http.Request, out interface{}) error {
	resp, err := c.do(req)
	if err != nil {
		return fmt.Errorf("failed to request: %w", err)
	}

	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		var errResp *openapi.VapiStdErrorsError
		var jb []byte
		var getErr error
		if errResp, jb, getErr = c.getErrorFromResponse(resp); getErr != nil {
			return fmt.Errorf("failed to get formatted error: %s. Response code: %d, Raw resnpose: %+v", getErr, resp.StatusCode, resp)
		}
		return fmt.Errorf("invalid status code (%d): %+v, Raw response: %+v, body: %s", resp.StatusCode, errResp, resp, jb)
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if jb, err := decodeBody(resp, out, c.Logger); err != nil {
		return fmt.Errorf("failed to decode response body (resp: %s): %w", jb, err)
	}

	return nil
}

var (
	errHTTPRequestDo     = "failed to http request"
	errHTTPCreateRequest = "failed to create request"
	errRequest           = "failed to do request"
	errCreatePostValue   = "failed to create post value"
	ErrNotFound          = fmt.Errorf("not found")
)

func (c *Client) do(req *http.Request) (*http.Response, error) {
	c.Logger.Printf("Do Request %+v\n", req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(errHTTPRequestDo+": %w", err)
	}

	return resp, nil
}

func (c *Client) getErrorFromResponse(resp *http.Response) (*openapi.VapiStdErrorsError, []byte, error) {
	var errResp openapi.VapiStdErrorsError
	var jb []byte
	var err error

	if jb, err = decodeBody(resp, &errResp, c.Logger); err != nil {
		return nil, jb, fmt.Errorf("failed to decode body %s: %w", jb, err)
	}
	return &errResp, jb, nil
}
