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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmServerCa() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmServerCaCreate,
		ReadContext:   resourceIbmSmHsmServerCaRead,
		UpdateContext: resourceIbmSmHsmServerCaUpdate,
		DeleteContext: resourceIbmSmHsmServerCaDelete,
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
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "destination_id"),
				Description:  "Destination ID.",
			},
			"path_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "path_id"),
				Description:  "Path ID.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "name"),
				Description:  "Human-readable display name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "description"),
				Description:  "Optional description. Omitted if not set.",
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "certificate"),
				Description:  "PEM-encoded Server CA certificate. Immutable.",
			},
			"expires_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiry parsed from certificate NotAfter field. System-managed. Populated on creation.",
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_server_ca", "status"),
				Description:  "Only one Server CA can be active per path. Starts inactive on creation.",
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
			"server_ca_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server CA ID. Immutable.",
			},
		},
	}
}

func ResourceIbmSmHsmServerCaValidator() *validate.ResourceValidator {
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
			Identifier:                 "path_id",
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
		validate.ValidateSchema{
			Identifier:                 "certificate",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^-----BEGIN CERTIFICATE-----[\s\S]+-----END CERTIFICATE-----\s*$`,
			MinValueLength:             1,
			MaxValueLength:             65536,
		},
		validate.ValidateSchema{
			Identifier:                 "status",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(active|inactive)$`,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_server_ca", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmServerCaCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "ibm_sm_hsm_server_ca", "create")
		return tfErr.GetDiag()
	}

	createInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathServerCaOptions{}

	createInstanceHsmDestinationPathServerCaOptions.SetInstanceID(d.Get("instance_id").(string))
	createInstanceHsmDestinationPathServerCaOptions.SetDestinationID(d.Get("destination_id").(string))
	createInstanceHsmDestinationPathServerCaOptions.SetPathID(d.Get("path_id").(string))
	createInstanceHsmDestinationPathServerCaOptions.SetName(d.Get("name").(string))
	createInstanceHsmDestinationPathServerCaOptions.SetCertificate(d.Get("certificate").(string))
	if _, ok := d.GetOk("description"); ok {
		createInstanceHsmDestinationPathServerCaOptions.SetDescription(d.Get("description").(string))
	}

	serverCa, response, err := secretsManagerInstanceManagementClient.CreateInstanceHsmDestinationPathServerCaWithContext(context, createInstanceHsmDestinationPathServerCaOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// The service may omit the id from the body; fall back to the Location header (".../server_cas/{id}").
	serverCaID := ""
	if serverCa != nil && !core.IsNil(serverCa.ID) && *serverCa.ID != "" {
		serverCaID = *serverCa.ID
	}
	if serverCaID == "" && response != nil {
		serverCaID = lastPathSegment(response.GetHeaders().Get("Location"))
	}
	if serverCaID == "" {
		statusCode := 0
		if response != nil {
			statusCode = response.StatusCode
		}
		err = fmt.Errorf("create server CA returned no id in the body or the Location header (HTTP status %d)", statusCode)
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", *createInstanceHsmDestinationPathServerCaOptions.InstanceID, *createInstanceHsmDestinationPathServerCaOptions.DestinationID, *createInstanceHsmDestinationPathServerCaOptions.PathID, serverCaID))

	if statusVal, ok := d.GetOk("status"); ok && statusVal.(string) == "active" {
		updateInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions{}
		updateInstanceHsmDestinationPathServerCaOptions.SetInstanceID(*createInstanceHsmDestinationPathServerCaOptions.InstanceID)
		updateInstanceHsmDestinationPathServerCaOptions.SetDestinationID(*createInstanceHsmDestinationPathServerCaOptions.DestinationID)
		updateInstanceHsmDestinationPathServerCaOptions.SetPathID(*createInstanceHsmDestinationPathServerCaOptions.PathID)
		updateInstanceHsmDestinationPathServerCaOptions.SetServerCaID(serverCaID)
		activeStatus := secretsmanagerinstancemanagementv2.ServerCaPatch_Status_Active
		updateInstanceHsmDestinationPathServerCaOptions.ServerCaPatch = map[string]interface{}{
			"status": activeStatus,
		}
		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationPathServerCaWithContext(context, updateInstanceHsmDestinationPathServerCaOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmHsmServerCaRead(context, d, meta)
}

// lastPathSegment returns the last "/"-separated segment of a URL or path, ignoring query and trailing slashes.
func lastPathSegment(s string) string {
	if i := strings.IndexAny(s, "?#"); i >= 0 {
		s = s[:i]
	}
	s = strings.TrimRight(s, "/")
	return s[strings.LastIndex(s, "/")+1:]
}

func resourceIbmSmHsmServerCaRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationPathServerCaOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationPathServerCaOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationPathServerCaOptions.SetPathID(parts[2])
	getInstanceHsmDestinationPathServerCaOptions.SetServerCaID(parts[3])

	serverCa, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathServerCaWithContext(context, getInstanceHsmDestinationPathServerCaOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_id", parts[0]); err != nil {
		err = fmt.Errorf("error setting instance_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-instance_id").GetDiag()
	}
	if err = d.Set("destination_id", parts[1]); err != nil {
		err = fmt.Errorf("error setting destination_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-destination_id").GetDiag()
	}
	if err = d.Set("path_id", parts[2]); err != nil {
		err = fmt.Errorf("error setting path_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-path_id").GetDiag()
	}
	if err = d.Set("name", serverCa.Name); err != nil {
		err = fmt.Errorf("error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-name").GetDiag()
	}
	if !core.IsNil(serverCa.Description) {
		if err = d.Set("description", serverCa.Description); err != nil {
			err = fmt.Errorf("error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("certificate", serverCa.Certificate); err != nil {
		err = fmt.Errorf("error setting certificate: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-certificate").GetDiag()
	}
	if err = d.Set("expires_at", flex.DateTimeToString(serverCa.ExpiresAt)); err != nil {
		err = fmt.Errorf("error setting expires_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-expires_at").GetDiag()
	}
	if err = d.Set("status", serverCa.Status); err != nil {
		err = fmt.Errorf("error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-status").GetDiag()
	}
	if err = d.Set("health_state", serverCa.HealthState); err != nil {
		err = fmt.Errorf("error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-health_state").GetDiag()
	}
	if !core.IsNil(serverCa.HealthAsOf) {
		if err = d.Set("health_as_of", flex.DateTimeToString(serverCa.HealthAsOf)); err != nil {
			err = fmt.Errorf("error setting health_as_of: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-health_as_of").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(serverCa.CreatedAt)); err != nil {
		err = fmt.Errorf("error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(serverCa.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(serverCa.UpdatedAt)); err != nil {
			err = fmt.Errorf("error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-updated_at").GetDiag()
		}
	}
	if err = d.Set("created_by", serverCa.CreatedBy); err != nil {
		err = fmt.Errorf("error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(serverCa.UpdatedBy) {
		if err = d.Set("updated_by", serverCa.UpdatedBy); err != nil {
			err = fmt.Errorf("error setting updated_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-updated_by").GetDiag()
		}
	}
	// The service may omit the id from the body; fall back to the last part of the resource ID.
	serverCaID := parts[3]
	if !core.IsNil(serverCa.ID) && *serverCa.ID != "" {
		serverCaID = *serverCa.ID
	}
	if err = d.Set("server_ca_id", serverCaID); err != nil {
		err = fmt.Errorf("error setting server_ca_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "read", "set-server_ca_id").GetDiag()
	}

	return nil
}

func resourceIbmSmHsmServerCaUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceHsmDestinationPathServerCaOptions.SetInstanceID(parts[0])
	updateInstanceHsmDestinationPathServerCaOptions.SetDestinationID(parts[1])
	updateInstanceHsmDestinationPathServerCaOptions.SetPathID(parts[2])
	updateInstanceHsmDestinationPathServerCaOptions.SetServerCaID(parts[3])

	hasChange := false

	patchVals := &secretsmanagerinstancemanagementv2.ServerCaPatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_server_ca", "update", "instance_id-forces-new").GetDiag()
	}
	if d.HasChange("destination_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "destination_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_server_ca", "update", "destination_id-forces-new").GetDiag()
	}
	if d.HasChange("path_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "path_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_server_ca", "update", "path_id-forces-new").GetDiag()
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

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateInstanceHsmDestinationPathServerCaOptions.ServerCaPatch = ResourceIbmSmHsmServerCaServerCaPatchAsPatch(patchVals, d)

		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationPathServerCaWithContext(context, updateInstanceHsmDestinationPathServerCaOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmHsmServerCaRead(context, d, meta)
}

func resourceIbmSmHsmServerCaDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathServerCaOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_server_ca", "delete", "sep-id-parts").GetDiag()
	}

	deleteInstanceHsmDestinationPathServerCaOptions.SetInstanceID(parts[0])
	deleteInstanceHsmDestinationPathServerCaOptions.SetDestinationID(parts[1])
	deleteInstanceHsmDestinationPathServerCaOptions.SetPathID(parts[2])
	deleteInstanceHsmDestinationPathServerCaOptions.SetServerCaID(parts[3])

	// The service refuses to delete an active Server CA and does not accept status "inactive".
	// If this CA is active, activate another Server CA on the path first (which demotes this one).
	getInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions{}
	getInstanceHsmDestinationPathServerCaOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationPathServerCaOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationPathServerCaOptions.SetPathID(parts[2])
	getInstanceHsmDestinationPathServerCaOptions.SetServerCaID(parts[3])
	serverCa, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathServerCaWithContext(context, getInstanceHsmDestinationPathServerCaOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(serverCa.Status) && *serverCa.Status == secretsmanagerinstancemanagementv2.ServerCa_Status_Active {
		listInstanceHsmDestinationPathServerCasOptions := &secretsmanagerinstancemanagementv2.ListInstanceHsmDestinationPathServerCasOptions{}
		listInstanceHsmDestinationPathServerCasOptions.SetInstanceID(parts[0])
		listInstanceHsmDestinationPathServerCasOptions.SetDestinationID(parts[1])
		listInstanceHsmDestinationPathServerCasOptions.SetPathID(parts[2])
		serverCaCollection, _, err := secretsManagerInstanceManagementClient.ListInstanceHsmDestinationPathServerCasWithContext(context, listInstanceHsmDestinationPathServerCasOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceHsmDestinationPathServerCasWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		otherID := ""
		for _, ca := range serverCaCollection.ServerCas {
			id := ""
			if !core.IsNil(ca.ID) && *ca.ID != "" {
				id = *ca.ID
			}
			if id != "" && id != parts[3] {
				otherID = id
				break
			}
		}

		if otherID == "" {
			// The only Server CA on the path cannot be deleted while active. It is removed together with its path or destination.
			d.SetId("")
			return diag.Diagnostics{{
				Severity: diag.Warning,
				Summary:  "Active Server CA was not deleted",
				Detail:   fmt.Sprintf("Server CA %s is the only Server CA on the path and is active, so the service cannot delete it. It was removed from the Terraform state and will be deleted with its path or destination.", parts[3]),
			}}
		}

		updateInstanceHsmDestinationPathServerCaOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathServerCaOptions{}
		updateInstanceHsmDestinationPathServerCaOptions.SetInstanceID(parts[0])
		updateInstanceHsmDestinationPathServerCaOptions.SetDestinationID(parts[1])
		updateInstanceHsmDestinationPathServerCaOptions.SetPathID(parts[2])
		updateInstanceHsmDestinationPathServerCaOptions.SetServerCaID(otherID)
		updateInstanceHsmDestinationPathServerCaOptions.ServerCaPatch = map[string]interface{}{
			"status": secretsmanagerinstancemanagementv2.ServerCaPatch_Status_Active,
		}
		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationPathServerCaWithContext(context, updateInstanceHsmDestinationPathServerCaOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	_, err = secretsManagerInstanceManagementClient.DeleteInstanceHsmDestinationPathServerCaWithContext(context, deleteInstanceHsmDestinationPathServerCaOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceHsmDestinationPathServerCaWithContext failed: %s", err.Error()), "ibm_sm_hsm_server_ca", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmSmHsmServerCaServerCaPatchAsPatch(patchVals *secretsmanagerinstancemanagementv2.ServerCaPatch, d *schema.ResourceData) map[string]interface{} {
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

	return patch
}
