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

func ResourceIBMIsInstanceSoftwareAttachment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsInstanceSoftwareAttachmentCreate,
		ReadContext:   resourceIBMIsInstanceSoftwareAttachmentRead,
		UpdateContext: resourceIBMIsInstanceSoftwareAttachmentUpdate,
		DeleteContext: resourceIBMIsInstanceSoftwareAttachmentDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_software_attachment", "instance_id"),
				Description:  "The virtual server instance identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_instance_software_attachment", "name"),
				Description:  "The name for this instance software attachment. The name is unique across all instance software attachments for the instance.",
			},
			"catalog_offering": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this instance software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plan": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The billing plan for the catalog offering version associated with this instance softwareattachment.If absent, no billing plan is associated with the catalog offering version (free).",
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
							Description: "The catalog offering version associated with this instance software attachment.",
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
				Description: "The date and time that the instance software attachment was created.",
			},
			"entitlement": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The entitlement for the licensed software for this instance software attachment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"licensed_software": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The licensed software for this instance software attachment entitlement.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sku": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SKU for this licensed software.",
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
				Description: "The URL for this instance software attachment.",
			},
			"lifecycle_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The lifecycle reasons for this instance software attachment (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A reason code for this lifecycle state:- `failed_registration`: the software instance's registration to Resource Controller,  which includes creation of any required software license(s), has failed. Delete the  instance and provision it again. If the problem persists, contact IBM Support.- `internal_error`: internal error (contact IBM support)- `pending_registration`: the software instance's registration to Resource Controller,  and the creation of any required software license(s), is being processed.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this lifecycle state.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A link to documentation about the reason for this lifecycle state.",
						},
					},
				},
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the instance software attachment.",
			},
			"offering_instance": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for the software offering instance registered with Resource Controller that is associated with the instance software attachment.",
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"is_instance_software_attachment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this instance software attachment.",
			},
		},
	}
}

func ResourceIBMIsInstanceSoftwareAttachmentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_instance_software_attachment", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsInstanceSoftwareAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceSoftwareAttachmentOptions := &vpcv1.UpdateInstanceSoftwareAttachmentOptions{}
	instanceSoftwareAttachmentPatch := &vpcv1.InstanceSoftwareAttachmentPatch{}
	updateInstanceSoftwareAttachmentOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("name"); ok {
		instanceSoftwareAttachmentPatch.Name = core.StringPtr(d.Get("name").(string))
	}
	instanceSoftwareAttachmentPatchAsPatch := ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentPatchAsPatch(instanceSoftwareAttachmentPatch, d)
	updateInstanceSoftwareAttachmentOptions.InstanceSoftwareAttachmentPatch = instanceSoftwareAttachmentPatchAsPatch
	instanceSoftwareAttachment, _, err := vpcClient.UpdateInstanceSoftwareAttachmentWithContext(context, updateInstanceSoftwareAttachmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceSoftwareAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_software_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *updateInstanceSoftwareAttachmentOptions.InstanceID, *instanceSoftwareAttachment.ID))

	return resourceIBMIsInstanceSoftwareAttachmentRead(context, d, meta)
}

func resourceIBMIsInstanceSoftwareAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceSoftwareAttachmentOptions := &vpcv1.GetInstanceSoftwareAttachmentOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "sep-id-parts").GetDiag()
	}

	getInstanceSoftwareAttachmentOptions.SetInstanceID(parts[0])
	getInstanceSoftwareAttachmentOptions.SetID(parts[1])

	instanceSoftwareAttachment, response, err := vpcClient.GetInstanceSoftwareAttachmentWithContext(context, getInstanceSoftwareAttachmentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceSoftwareAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_software_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(instanceSoftwareAttachment.Name) {
		if err = d.Set("name", instanceSoftwareAttachment.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(instanceSoftwareAttachment.CatalogOffering) {
		catalogOfferingMap, err := ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(instanceSoftwareAttachment.CatalogOffering)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "catalog_offering-to-map").GetDiag()
		}
		if err = d.Set("catalog_offering", []map[string]interface{}{catalogOfferingMap}); err != nil {
			err = fmt.Errorf("Error setting catalog_offering: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-catalog_offering").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(instanceSoftwareAttachment.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(instanceSoftwareAttachment.Entitlement) {
		entitlementMap, err := ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(instanceSoftwareAttachment.Entitlement)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "entitlement-to-map").GetDiag()
		}
		if err = d.Set("entitlement", []map[string]interface{}{entitlementMap}); err != nil {
			err = fmt.Errorf("Error setting entitlement: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-entitlement").GetDiag()
		}
	}
	if err = d.Set("href", instanceSoftwareAttachment.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range instanceSoftwareAttachment.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-lifecycle_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", instanceSoftwareAttachment.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-lifecycle_state").GetDiag()
	}
	if !core.IsNil(instanceSoftwareAttachment.OfferingInstance) {
		offeringInstanceMap, err := ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(instanceSoftwareAttachment.OfferingInstance)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "offering_instance-to-map").GetDiag()
		}
		if err = d.Set("offering_instance", []map[string]interface{}{offeringInstanceMap}); err != nil {
			err = fmt.Errorf("Error setting offering_instance: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-offering_instance").GetDiag()
		}
	}
	if err = d.Set("resource_type", instanceSoftwareAttachment.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("is_instance_software_attachment_id", instanceSoftwareAttachment.ID); err != nil {
		err = fmt.Errorf("Error setting is_instance_software_attachment_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "read", "set-is_instance_software_attachment_id").GetDiag()
	}

	return nil
}

func resourceIBMIsInstanceSoftwareAttachmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceSoftwareAttachmentOptions := &vpcv1.UpdateInstanceSoftwareAttachmentOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_instance_software_attachment", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceSoftwareAttachmentOptions.SetInstanceID(parts[0])
	updateInstanceSoftwareAttachmentOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.InstanceSoftwareAttachmentPatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_instance_software_attachment", "update", "instance_id-forces-new").GetDiag()
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
		updateInstanceSoftwareAttachmentOptions.InstanceSoftwareAttachmentPatch = ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateInstanceSoftwareAttachmentWithContext(context, updateInstanceSoftwareAttachmentOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceSoftwareAttachmentWithContext failed: %s", err.Error()), "ibm_is_instance_software_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsInstanceSoftwareAttachmentRead(context, d, meta)
}

func resourceIBMIsInstanceSoftwareAttachmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(model *vpcv1.InstanceSoftwareAttachmentCatalogOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Plan != nil {
		planMap, err := ResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model.Plan)
		if err != nil {
			return modelMap, err
		}
		modelMap["plan"] = []map[string]interface{}{planMap}
	}
	versionMap, err := ResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model *vpcv1.CatalogOfferingVersionPlanReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsInstanceSoftwareAttachmentDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model *vpcv1.CatalogOfferingVersionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(model *vpcv1.InstanceSoftwareAttachmentEntitlement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	licensedSoftware := []map[string]interface{}{}
	for _, licensedSoftwareItem := range model.LicensedSoftware {
		licensedSoftwareItemMap, err := ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(&licensedSoftwareItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		licensedSoftware = append(licensedSoftware, licensedSoftwareItemMap)
	}
	modelMap["licensed_software"] = licensedSoftware
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(model *vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["sku"] = *model.Sku
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(model *vpcv1.InstanceSoftwareAttachmentLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(model *vpcv1.InstanceSoftwareAttachmentOfferingInstance) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func ResourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentPatchAsPatch(patchVals *vpcv1.InstanceSoftwareAttachmentPatch, d *schema.ResourceData) map[string]interface{} {
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
