// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline

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
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/cdtektonpipelinev2"
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
			"trigger": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Tekton pipeline trigger object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_trigger_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "source trigger ID to clone from.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "name of the duplicated trigger.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trigger type.",
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
									"enum": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Options for SINGLE_SELECT property type.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"default": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Default option for SINGLE_SELECT property type.",
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
										Computed:    true,
										Description: "worker name.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "worker type.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
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
							Description: "flag whether the trigger is disabled.",
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
										Description: "Needed only for git trigger type. Branch name of the repo. Branch field doesn't coexist with pattern field.",
									},
									"pattern": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Needed only for git trigger type. Git branch or tag pattern to listen to. Please refer to https://github.com/micromatch/micromatch for pattern syntax.",
									},
									"blind_connection": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Needed only for git trigger type. Branch name of the repo.",
									},
									"hook_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Webhook ID.",
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
										Description: "If true, the trigger starts when tekton pipeline service receive a repo's 'push' git webhook event.",
									},
									"pull_request_closed": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger starts when tekton pipeline service receive a repo pull reqeust's 'close' git webhook event.",
									},
									"pull_request": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger starts when tekton pipeline service receive a repo pull reqeust's 'open' or 'update' git webhook event.",
									},
								},
							},
						},
						"service_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "UUID.",
						},
						"cron": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Needed only for timer trigger type. Cron expression for timer trigger. Maximum frequency is every 5 minutes.",
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
							Description: "Needed only for generic trigger type. Secret used to start generic trigger.",
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
					},
				},
			},
			"source_trigger_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "source trigger ID to clone from.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "name of the duplicated trigger.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trigger type.",
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
						"enum": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Options for SINGLE_SELECT property type.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"default": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Default option for SINGLE_SELECT property type.",
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
							Computed:    true,
							Description: "worker name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "worker type.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
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
				Description: "flag whether the trigger is disabled.",
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
							Description: "Needed only for git trigger type. Branch name of the repo. Branch field doesn't coexist with pattern field.",
						},
						"pattern": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Needed only for git trigger type. Git branch or tag pattern to listen to. Please refer to https://github.com/micromatch/micromatch for pattern syntax.",
						},
						"blind_connection": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Needed only for git trigger type. Branch name of the repo.",
						},
						"hook_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Webhook ID.",
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
							Description: "If true, the trigger starts when tekton pipeline service receive a repo's 'push' git webhook event.",
						},
						"pull_request_closed": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, the trigger starts when tekton pipeline service receive a repo pull reqeust's 'close' git webhook event.",
						},
						"pull_request": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "If true, the trigger starts when tekton pipeline service receive a repo pull reqeust's 'open' or 'update' git webhook event.",
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
				Description: "Needed only for timer trigger type. Cron expression for timer trigger. Maximum frequency is every 5 minutes.",
			},
			"timezone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Needed only for timer trigger type. Timezones for timer trigger.",
			},
			"secret": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Needed only for generic trigger type. Secret used to start generic trigger.",
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
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelineTriggerOptions := &cdtektonpipelinev2.CreateTektonPipelineTriggerOptions{}

	createTektonPipelineTriggerOptions.SetPipelineID(d.Get("pipeline_id").(string))
	if _, ok := d.GetOk("trigger"); ok {
		triggerModel, err := ResourceIBMTektonPipelineTriggerMapToTrigger(d.Get("trigger.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerOptions.SetTrigger(triggerModel)
	}

	triggerIntf, response, err := cdTektonPipelineClient.CreateTektonPipelineTriggerWithContext(context, createTektonPipelineTriggerOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTektonPipelineTriggerWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTektonPipelineTriggerWithContext failed %s\n%s", err, response))
	}

	trigger := triggerIntf.(*cdtektonpipelinev2.Trigger)
	d.SetId(fmt.Sprintf("%s/%s", *createTektonPipelineTriggerOptions.PipelineID, *trigger.ID))

	return ResourceIBMTektonPipelineTriggerRead(context, d, meta)
}

func ResourceIBMTektonPipelineTriggerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineTriggerOptions.SetPipelineID(parts[0])
	getTektonPipelineTriggerOptions.SetTriggerID(parts[1])

	triggerIntf, response, err := cdTektonPipelineClient.GetTektonPipelineTriggerWithContext(context, getTektonPipelineTriggerOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTektonPipelineTriggerWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelineTriggerWithContext failed %s\n%s", err, response))
	}

	trigger := triggerIntf.(*cdtektonpipelinev2.Trigger)
	if err = d.Set("pipeline_id", getTektonPipelineTriggerOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	// TODO: handle argument of type Trigger
	if err = d.Set("source_trigger_id", trigger.SourceTriggerID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting source_trigger_id: %s", err))
	}
	if err = d.Set("name", trigger.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("type", trigger.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
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
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerConcurrencyToMap(trigger.Concurrency)
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
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	updateTektonPipelineTriggerOptions := &cdtektonpipelinev2.UpdateTektonPipelineTriggerOptions{}

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
	if d.HasChange("trigger") {
		trigger := d.Get("trigger.0").(map[string]interface{})
		updateTektonPipelineTriggerOptions.SetTriggerPatch(trigger)
		hasChange = true
	}

	if hasChange {
		_, response, err := cdTektonPipelineClient.UpdateTektonPipelineTriggerWithContext(context, updateTektonPipelineTriggerOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateTektonPipelineTriggerWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateTektonPipelineTriggerWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMTektonPipelineTriggerRead(context, d, meta)
}

func ResourceIBMTektonPipelineTriggerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineTriggerOptions := &cdtektonpipelinev2.DeleteTektonPipelineTriggerOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineTriggerOptions.SetPipelineID(parts[0])
	deleteTektonPipelineTriggerOptions.SetTriggerID(parts[1])

	response, err := cdTektonPipelineClient.DeleteTektonPipelineTriggerWithContext(context, deleteTektonPipelineTriggerOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTektonPipelineTriggerWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTektonPipelineTriggerWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMTektonPipelineTriggerMapToTrigger(modelMap map[string]interface{}) (cdtektonpipelinev2.TriggerIntf, error) {
	model := &cdtektonpipelinev2.Trigger{}
	if modelMap["source_trigger_id"] != nil && modelMap["source_trigger_id"].(string) != "" {
		model.SourceTriggerID = core.StringPtr(modelMap["source_trigger_id"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["event_listener"] != nil && modelMap["event_listener"].(string) != "" {
		model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []cdtektonpipelinev2.TriggerPropertiesItem{}
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
	if modelMap["worker"] != nil && len(modelMap["worker"].([]interface{})) > 0 {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil && len(modelMap["concurrency"].([]interface{})) > 0 {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToConcurrency(modelMap["concurrency"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	if modelMap["disabled"] != nil {
		model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	}
	if modelMap["scm_source"] != nil && len(modelMap["scm_source"].([]interface{})) > 0 {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil && len(modelMap["events"].([]interface{})) > 0 {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil && modelMap["service_instance_id"].(string) != "" {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	if modelMap["cron"] != nil && modelMap["cron"].(string) != "" {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil && modelMap["timezone"].(string) != "" {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	if modelMap["secret"] != nil && len(modelMap["secret"].([]interface{})) > 0 {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerPropertiesItem(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerPropertiesItem, error) {
	model := &cdtektonpipelinev2.TriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["enum"] != nil {
		enum := []string{}
		for _, enumItem := range modelMap["enum"].([]interface{}) {
			enum = append(enum, enumItem.(string))
		}
		model.Enum = enum
	}
	if modelMap["default"] != nil && modelMap["default"].(string) != "" {
		model.Default = core.StringPtr(modelMap["default"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil && modelMap["path"].(string) != "" {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToWorker(modelMap map[string]interface{}) (*cdtektonpipelinev2.Worker, error) {
	model := &cdtektonpipelinev2.Worker{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToConcurrency(modelMap map[string]interface{}) (*cdtektonpipelinev2.Concurrency, error) {
	model := &cdtektonpipelinev2.Concurrency{}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerScmSource, error) {
	model := &cdtektonpipelinev2.TriggerScmSource{}
	if modelMap["url"] != nil && modelMap["url"].(string) != "" {
		model.URL = core.StringPtr(modelMap["url"].(string))
	}
	if modelMap["branch"] != nil && modelMap["branch"].(string) != "" {
		model.Branch = core.StringPtr(modelMap["branch"].(string))
	}
	if modelMap["pattern"] != nil && modelMap["pattern"].(string) != "" {
		model.Pattern = core.StringPtr(modelMap["pattern"].(string))
	}
	if modelMap["blind_connection"] != nil {
		model.BlindConnection = core.BoolPtr(modelMap["blind_connection"].(bool))
	}
	if modelMap["hook_id"] != nil && modelMap["hook_id"].(string) != "" {
		model.HookID = core.StringPtr(modelMap["hook_id"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToEvents(modelMap map[string]interface{}) (*cdtektonpipelinev2.Events, error) {
	model := &cdtektonpipelinev2.Events{}
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

func ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap map[string]interface{}) (*cdtektonpipelinev2.GenericSecret, error) {
	model := &cdtektonpipelinev2.GenericSecret{}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["source"] != nil && modelMap["source"].(string) != "" {
		model.Source = core.StringPtr(modelMap["source"].(string))
	}
	if modelMap["key_name"] != nil && modelMap["key_name"].(string) != "" {
		model.KeyName = core.StringPtr(modelMap["key_name"].(string))
	}
	if modelMap["algorithm"] != nil && modelMap["algorithm"].(string) != "" {
		model.Algorithm = core.StringPtr(modelMap["algorithm"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerDuplicateTrigger(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerDuplicateTrigger, error) {
	model := &cdtektonpipelinev2.TriggerDuplicateTrigger{}
	model.SourceTriggerID = core.StringPtr(modelMap["source_trigger_id"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerManualTrigger(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerManualTrigger, error) {
	model := &cdtektonpipelinev2.TriggerManualTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []cdtektonpipelinev2.TriggerManualTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerManualTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
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
	if modelMap["worker"] != nil && len(modelMap["worker"].([]interface{})) > 0 {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil && len(modelMap["concurrency"].([]interface{})) > 0 {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToConcurrency(modelMap["concurrency"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerManualTriggerPropertiesItem(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerManualTriggerPropertiesItem, error) {
	model := &cdtektonpipelinev2.TriggerManualTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["enum"] != nil {
		enum := []string{}
		for _, enumItem := range modelMap["enum"].([]interface{}) {
			enum = append(enum, enumItem.(string))
		}
		model.Enum = enum
	}
	if modelMap["default"] != nil && modelMap["default"].(string) != "" {
		model.Default = core.StringPtr(modelMap["default"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil && modelMap["path"].(string) != "" {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerScmTrigger(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerScmTrigger, error) {
	model := &cdtektonpipelinev2.TriggerScmTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []cdtektonpipelinev2.TriggerScmTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
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
	if modelMap["worker"] != nil && len(modelMap["worker"].([]interface{})) > 0 {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil && len(modelMap["concurrency"].([]interface{})) > 0 {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToConcurrency(modelMap["concurrency"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["scm_source"] != nil && len(modelMap["scm_source"].([]interface{})) > 0 {
		ScmSourceModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil && len(modelMap["events"].([]interface{})) > 0 {
		EventsModel, err := ResourceIBMTektonPipelineTriggerMapToEvents(modelMap["events"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["service_instance_id"] != nil && modelMap["service_instance_id"].(string) != "" {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerScmTriggerPropertiesItem(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerScmTriggerPropertiesItem, error) {
	model := &cdtektonpipelinev2.TriggerScmTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["enum"] != nil {
		enum := []string{}
		for _, enumItem := range modelMap["enum"].([]interface{}) {
			enum = append(enum, enumItem.(string))
		}
		model.Enum = enum
	}
	if modelMap["default"] != nil && modelMap["default"].(string) != "" {
		model.Default = core.StringPtr(modelMap["default"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil && modelMap["path"].(string) != "" {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerTimerTrigger(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerTimerTrigger, error) {
	model := &cdtektonpipelinev2.TriggerTimerTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []cdtektonpipelinev2.TriggerTimerTriggerPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerMapToTriggerTimerTriggerPropertiesItem(propertiesItem.(map[string]interface{}))
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
	if modelMap["worker"] != nil && len(modelMap["worker"].([]interface{})) > 0 {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil && len(modelMap["concurrency"].([]interface{})) > 0 {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToConcurrency(modelMap["concurrency"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["cron"] != nil && modelMap["cron"].(string) != "" {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil && modelMap["timezone"].(string) != "" {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerTimerTriggerPropertiesItem(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerTimerTriggerPropertiesItem, error) {
	model := &cdtektonpipelinev2.TriggerTimerTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["enum"] != nil {
		enum := []string{}
		for _, enumItem := range modelMap["enum"].([]interface{}) {
			enum = append(enum, enumItem.(string))
		}
		model.Enum = enum
	}
	if modelMap["default"] != nil && modelMap["default"].(string) != "" {
		model.Default = core.StringPtr(modelMap["default"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil && modelMap["path"].(string) != "" {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerGenericTrigger(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerGenericTrigger, error) {
	model := &cdtektonpipelinev2.TriggerGenericTrigger{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["properties"] != nil {
		properties := []cdtektonpipelinev2.TriggerGenericTriggerPropertiesItem{}
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
	if modelMap["worker"] != nil && len(modelMap["worker"].([]interface{})) > 0 {
		WorkerModel, err := ResourceIBMTektonPipelineTriggerMapToWorker(modelMap["worker"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["concurrency"] != nil && len(modelMap["concurrency"].([]interface{})) > 0 {
		ConcurrencyModel, err := ResourceIBMTektonPipelineTriggerMapToConcurrency(modelMap["concurrency"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Concurrency = ConcurrencyModel
	}
	model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	if modelMap["secret"] != nil && len(modelMap["secret"].([]interface{})) > 0 {
		SecretModel, err := ResourceIBMTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerMapToTriggerGenericTriggerPropertiesItem(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerGenericTriggerPropertiesItem, error) {
	model := &cdtektonpipelinev2.TriggerGenericTriggerPropertiesItem{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil && modelMap["value"].(string) != "" {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["enum"] != nil {
		enum := []string{}
		for _, enumItem := range modelMap["enum"].([]interface{}) {
			enum = append(enum, enumItem.(string))
		}
		model.Enum = enum
	}
	if modelMap["default"] != nil && modelMap["default"].(string) != "" {
		model.Default = core.StringPtr(modelMap["default"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["path"] != nil && modelMap["path"].(string) != "" {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerTriggerToMap(model cdtektonpipelinev2.TriggerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*cdtektonpipelinev2.TriggerDuplicateTrigger); ok {
		return ResourceIBMTektonPipelineTriggerTriggerDuplicateTriggerToMap(model.(*cdtektonpipelinev2.TriggerDuplicateTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerManualTrigger); ok {
		return ResourceIBMTektonPipelineTriggerTriggerManualTriggerToMap(model.(*cdtektonpipelinev2.TriggerManualTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerScmTrigger); ok {
		return ResourceIBMTektonPipelineTriggerTriggerScmTriggerToMap(model.(*cdtektonpipelinev2.TriggerScmTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerTimerTrigger); ok {
		return ResourceIBMTektonPipelineTriggerTriggerTimerTriggerToMap(model.(*cdtektonpipelinev2.TriggerTimerTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerGenericTrigger); ok {
		return ResourceIBMTektonPipelineTriggerTriggerGenericTriggerToMap(model.(*cdtektonpipelinev2.TriggerGenericTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.Trigger); ok {
		modelMap := make(map[string]interface{})
		model := model.(*cdtektonpipelinev2.Trigger)
		if model.SourceTriggerID != nil {
			modelMap["source_trigger_id"] = model.SourceTriggerID
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.Type != nil {
			modelMap["type"] = model.Type
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
			concurrencyMap, err := ResourceIBMTektonPipelineTriggerConcurrencyToMap(model.Concurrency)
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
		return nil, fmt.Errorf("Unrecognized cdtektonpipelinev2.TriggerIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelineTriggerTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Default != nil {
		modelMap["default"] = model.Default
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

func ResourceIBMTektonPipelineTriggerWorkerToMap(model *cdtektonpipelinev2.Worker) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelineTriggerConcurrencyToMap(model *cdtektonpipelinev2.Concurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerScmSourceToMap(model *cdtektonpipelinev2.TriggerScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.URL != nil {
		modelMap["url"] = model.URL
	}
	if model.Branch != nil {
		modelMap["branch"] = model.Branch
	}
	if model.Pattern != nil {
		modelMap["pattern"] = model.Pattern
	}
	if model.BlindConnection != nil {
		modelMap["blind_connection"] = model.BlindConnection
	}
	if model.HookID != nil {
		modelMap["hook_id"] = model.HookID
	}
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerEventsToMap(model *cdtektonpipelinev2.Events) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelineTriggerGenericSecretToMap(model *cdtektonpipelinev2.GenericSecret) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelineTriggerTriggerDuplicateTriggerToMap(model *cdtektonpipelinev2.TriggerDuplicateTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_trigger_id"] = model.SourceTriggerID
	modelMap["name"] = model.Name
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerManualTriggerToMap(model *cdtektonpipelinev2.TriggerManualTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerTriggerManualTriggerPropertiesItemToMap(&propertiesItem)
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
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerTriggerManualTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerManualTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Default != nil {
		modelMap["default"] = model.Default
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

func ResourceIBMTektonPipelineTriggerTriggerScmTriggerToMap(model *cdtektonpipelinev2.TriggerScmTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerTriggerScmTriggerPropertiesItemToMap(&propertiesItem)
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
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerConcurrencyToMap(model.Concurrency)
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

func ResourceIBMTektonPipelineTriggerTriggerScmTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerScmTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Default != nil {
		modelMap["default"] = model.Default
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

func ResourceIBMTektonPipelineTriggerTriggerTimerTriggerToMap(model *cdtektonpipelinev2.TriggerTimerTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := ResourceIBMTektonPipelineTriggerTriggerTimerTriggerPropertiesItemToMap(&propertiesItem)
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
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerConcurrencyToMap(model.Concurrency)
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

func ResourceIBMTektonPipelineTriggerTriggerTimerTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerTimerTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Default != nil {
		modelMap["default"] = model.Default
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

func ResourceIBMTektonPipelineTriggerTriggerGenericTriggerToMap(model *cdtektonpipelinev2.TriggerGenericTrigger) (map[string]interface{}, error) {
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
		concurrencyMap, err := ResourceIBMTektonPipelineTriggerConcurrencyToMap(model.Concurrency)
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

func ResourceIBMTektonPipelineTriggerTriggerGenericTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerGenericTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Default != nil {
		modelMap["default"] = model.Default
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
