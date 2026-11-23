// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx

import (
	"context"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &projectResource{}
	_ resource.ResourceWithConfigure   = &projectResource{}
	_ resource.ResourceWithImportState = &projectResource{}
)

func NewProjectResource() resource.Resource {
	return &projectResource{}
}

type projectResource struct {
	client *client
}

type projectModel struct {
	ID            types.String `tfsdk:"id"`
	Region        types.String `tfsdk:"region"`
	Endpoint      types.String `tfsdk:"endpoint"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Tags          types.List   `tfsdk:"tags"`
	COSInstanceID types.String `tfsdk:"cos_instance_id"`
	COSCRN        types.String `tfsdk:"cos_crn"`
	Delegated     types.Bool   `tfsdk:"storage_delegated"`
	RuntimeName   types.String `tfsdk:"runtime_name"`
	RuntimeGUID   types.String `tfsdk:"runtime_guid"`
	RuntimeCRN    types.String `tfsdk:"runtime_crn"`
	MarkSensitive types.Bool   `tfsdk:"mark_as_sensitive"`
	Location      types.String `tfsdk:"location"`
	ProjectURL    types.String `tfsdk:"project_url"`
}

// Wire types for the Data Platform projects API.
type projectStorage struct {
	Type        string `json:"type"`
	GUID        string `json:"guid"`
	ResourceCRN string `json:"resource_crn"`
	Delegated   bool   `json:"delegated"`
}

type projectCompute struct {
	Name        string         `json:"name"`
	GUID        string         `json:"guid"`
	Type        string         `json:"type"`
	CRN         string         `json:"crn"`
	Credentials map[string]any `json:"credentials,omitempty"`
}

type projectCreateBody struct {
	Name        string           `json:"name"`
	Generator   string           `json:"generator"`
	Type        string           `json:"type"`
	Storage     projectStorage   `json:"storage"`
	Description string           `json:"description"`
	Public      bool             `json:"public"`
	Tags        []string         `json:"tags"`
	Compute     []projectCompute `json:"compute"`
	Settings    map[string]any   `json:"settings"`
}

type projectUpdateBody struct {
	Name        string           `json:"name"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	Public      bool             `json:"public"`
	Compute     []projectCompute `json:"compute"`
}

type projectGetResponse struct {
	Metadata struct {
		GUID string `json:"guid"`
	} `json:"metadata"`
	Entity struct {
		Name        string           `json:"name"`
		Description string           `json:"description"`
		Tags        []string         `json:"tags"`
		Storage     projectStorage   `json:"storage"`
		Compute     []projectCompute `json:"compute"`
	} `json:"entity"`
}

func (r *projectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_watsonx_project"
}

func (r *projectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	replace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		Description: "Creates a watsonx.ai project backed by a Cloud Object Storage instance and bound to a watsonx.ai Runtime instance.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Project GUID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"region":   schema.StringAttribute{Required: true, PlanModifiers: replace, Description: "watsonx region. One of us-south, eu-gb, eu-de, jp-tok, au-syd, ca-tor."},
			"endpoint": schema.StringAttribute{Optional: true, PlanModifiers: replace, Description: "Override for the Data Platform API host. Defaults to the host for region."},
			"name":     schema.StringAttribute{Required: true, Description: "Project name. Must be unique in the account."},
			"description": schema.StringAttribute{
				Optional: true, Computed: true, Default: stringdefault.StaticString(""),
				Description: "Project description.",
			},
			"tags": schema.ListAttribute{
				ElementType: types.StringType, Optional: true,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				Description:   "Project tags. The API only accepts tags at create time, so a change recreates the project.",
			},
			"cos_instance_id": schema.StringAttribute{Required: true, PlanModifiers: replace, Description: "GUID of the Cloud Object Storage instance that stores project assets."},
			"cos_crn":         schema.StringAttribute{Required: true, PlanModifiers: replace, Description: "CRN of the Cloud Object Storage instance."},
			"storage_delegated": schema.BoolAttribute{
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				Description:   "Set to true when the COS instance is delegated with ibm_watsonx_storage_delegation.",
			},
			"runtime_name": schema.StringAttribute{Required: true, Description: "Name of the watsonx.ai Runtime instance."},
			"runtime_guid": schema.StringAttribute{Required: true, Description: "GUID of the watsonx.ai Runtime instance."},
			"runtime_crn":  schema.StringAttribute{Required: true, Description: "CRN of the watsonx.ai Runtime instance."},
			"mark_as_sensitive": schema.BoolAttribute{
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				Description:   "Restrict data in the project from leaving it. Can only be set at create time.",
			},
			"location": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Description: "API path of the project, for example /v2/projects/<guid>.",
			},
			"project_url": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Description: "Link to the project in the watsonx UI.",
			},
		},
	}
}

func (r *projectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, err := newClient(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Configure watsonx Client", err.Error())
		return
	}
	r.client = c
}

func (m *projectModel) compute() []projectCompute {
	return []projectCompute{{
		Name: m.RuntimeName.ValueString(),
		GUID: m.RuntimeGUID.ValueString(),
		Type: "machine_learning",
		CRN:  m.RuntimeCRN.ValueString(),
	}}
}

func (r *projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan projectModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	base, err := apiBase(plan.Region.ValueString(), plan.Endpoint.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}

	tags := []string{}
	if !plan.Tags.IsNull() {
		resp.Diagnostics.Append(plan.Tags.ElementsAs(ctx, &tags, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	body := projectCreateBody{
		Name:      plan.Name.ValueString(),
		Generator: "terraform-provider-ibm",
		Type:      "wx",
		Storage: projectStorage{
			Type:        "bmcos_object_storage",
			GUID:        plan.COSInstanceID.ValueString(),
			ResourceCRN: plan.COSCRN.ValueString(),
			Delegated:   plan.Delegated.ValueBool(),
		},
		Description: plan.Description.ValueString(),
		Public:      true,
		Tags:        tags,
		Compute:     plan.compute(),
		Settings: map[string]any{
			"access_restrictions": map[string]any{"data": plan.MarkSensitive.ValueBool()},
		},
	}

	var out struct {
		Location string `json:"location"`
	}
	if _, err := r.client.do(ctx, http.MethodPost, base+"/transactional/v2/projects?verify_unique_name=true", body, &out); err != nil {
		resp.Diagnostics.AddError("Unable to Create watsonx Project", err.Error())
		return
	}
	id := out.Location[strings.LastIndex(out.Location, "/")+1:]
	if id == "" {
		resp.Diagnostics.AddError("Unexpected Create Response", "The projects API did not return a location for the new project.")
		return
	}

	plan.ID = types.StringValue(id)
	plan.Location = types.StringValue(out.Location)
	plan.ProjectURL = types.StringValue(uiBase(plan.Region.ValueString()) + "/projects/" + id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *projectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state projectModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	base, err := apiBase(state.Region.ValueString(), state.Endpoint.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	var got projectGetResponse
	if _, err := r.client.do(ctx, http.MethodGet, base+"/v2/projects/"+state.ID.ValueString(), nil, &got); err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to Read watsonx Project", err.Error())
		return
	}

	// Only fields the user owns are refreshed; server-generated fields are ignored,
	// which is what ignore_all_server_changes did on the restapi resource.
	state.Name = types.StringValue(got.Entity.Name)
	state.Description = types.StringValue(got.Entity.Description)
	if got.Entity.Storage.GUID != "" {
		state.COSInstanceID = types.StringValue(got.Entity.Storage.GUID)
	}
	for _, c := range got.Entity.Compute {
		if c.Type == "machine_learning" {
			state.RuntimeName = types.StringValue(c.Name)
			state.RuntimeGUID = types.StringValue(c.GUID)
			state.RuntimeCRN = types.StringValue(c.CRN)
			break
		}
	}
	if state.Location.IsNull() || state.Location.ValueString() == "" {
		state.Location = types.StringValue("/v2/projects/" + state.ID.ValueString())
	}
	if state.ProjectURL.IsNull() || state.ProjectURL.ValueString() == "" {
		state.ProjectURL = types.StringValue(uiBase(state.Region.ValueString()) + "/projects/" + state.ID.ValueString())
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *projectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state projectModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	base, err := apiBase(plan.Region.ValueString(), plan.Endpoint.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	compute := plan.compute()
	compute[0].Credentials = map[string]any{}
	body := projectUpdateBody{
		Name:        plan.Name.ValueString(),
		Type:        "wx",
		Description: plan.Description.ValueString(),
		Public:      true,
		Compute:     compute,
	}
	if _, err := r.client.do(ctx, http.MethodPatch, base+"/v2/projects/"+state.ID.ValueString(), body, nil); err != nil {
		resp.Diagnostics.AddError("Unable to Update watsonx Project", err.Error())
		return
	}
	plan.ID, plan.Location, plan.ProjectURL = state.ID, state.Location, state.ProjectURL
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *projectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state projectModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	base, err := apiBase(state.Region.ValueString(), state.Endpoint.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	if _, err := r.client.do(ctx, http.MethodDelete, base+"/transactional/v2/projects/"+state.ID.ValueString(), nil, nil); err != nil && !isNotFound(err) {
		resp.Diagnostics.AddError("Unable to Delete watsonx Project", err.Error())
	}
}

// ImportState takes "<region>/<project_guid>". Storage and runtime fields are
// filled in by the next refresh.
func (r *projectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	region, id, ok := splitImportID(req.ID)
	if !ok {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected <region>/<project_guid>.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, pathOf("id"), id)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, pathOf("region"), region)...)
}
