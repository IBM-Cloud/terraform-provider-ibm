// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func DataSourceIBMSchematicsWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMSchematicsWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the workspace.  To find the workspace ID, use the `GET /v1/workspaces` API.",
			},
			"catalog_ref": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dry_run": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Dry run.",
						},
						"owning_account": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Owning account ID of the catalog.",
						},
						"item_icon_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to the icon of the software template in the IBM Cloud catalog.",
						},
						"item_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.",
						},
						"item_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the software that you chose to install from the IBM Cloud catalog.",
						},
						"item_readme_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to the readme file of the software template in the IBM Cloud catalog.",
						},
						"item_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to the software template in the IBM Cloud catalog.",
						},
						"launch_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to the dashboard to access your software.",
						},
						"offering_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the software template that you chose to install from the IBM Cloud catalog.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the workspace was created.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user ID that created the workspace.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The workspace CRN.",
			},
			"dependencies": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workspace dependencies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parents": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of workspace parents CRN identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"children": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of workspace children CRN identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the workspace.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the workspace.",
			},
			"last_health_check_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the last health check was performed by Schematics.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IBM Cloud location where your workspace was provisioned.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the workspace.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group the workspace was provisioned in.",
			},
			"runtime_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the provisioning engine, state file, and runtime logs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_cmd": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The command that was used to apply the Terraform template or IBM Cloud catalog software template.",
						},
						"engine_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The provisioning engine that was used to apply the Terraform template or IBM Cloud catalog software template.",
						},
						"engine_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the provisioning engine that was used.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID that was assigned to your Terraform template or IBM Cloud catalog software template.",
						},
						"log_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to access the logs that were created during the creation, update, or deletion of your IBM Cloud resources.",
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
							Description: "The URL where the Terraform statefile (`terraform.tfstate`) is stored. You can use the statefile to find an overview of IBM Cloud resources that were created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion jobs.",
						},
					},
				},
			},
			"shared_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the Target used by the templates originating from IBM Cloud catalog offerings. This information is not relevant when you create a workspace from your own Terraform template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target cluster name.",
						},
						"entitlement_keys": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The entitlement key that you want to use to install IBM Cloud entitled software.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.",
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the workspace.   **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.   **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.   **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.   **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.   **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply jobs to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.   **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.   **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.   **Stopped**: The Schematics plan, apply, or destroy job was cancelled manually.   **Template Error**: The Schematics template contains errors and cannot be processed.",
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of tags that are associated with the workspace.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"template_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the Terraform or IBM Cloud software template that you want to use.",
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
										Description: "Environment variable is hidden.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Environment variable name.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Environment variable is secure.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value for environment variable.",
									},
								},
							},
						},
						"folder": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The subfolder in your GitHub or GitLab repository where your Terraform template is stored. If your template is stored in the root directory, `.` is returned.",
						},
						"compact": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, to use the files from the specified folder & subfolder in your GitHub or GitLab repository and ignore the other folders in the repository.",
						},
						"has_githubtoken": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Has github token.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID that was assigned to your Terraform template or IBM Cloud catalog software template.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Terraform version that was used to run your Terraform code.",
						},
						"uninstall_script_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Uninstall script name.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `\"autoscaling: enabled: true minReplicas: 2\"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.",
						},
						"values_metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of input variables that are associated with the workspace.",
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"values_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API endpoint to access the input variables that you defined for your template.",
						},
						"variablestore": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Information about the input variables that your template uses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of your input variable.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the variable.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If set to `true`, the value of your input variable is protected and not returned in your API response.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "`Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html).<br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`,`object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Enter the value as a string for the primitive types such as `bool`, `number`, `string`, and `HCL` format for the complex variables, as you provide in a `.tfvars` file. **You need to enter escaped string of `HCL` format for the complex variable value**. For more information, about how to declare variables in a terraform configuration file and provide value to schematics, see [Providing values for the declared variables](https://cloud.ibm.com/docs/schematics?topic=schematics-create-tf-config#declare-variable).",
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
				Description: "Workspace template reference.",
			},
			"template_repo": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the Template repository used by the workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository branch.",
						},
						"full_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full repository URL.",
						},
						"has_uploadedgitrepotar": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Has uploaded Git repository tar.",
						},
						"release": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository release.",
						},
						"repo_sha_value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository SHA value.",
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository URL. For more information, about using `.netrc` in `env_values`, see [Usage of private module template](https://cloud.ibm.com/docs/schematics?topic=schematics-download-modules-pvt-git).",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source URL.",
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Terraform version that was used to run your Terraform code.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the workspace was last updated.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user ID that updated the workspace.",
			},
			"cart_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The associate cart order ID.",
			},
			"last_action_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the last Action performed on workspace.",
			},
			"last_activity_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of last Activity performed.",
			},
			"last_job": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Last job details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of last job.",
						},
						"job_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the last job.",
						},
						"job_status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the last job.",
						},
					},
				},
			},
			"workspace_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Response that indicate the status of the workspace as either frozen or locked.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"frozen": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, the workspace is frozen and changes to the workspace are disabled.",
						},
						"frozen_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the workspace was frozen.",
						},
						"frozen_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user ID that froze the workspace.",
						},
						"locked": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, the workspace is locked and disabled for changes.",
						},
						"locked_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user ID that initiated a resource-related job, such as applying or destroying resources, that locked the workspace.",
						},
						"locked_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp when the workspace was locked.",
						},
					},
				},
			},
			"workspace_status_msg": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the last job that ran against the workspace. -.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The success or error code that was returned for the last plan, apply, or destroy job that ran against your workspace.",
						},
						"status_msg": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The success or error message that was returned for the last plan, apply, or destroy job that ran against your workspace.",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMSchematicsWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

	getWorkspaceOptions.SetWID(d.Get("workspace_id").(string))

	workspaceResponse, response, err := schematicsClient.GetWorkspaceWithContext(context, getWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetWorkspaceWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getWorkspaceOptions.WID))

	catalogRef := []map[string]interface{}{}
	if workspaceResponse.CatalogRef != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceCatalogRefToMap(workspaceResponse.CatalogRef)
		if err != nil {
			return diag.FromErr(err)
		}
		catalogRef = append(catalogRef, modelMap)
	}
	if err = d.Set("catalog_ref", catalogRef); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_ref %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(workspaceResponse.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("created_by", workspaceResponse.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
	}

	if err = d.Set("crn", workspaceResponse.Crn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	dependencies := []map[string]interface{}{}
	if workspaceResponse.Dependencies != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceDependenciesToMap(workspaceResponse.Dependencies)
		if err != nil {
			return diag.FromErr(err)
		}
		dependencies = append(dependencies, modelMap)
	}
	if err = d.Set("dependencies", dependencies); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting dependencies %s", err))
	}

	if err = d.Set("description", workspaceResponse.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if err = d.Set("id", workspaceResponse.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}

	if err = d.Set("last_health_check_at", flex.DateTimeToString(workspaceResponse.LastHealthCheckAt)); err != nil {
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

	runtimeData := []map[string]interface{}{}
	if workspaceResponse.RuntimeData != nil {
		for _, modelItem := range workspaceResponse.RuntimeData {
			modelMap, err := DataSourceIBMSchematicsWorkspaceTemplateRunTimeDataResponseToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			runtimeData = append(runtimeData, modelMap)
		}
	}
	if err = d.Set("runtime_data", runtimeData); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting runtime_data %s", err))
	}

	sharedData := []map[string]interface{}{}
	if workspaceResponse.SharedData != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceSharedTargetDataResponseToMap(workspaceResponse.SharedData)
		if err != nil {
			return diag.FromErr(err)
		}
		sharedData = append(sharedData, modelMap)
	}
	if err = d.Set("shared_data", sharedData); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting shared_data %s", err))
	}

	if err = d.Set("status", workspaceResponse.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}

	templateData := []map[string]interface{}{}
	if workspaceResponse.TemplateData != nil {
		for _, modelItem := range workspaceResponse.TemplateData {
			modelMap, err := DataSourceIBMSchematicsWorkspaceTemplateSourceDataResponseToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			templateData = append(templateData, modelMap)
		}
	}
	if err = d.Set("template_data", templateData); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_data %s", err))
	}

	if err = d.Set("template_ref", workspaceResponse.TemplateRef); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_ref: %s", err))
	}

	templateRepo := []map[string]interface{}{}
	if workspaceResponse.TemplateRepo != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceTemplateRepoResponseToMap(workspaceResponse.TemplateRepo)
		if err != nil {
			return diag.FromErr(err)
		}
		templateRepo = append(templateRepo, modelMap)
	}
	if err = d.Set("template_repo", templateRepo); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_repo %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(workspaceResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	if err = d.Set("updated_by", workspaceResponse.UpdatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_by: %s", err))
	}

	if err = d.Set("cart_id", workspaceResponse.CartID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cart_id: %s", err))
	}

	if err = d.Set("last_action_name", workspaceResponse.LastActionName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_action_name: %s", err))
	}

	if err = d.Set("last_activity_id", workspaceResponse.LastActivityID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_activity_id: %s", err))
	}

	lastJob := []map[string]interface{}{}
	if workspaceResponse.LastJob != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceLastJobToMap(workspaceResponse.LastJob)
		if err != nil {
			return diag.FromErr(err)
		}
		lastJob = append(lastJob, modelMap)
	}
	if err = d.Set("last_job", lastJob); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_job %s", err))
	}

	workspaceStatus := []map[string]interface{}{}
	if workspaceResponse.WorkspaceStatus != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceWorkspaceStatusResponseToMap(workspaceResponse.WorkspaceStatus)
		if err != nil {
			return diag.FromErr(err)
		}
		workspaceStatus = append(workspaceStatus, modelMap)
	}
	if err = d.Set("workspace_status", workspaceStatus); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting workspace_status %s", err))
	}

	workspaceStatusMsg := []map[string]interface{}{}
	if workspaceResponse.WorkspaceStatusMsg != nil {
		modelMap, err := DataSourceIBMSchematicsWorkspaceWorkspaceStatusMessageToMap(workspaceResponse.WorkspaceStatusMsg)
		if err != nil {
			return diag.FromErr(err)
		}
		workspaceStatusMsg = append(workspaceStatusMsg, modelMap)
	}
	if err = d.Set("workspace_status_msg", workspaceStatusMsg); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting workspace_status_msg %s", err))
	}

	return nil
}

func DataSourceIBMSchematicsWorkspaceCatalogRefToMap(model *schematicsv1.CatalogRef) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DryRun != nil {
		modelMap["dry_run"] = *model.DryRun
	}
	if model.OwningAccount != nil {
		modelMap["owning_account"] = *model.OwningAccount
	}
	if model.ItemIconURL != nil {
		modelMap["item_icon_url"] = *model.ItemIconURL
	}
	if model.ItemID != nil {
		modelMap["item_id"] = *model.ItemID
	}
	if model.ItemName != nil {
		modelMap["item_name"] = *model.ItemName
	}
	if model.ItemReadmeURL != nil {
		modelMap["item_readme_url"] = *model.ItemReadmeURL
	}
	if model.ItemURL != nil {
		modelMap["item_url"] = *model.ItemURL
	}
	if model.LaunchURL != nil {
		modelMap["launch_url"] = *model.LaunchURL
	}
	if model.OfferingVersion != nil {
		modelMap["offering_version"] = *model.OfferingVersion
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceDependenciesToMap(model *schematicsv1.Dependencies) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Parents != nil {
		modelMap["parents"] = model.Parents
	}
	if model.Children != nil {
		modelMap["children"] = model.Children
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceTemplateRunTimeDataResponseToMap(model *schematicsv1.TemplateRunTimeDataResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EngineCmd != nil {
		modelMap["engine_cmd"] = *model.EngineCmd
	}
	if model.EngineName != nil {
		modelMap["engine_name"] = *model.EngineName
	}
	if model.EngineVersion != nil {
		modelMap["engine_version"] = *model.EngineVersion
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.LogStoreURL != nil {
		modelMap["log_store_url"] = *model.LogStoreURL
	}
	if model.OutputValues != nil {
	}
	if model.Resources != nil {
	}
	if model.StateStoreURL != nil {
		modelMap["state_store_url"] = *model.StateStoreURL
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceSharedTargetDataResponseToMap(model *schematicsv1.SharedTargetDataResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = *model.ClusterID
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = *model.ClusterName
	}
	if model.EntitlementKeys != nil {
	}
	if model.Namespace != nil {
		modelMap["namespace"] = *model.Namespace
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.ResourceGroupID != nil {
		modelMap["resource_group_id"] = *model.ResourceGroupID
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceTemplateSourceDataResponseToMap(model *schematicsv1.TemplateSourceDataResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnvValues != nil {
		envValues := []map[string]interface{}{}
		for _, envValuesItem := range model.EnvValues {
			envValuesItemMap, err := DataSourceIBMSchematicsWorkspaceEnvVariableResponseToMap(&envValuesItem)
			if err != nil {
				return modelMap, err
			}
			envValues = append(envValues, envValuesItemMap)
		}
		modelMap["env_values"] = envValues
	}
	if model.Folder != nil {
		modelMap["folder"] = *model.Folder
	}
	if model.Compact != nil {
		modelMap["compact"] = *model.Compact
	}
	if model.HasGithubtoken != nil {
		modelMap["has_githubtoken"] = *model.HasGithubtoken
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.UninstallScriptName != nil {
		modelMap["uninstall_script_name"] = *model.UninstallScriptName
	}
	if model.Values != nil {
		modelMap["values"] = *model.Values
	}
	if model.ValuesMetadata != nil {
	}
	if model.ValuesURL != nil {
		modelMap["values_url"] = *model.ValuesURL
	}
	if model.Variablestore != nil {
		variablestore := []map[string]interface{}{}
		for _, variablestoreItem := range model.Variablestore {
			variablestoreItemMap, err := DataSourceIBMSchematicsWorkspaceWorkspaceVariableResponseToMap(&variablestoreItem)
			if err != nil {
				return modelMap, err
			}
			variablestore = append(variablestore, variablestoreItemMap)
		}
		modelMap["variablestore"] = variablestore
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceEnvVariableResponseToMap(model *schematicsv1.EnvVariableResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hidden != nil {
		modelMap["hidden"] = *model.Hidden
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Secure != nil {
		modelMap["secure"] = *model.Secure
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceWorkspaceVariableResponseToMap(model *schematicsv1.WorkspaceVariableResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Secure != nil {
		modelMap["secure"] = *model.Secure
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceTemplateRepoResponseToMap(model *schematicsv1.TemplateRepoResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Branch != nil {
		modelMap["branch"] = *model.Branch
	}
	if model.FullURL != nil {
		modelMap["full_url"] = *model.FullURL
	}
	if model.HasUploadedgitrepotar != nil {
		modelMap["has_uploadedgitrepotar"] = *model.HasUploadedgitrepotar
	}
	if model.Release != nil {
		modelMap["release"] = *model.Release
	}
	if model.RepoShaValue != nil {
		modelMap["repo_sha_value"] = *model.RepoShaValue
	}
	if model.RepoURL != nil {
		modelMap["repo_url"] = *model.RepoURL
	}
	if model.URL != nil {
		modelMap["url"] = *model.URL
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceLastJobToMap(model *schematicsv1.LastJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.JobID != nil {
		modelMap["job_id"] = *model.JobID
	}
	if model.JobName != nil {
		modelMap["job_name"] = *model.JobName
	}
	if model.JobStatus != nil {
		modelMap["job_status"] = *model.JobStatus
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceWorkspaceStatusResponseToMap(model *schematicsv1.WorkspaceStatusResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Frozen != nil {
		modelMap["frozen"] = *model.Frozen
	}
	if model.FrozenAt != nil {
		modelMap["frozen_at"] = model.FrozenAt.String()
	}
	if model.FrozenBy != nil {
		modelMap["frozen_by"] = *model.FrozenBy
	}
	if model.Locked != nil {
		modelMap["locked"] = *model.Locked
	}
	if model.LockedBy != nil {
		modelMap["locked_by"] = *model.LockedBy
	}
	if model.LockedTime != nil {
		modelMap["locked_time"] = model.LockedTime.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsWorkspaceWorkspaceStatusMessageToMap(model *schematicsv1.WorkspaceStatusMessage) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMsg != nil {
		modelMap["status_msg"] = *model.StatusMsg
	}
	return modelMap, nil
}
