// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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

func DataSourceIBMIsPublicAddressRanges() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsPublicAddressRangesRead,

		Schema: map[string]*schema.Schema{
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filters the collection to resources with a `resource_group` property matching the specified identifier.",
			},
			"public_address_ranges": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of public address ranges.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this public address range.",
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
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this public address range. The name is unique across all public address ranges in the region.",
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
				},
			},
		},
	}
}

func dataSourceIBMIsPublicAddressRangesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_public_address_ranges", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	start := ""
	allrecs := []vpcv1.PublicAddressRange{}
	rgroup := ""

	if rg, ok := d.GetOk("resource_group"); ok {
		rgroup = rg.(string)
	}

	for {
		listPublicAddressRanges := &vpcv1.ListPublicAddressRangesOptions{}
		if start != "" {
			listPublicAddressRanges.Start = &start
		}
		if rgroup != "" {
			listPublicAddressRanges.ResourceGroupID = &rgroup
		}
		publicAddressRangeCollection, response, err := vpcClient.ListPublicAddressRangesWithContext(context, listPublicAddressRanges)
		if err != nil {
			log.Printf("[DEBUG] ListPublicAddressRangesWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("[ERROR] ListPublicAddressRangesWithContext failed %s\n%s", err, response))
		}
		start = flex.GetNext(publicAddressRangeCollection.Next)
		allrecs = append(allrecs, publicAddressRangeCollection.PublicAddressRanges...)
		if start == "" {
			break
		}
	}

	d.SetId(dataSourceIBMIsPublicAddressRangesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allrecs {
		modelMap, err := DataSourceIBMIsPublicAddressRangesPublicAddressRangeToMap(&modelItem, meta)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_public_address_ranges", "read", "PublicAddressRanges-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("public_address_ranges", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting public_address_ranges %s", err), "(Data) ibm_public_address_ranges", "read", "public_address_ranges-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsPublicAddressRangesID returns a reasonable ID for the list.
func dataSourceIBMIsPublicAddressRangesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsPublicAddressRangesPublicAddressRangeToMap(model *vpcv1.PublicAddressRange, meta interface{}) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["cidr"] = *model.CIDR
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["crn"] = *model.CRN
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["ipv4_address_count"] = flex.IntValue(model.Ipv4AddressCount)
	modelMap["lifecycle_state"] = *model.LifecycleState
	modelMap["name"] = *model.Name
	resourceGroupMap, err := DataSourceIBMIsPublicAddressRangesResourceGroupReferenceToMap(model.ResourceGroup)
	if err != nil {
		return modelMap, err
	}
	modelMap["resource_group"] = []map[string]interface{}{resourceGroupMap}
	modelMap["resource_type"] = *model.ResourceType
	if model.Target != nil {
		targetMap, err := DataSourceIBMIsPublicAddressRangesPublicAddressRangeTargetToMap(model.Target)
		if err != nil {
			return modelMap, err
		}
		modelMap["target"] = []map[string]interface{}{targetMap}
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *model.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc SSH Key (%s) user tags: %s", *model.ID, err)
	}
	modelMap[isPublicAddressRangeUserTags] = tags
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *model.CRN, "", isKeyAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource SSH Key (%s) access tags: %s", *model.ID, err)
	}
	modelMap[isPublicAddressRangeAccessTags] = accesstags
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangesResourceGroupReferenceToMap(model *vpcv1.ResourceGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangesPublicAddressRangeTargetToMap(model *vpcv1.PublicAddressRangeTarget) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	vpcMap, err := DataSourceIBMIsPublicAddressRangesVPCReferenceToMap(model.VPC)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc"] = []map[string]interface{}{vpcMap}
	zoneMap, err := DataSourceIBMIsPublicAddressRangesZoneReferenceToMap(model.Zone)
	if err != nil {
		return modelMap, err
	}
	modelMap["zone"] = []map[string]interface{}{zoneMap}
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangesVPCReferenceToMap(model *vpcv1.VPCReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsPublicAddressRangesDeletedToMap(model.Deleted)
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

func DataSourceIBMIsPublicAddressRangesDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsPublicAddressRangesZoneReferenceToMap(model *vpcv1.ZoneReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["name"] = *model.Name
	return modelMap, nil
}
