// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.2-f568b7d6-20261016-181021
*/

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmDestinations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmDestinationsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secrets Manager instance ID.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by destination state.",
			},
			"destinations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of destinations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination ID.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the destination resource.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination type.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional description.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the destination was created.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the destination was last updated.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IAM identity that created the destination.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed.",
						},
					},
				},
			},
			"total": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of destinations. Maximum 20 per instance.",
			},
		},
	}
}

func dataSourceIbmSmDestinationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_destinations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listInstanceDestinationsOptions := &secretsmanagerinstancemanagementv2.ListInstanceDestinationsOptions{}

	listInstanceDestinationsOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("state"); ok {
		listInstanceDestinationsOptions.SetState(d.Get("state").(string))
	}

	destinationCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceDestinationsWithContext(context, listInstanceDestinationsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceDestinationsWithContext failed: %s", err.Error()), "(Data) ibm_sm_destinations", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmDestinationsID(d))

	destinations := []map[string]interface{}{}
	for _, destinationsItem := range destinationCollection.Destinations {
		destinationsItemMap, err := DataSourceIbmSmDestinationsDestinationToMap(&destinationsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_destinations", "read", "destinations-to-map").GetDiag()
		}
		destinations = append(destinations, destinationsItemMap)
	}
	if err = d.Set("destinations", destinations); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting destinations: %s", err), "(Data) ibm_sm_destinations", "read", "set-destinations").GetDiag()
	}

	if err = d.Set("total", flex.IntValue(destinationCollection.Total)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting total: %s", err), "(Data) ibm_sm_destinations", "read", "set-total").GetDiag()
	}

	return nil
}

// dataSourceIbmSmDestinationsID returns a reasonable ID for the list.
func dataSourceIbmSmDestinationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmDestinationsDestinationToMap(model *secretsmanagerinstancemanagementv2.Destination) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	modelMap["name"] = *model.Name
	modelMap["type"] = *model.Type
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["state"] = *model.State
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}
