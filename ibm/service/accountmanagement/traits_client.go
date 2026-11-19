// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package accountmanagement

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// accountTraits is the body of GET and PATCH /v1/accounts/{account_id}/traits.
type accountTraits struct {
	FsReady              *bool `json:"fs_ready,omitempty"`
	TeamDirectoryEnabled *bool `json:"team_directory_enabled,omitempty"`
}

// traitsClient sends the traits calls through the provider's Account
// Management client, so it reuses its authenticator, retries and endpoint.
type traitsClient struct {
	client  *accountmanagementv4.AccountManagementV4
	account string
}

func newTraitsClient(providerData any) (*traitsClient, error) {
	session, ok := providerData.(conns.ClientSession)
	if !ok {
		return nil, fmt.Errorf("expected conns.ClientSession, got: %T. Please report this issue to the provider developers", providerData)
	}
	client, err := session.AccountManagementV4()
	if err != nil {
		return nil, err
	}
	tc := &traitsClient{client: client}
	if user, err := session.BluemixUserDetails(); err == nil && user != nil {
		tc.account = user.UserAccount
	}
	return tc, nil
}

// forEndpoint returns a copy of the client pointed at endpoint
func (c *traitsClient) forEndpoint(endpoint types.String) (*accountmanagementv4.AccountManagementV4, error) {
	if c == nil || c.client == nil {
		return nil, fmt.Errorf("the Account Management client is not configured")
	}
	if endpoint.IsNull() || endpoint.IsUnknown() || endpoint.ValueString() == "" {
		return c.client, nil
	}
	url := strings.TrimRight(endpoint.ValueString(), "/")
	switch {
	case strings.HasPrefix(url, "//"):
		url = "https:" + url
	case !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://"):
		url = "https://" + url
	}
	clone := c.client.Clone()
	if err := clone.SetServiceURL(url); err != nil {
		return nil, err
	}
	return clone, nil
}

// accountID returns the configured account ID, or the provider's account.
func (c *traitsClient) accountID(configured types.String) (string, error) {
	if !configured.IsNull() && !configured.IsUnknown() && configured.ValueString() != "" {
		return configured.ValueString(), nil
	}
	if c.account == "" {
		return "", fmt.Errorf("could not determine the account ID from the provider; set account_id")
	}
	return c.account, nil
}

func getTraits(ctx context.Context, client *accountmanagementv4.AccountManagementV4, accountID string) (*accountTraits, *core.DetailedResponse, error) {
	var out accountTraits
	response, err := traitsRequest(ctx, client, core.GET, accountID, nil, &out)
	return &out, response, err
}

func patchTraits(ctx context.Context, client *accountmanagementv4.AccountManagementV4, accountID string, body *accountTraits) error {
	_, err := traitsRequest(ctx, client, core.PATCH, accountID, body, nil)
	return err
}

func traitsRequest(ctx context.Context, client *accountmanagementv4.AccountManagementV4, method, accountID string, body, out any) (*core.DetailedResponse, error) {
	builder := core.NewRequestBuilder(method).WithContext(ctx)
	builder.EnableGzipCompression = client.GetEnableGzipCompression()
	if _, err := builder.ResolveRequestURL(client.Service.Options.URL, "/v1/accounts/{account_id}/traits", map[string]string{"account_id": accountID}); err != nil {
		return nil, err
	}
	builder.AddHeader("Accept", "application/json")
	if body != nil {
		builder.AddHeader("Content-Type", "application/json")
		if _, err := builder.SetBodyContentJSON(body); err != nil {
			return nil, err
		}
	}
	request, err := builder.Build()
	if err != nil {
		return nil, err
	}
	if out == nil {
		return client.Service.Request(request, nil)
	}
	return client.Service.Request(request, out)
}

func isNotFound(r *core.DetailedResponse) bool {
	return r != nil && r.StatusCode == http.StatusNotFound
}

func boolValue(p *bool) types.Bool {
	if p == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*p)
}
