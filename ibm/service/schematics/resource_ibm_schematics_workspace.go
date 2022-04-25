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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

const (
	schematicsWorkspaceName         = "name"
	schematicsWorkspaceDescription  = "description"
	schematicsWorkspaceTemplateType = "template_type"
)

func ResourceIBMSchematicsWorkspace() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSchematicsWorkspaceCreate,
		ReadContext:   ResourceIBMSchematicsWorkspaceRead,
		UpdateContext: ResourceIBMSchematicsWorkspaceUpdate,
		DeleteContext: ResourceIBMSchematicsWorkspaceDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"applied_shareddata_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of applied shared dataset ID.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"catalog_ref": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Information about the software template that you chose from the IBM Cloud catalog. This information is returned for IBM Cloud catalog offerings only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dry_run": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Dry run.",
						},
						"owning_account": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Owning account ID of the catalog.",
						},
						"item_icon_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URL to the icon of the software template in the IBM Cloud catalog.",
						},
						"item_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the software template that you chose to install from the IBM Cloud catalog. This software is provisioned with Schematics.",
						},
						"item_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the software that you chose to install from the IBM Cloud catalog.",
						},
						"item_readme_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URL to the readme file of the software template in the IBM Cloud catalog.",
						},
						"item_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URL to the software template in the IBM Cloud catalog.",
						},
						"launch_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URL to the dashboard to access your software.",
						},
						"offering_version": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The version of the software template that you chose to install from the IBM Cloud catalog.",
						},
					},
				},
			},
			"dependencies": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Workspace dependencies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parents": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of workspace parents CRN identifiers.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"children": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of workspace children CRN identifiers.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the workspace.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The location where you want to create your Schematics workspace and run the Schematics jobs. The location that you enter must match the API endpoint that you use. For example, if you use the Frankfurt API endpoint, you must specify `eu-de` as your location. If you use an API endpoint for a geography and you do not specify a location, Schematics determines the location based on availability.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of your workspace. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. When you create a workspace for your own Terraform template, consider including the microservice component that you set up with your Terraform template and the IBM Cloud environment where you want to deploy your resources in your name.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the resource group where you want to provision the workspace.",
			},
			"shared_data": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Information about the Target used by the templates originating from the  IBM Cloud catalog offerings. This information is not relevant for workspace created using your own Terraform template.",
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
							Description: "The ID of the cluster where you want to provision the resources of all IBM Cloud catalog templates that are included in the catalog offering.",
						},
						"cluster_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cluster name.",
						},
						"cluster_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cluster type.",
						},
						"entitlement_keys": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The entitlement key that you want to use to install IBM Cloud entitled software.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"namespace": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Kubernetes namespace or OpenShift project where the resources of all IBM Cloud catalog templates that are included in the catalog offering are deployed into.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IBM Cloud region that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.",
						},
						"resource_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the resource group that you want to use for the resources of all IBM Cloud catalog templates that are included in the catalog offering.",
						},
						"worker_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The cluster worker count.",
						},
						"worker_machine_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The cluster worker type.",
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of tags that are associated with the workspace.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"template_data": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Input data for the Template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"env_values": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of environment variables that you want to apply during the execution of a bash script or Terraform job. This field must be provided as a list of key-value pairs, for example, **TF_LOG=debug**. Each entry will be a map with one entry where `key is the environment variable name and value is value`. You can define environment variables for IBM Cloud catalog offerings that are provisioned by using a bash script. See [example to use special environment variable](https://cloud.ibm.com/docs/schematics?topic=schematics-set-parallelism#parallelism-example)  that are supported by Schematics.",
							Elem:        &schema.Schema{Type: schema.TypeMap},
						},
						"env_values_metadata": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Environment variables metadata.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Environment variable is hidden.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Environment variable name.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Environment variable is secure.",
									},
								},
							},
						},
						"folder": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The subfolder in your GitHub or GitLab repository where your Terraform template is stored.",
						},
						"compact": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, to use the files from the specified folder & subfolder in your GitHub or GitLab repository and ignore the other folders in the repository. For more information, see [Compact download for Schematics workspace](https://cloud.ibm.com/docs/schematics?topic=schematics-compact-download&interface=ui).",
						},
						"init_state_file": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The content of an existing Terraform statefile that you want to import in to your workspace. To get the content of a Terraform statefile for a specific Terraform template in an existing workspace, run `ibmcloud terraform state pull --id <workspace_id> --template <template_id>`.",
						},
						"injectors": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Array of injectable terraform blocks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tft_git_url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Git repo url hosting terraform template files.",
									},
									"tft_git_token": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Token to access the git repository (Optional).",
									},
									"tft_prefix": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Optional prefix word to append to files (Optional).",
									},
									"injection_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Injection type. Default is 'override'.",
									},
									"tft_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Terraform template name. Maps to folder name in git repo.",
									},
									"tft_parameters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Key name to replace.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Value to replace.",
												},
											},
										},
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Terraform version that you want to use to run your Terraform code. Enter `terraform_v1.1` to use Terraform version 1.1, and `terraform_v1.0` to use Terraform version 1.0. This is a required variable. Make sure that your Terraform config files are compatible with the Terraform version that you select. For more information, refer to [Terraform version](https://cloud.ibm.com/docs/schematics?topic=schematics-workspace-setup&interface=ui#create-workspace_ui).",
						},
						"uninstall_script_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Uninstall script name.",
						},
						"values": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A list of variable values that you want to apply during the Helm chart installation. The list must be provided in JSON format, such as `\"autoscaling: enabled: true minReplicas: 2\"`. The values that you define here override the default Helm chart values. This field is supported only for IBM Cloud catalog offerings that are provisioned by using the Terraform Helm provider.",
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
										Description: "The description of your input variable.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the variable.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If set to `true`, the value of your input variable is protected and not returned in your API response.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "`Terraform v0.11` supports `string`, `list`, `map` data type. For more information, about the syntax, see [Configuring input variables](https://www.terraform.io/docs/configuration-0-11/variables.html).<br> `Terraform v0.12` additionally, supports `bool`, `number` and complex data types such as `list(type)`, `map(type)`,`object({attribute name=type,..})`, `set(type)`, `tuple([type])`. For more information, about the syntax to use the complex data type, see [Configuring variables](https://www.terraform.io/docs/configuration/variables.html#type-constraints).",
									},
									"use_default": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Variable uses default value; and is not over-ridden.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
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
				Optional:    true,
				Description: "Workspace template ref.",
			},
			"template_repo": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Input variables for the Template repoository, while creating a workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"branch": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The repository branch.",
						},
						"release": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The repository release.",
						},
						"repo_sha_value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The repository SHA value.",
						},
						"repo_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The repository URL.",
						},
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The source URL.",
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
							Description: "If set to true, the workspace is frozen and changes to the workspace are disabled.",
						},
						"frozen_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The timestamp when the workspace was frozen.",
						},
						"frozen_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The user ID that froze the workspace.",
						},
						"locked": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If set to true, the workspace is locked and disabled for changes.",
						},
						"locked_by": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The user ID that initiated a resource-related job, such as applying or destroying resources, that locked the workspace.",
						},
						"locked_time": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The timestamp when the workspace was locked.",
						},
					},
				},
			},
			"x_github_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.",
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
			"last_health_check_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the last health check was performed by Schematics.",
			},
			"runtime_data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the provisioning engine, state file, and runtime logs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"engine_cmd": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The command that was used to apply the Terraform template or IBM Cloud catalog software template.",
						},
						"engine_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The provisioning engine that was used to apply the Terraform template or IBM Cloud catalog software template.",
						},
						"engine_version": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The version of the provisioning engine that was used.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID that was assigned to your Terraform template or IBM Cloud catalog software template.",
						},
						"log_store_url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URL to access the logs that were created during the creation, update, or deletion of your IBM Cloud resources.",
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
							Description: "The URL where the Terraform statefile (`terraform.tfstate`) is stored. You can use the statefile to find an overview of IBM Cloud resources that were created by Schematics. Schematics uses the statefile as an inventory list to determine future create, update, or deletion jobs.",
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the workspace.   **Active**: After you successfully ran your infrastructure code by applying your Terraform execution plan, the state of your workspace changes to `Active`.   **Connecting**: Schematics tries to connect to the template in your source repo. If successfully connected, the template is downloaded and metadata, such as input parameters, is extracted. After the template is downloaded, the state of the workspace changes to `Scanning`.   **Draft**: The workspace is created without a reference to a GitHub or GitLab repository.   **Failed**: If errors occur during the execution of your infrastructure code in IBM Cloud Schematics, your workspace status is set to `Failed`.   **Inactive**: The Terraform template was scanned successfully and the workspace creation is complete. You can now start running Schematics plan and apply jobs to provision the IBM Cloud resources that you specified in your template. If you have an `Active` workspace and decide to remove all your resources, your workspace is set to `Inactive` after all your resources are removed.   **In progress**: When you instruct IBM Cloud Schematics to run your infrastructure code by applying your Terraform execution plan, the status of our workspace changes to `In progress`.   **Scanning**: The download of the Terraform template is complete and vulnerability scanning started. If the scan is successful, the workspace state changes to `Inactive`. If errors in your template are found, the state changes to `Template Error`.   **Stopped**: The Schematics plan, apply, or destroy job was cancelled manually.   **Template Error**: The Schematics template contains errors and cannot be processed.",
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
							Optional:    true,
							Description: "ID of last job.",
						},
						"job_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the last job.",
						},
						"job_status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Status of the last job.",
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
							Optional:    true,
							Description: "The success or error code that was returned for the last plan, apply, or destroy job that ran against your workspace.",
						},
						"status_msg": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The success or error message that was returned for the last plan, apply, or destroy job that ran against your workspace.",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMSchematicsWorkspaceValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 schematicsWorkspaceName,
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Regexp:                     `^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
			Required:                   false})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 schematicsWorkspaceDescription,
			ValidateFunctionIdentifier: validate.StringLenBetween,
			Type:                       validate.TypeString,
			MinValueLength:             0,
			MaxValueLength:             2048,
			Optional:                   true})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 schematicsWorkspaceTemplateType,
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Regexp:                     `^terraform_v(?:0\.11|0\.12|0\.13|0\.14|0\.15|1\.0)(?:\.\d+)?$`,
			Default:                    "[]",
			Optional:                   true})

	ibmSchematicsWorkspaceResourceValidator := validate.ResourceValidator{ResourceName: "ibm_schematics_workspace", Schema: validateSchema}
	return &ibmSchematicsWorkspaceResourceValidator
}

func ResourceIBMSchematicsWorkspaceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createWorkspaceOptions := &schematicsv1.CreateWorkspaceOptions{}

	if _, ok := d.GetOk("applied_shareddata_ids"); ok {
		createWorkspaceOptions.SetAppliedShareddataIds(d.Get("applied_shareddata_ids").([]string))
	}
	if _, ok := d.GetOk("catalog_ref"); ok {
		catalogRefModel, err := ResourceIBMSchematicsWorkspaceMapToCatalogRef(d.Get("catalog_ref.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createWorkspaceOptions.SetCatalogRef(catalogRefModel)
	}
	if _, ok := d.GetOk("dependencies"); ok {
		dependenciesModel, err := ResourceIBMSchematicsWorkspaceMapToDependencies(d.Get("dependencies.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createWorkspaceOptions.SetDependencies(dependenciesModel)
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
		sharedDataModel, err := ResourceIBMSchematicsWorkspaceMapToSharedTargetData(d.Get("shared_data.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createWorkspaceOptions.SetSharedData(sharedDataModel)
	}
	if _, ok := d.GetOk("tags"); ok {
		createWorkspaceOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("template_data"); ok {
		var templateData []schematicsv1.TemplateSourceDataRequest
		for _, e := range d.Get("template_data").([]interface{}) {
			value := e.(map[string]interface{})
			templateDataItem, err := ResourceIBMSchematicsWorkspaceMapToTemplateSourceDataRequest(value)
			if err != nil {
				return diag.FromErr(err)
			}
			templateData = append(templateData, *templateDataItem)
		}
		createWorkspaceOptions.SetTemplateData(templateData)
	}
	if _, ok := d.GetOk("template_ref"); ok {
		createWorkspaceOptions.SetTemplateRef(d.Get("template_ref").(string))
	}
	if _, ok := d.GetOk("template_repo"); ok {
		templateRepoModel, err := ResourceIBMSchematicsWorkspaceMapToTemplateRepoRequest(d.Get("template_repo.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createWorkspaceOptions.SetTemplateRepo(templateRepoModel)
	}
	if _, ok := d.GetOk("type"); ok {
		createWorkspaceOptions.SetType(d.Get("type").([]string))
	}
	if _, ok := d.GetOk("workspace_status"); ok {
		workspaceStatusModel, err := ResourceIBMSchematicsWorkspaceMapToWorkspaceStatusRequest(d.Get("workspace_status.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createWorkspaceOptions.SetWorkspaceStatus(workspaceStatusModel)
	}
	if _, ok := d.GetOk("x_github_token"); ok {
		createWorkspaceOptions.SetXGithubToken(d.Get("x_github_token").(string))
	}

	workspaceResponse, response, err := schematicsClient.CreateWorkspaceWithContext(context, createWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateWorkspaceWithContext failed %s\n%s", err, response))
	}

	d.SetId(*workspaceResponse.ID)

	return ResourceIBMSchematicsWorkspaceRead(context, d, meta)
}

func ResourceIBMSchematicsWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
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
		return diag.FromErr(fmt.Errorf("GetWorkspaceWithContext failed %s\n%s", err, response))
	}

	if workspaceResponse.AppliedShareddataIds != nil {
		if err = d.Set("applied_shareddata_ids", workspaceResponse.AppliedShareddataIds); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting applied_shareddata_ids: %s", err))
		}
	}
	if workspaceResponse.CatalogRef != nil {
		catalogRefMap, err := ResourceIBMSchematicsWorkspaceCatalogRefToMap(workspaceResponse.CatalogRef)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("catalog_ref", []map[string]interface{}{catalogRefMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting catalog_ref: %s", err))
		}
	}
	if workspaceResponse.Dependencies != nil {
		dependenciesMap, err := ResourceIBMSchematicsWorkspaceDependenciesToMap(workspaceResponse.Dependencies)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("dependencies", []map[string]interface{}{dependenciesMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting dependencies: %s", err))
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
	if workspaceResponse.SharedData != nil {
		sharedDataMap, err := ResourceIBMSchematicsWorkspaceSharedTargetDataToMap(workspaceResponse.SharedData)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("shared_data", []map[string]interface{}{sharedDataMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting shared_data: %s", err))
		}
	}
	if workspaceResponse.Tags != nil {
		if err = d.Set("tags", workspaceResponse.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	templateData := []map[string]interface{}{}
	if workspaceResponse.TemplateData != nil {
		for _, templateDataItem := range workspaceResponse.TemplateData {
			templateDataItemMap, err := ResourceIBMSchematicsWorkspaceTemplateSourceDataRequestToMap(&templateDataItem)
			if err != nil {
				return diag.FromErr(err)
			}
			templateData = append(templateData, templateDataItemMap)
		}
	}
	if err = d.Set("template_data", templateData); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_data: %s", err))
	}
	if err = d.Set("template_ref", workspaceResponse.TemplateRef); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting template_ref: %s", err))
	}
	if workspaceResponse.TemplateRepo != nil {
		templateRepoMap, err := ResourceIBMSchematicsWorkspaceTemplateRepoRequestToMap(workspaceResponse.TemplateRepo)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("template_repo", []map[string]interface{}{templateRepoMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting template_repo: %s", err))
		}
	}
	if workspaceResponse.Type != nil {
		if err = d.Set("type", workspaceResponse.Type); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
		}
	}
	if workspaceResponse.WorkspaceStatus != nil {
		workspaceStatusMap, err := ResourceIBMSchematicsWorkspaceWorkspaceStatusRequestToMap(workspaceResponse.WorkspaceStatus)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("workspace_status", []map[string]interface{}{workspaceStatusMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting workspace_status: %s", err))
		}
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
	if err = d.Set("last_health_check_at", flex.DateTimeToString(workspaceResponse.LastHealthCheckAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_health_check_at: %s", err))
	}
	runtimeData := []map[string]interface{}{}
	if workspaceResponse.RuntimeData != nil {
		for _, runtimeDataItem := range workspaceResponse.RuntimeData {
			runtimeDataItemMap, err := ResourceIBMSchematicsWorkspaceTemplateRunTimeDataResponseToMap(&runtimeDataItem)
			if err != nil {
				return diag.FromErr(err)
			}
			runtimeData = append(runtimeData, runtimeDataItemMap)
		}
	}
	if err = d.Set("runtime_data", runtimeData); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting runtime_data: %s", err))
	}
	if err = d.Set("status", workspaceResponse.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
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
	if workspaceResponse.LastJob != nil {
		lastJobMap, err := ResourceIBMSchematicsWorkspaceLastJobToMap(workspaceResponse.LastJob)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("last_job", []map[string]interface{}{lastJobMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_job: %s", err))
		}
	}
	if workspaceResponse.WorkspaceStatusMsg != nil {
		workspaceStatusMsgMap, err := ResourceIBMSchematicsWorkspaceWorkspaceStatusMessageToMap(workspaceResponse.WorkspaceStatusMsg)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("workspace_status_msg", []map[string]interface{}{workspaceStatusMsgMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting workspace_status_msg: %s", err))
		}
	}

	return nil
}

func ResourceIBMSchematicsWorkspaceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
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
		catalogRef, err := ResourceIBMSchematicsWorkspaceMapToCatalogRef(d.Get("catalog_ref.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateWorkspaceOptions.SetCatalogRef(catalogRef)
		hasChange = true
	}
	if d.HasChange("dependencies") {
		dependencies, err := ResourceIBMSchematicsWorkspaceMapToDependencies(d.Get("dependencies.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateWorkspaceOptions.SetDependencies(dependencies)
		hasChange = true
	}
	if d.HasChange("description") {
		updateWorkspaceOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("name") {
		updateWorkspaceOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("shared_data") {
		sharedData, err := ResourceIBMSchematicsWorkspaceMapToSharedTargetData(d.Get("shared_data.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateWorkspaceOptions.SetSharedData(sharedData)
		hasChange = true
	}
	if d.HasChange("tags") {
		// TODO: handle Tags of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("template_data") {
		// TODO: handle TemplateData of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("template_repo") {
		templateRepo, err := ResourceIBMSchematicsWorkspaceMapToTemplateRepoUpdateRequest(d.Get("template_repo.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateWorkspaceOptions.SetTemplateRepo(templateRepo)
		hasChange = true
	}
	if d.HasChange("type") {
		// TODO: handle Type of type TypeList -- not primitive, not model
		hasChange = true
	}
	if d.HasChange("workspace_status") {
		workspaceStatus, err := ResourceIBMSchematicsWorkspaceMapToWorkspaceStatusUpdateRequest(d.Get("workspace_status.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateWorkspaceOptions.SetWorkspaceStatus(workspaceStatus)
		hasChange = true
	}

	if hasChange {
		_, response, err := schematicsClient.UpdateWorkspaceWithContext(context, updateWorkspaceOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateWorkspaceWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateWorkspaceWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMSchematicsWorkspaceRead(context, d, meta)
}

func ResourceIBMSchematicsWorkspaceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteWorkspaceOptions := &schematicsv1.DeleteWorkspaceOptions{}

	deleteWorkspaceOptions.SetWID(d.Id())

	_, response, err := schematicsClient.DeleteWorkspaceWithContext(context, deleteWorkspaceOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteWorkspaceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteWorkspaceWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMSchematicsWorkspaceMapToCatalogRef(modelMap map[string]interface{}) (*schematicsv1.CatalogRef, error) {
	model := &schematicsv1.CatalogRef{}
	if modelMap["dry_run"] != nil {
		model.DryRun = core.BoolPtr(modelMap["dry_run"].(bool))
	}
	if modelMap["owning_account"] != nil && modelMap["owning_account"].(string) != "" {
		model.OwningAccount = core.StringPtr(modelMap["owning_account"].(string))
	}
	if modelMap["item_icon_url"] != nil && modelMap["item_icon_url"].(string) != "" {
		model.ItemIconURL = core.StringPtr(modelMap["item_icon_url"].(string))
	}
	if modelMap["item_id"] != nil && modelMap["item_id"].(string) != "" {
		model.ItemID = core.StringPtr(modelMap["item_id"].(string))
	}
	if modelMap["item_name"] != nil && modelMap["item_name"].(string) != "" {
		model.ItemName = core.StringPtr(modelMap["item_name"].(string))
	}
	if modelMap["item_readme_url"] != nil && modelMap["item_readme_url"].(string) != "" {
		model.ItemReadmeURL = core.StringPtr(modelMap["item_readme_url"].(string))
	}
	if modelMap["item_url"] != nil && modelMap["item_url"].(string) != "" {
		model.ItemURL = core.StringPtr(modelMap["item_url"].(string))
	}
	if modelMap["launch_url"] != nil && modelMap["launch_url"].(string) != "" {
		model.LaunchURL = core.StringPtr(modelMap["launch_url"].(string))
	}
	if modelMap["offering_version"] != nil && modelMap["offering_version"].(string) != "" {
		model.OfferingVersion = core.StringPtr(modelMap["offering_version"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToDependencies(modelMap map[string]interface{}) (*schematicsv1.Dependencies, error) {
	model := &schematicsv1.Dependencies{}
	if modelMap["parents"] != nil {
		parents := []string{}
		for _, parentsItem := range modelMap["parents"].([]interface{}) {
			parents = append(parents, parentsItem.(string))
		}
		model.Parents = parents
	}
	if modelMap["children"] != nil {
		children := []string{}
		for _, childrenItem := range modelMap["children"].([]interface{}) {
			children = append(children, childrenItem.(string))
		}
		model.Children = children
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToSharedTargetData(modelMap map[string]interface{}) (*schematicsv1.SharedTargetData, error) {
	model := &schematicsv1.SharedTargetData{}
	if modelMap["cluster_created_on"] != nil && modelMap["cluster_created_on"].(string) != "" {
		model.ClusterCreatedOn = core.StringPtr(modelMap["cluster_created_on"].(string))
	}
	if modelMap["cluster_id"] != nil && modelMap["cluster_id"].(string) != "" {
		model.ClusterID = core.StringPtr(modelMap["cluster_id"].(string))
	}
	if modelMap["cluster_name"] != nil && modelMap["cluster_name"].(string) != "" {
		model.ClusterName = core.StringPtr(modelMap["cluster_name"].(string))
	}
	if modelMap["cluster_type"] != nil && modelMap["cluster_type"].(string) != "" {
		model.ClusterType = core.StringPtr(modelMap["cluster_type"].(string))
	}
	if modelMap["entitlement_keys"] != nil {
		entitlementKeys := []interface{}{}
		for _, entitlementKeysItem := range modelMap["entitlement_keys"].([]interface{}) {
			entitlementKeys = append(entitlementKeys, entitlementKeysItem)
		}
		model.EntitlementKeys = entitlementKeys
	}
	if modelMap["namespace"] != nil && modelMap["namespace"].(string) != "" {
		model.Namespace = core.StringPtr(modelMap["namespace"].(string))
	}
	if modelMap["region"] != nil && modelMap["region"].(string) != "" {
		model.Region = core.StringPtr(modelMap["region"].(string))
	}
	if modelMap["resource_group_id"] != nil && modelMap["resource_group_id"].(string) != "" {
		model.ResourceGroupID = core.StringPtr(modelMap["resource_group_id"].(string))
	}
	if modelMap["worker_count"] != nil {
		model.WorkerCount = core.Int64Ptr(int64(modelMap["worker_count"].(int)))
	}
	if modelMap["worker_machine_type"] != nil && modelMap["worker_machine_type"].(string) != "" {
		model.WorkerMachineType = core.StringPtr(modelMap["worker_machine_type"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToTemplateSourceDataRequest(modelMap map[string]interface{}) (*schematicsv1.TemplateSourceDataRequest, error) {
	model := &schematicsv1.TemplateSourceDataRequest{}
	if modelMap["env_values"] != nil {
		envValues := []interface{}{}
		for _, envValuesItem := range modelMap["env_values"].([]interface{}) {
			envValues = append(envValues, envValuesItem)
		}
		model.EnvValues = envValues
	}
	if modelMap["env_values_metadata"] != nil {
		envValuesMetadata := []schematicsv1.EnvironmentValuesMetadata{}
		for _, envValuesMetadataItem := range modelMap["env_values_metadata"].([]interface{}) {
			envValuesMetadataItemModel, err := ResourceIBMSchematicsWorkspaceMapToEnvironmentValuesMetadata(envValuesMetadataItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			envValuesMetadata = append(envValuesMetadata, *envValuesMetadataItemModel)
		}
		model.EnvValuesMetadata = envValuesMetadata
	}
	if modelMap["folder"] != nil && modelMap["folder"].(string) != "" {
		model.Folder = core.StringPtr(modelMap["folder"].(string))
	}
	if modelMap["compact"] != nil {
		model.Compact = core.BoolPtr(modelMap["compact"].(bool))
	}
	if modelMap["init_state_file"] != nil && modelMap["init_state_file"].(string) != "" {
		model.InitStateFile = core.StringPtr(modelMap["init_state_file"].(string))
	}
	if modelMap["injectors"] != nil {
		injectors := []schematicsv1.InjectTerraformTemplateInner{}
		for _, injectorsItem := range modelMap["injectors"].([]interface{}) {
			injectorsItemModel, err := ResourceIBMSchematicsWorkspaceMapToInjectTerraformTemplateInner(injectorsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			injectors = append(injectors, *injectorsItemModel)
		}
		model.Injectors = injectors
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["uninstall_script_name"] != nil && modelMap["uninstall_script_name"].(string) != "" {
		model.UninstallScriptName = core.StringPtr(modelMap["uninstall_script_name"].(string))
	}
	if modelMap["values"] != nil && modelMap["values"].(string) != "" {
		model.Values = core.StringPtr(modelMap["values"].(string))
	}
	if modelMap["values_metadata"] != nil {
		valuesMetadata := []interface{}{}
		for _, valuesMetadataItem := range modelMap["values_metadata"].([]interface{}) {
			valuesMetadata = append(valuesMetadata, valuesMetadataItem)
		}
		model.ValuesMetadata = valuesMetadata
	}
	if modelMap["variablestore"] != nil {
		variablestore := []schematicsv1.WorkspaceVariableRequest{}
		for _, variablestoreItem := range modelMap["variablestore"].([]interface{}) {
			variablestoreItemModel, err := ResourceIBMSchematicsWorkspaceMapToWorkspaceVariableRequest(variablestoreItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			variablestore = append(variablestore, *variablestoreItemModel)
		}
		model.Variablestore = variablestore
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToEnvironmentValuesMetadata(modelMap map[string]interface{}) (*schematicsv1.EnvironmentValuesMetadata, error) {
	model := &schematicsv1.EnvironmentValuesMetadata{}
	if modelMap["hidden"] != nil {
		model.Hidden = core.BoolPtr(modelMap["hidden"].(bool))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["secure"] != nil {
		model.Secure = core.BoolPtr(modelMap["secure"].(bool))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToInjectTerraformTemplateInner(modelMap map[string]interface{}) (*schematicsv1.InjectTerraformTemplateInner, error) {
	model := &schematicsv1.InjectTerraformTemplateInner{}
	if modelMap["tft_git_url"] != nil && modelMap["tft_git_url"].(string) != "" {
		model.TftGitURL = core.StringPtr(modelMap["tft_git_url"].(string))
	}
	if modelMap["tft_git_token"] != nil && modelMap["tft_git_token"].(string) != "" {
		model.TftGitToken = core.StringPtr(modelMap["tft_git_token"].(string))
	}
	if modelMap["tft_prefix"] != nil && modelMap["tft_prefix"].(string) != "" {
		model.TftPrefix = core.StringPtr(modelMap["tft_prefix"].(string))
	}
	if modelMap["injection_type"] != nil && modelMap["injection_type"].(string) != "" {
		model.InjectionType = core.StringPtr(modelMap["injection_type"].(string))
	}
	if modelMap["tft_name"] != nil && modelMap["tft_name"].(string) != "" {
		model.TftName = core.StringPtr(modelMap["tft_name"].(string))
	}
	if modelMap["tft_parameters"] != nil {
		tftParameters := []schematicsv1.InjectTerraformTemplateInnerTftParametersItem{}
		for _, tftParametersItem := range modelMap["tft_parameters"].([]interface{}) {
			tftParametersItemModel, err := ResourceIBMSchematicsWorkspaceMapToInjectTerraformTemplateInnerTftParametersItem(tftParametersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			tftParameters = append(tftParameters, *tftParametersItemModel)
		}
		model.TftParameters = tftParameters
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToInjectTerraformTemplateInnerTftParametersItem(modelMap map[string]interface{}) (*schematicsv1.InjectTerraformTemplateInnerTftParametersItem, error) {
	model := &schematicsv1.InjectTerraformTemplateInnerTftParametersItem{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToWorkspaceVariableRequest(modelMap map[string]interface{}) (*schematicsv1.WorkspaceVariableRequest, error) {
	model := &schematicsv1.WorkspaceVariableRequest{}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["secure"] != nil {
		model.Secure = core.BoolPtr(modelMap["secure"].(bool))
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["use_default"] != nil {
		model.UseDefault = core.BoolPtr(modelMap["use_default"].(bool))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToTemplateRepoRequest(modelMap map[string]interface{}) (*schematicsv1.TemplateRepoRequest, error) {
	model := &schematicsv1.TemplateRepoRequest{}
	if modelMap["branch"] != nil && modelMap["branch"].(string) != "" {
		model.Branch = core.StringPtr(modelMap["branch"].(string))
	}
	if modelMap["release"] != nil && modelMap["release"].(string) != "" {
		model.Release = core.StringPtr(modelMap["release"].(string))
	}
	if modelMap["repo_sha_value"] != nil && modelMap["repo_sha_value"].(string) != "" {
		model.RepoShaValue = core.StringPtr(modelMap["repo_sha_value"].(string))
	}
	if modelMap["repo_url"] != nil && modelMap["repo_url"].(string) != "" {
		model.RepoURL = core.StringPtr(modelMap["repo_url"].(string))
	}
	if modelMap["url"] != nil && modelMap["url"].(string) != "" {
		model.URL = core.StringPtr(modelMap["url"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToTemplateRepoUpdateRequest(modelMap map[string]interface{}) (*schematicsv1.TemplateRepoUpdateRequest, error) {
	model := &schematicsv1.TemplateRepoUpdateRequest{}
	if modelMap["branch"] != nil && modelMap["branch"].(string) != "" {
		model.Branch = core.StringPtr(modelMap["branch"].(string))
	}
	if modelMap["release"] != nil && modelMap["release"].(string) != "" {
		model.Release = core.StringPtr(modelMap["release"].(string))
	}
	if modelMap["repo_sha_value"] != nil && modelMap["repo_sha_value"].(string) != "" {
		model.RepoShaValue = core.StringPtr(modelMap["repo_sha_value"].(string))
	}
	if modelMap["repo_url"] != nil && modelMap["repo_url"].(string) != "" {
		model.RepoURL = core.StringPtr(modelMap["repo_url"].(string))
	}
	if modelMap["url"] != nil && modelMap["url"].(string) != "" {
		model.URL = core.StringPtr(modelMap["url"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToWorkspaceStatusRequest(modelMap map[string]interface{}) (*schematicsv1.WorkspaceStatusRequest, error) {
	model := &schematicsv1.WorkspaceStatusRequest{}
	if modelMap["frozen"] != nil {
		model.Frozen = core.BoolPtr(modelMap["frozen"].(bool))
	}
	if modelMap["frozen_at"] != nil {

	}
	if modelMap["frozen_by"] != nil && modelMap["frozen_by"].(string) != "" {
		model.FrozenBy = core.StringPtr(modelMap["frozen_by"].(string))
	}
	if modelMap["locked"] != nil {
		model.Locked = core.BoolPtr(modelMap["locked"].(bool))
	}
	if modelMap["locked_by"] != nil && modelMap["locked_by"].(string) != "" {
		model.LockedBy = core.StringPtr(modelMap["locked_by"].(string))
	}
	if modelMap["locked_time"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceMapToWorkspaceStatusUpdateRequest(modelMap map[string]interface{}) (*schematicsv1.WorkspaceStatusUpdateRequest, error) {
	model := &schematicsv1.WorkspaceStatusUpdateRequest{}
	if modelMap["frozen"] != nil {
		model.Frozen = core.BoolPtr(modelMap["frozen"].(bool))
	}
	if modelMap["frozen_at"] != nil {

	}
	if modelMap["frozen_by"] != nil && modelMap["frozen_by"].(string) != "" {
		model.FrozenBy = core.StringPtr(modelMap["frozen_by"].(string))
	}
	if modelMap["locked"] != nil {
		model.Locked = core.BoolPtr(modelMap["locked"].(bool))
	}
	if modelMap["locked_by"] != nil && modelMap["locked_by"].(string) != "" {
		model.LockedBy = core.StringPtr(modelMap["locked_by"].(string))
	}
	if modelMap["locked_time"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsWorkspaceCatalogRefToMap(model *schematicsv1.CatalogRef) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DryRun != nil {
		modelMap["dry_run"] = model.DryRun
	}
	if model.OwningAccount != nil {
		modelMap["owning_account"] = model.OwningAccount
	}
	if model.ItemIconURL != nil {
		modelMap["item_icon_url"] = model.ItemIconURL
	}
	if model.ItemID != nil {
		modelMap["item_id"] = model.ItemID
	}
	if model.ItemName != nil {
		modelMap["item_name"] = model.ItemName
	}
	if model.ItemReadmeURL != nil {
		modelMap["item_readme_url"] = model.ItemReadmeURL
	}
	if model.ItemURL != nil {
		modelMap["item_url"] = model.ItemURL
	}
	if model.LaunchURL != nil {
		modelMap["launch_url"] = model.LaunchURL
	}
	if model.OfferingVersion != nil {
		modelMap["offering_version"] = model.OfferingVersion
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceDependenciesToMap(model *schematicsv1.Dependencies) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Parents != nil {
		modelMap["parents"] = model.Parents
	}
	if model.Children != nil {
		modelMap["children"] = model.Children
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceSharedTargetDataToMap(model *schematicsv1.SharedTargetDataResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterID != nil {
		modelMap["cluster_id"] = model.ClusterID
	}
	if model.ClusterName != nil {
		modelMap["cluster_name"] = model.ClusterName
	}
	if model.EntitlementKeys != nil {
		entitlementKeys := []interface{}{}
		for _, entitlementKeysItem := range model.EntitlementKeys {
			entitlementKeys = append(entitlementKeys, entitlementKeysItem)
		}
		modelMap["entitlement_keys"] = entitlementKeys
	}
	if model.Namespace != nil {
		modelMap["namespace"] = model.Namespace
	}
	if model.Region != nil {
		modelMap["region"] = model.Region
	}
	if model.ResourceGroupID != nil {
		modelMap["resource_group_id"] = model.ResourceGroupID
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceTemplateSourceDataRequestToMap(model *schematicsv1.TemplateSourceDataResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EnvValues != nil {
		envValues := []schematicsv1.EnvVariableResponse{}
		for _, envValuesItem := range model.EnvValues {
			envValues = append(envValues, envValuesItem)
		}
		modelMap["env_values"] = envValues
	}
	if model.Folder != nil {
		modelMap["folder"] = model.Folder
	}
	if model.Compact != nil {
		modelMap["compact"] = model.Compact
	}
	if model.HasGithubtoken != nil {
		modelMap["compact"] = model.HasGithubtoken
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.UninstallScriptName != nil {
		modelMap["uninstall_script_name"] = model.UninstallScriptName
	}
	if model.Values != nil {
		modelMap["values"] = model.Values
	}
	if model.ValuesMetadata != nil {
		valuesMetadata := []interface{}{}
		for _, valuesMetadataItem := range model.ValuesMetadata {
			valuesMetadata = append(valuesMetadata, valuesMetadataItem)
		}
		modelMap["values_metadata"] = valuesMetadata
	}
	if model.Variablestore != nil {
		variablestore := []interface{}{}
		for _, variablestoreItem := range model.Variablestore {
			variablestoreItemMap, err := ResourceIBMSchematicsWorkspaceWorkspaceVariableResponseToMap(&variablestoreItem)
			if err != nil {
				return modelMap, err
			}
			variablestore = append(variablestore, variablestoreItemMap)
		}
		modelMap["variablestore"] = variablestore
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceEnvironmentValuesMetadataToMap(model *schematicsv1.EnvironmentValuesMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Hidden != nil {
		modelMap["hidden"] = model.Hidden
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Secure != nil {
		modelMap["secure"] = model.Secure
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceInjectTerraformTemplateInnerToMap(model *schematicsv1.InjectTerraformTemplateInner) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TftGitURL != nil {
		modelMap["tft_git_url"] = model.TftGitURL
	}
	if model.TftGitToken != nil {
		modelMap["tft_git_token"] = model.TftGitToken
	}
	if model.TftPrefix != nil {
		modelMap["tft_prefix"] = model.TftPrefix
	}
	if model.InjectionType != nil {
		modelMap["injection_type"] = model.InjectionType
	}
	if model.TftName != nil {
		modelMap["tft_name"] = model.TftName
	}
	if model.TftParameters != nil {
		tftParameters := []map[string]interface{}{}
		for _, tftParametersItem := range model.TftParameters {
			tftParametersItemMap, err := ResourceIBMSchematicsWorkspaceInjectTerraformTemplateInnerTftParametersItemToMap(&tftParametersItem)
			if err != nil {
				return modelMap, err
			}
			tftParameters = append(tftParameters, tftParametersItemMap)
		}
		modelMap["tft_parameters"] = tftParameters
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceInjectTerraformTemplateInnerTftParametersItemToMap(model *schematicsv1.InjectTerraformTemplateInnerTftParametersItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceWorkspaceVariableRequestToMap(model *schematicsv1.WorkspaceVariableRequest) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Secure != nil {
		modelMap["secure"] = model.Secure
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = model.UseDefault
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceWorkspaceVariableResponseToMap(model *schematicsv1.WorkspaceVariableResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Secure != nil {
		modelMap["secure"] = model.Secure
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceTemplateRepoRequestToMap(model *schematicsv1.TemplateRepoResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Branch != nil {
		modelMap["branch"] = model.Branch
	}
	if model.Release != nil {
		modelMap["release"] = model.Release
	}
	if model.RepoShaValue != nil {
		modelMap["repo_sha_value"] = model.RepoShaValue
	}
	if model.RepoURL != nil {
		modelMap["repo_url"] = model.RepoURL
	}
	if model.URL != nil {
		modelMap["url"] = model.URL
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceWorkspaceStatusRequestToMap(model *schematicsv1.WorkspaceStatusResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Frozen != nil {
		modelMap["frozen"] = model.Frozen
	}
	if model.FrozenAt != nil {
		modelMap["frozen_at"] = model.FrozenAt.String()
	}
	if model.FrozenBy != nil {
		modelMap["frozen_by"] = model.FrozenBy
	}
	if model.Locked != nil {
		modelMap["locked"] = model.Locked
	}
	if model.LockedBy != nil {
		modelMap["locked_by"] = model.LockedBy
	}
	if model.LockedTime != nil {
		modelMap["locked_time"] = model.LockedTime.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceTemplateRunTimeDataResponseToMap(model *schematicsv1.TemplateRunTimeDataResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EngineCmd != nil {
		modelMap["engine_cmd"] = model.EngineCmd
	}
	if model.EngineName != nil {
		modelMap["engine_name"] = model.EngineName
	}
	if model.EngineVersion != nil {
		modelMap["engine_version"] = model.EngineVersion
	}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.LogStoreURL != nil {
		modelMap["log_store_url"] = model.LogStoreURL
	}
	if model.OutputValues != nil {
		outputValues := []interface{}{}
		for _, outputValuesItem := range model.OutputValues {
			outputValues = append(outputValues, outputValuesItem)
		}
		modelMap["output_values"] = outputValues
	}
	if model.Resources != nil {
		resources := []interface{}{}
		for _, resourcesItem := range model.Resources {
			resources = append(resources, resourcesItem)
		}
		modelMap["resources"] = resources
	}
	if model.StateStoreURL != nil {
		modelMap["state_store_url"] = model.StateStoreURL
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceLastJobToMap(model *schematicsv1.LastJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.JobID != nil {
		modelMap["job_id"] = model.JobID
	}
	if model.JobName != nil {
		modelMap["job_name"] = model.JobName
	}
	if model.JobStatus != nil {
		modelMap["job_status"] = model.JobStatus
	}
	return modelMap, nil
}

func ResourceIBMSchematicsWorkspaceWorkspaceStatusMessageToMap(model *schematicsv1.WorkspaceStatusMessage) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMsg != nil {
		modelMap["status_msg"] = model.StatusMsg
	}
	return modelMap, nil
}
