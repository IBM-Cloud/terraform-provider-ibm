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

func DataSourceIbmSmHsmClientCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmHsmClientCertificateRead,

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
			"cert_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The identifier of the Client Certificate to operate on. Matches the `id` field in the `ClientCertificate` response schema.",
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
	}
}

func dataSourceIbmSmHsmClientCertificateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_hsm_client_certificate", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions{}

	getInstanceHsmDestinationClientCertificateOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceHsmDestinationClientCertificateOptions.SetDestinationID(d.Get("destination_id").(string))
	getInstanceHsmDestinationClientCertificateOptions.SetCertID(d.Get("cert_id").(string))

	clientCertificate, _, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationClientCertificateWithContext(context, getInstanceHsmDestinationClientCertificateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "(Data) ibm_sm_hsm_client_certificate", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *getInstanceHsmDestinationClientCertificateOptions.InstanceID, *getInstanceHsmDestinationClientCertificateOptions.DestinationID, *getInstanceHsmDestinationClientCertificateOptions.CertID))

	if !core.IsNil(clientCertificate.Href) {
		if err = d.Set("href", clientCertificate.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting href: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-href").GetDiag()
		}
	}

	if err = d.Set("name", clientCertificate.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting name: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-name").GetDiag()
	}

	if !core.IsNil(clientCertificate.Description) {
		if err = d.Set("description", clientCertificate.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting description: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(clientCertificate.Status) {
		if err = d.Set("status", clientCertificate.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting status: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-status").GetDiag()
		}
	}

	if err = d.Set("provisioning_state", clientCertificate.ProvisioningState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting provisioning_state: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-provisioning_state").GetDiag()
	}

	if err = d.Set("health_state", clientCertificate.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_state: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-health_state").GetDiag()
	}

	if err = d.Set("health_as_of", flex.DateTimeToString(clientCertificate.HealthAsOf)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting health_as_of: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-health_as_of").GetDiag()
	}

	if err = d.Set("csr", clientCertificate.Csr); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting csr: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-csr").GetDiag()
	}

	if err = d.Set("csr_cn", clientCertificate.CsrCn); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting csr_cn: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-csr_cn").GetDiag()
	}

	if !core.IsNil(clientCertificate.SignedCert) {
		if err = d.Set("signed_cert", clientCertificate.SignedCert); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting signed_cert: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-signed_cert").GetDiag()
		}
	}

	if !core.IsNil(clientCertificate.ExpiresAt) {
		if err = d.Set("expires_at", flex.DateTimeToString(clientCertificate.ExpiresAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting expires_at: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-expires_at").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(clientCertificate.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_at: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(clientCertificate.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(clientCertificate.UpdatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_at: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-updated_at").GetDiag()
		}
	}

	if err = d.Set("created_by", clientCertificate.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting created_by: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(clientCertificate.UpdatedBy) {
		if err = d.Set("updated_by", clientCertificate.UpdatedBy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("error setting updated_by: %s", err), "(Data) ibm_sm_hsm_client_certificate", "read", "set-updated_by").GetDiag()
		}
	}

	return nil
}
