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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmDestinationHealth() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmDestinationHealthRead,

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
				Description: "URL of the health snapshot endpoint for this destination.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.",
			},
			"health_as_of": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of this health check — always the current time of the live probe.",
			},
			"client_certificates": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "All client certificates for this destination (active, inactive, csr_generated). health_state freshly evaluated for all certs with a signed_cert.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Client certificate ID.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AS OF health state of the client certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.- `evaluating`: A health check is in progress for this client certificate.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.",
						},
					},
				},
			},
			"ha_groups": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "All HSM HA Groups for this destination. Empty array if none exist.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HA Group ID.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate).",
						},
					},
				},
			},
			"paths": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "All paths for this destination. Empty array if none exist.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Path ID.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable display name.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. Health check not yet performed, or path is deleting/deletion_failed.- `evaluating`: A health check is in progress for this path.- `in_progress`: An apply operation is running across the destination.- `no_valid_server_ca`: Pre-check failed — no Server CA is active, or the active Server CA has expired.- `not_configured`: Path is not yet provisioned (pending_account_policy, provisioning, pending_approval, approval_declined, crn_already_in_use, vpe_provisioning_failed) — AND active Server CA is valid. If active CA is missing or expired, no_valid_server_ca takes priority.- `healthy`: Active Server CA valid AND connectivity check to the HSM succeeded.- `unhealthy`: Active Server CA valid AND connectivity check to the HSM failed.",
						},
						"server_cas": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "All Server CAs for this path (active and inactive). Empty array if none.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Server CA ID.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Human-readable display name.",
									},
									"health_state": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "AS OF health state of the Server CA certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation. Health check not yet performed.- `evaluating`: A health check is in progress for this Server CA.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmHsmDestinationHealthRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destination_health", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getHsmDestinationHealthSnapshotOptions := &secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions{}

	getHsmDestinationHealthSnapshotOptions.SetInstanceID(d.Get("instance_id").(string))
	getHsmDestinationHealthSnapshotOptions.SetDestinationID(d.Get("destination_id").(string))

	destinationHealthSnapshot, _, err := secretsManagerInstanceManagementClient.GetHsmDestinationHealthSnapshotWithContext(context, getHsmDestinationHealthSnapshotOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetHsmDestinationHealthSnapshotWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_destination_health", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmHsmDestinationHealthID(d))

	if !core.IsNil(destinationHealthSnapshot.Href) {
		if err = d.Set("href", destinationHealthSnapshot.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting href: %s", err), "(Data) ibm_sm_hsm_destination_health", "read", "set-href").GetDiag()
		}
	}

	if err = d.Set("health_state", destinationHealthSnapshot.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_state: %s", err), "(Data) ibm_sm_hsm_destination_health", "read", "set-health_state").GetDiag()
	}

	if err = d.Set("health_as_of", flex.DateTimeToString(destinationHealthSnapshot.HealthAsOf)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_as_of: %s", err), "(Data) ibm_sm_hsm_destination_health", "read", "set-health_as_of").GetDiag()
	}

	clientCertificates := []map[string]interface{}{}
	for _, clientCertificatesItem := range destinationHealthSnapshot.ClientCertificates {
		clientCertificatesItemMap, err := DataSourceIbmSmHsmDestinationHealthHealthSnapshotClientCertToMap(&clientCertificatesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destination_health", "read", "client_certificates-to-map").GetDiag()
		}
		clientCertificates = append(clientCertificates, clientCertificatesItemMap)
	}
	if err = d.Set("client_certificates", clientCertificates); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting client_certificates: %s", err), "(Data) ibm_sm_hsm_destination_health", "read", "set-client_certificates").GetDiag()
	}

	haGroups := []map[string]interface{}{}
	for _, haGroupsItem := range destinationHealthSnapshot.HaGroups {
		haGroupsItemMap, err := DataSourceIbmSmHsmDestinationHealthHealthSnapshotHSMHAGroupToMap(&haGroupsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destination_health", "read", "ha_groups-to-map").GetDiag()
		}
		haGroups = append(haGroups, haGroupsItemMap)
	}
	if err = d.Set("ha_groups", haGroups); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting ha_groups: %s", err), "(Data) ibm_sm_hsm_destination_health", "read", "set-ha_groups").GetDiag()
	}

	paths := []map[string]interface{}{}
	for _, pathsItem := range destinationHealthSnapshot.Paths {
		pathsItemMap, err := DataSourceIbmSmHsmDestinationHealthHealthSnapshotConnectionToMap(&pathsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_destination_health", "read", "paths-to-map").GetDiag()
		}
		paths = append(paths, pathsItemMap)
	}
	if err = d.Set("paths", paths); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting paths: %s", err), "(Data) ibm_sm_hsm_destination_health", "read", "set-paths").GetDiag()
	}

	return nil
}

// dataSourceIbmSmHsmDestinationHealthID returns a reasonable ID for the list.
func dataSourceIbmSmHsmDestinationHealthID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmHsmDestinationHealthHealthSnapshotClientCertToMap(model *secretsmanagerinstancemanagementv2.HealthSnapshotClientCert) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	modelMap["health_state"] = *model.HealthState
	return modelMap, nil
}

func DataSourceIbmSmHsmDestinationHealthHealthSnapshotHSMHAGroupToMap(model *secretsmanagerinstancemanagementv2.HealthSnapshotHSMHAGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	modelMap["health_state"] = *model.HealthState
	return modelMap, nil
}

func DataSourceIbmSmHsmDestinationHealthHealthSnapshotConnectionToMap(model *secretsmanagerinstancemanagementv2.HealthSnapshotConnection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	modelMap["health_state"] = *model.HealthState
	serverCas := []map[string]interface{}{}
	for _, serverCasItem := range model.ServerCas {
		serverCasItemMap, err := DataSourceIbmSmHsmDestinationHealthHealthSnapshotServerCaToMap(&serverCasItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		serverCas = append(serverCas, serverCasItemMap)
	}
	modelMap["server_cas"] = serverCas
	return modelMap, nil
}

func DataSourceIbmSmHsmDestinationHealthHealthSnapshotServerCaToMap(model *secretsmanagerinstancemanagementv2.HealthSnapshotServerCa) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	modelMap["health_state"] = *model.HealthState
	return modelMap, nil
}
