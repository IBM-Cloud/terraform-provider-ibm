// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package logsrouter

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouting"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
)

// Manually add the function to handle duplicate "ibm_logs_router_targets" name between logs router v1 and v3.
// After an account migrates from logs router v1 to v3, v1 resources will stop working.
// Determine the data source version based on api endpoint.
func DataSourceIBMLogsRouterTargetsByApiEndpoint() *schema.Resource {
	apiEndpoint := os.Getenv("IBMCLOUD_LOGS_ROUTING_API_ENDPOINT")
	if strings.Contains(apiEndpoint, "/api/v3") {
		return DataSourceIBMLogsRouterTargets() // v3
	} else {
		return logsrouting.DataSourceIBMLogsRouterTargets() // v1
	}
}

func DataSourceIBMLogsRouterTargets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMLogsRouterTargetsRead,

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
							Description: "Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Logs Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/logs-router?topic=logs-router-target-monitoring&interface=ui#target-monitoring-ui) for details.",
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
							Description: "The status of the write attempt to the target with the provided endpoint parameters.",
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
						"managed_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Present when the target is enterprise-managed (`managed_by: enterprise`). For account-managed targets this field is omitted.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMLogsRouterTargetsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_router_targets", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listTargetsOptions := &logsrouterv3.ListTargetsOptions{}

	targetCollection, _, err := logsRouterClient.ListTargetsWithContext(context, listTargetsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListTargetsWithContext failed: %s", err.Error()), "(Data) ibm_logs_router_targets", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchTargets []logsrouterv3.Target
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range targetCollection.Targets {
			if *data.Name == name {
				matchTargets = append(matchTargets, data)
			}
		}
	} else {
		matchTargets = targetCollection.Targets
	}
	targetCollection.Targets = matchTargets

	if suppliedFilter {
		if len(targetCollection.Targets) == 0 {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("no Targets found with name %s", name), "(Data) ibm_logs_router_targets", "read", "no-collection-found").GetDiag()
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMLogsRouterTargetsID(d))
	}

	targets := []map[string]interface{}{}
	for _, targetsItem := range targetCollection.Targets {
		targetsItemMap, err := DataSourceIBMLogsRouterTargetsTargetToMap(&targetsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_router_targets", "read", "targets-to-map").GetDiag()
		}
		targets = append(targets, targetsItemMap)
	}
	if err = d.Set("targets", targets); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting targets: %s", err), "(Data) ibm_logs_router_targets", "read", "set-targets").GetDiag()
	}

	return nil
}

// dataSourceIBMLogsRouterTargetsID returns a reasonable ID for the list.
func dataSourceIBMLogsRouterTargetsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMLogsRouterTargetsTargetToMap(model *logsrouterv3.Target) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["crn"] = *model.CRN
	modelMap["destination_crn"] = *model.DestinationCRN
	modelMap["target_type"] = *model.TargetType
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	writeStatusMap, err := DataSourceIBMLogsRouterTargetsWriteStatusToMap(model.WriteStatus)
	if err != nil {
		return modelMap, err
	}
	modelMap["write_status"] = []map[string]interface{}{writeStatusMap}
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()

	if model.ManagedBy != nil {
		modelMap["managed_by"] = *model.ManagedBy
	}

	return modelMap, nil
}

func DataSourceIBMLogsRouterTargetsWriteStatusToMap(model *logsrouterv3.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = *model.ReasonForLastFailure
	}
	return modelMap, nil
}
