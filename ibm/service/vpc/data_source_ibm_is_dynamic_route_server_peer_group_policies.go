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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicies() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server peer group identifier.",
			},
			"sort": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "-created_at",
				Description: "Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order.",
			},
			"policies": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of policies for the dynamic route server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the peer group policy was created or last updated.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this dynamic route server peer group policy.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dynamic route server peer group policy.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this dynamic route server peer group policy. The name must not be used by another peer group policy. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state used for this dynamic route server peer group policy:- `disabled`: The peer group policy is disabled, and dynamic route server   will not apply the configured policy rules.- `enabled`: The peer group policy is enabled, and dynamic route server will  apply the configured policy rules.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of dynamic route server peer group policy:- `custom_routes`: A policy used for custom routes.  The custom routes are advertised to other peer groups based on the policies  configured.- `learned_routes`:  A policy used for learned routes.  Learned routes are advertised to other peer groups based on the policies  configured.- `vpc_address_prefixes`: A policy used for advertising VPC address prefixes.  The VPC address prefixes are advertised to other peer groups based on the policies  configured.- `vpc_routing_tables`:  A policy used for updating VPC routing tables.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"custom_routes": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The custom routes to advertise.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The destination CIDR of the route. The host identifier in the CIDR must be zero.",
									},
								},
							},
						},
						"excluded_prefixes": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of prefixes that are excluded from monitoring by the dynamic route server when learning routes from connected peer groups. They are applied in addition to any `excluded_prefixes` defined on the peer group. Each excluded prefix must be a subset of the `monitored_prefixes` configured in the peer group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ge": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum prefix length to match. If non-zero, the prefix watcher will only exclude routes within the `prefix` that have a prefix length greater than or equal to this value.If zero, `ge` match filtering will not be applied.If non-zero, `ge` must be:- Greater than or equal to the network `prefix` length.- Less than or equal to 32.If both `ge` and `le` are non-zero, `ge` must be less than or equal to `le`.",
									},
									"le": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum prefix length to match. If non-zero, the prefix watcher will only match routes within the `prefix` that have a prefix length less than or equal to this value.If zero, `le` match filtering will not be applied.If non-zero, `le` value must be:- Less than or equal to the network `prefix` length.- Greater than `ge`.",
									},
									"prefix": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The prefix to be excluded.",
									},
								},
							},
						},
						"peer_groups": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The peer groups from which the routes are being learned and then advertised back to the peer group on which this peer group policy has been applied.",
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
						"route_delete_delay": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of seconds to wait before deleting a route from `routing_tables` when the`status` of the route's associated peer is `down`, or after it is no longer advertised by peers in the peer group.",
						},
						"routing_tables": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The routing tables to update with the learned routes.When `next_hops` specifies multiple next hop addresses, individual routes are added for each address according to the `priority` value of the peer with that next hop `endpoint.address`:- For ingress routing tables, the `priority` value of the peer is used.- For egress routing tables, the route `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).",
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
				},
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listDynamicRouteServerPeerGroupPoliciesOptions := &vpcv1.ListDynamicRouteServerPeerGroupPoliciesOptions{}

	listDynamicRouteServerPeerGroupPoliciesOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	listDynamicRouteServerPeerGroupPoliciesOptions.SetDynamicRouteServerPeerGroupID(d.Get("dynamic_route_server_peer_group_id").(string))
	if _, ok := d.GetOk("sort"); ok {
		listDynamicRouteServerPeerGroupPoliciesOptions.SetSort(d.Get("sort").(string))
	}

	var pager *vpcv1.DynamicRouteServerPeerGroupPoliciesPager
	pager, err = vpcClient.NewDynamicRouteServerPeerGroupPoliciesPager(listDynamicRouteServerPeerGroupPoliciesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerPeerGroupPoliciesPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read", "DynamicRouteServers-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("policies", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting policies %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policies", "read", "policies-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesID returns a reasonable ID for the list.
func dataSourceIBMIsDynamicRouteServerPeerGroupPoliciesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicy); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPolicy)
		modelMap["created_at"] = model.CreatedAt.String()
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		modelMap["name"] = *model.Name
		if model.ResourceType != nil {
			modelMap["resource_type"] = *model.ResourceType
		}
		modelMap["state"] = *model.State
		modelMap["type"] = *model.Type
		if model.CustomRoutes != nil {
			customRoutes := []map[string]interface{}{}
			for _, customRoutesItem := range model.CustomRoutes {
				customRoutesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&customRoutesItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				customRoutes = append(customRoutes, customRoutesItemMap)
			}
			modelMap["custom_routes"] = customRoutes
		}
		if model.ExcludedPrefixes != nil {
			excludedPrefixes := []map[string]interface{}{}
			for _, excludedPrefixesItem := range model.ExcludedPrefixes {
				excludedPrefixesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&excludedPrefixesItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				excludedPrefixes = append(excludedPrefixes, excludedPrefixesItemMap)
			}
			modelMap["excluded_prefixes"] = excludedPrefixes
		}
		if model.PeerGroups != nil {
			peerGroups := []map[string]interface{}{}
			for _, peerGroupsItem := range model.PeerGroups {
				peerGroupsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(&peerGroupsItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				peerGroups = append(peerGroups, peerGroupsItemMap)
			}
			modelMap["peer_groups"] = peerGroups
		}
		if model.NextHops != nil {
			nextHops := []map[string]interface{}{}
			for _, nextHopsItem := range model.NextHops {
				nextHopsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(nextHopsItem)
				if err != nil {
					return modelMap, err
				}
				nextHops = append(nextHops, nextHopsItemMap)
			}
			modelMap["next_hops"] = nextHops
		}
		if model.RouteDeleteDelay != nil {
			modelMap["route_delete_delay"] = flex.IntValue(model.RouteDeleteDelay)
		}
		if model.RoutingTables != nil {
			routingTables := []map[string]interface{}{}
			for _, routingTablesItem := range model.RoutingTables {
				routingTablesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&routingTablesItem) // #nosec G601
				if err != nil {
					return modelMap, err
				}
				routingTables = append(routingTables, routingTablesItemMap)
			}
			modelMap["routing_tables"] = routingTables
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPolicyIntf subtype encountered")
	}
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["destination"] = *model.Destination
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Ge != nil {
		modelMap["ge"] = flex.IntValue(model.Ge)
	}
	if model.Le != nil {
		modelMap["le"] = flex.IntValue(model.Le)
	}
	modelMap["prefix"] = *model.Prefix
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHop); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHop)
		if model.Deleted != nil {
			deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model.Deleted)
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
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf subtype encountered")
	}
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["advertise"] = *model.Advertise
	vpcRoutingTableMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(model.VPCRoutingTable)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRoutesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutes) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["created_at"] = model.CreatedAt.String()
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	modelMap["name"] = *model.Name
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	modelMap["state"] = *model.State
	customRoutes := []map[string]interface{}{}
	for _, customRoutesItem := range model.CustomRoutes {
		customRoutesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&customRoutesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		customRoutes = append(customRoutes, customRoutesItemMap)
	}
	modelMap["custom_routes"] = customRoutes
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyLearnedRoutesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyLearnedRoutes) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["created_at"] = model.CreatedAt.String()
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	modelMap["name"] = *model.Name
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	modelMap["state"] = *model.State
	excludedPrefixes := []map[string]interface{}{}
	for _, excludedPrefixesItem := range model.ExcludedPrefixes {
		excludedPrefixesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&excludedPrefixesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		excludedPrefixes = append(excludedPrefixes, excludedPrefixesItemMap)
	}
	modelMap["excluded_prefixes"] = excludedPrefixes
	peerGroups := []map[string]interface{}{}
	for _, peerGroupsItem := range model.PeerGroups {
		peerGroupsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupReferenceToMap(&peerGroupsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		peerGroups = append(peerGroups, peerGroupsItemMap)
	}
	modelMap["peer_groups"] = peerGroups
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCAddressPrefixes) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["created_at"] = model.CreatedAt.String()
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	modelMap["name"] = *model.Name
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	modelMap["state"] = *model.State
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTables) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["created_at"] = model.CreatedAt.String()
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	modelMap["name"] = *model.Name
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	modelMap["state"] = *model.State
	excludedPrefixes := []map[string]interface{}{}
	for _, excludedPrefixesItem := range model.ExcludedPrefixes {
		excludedPrefixesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&excludedPrefixesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		excludedPrefixes = append(excludedPrefixes, excludedPrefixesItemMap)
	}
	modelMap["excluded_prefixes"] = excludedPrefixes
	nextHops := []map[string]interface{}{}
	for _, nextHopsItem := range model.NextHops {
		nextHopsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyNextHopToMap(nextHopsItem)
		if err != nil {
			return modelMap, err
		}
		nextHops = append(nextHops, nextHopsItemMap)
	}
	modelMap["next_hops"] = nextHops
	modelMap["route_delete_delay"] = flex.IntValue(model.RouteDeleteDelay)
	routingTables := []map[string]interface{}{}
	for _, routingTablesItem := range model.RoutingTables {
		routingTablesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPoliciesDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&routingTablesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		routingTables = append(routingTables, routingTablesItemMap)
	}
	modelMap["routing_tables"] = routingTables
	modelMap["type"] = *model.Type
	return modelMap, nil
}
