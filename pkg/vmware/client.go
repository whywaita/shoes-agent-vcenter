package vmware

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"
)

// Client is client for vCenter REST API
type Client struct {
	//API openapi.APIClient

	BaseURL  *url.URL
	Username string
	Password string

	SessionID string

	HTTPClient *http.Client
	Logger     *log.Logger
}

var (
	version   = "dirty"
	userAgent = fmt.Sprintf("shoes-agent-vcenter-per-host/%s", version)
)

// httpMu lock to create *http.Request while to call generateSession
var (
	httpMu sync.RWMutex
)

// NewClient create a client
func NewClient(vCenterIP, username, password string, logger *log.Logger) (*Client, error) {
	// validate
	if vCenterIP == "" || username == "" || password == "" {
		return nil, fmt.Errorf("must be set vCenterIP and username and password")
	}

	u, err := url.Parse(vCenterIP)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vCenter IP address: %w", err)
	}
	u.Scheme = "https"

	if logger == nil {
		l := log.New(io.Discard, "", log.LstdFlags)
		logger = l
	}

	httpClient := http.DefaultClient
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	sessionID, err := generateSessionID(vCenterIP, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to generate session ID: %w", err)
	}

	return &Client{
		BaseURL:    u,
		Username:   username,
		Password:   password,
		SessionID:  sessionID,
		HTTPClient: httpClient,
		Logger:     logger,
	}, nil
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	httpMu.RLock()
	u := *c.BaseURL
	u.Path = path.Join(c.BaseURL.Path, spath)
	httpMu.RUnlock()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create new HTTP request: %w", err)
	}

	req.Header.Set("vmware-api-session-id", c.SessionID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}
