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
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIBMCdTektonPipelineTrigger() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIBMCdTektonPipelineTriggerCreate,
		ReadContext:     resourceIBMCdTektonPipelineTriggerRead,
		UpdateContext:   resourceIBMCdTektonPipelineTriggerUpdate,
		DeleteContext:   resourceIBMCdTektonPipelineTriggerDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "pipeline_id"),
				Description: "The Tekton pipeline ID.",
			},
			"trigger": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "Tekton pipeline trigger.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_trigger_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the trigger to duplicate. Only needed when duplicating a trigger.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Trigger name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Trigger type.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "API URL for interacting with the trigger.",
						},
						"event_listener": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "ID.",
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
							Description: "Worker used to run the trigger. If not specified the trigger will use the default pipeline worker.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Name of the worker. Computed based on the worker ID.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Type of the worker. Computed based on the worker ID.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "ID of the worker.",
									},
								},
							},
						},
						"max_concurrent_runs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Defines the maximum number of concurrent runs for this trigger. Omit this property to disable the concurrency limit.",
						},
						"disabled": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Flag whether the trigger is disabled. If omitted the trigger is enabled by default.",
						},
						"scm_source": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "SCM source repository for a Git trigger. Only needed for Git triggers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "URL of the repository to which the trigger is listening.",
									},
									"branch": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of a branch from the repo. One of branch or tag must be specified, but only one or the other.",
									},
									"pattern": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Git branch or tag pattern to listen to. Please refer to https://github.com/micromatch/micromatch for pattern syntax.",
									},
									"blind_connection": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Set this boolean to true if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide. False by default.",
									},
									"hook_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the webhook from the repo. Computed upon creation of the trigger.",
									},
									"service_instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the repository service instance.",
									},
								},
							},
						},
						"events": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Only needed for Git triggers. Events object defines the events to which this Git trigger listens.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"push": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger listens for 'push' Git webhook events.",
									},
									"pull_request_closed": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger listens for 'close pull request' Git webhook events.",
									},
									"pull_request": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "If true, the trigger listens for 'open pull request' or 'update pull request' Git webhook events.",
									},
								},
							},
						},
						"cron": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Only needed for timer triggers. Cron expression for timer trigger. Maximum frequency is every 5 minutes.",
						},
						"timezone": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Only needed for timer triggers. Timezone for timer trigger.",
						},
						"secret": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Only needed for generic webhook trigger type. Secret used to start generic webhook trigger.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Secret type.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										DiffSuppressFunc: flex.SuppressGenericWebhookRawSecret,
										Description: "Secret value, not needed if secret type is `internal_validation`.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret location, not needed if secret type is `internal_validation`.",
									},
									"key_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Secret name, not needed if type is `internal_validation`.",
									},
									"algorithm": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Algorithm used for `digest_matches` secret type. Only needed for `digest_matches` secret type.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func ResourceIBMCdTektonPipelineTriggerValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_cd_tekton_pipeline_trigger", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMCdTektonPipelineTriggerCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelineTriggerOptions := &cdtektonpipelinev2.CreateTektonPipelineTriggerOptions{}

	createTektonPipelineTriggerOptions.SetPipelineID(d.Get("pipeline_id").(string))
	if _, ok := d.GetOk("trigger"); ok {
		triggerModel, err := resourceIBMCdTektonPipelineTriggerMapToTrigger(d.Get("trigger.0").(map[string]interface{}))
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

	return resourceIBMCdTektonPipelineTriggerRead(context, d, meta)
}

func resourceIBMCdTektonPipelineTriggerRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	if err = d.Set("pipeline_id", getTektonPipelineTriggerOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}

	trigger, err := resourceIBMCdTektonPipelineTriggerTriggerToMap(triggerIntf)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("trigger", []map[string]interface{}{trigger}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting trigger: %s", err))
	}

	return nil
}

func resourceIBMCdTektonPipelineTriggerUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	patchVals := &cdtektonpipelinev2.TriggerPatch{}
	if d.HasChange("pipeline_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation." +
				" The resource must be re-created to update this property.", "pipeline_id"))
	}
	if d.HasChange("trigger.0.name") {
		patchVals.Name = core.StringPtr(d.Get("trigger.0.name").(string))
		hasChange = true
	}
	if d.HasChange("trigger.0.events") {
		events, err := resourceIBMCdTektonPipelineTriggerMapToEvents(d.Get("trigger.0.events").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Events = events
		hasChange = true
	}
	if d.HasChange("trigger.0.event_listener") {
		patchVals.EventListener = core.StringPtr(d.Get("trigger.0.event_listener").(string))
		hasChange = true
	}
	if d.HasChange("trigger.0.tags") {
		tags := []string{}
		for _, tagsItem := range d.Get("trigger.0.tags").([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		patchVals.Tags = tags
		hasChange = true
	}
	if d.HasChange("trigger.0.worker") {
		worker, err := resourceIBMCdTektonPipelineTriggerMapToWorker(d.Get("trigger.0.worker").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Worker = worker
		hasChange = true
	}
	if d.HasChange("trigger.0.max_concurrent_runs") {
		patchVals.MaxConcurrentRuns = core.Int64Ptr(int64(d.Get("trigger.0.max_concurrent_runs").(int)))
		hasChange = true
	}
	if d.HasChange("trigger.0.secret") {
		secret, err := resourceIBMCdTektonPipelineTriggerMapToGenericSecret(d.Get("trigger.0.secret").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Secret = secret
		hasChange = true
	}
	if d.HasChange("trigger.0.scm_source") {
		scmSource, err := resourceIBMCdTektonPipelineTriggerMapToTriggerScmSource(d.Get("trigger.0.scm_source").([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.ScmSource = scmSource
		hasChange = true
	}
	if d.HasChange("trigger.0.cron") {
		patchVals.Cron = core.StringPtr(d.Get("trigger.0.cron").(string))
		hasChange = true
	}
	if d.HasChange("trigger.0.timezone") {
		patchVals.Timezone = core.StringPtr(d.Get("trigger.0.timezone").(string))
		hasChange = true
	}
	if d.HasChange("trigger.0.disabled") {
		patchVals.Disabled = core.BoolPtr(d.Get("trigger.0.disabled").(bool))
		hasChange = true
	}

	if hasChange {
		updateTektonPipelineTriggerOptions.TriggerPatch, _ = patchVals.AsPatch()
		_, response, err := cdTektonPipelineClient.UpdateTektonPipelineTriggerWithContext(context, updateTektonPipelineTriggerOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateTektonPipelineTriggerWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateTektonPipelineTriggerWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIBMCdTektonPipelineTriggerRead(context, d, meta)
}

func resourceIBMCdTektonPipelineTriggerDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIBMCdTektonPipelineTriggerMapToTrigger(modelMap map[string]interface{}) (cdtektonpipelinev2.TriggerIntf, error) {
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
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	if modelMap["event_listener"] != nil && modelMap["event_listener"].(string) != "" {
		model.EventListener = core.StringPtr(modelMap["event_listener"].(string))
	}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["tags"] != nil {
		tags := []string{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		model.Tags = tags
	}
	if modelMap["worker"] != nil && len(modelMap["worker"].([]interface{})) > 0 {
		WorkerModel, err := resourceIBMCdTektonPipelineTriggerMapToWorker(modelMap["worker"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Worker = WorkerModel
	}
	if modelMap["max_concurrent_runs"] != nil {
		model.MaxConcurrentRuns = core.Int64Ptr(int64(modelMap["max_concurrent_runs"].(int)))
	}
	if modelMap["disabled"] != nil {
		model.Disabled = core.BoolPtr(modelMap["disabled"].(bool))
	}
	if modelMap["scm_source"] != nil && len(modelMap["scm_source"].([]interface{})) > 0 {
		ScmSourceModel, err := resourceIBMCdTektonPipelineTriggerMapToTriggerScmSource(modelMap["scm_source"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ScmSource = ScmSourceModel
	}
	if modelMap["events"] != nil && len(modelMap["events"].([]interface{})) > 0 {
		EventsModel, err := resourceIBMCdTektonPipelineTriggerMapToEvents(modelMap["events"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Events = EventsModel
	}
	if modelMap["cron"] != nil && modelMap["cron"].(string) != "" {
		model.Cron = core.StringPtr(modelMap["cron"].(string))
	}
	if modelMap["timezone"] != nil && modelMap["timezone"].(string) != "" {
		model.Timezone = core.StringPtr(modelMap["timezone"].(string))
	}
	if modelMap["secret"] != nil && len(modelMap["secret"].([]interface{})) > 0 {
		SecretModel, err := resourceIBMCdTektonPipelineTriggerMapToGenericSecret(modelMap["secret"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Secret = SecretModel
	}
	return model, nil
}

func resourceIBMCdTektonPipelineTriggerMapToWorker(modelMap map[string]interface{}) (*cdtektonpipelinev2.Worker, error) {
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

func resourceIBMCdTektonPipelineTriggerMapToTriggerScmSource(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerScmSource, error) {
	model := &cdtektonpipelinev2.TriggerScmSource{}
	model.URL = core.StringPtr(modelMap["url"].(string))
	if modelMap["branch"] != nil && modelMap["branch"].(string) != "" {
		model.Branch = core.StringPtr(modelMap["branch"].(string))
	}
	if modelMap["pattern"] != nil && modelMap["pattern"].(string) != "" {
		model.Pattern = core.StringPtr(modelMap["pattern"].(string))
	}
	if modelMap["service_instance_id"] != nil && modelMap["service_instance_id"].(string) != "" {
		model.ServiceInstanceID = core.StringPtr(modelMap["service_instance_id"].(string))
	}
	return model, nil
}

func resourceIBMCdTektonPipelineTriggerMapToEvents(modelMap map[string]interface{}) (*cdtektonpipelinev2.Events, error) {
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

func resourceIBMCdTektonPipelineTriggerMapToGenericSecret(modelMap map[string]interface{}) (*cdtektonpipelinev2.GenericSecret, error) {
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

func resourceIBMCdTektonPipelineTriggerTriggerToMap(model cdtektonpipelinev2.TriggerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*cdtektonpipelinev2.TriggerDuplicateTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerTriggerDuplicateTriggerToMap(model.(*cdtektonpipelinev2.TriggerDuplicateTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerManualTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerTriggerManualTriggerToMap(model.(*cdtektonpipelinev2.TriggerManualTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerScmTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerTriggerScmTriggerToMap(model.(*cdtektonpipelinev2.TriggerScmTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerTimerTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerTriggerTimerTriggerToMap(model.(*cdtektonpipelinev2.TriggerTimerTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerGenericTrigger); ok {
		return resourceIBMCdTektonPipelineTriggerTriggerGenericTriggerToMap(model.(*cdtektonpipelinev2.TriggerGenericTrigger))
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
		if model.Href != nil {
			modelMap["href"] = model.Href
		}
		if model.EventListener != nil {
			modelMap["event_listener"] = model.EventListener
		}
		if model.ID != nil {
			modelMap["id"] = model.ID
		}
		if model.Tags != nil {
			modelMap["tags"] = model.Tags
		}
		if model.Worker != nil {
			workerMap, err := resourceIBMCdTektonPipelineTriggerWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.MaxConcurrentRuns != nil {
			modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
		}
		if model.Disabled != nil {
			modelMap["disabled"] = model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := resourceIBMCdTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := resourceIBMCdTektonPipelineTriggerEventsToMap(model.Events)
			if err != nil {
				return modelMap, err
			}
			modelMap["events"] = []map[string]interface{}{eventsMap}
		}
		if model.Cron != nil {
			modelMap["cron"] = model.Cron
		}
		if model.Timezone != nil {
			modelMap["timezone"] = model.Timezone
		}
		if model.Secret != nil {
			secretMap, err := resourceIBMCdTektonPipelineTriggerGenericSecretToMap(model.Secret)
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

func resourceIBMCdTektonPipelineTriggerWorkerToMap(model *cdtektonpipelinev2.Worker) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerTriggerScmSourceToMap(model *cdtektonpipelinev2.TriggerScmSource) (map[string]interface{}, error) {
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
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = model.ServiceInstanceID
	}
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerEventsToMap(model *cdtektonpipelinev2.Events) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerGenericSecretToMap(model *cdtektonpipelinev2.GenericSecret) (map[string]interface{}, error) {
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

func resourceIBMCdTektonPipelineTriggerTriggerDuplicateTriggerToMap(model *cdtektonpipelinev2.TriggerDuplicateTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["source_trigger_id"] = model.SourceTriggerID
	modelMap["name"] = model.Name
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerTriggerManualTriggerToMap(model *cdtektonpipelinev2.TriggerManualTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := resourceIBMCdTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	modelMap["disabled"] = model.Disabled
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerTriggerScmTriggerToMap(model *cdtektonpipelinev2.TriggerScmTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := resourceIBMCdTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	modelMap["disabled"] = model.Disabled
	if model.ScmSource != nil {
		scmSourceMap, err := resourceIBMCdTektonPipelineTriggerTriggerScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.Events != nil {
		eventsMap, err := resourceIBMCdTektonPipelineTriggerEventsToMap(model.Events)
		if err != nil {
			return modelMap, err
		}
		modelMap["events"] = []map[string]interface{}{eventsMap}
	}
	return modelMap, nil
}

func resourceIBMCdTektonPipelineTriggerTriggerTimerTriggerToMap(model *cdtektonpipelinev2.TriggerTimerTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := resourceIBMCdTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
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

func resourceIBMCdTektonPipelineTriggerTriggerGenericTriggerToMap(model *cdtektonpipelinev2.TriggerGenericTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	modelMap["event_listener"] = model.EventListener
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Tags != nil {
		modelMap["tags"] = model.Tags
	}
	if model.Worker != nil {
		workerMap, err := resourceIBMCdTektonPipelineTriggerWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = flex.IntValue(model.MaxConcurrentRuns)
	}
	modelMap["disabled"] = model.Disabled
	if model.Secret != nil {
		secretMap, err := resourceIBMCdTektonPipelineTriggerGenericSecretToMap(model.Secret)
		if err != nil {
			return modelMap, err
		}
		modelMap["secret"] = []map[string]interface{}{secretMap}
	}
	return modelMap, nil
}
