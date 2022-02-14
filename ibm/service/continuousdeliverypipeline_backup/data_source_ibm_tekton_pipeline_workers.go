// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package continuousdeliverypipeline

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"
)

func DataSourceIBMTektonPipelineWorkers() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMTektonPipelineWorkersRead,

		Schema: map[string]*schema.Schema{
			"pipeline_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The tekton pipeline ID.",
			},
			"workers": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workers list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "worker name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "worker type.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID.",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMTektonPipelineWorkersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	continuousDeliveryPipelineClient, err := meta.(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return diag.FromErr(err)
	}

	listTektonPipelineWorkersOptions := &continuousdeliverypipelinev2.ListTektonPipelineWorkersOptions{}

	listTektonPipelineWorkersOptions.SetPipelineID(d.Get("pipeline_id").(string))

	workers, response, err := continuousDeliveryPipelineClient.ListTektonPipelineWorkersWithContext(context, listTektonPipelineWorkersOptions)
	if err != nil {
		log.Printf("[DEBUG] ListTektonPipelineWorkersWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListTektonPipelineWorkersWithContext failed %s\n%s", err, response))
	}

	d.SetId(DataSourceIBMTektonPipelineWorkersID(d))

	workersL := []map[string]interface{}{}
	if workers.Workers != nil {
		for _, modelItem := range workers.Workers {
			modelMap, err := DataSourceIBMTektonPipelineWorkersWorkerToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			workersL = append(workersL, modelMap)
		}
	}
	if err = d.Set("workers", workersL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting workers %s", err))
	}

	return nil
}

// DataSourceIBMTektonPipelineWorkersID returns a reasonable ID for the list.
func DataSourceIBMTektonPipelineWorkersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMTektonPipelineWorkersWorkerToMap(model *continuousdeliverypipelinev2.Worker) (map[string]interface{}, error) {
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
