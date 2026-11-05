// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.0-af5c89a5-20240617-153232
 */

package configurationaggregator

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
)

func DataSourceIbmConfigAggregatorResourceCollectionStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmConfigAggregatorResourceCollectionStatusRead,

		Schema: map[string]*schema.Schema{
			"last_config_refresh_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp at which the configuration was last refreshed.",
			},
			"latest_reconciliation_mode": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Latest collection mode.",
			},
			"reconciliation_modes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Per reconciliation mode details. Only the modes that have run are present; the object is empty if no reconciliation has completed yet.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"automatic": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Status details of a reconciliation mode.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_refresh_time": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp at which the configuration was last refreshed by this reconciliation mode.",
									},
								},
							},
						},
						"manual": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Status details of a reconciliation mode.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_refresh_time": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp at which the configuration was last refreshed by this reconciliation mode.",
									},
								},
							},
						},
						"realtime": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Status details of a reconciliation mode.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"last_refresh_time": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp at which the configuration was last refreshed by this reconciliation mode.",
									},
								},
							},
						},
					},
				},
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the resource collection.",
			},
		},
	}
}

func dataSourceIbmConfigAggregatorResourceCollectionStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	configurationAggregatorClient, err := meta.(conns.ClientSession).ConfigurationAggregatorV1()
	region := getConfigurationInstanceRegion(configurationAggregatorClient, d)
	instanceId := d.Get("instance_id").(string)
	configurationAggregatorClient = getClientWithConfigurationInstanceEndpoint(configurationAggregatorClient, instanceId, region)
	if err != nil {
		// Error is coming from SDK client, so it doesn't need to be discriminated.
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_resource_collection_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getResourceCollectionStatusOptions := &configurationaggregatorv1.GetResourceCollectionStatusOptions{}

	statusResponse, _, err := configurationAggregatorClient.GetResourceCollectionStatusWithContext(context, getResourceCollectionStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetResourceCollectionStatusWithContext failed: %s", err.Error()), "(Data) ibm_config_aggregator_resource_collection_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmConfigAggregatorResourceCollectionStatusID(d))

	if err = d.Set("last_config_refresh_time", flex.DateTimeToString(statusResponse.LastConfigRefreshTime)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_config_refresh_time: %s", err), "(Data) ibm_config_aggregator_resource_collection_status", "read", "set-last_config_refresh_time").GetDiag()
	}

	if err = d.Set("latest_reconciliation_mode", statusResponse.LatestReconciliationMode); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting latest_reconciliation_mode: %s", err), "(Data) ibm_config_aggregator_resource_collection_status", "read", "set-latest_reconciliation_mode").GetDiag()
	}

	reconciliationModes := []map[string]interface{}{}
	if statusResponse.ReconciliationModes != nil {
		modelMap, err := DataSourceIbmConfigAggregatorResourceCollectionStatusStatusResponseReconciliationModesToMap(statusResponse.ReconciliationModes)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_resource_collection_status", "read", "reconciliation_modes-to-map").GetDiag()
		}
		reconciliationModes = append(reconciliationModes, modelMap)
	}
	if err = d.Set("reconciliation_modes", reconciliationModes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting reconciliation_modes: %s", err), "(Data) ibm_config_aggregator_resource_collection_status", "read", "set-reconciliation_modes").GetDiag()
	}

	if err = d.Set("status", statusResponse.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_config_aggregator_resource_collection_status", "read", "set-status").GetDiag()
	}

	return nil
}

// dataSourceIbmConfigAggregatorResourceCollectionStatusID returns a reasonable ID for the list.
func dataSourceIbmConfigAggregatorResourceCollectionStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmConfigAggregatorResourceCollectionStatusStatusResponseReconciliationModesToMap(model *configurationaggregatorv1.StatusResponseReconciliationModes) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Automatic != nil {
		modelMap["automatic"] = []map[string]interface{}{DataSourceIbmConfigAggregatorResourceCollectionStatusReconciliationModeStatusToMap(model.Automatic)}
	}
	if model.Manual != nil {
		modelMap["manual"] = []map[string]interface{}{DataSourceIbmConfigAggregatorResourceCollectionStatusReconciliationModeStatusToMap(model.Manual)}
	}
	if model.Realtime != nil {
		modelMap["realtime"] = []map[string]interface{}{DataSourceIbmConfigAggregatorResourceCollectionStatusReconciliationModeStatusToMap(model.Realtime)}
	}
	return modelMap, nil
}

func DataSourceIbmConfigAggregatorResourceCollectionStatusReconciliationModeStatusToMap(model *configurationaggregatorv1.ReconciliationModeStatus) map[string]interface{} {
	modelMap := make(map[string]interface{})
	if model.LastRefreshTime != nil {
		modelMap["last_refresh_time"] = model.LastRefreshTime.String()
	}
	return modelMap
}
