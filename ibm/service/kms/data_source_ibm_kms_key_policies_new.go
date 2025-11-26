package kms

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	kp "github.com/IBM/keyprotect-go-client"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &KMSKeyPoliciesDataSource{}

type KMSKeyPoliciesDataSource struct {
	client interface{}
}

func NewKMSKeyPoliciesDataSource(client interface{}) datasource.DataSource {
	return &KMSKeyPoliciesDataSource{
		client: client,
	}
}

func (d *KMSKeyPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kms_key_policies_new"
}

func (d *KMSKeyPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "KMS key policies data source",

		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Key protect or hpcs instance GUID",
				Required:            true,
			},
			"endpoint_type": schema.StringAttribute{
				MarkdownDescription: "public or private",
				Optional:            true,
				Computed:            true,
			},
			"key_id": schema.StringAttribute{
				MarkdownDescription: "Key ID of the Key",
				Optional:            true,
			},
			"alias": schema.StringAttribute{
				MarkdownDescription: "Alias of the Key",
				Optional:            true,
			},
			"policies": schema.ListAttribute{
				MarkdownDescription: "Creates or updates one or more policies for the specified key",
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rotation": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"id":               types.StringType,
									"crn":              types.StringType,
									"created_by":       types.StringType,
									"creation_date":    types.StringType,
									"updated_by":       types.StringType,
									"last_update_date": types.StringType,
									"interval_month":   types.Int64Type,
									"enabled":          types.BoolType,
								},
							},
						},
						"dual_auth_delete": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"id":               types.StringType,
									"crn":              types.StringType,
									"created_by":       types.StringType,
									"creation_date":    types.StringType,
									"updated_by":       types.StringType,
									"last_update_date": types.StringType,
									"enabled":          types.BoolType,
								},
							},
						},
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

func (d *KMSKeyPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data KMSKeyPoliciesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	instanceID := getInstanceIDFromCRN(data.InstanceID.ValueString())
	api, _, err := d.populateKPClientForPolicies(data, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	var id string
	if !data.KeyID.IsNull() && data.KeyID.ValueString() != "" {
		id = data.KeyID.ValueString()
	} else if !data.Alias.IsNull() && data.Alias.ValueString() != "" {
		id = data.Alias.ValueString()
		key, err := api.GetKey(ctx, id)
		if err != nil {
			resp.Diagnostics.AddError("Failed to get Key", err.Error())
			return
		}
		data.KeyID = basetypes.NewStringValue(key.ID)
	}

	policies, err := api.GetPolicies(ctx, id)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read policies", err.Error())
		return
	}

	if len(policies) == 0 {
		log.Printf("No Policy Configurations read\n")
	}

	// Convert policies using flex helper
	policiesFlat := flex.FlattenKeyPolicies(policies)
	policyList := make([]attr.Value, 0)

	if policiesFlat != nil && len(policiesFlat) > 0 {
		for _, policyMap := range policiesFlat {

			rotationList := types.ListNull(types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"id":               types.StringType,
					"crn":              types.StringType,
					"created_by":       types.StringType,
					"creation_date":    types.StringType,
					"updated_by":       types.StringType,
					"last_update_date": types.StringType,
					"interval_month":   types.Int64Type,
					"enabled":          types.BoolType,
				},
			})

			dualAuthList := types.ListNull(types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"id":               types.StringType,
					"crn":              types.StringType,
					"created_by":       types.StringType,
					"creation_date":    types.StringType,
					"updated_by":       types.StringType,
					"last_update_date": types.StringType,
					"enabled":          types.BoolType,
				},
			})

			if rotation, ok := policyMap["rotation"].([]interface{}); ok && len(rotation) > 0 {
				rotItems := make([]attr.Value, 0)
				for _, r := range rotation {
					rm := r.(map[string]interface{})
					rotObj, _ := types.ObjectValue(
						map[string]attr.Type{
							"id":               types.StringType,
							"crn":              types.StringType,
							"created_by":       types.StringType,
							"creation_date":    types.StringType,
							"updated_by":       types.StringType,
							"last_update_date": types.StringType,
							"interval_month":   types.Int64Type,
							"enabled":          types.BoolType,
						},
						map[string]attr.Value{
							"id":               basetypes.NewStringValue(rm["id"].(string)),
							"crn":              basetypes.NewStringValue(rm["crn"].(string)),
							"created_by":       basetypes.NewStringValue(rm["created_by"].(string)),
							"creation_date":    basetypes.NewStringValue(rm["creation_date"].(string)),
							"updated_by":       basetypes.NewStringValue(rm["updated_by"].(string)),
							"last_update_date": basetypes.NewStringValue(rm["last_update_date"].(string)),
							"interval_month":   basetypes.NewInt64Value(int64(rm["interval_month"].(int))),
							"enabled":          basetypes.NewBoolValue(rm["enabled"].(bool)),
						},
					)
					rotItems = append(rotItems, rotObj)
				}
				rotationList, _ = types.ListValue(types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":               types.StringType,
						"crn":              types.StringType,
						"created_by":       types.StringType,
						"creation_date":    types.StringType,
						"updated_by":       types.StringType,
						"last_update_date": types.StringType,
						"interval_month":   types.Int64Type,
						"enabled":          types.BoolType,
					},
				}, rotItems)
			}

			if dualAuth, ok := policyMap["dual_auth_delete"].([]interface{}); ok && len(dualAuth) > 0 {
				dualItems := make([]attr.Value, 0)
				for _, da := range dualAuth {
					dam := da.(map[string]interface{})
					dualObj, _ := types.ObjectValue(
						map[string]attr.Type{
							"id":               types.StringType,
							"crn":              types.StringType,
							"created_by":       types.StringType,
							"creation_date":    types.StringType,
							"updated_by":       types.StringType,
							"last_update_date": types.StringType,
							"enabled":          types.BoolType,
						},
						map[string]attr.Value{
							"id":               basetypes.NewStringValue(dam["id"].(string)),
							"crn":              basetypes.NewStringValue(dam["crn"].(string)),
							"created_by":       basetypes.NewStringValue(dam["created_by"].(string)),
							"creation_date":    basetypes.NewStringValue(dam["creation_date"].(string)),
							"updated_by":       basetypes.NewStringValue(dam["updated_by"].(string)),
							"last_update_date": basetypes.NewStringValue(dam["last_update_date"].(string)),
							"enabled":          basetypes.NewBoolValue(dam["enabled"].(bool)),
						},
					)
					dualItems = append(dualItems, dualObj)
				}
				dualAuthList, _ = types.ListValue(types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":               types.StringType,
						"crn":              types.StringType,
						"created_by":       types.StringType,
						"creation_date":    types.StringType,
						"updated_by":       types.StringType,
						"last_update_date": types.StringType,
						"enabled":          types.BoolType,
					},
				}, dualItems)
			}

			policyObj, _ := types.ObjectValue(
				map[string]attr.Type{
					"rotation": types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"id":               types.StringType,
								"crn":              types.StringType,
								"created_by":       types.StringType,
								"creation_date":    types.StringType,
								"updated_by":       types.StringType,
								"last_update_date": types.StringType,
								"interval_month":   types.Int64Type,
								"enabled":          types.BoolType,
							},
						},
					},
					"dual_auth_delete": types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"id":               types.StringType,
								"crn":              types.StringType,
								"created_by":       types.StringType,
								"creation_date":    types.StringType,
								"updated_by":       types.StringType,
								"last_update_date": types.StringType,
								"enabled":          types.BoolType,
							},
						},
					},
				},
				map[string]attr.Value{
					"rotation":         rotationList,
					"dual_auth_delete": dualAuthList,
				},
			)
			policyList = append(policyList, policyObj)
		}
	}

	data.Policies, _ = types.ListValue(types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"rotation": types.ListType{
				ElemType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":               types.StringType,
						"crn":              types.StringType,
						"created_by":       types.StringType,
						"creation_date":    types.StringType,
						"updated_by":       types.StringType,
						"last_update_date": types.StringType,
						"interval_month":   types.Int64Type,
						"enabled":          types.BoolType,
					},
				},
			},
			"dual_auth_delete": types.ListType{
				ElemType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":               types.StringType,
						"crn":              types.StringType,
						"created_by":       types.StringType,
						"creation_date":    types.StringType,
						"updated_by":       types.StringType,
						"last_update_date": types.StringType,
						"enabled":          types.BoolType,
					},
				},
			},
		},
	}, policyList)

	data.Id = basetypes.NewStringValue(instanceID)
	data.InstanceID = basetypes.NewStringValue(instanceID)

	if data.EndpointType.IsNull() {
		data.EndpointType = basetypes.NewStringValue("public")
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *KMSKeyPoliciesDataSource) populateKPClientForPolicies(model KMSKeyPoliciesDataSourceModel, instanceID string) (*kp.Client, *string, error) {
	kpAPI, err := d.client.(conns.ClientSession).KeyManagementAPI()
	if err != nil {
		return nil, nil, err
	}

	var endpointType string
	if !model.EndpointType.IsNull() {
		endpointType = model.EndpointType.ValueString()
	} else {
		endpointType = "public"
	}

	rsConClient, err := d.client.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, nil, err
	}

	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}

	instanceData, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
	if err != nil || instanceData == nil {
		return nil, nil, flex.FmtErrorf("[ERROR] Error retrieving resource instance: %s with resp code: %s", err, resp)
	}

	extensions := instanceData.Extensions
	kpAPI.URL, err = KmsEndpointURL(kpAPI, endpointType, extensions)
	if err != nil {
		return nil, nil, err
	}

	kpAPI.Config.InstanceID = instanceID
	return kpAPI, instanceData.CRN, nil
}

type KMSKeyPoliciesDataSourceModel struct {
	InstanceID   types.String `tfsdk:"instance_id"`
	EndpointType types.String `tfsdk:"endpoint_type"`
	KeyID        types.String `tfsdk:"key_id"`
	Alias        types.String `tfsdk:"alias"`
	Policies     types.List   `tfsdk:"policies"`
	Id           types.String `tfsdk:"id"`
}
