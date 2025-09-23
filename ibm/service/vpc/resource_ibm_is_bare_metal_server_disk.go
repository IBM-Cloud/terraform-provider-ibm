// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
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

			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the disk was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this bare metal server disk.",
			},
			"interface_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The disk attachment interface used:- `fcp`: Fiber Channel Protocol- `sata`: Serial Advanced Technology Attachment- `nvme`: Non-Volatile Memory ExpressThe enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"size": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size of the disk in GB (gigabytes).",
			},

			"allowed_use": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The usage constraints to match against the requested instance or bare metal server properties to determine compatibility.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bare_metal_server": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An image can only be used for bare metal instantiation if this expression resolves to true.",
						},
						"api_version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The API version with which to evaluate the expressions.",
						},
					},
				},
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

	if err = d.Set("created_at", disk.CreatedAt.String()); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("href", *disk.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-href").GetDiag()
	}

	if err = d.Set("interface_type", *disk.InterfaceType); err != nil {
		err = fmt.Errorf("Error setting interface_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-interface_type").GetDiag()
	}

	if err = d.Set("resource_type", *disk.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-resource_type").GetDiag()
	}

	if err = d.Set("size", *disk.Size); err != nil {
		err = fmt.Errorf("Error setting size: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-interface_type").GetDiag()
	}

	allowedUses := []map[string]interface{}{}
	if disk.AllowedUse != nil {
		modelMap, err := ResourceceIBMIsBareMetalServerDiskAllowedUseToMap(disk.AllowedUse)
		if err != nil {
			err = fmt.Errorf("Error setting allowed_use: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-allowed_use").GetDiag()
		}
		allowedUses = append(allowedUses, modelMap)
	}
	if err = d.Set("allowed_use", allowedUses); err != nil {
		err = fmt.Errorf("Error setting allowed_use: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_disk", "read", "set-allowed_use").GetDiag()
	}
	return nil
}

func ResourceceIBMIsBareMetalServerDiskAllowedUseToMap(model *vpcv1.BareMetalServerDiskAllowedUse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.BareMetalServer != nil {
		modelMap["bare_metal_server"] = *model.BareMetalServer
	}
	if model.ApiVersion != nil {
		modelMap["api_version"] = *model.ApiVersion
	}
	return modelMap, nil
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
