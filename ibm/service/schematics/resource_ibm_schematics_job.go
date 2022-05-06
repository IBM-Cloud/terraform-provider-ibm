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

func ResourceIBMSchematicsJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMSchematicsJobCreate,
		ReadContext:   ResourceIBMSchematicsJobRead,
		UpdateContext: ResourceIBMSchematicsJobUpdate,
		DeleteContext: ResourceIBMSchematicsJobDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"refresh_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IAM refresh token for the user or service identity.  **Retrieving refresh token**:   * Use `export IBMCLOUD_API_KEY=<ibmcloud_api_key>`, and execute `curl -X POST \"https://iam.cloud.ibm.com/identity/token\" -H \"Content-Type: application/x-www-form-urlencoded\" -d \"grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=$IBMCLOUD_API_KEY\" -u bx:bx`.   * For more information, about creating IAM access token and API Docs, refer, [IAM access token](/apidocs/iam-identity-token-api#gettoken-password) and [Create API key](/apidocs/iam-identity-token-api#create-api-key).    **Limitation**:   * If the token is expired, you can use `refresh token` to get a new IAM access token.   * The `refresh_token` parameter cannot be used to retrieve a new IAM access token.   * When the IAM access token is about to expire, use the API key to create a new access token.",
			},
			"command_object": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_job", "command_object"),
				Description:  "Name of the Schematics automation resource.",
			},
			"command_object_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Job command object id (workspace-id, action-id).",
			},
			"command_name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_job", "command_name"),
				Description:  "Schematics job command name.",
			},
			"command_parameter": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Schematics job command parameter (playbook-name).",
			},
			"command_options": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Command line options for the command.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"job_inputs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Job inputs used by Action or Workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"job_env_settings": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Environment variables used by the Job while performing Action or Workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of aliases for the variable name.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "User defined tags, while running the job.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_schematics_job", "location"),
				Description:  "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Job Status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"position_in_queue": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: "Position of job in pending queue.",
						},
						"total_in_queue": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Description: "Total no. of jobs in pending queue.",
						},
						"workspace_job_status": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Workspace Job Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workspace_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Workspace name.",
									},
									"status_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Status of Jobs.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Workspace job status message (eg. App1_Setup_Pending, for a 'Setup' flow in the 'App1' Workspace).",
									},
									"flow_status": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Environment Flow JOB Status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"flow_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "flow id.",
												},
												"flow_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "flow name.",
												},
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Flow Job status message - to be displayed along with the status_code;.",
												},
												"workitems": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Environment's individual workItem status details;.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"workspace_id": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Workspace id.",
															},
															"workspace_name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "workspace name.",
															},
															"job_id": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "workspace job id.",
															},
															"status_code": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Status of Jobs.",
															},
															"status_message": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "workitem job status message;.",
															},
															"updated_at": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "workitem job status updation timestamp.",
															},
														},
													},
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"template_status": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Workspace Flow Template job status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"template_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Template Id.",
												},
												"template_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Template name.",
												},
												"flow_index": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Index of the template in the Flow.",
												},
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Template job status message (eg. VPCt1_Apply_Pending, for a 'VPCt1' Template).",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
									"commands": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of terraform commands executed and their status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Name of the command.",
												},
												"outcome": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "outcome of the command.",
												},
											},
										},
									},
								},
							},
						},
						"action_job_status": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Action Job Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Action name.",
									},
									"status_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Status of Jobs.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Action Job status message - to be displayed along with the action_status_code.",
									},
									"bastion_status_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Status of Resources.",
									},
									"bastion_status_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Bastion status message - to be displayed along with the bastion_status_code;.",
									},
									"targets_status_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Status of Resources.",
									},
									"targets_status_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Aggregated status message for all target resources,  to be displayed along with the targets_status_code;.",
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"system_job_status": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "System Job Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_status_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "System job message.",
									},
									"system_status_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Status of Jobs.",
									},
									"schematics_resource_status": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "job staus for each schematics resource.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "system job status message.",
												},
												"schematics_resource_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "id for each resource which is targeted as a part of system job.",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"flow_job_status": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Environment Flow JOB Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flow_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "flow id.",
									},
									"flow_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "flow name.",
									},
									"status_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Status of Jobs.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Flow Job status message - to be displayed along with the status_code;.",
									},
									"workitems": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Environment's individual workItem status details;.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"workspace_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Workspace id.",
												},
												"workspace_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "workspace name.",
												},
												"job_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "workspace job id.",
												},
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "workitem job status message;.",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "workitem job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
					},
				},
			},
			"data": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Job data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of Job.",
						},
						"workspace_job_data": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Workspace Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workspace_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Workspace name.",
									},
									"flow_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Flow Id.",
									},
									"flow_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Flow name.",
									},
									"inputs": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Input variables data used by the Workspace Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of aliases for the variable name.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"outputs": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Output variables data from the Workspace Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of aliases for the variable name.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"settings": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Environment variables used by all the templates in the Workspace.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of aliases for the variable name.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"template_data": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Input / output data of the Template in the Workspace Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"template_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Template Id.",
												},
												"template_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Template name.",
												},
												"flow_index": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Index of the template in the Flow.",
												},
												"inputs": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Job inputs used by the Templates.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"outputs": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Job output from the Templates.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"settings": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Environment variables used by the template.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"action_job_data": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Action Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Flow name.",
									},
									"inputs": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Input variables data used by the Action Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of aliases for the variable name.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"outputs": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Output variables data from the Action Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of aliases for the variable name.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"settings": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Environment variables used by all the templates in the Action.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of aliases for the variable name.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Optional:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem:        &schema.Schema{Type: schema.TypeString},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
									"inventory_record": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Complete inventory resource details with user inputs and system generated data.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The unique name of your Inventory.  The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Inventory id.",
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The description of your Inventory.  The description can be up to 2048 characters long in size.",
												},
												"location": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
												},
												"resource_group": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Resource-group name for the Inventory definition.  By default, Inventory will be created in Default Resource Group.",
												},
												"created_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Inventory creation time.",
												},
												"created_by": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Email address of user who created the Inventory.",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Inventory updation time.",
												},
												"updated_by": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
													Description: "Email address of user who updated the Inventory.",
												},
												"inventories_ini": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Input inventory of host and host group for the playbook,  in the .ini file format.",
												},
												"resource_queries": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"materialized_inventory": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Materialized inventory details used by the Action Job, in .ini format.",
									},
								},
							},
						},
						"system_job_data": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Controls Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Key ID for which key event is generated.",
									},
									"schematics_resource_id": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of the schematics resource id.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"flow_job_data": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Flow Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flow_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Flow ID.",
									},
									"flow_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Flow Name.",
									},
									"workitems": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Job data used by each workitem Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"command_object_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "command object id.",
												},
												"command_object_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "command object name.",
												},
												"layers": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "layer name.",
												},
												"source_type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Type of source for the Template.",
												},
												"source": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Source of templates, playbooks, or controls.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"source_type": &schema.Schema{
																Type:        schema.TypeString,
																Required:    true,
																Description: "Type of source for the Template.",
															},
															"git": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "The connection details to the Git source repository.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"computed_git_repo_url": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**.",
																		},
																		"git_repo_url": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The URL to the Git repository that can be used to clone the template.",
																		},
																		"git_token": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The Personal Access Token (PAT) to connect to the Git URLs.",
																		},
																		"git_repo_folder": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The name of the folder in the Git repository, that contains the template.",
																		},
																		"git_release": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The name of the release tag that are used to fetch the Git repository.",
																		},
																		"git_branch": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The name of the branch that are used to fetch the Git repository.",
																		},
																	},
																},
															},
															"catalog": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "The connection details to the IBM Cloud Catalog source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"catalog_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The name of the private catalog.",
																		},
																		"offering_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The name of an offering in the IBM Cloud Catalog.",
																		},
																		"offering_version": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The version string of an offering in the IBM Cloud Catalog.",
																		},
																		"offering_kind": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The type of an offering, in the IBM Cloud Catalog.",
																		},
																		"offering_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The ID of an offering in the IBM Cloud Catalog.",
																		},
																		"offering_version_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The ID of an offering version the IBM Cloud Catalog.",
																		},
																		"offering_repo_url": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The repository URL of an offering, in the IBM Cloud Catalog.",
																		},
																	},
																},
															},
														},
													},
												},
												"inputs": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Input variables data for the workItem used in FlowJob.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"outputs": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Output variables for the workItem.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"settings": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Environment variables for the workItem.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Optional:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem:        &schema.Schema{Type: schema.TypeString},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Optional:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"last_job": &schema.Schema{
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Status of the last job executed by the workitem.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"command_object": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Name of the Schematics automation resource.",
															},
															"command_object_name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "command object name (workspace_name/action_name).",
															},
															"command_object_id": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Workitem command object id, maps to workspace_id or action_id.",
															},
															"command_name": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Schematics job command name.",
															},
															"job_id": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Workspace job id.",
															},
															"job_status": &schema.Schema{
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Status of Jobs.",
															},
														},
													},
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
					},
				},
			},
			"bastion": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Describes a bastion resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Bastion Name(Unique).",
						},
						"host": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Reference to the Inventory resource definition.",
						},
					},
				},
			},
			"log_summary": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Job log summary record.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Workspace Id.",
						},
						"job_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of Job.",
						},
						"log_start_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Job log start timestamp.",
						},
						"log_analyzed_till": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Job log update timestamp.",
						},
						"elapsed_time": &schema.Schema{
							Type:        schema.TypeFloat,
							Optional:    true,
							Computed:    true,
							Description: "Job log elapsed time (log_analyzed_till - log_start_at).",
						},
						"log_errors": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Job log errors.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error_code": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Error code in the Log.",
									},
									"error_msg": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Summary error message in the log.",
									},
									"error_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Description: "Number of occurrence.",
									},
								},
							},
						},
						"repo_download_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Repo download Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scanned_file_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of files scanned.",
									},
									"quarantined_file_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of files quarantined.",
									},
									"detected_filetype": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Detected template or data file type.",
									},
									"inputs_count": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Number of inputs detected.",
									},
									"outputs_count": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Number of outputs detected.",
									},
								},
							},
						},
						"workspace_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Workspace Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resources_add": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of resources add.",
									},
									"resources_modify": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of resources modify.",
									},
									"resources_destroy": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of resources destroy.",
									},
								},
							},
						},
						"flow_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Flow Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workitems_completed": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of workitems completed successfully.",
									},
									"workitems_pending": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of workitems pending in the flow.",
									},
									"workitems_failed": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "Number of workitems failed.",
									},
									"workitems": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"workspace_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "workspace ID.",
												},
												"job_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "workspace JOB ID.",
												},
												"resources_add": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Computed:    true,
													Description: "Number of resources add.",
												},
												"resources_modify": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Computed:    true,
													Description: "Number of resources modify.",
												},
												"resources_destroy": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Computed:    true,
													Description: "Number of resources destroy.",
												},
												"log_url": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Log url for job.",
												},
											},
										},
									},
								},
							},
						},
						"action_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Flow Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "number of targets or hosts.",
									},
									"task_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "number of tasks in playbook.",
									},
									"play_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "number of plays in playbook.",
									},
									"recap": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Recap records.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target": &schema.Schema{
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of target or host name.",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"ok": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "Number of OK.",
												},
												"changed": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "Number of changed.",
												},
												"failed": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "Number of failed.",
												},
												"skipped": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "Number of skipped.",
												},
												"unreachable": &schema.Schema{
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "Number of unreachable.",
												},
											},
										},
									},
								},
							},
						},
						"system_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "System Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Computed:    true,
										Description: "number of targets or hosts.",
									},
									"success": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Description: "Number of passed.",
									},
									"failed": &schema.Schema{
										Type:        schema.TypeFloat,
										Optional:    true,
										Description: "Number of failed.",
									},
								},
							},
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job name, uniquely derived from the related Workspace or Action.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of your job is derived from the related action or workspace.  The description can be up to 2048 characters long in size.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource-group name derived from the related Workspace or Action.",
			},
			"submitted_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job submission time.",
			},
			"submitted_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of user who submitted the job.",
			},
			"start_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job start time.",
			},
			"end_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job end time.",
			},
			"duration": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Duration of job execution; example 40 sec.",
			},
			"log_store_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job log store URL.",
			},
			"state_store_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job state store URL.",
			},
			"results_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job results store URL.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job status updation timestamp.",
			},
			"job_runner_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the Job Runner.",
			},
		},
	}
}

func ResourceIBMSchematicsJobValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "command_object",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "action, environment, system, workspace",
		},
		validate.ValidateSchema{
			Identifier:                 "command_name",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "ansible_playbook_check, ansible_playbook_run, create_action, create_cart, create_environment, create_workspace, delete_action, delete_environment, delete_workspace, environment_init, environment_install, environment_uninstall, patch_action, patch_workspace, put_action, put_environment, put_workspace, repository_process, system_key_delete, system_key_disable, system_key_enable, system_key_restore, system_key_rotate, terraform_commands, workspace_apply, workspace_destroy, workspace_plan, workspace_refresh",
		},
		validate.ValidateSchema{
			Identifier:                 "location",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "eu-de, eu-gb, us-east, us-south",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_schematics_job", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMSchematicsJobCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createJobOptions := &schematicsv1.CreateJobOptions{}

	createJobOptions.SetRefreshToken(d.Get("refresh_token").(string))
	if _, ok := d.GetOk("command_object"); ok {
		createJobOptions.SetCommandObject(d.Get("command_object").(string))
	}
	if _, ok := d.GetOk("command_object_id"); ok {
		createJobOptions.SetCommandObjectID(d.Get("command_object_id").(string))
	}
	if _, ok := d.GetOk("command_name"); ok {
		createJobOptions.SetCommandName(d.Get("command_name").(string))
	}
	if _, ok := d.GetOk("command_parameter"); ok {
		createJobOptions.SetCommandParameter(d.Get("command_parameter").(string))
	}
	if _, ok := d.GetOk("command_options"); ok {
		createJobOptions.SetCommandOptions(d.Get("command_options").([]string))
	}
	if _, ok := d.GetOk("job_inputs"); ok {
		var jobInputs []schematicsv1.VariableData
		for _, e := range d.Get("job_inputs").([]interface{}) {
			value := e.(map[string]interface{})
			jobInputsItem, err := ResourceIBMSchematicsJobMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			jobInputs = append(jobInputs, *jobInputsItem)
		}
		createJobOptions.SetInputs(jobInputs)
	}
	if _, ok := d.GetOk("job_env_settings"); ok {
		var jobEnvSettings []schematicsv1.VariableData
		for _, e := range d.Get("job_env_settings").([]interface{}) {
			value := e.(map[string]interface{})
			jobEnvSettingsItem, err := ResourceIBMSchematicsJobMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			jobEnvSettings = append(jobEnvSettings, *jobEnvSettingsItem)
		}
		createJobOptions.SetSettings(jobEnvSettings)
	}
	if _, ok := d.GetOk("tags"); ok {
		createJobOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("location"); ok {
		createJobOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("status"); ok {
		statusModel, err := ResourceIBMSchematicsJobMapToJobStatus(d.Get("status.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createJobOptions.SetStatus(statusModel)
	}
	if _, ok := d.GetOk("data"); ok {
		dataModel, err := ResourceIBMSchematicsJobMapToJobData(d.Get("data.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createJobOptions.SetData(dataModel)
	}
	if _, ok := d.GetOk("bastion"); ok {
		bastionModel, err := ResourceIBMSchematicsJobMapToBastionResourceDefinition(d.Get("bastion.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createJobOptions.SetBastion(bastionModel)
	}
	if _, ok := d.GetOk("log_summary"); ok {
		logSummaryModel, err := ResourceIBMSchematicsJobMapToJobLogSummary(d.Get("log_summary.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createJobOptions.SetLogSummary(logSummaryModel)
	}

	job, response, err := schematicsClient.CreateJobWithContext(context, createJobOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateJobWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateJobWithContext failed %s\n%s", err, response))
	}

	d.SetId(*job.ID)

	return ResourceIBMSchematicsJobRead(context, d, meta)
}

func ResourceIBMSchematicsJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getJobOptions := &schematicsv1.GetJobOptions{}

	getJobOptions.SetJobID(d.Id())

	job, response, err := schematicsClient.GetJobWithContext(context, getJobOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetJobWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetJobWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("command_object", job.CommandObject); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting command_object: %s", err))
	}
	if err = d.Set("command_object_id", job.CommandObjectID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting command_object_id: %s", err))
	}
	if err = d.Set("command_name", job.CommandName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting command_name: %s", err))
	}
	if err = d.Set("command_parameter", job.CommandParameter); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting command_parameter: %s", err))
	}
	if job.CommandOptions != nil {
		if err = d.Set("command_options", job.CommandOptions); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting command_options: %s", err))
		}
	}
	jobInputs := []map[string]interface{}{}
	if job.Inputs != nil {
		for _, jobInputsItem := range job.Inputs {
			jobInputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&jobInputsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			jobInputs = append(jobInputs, jobInputsItemMap)
		}
	}
	if err = d.Set("job_inputs", jobInputs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_inputs: %s", err))
	}
	jobEnvSettings := []map[string]interface{}{}
	if job.Settings != nil {
		for _, jobEnvSettingsItem := range job.Settings {
			jobEnvSettingsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&jobEnvSettingsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			jobEnvSettings = append(jobEnvSettings, jobEnvSettingsItemMap)
		}
	}
	if err = d.Set("job_env_settings", jobEnvSettings); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_env_settings: %s", err))
	}
	if job.Tags != nil {
		if err = d.Set("tags", job.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	if err = d.Set("location", job.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if job.Status != nil {
		statusMap, err := ResourceIBMSchematicsJobJobStatusToMap(job.Status)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("status", []map[string]interface{}{statusMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
		}
	}
	if job.Data != nil {
		dataMap, err := ResourceIBMSchematicsJobJobDataToMap(job.Data)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("data", []map[string]interface{}{dataMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting data: %s", err))
		}
	}
	if job.Bastion != nil {
		bastionMap, err := ResourceIBMSchematicsJobBastionResourceDefinitionToMap(job.Bastion)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("bastion", []map[string]interface{}{bastionMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting bastion: %s", err))
		}
	}
	if job.LogSummary != nil {
		logSummaryMap, err := ResourceIBMSchematicsJobJobLogSummaryToMap(job.LogSummary)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("log_summary", []map[string]interface{}{logSummaryMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting log_summary: %s", err))
		}
	}
	if err = d.Set("name", job.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("description", job.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("resource_group", job.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}
	if err = d.Set("submitted_at", flex.DateTimeToString(job.SubmittedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting submitted_at: %s", err))
	}
	if err = d.Set("submitted_by", job.SubmittedBy); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting submitted_by: %s", err))
	}
	if err = d.Set("start_at", flex.DateTimeToString(job.StartAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting start_at: %s", err))
	}
	if err = d.Set("end_at", flex.DateTimeToString(job.EndAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting end_at: %s", err))
	}
	if err = d.Set("duration", job.Duration); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting duration: %s", err))
	}
	if err = d.Set("log_store_url", job.LogStoreURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting log_store_url: %s", err))
	}
	if err = d.Set("state_store_url", job.StateStoreURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state_store_url: %s", err))
	}
	if err = d.Set("results_url", job.ResultsURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting results_url: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(job.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("job_runner_id", job.JobRunnerID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_runner_id: %s", err))
	}

	return nil
}

func ResourceIBMSchematicsJobUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateJobOptions := &schematicsv1.UpdateJobOptions{}

	updateJobOptions.SetJobID(d.Id())
	updateJobOptions.SetRefreshToken(d.Get("refresh_token").(string))
	if _, ok := d.GetOk("command_object"); ok {
		updateJobOptions.SetCommandObject(d.Get("command_object").(string))
	}
	if _, ok := d.GetOk("command_object_id"); ok {
		updateJobOptions.SetCommandObjectID(d.Get("command_object_id").(string))
	}
	if _, ok := d.GetOk("command_name"); ok {
		updateJobOptions.SetCommandName(d.Get("command_name").(string))
	}
	if _, ok := d.GetOk("command_parameter"); ok {
		updateJobOptions.SetCommandParameter(d.Get("command_parameter").(string))
	}
	if _, ok := d.GetOk("command_options"); ok {
		updateJobOptions.SetCommandOptions(d.Get("command_options").([]string))
	}
	if _, ok := d.GetOk("job_inputs"); ok {
		var jobInputs []schematicsv1.VariableData
		for _, e := range d.Get("job_inputs").([]interface{}) {
			value := e.(map[string]interface{})
			jobInputsItem, err := ResourceIBMSchematicsJobMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			jobInputs = append(jobInputs, *jobInputsItem)
		}
		updateJobOptions.SetInputs(jobInputs)
	}
	if _, ok := d.GetOk("job_env_settings"); ok {
		var jobEnvSettings []schematicsv1.VariableData
		for _, e := range d.Get("job_env_settings").([]interface{}) {
			value := e.(map[string]interface{})
			jobEnvSettingsItem, err := ResourceIBMSchematicsJobMapToVariableData(value)
			if err != nil {
				return diag.FromErr(err)
			}
			jobEnvSettings = append(jobEnvSettings, *jobEnvSettingsItem)
		}
		updateJobOptions.SetSettings(jobEnvSettings)
	}
	if _, ok := d.GetOk("tags"); ok {
		updateJobOptions.SetTags(d.Get("tags").([]string))
	}
	if _, ok := d.GetOk("location"); ok {
		updateJobOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("status"); ok {
		status, err := ResourceIBMSchematicsJobMapToJobStatus(d.Get("status.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateJobOptions.SetStatus(status)
	}
	if _, ok := d.GetOk("data"); ok {
		data, err := ResourceIBMSchematicsJobMapToJobData(d.Get("data.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateJobOptions.SetData(data)
	}
	if _, ok := d.GetOk("bastion"); ok {
		bastion, err := ResourceIBMSchematicsJobMapToBastionResourceDefinition(d.Get("bastion.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateJobOptions.SetBastion(bastion)
	}
	if _, ok := d.GetOk("log_summary"); ok {
		logSummary, err := ResourceIBMSchematicsJobMapToJobLogSummary(d.Get("log_summary.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateJobOptions.SetLogSummary(logSummary)
	}

	_, response, err := schematicsClient.UpdateJobWithContext(context, updateJobOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateJobWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateJobWithContext failed %s\n%s", err, response))
	}

	return ResourceIBMSchematicsJobRead(context, d, meta)
}

func ResourceIBMSchematicsJobDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteJobOptions := &schematicsv1.DeleteJobOptions{}

	deleteJobOptions.SetJobID(d.Id())

	response, err := schematicsClient.DeleteJobWithContext(context, deleteJobOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteJobWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteJobWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMSchematicsJobMapToVariableData(modelMap map[string]interface{}) (*schematicsv1.VariableData, error) {
	model := &schematicsv1.VariableData{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["use_default"] != nil {
		model.UseDefault = core.BoolPtr(modelMap["use_default"].(bool))
	}
	if modelMap["metadata"] != nil && len(modelMap["metadata"].([]interface{})) > 0 {
		MetadataModel, err := ResourceIBMSchematicsJobMapToVariableMetadata(modelMap["metadata"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Metadata = MetadataModel
	}
	if modelMap["link"] != nil && modelMap["link"].(string) != "" {
		model.Link = core.StringPtr(modelMap["link"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToVariableMetadata(modelMap map[string]interface{}) (*schematicsv1.VariableMetadata, error) {
	model := &schematicsv1.VariableMetadata{}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["aliases"] != nil {
		aliases := []string{}
		for _, aliasesItem := range modelMap["aliases"].([]interface{}) {
			aliases = append(aliases, aliasesItem.(string))
		}
		model.Aliases = aliases
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["cloud_data_type"] != nil && modelMap["cloud_data_type"].(string) != "" {
		model.CloudDataType = core.StringPtr(modelMap["cloud_data_type"].(string))
	}
	if modelMap["default_value"] != nil && modelMap["default_value"].(string) != "" {
		model.DefaultValue = core.StringPtr(modelMap["default_value"].(string))
	}
	if modelMap["link_status"] != nil && modelMap["link_status"].(string) != "" {
		model.LinkStatus = core.StringPtr(modelMap["link_status"].(string))
	}
	if modelMap["secure"] != nil {
		model.Secure = core.BoolPtr(modelMap["secure"].(bool))
	}
	if modelMap["immutable"] != nil {
		model.Immutable = core.BoolPtr(modelMap["immutable"].(bool))
	}
	if modelMap["hidden"] != nil {
		model.Hidden = core.BoolPtr(modelMap["hidden"].(bool))
	}
	if modelMap["required"] != nil {
		model.Required = core.BoolPtr(modelMap["required"].(bool))
	}
	if modelMap["options"] != nil {
		options := []string{}
		for _, optionsItem := range modelMap["options"].([]interface{}) {
			options = append(options, optionsItem.(string))
		}
		model.Options = options
	}
	if modelMap["min_value"] != nil {
		model.MinValue = core.Int64Ptr(int64(modelMap["min_value"].(int)))
	}
	if modelMap["max_value"] != nil {
		model.MaxValue = core.Int64Ptr(int64(modelMap["max_value"].(int)))
	}
	if modelMap["min_length"] != nil {
		model.MinLength = core.Int64Ptr(int64(modelMap["min_length"].(int)))
	}
	if modelMap["max_length"] != nil {
		model.MaxLength = core.Int64Ptr(int64(modelMap["max_length"].(int)))
	}
	if modelMap["matches"] != nil && modelMap["matches"].(string) != "" {
		model.Matches = core.StringPtr(modelMap["matches"].(string))
	}
	if modelMap["position"] != nil {
		model.Position = core.Int64Ptr(int64(modelMap["position"].(int)))
	}
	if modelMap["group_by"] != nil && modelMap["group_by"].(string) != "" {
		model.GroupBy = core.StringPtr(modelMap["group_by"].(string))
	}
	if modelMap["source"] != nil && modelMap["source"].(string) != "" {
		model.Source = core.StringPtr(modelMap["source"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatus(modelMap map[string]interface{}) (*schematicsv1.JobStatus, error) {
	model := &schematicsv1.JobStatus{}
	if modelMap["position_in_queue"] != nil {
		model.PositionInQueue = core.Float64Ptr(modelMap["position_in_queue"].(float64))
	}
	if modelMap["total_in_queue"] != nil {
		model.TotalInQueue = core.Float64Ptr(modelMap["total_in_queue"].(float64))
	}
	if modelMap["workspace_job_status"] != nil && len(modelMap["workspace_job_status"].([]interface{})) > 0 {
		WorkspaceJobStatusModel, err := ResourceIBMSchematicsJobMapToJobStatusWorkspace(modelMap["workspace_job_status"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WorkspaceJobStatus = WorkspaceJobStatusModel
	}
	if modelMap["action_job_status"] != nil && len(modelMap["action_job_status"].([]interface{})) > 0 {
		ActionJobStatusModel, err := ResourceIBMSchematicsJobMapToJobStatusAction(modelMap["action_job_status"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActionJobStatus = ActionJobStatusModel
	}
	if modelMap["system_job_status"] != nil && len(modelMap["system_job_status"].([]interface{})) > 0 {
		SystemJobStatusModel, err := ResourceIBMSchematicsJobMapToJobStatusSystem(modelMap["system_job_status"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SystemJobStatus = SystemJobStatusModel
	}
	if modelMap["flow_job_status"] != nil && len(modelMap["flow_job_status"].([]interface{})) > 0 {
		FlowJobStatusModel, err := ResourceIBMSchematicsJobMapToJobStatusFlow(modelMap["flow_job_status"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FlowJobStatus = FlowJobStatusModel
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusWorkspace(modelMap map[string]interface{}) (*schematicsv1.JobStatusWorkspace, error) {
	model := &schematicsv1.JobStatusWorkspace{}
	if modelMap["workspace_name"] != nil && modelMap["workspace_name"].(string) != "" {
		model.WorkspaceName = core.StringPtr(modelMap["workspace_name"].(string))
	}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	if modelMap["flow_status"] != nil && len(modelMap["flow_status"].([]interface{})) > 0 {
		FlowStatusModel, err := ResourceIBMSchematicsJobMapToJobStatusFlow(modelMap["flow_status"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FlowStatus = FlowStatusModel
	}
	if modelMap["template_status"] != nil {
		templateStatus := []schematicsv1.JobStatusTemplate{}
		for _, templateStatusItem := range modelMap["template_status"].([]interface{}) {
			templateStatusItemModel, err := ResourceIBMSchematicsJobMapToJobStatusTemplate(templateStatusItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			templateStatus = append(templateStatus, *templateStatusItemModel)
		}
		model.TemplateStatus = templateStatus
	}
	if modelMap["updated_at"] != nil {

	}
	if modelMap["commands"] != nil {
		commands := []schematicsv1.CommandsInfo{}
		for _, commandsItem := range modelMap["commands"].([]interface{}) {
			commandsItemModel, err := ResourceIBMSchematicsJobMapToCommandsInfo(commandsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			commands = append(commands, *commandsItemModel)
		}
		model.Commands = commands
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusFlow(modelMap map[string]interface{}) (*schematicsv1.JobStatusFlow, error) {
	model := &schematicsv1.JobStatusFlow{}
	if modelMap["flow_id"] != nil && modelMap["flow_id"].(string) != "" {
		model.FlowID = core.StringPtr(modelMap["flow_id"].(string))
	}
	if modelMap["flow_name"] != nil && modelMap["flow_name"].(string) != "" {
		model.FlowName = core.StringPtr(modelMap["flow_name"].(string))
	}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	if modelMap["workitems"] != nil {
		workitems := []schematicsv1.JobStatusWorkitem{}
		for _, workitemsItem := range modelMap["workitems"].([]interface{}) {
			workitemsItemModel, err := ResourceIBMSchematicsJobMapToJobStatusWorkitem(workitemsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			workitems = append(workitems, *workitemsItemModel)
		}
		model.Workitems = workitems
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusWorkitem(modelMap map[string]interface{}) (*schematicsv1.JobStatusWorkitem, error) {
	model := &schematicsv1.JobStatusWorkitem{}
	if modelMap["workspace_id"] != nil && modelMap["workspace_id"].(string) != "" {
		model.WorkspaceID = core.StringPtr(modelMap["workspace_id"].(string))
	}
	if modelMap["workspace_name"] != nil && modelMap["workspace_name"].(string) != "" {
		model.WorkspaceName = core.StringPtr(modelMap["workspace_name"].(string))
	}
	if modelMap["job_id"] != nil && modelMap["job_id"].(string) != "" {
		model.JobID = core.StringPtr(modelMap["job_id"].(string))
	}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusTemplate(modelMap map[string]interface{}) (*schematicsv1.JobStatusTemplate, error) {
	model := &schematicsv1.JobStatusTemplate{}
	if modelMap["template_id"] != nil && modelMap["template_id"].(string) != "" {
		model.TemplateID = core.StringPtr(modelMap["template_id"].(string))
	}
	if modelMap["template_name"] != nil && modelMap["template_name"].(string) != "" {
		model.TemplateName = core.StringPtr(modelMap["template_name"].(string))
	}
	if modelMap["flow_index"] != nil {
		model.FlowIndex = core.Int64Ptr(int64(modelMap["flow_index"].(int)))
	}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToCommandsInfo(modelMap map[string]interface{}) (*schematicsv1.CommandsInfo, error) {
	model := &schematicsv1.CommandsInfo{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["outcome"] != nil && modelMap["outcome"].(string) != "" {
		model.Outcome = core.StringPtr(modelMap["outcome"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusAction(modelMap map[string]interface{}) (*schematicsv1.JobStatusAction, error) {
	model := &schematicsv1.JobStatusAction{}
	if modelMap["action_name"] != nil && modelMap["action_name"].(string) != "" {
		model.ActionName = core.StringPtr(modelMap["action_name"].(string))
	}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	if modelMap["bastion_status_code"] != nil && modelMap["bastion_status_code"].(string) != "" {
		model.BastionStatusCode = core.StringPtr(modelMap["bastion_status_code"].(string))
	}
	if modelMap["bastion_status_message"] != nil && modelMap["bastion_status_message"].(string) != "" {
		model.BastionStatusMessage = core.StringPtr(modelMap["bastion_status_message"].(string))
	}
	if modelMap["targets_status_code"] != nil && modelMap["targets_status_code"].(string) != "" {
		model.TargetsStatusCode = core.StringPtr(modelMap["targets_status_code"].(string))
	}
	if modelMap["targets_status_message"] != nil && modelMap["targets_status_message"].(string) != "" {
		model.TargetsStatusMessage = core.StringPtr(modelMap["targets_status_message"].(string))
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusSystem(modelMap map[string]interface{}) (*schematicsv1.JobStatusSystem, error) {
	model := &schematicsv1.JobStatusSystem{}
	if modelMap["system_status_message"] != nil && modelMap["system_status_message"].(string) != "" {
		model.SystemStatusMessage = core.StringPtr(modelMap["system_status_message"].(string))
	}
	if modelMap["system_status_code"] != nil && modelMap["system_status_code"].(string) != "" {
		model.SystemStatusCode = core.StringPtr(modelMap["system_status_code"].(string))
	}
	if modelMap["schematics_resource_status"] != nil {
		schematicsResourceStatus := []schematicsv1.JobStatusSchematicsResources{}
		for _, schematicsResourceStatusItem := range modelMap["schematics_resource_status"].([]interface{}) {
			schematicsResourceStatusItemModel, err := ResourceIBMSchematicsJobMapToJobStatusSchematicsResources(schematicsResourceStatusItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			schematicsResourceStatus = append(schematicsResourceStatus, *schematicsResourceStatusItemModel)
		}
		model.SchematicsResourceStatus = schematicsResourceStatus
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobStatusSchematicsResources(modelMap map[string]interface{}) (*schematicsv1.JobStatusSchematicsResources, error) {
	model := &schematicsv1.JobStatusSchematicsResources{}
	if modelMap["status_code"] != nil && modelMap["status_code"].(string) != "" {
		model.StatusCode = core.StringPtr(modelMap["status_code"].(string))
	}
	if modelMap["status_message"] != nil && modelMap["status_message"].(string) != "" {
		model.StatusMessage = core.StringPtr(modelMap["status_message"].(string))
	}
	if modelMap["schematics_resource_id"] != nil && modelMap["schematics_resource_id"].(string) != "" {
		model.SchematicsResourceID = core.StringPtr(modelMap["schematics_resource_id"].(string))
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobData(modelMap map[string]interface{}) (*schematicsv1.JobData, error) {
	model := &schematicsv1.JobData{}
	model.JobType = core.StringPtr(modelMap["job_type"].(string))
	if modelMap["workspace_job_data"] != nil && len(modelMap["workspace_job_data"].([]interface{})) > 0 {
		WorkspaceJobDataModel, err := ResourceIBMSchematicsJobMapToJobDataWorkspace(modelMap["workspace_job_data"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WorkspaceJobData = WorkspaceJobDataModel
	}
	if modelMap["action_job_data"] != nil && len(modelMap["action_job_data"].([]interface{})) > 0 {
		ActionJobDataModel, err := ResourceIBMSchematicsJobMapToJobDataAction(modelMap["action_job_data"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActionJobData = ActionJobDataModel
	}
	if modelMap["system_job_data"] != nil && len(modelMap["system_job_data"].([]interface{})) > 0 {
		SystemJobDataModel, err := ResourceIBMSchematicsJobMapToJobDataSystem(modelMap["system_job_data"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SystemJobData = SystemJobDataModel
	}
	if modelMap["flow_job_data"] != nil && len(modelMap["flow_job_data"].([]interface{})) > 0 {
		FlowJobDataModel, err := ResourceIBMSchematicsJobMapToJobDataFlow(modelMap["flow_job_data"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FlowJobData = FlowJobDataModel
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataWorkspace(modelMap map[string]interface{}) (*schematicsv1.JobDataWorkspace, error) {
	model := &schematicsv1.JobDataWorkspace{}
	if modelMap["workspace_name"] != nil && modelMap["workspace_name"].(string) != "" {
		model.WorkspaceName = core.StringPtr(modelMap["workspace_name"].(string))
	}
	if modelMap["flow_id"] != nil && modelMap["flow_id"].(string) != "" {
		model.FlowID = core.StringPtr(modelMap["flow_id"].(string))
	}
	if modelMap["flow_name"] != nil && modelMap["flow_name"].(string) != "" {
		model.FlowName = core.StringPtr(modelMap["flow_name"].(string))
	}
	if modelMap["inputs"] != nil {
		inputs := []schematicsv1.VariableData{}
		for _, inputsItem := range modelMap["inputs"].([]interface{}) {
			inputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(inputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			inputs = append(inputs, *inputsItemModel)
		}
		model.Inputs = inputs
	}
	if modelMap["outputs"] != nil {
		outputs := []schematicsv1.VariableData{}
		for _, outputsItem := range modelMap["outputs"].([]interface{}) {
			outputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(outputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			outputs = append(outputs, *outputsItemModel)
		}
		model.Outputs = outputs
	}
	if modelMap["settings"] != nil {
		settings := []schematicsv1.VariableData{}
		for _, settingsItem := range modelMap["settings"].([]interface{}) {
			settingsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(settingsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			settings = append(settings, *settingsItemModel)
		}
		model.Settings = settings
	}
	if modelMap["template_data"] != nil {
		templateData := []schematicsv1.JobDataTemplate{}
		for _, templateDataItem := range modelMap["template_data"].([]interface{}) {
			templateDataItemModel, err := ResourceIBMSchematicsJobMapToJobDataTemplate(templateDataItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			templateData = append(templateData, *templateDataItemModel)
		}
		model.TemplateData = templateData
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataTemplate(modelMap map[string]interface{}) (*schematicsv1.JobDataTemplate, error) {
	model := &schematicsv1.JobDataTemplate{}
	if modelMap["template_id"] != nil && modelMap["template_id"].(string) != "" {
		model.TemplateID = core.StringPtr(modelMap["template_id"].(string))
	}
	if modelMap["template_name"] != nil && modelMap["template_name"].(string) != "" {
		model.TemplateName = core.StringPtr(modelMap["template_name"].(string))
	}
	if modelMap["flow_index"] != nil {
		model.FlowIndex = core.Int64Ptr(int64(modelMap["flow_index"].(int)))
	}
	if modelMap["inputs"] != nil {
		inputs := []schematicsv1.VariableData{}
		for _, inputsItem := range modelMap["inputs"].([]interface{}) {
			inputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(inputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			inputs = append(inputs, *inputsItemModel)
		}
		model.Inputs = inputs
	}
	if modelMap["outputs"] != nil {
		outputs := []schematicsv1.VariableData{}
		for _, outputsItem := range modelMap["outputs"].([]interface{}) {
			outputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(outputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			outputs = append(outputs, *outputsItemModel)
		}
		model.Outputs = outputs
	}
	if modelMap["settings"] != nil {
		settings := []schematicsv1.VariableData{}
		for _, settingsItem := range modelMap["settings"].([]interface{}) {
			settingsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(settingsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			settings = append(settings, *settingsItemModel)
		}
		model.Settings = settings
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataAction(modelMap map[string]interface{}) (*schematicsv1.JobDataAction, error) {
	model := &schematicsv1.JobDataAction{}
	if modelMap["action_name"] != nil && modelMap["action_name"].(string) != "" {
		model.ActionName = core.StringPtr(modelMap["action_name"].(string))
	}
	if modelMap["inputs"] != nil {
		inputs := []schematicsv1.VariableData{}
		for _, inputsItem := range modelMap["inputs"].([]interface{}) {
			inputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(inputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			inputs = append(inputs, *inputsItemModel)
		}
		model.Inputs = inputs
	}
	if modelMap["outputs"] != nil {
		outputs := []schematicsv1.VariableData{}
		for _, outputsItem := range modelMap["outputs"].([]interface{}) {
			outputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(outputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			outputs = append(outputs, *outputsItemModel)
		}
		model.Outputs = outputs
	}
	if modelMap["settings"] != nil {
		settings := []schematicsv1.VariableData{}
		for _, settingsItem := range modelMap["settings"].([]interface{}) {
			settingsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(settingsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			settings = append(settings, *settingsItemModel)
		}
		model.Settings = settings
	}
	if modelMap["updated_at"] != nil {

	}
	if modelMap["inventory_record"] != nil && len(modelMap["inventory_record"].([]interface{})) > 0 {
		InventoryRecordModel, err := ResourceIBMSchematicsJobMapToInventoryResourceRecord(modelMap["inventory_record"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.InventoryRecord = InventoryRecordModel
	}
	if modelMap["materialized_inventory"] != nil && modelMap["materialized_inventory"].(string) != "" {
		model.MaterializedInventory = core.StringPtr(modelMap["materialized_inventory"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToInventoryResourceRecord(modelMap map[string]interface{}) (*schematicsv1.InventoryResourceRecord, error) {
	model := &schematicsv1.InventoryResourceRecord{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["location"] != nil && modelMap["location"].(string) != "" {
		model.Location = core.StringPtr(modelMap["location"].(string))
	}
	if modelMap["resource_group"] != nil && modelMap["resource_group"].(string) != "" {
		model.ResourceGroup = core.StringPtr(modelMap["resource_group"].(string))
	}
	if modelMap["created_at"] != nil {

	}
	if modelMap["created_by"] != nil && modelMap["created_by"].(string) != "" {
		model.CreatedBy = core.StringPtr(modelMap["created_by"].(string))
	}
	if modelMap["updated_at"] != nil {

	}
	if modelMap["updated_by"] != nil && modelMap["updated_by"].(string) != "" {
		model.UpdatedBy = core.StringPtr(modelMap["updated_by"].(string))
	}
	if modelMap["inventories_ini"] != nil && modelMap["inventories_ini"].(string) != "" {
		model.InventoriesIni = core.StringPtr(modelMap["inventories_ini"].(string))
	}
	if modelMap["resource_queries"] != nil {
		resourceQueries := []string{}
		for _, resourceQueriesItem := range modelMap["resource_queries"].([]interface{}) {
			resourceQueries = append(resourceQueries, resourceQueriesItem.(string))
		}
		model.ResourceQueries = resourceQueries
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataSystem(modelMap map[string]interface{}) (*schematicsv1.JobDataSystem, error) {
	model := &schematicsv1.JobDataSystem{}
	if modelMap["key_id"] != nil && modelMap["key_id"].(string) != "" {
		model.KeyID = core.StringPtr(modelMap["key_id"].(string))
	}
	if modelMap["schematics_resource_id"] != nil {
		schematicsResourceID := []string{}
		for _, schematicsResourceIDItem := range modelMap["schematics_resource_id"].([]interface{}) {
			schematicsResourceID = append(schematicsResourceID, schematicsResourceIDItem.(string))
		}
		model.SchematicsResourceID = schematicsResourceID
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataFlow(modelMap map[string]interface{}) (*schematicsv1.JobDataFlow, error) {
	model := &schematicsv1.JobDataFlow{}
	if modelMap["flow_id"] != nil && modelMap["flow_id"].(string) != "" {
		model.FlowID = core.StringPtr(modelMap["flow_id"].(string))
	}
	if modelMap["flow_name"] != nil && modelMap["flow_name"].(string) != "" {
		model.FlowName = core.StringPtr(modelMap["flow_name"].(string))
	}
	if modelMap["workitems"] != nil {
		workitems := []schematicsv1.JobDataWorkItem{}
		for _, workitemsItem := range modelMap["workitems"].([]interface{}) {
			workitemsItemModel, err := ResourceIBMSchematicsJobMapToJobDataWorkItem(workitemsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			workitems = append(workitems, *workitemsItemModel)
		}
		model.Workitems = workitems
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataWorkItem(modelMap map[string]interface{}) (*schematicsv1.JobDataWorkItem, error) {
	model := &schematicsv1.JobDataWorkItem{}
	if modelMap["command_object_id"] != nil && modelMap["command_object_id"].(string) != "" {
		model.CommandObjectID = core.StringPtr(modelMap["command_object_id"].(string))
	}
	if modelMap["command_object_name"] != nil && modelMap["command_object_name"].(string) != "" {
		model.CommandObjectName = core.StringPtr(modelMap["command_object_name"].(string))
	}
	if modelMap["layers"] != nil && modelMap["layers"].(string) != "" {
		model.Layers = core.StringPtr(modelMap["layers"].(string))
	}
	if modelMap["source_type"] != nil && modelMap["source_type"].(string) != "" {
		model.SourceType = core.StringPtr(modelMap["source_type"].(string))
	}
	if modelMap["source"] != nil && len(modelMap["source"].([]interface{})) > 0 {
		SourceModel, err := ResourceIBMSchematicsJobMapToExternalSource(modelMap["source"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Source = SourceModel
	}
	if modelMap["inputs"] != nil {
		inputs := []schematicsv1.VariableData{}
		for _, inputsItem := range modelMap["inputs"].([]interface{}) {
			inputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(inputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			inputs = append(inputs, *inputsItemModel)
		}
		model.Inputs = inputs
	}
	if modelMap["outputs"] != nil {
		outputs := []schematicsv1.VariableData{}
		for _, outputsItem := range modelMap["outputs"].([]interface{}) {
			outputsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(outputsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			outputs = append(outputs, *outputsItemModel)
		}
		model.Outputs = outputs
	}
	if modelMap["settings"] != nil {
		settings := []schematicsv1.VariableData{}
		for _, settingsItem := range modelMap["settings"].([]interface{}) {
			settingsItemModel, err := ResourceIBMSchematicsJobMapToVariableData(settingsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			settings = append(settings, *settingsItemModel)
		}
		model.Settings = settings
	}
	if modelMap["last_job"] != nil && len(modelMap["last_job"].([]interface{})) > 0 {
		LastJobModel, err := ResourceIBMSchematicsJobMapToJobDataWorkItemLastJob(modelMap["last_job"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.LastJob = LastJobModel
	}
	if modelMap["updated_at"] != nil {

	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToExternalSource(modelMap map[string]interface{}) (*schematicsv1.ExternalSource, error) {
	model := &schematicsv1.ExternalSource{}
	model.SourceType = core.StringPtr(modelMap["source_type"].(string))
	if modelMap["git"] != nil && len(modelMap["git"].([]interface{})) > 0 {
		GitModel, err := ResourceIBMSchematicsJobMapToExternalSourceGit(modelMap["git"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Git = GitModel
	}
	if modelMap["catalog"] != nil && len(modelMap["catalog"].([]interface{})) > 0 {
		CatalogModel, err := ResourceIBMSchematicsJobMapToExternalSourceCatalog(modelMap["catalog"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Catalog = CatalogModel
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToExternalSourceGit(modelMap map[string]interface{}) (*schematicsv1.ExternalSourceGit, error) {
	model := &schematicsv1.ExternalSourceGit{}
	if modelMap["computed_git_repo_url"] != nil && modelMap["computed_git_repo_url"].(string) != "" {
		model.ComputedGitRepoURL = core.StringPtr(modelMap["computed_git_repo_url"].(string))
	}
	if modelMap["git_repo_url"] != nil && modelMap["git_repo_url"].(string) != "" {
		model.GitRepoURL = core.StringPtr(modelMap["git_repo_url"].(string))
	}
	if modelMap["git_token"] != nil && modelMap["git_token"].(string) != "" {
		model.GitToken = core.StringPtr(modelMap["git_token"].(string))
	}
	if modelMap["git_repo_folder"] != nil && modelMap["git_repo_folder"].(string) != "" {
		model.GitRepoFolder = core.StringPtr(modelMap["git_repo_folder"].(string))
	}
	if modelMap["git_release"] != nil && modelMap["git_release"].(string) != "" {
		model.GitRelease = core.StringPtr(modelMap["git_release"].(string))
	}
	if modelMap["git_branch"] != nil && modelMap["git_branch"].(string) != "" {
		model.GitBranch = core.StringPtr(modelMap["git_branch"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToExternalSourceCatalog(modelMap map[string]interface{}) (*schematicsv1.ExternalSourceCatalog, error) {
	model := &schematicsv1.ExternalSourceCatalog{}
	if modelMap["catalog_name"] != nil && modelMap["catalog_name"].(string) != "" {
		model.CatalogName = core.StringPtr(modelMap["catalog_name"].(string))
	}
	if modelMap["offering_name"] != nil && modelMap["offering_name"].(string) != "" {
		model.OfferingName = core.StringPtr(modelMap["offering_name"].(string))
	}
	if modelMap["offering_version"] != nil && modelMap["offering_version"].(string) != "" {
		model.OfferingVersion = core.StringPtr(modelMap["offering_version"].(string))
	}
	if modelMap["offering_kind"] != nil && modelMap["offering_kind"].(string) != "" {
		model.OfferingKind = core.StringPtr(modelMap["offering_kind"].(string))
	}
	if modelMap["offering_id"] != nil && modelMap["offering_id"].(string) != "" {
		model.OfferingID = core.StringPtr(modelMap["offering_id"].(string))
	}
	if modelMap["offering_version_id"] != nil && modelMap["offering_version_id"].(string) != "" {
		model.OfferingVersionID = core.StringPtr(modelMap["offering_version_id"].(string))
	}
	if modelMap["offering_repo_url"] != nil && modelMap["offering_repo_url"].(string) != "" {
		model.OfferingRepoURL = core.StringPtr(modelMap["offering_repo_url"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobDataWorkItemLastJob(modelMap map[string]interface{}) (*schematicsv1.JobDataWorkItemLastJob, error) {
	model := &schematicsv1.JobDataWorkItemLastJob{}
	if modelMap["command_object"] != nil && modelMap["command_object"].(string) != "" {
		model.CommandObject = core.StringPtr(modelMap["command_object"].(string))
	}
	if modelMap["command_object_name"] != nil && modelMap["command_object_name"].(string) != "" {
		model.CommandObjectName = core.StringPtr(modelMap["command_object_name"].(string))
	}
	if modelMap["command_object_id"] != nil && modelMap["command_object_id"].(string) != "" {
		model.CommandObjectID = core.StringPtr(modelMap["command_object_id"].(string))
	}
	if modelMap["command_name"] != nil && modelMap["command_name"].(string) != "" {
		model.CommandName = core.StringPtr(modelMap["command_name"].(string))
	}
	if modelMap["job_id"] != nil && modelMap["job_id"].(string) != "" {
		model.JobID = core.StringPtr(modelMap["job_id"].(string))
	}
	if modelMap["job_status"] != nil && modelMap["job_status"].(string) != "" {
		model.JobStatus = core.StringPtr(modelMap["job_status"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToBastionResourceDefinition(modelMap map[string]interface{}) (*schematicsv1.BastionResourceDefinition, error) {
	model := &schematicsv1.BastionResourceDefinition{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["host"] != nil && modelMap["host"].(string) != "" {
		model.Host = core.StringPtr(modelMap["host"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummary(modelMap map[string]interface{}) (*schematicsv1.JobLogSummary, error) {
	model := &schematicsv1.JobLogSummary{}
	if modelMap["job_id"] != nil && modelMap["job_id"].(string) != "" {
		model.JobID = core.StringPtr(modelMap["job_id"].(string))
	}
	if modelMap["job_type"] != nil && modelMap["job_type"].(string) != "" {
		model.JobType = core.StringPtr(modelMap["job_type"].(string))
	}
	if modelMap["log_start_at"] != nil {

	}
	if modelMap["log_analyzed_till"] != nil {

	}
	if modelMap["elapsed_time"] != nil {
		model.ElapsedTime = core.Float64Ptr(modelMap["elapsed_time"].(float64))
	}
	if modelMap["log_errors"] != nil {
		logErrors := []schematicsv1.JobLogSummaryLogErrors{}
		for _, logErrorsItem := range modelMap["log_errors"].([]interface{}) {
			logErrorsItemModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryLogErrors(logErrorsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			logErrors = append(logErrors, *logErrorsItemModel)
		}
		model.LogErrors = logErrors
	}
	if modelMap["repo_download_job"] != nil && len(modelMap["repo_download_job"].([]interface{})) > 0 {
		RepoDownloadJobModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryRepoDownloadJob(modelMap["repo_download_job"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RepoDownloadJob = RepoDownloadJobModel
	}
	if modelMap["workspace_job"] != nil && len(modelMap["workspace_job"].([]interface{})) > 0 {
		WorkspaceJobModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryWorkspaceJob(modelMap["workspace_job"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.WorkspaceJob = WorkspaceJobModel
	}
	if modelMap["flow_job"] != nil && len(modelMap["flow_job"].([]interface{})) > 0 {
		FlowJobModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryFlowJob(modelMap["flow_job"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.FlowJob = FlowJobModel
	}
	if modelMap["action_job"] != nil && len(modelMap["action_job"].([]interface{})) > 0 {
		ActionJobModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryActionJob(modelMap["action_job"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ActionJob = ActionJobModel
	}
	if modelMap["system_job"] != nil && len(modelMap["system_job"].([]interface{})) > 0 {
		SystemJobModel, err := ResourceIBMSchematicsJobMapToJobLogSummarySystemJob(modelMap["system_job"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SystemJob = SystemJobModel
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryLogErrors(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryLogErrors, error) {
	model := &schematicsv1.JobLogSummaryLogErrors{}
	if modelMap["error_code"] != nil && modelMap["error_code"].(string) != "" {
		model.ErrorCode = core.StringPtr(modelMap["error_code"].(string))
	}
	if modelMap["error_msg"] != nil && modelMap["error_msg"].(string) != "" {
		model.ErrorMsg = core.StringPtr(modelMap["error_msg"].(string))
	}
	if modelMap["error_count"] != nil {
		model.ErrorCount = core.Float64Ptr(modelMap["error_count"].(float64))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryRepoDownloadJob(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryRepoDownloadJob, error) {
	model := &schematicsv1.JobLogSummaryRepoDownloadJob{}
	if modelMap["scanned_file_count"] != nil {
		model.ScannedFileCount = core.Float64Ptr(modelMap["scanned_file_count"].(float64))
	}
	if modelMap["quarantined_file_count"] != nil {
		model.QuarantinedFileCount = core.Float64Ptr(modelMap["quarantined_file_count"].(float64))
	}
	if modelMap["detected_filetype"] != nil && modelMap["detected_filetype"].(string) != "" {
		model.DetectedFiletype = core.StringPtr(modelMap["detected_filetype"].(string))
	}
	if modelMap["inputs_count"] != nil && modelMap["inputs_count"].(string) != "" {
		model.InputsCount = core.StringPtr(modelMap["inputs_count"].(string))
	}
	if modelMap["outputs_count"] != nil && modelMap["outputs_count"].(string) != "" {
		model.OutputsCount = core.StringPtr(modelMap["outputs_count"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryWorkspaceJob(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryWorkspaceJob, error) {
	model := &schematicsv1.JobLogSummaryWorkspaceJob{}
	if modelMap["resources_add"] != nil {
		model.ResourcesAdd = core.Float64Ptr(modelMap["resources_add"].(float64))
	}
	if modelMap["resources_modify"] != nil {
		model.ResourcesModify = core.Float64Ptr(modelMap["resources_modify"].(float64))
	}
	if modelMap["resources_destroy"] != nil {
		model.ResourcesDestroy = core.Float64Ptr(modelMap["resources_destroy"].(float64))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryFlowJob(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryFlowJob, error) {
	model := &schematicsv1.JobLogSummaryFlowJob{}
	if modelMap["workitems_completed"] != nil {
		model.WorkitemsCompleted = core.Float64Ptr(modelMap["workitems_completed"].(float64))
	}
	if modelMap["workitems_pending"] != nil {
		model.WorkitemsPending = core.Float64Ptr(modelMap["workitems_pending"].(float64))
	}
	if modelMap["workitems_failed"] != nil {
		model.WorkitemsFailed = core.Float64Ptr(modelMap["workitems_failed"].(float64))
	}
	if modelMap["workitems"] != nil {
		workitems := []schematicsv1.JobLogSummaryWorkitems{}
		for _, workitemsItem := range modelMap["workitems"].([]interface{}) {
			workitemsItemModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryWorkitems(workitemsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			workitems = append(workitems, *workitemsItemModel)
		}
		model.Workitems = workitems
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryWorkitems(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryWorkitems, error) {
	model := &schematicsv1.JobLogSummaryWorkitems{}
	if modelMap["workspace_id"] != nil && modelMap["workspace_id"].(string) != "" {
		model.WorkspaceID = core.StringPtr(modelMap["workspace_id"].(string))
	}
	if modelMap["job_id"] != nil && modelMap["job_id"].(string) != "" {
		model.JobID = core.StringPtr(modelMap["job_id"].(string))
	}
	if modelMap["resources_add"] != nil {
		model.ResourcesAdd = core.Float64Ptr(modelMap["resources_add"].(float64))
	}
	if modelMap["resources_modify"] != nil {
		model.ResourcesModify = core.Float64Ptr(modelMap["resources_modify"].(float64))
	}
	if modelMap["resources_destroy"] != nil {
		model.ResourcesDestroy = core.Float64Ptr(modelMap["resources_destroy"].(float64))
	}
	if modelMap["log_url"] != nil && modelMap["log_url"].(string) != "" {
		model.LogURL = core.StringPtr(modelMap["log_url"].(string))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryActionJob(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryActionJob, error) {
	model := &schematicsv1.JobLogSummaryActionJob{}
	if modelMap["target_count"] != nil {
		model.TargetCount = core.Float64Ptr(modelMap["target_count"].(float64))
	}
	if modelMap["task_count"] != nil {
		model.TaskCount = core.Float64Ptr(modelMap["task_count"].(float64))
	}
	if modelMap["play_count"] != nil {
		model.PlayCount = core.Float64Ptr(modelMap["play_count"].(float64))
	}
	if modelMap["recap"] != nil && len(modelMap["recap"].([]interface{})) > 0 {
		RecapModel, err := ResourceIBMSchematicsJobMapToJobLogSummaryActionJobRecap(modelMap["recap"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Recap = RecapModel
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummaryActionJobRecap(modelMap map[string]interface{}) (*schematicsv1.JobLogSummaryActionJobRecap, error) {
	model := &schematicsv1.JobLogSummaryActionJobRecap{}
	if modelMap["target"] != nil {
		target := []string{}
		for _, targetItem := range modelMap["target"].([]interface{}) {
			target = append(target, targetItem.(string))
		}
		model.Target = target
	}
	if modelMap["ok"] != nil {
		model.Ok = core.Float64Ptr(modelMap["ok"].(float64))
	}
	if modelMap["changed"] != nil {
		model.Changed = core.Float64Ptr(modelMap["changed"].(float64))
	}
	if modelMap["failed"] != nil {
		model.Failed = core.Float64Ptr(modelMap["failed"].(float64))
	}
	if modelMap["skipped"] != nil {
		model.Skipped = core.Float64Ptr(modelMap["skipped"].(float64))
	}
	if modelMap["unreachable"] != nil {
		model.Unreachable = core.Float64Ptr(modelMap["unreachable"].(float64))
	}
	return model, nil
}

func ResourceIBMSchematicsJobMapToJobLogSummarySystemJob(modelMap map[string]interface{}) (*schematicsv1.JobLogSummarySystemJob, error) {
	model := &schematicsv1.JobLogSummarySystemJob{}
	if modelMap["target_count"] != nil {
		model.TargetCount = core.Float64Ptr(modelMap["target_count"].(float64))
	}
	if modelMap["success"] != nil {
		model.Success = core.Float64Ptr(modelMap["success"].(float64))
	}
	if modelMap["failed"] != nil {
		model.Failed = core.Float64Ptr(modelMap["failed"].(float64))
	}
	return model, nil
}

func ResourceIBMSchematicsJobVariableDataToMap(model *schematicsv1.VariableData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = model.UseDefault
	}
	if model.Metadata != nil {
		metadataMap, err := ResourceIBMSchematicsJobVariableMetadataToMap(model.Metadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["metadata"] = []map[string]interface{}{metadataMap}
	}
	if model.Link != nil {
		modelMap["link"] = model.Link
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobVariableMetadataToMap(model *schematicsv1.VariableMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.Aliases != nil {
		modelMap["aliases"] = model.Aliases
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.CloudDataType != nil {
		modelMap["cloud_data_type"] = model.CloudDataType
	}
	if model.DefaultValue != nil {
		modelMap["default_value"] = model.DefaultValue
	}
	if model.LinkStatus != nil {
		modelMap["link_status"] = model.LinkStatus
	}
	if model.Secure != nil {
		modelMap["secure"] = model.Secure
	}
	if model.Immutable != nil {
		modelMap["immutable"] = model.Immutable
	}
	if model.Hidden != nil {
		modelMap["hidden"] = model.Hidden
	}
	if model.Required != nil {
		modelMap["required"] = model.Required
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	if model.MinValue != nil {
		modelMap["min_value"] = flex.IntValue(model.MinValue)
	}
	if model.MaxValue != nil {
		modelMap["max_value"] = flex.IntValue(model.MaxValue)
	}
	if model.MinLength != nil {
		modelMap["min_length"] = flex.IntValue(model.MinLength)
	}
	if model.MaxLength != nil {
		modelMap["max_length"] = flex.IntValue(model.MaxLength)
	}
	if model.Matches != nil {
		modelMap["matches"] = model.Matches
	}
	if model.Position != nil {
		modelMap["position"] = flex.IntValue(model.Position)
	}
	if model.GroupBy != nil {
		modelMap["group_by"] = model.GroupBy
	}
	if model.Source != nil {
		modelMap["source"] = model.Source
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusToMap(model *schematicsv1.JobStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PositionInQueue != nil {
		modelMap["position_in_queue"] = model.PositionInQueue
	}
	if model.TotalInQueue != nil {
		modelMap["total_in_queue"] = model.TotalInQueue
	}
	if model.WorkspaceJobStatus != nil {
		workspaceJobStatusMap, err := ResourceIBMSchematicsJobJobStatusWorkspaceToMap(model.WorkspaceJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["workspace_job_status"] = []map[string]interface{}{workspaceJobStatusMap}
	}
	if model.ActionJobStatus != nil {
		actionJobStatusMap, err := ResourceIBMSchematicsJobJobStatusActionToMap(model.ActionJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_job_status"] = []map[string]interface{}{actionJobStatusMap}
	}
	if model.SystemJobStatus != nil {
		systemJobStatusMap, err := ResourceIBMSchematicsJobJobStatusSystemToMap(model.SystemJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_job_status"] = []map[string]interface{}{systemJobStatusMap}
	}
	if model.FlowJobStatus != nil {
		flowJobStatusMap, err := ResourceIBMSchematicsJobJobStatusFlowToMap(model.FlowJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_job_status"] = []map[string]interface{}{flowJobStatusMap}
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusWorkspaceToMap(model *schematicsv1.JobStatusWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = model.WorkspaceName
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.FlowStatus != nil {
		flowStatusMap, err := ResourceIBMSchematicsJobJobStatusFlowToMap(model.FlowStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_status"] = []map[string]interface{}{flowStatusMap}
	}
	if model.TemplateStatus != nil {
		templateStatus := []map[string]interface{}{}
		for _, templateStatusItem := range model.TemplateStatus {
			templateStatusItemMap, err := ResourceIBMSchematicsJobJobStatusTemplateToMap(&templateStatusItem)
			if err != nil {
				return modelMap, err
			}
			templateStatus = append(templateStatus, templateStatusItemMap)
		}
		modelMap["template_status"] = templateStatus
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.Commands != nil {
		commands := []map[string]interface{}{}
		for _, commandsItem := range model.Commands {
			commandsItemMap, err := ResourceIBMSchematicsJobCommandsInfoToMap(&commandsItem)
			if err != nil {
				return modelMap, err
			}
			commands = append(commands, commandsItemMap)
		}
		modelMap["commands"] = commands
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusFlowToMap(model *schematicsv1.JobStatusFlow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FlowID != nil {
		modelMap["flow_id"] = model.FlowID
	}
	if model.FlowName != nil {
		modelMap["flow_name"] = model.FlowName
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.Workitems != nil {
		workitems := []map[string]interface{}{}
		for _, workitemsItem := range model.Workitems {
			workitemsItemMap, err := ResourceIBMSchematicsJobJobStatusWorkitemToMap(&workitemsItem)
			if err != nil {
				return modelMap, err
			}
			workitems = append(workitems, workitemsItemMap)
		}
		modelMap["workitems"] = workitems
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusWorkitemToMap(model *schematicsv1.JobStatusWorkitem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = model.WorkspaceID
	}
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = model.WorkspaceName
	}
	if model.JobID != nil {
		modelMap["job_id"] = model.JobID
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusTemplateToMap(model *schematicsv1.JobStatusTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TemplateID != nil {
		modelMap["template_id"] = model.TemplateID
	}
	if model.TemplateName != nil {
		modelMap["template_name"] = model.TemplateName
	}
	if model.FlowIndex != nil {
		modelMap["flow_index"] = flex.IntValue(model.FlowIndex)
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobCommandsInfoToMap(model *schematicsv1.CommandsInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Outcome != nil {
		modelMap["outcome"] = model.Outcome
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusActionToMap(model *schematicsv1.JobStatusAction) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActionName != nil {
		modelMap["action_name"] = model.ActionName
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.BastionStatusCode != nil {
		modelMap["bastion_status_code"] = model.BastionStatusCode
	}
	if model.BastionStatusMessage != nil {
		modelMap["bastion_status_message"] = model.BastionStatusMessage
	}
	if model.TargetsStatusCode != nil {
		modelMap["targets_status_code"] = model.TargetsStatusCode
	}
	if model.TargetsStatusMessage != nil {
		modelMap["targets_status_message"] = model.TargetsStatusMessage
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusSystemToMap(model *schematicsv1.JobStatusSystem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SystemStatusMessage != nil {
		modelMap["system_status_message"] = model.SystemStatusMessage
	}
	if model.SystemStatusCode != nil {
		modelMap["system_status_code"] = model.SystemStatusCode
	}
	if model.SchematicsResourceStatus != nil {
		schematicsResourceStatus := []map[string]interface{}{}
		for _, schematicsResourceStatusItem := range model.SchematicsResourceStatus {
			schematicsResourceStatusItemMap, err := ResourceIBMSchematicsJobJobStatusSchematicsResourcesToMap(&schematicsResourceStatusItem)
			if err != nil {
				return modelMap, err
			}
			schematicsResourceStatus = append(schematicsResourceStatus, schematicsResourceStatusItemMap)
		}
		modelMap["schematics_resource_status"] = schematicsResourceStatus
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobStatusSchematicsResourcesToMap(model *schematicsv1.JobStatusSchematicsResources) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = model.StatusMessage
	}
	if model.SchematicsResourceID != nil {
		modelMap["schematics_resource_id"] = model.SchematicsResourceID
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataToMap(model *schematicsv1.JobData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["job_type"] = model.JobType
	if model.WorkspaceJobData != nil {
		workspaceJobDataMap, err := ResourceIBMSchematicsJobJobDataWorkspaceToMap(model.WorkspaceJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["workspace_job_data"] = []map[string]interface{}{workspaceJobDataMap}
	}
	if model.ActionJobData != nil {
		actionJobDataMap, err := ResourceIBMSchematicsJobJobDataActionToMap(model.ActionJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_job_data"] = []map[string]interface{}{actionJobDataMap}
	}
	if model.SystemJobData != nil {
		systemJobDataMap, err := ResourceIBMSchematicsJobJobDataSystemToMap(model.SystemJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_job_data"] = []map[string]interface{}{systemJobDataMap}
	}
	if model.FlowJobData != nil {
		flowJobDataMap, err := ResourceIBMSchematicsJobJobDataFlowToMap(model.FlowJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_job_data"] = []map[string]interface{}{flowJobDataMap}
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataWorkspaceToMap(model *schematicsv1.JobDataWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = model.WorkspaceName
	}
	if model.FlowID != nil {
		modelMap["flow_id"] = model.FlowID
	}
	if model.FlowName != nil {
		modelMap["flow_name"] = model.FlowName
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&inputsItem)
			if err != nil {
				return modelMap, err
			}
			inputs = append(inputs, inputsItemMap)
		}
		modelMap["inputs"] = inputs
	}
	if model.Outputs != nil {
		outputs := []map[string]interface{}{}
		for _, outputsItem := range model.Outputs {
			outputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&outputsItem)
			if err != nil {
				return modelMap, err
			}
			outputs = append(outputs, outputsItemMap)
		}
		modelMap["outputs"] = outputs
	}
	if model.Settings != nil {
		settings := []map[string]interface{}{}
		for _, settingsItem := range model.Settings {
			settingsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&settingsItem)
			if err != nil {
				return modelMap, err
			}
			settings = append(settings, settingsItemMap)
		}
		modelMap["settings"] = settings
	}
	if model.TemplateData != nil {
		templateData := []map[string]interface{}{}
		for _, templateDataItem := range model.TemplateData {
			templateDataItemMap, err := ResourceIBMSchematicsJobJobDataTemplateToMap(&templateDataItem)
			if err != nil {
				return modelMap, err
			}
			templateData = append(templateData, templateDataItemMap)
		}
		modelMap["template_data"] = templateData
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataTemplateToMap(model *schematicsv1.JobDataTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TemplateID != nil {
		modelMap["template_id"] = model.TemplateID
	}
	if model.TemplateName != nil {
		modelMap["template_name"] = model.TemplateName
	}
	if model.FlowIndex != nil {
		modelMap["flow_index"] = flex.IntValue(model.FlowIndex)
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&inputsItem)
			if err != nil {
				return modelMap, err
			}
			inputs = append(inputs, inputsItemMap)
		}
		modelMap["inputs"] = inputs
	}
	if model.Outputs != nil {
		outputs := []map[string]interface{}{}
		for _, outputsItem := range model.Outputs {
			outputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&outputsItem)
			if err != nil {
				return modelMap, err
			}
			outputs = append(outputs, outputsItemMap)
		}
		modelMap["outputs"] = outputs
	}
	if model.Settings != nil {
		settings := []map[string]interface{}{}
		for _, settingsItem := range model.Settings {
			settingsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&settingsItem)
			if err != nil {
				return modelMap, err
			}
			settings = append(settings, settingsItemMap)
		}
		modelMap["settings"] = settings
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataActionToMap(model *schematicsv1.JobDataAction) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActionName != nil {
		modelMap["action_name"] = model.ActionName
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&inputsItem)
			if err != nil {
				return modelMap, err
			}
			inputs = append(inputs, inputsItemMap)
		}
		modelMap["inputs"] = inputs
	}
	if model.Outputs != nil {
		outputs := []map[string]interface{}{}
		for _, outputsItem := range model.Outputs {
			outputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&outputsItem)
			if err != nil {
				return modelMap, err
			}
			outputs = append(outputs, outputsItemMap)
		}
		modelMap["outputs"] = outputs
	}
	if model.Settings != nil {
		settings := []map[string]interface{}{}
		for _, settingsItem := range model.Settings {
			settingsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&settingsItem)
			if err != nil {
				return modelMap, err
			}
			settings = append(settings, settingsItemMap)
		}
		modelMap["settings"] = settings
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.InventoryRecord != nil {
		inventoryRecordMap, err := ResourceIBMSchematicsJobInventoryResourceRecordToMap(model.InventoryRecord)
		if err != nil {
			return modelMap, err
		}
		modelMap["inventory_record"] = []map[string]interface{}{inventoryRecordMap}
	}
	if model.MaterializedInventory != nil {
		modelMap["materialized_inventory"] = model.MaterializedInventory
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobInventoryResourceRecordToMap(model *schematicsv1.InventoryResourceRecord) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Location != nil {
		modelMap["location"] = model.Location
	}
	if model.ResourceGroup != nil {
		modelMap["resource_group"] = model.ResourceGroup
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = model.CreatedBy
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.UpdatedBy != nil {
		modelMap["updated_by"] = model.UpdatedBy
	}
	if model.InventoriesIni != nil {
		modelMap["inventories_ini"] = model.InventoriesIni
	}
	if model.ResourceQueries != nil {
		modelMap["resource_queries"] = model.ResourceQueries
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataSystemToMap(model *schematicsv1.JobDataSystem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.KeyID != nil {
		modelMap["key_id"] = model.KeyID
	}
	if model.SchematicsResourceID != nil {
		modelMap["schematics_resource_id"] = model.SchematicsResourceID
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataFlowToMap(model *schematicsv1.JobDataFlow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FlowID != nil {
		modelMap["flow_id"] = model.FlowID
	}
	if model.FlowName != nil {
		modelMap["flow_name"] = model.FlowName
	}
	if model.Workitems != nil {
		workitems := []map[string]interface{}{}
		for _, workitemsItem := range model.Workitems {
			workitemsItemMap, err := ResourceIBMSchematicsJobJobDataWorkItemToMap(&workitemsItem)
			if err != nil {
				return modelMap, err
			}
			workitems = append(workitems, workitemsItemMap)
		}
		modelMap["workitems"] = workitems
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataWorkItemToMap(model *schematicsv1.JobDataWorkItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CommandObjectID != nil {
		modelMap["command_object_id"] = model.CommandObjectID
	}
	if model.CommandObjectName != nil {
		modelMap["command_object_name"] = model.CommandObjectName
	}
	if model.Layers != nil {
		modelMap["layers"] = model.Layers
	}
	if model.SourceType != nil {
		modelMap["source_type"] = model.SourceType
	}
	if model.Source != nil {
		sourceMap, err := ResourceIBMSchematicsJobExternalSourceToMap(model.Source)
		if err != nil {
			return modelMap, err
		}
		modelMap["source"] = []map[string]interface{}{sourceMap}
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&inputsItem)
			if err != nil {
				return modelMap, err
			}
			inputs = append(inputs, inputsItemMap)
		}
		modelMap["inputs"] = inputs
	}
	if model.Outputs != nil {
		outputs := []map[string]interface{}{}
		for _, outputsItem := range model.Outputs {
			outputsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&outputsItem)
			if err != nil {
				return modelMap, err
			}
			outputs = append(outputs, outputsItemMap)
		}
		modelMap["outputs"] = outputs
	}
	if model.Settings != nil {
		settings := []map[string]interface{}{}
		for _, settingsItem := range model.Settings {
			settingsItemMap, err := ResourceIBMSchematicsJobVariableDataToMap(&settingsItem)
			if err != nil {
				return modelMap, err
			}
			settings = append(settings, settingsItemMap)
		}
		modelMap["settings"] = settings
	}
	if model.LastJob != nil {
		lastJobMap, err := ResourceIBMSchematicsJobJobDataWorkItemLastJobToMap(model.LastJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_job"] = []map[string]interface{}{lastJobMap}
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobExternalSourceToMap(model *schematicsv1.ExternalSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_type"] = model.SourceType
	if model.Git != nil {
		gitMap, err := ResourceIBMSchematicsJobExternalSourceGitToMap(model.Git)
		if err != nil {
			return modelMap, err
		}
		modelMap["git"] = []map[string]interface{}{gitMap}
	}
	if model.Catalog != nil {
		catalogMap, err := ResourceIBMSchematicsJobExternalSourceCatalogToMap(model.Catalog)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog"] = []map[string]interface{}{catalogMap}
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobExternalSourceGitToMap(model *schematicsv1.ExternalSourceGit) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ComputedGitRepoURL != nil {
		modelMap["computed_git_repo_url"] = model.ComputedGitRepoURL
	}
	if model.GitRepoURL != nil {
		modelMap["git_repo_url"] = model.GitRepoURL
	}
	if model.GitToken != nil {
		modelMap["git_token"] = model.GitToken
	}
	if model.GitRepoFolder != nil {
		modelMap["git_repo_folder"] = model.GitRepoFolder
	}
	if model.GitRelease != nil {
		modelMap["git_release"] = model.GitRelease
	}
	if model.GitBranch != nil {
		modelMap["git_branch"] = model.GitBranch
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobExternalSourceCatalogToMap(model *schematicsv1.ExternalSourceCatalog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogName != nil {
		modelMap["catalog_name"] = model.CatalogName
	}
	if model.OfferingName != nil {
		modelMap["offering_name"] = model.OfferingName
	}
	if model.OfferingVersion != nil {
		modelMap["offering_version"] = model.OfferingVersion
	}
	if model.OfferingKind != nil {
		modelMap["offering_kind"] = model.OfferingKind
	}
	if model.OfferingID != nil {
		modelMap["offering_id"] = model.OfferingID
	}
	if model.OfferingVersionID != nil {
		modelMap["offering_version_id"] = model.OfferingVersionID
	}
	if model.OfferingRepoURL != nil {
		modelMap["offering_repo_url"] = model.OfferingRepoURL
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobDataWorkItemLastJobToMap(model *schematicsv1.JobDataWorkItemLastJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CommandObject != nil {
		modelMap["command_object"] = model.CommandObject
	}
	if model.CommandObjectName != nil {
		modelMap["command_object_name"] = model.CommandObjectName
	}
	if model.CommandObjectID != nil {
		modelMap["command_object_id"] = model.CommandObjectID
	}
	if model.CommandName != nil {
		modelMap["command_name"] = model.CommandName
	}
	if model.JobID != nil {
		modelMap["job_id"] = model.JobID
	}
	if model.JobStatus != nil {
		modelMap["job_status"] = model.JobStatus
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobBastionResourceDefinitionToMap(model *schematicsv1.BastionResourceDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Host != nil {
		modelMap["host"] = model.Host
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryToMap(model *schematicsv1.JobLogSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.JobID != nil {
		modelMap["job_id"] = model.JobID
	}
	if model.JobType != nil {
		modelMap["job_type"] = model.JobType
	}
	if model.LogStartAt != nil {
		modelMap["log_start_at"] = model.LogStartAt.String()
	}
	if model.LogAnalyzedTill != nil {
		modelMap["log_analyzed_till"] = model.LogAnalyzedTill.String()
	}
	if model.ElapsedTime != nil {
		modelMap["elapsed_time"] = model.ElapsedTime
	}
	if model.LogErrors != nil {
		logErrors := []map[string]interface{}{}
		for _, logErrorsItem := range model.LogErrors {
			logErrorsItemMap, err := ResourceIBMSchematicsJobJobLogSummaryLogErrorsToMap(&logErrorsItem)
			if err != nil {
				return modelMap, err
			}
			logErrors = append(logErrors, logErrorsItemMap)
		}
		modelMap["log_errors"] = logErrors
	}
	if model.RepoDownloadJob != nil {
		repoDownloadJobMap, err := ResourceIBMSchematicsJobJobLogSummaryRepoDownloadJobToMap(model.RepoDownloadJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["repo_download_job"] = []map[string]interface{}{repoDownloadJobMap}
	}
	if model.WorkspaceJob != nil {
		workspaceJobMap, err := ResourceIBMSchematicsJobJobLogSummaryWorkspaceJobToMap(model.WorkspaceJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["workspace_job"] = []map[string]interface{}{workspaceJobMap}
	}
	if model.FlowJob != nil {
		flowJobMap, err := ResourceIBMSchematicsJobJobLogSummaryFlowJobToMap(model.FlowJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_job"] = []map[string]interface{}{flowJobMap}
	}
	if model.ActionJob != nil {
		actionJobMap, err := ResourceIBMSchematicsJobJobLogSummaryActionJobToMap(model.ActionJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_job"] = []map[string]interface{}{actionJobMap}
	}
	if model.SystemJob != nil {
		systemJobMap, err := ResourceIBMSchematicsJobJobLogSummarySystemJobToMap(model.SystemJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_job"] = []map[string]interface{}{systemJobMap}
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryLogErrorsToMap(model *schematicsv1.JobLogSummaryLogErrors) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorCode != nil {
		modelMap["error_code"] = model.ErrorCode
	}
	if model.ErrorMsg != nil {
		modelMap["error_msg"] = model.ErrorMsg
	}
	if model.ErrorCount != nil {
		modelMap["error_count"] = model.ErrorCount
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryRepoDownloadJobToMap(model *schematicsv1.JobLogSummaryRepoDownloadJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ScannedFileCount != nil {
		modelMap["scanned_file_count"] = model.ScannedFileCount
	}
	if model.QuarantinedFileCount != nil {
		modelMap["quarantined_file_count"] = model.QuarantinedFileCount
	}
	if model.DetectedFiletype != nil {
		modelMap["detected_filetype"] = model.DetectedFiletype
	}
	if model.InputsCount != nil {
		modelMap["inputs_count"] = model.InputsCount
	}
	if model.OutputsCount != nil {
		modelMap["outputs_count"] = model.OutputsCount
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryWorkspaceJobToMap(model *schematicsv1.JobLogSummaryWorkspaceJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourcesAdd != nil {
		modelMap["resources_add"] = model.ResourcesAdd
	}
	if model.ResourcesModify != nil {
		modelMap["resources_modify"] = model.ResourcesModify
	}
	if model.ResourcesDestroy != nil {
		modelMap["resources_destroy"] = model.ResourcesDestroy
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryFlowJobToMap(model *schematicsv1.JobLogSummaryFlowJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkitemsCompleted != nil {
		modelMap["workitems_completed"] = model.WorkitemsCompleted
	}
	if model.WorkitemsPending != nil {
		modelMap["workitems_pending"] = model.WorkitemsPending
	}
	if model.WorkitemsFailed != nil {
		modelMap["workitems_failed"] = model.WorkitemsFailed
	}
	if model.Workitems != nil {
		workitems := []map[string]interface{}{}
		for _, workitemsItem := range model.Workitems {
			workitemsItemMap, err := ResourceIBMSchematicsJobJobLogSummaryWorkitemsToMap(&workitemsItem)
			if err != nil {
				return modelMap, err
			}
			workitems = append(workitems, workitemsItemMap)
		}
		modelMap["workitems"] = workitems
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryWorkitemsToMap(model *schematicsv1.JobLogSummaryWorkitems) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = model.WorkspaceID
	}
	if model.JobID != nil {
		modelMap["job_id"] = model.JobID
	}
	if model.ResourcesAdd != nil {
		modelMap["resources_add"] = model.ResourcesAdd
	}
	if model.ResourcesModify != nil {
		modelMap["resources_modify"] = model.ResourcesModify
	}
	if model.ResourcesDestroy != nil {
		modelMap["resources_destroy"] = model.ResourcesDestroy
	}
	if model.LogURL != nil {
		modelMap["log_url"] = model.LogURL
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryActionJobToMap(model *schematicsv1.JobLogSummaryActionJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetCount != nil {
		modelMap["target_count"] = model.TargetCount
	}
	if model.TaskCount != nil {
		modelMap["task_count"] = model.TaskCount
	}
	if model.PlayCount != nil {
		modelMap["play_count"] = model.PlayCount
	}
	if model.Recap != nil {
		recapMap, err := ResourceIBMSchematicsJobJobLogSummaryActionJobRecapToMap(model.Recap)
		if err != nil {
			return modelMap, err
		}
		modelMap["recap"] = []map[string]interface{}{recapMap}
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummaryActionJobRecapToMap(model *schematicsv1.JobLogSummaryActionJobRecap) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Target != nil {
		modelMap["target"] = model.Target
	}
	if model.Ok != nil {
		modelMap["ok"] = model.Ok
	}
	if model.Changed != nil {
		modelMap["changed"] = model.Changed
	}
	if model.Failed != nil {
		modelMap["failed"] = model.Failed
	}
	if model.Skipped != nil {
		modelMap["skipped"] = model.Skipped
	}
	if model.Unreachable != nil {
		modelMap["unreachable"] = model.Unreachable
	}
	return modelMap, nil
}

func ResourceIBMSchematicsJobJobLogSummarySystemJobToMap(model *schematicsv1.JobLogSummarySystemJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetCount != nil {
		modelMap["target_count"] = model.TargetCount
	}
	if model.Success != nil {
		modelMap["success"] = model.Success
	}
	if model.Failed != nil {
		modelMap["failed"] = model.Failed
	}
	return modelMap, nil
}
