// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMISVPNGateway() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNGatewayRead,

		Schema: map[string]*schema.Schema{
			isVPNGatewayID: {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"vpn_gateway_name", isVPNGatewayID},
				Description:  "The VPN gateway identifier.",
			},
			"vpn_gateway_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"vpn_gateway_name", isVPNGatewayID},
				Description:  "The VPN gateway name.",
			},
			"connections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Connections for this VPN gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
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
							Description: "The VPN connection's canonical URL.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPN gateway connection.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this VPN connection.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this VPN gateway was created.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The VPN gateway's CRN.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The VPN gateway's canonical URL.",
			},
			"members": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The members for the VPN gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The private IP address assigned to the VPN gateway member. This property will be present only when the VPN gateway status is`available`.",
						},
						"public_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public IP address assigned to the VPN gateway member.",
						},
						"role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The high availability role assigned to the VPN gateway member.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN gateway member.",
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
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user-defined name for this VPN gateway.",
			},
			"resource_group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this VPN gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group.",
						},
					},
				},
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the VPN gateway.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource.- `ok`: Healthy- `degraded`: Suffering from compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.",
			},
			"health_reasons": {
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
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the VPN route.",
			},
			"lifecycle_reasons": {
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
			"subnet": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this subnet.",
						},
						"deleted": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and provides some supplementary information.",
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
							Description: "The URL for this subnet.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this subnet.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this subnet.",
						},
					},
				},
			},
			"mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Route mode VPN gateway.",
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
	}
}

func dataSourceIBMIsVPNGatewayRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpn_gateway", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpn_gateway_name := d.Get("vpn_gateway_name").(string)
	vpn_gateway_id := d.Get("vpn_gateway").(string)
	vpnGateway := &vpcv1.VPNGateway{}
	if vpn_gateway_id != "" {
		getVPNGatewayOptions := &vpcv1.GetVPNGatewayOptions{}

		getVPNGatewayOptions.SetID(vpn_gateway_id)

		vpnGatewayIntf, _, err := vpcClient.GetVPNGatewayWithContext(context, getVPNGatewayOptions)
		if err != nil || vpnGatewayIntf.(*vpcv1.VPNGateway) == nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayWithContext failed: %s", err.Error()), "(Data) ibm_is_vpn_gateway", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		vpnGateway = vpnGatewayIntf.(*vpcv1.VPNGateway)
	} else {
		listvpnGWOptions := vpcClient.NewListVPNGatewaysOptions()

		start := ""
		allrecs := []vpcv1.VPNGatewayIntf{}
		for {
			if start != "" {
				listvpnGWOptions.Start = &start
			}
			availableVPNGateways, _, err := vpcClient.ListVPNGatewaysWithContext(context, listvpnGWOptions)
			if err != nil || availableVPNGateways == nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNGatewaysWithContext failed: %s", err.Error()), "(Data) ibm_is_vpn_gateway", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(availableVPNGateways.Next)
			allrecs = append(allrecs, availableVPNGateways.VPNGateways...)
			if start == "" {
				break
			}
		}
		vpn_gateway_found := false
		for _, vpnGatewayIntfItem := range allrecs {
			if *vpnGatewayIntfItem.(*vpcv1.VPNGateway).Name == vpn_gateway_name {
				vpnGateway = vpnGatewayIntfItem.(*vpcv1.VPNGateway)
				vpn_gateway_found = true
				break
			}
		}
		if !vpn_gateway_found {
			err = fmt.Errorf("No vpn gateway found with given name %s", vpn_gateway_name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNGatewaysWithContext failed: %s", err.Error()), "(Data) ibm_is_vpn_gateway", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	d.SetId(*vpnGateway.ID)

	if vpnGateway.Connections != nil {
		err = d.Set("connections", dataSourceVPNGatewayFlattenConnections(vpnGateway.Connections))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connections: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-connections").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(vpnGateway.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("crn", vpnGateway.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-crn").GetDiag()
	}
	if err = d.Set("href", vpnGateway.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-href").GetDiag()
	}

	// if vpnGateway.Members != nil {
	// 	err = d.Set("members", dataSourceVPNGatewayFlattenMembers(vpnGateway.Members))
	// 	if err != nil {
	// 		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting members: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-members").GetDiag()
	// 	}
	// }
	// regional vpn

	if !core.IsNil(vpnGateway.AvailabilityMode) {
		if err = d.Set("availability_mode", vpnGateway.AvailabilityMode); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting availability_mode: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-availability_mode").GetDiag()
		}
	}
	members := []map[string]interface{}{}
	for _, membersItem := range vpnGateway.Members {
		membersItemMap, err := DataSourceIBMIsVPNGatewayVPNGatewayMemberToMap(&membersItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpn_gateway", "read", "members-to-map").GetDiag()
		}
		members = append(members, membersItemMap)
	}
	if err = d.Set("members", members); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting members: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-members").GetDiag()
	}
	if err = d.Set("name", vpnGateway.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-name").GetDiag()
	}

	if vpnGateway.ResourceGroup != nil {
		err = d.Set("resource_group", dataSourceVPNGatewayFlattenResourceGroup(*vpnGateway.ResourceGroup))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-resource_group").GetDiag()
		}
	}
	if err = d.Set("resource_type", vpnGateway.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("health_state", vpnGateway.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_state: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-health_state").GetDiag()
	}
	if err := d.Set("health_reasons", resourceVPNGatewayRouteFlattenHealthReasons(vpnGateway.HealthReasons)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_reasons: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-health_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", vpnGateway.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-lifecycle_state").GetDiag()
	}
	if err := d.Set("lifecycle_reasons", resourceVPNGatewayFlattenLifecycleReasons(vpnGateway.LifecycleReasons)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-lifecycle_reasons").GetDiag()
	}
	if vpnGateway.Subnet != nil {
		err = d.Set("subnet", dataSourceVPNGatewayFlattenSubnet(*vpnGateway.Subnet))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnet: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-subnet").GetDiag()
		}
	}
	if err = d.Set("mode", vpnGateway.Mode); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mode: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-mode").GetDiag()
	}
	if vpnGateway.AdvertisedCIDRs != nil {
		if err = d.Set("advertised_cidrs", vpnGateway.AdvertisedCIDRs); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting advertised_cidrs: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-advertised_cidrs").GetDiag()
		}
	}
	if vpnGateway.LocalAsn != nil {
		if err = d.Set("local_asn", vpnGateway.LocalAsn); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting local_asn: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-local_asn").GetDiag()
		}
	}
	if vpnGateway.VPC != nil {
		err = d.Set("vpc", dataSourceVPNGatewayFlattenVPC(vpnGateway.VPC))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpc: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-vpc").GetDiag()
		}
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *vpnGateway.CRN, "", isUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource vpc VPN Gateway (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(isVPNGatewayTags, tags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-tags").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *vpnGateway.CRN, "", isAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource VPC VPN Gateway (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(isVPNGatewayAccessTags, accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_vpn_gateway", "read", "set-access_tags").GetDiag()
	}
	return nil
}

func dataSourceVPNGatewayFlattenConnections(result []vpcv1.VPNGatewayConnectionReference) (connections []map[string]interface{}) {
	for _, connectionsItem := range result {
		connections = append(connections, dataSourceVPNGatewayConnectionsToMap(connectionsItem))
	}

	return connections
}

func dataSourceVPNGatewayConnectionsToMap(connectionsItem vpcv1.VPNGatewayConnectionReference) (connectionsMap map[string]interface{}) {
	connectionsMap = map[string]interface{}{}

	if connectionsItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVPNGatewayConnectionsDeletedToMap(*connectionsItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		connectionsMap["deleted"] = deletedList
	}
	if connectionsItem.Href != nil {
		connectionsMap["href"] = connectionsItem.Href
	}
	if connectionsItem.ID != nil {
		connectionsMap["id"] = connectionsItem.ID
	}
	if connectionsItem.Name != nil {
		connectionsMap["name"] = connectionsItem.Name
	}
	if connectionsItem.ResourceType != nil {
		connectionsMap["resource_type"] = connectionsItem.ResourceType
	}

	return connectionsMap
}

func dataSourceVPNGatewayConnectionsDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceVPNGatewayFlattenMembers(result []vpcv1.VPNGatewayMember) (members []map[string]interface{}) {
	for _, membersItem := range result {
		members = append(members, dataSourceVPNGatewayMembersToMap(membersItem))
	}

	return members
}

func dataSourceVPNGatewayMembersToMap(membersItem vpcv1.VPNGatewayMember) (membersMap map[string]interface{}) {
	membersMap = map[string]interface{}{}

	if membersItem.PrivateIP != nil && membersItem.PrivateIP.Address != nil {
		membersMap["private_ip_address"] = membersItem.PrivateIP.Address
	}
	if membersItem.PublicIP != nil {
		membersMap["public_ip_address"] = membersItem.PublicIP.Address
	}
	if membersItem.Role != nil {
		membersMap["role"] = membersItem.Role
	}

	return membersMap
}

func dataSourceVPNGatewayFlattenResourceGroup(result vpcv1.ResourceGroupReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVPNGatewayResourceGroupToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVPNGatewayResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap["href"] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap["id"] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap["name"] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func dataSourceVPNGatewayFlattenSubnet(result vpcv1.SubnetReference) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceVPNGatewaySubnetToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceVPNGatewaySubnetToMap(subnetItem vpcv1.SubnetReference) (subnetMap map[string]interface{}) {
	subnetMap = map[string]interface{}{}

	if subnetItem.CRN != nil {
		subnetMap["crn"] = subnetItem.CRN
	}
	if subnetItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVPNGatewaySubnetDeletedToMap(*subnetItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		subnetMap["deleted"] = deletedList
	}
	if subnetItem.Href != nil {
		subnetMap["href"] = subnetItem.Href
	}
	if subnetItem.ID != nil {
		subnetMap["id"] = subnetItem.ID
	}
	if subnetItem.Name != nil {
		subnetMap["name"] = subnetItem.Name
	}

	return subnetMap
}

func dataSourceVPNGatewaySubnetDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func dataSourceVPNGatewayFlattenVPC(result *vpcv1.VPCReference) (vpcs []map[string]interface{}) {
	vpcs = append(vpcs, dataSourceVPNGatewayVpcToMap(*result))
	return vpcs
}

func dataSourceVPNGatewayVpcToMap(vpcItem vpcv1.VPCReference) (vpcsMap map[string]interface{}) {
	vpcsMap = map[string]interface{}{}

	if vpcItem.CRN != nil {
		vpcsMap["crn"] = vpcItem.CRN
	}
	if vpcItem.Deleted != nil {
		deletedList := []map[string]interface{}{}
		deletedMap := dataSourceVPNGatewayVpcDeletedToMap(*vpcItem.Deleted)
		deletedList = append(deletedList, deletedMap)
		vpcsMap["deleted"] = deletedList
	}
	if vpcItem.Href != nil {
		vpcsMap["href"] = vpcItem.Href
	}
	if vpcItem.ID != nil {
		vpcsMap["id"] = vpcItem.ID
	}
	if vpcItem.Name != nil {
		vpcsMap["name"] = vpcItem.Name
	}

	return vpcsMap

}

func dataSourceVPNGatewayVpcDeletedToMap(deletedItem vpcv1.Deleted) (deletedMap map[string]interface{}) {
	deletedMap = map[string]interface{}{}

	if deletedItem.MoreInfo != nil {
		deletedMap["more_info"] = deletedItem.MoreInfo
	}

	return deletedMap
}

func DataSourceIBMIsVPNGatewayVPNGatewayMemberToMap(model *vpcv1.VPNGatewayMember) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	healthReasons := []map[string]interface{}{}
	for _, healthReasonsItem := range model.HealthReasons {
		healthReasonsItemMap, err := DataSourceIBMIsVPNGatewayVPNGatewayMemberHealthReasonToMap(&healthReasonsItem) // #nosec G601
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
		lifecycleReasonsItemMap, err := DataSourceIBMIsVPNGatewayVPNGatewayMemberLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	modelMap["lifecycle_state"] = *model.LifecycleState
	privateIPMap, err := DataSourceIBMIsVPNGatewayReservedIPReferenceVPNGatewayContextToMap(model.PrivateIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["private_ip"] = []map[string]interface{}{privateIPMap}
	publicIPMap, err := DataSourceIBMIsVPNGatewayIPToMap(model.PublicIP)
	if err != nil {
		return modelMap, err
	}
	modelMap["public_ip"] = []map[string]interface{}{publicIPMap}
	modelMap["role"] = *model.Role
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewayVPNGatewayMemberHealthReasonToMap(model *vpcv1.VPNGatewayMemberHealthReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewayVPNGatewayMemberLifecycleReasonToMap(model *vpcv1.VPNGatewayMemberLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewayReservedIPReferenceVPNGatewayContextToMap(model *vpcv1.ReservedIPReferenceVPNGatewayMemberContext) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsVPNGatewayDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	subnetMap, err := DataSourceIBMIsVPNGatewaySubnetReferenceToMap(model.Subnet)
	if err != nil {
		return modelMap, err
	}
	modelMap["subnet"] = []map[string]interface{}{subnetMap}
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewaySubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsVPNGatewayDeletedToMap(model.Deleted)
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

func DataSourceIBMIsVPNGatewayIPToMap(model *vpcv1.IP) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["address"] = *model.Address
	return modelMap, nil
}

func DataSourceIBMIsVPNGatewayDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}
