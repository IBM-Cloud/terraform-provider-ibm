// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

import (
	"context"
	"encoding/json"
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

func ResourceIbmProjectConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProjectConfigCreate,
		ReadContext:   resourceIbmProjectConfigRead,
		UpdateContext: resourceIbmProjectConfigUpdate,
		DeleteContext: resourceIbmProjectConfigDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project_config", "project_id"),
				Description:  "The unique project ID.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project_config", "name"),
				Description:  "The name of the configuration.",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// ValidateFunc: validate.InvokeValidator("ibm_project_config", "description"),
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
				Type:     schema.TypeString,
				Optional: true,
				// ValidateFunc: validate.InvokeValidator("ibm_project_config", "locator_id"),
				Description: "A dotted value of catalogID.versionID.",
			},
			"input": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The input variables for the configuration definition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"setting": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version of the configuration.",
			},
			"is_draft": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The flag that indicates whether the version of the configuration is draft, or active.",
			},
			"needs_attention_state": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The needs attention state of a configuration.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the configuration.",
			},
			"update_available": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The flag that indicates whether a configuration update is available.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"last_approved": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
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
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"cra_logs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
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
							Elem:        &schema.Schema{Type: schema.TypeString},
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
				Computed:    true,
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
				Computed:    true,
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
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"apply_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the apply jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"destroy_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the destroy jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"message_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The message summaries of jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"plan_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of plan jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"apply_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of apply jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"destroy_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of destroy jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"install_job": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
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
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"apply_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the apply jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"destroy_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the destroy jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"message_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The message summaries of jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"plan_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of plan jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"apply_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of apply jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"destroy_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of destroy jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"uninstall_job": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
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
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"apply_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the apply jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"destroy_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The summary of the destroy jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"message_summary": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The message summaries of jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"plan_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of plan jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"apply_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of apply jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"destroy_messages": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The messages of destroy jobs on the configuration.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"output": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
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
				Computed:    true,
				Description: "The type of a project configuration manual property.",
			},
			"definition": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Schematics template property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
								Schema: map[string]*schema.Schema{},
							},
						},
						"setting": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
			"project_config_canonical_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
			},
		},
	}
}

func ResourceIbmProjectConfigValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[\.\-0-9a-zA-Z]+$`,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$`,
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
		validate.ValidateSchema{
			Identifier:                 "locator_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(?!\s)(?!.*\s$)[\.0-9a-z-A-Z_-]+$`,
			MinValueLength:             1,
			MaxValueLength:             512,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_project_config", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProjectConfigCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createConfigOptions := &projectv1.CreateConfigOptions{}

	createConfigOptions.SetProjectID(d.Get("project_id").(string))
	if _, ok := d.GetOk("name"); ok {
		createConfigOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createConfigOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("labels"); ok {
		var labels []string
		for _, v := range d.Get("labels").([]interface{}) {
			labelsItem := v.(string)
			labels = append(labels, labelsItem)
		}
		createConfigOptions.SetLabels(labels)
	}
	if _, ok := d.GetOk("authorizations"); ok {
		authorizationsModel, err := resourceIbmProjectConfigMapToProjectConfigAuth(d.Get("authorizations.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createConfigOptions.SetAuthorizations(authorizationsModel)
	}
	if _, ok := d.GetOk("compliance_profile"); ok {
		complianceProfileModel, err := resourceIbmProjectConfigMapToProjectConfigComplianceProfile(d.Get("compliance_profile.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createConfigOptions.SetComplianceProfile(complianceProfileModel)
	}
	if _, ok := d.GetOk("locator_id"); ok {
		createConfigOptions.SetLocatorID(d.Get("locator_id").(string))
	}
	if _, ok := d.GetOk("input"); ok {
		inputModel, err := resourceIbmProjectConfigMapToInputVariable(d.Get("input.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createConfigOptions.SetInput(inputModel)
	}
	if _, ok := d.GetOk("setting"); ok {
		settingModel, err := resourceIbmProjectConfigMapToProjectConfigSetting(d.Get("setting.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createConfigOptions.SetSetting(settingModel)
	}

	projectConfigCanonical, response, err := projectClient.CreateConfigWithContext(context, createConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateConfigWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createConfigOptions.ProjectID, *projectConfigCanonical.ID))

	return resourceIbmProjectConfigRead(context, d, meta)
}

func resourceIbmProjectConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getConfigOptions := &projectv1.GetConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getConfigOptions.SetProjectID(parts[0])
	getConfigOptions.SetID(parts[1])

	projectConfigCanonical, response, err := projectClient.GetConfigWithContext(context, getConfigOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetConfigWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("project_id", projectConfigCanonical.ProjectID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_id: %s", err))
	}
	if !core.IsNil(projectConfigCanonical.Name) {
		if err = d.Set("name", projectConfigCanonical.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Description) {
		if err = d.Set("description", projectConfigCanonical.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Labels) {
		if err = d.Set("labels", projectConfigCanonical.Labels); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting labels: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Authorizations) {
		authorizationsMap, err := resourceIbmProjectConfigProjectConfigAuthToMap(projectConfigCanonical.Authorizations)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("authorizations", []map[string]interface{}{authorizationsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting authorizations: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.ComplianceProfile) {
		complianceProfileMap, err := resourceIbmProjectConfigProjectConfigComplianceProfileToMap(projectConfigCanonical.ComplianceProfile)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("compliance_profile", []map[string]interface{}{complianceProfileMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting compliance_profile: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.LocatorID) {
		if err = d.Set("locator_id", projectConfigCanonical.LocatorID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting locator_id: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Input) {
		inputMap, err := resourceIbmProjectConfigInputVariableToMap(projectConfigCanonical.Input)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("input", []map[string]interface{}{inputMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting input: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Setting) {
		settingMap, err := resourceIbmProjectConfigProjectConfigSettingToMap(projectConfigCanonical.Setting)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("setting", []map[string]interface{}{settingMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting setting: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Version) {
		if err = d.Set("version", flex.IntValue(projectConfigCanonical.Version)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.IsDraft) {
		if err = d.Set("is_draft", projectConfigCanonical.IsDraft); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting is_draft: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.NeedsAttentionState) {
		needsAttentionState := []interface{}{}
		for _, needsAttentionStateItem := range projectConfigCanonical.NeedsAttentionState {
			needsAttentionState = append(needsAttentionState, needsAttentionStateItem)
		}
		if err = d.Set("needs_attention_state", needsAttentionState); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting needs_attention_state: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.State) {
		if err = d.Set("state", projectConfigCanonical.State); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.UpdateAvailable) {
		if err = d.Set("update_available", projectConfigCanonical.UpdateAvailable); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting update_available: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(projectConfigCanonical.CreatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(projectConfigCanonical.UpdatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.LastApproved) {
		lastApprovedMap, err := resourceIbmProjectConfigProjectConfigMetadataLastApprovedToMap(projectConfigCanonical.LastApproved)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("last_approved", []map[string]interface{}{lastApprovedMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_approved: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.LastSave) {
		if err = d.Set("last_save", flex.DateTimeToString(projectConfigCanonical.LastSave)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_save: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.CraLogs) {
		craLogsMap, err := resourceIbmProjectConfigProjectConfigMetadataCraLogsToMap(projectConfigCanonical.CraLogs)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("cra_logs", []map[string]interface{}{craLogsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cra_logs: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.CostEstimate) {
		costEstimateMap, err := resourceIbmProjectConfigProjectConfigMetadataCostEstimateToMap(projectConfigCanonical.CostEstimate)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("cost_estimate", []map[string]interface{}{costEstimateMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cost_estimate: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.CheckJob) {
		checkJobMap, err := resourceIbmProjectConfigActionJobWithSummaryAndHrefToMap(projectConfigCanonical.CheckJob)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("check_job", []map[string]interface{}{checkJobMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting check_job: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.InstallJob) {
		installJobMap, err := resourceIbmProjectConfigActionJobWithSummaryAndHrefToMap(projectConfigCanonical.InstallJob)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("install_job", []map[string]interface{}{installJobMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting install_job: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.UninstallJob) {
		uninstallJobMap, err := resourceIbmProjectConfigActionJobWithSummaryAndHrefToMap(projectConfigCanonical.UninstallJob)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("uninstall_job", []map[string]interface{}{uninstallJobMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting uninstall_job: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Output) {
		output := []map[string]interface{}{}
		for _, outputItem := range projectConfigCanonical.Output {
			outputItemMap, err := resourceIbmProjectConfigOutputValueToMap(&outputItem)
			if err != nil {
				return diag.FromErr(err)
			}
			output = append(output, outputItemMap)
		}
		if err = d.Set("output", output); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting output: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Type) {
		if err = d.Set("type", projectConfigCanonical.Type); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.Definition) {
		definitionMap, err := resourceIbmProjectConfigProjectConfigDefinitionResponseTerraformToMap(projectConfigCanonical.Definition)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("definition", []map[string]interface{}{definitionMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting definition: %s", err))
		}
	}
	if !core.IsNil(projectConfigCanonical.ID) {
		if err = d.Set("project_config_canonical_id", projectConfigCanonical.ID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting project_config_canonical_id: %s", err))
		}
	}

	return nil
}

func resourceIbmProjectConfigUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateConfigOptions := &projectv1.UpdateConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	updateConfigOptions.SetProjectID(parts[0])
	updateConfigOptions.SetID(parts[1])

	hasChange := false

	if d.HasChange("project_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "project_id"))
	}
	if d.HasChange("name") {
		updateConfigOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("description") {
		updateConfigOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("labels") {
		var labels []string
		for _, v := range d.Get("labels").([]interface{}) {
			labelsItem := v.(string)
			labels = append(labels, labelsItem)
		}
		updateConfigOptions.SetLabels(labels)
		hasChange = true
	}
	if d.HasChange("authorizations") {
		authorizations, err := resourceIbmProjectConfigMapToProjectConfigAuth(d.Get("authorizations.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateConfigOptions.SetAuthorizations(authorizations)
		hasChange = true
	}
	if d.HasChange("compliance_profile") {
		complianceProfile, err := resourceIbmProjectConfigMapToProjectConfigComplianceProfile(d.Get("compliance_profile.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateConfigOptions.SetComplianceProfile(complianceProfile)
		hasChange = true
	}
	if d.HasChange("locator_id") {
		updateConfigOptions.SetLocatorID(d.Get("locator_id").(string))
		hasChange = true
	}
	if d.HasChange("input") {
		input, err := resourceIbmProjectConfigMapToInputVariable(d.Get("input.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateConfigOptions.SetInput(input)
		hasChange = true
	}
	if d.HasChange("setting") {
		setting, err := resourceIbmProjectConfigMapToProjectConfigSetting(d.Get("setting.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateConfigOptions.SetSetting(setting)
		hasChange = true
	}

	if hasChange {
		_, response, err := projectClient.UpdateConfigWithContext(context, updateConfigOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateConfigWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateConfigWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmProjectConfigRead(context, d, meta)
}

func resourceIbmProjectConfigDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteConfigOptions := &projectv1.DeleteConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteConfigOptions.SetProjectID(parts[0])
	deleteConfigOptions.SetID(parts[1])

	_, response, err := projectClient.DeleteConfigWithContext(context, deleteConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteConfigWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmProjectConfigMapToProjectConfigAuth(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuth, error) {
	model := &projectv1.ProjectConfigAuth{}
	if modelMap["trusted_profile"] != nil && len(modelMap["trusted_profile"].([]interface{})) > 0 {
		TrustedProfileModel, err := resourceIbmProjectConfigMapToProjectConfigAuthTrustedProfile(modelMap["trusted_profile"].([]interface{})[0].(map[string]interface{}))
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

func resourceIbmProjectConfigMapToProjectConfigAuthTrustedProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuthTrustedProfile, error) {
	model := &projectv1.ProjectConfigAuthTrustedProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["target_iam_id"] != nil && modelMap["target_iam_id"].(string) != "" {
		model.TargetIamID = core.StringPtr(modelMap["target_iam_id"].(string))
	}
	return model, nil
}

func resourceIbmProjectConfigMapToProjectConfigComplianceProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigComplianceProfile, error) {
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

func resourceIbmProjectConfigMapToInputVariable(modelMap map[string]interface{}) (*projectv1.InputVariable, error) {
	model := &projectv1.InputVariable{}
	return model, nil
}

func resourceIbmProjectConfigMapToProjectConfigSetting(modelMap map[string]interface{}) (*projectv1.ProjectConfigSetting, error) {
	model := &projectv1.ProjectConfigSetting{}
	return model, nil
}

func resourceIbmProjectConfigProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfile != nil {
		trustedProfileMap, err := resourceIbmProjectConfigProjectConfigAuthTrustedProfileToMap(model.TrustedProfile)
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

func resourceIbmProjectConfigProjectConfigAuthTrustedProfileToMap(model *projectv1.ProjectConfigAuthTrustedProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.TargetIamID != nil {
		modelMap["target_iam_id"] = model.TargetIamID
	}
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigComplianceProfileToMap(model *projectv1.ProjectConfigComplianceProfile) (map[string]interface{}, error) {
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

func resourceIbmProjectConfigInputVariableToMap(model *projectv1.InputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigSettingToMap(model *projectv1.ProjectConfigSetting) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigMetadataLastApprovedToMap(model *projectv1.ProjectConfigMetadataLastApproved) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_forced"] = model.IsForced
	if model.Comment != nil {
		modelMap["comment"] = model.Comment
	}
	modelMap["timestamp"] = model.Timestamp.String()
	modelMap["user_id"] = model.UserID
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigMetadataCraLogsToMap(model *projectv1.ProjectConfigMetadataCraLogs) (map[string]interface{}, error) {
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

func resourceIbmProjectConfigProjectConfigMetadataCostEstimateToMap(model *projectv1.ProjectConfigMetadataCostEstimate) (map[string]interface{}, error) {
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

func resourceIbmProjectConfigActionJobWithSummaryAndHrefToMap(model *projectv1.ActionJobWithSummaryAndHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Summary != nil {
		summaryMap, err := resourceIbmProjectConfigActionJobSummaryToMap(model.Summary)
		if err != nil {
			return modelMap, err
		}
		modelMap["summary"] = []map[string]interface{}{summaryMap}
	}
	return modelMap, nil
}

func resourceIbmProjectConfigActionJobSummaryToMap(model *projectv1.ActionJobSummary) (map[string]interface{}, error) {
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

func resourceIbmProjectConfigOutputValueToMap(model *projectv1.OutputValue) (map[string]interface{}, error) {
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

func resourceIbmProjectConfigProjectConfigDefinitionResponseTerraformToMap(model *projectv1.ProjectConfigDefinitionResponseTerraform) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Output != nil {
		output := []map[string]interface{}{}
		for _, outputItem := range model.Output {
			outputItemMap, err := resourceIbmProjectConfigOutputValueToMap(&outputItem)
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
		authorizationsMap, err := resourceIbmProjectConfigProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := resourceIbmProjectConfigProjectConfigComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	if model.LocatorID != nil {
		modelMap["locator_id"] = model.LocatorID
	}
	if model.Input != nil {
		inputMap, err := resourceIbmProjectConfigInputVariableToMap(model.Input)
		if err != nil {
			return modelMap, err
		}
		modelMap["input"] = []map[string]interface{}{inputMap}
	}
	if model.Setting != nil {
		settingMap, err := resourceIbmProjectConfigProjectConfigSettingToMap(model.Setting)
		if err != nil {
			return modelMap, err
		}
		modelMap["setting"] = []map[string]interface{}{settingMap}
	}
	return modelMap, nil
}
