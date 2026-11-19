// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package accountmanagement

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                   = &accountTraitsResource{}
	_ resource.ResourceWithConfigure      = &accountTraitsResource{}
	_ resource.ResourceWithImportState    = &accountTraitsResource{}
	_ resource.ResourceWithValidateConfig = &accountTraitsResource{}
)

// NewAccountTraitsResource PATCHes
// /v1/accounts/{account_id}/traits.
func NewAccountTraitsResource() resource.Resource {
	return &accountTraitsResource{}
}

type accountTraitsResource struct {
	client *traitsClient
}

type accountTraitsModel struct {
	ID                   types.String `tfsdk:"id"`
	AccountID            types.String `tfsdk:"account_id"`
	Endpoint             types.String `tfsdk:"endpoint"`
	FsReady              types.Bool   `tfsdk:"fs_ready"`
	TeamDirectoryEnabled types.Bool   `tfsdk:"team_directory_enabled"`
}

func (r *accountTraitsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_traits"
}

func (r *accountTraitsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages account traits through the Account Management API (PATCH /v1/accounts/{account_id}/traits): use of Financial Services Validated products, and user list visibility. Only the traits you set are managed. Traits are account settings, so destroying this resource leaves them as they are.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Description: "Same as account_id.",
			},
			"account_id": schema.StringAttribute{
				Optional: true, Computed: true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured(), stringplanmodifier.UseStateForUnknown()},
				Description:   "Account ID. Defaults to the account of the provider's credentials.",
			},
			"endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Account Management API host, for example 'accounts.cloud.ibm.com' or 'private.us-south.accounts.cloud.ibm.com'. Defaults to the endpoint the provider is configured with.",
			},
			"fs_ready": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable use of Financial Services Validated products in the account. Not managed when omitted.",
			},
			"team_directory_enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Let users in the account see the user list (team directory). Set false to restrict user list visibility. Not managed when omitted.",
			},
		},
	}
}

func (r *accountTraitsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var cfg accountTraitsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &cfg)...)
	if resp.Diagnostics.HasError() || cfg.FsReady.IsUnknown() || cfg.TeamDirectoryEnabled.IsUnknown() {
		return
	}
	if cfg.FsReady.IsNull() && cfg.TeamDirectoryEnabled.IsNull() {
		resp.Diagnostics.AddAttributeError(path.Root("fs_ready"), "Missing Trait", "Set at least one of fs_ready or team_directory_enabled.")
	}
}

func (r *accountTraitsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, err := newTraitsClient(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Create Account Management Client", err.Error())
		return
	}
	r.client = c
}

func (r *accountTraitsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan accountTraitsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(r.apply(ctx, &plan)...)
	if !resp.Diagnostics.HasError() {
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	}
}

func (r *accountTraitsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan accountTraitsModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(r.apply(ctx, &plan)...)
	if !resp.Diagnostics.HasError() {
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	}
}

func (r *accountTraitsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state accountTraitsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	client, err := r.client.forEndpoint(state.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Account Management Endpoint", err.Error())
		return
	}
	traits, response, err := getTraits(ctx, client, state.AccountID.ValueString())
	if err != nil {
		if isNotFound(response) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to Read Account Traits", err.Error())
		return
	}
	applyManagedTraits(&state, traits)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Delete only removes the resource from state. Traits are account settings
// with no "unset" state, so the account keeps its last values.
func (r *accountTraitsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

// ImportState takes the account ID.
func (r *accountTraitsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), req.ID)...)
}

func (r *accountTraitsResource) apply(ctx context.Context, m *accountTraitsModel) (diags diag.Diagnostics) {
	client, err := r.client.forEndpoint(m.Endpoint)
	if err != nil {
		diags.AddError("Invalid Account Management Endpoint", err.Error())
		return
	}
	accountID, err := r.client.accountID(m.AccountID)
	if err != nil {
		diags.AddError("Unknown Account", err.Error())
		return
	}
	body := &accountTraits{}
	if !m.FsReady.IsNull() {
		v := m.FsReady.ValueBool()
		body.FsReady = &v
	}
	if !m.TeamDirectoryEnabled.IsNull() {
		v := m.TeamDirectoryEnabled.ValueBool()
		body.TeamDirectoryEnabled = &v
	}
	if err := patchTraits(ctx, client, accountID, body); err != nil {
		diags.AddError("Unable to Update Account Traits", err.Error())
		return
	}
	m.AccountID = types.StringValue(accountID)
	m.ID = m.AccountID
	return
}

// applyManagedTraits refreshes only the traits the config manages, so a trait
// changed outside Terraform shows as drift and an unmanaged one never does.
func applyManagedTraits(m *accountTraitsModel, t *accountTraits) {
	if !m.FsReady.IsNull() && t.FsReady != nil {
		m.FsReady = boolValue(t.FsReady)
	}
	if !m.TeamDirectoryEnabled.IsNull() && t.TeamDirectoryEnabled != nil {
		m.TeamDirectoryEnabled = boolValue(t.TeamDirectoryEnabled)
	}
}
