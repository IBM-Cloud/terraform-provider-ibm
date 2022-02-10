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

func ResourceIBMTektonPipelineProperty() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMTektonPipelinePropertyCreate,
		ReadContext:   ResourceIBMTektonPipelinePropertyRead,
		UpdateContext: ResourceIBMTektonPipelinePropertyUpdate,
		DeleteContext: ResourceIBMTektonPipelinePropertyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_tekton_pipeline_property", "pipeline_id"),
				Description:  "The tekton pipeline ID.",
			},
			"create_tekton_pipeline_properties_request": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"env_properties": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Pipeline properties list.",
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
								},
							},
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
		},
	}
}

func ResourceIBMTektonPipelinePropertyValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_tekton_pipeline_property", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMTektonPipelinePropertyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelinePropertiesOptions := &continuousdeliverypipelinev2.CreateTektonPipelinePropertiesOptions{}

	createTektonPipelinePropertiesOptions.SetPipelineID(d.Get("pipeline_id").(string))
	if _, ok := d.GetOk("create_tekton_pipeline_properties_request"); ok {
		createTektonPipelinePropertiesRequest, err := ResourceIBMTektonPipelinePropertyMapToCreateTektonPipelinePropertiesRequest(d.Get("create_tekton_pipeline_properties_request.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelinePropertiesOptions.SetCreateTektonPipelinePropertiesRequest(createTektonPipelinePropertiesRequest)
	}

	envProperties, response, err := continuousDeliveryPipelineClient.CreateTektonPipelinePropertiesWithContext(context, createTektonPipelinePropertiesOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTektonPipelinePropertiesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTektonPipelinePropertiesWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createTektonPipelinePropertiesOptions.PipelineID, *property.Name))

	return ResourceIBMTektonPipelinePropertyRead(context, d, meta)
}

func ResourceIBMTektonPipelinePropertyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelinePropertyOptions := &continuousdeliverypipelinev2.GetTektonPipelinePropertyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
	getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

	property, response, err := continuousDeliveryPipelineClient.GetTektonPipelinePropertyWithContext(context, getTektonPipelinePropertyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTektonPipelinePropertyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelinePropertyWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("pipeline_id", getTektonPipelinePropertyOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	// TODO: handle argument of type CreateTektonPipelinePropertiesRequest
	if err = d.Set("value", property.Value); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting value: %s", err))
	}
	if err = d.Set("options", property.Options); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting options: %s", err))
	}
	if err = d.Set("type", property.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("path", property.Path); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting path: %s", err))
	}

	return nil
}

func ResourceIBMTektonPipelinePropertyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTektonPipelinePropertyOptions := &continuousdeliverypipelinev2.ReplaceTektonPipelinePropertyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTektonPipelinePropertyOptions.SetPipelineID(parts[0])
	replaceTektonPipelinePropertyOptions.SetPropertyName(parts[1])

	hasChange := false

	if d.HasChange("pipeline_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "pipeline_id"))
	}
	if d.HasChange("create_tekton_pipeline_properties_request") {
		createTektonPipelinePropertiesRequest, err := ResourceIBMTektonPipelinePropertyMapToCreateTektonPipelinePropertiesRequest(d.Get("create_tekton_pipeline_properties_request.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		replaceTektonPipelinePropertyOptions.SetCreateTektonPipelinePropertiesRequest(createTektonPipelinePropertiesRequest)
		hasChange = true
	}

	if hasChange {
		_, response, err := continuousDeliveryPipelineClient.ReplaceTektonPipelinePropertyWithContext(context, replaceTektonPipelinePropertyOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceTektonPipelinePropertyWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ReplaceTektonPipelinePropertyWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMTektonPipelinePropertyRead(context, d, meta)
}

func ResourceIBMTektonPipelinePropertyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelinePropertyOptions := &continuousdeliverypipelinev2.DeleteTektonPipelinePropertyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelinePropertyOptions.SetPipelineID(parts[0])
	deleteTektonPipelinePropertyOptions.SetPropertyName(parts[1])

	response, err := continuousDeliveryPipelineClient.DeleteTektonPipelinePropertyWithContext(context, deleteTektonPipelinePropertyOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTektonPipelinePropertyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTektonPipelinePropertyWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMTektonPipelinePropertyMapToCreateTektonPipelinePropertiesRequest(modelMap map[string]interface{}) (continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestIntf, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequest{}
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
	if modelMap["env_properties"] != nil {
		envProperties := []continuousdeliverypipelinev2.Property{}
		for _, envPropertiesItem := range modelMap["env_properties"].([]interface{}) {
			envPropertiesItemModel, err := ResourceIBMTektonPipelinePropertyMapToProperty(envPropertiesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			envProperties = append(envProperties, *envPropertiesItemModel)
		}
		model.EnvProperties = envProperties
	}
	return model, nil
}

func ResourceIBMTektonPipelinePropertyMapToProperty(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.Property, error) {
	model := &continuousdeliverypipelinev2.Property{}
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

func ResourceIBMTektonPipelinePropertyMapToCreateTektonPipelinePropertiesRequestProperty(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestProperty, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestProperty{}
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

func ResourceIBMTektonPipelinePropertyMapToCreateTektonPipelinePropertiesRequestEnvProperties(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestEnvProperties, error) {
	model := &continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestEnvProperties{}
	envProperties := []continuousdeliverypipelinev2.Property{}
	for _, envPropertiesItem := range modelMap["env_properties"].([]interface{}) {
		envPropertiesItemModel, err := ResourceIBMTektonPipelinePropertyMapToProperty(envPropertiesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		envProperties = append(envProperties, *envPropertiesItemModel)
	}
	model.EnvProperties = envProperties
	return model, nil
}

func ResourceIBMTektonPipelinePropertyCreateTektonPipelinePropertiesRequestToMap(model continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestIntf) (map[string]interface{}, error) {
	if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestProperty); ok {
		return ResourceIBMTektonPipelinePropertyCreateTektonPipelinePropertiesRequestPropertyToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestProperty))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestEnvProperties); ok {
		return ResourceIBMTektonPipelinePropertyCreateTektonPipelinePropertiesRequestEnvPropertiesToMap(model.(*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestEnvProperties))
	} else if _, ok := model.(*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequest); ok {
		modelMap := make(map[string]interface{})
		model := model.(*continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequest)
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
		if model.EnvProperties != nil {
			envProperties := []map[string]interface{}{}
			for _, envPropertiesItem := range model.EnvProperties {
				envPropertiesItemMap, err := ResourceIBMTektonPipelinePropertyPropertyToMap(&envPropertiesItem)
				if err != nil {
					return modelMap, err
				}
				envProperties = append(envProperties, envPropertiesItemMap)
			}
			modelMap["env_properties"] = envProperties
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestIntf subtype encountered")
	}
}

func ResourceIBMTektonPipelinePropertyPropertyToMap(model *continuousdeliverypipelinev2.Property) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelinePropertyCreateTektonPipelinePropertiesRequestPropertyToMap(model *continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestProperty) (map[string]interface{}, error) {
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

func ResourceIBMTektonPipelinePropertyCreateTektonPipelinePropertiesRequestEnvPropertiesToMap(model *continuousdeliverypipelinev2.CreateTektonPipelinePropertiesRequestEnvProperties) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	envProperties := []map[string]interface{}{}
	for _, envPropertiesItem := range model.EnvProperties {
		envPropertiesItemMap, err := ResourceIBMTektonPipelinePropertyPropertyToMap(&envPropertiesItem)
		if err != nil {
			return modelMap, err
		}
		envProperties = append(envProperties, envPropertiesItemMap)
	}
	modelMap["env_properties"] = envProperties
	return modelMap, nil
}
