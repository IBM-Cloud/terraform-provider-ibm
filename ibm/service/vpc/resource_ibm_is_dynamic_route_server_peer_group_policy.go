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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsDynamicRouteServerPeerGroupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerPeerGroupPolicyCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerPeerGroupPolicyRead,
		UpdateContext: resourceIBMIsDynamicRouteServerPeerGroupPolicyUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerPeerGroupPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"dynamic_route_server_peer_group_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "dynamic_route_server_peer_group_id"),
				Description:  "The dynamic route server peer group identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "name"),
				Description:  "The name for this dynamic route server peer group policy. The name must not be used by another peer group policy. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"state": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "enabled",
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "state"),
				Description:  "The state used for this dynamic route server peer group policy:- `disabled`: The peer group policy is disabled, and dynamic route server   will not apply the configured policy rules.- `enabled`: The peer group policy is enabled, and dynamic route server will  apply the configured policy rules.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "type"),
				Description:  "The type of dynamic route server peer group policy:- `custom_routes`: A policy used for custom routes.  The custom routes are advertised to other peer groups based on the policies  configured.- `learned_routes`:  A policy used for learned routes.  Learned routes are advertised to other peer groups based on the policies  configured.- `vpc_address_prefixes`: A policy used for advertising VPC address prefixes.  The VPC address prefixes are advertised to other peer groups based on the policies  configured.- `vpc_routing_tables`:  A policy used for updating VPC routing tables.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"custom_routes": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The custom routes to advertise.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The destination CIDR of the route. The host identifier in the CIDR must be zero.",
						},
					},
				},
			},
			"excluded_prefixes": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of prefixes that are excluded from monitoring by the dynamic route server when learning routes from connected peer groups. They are applied in addition to any `excluded_prefixes` defined on the peer group. Each excluded prefix must be a subset of the `monitored_prefixes` configured in the peer group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ge": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The minimum prefix length to match. If non-zero, the prefix watcher will only exclude routes within the `prefix` that have a prefix length greater than or equal to this value.If zero, `ge` match filtering will not be applied.If non-zero, `ge` must be:- Greater than or equal to the network `prefix` length.- Less than or equal to 32.If both `ge` and `le` are non-zero, `ge` must be less than or equal to `le`.",
						},
						"le": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The maximum prefix length to match. If non-zero, the prefix watcher will only match routes within the `prefix` that have a prefix length less than or equal to this value.If zero, `le` match filtering will not be applied.If non-zero, `le` value must be:- Less than or equal to the network `prefix` length.- Greater than `ge`.",
						},
						"prefix": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The prefix to be excluded.",
						},
					},
				},
			},
			"peer_groups": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The peer groups from which the routes are being learned and then advertised back to the peer group on which this peer group policy has been applied.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "A link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The URL for this dynamic route server peer group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID for this dynamic route server peer group peer.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name for this dynamic route server peer group.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"next_hops": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The next hop resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "A link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The URL for this dynamic route server peer group.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID for this dynamic route server peer group peer.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name for this dynamic route server peer group.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"route_delete_delay": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_peer_group_policy", "route_delete_delay"),
				Description:  "The number of seconds to wait before deleting a route from `routing_tables` when the`status` of the route's associated peer is `down`, or after it is no longer advertised by peers in the peer group.",
			},
			"routing_tables": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The routing tables to update with the learned routes.When `next_hops` specifies multiple next hop addresses, individual routes are added for each address according to the `priority` value of the peer with that next hop `endpoint.address`:- For ingress routing tables, the `priority` value of the peer is used.- For egress routing tables, the route `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.",
						},
						"vpc_routing_table": &schema.Schema{
							Type:     schema.TypeList,
							MinItems: 1,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The CRN for this VPC routing table.",
									},
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"more_info": &schema.Schema{
													Type:        schema.TypeString,
													Required:    true,
													Description: "A link to documentation about deleted resources.",
												},
											},
										},
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The URL for this routing table.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The unique identifier for this routing table.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name for this routing table. The name is unique across all routing tables for the VPC.",
									},
									"resource_type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The resource type.",
									},
								},
							},
						},
					},
				},
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
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"is_dynamic_route_server_peer_group_policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server peer group policy.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "dynamic_route_server_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "dynamic_route_server_peer_group_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "state",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "disabled, enabled",
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "custom_routes, learned_routes, vpc_address_prefixes, vpc_routing_tables",
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "route_delete_delay",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "60",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_peer_group_policy", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.CreateDynamicRouteServerPeerGroupPolicyOptions{}

	if _, ok := d.GetOk("name"); ok {
		bodyModelMap["name"] = d.Get("name")
	}
	if _, ok := d.GetOk("state"); ok {
		bodyModelMap["state"] = d.Get("state")
	}
	bodyModelMap["type"] = d.Get("type")
	if _, ok := d.GetOk("custom_routes"); ok {
		bodyModelMap["custom_routes"] = d.Get("custom_routes")
	}
	if _, ok := d.GetOk("excluded_prefixes"); ok {
		bodyModelMap["excluded_prefixes"] = d.Get("excluded_prefixes")
	}
	if _, ok := d.GetOk("peer_groups"); ok {
		bodyModelMap["peer_groups"] = d.Get("peer_groups")
	}
	if _, ok := d.GetOk("next_hops"); ok {
		bodyModelMap["next_hops"] = d.Get("next_hops")
	}
	if _, ok := d.GetOk("route_delete_delay"); ok {
		bodyModelMap["route_delete_delay"] = d.Get("route_delete_delay")
	}
	if _, ok := d.GetOk("routing_tables"); ok {
		bodyModelMap["routing_tables"] = d.Get("routing_tables")
	}
	createDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	createDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(d.Get("dynamic_route_server_peer_group_id").(string))
	convertedModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "create", "parse-request-body").GetDiag()
	}
	createDynamicRouteServerPeerGroupPolicyOptions.DynamicRouteServerPeerGroupPolicyPrototype = convertedModel

	dynamicRouteServerPeerGroupPolicyIntf, _, err := vpcClient.CreateDynamicRouteServerPeerGroupPolicyWithContext(context, createDynamicRouteServerPeerGroupPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerPeerGroupPolicy := dynamicRouteServerPeerGroupPolicyIntf.(*vpcv1.DynamicRouteServerPeerGroupPolicy)
	d.SetId(fmt.Sprintf("%s/%s/%s", *createDynamicRouteServerPeerGroupPolicyOptions.DynamicRouteServerID, *createDynamicRouteServerPeerGroupPolicyOptions.DynamicRouteServerPeerGroupID, *dynamicRouteServerPeerGroupPolicy.ID))

	return resourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.GetDynamicRouteServerPeerGroupPolicyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(parts[1])
	getDynamicRouteServerPeerGroupPolicyOptions.SetID(parts[2])

	dynamicRouteServerPeerGroupPolicyIntf, response, err := vpcClient.GetDynamicRouteServerPeerGroupPolicyWithContext(context, getDynamicRouteServerPeerGroupPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerPeerGroupPolicy := dynamicRouteServerPeerGroupPolicyIntf.(*vpcv1.DynamicRouteServerPeerGroupPolicy)
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.Name) {
		if err = d.Set("name", dynamicRouteServerPeerGroupPolicy.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.State) {
		if err = d.Set("state", dynamicRouteServerPeerGroupPolicy.State); err != nil {
			err = fmt.Errorf("Error setting state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-state").GetDiag()
		}
	}
	if err = d.Set("type", dynamicRouteServerPeerGroupPolicy.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-type").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.CustomRoutes) {
		customRoutes := []map[string]interface{}{}
		for _, customRoutesItem := range dynamicRouteServerPeerGroupPolicy.CustomRoutes {
			customRoutesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(&customRoutesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "custom_routes-to-map").GetDiag()
			}
			customRoutes = append(customRoutes, customRoutesItemMap)
		}
		if err = d.Set("custom_routes", customRoutes); err != nil {
			err = fmt.Errorf("Error setting custom_routes: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-custom_routes").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.ExcludedPrefixes) {
		excludedPrefixes := []map[string]interface{}{}
		for _, excludedPrefixesItem := range dynamicRouteServerPeerGroupPolicy.ExcludedPrefixes {
			excludedPrefixesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(&excludedPrefixesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "excluded_prefixes-to-map").GetDiag()
			}
			excludedPrefixes = append(excludedPrefixes, excludedPrefixesItemMap)
		}
		if err = d.Set("excluded_prefixes", excludedPrefixes); err != nil {
			err = fmt.Errorf("Error setting excluded_prefixes: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-excluded_prefixes").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.PeerGroups) {
		peerGroups := []map[string]interface{}{}
		for _, peerGroupsItem := range dynamicRouteServerPeerGroupPolicy.PeerGroups {
			peerGroupsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(&peerGroupsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "peer_groups-to-map").GetDiag()
			}
			peerGroups = append(peerGroups, peerGroupsItemMap)
		}
		if err = d.Set("peer_groups", peerGroups); err != nil {
			err = fmt.Errorf("Error setting peer_groups: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-peer_groups").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.NextHops) {
		nextHops := []map[string]interface{}{}
		for _, nextHopsItem := range dynamicRouteServerPeerGroupPolicy.NextHops {
			nextHopsItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(nextHopsItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "next_hops-to-map").GetDiag()
			}
			nextHops = append(nextHops, nextHopsItemMap)
		}
		if err = d.Set("next_hops", nextHops); err != nil {
			err = fmt.Errorf("Error setting next_hops: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-next_hops").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.RouteDeleteDelay) {
		if err = d.Set("route_delete_delay", flex.IntValue(dynamicRouteServerPeerGroupPolicy.RouteDeleteDelay)); err != nil {
			err = fmt.Errorf("Error setting route_delete_delay: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-route_delete_delay").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.RoutingTables) {
		routingTables := []map[string]interface{}{}
		for _, routingTablesItem := range dynamicRouteServerPeerGroupPolicy.RoutingTables {
			routingTablesItemMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(&routingTablesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "routing_tables-to-map").GetDiag()
			}
			routingTables = append(routingTables, routingTablesItemMap)
		}
		if err = d.Set("routing_tables", routingTables); err != nil {
			err = fmt.Errorf("Error setting routing_tables: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-routing_tables").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerPeerGroupPolicy.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.Href) {
		if err = d.Set("href", dynamicRouteServerPeerGroupPolicy.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-href").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.ResourceType) {
		if err = d.Set("resource_type", dynamicRouteServerPeerGroupPolicy.ResourceType); err != nil {
			err = fmt.Errorf("Error setting resource_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-resource_type").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerPeerGroupPolicy.ID) {
		if err = d.Set("is_dynamic_route_server_peer_group_policy_id", dynamicRouteServerPeerGroupPolicy.ID); err != nil {
			err = fmt.Errorf("Error setting is_dynamic_route_server_peer_group_policy_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-is_dynamic_route_server_peer_group_policy_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_peer_group_policy", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.UpdateDynamicRouteServerPeerGroupPolicyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "sep-id-parts").GetDiag()
	}

	updateDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(parts[0])
	updateDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(parts[1])
	updateDynamicRouteServerPeerGroupPolicyOptions.SetID(parts[2])

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerPeerGroupPolicyPatch{}
	if d.HasChange("dynamic_route_server_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_peer_group_policy", "update", "dynamic_route_server_id-forces-new").GetDiag()
	}
	if d.HasChange("dynamic_route_server_peer_group_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_peer_group_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_peer_group_policy", "update", "dynamic_route_server_peer_group_id-forces-new").GetDiag()
	}
	if d.HasChange("custom_routes") {
		var customRoutes []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype
		for _, v := range d.Get("custom_routes").([]interface{}) {
			value := v.(map[string]interface{})
			customRoutesItem, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-custom_routes").GetDiag()
			}
			customRoutes = append(customRoutes, *customRoutesItem)
		}
		patchVals.CustomRoutes = customRoutes
		hasChange = true
	}
	if d.HasChange("excluded_prefixes") {
		var excludedPrefixes []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype
		for _, v := range d.Get("excluded_prefixes").([]interface{}) {
			value := v.(map[string]interface{})
			excludedPrefixesItem, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-excluded_prefixes").GetDiag()
			}
			excludedPrefixes = append(excludedPrefixes, *excludedPrefixesItem)
		}
		patchVals.ExcludedPrefixes = excludedPrefixes
		hasChange = true
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("next_hops") {
		var nextHops []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototype
		for _, v := range d.Get("next_hops").([]interface{}) {
			value := v.(map[string]interface{})
			nextHopsItem, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-next_hops").GetDiag()
			}
			nextHops = append(nextHops, *nextHopsItem)
		}
		patchVals.NextHops = nextHops
		hasChange = true
	}
	if d.HasChange("peer_groups") {
		var peerGroups []vpcv1.DynamicRouteServerPeerGroupIdentity
		for _, v := range d.Get("peer_groups").([]interface{}) {
			value := v.(map[string]interface{})
			peerGroupsItem, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-peer_groups").GetDiag()
			}
			peerGroups = append(peerGroups, *peerGroupsItem)
		}
		patchVals.PeerGroups = peerGroups
		hasChange = true
	}
	if d.HasChange("route_delete_delay") {
		newRouteDeleteDelay := int64(d.Get("route_delete_delay").(int))
		patchVals.RouteDeleteDelay = &newRouteDeleteDelay
		hasChange = true
	}
	if d.HasChange("routing_tables") {
		var routingTables []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch
		for _, v := range d.Get("routing_tables").([]interface{}) {
			value := v.(map[string]interface{})
			routingTablesItem, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "update", "parse-routing_tables").GetDiag()
			}
			routingTables = append(routingTables, *routingTablesItem)
		}
		patchVals.RoutingTables = routingTables
		hasChange = true
	}
	if d.HasChange("state") {
		newState := d.Get("state").(string)
		patchVals.State = &newState
		hasChange = true
	}
	updateDynamicRouteServerPeerGroupPolicyOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateDynamicRouteServerPeerGroupPolicyOptions.DynamicRouteServerPeerGroupPolicyPatch = ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerPeerGroupPolicyWithContext(context, updateDynamicRouteServerPeerGroupPolicyOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerPeerGroupPolicyRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerPeerGroupPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDynamicRouteServerPeerGroupPolicyOptions := &vpcv1.DeleteDynamicRouteServerPeerGroupPolicyOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_peer_group_policy", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerPeerGroupPolicyOptions.SetDynamicRouteServerPeerGroupID(parts[1])
	deleteDynamicRouteServerPeerGroupPolicyOptions.SetID(parts[2])

	deleteDynamicRouteServerPeerGroupPolicyOptions.SetIfMatch(d.Get("etag").(string))

	_, _, err = vpcClient.DeleteDynamicRouteServerPeerGroupPolicyWithContext(context, deleteDynamicRouteServerPeerGroupPolicyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerPeerGroupPolicyWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_peer_group_policy", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}
	model.Destination = core.StringPtr(modelMap["destination"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
	if modelMap["ge"] != nil {
		model.Ge = core.Int64Ptr(int64(modelMap["ge"].(int)))
	}
	if modelMap["le"] != nil {
		model.Le = core.Int64Ptr(int64(modelMap["le"].(int)))
	}
	model.Prefix = core.StringPtr(modelMap["prefix"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupIdentityIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupIdentityByID, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupIdentityByHref, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototype{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{}
	if modelMap["advertise"] != nil {
		model.Advertise = core.BoolPtr(modelMap["advertise"].(bool))
	}
	VPCRoutingTableModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(modelMap["vpc_routing_table"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.VPCRoutingTable = VPCRoutingTableModel
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(modelMap map[string]interface{}) (vpcv1.RoutingTableIdentityIntf, error) {
	model := &vpcv1.RoutingTableIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["crn"] != nil && modelMap["crn"].(string) != "" {
		model.CRN = core.StringPtr(modelMap["crn"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByID(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByID, error) {
	model := &vpcv1.RoutingTableIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByCRN, error) {
	model := &vpcv1.RoutingTableIdentityByCRN{}
	model.CRN = core.StringPtr(modelMap["crn"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentityByHref(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByHref, error) {
	model := &vpcv1.RoutingTableIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatch{}
	if modelMap["advertise"] != nil {
		model.Advertise = core.BoolPtr(modelMap["advertise"].(bool))
	}
	VPCRoutingTableModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToRoutingTableIdentity(modelMap["vpc_routing_table"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.VPCRoutingTable = VPCRoutingTableModel
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototype{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["custom_routes"] != nil {
		customRoutes := []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}
		for _, customRoutesItem := range modelMap["custom_routes"].([]interface{}) {
			customRoutesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(customRoutesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			customRoutes = append(customRoutes, *customRoutesItemModel)
		}
		model.CustomRoutes = customRoutes
	}
	if modelMap["excluded_prefixes"] != nil {
		excludedPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
		for _, excludedPrefixesItem := range modelMap["excluded_prefixes"].([]interface{}) {
			excludedPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(excludedPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			excludedPrefixes = append(excludedPrefixes, *excludedPrefixesItemModel)
		}
		model.ExcludedPrefixes = excludedPrefixes
	}
	if modelMap["peer_groups"] != nil {
		peerGroups := []vpcv1.DynamicRouteServerPeerGroupIdentityIntf{}
		for _, peerGroupsItem := range modelMap["peer_groups"].([]interface{}) {
			peerGroupsItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(peerGroupsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			peerGroups = append(peerGroups, peerGroupsItemModel)
		}
		model.PeerGroups = peerGroups
	}
	if modelMap["next_hops"] != nil {
		nextHops := []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf{}
		for _, nextHopsItem := range modelMap["next_hops"].([]interface{}) {
			nextHopsItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(nextHopsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			nextHops = append(nextHops, nextHopsItemModel)
		}
		model.NextHops = nextHops
	}
	if modelMap["route_delete_delay"] != nil {
		model.RouteDeleteDelay = core.Int64Ptr(int64(modelMap["route_delete_delay"].(int)))
	}
	if modelMap["routing_tables"] != nil {
		routingTables := []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{}
		for _, routingTablesItem := range modelMap["routing_tables"].([]interface{}) {
			routingTablesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(routingTablesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			routingTables = append(routingTables, *routingTablesItemModel)
		}
		model.RoutingTables = routingTables
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyCustomRoutesPrototype{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	customRoutes := []vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoutePrototype{}
	for _, customRoutesItem := range modelMap["custom_routes"].([]interface{}) {
		customRoutesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyCustomRoutePrototype(customRoutesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		customRoutes = append(customRoutes, *customRoutesItemModel)
	}
	model.CustomRoutes = customRoutes
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyLearnedRoutesPrototype{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	if modelMap["excluded_prefixes"] != nil {
		excludedPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
		for _, excludedPrefixesItem := range modelMap["excluded_prefixes"].([]interface{}) {
			excludedPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(excludedPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			excludedPrefixes = append(excludedPrefixes, *excludedPrefixesItemModel)
		}
		model.ExcludedPrefixes = excludedPrefixes
	}
	peerGroups := []vpcv1.DynamicRouteServerPeerGroupIdentityIntf{}
	for _, peerGroupsItem := range modelMap["peer_groups"].([]interface{}) {
		peerGroupsItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupIdentity(peerGroupsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		peerGroups = append(peerGroups, peerGroupsItemModel)
	}
	model.PeerGroups = peerGroups
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCAddressPrefixesPrototype{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype, error) {
	model := &vpcv1.DynamicRouteServerPeerGroupPolicyPrototypeDynamicRouteServerPeerGroupPolicyVPCRoutingTablesPrototype{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	if modelMap["excluded_prefixes"] != nil {
		excludedPrefixes := []vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype{}
		for _, excludedPrefixesItem := range modelMap["excluded_prefixes"].([]interface{}) {
			excludedPrefixesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype(excludedPrefixesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			excludedPrefixes = append(excludedPrefixes, *excludedPrefixesItemModel)
		}
		model.ExcludedPrefixes = excludedPrefixes
	}
	nextHops := []vpcv1.DynamicRouteServerPeerGroupPolicyNextHopPrototypeIntf{}
	for _, nextHopsItem := range modelMap["next_hops"].([]interface{}) {
		nextHopsItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyNextHopPrototype(nextHopsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		nextHops = append(nextHops, nextHopsItemModel)
	}
	model.NextHops = nextHops
	if modelMap["route_delete_delay"] != nil {
		model.RouteDeleteDelay = core.Int64Ptr(int64(modelMap["route_delete_delay"].(int)))
	}
	routingTables := []vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype{}
	for _, routingTablesItem := range modelMap["routing_tables"].([]interface{}) {
		routingTablesItemModel, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyMapToDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePrototype(routingTablesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		routingTables = append(routingTables, *routingTablesItemModel)
	}
	model.RoutingTables = routingTables
	model.Type = core.StringPtr(modelMap["type"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyCustomRouteToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyCustomRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["destination"] = *model.Destination
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeToMap(model *vpcv1.DynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototype) (map[string]interface{}, error) {
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

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopToMap(model vpcv1.DynamicRouteServerPeerGroupPolicyNextHopIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference); ok {
		return ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHop); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerPeerGroupPolicyNextHop)
		if model.Deleted != nil {
			deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTableToMap(model *vpcv1.DynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["advertise"] = *model.Advertise
	vpcRoutingTableMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model.VPCRoutingTable)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerPeerGroupPolicyDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyPatchAsPatch(patchVals *vpcv1.DynamicRouteServerPeerGroupPolicyPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "custom_routes"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["custom_routes"] = nil
	} else if !exists {
		delete(patch, "custom_routes")
	}
	path = "excluded_prefixes"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["excluded_prefixes"] = nil
	} else if exists && patch["excluded_prefixes"] != nil {
		excluded_prefixesList := patch["excluded_prefixes"].([]map[string]interface{})
		for i, excluded_prefixesItem := range excluded_prefixesList {
			ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeAsPatch(excluded_prefixesItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "excluded_prefixes")
	}
	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}
	path = "next_hops"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["next_hops"] = nil
	} else if exists && patch["next_hops"] != nil {
		next_hopsList := patch["next_hops"].([]map[string]interface{})
		for i, next_hopsItem := range next_hopsList {
			ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopPrototypeAsPatch(next_hopsItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "next_hops")
	}
	path = "peer_groups"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["peer_groups"] = nil
	} else if exists && patch["peer_groups"] != nil {
		peer_groupsList := patch["peer_groups"].([]map[string]interface{})
		for i, peer_groupsItem := range peer_groupsList {
			ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupIdentityAsPatch(peer_groupsItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "peer_groups")
	}
	path = "route_delete_delay"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["route_delete_delay"] = nil
	} else if !exists {
		delete(patch, "route_delete_delay")
	}
	path = "routing_tables"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["routing_tables"] = nil
	} else if exists && patch["routing_tables"] != nil {
		routing_tablesList := patch["routing_tables"].([]map[string]interface{})
		for i, routing_tablesItem := range routing_tablesList {
			ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatchAsPatch(routing_tablesItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "routing_tables")
	}
	path = "state"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["state"] = nil
	} else if !exists {
		delete(patch, "state")
	}

	return patch
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyVPCRoutingTablesRoutingTablePatchAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".advertise"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["advertise"] = nil
	} else if !exists {
		delete(patch, "advertise")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupIdentityAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".id"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["id"] = nil
	} else if !exists {
		delete(patch, "id")
	}
	path = rootPath + ".href"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["href"] = nil
	} else if !exists {
		delete(patch, "href")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPolicyNextHopPrototypeAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".id"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["id"] = nil
	} else if !exists {
		delete(patch, "id")
	}
	path = rootPath + ".href"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["href"] = nil
	} else if !exists {
		delete(patch, "href")
	}
}

func ResourceIBMIsDynamicRouteServerPeerGroupPolicyDynamicRouteServerPeerGroupPrefixWatcherExcludedPrefixPrototypeAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".ge"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["ge"] = nil
	} else if !exists {
		delete(patch, "ge")
	}
	path = rootPath + ".le"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["le"] = nil
	} else if !exists {
		delete(patch, "le")
	}
}
