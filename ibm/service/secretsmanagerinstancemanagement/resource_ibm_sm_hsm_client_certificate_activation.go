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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmClientCertificateActivation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmClientCertificateActivationCreate,
		ReadContext:   resourceIbmSmHsmClientCertificateActivationRead,
		UpdateContext: resourceIbmSmHsmClientCertificateActivationUpdate,
		DeleteContext: resourceIbmSmHsmClientCertificateActivationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The region of the Secrets Manager instance.",
			},
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate_activation", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate_activation", "destination_id"),
				Description:  "Destination ID.",
			},
			"cert_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate_activation", "cert_id"),
				Description:  "Client certificate ID.",
			},
			"signed_cert": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Signed client certificate PEM.",
			},
			"provisioning_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Lifecycle state of the activated client certificate.",
			},
			"expires_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiry date of the activated client certificate.",
			},
		},
	}
}

func ResourceIbmSmHsmClientCertificateActivationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "destination_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "cert_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_client_certificate_activation", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmClientCertificateActivationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate_activation", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "ibm_sm_hsm_client_certificate_activation", "create")
		return tfErr.GetDiag()
	}

	updateOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions{}
	updateOptions.SetInstanceID(d.Get("instance_id").(string))
	updateOptions.SetDestinationID(d.Get("destination_id").(string))
	updateOptions.SetCertID(d.Get("cert_id").(string))

	signedCert := d.Get("signed_cert").(string)
	status := secretsmanagerinstancemanagementv2.ClientCertificate_ProvisioningState_Active
	patchVals := map[string]interface{}{
		"signed_cert": signedCert,
		"status":      status,
	}
	updateOptions.SetClientCertificatePatch(patchVals)

	_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationClientCertificateWithContext(context, updateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "ibm_sm_hsm_client_certificate_activation", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *updateOptions.InstanceID, *updateOptions.DestinationID, *updateOptions.CertID))

	return resourceIbmSmHsmClientCertificateActivationRead(context, d, meta)
}

func resourceIbmSmHsmClientCertificateActivationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate_activation", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate_activation", "read", "sep-id-parts").GetDiag()
	}

	getOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions{}
	getOptions.SetInstanceID(parts[0])
	getOptions.SetDestinationID(parts[1])
	getOptions.SetCertID(parts[2])

	cert, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationClientCertificateWithContext(context, getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "ibm_sm_hsm_client_certificate_activation", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if cert.ProvisioningState != nil {
		if err = d.Set("provisioning_state", *cert.ProvisioningState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate_activation", "read", "set-provisioning_state").GetDiag()
		}
	}
	if cert.ExpiresAt != nil {
		if err = d.Set("expires_at", cert.ExpiresAt.String()); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate_activation", "read", "set-expires_at").GetDiag()
		}
	}
	return nil
}

func resourceIbmSmHsmClientCertificateActivationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceIbmSmHsmClientCertificateActivationCreate(context, d, meta)
}

func resourceIbmSmHsmClientCertificateActivationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not delete the certificate; it only tracks activation.
	d.SetId("")
	return nil
}
