// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.99.1-daeb6e46-20250131-173156
 */

package catalogmanagement

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func ResourceIBMCmAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCmAccountCreate,
		ReadContext:   resourceIBMCmAccountRead,
		UpdateContext: resourceIBMCmAccountUpdate,
		DeleteContext: resourceIBMCmAccountDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloudant revision.",
			},
			"hide_ibm_cloud_catalog": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Hide the public catalog in this account.",
			},
			"account_filters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Filters for account and catalog filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_all": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "-> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.",
						},
						"category_filters": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Optional:    true,
							Description: "Filter against offering categories with dynamic keys.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"category_name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Name of this category",
									},
									"include": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Optional:    true,
										Description: "Whether to include the category in the catalog filter.",
									},
									"filter": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Optional:    true,
										Description: "Filter terms related to the category.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Optional:    true,
													Description: "List of filter terms for the category.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"id_filters": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Offering filter terms.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"exclude": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "Offering filter terms.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"filter_terms": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"region_filter": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Region filter string.",
			},
			"terraform_engines": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: "List of terraform engines configured for this account.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "User provided name for the specified engine.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The terraform engine type. The only one supported at the moment is terraform-enterprise.",
						},
						"public_endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The public endpoint for the engine instance.",
						},
						"private_endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The private endpoint for the engine instance.",
						},
						"api_token": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Sensitive:        true,
							Description:      "The api key used to access the engine instance.",
							DiffSuppressFunc: flex.ApplyOnce,
						},
						"da_creation": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "The settings that determines how deployable architectures are auto-created from workspaces in the terraform engine.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Determines whether deployable architectures are auto-created from workspaces in the engine.",
									},
									"default_private_catalog_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Default private catalog to create the deployable architectures in.",
									},
									"polling_info": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										MaxItems:    1,
										Description: "Determines which workspace scope to query to auto-create deployable architectures from.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"scopes": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Description: "List of scopes to auto-create deployable architectures from workspaces in the engine.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "Identifier for the specified type in the scope.",
															},
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "Scope to auto-create deployable architectures from. The supported scopes today are workspace, org, and project.",
															},
														},
													},
												},
												"last_polling_status": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Last polling status of the engine scope.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"code": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Status code of the last polling attempt.",
															},
															"message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Status message from the last polling attempt.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceIBMCmAccountCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{}

	account, _, err := catalogManagementClient.GetCatalogAccountWithContext(context, getCatalogAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetCatalogAccountWithContext failed: %s", err.Error()), "ibm_cm_account", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*account.ID)

	// Call the Update function to ensure that the resource is in sync with the configuration
	return resourceIBMCmAccountUpdate(context, d, meta)
}

func resourceIBMCmAccountRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{}

	account, response, err := catalogManagementClient.GetCatalogAccountWithContext(context, getCatalogAccountOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetCatalogAccountWithContext failed: %s", err.Error()), "ibm_cm_account", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(account.Rev) {
		if err = d.Set("rev", account.Rev); err != nil {
			err = fmt.Errorf("error setting rev: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "set-rev").GetDiag()
		}
	}
	if !core.IsNil(account.HideIBMCloudCatalog) {
		if err = d.Set("hide_ibm_cloud_catalog", account.HideIBMCloudCatalog); err != nil {
			err = fmt.Errorf("error setting hide_ibm_cloud_catalog: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "set-hide_ibm_cloud_catalog").GetDiag()
		}
	}
	if !core.IsNil(account.AccountFilters) {
		accountFiltersMap, err := ResourceIBMCmAccountFiltersToMap(account.AccountFilters)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "account_filters-to-map").GetDiag()
		}
		if err = d.Set("account_filters", []map[string]interface{}{accountFiltersMap}); err != nil {
			err = fmt.Errorf("error setting account_filters: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "set-account_filters").GetDiag()
		}
	}
	if !core.IsNil(account.RegionFilter) {
		if err = d.Set("region_filter", account.RegionFilter); err != nil {
			err = fmt.Errorf("error setting region_filter: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "set-region_filter").GetDiag()
		}
	}
	if !core.IsNil(account.TerraformEngines) {
		terraformEngines := []map[string]interface{}{}
		for _, terraformEnginesItem := range account.TerraformEngines {
			terraformEnginesItemMap, err := ResourceIBMCmAccountTerraformEnginesToMap(&terraformEnginesItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "terraform_engines-to-map").GetDiag()
			}
			terraformEngines = append(terraformEngines, terraformEnginesItemMap)
		}
		if err = d.Set("terraform_engines", terraformEngines); err != nil {
			err = fmt.Errorf("Error setting terraform_engines: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_cm_account", "read", "set-terraform_engines").GetDiag()
		}
	}

	return nil
}

func resourceIBMCmAccountUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_cm_version", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getCatalogAccountOptions := &catalogmanagementv1.GetCatalogAccountOptions{}

	account, response, err := catalogManagementClient.GetCatalogAccountWithContext(context, getCatalogAccountOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetCatalogAccountWithContext failed: %s", err.Error()), "ibm_cm_account", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateCatalogAccountOptions := &catalogmanagementv1.UpdateCatalogAccountOptions{}
	updateCatalogAccountOptions.SetID(*account.ID)
	updateCatalogAccountOptions.SetRev(*account.Rev)

	if d.HasChange("hide_IBM_cloud_catalog") {
		if v, ok := d.GetOk("hide_IBM_cloud_catalog"); ok {
			updateCatalogAccountOptions.SetHideIBMCloudCatalog(v.(bool))
		}
	} else if account.HideIBMCloudCatalog != nil {
		updateCatalogAccountOptions.SetHideIBMCloudCatalog(*account.HideIBMCloudCatalog)
	}

	if d.HasChange("region_filter") {
		updateCatalogAccountOptions.SetRegionFilter(d.Get("region_filter").(string))
	} else if account.RegionFilter != nil {
		updateCatalogAccountOptions.SetRegionFilter(*account.RegionFilter)
	}

	if d.HasChange("account_filters") {
		if v, ok := d.GetOk("account_filters"); ok {
			accountFilters, err := resourceIBMCmAccountFiltersMapToFilters(v.([]interface{})[0].(map[string]interface{}))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_cm_account", "update")
				log.Printf("[DEBUG]\\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			updateCatalogAccountOptions.SetAccountFilters(accountFilters)
		}
	} else if account.AccountFilters != nil {
		updateCatalogAccountOptions.SetAccountFilters(account.AccountFilters)
	}

	if d.HasChange("terraform_engines") {
		if v, ok := d.GetOk("terraform_engines"); ok {
			terraformEngines, err := resourceIBMCmAccountTerraformEnginesMapToStruct(v.([]interface{}))
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_cm_account", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			updateCatalogAccountOptions.SetTerraformEngines(terraformEngines)
		}
	} else if account.TerraformEngines != nil {
		updateCatalogAccountOptions.SetTerraformEngines(account.TerraformEngines)
	}

	_, response, err = catalogManagementClient.UpdateCatalogAccountWithContext(context, updateCatalogAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateCatalogAccountWithContext failed %s\n%s", err, response), "ibm_cm_object", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMCmAccountRead(context, d, meta)
}

func resourceIBMCmAccountDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIBMCmAccountFiltersToMap(model *catalogmanagementv1.Filters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IncludeAll != nil {
		modelMap["include_all"] = *model.IncludeAll
	}
	if model.CategoryFilters != nil {
		var categoryFiltersList []map[string]interface{}
		for k, category := range model.CategoryFilters {
			categoryFilterMap, err := ResourceIBMCmAccountCategoryFilterToMap(k, &category)
			if err != nil {
				return modelMap, err
			}
			categoryFiltersList = append(categoryFiltersList, categoryFilterMap)
		}
		modelMap["category_filters"] = categoryFiltersList
	}
	if model.IDFilters != nil {
		idFiltersMap, err := ResourceIBMCmAccountIDFilterToMap(model.IDFilters)
		if err != nil {
			return modelMap, err
		}
		modelMap["id_filters"] = []map[string]interface{}{idFiltersMap}
	}
	return modelMap, nil
}

func ResourceIBMCmAccountCategoryFilterToMap(key string, model *catalogmanagementv1.CategoryFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if key != "" {
		modelMap["category_name"] = key
	}
	if model.Include != nil {
		modelMap["include"] = *model.Include
	}
	if model.Filter != nil {
		filterMap, err := ResourceIBMCmAccountFilterTermsToMap(model.Filter)
		if err != nil {
			return modelMap, err
		}
		modelMap["filter"] = []map[string]interface{}{filterMap}
	}
	return modelMap, nil
}

func ResourceIBMCmAccountFilterTermsToMap(model *catalogmanagementv1.FilterTerms) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FilterTerms != nil {
		modelMap["filter_terms"] = model.FilterTerms
	}
	return modelMap, nil
}

func ResourceIBMCmAccountIDFilterToMap(model *catalogmanagementv1.IDFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Include != nil {
		includeMap, err := ResourceIBMCmAccountFilterTermsToMap(model.Include)
		if err != nil {
			return modelMap, err
		}
		modelMap["include"] = []map[string]interface{}{includeMap}
	}
	if model.Exclude != nil {
		excludeMap, err := ResourceIBMCmAccountFilterTermsToMap(model.Exclude)
		if err != nil {
			return modelMap, err
		}
		modelMap["exclude"] = []map[string]interface{}{excludeMap}
	}
	return modelMap, nil
}

func ResourceIBMCmAccountTerraformEnginesToMap(model *catalogmanagementv1.TerraformEngines) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.PublicEndpoint != nil {
		modelMap["public_endpoint"] = *model.PublicEndpoint
	}
	if model.PrivateEndpoint != nil {
		modelMap["private_endpoint"] = *model.PrivateEndpoint
	}
	if model.APIToken != nil {
		modelMap["api_token"] = *model.APIToken
	}
	if model.DaCreation != nil {
		daCreationMap, err := ResourceIBMCmAccountDaCreationToMap(model.DaCreation)
		if err != nil {
			return modelMap, err
		}
		modelMap["da_creation"] = []map[string]interface{}{daCreationMap}
	}
	return modelMap, nil
}

func ResourceIBMCmAccountDaCreationToMap(model *catalogmanagementv1.DaCreation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.DefaultPrivateCatalogID != nil {
		modelMap["default_private_catalog_id"] = *model.DefaultPrivateCatalogID
	}
	if model.PollingInfo != nil {
		pollingInfoMap, err := ResourceIBMCmAccountPollingInfoToMap(model.PollingInfo)
		if err != nil {
			return modelMap, err
		}
		modelMap["polling_info"] = []map[string]interface{}{pollingInfoMap}
	}
	return modelMap, nil
}

func ResourceIBMCmAccountPollingInfoToMap(model *catalogmanagementv1.PollingInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Scopes != nil {
		scopes := []map[string]interface{}{}
		for _, scopesItem := range model.Scopes {
			scopesItemMap, err := ResourceIBMCmAccountTerraformEngineScopeToMap(&scopesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			scopes = append(scopes, scopesItemMap)
		}
		modelMap["scopes"] = scopes
	}
	if model.LastPollingStatus != nil {
		lastPollingStatusMap, err := ResourceIBMCmAccountPollingInfoLastPollingStatusToMap(model.LastPollingStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_polling_status"] = []map[string]interface{}{lastPollingStatusMap}
	}
	return modelMap, nil
}

func ResourceIBMCmAccountTerraformEngineScopeToMap(model *catalogmanagementv1.TerraformEngineScope) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	return modelMap, nil
}

func ResourceIBMCmAccountPollingInfoLastPollingStatusToMap(model *catalogmanagementv1.PollingInfoLastPollingStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Code != nil {
		modelMap["code"] = flex.IntValue(model.Code)
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}

func resourceIBMCmAccountFiltersMapToFilters(modelMap map[string]interface{}) (*catalogmanagementv1.Filters, error) {
	model := &catalogmanagementv1.Filters{}
	if modelMap["include_all"] != nil {
		model.IncludeAll = core.BoolPtr(modelMap["include_all"].(bool))
	}
	if modelMap["id_filters"] != nil && len(modelMap["id_filters"].([]interface{})) > 0 {
		var IDFiltersModel *catalogmanagementv1.IDFilter
		var err error
		if modelMap["id_filters"].([]interface{})[0] != nil {
			IDFiltersModel, err = resourceIBMCmCatalogMapToIDFilter(modelMap["id_filters"].([]interface{})[0].(map[string]interface{}))
			if err != nil {
				return model, err
			}
		}
		model.IDFilters = IDFiltersModel
	}
	if modelMap["category_filters"] != nil {
		categoryFiltersList := modelMap["category_filters"].([]interface{})
		categoryFilters := make(map[string]catalogmanagementv1.CategoryFilter) // Initialize the map for category filters

		for _, item := range categoryFiltersList {
			categoryFilterMap := item.(map[string]interface{})
			categoryName := categoryFilterMap["category_name"].(string) // Extract category_name as the map key

			// Convert the category filter to the appropriate struct
			categoryFilter, err := resourceIBMCmAccountMapToCategoryFilter(categoryFilterMap)
			if err != nil {
				return model, err
			}

			// Add the category filter to the map using category_name as the key
			categoryFilters[categoryName] = *categoryFilter
		}

		// Assign the map to the model
		model.CategoryFilters = categoryFilters
	}
	return model, nil
}

func resourceIBMCmAccountMapToCategoryFilter(modelMap map[string]interface{}) (*catalogmanagementv1.CategoryFilter, error) {
	model := &catalogmanagementv1.CategoryFilter{}
	if modelMap["include"] != nil {
		model.Include = core.BoolPtr(modelMap["include"].(bool))
	}
	if modelMap["filter"] != nil && len(modelMap["filter"].([]interface{})) > 0 {
		FilterModel, err := resourceIBMCmCatalogMapToFilterTerms(modelMap["filter"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Filter = FilterModel
	}
	return model, nil
}

func resourceIBMCmAccountTerraformEnginesMapToStruct(enginesList []interface{}) ([]catalogmanagementv1.TerraformEngines, error) {
	var engines []catalogmanagementv1.TerraformEngines

	for _, item := range enginesList {
		engineMap := item.(map[string]interface{})
		engine := catalogmanagementv1.TerraformEngines{}

		if engineMap["name"] != nil {
			engine.Name = core.StringPtr(engineMap["name"].(string))
		}
		if engineMap["type"] != nil {
			engine.Type = core.StringPtr(engineMap["type"].(string))
		}
		if engineMap["public_endpoint"] != nil {
			engine.PublicEndpoint = core.StringPtr(engineMap["public_endpoint"].(string))
		}
		if engineMap["private_endpoint"] != nil {
			engine.PrivateEndpoint = core.StringPtr(engineMap["private_endpoint"].(string))
		}
		if engineMap["api_token"] != nil {
			engine.APIToken = core.StringPtr(engineMap["api_token"].(string))
		}

		if engineMap["da_creation"] != nil && len(engineMap["da_creation"].([]interface{})) > 0 {
			daCreationMap := engineMap["da_creation"].([]interface{})[0].(map[string]interface{})
			daCreation, err := resourceIBMCmTerraformEngineMapToDaCreation(daCreationMap)
			if err != nil {
				return nil, err
			}
			engine.DaCreation = daCreation
		}

		engines = append(engines, engine)
	}

	return engines, nil
}

func resourceIBMCmTerraformEngineMapToDaCreation(modelMap map[string]interface{}) (*catalogmanagementv1.DaCreation, error) {
	model := &catalogmanagementv1.DaCreation{}

	if modelMap["enabled"] != nil {
		model.Enabled = core.BoolPtr(modelMap["enabled"].(bool))
	}
	if modelMap["default_private_catalog_id"] != nil {
		model.DefaultPrivateCatalogID = core.StringPtr(modelMap["default_private_catalog_id"].(string))
	}
	if modelMap["polling_info"] != nil && len(modelMap["polling_info"].([]interface{})) > 0 {
		pollingInfoMap := modelMap["polling_info"].([]interface{})[0].(map[string]interface{})
		pollingInfo, err := resourceIBMCmTerraformEngineMapToPollingInfo(pollingInfoMap)
		if err != nil {
			return model, err
		}
		model.PollingInfo = pollingInfo
	}

	return model, nil
}

func resourceIBMCmTerraformEngineMapToPollingInfo(modelMap map[string]interface{}) (*catalogmanagementv1.PollingInfo, error) {
	model := &catalogmanagementv1.PollingInfo{}

	if modelMap["scopes"] != nil {
		scopeList := modelMap["scopes"].([]interface{})
		var scopes []catalogmanagementv1.TerraformEngineScope

		for _, s := range scopeList {
			scopeMap := s.(map[string]interface{})
			scope := catalogmanagementv1.TerraformEngineScope{}

			if scopeMap["name"] != nil {
				scope.Name = core.StringPtr(scopeMap["name"].(string))
			}
			if scopeMap["type"] != nil {
				scope.Type = core.StringPtr(scopeMap["type"].(string))
			}

			scopes = append(scopes, scope)
		}
		model.Scopes = scopes
	}

	return model, nil
}
