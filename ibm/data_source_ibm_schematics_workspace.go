/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func dataSourceIBMSchematicsWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSchematicsWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The workspace ID for the workspace that you want to query.  You can run the GET /workspaces call if you need to look up the  workspace IDs in your IBM Cloud account.",
			},
			"applied_shareddata_ids": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of applied shared dataset id.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"catalog_ref": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "CatalogRef -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dry_run": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Dry run.",
						},
						"item_icon_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item icon url.",
						},
						"item_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item id.",
						},
						"item_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item name.",
						},
						"item_readme_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item readme url.",
						},
						"item_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item url.",
						},
						"launch_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item launch url.",
						},
						"offering_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Catalog item offering version.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace created at.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace created by.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace CRN.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace description.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace id.",
			},
			"last_health_check_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last health checked at.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace location.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace name.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace resource group.",
			},
			"runtime_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace runtime data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_cmd": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine command.",
						},
						"engine_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine name.",
						},
						"engine_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Engine version.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template id.",
						},
						"log_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Log store url.",
						},
						"output_values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of Output values.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"resources": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of resources.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"state_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State store URL.",
						},
					},
				},
			},
			"shared_data": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "SharedTargetDataResponse -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target cluster id.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target cluster name.",
						},
						"entitlement_keys": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Entitlement keys.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target namespace.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target region.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target resource group id.",
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace status type.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace tags.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"template_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace template data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"env_values": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of environment values.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Env variable is hidden.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Env variable name.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Env variable is secure.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value for env variable.",
									},
								},
							},
						},
						"folder": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Folder name.",
						},
						"has_githubtoken": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Has github token.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template id.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template tyoe.",
						},
						"uninstall_script_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Uninstall script name.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Values.",
						},
						"values_metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of values metadata.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"values_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Values URL.",
						},
						"variablestore": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VariablesResponse -.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Variable descrption.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Variable name.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Variable is secure.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Variable type.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the Variable.",
									},
								},
							},
						},
					},
				},
			},
			"template_ref": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace template ref.",
			},
			"template_repo": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "TemplateRepoResponse -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Repo branch.",
						},
						"full_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full repo URL.",
						},
						"has_uploadedgitrepotar": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Has uploaded git repo tar.",
						},
						"release": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Repo release.",
						},
						"repo_sha_value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Repo SHA value.",
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Repo URL.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source URL.",
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Workspace type.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace updated at.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace updated by.",
			},
			"workspace_status": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "WorkspaceStatusResponse -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"frozen": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Frozen status.",
						},
						"frozen_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Frozen at.",
						},
						"frozen_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Frozen by.",
						},
						"locked": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Locked status.",
						},
						"locked_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Locked by.",
						},
						"locked_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Locked at.",
						},
					},
				},
			},
			"workspace_status_msg": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "WorkspaceStatusMessage -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status code.",
						},
						"status_msg": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status message.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMSchematicsWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

	getWorkspaceOptions.SetWID(d.Get("workspace_id").(string))

	workspaceResponse, response, err := schematicsClient.GetWorkspaceWithContext(context, getWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*workspaceResponse.ID)
	if err = d.Set("applied_shareddata_ids", workspaceResponse.AppliedShareddataIds); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting applied_shareddata_ids: %s", err))
	}

	if workspaceResponse.CatalogRef != nil {
		err = d.Set("catalog_ref", dataSourceWorkspaceResponseFlattenCatalogRef(*workspaceResponse.CatalogRef))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting catalog_ref %s", err))
		}
	}
	if err = d.Set("created_at", workspaceResponse.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by", workspaceResponse.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("crn", workspaceResponse.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("description", workspaceResponse.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("id", workspaceResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}
	if err = d.Set("last_health_check_at", workspaceResponse.LastHealthCheckAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_health_check_at: %s", err))
	}
	if err = d.Set("location", workspaceResponse.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if err = d.Set("name", workspaceResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("resource_group", workspaceResponse.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}

	if workspaceResponse.RuntimeData != nil {
		err = d.Set("runtime_data", dataSourceWorkspaceResponseFlattenRuntimeData(workspaceResponse.RuntimeData))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting runtime_data %s", err))
		}
	}

	if workspaceResponse.SharedData != nil {
		err = d.Set("shared_data", dataSourceWorkspaceResponseFlattenSharedData(*workspaceResponse.SharedData))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting shared_data %s", err))
		}
	}
	if err = d.Set("status", workspaceResponse.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}
	if err = d.Set("tags", workspaceResponse.Tags); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
	}

	if workspaceResponse.TemplateData != nil {
		err = d.Set("template_data", dataSourceWorkspaceResponseFlattenTemplateData(workspaceResponse.TemplateData))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting template_data %s", err))
		}
	}
	if err = d.Set("template_ref", workspaceResponse.TemplateRef); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_ref: %s", err))
	}

	if workspaceResponse.TemplateRepo != nil {
		err = d.Set("template_repo", dataSourceWorkspaceResponseFlattenTemplateRepo(*workspaceResponse.TemplateRepo))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting template_repo %s", err))
		}
	}
	if err = d.Set("type", workspaceResponse.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("updated_at", workspaceResponse.UpdatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("updated_by", workspaceResponse.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}

	if workspaceResponse.WorkspaceStatus != nil {
		err = d.Set("workspace_status", dataSourceWorkspaceResponseFlattenWorkspaceStatus(*workspaceResponse.WorkspaceStatus))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting workspace_status %s", err))
		}
	}

	if workspaceResponse.WorkspaceStatusMsg != nil {
		err = d.Set("workspace_status_msg", dataSourceWorkspaceResponseFlattenWorkspaceStatusMsg(*workspaceResponse.WorkspaceStatusMsg))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting workspace_status_msg %s", err))
		}
	}

	return nil
}

func dataSourceWorkspaceResponseFlattenCatalogRef(result schematicsv1.CatalogRef) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseCatalogRefToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseCatalogRefToMap(catalogRefItem schematicsv1.CatalogRef) (catalogRefMap map[string]interface{}) {
	catalogRefMap = map[string]interface{}{}

	if catalogRefItem.DryRun != nil {
		catalogRefMap["dry_run"] = catalogRefItem.DryRun
	}
	if catalogRefItem.ItemIconURL != nil {
		catalogRefMap["item_icon_url"] = catalogRefItem.ItemIconURL
	}
	if catalogRefItem.ItemID != nil {
		catalogRefMap["item_id"] = catalogRefItem.ItemID
	}
	if catalogRefItem.ItemName != nil {
		catalogRefMap["item_name"] = catalogRefItem.ItemName
	}
	if catalogRefItem.ItemReadmeURL != nil {
		catalogRefMap["item_readme_url"] = catalogRefItem.ItemReadmeURL
	}
	if catalogRefItem.ItemURL != nil {
		catalogRefMap["item_url"] = catalogRefItem.ItemURL
	}
	if catalogRefItem.LaunchURL != nil {
		catalogRefMap["launch_url"] = catalogRefItem.LaunchURL
	}
	if catalogRefItem.OfferingVersion != nil {
		catalogRefMap["offering_version"] = catalogRefItem.OfferingVersion
	}

	return catalogRefMap
}


func dataSourceWorkspaceResponseFlattenRuntimeData(result []schematicsv1.TemplateRunTimeDataResponse) (runtimeData []map[string]interface{}) {
	for _, runtimeDataItem := range result {
		runtimeData = append(runtimeData, dataSourceWorkspaceResponseRuntimeDataToMap(runtimeDataItem))
	}

	return runtimeData
}

func dataSourceWorkspaceResponseRuntimeDataToMap(runtimeDataItem schematicsv1.TemplateRunTimeDataResponse) (runtimeDataMap map[string]interface{}) {
	runtimeDataMap = map[string]interface{}{}

	if runtimeDataItem.EngineCmd != nil {
		runtimeDataMap["engine_cmd"] = runtimeDataItem.EngineCmd
	}
	if runtimeDataItem.EngineName != nil {
		runtimeDataMap["engine_name"] = runtimeDataItem.EngineName
	}
	if runtimeDataItem.EngineVersion != nil {
		runtimeDataMap["engine_version"] = runtimeDataItem.EngineVersion
	}
	if runtimeDataItem.ID != nil {
		runtimeDataMap["id"] = runtimeDataItem.ID
	}
	if runtimeDataItem.LogStoreURL != nil {
		runtimeDataMap["log_store_url"] = runtimeDataItem.LogStoreURL
	}
	if runtimeDataItem.OutputValues != nil {
		runtimeDataMap["output_values"] = runtimeDataItem.OutputValues
	}
	if runtimeDataItem.Resources != nil {
		runtimeDataMap["resources"] = runtimeDataItem.Resources
	}
	if runtimeDataItem.StateStoreURL != nil {
		runtimeDataMap["state_store_url"] = runtimeDataItem.StateStoreURL
	}

	return runtimeDataMap
}


func dataSourceWorkspaceResponseFlattenSharedData(result schematicsv1.SharedTargetDataResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseSharedDataToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseSharedDataToMap(sharedDataItem schematicsv1.SharedTargetDataResponse) (sharedDataMap map[string]interface{}) {
	sharedDataMap = map[string]interface{}{}

	if sharedDataItem.ClusterID != nil {
		sharedDataMap["cluster_id"] = sharedDataItem.ClusterID
	}
	if sharedDataItem.ClusterName != nil {
		sharedDataMap["cluster_name"] = sharedDataItem.ClusterName
	}
	if sharedDataItem.EntitlementKeys != nil {
		sharedDataMap["entitlement_keys"] = sharedDataItem.EntitlementKeys
	}
	if sharedDataItem.Namespace != nil {
		sharedDataMap["namespace"] = sharedDataItem.Namespace
	}
	if sharedDataItem.Region != nil {
		sharedDataMap["region"] = sharedDataItem.Region
	}
	if sharedDataItem.ResourceGroupID != nil {
		sharedDataMap["resource_group_id"] = sharedDataItem.ResourceGroupID
	}

	return sharedDataMap
}


func dataSourceWorkspaceResponseFlattenTemplateData(result []schematicsv1.TemplateSourceDataResponse) (templateData []map[string]interface{}) {
	for _, templateDataItem := range result {
		templateData = append(templateData, dataSourceWorkspaceResponseTemplateDataToMap(templateDataItem))
	}

	return templateData
}

func dataSourceWorkspaceResponseTemplateDataToMap(templateDataItem schematicsv1.TemplateSourceDataResponse) (templateDataMap map[string]interface{}) {
	templateDataMap = map[string]interface{}{}

	
	if templateDataItem.EnvValues != nil {
		envValuesList := []map[string]interface{}{}
		for _, envValuesItem := range templateDataItem.EnvValues {
			envValuesList = append(envValuesList, dataSourceWorkspaceResponseTemplateDataEnvValuesToMap(envValuesItem))
		}
		templateDataMap["env_values"] = envValuesList
	}
	if templateDataItem.Folder != nil {
		templateDataMap["folder"] = templateDataItem.Folder
	}
	if templateDataItem.HasGithubtoken != nil {
		templateDataMap["has_githubtoken"] = templateDataItem.HasGithubtoken
	}
	if templateDataItem.ID != nil {
		templateDataMap["id"] = templateDataItem.ID
	}
	if templateDataItem.Type != nil {
		templateDataMap["type"] = templateDataItem.Type
	}
	if templateDataItem.UninstallScriptName != nil {
		templateDataMap["uninstall_script_name"] = templateDataItem.UninstallScriptName
	}
	if templateDataItem.Values != nil {
		templateDataMap["values"] = templateDataItem.Values
	}
	if templateDataItem.ValuesMetadata != nil {
		templateDataMap["values_metadata"] = templateDataItem.ValuesMetadata
	}
	if templateDataItem.ValuesURL != nil {
		templateDataMap["values_url"] = templateDataItem.ValuesURL
	}
	if templateDataItem.Variablestore != nil {
		variablestoreList := []map[string]interface{}{}
		for _, variablestoreItem := range templateDataItem.Variablestore {
			variablestoreList = append(variablestoreList, dataSourceWorkspaceResponseTemplateDataVariablestoreToMap(variablestoreItem))
		}
		templateDataMap["variablestore"] = variablestoreList
	}

	return templateDataMap
}

func dataSourceWorkspaceResponseTemplateDataEnvValuesToMap(envValuesItem schematicsv1.EnvVariableResponse) (envValuesMap map[string]interface{}) {
	envValuesMap = map[string]interface{}{}

	if envValuesItem.Hidden != nil {
		envValuesMap["hidden"] = envValuesItem.Hidden
	}
	if envValuesItem.Name != nil {
		envValuesMap["name"] = envValuesItem.Name
	}
	if envValuesItem.Secure != nil {
		envValuesMap["secure"] = envValuesItem.Secure
	}
	if envValuesItem.Value != nil {
		envValuesMap["value"] = envValuesItem.Value
	}

	return envValuesMap
}


func dataSourceWorkspaceResponseTemplateDataVariablestoreToMap(variablestoreItem schematicsv1.WorkspaceVariableResponse) (variablestoreMap map[string]interface{}) {
	variablestoreMap = map[string]interface{}{}

	if variablestoreItem.Description != nil {
		variablestoreMap["description"] = variablestoreItem.Description
	}
	if variablestoreItem.Name != nil {
		variablestoreMap["name"] = variablestoreItem.Name
	}
	if variablestoreItem.Secure != nil {
		variablestoreMap["secure"] = variablestoreItem.Secure
	}
	if variablestoreItem.Type != nil {
		variablestoreMap["type"] = variablestoreItem.Type
	}
	if variablestoreItem.Value != nil {
		variablestoreMap["value"] = variablestoreItem.Value
	}

	return variablestoreMap
}



func dataSourceWorkspaceResponseFlattenTemplateRepo(result schematicsv1.TemplateRepoResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseTemplateRepoToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseTemplateRepoToMap(templateRepoItem schematicsv1.TemplateRepoResponse) (templateRepoMap map[string]interface{}) {
	templateRepoMap = map[string]interface{}{}

	if templateRepoItem.Branch != nil {
		templateRepoMap["branch"] = templateRepoItem.Branch
	}
	if templateRepoItem.FullURL != nil {
		templateRepoMap["full_url"] = templateRepoItem.FullURL
	}
	if templateRepoItem.HasUploadedgitrepotar != nil {
		templateRepoMap["has_uploadedgitrepotar"] = templateRepoItem.HasUploadedgitrepotar
	}
	if templateRepoItem.Release != nil {
		templateRepoMap["release"] = templateRepoItem.Release
	}
	if templateRepoItem.RepoShaValue != nil {
		templateRepoMap["repo_sha_value"] = templateRepoItem.RepoShaValue
	}
	if templateRepoItem.RepoURL != nil {
		templateRepoMap["repo_url"] = templateRepoItem.RepoURL
	}
	if templateRepoItem.URL != nil {
		templateRepoMap["url"] = templateRepoItem.URL
	}

	return templateRepoMap
}


func dataSourceWorkspaceResponseFlattenWorkspaceStatus(result schematicsv1.WorkspaceStatusResponse) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseWorkspaceStatusToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseWorkspaceStatusToMap(workspaceStatusItem schematicsv1.WorkspaceStatusResponse) (workspaceStatusMap map[string]interface{}) {
	workspaceStatusMap = map[string]interface{}{}

	if workspaceStatusItem.Frozen != nil {
		workspaceStatusMap["frozen"] = workspaceStatusItem.Frozen
	}
	if workspaceStatusItem.FrozenAt != nil {
		workspaceStatusMap["frozen_at"] = workspaceStatusItem.FrozenAt.String()
	}
	if workspaceStatusItem.FrozenBy != nil {
		workspaceStatusMap["frozen_by"] = workspaceStatusItem.FrozenBy
	}
	if workspaceStatusItem.Locked != nil {
		workspaceStatusMap["locked"] = workspaceStatusItem.Locked
	}
	if workspaceStatusItem.LockedBy != nil {
		workspaceStatusMap["locked_by"] = workspaceStatusItem.LockedBy
	}
	if workspaceStatusItem.LockedTime != nil {
		workspaceStatusMap["locked_time"] = workspaceStatusItem.LockedTime.String()
	}

	return workspaceStatusMap
}


func dataSourceWorkspaceResponseFlattenWorkspaceStatusMsg(result schematicsv1.WorkspaceStatusMessage) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceWorkspaceResponseWorkspaceStatusMsgToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceWorkspaceResponseWorkspaceStatusMsgToMap(workspaceStatusMsgItem schematicsv1.WorkspaceStatusMessage) (workspaceStatusMsgMap map[string]interface{}) {
	workspaceStatusMsgMap = map[string]interface{}{}

	if workspaceStatusMsgItem.StatusCode != nil {
		workspaceStatusMsgMap["status_code"] = workspaceStatusMsgItem.StatusCode
	}
	if workspaceStatusMsgItem.StatusMsg != nil {
		workspaceStatusMsgMap["status_msg"] = workspaceStatusMsgItem.StatusMsg
	}

	return workspaceStatusMsgMap
}

