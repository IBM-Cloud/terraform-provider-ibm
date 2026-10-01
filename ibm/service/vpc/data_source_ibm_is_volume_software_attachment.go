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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVolumeSoftwareAttachment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVolumeSoftwareAttachmentRead,

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The volume identifier.",
			},
			"is_volume_software_attachment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The volume software attachment identifier.",
			},
			"catalog_offering": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this volume software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plan": &schema.Schema{
							Type:        schema.TypeList,
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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this volume software attachment. The name is unique across all software attachments for the volume.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
		},
	}
}

func dataSourceIBMIsVolumeSoftwareAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_software_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeSoftwareAttachmentOptions := &vpcv1.GetVolumeSoftwareAttachmentOptions{}

	getVolumeSoftwareAttachmentOptions.SetVolumeID(d.Get("volume_id").(string))
	getVolumeSoftwareAttachmentOptions.SetID(d.Get("is_volume_software_attachment_id").(string))

	volumeSoftwareAttachment, _, err := vpcClient.GetVolumeSoftwareAttachmentWithContext(context, getVolumeSoftwareAttachmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeSoftwareAttachmentWithContext failed: %s", err.Error()), "(Data) ibm_is_volume_software_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getVolumeSoftwareAttachmentOptions.VolumeID, *getVolumeSoftwareAttachmentOptions.ID))

	if !core.IsNil(volumeSoftwareAttachment.CatalogOffering) {
		catalogOffering := []map[string]interface{}{}
		catalogOfferingMap, err := DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentCatalogOfferingToMap(volumeSoftwareAttachment.CatalogOffering)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_software_attachment", "read", "catalog_offering-to-map").GetDiag()
		}
		catalogOffering = append(catalogOffering, catalogOfferingMap)
		if err = d.Set("catalog_offering", catalogOffering); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_volume_software_attachment", "read", "set-catalog_offering").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(volumeSoftwareAttachment.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_volume_software_attachment", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(volumeSoftwareAttachment.Entitlement) {
		entitlement := []map[string]interface{}{}
		entitlementMap, err := DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementToMap(volumeSoftwareAttachment.Entitlement)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_software_attachment", "read", "entitlement-to-map").GetDiag()
		}
		entitlement = append(entitlement, entitlementMap)
		if err = d.Set("entitlement", entitlement); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entitlement: %s", err), "(Data) ibm_is_volume_software_attachment", "read", "set-entitlement").GetDiag()
		}
	}

	if err = d.Set("href", volumeSoftwareAttachment.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_volume_software_attachment", "read", "set-href").GetDiag()
	}

	if err = d.Set("name", volumeSoftwareAttachment.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_volume_software_attachment", "read", "set-name").GetDiag()
	}

	if err = d.Set("resource_type", volumeSoftwareAttachment.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_volume_software_attachment", "read", "set-resource_type").GetDiag()
	}

	return nil
}

func DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentCatalogOfferingToMap(model *vpcv1.VolumeSoftwareAttachmentCatalogOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Plan != nil {
		planMap, err := DataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model.Plan)
		if err != nil {
			return modelMap, err
		}
		modelMap["plan"] = []map[string]interface{}{planMap}
	}
	versionMap, err := DataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func DataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model *vpcv1.CatalogOfferingVersionPlanReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsVolumeSoftwareAttachmentDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsVolumeSoftwareAttachmentDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsVolumeSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model *vpcv1.CatalogOfferingVersionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementToMap(model *vpcv1.VolumeSoftwareAttachmentEntitlement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	licensableSoftware := []map[string]interface{}{}
	for _, licensableSoftwareItem := range model.LicensableSoftware {
		licensableSoftwareItemMap, err := DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(&licensableSoftwareItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		licensableSoftware = append(licensableSoftware, licensableSoftwareItemMap)
	}
	modelMap["licensable_software"] = licensableSoftware
	return modelMap, nil
}

func DataSourceIBMIsVolumeSoftwareAttachmentVolumeSoftwareAttachmentEntitlementLicensableSoftwareToMap(model *vpcv1.VolumeSoftwareAttachmentEntitlementLicensableSoftware) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["sku"] = *model.Sku
	return modelMap, nil
}
