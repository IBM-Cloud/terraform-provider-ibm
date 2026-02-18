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

func DataSourceIBMIsInstanceSoftwareAttachment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsInstanceSoftwareAttachmentRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual server instance identifier.",
			},
			"is_instance_software_attachment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance software attachment identifier.",
			},
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
	}
}

func dataSourceIBMIsInstanceSoftwareAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceSoftwareAttachmentOptions := &vpcv1.GetInstanceSoftwareAttachmentOptions{}

	getInstanceSoftwareAttachmentOptions.SetInstanceID(d.Get("instance_id").(string))
	getInstanceSoftwareAttachmentOptions.SetID(d.Get("is_instance_software_attachment_id").(string))

	instanceSoftwareAttachment, _, err := vpcClient.GetInstanceSoftwareAttachmentWithContext(context, getInstanceSoftwareAttachmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceSoftwareAttachmentWithContext failed: %s", err.Error()), "(Data) ibm_is_instance_software_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getInstanceSoftwareAttachmentOptions.InstanceID, *getInstanceSoftwareAttachmentOptions.ID))

	if !core.IsNil(instanceSoftwareAttachment.CatalogOffering) {
		catalogOffering := []map[string]interface{}{}
		catalogOfferingMap, err := DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(instanceSoftwareAttachment.CatalogOffering)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachment", "read", "catalog_offering-to-map").GetDiag()
		}
		catalogOffering = append(catalogOffering, catalogOfferingMap)
		if err = d.Set("catalog_offering", catalogOffering); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting catalog_offering: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-catalog_offering").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(instanceSoftwareAttachment.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(instanceSoftwareAttachment.Entitlement) {
		entitlement := []map[string]interface{}{}
		entitlementMap, err := DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(instanceSoftwareAttachment.Entitlement)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachment", "read", "entitlement-to-map").GetDiag()
		}
		entitlement = append(entitlement, entitlementMap)
		if err = d.Set("entitlement", entitlement); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entitlement: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-entitlement").GetDiag()
		}
	}

	if err = d.Set("href", instanceSoftwareAttachment.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-href").GetDiag()
	}

	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range instanceSoftwareAttachment.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachment", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-lifecycle_reasons").GetDiag()
	}

	if err = d.Set("lifecycle_state", instanceSoftwareAttachment.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-lifecycle_state").GetDiag()
	}

	if err = d.Set("name", instanceSoftwareAttachment.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-name").GetDiag()
	}

	if !core.IsNil(instanceSoftwareAttachment.OfferingInstance) {
		offeringInstance := []map[string]interface{}{}
		offeringInstanceMap, err := DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(instanceSoftwareAttachment.OfferingInstance)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_instance_software_attachment", "read", "offering_instance-to-map").GetDiag()
		}
		offeringInstance = append(offeringInstance, offeringInstanceMap)
		if err = d.Set("offering_instance", offeringInstance); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting offering_instance: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-offering_instance").GetDiag()
		}
	}

	if err = d.Set("resource_type", instanceSoftwareAttachment.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_instance_software_attachment", "read", "set-resource_type").GetDiag()
	}

	return nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentCatalogOfferingToMap(model *vpcv1.InstanceSoftwareAttachmentCatalogOffering) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Plan != nil {
		planMap, err := DataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model.Plan)
		if err != nil {
			return modelMap, err
		}
		modelMap["plan"] = []map[string]interface{}{planMap}
	}
	versionMap, err := DataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model.Version)
	if err != nil {
		return modelMap, err
	}
	modelMap["version"] = []map[string]interface{}{versionMap}
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionPlanReferenceToMap(model *vpcv1.CatalogOfferingVersionPlanReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsInstanceSoftwareAttachmentDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentCatalogOfferingVersionReferenceToMap(model *vpcv1.CatalogOfferingVersionReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementToMap(model *vpcv1.InstanceSoftwareAttachmentEntitlement) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	licensedSoftware := []map[string]interface{}{}
	for _, licensedSoftwareItem := range model.LicensedSoftware {
		licensedSoftwareItemMap, err := DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(&licensedSoftwareItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		licensedSoftware = append(licensedSoftware, licensedSoftwareItemMap)
	}
	modelMap["licensed_software"] = licensedSoftware
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentEntitlementLicensedSoftwareToMap(model *vpcv1.InstanceSoftwareAttachmentEntitlementLicensedSoftware) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["sku"] = *model.Sku
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentLifecycleReasonToMap(model *vpcv1.InstanceSoftwareAttachmentLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsInstanceSoftwareAttachmentInstanceSoftwareAttachmentOfferingInstanceToMap(model *vpcv1.InstanceSoftwareAttachmentOfferingInstance) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}
