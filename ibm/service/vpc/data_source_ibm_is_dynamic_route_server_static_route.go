// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerStaticRoute() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerStaticRouteRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"is_dynamic_route_server_static_route_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server static route identifier.",
			},
			"added_routes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routes added to the VPC routing tables for this dynamic route server static route.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_routing_table_route": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The route reference for the applied route, If absent the routecould not be added to the routing table.For more information, see[Dynamic Route Server Static Route Failure](https://cloud.ibm.com/docs/__TBD__).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
										Description: "The URL for this route.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this route.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this route. The name is unique across all routes in the routing table.",
									},
								},
							},
						},
					},
				},
			},
			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The destination CIDR of the route. The host identifier in the CIDR must be zero.At most two routes per `zone` in a VPC routing table can have the same `destination` and `priority`, and only if both routes have an `action` of `deliver` and the`next_hop` is an IP address.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server static routes.",
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
				Description: "The lifecycle state of the static route.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this static route. The name must not be used by another static route for this dynamic route server. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"next_hops": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The next hop resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Description: "The URL for this dynamic route server peer group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID for this dynamic route server peer group peer.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this dynamic route server peer group.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"route_delete_delay": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of seconds to wait before deleting a route from `routing_tables` when the `status` of the route's associated peer is `down`.",
			},
			"routing_tables": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routing tables to update.A route is added to each routing table for each peer in the `next_hop` resource. Each route uses its associated peer's `endpoint.address` for the route's `next_hop`, and:- For ingress routing tables, the route's `priority` uses the peer's `priority`.- For egress routing tables, the route's `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.",
						},
						"vpc_routing_table": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this VPC routing table.",
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
										Description: "The URL for this routing table.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this routing table.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this routing table. The name is unique across all routing tables for the VPC.",
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
		},
	}
}

func dataSourceIBMIsDynamicRouteServerStaticRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_static_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerStaticRouteOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}

	getDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	getDynamicRouteServerStaticRouteOptions.SetID(d.Get("is_dynamic_route_server_static_route_id").(string))

	dynamicRouteServerStaticRoute, _, err := vpcClient.GetDynamicRouteServerStaticRouteWithContext(context, getDynamicRouteServerStaticRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_static_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getDynamicRouteServerStaticRouteOptions.DynamicRouteServerID, *getDynamicRouteServerStaticRouteOptions.ID))

	addedRoutes := []map[string]interface{}{}
	for _, addedRoutesItem := range dynamicRouteServerStaticRoute.AddedRoutes {
		addedRoutesItemMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&addedRoutesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_static_route", "read", "added_routes-to-map").GetDiag()
		}
		addedRoutes = append(addedRoutes, addedRoutesItemMap)
	}
	if err = d.Set("added_routes", addedRoutes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting added_routes: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-added_routes").GetDiag()
	}

	if err = d.Set("destination", dynamicRouteServerStaticRoute.Destination); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-destination").GetDiag()
	}

	if err = d.Set("href", dynamicRouteServerStaticRoute.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-href").GetDiag()
	}

	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range dynamicRouteServerStaticRoute.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_static_route", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-lifecycle_reasons").GetDiag()
	}

	if !core.IsNil(dynamicRouteServerStaticRoute.LifecycleState) {
		if err = d.Set("lifecycle_state", dynamicRouteServerStaticRoute.LifecycleState); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-lifecycle_state").GetDiag()
		}
	}

	if err = d.Set("name", dynamicRouteServerStaticRoute.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-name").GetDiag()
	}

	nextHops := []map[string]interface{}{}
	for _, nextHopsItem := range dynamicRouteServerStaticRoute.NextHops {
		nextHopsItemMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(nextHopsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_static_route", "read", "next_hops-to-map").GetDiag()
		}
		nextHops = append(nextHops, nextHopsItemMap)
	}
	if err = d.Set("next_hops", nextHops); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting next_hops: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-next_hops").GetDiag()
	}

	if err = d.Set("resource_type", dynamicRouteServerStaticRoute.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-resource_type").GetDiag()
	}

	if err = d.Set("route_delete_delay", flex.IntValue(dynamicRouteServerStaticRoute.RouteDeleteDelay)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_delete_delay: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-route_delete_delay").GetDiag()
	}

	routingTables := []map[string]interface{}{}
	for _, routingTablesItem := range dynamicRouteServerStaticRoute.RoutingTables {
		routingTablesItemMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&routingTablesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_static_route", "read", "routing_tables-to-map").GetDiag()
		}
		routingTables = append(routingTables, routingTablesItemMap)
	}
	if err = d.Set("routing_tables", routingTables); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting routing_tables: %s", err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-routing_tables").GetDiag()
	}

	return nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model *vpcv1.DynamicRouteServerStaticRouteAddedRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VPCRoutingTableRoute != nil {
		vpcRoutingTableRouteMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model.VPCRoutingTableRoute)
		if err != nil {
			return modelMap, err
		}
		modelMap["vpc_routing_table_route"] = []map[string]interface{}{vpcRoutingTableRouteMap}
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model *vpcv1.RouteReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model vpcv1.DynamicRouteServerStaticRouteNextHopIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference); ok {
		return DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model.(*vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerStaticRouteNextHop); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerStaticRouteNextHop)
		if model.Deleted != nil {
			deletedMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
			if err != nil {
				return modelMap, err
			}
			modelMap["deleted"] = []map[string]interface{}{deletedMap}
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.ResourceType != nil {
			modelMap["resource_type"] = *model.ResourceType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerStaticRouteNextHopIntf subtype encountered")
	}
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model *vpcv1.DynamicRouteServerStaticRouteRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["advertise"] = *model.Advertise
	vpcRoutingTableMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model.VPCRoutingTable)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
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
