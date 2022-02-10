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

func ResourceIBMTektonPipelineTriggerProperty() *schema.Resource {
	return &schema.Resource{
		CreateContext:   ResourceIBMTektonPipelineTriggerPropertyCreate,
		ReadContext:     ResourceIBMTektonPipelineTriggerPropertyRead,
		UpdateContext:   ResourceIBMTektonPipelineTriggerPropertyUpdate,
		DeleteContext:   ResourceIBMTektonPipelineTriggerPropertyDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_tekton_pipeline_trigger_property", "pipeline_id"),
				Description: "The tekton pipeline ID.",
			},
			"trigger_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_tekton_pipeline_trigger_property", "trigger_id"),
				Description: "The trigger ID.",
			},
			"create_tekton_pipeline_trigger_properties_request": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"properties": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Trigger properties list.",
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
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
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
							Optional:    true,
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
			"value": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String format property value.",
			},
			"options": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Options for SINGLE_SELECT property type.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Property type.",
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "property path for INTEGRATION type properties.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "General href URL.",
			},
		},
	}
}

func ResourceIBMTektonPipelineTriggerPropertyValidator() *validate.ResourceValidator {
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
		validate.ValidateSchema{
			Identifier:                 "trigger_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z]+$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_tekton_pipeline_trigger_property", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMTektonPipelineTriggerPropertyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelineTriggerPropertiesOptions := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesOptions{}

	createTektonPipelineTriggerPropertiesOptions.SetPipelineID(d.Get("pipeline_id").(string))
	createTektonPipelineTriggerPropertiesOptions.SetTriggerID(d.Get("trigger_id").(string))
	if _, ok := d.GetOk("create_tekton_pipeline_trigger_properties_request"); ok {
		createTektonPipelineTriggerPropertiesRequest, err := ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequest(d.Get("create_tekton_pipeline_trigger_properties_request.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineTriggerPropertiesOptions.SetCreateTektonPipelineTriggerPropertiesRequest(createTektonPipelineTriggerPropertiesRequest)
	}

	triggerProperties, response, err := continuousDeliveryPipelineClient.CreateTektonPipelineTriggerPropertiesWithContext(context, createTektonPipelineTriggerPropertiesOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTektonPipelineTriggerPropertiesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTektonPipelineTriggerPropertiesWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createTektonPipelineTriggerPropertiesOptions.PipelineID, *createTektonPipelineTriggerPropertiesOptions.TriggerID, *createTektonPipelineTriggerRequestTriggersTriggersItemTriggerScmTriggerPropertiesItem.Name))

	return ResourceIBMTektonPipelineTriggerPropertyRead(context, d, meta)
}

func ResourceIBMTektonPipelineTriggerPropertyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineTriggerPropertyOptions := &continuousdeliverypipelinev2.GetTektonPipelineTriggerPropertyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
	getTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
	getTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

	property, response, err := continuousDeliveryPipelineClient.GetTektonPipelineTriggerPropertyWithContext(context, getTektonPipelineTriggerPropertyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTektonPipelineTriggerPropertyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelineTriggerPropertyWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("pipeline_id", getTektonPipelineTriggerPropertyOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	if err = d.Set("trigger_id", getTektonPipelineTriggerPropertyOptions.TriggerID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting trigger_id: %s", err))
	}
	// TODO: handle argument of type CreateTektonPipelineTriggerPropertiesRequest
	if err = d.Set("value", triggerProperty.Value); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting value: %s", err))
	}
	if err = d.Set("options", triggerProperty.Options); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting options: %s", err))
	}
	if err = d.Set("type", triggerProperty.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("path", triggerProperty.Path); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting path: %s", err))
	}
	if err = d.Set("href", triggerProperty.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	return nil
}

func ResourceIBMTektonPipelineTriggerPropertyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTektonPipelineTriggerPropertyOptions := &continuousdeliverypipelinev2.ReplaceTektonPipelineTriggerPropertyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
	replaceTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
	replaceTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

	hasChange := false

	if d.HasChange("pipeline_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation." +
				" The resource must be re-created to update this property.", "pipeline_id"))
	}
	if d.HasChange("trigger_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation." +
				" The resource must be re-created to update this property.", "trigger_id"))
	}
	if d.HasChange("create_tekton_pipeline_trigger_properties_request") {
		createTektonPipelineTriggerPropertiesRequest, err := ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequest(d.Get("create_tekton_pipeline_trigger_properties_request.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		replaceTektonPipelineTriggerPropertyOptions.SetCreateTektonPipelineTriggerPropertiesRequest(createTektonPipelineTriggerPropertiesRequest)
		hasChange = true
	}

	if hasChange {
		_, response, err := continuousDeliveryPipelineClient.ReplaceTektonPipelineTriggerPropertyWithContext(context, replaceTektonPipelineTriggerPropertyOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceTektonPipelineTriggerPropertyWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ReplaceTektonPipelineTriggerPropertyWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMTektonPipelineTriggerPropertyRead(context, d, meta)
}

func ResourceIBMTektonPipelineTriggerPropertyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineTriggerPropertyOptions := &continuousdeliverypipelinev2.DeleteTektonPipelineTriggerPropertyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
	deleteTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
	deleteTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

	response, err := continuousDeliveryPipelineClient.DeleteTektonPipelineTriggerPropertyWithContext(context, deleteTektonPipelineTriggerPropertyOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTektonPipelineTriggerPropertyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTektonPipelineTriggerPropertyWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequest(modelMap map[string]interface{}) (continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestIntf, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequest{}
	if modelMap["properties"] != nil {
		properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestPropertiesItem{}
		for _, propertiesItem := range modelMap["properties"].([]interface{}) {
			propertiesItemModel, err := ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequestPropertiesItem(propertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			properties = append(properties, *propertiesItemModel)
		}
		model.Properties = properties
	}
	if modelMap["name"] != nil {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["value"] != nil {
		model.Value = core.StringPtr(modelMap["value"].(string))
	}
	if modelMap["options"] != nil {
	
	}
	if modelMap["type"] != nil {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["path"] != nil {
		model.Path = core.StringPtr(modelMap["path"].(string))
	}
	return model, nil
}

func ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequestPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestPropertiesItem{}
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

func ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequestTriggerProperties(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerProperties, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerProperties{}
	properties := []continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItem{}
	for _, propertiesItem := range modelMap["properties"].([]interface{}) {
		propertiesItemModel, err := ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItem(propertiesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		properties = append(properties, *propertiesItemModel)
	}
	model.Properties = properties
	return model, nil
}

func ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItem(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItem, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItem{}
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

func ResourceIBMTektonPipelineTriggerPropertyMapToCreateTektonPipelineTriggerPropertiesRequestProperty(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestProperty, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestProperty{}
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
	return model, nil
}

func ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestToMap(model continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestIntf) (map[string]interface{}, error) {
	if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerProperties); ok {
		return ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerProperties))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestProperty); ok {
		return ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestPropertyToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestProperty))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequest); ok {
		modelMap := make(map[string]interface{})
		model := model.(*continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequest)
		if model.Properties != nil {
			properties := []map[string]interface{}{}
			for _, propertiesItem := range model.Properties {
				propertiesItemMap, err := ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestPropertiesItemToMap(&propertiesItem)
				if err != nil {
					return modelMap, err
				}
				properties = append(properties, propertiesItemMap)
			}
			modelMap["properties"] = properties
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.Value != nil {
			modelMap["value"] = model.Value
		}
		if model.Options != nil {
			modelMap["options"] = model.Options
		}
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		if model.Path != nil {
			modelMap["path"] = model.Path
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestPropertiesItem) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerProperties) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	properties := []map[string]interface{}{}
	for _, propertiesItem := range model.Properties {
		propertiesItemMap, err := ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItemToMap(&propertiesItem)
		if err != nil {
			return modelMap, err
		}
		properties = append(properties, propertiesItemMap)
	}
	modelMap["properties"] = properties
	return modelMap, nil
}

func ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItemToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestTriggerPropertiesPropertiesItem) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelineTriggerPropertyCreateTektonPipelineTriggerPropertiesRequestPropertyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelineTriggerPropertiesRequestProperty) (map[string]interface{}, error) {
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
	return modelMap, nil
}
