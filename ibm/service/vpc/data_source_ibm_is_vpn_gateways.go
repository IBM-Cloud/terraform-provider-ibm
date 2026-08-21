// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isvpnGateways            = "vpn_gateways"
	isVPNGatewayResourceType = "resource_type"
	isVPNGatewayCrn          = "crn"
)

func DataSourceIBMISVPNGateways() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMVPNGatewaysRead,

		Schema: map[string]*schema.Schema{
			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier of the resource group this vpn gateway belongs to",
			},
			"mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The mode of this vpn gateway.",
			},
			isvpnGateways: {
				Type:        schema.TypeList,
				Description: "Collection of VPN Gateways",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isVPNGatewayName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN Gateway instance name",
						},
						isVPNGatewayCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that this VPN gateway was created",
						},
						isVPNGatewayCrn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPN gateway's CRN",
						},
						// regional vpn
						"members": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The members for the VPN gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public IP address assigned to the VPN gateway member",
									},

									"private_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private IP address assigned to the VPN gateway member",
									},
									"role": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The high availability role assigned to the VPN gateway member",
									},

									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the VPN gateway member",
									},
									"health_reasons": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The reasons for the current `health_state` (if any).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A reason code for this health state:- `cannot_reserve_ip_address`: IP address exhaustion (release addresses on the VPN's  subnet)- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
												},
												"message": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "An explanation of the reason for this health state.",
												},
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A link to documentation about the reason for this health state.",
												},
											},
										},
									},
									"health_state": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The health of this resource:- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A resource with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` resource may also have this state.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this VPN gateway member.",
									},
									"lifecycle_reasons": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The reasons for the current `lifecycle_state` (if any).",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
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
										Description: "The lifecycle state of the VPN gateway member.",
									},
									"private_ip": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The reserved IP address assigned to the VPN gateway member.This property will be present only when the VPN gateway status is `available`.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
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
												"href": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URL for this reserved IP.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique identifier for this reserved IP.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name for this reserved IP. The name is unique across all reserved IPs in a subnet.",
												},
												"resource_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource type.",
												},
												"subnet": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"crn": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The CRN for this subnet.",
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
															"href": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The URL for this subnet.",
															},
															"id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The unique identifier for this subnet.",
															},
															"name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name for this subnet. The name is unique across all subnets in the VPC.",
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
										},
									},
									"public_ip": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The public IP address assigned to the VPN gateway member.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.",
												},
											},
										},
									},
								},
							},
						},

						"availability_mode": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The availability mode of the VPN gateway:- `zonal`: The availability of this VPN gateway is limited only to a single zone of a  given region as provided by the `zone` of the VPN gateway.",
						},
						isVPNGatewayResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},

						isVPNGatewayStatus: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN gateway",
						},
						isVPNGatewayHealthState: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The health of this resource.- `ok`: Healthy- `degraded`: Suffering from compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.",
						},
						isVPNGatewayHealthReasons: {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A snake case string succinctly identifying the reason for this health state.",
									},

									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this health state.",
									},

									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about the reason for this health state.",
									},
								},
							},
						},
						isVPNGatewayLifecycleState: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lifecycle state of the VPN route.",
						},
						isVPNGatewayLifecycleReasons: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current lifecycle_state (if any).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A snake case string succinctly identifying the reason for this lifecycle state.",
									},

									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this lifecycle state.",
									},

									"more_info": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about the reason for this lifecycle state.",
									},
								},
							},
						},
						isVPNGatewaySubnet: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPNGateway subnet info",
						},
						isVPNGatewayResourceGroup: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "resource group identifiers ",
						},
						isVPNGatewayMode: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: " VPN gateway mode(policy/route) ",
						},
						isVPNGatewayLocalAsn: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The local autonomous system number (ASN) for this VPN gateway and its connections.",
						},
						isVPNGatewayAdvertisedCidrs: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The additional CIDRs advertised through any enabled routing protocol (for example, BGP). The routing protocol will advertise routes with these CIDRs and VPC prefixes as route destinations.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vpc": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VPC for the VPN Gateway",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this VPC.",
									},
									"deleted": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this VPC.",
									},
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this VPC.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique user-defined name for this VPC.",
									},
								},
							},
						},
						isVPNGatewayTags: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "VPN Gateway tags list",
						},
						isVPNGatewayAccessTags: {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Set:         flex.ResourceIBMVPCHash,
							Description: "List of access management tags",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMVPNGatewaysRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpn_gateways", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listvpnGWOptions := sess.NewListVPNGatewaysOptions()
	if resgroupintf, ok := d.GetOk("resource_group"); ok {
		resGroup := resgroupintf.(string)
		listvpnGWOptions.ResourceGroupID = &resGroup
	}
	if modeIntf, ok := d.GetOk("mode"); ok {
		mode := modeIntf.(string)
		listvpnGWOptions.Mode = &mode
	}
	start := ""
	allrecs := []vpcv1.VPNGatewayIntf{}
	for {
		if start != "" {
			listvpnGWOptions.Start = &start
		}
		availableVPNGateways, _, err := sess.ListVPNGatewaysWithContext(context, listvpnGWOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNGatewaysWithContext failed %s", err), "(Data) ibm_is_vpn_gateways", "read")
			log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		start = flex.GetNext(availableVPNGateways.Next)
		allrecs = append(allrecs, availableVPNGateways.VPNGateways...)
		if start == "" {
			break
		}
	}

	vpngateways := make([]map[string]interface{}, 0)
	for _, instance := range allrecs {
		gateway := map[string]interface{}{}
		data := instance.(*vpcv1.VPNGateway)
		gateway[isVPNGatewayName] = *data.Name
		gateway[isVPNGatewayCreatedAt] = data.CreatedAt.String()
		gateway[isVPNGatewayResourceType] = *data.ResourceType
		gateway[isVPNGatewayHealthState] = *data.HealthState
		gateway[isVPNGatewayHealthReasons] = resourceVPNGatewayRouteFlattenHealthReasons(data.HealthReasons)
		gateway[isVPNGatewayLifecycleState] = *data.LifecycleState
		gateway[isVPNGatewayLifecycleReasons] = resourceVPNGatewayFlattenLifecycleReasons(data.LifecycleReasons)
		gateway[isVPNGatewayMode] = *data.Mode
		if data.LocalAsn != nil {
			gateway[isVPNGatewayLocalAsn] = *data.LocalAsn
		}
		// regional vpn
		gateway["availability_mode"] = *data.AvailabilityMode

		members := []map[string]interface{}{}
		for _, membersItem := range data.Members {
			membersItemMap, err := DataSourceIBMIsVPNGatewaysVPNGatewayMemberToMap(&membersItem) // #nosec G601
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DataSourceIBMIsVPNGatewaysVPNGatewayMemberToMap failed %s", err), "(Data) ibm_is_vpn_gateways", "read")
				log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			members = append(members, membersItemMap)
		}
		gateway["members"] = members
		if data.AdvertisedCIDRs != nil {
			gateway[isVPNGatewayAdvertisedCidrs] = data.AdvertisedCIDRs
		}
		gateway[isVPNGatewayResourceGroup] = *data.ResourceGroup.ID
		gateway[isVPNGatewaySubnet] = *data.Subnet.ID
		gateway[isVPNGatewayCrn] = *data.CRN
		tags, err := flex.GetGlobalTagsUsingCRN(meta, *data.CRN, "", isUserTagType)
		if err != nil {
			log.Printf(
				"Error on get of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
		}
		gateway[isVPNGatewayTags] = tags

		accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *data.CRN, "", isAccessTagType)
		if err != nil {
			log.Printf(
				"Error on get of resource VPC VPN Gateway (%s) access tags: %s", d.Id(), err)
		}
		gateway[isVPNGatewayAccessTags] = accesstags
		// if data.Members != nil {
		// 	vpcMembersIpsList := make([]map[string]interface{}, 0)
		// 	for _, memberIP := range data.Members {
		// 		currentMemberIP := map[string]interface{}{}
		// 		if memberIP.PublicIP != nil {
		// 			currentMemberIP["address"] = *memberIP.PublicIP.Address
		// 			currentMemberIP["role"] = *memberIP.Role
		// 			vpcMembersIpsList = append(vpcMembersIpsList, currentMemberIP)
		// 		}
		// 		if memberIP.PrivateIP != nil && memberIP.PrivateIP.Address != nil {
		// 			currentMemberIP["private_address"] = *memberIP.PrivateIP.Address
		// 		}
		// 	}
		// 	gateway[isVPNGatewayMembers] = vpcMembersIpsList
		// }

		if data.VPC != nil {
			vpcList := []map[string]interface{}{}
			vpcList = append(vpcList, dataSourceVPNServerCollectionVPNGatewayVpcReferenceToMap(data.VPC))
			gateway["vpc"] = vpcList
		}

		vpngateways = append(vpngateways, gateway)
	}

	d.SetId(dataSourceIBMVPNGatewaysID(d))
	if err = d.Set("vpn_gateways", vpngateways); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpn_gateways %s", err), "(Data) ibm_is_vpn_gateways", "read", "vpn_gateways-set").GetDiag()
	}
	return nil
}

// dataSourceIBMVPNGatewaysID returns a reasonable ID  list.
func dataSourceIBMVPNGatewaysID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceVPNServerCollectionVPNGatewayVpcReferenceToMap(vpcsItem *vpcv1.VPCReference) (vpcsMap map[string]interface{}) {
	vpcsMap = map[string]interface{}{}

	if vpcsItem.CRN != nil {
		vpcsMap["crn"] = vpcsItem.CRN
	}
	if vpcsItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVPNGatewayCollectionVpcsDeletedToMap(*vpcsItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		vpcsMap["deleted"] = deletedList
	}
	if vpcsItem.Href != nil {
		vpcsMap["href"] = vpcsItem.Href
	}
	if vpcsItem.ID != nil {
		vpcsMap["id"] = vpcsItem.ID
	}
	if vpcsItem.Name != nil {
		vpcsMap["name"] = vpcsItem.Name
	}

	return vpcsMap
}

func dataSourceVPNGatewayCollectionVpcsDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func DataSourceIBMIsVPNGatewaysVPNGatewayMemberToMap(model *vpcv1.VPNGatewayMember) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	healthReasons := []map[string]interface{}{}
	for _, healthReasonsItem := range model.HealthReasons {
		healthReasonsItemMap, err := DataSourceIBMIsVPNGatewaysVPNGatewayMemberHealthReasonToMap(&healthReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		healthReasons = append(healthReasons, healthReasonsItemMap)
	}
	modelMap["health_reasons"] = healthReasons
	modelMap["health_state"] = *model.HealthState
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsVPNGatewaysVPNGatewayMemberLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	modelMap["lifecycle_state"] = *model.LifecycleState
	privateIPMap, err := DataSourceIBMIsVPNGatewaysReservedIPReferenceVPNGatewayContextToMap(model.PrivateIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["private_ip"] = []map[string]interface{}{privateIPMap}
	publicIPMap, err := DataSourceIBMIsVPNGatewaysIPToMap(model.PublicIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["public_ip"] = []map[string]interface{}{publicIPMap}
	modelMap["role"] = *model.Role
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewaysVPNGatewayMemberHealthReasonToMap(model *vpcv1.VPNGatewayMemberHealthReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewaysVPNGatewayMemberLifecycleReasonToMap(model *vpcv1.VPNGatewayMemberLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewaysReservedIPReferenceVPNGatewayContextToMap(model *vpcv1.ReservedIPReferenceVPNGatewayMemberContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsVPNGatewaysDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	subnetMap, err := DataSourceIBMIsVPNGatewaysSubnetReferenceToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewaysSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsVPNGatewaysDeletedToMap(model.Deleted)
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

func DataSourceIBMIsVPNGatewaysIPToMap(model *vpcv1.IP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}
func DataSourceIBMIsVPNGatewaysDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}
