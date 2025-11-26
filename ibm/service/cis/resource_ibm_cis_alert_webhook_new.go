package cis

import (
	"context"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.Resource = &CISWebhookResource{}
var _ resource.ResourceWithImportState = &CISWebhookResource{}

type CISWebhookResource struct {
	client interface{}
}

func NewCISWebhookResource(client interface{}) resource.Resource {
	return &CISWebhookResource{
		client: client,
	}
}

func (r *CISWebhookResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cis_alert_webhook_new"
}

func (r *CISWebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CIS alert webhook resource",

		Attributes: map[string]schema.Attribute{
			"cis_id": schema.StringAttribute{
				MarkdownDescription: "CIS instance crn",
				Required:            true,
			},
			"webhook_id": schema.StringAttribute{
				MarkdownDescription: "Webhook ID",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Webhook Name",
				Required:            true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "Webhook URL",
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Webhook Type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"secret": schema.StringAttribute{
				MarkdownDescription: "API key needed to use the webhook",
				Optional:            true,
				Sensitive:           true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Webhook identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *CISWebhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CISWebhookResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sess, err := r.client.(conns.ClientSession).CisWebhookSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS webhook session", err.Error())
		return
	}

	crn := plan.CisID.ValueString()
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewCreateAlertWebhookOptions()

	if !plan.Name.IsNull() {
		opt.SetName(plan.Name.ValueString())
	}
	if !plan.URL.IsNull() {
		opt.SetURL(plan.URL.ValueString())
	}
	if !plan.Secret.IsNull() {
		opt.SetSecret(plan.Secret.ValueString())
	}

	result, _, err := sess.CreateAlertWebhook(opt)
	if err != nil || result == nil {
		resp.Diagnostics.AddError("Error creating webhook", err.Error())
		return
	}

	plan.Id = basetypes.NewStringValue(flex.ConvertCisToTfTwoVar(*result.Result.ID, crn))
	plan.WebhookID = basetypes.NewStringValue(*result.Result.ID)

	// Read to populate full state
	err = r.readWebhook(ctx, &plan, crn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading webhook after create", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CISWebhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CISWebhookResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	webhookID, crn, err := flex.ConvertTftoCisTwoVar(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error parsing ID", err.Error())
		return
	}

	sess, err := r.client.(conns.ClientSession).CisWebhookSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS webhook session", err.Error())
		return
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewGetWebhookOptions(webhookID)

	result, response, err := sess.GetWebhook(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error getting webhook", err.Error())
		return
	}

	state.CisID = basetypes.NewStringValue(crn)
	state.WebhookID = basetypes.NewStringValue(*result.Result.ID)
	state.Name = basetypes.NewStringValue(*result.Result.Name)
	state.URL = basetypes.NewStringValue(*result.Result.URL)
	state.Type = basetypes.NewStringValue(*result.Result.Type)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CISWebhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CISWebhookResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sess, err := r.client.(conns.ClientSession).CisWebhookSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS webhook session", err.Error())
		return
	}

	webhookID, crn, err := flex.ConvertTftoCisTwoVar(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error parsing ID", err.Error())
		return
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewUpdateAlertWebhookOptions(webhookID)

	if !plan.Name.IsNull() {
		opt.SetName(plan.Name.ValueString())
	}
	if !plan.URL.IsNull() {
		opt.SetURL(plan.URL.ValueString())
	}
	if !plan.Secret.IsNull() {
		opt.SetSecret(plan.Secret.ValueString())
	}

	result, _, err := sess.UpdateAlertWebhook(opt)
	if err != nil || result == nil {
		resp.Diagnostics.AddError("Error updating webhook", err.Error())
		return
	}

	// Read to get latest state
	err = r.readWebhook(ctx, &plan, crn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading webhook after update", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CISWebhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CISWebhookResourceModel
	req.State.Get(ctx, &state)

	sess, err := r.client.(conns.ClientSession).CisWebhookSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS webhook session", err.Error())
		return
	}

	webhookID, crn, err := flex.ConvertTftoCisTwoVar(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error parsing ID", err.Error())
		return
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewDeleteWebhookOptions(webhookID)

	_, response, err := sess.DeleteWebhook(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error deleting webhook", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *CISWebhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CISWebhookResource) readWebhook(ctx context.Context, model *CISWebhookResourceModel, crn string) error {
	sess, err := r.client.(conns.ClientSession).CisWebhookSession()
	if err != nil {
		return err
	}

	webhookID, _, err := flex.ConvertTftoCisTwoVar(model.Id.ValueString())
	if err != nil {
		return err
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewGetWebhookOptions(webhookID)
	result, _, err := sess.GetWebhook(opt)
	if err != nil {
		return err
	}

	model.CisID = basetypes.NewStringValue(crn)
	model.WebhookID = basetypes.NewStringValue(*result.Result.ID)
	model.Name = basetypes.NewStringValue(*result.Result.Name)
	model.URL = basetypes.NewStringValue(*result.Result.URL)
	model.Type = basetypes.NewStringValue(*result.Result.Type)

	return nil
}

type CISWebhookResourceModel struct {
	CisID     types.String `tfsdk:"cis_id"`
	WebhookID types.String `tfsdk:"webhook_id"`
	Name      types.String `tfsdk:"name"`
	URL       types.String `tfsdk:"url"`
	Type      types.String `tfsdk:"type"`
	Secret    types.String `tfsdk:"secret"`
	Id        types.String `tfsdk:"id"`
}
