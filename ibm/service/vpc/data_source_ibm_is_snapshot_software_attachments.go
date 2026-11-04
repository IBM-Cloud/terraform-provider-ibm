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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsSnapshotSoftwareAttachments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsSnapshotSoftwareAttachmentsRead,

		Schema: map[string]*schema.Schema{
			"snapshot_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The snapshot identifier.",
			},
			"software_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The software attachments for the snapshot.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_offering": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this snapshot software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"plan": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The billing plan for the catalog offering version associated with this snapshot softwareattachment.If absent, no billing plan is associated with the catalog offering version (free).",
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
										Description: "The catalog offering version associated with this snapshot software attachment.",
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
							Description: "The date and time that the snapshot software attachment was created.",
						},
						"entitlement": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The entitlement for the snapshot software attachment's licensable software.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"licensable_software": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The licensable software for this snapshot software attachment entitlement. The software will be licensed when an instance is provisioned from this snapshot.",
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
							Description: "The URL for this snapshot software attachment.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this snapshot software attachment.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this snapshot software attachment. The name is unique across all software attachments for the snapshot.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsSnapshotSoftwareAttachmentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot_software_attachments", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listSnapshotSoftwareAttachmentsOptions := &vpcv1.ListSnapshotSoftwareAttachmentsOptions{}

	listSnapshotSoftwareAttachmentsOptions.SetSnapshotID(d.Get("snapshot_id").(string))

	snapshotSoftwareAttachmentCollection, _, err := vpcClient.ListSnapshotSoftwareAttachmentsWithContext(context, listSnapshotSoftwareAttachmentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListSnapshotSoftwareAttachmentsWithContext failed: %s", err.Error()), "(Data) ibm_is_snapshot_software_attachments", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsSnapshotSoftwareAttachmentsID(d))

	softwareAttachments := []map[string]interface{}{}
	for _, softwareAttachmentsItem := range snapshotSoftwareAttachmentCollection.SoftwareAttachments {
		softwareAttachmentsItemMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentToMap(&softwareAttachmentsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_snapshot_software_attachments", "read", "software_attachments-to-map").GetDiag()
		}
		softwareAttachments = append(softwareAttachments, softwareAttachmentsItemMap)
	}
	if err = d.Set("software_attachments", softwareAttachments); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting software_attachments: %s", err), "(Data) ibm_is_snapshot_software_attachments", "read", "set-software_attachments").GetDiag()
	}

	return nil
}

// dataSourceIBMIsSnapshotSoftwareAttachmentsID returns a reasonable ID for the list.
func dataSourceIBMIsSnapshotSoftwareAttachmentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentToMap(model *vpcv1.SnapshotSoftwareAttachment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogOffering != nil {
		catalogOfferingMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentCatalogOfferingToMap(model.CatalogOffering)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog_offering"] = []map[string]interface{}{catalogOfferingMap}
	}
	modelMap["created_at"] = model.CreatedAt.String()
	if model.Entitlement != nil {
		entitlementMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementToMap(model.Entitlement)
		if err != nil {
			return modelMap, err
		}
		modelMap["entitlement"] = []map[string]interface{}{entitlementMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentCatalogOfferingToMap(model *vpcv1.SnapshotSoftwareAttachmentCatalogOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Plan != nil {
		planMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(model.Plan)
		if err != nil {
			return modelMap, err
		}
		modelMap["plan"] = []map[string]interface{}{planMap}
	}
	versionMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(model *vpcv1.CatalogOfferingVersionPlanReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(model *vpcv1.CatalogOfferingVersionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementToMap(model *vpcv1.SnapshotSoftwareAttachmentEntitlement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	licensableSoftware := []map[string]interface{}{}
	for _, licensableSoftwareItem := range model.LicensableSoftware {
		licensableSoftwareItemMap, err := DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementLicensableSoftwareToMap(&licensableSoftwareItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		licensableSoftware = append(licensableSoftware, licensableSoftwareItemMap)
	}
	modelMap["licensable_software"] = licensableSoftware
	return modelMap, nil
}

func DataSourceIBMIsSnapshotSoftwareAttachmentsSnapshotSoftwareAttachmentEntitlementLicensableSoftwareToMap(model *vpcv1.SnapshotSoftwareAttachmentEntitlementLicensableSoftware) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["sku"] = *model.Sku
	return modelMap, nil
}
