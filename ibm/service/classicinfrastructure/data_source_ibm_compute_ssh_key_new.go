// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
//Plugin Framework migration from data_source_ibm_compute_ssh_key.go

package classicinfrastructure

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &ComputeSshKeyDataSource{}
var _ datasource.DataSourceWithConfigure = &ComputeSshKeyDataSource{}

type ComputeSshKeyDataSource struct {
	client interface{}
}

func NewComputeSshKeyDataSource(client interface{}) datasource.DataSource {
	return &ComputeSshKeyDataSource{
		client: client,
	}
}

func (d *ComputeSshKeyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_computesshkey_new"
}

func (d *ComputeSshKeyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Plugin Framework implementation of ComputeSshKey data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Data source identifier",
			},
			// TODO: Add remaining schema attributes from SDKv2 implementation
		},
	}
}

func (d *ComputeSshKeyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData
}

func (d *ComputeSshKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ComputeSshKeyDataSourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Implement Read logic from SDKv2 data source
	log.Printf("[DEBUG] Reading ComputeSshKey data source")

	resp.Diagnostics.AddError(
		"Not Yet Implemented",
		"This data source has not been fully migrated to Plugin Framework yet.",
	)
}

type ComputeSshKeyDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	// TODO: Add remaining model fields from SDKv2 schema
}
