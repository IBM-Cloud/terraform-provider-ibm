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

func DataSourceIBMIsInstanceSoftwareAttachments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsInstanceSoftwareAttachmentsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual server instance identifier.",
			},
			"software_attachments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The software attachments for the instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_offering": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this instance software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"plan": &schema.Schema{
										Type:        schema.TypeList,
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
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this instance software attachment.",
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
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this instance software attachment. The name is unique across all instance software attachments for the instance.",
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
					},
				},
			},
		},
	}
}

func dataSourceIBMIsInstanceSoftwareAttachmentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachments", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listInstanceSoftwareAttachmentsOptions := &vpcv1.ListInstanceSoftwareAttachmentsOptions{}

	listInstanceSoftwareAttachmentsOptions.SetInstanceID(d.Get("instance_id").(string))

	instanceSoftwareAttachmentCollection, _, err := vpcClient.ListInstanceSoftwareAttachmentsWithContext(context, listInstanceSoftwareAttachmentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListInstanceSoftwareAttachmentsWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_software_attachments", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsInstanceSoftwareAttachmentsID(d))

	softwareAttachments := []map[string]interface{}{}
	for _, softwareAttachmentsItem := range instanceSoftwareAttachmentCollection.SoftwareAttachments {
		softwareAttachmentsItemMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentToMap(&softwareAttachmentsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachments", "read", "software_attachments-to-map").GetDiag()
		}
		softwareAttachments = append(softwareAttachments, softwareAttachmentsItemMap)
	}
	if err = d.Set("software_attachments", softwareAttachments); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting software_attachments: %s", err), "(Data) ibm_is_instance_software_attachments", "read", "set-software_attachments").GetDiag()
	}

	return nil
}

// dataSourceIBMIsInstanceSoftwareAttachmentsID returns a reasonable ID for the list.
func dataSourceIBMIsInstanceSoftwareAttachmentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentToMap(model *vpcv1.InstanceSoftwareAttachment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogOffering != nil {
		catalogOfferingMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentCatalogOfferingToMap(model.CatalogOffering)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog_offering"] = []map[string]interface{}{catalogOfferingMap}
	}
	modelMap["created_at"] = model.CreatedAt.String()
	if model.Entitlement != nil {
		entitlementMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentEntitlementToMap(model.Entitlement)
		if err != nil {
			return modelMap, err
		}
		modelMap["entitlement"] = []map[string]interface{}{entitlementMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	modelMap["lifecycle_state"] = *model.LifecycleState
	modelMap["name"] = *model.Name
	if model.OfferingInstance != nil {
		offeringInstanceMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentOfferingInstanceToMap(model.OfferingInstance)
		if err != nil {
			return modelMap, err
		}
		modelMap["offering_instance"] = []map[string]interface{}{offeringInstanceMap}
	}
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentCatalogOfferingToMap(model *vpcv1.InstanceSoftwareAttachmentCatalogOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Plan != nil {
		planMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(model.Plan)
		if err != nil {
			return modelMap, err
		}
		modelMap["plan"] = []map[string]interface{}{planMap}
	}
	versionMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsCatalogOfferingVersionPlanReferenceToMap(model *vpcv1.CatalogOfferingVersionPlanReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsCatalogOfferingVersionReferenceToMap(model *vpcv1.CatalogOfferingVersionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentEntitlementToMap(model *vpcv1.InstanceSoftwareAttachmentEntitlement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	licensedSoftware := []map[string]interface{}{}
	for _, licensedSoftwareItem := range model.LicensedSoftware {
		licensedSoftwareItemMap, err := DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(&licensedSoftwareItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		licensedSoftware = append(licensedSoftware, licensedSoftwareItemMap)
	}
	modelMap["licensed_software"] = licensedSoftware
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(model *vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["sku"] = *model.Sku
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentLifecycleReasonToMap(model *vpcv1.InstanceSoftwareAttachmentLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentsInstanceSoftwareAttachmentOfferingInstanceToMap(model *vpcv1.InstanceSoftwareAttachmentOfferingInstance) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}
