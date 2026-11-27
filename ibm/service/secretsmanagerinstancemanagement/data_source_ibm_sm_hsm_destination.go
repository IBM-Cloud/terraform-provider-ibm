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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmDestination() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmDestinationRead,

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
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination ID.",
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
	}
}

func dataSourceIbmSmHsmDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "(Data) ibm_sm_hsm_destination", "read")
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions{}

	getInstanceHsmDestinationOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceHsmDestinationOptions.SetDestinationID(d.Get("destination_id").(string))

	hSMDestination, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationWithContext(context, getInstanceHsmDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getInstanceHsmDestinationOptions.InstanceID, *getInstanceHsmDestinationOptions.DestinationID))

	if !core.IsNil(hSMDestination.Href) {
		if err = d.Set("href", hSMDestination.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting href: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-href").GetDiag()
		}
	}

	if !core.IsNil(hSMDestination.AccountID) {
		if err = d.Set("account_id", hSMDestination.AccountID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting account_id: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-account_id").GetDiag()
		}
	}

	if err = d.Set("name", hSMDestination.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting name: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-name").GetDiag()
	}

	if !core.IsNil(hSMDestination.Description) {
		if err = d.Set("description", hSMDestination.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting description: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-description").GetDiag()
		}
	}

	if err = d.Set("vendor", hSMDestination.Vendor); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting vendor: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-vendor").GetDiag()
	}

	if err = d.Set("provisioning_state", hSMDestination.ProvisioningState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting provisioning_state: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-provisioning_state").GetDiag()
	}

	if err = d.Set("health_state", hSMDestination.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_state: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-health_state").GetDiag()
	}

	if err = d.Set("health_as_of", flex.DateTimeToString(hSMDestination.HealthAsOf)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_as_of: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-health_as_of").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(hSMDestination.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_at: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(hSMDestination.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(hSMDestination.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_at: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-updated_at").GetDiag()
		}
	}

	if err = d.Set("created_by", hSMDestination.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_by: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(hSMDestination.UpdatedBy) {
		if err = d.Set("updated_by", hSMDestination.UpdatedBy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_by: %s", err), "(Data) ibm_sm_hsm_destination", "read", "set-updated_by").GetDiag()
		}
	}

	return nil
}
