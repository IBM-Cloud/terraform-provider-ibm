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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/go-openapi/strfmt"
)

func resourceIBMSchematicsWorkspace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMSchematicsWorkspaceCreate,
		ReadContext:   resourceIBMSchematicsWorkspaceRead,
		UpdateContext: resourceIBMSchematicsWorkspaceUpdate,
		DeleteContext: resourceIBMSchematicsWorkspaceDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"applied_shareddata_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of applied shared dataset id.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"catalog_ref": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "CatalogRef -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dry_run": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Dry run.",
						},
						"item_icon_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item icon url.",
						},
						"item_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item id.",
						},
						"item_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item name.",
						},
						"item_readme_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item readme url.",
						},
						"item_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item url.",
						},
						"launch_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item launch url.",
						},
						"offering_version": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Catalog item offering version.",
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Workspace description.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Workspace location.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Workspace name.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Workspace resource group.",
			},
			"shared_data": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "SharedTargetData -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_created_on": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster created on.",
						},
						"cluster_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster id.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster name.",
						},
						"cluster_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster type.",
						},
						"entitlement_keys": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Entitlement keys.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Target namespace.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Target region.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Target resource group id.",
						},
						"worker_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Cluster worker count.",
						},
						"worker_machine_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster worker type.",
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Workspace tags.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"template_data": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "TemplateData -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"env_values": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "EnvVariableRequest ..",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"folder": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Folder name.",
						},
						"init_state_file": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Init state file.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Template type.",
						},
						"uninstall_script_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Uninstall script name.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Value.",
						},
						"values_metadata": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "List of values metadata.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"variablestore": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "VariablesRequest -.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Variable description.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Variable name.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Variable is secure.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Variable type.",
									},
									"use_default": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Variable uses default value; and is not over-ridden.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
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
				Optional:    true,
				Description: "Workspace template ref.",
			},
			"template_repo": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "TemplateRepoRequest -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repo branch.",
						},
						"release": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repo release.",
						},
						"repo_sha_value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repo SHA value.",
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Repo URL.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Source URL.",
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of Workspace type.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"workspace_status": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "WorkspaceStatusRequest -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"frozen": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Frozen status.",
						},
						"frozen_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Frozen at.",
						},
						"frozen_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Frozen by.",
						},
						"locked": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Locked status.",
						},
						"locked_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Locked by.",
						},
						"locked_time": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Locked at.",
						},
					},
				},
			},
			"x_github_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The github token associated with the GIT. Required for cloning of repo.",
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
			"last_health_check_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last health checked at.",
			},
			"runtime_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace runtime data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_cmd": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Engine command.",
						},
						"engine_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Engine name.",
						},
						"engine_version": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Engine version.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Template id.",
						},
						"log_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Log store url.",
						},
						"output_values": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Output values.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"resources": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of resources.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"state_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "State store URL.",
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Workspace status type.",
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
			"workspace_status_msg": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "WorkspaceStatusMessage -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_code": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Status code.",
						},
						"status_msg": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Status message.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMSchematicsWorkspaceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createWorkspaceOptions := &schematicsv1.CreateWorkspaceOptions{}

	if _, ok := d.GetOk("applied_shareddata_ids"); ok {
		createWorkspaceOptions.SetAppliedShareddataIds(expandStringList(d.Get("applied_shareddata_ids").([]interface{})))
	}
	if _, ok := d.GetOk("catalog_ref"); ok {
		catalogRef := resourceIBMSchematicsWorkspaceMapToCatalogRef(d.Get("catalog_ref.0").(map[string]interface{}))
		createWorkspaceOptions.SetCatalogRef(&catalogRef)
	}
	if _, ok := d.GetOk("description"); ok {
		createWorkspaceOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		createWorkspaceOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("name"); ok {
		createWorkspaceOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		createWorkspaceOptions.SetResourceGroup(d.Get("resource_group").(string))
	}
	if _, ok := d.GetOk("shared_data"); ok {
		sharedData := resourceIBMSchematicsWorkspaceMapToSharedTargetData(d.Get("shared_data.0").(map[string]interface{}))
		createWorkspaceOptions.SetSharedData(&sharedData)
	}
	if _, ok := d.GetOk("tags"); ok {
		createWorkspaceOptions.SetTags(expandStringList(d.Get("tags").([]interface{})))
	}
	if _, ok := d.GetOk("template_data"); ok {
		var templateData []schematicsv1.TemplateSourceDataRequest
		for _, e := range d.Get("template_data").([]interface{}) {
			value := e.(map[string]interface{})
			templateDataItem := resourceIBMSchematicsWorkspaceMapToTemplateSourceDataRequest(value)
			templateData = append(templateData, templateDataItem)
		}
		createWorkspaceOptions.SetTemplateData(templateData)
	}
	if _, ok := d.GetOk("template_ref"); ok {
		createWorkspaceOptions.SetTemplateRef(d.Get("template_ref").(string))
	}
	if _, ok := d.GetOk("template_repo"); ok {
		templateRepo := resourceIBMSchematicsWorkspaceMapToTemplateRepoRequest(d.Get("template_repo.0").(map[string]interface{}))
		createWorkspaceOptions.SetTemplateRepo(&templateRepo)
	}
	if _, ok := d.GetOk("type"); ok {
		createWorkspaceOptions.SetType(expandStringList(d.Get("type").([]interface{})))
	}
	if _, ok := d.GetOk("workspace_status"); ok {
		workspaceStatus := resourceIBMSchematicsWorkspaceMapToWorkspaceStatusRequest(d.Get("workspace_status.0").(map[string]interface{}))
		createWorkspaceOptions.SetWorkspaceStatus(&workspaceStatus)
	}
	if _, ok := d.GetOk("x_github_token"); ok {
		createWorkspaceOptions.SetXGithubToken(d.Get("x_github_token").(string))
	}

	workspaceResponse, response, err := schematicsClient.CreateWorkspaceWithContext(context, createWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*workspaceResponse.ID)

	return resourceIBMSchematicsWorkspaceRead(context, d, meta)
}

func resourceIBMSchematicsWorkspaceMapToCatalogRef(catalogRefMap map[string]interface{}) schematicsv1.CatalogRef {
	catalogRef := schematicsv1.CatalogRef{}

	if catalogRefMap["dry_run"] != nil {
		catalogRef.DryRun = core.BoolPtr(catalogRefMap["dry_run"].(bool))
	}
	if catalogRefMap["item_icon_url"] != nil {
		catalogRef.ItemIconURL = core.StringPtr(catalogRefMap["item_icon_url"].(string))
	}
	if catalogRefMap["item_id"] != nil {
		catalogRef.ItemID = core.StringPtr(catalogRefMap["item_id"].(string))
	}
	if catalogRefMap["item_name"] != nil {
		catalogRef.ItemName = core.StringPtr(catalogRefMap["item_name"].(string))
	}
	if catalogRefMap["item_readme_url"] != nil {
		catalogRef.ItemReadmeURL = core.StringPtr(catalogRefMap["item_readme_url"].(string))
	}
	if catalogRefMap["item_url"] != nil {
		catalogRef.ItemURL = core.StringPtr(catalogRefMap["item_url"].(string))
	}
	if catalogRefMap["launch_url"] != nil {
		catalogRef.LaunchURL = core.StringPtr(catalogRefMap["launch_url"].(string))
	}
	if catalogRefMap["offering_version"] != nil {
		catalogRef.OfferingVersion = core.StringPtr(catalogRefMap["offering_version"].(string))
	}

	return catalogRef
}

func resourceIBMSchematicsWorkspaceMapToSharedTargetData(sharedTargetDataMap map[string]interface{}) schematicsv1.SharedTargetData {
	sharedTargetData := schematicsv1.SharedTargetData{}

	if sharedTargetDataMap["cluster_created_on"] != nil {
		sharedTargetData.ClusterCreatedOn = core.StringPtr(sharedTargetDataMap["cluster_created_on"].(string))
	}
	if sharedTargetDataMap["cluster_id"] != nil {
		sharedTargetData.ClusterID = core.StringPtr(sharedTargetDataMap["cluster_id"].(string))
	}
	if sharedTargetDataMap["cluster_name"] != nil {
		sharedTargetData.ClusterName = core.StringPtr(sharedTargetDataMap["cluster_name"].(string))
	}
	if sharedTargetDataMap["cluster_type"] != nil {
		sharedTargetData.ClusterType = core.StringPtr(sharedTargetDataMap["cluster_type"].(string))
	}
	if sharedTargetDataMap["entitlement_keys"] != nil {
		entitlementKeys := []interface{}{}
		for _, entitlementKeysItem := range sharedTargetDataMap["entitlement_keys"].([]interface{}) {
			entitlementKeys = append(entitlementKeys, entitlementKeysItem.(interface{}))
		}
		sharedTargetData.EntitlementKeys = entitlementKeys
	}
	if sharedTargetDataMap["namespace"] != nil {
		sharedTargetData.Namespace = core.StringPtr(sharedTargetDataMap["namespace"].(string))
	}
	if sharedTargetDataMap["region"] != nil {
		sharedTargetData.Region = core.StringPtr(sharedTargetDataMap["region"].(string))
	}
	if sharedTargetDataMap["resource_group_id"] != nil {
		sharedTargetData.ResourceGroupID = core.StringPtr(sharedTargetDataMap["resource_group_id"].(string))
	}
	if sharedTargetDataMap["worker_count"] != nil {
		sharedTargetData.WorkerCount = core.Int64Ptr(int64(sharedTargetDataMap["worker_count"].(int)))
	}
	if sharedTargetDataMap["worker_machine_type"] != nil {
		sharedTargetData.WorkerMachineType = core.StringPtr(sharedTargetDataMap["worker_machine_type"].(string))
	}

	return sharedTargetData
}

func resourceIBMSchematicsWorkspaceMapToTemplateSourceDataRequest(templateSourceDataRequestMap map[string]interface{}) schematicsv1.TemplateSourceDataRequest {
	templateSourceDataRequest := schematicsv1.TemplateSourceDataRequest{}

	if templateSourceDataRequestMap["env_values"] != nil {
		envValues := []interface{}{}
		for _, envValuesItem := range templateSourceDataRequestMap["env_values"].([]interface{}) {
			envValues = append(envValues, envValuesItem.(interface{}))
		}
		templateSourceDataRequest.EnvValues = envValues
	}
	if templateSourceDataRequestMap["folder"] != nil {
		templateSourceDataRequest.Folder = core.StringPtr(templateSourceDataRequestMap["folder"].(string))
	}
	if templateSourceDataRequestMap["init_state_file"] != nil {
		templateSourceDataRequest.InitStateFile = core.StringPtr(templateSourceDataRequestMap["init_state_file"].(string))
	}
	if templateSourceDataRequestMap["type"] != nil {
		templateSourceDataRequest.Type = core.StringPtr(templateSourceDataRequestMap["type"].(string))
	}
	if templateSourceDataRequestMap["uninstall_script_name"] != nil {
		templateSourceDataRequest.UninstallScriptName = core.StringPtr(templateSourceDataRequestMap["uninstall_script_name"].(string))
	}
	if templateSourceDataRequestMap["values"] != nil {
		templateSourceDataRequest.Values = core.StringPtr(templateSourceDataRequestMap["values"].(string))
	}
	if templateSourceDataRequestMap["values_metadata"] != nil {
		valuesMetadata := []interface{}{}
		for _, valuesMetadataItem := range templateSourceDataRequestMap["values_metadata"].([]interface{}) {
			valuesMetadata = append(valuesMetadata, valuesMetadataItem.(interface{}))
		}
		templateSourceDataRequest.ValuesMetadata = valuesMetadata
	}
	if templateSourceDataRequestMap["variablestore"] != nil {
		variablestore := []schematicsv1.WorkspaceVariableRequest{}
		for _, variablestoreItem := range templateSourceDataRequestMap["variablestore"].([]interface{}) {
			variablestoreItemModel := resourceIBMSchematicsWorkspaceMapToWorkspaceVariableRequest(variablestoreItem.(map[string]interface{}))
			variablestore = append(variablestore, variablestoreItemModel)
		}
		templateSourceDataRequest.Variablestore = variablestore
	}

	return templateSourceDataRequest
}

func resourceIBMSchematicsWorkspaceMapToWorkspaceVariableRequest(workspaceVariableRequestMap map[string]interface{}) schematicsv1.WorkspaceVariableRequest {
	workspaceVariableRequest := schematicsv1.WorkspaceVariableRequest{}

	if workspaceVariableRequestMap["description"] != nil {
		workspaceVariableRequest.Description = core.StringPtr(workspaceVariableRequestMap["description"].(string))
	}
	if workspaceVariableRequestMap["name"] != nil {
		workspaceVariableRequest.Name = core.StringPtr(workspaceVariableRequestMap["name"].(string))
	}
	if workspaceVariableRequestMap["secure"] != nil {
		workspaceVariableRequest.Secure = core.BoolPtr(workspaceVariableRequestMap["secure"].(bool))
	}
	if workspaceVariableRequestMap["type"] != nil {
		workspaceVariableRequest.Type = core.StringPtr(workspaceVariableRequestMap["type"].(string))
	}
	if workspaceVariableRequestMap["use_default"] != nil {
		workspaceVariableRequest.UseDefault = core.BoolPtr(workspaceVariableRequestMap["use_default"].(bool))
	}
	if workspaceVariableRequestMap["value"] != nil {
		workspaceVariableRequest.Value = core.StringPtr(workspaceVariableRequestMap["value"].(string))
	}

	return workspaceVariableRequest
}

func resourceIBMSchematicsWorkspaceMapToTemplateRepoRequest(templateRepoRequestMap map[string]interface{}) schematicsv1.TemplateRepoRequest {
	templateRepoRequest := schematicsv1.TemplateRepoRequest{}

	if templateRepoRequestMap["branch"] != nil {
		templateRepoRequest.Branch = core.StringPtr(templateRepoRequestMap["branch"].(string))
	}
	if templateRepoRequestMap["release"] != nil {
		templateRepoRequest.Release = core.StringPtr(templateRepoRequestMap["release"].(string))
	}
	if templateRepoRequestMap["repo_sha_value"] != nil {
		templateRepoRequest.RepoShaValue = core.StringPtr(templateRepoRequestMap["repo_sha_value"].(string))
	}
	if templateRepoRequestMap["repo_url"] != nil {
		templateRepoRequest.RepoURL = core.StringPtr(templateRepoRequestMap["repo_url"].(string))
	}
	if templateRepoRequestMap["url"] != nil {
		templateRepoRequest.URL = core.StringPtr(templateRepoRequestMap["url"].(string))
	}

	return templateRepoRequest
}

func resourceIBMSchematicsWorkspaceMapToWorkspaceStatusRequest(workspaceStatusRequestMap map[string]interface{}) schematicsv1.WorkspaceStatusRequest {
	workspaceStatusRequest := schematicsv1.WorkspaceStatusRequest{}

	if workspaceStatusRequestMap["frozen"] != nil {
		workspaceStatusRequest.Frozen = core.BoolPtr(workspaceStatusRequestMap["frozen"].(bool))
	}
	if workspaceStatusRequestMap["frozen_at"] != nil {
		frozenAt, err := strfmt.ParseDateTime(workspaceStatusRequestMap["frozen_at"].(string))
		if err != nil {
			workspaceStatusRequest.FrozenAt = &frozenAt
		}
	}
	if workspaceStatusRequestMap["frozen_by"] != nil {
		workspaceStatusRequest.FrozenBy = core.StringPtr(workspaceStatusRequestMap["frozen_by"].(string))
	}
	if workspaceStatusRequestMap["locked"] != nil {
		workspaceStatusRequest.Locked = core.BoolPtr(workspaceStatusRequestMap["locked"].(bool))
	}
	if workspaceStatusRequestMap["locked_by"] != nil {
		workspaceStatusRequest.LockedBy = core.StringPtr(workspaceStatusRequestMap["locked_by"].(string))
	}
	if workspaceStatusRequestMap["locked_time"] != nil {
		lockedTime, err := strfmt.ParseDateTime(workspaceStatusRequestMap["locked_time"].(string))
		if err != nil {
			workspaceStatusRequest.LockedTime = &lockedTime
		}
	}

	return workspaceStatusRequest
}

func resourceIBMSchematicsWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

	getWorkspaceOptions.SetWID(d.Id())

	workspaceResponse, response, err := schematicsClient.GetWorkspaceWithContext(context, getWorkspaceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if workspaceResponse.AppliedShareddataIds != nil {
		if err = d.Set("applied_shareddata_ids", workspaceResponse.AppliedShareddataIds); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting applied_shareddata_ids: %s", err))
		}
	}
	if workspaceResponse.CatalogRef != nil {
		catalogRefMap := resourceIBMSchematicsWorkspaceCatalogRefToMap(*workspaceResponse.CatalogRef)
		if err = d.Set("catalog_ref", []map[string]interface{}{catalogRefMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting catalog_ref: %s", err))
		}
	}
	if err = d.Set("description", workspaceResponse.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
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
	if _, ok := d.GetOk("shared_data"); ok {
		if workspaceResponse.SharedData != nil {
			sharedDataMap := resourceIBMSchematicsWorkspaceSharedTargetDataToMap(*workspaceResponse.SharedData)
			if err = d.Set("shared_data", []map[string]interface{}{sharedDataMap}); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting shared_data: %s", err))
			}
		}
	}
	if workspaceResponse.Tags != nil {
		if err = d.Set("tags", workspaceResponse.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	if workspaceResponse.TemplateData != nil {
		templateData := []map[string]interface{}{}
		for _, templateDataItem := range workspaceResponse.TemplateData {
			templateDataItemMap := resourceIBMSchematicsWorkspaceTemplateSourceDataRequestToMap(templateDataItem)
			templateData = append(templateData, templateDataItemMap)
		}
		if err = d.Set("template_data", templateData); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting template_data: %s", err))
		}
	}
	if err = d.Set("template_ref", workspaceResponse.TemplateRef); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_ref: %s", err))
	}
	if _, ok := d.GetOk("template_repo"); ok {
		if workspaceResponse.TemplateRepo != nil {
			templateRepoMap := resourceIBMSchematicsWorkspaceTemplateRepoRequestToMap(*workspaceResponse.TemplateRepo)
			if err = d.Set("template_repo", []map[string]interface{}{templateRepoMap}); err != nil {
				return diag.FromErr(fmt.Errorf("Error reading template_repo: %s", err))
			}
		}
	}
	if workspaceResponse.Type != nil {
		if err = d.Set("type", workspaceResponse.Type); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
		}
	}
	if workspaceResponse.WorkspaceStatus != nil {
		workspaceStatusMap := resourceIBMSchematicsWorkspaceWorkspaceStatusRequestToMap(*workspaceResponse.WorkspaceStatus)
		if err = d.Set("workspace_status", []map[string]interface{}{workspaceStatusMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting workspace_status: %s", err))
		}
	}
	if err = d.Set("x_github_token", workspaceResponse.XGithubToken); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting x_github_token: %s", err))
	}
	if workspaceResponse.CreatedAt != nil {
		if err = d.Set("created_at", workspaceResponse.CreatedAt.String()); err != nil {
			return diag.FromErr(fmt.Errorf("Error reading created_at: %s", err))
		}
	}
	if err = d.Set("created_by", workspaceResponse.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}
	if err = d.Set("crn", workspaceResponse.Crn); err != nil {
		return diag.FromErr(fmt.Errorf("Error reading crn: %s", err))
	}
	if workspaceResponse.LastHealthCheckAt != nil {
		if err = d.Set("last_health_check_at", workspaceResponse.LastHealthCheckAt.String()); err != nil {
			return diag.FromErr(fmt.Errorf("Error reading last_health_check_at: %s", err))
		}
	}
	if workspaceResponse.RuntimeData != nil {
		runtimeData := []map[string]interface{}{}
		for _, runtimeDataItem := range workspaceResponse.RuntimeData {
			runtimeDataItemMap := resourceIBMSchematicsWorkspaceTemplateRunTimeDataResponseToMap(runtimeDataItem)
			runtimeData = append(runtimeData, runtimeDataItemMap)
		}
		if err = d.Set("runtime_data", runtimeData); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting runtime_data: %s", err))
		}
	}
	if err = d.Set("status", workspaceResponse.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}
	if workspaceResponse.UpdatedAt != nil {
		if err = d.Set("updated_at", workspaceResponse.UpdatedAt.String()); err != nil {
			return diag.FromErr(fmt.Errorf("Error reading updated_at: %s", err))
		}
	}
	if err = d.Set("updated_by", workspaceResponse.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}
	if workspaceResponse.WorkspaceStatusMsg != nil {
		workspaceStatusMsgMap := resourceIBMSchematicsWorkspaceWorkspaceStatusMessageToMap(*workspaceResponse.WorkspaceStatusMsg)
		if err = d.Set("workspace_status_msg", []map[string]interface{}{workspaceStatusMsgMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting workspace_status_msg: %s", err))
		}
	}

	return nil
}

func resourceIBMSchematicsWorkspaceCatalogRefToMap(catalogRef schematicsv1.CatalogRef) map[string]interface{} {
	catalogRefMap := map[string]interface{}{}

	catalogRefMap["dry_run"] = catalogRef.DryRun
	catalogRefMap["item_icon_url"] = catalogRef.ItemIconURL
	catalogRefMap["item_id"] = catalogRef.ItemID
	catalogRefMap["item_name"] = catalogRef.ItemName
	catalogRefMap["item_readme_url"] = catalogRef.ItemReadmeURL
	catalogRefMap["item_url"] = catalogRef.ItemURL
	catalogRefMap["launch_url"] = catalogRef.LaunchURL
	catalogRefMap["offering_version"] = catalogRef.OfferingVersion

	return catalogRefMap
}

func resourceIBMSchematicsWorkspaceSharedTargetDataToMap(sharedTargetData schematicsv1.SharedTargetData) map[string]interface{} {
	sharedTargetDataMap := map[string]interface{}{}

	sharedTargetDataMap["cluster_created_on"] = sharedTargetData.ClusterCreatedOn
	sharedTargetDataMap["cluster_id"] = sharedTargetData.ClusterID
	sharedTargetDataMap["cluster_name"] = sharedTargetData.ClusterName
	sharedTargetDataMap["cluster_type"] = sharedTargetData.ClusterType
	if sharedTargetData.EntitlementKeys != nil {
		entitlementKeys := []interface{}{}
		for _, entitlementKeysItem := range sharedTargetData.EntitlementKeys {
			entitlementKeys = append(entitlementKeys, entitlementKeysItem)
		}
		sharedTargetDataMap["entitlement_keys"] = entitlementKeys
	}
	sharedTargetDataMap["namespace"] = sharedTargetData.Namespace
	sharedTargetDataMap["region"] = sharedTargetData.Region
	sharedTargetDataMap["resource_group_id"] = sharedTargetData.ResourceGroupID
	sharedTargetDataMap["worker_count"] = intValue(sharedTargetData.WorkerCount)
	sharedTargetDataMap["worker_machine_type"] = sharedTargetData.WorkerMachineType

	return sharedTargetDataMap
}

func resourceIBMSchematicsWorkspaceTemplateSourceDataRequestToMap(templateSourceDataRequest schematicsv1.TemplateSourceDataRequest) map[string]interface{} {
	templateSourceDataRequestMap := map[string]interface{}{}

	if templateSourceDataRequest.EnvValues != nil {
		envValues := []interface{}{}
		for _, envValuesItem := range templateSourceDataRequest.EnvValues {
			envValues = append(envValues, envValuesItem)
		}
		templateSourceDataRequestMap["env_values"] = envValues
	}
	templateSourceDataRequestMap["folder"] = templateSourceDataRequest.Folder
	templateSourceDataRequestMap["init_state_file"] = templateSourceDataRequest.InitStateFile
	templateSourceDataRequestMap["type"] = templateSourceDataRequest.Type
	templateSourceDataRequestMap["uninstall_script_name"] = templateSourceDataRequest.UninstallScriptName
	templateSourceDataRequestMap["values"] = templateSourceDataRequest.Values
	if templateSourceDataRequest.ValuesMetadata != nil {
		valuesMetadata := []interface{}{}
		for _, valuesMetadataItem := range templateSourceDataRequest.ValuesMetadata {
			valuesMetadata = append(valuesMetadata, valuesMetadataItem)
		}
		templateSourceDataRequestMap["values_metadata"] = valuesMetadata
	}
	if templateSourceDataRequest.Variablestore != nil {
		variablestore := []map[string]interface{}{}
		for _, variablestoreItem := range templateSourceDataRequest.Variablestore {
			variablestoreItemMap := resourceIBMSchematicsWorkspaceWorkspaceVariableRequestToMap(variablestoreItem)
			variablestore = append(variablestore, variablestoreItemMap)
			// TODO: handle Variablestore of type TypeList -- list of non-primitive, not model items
		}
		templateSourceDataRequestMap["variablestore"] = variablestore
	}

	return templateSourceDataRequestMap
}

func resourceIBMSchematicsWorkspaceWorkspaceVariableRequestToMap(workspaceVariableRequest schematicsv1.WorkspaceVariableRequest) map[string]interface{} {
	workspaceVariableRequestMap := map[string]interface{}{}

	workspaceVariableRequestMap["description"] = workspaceVariableRequest.Description
	workspaceVariableRequestMap["name"] = workspaceVariableRequest.Name
	workspaceVariableRequestMap["secure"] = workspaceVariableRequest.Secure
	workspaceVariableRequestMap["type"] = workspaceVariableRequest.Type
	workspaceVariableRequestMap["use_default"] = workspaceVariableRequest.UseDefault
	workspaceVariableRequestMap["value"] = workspaceVariableRequest.Value

	return workspaceVariableRequestMap
}

func resourceIBMSchematicsWorkspaceTemplateRepoRequestToMap(templateRepoRequest schematicsv1.TemplateRepoRequest) map[string]interface{} {
	templateRepoRequestMap := map[string]interface{}{}

	templateRepoRequestMap["branch"] = templateRepoRequest.Branch
	templateRepoRequestMap["release"] = templateRepoRequest.Release
	templateRepoRequestMap["repo_sha_value"] = templateRepoRequest.RepoShaValue
	templateRepoRequestMap["repo_url"] = templateRepoRequest.RepoURL
	templateRepoRequestMap["url"] = templateRepoRequest.URL

	return templateRepoRequestMap
}

func resourceIBMSchematicsWorkspaceWorkspaceStatusRequestToMap(workspaceStatusRequest schematicsv1.WorkspaceStatusRequest) map[string]interface{} {
	workspaceStatusRequestMap := map[string]interface{}{}

	workspaceStatusRequestMap["frozen"] = workspaceStatusRequest.Frozen
	workspaceStatusRequestMap["frozen_at"] = workspaceStatusRequest.FrozenAt.String()
	workspaceStatusRequestMap["frozen_by"] = workspaceStatusRequest.FrozenBy
	workspaceStatusRequestMap["locked"] = workspaceStatusRequest.Locked
	workspaceStatusRequestMap["locked_by"] = workspaceStatusRequest.LockedBy
	workspaceStatusRequestMap["locked_time"] = workspaceStatusRequest.LockedTime.String()

	return workspaceStatusRequestMap
}

func resourceIBMSchematicsWorkspaceTemplateRunTimeDataResponseToMap(templateRunTimeDataResponse schematicsv1.TemplateRunTimeDataResponse) map[string]interface{} {
	templateRunTimeDataResponseMap := map[string]interface{}{}

	templateRunTimeDataResponseMap["engine_cmd"] = templateRunTimeDataResponse.EngineCmd
	templateRunTimeDataResponseMap["engine_name"] = templateRunTimeDataResponse.EngineName
	templateRunTimeDataResponseMap["engine_version"] = templateRunTimeDataResponse.EngineVersion
	templateRunTimeDataResponseMap["id"] = templateRunTimeDataResponse.ID
	templateRunTimeDataResponseMap["log_store_url"] = templateRunTimeDataResponse.LogStoreURL
	if templateRunTimeDataResponse.OutputValues != nil {
		outputValues := []interface{}{}
		for _, outputValuesItem := range templateRunTimeDataResponse.OutputValues {
			outputValues = append(outputValues, outputValuesItem)
		}
		templateRunTimeDataResponseMap["output_values"] = outputValues
	}
	if templateRunTimeDataResponse.Resources != nil {
		resources := []interface{}{}
		for _, resourcesItem := range templateRunTimeDataResponse.Resources {
			resources = append(resources, resourcesItem)
		}
		templateRunTimeDataResponseMap["resources"] = resources
	}
	templateRunTimeDataResponseMap["state_store_url"] = templateRunTimeDataResponse.StateStoreURL

	return templateRunTimeDataResponseMap
}

func resourceIBMSchematicsWorkspaceWorkspaceStatusMessageToMap(workspaceStatusMessage schematicsv1.WorkspaceStatusMessage) map[string]interface{} {
	workspaceStatusMessageMap := map[string]interface{}{}

	workspaceStatusMessageMap["status_code"] = workspaceStatusMessage.StatusCode
	workspaceStatusMessageMap["status_msg"] = workspaceStatusMessage.StatusMsg

	return workspaceStatusMessageMap
}

func resourceIBMSchematicsWorkspaceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateWorkspaceOptions := &schematicsv1.UpdateWorkspaceOptions{}

	updateWorkspaceOptions.SetWID(d.Id())

	hasChange := false

	if d.HasChange("applied_shareddata_ids") {
		// TODO: handle AppliedShareddataIds of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("catalog_ref") {
		catalogRef := resourceIBMSchematicsWorkspaceMapToCatalogRef(d.Get("catalog_ref.0").(map[string]interface{}))
		updateWorkspaceOptions.SetCatalogRef(&catalogRef)
		hasChange = true
	}
	if d.HasChange("description") {
		updateWorkspaceOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("location") {
		updateWorkspaceOptions.SetLocation(d.Get("location").(string))
		hasChange = true
	}
	if d.HasChange("name") {
		updateWorkspaceOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("resource_group") {
		updateWorkspaceOptions.SetResourceGroup(d.Get("resource_group").(string))
		hasChange = true
	}
	if d.HasChange("shared_data") {
		sharedData := resourceIBMSchematicsWorkspaceMapToSharedTargetData(d.Get("shared_data.0").(map[string]interface{}))
		updateWorkspaceOptions.SetSharedData(&sharedData)
		hasChange = true
	}
	if d.HasChange("tags") {
		updateWorkspaceOptions.SetTags(expandStringList(d.Get("tags").([]interface{})))
		hasChange = true
	}
	if d.HasChange("template_data") {
		var templateData []schematicsv1.TemplateSourceDataRequest
		for _, e := range d.Get("template_data").([]interface{}) {
			value := e.(map[string]interface{})
			templateDataItem := resourceIBMSchematicsWorkspaceMapToTemplateSourceDataRequest(value)
			templateData = append(templateData, templateDataItem)
		}
		updateWorkspaceOptions.SetTemplateData(templateData)
		hasChange = true
	}
	if d.HasChange("template_repo") {
		templateRepo := resourceIBMSchematicsWorkspaceMapToTemplateRepoRequest(d.Get("template_repo.0").(map[string]interface{}))
		updateWorkspaceOptions.SetTemplateRepo(&templateRepo)
		hasChange = true
	}
	if d.HasChange("type") {
		updateWorkspaceOptions.SetType(expandStringList(d.Get("type").([]interface{})))
		hasChange = true
	}
	if d.HasChange("workspace_status") {
		workspaceStatus := resourceIBMSchematicsWorkspaceMapToWorkspaceStatusRequest(d.Get("workspace_status.0").(map[string]interface{}))
		updateWorkspaceOptions.SetWorkspaceStatus(&workspaceStatus)
		hasChange = true
	}
	if d.HasChange("x_github_token") {
		updateWorkspaceOptions.SetXGithubToken(d.Get("x_github_token").(string))
		hasChange = true
	}

	if hasChange {
		_, response, err := schematicsClient.UpdateWorkspaceWithContext(context, updateWorkspaceOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateWorkspaceWithContext failed %s\n%s", err, response)
			return diag.FromErr(err)
		}
	}

	return resourceIBMSchematicsWorkspaceRead(context, d, meta)
}

func resourceIBMSchematicsWorkspaceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteWorkspaceOptions := &schematicsv1.DeleteWorkspaceOptions{}

	deleteWorkspaceOptions.SetWID(d.Id())

	_, response, err := schematicsClient.DeleteWorkspaceWithContext(context, deleteWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
