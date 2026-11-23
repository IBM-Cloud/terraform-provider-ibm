// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &projectDataSource{}
	_ datasource.DataSourceWithConfigure = &projectDataSource{}
)

func NewProjectDataSource() datasource.DataSource {
	return &projectDataSource{}
}

type projectDataSource struct {
	client *client
}

type projectDataModel struct {
	ID            types.String `tfsdk:"id"`
	Region        types.String `tfsdk:"region"`
	Endpoint      types.String `tfsdk:"endpoint"`
	ProjectID     types.String `tfsdk:"project_id"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Tags          types.List   `tfsdk:"tags"`
	COSInstanceID types.String `tfsdk:"cos_instance_id"`
	COSCRN        types.String `tfsdk:"cos_crn"`
	ProjectURL    types.String `tfsdk:"project_url"`
	RawJSON       types.String `tfsdk:"project_json"`
}

func (d *projectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_watsonx_project"
}

func (d *projectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads a watsonx.ai project by GUID.",
		Attributes: map[string]schema.Attribute{
			"id":              schema.StringAttribute{Computed: true, Description: "Project GUID."},
			"region":          schema.StringAttribute{Required: true, Description: "watsonx region. One of us-south, eu-gb, eu-de, jp-tok, au-syd, ca-tor."},
			"endpoint":        schema.StringAttribute{Optional: true, Description: "Override for the Data Platform API host."},
			"project_id":      schema.StringAttribute{Required: true, Description: "Project GUID to look up."},
			"name":            schema.StringAttribute{Computed: true, Description: "Project name."},
			"description":     schema.StringAttribute{Computed: true, Description: "Project description."},
			"tags":            schema.ListAttribute{Computed: true, ElementType: types.StringType, Description: "Project tags."},
			"cos_instance_id": schema.StringAttribute{Computed: true, Description: "GUID of the backing Cloud Object Storage instance."},
			"cos_crn":         schema.StringAttribute{Computed: true, Description: "CRN of the backing Cloud Object Storage instance."},
			"project_url":     schema.StringAttribute{Computed: true, Description: "Link to the project in the watsonx UI."},
			"project_json":    schema.StringAttribute{Computed: true, Description: "Full project record as returned by the API, for fields not modelled above. Same shape as api_response on data.restapi_object."},
		},
	}
}

func (d *projectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, err := newClient(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Configure watsonx Client", err.Error())
		return
	}
	d.client = c
}

func (d *projectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data projectDataModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	base, err := apiBase(data.Region.ValueString(), data.Endpoint.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}

	var list struct {
		Resources []json.RawMessage `json:"resources"`
	}
	id := data.ProjectID.ValueString()
	if _, err := d.client.do(ctx, http.MethodGet, base+"/v2/projects?project_ids="+id, nil, &list); err != nil {
		resp.Diagnostics.AddError("Unable to Read watsonx Project", err.Error())
		return
	}

	for _, raw := range list.Resources {
		var p projectGetResponse
		if err := json.Unmarshal(raw, &p); err != nil || p.Metadata.GUID != id {
			continue
		}
		tags, diags := types.ListValueFrom(ctx, types.StringType, p.Entity.Tags)
		resp.Diagnostics.Append(diags...)
		data.ID = types.StringValue(id)
		data.Name = types.StringValue(p.Entity.Name)
		data.Description = types.StringValue(p.Entity.Description)
		data.Tags = tags
		data.COSInstanceID = types.StringValue(p.Entity.Storage.GUID)
		data.COSCRN = types.StringValue(p.Entity.Storage.ResourceCRN)
		data.ProjectURL = types.StringValue(uiBase(data.Region.ValueString()) + "/projects/" + id)
		data.RawJSON = types.StringValue(string(raw))
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}
	resp.Diagnostics.AddError("Project Not Found", fmt.Sprintf("No watsonx project with GUID %q in region %s.", id, data.Region.ValueString()))
}
