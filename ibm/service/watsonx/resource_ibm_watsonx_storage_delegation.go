// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &storageDelegationResource{}
	_ resource.ResourceWithConfigure   = &storageDelegationResource{}
	_ resource.ResourceWithImportState = &storageDelegationResource{}
)

func NewStorageDelegationResource() resource.Resource {
	return &storageDelegationResource{}
}

type storageDelegationResource struct {
	client *client
}

type storageDelegationModel struct {
	ID            types.String `tfsdk:"id"`
	Region        types.String `tfsdk:"region"`
	Endpoint      types.String `tfsdk:"endpoint"`
	COSInstanceID types.String `tfsdk:"cos_instance_id"`
	KMSKeyCRN     types.String `tfsdk:"kms_key_crn"`
	Catalogs      types.Bool   `tfsdk:"catalogs"`
	Projects      types.Bool   `tfsdk:"projects"`
}

type storageDelegationBody struct {
	COSInstanceID string `json:"cos_instance_id,omitempty"`
	KMSKeyCRN     string `json:"kms_key_crn,omitempty"`
	Catalogs      *bool  `json:"catalogs,omitempty"`
	Projects      *bool  `json:"projects,omitempty"`
}

func (r *storageDelegationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_watsonx_storage_delegation"
}

func (r *storageDelegationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Delegates a Cloud Object Storage instance to watsonx so projects and catalogs created on it are encrypted with your Key Protect or HPCS key. Create the COS to KMS authorization policy before this resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Same as cos_instance_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"region": schema.StringAttribute{
				Required:      true,
				Description:   "watsonx region. One of us-south, eu-gb, eu-de, jp-tok, au-syd, ca-tor.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"endpoint": schema.StringAttribute{
				Optional:      true,
				Description:   "Override for the Data Platform UI host that serves /api/rest/v1/storage-delegations. Defaults to the host for region.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cos_instance_id": schema.StringAttribute{
				Required:      true,
				Description:   "GUID of the Cloud Object Storage instance.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"kms_key_crn": schema.StringAttribute{
				Required:    true,
				Description: "CRN of the Key Protect or Hyper Protect Crypto Services root key.",
			},
			"catalogs": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Delegate the storage for catalogs. Default: true.",
			},
			"projects": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Delegate the storage for projects. Default: true.",
			},
		},
	}
}

func (r *storageDelegationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *storageDelegationResource) collectionURL(m *storageDelegationModel) (string, error) {
	base := ""
	if ep := m.Endpoint.ValueString(); ep != "" {
		base = ensureHTTPS(ep)
	} else {
		base = uiBase(m.Region.ValueString())
	}
	if base == "" {
		return "", fmt.Errorf("region %q is not supported, use one of %v or set endpoint", m.Region.ValueString(), supportedRegions())
	}
	return base + "/api/rest/v1/storage-delegations", nil
}

func (r *storageDelegationResource) body(m *storageDelegationModel) storageDelegationBody {
	catalogs, projects := m.Catalogs.ValueBool(), m.Projects.ValueBool()
	return storageDelegationBody{
		COSInstanceID: m.COSInstanceID.ValueString(),
		KMSKeyCRN:     m.KMSKeyCRN.ValueString(),
		Catalogs:      &catalogs,
		Projects:      &projects,
	}
}

func (r *storageDelegationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan storageDelegationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	url, err := r.collectionURL(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	if _, err := r.client.do(ctx, http.MethodPost, url, r.body(&plan), nil); err != nil {
		resp.Diagnostics.AddError("Unable to Create Storage Delegation", err.Error())
		return
	}
	plan.ID = plan.COSInstanceID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *storageDelegationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state storageDelegationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	url, err := r.collectionURL(&state)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	var got storageDelegationBody
	if _, err := r.client.do(ctx, http.MethodGet, url+"/"+state.COSInstanceID.ValueString(), nil, &got); err != nil {
		if isNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to Read Storage Delegation", err.Error())
		return
	}
	if got.KMSKeyCRN != "" {
		state.KMSKeyCRN = types.StringValue(got.KMSKeyCRN)
	}
	if got.Catalogs != nil {
		state.Catalogs = types.BoolValue(*got.Catalogs)
	}
	if got.Projects != nil {
		state.Projects = types.BoolValue(*got.Projects)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *storageDelegationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan storageDelegationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	url, err := r.collectionURL(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	if _, err := r.client.do(ctx, http.MethodPatch, url+"/"+plan.COSInstanceID.ValueString(), r.body(&plan), nil); err != nil {
		resp.Diagnostics.AddError("Unable to Update Storage Delegation", err.Error())
		return
	}
	plan.ID = plan.COSInstanceID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *storageDelegationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state storageDelegationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	url, err := r.collectionURL(&state)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Region", err.Error())
		return
	}
	if _, err := r.client.do(ctx, http.MethodDelete, url+"/"+state.COSInstanceID.ValueString(), nil, nil); err != nil && !isNotFound(err) {
		resp.Diagnostics.AddError("Unable to Delete Storage Delegation", err.Error())
	}
}

// ImportState takes "<region>/<cos_instance_id>".
func (r *storageDelegationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	region, id, ok := splitImportID(req.ID)
	if !ok {
		resp.Diagnostics.AddError("Invalid Import ID", "Expected <region>/<cos_instance_id>.")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, pathOf("id"), id)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, pathOf("cos_instance_id"), id)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, pathOf("region"), region)...)
}
