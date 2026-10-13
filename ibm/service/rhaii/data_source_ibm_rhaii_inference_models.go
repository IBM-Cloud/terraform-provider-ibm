// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &rhaiiInferenceModelsDataSource{}
	_ datasource.DataSourceWithConfigure = &rhaiiInferenceModelsDataSource{}
)

// NewRhaiiInferenceModelsDataSource returns the ibm_rhaii_inference_models data
// source, which lists the inference models of a Red Hat AI Inference project.
func NewRhaiiInferenceModelsDataSource() datasource.DataSource {
	return &rhaiiInferenceModelsDataSource{}
}

type rhaiiInferenceModelsDataSource struct {
	meta conns.ClientSession
}

type rhaiiInferenceModelsDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	ProjectID types.String `tfsdk:"project_id"`
	Location  types.String `tfsdk:"location"`
	Models    types.List   `tfsdk:"models"`
}

func rhaiiInferenceModelSummaryAttrTypes() map[string]attr.Type {
	t := rhaiiInferenceModelMetadataAttrTypes()
	t["id"] = types.StringType
	t["object"] = types.StringType
	t["created"] = types.Int64Type
	t["owned_by"] = types.StringType
	return t
}

func (d *rhaiiInferenceModelsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rhaii_inference_models"
}

func (d *rhaiiInferenceModelsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	modelAttrs := rhaiiInferenceModelMetadataAttributes()
	modelAttrs["id"] = schema.StringAttribute{Computed: true, Description: "The ID of the model. Use this value as `model` in chat completion requests."}
	modelAttrs["object"] = schema.StringAttribute{Computed: true, Description: "The object type, `model`."}
	modelAttrs["created"] = schema.Int64Attribute{Computed: true, Description: "The Unix timestamp in seconds when the model was created."}
	modelAttrs["owned_by"] = schema.StringAttribute{Computed: true, Description: "The owner of the model."}

	resp.Schema = schema.Schema{
		Description: "Lists the inference models that are available in a Red Hat AI Inference project.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{Computed: true, Description: "The ID of the project."},
			"project_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the Red Hat AI Inference project.",
			},
			"location": schema.StringAttribute{
				Optional:    true,
				Description: "The region of the project. When not set, the endpoint from the provider configuration is used.",
			},
			"models": schema.ListNestedAttribute{
				Computed:     true,
				Description:  "The inference models that are available in the project.",
				NestedObject: schema.NestedAttributeObject{Attributes: modelAttrs},
			},
		},
	}
}

func (d *rhaiiInferenceModelsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	meta, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", fmt.Sprintf("Expected conns.ClientSession, got %T.", req.ProviderData))
		return
	}
	d.meta = meta
}

func (d *rhaiiInferenceModelsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m rhaiiInferenceModelsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := rhaiiClientForLocation(d.meta, m.Location.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error initializing the Red Hat AI Inference client", err.Error())
		return
	}
	projectID := m.ProjectID.ValueString()
	collection, _, err := client.ListInferenceModelsWithContext(ctx, client.NewListInferenceModelsOptions(projectID))
	if err != nil {
		resp.Diagnostics.AddError("Error listing inference models", fmt.Sprintf("ListInferenceModelsWithContext failed: %s", err))
		return
	}

	elemType := types.ObjectType{AttrTypes: rhaiiInferenceModelSummaryAttrTypes()}
	models := make([]attr.Value, 0, len(collection.Data))
	for _, model := range collection.Data {
		meta, diags := flattenRhaiiInferenceModelMetadata(model.CustomMetadata)
		resp.Diagnostics.Append(diags...)
		values := meta.metadataAttrValues()
		values["id"] = types.StringValue(strDeref(model.ID))
		values["object"] = types.StringValue(strDeref(model.Object))
		values["owned_by"] = types.StringValue(strDeref(model.OwnedBy))
		values["created"] = types.Int64Null()
		if model.Created != nil {
			values["created"] = types.Int64Value(*model.Created)
		}
		obj, diags := types.ObjectValue(elemType.AttrTypes, values)
		resp.Diagnostics.Append(diags...)
		models = append(models, obj)
	}
	list, diags := types.ListValue(elemType, models)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	m.ID = types.StringValue(projectID)
	m.Models = list
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
