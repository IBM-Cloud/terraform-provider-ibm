// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.2-f568b7d6-20261016-181021
*/

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmSmDestinationCreate,
		ReadContext:     resourceIbmSmDestinationRead,
		UpdateContext:   resourceIbmSmDestinationUpdate,
		DeleteContext:   resourceIbmSmDestinationDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_destination", "instance_id"),
				Description: "Secrets Manager instance ID.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_destination", "name"),
				Description: "Destination name.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_destination", "type"),
				Description: "Destination type.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_destination", "description"),
				Description: "Optional description.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_destination", "crn"),
				Description: "IBM Cloud Database service instance CRN.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the destination resource.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the destination was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the destination was last updated.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM identity that created the destination.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination ID.",
			},
		},
	}
}

func ResourceIbmSmDestinationValidator() *validate.ResourceValidator {
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
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$`,
			MinValueLength:             2,
			MaxValueLength:             256,
		},
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "ibm_cloud_database",
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `(.*?)`,
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_destination", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmDestinationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.CreateInstanceDestinationOptions{}

	bodyModelMap["name"] = d.Get("name")
	bodyModelMap["type"] = d.Get("type")
	if _, ok := d.GetOk("description"); ok {
		bodyModelMap["description"] = d.Get("description")
	}
	bodyModelMap["crn"] = d.Get("crn")
	createInstanceDestinationOptions.SetInstanceID(d.Get("instance_id").(string))
	convertedModel, err := ResourceIbmSmDestinationMapToCreateDestinationRequest(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "create", "parse-request-body").GetDiag()
	}
	createInstanceDestinationOptions.DestinationPrototype = convertedModel

	ibmCloudDatabaseDestination, _, err := secretsManagerInstanceManagementClient.CreateInstanceDestinationWithContext(context, createInstanceDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceDestinationWithContext failed: %s", err.Error()), "ibm_sm_destination", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createInstanceDestinationOptions.InstanceID, string(*ibmCloudDatabaseDestination.ID)))

	_, err = waitForIbmSmDestinationCreate(d, meta)
	if err != nil {
		errMsg := fmt.Sprintf("error waiting for resource IbmSmDestination (%s) to be created: %s", d.Id(), err)
		return flex.DiscriminatedTerraformErrorf(err, errMsg, "ibm_sm_destination", "create", "wait-for-state").GetDiag()
	}

	return resourceIbmSmDestinationRead(context, d, meta)
}

func waitForIbmSmDestinationCreate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return false, err
	}
	getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getInstanceDestinationOptions.SetInstanceID(parts[0])
	getInstanceDestinationOptions.SetDestinationID(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"not_started", "provisioning"},
		Target:  []string{"succeeded"},
		Refresh: func() (interface{}, string, error) {
			stateObj, response, err := secretsManagerInstanceManagementClient.GetInstanceDestination(getInstanceDestinationOptions)
			if err != nil {
				if sdkErr, ok := err.(*core.SDKProblem); ok && response.GetStatusCode() == 404 {
					sdkErr.Summary = fmt.Sprintf("the instance %s does not exist anymore: %s", "getInstanceDestinationOptions", err)
					return nil, "", sdkErr
				}
				return nil, "", err
			}
			failStates := map[string]bool{"failed": true}
			if failStates[*stateObj.State] {
				return stateObj, *stateObj.State, fmt.Errorf("the instance %s failed: %s", "getInstanceDestinationOptions", err)
			}
			return stateObj, *stateObj.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func resourceIbmSmDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "sep-id-parts").GetDiag()
	}

	getInstanceDestinationOptions.SetInstanceID(parts[0])
	getInstanceDestinationOptions.SetDestinationID(parts[1])

	ibmCloudDatabaseDestination, response, err := secretsManagerInstanceManagementClient.GetInstanceDestinationWithContext(context, getInstanceDestinationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceDestinationWithContext failed: %s", err.Error()), "ibm_sm_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", ibmCloudDatabaseDestination.Name); err != nil {
		err = fmt.Errorf("error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-name").GetDiag()
	}
	if err = d.Set("type", ibmCloudDatabaseDestination.Type); err != nil {
		err = fmt.Errorf("error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-type").GetDiag()
	}
	if !core.IsNil(ibmCloudDatabaseDestination.Description) {
		if err = d.Set("description", ibmCloudDatabaseDestination.Description); err != nil {
			err = fmt.Errorf("error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("crn", ibmCloudDatabaseDestination.Crn); err != nil {
		err = fmt.Errorf("error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(ibmCloudDatabaseDestination.Href) {
		if err = d.Set("href", ibmCloudDatabaseDestination.Href); err != nil {
			err = fmt.Errorf("error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-href").GetDiag()
		}
	}
	if err = d.Set("state", ibmCloudDatabaseDestination.State); err != nil {
		err = fmt.Errorf("error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-state").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(ibmCloudDatabaseDestination.CreatedAt)); err != nil {
		err = fmt.Errorf("error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(ibmCloudDatabaseDestination.UpdatedAt)); err != nil {
		err = fmt.Errorf("error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-updated_at").GetDiag()
	}
	if !core.IsNil(ibmCloudDatabaseDestination.CreatedBy) {
		if err = d.Set("created_by", ibmCloudDatabaseDestination.CreatedBy); err != nil {
			err = fmt.Errorf("error setting created_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-created_by").GetDiag()
		}
	}
	if !core.IsNil(ibmCloudDatabaseDestination.Message) {
		if err = d.Set("message", ibmCloudDatabaseDestination.Message); err != nil {
			err = fmt.Errorf("error setting message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-message").GetDiag()
		}
	}
	if err = d.Set("destination_id", ibmCloudDatabaseDestination.ID); err != nil {
		err = fmt.Errorf("error setting destination_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "read", "set-destination_id").GetDiag()
	}

	return nil
}

func resourceIbmSmDestinationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceDestinationOptions.SetInstanceID(parts[0])
	updateInstanceDestinationOptions.SetDestinationID(parts[1])

	hasChange := false

	patchVals := &secretsmanagerinstancemanagementv2.DestinationPatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation." +
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_destination", "update", "instance_id-forces-new").GetDiag()
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
		updateInstanceDestinationOptions.DestinationPatch = ResourceIbmSmDestinationDestinationPatchAsPatch(patchVals, d)

		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceDestinationWithContext(context, updateInstanceDestinationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceDestinationWithContext failed: %s", err.Error()), "ibm_sm_destination", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmDestinationRead(context, d, meta)
}

func resourceIbmSmDestinationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_destination", "delete", "sep-id-parts").GetDiag()
	}

	deleteInstanceDestinationOptions.SetInstanceID(parts[0])
	deleteInstanceDestinationOptions.SetDestinationID(parts[1])

	_, _, err = secretsManagerInstanceManagementClient.DeleteInstanceDestinationWithContext(context, deleteInstanceDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceDestinationWithContext failed: %s", err.Error()), "ibm_sm_destination", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = waitForIbmSmDestinationDelete(d, meta)
	if err != nil {
		errMsg := fmt.Sprintf("error waiting for resource IbmSmDestination (%s) to be deleted: %s", d.Id(), err)
		return flex.DiscriminatedTerraformErrorf(err, errMsg, "ibm_sm_destination", "delete", "wait-for-state").GetDiag()
	}

	d.SetId("")

	return nil
}

func waitForIbmSmDestinationDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		return false, err
	}
	getInstanceDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getInstanceDestinationOptions.SetInstanceID(parts[0])
	getInstanceDestinationOptions.SetDestinationID(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"deleting"},
		Target:  []string{"success", "succeeded", "valid", "active", "done", "normal", "ready", "provisioned", "available", "deployed", "healthy", "pending_reclamation", "deleted", "removed", "stable", "suspended"},
		Refresh: func() (interface{}, string, error) {
			stateObj, response, err := secretsManagerInstanceManagementClient.GetInstanceDestination(getInstanceDestinationOptions)
			if err != nil {
				if _, ok := err.(*core.SDKProblem); ok && response.GetStatusCode() == 404 {
					return nil, "success", nil
				}
				return nil, "", err
			}
			failStates := map[string]bool{"failed": true}
			if failStates[*stateObj.State] {
				return stateObj, *stateObj.State, fmt.Errorf("the instance %s failed: %s", "getInstanceDestinationOptions", err)
			}
			return stateObj, *stateObj.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func ResourceIbmSmDestinationMapToCreateDestinationRequest(modelMap map[string]interface{}) (secretsmanagerinstancemanagementv2.CreateDestinationRequestIntf, error) {
	discValue, ok := modelMap["type"]
	if !ok {
		return nil, fmt.Errorf("discriminator property 'type' not found in map")
	}
	switch discValue {
	case "ibm_cloud_database":
		return ResourceIbmSmDestinationMapToCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(modelMap)
	default:
		return nil, fmt.Errorf("unexpected value for discriminator property 'type' found in map: '%s'", discValue)
	}
}

func ResourceIbmSmDestinationMapToCreateDestinationRequestIbmCloudDatabaseDestinationPrototype(modelMap map[string]interface{}) (*secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype, error) {
	model := &secretsmanagerinstancemanagementv2.CreateDestinationRequestIbmCloudDatabaseDestinationPrototype{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	model.Crn = core.StringPtr(modelMap["crn"].(string))
	return model, nil
}

func ResourceIbmSmDestinationDestinationPatchAsPatch(patchVals *secretsmanagerinstancemanagementv2.DestinationPatch, d *schema.ResourceData) map[string]interface{} {
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
