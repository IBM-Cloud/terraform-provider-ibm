// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMIsBareMetalServerDisk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISBareMetalServerDiskCreate,
		ReadContext:   resourceIBMISBareMetalServerDiskRead,
		UpdateContext: resourceIBMISBareMetalServerDiskUpdate,
		DeleteContext: resourceIBMISBareMetalServerDiskDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bare metal server identifier",
			},
			isBareMetalServerDisk: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bare metal server disk identifier",
			},

			isBareMetalServerDiskName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Bare metal server disk name",
				ValidateFunc: validate.InvokeValidator("ibm_is_bare_metal_server_disk", isBareMetalServerDiskName),
			},
		},
	}
}

func ResourceIBMIsBareMetalServerDiskValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isBareMetalServerDiskName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	ibmISBareMetalServerDiskResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_bare_metal_server_disk", Schema: validateSchema}
	return &ibmISBareMetalServerDiskResourceValidator
}

func resourceIBMISBareMetalServerDiskCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var bareMetalServerId, diskId, diskName string
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}
	if bmsDiskId, ok := d.GetOk(isBareMetalServerDisk); ok {
		diskId = bmsDiskId.(string)
	}
	if bmsDiskName, ok := d.GetOk(isBareMetalServerDiskName); ok {
		diskName = bmsDiskName.(string)
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.UpdateBareMetalServerDiskOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &diskId,
	}
	diskPatchModel := &vpcv1.BareMetalServerDiskPatch{
		Name: &diskName,
	}
	diskPatch, err := diskPatchModel.AsPatch()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("diskPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_bare_metal_server_disk", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options.BareMetalServerDiskPatch = diskPatch
	disk, _, err := sess.UpdateBareMetalServerDiskWithContext(context, options)
	if err != nil || disk == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateBareMetalServerNetworkInterfaceWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_disk", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*disk.ID)
	diagErr := bareMetalServerDiskGet(context, d, sess, bareMetalServerId, diskId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func bareMetalServerDiskGet(context context.Context, d *schema.ResourceData, sess *vpcv1.VpcV1, bareMetalServerId, diskId string) diag.Diagnostics {

	options := &vpcv1.GetBareMetalServerDiskOptions{
		BareMetalServerID: &bareMetalServerId,
		ID:                &diskId,
	}
	disk, response, err := sess.GetBareMetalServerDiskWithContext(context, options)
	if err != nil || disk == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerDiskWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_disk", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set(isBareMetalServerID, bareMetalServerId); err != nil {
		err = fmt.Errorf("Error setting bare_metal_server: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-bare_metal_server").GetDiag()
	}
	if err = d.Set(isBareMetalServerDisk, *disk.ID); err != nil {
		err = fmt.Errorf("Error setting disk: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-disk").GetDiag()
	}
	if err = d.Set(isBareMetalServerDiskName, *disk.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-name").GetDiag()
	}
	return nil
}

func resourceIBMISBareMetalServerDiskRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var bareMetalServerId, diskId string
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}
	if bmsDiskId, ok := d.GetOk(isBareMetalServerDisk); ok {
		diskId = bmsDiskId.(string)
	}
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	diagErr := bareMetalServerDiskGet(context, d, sess, bareMetalServerId, diskId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func resourceIBMISBareMetalServerDiskUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	if d.HasChange(isBareMetalServerDiskName) {
		var bareMetalServerId, diskId, diskName string
		if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
			bareMetalServerId = bmsId.(string)
		}
		if bmsDiskId, ok := d.GetOk(isBareMetalServerDisk); ok {
			diskId = bmsDiskId.(string)
		}
		if bmsDiskName, ok := d.GetOk(isBareMetalServerDiskName); ok {
			diskName = bmsDiskName.(string)
		}

		sess, err := vpcClient(meta)
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "update", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options := &vpcv1.UpdateBareMetalServerDiskOptions{
			BareMetalServerID: &bareMetalServerId,
			ID:                &diskId,
		}
		diskPatchModel := &vpcv1.BareMetalServerDiskPatch{
			Name: &diskName,
		}
		diskPatch, err := diskPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("diskPatchModel.AsPatch() failed: %s", err.Error()), "ibm_is_bare_metal_server_disk", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.BareMetalServerDiskPatch = diskPatch
		disk, _, err := sess.UpdateBareMetalServerDiskWithContext(context, options)
		if err != nil || disk == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateBareMetalServerDiskWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_disk", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		diagErr := bareMetalServerDiskGet(context, d, sess, bareMetalServerId, diskId)
		if diagErr != nil {
			return diagErr
		}
	}
	return nil
}

func resourceIBMISBareMetalServerDiskDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")

	return nil
}
