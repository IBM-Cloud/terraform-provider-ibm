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

func DataSourceIbmSmHsmClientCertificates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmClientCertificatesRead,

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
			"client_certificates": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of client certificates for this destination.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate ID. Immutable.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the client certificate resource.",
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
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Request-only. Set to `active` in a PATCH to activate this certificate. Never returned; read the current state from `provisioning_state`.",
						},
						"provisioning_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Lifecycle state of the client certificate:- `csr_generated`: CSR and private key generated successfully. `csr` and `csr_cn` are populated. Initial state on POST. Waiting for signed_cert to be uploaded via PATCH.- `inactive`: signed_cert has been uploaded but the certificate has not been activated yet.- `active`: This certificate is the active client certificate for the destination. Only one can be active at a time.",
						},
						"health_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AS OF health state of the client certificate. Written exclusively by POST …/health:- `unknown`: Initial state on creation, and for csr_generated certs with no signed_cert.- `evaluating`: A health check is in progress for this client certificate.- `valid`: Certificate is within its validity period (expires_at is in the future).- `expired`: Certificate has passed its expires_at date.",
						},
						"health_as_of": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.",
						},
						"csr": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PEM-encoded CSR. Platform-generated. Immutable once set. Populated on creation (always present once the resource exists).",
						},
						"csr_cn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Common Name embedded in the CSR. Set to the SM instance UUID at generation time. Immutable once set. Populated on creation (always present once the resource exists).",
						},
						"signed_cert": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CA-signed certificate (PEM). Omitted until uploaded via PATCH.",
						},
						"expires_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parsed from signed_cert NotAfter field. Omitted until signed_cert is uploaded.",
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

func dataSourceIbmSmHsmClientCertificatesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_client_certificates", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "(Data) ibm_sm_hsm_client_certificates", "read")
		return tfErr.GetDiag()
	}

	listInstanceHsmDestinationClientCertificatesOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationClientCertificatesOptions{}

	listInstanceHsmDestinationClientCertificatesOptions.SetInstanceID(d.Get("instance_id").(string))
	listInstanceHsmDestinationClientCertificatesOptions.SetDestinationID(d.Get("destination_id").(string))

	clientCertificateCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceHsmDestinationClientCertificatesWithContext(context, listInstanceHsmDestinationClientCertificatesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceHsmDestinationClientCertificatesWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_client_certificates", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmHsmClientCertificatesID(d))

	clientCertificates := []map[string]interface{}{}
	for _, clientCertificatesItem := range clientCertificateCollection.ClientCertificates {
		clientCertificatesItemMap, err := DataSourceIbmSmHsmClientCertificatesClientCertificateToMap(&clientCertificatesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_client_certificates", "read", "client_certificates-to-map").GetDiag()
		}
		clientCertificates = append(clientCertificates, clientCertificatesItemMap)
	}
	if err = d.Set("client_certificates", clientCertificates); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting client_certificates: %s", err), "(Data) ibm_sm_hsm_client_certificates", "read", "set-client_certificates").GetDiag()
	}

	return nil
}

// dataSourceIbmSmHsmClientCertificatesID returns a reasonable ID for the list.
func dataSourceIbmSmHsmClientCertificatesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmHsmClientCertificatesClientCertificateToMap(model *secretsmanagerinstancemanagementv2.ClientCertificate) (map[string]interface{}, error) {
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
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	modelMap["provisioning_state"] = *model.ProvisioningState
	modelMap["health_state"] = *model.HealthState
	modelMap["health_as_of"] = model.HealthAsOf.String()
	modelMap["csr"] = *model.Csr
	modelMap["csr_cn"] = *model.CsrCn
	if model.SignedCert != nil {
		modelMap["signed_cert"] = *model.SignedCert
	}
	if model.ExpiresAt != nil {
		modelMap["expires_at"] = model.ExpiresAt.String()
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
