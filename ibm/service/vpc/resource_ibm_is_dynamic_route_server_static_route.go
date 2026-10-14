// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isDynamicRouteServerStaticRouteStable   = "stable"
	isDynamicRouteServerStaticRouteFailed   = "failed"
	isDynamicRouteServerStaticRouteDeleted  = "deleted"
	isDynamicRouteServerStaticRoutePending  = "pending"
	isDynamicRouteServerStaticRouteDeleting = "deleting"
	isDynamicRouteServerStaticRouteUpdating = "updating"
	isDynamicRouteServerStaticRouteWaiting  = "waiting"
)

func ResourceIBMIsDynamicRouteServerStaticRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerStaticRouteCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerStaticRouteRead,
		UpdateContext: resourceIBMIsDynamicRouteServerStaticRouteUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerStaticRouteDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

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
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "destination"),
				Description:  "The destination CIDR of the route. The host identifier in the CIDR must be zero.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "name"),
				Description:  "The name for this static route. The name must not be used by another static route for this dynamic route server. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"next_hops": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "The next hop peer groups. Each item is identified by `id` or `href`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The unique identifier for this dynamic route server peer group.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The URL for this dynamic route server peer group.",
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
					},
				},
			},
			"route_delete_delay": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      0,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_static_route", "route_delete_delay"),
				Description:  "The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`.",
			},
			"routing_tables": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "The routing tables to update. A route is added to each routing table for each peer in the `next_hops` peer groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advertise": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.",
						},
						"vpc_routing_table": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The VPC routing table, identified by `id`, `href` or `crn`.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The unique identifier for this routing table.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The URL for this routing table.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The CRN for this routing table.",
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
							Computed:    true,
							Description: "The route reference for the applied route. If absent, the route could not be added to the routing table.",
							Elem: &schema.Resource{
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
										Description: "The name for this route.",
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
				Description: "The URL for this dynamic route server static route.",
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
				Description: "The lifecycle state of the static route.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"dynamic_route_server_static_route_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server static route.",
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

func resourceIBMIsDynamicRouteServerStaticRouteNextHopsFromList(list []interface{}) ([]vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeIntf, error) {
	nextHops := []vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeIntf{}
	for _, v := range list {
		value, _ := v.(map[string]interface{})
		item, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(value)
		if err != nil {
			return nil, err
		}
		nextHops = append(nextHops, item)
	}
	return nextHops, nil
}

func resourceIBMIsDynamicRouteServerStaticRouteRoutingTablesFromList(list []interface{}) ([]vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype, error) {
	routingTables := []vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype{}
	for _, v := range list {
		value, _ := v.(map[string]interface{})
		item, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(value)
		if err != nil {
			return nil, err
		}
		routingTables = append(routingTables, *item)
	}
	return routingTables, nil
}

func resourceIBMIsDynamicRouteServerStaticRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	createOptions := &vpcv1.CreateDynamicRouteServerStaticRouteOptions{}
	createOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	createOptions.SetDestination(d.Get("destination").(string))

	nextHops, err := resourceIBMIsDynamicRouteServerStaticRouteNextHopsFromList(d.Get("next_hops").([]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "create", "parse-next_hops").GetDiag()
	}
	createOptions.SetNextHops(nextHops)

	routingTables, err := resourceIBMIsDynamicRouteServerStaticRouteRoutingTablesFromList(d.Get("routing_tables").([]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "create", "parse-routing_tables").GetDiag()
	}
	createOptions.SetRoutingTables(routingTables)

	if _, ok := d.GetOk("name"); ok {
		createOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("route_delete_delay"); ok {
		createOptions.SetRouteDeleteDelay(int64(d.Get("route_delete_delay").(int)))
	}

	staticRoute, _, err := vpcClient.CreateDynamicRouteServerStaticRouteWithContext(context, createOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", dynamicRouteServerID, *staticRoute.ID))

	_, err = isWaitForDynamicRouteServerStaticRouteStable(vpcClient, dynamicRouteServerID, *staticRoute.ID, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Waiting for dynamic route server static route to become stable failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIsDynamicRouteServerStaticRouteRead(context, d, meta)
}

// resourceIBMIsDynamicRouteServerStaticRouteParseID splits "<drs>/<static_route>".
func resourceIBMIsDynamicRouteServerStaticRouteParseID(id string) (string, string, error) {
	parts, err := flex.SepIdParts(id, "/")
	if err != nil {
		return "", "", err
	}
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid ID %q, expected <dynamic_route_server_id>/<static_route_id>", id)
	}
	return parts[0], parts[1], nil
}

func resourceIBMIsDynamicRouteServerStaticRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	dynamicRouteServerID, staticRouteID, err := resourceIBMIsDynamicRouteServerStaticRouteParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "sep-id-parts").GetDiag()
	}

	getOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}
	getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	getOptions.SetID(staticRouteID)

	staticRoute, response, err := vpcClient.GetDynamicRouteServerStaticRouteWithContext(context, getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// On import only the ID is known; derive the parent identifier from it.
	if err = d.Set("dynamic_route_server_id", dynamicRouteServerID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dynamic_route_server_id: %s", err), "ibm_is_dynamic_route_server_static_route", "read", "set-dynamic_route_server_id").GetDiag()
	}

	modelMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(staticRoute)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "read", "static-route-to-map").GetDiag()
	}
	for key, value := range modelMap {
		if key == "id" {
			key = "dynamic_route_server_static_route_id"
		}
		if err = d.Set(key, value); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "ibm_is_dynamic_route_server_static_route", "read", "set-"+key).GetDiag()
		}
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

	dynamicRouteServerID, staticRouteID, err := resourceIBMIsDynamicRouteServerStaticRouteParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "sep-id-parts").GetDiag()
	}

	hasChange := false
	patchVals := &vpcv1.DynamicRouteServerStaticRoutePatch{}

	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("next_hops") {
		nextHops, err := resourceIBMIsDynamicRouteServerStaticRouteNextHopsFromList(d.Get("next_hops").([]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "parse-next_hops").GetDiag()
		}
		patchVals.NextHops = nextHops
		hasChange = true
	}
	if d.HasChange("route_delete_delay") {
		patchVals.RouteDeleteDelay = core.Int64Ptr(int64(d.Get("route_delete_delay").(int)))
		hasChange = true
	}
	if d.HasChange("routing_tables") {
		routingTables, err := resourceIBMIsDynamicRouteServerStaticRouteRoutingTablesFromList(d.Get("routing_tables").([]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "parse-routing_tables").GetDiag()
		}
		patchVals.RoutingTables = routingTables
		hasChange = true
	}

	if hasChange {
		patch, err := patchVals.AsPatch()
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "update", "as-patch").GetDiag()
		}

		updateOptions := &vpcv1.UpdateDynamicRouteServerStaticRouteOptions{}
		updateOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		updateOptions.SetID(staticRouteID)
		updateOptions.SetIfMatch(d.Get("etag").(string))
		updateOptions.DynamicRouteServerStaticRoutePatch = patch

		_, _, err = vpcClient.UpdateDynamicRouteServerStaticRouteWithContext(context, updateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		_, err = isWaitForDynamicRouteServerStaticRouteStable(vpcClient, dynamicRouteServerID, staticRouteID, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Waiting for dynamic route server static route to become stable failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "update")
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

	dynamicRouteServerID, staticRouteID, err := resourceIBMIsDynamicRouteServerStaticRouteParseID(d.Id())
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_static_route", "delete", "sep-id-parts").GetDiag()
	}

	deleteOptions := &vpcv1.DeleteDynamicRouteServerStaticRouteOptions{}
	deleteOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	deleteOptions.SetID(staticRouteID)
	if etag, ok := d.GetOk("etag"); ok && etag.(string) != "" {
		deleteOptions.SetIfMatch(etag.(string))
	}

	_, response, err := vpcClient.DeleteDynamicRouteServerStaticRouteWithContext(context, deleteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForDynamicRouteServerStaticRouteDeleted(vpcClient, dynamicRouteServerID, staticRouteID, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Waiting for dynamic route server static route to be deleted failed: %s", err.Error()), "ibm_is_dynamic_route_server_static_route", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func isDynamicRouteServerStaticRouteRefreshFunc(sess *vpcv1.VpcV1, dynamicRouteServerID, staticRouteID string, forDelete bool) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		getOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}
		getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
		getOptions.SetID(staticRouteID)
		staticRoute, response, err := sess.GetDynamicRouteServerStaticRoute(getOptions)
		if err != nil {
			if forDelete && response != nil && response.StatusCode == 404 {
				return struct{}{}, isDynamicRouteServerStaticRouteDeleted, nil
			}
			return nil, "", fmt.Errorf("Error getting dynamic route server static route: %s\n%s", err, response)
		}
		state := ""
		if staticRoute.LifecycleState != nil {
			state = *staticRoute.LifecycleState
		}
		if state == isDynamicRouteServerStaticRouteFailed {
			return staticRoute, state, fmt.Errorf("dynamic route server static route (%s) entered the failed lifecycle state", staticRouteID)
		}
		return staticRoute, state, nil
	}
}

func isWaitForDynamicRouteServerStaticRouteStable(sess *vpcv1.VpcV1, dynamicRouteServerID, staticRouteID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server static route (%s) to be stable.", staticRouteID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerStaticRoutePending, isDynamicRouteServerStaticRouteUpdating, isDynamicRouteServerStaticRouteWaiting, ""},
		Target:     []string{isDynamicRouteServerStaticRouteStable},
		Refresh:    isDynamicRouteServerStaticRouteRefreshFunc(sess, dynamicRouteServerID, staticRouteID, false),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isWaitForDynamicRouteServerStaticRouteDeleted(sess *vpcv1.VpcV1, dynamicRouteServerID, staticRouteID string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for dynamic route server static route (%s) to be deleted.", staticRouteID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isDynamicRouteServerStaticRouteDeleting, isDynamicRouteServerStaticRouteStable, isDynamicRouteServerStaticRoutePending, isDynamicRouteServerStaticRouteUpdating, isDynamicRouteServerStaticRouteWaiting, ""},
		Target:     []string{isDynamicRouteServerStaticRouteDeleted},
		Refresh:    isDynamicRouteServerStaticRouteRefreshFunc(sess, dynamicRouteServerID, staticRouteID, true),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

// ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap flattens a static route into the
// resource attribute map (keys match the schema; "id" is the static route identifier).
func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(model *vpcv1.DynamicRouteServerStaticRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("dynamic route server static route is nil")
	}
	addedRoutes := []map[string]interface{}{}
	for _, item := range model.AddedRoutes {
		itemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		addedRoutes = append(addedRoutes, itemMap)
	}
	modelMap["added_routes"] = addedRoutes
	if model.Destination != nil {
		modelMap["destination"] = *model.Destination
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, item := range model.LifecycleReasons {
		itemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, itemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	if model.LifecycleState != nil {
		modelMap["lifecycle_state"] = *model.LifecycleState
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	nextHops := []map[string]interface{}{}
	for _, item := range model.NextHops {
		itemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(item)
		if err != nil {
			return modelMap, err
		}
		nextHops = append(nextHops, itemMap)
	}
	modelMap["next_hops"] = nextHops
	if model.ResourceType != nil {
		modelMap["resource_type"] = *model.ResourceType
	}
	modelMap["route_delete_delay"] = 0
	if model.RouteDeleteDelay != nil {
		modelMap["route_delete_delay"] = flex.IntValue(model.RouteDeleteDelay)
	}
	routingTables := []map[string]interface{}{}
	for _, item := range model.RoutingTables {
		itemMap, err := ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(&item) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		routingTables = append(routingTables, itemMap)
	}
	modelMap["routing_tables"] = routingTables
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototype(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeIntf, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(modelMap)
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityIntf, error) {
	id, _ := modelMap["id"].(string)
	href, _ := modelMap["href"].(string)
	if id != "" {
		return ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(modelMap)
	}
	if href != "" {
		return ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(modelMap)
	}
	return nil, fmt.Errorf("next_hops item requires id or href")
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("next hop identity requires id")
	}
	return &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("next hop identity requires href")
	}
	return &vpcv1.DynamicRouteServerStaticRouteNextHopPrototypeDynamicRouteServerPeerGroupIdentityDynamicRouteServerPeerGroupIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToDynamicRouteServerStaticRouteRoutingTablePrototype(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype, error) {
	model := &vpcv1.DynamicRouteServerStaticRouteRoutingTablePrototype{}
	advertise, _ := modelMap["advertise"].(bool)
	model.Advertise = core.BoolPtr(advertise)
	list, _ := modelMap["vpc_routing_table"].([]interface{})
	if len(list) == 0 || list[0] == nil {
		return model, fmt.Errorf("routing_tables item requires vpc_routing_table")
	}
	identityMap, ok := list[0].(map[string]interface{})
	if !ok {
		return model, fmt.Errorf("routing_tables item requires vpc_routing_table")
	}
	identity, err := ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(identityMap)
	if err != nil {
		return model, err
	}
	model.VPCRoutingTable = identity
	return model, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentity(modelMap map[string]interface{}) (vpcv1.RoutingTableIdentityIntf, error) {
	id, _ := modelMap["id"].(string)
	crn, _ := modelMap["crn"].(string)
	href, _ := modelMap["href"].(string)
	if id != "" {
		return ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(modelMap)
	}
	if crn != "" {
		return ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(modelMap)
	}
	if href != "" {
		return ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(modelMap)
	}
	return nil, fmt.Errorf("vpc_routing_table requires id, crn or href")
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByID(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByID, error) {
	id, _ := modelMap["id"].(string)
	if id == "" {
		return nil, fmt.Errorf("routing table identity requires id")
	}
	return &vpcv1.RoutingTableIdentityByID{ID: core.StringPtr(id)}, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByCRN, error) {
	crn, _ := modelMap["crn"].(string)
	if crn == "" {
		return nil, fmt.Errorf("routing table identity requires crn")
	}
	return &vpcv1.RoutingTableIdentityByCRN{CRN: core.StringPtr(crn)}, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteMapToRoutingTableIdentityByHref(modelMap map[string]interface{}) (*vpcv1.RoutingTableIdentityByHref, error) {
	href, _ := modelMap["href"].(string)
	if href == "" {
		return nil, fmt.Errorf("routing table identity requires href")
	}
	return &vpcv1.RoutingTableIdentityByHref{Href: core.StringPtr(href)}, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model vpcv1.DynamicRouteServerStaticRouteNextHopIntf) (map[string]interface{}, error) {
	switch nextHop := model.(type) {
	case *vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference:
		return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(nextHop)
	case *vpcv1.DynamicRouteServerStaticRouteNextHop:
		if nextHop == nil {
			return nil, fmt.Errorf("next hop is nil")
		}
		return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(&vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference{
			Deleted:      nextHop.Deleted,
			Href:         nextHop.Href,
			ID:           nextHop.ID,
			Name:         nextHop.Name,
			ResourceType: nextHop.ResourceType,
		})
	default:
		return nil, fmt.Errorf("Unrecognized vpcv1.DynamicRouteServerStaticRouteNextHopIntf subtype encountered")
	}
}

func ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("deleted is nil")
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("next hop is nil")
	}
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
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model *vpcv1.DynamicRouteServerStaticRouteRoutingTable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("routing table is nil")
	}
	modelMap["advertise"] = false
	if model.Advertise != nil {
		modelMap["advertise"] = *model.Advertise
	}
	if model.VPCRoutingTable != nil {
		vpcRoutingTableMap, err := ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model.VPCRoutingTable)
		if err != nil {
			return modelMap, err
		}
		modelMap["vpc_routing_table"] = []map[string]interface{}{vpcRoutingTableMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("routing table reference is nil")
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
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
}

func ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model *vpcv1.DynamicRouteServerStaticRouteAddedRoute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("added route is nil")
	}
	if model.VPCRoutingTableRoute != nil {
		routeMap, err := ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model.VPCRoutingTableRoute)
		if err != nil {
			return modelMap, err
		}
		modelMap["vpc_routing_table_route"] = []map[string]interface{}{routeMap}
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model *vpcv1.RouteReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("route reference is nil")
	}
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
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model == nil {
		return modelMap, fmt.Errorf("lifecycle reason is nil")
	}
	if model.Code != nil {
		modelMap["code"] = *model.Code
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}
