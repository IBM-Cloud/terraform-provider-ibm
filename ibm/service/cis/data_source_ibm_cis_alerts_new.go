package cis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/alertsv1"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &CISAlertsDataSource{}

type CISAlertsDataSource struct {
	client interface{}
}

func NewCISAlertsDataSource(client interface{}) datasource.DataSource {
	return &CISAlertsDataSource{
		client: client,
	}
}

func (d *CISAlertsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cis_alerts_new"
}

func (d *CISAlertsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CIS alerts data source",

		Attributes: map[string]schema.Attribute{
			"cis_id": schema.StringAttribute{
				MarkdownDescription: "CIS instance crn",
				Required:            true,
			},
			"alert_policies": schema.ListAttribute{
				MarkdownDescription: "Container for response information",
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"policy_id":   types.StringType,
						"name":        types.StringType,
						"description": types.StringType,
						"enabled":     types.BoolType,
						"alert_type":  types.StringType,
						"mechanisms": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"email":    types.SetType{ElemType: types.StringType},
									"webhooks": types.SetType{ElemType: types.StringType},
								},
							},
						},
						"filters":    types.StringType,
						"conditions": types.StringType,
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Data source identifier",
			},
		},
	}
}

func (d *CISAlertsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CISAlertsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sess, err := d.client.(conns.ClientSession).CisAlertsSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS alerts session", err.Error())
		return
	}

	crn := data.CisID.ValueString()
	sess.Crn = core.StringPtr(crn)

	opt := sess.NewGetAlertPoliciesOptions()
	result, _, err := sess.GetAlertPolicies(opt)
	if err != nil {
		resp.Diagnostics.AddError("Error getting alert policies", err.Error())
		return
	}

	alertList := make([]attr.Value, 0)
	for _, alertObj := range result.Result {
		// Convert mechanisms
		mechanismsList := make([]attr.Value, 0)
		mechanismObj := d.dataflattenCISMechanism(*alertObj.Mechanisms)
		mechanismsList = append(mechanismsList, mechanismObj)

		mechanismsListValue, _ := types.ListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"email":    types.SetType{ElemType: types.StringType},
					"webhooks": types.SetType{ElemType: types.StringType},
				},
			},
			mechanismsList,
		)

		filterOpt, err := json.Marshal(alertObj.Filters)
		if err != nil {
			resp.Diagnostics.AddError("Error marshalling filters", err.Error())
			return
		}
		normalizedFilter, _ := flex.NormalizeJSONString(string(filterOpt))

		conditionsOpt, err := json.Marshal(alertObj.Conditions)
		if err != nil {
			resp.Diagnostics.AddError("Error marshalling conditions", err.Error())
			return
		}
		normalizedConditions, _ := flex.NormalizeJSONString(string(conditionsOpt))

		alertOutput, _ := types.ObjectValue(
			map[string]attr.Type{
				"policy_id":   types.StringType,
				"name":        types.StringType,
				"description": types.StringType,
				"enabled":     types.BoolType,
				"alert_type":  types.StringType,
				"mechanisms": types.ListType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"email":    types.SetType{ElemType: types.StringType},
							"webhooks": types.SetType{ElemType: types.StringType},
						},
					},
				},
				"filters":    types.StringType,
				"conditions": types.StringType,
			},
			map[string]attr.Value{
				"policy_id":   basetypes.NewStringValue(*alertObj.ID),
				"name":        basetypes.NewStringValue(*alertObj.Name),
				"description": basetypes.NewStringValue(*alertObj.Description),
				"enabled":     basetypes.NewBoolValue(*alertObj.Enabled),
				"alert_type":  basetypes.NewStringValue(*alertObj.AlertType),
				"mechanisms":  mechanismsListValue,
				"filters":     basetypes.NewStringValue(normalizedFilter),
				"conditions":  basetypes.NewStringValue(normalizedConditions),
			},
		)

		alertList = append(alertList, alertOutput)
	}

	data.AlertPolicies, _ = types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"policy_id":   types.StringType,
				"name":        types.StringType,
				"description": types.StringType,
				"enabled":     types.BoolType,
				"alert_type":  types.StringType,
				"mechanisms": types.ListType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"email":    types.SetType{ElemType: types.StringType},
							"webhooks": types.SetType{ElemType: types.StringType},
						},
					},
				},
				"filters":    types.StringType,
				"conditions": types.StringType,
			},
		},
		alertList,
	)

	data.Id = basetypes.NewStringValue(time.Now().UTC().String())
	data.CisID = basetypes.NewStringValue(crn)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *CISAlertsDataSource) dataflattenCISMechanism(mechanism alertsv1.ListAlertPoliciesRespResultItemMechanisms) attr.Value {
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

type CISAlertsDataSourceModel struct {
	CisID         types.String `tfsdk:"cis_id"`
	AlertPolicies types.List   `tfsdk:"alert_policies"`
	Id            types.String `tfsdk:"id"`
}
