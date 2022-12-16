// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func DataSourceIBMMetricsRouterTargets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMMetricsRouterTargetsRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the target resource.",
			},
			"targets": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of target resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UUID of the target resource.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the target resource.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crn of the target resource.",
						},
						"destination_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the destination service instance or resource.",
						},
						"target_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the target.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Include this optional field if you used it to create a target in a different region other than the one you are connected.",
						},
						"write_status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The status of the write attempt to the target with the provided destination info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status such as failed or success.",
									},
									"last_failure": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp of the failure.",
									},
									"reason_for_last_failure": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Detailed description of the cause of the failure.",
									},
								},
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp of the target creation time.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp of the target last updated time.",
						},
						"api_version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The API version of the target.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMMetricsRouterTargetsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	listTargetsOptions := &metricsrouterv3.ListTargetsOptions{}

	targetList, response, err := metricsRouterClient.ListTargetsWithContext(context, listTargetsOptions)
	if err != nil {
		log.Printf("[DEBUG] ListTargetsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListTargetsWithContext failed %s\n%s", err, response))
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchTargets []metricsrouterv3.Target
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range targetList.Targets {
			if *data.Name == name {
				matchTargets = append(matchTargets, data)
			}
		}
	} else {
		matchTargets = targetList.Targets
	}
	targetList.Targets = matchTargets

	if suppliedFilter {
		if len(targetList.Targets) == 0 {
			return diag.FromErr(fmt.Errorf("no Targets found with name %s", name))
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMMetricsRouterTargetsID(d))
	}

	targets := []map[string]interface{}{}
	if targetList.Targets != nil {
		for _, modelItem := range targetList.Targets {
			modelMap, err := dataSourceIBMMetricsRouterTargetsTargetToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			targets = append(targets, modelMap)
		}
	}
	if err = d.Set("targets", targets); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting targets %s", err))
	}

	return nil
}

// dataSourceIBMMetricsRouterTargetsID returns a reasonable ID for the list.
func dataSourceIBMMetricsRouterTargetsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIBMMetricsRouterTargetsTargetToMap(model *metricsrouterv3.Target) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.DestinationCRN != nil {
		modelMap["destination_crn"] = *model.DestinationCRN
	}
	if model.TargetType != nil {
		modelMap["target_type"] = *model.TargetType
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.WriteStatus != nil {
		writeStatusMap, err := dataSourceIBMMetricsRouterTargetsWriteStatusToMap(model.WriteStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["write_status"] = []map[string]interface{}{writeStatusMap}
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.APIVersion != nil {
		modelMap["api_version"] = *model.APIVersion
	}
	return modelMap, nil
}

func dataSourceIBMMetricsRouterTargetsWriteStatusToMap(model *metricsrouterv3.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = *model.ReasonForLastFailure
	}
	return modelMap, nil
}
