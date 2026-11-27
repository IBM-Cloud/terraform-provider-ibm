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

func DataSourceIbmSmHsmPaths() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmPathsRead,

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
			"paths": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Paths for this destination. At most 20 paths may exist per destination.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Path ID. Immutable. System-managed.",
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
				},
			},
		},
	}
}

func dataSourceIbmSmHsmPathsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_paths", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "(Data) ibm_sm_hsm_paths", "read")
		return tfErr.GetDiag()
	}

	listInstanceHsmDestinationPathsOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathsOptions{}

	listInstanceHsmDestinationPathsOptions.SetInstanceID(d.Get("instance_id").(string))
	listInstanceHsmDestinationPathsOptions.SetDestinationID(d.Get("destination_id").(string))

	pathCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceHsmDestinationPathsWithContext(context, listInstanceHsmDestinationPathsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceHsmDestinationPathsWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_paths", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmHsmPathsID(d))

	paths := []map[string]interface{}{}
	for _, pathsItem := range pathCollection.Paths {
		pathsItemMap, err := DataSourceIbmSmHsmPathsPathToMap(&pathsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_paths", "read", "paths-to-map").GetDiag()
		}
		paths = append(paths, pathsItemMap)
	}
	if err = d.Set("paths", paths); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting paths: %s", err), "(Data) ibm_sm_hsm_paths", "read", "set-paths").GetDiag()
	}

	return nil
}

// dataSourceIbmSmHsmPathsID returns a reasonable ID for the list.
func dataSourceIbmSmHsmPathsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmHsmPathsPathToMap(model *secretsmanagerinstancemanagementv2.Path) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["crn"] = *model.Crn
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
