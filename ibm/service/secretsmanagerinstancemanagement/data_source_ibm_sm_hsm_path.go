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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmPath() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmPathRead,

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
			"path_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path ID.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the path resource.",
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
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of the gateway through which the HSM is reached. Immutable.",
			},
			"provisioning_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Lifecycle state of the path:- `pending_server_ca`: Path created. Waiting for a Server CA to be added and activated.- `pending_account_policy`: Server CA active. Customer must review and confirm the Secrets Manager account policy for the gateway before proceeding.- `provisioning`: Connectivity setup in progress (async). Triggered by POST /build.- `pending_approval`: Connectivity created. Waiting for customer approval.- `completed`: Connectivity established. Path fully operational.- `vpe_provisioning_failed`: Connectivity setup failed. Call POST /reset then retry POST /build.- `crn_already_in_use`: Connectivity setup failed — the gateway CRN is already in use by another Secrets Manager instance in the same region. Obtain a new CRN and update via PATCH, then call POST /reset.- `approval_declined`: Customer declined the pending approval. Call POST /reset.- `deleting`: Delete initiated; async cleanup in progress.- `deletion_failed`: Async deletion failed. Retry DELETE to re-initiate.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.- `evaluating`: A health check is in progress for this path.- `in_progress`: An apply operation is running across the destination.- `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.- `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval, approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is missing or expired, no_valid_server_ca takes priority.- `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.- `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.",
			},
			"health_as_of": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed field.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed field. Omitted until first update.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed field.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed field. Omitted until first update.",
			},
		},
	}
}

func dataSourceIbmSmHsmPathRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_path", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}

	getInstanceHsmDestinationPathOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceHsmDestinationPathOptions.SetDestinationID(d.Get("destination_id").(string))
	getInstanceHsmDestinationPathOptions.SetPathID(d.Get("path_id").(string))

	path, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathWithContext(context, getInstanceHsmDestinationPathOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_path", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *getInstanceHsmDestinationPathOptions.InstanceID, *getInstanceHsmDestinationPathOptions.DestinationID, *getInstanceHsmDestinationPathOptions.PathID))

	if !core.IsNil(path.Href) {
		if err = d.Set("href", path.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting href: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-href").GetDiag()
		}
	}

	if err = d.Set("name", path.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting name: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-name").GetDiag()
	}

	if !core.IsNil(path.Description) {
		if err = d.Set("description", path.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting description: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-description").GetDiag()
		}
	}

	if err = d.Set("crn", path.Crn); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting crn: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-crn").GetDiag()
	}

	if err = d.Set("provisioning_state", path.ProvisioningState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting provisioning_state: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-provisioning_state").GetDiag()
	}

	if err = d.Set("health_state", path.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_state: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-health_state").GetDiag()
	}

	if err = d.Set("health_as_of", flex.DateTimeToString(path.HealthAsOf)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_as_of: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-health_as_of").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(path.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_at: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(path.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(path.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_at: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-updated_at").GetDiag()
		}
	}

	if err = d.Set("created_by", path.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_by: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(path.UpdatedBy) {
		if err = d.Set("updated_by", path.UpdatedBy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_by: %s", err), "(Data) ibm_sm_hsm_path", "read", "set-updated_by").GetDiag()
		}
	}

	return nil
}
