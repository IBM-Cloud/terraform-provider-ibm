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
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIBMCdTektonPipeline() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCdTektonPipelineCreate,
		ReadContext:   resourceIBMCdTektonPipelineRead,
		UpdateContext: resourceIBMCdTektonPipelineUpdate,
		DeleteContext: resourceIBMCdTektonPipelineDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"worker": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Worker object with worker ID only.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"pipeline_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "String.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Pipeline status.",
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID.",
			},
			"toolchain": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Toolchain object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "UUID.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN for the toolchain that containing the tekton pipeline.",
						},
					},
				},
			},
			"definitions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Definition list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scm_source": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Scm source for tekton pipeline defintion.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "General href URL.",
									},
									"branch": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A branch of the repo, branch field doesn't coexist with tag field.",
									},
									"tag": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A tag of the repo.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The path to the definitions yaml files.",
									},
								},
							},
						},
						"service_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "UUID.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "UUID.",
						},
					},
				},
			},
			"properties": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tekton pipeline's environment properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Property name.",
						},
						"value": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: flex.SuppressPipelinePropertyRawSecret,
							Description:      "String format property value.",
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
					},
				},
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Standard RFC 3339 Date Time String.",
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Standard RFC 3339 Date Time String.",
			},
			"pipeline_definition": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tekton pipeline definition document detail object. If this property is absent, the pipeline has no definitions added.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The state of pipeline definition status.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "UUID.",
						},
					},
				},
			},
			"triggers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tekton pipeline triggers list.",
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
							Computed:    true,
							Description: "Id.",
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
										ForceNew:    true,
										Description: "Property name.",
									},
									"value": &schema.Schema{
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: flex.SuppressTriggerPropertyRawSecret,
										Description:      "String format property value.",
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
										Description: "Id.",
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
										Required:    true,
										ForceNew:    true,
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
										Type:             schema.TypeString,
										Optional:         true,
										DiffSuppressFunc: flex.SuppressGenericWebhookRawSecret,
										Description:      "Secret value, not needed if secret type is \"internalValidation\".",
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
			"html_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Dashboard URL of this pipeline.",
			},
			"build_number": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The latest pipeline run build number. If this property is absent, the pipeline hasn't had any pipelineRuns.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag whether this pipeline enabled.",
			},
		},
	}
}

func resourceIBMCdTektonPipelineCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelineOptions := &cdtektonpipelinev2.CreateTektonPipelineOptions{}

	if _, ok := d.GetOk("worker"); ok {
		workerModel, err := resourceIBMCdTektonPipelineMapToWorkerWithID(d.Get("worker.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineOptions.SetWorker(workerModel)
	}

	if _, ok := d.GetOk("pipeline_id"); ok {
		createTektonPipelineOptions.SetID(d.Get("pipeline_id").(string))
	}
	tektonPipeline, response, err := cdTektonPipelineClient.CreateTektonPipelineWithContext(context, createTektonPipelineOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTektonPipelineWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTektonPipelineWithContext failed %s\n%s", err, response))
	}

	d.SetId(*tektonPipeline.ID)

	return resourceIBMCdTektonPipelineRead(context, d, meta)
}

func resourceIBMCdTektonPipelineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{}

	getTektonPipelineOptions.SetID(d.Id())

	tektonPipeline, response, err := cdTektonPipelineClient.GetTektonPipelineWithContext(context, getTektonPipelineOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTektonPipelineWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelineWithContext failed %s\n%s", err, response))
	}

	if tektonPipeline.Worker != nil {
		workerMap, err := resourceIBMCdTektonPipelineWorkerWithIDToMap(tektonPipeline.Worker)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("worker", []map[string]interface{}{workerMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting worker: %s", err))
		}
	}
	if err = d.Set("pipeline_id", tektonPipeline.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	if err = d.Set("name", tektonPipeline.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("status", tektonPipeline.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}
	if err = d.Set("resource_group_id", tektonPipeline.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	toolchainMap, err := resourceIBMCdTektonPipelineToolchainToMap(tektonPipeline.Toolchain)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("toolchain", []map[string]interface{}{toolchainMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain: %s", err))
	}
	definitions := []map[string]interface{}{}
	for _, definitionsItem := range tektonPipeline.Definitions {
		definitionsItemMap, err := resourceIBMCdTektonPipelineDefinitionToMap(&definitionsItem)
		if err != nil {
			return diag.FromErr(err)
		}
		definitions = append(definitions, definitionsItemMap)
	}
	if err = d.Set("definitions", definitions); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definitions: %s", err))
	}
	properties := []map[string]interface{}{}
	for _, propertiesItem := range tektonPipeline.Properties {
		propertiesItemMap, err := resourceIBMCdTektonPipelinePropertyToMap(&propertiesItem)
		if err != nil {
			return diag.FromErr(err)
		}
		properties = append(properties, propertiesItemMap)
	}
	if err = d.Set("properties", properties); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting properties: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(tektonPipeline.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("created", flex.DateTimeToString(tektonPipeline.Created)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if tektonPipeline.PipelineDefinition != nil {
		pipelineDefinitionMap, err := resourceIBMCdTektonPipelineTektonPipelinePipelineDefinitionToMap(tektonPipeline.PipelineDefinition)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("pipeline_definition", []map[string]interface{}{pipelineDefinitionMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting pipeline_definition: %s", err))
		}
	}
	triggers := []map[string]interface{}{}
	for _, triggersItem := range tektonPipeline.Triggers {
		triggersItemMap, err := resourceIBMCdTektonPipelineTriggerToMap(triggersItem)
		if err != nil {
			return diag.FromErr(err)
		}
		triggers = append(triggers, triggersItemMap)
	}
	if err = d.Set("triggers", triggers); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting triggers: %s", err))
	}
	if err = d.Set("html_url", tektonPipeline.HTMLURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting html_url: %s", err))
	}
	if err = d.Set("build_number", flex.IntValue(tektonPipeline.BuildNumber)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting build_number: %s", err))
	}
	if err = d.Set("enabled", tektonPipeline.Enabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting enabled: %s", err))
	}

	return nil
}

func resourceIBMCdTektonPipelineUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	updateTektonPipelineOptions := &cdtektonpipelinev2.UpdateTektonPipelineOptions{}

	updateTektonPipelineOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("worker") {
		worker, err := resourceIBMCdTektonPipelineMapToWorkerWithID(d.Get("worker.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateTektonPipelineOptions.SetWorker(worker)
		hasChange = true
	}

	if hasChange {
		_, response, err := cdTektonPipelineClient.UpdateTektonPipelineWithContext(context, updateTektonPipelineOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateTektonPipelineWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateTektonPipelineWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIBMCdTektonPipelineRead(context, d, meta)
}

func resourceIBMCdTektonPipelineDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineOptions := &cdtektonpipelinev2.DeleteTektonPipelineOptions{}

	deleteTektonPipelineOptions.SetID(d.Id())

	response, err := cdTektonPipelineClient.DeleteTektonPipelineWithContext(context, deleteTektonPipelineOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTektonPipelineWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTektonPipelineWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIBMCdTektonPipelineMapToWorkerWithID(modelMap map[string]interface{}) (*cdtektonpipelinev2.WorkerWithID, error) {
	model := &cdtektonpipelinev2.WorkerWithID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func resourceIBMCdTektonPipelineWorkerWithIDToMap(model *cdtektonpipelinev2.Worker) (map[string]interface{}, error) {
	// TODO we alter cdtektonpipelinev2.WorkerWithID to cdtektonpipelinev2.Worker in func params. Determine why and if we can fix it
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	return modelMap, nil
}

func resourceIBMCdTektonPipelineToolchainToMap(model *cdtektonpipelinev2.Toolchain) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	modelMap["crn"] = model.CRN
	return modelMap, nil
}

func resourceIBMCdTektonPipelineDefinitionToMap(model *cdtektonpipelinev2.Definition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	scmSourceMap, err := resourceIBMCdTektonPipelineDefinitionScmSourceToMap(model.ScmSource)
	if err != nil {
		return modelMap, err
	}
	modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	modelMap["service_instance_id"] = model.ServiceInstanceID
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	return modelMap, nil
}

func resourceIBMCdTektonPipelineDefinitionScmSourceToMap(model *cdtektonpipelinev2.DefinitionScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["url"] = model.URL
	if model.Branch != nil {
		modelMap["branch"] = model.Branch
	}
	if model.Tag != nil {
		modelMap["tag"] = model.Tag
	}
	modelMap["path"] = model.Path
	return modelMap, nil
}

func resourceIBMCdTektonPipelinePropertyToMap(model *cdtektonpipelinev2.Property) (map[string]interface{}, error) {
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
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTektonPipelinePipelineDefinitionToMap(model *cdtektonpipelinev2.TektonPipelinePipelineDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerToMap(model cdtektonpipelinev2.TriggerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*cdtektonpipelinev2.TriggerDuplicateTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerDuplicateTriggerToMap(model.(*cdtektonpipelinev2.TriggerDuplicateTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerManualTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerManualTriggerToMap(model.(*cdtektonpipelinev2.TriggerManualTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerScmTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerScmTriggerToMap(model.(*cdtektonpipelinev2.TriggerScmTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerTimerTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerTimerTriggerToMap(model.(*cdtektonpipelinev2.TriggerTimerTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerGenericTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerGenericTriggerToMap(model.(*cdtektonpipelinev2.TriggerGenericTrigger))
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
				propertiesItemMap, err := resourceIBMCdTektonPipelineTriggerPropertiesItemToMap(&propertiesItem)
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
			workerMap, err := resourceIBMCdTektonPipelineWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.Concurrency != nil {
			concurrencyMap, err := resourceIBMCdTektonPipelineConcurrencyToMap(model.Concurrency)
			if err != nil {
				return modelMap, err
			}
			modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
		}
		if model.Disabled != nil {
			modelMap["disabled"] = model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := resourceIBMCdTektonPipelineTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := resourceIBMCdTektonPipelineEventsToMap(model.Events)
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
			secretMap, err := resourceIBMCdTektonPipelineGenericSecretToMap(model.Secret)
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

func resourceIBMCdTektonPipelineTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerPropertiesItem) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineWorkerToMap(model *cdtektonpipelinev2.Worker) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineConcurrencyToMap(model *cdtektonpipelinev2.Concurrency) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerScmSourceToMap(model *cdtektonpipelinev2.TriggerScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["url"] = model.URL
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

func resourceIBMCdTektonPipelineEventsToMap(model *cdtektonpipelinev2.Events) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineGenericSecretToMap(model *cdtektonpipelinev2.GenericSecret) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerDuplicateTriggerToMap(model *cdtektonpipelinev2.TriggerDuplicateTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_trigger_id"] = model.SourceTriggerID
	modelMap["name"] = model.Name
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerManualTriggerToMap(model *cdtektonpipelinev2.TriggerManualTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := resourceIBMCdTektonPipelineTriggerManualTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := resourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := resourceIBMCdTektonPipelineConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerManualTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerManualTriggerPropertiesItem) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerScmTriggerToMap(model *cdtektonpipelinev2.TriggerScmTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := resourceIBMCdTektonPipelineTriggerScmTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := resourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := resourceIBMCdTektonPipelineConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.ScmSource != nil {
		scmSourceMap, err := resourceIBMCdTektonPipelineTriggerScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.Events != nil {
		eventsMap, err := resourceIBMCdTektonPipelineEventsToMap(model.Events)
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

func resourceIBMCdTektonPipelineTriggerScmTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerScmTriggerPropertiesItem) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerTimerTriggerToMap(model *cdtektonpipelinev2.TriggerTimerTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := resourceIBMCdTektonPipelineTriggerTimerTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := resourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := resourceIBMCdTektonPipelineConcurrencyToMap(model.Concurrency)
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

func resourceIBMCdTektonPipelineTriggerTimerTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerTimerTriggerPropertiesItem) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerGenericTriggerToMap(model *cdtektonpipelinev2.TriggerGenericTrigger) (map[string]interface{}, error) {
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
			propertiesItemMap, err := resourceIBMCdTektonPipelineTriggerGenericTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := resourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.Concurrency != nil {
		concurrencyMap, err := resourceIBMCdTektonPipelineConcurrencyToMap(model.Concurrency)
		if err != nil {
			return modelMap, err
		}
		modelMap["concurrency"] = []map[string]interface{}{concurrencyMap}
	}
	modelMap["disabled"] = model.Disabled
	if model.Secret != nil {
		secretMap, err := resourceIBMCdTektonPipelineGenericSecretToMap(model.Secret)
		if err != nil {
			return modelMap, err
		}
		modelMap["secret"] = []map[string]interface{}{secretMap}
	}
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerGenericTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerGenericTriggerPropertiesItem) (map[string]interface{}, error) {
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
