// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsVolumeSoftwareAttachment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVolumeSoftwareAttachmentCreate,
		ReadContext:   resourceIBMIsVolumeSoftwareAttachmentRead,
		UpdateContext: resourceIBMIsVolumeSoftwareAttachmentUpdate,
		DeleteContext: resourceIBMIsVolumeSoftwareAttachmentDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_software_attachment", "volume_id"),
				Description:  "The volume identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_software_attachment", "name"),
				Description:  "The name for this volume software attachment. The name is unique across all software attachments for the volume.",
			},
			"catalog_offering": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this volume software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plan": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The billing plan for the catalog offering version associated with this volume softwareattachment.If absent, no billing plan is associated with the catalog offering version (free).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this[catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user) offering version's billing plan.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A link to documentation about deleted resources.",
												},
											},
										},
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The catalog offering version associated with this volume software attachment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this version of a[catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user) offering.",
									},
								},
							},
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume software attachment was created.",
			},
			"entitlement": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The entitlement for the volume software attachment's licensable software.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"licensable_software": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The licensable software for this volume software attachment entitlement. The software will be licensed when an instance is provisioned from this volume.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sku": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SKU for this licensable software.",
									},
								},
							},
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this volume software attachment.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"is_volume_software_attachment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this volume software attachment.",
			},
		},
	}
}

func ResourceIBMIsVolumeSoftwareAttachmentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "volume_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_volume_software_attachment", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsVolumeSoftwareAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateVolumeSoftwareAttachmentOptions := &vpcv1.UpdateVolumeSoftwareAttachmentOptions{}
	volumeSoftwareAttachmentPatch := &vpcv1.VolumeSoftwareAttachmentPatch{}
	updateVolumeSoftwareAttachmentOptions.SetVolumeID(d.Get("volume_id").(string))

	if _, ok := d.GetOk("name"); ok {
		volumeSoftwareAttachmentPatch.Name = core.StringPtr(d.Get("name").(string))
	}

	volumeSoftwareAttachmentPatchAsPatch := ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentPatchAsPatch(volumeSoftwareAttachmentPatch, d)
	updateVolumeSoftwareAttachmentOptions.VolumeSoftwareAttachmentPatch = volumeSoftwareAttachmentPatchAsPatch
	volumeSoftwareAttachment, _, err := vpcClient.UpdateVolumeSoftwareAttachmentWithContext(context, updateVolumeSoftwareAttachmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVolumeSoftwareAttachmentWithContext failed: %s", err.Error()), "ibm_is_volume_software_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *updateVolumeSoftwareAttachmentOptions.VolumeID, *volumeSoftwareAttachment.ID))

	return resourceIBMIsVolumeSoftwareAttachmentRead(context, d, meta)
}

func resourceIBMIsVolumeSoftwareAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeSoftwareAttachmentOptions := &vpcv1.GetVolumeSoftwareAttachmentOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "sep-id-parts").GetDiag()
	}

	getVolumeSoftwareAttachmentOptions.SetVolumeID(parts[0])
	getVolumeSoftwareAttachmentOptions.SetID(parts[1])

	volumeSoftwareAttachment, response, err := vpcClient.GetVolumeSoftwareAttachmentWithContext(context, getVolumeSoftwareAttachmentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeSoftwareAttachmentWithContext failed: %s", err.Error()), "ibm_is_volume_software_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(volumeSoftwareAttachment.Name) {
		if err = d.Set("name", volumeSoftwareAttachment.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(volumeSoftwareAttachment.CatalogOffering) {
		catalogOfferingMap, err := ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentCatalogOfferingToMap(volumeSoftwareAttachment.CatalogOffering)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "catalog_offering-to-map").GetDiag()
		}
		if err = d.Set("catalog_offering", []map[string]interface{}{catalogOfferingMap}); err != nil {
			err = fmt.Errorf("Error setting catalog_offering: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-catalog_offering").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(volumeSoftwareAttachment.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(volumeSoftwareAttachment.Entitlement) {
		entitlementMap, err := ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementToMap(volumeSoftwareAttachment.Entitlement)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "entitlement-to-map").GetDiag()
		}
		if err = d.Set("entitlement", []map[string]interface{}{entitlementMap}); err != nil {
			err = fmt.Errorf("Error setting entitlement: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-entitlement").GetDiag()
		}
	}
	if err = d.Set("href", volumeSoftwareAttachment.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", volumeSoftwareAttachment.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("is_volume_software_attachment_id", volumeSoftwareAttachment.ID); err != nil {
		err = fmt.Errorf("Error setting is_volume_software_attachment_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "read", "set-is_volume_software_attachment_id").GetDiag()
	}

	return nil
}

func resourceIBMIsVolumeSoftwareAttachmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateVolumeSoftwareAttachmentOptions := &vpcv1.UpdateVolumeSoftwareAttachmentOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_software_attachment", "update", "sep-id-parts").GetDiag()
	}

	updateVolumeSoftwareAttachmentOptions.SetVolumeID(parts[0])
	updateVolumeSoftwareAttachmentOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.VolumeSoftwareAttachmentPatch{}
	if d.HasChange("volume_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "volume_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_volume_software_attachment", "update", "volume_id-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateVolumeSoftwareAttachmentOptions.VolumeSoftwareAttachmentPatch = ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateVolumeSoftwareAttachmentWithContext(context, updateVolumeSoftwareAttachmentOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeSoftwareAttachmentWithContext failed: %s", err.Error()), "ibm_is_volume_software_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsVolumeSoftwareAttachmentRead(context, d, meta)
}

func resourceIBMIsVolumeSoftwareAttachmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentCatalogOfferingToMap(model *vpcv1.VolumeSoftwareAttachmentCatalogOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Plan != nil {
		planMap, err := ResourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model.Plan)
		if err != nil {
			return modelMap, err
		}
		modelMap["plan"] = []map[string]interface{}{planMap}
	}
	versionMap, err := ResourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func ResourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model *vpcv1.CatalogOfferingVersionPlanReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsVolumeSoftwareAttachmentDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	return modelMap, nil
}

func ResourceIBMIsVolumeSoftwareAttachmentDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model *vpcv1.CatalogOfferingVersionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementToMap(model *vpcv1.VolumeSoftwareAttachmentEntitlement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	licensableSoftware := []map[string]interface{}{}
	for _, licensableSoftwareItem := range model.LicensableSoftware {
		licensableSoftwareItemMap, err := ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(&licensableSoftwareItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		licensableSoftware = append(licensableSoftware, licensableSoftwareItemMap)
	}
	modelMap["licensable_software"] = licensableSoftware
	return modelMap, nil
}

func ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(model *vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["sku"] = *model.Sku
	return modelMap, nil
}

func ResourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentPatchAsPatch(patchVals *vpcv1.VolumeSoftwareAttachmentPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}

	return patch
}
