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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmClientCertificate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmClientCertificateCreate,
		ReadContext:   resourceIbmSmHsmClientCertificateRead,
		UpdateContext: resourceIbmSmHsmClientCertificateUpdate,
		DeleteContext: resourceIbmSmHsmClientCertificateDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate", "destination_id"),
				Description:  "Destination ID.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate", "name"),
				Description:  "Human-readable display name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_client_certificate", "description"),
				Description:  "Optional description. Omitted if not set.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the client certificate resource.",
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
			"cert_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate ID. Immutable.",
			},
		},
	}
}

func ResourceIbmSmHsmClientCertificateValidator() *validate.ResourceValidator {
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
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$`,
			MinValueLength:             2,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(.*?)$`,
			MinValueLength:             0,
			MaxValueLength:             1024,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_client_certificate", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmClientCertificateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationClientCertificateOptions{}

	createInstanceHsmDestinationClientCertificateOptions.SetInstanceID(d.Get("instance_id").(string))
	createInstanceHsmDestinationClientCertificateOptions.SetDestinationID(d.Get("destination_id").(string))
	createInstanceHsmDestinationClientCertificateOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		createInstanceHsmDestinationClientCertificateOptions.SetDescription(d.Get("description").(string))
	}

	clientCertificate, _, err := secretsManagerInstanceManagementClient.CreateInstanceHsmDestinationClientCertificateWithContext(context, createInstanceHsmDestinationClientCertificateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "ibm_sm_hsm_client_certificate", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createInstanceHsmDestinationClientCertificateOptions.InstanceID, *createInstanceHsmDestinationClientCertificateOptions.DestinationID, *clientCertificate.ID))

	return resourceIbmSmHsmClientCertificateRead(context, d, meta)
}

func resourceIbmSmHsmClientCertificateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationClientCertificateOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationClientCertificateOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationClientCertificateOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationClientCertificateOptions.SetCertID(parts[2])

	clientCertificate, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationClientCertificateWithContext(context, getInstanceHsmDestinationClientCertificateOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "ibm_sm_hsm_client_certificate", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_id", parts[0]); err != nil {
		err = fmt.Errorf("error setting instance_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-instance_id").GetDiag()
	}
	if err = d.Set("destination_id", parts[1]); err != nil {
		err = fmt.Errorf("error setting destination_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-destination_id").GetDiag()
	}
	if err = d.Set("name", clientCertificate.Name); err != nil {
		err = fmt.Errorf("error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-name").GetDiag()
	}
	if !core.IsNil(clientCertificate.Description) {
		if err = d.Set("description", clientCertificate.Description); err != nil {
			err = fmt.Errorf("error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(clientCertificate.Href) {
		if err = d.Set("href", clientCertificate.Href); err != nil {
			err = fmt.Errorf("error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-href").GetDiag()
		}
	}
	if !core.IsNil(clientCertificate.Status) {
		if err = d.Set("status", clientCertificate.Status); err != nil {
			err = fmt.Errorf("error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-status").GetDiag()
		}
	}
	if err = d.Set("provisioning_state", clientCertificate.ProvisioningState); err != nil {
		err = fmt.Errorf("error setting provisioning_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-provisioning_state").GetDiag()
	}
	if err = d.Set("health_state", clientCertificate.HealthState); err != nil {
		err = fmt.Errorf("error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("health_as_of", flex.DateTimeToString(clientCertificate.HealthAsOf)); err != nil {
		err = fmt.Errorf("error setting health_as_of: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-health_as_of").GetDiag()
	}
	if err = d.Set("csr", clientCertificate.Csr); err != nil {
		err = fmt.Errorf("error setting csr: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-csr").GetDiag()
	}
	if err = d.Set("csr_cn", clientCertificate.CsrCn); err != nil {
		err = fmt.Errorf("error setting csr_cn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-csr_cn").GetDiag()
	}
	if !core.IsNil(clientCertificate.SignedCert) {
		if err = d.Set("signed_cert", clientCertificate.SignedCert); err != nil {
			err = fmt.Errorf("error setting signed_cert: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-signed_cert").GetDiag()
		}
	}
	if !core.IsNil(clientCertificate.ExpiresAt) {
		if err = d.Set("expires_at", flex.DateTimeToString(clientCertificate.ExpiresAt)); err != nil {
			err = fmt.Errorf("error setting expires_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-expires_at").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(clientCertificate.CreatedAt)); err != nil {
		err = fmt.Errorf("error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(clientCertificate.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(clientCertificate.UpdatedAt)); err != nil {
			err = fmt.Errorf("error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-updated_at").GetDiag()
		}
	}
	if err = d.Set("created_by", clientCertificate.CreatedBy); err != nil {
		err = fmt.Errorf("error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(clientCertificate.UpdatedBy) {
		if err = d.Set("updated_by", clientCertificate.UpdatedBy); err != nil {
			err = fmt.Errorf("error setting updated_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-updated_by").GetDiag()
		}
	}
	if !core.IsNil(clientCertificate.ID) {
		if err = d.Set("cert_id", clientCertificate.ID); err != nil {
			err = fmt.Errorf("error setting cert_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "read", "set-cert_id").GetDiag()
		}
	}

	return nil
}

func resourceIbmSmHsmClientCertificateUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationClientCertificateOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceHsmDestinationClientCertificateOptions.SetInstanceID(parts[0])
	updateInstanceHsmDestinationClientCertificateOptions.SetDestinationID(parts[1])
	updateInstanceHsmDestinationClientCertificateOptions.SetCertID(parts[2])

	hasChange := false

	patchVals := &secretsmanagerinstancemanagementv2.ClientCertificatePatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_client_certificate", "update", "instance_id-forces-new").GetDiag()
	}
	if d.HasChange("destination_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "destination_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_client_certificate", "update", "destination_id-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("description") {
		newDescription := d.Get("description").(string)
		patchVals.Description = &newDescription
		hasChange = true
	}
	if d.HasChange("status") {
		newStatus := d.Get("status").(string)
		patchVals.Status = &newStatus
		hasChange = true
	}
	if d.HasChange("signed_cert") {
		newSignedCert := d.Get("signed_cert").(string)
		patchVals.SignedCert = &newSignedCert
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateInstanceHsmDestinationClientCertificateOptions.ClientCertificatePatch = ResourceIbmSmHsmClientCertificateClientCertificatePatchAsPatch(patchVals, d)

		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationClientCertificateWithContext(context, updateInstanceHsmDestinationClientCertificateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "ibm_sm_hsm_client_certificate", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmHsmClientCertificateRead(context, d, meta)
}

func resourceIbmSmHsmClientCertificateDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceHsmDestinationClientCertificateOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationClientCertificateOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_client_certificate", "delete", "sep-id-parts").GetDiag()
	}

	deleteInstanceHsmDestinationClientCertificateOptions.SetInstanceID(parts[0])
	deleteInstanceHsmDestinationClientCertificateOptions.SetDestinationID(parts[1])
	deleteInstanceHsmDestinationClientCertificateOptions.SetCertID(parts[2])

	_, err = secretsManagerInstanceManagementClient.DeleteInstanceHsmDestinationClientCertificateWithContext(context, deleteInstanceHsmDestinationClientCertificateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceHsmDestinationClientCertificateWithContext failed: %s", err.Error()), "ibm_sm_hsm_client_certificate", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmSmHsmClientCertificateClientCertificatePatchAsPatch(patchVals *secretsmanagerinstancemanagementv2.ClientCertificatePatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}
	path = "description"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["description"] = nil
	} else if !exists {
		delete(patch, "description")
	}
	path = "status"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["status"] = nil
	} else if !exists {
		delete(patch, "status")
	}
	path = "signed_cert"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["signed_cert"] = nil
	} else if !exists {
		delete(patch, "signed_cert")
	}

	return patch
}
