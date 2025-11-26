package cis

import (
	"context"
	"encoding/json"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/alertsv1"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ resource.Resource = &CISAlertResource{}
var _ resource.ResourceWithImportState = &CISAlertResource{}

type CISAlertResource struct {
	client interface{}
}

func NewCISAlertResource(client interface{}) resource.Resource {
	return &CISAlertResource{
		client: client,
	}
}

func (r *CISAlertResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cis_alert_new"
}

func (r *CISAlertResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CIS alert resource",

		Attributes: map[string]schema.Attribute{
			"cis_id": schema.StringAttribute{
				MarkdownDescription: "CIS instance crn",
				Required:            true,
			},
			"policy_id": schema.StringAttribute{
				MarkdownDescription: "Identifier of the Alert Policy",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Policy name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Policy Description",
				Optional:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Is the alert policy active",
				Required:            true,
			},
			"alert_type": schema.StringAttribute{
				MarkdownDescription: "Condition for the alert",
				Required:            true,
			},
			"mechanisms": schema.ListAttribute{
				MarkdownDescription: "Delivery mechanisms for the alert",
				Required:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"email":    types.SetType{ElemType: types.StringType},
						"webhooks": types.SetType{ElemType: types.StringType},
					},
				},
			},
			"filters": schema.StringAttribute{
				MarkdownDescription: "Filters based on filter type (JSON string)",
				Optional:            true,
			},
			"conditions": schema.StringAttribute{
				MarkdownDescription: "Conditions based on filter type (JSON string)",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Alert identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *CISAlertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CISAlertResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sess, err := r.client.(conns.ClientSession).CisAlertsSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS alerts session", err.Error())
		return
	}

	crn := plan.CisID.ValueString()
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewCreateAlertPolicyOptions()

	if !plan.Name.IsNull() {
		opt.SetName(plan.Name.ValueString())
	}
	if !plan.Description.IsNull() {
		opt.SetDescription(plan.Description.ValueString())
	}
	if !plan.Enabled.IsNull() {
		opt.SetEnabled(plan.Enabled.ValueBool())
	}
	if !plan.AlertType.IsNull() {
		opt.SetAlertType(plan.AlertType.ValueString())
	}
	if !plan.Filters.IsNull() && plan.Filters.ValueString() != "" {
		var filter interface{}
		json.Unmarshal([]byte(plan.Filters.ValueString()), &filter)
		opt.Filters = filter
	}

	mechanismsOpt := &alertsv1.CreateAlertPolicyInputMechanisms{}
	if !plan.Mechanisms.IsNull() && len(plan.Mechanisms.Elements()) > 0 {
		mechanismsList := make([]CISMechanismModel, 0, len(plan.Mechanisms.Elements()))
		resp.Diagnostics.Append(plan.Mechanisms.ElementsAs(ctx, &mechanismsList, false)...)

		if len(mechanismsList) > 0 {
			mechanism := mechanismsList[0]

			if !mechanism.Webhooks.IsNull() && len(mechanism.Webhooks.Elements()) > 0 {
				webhookList := make([]string, 0, len(mechanism.Webhooks.Elements()))
				resp.Diagnostics.Append(mechanism.Webhooks.ElementsAs(ctx, &webhookList, false)...)

				var webhookarray = make([]alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem, len(webhookList))
				for k, w := range webhookList {
					webhookarray[k] = alertsv1.CreateAlertPolicyInputMechanismsWebhooksItem{
						ID: &w,
					}
				}
				mechanismsOpt.Webhooks = webhookarray
			}

			if !mechanism.Email.IsNull() && len(mechanism.Email.Elements()) > 0 {
				emailList := make([]string, 0, len(mechanism.Email.Elements()))
				resp.Diagnostics.Append(mechanism.Email.ElementsAs(ctx, &emailList, false)...)

				var emailarray = make([]alertsv1.CreateAlertPolicyInputMechanismsEmailItem, len(emailList))
				for k, e := range emailList {
					emailarray[k] = alertsv1.CreateAlertPolicyInputMechanismsEmailItem{
						ID: &e,
					}
				}
				mechanismsOpt.Email = emailarray
			}
		}
	}
	opt.Mechanisms = mechanismsOpt

	result, _, err := sess.CreateAlertPolicy(opt)
	if err != nil || result == nil {
		resp.Diagnostics.AddError("Error creating alert policy", err.Error())
		return
	}

	plan.Id = basetypes.NewStringValue(flex.ConvertCisToTfTwoVar(*result.Result.ID, crn))
	plan.PolicyID = basetypes.NewStringValue(*result.Result.ID)

	// Read to populate full state
	err = r.readAlert(ctx, &plan, crn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading alert after create", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CISAlertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CISAlertResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	alertID, crn, err := flex.ConvertTftoCisTwoVar(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error parsing ID", err.Error())
		return
	}

	sess, err := r.client.(conns.ClientSession).CisAlertsSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS alerts session", err.Error())
		return
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewGetAlertPolicyOptions(alertID)
	result, response, err := sess.GetAlertPolicy(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error getting alert policy", err.Error())
		return
	}

	state.CisID = basetypes.NewStringValue(crn)
	state.PolicyID = basetypes.NewStringValue(*result.Result.ID)
	state.Name = basetypes.NewStringValue(*result.Result.Name)
	state.Description = basetypes.NewStringValue(*result.Result.Description)
	state.Enabled = basetypes.NewBoolValue(*result.Result.Enabled)
	state.AlertType = basetypes.NewStringValue(*result.Result.AlertType)

	// Convert mechanisms
	mechanismsList := make([]attr.Value, 0)
	mechanismObj := r.flattenCISMechanism(*result.Result.Mechanisms)
	mechanismsList = append(mechanismsList, mechanismObj)

	state.Mechanisms, _ = types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email":    types.SetType{ElemType: types.StringType},
				"webhooks": types.SetType{ElemType: types.StringType},
			},
		},
		mechanismsList,
	)

	filterOpt, err := json.Marshal(result.Result.Filters)
	if err != nil {
		resp.Diagnostics.AddError("Error marshalling filters", err.Error())
		return
	}
	normalizedFilter, _ := flex.NormalizeJSONString(string(filterOpt))
	state.Filters = basetypes.NewStringValue(normalizedFilter)

	conditionsOpt, err := json.Marshal(result.Result.Conditions)
	if err != nil {
		resp.Diagnostics.AddError("Error marshalling conditions", err.Error())
		return
	}
	normalizedConditions, _ := flex.NormalizeJSONString(string(conditionsOpt))
	state.Conditions = basetypes.NewStringValue(normalizedConditions)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CISAlertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CISAlertResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sess, err := r.client.(conns.ClientSession).CisAlertsSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS alerts session", err.Error())
		return
	}

	alertID, crn, err := flex.ConvertTftoCisTwoVar(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error parsing ID", err.Error())
		return
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewUpdateAlertPolicyOptions(alertID)

	if !plan.Name.IsNull() {
		opt.SetName(plan.Name.ValueString())
	}
	if !plan.Description.IsNull() {
		opt.SetDescription(plan.Description.ValueString())
	}
	if !plan.Enabled.IsNull() {
		opt.SetEnabled(plan.Enabled.ValueBool())
	}
	if !plan.AlertType.IsNull() {
		opt.SetAlertType(plan.AlertType.ValueString())
	}
	if !plan.Conditions.IsNull() && plan.Conditions.ValueString() != "" {
		var condition interface{}
		json.Unmarshal([]byte(plan.Conditions.ValueString()), &condition)
		opt.Conditions = condition
	}
	if !plan.Filters.IsNull() && plan.Filters.ValueString() != "" {
		var filter interface{}
		json.Unmarshal([]byte(plan.Filters.ValueString()), &filter)
		opt.Filters = filter
	}

	mechanismsOpt := &alertsv1.UpdateAlertPolicyInputMechanisms{}
	if !plan.Mechanisms.IsNull() && len(plan.Mechanisms.Elements()) > 0 {
		mechanismsList := make([]CISMechanismModel, 0, len(plan.Mechanisms.Elements()))
		resp.Diagnostics.Append(plan.Mechanisms.ElementsAs(ctx, &mechanismsList, false)...)

		if len(mechanismsList) > 0 {
			mechanism := mechanismsList[0]

			if !mechanism.Webhooks.IsNull() && len(mechanism.Webhooks.Elements()) > 0 {
				webhookList := make([]string, 0, len(mechanism.Webhooks.Elements()))
				resp.Diagnostics.Append(mechanism.Webhooks.ElementsAs(ctx, &webhookList, false)...)

				var webhookarray = make([]alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem, len(webhookList))
				for k, w := range webhookList {
					webhookarray[k] = alertsv1.UpdateAlertPolicyInputMechanismsWebhooksItem{
						ID: &w,
					}
				}
				mechanismsOpt.Webhooks = webhookarray
			}

			if !mechanism.Email.IsNull() && len(mechanism.Email.Elements()) > 0 {
				emailList := make([]string, 0, len(mechanism.Email.Elements()))
				resp.Diagnostics.Append(mechanism.Email.ElementsAs(ctx, &emailList, false)...)

				var emailarray = make([]alertsv1.UpdateAlertPolicyInputMechanismsEmailItem, len(emailList))
				for k, e := range emailList {
					emailarray[k] = alertsv1.UpdateAlertPolicyInputMechanismsEmailItem{
						ID: &e,
					}
				}
				mechanismsOpt.Email = emailarray
			}
		}
	}
	opt.Mechanisms = mechanismsOpt

	result, _, err := sess.UpdateAlertPolicy(opt)
	if err != nil || result == nil {
		resp.Diagnostics.AddError("Error updating alert policy", err.Error())
		return
	}

	// Read to get latest state
	err = r.readAlert(ctx, &plan, crn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading alert after update", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CISAlertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CISAlertResourceModel
	req.State.Get(ctx, &state)

	sess, err := r.client.(conns.ClientSession).CisAlertsSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS alerts session", err.Error())
		return
	}

	alertID, crn, err := flex.ConvertTftoCisTwoVar(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error parsing ID", err.Error())
		return
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewDeleteAlertPolicyOptions(alertID)
	_, response, err := sess.DeleteAlertPolicy(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error deleting alert", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *CISAlertResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CISAlertResource) readAlert(ctx context.Context, model *CISAlertResourceModel, crn string) error {
	sess, err := r.client.(conns.ClientSession).CisAlertsSession()
	if err != nil {
		return err
	}

	alertID, _, err := flex.ConvertTftoCisTwoVar(model.Id.ValueString())
	if err != nil {
		return err
	}

	sess.Crn = core.StringPtr(crn)
	opt := sess.NewGetAlertPolicyOptions(alertID)
	result, _, err := sess.GetAlertPolicy(opt)
	if err != nil {
		return err
	}

	model.CisID = basetypes.NewStringValue(crn)
	model.PolicyID = basetypes.NewStringValue(*result.Result.ID)
	model.Name = basetypes.NewStringValue(*result.Result.Name)
	model.Description = basetypes.NewStringValue(*result.Result.Description)
	model.Enabled = basetypes.NewBoolValue(*result.Result.Enabled)
	model.AlertType = basetypes.NewStringValue(*result.Result.AlertType)

	// Convert mechanisms
	mechanismsList := make([]attr.Value, 0)
	mechanismObj := r.flattenCISMechanism(*result.Result.Mechanisms)
	mechanismsList = append(mechanismsList, mechanismObj)

	model.Mechanisms, _ = types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email":    types.SetType{ElemType: types.StringType},
				"webhooks": types.SetType{ElemType: types.StringType},
			},
		},
		mechanismsList,
	)

	filterOpt, _ := json.Marshal(result.Result.Filters)
	normalizedFilter, _ := flex.NormalizeJSONString(string(filterOpt))
	model.Filters = basetypes.NewStringValue(normalizedFilter)

	conditionsOpt, _ := json.Marshal(result.Result.Conditions)
	normalizedConditions, _ := flex.NormalizeJSONString(string(conditionsOpt))
	model.Conditions = basetypes.NewStringValue(normalizedConditions)

	return nil
}

func (r *CISAlertResource) flattenCISMechanism(mechanism alertsv1.GetAlertPolicyRespResultMechanisms) attr.Value {
	emailoutput := make([]attr.Value, 0)
	webhookoutput := make([]attr.Value, 0)

	for _, mech := range mechanism.Email {
		emailoutput = append(emailoutput, basetypes.NewStringValue(*mech.ID))
	}

	for _, mech := range mechanism.Webhooks {
		webhookoutput = append(webhookoutput, basetypes.NewStringValue(*mech.ID))
	}

	emailSet, _ := types.SetValue(types.StringType, emailoutput)
	webhookSet, _ := types.SetValue(types.StringType, webhookoutput)

	mechanismObj, _ := types.ObjectValue(
		map[string]attr.Type{
			"email":    types.SetType{ElemType: types.StringType},
			"webhooks": types.SetType{ElemType: types.StringType},
		},
		map[string]attr.Value{
			"email":    emailSet,
			"webhooks": webhookSet,
		},
	)

	return mechanismObj
}

type CISAlertResourceModel struct {
	CisID       types.String `tfsdk:"cis_id"`
	PolicyID    types.String `tfsdk:"policy_id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	AlertType   types.String `tfsdk:"alert_type"`
	Mechanisms  types.List   `tfsdk:"mechanisms"`
	Filters     types.String `tfsdk:"filters"`
	Conditions  types.String `tfsdk:"conditions"`
	Id          types.String `tfsdk:"id"`
}

type CISMechanismModel struct {
	Email    types.Set `tfsdk:"email"`
	Webhooks types.Set `tfsdk:"webhooks"`
}
