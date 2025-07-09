// Copyright IBM Corp. 2021, 2022 All Rights Reserved.
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

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	rtRoutes = "routes"
	rtCrn    = "crn"
)

func DataSourceIBMIsVPCRoutingTable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPCRoutingTableRead,

		Schema: map[string]*schema.Schema{
			isVpcID: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPC identifier.",
			},
			rName: &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				AtLeastOneOf:  []string{rName, isRoutingTableID},
				ConflictsWith: []string{isRoutingTableID},
				Description:   "The user-defined name for this routing table.",
			},
			isRoutingTableAcceptRoutesFrom: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The filters specifying the resources that may create routes in this routing table.At present, only the `resource_type` filter is permitted, and only the `vpn_gateway` value is supported, but filter support is expected to expand in the future.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			isRoutingTableID: &schema.Schema{
				Type:          schema.TypeString,
				AtLeastOneOf:  []string{rName, isRoutingTableID},
				ConflictsWith: []string{rName},
				Optional:      true,
				Description:   "The routing table identifier.",
			},
			rtCrn: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The routing table CRN.",
			},
			"advertise_routes_to": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The ingress sources to advertise routes to. Routes in the table with `advertise` enabled will be advertised to these sources.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			rtCreateAt: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this routing table was created.",
			},
			rtHref: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this routing table.",
			},
			rtIsDefault: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this is the default routing table for this VPC.",
			},
			rtLifecycleState: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the routing table.",
			},
			rtResourceType: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			rtRouteDirectLinkIngress: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this routing table is used to route traffic that originates from[Direct Link](https://cloud.ibm.com/docs/dl/) to this VPC.Incoming traffic will be routed according to the routing table with one exception: routes with an `action` of `deliver` are treated as `drop` unless the `next_hop` is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a `next_hop` of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.",
			},
			rtRouteInternetIngress: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this routing table is used to route traffic that originates from the internet.Incoming traffic will be routed according to the routing table with two exceptions:- Traffic destined for IP addresses associated with public gateways will not be  subject to routes in this routing table.- Routes with an action of deliver are treated as drop unless the `next_hop` is an  IP address bound to a network interface on a subnet in the route's `zone`.  Therefore, if an incoming packet matches a route with a `next_hop` of an  internet-bound IP address or a VPN gateway connection, the packet will be dropped.",
			},
			rtRouteTransitGatewayIngress: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this routing table is used to route traffic that originates from from [Transit Gateway](https://cloud.ibm.com/cloud/transit-gateway/) to this VPC.Incoming traffic will be routed according to the routing table with one exception: routes with an `action` of `deliver` are treated as `drop` unless the `next_hop` is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a `next_hop` of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.",
			},
			rtRouteVPCZoneIngress: &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this routing table is used to route traffic that originates from subnets in other zones in this VPC.Incoming traffic will be routed according to the routing table with one exception: routes with an `action` of `deliver` are treated as `drop` unless the `next_hop` is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a `next_hop` of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.",
			},
			rtRoutes: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routes for this routing table.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						rDeleted: &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									rMoreInfo: &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						rtHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this route.",
						},
						rId: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this route.",
						},
						rName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this route.",
						},
					},
				},
			},
			rtSubnets: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The subnets to which this routing table is attached.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						rtCrn: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this subnet.",
						},
						rDeleted: &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									rMoreInfo: &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						rtHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this subnet.",
						},
						rId: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this subnet.",
						},
						rName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this subnet.",
						},
					},
				},
			},
			rtResourceGroup: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resource group for this volume.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						rtResourceGroupHref: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						rtResourceGroupId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this resource group.",
						},
						rtResourceGroupName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this resource group.",
						},
					},
				},
			},
			rtTags: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      flex.ResourceIBMVPCHash,
			},
			rtAccessTags: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access tags",
			},
		},
	}
}

func dataSourceIBMIsVPCRoutingTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpc_routing_table", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := d.Get(isVpcID).(string)
	rtId := d.Get(isRoutingTableID).(string)
	routingTableName := d.Get(rName).(string)
	var routingTable *vpcv1.RoutingTable
	if rtId != "" {
		getVPCRoutingTableOptions := &vpcv1.GetVPCRoutingTableOptions{
			VPCID: &vpcID,
			ID:    &rtId,
		}

		rt, _, err := vpcClient.GetVPCRoutingTableWithContext(ctx, getVPCRoutingTableOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCRoutingTableWithContext failed: %s", err.Error()), "(Data) ibm_is_vpc_routing_table", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		routingTable = rt
	} else {
		start := ""
		allrecs := []vpcv1.RoutingTable{}
		for {
			listOptions := &vpcv1.ListVPCRoutingTablesOptions{
				VPCID: &vpcID,
			}
			if start != "" {
				listOptions.Start = &start
			}
			result, _, err := vpcClient.ListVPCRoutingTablesWithContext(ctx, listOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPCRoutingTablesWithContext failed: %s", err.Error()), "(Data) ibm_is_vpc_routing_table", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(result.Next)
			allrecs = append(allrecs, result.RoutingTables...)
			if start == "" {
				break
			}
		}
		for _, r := range allrecs {
			if *r.Name == routingTableName {
				routingTable = &r
			}
		}
		if routingTable == nil {
			err = fmt.Errorf("Provided routing table %s cannot be found in the vpc %s", routingTableName, vpcID)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPCRoutingTablesWithContext failed: %s", err.Error()), "(Data) ibm_is_vpc_routing_table", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId(*routingTable.ID)

	if err = d.Set("created_at", flex.DateTimeToString(routingTable.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-created_at").GetDiag()
	}
	acceptRoutesFromInfo := make([]map[string]interface{}, 0)
	if routingTable.AcceptRoutesFrom != nil {
		for _, AcceptRoutesFrom := range routingTable.AcceptRoutesFrom {
			l := map[string]interface{}{}
			if AcceptRoutesFrom.ResourceType != nil {
				l["resource_type"] = *AcceptRoutesFrom.ResourceType
				acceptRoutesFromInfo = append(acceptRoutesFromInfo, l)
			}
		}
	}
	if err = d.Set(isRoutingTableAcceptRoutesFrom, acceptRoutesFromInfo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting accept_routes_from: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-accept_routes_from").GetDiag()
	}

	if err = d.Set(isRoutingTableID, routingTable.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting routing_table: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-routing_table").GetDiag()
	}

	if err = d.Set(rtHref, routingTable.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-href").GetDiag()
	}

	if err = d.Set("is_default", routingTable.IsDefault); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_default: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-is_default").GetDiag()
	}

	if err = d.Set("lifecycle_state", routingTable.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("name", routingTable.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-name").GetDiag()
	}

	if err = d.Set("resource_type", routingTable.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("route_direct_link_ingress", routingTable.RouteDirectLinkIngress); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_direct_link_ingress: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-route_direct_link_ingress").GetDiag()
	}

	if err = d.Set("route_internet_ingress", routingTable.RouteInternetIngress); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_internet_ingress: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-route_internet_ingress").GetDiag()
	}
	if err = d.Set("route_transit_gateway_ingress", routingTable.RouteTransitGatewayIngress); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_transit_gateway_ingress: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-route_transit_gateway_ingress").GetDiag()
	}

	if err = d.Set("route_vpc_zone_ingress", routingTable.RouteVPCZoneIngress); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_vpc_zone_ingress: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-route_vpc_zone_ingress").GetDiag()
	}
	if err = d.Set("advertise_routes_to", routingTable.AdvertiseRoutesTo); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting advertise_routes_to: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-advertise_routes_to").GetDiag()
	}

	if err = d.Set(rtCrn, routingTable.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-crn").GetDiag()
	}

	routes := []map[string]interface{}{}
	if routingTable.Routes != nil {
		for _, modelItem := range routingTable.Routes {
			modelMap, err := dataSourceIBMIBMIsVPCRoutingTableRouteReferenceToMap(&modelItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpc_routing_table", "read", "routes-to-map").GetDiag()
			}
			routes = append(routes, modelMap)
		}
	}
	if err = d.Set(rtRoutes, routes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting routes: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-routes").GetDiag()
	}

	subnets := []map[string]interface{}{}
	if routingTable.Subnets != nil {
		for _, modelItem := range routingTable.Subnets {
			modelMap, err := dataSourceIBMIBMIsVPCRoutingTableSubnetReferenceToMap(&modelItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpc_routing_table", "read", "subnets-to-map").GetDiag()
			}
			subnets = append(subnets, modelMap)
		}
	}
	if err = d.Set(rtSubnets, subnets); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting subnets: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-subnets").GetDiag()
	}

	resourceGroupList := []map[string]interface{}{}
	if routingTable.ResourceGroup != nil {
		resourceGroupMap := routingTableResourceGroupToMap(*routingTable.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
	}
	if err = d.Set(rtResourceGroup, resourceGroupList); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-resource_group").GetDiag()
	}

	tags, err := flex.GetGlobalTagsUsingCRN(meta, *routingTable.CRN, "", rtUserTagType)
	if err != nil {
		log.Printf(
			"An error occured during reading of routing table (%s) tags : %s", d.Id(), err)
	}
	if err = d.Set(rtTags, tags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-tags").GetDiag()
	}

	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *routingTable.CRN, "", rtAccessTagType)
	if err != nil {
		log.Printf(
			"An error occured reading access tags for routing table (%s) : %s", d.Id(), err)
	}
	if err = d.Set(rtAccessTags, accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "(Data) ibm_is_vpc_routing_table", "read", "set-access_tags").GetDiag()
	}

	return nil
}

func dataSourceIBMIBMIsVPCRoutingTableRouteReferenceToMap(model *vpcv1.RouteReference) (map[string]interface{}, error) {
	modelMap := map[string]interface{}{}
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIBMIsVPCRoutingTableRouteReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap[rDeleted] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap[rtHref] = *model.Href
	}
	if model.ID != nil {
		modelMap[rId] = *model.ID
	}
	if model.Name != nil {
		modelMap[rName] = *model.Name
	}
	return modelMap, nil
}

func dataSourceIBMIBMIsVPCRoutingTableRouteReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := map[string]interface{}{}
	if model.MoreInfo != nil {
		modelMap[rMoreInfo] = *model.MoreInfo
	}
	return modelMap, nil
}

func dataSourceIBMIBMIsVPCRoutingTableSubnetReferenceToMap(model *vpcv1.SubnetReference) (map[string]interface{}, error) {
	modelMap := map[string]interface{}{}
	if model.CRN != nil {
		modelMap[rtCrn] = *model.CRN
	}
	if model.Deleted != nil {
		deletedMap, err := dataSourceIBMIBMIsVPCRoutingTableSubnetReferenceDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap[rDeleted] = []map[string]interface{}{deletedMap}
	}
	if model.Href != nil {
		modelMap[rtHref] = *model.Href
	}
	if model.ID != nil {
		modelMap[rId] = *model.ID
	}
	if model.Name != nil {
		modelMap[rName] = *model.Name
	}
	return modelMap, nil
}

func dataSourceIBMIBMIsVPCRoutingTableSubnetReferenceDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := map[string]interface{}{}
	if model.MoreInfo != nil {
		modelMap[rMoreInfo] = *model.MoreInfo
	}
	return modelMap, nil
}
