// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	floatingIPName                  = "name"
	floatingIPAddress               = "address"
	floatingIPStatus                = "status"
	floatingIPZone                  = "zone"
	floatingIPTarget                = "target"
	floatingIPTargets               = "target_list"
	floatingIPTargetsHref           = "href"
	floatingIPTargetsCrn            = "crn"
	floatingIPTargetsDeleted        = "deleted"
	floatingIPTargetsMoreInfo       = "more_info"
	floatingIPTargetsId             = "id"
	floatingIPTargetsName           = "name"
	floatingIPTargetsResourceType   = "resource_type"
	floatingIPTags                  = "tags"
	floatingIPCRN                   = "crn"
	floatingIpPrimaryIP             = "primary_ip"
	floatingIpPrimaryIpAddress      = "address"
	floatingIpPrimaryIpHref         = "href"
	floatingIpPrimaryIpName         = "name"
	floatingIpPrimaryIpId           = "reserved_ip"
	floatingIpPrimaryIpResourceType = "resource_type"
)

func DataSourceIBMISFloatingIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISFloatingIPRead,

		Schema: map[string]*schema.Schema{

			floatingIPName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the floating IP",
			},

			floatingIPAddress: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP address",
			},

			floatingIPStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP status",
			},

			floatingIPZone: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name",
			},

			floatingIPTarget: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target info",
			},

			floatingIPTargets: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The target of this floating IP.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						floatingIPTargetsDeleted: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									floatingIPTargetsMoreInfo: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						floatingIPTargetsHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this network interface.",
						},
						floatingIPTargetsId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this network interface.",
						},
						floatingIPTargetsName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this network interface.",
						},
						floatingIpPrimaryIP: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									floatingIpPrimaryIpAddress: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address to reserve, which must not already be reserved on the subnet.",
									},
									floatingIpPrimaryIpHref: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this reserved IP",
									},
									floatingIpPrimaryIpName: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this reserved IP. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the subnet the reserved IP resides in. ",
									},
									floatingIpPrimaryIpId: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identifies a reserved IP by a unique property.",
									},
									floatingIpPrimaryIpResourceType: {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The resource type",
									},
								},
							},
						},
						floatingIPTargetsResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						floatingIPTargetsCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this public gateway.",
						},
					},
				},
			},

			floatingIPCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Floating IP crn",
			},

			floatingIPTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "Floating IP tags",
			},

			isFloatingIPAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func dataSourceIBMISFloatingIPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	floatingIPName := d.Get(isFloatingIPName).(string)
	diag := floatingIPGet(ctx, d, meta, floatingIPName)
	if diag != nil {
		return diag
	}
	return nil
}

func floatingIPGet(ctx context.Context, d *schema.ResourceData, meta interface{}, name string) diag.Diagnostics { // Changed return type
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_ibm_is_floating_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	start := ""
	allFloatingIPs := []vpcv1.FloatingIP{}
	for {
		floatingIPOptions := &vpcv1.ListFloatingIpsOptions{}
		if start != "" {
			floatingIPOptions.Start = &start
		}
		floatingIPs, _, err := vpcClient.ListFloatingIpsWithContext(ctx, floatingIPOptions) // Use WithContext
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListFloatingIpsWithContext failed: %s", err.Error()), "(Data) ibm_ibm_is_floating_ip", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(floatingIPs.Next)
		allFloatingIPs = append(allFloatingIPs, floatingIPs.FloatingIps...)
		if start == "" {
			break
		}
	}

	for _, floatingIP := range allFloatingIPs {
		if *floatingIP.Name == name {

			if err = d.Set("name", floatingIP.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-name").GetDiag()
			}
			if err = d.Set("address", floatingIP.Address); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting address: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-address").GetDiag()
			}
			if err = d.Set("status", floatingIP.Status); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-status").GetDiag()
			}
			if err = d.Set(floatingIPZone, *floatingIP.Zone.Name); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-zone").GetDiag()
			}
			if err = d.Set("crn", floatingIP.CRN); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-crn").GetDiag()
			}

			if floatingIP.Target != nil {
				targetId, targetMap := dataSourceFloatingIPCollectionFloatingIpTargetToMap(floatingIP.Target)
				if err = d.Set(floatingIPTarget, targetId); err != nil { // We don't use targetID, it's not even useful in set
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-target").GetDiag()
				}
				targetList := []map[string]interface{}{}
				targetList = append(targetList, targetMap)
				if err = d.Set(floatingIPTargets, targetList); err != nil {
					return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target_list: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-target_list").GetDiag()
				}
			}

			tags, err := flex.GetGlobalTagsUsingCRN(meta, *floatingIP.CRN, "", isUserTagType)
			if err != nil {
				log.Printf("Error on get of vpc Floating IP (%s) tags: %s", *floatingIP.Address, err)
			}
			if err = d.Set(floatingIPTags, tags); err != nil { // Use d.Set and check error
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-tags").GetDiag()
			}

			accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *floatingIP.CRN, "", isAccessTagType)
			if err != nil {
				log.Printf(
					"Error on get of resource floating ip (%s) access tags: %s", d.Id(), err)
			}
			if err = d.Set(isFloatingIPAccessTags, accesstags); err != nil { // Use d.Set and check error
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_ibm_is_floating_ip", "read", "set-access_tags").GetDiag()
			}
			d.SetId(*floatingIP.ID)

			return nil
		}
	}

	err = fmt.Errorf("No floatingIP found with name %s", name)
	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_floating_ip", "read", "not-found").GetDiag()

}

func dataSourceFloatingIPCollectionFloatingIpTargetToMap(targetItemIntf vpcv1.FloatingIPTargetIntf) (targetId string, targetMap map[string]interface{}) {
	targetMap = map[string]interface{}{}
	targetId = ""
	switch reflect.TypeOf(targetItemIntf).String() {
	case "*vpcv1.FloatingIPTargetNetworkInterfaceReference":
		{
			targetItem := targetItemIntf.(*vpcv1.FloatingIPTargetNetworkInterfaceReference)
			targetId = *targetItem.ID
			if targetItem.Deleted != nil {
				deletedList := []map[string]interface{}{}
				deletedMap := dataSourceFloatingIPTargetNicDeletedToMap(*targetItem.Deleted)
				deletedList = append(deletedList, deletedMap)
				targetMap[floatingIPTargetsDeleted] = deletedList
			}
			if targetItem.Href != nil {
				targetMap[floatingIPTargetsHref] = targetItem.Href
			}
			if targetItem.ID != nil {
				targetMap[floatingIPTargetsId] = targetItem.ID
			}
			if targetItem.Name != nil {
				targetMap[floatingIPTargetsName] = targetItem.Name
			}
			if targetItem.PrimaryIP != nil {
				primaryIpList := make([]map[string]interface{}, 0)
				currentIP := map[string]interface{}{}
				if targetItem.PrimaryIP.Address != nil {
					currentIP[floatingIpPrimaryIpAddress] = *targetItem.PrimaryIP.Address
				}
				if targetItem.PrimaryIP.Href != nil {
					currentIP[floatingIpPrimaryIpHref] = *targetItem.PrimaryIP.Href
				}
				if targetItem.PrimaryIP.Name != nil {
					currentIP[floatingIpPrimaryIpName] = *targetItem.PrimaryIP.Name
				}
				if targetItem.PrimaryIP.ID != nil {
					currentIP[floatingIpPrimaryIpId] = *targetItem.PrimaryIP.ID
				}
				if targetItem.PrimaryIP.ResourceType != nil {
					currentIP[floatingIpPrimaryIpResourceType] = *targetItem.PrimaryIP.ResourceType
				}
				primaryIpList = append(primaryIpList, currentIP)
				targetMap[floatingIpPrimaryIP] = primaryIpList
			}
			if targetItem.ResourceType != nil {
				targetMap[floatingIPTargetsResourceType] = targetItem.ResourceType
			}
		}
	case "*vpcv1.FloatingIPTargetPublicGatewayReference":
		{
			targetItem := targetItemIntf.(*vpcv1.FloatingIPTargetPublicGatewayReference)
			targetId = *targetItem.ID
			if targetItem.Deleted != nil {
				deletedList := []map[string]interface{}{}
				deletedMap := dataSourceFloatingIPTargetPgDeletedToMap(*targetItem.Deleted)
				deletedList = append(deletedList, deletedMap)
				targetMap[floatingIPTargetsDeleted] = deletedList
			}
			if targetItem.Href != nil {
				targetMap[floatingIPTargetsHref] = targetItem.Href
			}
			if targetItem.ID != nil {
				targetMap[floatingIPTargetsId] = targetItem.ID
			}
			if targetItem.Name != nil {
				targetMap[floatingIPTargetsName] = targetItem.Name
			}
			if targetItem.ResourceType != nil {
				targetMap[floatingIPTargetsResourceType] = targetItem.ResourceType
			}
			if targetItem.CRN != nil {
				targetMap[floatingIPTargetsCrn] = targetItem.CRN
			}
		}
	case "*vpcv1.FloatingIPTarget":
		{
			targetItem := targetItemIntf.(*vpcv1.FloatingIPTarget)
			targetId = *targetItem.ID
			if targetItem.Deleted != nil {
				deletedList := []map[string]interface{}{}
				deletedMap := dataSourceFloatingIPTargetNicDeletedToMap(*targetItem.Deleted)
				deletedList = append(deletedList, deletedMap)
				targetMap[floatingIPTargetsDeleted] = deletedList
			}
			if targetItem.Href != nil {
				targetMap[floatingIPTargetsHref] = targetItem.Href
			}
			if targetItem.ID != nil {
				targetMap[floatingIPTargetsId] = targetItem.ID
			}
			if targetItem.Name != nil {
				targetMap[floatingIPTargetsName] = targetItem.Name
			}
			if targetItem.PrimaryIP != nil && targetItem.PrimaryIP.Address != nil {
				primaryIpList := make([]map[string]interface{}, 0)
				currentIP := map[string]interface{}{}
				if targetItem.PrimaryIP.Address != nil {
					currentIP[floatingIpPrimaryIpAddress] = *targetItem.PrimaryIP.Address
				}
				if targetItem.PrimaryIP.Href != nil {
					currentIP[floatingIpPrimaryIpHref] = *targetItem.PrimaryIP.Href
				}
				if targetItem.PrimaryIP.Name != nil {
					currentIP[floatingIpPrimaryIpName] = *targetItem.PrimaryIP.Name
				}
				if targetItem.PrimaryIP.ID != nil {
					currentIP[floatingIpPrimaryIpId] = *targetItem.PrimaryIP.ID
				}
				if targetItem.PrimaryIP.ResourceType != nil {
					currentIP[floatingIpPrimaryIpResourceType] = *targetItem.PrimaryIP.ResourceType
				}
				primaryIpList = append(primaryIpList, currentIP)
				targetMap[floatingIpPrimaryIP] = primaryIpList
			}
			if targetItem.ResourceType != nil {
				targetMap[floatingIPTargetsResourceType] = targetItem.ResourceType
			}
			if targetItem.CRN != nil {
				targetMap[floatingIPTargetsCrn] = targetItem.CRN
			}
		}
	}

	return targetId, targetMap
}

func dataSourceFloatingIPTargetNicDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap[floatingIPTargetsMoreInfo] = deletedItem.MoreInfo
	}

	return deletedMap
}
func dataSourceFloatingIPTargetPgDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap[floatingIPTargetsMoreInfo] = deletedItem.MoreInfo
	}

	return deletedMap
}
