// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ()

func ResourceIBMISDedicatedHostDiskManagement() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMisDedicatedHostDiskManagementCreate,
		ReadContext:   resourceIBMisDedicatedHostDiskManagementRead,
		UpdateContext: resourceIBMisDedicatedHostDiskManagementUpdate,
		DeleteContext: resourceIBMisDedicatedHostDiskManagementDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dedicated_host": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the dedicated host for which disks has to be managed",
			},
			"disks": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Disk information that has to be updated.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The unique identifier for this disk.",
						},
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_dedicated_host_disk_management", "name"),
							Description:  "The user-defined name for this disk. The disk will be updated with this new name",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMISDedicatedHostDiskManagementValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	ibmISDedicatedHostDiskManagementValidator := validate.ResourceValidator{ResourceName: "ibm_is_dedicated_host_disk_management", Schema: validateSchema}
	return &ibmISDedicatedHostDiskManagementValidator
}

func resourceIBMisDedicatedHostDiskManagementCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dedicated_host_disk_management", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	dedicatedhost := d.Get("dedicated_host").(string)
	disks := d.Get("disks")
	diskUpdate := disks.([]interface{})

	for _, disk := range diskUpdate {
		diskItem := disk.(map[string]interface{})
		namestr := diskItem["name"].(string)
		diskid := diskItem["id"].(string)

		updateDedicatedHostDiskOptions := &vpcv1.UpdateDedicatedHostDiskOptions{}
		updateDedicatedHostDiskOptions.SetDedicatedHostID(dedicatedhost)
		updateDedicatedHostDiskOptions.SetID(diskid)
		dedicatedHostDiskPatchModel := &vpcv1.DedicatedHostDiskPatch{
			Name: &namestr,
		}

		dedicatedHostDiskPatch, err := dedicatedHostDiskPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("[ERROR] Error calling asPatch for DedicatedHostDiskPatch: %s", err), "ibm_is_dedicated_host_disk_management", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateDedicatedHostDiskOptions.SetDedicatedHostDiskPatch(dedicatedHostDiskPatch)

		_, _, err = vpcClient.UpdateDedicatedHostDiskWithContext(ctx, updateDedicatedHostDiskOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("[ERROR] Error calling UpdateDedicatedHostDisk: %s", err), "ibm_is_dedicated_host_disk_management", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

	}
	d.SetId(dedicatedhost)
	return resourceIBMisDedicatedHostDiskManagementRead(ctx, d, meta)
}

func resourceIBMisDedicatedHostDiskManagementUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dedicated_host_disk_management", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange("disks") && !d.IsNewResource() {

		disks := d.Get("disks")
		diskUpdate := disks.([]interface{})

		for _, disk := range diskUpdate {
			diskItem := disk.(map[string]interface{})
			namestr := diskItem["name"].(string)
			diskid := diskItem["id"].(string)
			updateDedicatedHostDiskOptions := &vpcv1.UpdateDedicatedHostDiskOptions{}
			updateDedicatedHostDiskOptions.SetDedicatedHostID(d.Id())
			updateDedicatedHostDiskOptions.SetID(diskid)
			dedicatedHostDiskPatchModel := &vpcv1.DedicatedHostDiskPatch{
				Name: &namestr,
			}

			dedicatedHostDiskPatch, err := dedicatedHostDiskPatchModel.AsPatch()
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("[ERROR] Error calling asPatch for DedicatedHostDiskPatch: %s", err), "ibm_is_dedicated_host_disk_management", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			updateDedicatedHostDiskOptions.SetDedicatedHostDiskPatch(dedicatedHostDiskPatch)

			_, response, err := vpcClient.UpdateDedicatedHostDiskWithContext(ctx, updateDedicatedHostDiskOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("[ERROR] Error updating dedicated host disk: %s %s", err, response), "ibm_is_dedicated_host_disk_management", "update")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}

		}

	}
	return resourceIBMisDedicatedHostDiskManagementRead(ctx, d, meta)
}

func resourceIBMisDedicatedHostDiskManagementDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.SetId("")
	return nil
}

func resourceIBMisDedicatedHostDiskManagementRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.Set("dedicated_host", d.Id())

	return nil
}
