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

func DataSourceIBMSchematicsJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMSchematicsJobRead,

		Schema: map[string]*schema.Schema{
			"job_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Job Id. Use `GET /v2/jobs` API to look up the Job Ids in your IBM Cloud account.",
			},
			"command_object": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Schematics automation resource.",
			},
			"command_object_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job command object id (workspace-id, action-id).",
			},
			"command_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Schematics job command name.",
			},
			"command_parameter": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Schematics job command parameter (playbook-name).",
			},
			"command_options": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Command line options for the command.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"job_inputs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Job inputs used by Action or Workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"job_env_settings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Environment variables used by the Job while performing Action or Workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the variable. For example, `name = \"inventory username\"`.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
						},
						"use_default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "An user editable metadata for the variables.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the variable.",
									},
									"aliases": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of aliases for the variable name.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the meta data.",
									},
									"cloud_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
									},
									"default_value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value for the variable only if the override value is not specified.",
									},
									"link_status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the link.",
									},
									"secure": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable secure or sensitive ?.",
									},
									"immutable": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is the variable readonly ?.",
									},
									"hidden": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If **true**, the variable is not displayed on UI or Command line.",
									},
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If the variable required?.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"min_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum value of the variable. Applicable for the integer type.",
									},
									"max_value": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum value of the variable. Applicable for the integer type.",
									},
									"min_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum length of the variable value. Applicable for the string type.",
									},
									"max_length": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum length of the variable value. Applicable for the string type.",
									},
									"matches": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The regex for the variable value.",
									},
									"position": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The relative position of this variable in a list.",
									},
									"group_by": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The display name of the group this variable belongs to.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of this meta-data.",
									},
								},
							},
						},
						"link": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reference link to the variable value By default the expression points to `$self.value`.",
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User defined tags, while running the job.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Job ID.",
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
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
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
			"status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Job Status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"position_in_queue": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Position of job in pending queue.",
						},
						"total_in_queue": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Total no. of jobs in pending queue.",
						},
						"workspace_job_status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Workspace Job Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workspace_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Workspace name.",
									},
									"status_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of Jobs.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Workspace job status message (eg. App1_Setup_Pending, for a 'Setup' flow in the 'App1' Workspace).",
									},
									"flow_status": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Environment Flow JOB Status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"flow_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "flow id.",
												},
												"flow_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "flow name.",
												},
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Flow Job status message - to be displayed along with the status_code;.",
												},
												"workitems": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Environment's individual workItem status details;.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"workspace_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Workspace id.",
															},
															"workspace_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "workspace name.",
															},
															"job_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "workspace job id.",
															},
															"status_code": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Status of Jobs.",
															},
															"status_message": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "workitem job status message;.",
															},
															"updated_at": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "workitem job status updation timestamp.",
															},
														},
													},
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"template_status": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Workspace Flow Template job status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"template_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Template Id.",
												},
												"template_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Template name.",
												},
												"flow_index": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Index of the template in the Flow.",
												},
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Template job status message (eg. VPCt1_Apply_Pending, for a 'VPCt1' Template).",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Job status updation timestamp.",
									},
									"commands": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of terraform commands executed and their status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Name of the command.",
												},
												"outcome": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "Action Job Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action name.",
									},
									"status_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of Jobs.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action Job status message - to be displayed along with the action_status_code.",
									},
									"bastion_status_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of Resources.",
									},
									"bastion_status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Bastion status message - to be displayed along with the bastion_status_code;.",
									},
									"targets_status_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of Resources.",
									},
									"targets_status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Aggregated status message for all target resources,  to be displayed along with the targets_status_code;.",
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"system_job_status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "System Job Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "System job message.",
									},
									"system_status_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of Jobs.",
									},
									"schematics_resource_status": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "job staus for each schematics resource.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "system job status message.",
												},
												"schematics_resource_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "id for each resource which is targeted as a part of system job.",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"flow_job_status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Environment Flow JOB Status.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flow_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "flow id.",
									},
									"flow_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "flow name.",
									},
									"status_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of Jobs.",
									},
									"status_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Flow Job status message - to be displayed along with the status_code;.",
									},
									"workitems": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Environment's individual workItem status details;.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"workspace_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Workspace id.",
												},
												"workspace_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "workspace name.",
												},
												"job_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "workspace job id.",
												},
												"status_code": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status of Jobs.",
												},
												"status_message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "workitem job status message;.",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "workitem job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
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
				Computed:    true,
				Description: "Job data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of Job.",
						},
						"workspace_job_data": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Workspace Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workspace_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Workspace name.",
									},
									"flow_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Flow Id.",
									},
									"flow_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Flow name.",
									},
									"inputs": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Input variables data used by the Workspace Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of aliases for the variable name.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"outputs": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Output variables data from the Workspace Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of aliases for the variable name.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"settings": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Environment variables used by all the templates in the Workspace.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of aliases for the variable name.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"template_data": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Input / output data of the Template in the Workspace Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"template_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Template Id.",
												},
												"template_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Template name.",
												},
												"flow_index": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Index of the template in the Flow.",
												},
												"inputs": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Job inputs used by the Templates.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"outputs": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Job output from the Templates.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"settings": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Environment variables used by the template.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"action_job_data": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Action Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Flow name.",
									},
									"inputs": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Input variables data used by the Action Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of aliases for the variable name.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"outputs": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Output variables data from the Action Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of aliases for the variable name.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"settings": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Environment variables used by all the templates in the Action.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the variable. For example, `name = \"inventory username\"`.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
												},
												"use_default": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
												},
												"metadata": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "An user editable metadata for the variables.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of the variable.",
															},
															"aliases": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of aliases for the variable name.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"description": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The description of the meta data.",
															},
															"cloud_data_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
															},
															"default_value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Default value for the variable only if the override value is not specified.",
															},
															"link_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The status of the link.",
															},
															"secure": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable secure or sensitive ?.",
															},
															"immutable": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Is the variable readonly ?.",
															},
															"hidden": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If **true**, the variable is not displayed on UI or Command line.",
															},
															"required": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "If the variable required?.",
															},
															"options": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"min_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum value of the variable. Applicable for the integer type.",
															},
															"max_value": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum value of the variable. Applicable for the integer type.",
															},
															"min_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The minimum length of the variable value. Applicable for the string type.",
															},
															"max_length": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The maximum length of the variable value. Applicable for the string type.",
															},
															"matches": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The regex for the variable value.",
															},
															"position": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The relative position of this variable in a list.",
															},
															"group_by": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The display name of the group this variable belongs to.",
															},
															"source": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The source of this meta-data.",
															},
														},
													},
												},
												"link": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reference link to the variable value By default the expression points to `$self.value`.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Job status updation timestamp.",
									},
									"inventory_record": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Complete inventory resource details with user inputs and system generated data.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique name of your Inventory.  The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Inventory id.",
												},
												"description": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of your Inventory.  The description can be up to 2048 characters long in size.",
												},
												"location": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.",
												},
												"resource_group": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Resource-group name for the Inventory definition.  By default, Inventory will be created in Default Resource Group.",
												},
												"created_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Inventory creation time.",
												},
												"created_by": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Email address of user who created the Inventory.",
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Inventory updation time.",
												},
												"updated_by": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Email address of user who updated the Inventory.",
												},
												"inventories_ini": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Input inventory of host and host group for the playbook,  in the .ini file format.",
												},
												"resource_queries": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"materialized_inventory": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Materialized inventory details used by the Action Job, in .ini format.",
									},
								},
							},
						},
						"system_job_data": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Controls Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Key ID for which key event is generated.",
									},
									"schematics_resource_id": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of the schematics resource id.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Job status updation timestamp.",
									},
								},
							},
						},
						"flow_job_data": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Flow Job data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flow_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Flow ID.",
									},
									"flow_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Flow Name.",
									},
									"workitems": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Job data used by each workitem Job.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"command_object_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "command object id.",
												},
												"command_object_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "command object name.",
												},
												"layers": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "layer name.",
												},
												"source_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Type of source for the Template.",
												},
												"source": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Source of templates, playbooks, or controls.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"source_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Type of source for the Template.",
															},
															"git": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The connection details to the Git source repository.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"computed_git_repo_url": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**.",
																		},
																		"git_repo_url": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The URL to the Git repository that can be used to clone the template.",
																		},
																		"git_token": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The Personal Access Token (PAT) to connect to the Git URLs.",
																		},
																		"git_repo_folder": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name of the folder in the Git repository, that contains the template.",
																		},
																		"git_release": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name of the release tag that are used to fetch the Git repository.",
																		},
																		"git_branch": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name of the branch that are used to fetch the Git repository.",
																		},
																	},
																},
															},
															"catalog": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The connection details to the IBM Cloud Catalog source.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"catalog_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name of the private catalog.",
																		},
																		"offering_name": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The name of an offering in the IBM Cloud Catalog.",
																		},
																		"offering_version": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The version string of an offering in the IBM Cloud Catalog.",
																		},
																		"offering_kind": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The type of an offering, in the IBM Cloud Catalog.",
																		},
																		"offering_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The ID of an offering in the IBM Cloud Catalog.",
																		},
																		"offering_version_id": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The ID of an offering version the IBM Cloud Catalog.",
																		},
																		"offering_repo_url": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
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
													Computed:    true,
													Description: "Input variables data for the workItem used in FlowJob.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"outputs": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Output variables for the workItem.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"settings": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Environment variables for the workItem.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of the variable. For example, `name = \"inventory username\"`.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value for the variable or reference to the value. For example, `value = \"<provide your ssh_key_value with \n>\"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.",
															},
															"use_default": &schema.Schema{
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.",
															},
															"metadata": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "An user editable metadata for the variables.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Type of the variable.",
																		},
																		"aliases": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of aliases for the variable name.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"description": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The description of the meta data.",
																		},
																		"cloud_data_type": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Cloud data type of the variable. eg. resource_group_id, region, vpc_id.",
																		},
																		"default_value": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Default value for the variable only if the override value is not specified.",
																		},
																		"link_status": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The status of the link.",
																		},
																		"secure": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable secure or sensitive ?.",
																		},
																		"immutable": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Is the variable readonly ?.",
																		},
																		"hidden": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If **true**, the variable is not displayed on UI or Command line.",
																		},
																		"required": &schema.Schema{
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "If the variable required?.",
																		},
																		"options": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.",
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},
																		"min_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum value of the variable. Applicable for the integer type.",
																		},
																		"max_value": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum value of the variable. Applicable for the integer type.",
																		},
																		"min_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The minimum length of the variable value. Applicable for the string type.",
																		},
																		"max_length": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The maximum length of the variable value. Applicable for the string type.",
																		},
																		"matches": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The regex for the variable value.",
																		},
																		"position": &schema.Schema{
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The relative position of this variable in a list.",
																		},
																		"group_by": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The display name of the group this variable belongs to.",
																		},
																		"source": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The source of this meta-data.",
																		},
																	},
																},
															},
															"link": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The reference link to the variable value By default the expression points to `$self.value`.",
															},
														},
													},
												},
												"last_job": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Status of the last job executed by the workitem.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"command_object": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Name of the Schematics automation resource.",
															},
															"command_object_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "command object name (workspace_name/action_name).",
															},
															"command_object_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Workitem command object id, maps to workspace_id or action_id.",
															},
															"command_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Schematics job command name.",
															},
															"job_id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Workspace job id.",
															},
															"job_status": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Status of Jobs.",
															},
														},
													},
												},
												"updated_at": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Job status updation timestamp.",
												},
											},
										},
									},
									"updated_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
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
				Computed:    true,
				Description: "Describes a bastion resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bastion Name(Unique).",
						},
						"host": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reference to the Inventory resource definition.",
						},
					},
				},
			},
			"log_summary": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Job log summary record.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"job_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workspace Id.",
						},
						"job_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of Job.",
						},
						"log_start_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Job log start timestamp.",
						},
						"log_analyzed_till": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Job log update timestamp.",
						},
						"elapsed_time": &schema.Schema{
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Job log elapsed time (log_analyzed_till - log_start_at).",
						},
						"log_errors": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Job log errors.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error_code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Error code in the Log.",
									},
									"error_msg": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Summary error message in the log.",
									},
									"error_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of occurrence.",
									},
								},
							},
						},
						"repo_download_job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Repo download Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"scanned_file_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of files scanned.",
									},
									"quarantined_file_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of files quarantined.",
									},
									"detected_filetype": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Detected template or data file type.",
									},
									"inputs_count": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Number of inputs detected.",
									},
									"outputs_count": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Number of outputs detected.",
									},
								},
							},
						},
						"workspace_job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Workspace Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resources_add": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of resources add.",
									},
									"resources_modify": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of resources modify.",
									},
									"resources_destroy": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of resources destroy.",
									},
								},
							},
						},
						"flow_job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Flow Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workitems_completed": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of workitems completed successfully.",
									},
									"workitems_pending": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of workitems pending in the flow.",
									},
									"workitems_failed": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of workitems failed.",
									},
									"workitems": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"workspace_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "workspace ID.",
												},
												"job_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "workspace JOB ID.",
												},
												"resources_add": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of resources add.",
												},
												"resources_modify": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of resources modify.",
												},
												"resources_destroy": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of resources destroy.",
												},
												"log_url": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "Flow Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "number of targets or hosts.",
									},
									"task_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "number of tasks in playbook.",
									},
									"play_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "number of plays in playbook.",
									},
									"recap": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Recap records.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of target or host name.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"ok": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of OK.",
												},
												"changed": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of changed.",
												},
												"failed": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of failed.",
												},
												"skipped": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "Number of skipped.",
												},
												"unreachable": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
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
							Computed:    true,
							Description: "System Job log summary.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_count": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "number of targets or hosts.",
									},
									"success": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of passed.",
									},
									"failed": &schema.Schema{
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Number of failed.",
									},
								},
							},
						},
					},
				},
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

func DataSourceIBMSchematicsJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	schematicsClient, err := meta.(conns.ClientSession).SchematicsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getJobOptions := &schematicsv1.GetJobOptions{}

	getJobOptions.SetJobID(d.Get("job_id").(string))

	job, response, err := schematicsClient.GetJobWithContext(context, getJobOptions)
	if err != nil {
		log.Printf("[DEBUG] GetJobWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetJobWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getJobOptions.JobID))

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

	jobInputs := []map[string]interface{}{}
	if job.Inputs != nil {
		for _, modelItem := range job.Inputs {
			modelMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			jobInputs = append(jobInputs, modelMap)
		}
	}
	if err = d.Set("job_inputs", jobInputs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_inputs %s", err))
	}

	jobEnvSettings := []map[string]interface{}{}
	if job.Settings != nil {
		for _, modelItem := range job.Settings {
			modelMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			jobEnvSettings = append(jobEnvSettings, modelMap)
		}
	}
	if err = d.Set("job_env_settings", jobEnvSettings); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting job_env_settings %s", err))
	}

	if err = d.Set("id", job.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}

	if err = d.Set("name", job.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("description", job.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if err = d.Set("location", job.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
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

	status := []map[string]interface{}{}
	if job.Status != nil {
		modelMap, err := DataSourceIBMSchematicsJobJobStatusToMap(job.Status)
		if err != nil {
			return diag.FromErr(err)
		}
		status = append(status, modelMap)
	}
	if err = d.Set("status", status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status %s", err))
	}

	data := []map[string]interface{}{}
	if job.Data != nil {
		modelMap, err := DataSourceIBMSchematicsJobJobDataToMap(job.Data)
		if err != nil {
			return diag.FromErr(err)
		}
		data = append(data, modelMap)
	}
	if err = d.Set("data", data); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting data %s", err))
	}

	bastion := []map[string]interface{}{}
	if job.Bastion != nil {
		modelMap, err := DataSourceIBMSchematicsJobBastionResourceDefinitionToMap(job.Bastion)
		if err != nil {
			return diag.FromErr(err)
		}
		bastion = append(bastion, modelMap)
	}
	if err = d.Set("bastion", bastion); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting bastion %s", err))
	}

	logSummary := []map[string]interface{}{}
	if job.LogSummary != nil {
		modelMap, err := DataSourceIBMSchematicsJobJobLogSummaryToMap(job.LogSummary)
		if err != nil {
			return diag.FromErr(err)
		}
		logSummary = append(logSummary, modelMap)
	}
	if err = d.Set("log_summary", logSummary); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting log_summary %s", err))
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

func DataSourceIBMSchematicsJobVariableDataToMap(model *schematicsv1.VariableData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.UseDefault != nil {
		modelMap["use_default"] = *model.UseDefault
	}
	if model.Metadata != nil {
		metadataMap, err := DataSourceIBMSchematicsJobVariableMetadataToMap(model.Metadata)
		if err != nil {
			return modelMap, err
		}
		modelMap["metadata"] = []map[string]interface{}{metadataMap}
	}
	if model.Link != nil {
		modelMap["link"] = *model.Link
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobVariableMetadataToMap(model *schematicsv1.VariableMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Aliases != nil {
		modelMap["aliases"] = model.Aliases
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.CloudDataType != nil {
		modelMap["cloud_data_type"] = *model.CloudDataType
	}
	if model.DefaultValue != nil {
		modelMap["default_value"] = *model.DefaultValue
	}
	if model.LinkStatus != nil {
		modelMap["link_status"] = *model.LinkStatus
	}
	if model.Secure != nil {
		modelMap["secure"] = *model.Secure
	}
	if model.Immutable != nil {
		modelMap["immutable"] = *model.Immutable
	}
	if model.Hidden != nil {
		modelMap["hidden"] = *model.Hidden
	}
	if model.Required != nil {
		modelMap["required"] = *model.Required
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	if model.MinValue != nil {
		modelMap["min_value"] = *model.MinValue
	}
	if model.MaxValue != nil {
		modelMap["max_value"] = *model.MaxValue
	}
	if model.MinLength != nil {
		modelMap["min_length"] = *model.MinLength
	}
	if model.MaxLength != nil {
		modelMap["max_length"] = *model.MaxLength
	}
	if model.Matches != nil {
		modelMap["matches"] = *model.Matches
	}
	if model.Position != nil {
		modelMap["position"] = *model.Position
	}
	if model.GroupBy != nil {
		modelMap["group_by"] = *model.GroupBy
	}
	if model.Source != nil {
		modelMap["source"] = *model.Source
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobStatusToMap(model *schematicsv1.JobStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PositionInQueue != nil {
		modelMap["position_in_queue"] = *model.PositionInQueue
	}
	if model.TotalInQueue != nil {
		modelMap["total_in_queue"] = *model.TotalInQueue
	}
	if model.WorkspaceJobStatus != nil {
		workspaceJobStatusMap, err := DataSourceIBMSchematicsJobJobStatusWorkspaceToMap(model.WorkspaceJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["workspace_job_status"] = []map[string]interface{}{workspaceJobStatusMap}
	}
	if model.ActionJobStatus != nil {
		actionJobStatusMap, err := DataSourceIBMSchematicsJobJobStatusActionToMap(model.ActionJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_job_status"] = []map[string]interface{}{actionJobStatusMap}
	}
	if model.SystemJobStatus != nil {
		systemJobStatusMap, err := DataSourceIBMSchematicsJobJobStatusSystemToMap(model.SystemJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_job_status"] = []map[string]interface{}{systemJobStatusMap}
	}
	if model.FlowJobStatus != nil {
		flowJobStatusMap, err := DataSourceIBMSchematicsJobJobStatusFlowToMap(model.FlowJobStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_job_status"] = []map[string]interface{}{flowJobStatusMap}
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobStatusWorkspaceToMap(model *schematicsv1.JobStatusWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = *model.WorkspaceName
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.FlowStatus != nil {
		flowStatusMap, err := DataSourceIBMSchematicsJobJobStatusFlowToMap(model.FlowStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_status"] = []map[string]interface{}{flowStatusMap}
	}
	if model.TemplateStatus != nil {
		templateStatus := []map[string]interface{}{}
		for _, templateStatusItem := range model.TemplateStatus {
			templateStatusItemMap, err := DataSourceIBMSchematicsJobJobStatusTemplateToMap(&templateStatusItem)
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
			commandsItemMap, err := DataSourceIBMSchematicsJobCommandsInfoToMap(&commandsItem)
			if err != nil {
				return modelMap, err
			}
			commands = append(commands, commandsItemMap)
		}
		modelMap["commands"] = commands
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobStatusFlowToMap(model *schematicsv1.JobStatusFlow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FlowID != nil {
		modelMap["flow_id"] = *model.FlowID
	}
	if model.FlowName != nil {
		modelMap["flow_name"] = *model.FlowName
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.Workitems != nil {
		workitems := []map[string]interface{}{}
		for _, workitemsItem := range model.Workitems {
			workitemsItemMap, err := DataSourceIBMSchematicsJobJobStatusWorkitemToMap(&workitemsItem)
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

func DataSourceIBMSchematicsJobJobStatusWorkitemToMap(model *schematicsv1.JobStatusWorkitem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = *model.WorkspaceID
	}
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = *model.WorkspaceName
	}
	if model.JobID != nil {
		modelMap["job_id"] = *model.JobID
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobStatusTemplateToMap(model *schematicsv1.JobStatusTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TemplateID != nil {
		modelMap["template_id"] = *model.TemplateID
	}
	if model.TemplateName != nil {
		modelMap["template_name"] = *model.TemplateName
	}
	if model.FlowIndex != nil {
		modelMap["flow_index"] = *model.FlowIndex
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobCommandsInfoToMap(model *schematicsv1.CommandsInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Outcome != nil {
		modelMap["outcome"] = *model.Outcome
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobStatusActionToMap(model *schematicsv1.JobStatusAction) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActionName != nil {
		modelMap["action_name"] = *model.ActionName
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.BastionStatusCode != nil {
		modelMap["bastion_status_code"] = *model.BastionStatusCode
	}
	if model.BastionStatusMessage != nil {
		modelMap["bastion_status_message"] = *model.BastionStatusMessage
	}
	if model.TargetsStatusCode != nil {
		modelMap["targets_status_code"] = *model.TargetsStatusCode
	}
	if model.TargetsStatusMessage != nil {
		modelMap["targets_status_message"] = *model.TargetsStatusMessage
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobStatusSystemToMap(model *schematicsv1.JobStatusSystem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SystemStatusMessage != nil {
		modelMap["system_status_message"] = *model.SystemStatusMessage
	}
	if model.SystemStatusCode != nil {
		modelMap["system_status_code"] = *model.SystemStatusCode
	}
	if model.SchematicsResourceStatus != nil {
		schematicsResourceStatus := []map[string]interface{}{}
		for _, schematicsResourceStatusItem := range model.SchematicsResourceStatus {
			schematicsResourceStatusItemMap, err := DataSourceIBMSchematicsJobJobStatusSchematicsResourcesToMap(&schematicsResourceStatusItem)
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

func DataSourceIBMSchematicsJobJobStatusSchematicsResourcesToMap(model *schematicsv1.JobStatusSchematicsResources) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.StatusCode != nil {
		modelMap["status_code"] = *model.StatusCode
	}
	if model.StatusMessage != nil {
		modelMap["status_message"] = *model.StatusMessage
	}
	if model.SchematicsResourceID != nil {
		modelMap["schematics_resource_id"] = *model.SchematicsResourceID
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobDataToMap(model *schematicsv1.JobData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.JobType != nil {
		modelMap["job_type"] = *model.JobType
	}
	if model.WorkspaceJobData != nil {
		workspaceJobDataMap, err := DataSourceIBMSchematicsJobJobDataWorkspaceToMap(model.WorkspaceJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["workspace_job_data"] = []map[string]interface{}{workspaceJobDataMap}
	}
	if model.ActionJobData != nil {
		actionJobDataMap, err := DataSourceIBMSchematicsJobJobDataActionToMap(model.ActionJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_job_data"] = []map[string]interface{}{actionJobDataMap}
	}
	if model.SystemJobData != nil {
		systemJobDataMap, err := DataSourceIBMSchematicsJobJobDataSystemToMap(model.SystemJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_job_data"] = []map[string]interface{}{systemJobDataMap}
	}
	if model.FlowJobData != nil {
		flowJobDataMap, err := DataSourceIBMSchematicsJobJobDataFlowToMap(model.FlowJobData)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_job_data"] = []map[string]interface{}{flowJobDataMap}
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobDataWorkspaceToMap(model *schematicsv1.JobDataWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceName != nil {
		modelMap["workspace_name"] = *model.WorkspaceName
	}
	if model.FlowID != nil {
		modelMap["flow_id"] = *model.FlowID
	}
	if model.FlowName != nil {
		modelMap["flow_name"] = *model.FlowName
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&inputsItem)
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
			outputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&outputsItem)
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
			settingsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&settingsItem)
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
			templateDataItemMap, err := DataSourceIBMSchematicsJobJobDataTemplateToMap(&templateDataItem)
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

func DataSourceIBMSchematicsJobJobDataTemplateToMap(model *schematicsv1.JobDataTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TemplateID != nil {
		modelMap["template_id"] = *model.TemplateID
	}
	if model.TemplateName != nil {
		modelMap["template_name"] = *model.TemplateName
	}
	if model.FlowIndex != nil {
		modelMap["flow_index"] = *model.FlowIndex
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&inputsItem)
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
			outputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&outputsItem)
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
			settingsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&settingsItem)
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

func DataSourceIBMSchematicsJobJobDataActionToMap(model *schematicsv1.JobDataAction) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ActionName != nil {
		modelMap["action_name"] = *model.ActionName
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&inputsItem)
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
			outputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&outputsItem)
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
			settingsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&settingsItem)
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
		inventoryRecordMap, err := DataSourceIBMSchematicsJobInventoryResourceRecordToMap(model.InventoryRecord)
		if err != nil {
			return modelMap, err
		}
		modelMap["inventory_record"] = []map[string]interface{}{inventoryRecordMap}
	}
	if model.MaterializedInventory != nil {
		modelMap["materialized_inventory"] = *model.MaterializedInventory
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobInventoryResourceRecordToMap(model *schematicsv1.InventoryResourceRecord) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Location != nil {
		modelMap["location"] = *model.Location
	}
	if model.ResourceGroup != nil {
		modelMap["resource_group"] = *model.ResourceGroup
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.UpdatedBy != nil {
		modelMap["updated_by"] = *model.UpdatedBy
	}
	if model.InventoriesIni != nil {
		modelMap["inventories_ini"] = *model.InventoriesIni
	}
	if model.ResourceQueries != nil {
		modelMap["resource_queries"] = model.ResourceQueries
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobDataSystemToMap(model *schematicsv1.JobDataSystem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.KeyID != nil {
		modelMap["key_id"] = *model.KeyID
	}
	if model.SchematicsResourceID != nil {
		modelMap["schematics_resource_id"] = model.SchematicsResourceID
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobDataFlowToMap(model *schematicsv1.JobDataFlow) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.FlowID != nil {
		modelMap["flow_id"] = *model.FlowID
	}
	if model.FlowName != nil {
		modelMap["flow_name"] = *model.FlowName
	}
	if model.Workitems != nil {
		workitems := []map[string]interface{}{}
		for _, workitemsItem := range model.Workitems {
			workitemsItemMap, err := DataSourceIBMSchematicsJobJobDataWorkItemToMap(&workitemsItem)
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

func DataSourceIBMSchematicsJobJobDataWorkItemToMap(model *schematicsv1.JobDataWorkItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CommandObjectID != nil {
		modelMap["command_object_id"] = *model.CommandObjectID
	}
	if model.CommandObjectName != nil {
		modelMap["command_object_name"] = *model.CommandObjectName
	}
	if model.Layers != nil {
		modelMap["layers"] = *model.Layers
	}
	if model.SourceType != nil {
		modelMap["source_type"] = *model.SourceType
	}
	if model.Source != nil {
		sourceMap, err := DataSourceIBMSchematicsJobExternalSourceToMap(model.Source)
		if err != nil {
			return modelMap, err
		}
		modelMap["source"] = []map[string]interface{}{sourceMap}
	}
	if model.Inputs != nil {
		inputs := []map[string]interface{}{}
		for _, inputsItem := range model.Inputs {
			inputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&inputsItem)
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
			outputsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&outputsItem)
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
			settingsItemMap, err := DataSourceIBMSchematicsJobVariableDataToMap(&settingsItem)
			if err != nil {
				return modelMap, err
			}
			settings = append(settings, settingsItemMap)
		}
		modelMap["settings"] = settings
	}
	if model.LastJob != nil {
		lastJobMap, err := DataSourceIBMSchematicsJobJobDataWorkItemLastJobToMap(model.LastJob)
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

func DataSourceIBMSchematicsJobExternalSourceToMap(model *schematicsv1.ExternalSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SourceType != nil {
		modelMap["source_type"] = *model.SourceType
	}
	if model.Git != nil {
		gitMap, err := DataSourceIBMSchematicsJobExternalSourceGitToMap(model.Git)
		if err != nil {
			return modelMap, err
		}
		modelMap["git"] = []map[string]interface{}{gitMap}
	}
	if model.Catalog != nil {
		catalogMap, err := DataSourceIBMSchematicsJobExternalSourceCatalogToMap(model.Catalog)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog"] = []map[string]interface{}{catalogMap}
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobExternalSourceGitToMap(model *schematicsv1.ExternalSourceGit) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ComputedGitRepoURL != nil {
		modelMap["computed_git_repo_url"] = *model.ComputedGitRepoURL
	}
	if model.GitRepoURL != nil {
		modelMap["git_repo_url"] = *model.GitRepoURL
	}
	if model.GitToken != nil {
		modelMap["git_token"] = *model.GitToken
	}
	if model.GitRepoFolder != nil {
		modelMap["git_repo_folder"] = *model.GitRepoFolder
	}
	if model.GitRelease != nil {
		modelMap["git_release"] = *model.GitRelease
	}
	if model.GitBranch != nil {
		modelMap["git_branch"] = *model.GitBranch
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobExternalSourceCatalogToMap(model *schematicsv1.ExternalSourceCatalog) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogName != nil {
		modelMap["catalog_name"] = *model.CatalogName
	}
	if model.OfferingName != nil {
		modelMap["offering_name"] = *model.OfferingName
	}
	if model.OfferingVersion != nil {
		modelMap["offering_version"] = *model.OfferingVersion
	}
	if model.OfferingKind != nil {
		modelMap["offering_kind"] = *model.OfferingKind
	}
	if model.OfferingID != nil {
		modelMap["offering_id"] = *model.OfferingID
	}
	if model.OfferingVersionID != nil {
		modelMap["offering_version_id"] = *model.OfferingVersionID
	}
	if model.OfferingRepoURL != nil {
		modelMap["offering_repo_url"] = *model.OfferingRepoURL
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobDataWorkItemLastJobToMap(model *schematicsv1.JobDataWorkItemLastJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CommandObject != nil {
		modelMap["command_object"] = *model.CommandObject
	}
	if model.CommandObjectName != nil {
		modelMap["command_object_name"] = *model.CommandObjectName
	}
	if model.CommandObjectID != nil {
		modelMap["command_object_id"] = *model.CommandObjectID
	}
	if model.CommandName != nil {
		modelMap["command_name"] = *model.CommandName
	}
	if model.JobID != nil {
		modelMap["job_id"] = *model.JobID
	}
	if model.JobStatus != nil {
		modelMap["job_status"] = *model.JobStatus
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobBastionResourceDefinitionToMap(model *schematicsv1.BastionResourceDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Host != nil {
		modelMap["host"] = *model.Host
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryToMap(model *schematicsv1.JobLogSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.JobID != nil {
		modelMap["job_id"] = *model.JobID
	}
	if model.JobType != nil {
		modelMap["job_type"] = *model.JobType
	}
	if model.LogStartAt != nil {
		modelMap["log_start_at"] = model.LogStartAt.String()
	}
	if model.LogAnalyzedTill != nil {
		modelMap["log_analyzed_till"] = model.LogAnalyzedTill.String()
	}
	if model.ElapsedTime != nil {
		modelMap["elapsed_time"] = *model.ElapsedTime
	}
	if model.LogErrors != nil {
		logErrors := []map[string]interface{}{}
		for _, logErrorsItem := range model.LogErrors {
			logErrorsItemMap, err := DataSourceIBMSchematicsJobJobLogSummaryLogErrorsToMap(&logErrorsItem)
			if err != nil {
				return modelMap, err
			}
			logErrors = append(logErrors, logErrorsItemMap)
		}
		modelMap["log_errors"] = logErrors
	}
	if model.RepoDownloadJob != nil {
		repoDownloadJobMap, err := DataSourceIBMSchematicsJobJobLogSummaryRepoDownloadJobToMap(model.RepoDownloadJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["repo_download_job"] = []map[string]interface{}{repoDownloadJobMap}
	}
	if model.WorkspaceJob != nil {
		workspaceJobMap, err := DataSourceIBMSchematicsJobJobLogSummaryWorkspaceJobToMap(model.WorkspaceJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["workspace_job"] = []map[string]interface{}{workspaceJobMap}
	}
	if model.FlowJob != nil {
		flowJobMap, err := DataSourceIBMSchematicsJobJobLogSummaryFlowJobToMap(model.FlowJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["flow_job"] = []map[string]interface{}{flowJobMap}
	}
	if model.ActionJob != nil {
		actionJobMap, err := DataSourceIBMSchematicsJobJobLogSummaryActionJobToMap(model.ActionJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_job"] = []map[string]interface{}{actionJobMap}
	}
	if model.SystemJob != nil {
		systemJobMap, err := DataSourceIBMSchematicsJobJobLogSummarySystemJobToMap(model.SystemJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["system_job"] = []map[string]interface{}{systemJobMap}
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryLogErrorsToMap(model *schematicsv1.JobLogSummaryLogErrors) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorCode != nil {
		modelMap["error_code"] = *model.ErrorCode
	}
	if model.ErrorMsg != nil {
		modelMap["error_msg"] = *model.ErrorMsg
	}
	if model.ErrorCount != nil {
		modelMap["error_count"] = *model.ErrorCount
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryRepoDownloadJobToMap(model *schematicsv1.JobLogSummaryRepoDownloadJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ScannedFileCount != nil {
		modelMap["scanned_file_count"] = *model.ScannedFileCount
	}
	if model.QuarantinedFileCount != nil {
		modelMap["quarantined_file_count"] = *model.QuarantinedFileCount
	}
	if model.DetectedFiletype != nil {
		modelMap["detected_filetype"] = *model.DetectedFiletype
	}
	if model.InputsCount != nil {
		modelMap["inputs_count"] = *model.InputsCount
	}
	if model.OutputsCount != nil {
		modelMap["outputs_count"] = *model.OutputsCount
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryWorkspaceJobToMap(model *schematicsv1.JobLogSummaryWorkspaceJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourcesAdd != nil {
		modelMap["resources_add"] = *model.ResourcesAdd
	}
	if model.ResourcesModify != nil {
		modelMap["resources_modify"] = *model.ResourcesModify
	}
	if model.ResourcesDestroy != nil {
		modelMap["resources_destroy"] = *model.ResourcesDestroy
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryFlowJobToMap(model *schematicsv1.JobLogSummaryFlowJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkitemsCompleted != nil {
		modelMap["workitems_completed"] = *model.WorkitemsCompleted
	}
	if model.WorkitemsPending != nil {
		modelMap["workitems_pending"] = *model.WorkitemsPending
	}
	if model.WorkitemsFailed != nil {
		modelMap["workitems_failed"] = *model.WorkitemsFailed
	}
	if model.Workitems != nil {
		workitems := []map[string]interface{}{}
		for _, workitemsItem := range model.Workitems {
			workitemsItemMap, err := DataSourceIBMSchematicsJobJobLogSummaryWorkitemsToMap(&workitemsItem)
			if err != nil {
				return modelMap, err
			}
			workitems = append(workitems, workitemsItemMap)
		}
		modelMap["workitems"] = workitems
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryWorkitemsToMap(model *schematicsv1.JobLogSummaryWorkitems) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.WorkspaceID != nil {
		modelMap["workspace_id"] = *model.WorkspaceID
	}
	if model.JobID != nil {
		modelMap["job_id"] = *model.JobID
	}
	if model.ResourcesAdd != nil {
		modelMap["resources_add"] = *model.ResourcesAdd
	}
	if model.ResourcesModify != nil {
		modelMap["resources_modify"] = *model.ResourcesModify
	}
	if model.ResourcesDestroy != nil {
		modelMap["resources_destroy"] = *model.ResourcesDestroy
	}
	if model.LogURL != nil {
		modelMap["log_url"] = *model.LogURL
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryActionJobToMap(model *schematicsv1.JobLogSummaryActionJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetCount != nil {
		modelMap["target_count"] = *model.TargetCount
	}
	if model.TaskCount != nil {
		modelMap["task_count"] = *model.TaskCount
	}
	if model.PlayCount != nil {
		modelMap["play_count"] = *model.PlayCount
	}
	if model.Recap != nil {
		recapMap, err := DataSourceIBMSchematicsJobJobLogSummaryActionJobRecapToMap(model.Recap)
		if err != nil {
			return modelMap, err
		}
		modelMap["recap"] = []map[string]interface{}{recapMap}
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummaryActionJobRecapToMap(model *schematicsv1.JobLogSummaryActionJobRecap) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Target != nil {
		modelMap["target"] = model.Target
	}
	if model.Ok != nil {
		modelMap["ok"] = *model.Ok
	}
	if model.Changed != nil {
		modelMap["changed"] = *model.Changed
	}
	if model.Failed != nil {
		modelMap["failed"] = *model.Failed
	}
	if model.Skipped != nil {
		modelMap["skipped"] = *model.Skipped
	}
	if model.Unreachable != nil {
		modelMap["unreachable"] = *model.Unreachable
	}
	return modelMap, nil
}

func DataSourceIBMSchematicsJobJobLogSummarySystemJobToMap(model *schematicsv1.JobLogSummarySystemJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetCount != nil {
		modelMap["target_count"] = *model.TargetCount
	}
	if model.Success != nil {
		modelMap["success"] = *model.Success
	}
	if model.Failed != nil {
		modelMap["failed"] = *model.Failed
	}
	return modelMap, nil
}
