// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package watsonx holds Plugin Framework resources for the watsonx / Data Platform
// APIs that terraform-ibm-watsonx-ai used to reach through Mastercard/restapi.
package watsonx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
)

var (
	apiHosts = map[string]string{
		"us-south": "api.dataplatform.cloud.ibm.com",
		"eu-gb":    "api.eu-gb.dataplatform.cloud.ibm.com",
		"eu-de":    "api.eu-de.dataplatform.cloud.ibm.com",
		"jp-tok":   "api.jp-tok.dataplatform.cloud.ibm.com",
		"au-syd":   "api.au-syd.dai.cloud.ibm.com",
		"ca-tor":   "api.ca-tor.dai.cloud.ibm.com",
	}
	uiHosts = map[string]string{
		"us-south": "dataplatform.cloud.ibm.com",
		"eu-gb":    "eu-gb.dataplatform.cloud.ibm.com",
		"eu-de":    "eu-de.dataplatform.cloud.ibm.com",
		"jp-tok":   "jp-tok.dataplatform.cloud.ibm.com",
		"au-syd":   "au-syd.dai.cloud.ibm.com",
		"ca-tor":   "ca-tor.dai.cloud.ibm.com",
	}
)

func supportedRegions() []string {
	return []string{"us-south", "eu-gb", "eu-de", "jp-tok", "au-syd", "ca-tor"}
}

// apiError carries the HTTP status so callers can treat 404 as "gone".
type apiError struct {
	StatusCode int
	Method     string
	URL        string
	Body       string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%s %s returned HTTP %d: %s", e.Method, e.URL, e.StatusCode, e.Body)
}

func isNotFound(err error) bool {
	if ae, ok := err.(*apiError); ok {
		return ae.StatusCode == http.StatusNotFound
	}
	return false
}

// client is a thin JSON client that reuses the provider's IAM session, so no token
// ever lands in Terraform state.
type client struct {
	token func() (string, error)
	http  *http.Client
}

func newClient(providerData any) (*client, error) {
	session, ok := providerData.(conns.ClientSession)
	if !ok {
		return nil, fmt.Errorf("expected conns.ClientSession, got: %T. Please report this issue to the provider developers", providerData)
	}
	token := func() (string, error) {
		bmx, err := session.BluemixSession()
		if err != nil {
			return "", err
		}
		return bmx.Config.IAMAccessToken, nil
	}
	return &client{token: token, http: &http.Client{Timeout: 2 * time.Minute}}, nil
}

func apiBase(region, endpoint string) (string, error) {
	if endpoint != "" {
		return strings.TrimRight(ensureHTTPS(endpoint), "/"), nil
	}
	host, ok := apiHosts[region]
	if !ok {
		return "", fmt.Errorf("region %q is not supported, use one of %v or set endpoint", region, supportedRegions())
	}
	return "https://" + host, nil
}

func uiBase(region string) string {
	if host, ok := uiHosts[region]; ok {
		return "https://" + host
	}
	return ""
}

func ensureHTTPS(u string) string {
	if strings.HasPrefix(u, "https://") || strings.HasPrefix(u, "http://") {
		return u
	}
	return "https://" + strings.TrimPrefix(u, "//")
}

// do sends one request. out may be nil. It returns the response headers so callers
// can read Location after a create.
func (c *client) do(ctx context.Context, method, url string, in, out any) (http.Header, error) {
	token, err := c.token()
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(token, "Bearer ") {
		token = "Bearer " + token
	}

	var body io.Reader
	if in != nil {
		b, err := json.Marshal(in)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "terraform-provider-ibm/"+version.Version)
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return resp.Header, &apiError{StatusCode: resp.StatusCode, Method: method, URL: url, Body: string(raw)}
	}
	if out != nil && len(raw) > 0 {
		if err := json.Unmarshal(raw, out); err != nil {
			return resp.Header, fmt.Errorf("decoding response of %s %s: %w", method, url, err)
		}
	}
	return resp.Header, nil
}
