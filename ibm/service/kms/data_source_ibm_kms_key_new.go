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

var _ datasource.DataSource = &KMSKeyDataSource{}

type KMSKeyDataSource struct {
	client interface{}
}

func NewKMSKeyDataSource(client interface{}) datasource.DataSource {
	return &KMSKeyDataSource{
		client: client,
	}
}

func (d *KMSKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kms_key_new"
}

func (d *KMSKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "KMS key data source",

		Attributes: map[string]schema.Attribute{
			"instance_id": schema.StringAttribute{
				MarkdownDescription: "Key protect or hpcs instance GUID",
				Required:            true,
			},
			"limit": schema.Int64Attribute{
				MarkdownDescription: "Limit till the keys to be fetched",
				Optional:            true,
			},
			"key_id": schema.StringAttribute{
				MarkdownDescription: "Key ID",
				Optional:            true,
			},
			"key_name": schema.StringAttribute{
				MarkdownDescription: "The name of the key to be fetched",
				Optional:            true,
			},
			"alias": schema.StringAttribute{
				MarkdownDescription: "The alias associated with the key",
				Optional:            true,
			},
			"endpoint_type": schema.StringAttribute{
				MarkdownDescription: "public or private",
				Optional:            true,
				Computed:            true,
			},
			"keys": schema.ListAttribute{
				MarkdownDescription: "List of keys",
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"aliases":      types.ListType{ElemType: types.StringType},
						"name":         types.StringType,
						"key_ring_id":  types.StringType,
						"crn":          types.StringType,
						"id":           types.StringType,
						"description":  types.StringType,
						"standard_key": types.BoolType,
						"policies": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"rotation": types.ListType{
										ElemType: types.ObjectType{
											AttrTypes: map[string]attr.Type{
												"id":               types.StringType,
												"created_by":       types.StringType,
												"creation_date":    types.StringType,
												"crn":              types.StringType,
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
												"created_by":       types.StringType,
												"creation_date":    types.StringType,
												"crn":              types.StringType,
												"updated_by":       types.StringType,
												"last_update_date": types.StringType,
												"enabled":          types.BoolType,
											},
										},
									},
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

func (d *KMSKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data KMSKeyDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	instanceID := getInstanceIDFromCRN(data.InstanceID.ValueString())
	api, _, err := d.populateKPClientForDataSource(data, instanceID)
	if err != nil {
		resp.Diagnostics.AddError("Error creating KP client", err.Error())
		return
	}

	var totalKeys []kp.Key

	if !data.KeyName.IsNull() && data.KeyName.ValueString() != "" {
		limit := 0
		if !data.Limit.IsNull() {
			limit = int(data.Limit.ValueInt64())
		}
		offset := 0
		pageSize := 200

		if limit == 0 {
			keys, err := api.GetKeys(ctx, 0, offset)
			if err != nil {
				resp.Diagnostics.AddError("Error getting keys", err.Error())
				return
			}
			totalKeys = append(totalKeys, keys.Keys...)
		} else {
			for {
				if offset < limit {
					if (limit - offset) < pageSize {
						keys, err := api.GetKeys(ctx, (limit - offset), offset)
						if err != nil {
							resp.Diagnostics.AddError("Error getting keys", err.Error())
							return
						}
						totalKeys = append(totalKeys, keys.Keys...)
						break
					} else {
						keys, err := api.GetKeys(ctx, pageSize, offset)
						if err != nil {
							resp.Diagnostics.AddError("Error getting keys", err.Error())
							return
						}
						numOfKeysFetched := keys.Metadata.NumberOfKeys
						totalKeys = append(totalKeys, keys.Keys...)
						if numOfKeysFetched < pageSize || offset+pageSize == limit {
							break
						}
						offset = offset + pageSize
					}
				}
			}
		}

		if len(totalKeys) == 0 {
			resp.Diagnostics.AddError("No keys found", "No keys in instance "+instanceID)
			return
		}

		var matchKeys []kp.Key
		keyName := data.KeyName.ValueString()
		for _, keyData := range totalKeys {
			if keyData.Name == keyName {
				matchKeys = append(matchKeys, keyData)
			}
		}

		if len(matchKeys) == 0 {
			resp.Diagnostics.AddError("No keys found", "No keys with name "+keyName+" in instance "+instanceID)
			return
		}

		keyList, err := d.buildKeyList(ctx, api, matchKeys)
		if err != nil {
			resp.Diagnostics.AddError("Error building key list", err.Error())
			return
		}

		data.Keys = keyList
		data.Id = basetypes.NewStringValue(instanceID)

	} else if !data.KeyID.IsNull() && data.KeyID.ValueString() != "" {
		key, err := api.GetKey(ctx, data.KeyID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error getting key", err.Error())
			return
		}

		keyList, err := d.buildKeyList(ctx, api, []kp.Key{*key})
		if err != nil {
			resp.Diagnostics.AddError("Error building key list", err.Error())
			return
		}

		data.Keys = keyList
		data.Id = basetypes.NewStringValue(instanceID)

	} else if !data.Alias.IsNull() && data.Alias.ValueString() != "" {
		aliasName := data.Alias.ValueString()
		key, err := api.GetKey(ctx, aliasName)
		if err != nil {
			resp.Diagnostics.AddError("Error getting key by alias", err.Error())
			return
		}

		keyList, err := d.buildKeyList(ctx, api, []kp.Key{*key})
		if err != nil {
			resp.Diagnostics.AddError("Error building key list", err.Error())
			return
		}

		data.Keys = keyList
		data.Id = basetypes.NewStringValue(instanceID)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *KMSKeyDataSource) buildKeyList(ctx context.Context, api *kp.Client, keys []kp.Key) (types.List, error) {
	keyList := make([]attr.Value, 0)

	for _, key := range keys {
		aliases := make([]attr.Value, 0)
		for _, alias := range key.Aliases {
			aliases = append(aliases, basetypes.NewStringValue(alias))
		}

		aliasList, _ := types.ListValue(types.StringType, aliases)

		policies, err := api.GetPolicies(ctx, key.ID)
		if err != nil {
			return types.List{}, err
		}

		policiesList := types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"rotation": types.ListType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"id":               types.StringType,
							"created_by":       types.StringType,
							"creation_date":    types.StringType,
							"crn":              types.StringType,
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
							"created_by":       types.StringType,
							"creation_date":    types.StringType,
							"crn":              types.StringType,
							"updated_by":       types.StringType,
							"last_update_date": types.StringType,
							"enabled":          types.BoolType,
						},
					},
				},
			},
		})

		if len(policies) > 0 {
			// Convert policies using flex helper
			policiesFlat := flex.FlattenKeyPolicies(policies)
			if policiesFlat != nil && len(policiesFlat) > 0 {
				policyList := make([]attr.Value, 0)
				for _, policyMap := range policiesFlat {

					rotationList := types.ListNull(types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"id":               types.StringType,
							"created_by":       types.StringType,
							"creation_date":    types.StringType,
							"crn":              types.StringType,
							"updated_by":       types.StringType,
							"last_update_date": types.StringType,
							"interval_month":   types.Int64Type,
							"enabled":          types.BoolType,
						},
					})

					dualAuthList := types.ListNull(types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"id":               types.StringType,
							"created_by":       types.StringType,
							"creation_date":    types.StringType,
							"crn":              types.StringType,
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
									"created_by":       types.StringType,
									"creation_date":    types.StringType,
									"crn":              types.StringType,
									"updated_by":       types.StringType,
									"last_update_date": types.StringType,
									"interval_month":   types.Int64Type,
									"enabled":          types.BoolType,
								},
								map[string]attr.Value{
									"id":               basetypes.NewStringValue(rm["id"].(string)),
									"created_by":       basetypes.NewStringValue(rm["created_by"].(string)),
									"creation_date":    basetypes.NewStringValue(rm["creation_date"].(string)),
									"crn":              basetypes.NewStringValue(rm["crn"].(string)),
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
								"created_by":       types.StringType,
								"creation_date":    types.StringType,
								"crn":              types.StringType,
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
									"created_by":       types.StringType,
									"creation_date":    types.StringType,
									"crn":              types.StringType,
									"updated_by":       types.StringType,
									"last_update_date": types.StringType,
									"enabled":          types.BoolType,
								},
								map[string]attr.Value{
									"id":               basetypes.NewStringValue(dam["id"].(string)),
									"created_by":       basetypes.NewStringValue(dam["created_by"].(string)),
									"creation_date":    basetypes.NewStringValue(dam["creation_date"].(string)),
									"crn":              basetypes.NewStringValue(dam["crn"].(string)),
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
								"created_by":       types.StringType,
								"creation_date":    types.StringType,
								"crn":              types.StringType,
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
										"created_by":       types.StringType,
										"creation_date":    types.StringType,
										"crn":              types.StringType,
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
										"created_by":       types.StringType,
										"creation_date":    types.StringType,
										"crn":              types.StringType,
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

				policiesList, _ = types.ListValue(types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"rotation": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"id":               types.StringType,
									"created_by":       types.StringType,
									"creation_date":    types.StringType,
									"crn":              types.StringType,
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
									"created_by":       types.StringType,
									"creation_date":    types.StringType,
									"crn":              types.StringType,
									"updated_by":       types.StringType,
									"last_update_date": types.StringType,
									"enabled":          types.BoolType,
								},
							},
						},
					},
				}, policyList)
			}
		}

		if len(policies) == 0 {
			log.Printf("No Policy Configurations read\n")
		}

		keyObj, _ := types.ObjectValue(
			map[string]attr.Type{
				"aliases":      types.ListType{ElemType: types.StringType},
				"name":         types.StringType,
				"key_ring_id":  types.StringType,
				"crn":          types.StringType,
				"id":           types.StringType,
				"description":  types.StringType,
				"standard_key": types.BoolType,
				"policies": types.ListType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"rotation": types.ListType{
								ElemType: types.ObjectType{
									AttrTypes: map[string]attr.Type{
										"id":               types.StringType,
										"created_by":       types.StringType,
										"creation_date":    types.StringType,
										"crn":              types.StringType,
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
										"created_by":       types.StringType,
										"creation_date":    types.StringType,
										"crn":              types.StringType,
										"updated_by":       types.StringType,
										"last_update_date": types.StringType,
										"enabled":          types.BoolType,
									},
								},
							},
						},
					},
				},
			},
			map[string]attr.Value{
				"id":           basetypes.NewStringValue(key.ID),
				"name":         basetypes.NewStringValue(key.Name),
				"crn":          basetypes.NewStringValue(key.CRN),
				"standard_key": basetypes.NewBoolValue(key.Extractable),
				"description":  basetypes.NewStringValue(key.Description),
				"aliases":      aliasList,
				"key_ring_id":  basetypes.NewStringValue(key.KeyRingID),
				"policies":     policiesList,
			},
		)

		keyList = append(keyList, keyObj)
	}

	listValue, diags := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"aliases":      types.ListType{ElemType: types.StringType},
				"name":         types.StringType,
				"key_ring_id":  types.StringType,
				"crn":          types.StringType,
				"id":           types.StringType,
				"description":  types.StringType,
				"standard_key": types.BoolType,
				"policies": types.ListType{
					ElemType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"rotation": types.ListType{
								ElemType: types.ObjectType{
									AttrTypes: map[string]attr.Type{
										"id":               types.StringType,
										"created_by":       types.StringType,
										"creation_date":    types.StringType,
										"crn":              types.StringType,
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
										"created_by":       types.StringType,
										"creation_date":    types.StringType,
										"crn":              types.StringType,
										"updated_by":       types.StringType,
										"last_update_date": types.StringType,
										"enabled":          types.BoolType,
									},
								},
							},
						},
					},
				},
			},
		},
		keyList,
	)
	if diags.HasError() {
		return types.List{}, flex.FmtErrorf("Error creating list value: %v", diags)
	}
	return listValue, nil
}

func (d *KMSKeyDataSource) populateKPClientForDataSource(model KMSKeyDataSourceModel, instanceID string) (*kp.Client, *string, error) {
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

type KMSKeyDataSourceModel struct {
	InstanceID   types.String `tfsdk:"instance_id"`
	Limit        types.Int64  `tfsdk:"limit"`
	KeyID        types.String `tfsdk:"key_id"`
	KeyName      types.String `tfsdk:"key_name"`
	Alias        types.String `tfsdk:"alias"`
	EndpointType types.String `tfsdk:"endpoint_type"`
	Keys         types.List   `tfsdk:"keys"`
	Id           types.String `tfsdk:"id"`
}
