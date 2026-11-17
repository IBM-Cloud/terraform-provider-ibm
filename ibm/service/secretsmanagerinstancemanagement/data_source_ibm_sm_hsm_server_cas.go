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
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmHsmServerCas() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmServerCasRead,

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
			"server_cas": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of Server CA certificates for this path.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server CA ID. Immutable.",
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
				},
			},
		},
	}
}

func dataSourceIbmSmHsmServerCasRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_server_cas", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listInstanceHsmDestinationPathServerCasOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions{}

	listInstanceHsmDestinationPathServerCasOptions.SetInstanceID(d.Get("instance_id").(string))
	listInstanceHsmDestinationPathServerCasOptions.SetDestinationID(d.Get("destination_id").(string))
	listInstanceHsmDestinationPathServerCasOptions.SetPathID(d.Get("path_id").(string))

	serverCaCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceHsmDestinationPathServerCasWithContext(context, listInstanceHsmDestinationPathServerCasOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceHsmDestinationPathServerCasWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_server_cas", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmHsmServerCasID(d))

	serverCas := []map[string]interface{}{}
	for _, serverCasItem := range serverCaCollection.ServerCas {
		serverCasItemMap, err := DataSourceIbmSmHsmServerCasServerCaToMap(&serverCasItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_server_cas", "read", "server_cas-to-map").GetDiag()
		}
		serverCas = append(serverCas, serverCasItemMap)
	}
	if err = d.Set("server_cas", serverCas); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting server_cas: %s", err), "(Data) ibm_sm_hsm_server_cas", "read", "set-server_cas").GetDiag()
	}

	return nil
}

// dataSourceIbmSmHsmServerCasID returns a reasonable ID for the list.
func dataSourceIbmSmHsmServerCasID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmHsmServerCasServerCaToMap(model *secretsmanagerinstancemanagementv2.ServerCa) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil && *model.ID != "" {
		modelMap["id"] = *model.ID
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Certificate != nil {
		modelMap["certificate"] = *model.Certificate
	}
	modelMap["expires_at"] = model.ExpiresAt.String()
	modelMap["status"] = *model.Status
	modelMap["health_state"] = *model.HealthState
	if model.HealthAsOf != nil {
		modelMap["health_as_of"] = model.HealthAsOf.String()
	}
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
