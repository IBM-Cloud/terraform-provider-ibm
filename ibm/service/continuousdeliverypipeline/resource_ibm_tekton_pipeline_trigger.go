// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package continuousdeliverypipeline

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
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"
)

func ResourceIBMTektonPipelineTrigger() *schema.Resource {
	return &schema.Resource{
		CreateContext:   ResourceIBMTektonPipelineTriggerCreate,
		ReadContext:     ResourceIBMTektonPipelineTriggerRead,
		UpdateContext:   ResourceIBMTektonPipelineTriggerUpdate,
		DeleteContext:   ResourceIBMTektonPipelineTriggerDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_tekton_pipeline_trigger", "pipeline_id"),
				Description: "The tekton pipeline ID.",
			},
			"create_tekton_pipeline_trigger_request": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trigger type.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trigger name.",
						},
						"event_listener": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Event listener name.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UUID.",
						},
						"properties": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Trigger properties.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Property name.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "String format property value.",
									},
									"options": &schema.Schema{
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Options for SINGLE_SELECT property type.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Property type.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "property path for INTEGRATION type properties.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "General href URL.",
									},
								},
							},
						},
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Trigger tags array.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"worker": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Trigger worker used to run the trigger, the trigger worker overrides the default pipeline worker.If not exist, this trigger uses default pipeline worker.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "worker name.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "worker type.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "ID.",
									},
								},
							},
						},
						"concurrency": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Concurrency object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_concurrent_runs": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Defines the maximum number of concurrent runs for this trigger.",
									},
								},
							},
						},
						"disabled": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Defines if this trigger is disabled.",
						},
						"scm_source": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Scm source for git type tekton pipeline trigger.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Needed only for git trigger type. Repo URL that listening to.",
									},
									"branch": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Needed only for git trigger type. Branch name of the repo.",
									},
									"blind_connection": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Needed only for git trigger type. Branch name of the repo.",
									},
									"hook_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Webhook Id.",
									},
								},
							},
						},
						"events": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Needed only for git trigger type. Events object defines the events this git trigger listening to.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"push": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger will start when a 'push' event received.",
									},
									"pull_request_closed": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger will start when a pull request 'close' event received.",
									},
									"pull_request": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger will start when a pull request 'open' or 'update' event received.",
									},
								},
							},
						},
						"service_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UUID.",
						},
						"cron": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Needed only for timer trigger type. Cron expression for timer trigger.",
						},
						"timezone": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Needed only for timer trigger type. Timezones for timer trigger.",
						},
						"secret": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Needed only for generic trigger type. Secret used to execute generic trigger.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret type.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret value, not needed if secret type is \"internalValidation\".",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret location, not needed if secret type is \"internalValidation\".",
									},
									"key_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret name, not needed if type is \"internalValidation\".",
									},
									"algorithm": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Algorithm used for \"digestMatches\" secret type.",
									},
								},
							},
						},
						"triggers": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Tekton pipeline triggers list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Trigger type.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Trigger name.",
									},
									"event_listener": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Event listener name.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "UUID.",
									},
									"properties": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Trigger properties.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Property name.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "String format property value.",
												},
												"options": &schema.Schema{
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "Options for SINGLE_SELECT property type.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "Property type.",
												},
												"path": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "property path for INTEGRATION type properties.",
												},
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "General href URL.",
												},
											},
										},
									},
									"tags": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Trigger tags array.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"worker": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Trigger worker used to run the trigger, the trigger worker overrides the default pipeline worker.If not exist, this trigger uses default pipeline worker.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "worker name.",
												},
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "worker type.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "ID.",
												},
											},
										},
									},
									"concurrency": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Concurrency object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"max_concurrent_runs": &schema.Schema{
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Defines the maximum number of concurrent runs for this trigger.",
												},
											},
										},
									},
									"disabled": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Defines if this trigger is disabled.",
									},
									"scm_source": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Scm source for git type tekton pipeline trigger.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"url": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Needed only for git trigger type. Repo URL that listening to.",
												},
												"branch": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Needed only for git trigger type. Branch name of the repo.",
												},
												"blind_connection": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Needed only for git trigger type. Branch name of the repo.",
												},
												"hook_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Webhook Id.",
												},
											},
										},
									},
									"events": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Needed only for git trigger type. Events object defines the events this git trigger listening to.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"push": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If true, the trigger will start when a 'push' event received.",
												},
												"pull_request_closed": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If true, the trigger will start when a pull request 'close' event received.",
												},
												"pull_request": &schema.Schema{
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "If true, the trigger will start when a pull request 'open' or 'update' event received.",
												},
											},
										},
									},
									"service_instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "UUID.",
									},
									"cron": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Needed only for timer trigger type. Cron expression for timer trigger.",
									},
									"timezone": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Needed only for timer trigger type. Timezones for timer trigger.",
									},
									"secret": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Needed only for generic trigger type. Secret used to execute generic trigger.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Secret type.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Secret value, not needed if secret type is \"internalValidation\".",
												},
												"source": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Secret location, not needed if secret type is \"internalValidation\".",
												},
												"key_name": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Secret name, not needed if type is \"internalValidation\".",
												},
												"algorithm": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Algorithm used for \"digestMatches\" secret type.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "General href URL.",
									},
								},
							},
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trigger type.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trigger name.",
			},
			"event_listener": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Event listener name.",
			},
			"properties": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Trigger properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Property name.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "String format property value.",
						},
						"options": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Options for SINGLE_SELECT property type.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Property type.",
						},
						"path": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "property path for INTEGRATION type properties.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "General href URL.",
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Trigger tags array.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"worker": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Trigger worker used to run the trigger, the trigger worker overrides the default pipeline worker.If not exist, this trigger uses default pipeline worker.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "worker name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "worker type.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID.",
						},
					},
				},
			},
			"concurrency": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Concurrency object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_concurrent_runs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Defines the maximum number of concurrent runs for this trigger.",
						},
					},
				},
			},
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Defines if this trigger is disabled.",
			},
			"scm_source": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Scm source for git type tekton pipeline trigger.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Needed only for git trigger type. Repo URL that listening to.",
						},
						"branch": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Needed only for git trigger type. Branch name of the repo.",
						},
						"blind_connection": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Needed only for git trigger type. Branch name of the repo.",
						},
						"hook_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Webhook Id.",
						},
					},
				},
			},
			"events": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Needed only for git trigger type. Events object defines the events this git trigger listening to.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"push": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, the trigger will start when a 'push' event received.",
						},
						"pull_request_closed": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, the trigger will start when a pull request 'close' event received.",
						},
						"pull_request": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, the trigger will start when a pull request 'open' or 'update' event received.",
						},
					},
				},
			},
			"service_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID.",
			},
			"cron": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Needed only for timer trigger type. Cron expression for timer trigger.",
			},
			"timezone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Needed only for timer trigger type. Timezones for timer trigger.",
			},
			"secret": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Needed only for generic trigger type. Secret used to execute generic trigger.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secret type.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secret value, not needed if secret type is \"internalValidation\".",
						},
						"source": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secret location, not needed if secret type is \"internalValidation\".",
						},
						"key_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secret name, not needed if type is \"internalValidation\".",
						},
						"algorithm": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Algorithm used for \"digestMatches\" secret type.",
						},
					},
				},
			},
			"trigger_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID.",
			},
		},
	}
}

func ResourceIBMTektonPipelineTriggerValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "pipeline_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z]+$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_tekton_pipeline_trigger", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMTektonPipelineTriggerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelineTriggerOptions := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerOptions{}

	createTektonPipelineTriggerOptions.SetPipelineID(d.Get("pipeline_id").(string))
	if _, ok := d.GetOk("create_tekton_pipeline_trigger_request"); ok {
		createTektonPipelineTriggerRequest, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequest(d.Get("create_tekton_pipeline_trigger_request.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerOptions.SetCreateTektonPipelineTriggerRequest(createTektonPipelineTriggerRequest)
	}

	triggers, response, err := continuousDeliveryPipelineClient.CreateTektonPipelineTriggerWithContext(context, createTektonPipelineTriggerOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTektonPipelineTriggerWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTektonPipelineTriggerWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createTektonPipelineTriggerOptions.PipelineID, *trigger.ID))

	return ResourceIBMTektonPipelineTriggerRead(context, d, meta)
}

func ResourceIBMTektonPipelineTriggerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineTriggerOptions := &continuousdeliverypipelinev2.GetTektonPipelineTriggerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineTriggerOptions.SetPipelineID(parts[0])
	getTektonPipelineTriggerOptions.SetTriggerID(parts[1])

	trigger, response, err := continuousDeliveryPipelineClient.GetTektonPipelineTriggerWithContext(context, getTektonPipelineTriggerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTektonPipelineTriggerWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelineTriggerWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("pipeline_id", getTektonPipelineTriggerOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	// TODO: handle argument of type CreateTektonPipelineTriggerRequest
	if err = d.Set("type", trigger.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("name", trigger.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("event_listener", trigger.EventListener); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting event_listener: %s", err))
	}
	
	properties := []map[string]interface{}{}
	if trigger.Properties != nil {
		for _, propertiesItem := range trigger.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			properties = append(properties, propertiesItemMap)
		}
	}
	if err = d.Set("properties", properties); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting properties: %s", err))
	}
	
	if trigger.Tags != nil {
		if err = d.Set("tags", trigger.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	if trigger.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(trigger.Worker)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("worker", []map[string]interface{}{workerMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting worker: %s", err))
		}
	}
	if trigger.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerTriggerConcurrencyToMap(trigger.Concurrency)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("concurrency", []map[string]interface{}{concurrencyMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting concurrency: %s", err))
		}
	}
	if err = d.Set("disabled", trigger.Disabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disabled: %s", err))
	}
	if trigger.ScmSource != nil {
		scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(trigger.ScmSource)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("scm_source", []map[string]interface{}{scmSourceMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scm_source: %s", err))
		}
	}
	if trigger.Events != nil {
		eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(trigger.Events)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("events", []map[string]interface{}{eventsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting events: %s", err))
		}
	}
	if err = d.Set("service_instance_id", trigger.ServiceInstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting service_instance_id: %s", err))
	}
	if err = d.Set("cron", trigger.Cron); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cron: %s", err))
	}
	if err = d.Set("timezone", trigger.Timezone); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting timezone: %s", err))
	}
	if trigger.Secret != nil {
		secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(trigger.Secret)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("secret", []map[string]interface{}{secretMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting secret: %s", err))
		}
	}
	if err = d.Set("trigger_id", trigger.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting trigger_id: %s", err))
	}

	return nil
}

func ResourceIBMTektonPipelineTriggerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	updateTektonPipelineTriggerOptions := &continuousdeliverypipelinev2.UpdateTektonPipelineTriggerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	updateTektonPipelineTriggerOptions.SetPipelineID(parts[0])
	updateTektonPipelineTriggerOptions.SetTriggerID(parts[1])

	hasChange := false

	if d.HasChange("pipeline_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation." +
				" The resource must be re-created to update this property.", "pipeline_id"))
	}
	if d.HasChange("create_tekton_pipeline_trigger_request") {
		createTektonPipelineTriggerRequest, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequest(d.Get("create_tekton_pipeline_trigger_request.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateTektonPipelineTriggerOptions.SetCreateTektonPipelineTriggerRequest(createTektonPipelineTriggerRequest)
		hasChange = true
	}

	if hasChange {
		_, response, err := continuousDeliveryPipelineClient.UpdateTektonPipelineTriggerWithContext(context, updateTektonPipelineTriggerOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateTektonPipelineTriggerWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateTektonPipelineTriggerWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMTektonPipelineTriggerRead(context, d, meta)
}

func ResourceIBMTektonPipelineTriggerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineTriggerOptions := &continuousdeliverypipelinev2.DeleteTektonPipelineTriggerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineTriggerOptions.SetPipelineID(parts[0])
	deleteTektonPipelineTriggerOptions.SetTriggerID(parts[1])

	response, err := continuousDeliveryPipelineClient.DeleteTektonPipelineTriggerWithContext(context, deleteTektonPipelineTriggerOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTektonPipelineTriggerWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTektonPipelineTriggerWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequest(modelMap map[string]interface{}) (continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestIntf, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequest{}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["name"] != nil {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["event_listener"] != nil {
		model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	}
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	if modelMap["disabled"] != nil {
		model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	}
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	if modelMap["triggers"] != nil {
		triggers := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemIntf{}
		for _, triggersItem := range modelMap["triggers"].([]interface{}) {
			triggersItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItem(triggersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			triggers = append(triggers, triggersItemModel)
		}
		model.Triggers = triggers
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToWorker(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.Worker, error) {
	model := &continuousdeliverypipelinev2.Worker{}
	if modelMap["name"] != nil {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.TriggerScmSource, error) {
	model := &continuousdeliverypipelinev2.TriggerScmSource{}
	if modelMap["url"] != nil {
		model.URL = core.StringPtr(modelMap["url"].(string))
	}
	if modelMap["branch"] != nil {
		model.Branch = core.StringPtr(modelMap["branch"].(string))
	}
	if modelMap["blind_connection"] != nil {
		model.BlindConnection = core.BoolPtr(modelMap["blind_connection"].(bool))
	}
	if modelMap["hook_id"] != nil {
		model.HookID = core.StringPtr(modelMap["hook_id"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToEvents(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.Events, error) {
	model := &continuousdeliverypipelinev2.Events{}
	if modelMap["push"] != nil {
		model.Push = core.BoolPtr(modelMap["push"].(bool))
	}
	if modelMap["pull_request_closed"] != nil {
		model.PullRequestClosed = core.BoolPtr(modelMap["pull_request_closed"].(bool))
	}
	if modelMap["pull_request"] != nil {
		model.PullRequest = core.BoolPtr(modelMap["pull_request"].(bool))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.GenericSecret, error) {
	model := &continuousdeliverypipelinev2.GenericSecret{}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["source"] != nil {
		model.Source = core.StringPtr(modelMap["source"].(string))
	}
	if modelMap["key_name"] != nil {
		model.KeyName = core.StringPtr(modelMap["key_name"].(string))
	}
	if modelMap["algorithm"] != nil {
		model.Algorithm = core.StringPtr(modelMap["algorithm"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItem(modelMap map[string]interface{}) (continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemIntf, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItem{}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["name"] != nil {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["event_listener"] != nil {
		model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	}
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	if modelMap["disabled"] != nil {
		model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	}
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTrigger(modelMap map[string]interface{}) (continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerIntf, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTrigger{}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["name"] != nil {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["event_listener"] != nil {
		model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	}
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	if modelMap["disabled"] != nil {
		model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	}
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerManualTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerScmTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerTimerTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerGenericTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggers(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggers, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggers{}
	triggers := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemIntf{}
	for _, triggersItem := range modelMap["triggers"].([]interface{}) {
		triggersItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItem(triggersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		triggers = append(triggers, triggersItemModel)
	}
	model.Triggers = triggers
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItem(modelMap map[string]interface{}) (continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemIntf, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItem{}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["name"] != nil {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["event_listener"] != nil {
		model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	}
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.TriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	if modelMap["disabled"] != nil {
		model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	}
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.TriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.TriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.TriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.TriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["scm_source"] != nil {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["cron"] != nil {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTrigger(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTrigger, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTrigger{}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.TriggerGenericTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerGenericTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerGenericTriggerConcurrency(modelMap["concurrency"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["secret"] != nil {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerGenericTriggerPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.TriggerGenericTriggerPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.TriggerGenericTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerGenericTriggerConcurrency(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.TriggerGenericTriggerConcurrency, error) {
	model := &continuousdeliverypipelinev2.TriggerGenericTriggerConcurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestToMap(model continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestIntf) (map[string]interface{}, error) {
	if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggers); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggers))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequest); ok {
		modelMap := make(map[string]interface{})
		model := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequest)
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.EventListener != nil {
			modelMap["event_listener"] = model.EventListener
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Properties != nil {
			properties := []map[string]interface{}{}
			for _, propertiesItem := range model.Properties {
				propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestPropertiesItemToMap(&propertiesItem)
				if err != nil {
					return modelMap, err
				}
				properties = append(properties, propertiesItemMap)
			}
			modelMap["properties"] = properties
		}
		if model.Tags != nil {
			modelMap["tags"] = model.Tags
		}
		if model.Worker != nil {
			workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.Concurrency != nil {
			concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestConcurrencyToMap(model.Concurrency)
			if err != nil {
				return modelMap, err
			}
			modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
		}
		if model.Disabled != nil {
			modelMap["disabled"] = model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
			if err != nil {
				return modelMap, err
			}
			modelMap["events"] = []map[string]interface{}{eventsMap}
		}
		if model.ServiceInstanceID != nil {
			modelMap["service_instance_id"] = model.ServiceInstanceID
		}
		if model.Cron != nil {
			modelMap["cron"] = model.Cron
		}
		if model.Timezone != nil {
			modelMap["timezone"] = model.Timezone
		}
		if model.Secret != nil {
			secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
			if err != nil {
				return modelMap, err
			}
			modelMap["secret"] = []map[string]interface{}{secretMap}
		}
		if model.Triggers != nil {
			triggers := []map[string]interface{}{}
			for _, triggersItem := range model.Triggers {
				triggersItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemToMap(triggersItem)
				if err != nil {
					return modelMap, err
				}
				triggers = append(triggers, triggersItemMap)
			}
			modelMap["triggers"] = triggers
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerWorkerToMap(model *continuousdeliverypipelinev2.Worker) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	modelMap["id"] = model.ID
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model *continuousdeliverypipelinev2.TriggerScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.URL != nil {
		modelMap["url"] = model.URL
	}
	if model.Branch != nil {
		modelMap["branch"] = model.Branch
	}
	if model.BlindConnection != nil {
		modelMap["blind_connection"] = model.BlindConnection
	}
	if model.HookID != nil {
		modelMap["hook_id"] = model.HookID
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerEventsToMap(model *continuousdeliverypipelinev2.Events) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Push != nil {
		modelMap["push"] = model.Push
	}
	if model.PullRequestClosed != nil {
		modelMap["pull_request_closed"] = model.PullRequestClosed
	}
	if model.PullRequest != nil {
		modelMap["pull_request"] = model.PullRequest
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerGenericSecretToMap(model *continuousdeliverypipelinev2.GenericSecret) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Source != nil {
		modelMap["source"] = model.Source
	}
	if model.KeyName != nil {
		modelMap["key_name"] = model.KeyName
	}
	if model.Algorithm != nil {
		modelMap["algorithm"] = model.Algorithm
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemToMap(model continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemIntf) (map[string]interface{}, error) {
	if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItem); ok {
		modelMap := make(map[string]interface{})
		model := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItem)
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.EventListener != nil {
			modelMap["event_listener"] = model.EventListener
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Properties != nil {
			properties := []map[string]interface{}{}
			for _, propertiesItem := range model.Properties {
				propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemPropertiesItemToMap(&propertiesItem)
				if err != nil {
					return modelMap, err
				}
				properties = append(properties, propertiesItemMap)
			}
			modelMap["properties"] = properties
		}
		if model.Tags != nil {
			modelMap["tags"] = model.Tags
		}
		if model.Worker != nil {
			workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.Concurrency != nil {
			concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemConcurrencyToMap(model.Concurrency)
			if err != nil {
				return modelMap, err
			}
			modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
		}
		if model.Disabled != nil {
			modelMap["disabled"] = model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
			if err != nil {
				return modelMap, err
			}
			modelMap["events"] = []map[string]interface{}{eventsMap}
		}
		if model.ServiceInstanceID != nil {
			modelMap["service_instance_id"] = model.ServiceInstanceID
		}
		if model.Cron != nil {
			modelMap["cron"] = model.Cron
		}
		if model.Timezone != nil {
			modelMap["timezone"] = model.Timezone
		}
		if model.Secret != nil {
			secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
			if err != nil {
				return modelMap, err
			}
			modelMap["secret"] = []map[string]interface{}{secretMap}
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerManualTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.ScmSource != nil {
		scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.Events != nil {
		eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
		if err != nil {
			return modelMap, err
		}
		modelMap["events"] = []map[string]interface{}{eventsMap}
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = model.ServiceInstanceID
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerScmTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Cron != nil {
		modelMap["cron"] = model.Cron
	}
	if model.Timezone != nil {
		modelMap["timezone"] = model.Timezone
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerTimerTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Secret != nil {
		secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
		if err != nil {
			return modelMap, err
		}
		modelMap["secret"] = []map[string]interface{}{secretMap}
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersItemTriggerGenericTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerToMap(model continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTrigger); ok {
		modelMap := make(map[string]interface{})
		model := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTrigger)
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.EventListener != nil {
			modelMap["event_listener"] = model.EventListener
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Properties != nil {
			properties := []map[string]interface{}{}
			for _, propertiesItem := range model.Properties {
				propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerPropertiesItemToMap(&propertiesItem)
				if err != nil {
					return modelMap, err
				}
				properties = append(properties, propertiesItemMap)
			}
			modelMap["properties"] = properties
		}
		if model.Tags != nil {
			modelMap["tags"] = model.Tags
		}
		if model.Worker != nil {
			workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.Concurrency != nil {
			concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerConcurrencyToMap(model.Concurrency)
			if err != nil {
				return modelMap, err
			}
			modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
		}
		if model.Disabled != nil {
			modelMap["disabled"] = model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
			if err != nil {
				return modelMap, err
			}
			modelMap["events"] = []map[string]interface{}{eventsMap}
		}
		if model.ServiceInstanceID != nil {
			modelMap["service_instance_id"] = model.ServiceInstanceID
		}
		if model.Cron != nil {
			modelMap["cron"] = model.Cron
		}
		if model.Timezone != nil {
			modelMap["timezone"] = model.Timezone
		}
		if model.Secret != nil {
			secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
			if err != nil {
				return modelMap, err
			}
			modelMap["secret"] = []map[string]interface{}{secretMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerManualTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.ScmSource != nil {
		scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.Events != nil {
		eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
		if err != nil {
			return modelMap, err
		}
		modelMap["events"] = []map[string]interface{}{eventsMap}
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = model.ServiceInstanceID
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerScmTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Cron != nil {
		modelMap["cron"] = model.Cron
	}
	if model.Timezone != nil {
		modelMap["timezone"] = model.Timezone
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerTimerTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Secret != nil {
		secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
		if err != nil {
			return modelMap, err
		}
		modelMap["secret"] = []map[string]interface{}{secretMap}
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggerTriggerGenericTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggers) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	triggers := []map[string]interface{}{}
	for _, triggersItem := range model.Triggers {
		triggersItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemToMap(triggersItem)
		if err != nil {
			return modelMap, err
		}
		triggers = append(triggers, triggersItemMap)
	}
	modelMap["triggers"] = triggers
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemToMap(model continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemIntf) (map[string]interface{}, error) {
	if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTrigger); ok {
		return ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTriggerToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTrigger))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItem); ok {
		modelMap := make(map[string]interface{})
		model := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItem)
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.EventListener != nil {
			modelMap["event_listener"] = model.EventListener
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Properties != nil {
			properties := []map[string]interface{}{}
			for _, propertiesItem := range model.Properties {
				propertiesItemMap, err := ResourceIBMTektonPipelineTriggerTriggerPropertiesItemToMap(&propertiesItem)
				if err != nil {
					return modelMap, err
				}
				properties = append(properties, propertiesItemMap)
			}
			modelMap["properties"] = properties
		}
		if model.Tags != nil {
			modelMap["tags"] = model.Tags
		}
		if model.Worker != nil {
			workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.Concurrency != nil {
			concurrencyMap, err := ResourceIBMTektonPipelineTriggerTriggerConcurrencyToMap(model.Concurrency)
			if err != nil {
				return modelMap, err
			}
			modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
		}
		if model.Disabled != nil {
			modelMap["disabled"] = model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
			if err != nil {
				return modelMap, err
			}
			modelMap["events"] = []map[string]interface{}{eventsMap}
		}
		if model.ServiceInstanceID != nil {
			modelMap["service_instance_id"] = model.ServiceInstanceID
		}
		if model.Cron != nil {
			modelMap["cron"] = model.Cron
		}
		if model.Timezone != nil {
			modelMap["timezone"] = model.Timezone
		}
		if model.Secret != nil {
			secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
			if err != nil {
				return modelMap, err
			}
			modelMap["secret"] = []map[string]interface{}{secretMap}
		}
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelineTriggerTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.TriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.TriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerManualTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.ScmSource != nil {
		scmSourceMap, err := ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.Events != nil {
		eventsMap, err := ResourceIBMTektonPipelineTriggerEventsToMap(model.Events)
		if err != nil {
			return modelMap, err
		}
		modelMap["events"] = []map[string]interface{}{eventsMap}
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = model.ServiceInstanceID
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Cron != nil {
		modelMap["cron"] = model.Cron
	}
	if model.Timezone != nil {
		modelMap["timezone"] = model.Timezone
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerTimerTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerCreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTriggerToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerRequestTriggersTriggersItemTriggerGenericTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerTriggerGenericTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return modelMap, err
			}
			properties = append(properties, propertiesItemMap)
		}
		modelMap["properties"] = properties
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := ResourceIBMTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerTriggerGenericTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Secret != nil {
		secretMap, err := ResourceIBMTektonPipelineTriggerGenericSecretToMap(model.Secret)
		if err != nil {
			return modelMap, err
		}
		modelMap["secret"] = []map[string]interface{}{secretMap}
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerGenericTriggerPropertiesItemToMap(model *continuousdeliverypipelinev2.TriggerGenericTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Options != nil {
		modelMap["options"] = model.Options
	}
	modelMap["type"] = model.Type
	if model.Path != nil {
		modelMap["path"] = model.Path
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerGenericTriggerConcurrencyToMap(model *continuousdeliverypipelinev2.TriggerGenericTriggerConcurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}
