// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

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
	"github.com/IBM/project-go-sdk/projectv1"
)

func ResourceIbmProject() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmProjectCreate,
		ReadContext:     resourceIbmProjectRead,
		UpdateContext:   resourceIbmProjectUpdate,
		DeleteContext:   resourceIbmProjectDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "resource_group"),
				Description: "The resource group where the project's data and tools are created.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "location"),
				Description: "The location where the project's data and tools are created.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "name"),
				Description: "The name of the project.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "description"),
				Description: "A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.",
			},
			"destroy_on_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "The policy that indicates whether the resources are destroyed or not when a project is deleted.",
			},
			"configs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the configuration.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of the project configuration.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A collection of configuration labels.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"authorizations": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trusted_profile": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The trusted profile for authorizations.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The unique ID.",
												},
												"target_iam_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The unique ID.",
												},
											},
										},
									},
									"method": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.",
									},
									"api_key": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IBM Cloud API Key.",
									},
								},
							},
						},
						"compliance_profile": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The profile required for compliance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
									"instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
									"instance_location": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The location of the compliance instance.",
									},
									"attachment_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
									"profile_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the compliance profile.",
									},
								},
							},
						},
						"locator_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A dotted value of catalogID.versionID.",
						},
						"input": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The input variables for the configuration definition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
								},
							},
						},
						"setting": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
								},
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
						},
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique ID.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The version of the configuration.",
						},
						"is_draft": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The flag that indicates whether the version of the configuration is draft, or active.",
						},
						"needs_attention_state": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The needs attention state of a configuration.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The state of the configuration.",
						},
						"update_available": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The flag that indicates whether a configuration update is available.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"last_approved": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The last approved metadata of the configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"is_forced": &schema.Schema{
										Type:        schema.TypeBool,
										Required:    true,
										Description: "The flag that indicates whether the approval was forced approved.",
									},
									"comment": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The comment left by the user who approved the configuration.",
									},
									"timestamp": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The unique ID.",
									},
								},
							},
						},
						"last_save": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"cra_logs": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The Code Risk Analyzer logs of the configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cra_version": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The version of the Code Risk Analyzer logs of the configuration.",
									},
									"schema_version": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The schema version of Code Risk Analyzer logs of the configuration.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The status of the Code Risk Analyzer logs of the configuration.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the Code Risk Analyzer logs of the configuration.",
										Elem: &schema.Schema{Type: schema.TypeString},
									},
									"timestamp": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
									},
								},
							},
						},
						"cost_estimate": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The cost estimate of the configuration.It only exists after the first configuration validation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The version of the cost estimate of the configuration.",
									},
									"currency": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The currency of the cost estimate of the configuration.",
									},
									"total_hourly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The total hourly cost estimate of the configuration.",
									},
									"total_monthly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The total monthly cost estimate of the configuration.",
									},
									"past_total_hourly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The past total hourly cost estimate of the configuration.",
									},
									"past_total_monthly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The past total monthly cost estimate of the configuration.",
									},
									"diff_total_hourly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The difference between current and past total hourly cost estimates of the configuration.",
									},
									"diff_total_monthly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The difference between current and past total monthly cost estimates of the configuration.",
									},
									"time_generated": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
								},
							},
						},
						"check_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The action job performed on the project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A relative URL.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The summaries of jobs that were performed on the configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plan_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the plan jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"apply_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the apply jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"destroy_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the destroy jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"message_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The message summaries of jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"plan_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of plan jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"apply_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of apply jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"destroy_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of destroy jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
						"install_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The action job performed on the project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A relative URL.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The summaries of jobs that were performed on the configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plan_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the plan jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"apply_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the apply jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"destroy_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the destroy jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"message_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The message summaries of jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"plan_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of plan jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"apply_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of apply jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"destroy_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of destroy jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
						"uninstall_job": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The action job performed on the project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A relative URL.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The summaries of jobs that were performed on the configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plan_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the plan jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"apply_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the apply jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"destroy_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The summary of the destroy jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"message_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The message summaries of jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"plan_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of plan jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"apply_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of apply jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
												"destroy_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "The messages of destroy jobs on the configuration.",
													Elem: &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
								},
							},
						},
						"output": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The outputs of a Schematics template property.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The variable name.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A short explanation of the output value.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Can be any value - a string, number, boolean, array, or object.",
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of a project configuration manual property.",
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An IBM Cloud resource name, which uniquely identifies a resource.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"cumulative_needs_attention_view": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The cumulative list of needs attention items for a project. If the view is successfully retrieved, an array which could be empty is returned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The event name.",
						},
						"event_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A unique ID for that individual event.",
						},
						"config_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A unique ID for the configuration.",
						},
						"config_version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The version number of the configuration.",
						},
					},
				},
			},
			"cumulative_needs_attention_view_error": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True indicates that the fetch of the needs attention items failed. It only exists if there was an error while retrieving the cumulative needs attention view.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project status value.",
			},
			"event_notifications_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the event notifications instance if one is connected to this project.",
			},
		},
	}
}

func ResourceIbmProjectValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "resource_group",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^$|^(?!\s)(?!.*\s$)[^'"` + "`" + `<>{}\x00-\x1F]*$`,
			MinValueLength:             0,
			MaxValueLength:             40,
		},
		validate.ValidateSchema{
			Identifier:                 "location",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^$|^(us-south|us-east|eu-gb|eu-de)$`,
			MinValueLength:             0,
			MaxValueLength:             12,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(?!\s)(?!.*\s$)[^'"` + "`" + `<>{}\x00-\x1F]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^$|^(?!\s).*\S$`,
			MinValueLength:             0,
			MaxValueLength:             1024,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_project", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProjectCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createProjectOptions := &projectv1.CreateProjectOptions{}

	createProjectOptions.SetResourceGroup(d.Get("resource_group").(string))
	createProjectOptions.SetLocation(d.Get("location").(string))
	if _, ok := d.GetOk("name"); ok {
		createProjectOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createProjectOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("destroy_on_delete"); ok {
		createProjectOptions.SetDestroyOnDelete(d.Get("destroy_on_delete").(bool))
	}
	if _, ok := d.GetOk("configs"); ok {
		var configs []projectv1.ProjectConfigTerraform
		for _, v := range d.Get("configs").([]interface{}) {
			value := v.(map[string]interface{})
			configsItem, err := resourceIbmProjectMapToProjectConfigTerraform(value)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, *configsItem)
		}
		createProjectOptions.SetConfigs(configs)
	}

	projectCanonical, response, err := projectClient.CreateProjectWithContext(context, createProjectOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateProjectWithContext failed %s\n%s", err, response))
	}

	d.SetId(*projectCanonical.ID)

	return resourceIbmProjectRead(context, d, meta)
}

func resourceIbmProjectRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProjectOptions := &projectv1.GetProjectOptions{}

	getProjectOptions.SetID(d.Id())

	projectCanonical, response, err := projectClient.GetProjectWithContext(context, getProjectOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProjectWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("resource_group", projectCanonical.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}
	if err = d.Set("location", projectCanonical.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if !core.IsNil(projectCanonical.Name) {
		if err = d.Set("name", projectCanonical.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.Description) {
		if err = d.Set("description", projectCanonical.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.DestroyOnDelete) {
		if err = d.Set("destroy_on_delete", projectCanonical.DestroyOnDelete); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting destroy_on_delete: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.Configs) {
		configs := []map[string]interface{}{}
		for _, configsItem := range projectCanonical.Configs {
			configsItemMap, err := resourceIbmProjectProjectConfigCanonicalToMap(&configsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, configsItemMap)
		}
		if err = d.Set("configs", configs); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting configs: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.Crn) {
		if err = d.Set("crn", projectCanonical.Crn); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(projectCanonical.CreatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.CumulativeNeedsAttentionView) {
		cumulativeNeedsAttentionView := []map[string]interface{}{}
		for _, cumulativeNeedsAttentionViewItem := range projectCanonical.CumulativeNeedsAttentionView {
			cumulativeNeedsAttentionViewItemMap, err := resourceIbmProjectCumulativeNeedsAttentionToMap(&cumulativeNeedsAttentionViewItem)
			if err != nil {
				return diag.FromErr(err)
			}
			cumulativeNeedsAttentionView = append(cumulativeNeedsAttentionView, cumulativeNeedsAttentionViewItemMap)
		}
		if err = d.Set("cumulative_needs_attention_view", cumulativeNeedsAttentionView); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cumulative_needs_attention_view: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.CumulativeNeedsAttentionViewError) {
		if err = d.Set("cumulative_needs_attention_view_error", projectCanonical.CumulativeNeedsAttentionViewError); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cumulative_needs_attention_view_error: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.State) {
		if err = d.Set("state", projectCanonical.State); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
		}
	}
	if !core.IsNil(projectCanonical.EventNotificationsCrn) {
		if err = d.Set("event_notifications_crn", projectCanonical.EventNotificationsCrn); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting event_notifications_crn: %s", err))
		}
	}

	return nil
}

func resourceIbmProjectUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProjectOptions := &projectv1.UpdateProjectOptions{}

	updateProjectOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("name") {
		updateProjectOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateProjectOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("destroy_on_delete") {
		updateProjectOptions.SetDestroyOnDelete(d.Get("destroy_on_delete").(bool))
		hasChange = true
	}
	if d.HasChange("configs") {
		var configs []projectv1.ProjectConfigTerraform
		for _, v := range d.Get("configs").([]interface{}) {
			value := v.(map[string]interface{})
			configsItem, err := resourceIbmProjectMapToProjectConfigTerraform(value)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, *configsItem)
		}
		updateProjectOptions.SetConfigs(configs)
		hasChange = true
	}

	if hasChange {
		_, response, err := projectClient.UpdateProjectWithContext(context, updateProjectOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateProjectWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateProjectWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmProjectRead(context, d, meta)
}

func resourceIbmProjectDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProjectOptions := &projectv1.DeleteProjectOptions{}

	deleteProjectOptions.SetID(d.Id())

	response, err := projectClient.DeleteProjectWithContext(context, deleteProjectOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteProjectWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmProjectMapToProjectConfigTerraform(modelMap map[string]interface{}) (*projectv1.ProjectConfigTerraform, error) {
	model := &projectv1.ProjectConfigTerraform{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["labels"] != nil {
		labels := []string{}
		for _, labelsItem := range modelMap["labels"].([]interface{}) {
			labels = append(labels, labelsItem.(string))
		}
		model.Labels = labels
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := resourceIbmProjectMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := resourceIbmProjectMapToProjectConfigComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	if modelMap["locator_id"] != nil && modelMap["locator_id"].(string) != "" {
		model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	}
	if modelMap["input"] != nil && len(modelMap["input"].([]interface{})) > 0 {
		InputModel, err := resourceIbmProjectMapToInputVariable(modelMap["input"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Input = InputModel
	}
	if modelMap["setting"] != nil && len(modelMap["setting"].([]interface{})) > 0 {
		SettingModel, err := resourceIbmProjectMapToProjectConfigSetting(modelMap["setting"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Setting = SettingModel
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigAuth(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuth, error) {
	model := &projectv1.ProjectConfigAuth{}
	if modelMap["trusted_profile"] != nil && len(modelMap["trusted_profile"].([]interface{})) > 0 {
		TrustedProfileModel, err := resourceIbmProjectMapToProjectConfigAuthTrustedProfile(modelMap["trusted_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TrustedProfile = TrustedProfileModel
	}
	if modelMap["method"] != nil && modelMap["method"].(string) != "" {
		model.Method = core.StringPtr(modelMap["method"].(string))
	}
	if modelMap["api_key"] != nil && modelMap["api_key"].(string) != "" {
		model.ApiKey = core.StringPtr(modelMap["api_key"].(string))
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigAuthTrustedProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuthTrustedProfile, error) {
	model := &projectv1.ProjectConfigAuthTrustedProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["target_iam_id"] != nil && modelMap["target_iam_id"].(string) != "" {
		model.TargetIamID = core.StringPtr(modelMap["target_iam_id"].(string))
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigComplianceProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigComplianceProfile, error) {
	model := &projectv1.ProjectConfigComplianceProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["instance_id"] != nil && modelMap["instance_id"].(string) != "" {
		model.InstanceID = core.StringPtr(modelMap["instance_id"].(string))
	}
	if modelMap["instance_location"] != nil && modelMap["instance_location"].(string) != "" {
		model.InstanceLocation = core.StringPtr(modelMap["instance_location"].(string))
	}
	if modelMap["attachment_id"] != nil && modelMap["attachment_id"].(string) != "" {
		model.AttachmentID = core.StringPtr(modelMap["attachment_id"].(string))
	}
	if modelMap["profile_name"] != nil && modelMap["profile_name"].(string) != "" {
		model.ProfileName = core.StringPtr(modelMap["profile_name"].(string))
	}
	return model, nil
}

func resourceIbmProjectMapToInputVariable(modelMap map[string]interface{}) (*projectv1.InputVariable, error) {
	model := &projectv1.InputVariable{}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigSetting(modelMap map[string]interface{}) (*projectv1.ProjectConfigSetting, error) {
	model := &projectv1.ProjectConfigSetting{}
	return model, nil
}

func resourceIbmProjectProjectConfigCanonicalToMap(model *projectv1.ProjectConfigCanonical) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Authorizations != nil {
		authorizationsMap, err := resourceIbmProjectProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := resourceIbmProjectProjectConfigComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	if model.LocatorID != nil {
		modelMap["locator_id"] = model.LocatorID
	}
	if model.Input != nil {
		inputMap, err := resourceIbmProjectInputVariableToMap(model.Input)
		if err != nil {
			return modelMap, err
		}
		modelMap["input"] = []map[string]interface{}{inputMap}
	}
	if model.Setting != nil {
		settingMap, err := resourceIbmProjectProjectConfigSettingToMap(model.Setting)
		if err != nil {
			return modelMap, err
		}
		modelMap["setting"] = []map[string]interface{}{settingMap}
	}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.ProjectID != nil {
		modelMap["project_id"] = model.ProjectID
	}
	if model.Version != nil {
		modelMap["version"] = flex.IntValue(model.Version)
	}
	if model.IsDraft != nil {
		modelMap["is_draft"] = model.IsDraft
	}
	if model.NeedsAttentionState != nil {
		modelMap["needs_attention_state"] = model.NeedsAttentionState
	}
	if model.State != nil {
		modelMap["state"] = model.State
	}
	if model.UpdateAvailable != nil {
		modelMap["update_available"] = model.UpdateAvailable
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.LastApproved != nil {
		lastApprovedMap, err := resourceIbmProjectProjectConfigMetadataLastApprovedToMap(model.LastApproved)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_approved"] = []map[string]interface{}{lastApprovedMap}
	}
	if model.LastSave != nil {
		modelMap["last_save"] = model.LastSave.String()
	}
	if model.CraLogs != nil {
		craLogsMap, err := resourceIbmProjectProjectConfigMetadataCraLogsToMap(model.CraLogs)
		if err != nil {
			return modelMap, err
		}
		modelMap["cra_logs"] = []map[string]interface{}{craLogsMap}
	}
	if model.CostEstimate != nil {
		costEstimateMap, err := resourceIbmProjectProjectConfigMetadataCostEstimateToMap(model.CostEstimate)
		if err != nil {
			return modelMap, err
		}
		modelMap["cost_estimate"] = []map[string]interface{}{costEstimateMap}
	}
	if model.CheckJob != nil {
		checkJobMap, err := resourceIbmProjectActionJobWithSummaryAndHrefToMap(model.CheckJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["check_job"] = []map[string]interface{}{checkJobMap}
	}
	if model.InstallJob != nil {
		installJobMap, err := resourceIbmProjectActionJobWithSummaryAndHrefToMap(model.InstallJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["install_job"] = []map[string]interface{}{installJobMap}
	}
	if model.UninstallJob != nil {
		uninstallJobMap, err := resourceIbmProjectActionJobWithSummaryAndHrefToMap(model.UninstallJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["uninstall_job"] = []map[string]interface{}{uninstallJobMap}
	}
	if model.Output != nil {
		output := []map[string]interface{}{}
		for _, outputItem := range model.Output {
			outputItemMap, err := resourceIbmProjectOutputValueToMap(&outputItem)
			if err != nil {
				return modelMap, err
			}
			output = append(output, outputItemMap)
		}
		modelMap["output"] = output
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfile != nil {
		trustedProfileMap, err := resourceIbmProjectProjectConfigAuthTrustedProfileToMap(model.TrustedProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["trusted_profile"] = []map[string]interface{}{trustedProfileMap}
	}
	if model.Method != nil {
		modelMap["method"] = model.Method
	}
	if model.ApiKey != nil {
		modelMap["api_key"] = model.ApiKey
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigAuthTrustedProfileToMap(model *projectv1.ProjectConfigAuthTrustedProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.TargetIamID != nil {
		modelMap["target_iam_id"] = model.TargetIamID
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigComplianceProfileToMap(model *projectv1.ProjectConfigComplianceProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.InstanceID != nil {
		modelMap["instance_id"] = model.InstanceID
	}
	if model.InstanceLocation != nil {
		modelMap["instance_location"] = model.InstanceLocation
	}
	if model.AttachmentID != nil {
		modelMap["attachment_id"] = model.AttachmentID
	}
	if model.ProfileName != nil {
		modelMap["profile_name"] = model.ProfileName
	}
	return modelMap, nil
}

func resourceIbmProjectInputVariableToMap(model *projectv1.InputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func resourceIbmProjectProjectConfigSettingToMap(model *projectv1.ProjectConfigSetting) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func resourceIbmProjectProjectConfigMetadataLastApprovedToMap(model *projectv1.ProjectConfigMetadataLastApproved) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_forced"] = model.IsForced
	if model.Comment != nil {
		modelMap["comment"] = model.Comment
	}
	modelMap["timestamp"] = model.Timestamp.String()
	modelMap["user_id"] = model.UserID
	return modelMap, nil
}

func resourceIbmProjectProjectConfigMetadataCraLogsToMap(model *projectv1.ProjectConfigMetadataCraLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CraVersion != nil {
		modelMap["cra_version"] = model.CraVersion
	}
	if model.SchemaVersion != nil {
		modelMap["schema_version"] = model.SchemaVersion
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.Summary != nil {
		summary := make(map[string]interface{})
		for k, v := range model.Summary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			summary[k] = string(bytes)
		}
		modelMap["summary"] = summary
	}
	if model.Timestamp != nil {
		modelMap["timestamp"] = model.Timestamp.String()
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigMetadataCostEstimateToMap(model *projectv1.ProjectConfigMetadataCostEstimate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Version != nil {
		modelMap["version"] = model.Version
	}
	if model.Currency != nil {
		modelMap["currency"] = model.Currency
	}
	if model.TotalHourlyCost != nil {
		modelMap["total_hourly_cost"] = model.TotalHourlyCost
	}
	if model.TotalMonthlyCost != nil {
		modelMap["total_monthly_cost"] = model.TotalMonthlyCost
	}
	if model.PastTotalHourlyCost != nil {
		modelMap["past_total_hourly_cost"] = model.PastTotalHourlyCost
	}
	if model.PastTotalMonthlyCost != nil {
		modelMap["past_total_monthly_cost"] = model.PastTotalMonthlyCost
	}
	if model.DiffTotalHourlyCost != nil {
		modelMap["diff_total_hourly_cost"] = model.DiffTotalHourlyCost
	}
	if model.DiffTotalMonthlyCost != nil {
		modelMap["diff_total_monthly_cost"] = model.DiffTotalMonthlyCost
	}
	if model.TimeGenerated != nil {
		modelMap["time_generated"] = model.TimeGenerated.String()
	}
	if model.UserID != nil {
		modelMap["user_id"] = model.UserID
	}
	return modelMap, nil
}

func resourceIbmProjectActionJobWithSummaryAndHrefToMap(model *projectv1.ActionJobWithSummaryAndHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Summary != nil {
		summaryMap, err := resourceIbmProjectActionJobSummaryToMap(model.Summary)
		if err != nil {
			return modelMap, err
		}
		modelMap["summary"] = []map[string]interface{}{summaryMap}
	}
	return modelMap, nil
}

func resourceIbmProjectActionJobSummaryToMap(model *projectv1.ActionJobSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PlanSummary != nil {
		planSummary := make(map[string]interface{})
		for k, v := range model.PlanSummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			planSummary[k] = string(bytes)
		}
		modelMap["plan_summary"] = planSummary
	}
	if model.ApplySummary != nil {
		applySummary := make(map[string]interface{})
		for k, v := range model.ApplySummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			applySummary[k] = string(bytes)
		}
		modelMap["apply_summary"] = applySummary
	}
	if model.DestroySummary != nil {
		destroySummary := make(map[string]interface{})
		for k, v := range model.DestroySummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			destroySummary[k] = string(bytes)
		}
		modelMap["destroy_summary"] = destroySummary
	}
	if model.MessageSummary != nil {
		messageSummary := make(map[string]interface{})
		for k, v := range model.MessageSummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			messageSummary[k] = string(bytes)
		}
		modelMap["message_summary"] = messageSummary
	}
	if model.PlanMessages != nil {
		planMessages := make(map[string]interface{})
		for k, v := range model.PlanMessages {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			planMessages[k] = string(bytes)
		}
		modelMap["plan_messages"] = planMessages
	}
	if model.ApplyMessages != nil {
		applyMessages := make(map[string]interface{})
		for k, v := range model.ApplyMessages {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			applyMessages[k] = string(bytes)
		}
		modelMap["apply_messages"] = applyMessages
	}
	if model.DestroyMessages != nil {
		destroyMessages := make(map[string]interface{})
		for k, v := range model.DestroyMessages {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			destroyMessages[k] = string(bytes)
		}
		modelMap["destroy_messages"] = destroyMessages
	}
	return modelMap, nil
}

func resourceIbmProjectOutputValueToMap(model *projectv1.OutputValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func resourceIbmProjectCumulativeNeedsAttentionToMap(model *projectv1.CumulativeNeedsAttention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Event != nil {
		modelMap["event"] = model.Event
	}
	if model.EventID != nil {
		modelMap["event_id"] = model.EventID
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.ConfigVersion != nil {
		modelMap["config_version"] = flex.IntValue(model.ConfigVersion)
	}
	return modelMap, nil
}
