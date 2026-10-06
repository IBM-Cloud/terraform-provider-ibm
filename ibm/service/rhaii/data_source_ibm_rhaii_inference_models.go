// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const rhaiiInferenceModelsDataSourceName = "ibm_rhaii_inference_models"

// DataSourceIBMRhaiiInferenceModels lists the inference models that are
// available in a Red Hat AI Inference project.
func DataSourceIBMRhaiiInferenceModels() *schema.Resource {
	modelSchema := map[string]*schema.Schema{
		"id":       {Type: schema.TypeString, Computed: true, Description: "The ID of the model. Use this value as `model` in chat completion requests."},
		"object":   {Type: schema.TypeString, Computed: true, Description: "The object type, `model`."},
		"created":  {Type: schema.TypeInt, Computed: true, Description: "The Unix timestamp in seconds when the model was created."},
		"owned_by": {Type: schema.TypeString, Computed: true, Description: "The owner of the model."},
	}
	for k, v := range rhaiiInferenceModelMetadataSchema(false) {
		modelSchema[k] = v
	}

	return &schema.Resource{
		ReadContext: dataSourceIBMRhaiiInferenceModelsRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the Red Hat AI Inference project.",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region of the project. When not set, the endpoint from the provider configuration is used.",
			},
			"models": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The inference models that are available in the project.",
				Elem:        &schema.Resource{Schema: modelSchema},
			},
		},
	}
}

func dataSourceIBMRhaiiInferenceModelsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, err := rhaiiClientForLocation(meta, d.Get("location").(string))
	if err != nil {
		return rhaiiDiagFor(err, rhaiiInferenceModelsDataSourceName, "read", "initialize-client")
	}

	projectID := d.Get("project_id").(string)
	collection, _, err := client.ListInferenceModelsWithContext(ctx, client.NewListInferenceModelsOptions(projectID))
	if err != nil {
		return rhaiiDiagFor(fmt.Errorf("ListInferenceModelsWithContext failed: %s", err), rhaiiInferenceModelsDataSourceName, "read", "list-models")
	}

	models := make([]interface{}, 0, len(collection.Data))
	for _, model := range collection.Data {
		m, err := flattenRhaiiInferenceModelMetadata(model.CustomMetadata, false)
		if err != nil {
			return rhaiiDiagFor(err, rhaiiInferenceModelsDataSourceName, "read", "flatten-model")
		}
		m["id"] = strDeref(model.ID)
		m["object"] = strDeref(model.Object)
		m["owned_by"] = strDeref(model.OwnedBy)
		if model.Created != nil {
			m["created"] = int(*model.Created)
		}
		models = append(models, m)
	}

	d.SetId(projectID)
	if err := d.Set("models", models); err != nil {
		return rhaiiDiagFor(fmt.Errorf("error setting models: %s", err), rhaiiInferenceModelsDataSourceName, "read", "set-models")
	}
	return nil
}
