package vmware

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

// generateSessionID generate a ID of session
func generateSessionID(ipAddress, username, password string) (string, error) {
	spath := "/rest/com/vmware/cis/session"

	endpoint := fmt.Sprintf("https://%s%s", ipAddress, spath)
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Accept", "application/json")

	client := http.DefaultClient
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to do http request: %w", err)
	}

	type resultSession struct {
		Value string `json:"value"`
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var r resultSession
	if err := decoder.Decode(&r); err != nil {
		return "", fmt.Errorf("failed to decode body: %w", err)
	}

	return r.Value, nil
}
