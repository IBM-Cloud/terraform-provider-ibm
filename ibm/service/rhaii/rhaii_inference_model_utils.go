// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// rhaiiInferenceModelMetadataSchema returns the attributes built from the IBM
// metadata of an inference model. withModelCard adds the model card, which can
// be large and is only returned by the single model data source.
func rhaiiInferenceModelMetadataSchema(withModelCard bool) map[string]*schema.Schema {
	computed := func(t schema.ValueType, description string) *schema.Schema {
		return &schema.Schema{Type: t, Computed: true, Description: description}
	}
	s := map[string]*schema.Schema{
		"name":         computed(schema.TypeString, "The name of the model."),
		"display_name": computed(schema.TypeString, "The display name of the model."),
		"state":        computed(schema.TypeString, "The state of the model, for example `model_loaded`."),
		"status":       computed(schema.TypeString, "The status of the model, for example `loaded`."),
		"total_params": computed(schema.TypeInt, "The total number of parameters of the model."),
		"config_json":  computed(schema.TypeString, "The model configuration (architecture details and hyperparameters) as a JSON string."),
		"pricing": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The pricing measures of the model.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"input_measure":  computed(schema.TypeString, "The billing measure for input tokens."),
					"output_measure": computed(schema.TypeString, "The billing measure for output tokens."),
				},
			},
		},
		"validation": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The validation information of a custom model.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"status":       computed(schema.TypeString, "The validation status."),
					"error":        computed(schema.TypeString, "The validation error message, if any."),
					"completed_at": computed(schema.TypeString, "The date when the validation completed."),
				},
			},
		},
	}
	if withModelCard {
		s["model_card"] = computed(schema.TypeString, "The model card in markdown format.")
		s["source_json"] = computed(schema.TypeString, "The source of a custom model (huggingface, ibmcos, oci or s3) as a JSON string.")
	}
	return s
}

// flattenRhaiiInferenceModelMetadata flattens the IBM metadata of a model into a
// map whose keys match rhaiiInferenceModelMetadataSchema.
func flattenRhaiiInferenceModelMetadata(m *rhaiiv1.InferenceModelMetadata, withModelCard bool) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	if m == nil {
		return out, nil
	}
	out["name"] = strDeref(m.Name)
	out["display_name"] = strDeref(m.DisplayName)
	out["state"] = strDeref(m.State)
	out["status"] = strDeref(m.Status)
	if m.TotalParams != nil {
		out["total_params"] = int(*m.TotalParams)
	}

	configJSON, err := rhaiiToJSON(m.Config)
	if err != nil {
		return nil, err
	}
	out["config_json"] = configJSON

	if m.Pricing != nil {
		out["pricing"] = []interface{}{map[string]interface{}{
			"input_measure":  strDeref(m.Pricing.InputMeasure),
			"output_measure": strDeref(m.Pricing.OutputMeasure),
		}}
	}
	if m.Validation != nil {
		out["validation"] = []interface{}{map[string]interface{}{
			"status":       strDeref(m.Validation.Status),
			"error":        strDeref(m.Validation.Error),
			"completed_at": strDeref(m.Validation.CompletedAt),
		}}
	}

	if withModelCard {
		out["model_card"] = strDeref(m.ModelCard)
		sourceJSON, err := rhaiiToJSON(m.Source)
		if err != nil {
			return nil, err
		}
		out["source_json"] = sourceJSON
	}
	return out, nil
}
