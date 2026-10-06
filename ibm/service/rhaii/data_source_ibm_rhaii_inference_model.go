// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const rhaiiInferenceModelDataSourceName = "ibm_rhaii_inference_model"

// DataSourceIBMRhaiiInferenceModel reads the details of one inference model of
// a Red Hat AI Inference project.
func DataSourceIBMRhaiiInferenceModel() *schema.Resource {
	s := map[string]*schema.Schema{
		"project_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The ID of the Red Hat AI Inference project.",
		},
		"model": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validate.InvokeDataSourceValidator(rhaiiInferenceModelDataSourceName, "model"),
			Description:  "The ID of the model, for example `llama-3-3-70b-instruct`.",
		},
		"location": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The region of the project. When not set, the endpoint from the provider configuration is used.",
		},
		"identifier":           {Type: schema.TypeString, Computed: true, Description: "The unique identifier of the model."},
		"provider_id":          {Type: schema.TypeString, Computed: true, Description: "The ID of the provider that owns the model."},
		"provider_resource_id": {Type: schema.TypeString, Computed: true, Description: "The ID of the model in the provider."},
		"type":                 {Type: schema.TypeString, Computed: true, Description: "The resource type, `model`."},
		"model_type":           {Type: schema.TypeString, Computed: true, Description: "The model type, for example `llm`, `embedding` or `rerank`."},
	}
	for k, v := range rhaiiInferenceModelMetadataSchema(true) {
		s[k] = v
	}

	return &schema.Resource{
		ReadContext: dataSourceIBMRhaiiInferenceModelRead,
		Schema:      s,
	}
}

// DataSourceIBMRhaiiInferenceModelValidator validates the arguments of ibm_rhaii_inference_model.
func DataSourceIBMRhaiiInferenceModelValidator() *validate.ResourceValidator {
	validateSchema := []validate.ValidateSchema{
		{
			Identifier:                 "model",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-z0-9-]+$`,
			MinValueLength:             1,
			MaxValueLength:             100,
		},
	}
	return &validate.ResourceValidator{ResourceName: rhaiiInferenceModelDataSourceName, Schema: validateSchema}
}

func dataSourceIBMRhaiiInferenceModelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, err := rhaiiClientForLocation(meta, d.Get("location").(string))
	if err != nil {
		return rhaiiDiagFor(err, rhaiiInferenceModelDataSourceName, "read", "initialize-client")
	}

	projectID := d.Get("project_id").(string)
	modelID := d.Get("model").(string)
	model, _, err := client.GetInferenceModelWithContext(ctx, client.NewGetInferenceModelOptions(projectID, modelID))
	if err != nil {
		return rhaiiDiagFor(fmt.Errorf("GetInferenceModelWithContext failed: %s", err), rhaiiInferenceModelDataSourceName, "read", "get-model")
	}

	values, err := flattenRhaiiInferenceModelMetadata(model.Metadata, true)
	if err != nil {
		return rhaiiDiagFor(err, rhaiiInferenceModelDataSourceName, "read", "flatten-model")
	}
	values["identifier"] = strDeref(model.Identifier)
	values["provider_id"] = strDeref(model.ProviderID)
	values["provider_resource_id"] = strDeref(model.ProviderResourceID)
	values["type"] = strDeref(model.Type)
	values["model_type"] = strDeref(model.ModelType)

	d.SetId(fmt.Sprintf("%s/%s", projectID, modelID))
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			return rhaiiDiagFor(fmt.Errorf("error setting %s: %s", k, err), rhaiiInferenceModelDataSourceName, "read", "set-attributes")
		}
	}
	return nil
}
