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

func DataSourceIbmSmHsmServerCa() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmServerCaRead,

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
			"server_ca_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The identifier of the Server CA to operate on. Matches the `id` field in the `ServerCa` response schema.",
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
			"certificate": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PEM-encoded Server CA certificate. Immutable.",
			},
			"expires_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiry parsed from certificate NotAfter field. System-managed. Populated on creation.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Only one Server CA can be active per path. Starts inactive on creation.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AS OF health state of the Server CA certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation. Health check not yet performed.- `evaluating`: A health check is in progress for this Server CA.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.",
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

func dataSourceIbmSmHsmServerCaRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_server_ca", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions{}

	getInstanceHsmDestinationPathServerCaOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceHsmDestinationPathServerCaOptions.SetDestinationID(d.Get("destination_id").(string))
	getInstanceHsmDestinationPathServerCaOptions.SetPathID(d.Get("path_id").(string))
	getInstanceHsmDestinationPathServerCaOptions.SetServerCaID(d.Get("server_ca_id").(string))

	serverCa, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathServerCaWithContext(context, getInstanceHsmDestinationPathServerCaOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_server_ca", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", *getInstanceHsmDestinationPathServerCaOptions.InstanceID, *getInstanceHsmDestinationPathServerCaOptions.DestinationID, *getInstanceHsmDestinationPathServerCaOptions.PathID, *getInstanceHsmDestinationPathServerCaOptions.ServerCaID))

	if err = d.Set("name", serverCa.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting name: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-name").GetDiag()
	}

	if !core.IsNil(serverCa.Description) {
		if err = d.Set("description", serverCa.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting description: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(serverCa.Certificate) {
		if err = d.Set("certificate", serverCa.Certificate); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting certificate: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-certificate").GetDiag()
		}
	}

	if err = d.Set("expires_at", flex.DateTimeToString(serverCa.ExpiresAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting expires_at: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-expires_at").GetDiag()
	}

	if err = d.Set("status", serverCa.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting status: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-status").GetDiag()
	}

	if err = d.Set("health_state", serverCa.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_state: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-health_state").GetDiag()
	}

	if !core.IsNil(serverCa.HealthAsOf) {
		if err = d.Set("health_as_of", flex.DateTimeToString(serverCa.HealthAsOf)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_as_of: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-health_as_of").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(serverCa.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_at: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(serverCa.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(serverCa.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_at: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-updated_at").GetDiag()
		}
	}

	if err = d.Set("created_by", serverCa.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_by: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(serverCa.UpdatedBy) {
		if err = d.Set("updated_by", serverCa.UpdatedBy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_by: %s", err), "(Data) ibm_sm_hsm_server_ca", "read", "set-updated_by").GetDiag()
		}
	}

	return nil
}
