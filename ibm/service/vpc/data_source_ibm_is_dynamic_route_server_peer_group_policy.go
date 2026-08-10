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

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPolicyRead,

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
			"is_dynamic_route_server_peer_group_policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server peer group policy identifier.",
			},
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
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}

	getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(d.Get("dynamic_route_server_peer_group_id").(string))
	getDynamicRouteServerPeerGroupPolicyOptions.SetID(d.Get("is_dynamic_route_server_peer_group_policy_id").(string))

	dynamicRouteServerPeerGroupPolicyIntf, _, err := vpcClient.GetDynamicRouteServerPeerGroupPolicyWithContext(context, getDynamicRouteServerPeerGroupPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	dynamicRouteServerPeerGroupPolicy := dynamicRouteServerPeerGroupPolicyIntf.(*vpcv1.DynamicRouteServerPeerGroupPolicy)

	d.SetId(fmt.Sprintf("%s/%s/%s", *getDynamicRouteServerPeerGroupPolicyOptions.DynamicRouteServerID, *getDynamicRouteServerPeerGroupPolicyOptions.DynamicRouteServerPeerGroupID, *getDynamicRouteServerPeerGroupPolicyOptions.ID))

	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerPeerGroupPolicy.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.Href) {
		if err = d.Set("href", dynamicRouteServerPeerGroupPolicy.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-href").GetDiag()
		}
	}

	if err = d.Set("name", dynamicRouteServerPeerGroupPolicy.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-name").GetDiag()
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.ResourceType) {
		if err = d.Set("resource_type", dynamicRouteServerPeerGroupPolicy.ResourceType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-resource_type").GetDiag()
		}
	}

	if err = d.Set("state", dynamicRouteServerPeerGroupPolicy.State); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting state: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-state").GetDiag()
	}

	if err = d.Set("type", dynamicRouteServerPeerGroupPolicy.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-type").GetDiag()
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.CustomRoutes) {
		customRoutes := []map[string]interface{}{}
		for _, customRoutesItem := range dynamicRouteServerPeerGroupPolicy.CustomRoutes {
			customRoutesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&customRoutesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "custom_routes-to-map").GetDiag()
			}
			customRoutes = append(customRoutes, customRoutesItemMap)
		}
		if err = d.Set("custom_routes", customRoutes); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting custom_routes: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-custom_routes").GetDiag()
		}
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.ExcludedPrefixes) {
		excludedPrefixes := []map[string]interface{}{}
		for _, excludedPrefixesItem := range dynamicRouteServerPeerGroupPolicy.ExcludedPrefixes {
			excludedPrefixesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&excludedPrefixesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "excluded_prefixes-to-map").GetDiag()
			}
			excludedPrefixes = append(excludedPrefixes, excludedPrefixesItemMap)
		}
		if err = d.Set("excluded_prefixes", excludedPrefixes); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting excluded_prefixes: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-excluded_prefixes").GetDiag()
		}
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.PeerGroups) {
		peerGroups := []map[string]interface{}{}
		for _, peerGroupsItem := range dynamicRouteServerPeerGroupPolicy.PeerGroups {
			peerGroupsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(&peerGroupsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "peer_groups-to-map").GetDiag()
			}
			peerGroups = append(peerGroups, peerGroupsItemMap)
		}
		if err = d.Set("peer_groups", peerGroups); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting peer_groups: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-peer_groups").GetDiag()
		}
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.NextHops) {
		nextHops := []map[string]interface{}{}
		for _, nextHopsItem := range dynamicRouteServerPeerGroupPolicy.NextHops {
			nextHopsItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(nextHopsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "next_hops-to-map").GetDiag()
			}
			nextHops = append(nextHops, nextHopsItemMap)
		}
		if err = d.Set("next_hops", nextHops); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting next_hops: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-next_hops").GetDiag()
		}
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.RouteDeleteDelay) {
		if err = d.Set("route_delete_delay", flex.IntValue(dynamicRouteServerPeerGroupPolicy.RouteDeleteDelay)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_delete_delay: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-route_delete_delay").GetDiag()
		}
	}

	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.RoutingTables) {
		routingTables := []map[string]interface{}{}
		for _, routingTablesItem := range dynamicRouteServerPeerGroupPolicy.RoutingTables {
			routingTablesItemMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&routingTablesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "routing_tables-to-map").GetDiag()
			}
			routingTables = append(routingTables, routingTablesItemMap)
		}
		if err = d.Set("routing_tables", routingTables); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting routing_tables: %s", err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-routing_tables").GetDiag()
		}
	}

	return nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["destination"] = *model.Destination
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) (map[string]interface{}, error) {
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference); ok {
		return DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHop); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHop)
		if model.Deleted != nil {
			deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["advertise"] = *model.Advertise
	vpcRoutingTableMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model.VPCRoutingTable)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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
