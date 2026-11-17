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

func ResourceIbmSmHsmHaGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmHaGroupCreate,
		ReadContext:   resourceIbmSmHsmHaGroupRead,
		UpdateContext: resourceIbmSmHsmHaGroupUpdate,
		DeleteContext: resourceIbmSmHsmHaGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_ha_group", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"destination_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_ha_group", "destination_id"),
				Description:  "Destination ID.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_ha_group", "name"),
				Description:  "Human-readable display name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_ha_group", "description"),
				Description:  "Optional description. Omitted if not set.",
			},
			"group_label": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_ha_group", "group_label"),
				Description:  "HA Group virtual token label. Immutable after creation.",
			},
			"group_serial": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_ha_group", "group_serial"),
				Description:  "HA Group serial number. Immutable after creation.",
			},
			"member_serials": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "Ordered list of HSM node member serial numbers for this HA Group. Max 100 unique members.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the HA Group resource.",
			},
			"provisioning_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Lifecycle state of the HA Group:- `completed`: HA Group has been created and its config has been staged in the HSM config file. Initial and only state.**Note:** New values may be added to this enumeration in future API versions. Clients MUST handle unknown values gracefully and MUST NOT fail on unrecognised states.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AS OF health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Initial state on creation. No health check has been performed yet.- `evaluating`: A health check is in progress for this resource.- `in_progress`: An apply operation is running across the destination.- `healthy`: Configured AND all member HSM nodes are healthy.- `degraded`: Configured AND any one member is unhealthy (but not all).- `unhealthy`: Configured AND all members are unhealthy or unknown.- `not_configured`: Prerequisites not met (destination needs ≥1 path with active non-expired Server CA AND ≥1 active non-expired client certificate).",
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
			"ha_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "HA Group ID. Immutable. System-managed.",
			},
		},
	}
}

func ResourceIbmSmHsmHaGroupValidator() *validate.ResourceValidator {
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
			Identifier:                 "group_label",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[ -~]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "group_serial",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_ha_group", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmHaGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationHaGroupOptions{}

	createInstanceHsmDestinationHaGroupOptions.SetInstanceID(d.Get("instance_id").(string))
	createInstanceHsmDestinationHaGroupOptions.SetDestinationID(d.Get("destination_id").(string))
	createInstanceHsmDestinationHaGroupOptions.SetName(d.Get("name").(string))
	createInstanceHsmDestinationHaGroupOptions.SetGroupLabel(d.Get("group_label").(string))
	createInstanceHsmDestinationHaGroupOptions.SetGroupSerial(d.Get("group_serial").(string))
	var memberSerials []string
	for _, v := range d.Get("member_serials").([]interface{}) {
		memberSerialsItem := v.(string)
		memberSerials = append(memberSerials, memberSerialsItem)
	}
	createInstanceHsmDestinationHaGroupOptions.SetMemberSerials(memberSerials)
	if _, ok := d.GetOk("description"); ok {
		createInstanceHsmDestinationHaGroupOptions.SetDescription(d.Get("description").(string))
	}

	hsmhaGroup, _, err := secretsManagerInstanceManagementClient.CreateInstanceHsmDestinationHaGroupWithContext(context, createInstanceHsmDestinationHaGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceHsmDestinationHaGroupWithContext failed: %s", err.Error()), "ibm_sm_hsm_ha_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createInstanceHsmDestinationHaGroupOptions.InstanceID, *createInstanceHsmDestinationHaGroupOptions.DestinationID, *hsmhaGroup.ID))

	return resourceIbmSmHsmHaGroupRead(context, d, meta)
}

func resourceIbmSmHsmHaGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationHaGroupOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationHaGroupOptions.SetDestinationID(parts[1])
	getInstanceHsmDestinationHaGroupOptions.SetHaGroupID(parts[2])

	hsmhaGroup, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationHaGroupWithContext(context, getInstanceHsmDestinationHaGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationHaGroupWithContext failed: %s", err.Error()), "ibm_sm_hsm_ha_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_id", parts[0]); err != nil {
		err = fmt.Errorf("error setting instance_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-instance_id").GetDiag()
	}
	if err = d.Set("destination_id", parts[1]); err != nil {
		err = fmt.Errorf("error setting destination_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-destination_id").GetDiag()
	}
	if err = d.Set("name", hsmhaGroup.Name); err != nil {
		err = fmt.Errorf("error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-name").GetDiag()
	}
	if !core.IsNil(hsmhaGroup.Description) {
		if err = d.Set("description", hsmhaGroup.Description); err != nil {
			err = fmt.Errorf("error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("group_label", hsmhaGroup.GroupLabel); err != nil {
		err = fmt.Errorf("error setting group_label: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-group_label").GetDiag()
	}
	if err = d.Set("group_serial", hsmhaGroup.GroupSerial); err != nil {
		err = fmt.Errorf("error setting group_serial: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-group_serial").GetDiag()
	}
	if err = d.Set("member_serials", hsmhaGroup.MemberSerials); err != nil {
		err = fmt.Errorf("error setting member_serials: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-member_serials").GetDiag()
	}
	if !core.IsNil(hsmhaGroup.Href) {
		if err = d.Set("href", hsmhaGroup.Href); err != nil {
			err = fmt.Errorf("error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-href").GetDiag()
		}
	}
	if err = d.Set("provisioning_state", hsmhaGroup.ProvisioningState); err != nil {
		err = fmt.Errorf("error setting provisioning_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-provisioning_state").GetDiag()
	}
	if err = d.Set("health_state", hsmhaGroup.HealthState); err != nil {
		err = fmt.Errorf("error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("health_as_of", flex.DateTimeToString(hsmhaGroup.HealthAsOf)); err != nil {
		err = fmt.Errorf("error setting health_as_of: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-health_as_of").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(hsmhaGroup.CreatedAt)); err != nil {
		err = fmt.Errorf("error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(hsmhaGroup.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(hsmhaGroup.UpdatedAt)); err != nil {
			err = fmt.Errorf("error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-updated_at").GetDiag()
		}
	}
	if err = d.Set("created_by", hsmhaGroup.CreatedBy); err != nil {
		err = fmt.Errorf("error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(hsmhaGroup.UpdatedBy) {
		if err = d.Set("updated_by", hsmhaGroup.UpdatedBy); err != nil {
			err = fmt.Errorf("error setting updated_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-updated_by").GetDiag()
		}
	}
	if !core.IsNil(hsmhaGroup.ID) {
		if err = d.Set("ha_group_id", hsmhaGroup.ID); err != nil {
			err = fmt.Errorf("error setting ha_group_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "read", "set-ha_group_id").GetDiag()
		}
	}

	return nil
}

func resourceIbmSmHsmHaGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationHaGroupOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceHsmDestinationHaGroupOptions.SetInstanceID(parts[0])
	updateInstanceHsmDestinationHaGroupOptions.SetDestinationID(parts[1])
	updateInstanceHsmDestinationHaGroupOptions.SetHaGroupID(parts[2])

	hasChange := false

	patchVals := &secretsmanagerinstancemanagementv2.HSMHAGroupPatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_ha_group", "update", "instance_id-forces-new").GetDiag()
	}
	if d.HasChange("destination_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "destination_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_ha_group", "update", "destination_id-forces-new").GetDiag()
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
	if d.HasChange("member_serials") {
		var memberSerials []string
		for _, v := range d.Get("member_serials").([]interface{}) {
			memberSerialsItem := v.(string)
			memberSerials = append(memberSerials, memberSerialsItem)
		}
		patchVals.MemberSerials = memberSerials
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateInstanceHsmDestinationHaGroupOptions.HSMHAGroupPatch = ResourceIbmSmHsmHaGroupHSMHAGroupPatchAsPatch(patchVals, d)

		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationHaGroupWithContext(context, updateInstanceHsmDestinationHaGroupOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationHaGroupWithContext failed: %s", err.Error()), "ibm_sm_hsm_ha_group", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmHsmHaGroupRead(context, d, meta)
}

func resourceIbmSmHsmHaGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceHsmDestinationHaGroupOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationHaGroupOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_ha_group", "delete", "sep-id-parts").GetDiag()
	}

	deleteInstanceHsmDestinationHaGroupOptions.SetInstanceID(parts[0])
	deleteInstanceHsmDestinationHaGroupOptions.SetDestinationID(parts[1])
	deleteInstanceHsmDestinationHaGroupOptions.SetHaGroupID(parts[2])

	_, err = secretsManagerInstanceManagementClient.DeleteInstanceHsmDestinationHaGroupWithContext(context, deleteInstanceHsmDestinationHaGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceHsmDestinationHaGroupWithContext failed: %s", err.Error()), "ibm_sm_hsm_ha_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmSmHsmHaGroupHSMHAGroupPatchAsPatch(patchVals *secretsmanagerinstancemanagementv2.HSMHAGroupPatch, d *schema.ResourceData) map[string]interface{} {
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
	path = "member_serials"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["member_serials"] = nil
	} else if !exists {
		delete(patch, "member_serials")
	}

	return patch
}
