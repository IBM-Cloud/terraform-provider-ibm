// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmDestinations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmDestinationsRead,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The region of the Secrets Manager instance.",
			},
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secrets Manager instance ID.",
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of HSM Destinations in the collection.",
			},
			"hsm_destinations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of HSM Destinations. At most 5 destinations may exist per SM instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination ID. Immutable. System-managed.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the destination resource.",
						},
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IBM Cloud account ID. System-managed. Required for account policy setup.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional description. Omitted if not set.",
						},
						"vendor": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HSM vendor. Immutable after creation.",
						},
						"provisioning_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Lifecycle state of the destination:- `completed`: Destination created and available for configuration. Initial state on POST.- `deleting`: Hard delete initiated; asynchronous cleanup in progress.- `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.",
						},
						"health_as_of": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Computed field. When the resource was created.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Computed field. Omitted until first update.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IAM identity of creator. Computed field.",
						},
						"updated_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IAM identity of last updater. Computed field. Omitted until first update.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmHsmDestinationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destinations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "(Data) ibm_sm_hsm_destinations", "read")
		return tfErr.GetDiag()
	}

	listInstanceHsmDestinationsOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationsOptions{}

	listInstanceHsmDestinationsOptions.SetInstanceID(d.Get("instance_id").(string))

	hSMDestinationCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceHsmDestinationsWithContext(context, listInstanceHsmDestinationsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceHsmDestinationsWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_destinations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmHsmDestinationsID(d))

	if err = d.Set("total_count", flex.IntValue(hSMDestinationCollection.TotalCount)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting total_count: %s", err), "(Data) ibm_sm_hsm_destinations", "read", "set-total_count").GetDiag()
	}

	hsmDestinations := []map[string]interface{}{}
	for _, hsmDestinationsItem := range hSMDestinationCollection.HsmDestinations {
		hsmDestinationsItemMap, err := DataSourceIbmSmHsmDestinationsHSMDestinationToMap(&hsmDestinationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destinations", "read", "hsm_destinations-to-map").GetDiag()
		}
		hsmDestinations = append(hsmDestinations, hsmDestinationsItemMap)
	}
	if err = d.Set("hsm_destinations", hsmDestinations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting hsm_destinations: %s", err), "(Data) ibm_sm_hsm_destinations", "read", "set-hsm_destinations").GetDiag()
	}

	return nil
}

// dataSourceIbmSmHsmDestinationsID returns a reasonable ID for the list.
func dataSourceIbmSmHsmDestinationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmHsmDestinationsHSMDestinationToMap(model *secretsmanagerinstancemanagementv2.HSMDestination) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["vendor"] = *model.Vendor
	modelMap["provisioning_state"] = *model.ProvisioningState
	modelMap["health_state"] = *model.HealthState
	modelMap["health_as_of"] = model.HealthAsOf.String()
	modelMap["created_at"] = model.CreatedAt.String()
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	modelMap["created_by"] = *model.CreatedBy
	if model.UpdatedBy != nil {
		modelMap["updated_by"] = *model.UpdatedBy
	}
	return modelMap, nil
}
