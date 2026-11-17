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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmPath() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmPathCreate,
		ReadContext:   resourceIbmSmHsmPathRead,
		UpdateContext: resourceIbmSmHsmPathUpdate,
		DeleteContext: resourceIbmSmHsmPathDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path", "destination_id"),
				Description:  "Destination ID.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path", "name"),
				Description:  "Human-readable display name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path", "description"),
				Description:  "Optional description. Omitted if not set.",
			},
			"crn": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_path", "crn"),
				Description:  "CRN of the gateway through which the HSM is reached. Immutable.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the path resource.",
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
			"path_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path ID. Immutable. System-managed.",
			},
		},
	}
}

func ResourceIbmSmHsmPathValidator() *validate.ResourceValidator {
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
		validate.ValidateSchema{
			Identifier:                 "crn",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$`,
			MinValueLength:             9,
			MaxValueLength:             512,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_path", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmPathCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationPathOptions{}

	createInstanceHsmDestinationPathOptions.SetInstanceID(d.Get("instance_id").(string))
	createInstanceHsmDestinationPathOptions.SetDestinationID(d.Get("destination_id").(string))
	createInstanceHsmDestinationPathOptions.SetName(d.Get("name").(string))
	createInstanceHsmDestinationPathOptions.SetCrn(d.Get("crn").(string))
	if _, ok := d.GetOk("description"); ok {
		createInstanceHsmDestinationPathOptions.SetDescription(d.Get("description").(string))
	}

	path, _, err := secretsManagerInstanceManagementClient.CreateInstanceHsmDestinationPathWithContext(context, createInstanceHsmDestinationPathOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "ibm_sm_hsm_path", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createInstanceHsmDestinationPathOptions.InstanceID, *createInstanceHsmDestinationPathOptions.DestinationID, *path.ID))

	return resourceIbmSmHsmPathRead(context, d, meta)
}

func resourceIbmSmHsmPathRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationPathOptions.SetPathID(parts[2])

	path, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPathWithContext(context, getInstanceHsmDestinationPathOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "ibm_sm_hsm_path", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_id", parts[0]); err != nil {
		err = fmt.Errorf("error setting instance_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-instance_id").GetDiag()
	}
	if err = d.Set("destination_id", parts[1]); err != nil {
		err = fmt.Errorf("error setting destination_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-destination_id").GetDiag()
	}
	if err = d.Set("name", path.Name); err != nil {
		err = fmt.Errorf("error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-name").GetDiag()
	}
	if !core.IsNil(path.Description) {
		if err = d.Set("description", path.Description); err != nil {
			err = fmt.Errorf("error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("crn", path.Crn); err != nil {
		err = fmt.Errorf("error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(path.Href) {
		if err = d.Set("href", path.Href); err != nil {
			err = fmt.Errorf("error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-href").GetDiag()
		}
	}
	if err = d.Set("provisioning_state", path.ProvisioningState); err != nil {
		err = fmt.Errorf("error setting provisioning_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-provisioning_state").GetDiag()
	}
	if err = d.Set("health_state", path.HealthState); err != nil {
		err = fmt.Errorf("error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("health_as_of", flex.DateTimeToString(path.HealthAsOf)); err != nil {
		err = fmt.Errorf("error setting health_as_of: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-health_as_of").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(path.CreatedAt)); err != nil {
		err = fmt.Errorf("error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(path.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(path.UpdatedAt)); err != nil {
			err = fmt.Errorf("error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-updated_at").GetDiag()
		}
	}
	if err = d.Set("created_by", path.CreatedBy); err != nil {
		err = fmt.Errorf("error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(path.UpdatedBy) {
		if err = d.Set("updated_by", path.UpdatedBy); err != nil {
			err = fmt.Errorf("error setting updated_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-updated_by").GetDiag()
		}
	}
	if !core.IsNil(path.ID) {
		if err = d.Set("path_id", path.ID); err != nil {
			err = fmt.Errorf("error setting path_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "read", "set-path_id").GetDiag()
		}
	}

	return nil
}

func resourceIbmSmHsmPathUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationPathOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
	updateInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
	updateInstanceHsmDestinationPathOptions.SetPathID(parts[2])

	hasChange := false

	patchVals := &secretsmanagerinstancemanagementv2.PathPatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_path", "update", "instance_id-forces-new").GetDiag()
	}
	if d.HasChange("destination_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "destination_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_path", "update", "destination_id-forces-new").GetDiag()
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

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateInstanceHsmDestinationPathOptions.PathPatch = ResourceIbmSmHsmPathPathPatchAsPatch(patchVals, d)

		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationPathWithContext(context, updateInstanceHsmDestinationPathOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "ibm_sm_hsm_path", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmHsmPathRead(context, d, meta)
}

func resourceIbmSmHsmPathDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationPathOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_path", "delete", "sep-id-parts").GetDiag()
	}

	deleteInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
	deleteInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
	deleteInstanceHsmDestinationPathOptions.SetPathID(parts[2])

	_, _, err = secretsManagerInstanceManagementClient.DeleteInstanceHsmDestinationPathWithContext(context, deleteInstanceHsmDestinationPathOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceHsmDestinationPathWithContext failed: %s", err.Error()), "ibm_sm_hsm_path", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = waitForIbmSmHsmPathDelete(d, meta)
	if err != nil {
		errMsg := fmt.Sprintf("error waiting for resource IbmSmHsmPath (%s) to be deleted: %s", d.Id(), err)
		return flex.DiscriminatedTerraformErrorf(err, errMsg, "ibm_sm_hsm_path", "delete", "wait-for-state").GetDiag()
	}

	d.SetId("")

	return nil
}

func waitForIbmSmHsmPathDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return false, err
	}
	getInstanceHsmDestinationPathOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getInstanceHsmDestinationPathOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationPathOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationPathOptions.SetPathID(parts[2])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"deleting"},
		Target:  []string{"success", "succeeded", "valid", "active", "done", "normal", "ready", "provisioned", "available", "deployed", "healthy", "pending_reclamation", "deleted", "removed", "stable", "suspended"},
		Refresh: func() (interface{}, string, error) {
			stateObj, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationPath(getInstanceHsmDestinationPathOptions)
			if err != nil {
				if _, ok := err.(*core.SDKProblem); ok && response != nil && response.GetStatusCode() == 404 {
					return response, "success", nil
				}
				return nil, "", err
			}
			failStates := map[string]bool{"deletion_failed": true}
			if failStates[*stateObj.ProvisioningState] {
				return stateObj, *stateObj.ProvisioningState, fmt.Errorf("the instance %s failed: %s", "getInstanceHsmDestinationPathOptions", err)
			}
			return stateObj, *stateObj.ProvisioningState, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func ResourceIbmSmHsmPathPathPatchAsPatch(patchVals *secretsmanagerinstancemanagementv2.PathPatch, d *schema.ResourceData) map[string]interface{} {
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

	return patch
}
