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

func ResourceIBMTektonPipelineDefinition() *schema.Resource {
	return &schema.Resource{
		CreateContext: ResourceIBMTektonPipelineDefinitionCreate,
		ReadContext:   ResourceIBMTektonPipelineDefinitionRead,
		UpdateContext: ResourceIBMTektonPipelineDefinitionUpdate,
		DeleteContext: ResourceIBMTektonPipelineDefinitionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_tekton_pipeline_definition", "pipeline_id"),
				Description:  "The tekton pipeline ID.",
			},
			"scm_source": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Scm source for tekton pipeline defintion.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "General href URL.",
						},
						"branch": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A branch of the repo.",
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
				Computed:    true,
				Description: "UUID.",
			},
			"definition_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID.",
			},
		},
	}
}

func ResourceIBMTektonPipelineDefinitionValidator() *validate.ResourceValidator {
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_tekton_pipeline_definition", Schema: validateSchema}
	return &resourceValidator
}

func ResourceIBMTektonPipelineDefinitionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	createTektonPipelineDefinitionOptions := &continuousdeliverypipelinev2.CreateTektonPipelineDefinitionOptions{}

	createTektonPipelineDefinitionOptions.SetPipelineID(d.Get("pipeline_id").(string))
	if _, ok := d.GetOk("scm_source"); ok {
		scmSource, err := ResourceIBMTektonPipelineDefinitionMapToDefinitionScmSource(d.Get("scm_source.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createTektonPipelineDefinitionOptions.SetScmSource(scmSource)
	}

	definition, response, err := continuousDeliveryPipelineClient.CreateTektonPipelineDefinitionWithContext(context, createTektonPipelineDefinitionOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTektonPipelineDefinitionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateTektonPipelineDefinitionWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createTektonPipelineDefinitionOptions.PipelineID, *definition.ID))

	return ResourceIBMTektonPipelineDefinitionRead(context, d, meta)
}

func ResourceIBMTektonPipelineDefinitionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineDefinitionOptions := &continuousdeliverypipelinev2.GetTektonPipelineDefinitionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
	getTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])

	definition, response, err := continuousDeliveryPipelineClient.GetTektonPipelineDefinitionWithContext(context, getTektonPipelineDefinitionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTektonPipelineDefinitionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTektonPipelineDefinitionWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("pipeline_id", getTektonPipelineDefinitionOptions.PipelineID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_id: %s", err))
	}
	if definition.ScmSource != nil {
		scmSourceMap, err := ResourceIBMTektonPipelineDefinitionDefinitionScmSourceToMap(definition.ScmSource)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("scm_source", []map[string]interface{}{scmSourceMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting scm_source: %s", err))
		}
	}
	if err = d.Set("service_instance_id", definition.ServiceInstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting service_instance_id: %s", err))
	}
	if err = d.Set("definition_id", definition.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definition_id: %s", err))
	}

	return nil
}

func ResourceIBMTektonPipelineDefinitionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTektonPipelineDefinitionOptions := &continuousdeliverypipelinev2.ReplaceTektonPipelineDefinitionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	replaceTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
	replaceTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])
	replaceTektonPipelineDefinitionOptions.SetServiceInstanceID(d.Get("service_instance_id").(string))
	replaceTektonPipelineDefinitionOptions.SetID(parts[1])

	hasChange := false

	if d.HasChange("pipeline_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "pipeline_id"))
	}
	if d.HasChange("scm_source") {
		scmSource, err := ResourceIBMTektonPipelineDefinitionMapToDefinitionScmSource(d.Get("scm_source.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		replaceTektonPipelineDefinitionOptions.SetScmSource(scmSource)
		hasChange = true
	}

	if hasChange {
		_, response, err := continuousDeliveryPipelineClient.ReplaceTektonPipelineDefinitionWithContext(context, replaceTektonPipelineDefinitionOptions)
		if err != nil {
			log.Printf("[DEBUG] ReplaceTektonPipelineDefinitionWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ReplaceTektonPipelineDefinitionWithContext failed %s\n%s", err, response))
		}
	}

	return ResourceIBMTektonPipelineDefinitionRead(context, d, meta)
}

func ResourceIBMTektonPipelineDefinitionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineDefinitionOptions := &continuousdeliverypipelinev2.DeleteTektonPipelineDefinitionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
	deleteTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])

	response, err := continuousDeliveryPipelineClient.DeleteTektonPipelineDefinitionWithContext(context, deleteTektonPipelineDefinitionOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTektonPipelineDefinitionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTektonPipelineDefinitionWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func ResourceIBMTektonPipelineDefinitionMapToDefinitionScmSource(modelMap map[string]interface{}) (*continuousdeliverypipelinev2.DefinitionScmSource, error) {
	model := &continuousdeliverypipelinev2.DefinitionScmSource{}
	model.URL = core.StringPtr(modelMap["url"].(string))
	model.Branch = core.StringPtr(modelMap["branch"].(string))
	model.Tag = core.StringPtr(modelMap["tag"].(string))
	model.Path = core.StringPtr(modelMap["path"].(string))
	return model, nil
}

func ResourceIBMTektonPipelineDefinitionDefinitionScmSourceToMap(model *continuousdeliverypipelinev2.DefinitionScmSource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["url"] = model.URL
	modelMap["branch"] = model.Branch
	modelMap["tag"] = model.Tag
	modelMap["path"] = model.Path
	return modelMap, nil
}
