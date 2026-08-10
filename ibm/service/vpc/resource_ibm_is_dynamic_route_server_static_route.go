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

func ResourceIBMIsDynamicRouteServerStaticRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerStaticRouteCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerStaticRouteRead,
		UpdateContext: resourceIBMIsDynamicRouteServerStaticRouteUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerStaticRouteDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"destination": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "destination"),
				Description:  "The destination CIDR of the route. The host identifier in the CIDR must be zero.At most two routes per `zone` in a VPC routing table can have the same `destination` and `priority`, and only if both routes have an `action` of `deliver` and the`next_hop` is an IP address.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "name"),
				Description:  "The name for this static route. The name must not be used by another static route for this dynamic route server. Names starting with ibm- are reserved for ibm managed policies, and are not allowed. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"next_hops": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
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
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "route_delete_delay"),
				Description:  "The number of seconds to wait before deleting a route from `routing_tables` when the `status` of the route's associated peer is `down`.",
			},
			"routing_tables": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The routing tables to update.A route is added to each routing table for each peer in the `next_hop` resource. Each route uses its associated peer's `endpoint.address` for the route's `next_hop`, and:- For ingress routing tables, the route's `priority` uses the peer's `priority`.- For egress routing tables, the route's `priority` is determined by  [Zone priority rules](https://cloud.ibm.com/docs/__TBD__).",
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
			"added_routes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routes added to the VPC routing tables for this dynamic route server static route.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_routing_table_route": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The route reference for the applied route, If absent the routecould not be added to the routing table.For more information, see[Dynamic Route Server Static Route Failure](https://cloud.ibm.com/docs/__TBD__).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deleted": &schema.Schema{
										Type:        schema.TypeList,
										Optional:    true,
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
							Optional:    true,
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
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"is_dynamic_route_server_static_route_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server static routes.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerStaticRouteValidator() *validate.ResourceValidator {
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
			Identifier:                 "destination",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$`,
			MinValueLength:             9,
			MaxValueLength:             18,
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
			Identifier:                 "route_delete_delay",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "0",
			MaxValue:                   "60",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_static_route", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerStaticRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDynamicRouteServerStaticRouteOptions := &vpcv1.CreateDynamicRouteServerStaticRouteOptions{}

	createDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	createDynamicRouteServerStaticRouteOptions.SetDestination(d.Get("destination").(string))
	var nextHops []vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeIntf
	for _, v := range d.Get("next_hops").([]interface{}) {
		value := v.(map[string]interface{})
		nextHopsItem, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "create", "parse-next_hops").GetDiag()
		}
		nextHops = append(nextHops, nextHopsItem)
	}
	createDynamicRouteServerStaticRouteOptions.SetNextHops(nextHops)
	var routingTables []vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype
	for _, v := range d.Get("routing_tables").([]interface{}) {
		value := v.(map[string]interface{})
		routingTablesItem, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "create", "parse-routing_tables").GetDiag()
		}
		routingTables = append(routingTables, *routingTablesItem)
	}
	createDynamicRouteServerStaticRouteOptions.SetRoutingTables(routingTables)
	if _, ok := d.GetOk("name"); ok {
		createDynamicRouteServerStaticRouteOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("route_delete_delay"); ok {
		createDynamicRouteServerStaticRouteOptions.SetRouteDeleteDelay(int64(d.Get("route_delete_delay").(int)))
	}

	dynamicRouteServerStaticRoute, _, err := vpcClient.CreateDynamicRouteServerStaticRouteWithContext(context, createDynamicRouteServerStaticRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createDynamicRouteServerStaticRouteOptions.DynamicRouteServerID, *dynamicRouteServerStaticRoute.ID))

	return resourceIBMIsDynamicRouteServerStaticRouteRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerStaticRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerStaticRouteOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerStaticRouteOptions.SetID(parts[1])

	dynamicRouteServerStaticRoute, response, err := vpcClient.GetDynamicRouteServerStaticRouteWithContext(context, getDynamicRouteServerStaticRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("destination", dynamicRouteServerStaticRoute.Destination); err != nil {
		err = fmt.Errorf("Error setting destination: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-destination").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerStaticRoute.Name) {
		if err = d.Set("name", dynamicRouteServerStaticRoute.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-name").GetDiag()
		}
	}
	nextHops := []map[string]interface{}{}
	for _, nextHopsItem := range dynamicRouteServerStaticRoute.NextHops {
		nextHopsItemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(nextHopsItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "next_hops-to-map").GetDiag()
		}
		nextHops = append(nextHops, nextHopsItemMap)
	}
	if err = d.Set("next_hops", nextHops); err != nil {
		err = fmt.Errorf("Error setting next_hops: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-next_hops").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerStaticRoute.RouteDeleteDelay) {
		if err = d.Set("route_delete_delay", flex.IntValue(dynamicRouteServerStaticRoute.RouteDeleteDelay)); err != nil {
			err = fmt.Errorf("Error setting route_delete_delay: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-route_delete_delay").GetDiag()
		}
	}
	routingTables := []map[string]interface{}{}
	for _, routingTablesItem := range dynamicRouteServerStaticRoute.RoutingTables {
		routingTablesItemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&routingTablesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "routing_tables-to-map").GetDiag()
		}
		routingTables = append(routingTables, routingTablesItemMap)
	}
	if err = d.Set("routing_tables", routingTables); err != nil {
		err = fmt.Errorf("Error setting routing_tables: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-routing_tables").GetDiag()
	}
	addedRoutes := []map[string]interface{}{}
	for _, addedRoutesItem := range dynamicRouteServerStaticRoute.AddedRoutes {
		addedRoutesItemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&addedRoutesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "added_routes-to-map").GetDiag()
		}
		addedRoutes = append(addedRoutes, addedRoutesItemMap)
	}
	if err = d.Set("added_routes", addedRoutes); err != nil {
		err = fmt.Errorf("Error setting added_routes: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-added_routes").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerStaticRoute.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range dynamicRouteServerStaticRoute.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-lifecycle_reasons").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerStaticRoute.LifecycleState) {
		if err = d.Set("lifecycle_state", dynamicRouteServerStaticRoute.LifecycleState); err != nil {
			err = fmt.Errorf("Error setting lifecycle_state: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-lifecycle_state").GetDiag()
		}
	}
	if err = d.Set("resource_type", dynamicRouteServerStaticRoute.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("is_dynamic_route_server_static_route_id", dynamicRouteServerStaticRoute.ID); err != nil {
		err = fmt.Errorf("Error setting is_dynamic_route_server_static_route_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "set-is_dynamic_route_server_static_route_id").GetDiag()
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_is_dynamic_route_server_static_route", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerStaticRouteUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateDynamicRouteServerStaticRouteOptions := &vpcv1.UpdateDynamicRouteServerStaticRouteOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "sep-id-parts").GetDiag()
	}

	updateDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(parts[0])
	updateDynamicRouteServerStaticRouteOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerStaticRoutePatch{}
	if d.HasChange("dynamic_route_server_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_static_route", "update", "dynamic_route_server_id-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("next_hops") {
		var nextHops []vpcv1.DynamicRouteServerStaticRouteNextHopPrototype
		for _, v := range d.Get("next_hops").([]interface{}) {
			value := v.(map[string]interface{})
			nextHopsItem, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "parse-next_hops").GetDiag()
			}
			nextHops = append(nextHops, *nextHopsItem)
		}
		patchVals.NextHops = nextHops
		hasChange = true
	}
	if d.HasChange("route_delete_delay") {
		newRouteDeleteDelay := int64(d.Get("route_delete_delay").(int))
		patchVals.RouteDeleteDelay = &newRouteDeleteDelay
		hasChange = true
	}
	if d.HasChange("routing_tables") {
		var routingTables []vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype
		for _, v := range d.Get("routing_tables").([]interface{}) {
			value := v.(map[string]interface{})
			routingTablesItem, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "parse-routing_tables").GetDiag()
			}
			routingTables = append(routingTables, *routingTablesItem)
		}
		patchVals.RoutingTables = routingTables
		hasChange = true
	}
	updateDynamicRouteServerStaticRouteOptions.SetIfMatch(d.Get("etag").(string))

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateDynamicRouteServerStaticRouteOptions.DynamicRouteServerStaticRoutePatch = ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRoutePatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerStaticRouteWithContext(context, updateDynamicRouteServerStaticRouteOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerStaticRouteRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerStaticRouteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDynamicRouteServerStaticRouteOptions := &vpcv1.DeleteDynamicRouteServerStaticRouteOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerStaticRouteOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerStaticRouteOptions.SetID(parts[1])

	deleteDynamicRouteServerStaticRouteOptions.SetIfMatch(d.Get("etag").(string))

	_, _, err = vpcClient.DeleteDynamicRouteServerStaticRouteWithContext(context, deleteDynamicRouteServerStaticRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeIntf, error) {
	model := &vpcv1.DynamicRouteServerStaticRouteNextHopPrototype{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityIntf, error) {
	model := &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID, error) {
	model := &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref, error) {
	model := &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype, error) {
	model := &vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype{}
	if modelMap["advertise"] != nil {
		model.Advertise = core.BoolPtr(modelMap["advertise"].(bool))
	}
	VPCRoutingTableModel, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(modelMap["vpc_routing_table"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.VPCRoutingTable = VPCRoutingTableModel
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(modelMap map[string]interface{}) (vpcv1.RoutingTableIdentityIntf, error) {
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

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByID, error) {
	model := &vpcv1.RoutingTableIdentityByID{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByCRN, error) {
	model := &vpcv1.RoutingTableIdentityByCRN{}
	model.CRN = core.StringPtr(modelMap["crn"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByHref, error) {
	model := &vpcv1.RoutingTableIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model vpcv1.DynamicRouteServerStaticRouteNextHopIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference); ok {
		return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model.(*vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference))
	} else if _, ok := model.(*vpcv1.DynamicRouteServerStaticRouteNextHop); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.DynamicRouteServerStaticRouteNextHop)
		if model.Deleted != nil {
			deletedMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model *vpcv1.DynamicRouteServerStaticRouteRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["advertise"] = *model.Advertise
	vpcRoutingTableMap, err := ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model.VPCRoutingTable)
	if err != nil {
		return modelMap, err
	}
	modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model *vpcv1.DynamicRouteServerStaticRouteAddedRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VPCRoutingTableRoute != nil {
		vpcRoutingTableRouteMap, err := ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model.VPCRoutingTableRoute)
		if err != nil {
			return modelMap, err
		}
		modelMap["vpc_routing_table_route"] = []map[string]interface{}{vpcRoutingTableRouteMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model *vpcv1.RouteReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model.Deleted)
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

func ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRoutePatchAsPatch(patchVals *vpcv1.DynamicRouteServerStaticRoutePatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

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
			ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopPrototypeAsPatch(next_hopsItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "next_hops")
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
			ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTablePrototypeAsPatch(routing_tablesItem, d, fmt.Sprintf("%s.%d", path, i))
		}
	} else if !exists {
		delete(patch, "routing_tables")
	}

	return patch
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTablePrototypeAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
	var path string

	path = rootPath + ".advertise"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["advertise"] = nil
	} else if !exists {
		delete(patch, "advertise")
	}
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopPrototypeAsPatch(patch map[string]interface{}, d *schema.ResourceData, rootPath string) {
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
