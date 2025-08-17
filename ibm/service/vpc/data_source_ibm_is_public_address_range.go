// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMIsPublicAddressRange() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsPublicAddressRangeRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The public address range identifier.",
			},

			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The unique user-defined name for this public-address-range",
			},

			"cidr": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The public IPv4 range, expressed in CIDR format.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the public address range was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this public address range.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this public address range.",
			},
			"ipv4_address_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of IPv4 addresses in this public address range.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the public address range.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this public address range.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this resource group.",
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"target": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target this public address range is bound to.If absent, this pubic address range is not bound to a target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The VPC this public address range is bound to.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this VPC.",
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
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this VPC.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this VPC.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this VPC. The name is unique across all VPCs in the region.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type.",
									},
								},
							},
						},
						"zone": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The zone this public address range resides in.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this zone.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The globally unique name for this zone.",
									},
								},
							},
						},
					},
				},
			},

			isPublicAddressRangeUserTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			isPublicAddressRangeAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},
		},
	}
}

func dataSourceIBMIsPublicAddressRangeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_public_address_range", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var publicAddressRange *vpcv1.PublicAddressRange

	if v, ok := d.GetOk("identifier"); ok {
		getPublicAddressRangeOptions := &vpcv1.GetPublicAddressRangeOptions{}

		getPublicAddressRangeOptions.SetID(v.(string))

		publicAddressRangeinfo, _, err := vpcClient.GetPublicAddressRangeWithContext(context, getPublicAddressRangeOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPublicAddressRangeWithContext failed: %s", err.Error()), "(Data) ibm_public_address_range", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		publicAddressRange = publicAddressRangeinfo
	} else if v, ok := d.GetOk("name"); ok {
		name := v.(string)
		start := ""
		allrecs := []vpcv1.PublicAddressRange{}

		for {
			listPublicAddressRanges := &vpcv1.ListPublicAddressRangesOptions{}
			if start != "" {
				listPublicAddressRanges.Start = &start
			}
			publicAddressRangeCollection, _, err := vpcClient.ListPublicAddressRangesWithContext(context, listPublicAddressRanges)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListPublicAddressRangesWithContext failed: %s", err.Error()), "(Data) ibm_public_address_range", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(publicAddressRangeCollection.Next)
			allrecs = append(allrecs, publicAddressRangeCollection.PublicAddressRanges...)
			if start == "" {
				break
			}
		}

		for _, publicAddressRangeInfo := range allrecs {
			if *publicAddressRangeInfo.Name == name {
				publicAddressRange = &publicAddressRangeInfo
				break
			}
		}
		if publicAddressRange == nil {
			log.Printf("[DEBUG] No publicAddressRange found with name %s", name)
			return diag.FromErr(fmt.Errorf("[ERROR] No Public Address Range found with name %s", name))
		}
	}
	if publicAddressRange == nil {
		log.Printf("[DEBUG] No publicAddressRange found")
		return diag.FromErr(fmt.Errorf("[ERROR] No Public Address Range found"))
	}

	d.SetId(*publicAddressRange.ID)

	if err = d.Set("cidr", publicAddressRange.CIDR); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cidr: %s", err), "(Data) ibm_public_address_range", "read", "set-cidr").GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(publicAddressRange.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_public_address_range", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("crn", publicAddressRange.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_public_address_range", "read", "set-crn").GetDiag()
	}

	if err = d.Set("href", publicAddressRange.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_public_address_range", "read", "set-href").GetDiag()
	}

	if err = d.Set("ipv4_address_count", flex.IntValue(publicAddressRange.Ipv4AddressCount)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ipv4_address_count: %s", err), "(Data) ibm_public_address_range", "read", "set-ipv4_address_count").GetDiag()
	}

	if err = d.Set("lifecycle_state", publicAddressRange.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_public_address_range", "read", "set-lifecycle_state").GetDiag()
	}

	if err = d.Set("name", publicAddressRange.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_public_address_range", "read", "set-name").GetDiag()
	}

	resourceGroup := []map[string]interface{}{}
	if publicAddressRange.ResourceGroup != nil {
		modelMap, err := DataSourceIBMIsPublicAddressRangeResourceGroupReferenceToMap(publicAddressRange.ResourceGroup)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_public_address_range", "read", "resource_group-to-map").GetDiag()
		}
		resourceGroup = append(resourceGroup, modelMap)
	}
	if err = d.Set("resource_group", resourceGroup); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_public_address_range", "read", "set-resource_group").GetDiag()
	}

	if err = d.Set("resource_type", publicAddressRange.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_public_address_range", "read", "set-resource_type").GetDiag()
	}

	target := []map[string]interface{}{}
	if publicAddressRange.Target != nil {
		modelMap, err := DataSourceIBMIsPublicAddressRangePublicAddressRangeTargetToMap(publicAddressRange.Target)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_public_address_range", "read", "target-to-map").GetDiag()
		}
		target = append(target, modelMap)
	}
	if err = d.Set("target", target); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_public_address_range", "read", "set-target").GetDiag()
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *publicAddressRange.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"An error occured during reading of subnet (%s) tags : %s", d.Id(), err)
	}
	d.Set(isPublicAddressRangeUserTags, tags)

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *publicAddressRange.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource subnet (%s) access tags: %s", d.Id(), err)
	}

	d.Set(isPublicAddressRangeAccessTags, accesstags)

	return nil
}

func DataSourceIBMIsPublicAddressRangeResourceGroupReferenceToMap(model *vpcv1.ResourceGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangePublicAddressRangeTargetToMap(model *vpcv1.PublicAddressRangeTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	vpcMap, err := DataSourceIBMIsPublicAddressRangeVPCReferenceToMap(model.VPC)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc"] = []map[string]interface{}{vpcMap}
	zoneMap, err := DataSourceIBMIsPublicAddressRangeZoneReferenceToMap(model.Zone)
	if err != nil {
		return modelMap, err
	}
	modelMap["zone"] = []map[string]interface{}{zoneMap}
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangeVPCReferenceToMap(model *vpcv1.VPCReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsPublicAddressRangeDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangeDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangeZoneReferenceToMap(model *vpcv1.ZoneReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	return modelMap, nil
}
