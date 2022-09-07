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
		CreateContext: resourceIBMCdTektonPipelineTriggerCreate,
		ReadContext:   resourceIBMCdTektonPipelineTriggerRead,
		UpdateContext: resourceIBMCdTektonPipelineTriggerUpdate,
		DeleteContext: resourceIBMCdTektonPipelineTriggerDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "pipeline_id"),
				Description:  "The Tekton pipeline ID.",
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "type"),
				Description:  "Trigger type.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "name"),
				Description:  "Trigger name.",
			},
			"event_listener": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "event_listener"),
				Description:  "Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.",
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
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: flex.SuppressGenericWebhookRawSecret,
							Description:      "Secret value, not needed if secret type is `internal_validation`.",
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
			"cron": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "cron"),
				Description:  "Only needed for timer triggers. Cron expression for timer trigger.",
			},
			"timezone": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_cd_tekton_pipeline_trigger", "timezone"),
				Description:  "Only needed for timer triggers. Timezone for timer trigger.",
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
							Computed:    true,
							Description: "True if the repository server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.",
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
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API URL for interacting with the trigger.",
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
							ForceNew:    true,
							Description: "Property name.",
						},
						"value": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: flex.SuppressTriggerPropertyRawSecret,
							Description:      "Property value. Can be empty and should be omitted for `single_select` property type.",
						},
						"enum": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Options for `single_select` property type. Only needed for `single_select` property type.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Property type.",
						},
						"path": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A dot notation path for `integration` type properties to select a value from the tool integration. If left blank the full tool integration data will be used.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "API URL for interacting with the trigger property.",
						},
					},
				},
			},
			"webhook_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Webhook URL that can be used to trigger pipeline runs.",
			},
			"trigger_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID.",
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
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "generic, manual, scm, timer",
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9][-0-9a-zA-Z_. ]{1,235}[a-zA-Z0-9]$`,
			MinValueLength:             1,
			MaxValueLength:             253,
		},
		validate.ValidateSchema{
			Identifier:                 "event_listener",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[-0-9a-zA-Z_.]{1,235}$`,
			MinValueLength:             1,
			MaxValueLength:             253,
		},
		validate.ValidateSchema{
			Identifier:                 "cron",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\*\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])) (\*|([0-9]|1[0-9]|2[0-3])|\*\/([0-9]|1[0-9]|2[0-3])) (\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\*\/([1-9]|1[0-9]|2[0-9]|3[0-1])) (\*|([1-9]|1[0-2])|\*\/([1-9]|1[0-2])) (\*|([0-6])|\*\/([0-6]))$`,
			MinValueLength:             5,
			MaxValueLength:             253,
		},
		validate.ValidateSchema{
			Identifier:                 "timezone",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[-0-9a-zA-Z_., \/]{1,234}$`,
			MinValueLength:             1,
			MaxValueLength:             253,
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
	if _, ok := d.GetOk("type"); ok {
		createTektonPipelineTriggerOptions.SetType(d.Get("type").(string))
	}
	if _, ok := d.GetOk("name"); ok {
		createTektonPipelineTriggerOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("event_listener"); ok {
		createTektonPipelineTriggerOptions.SetEventListener(d.Get("event_listener").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		tags := []string{}
		for _, tagsItem := range d.Get("tags").([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		createTektonPipelineTriggerOptions.SetTags(tags)
	}
	if _, ok := d.GetOk("worker"); ok {
		workerModel, err := resourceIBMCdTektonPipelineTriggerMapToWorker(d.Get("worker.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerOptions.SetWorker(workerModel)
	}
	if _, ok := d.GetOk("max_concurrent_runs"); ok {
		createTektonPipelineTriggerOptions.SetMaxConcurrentRuns(int64(d.Get("max_concurrent_runs").(int)))
	}
	if _, ok := d.GetOk("disabled"); ok {
		createTektonPipelineTriggerOptions.SetDisabled(d.Get("disabled").(bool))
	}
	if _, ok := d.GetOk("secret"); ok {
		secretModel, err := resourceIBMCdTektonPipelineTriggerMapToGenericSecret(d.Get("secret.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerOptions.SetSecret(secretModel)
	}
	if _, ok := d.GetOk("cron"); ok {
		createTektonPipelineTriggerOptions.SetCron(d.Get("cron").(string))
	}
	if _, ok := d.GetOk("timezone"); ok {
		createTektonPipelineTriggerOptions.SetTimezone(d.Get("timezone").(string))
	}
	if _, ok := d.GetOk("scm_source"); ok {
		scmSourceModel, err := resourceIBMCdTektonPipelineTriggerMapToTriggerScmSource(d.Get("scm_source.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerOptions.SetScmSource(scmSourceModel)
	}
	if _, ok := d.GetOk("events"); ok {
		eventsModel, err := resourceIBMCdTektonPipelineTriggerMapToEvents(d.Get("events.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerOptions.SetEvents(eventsModel)
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

	trigger := triggerIntf.(*cdtektonpipelinev2.Trigger)
	if err = d.Set("pipeline_id", getTektonPipelineTriggerOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	if err = d.Set("type", trigger.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("name", trigger.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("event_listener", trigger.EventListener); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting event_listener: %s", err))
	}
	if trigger.Tags != nil {
		if err = d.Set("tags", trigger.Tags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting tags: %s", err))
		}
	}
	if trigger.Worker != nil {
		workerMap, err := resourceIBMCdTektonPipelineTriggerWorkerToMap(trigger.Worker)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("worker", []map[string]interface{}{workerMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting worker: %s", err))
		}
	}
	if err = d.Set("max_concurrent_runs", flex.IntValue(trigger.MaxConcurrentRuns)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting max_concurrent_runs: %s", err))
	}
	if err = d.Set("disabled", trigger.Disabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disabled: %s", err))
	}
	if trigger.Secret != nil {
		secretMap, err := resourceIBMCdTektonPipelineTriggerGenericSecretToMap(trigger.Secret)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("secret", []map[string]interface{}{secretMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting secret: %s", err))
		}
	}
	if err = d.Set("cron", trigger.Cron); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cron: %s", err))
	}
	if err = d.Set("timezone", trigger.Timezone); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting timezone: %s", err))
	}
	if trigger.ScmSource != nil {
		scmSourceMap, err := resourceIBMCdTektonPipelineTriggerTriggerScmSourceToMap(trigger.ScmSource)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("scm_source", []map[string]interface{}{scmSourceMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scm_source: %s", err))
		}
	}
	if trigger.Events != nil {
		eventsMap, err := resourceIBMCdTektonPipelineTriggerEventsToMap(trigger.Events)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("events", []map[string]interface{}{eventsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting events: %s", err))
		}
	}
	if err = d.Set("href", trigger.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	properties := []map[string]interface{}{}
	if trigger.Properties != nil {
		for _, propertiesItem := range trigger.Properties {
			propertiesItemMap, err := resourceIBMCdTektonPipelineTriggerTriggerPropertiesItemToMap(&propertiesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			properties = append(properties, propertiesItemMap)
		}
	}
	if err = d.Set("properties", properties); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting properties: %s", err))
	}
	if err = d.Set("webhook_url", trigger.WebhookURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting webhook_url: %s", err))
	}
	if err = d.Set("trigger_id", trigger.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting trigger_id: %s", err))
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
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "pipeline_id"))
	}
	if d.HasChange("type") {
		patchVals.Type = core.StringPtr(d.Get("type").(string))
		hasChange = true
	}
	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("event_listener") {
		patchVals.EventListener = core.StringPtr(d.Get("event_listener").(string))
		hasChange = true
	}
	if d.HasChange("tags") {
		tags := []string{}
		for _, tagsItem := range d.Get("tags").([]interface{}) {
			tags = append(tags, tagsItem.(string))
		}
		patchVals.Tags = tags
		hasChange = true
	}
	if d.HasChange("worker") {
		worker, err := resourceIBMCdTektonPipelineTriggerMapToWorker(d.Get("worker.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Worker = worker
		hasChange = true
	}
	if d.HasChange("max_concurrent_runs") {
		patchVals.MaxConcurrentRuns = core.Int64Ptr(int64(d.Get("max_concurrent_runs").(int)))
		hasChange = true
	}
	if d.HasChange("disabled") {
		patchVals.Disabled = core.BoolPtr(d.Get("disabled").(bool))
		hasChange = true
	}
	if d.HasChange("secret") {
		secret, err := resourceIBMCdTektonPipelineTriggerMapToGenericSecret(d.Get("secret.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Secret = secret
		hasChange = true
	}
	if d.HasChange("cron") {
		patchVals.Cron = core.StringPtr(d.Get("cron").(string))
		hasChange = true
	}
	if d.HasChange("timezone") {
		patchVals.Timezone = core.StringPtr(d.Get("timezone").(string))
		hasChange = true
	}
	if d.HasChange("scm_source") {
		scmSource, err := resourceIBMCdTektonPipelineTriggerMapToTriggerScmSource(d.Get("scm_source.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.ScmSource = scmSource
		hasChange = true
	}
	if d.HasChange("events") {
		events, err := resourceIBMCdTektonPipelineTriggerMapToEvents(d.Get("events.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchVals.Events = events
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

func resourceIBMCdTektonPipelineTriggerMapToTriggerScmSource(modelMap map[string]interface{}) (*cdtektonpipelinev2.TriggerScmSource, error) {
	model := &cdtektonpipelinev2.TriggerScmSource{}
	model.URL = core.StringPtr(modelMap["url"].(string))
	if modelMap["branch"] != nil && modelMap["branch"].(string) != "" {
		model.Branch = core.StringPtr(modelMap["branch"].(string))
	}
	if modelMap["pattern"] != nil && modelMap["pattern"].(string) != "" {
		model.Pattern = core.StringPtr(modelMap["pattern"].(string))
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

func resourceIBMCdTektonPipelineTriggerTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
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
