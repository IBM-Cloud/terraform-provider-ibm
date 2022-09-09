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
)

func DataSourceIBMCdTektonPipeline() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMCdTektonPipelineRead,

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of current instance.",
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
							Computed:    true,
							Description: "UUID.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for the toolchain that contains the Tekton pipeline.",
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
							Computed:    true,
							Description: "SCM source for Tekton pipeline definition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URL of the definition repository.",
									},
									"branch": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A branch from the repo. One of branch or tag must be specified, but only one or the other.",
									},
									"tag": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A tag from the repo. One of branch or tag must be specified, but only one or the other.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The path to the definition's yaml files.",
									},
									"service_instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the SCM repository service instance.",
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
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
							Computed:    true,
							Description: "Property name.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property value.",
						},
						"enum": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Options for `single_select` property type. Only needed when using `single_select` property type.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property type.",
						},
						"path": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A dot notation path for `integration` type properties to select a value from the tool integration.",
						},
					},
				},
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Standard RFC 3339 Date Time String.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Standard RFC 3339 Date Time String.",
			},
			"triggers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tekton pipeline triggers list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "API URL for interacting with the trigger.",
						},
						"event_listener": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID.",
						},
						"properties": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Trigger properties.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Property name.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Property value. Can be empty and should be omitted for `single_select` property type.",
									},
									"enum": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Options for `single_select` property type. Only needed for `single_select` property type.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Property type.",
									},
									"path": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A dot notation path for `integration` type properties to select a value from the tool integration. If left blank the full tool integration data will be used.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "API URL for interacting with the trigger property.",
									},
								},
							},
						},
						"tags": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Trigger tags array.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"worker": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Worker used to run the trigger. If not specified the trigger will use the default pipeline worker.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the worker. Computed based on the worker ID.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the worker. Computed based on the worker ID.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the worker.",
									},
								},
							},
						},
						"max_concurrent_runs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Defines the maximum number of concurrent runs for this trigger. Omit this property to disable the concurrency limit.",
						},
						"disabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag whether the trigger is disabled. If omitted the trigger is enabled by default.",
						},
						"scm_source": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "SCM source repository for a Git trigger. Only needed for Git triggers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URL of the repository to which the trigger is listening.",
									},
									"branch": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of a branch from the repo. One of branch or tag must be specified, but only one or the other.",
									},
									"pattern": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
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
							Computed:    true,
							Description: "Only needed for Git triggers. Events object defines the events to which this Git trigger listens.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"push": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If true, the trigger listens for 'push' Git webhook events.",
									},
									"pull_request_closed": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If true, the trigger listens for 'close pull request' Git webhook events.",
									},
									"pull_request": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If true, the trigger listens for 'open pull request' or 'update pull request' Git webhook events.",
									},
								},
							},
						},
						"cron": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only needed for timer triggers. Cron expression for timer trigger. Maximum frequency is every 5 minutes.",
						},
						"timezone": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Only needed for timer triggers. Timezone for timer trigger.",
						},
						"secret": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Only needed for generic webhook trigger type. Secret used to start generic webhook trigger.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Secret type.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Secret value, not needed if secret type is `internal_validation`.",
									},
									"source": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Secret location, not needed if secret type is `internal_validation`.",
									},
									"key_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Secret name, not needed if type is `internal_validation`.",
									},
									"algorithm": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Algorithm used for `digest_matches` secret type. Only needed for `digest_matches` secret type.",
									},
								},
							},
						},
						"webhook_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook URL that can be used to trigger pipeline runs.",
						},
					},
				},
			},
			"worker": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Default pipeline worker used to run the pipeline.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the worker. Computed based on the worker ID.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the worker. Computed based on the worker ID.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the worker.",
						},
					},
				},
			},
			"runs_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL for this pipeline showing the list of pipeline runs.",
			},
			"build_number": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The latest pipeline run build number. If this property is absent, the pipeline hasn't had any pipeline runs.",
			},
			"enable_slack_notifications": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag whether to enable slack notifications for this pipeline. When enabled, pipeline run events will be published on all slack integration specified channels in the enclosing toolchain.",
			},
			"enable_partial_cloning": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag whether to enable partial cloning for this pipeline. When partial clone is enabled, only the files contained within the paths specified in definition repositories will be read and cloned. This means symbolic links may not work.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag whether this pipeline is enabled.",
			},
		},
	}
}

func dataSourceIBMCdTektonPipelineRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cdTektonPipelineClient, err := meta.(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{}

	getTektonPipelineOptions.SetID(d.Get("pipeline_id").(string))

	tektonPipeline, response, err := cdTektonPipelineClient.GetTektonPipelineWithContext(context, getTektonPipelineOptions)
	if err != nil {
		log.Printf("[DEBUG] GetTektonPipelineWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelineWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getTektonPipelineOptions.ID))

	if err = d.Set("name", tektonPipeline.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("status", tektonPipeline.Status); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting status: %s", err))
	}

	if err = d.Set("resource_group_id", tektonPipeline.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}

	toolchain := []map[string]interface{}{}
	if tektonPipeline.Toolchain != nil {
		modelMap, err := dataSourceIBMCdTektonPipelineToolchainToMap(tektonPipeline.Toolchain)
		if err != nil {
			return diag.FromErr(err)
		}
		toolchain = append(toolchain, modelMap)
	}
	if err = d.Set("toolchain", toolchain); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting toolchain %s", err))
	}

	definitions := []map[string]interface{}{}
	if tektonPipeline.Definitions != nil {
		for _, modelItem := range tektonPipeline.Definitions {
			modelMap, err := dataSourceIBMCdTektonPipelineDefinitionToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			definitions = append(definitions, modelMap)
		}
	}
	if err = d.Set("definitions", definitions); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definitions %s", err))
	}

	properties := []map[string]interface{}{}
	if tektonPipeline.Properties != nil {
		for _, modelItem := range tektonPipeline.Properties {
			modelMap, err := dataSourceIBMCdTektonPipelinePropertyToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			properties = append(properties, modelMap)
		}
	}
	if err = d.Set("properties", properties); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting properties %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(tektonPipeline.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(tektonPipeline.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	triggers := []map[string]interface{}{}
	if tektonPipeline.Triggers != nil {
		for _, modelItem := range tektonPipeline.Triggers {
			modelMap, err := dataSourceIBMCdTektonPipelineTriggerToMap(modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			triggers = append(triggers, modelMap)
		}
	}
	if err = d.Set("triggers", triggers); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting triggers %s", err))
	}

	worker := []map[string]interface{}{}
	if tektonPipeline.Worker != nil {
		modelMap, err := dataSourceIBMCdTektonPipelineWorkerToMap(tektonPipeline.Worker)
		if err != nil {
			return diag.FromErr(err)
		}
		worker = append(worker, modelMap)
	}
	if err = d.Set("worker", worker); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting worker %s", err))
	}

	if err = d.Set("runs_url", tektonPipeline.RunsURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting runs_url: %s", err))
	}

	if err = d.Set("build_number", flex.IntValue(tektonPipeline.BuildNumber)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting build_number: %s", err))
	}

	if err = d.Set("enable_slack_notifications", tektonPipeline.EnableSlackNotifications); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting enable_slack_notifications: %s", err))
	}

	if err = d.Set("enable_partial_cloning", tektonPipeline.EnablePartialCloning); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting enable_partial_cloning: %s", err))
	}

	if err = d.Set("enabled", tektonPipeline.Enabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting enabled: %s", err))
	}

	return nil
}

func dataSourceIBMCdTektonPipelineToolchainToMap(model *cdtektonpipelinev2.Toolchain) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineDefinitionToMap(model *cdtektonpipelinev2.Definition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ScmSource != nil {
		scmSourceMap, err := dataSourceIBMCdTektonPipelineDefinitionScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineDefinitionScmSourceToMap(model *cdtektonpipelinev2.DefinitionScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.URL != nil {
		modelMap["url"] = *model.URL
	}
	if model.Branch != nil {
		modelMap["branch"] = *model.Branch
	}
	if model.Tag != nil {
		modelMap["tag"] = *model.Tag
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = *model.ServiceInstanceID
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelinePropertyToMap(model *cdtektonpipelinev2.Property) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerToMap(model cdtektonpipelinev2.TriggerIntf) (map[string]interface{}, error) {
	if _, ok := model.(*cdtektonpipelinev2.TriggerManualTrigger); ok {
		return dataSourceIBMCdTektonPipelineTriggerManualTriggerToMap(model.(*cdtektonpipelinev2.TriggerManualTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerScmTrigger); ok {
		return dataSourceIBMCdTektonPipelineTriggerScmTriggerToMap(model.(*cdtektonpipelinev2.TriggerScmTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerTimerTrigger); ok {
		return dataSourceIBMCdTektonPipelineTriggerTimerTriggerToMap(model.(*cdtektonpipelinev2.TriggerTimerTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.TriggerGenericTrigger); ok {
		return dataSourceIBMCdTektonPipelineTriggerGenericTriggerToMap(model.(*cdtektonpipelinev2.TriggerGenericTrigger))
	} else if _, ok := model.(*cdtektonpipelinev2.Trigger); ok {
		modelMap := make(map[string]interface{})
		model := model.(*cdtektonpipelinev2.Trigger)
		if model.Type != nil {
			modelMap["type"] = *model.Type
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.EventListener != nil {
			modelMap["event_listener"] = *model.EventListener
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Properties != nil {
			properties := []map[string]interface{}{}
			for _, propertiesItem := range model.Properties {
				propertiesItemMap, err := dataSourceIBMCdTektonPipelineTriggerPropertiesItemToMap(&propertiesItem)
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
			workerMap, err := dataSourceIBMCdTektonPipelineWorkerToMap(model.Worker)
			if err != nil {
				return modelMap, err
			}
			modelMap["worker"] = []map[string]interface{}{workerMap}
		}
		if model.MaxConcurrentRuns != nil {
			modelMap["max_concurrent_runs"] = *model.MaxConcurrentRuns
		}
		if model.Disabled != nil {
			modelMap["disabled"] = *model.Disabled
		}
		if model.ScmSource != nil {
			scmSourceMap, err := dataSourceIBMCdTektonPipelineTriggerScmSourceToMap(model.ScmSource)
			if err != nil {
				return modelMap, err
			}
			modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
		}
		if model.Events != nil {
			eventsMap, err := dataSourceIBMCdTektonPipelineEventsToMap(model.Events)
			if err != nil {
				return modelMap, err
			}
			modelMap["events"] = []map[string]interface{}{eventsMap}
		}
		if model.Cron != nil {
			modelMap["cron"] = *model.Cron
		}
		if model.Timezone != nil {
			modelMap["timezone"] = *model.Timezone
		}
		if model.Secret != nil {
			secretMap, err := dataSourceIBMCdTektonPipelineGenericSecretToMap(model.Secret)
			if err != nil {
				return modelMap, err
			}
			modelMap["secret"] = []map[string]interface{}{secretMap}
		}
		if model.WebhookURL != nil {
			modelMap["webhook_url"] = *model.WebhookURL
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized cdtektonpipelinev2.TriggerIntf subtype encountered")
	}
}

func dataSourceIBMCdTektonPipelineTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineWorkerToMap(model *cdtektonpipelinev2.Worker) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerScmSourceToMap(model *cdtektonpipelinev2.TriggerScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.URL != nil {
		modelMap["url"] = *model.URL
	}
	if model.Branch != nil {
		modelMap["branch"] = *model.Branch
	}
	if model.Pattern != nil {
		modelMap["pattern"] = *model.Pattern
	}
	if model.BlindConnection != nil {
		modelMap["blind_connection"] = *model.BlindConnection
	}
	if model.HookID != nil {
		modelMap["hook_id"] = *model.HookID
	}
	if model.ServiceInstanceID != nil {
		modelMap["service_instance_id"] = *model.ServiceInstanceID
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineEventsToMap(model *cdtektonpipelinev2.Events) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Push != nil {
		modelMap["push"] = *model.Push
	}
	if model.PullRequestClosed != nil {
		modelMap["pull_request_closed"] = *model.PullRequestClosed
	}
	if model.PullRequest != nil {
		modelMap["pull_request"] = *model.PullRequest
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineGenericSecretToMap(model *cdtektonpipelinev2.GenericSecret) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Source != nil {
		modelMap["source"] = *model.Source
	}
	if model.KeyName != nil {
		modelMap["key_name"] = *model.KeyName
	}
	if model.Algorithm != nil {
		modelMap["algorithm"] = *model.Algorithm
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerManualTriggerToMap(model *cdtektonpipelinev2.TriggerManualTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.EventListener != nil {
		modelMap["event_listener"] = *model.EventListener
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := dataSourceIBMCdTektonPipelineTriggerManualTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := dataSourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = *model.MaxConcurrentRuns
	}
	if model.Disabled != nil {
		modelMap["disabled"] = *model.Disabled
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerManualTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerManualTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerScmTriggerToMap(model *cdtektonpipelinev2.TriggerScmTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.EventListener != nil {
		modelMap["event_listener"] = *model.EventListener
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := dataSourceIBMCdTektonPipelineTriggerScmTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := dataSourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = *model.MaxConcurrentRuns
	}
	if model.Disabled != nil {
		modelMap["disabled"] = *model.Disabled
	}
	if model.ScmSource != nil {
		scmSourceMap, err := dataSourceIBMCdTektonPipelineTriggerScmSourceToMap(model.ScmSource)
		if err != nil {
			return modelMap, err
		}
		modelMap["scm_source"] = []map[string]interface{}{scmSourceMap}
	}
	if model.Events != nil {
		eventsMap, err := dataSourceIBMCdTektonPipelineEventsToMap(model.Events)
		if err != nil {
			return modelMap, err
		}
		modelMap["events"] = []map[string]interface{}{eventsMap}
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerScmTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerScmTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerTimerTriggerToMap(model *cdtektonpipelinev2.TriggerTimerTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.EventListener != nil {
		modelMap["event_listener"] = *model.EventListener
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := dataSourceIBMCdTektonPipelineTriggerTimerTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := dataSourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = *model.MaxConcurrentRuns
	}
	if model.Disabled != nil {
		modelMap["disabled"] = *model.Disabled
	}
	if model.Cron != nil {
		modelMap["cron"] = *model.Cron
	}
	if model.Timezone != nil {
		modelMap["timezone"] = *model.Timezone
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerTimerTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerTimerTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerGenericTriggerToMap(model *cdtektonpipelinev2.TriggerGenericTrigger) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.EventListener != nil {
		modelMap["event_listener"] = *model.EventListener
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Properties != nil {
		properties := []map[string]interface{}{}
		for _, propertiesItem := range model.Properties {
			propertiesItemMap, err := dataSourceIBMCdTektonPipelineTriggerGenericTriggerPropertiesItemToMap(&propertiesItem)
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
		workerMap, err := dataSourceIBMCdTektonPipelineWorkerToMap(model.Worker)
		if err != nil {
			return modelMap, err
		}
		modelMap["worker"] = []map[string]interface{}{workerMap}
	}
	if model.MaxConcurrentRuns != nil {
		modelMap["max_concurrent_runs"] = *model.MaxConcurrentRuns
	}
	if model.Disabled != nil {
		modelMap["disabled"] = *model.Disabled
	}
	if model.Secret != nil {
		secretMap, err := dataSourceIBMCdTektonPipelineGenericSecretToMap(model.Secret)
		if err != nil {
			return modelMap, err
		}
		modelMap["secret"] = []map[string]interface{}{secretMap}
	}
	if model.WebhookURL != nil {
		modelMap["webhook_url"] = *model.WebhookURL
	}
	return modelMap, nil
}

func dataSourceIBMCdTektonPipelineTriggerGenericTriggerPropertiesItemToMap(model *cdtektonpipelinev2.TriggerGenericTriggerPropertiesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Value != nil {
		modelMap["value"] = *model.Value
	}
	if model.Enum != nil {
		modelMap["enum"] = model.Enum
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Path != nil {
		modelMap["path"] = *model.Path
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}
