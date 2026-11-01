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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmDestination() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmDestinationRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
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
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IBM Cloud Database service instance CRN.",
			},
		},
	}
}

func dataSourceIbmSmDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{}

	getInstanceDestinationOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceDestinationOptions.SetDestinationID(d.Get("destination_id").(string))

	ibmCloudDatabaseDestination, _, err := secretsManagerInstanceManagementClient.GetInstanceDestinationWithContext(context, getInstanceDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceDestinationWithContext failed: %s", err.Error()), "(Data) ibm_sm_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getInstanceDestinationOptions.InstanceID, *getInstanceDestinationOptions.DestinationID))

	if !core.IsNil(ibmCloudDatabaseDestination.Href) {
		if err = d.Set("href", ibmCloudDatabaseDestination.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting href: %s", err), "(Data) ibm_sm_destination", "read", "set-href").GetDiag()
		}
	}

	if err = d.Set("name", ibmCloudDatabaseDestination.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting name: %s", err), "(Data) ibm_sm_destination", "read", "set-name").GetDiag()
	}

	if err = d.Set("type", ibmCloudDatabaseDestination.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting type: %s", err), "(Data) ibm_sm_destination", "read", "set-type").GetDiag()
	}

	if !core.IsNil(ibmCloudDatabaseDestination.Description) {
		if err = d.Set("description", ibmCloudDatabaseDestination.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting description: %s", err), "(Data) ibm_sm_destination", "read", "set-description").GetDiag()
		}
	}

	if err = d.Set("state", ibmCloudDatabaseDestination.State); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting state: %s", err), "(Data) ibm_sm_destination", "read", "set-state").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(ibmCloudDatabaseDestination.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_at: %s", err), "(Data) ibm_sm_destination", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("updated_at", flex.DateTimeToString(ibmCloudDatabaseDestination.UpdatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_at: %s", err), "(Data) ibm_sm_destination", "read", "set-updated_at").GetDiag()
	}

	if !core.IsNil(ibmCloudDatabaseDestination.CreatedBy) {
		if err = d.Set("created_by", ibmCloudDatabaseDestination.CreatedBy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_by: %s", err), "(Data) ibm_sm_destination", "read", "set-created_by").GetDiag()
		}
	}

	if !core.IsNil(ibmCloudDatabaseDestination.Message) {
		if err = d.Set("message", ibmCloudDatabaseDestination.Message); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting message: %s", err), "(Data) ibm_sm_destination", "read", "set-message").GetDiag()
		}
	}

	if err = d.Set("crn", ibmCloudDatabaseDestination.Crn); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting crn: %s", err), "(Data) ibm_sm_destination", "read", "set-crn").GetDiag()
	}

	return nil
}
