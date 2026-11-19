// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package accountmanagement

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &accountTraitsDataSource{}
	_ datasource.DataSourceWithConfigure = &accountTraitsDataSource{}
)

// NewAccountTraitsDataSource reads account traits (GET /v1/accounts/{account_id}/traits).
func NewAccountTraitsDataSource() datasource.DataSource {
	return &accountTraitsDataSource{}
}

type accountTraitsDataSource struct {
	client *traitsClient
}

func (d *accountTraitsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_traits"
}

func (d *accountTraitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads account traits through the Account Management API: use of Financial Services Validated products, and user list visibility.",
		Attributes: map[string]schema.Attribute{
			"id":                     schema.StringAttribute{Computed: true, Description: "Same as account_id."},
			"account_id":             schema.StringAttribute{Optional: true, Computed: true, Description: "Account ID. Defaults to the account of the provider's credentials."},
			"endpoint":               schema.StringAttribute{Optional: true, Description: "Account Management API host. Defaults to the endpoint the provider is configured with."},
			"fs_ready":               schema.BoolAttribute{Computed: true, Description: "Whether Financial Services Validated products are enabled."},
			"team_directory_enabled": schema.BoolAttribute{Computed: true, Description: "Whether users can see the user list (team directory)."},
		},
	}
}

func (d *accountTraitsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, err := newTraitsClient(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Create Account Management Client", err.Error())
		return
	}
	d.client = c
}

func (d *accountTraitsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data accountTraitsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, err := d.client.forEndpoint(data.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Account Management Endpoint", err.Error())
		return
	}
	accountID, err := d.client.accountID(data.AccountID)
	if err != nil {
		resp.Diagnostics.AddError("Unknown Account", err.Error())
		return
	}
	traits, _, err := getTraits(ctx, client, accountID)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Read Account Traits", err.Error())
		return
	}
	data.AccountID = types.StringValue(accountID)
	data.ID = data.AccountID
	data.FsReady = boolValue(traits.FsReady)
	data.TeamDirectoryEnabled = boolValue(traits.TeamDirectoryEnabled)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
