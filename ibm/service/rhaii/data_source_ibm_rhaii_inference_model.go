// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &rhaiiInferenceModelDataSource{}
	_ datasource.DataSourceWithConfigure = &rhaiiInferenceModelDataSource{}
)

// NewRhaiiInferenceModelDataSource returns the ibm_rhaii_inference_model data
// source, which reads one inference model of a Red Hat AI Inference project.
func NewRhaiiInferenceModelDataSource() datasource.DataSource {
	return &rhaiiInferenceModelDataSource{}
}

type rhaiiInferenceModelDataSource struct {
	meta conns.ClientSession
}

type rhaiiInferenceModelDataSourceModel struct {
	ID                 types.String `tfsdk:"id"`
	ProjectID          types.String `tfsdk:"project_id"`
	Model              types.String `tfsdk:"model"`
	Location           types.String `tfsdk:"location"`
	Identifier         types.String `tfsdk:"identifier"`
	ProviderID         types.String `tfsdk:"provider_id"`
	ProviderResourceID types.String `tfsdk:"provider_resource_id"`
	Type               types.String `tfsdk:"type"`
	ModelType          types.String `tfsdk:"model_type"`
	ModelCard          types.String `tfsdk:"model_card"`
	SourceJSON         types.String `tfsdk:"source_json"`
	rhaiiInferenceModelMetadataModel
}

func (d *rhaiiInferenceModelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rhaii_inference_model"
}

func (d *rhaiiInferenceModelDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attrs := rhaiiInferenceModelMetadataAttributes()
	str := func(desc string) schema.Attribute { return schema.StringAttribute{Computed: true, Description: desc} }
	attrs["id"] = str("The ID of the data source, in `<project_id>/<model>` format.")
	attrs["project_id"] = schema.StringAttribute{Required: true, Description: "The ID of the Red Hat AI Inference project."}
	attrs["model"] = schema.StringAttribute{
		Required:    true,
		Description: "The ID of the model, for example `llama-3-3-70b-instruct`.",
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 100),
			stringvalidator.RegexMatches(rhaiiModelIDRegexp, "can contain only lower case letters, numbers and -"),
		},
	}
	attrs["location"] = schema.StringAttribute{Optional: true, Description: "The region of the project. When not set, the endpoint from the provider configuration is used."}
	attrs["identifier"] = str("The unique identifier of the model.")
	attrs["provider_id"] = str("The ID of the provider that owns the model.")
	attrs["provider_resource_id"] = str("The ID of the model in the provider.")
	attrs["type"] = str("The resource type, `model`.")
	attrs["model_type"] = str("The model type, for example `llm`, `embedding` or `rerank`.")
	attrs["model_card"] = str("The model card in markdown format.")
	attrs["source_json"] = str("The source of a custom model (huggingface, ibmcos, oci or s3) as a JSON string.")

	resp.Schema = schema.Schema{
		Description: "Reads an inference model of a Red Hat AI Inference project.",
		Attributes:  attrs,
	}
}

func (d *rhaiiInferenceModelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *rhaiiInferenceModelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m rhaiiInferenceModelDataSourceModel
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
	modelID := m.Model.ValueString()
	model, _, err := client.GetInferenceModelWithContext(ctx, client.NewGetInferenceModelOptions(projectID, modelID))
	if err != nil {
		resp.Diagnostics.AddError("Error reading the inference model", fmt.Sprintf("GetInferenceModelWithContext failed: %s", err))
		return
	}

	metadata, diags := flattenRhaiiInferenceModelMetadata(model.Metadata)
	resp.Diagnostics.Append(diags...)
	m.rhaiiInferenceModelMetadataModel = metadata
	m.ModelCard = types.StringValue("")
	m.SourceJSON = types.StringValue("")
	if model.Metadata != nil {
		m.ModelCard = types.StringValue(strDeref(model.Metadata.ModelCard))
		sourceJSON, err := rhaiiToJSON(model.Metadata.Source)
		if err != nil {
			resp.Diagnostics.AddError("Error converting the model source to JSON", err.Error())
		}
		m.SourceJSON = types.StringValue(sourceJSON)
	}
	m.ID = types.StringValue(fmt.Sprintf("%s/%s", projectID, modelID))
	m.Identifier = types.StringValue(strDeref(model.Identifier))
	m.ProviderID = types.StringValue(strDeref(model.ProviderID))
	m.ProviderResourceID = types.StringValue(strDeref(model.ProviderResourceID))
	m.Type = types.StringValue(strDeref(model.Type))
	m.ModelType = types.StringValue(strDeref(model.ModelType))

	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
