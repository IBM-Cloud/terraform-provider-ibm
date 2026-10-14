// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// dataSourceIBMIsDynamicRouteServerPeerGroupPolicyReferenceSchema is the computed schema for a peer group
// reference (peer_groups / next_hops items).
func dataSourceIBMIsDynamicRouteServerPeerGroupPolicyReferenceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deleted": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
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
				Description: "The unique identifier for this dynamic route server peer group.",
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
	}
}

// dataSourceIBMIsDynamicRouteServerPeerGroupPolicyItemSchema is the computed schema of one policy, shared by the
// single and list data sources.
func dataSourceIBMIsDynamicRouteServerPeerGroupPolicyItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"lifecycle_reasons": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The reasons for the current `lifecycle_state` (if any).",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"code": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "A reason code for this lifecycle state.",
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
			Description: "The lifecycle state of the dynamic route server peer group policy.",
		},
		"name": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name for this dynamic route server peer group policy.",
		},
		"resource_type": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The resource type.",
		},
		"state": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The state of this dynamic route server peer group policy. One of `enabled` or `disabled`.",
		},
		"type": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The type of dynamic route server peer group policy. One of `custom_routes`, `learned_routes`, `vpc_address_prefixes` or `vpc_routing_tables`.",
		},
		"custom_routes": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The custom routes to advertise (type `custom_routes`).",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"destination": &schema.Schema{
						Type:        schema.TypeString,
						Computed:    true,
						Description: "The destination CIDR of the route.",
					},
				},
			},
		},
		"excluded_prefixes": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The prefixes excluded from monitoring when learning routes from connected peer groups.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ge": &schema.Schema{
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "The minimum prefix length to match. Zero disables `ge` filtering.",
					},
					"le": &schema.Schema{
						Type:        schema.TypeInt,
						Computed:    true,
						Description: "The maximum prefix length to match. Zero disables `le` filtering.",
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
			Description: "The peer groups from which routes are learned (type `learned_routes`).",
			Elem:        dataSourceIBMIsDynamicRouteServerPeerGroupPolicyReferenceSchema(),
		},
		"next_hops": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The next hop peer groups (type `vpc_routing_tables`).",
			Elem:        dataSourceIBMIsDynamicRouteServerPeerGroupPolicyReferenceSchema(),
		},
		"route_delete_delay": &schema.Schema{
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`, or after it is no longer advertised.",
		},
		"routing_tables": &schema.Schema{
			Type:        schema.TypeList,
			Computed:    true,
			Description: "The routing tables to update with the learned routes (type `vpc_routing_tables`).",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"advertise": &schema.Schema{
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.",
					},
					"vpc_routing_table": &schema.Schema{
						Type:        schema.TypeList,
						Computed:    true,
						Description: "The VPC routing table.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"crn": &schema.Schema{
									Type:        schema.TypeString,
									Computed:    true,
									Description: "The CRN for this routing table.",
								},
								"deleted": &schema.Schema{
									Type:        schema.TypeList,
									Computed:    true,
									Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
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
									Description: "The name for this routing table.",
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
	}
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicy() *schema.Resource {
	policySchema := dataSourceIBMIsDynamicRouteServerPeerGroupPolicyItemSchema()
	policySchema["dynamic_route_server_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The dynamic route server identifier.",
	}
	policySchema["dynamic_route_server_peer_group_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The dynamic route server peer group identifier.",
	}
	policySchema["dynamic_route_server_peer_group_policy_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: "The dynamic route server peer group policy identifier.",
	}
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerPeerGroupPolicyRead,
		Schema:      policySchema,
	}
}

func dataSourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	peerGroupID := d.Get("dynamic_route_server_peer_group_id").(string)
	policyID := d.Get("dynamic_route_server_peer_group_policy_id").(string)

	getOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}
	getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	getOptions.SetDynamicRouteServerPeerGroupID(peerGroupID)
	getOptions.SetID(policyID)

	policyIntf, _, err := vpcClient.GetDynamicRouteServerPeerGroupPolicyWithContext(context, getOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	modelMap, err := DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(policyIntf)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "policy-to-map").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", dynamicRouteServerID, peerGroupID, policyID))

	for key, value := range modelMap {
		if key == "id" {
			continue
		}
		if err = d.Set(key, value); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "(Data) ibm_is_dynamic_route_server_peer_group_policy", "read", "set-"+key).GetDiag()
		}
	}

	return nil
}

// DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap flattens any policy
// subtype into the data source attribute map.
func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyIntf) (map[string]interface{}, error) {
	policy, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyToBase(model)
	if err != nil {
		return nil, err
	}
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyToMap(policy)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model)
}

func DataSourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model)
}
