// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	rID          = "route_id"
	rDestination = "destination"
	rAction      = "action"
	rNextHop     = "next_hop"
	rName        = "name"
	rZone        = "zone"
)

func ResourceIBMISVPCRoutingTableRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPCRoutingTableRouteCreate,
		ReadContext:   resourceIBMISVPCRoutingTableRouteRead,
		UpdateContext: resourceIBMISVPCRoutingTableRouteUpdate,
		DeleteContext: resourceIBMISVPCRoutingTableRouteDelete,
		Exists:        resourceIBMISVPCRoutingTableRouteExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			rtID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The routing table identifier.",
			},
			rtVpcID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC identifier.",
			},
			rDestination: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The destination of the route.",
			},
			rZone: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone to apply the route to. Traffic from subnets in this zone will be subject to this route.",
			},
			rNextHop: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "If action is deliver, the next hop that packets will be delivered to. For other action values, its address will be 0.0.0.0.",
			},
			rAction: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "deliver",
				Description:  "The action to perform with a packet matching the route.",
				ValidateFunc: validate.InvokeValidator("ibm_is_vpc_routing_table_route", rAction),
			},
			"advertise": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicates whether this route will be advertised to the ingress sources specified by the `advertise_routes_to` routing table property.",
			},
			rName: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     false,
				Computed:     true,
				Description:  "The user-defined name for this route.",
				ValidateFunc: validate.InvokeValidator("ibm_is_vpc_routing_table_route", rName),
			},
			rID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The routing table route identifier.",
			},
			rtHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table route Href",
			},
			rtCreateAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table route Created At",
			},
			"creator": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "If present, the resource that created the route. Routes with this property present cannot bedirectly deleted. All routes with an `origin` of `learned` or `service` will have thisproperty set, and future `origin` values may also have this property set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPN gateway's CRN.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPN gateway's canonical URL.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this VPN gateway.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user-defined name for this VPN gateway.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			rtLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table route Lifecycle State",
			},
			"priority": {
				Type:         schema.TypeInt,
				Computed:     true,
				Optional:     true,
				Description:  "The route's priority. Smaller values have higher priority.",
				ValidateFunc: validate.InvokeValidator("ibm_is_vpc_routing_table_route", "priority"),
			},
			rtOrigin: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The origin of this route.",
			},
		},
	}
}

func ResourceIBMISVPCRoutingTableRouteValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	actionAllowedValues := "delegate, delegate_vpc, deliver, drop"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 rtName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   false,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 rAction,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   false,
			AllowedValues:              actionAllowedValues})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "priority",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			MinValue:                   "0",
			MaxValue:                   "4"})
	ibmVPCRoutingTableRouteValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpc_routing_table_route", Schema: validateSchema}
	return &ibmVPCRoutingTableRouteValidator
}

func resourceIBMISVPCRoutingTableRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table_route", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := d.Get(rtVpcID).(string)
	tableID := d.Get(rtID).(string)
	destination := d.Get(rDestination).(string)
	zone := d.Get(rZone).(string)
	z := &vpcv1.ZoneIdentityByName{
		Name: core.StringPtr(zone),
	}

	createVpcRoutingTableRouteOptions := sess.NewCreateVPCRoutingTableRouteOptions(vpcID, tableID, destination, z)
	createVpcRoutingTableRouteOptions.SetZone(z)
	createVpcRoutingTableRouteOptions.SetDestination(destination)

	if add, ok := d.GetOk(rNextHop); ok {
		item := add.(string)
		if net.ParseIP(item) == nil {
			nhConnectionID := &vpcv1.RouteNextHopPrototype{
				ID: core.StringPtr(item),
			}
			createVpcRoutingTableRouteOptions.SetNextHop(nhConnectionID)
		} else {
			nh := &vpcv1.RouteNextHopPrototype{
				Address: core.StringPtr(item),
			}
			createVpcRoutingTableRouteOptions.SetNextHop(nh)
		}
	}

	if action, ok := d.GetOk(rAction); ok {
		routeAction := action.(string)
		createVpcRoutingTableRouteOptions.SetAction(routeAction)
	}

	if advertiseVal, ok := d.GetOk("advertise"); ok {
		advertise := advertiseVal.(bool)
		createVpcRoutingTableRouteOptions.SetAdvertise(advertise)
	}

	if name, ok := d.GetOk(rName); ok {
		routeName := name.(string)
		createVpcRoutingTableRouteOptions.SetName(routeName)
	}

	// Using GetOkExists to detet 0 as the possible values.
	if priority, ok := d.GetOkExists("priority"); ok {
		routePriority := priority.(int)
		createVpcRoutingTableRouteOptions.SetPriority(int64(routePriority))
	}

	route, _, err := sess.CreateVPCRoutingTableRouteWithContext(context, createVpcRoutingTableRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPCRoutingTableWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table_route", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", vpcID, tableID, *route.ID))
	d.Set(rID, *route.ID)
	return resourceIBMISVPCRoutingTableRouteRead(context, d, meta)
}

func resourceIBMISVPCRoutingTableRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idSet := strings.Split(d.Id(), "/")
	getVpcRoutingTableRouteOptions := sess.NewGetVPCRoutingTableRouteOptions(idSet[0], idSet[1], idSet[2])
	route, response, err := sess.GetVPCRoutingTableRouteWithContext(context, getVpcRoutingTableRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCRoutingTableRouteWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set(rID, *route.ID)
	if route.Advertise != nil {
		d.Set("Advertise", route.Advertise)
	}
	if err = d.Set(rName, *route.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-name").GetDiag()
	}
	if err = d.Set(rDestination, *route.Destination); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-destination").GetDiag()
	}

	if route.NextHop != nil {
		nexthop := route.NextHop.(*vpcv1.RouteNextHop)
		if nexthop.Address != nil {

			if err = d.Set(rNextHop, *nexthop.Address); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting next_hop: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-next_hop").GetDiag()
			}
		}
		if nexthop.ID != nil {
			if err = d.Set(rNextHop, *nexthop.ID); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting next_hop: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-next_hop").GetDiag()
			}
		}
	}
	if err = d.Set("origin", route.Origin); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting origin: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-origin").GetDiag()
	}
	if route.Zone != nil {
		if err = d.Set(rZone, *route.Zone.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting zone: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-zone").GetDiag()
		}
	}

	if err = d.Set(rtHref, route.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-href").GetDiag()
	}
	if err = d.Set(rtLifecycleState, route.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set(rtCreateAt, route.CreatedAt.String()); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-created_at").GetDiag()
	}

	creator := []map[string]interface{}{}
	if route.Creator != nil {
		mm, err := dataSourceIBMIsRouteCreatorToMap(route.Creator)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "creator-to-map").GetDiag()
		}
		creator = append(creator, mm)
	}

	if err = d.Set("creator", creator); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting creator: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-creator").GetDiag()
	}

	if err = d.Set("priority", route.Priority); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting priority: %s", err), "ibm_is_vpc_routing_table_route", "read", "set-priority").GetDiag()
	}
	return nil
}

func resourceIBMISVPCRoutingTableRouteUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table_route", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idSet := strings.Split(d.Id(), "/")
	hasChange := false
	routePatch := make(map[string]interface{})
	updateVpcRoutingTableRouteOptions := sess.NewUpdateVPCRoutingTableRouteOptions(idSet[0], idSet[1], idSet[2], routePatch)

	// Construct an instance of the RoutePatch model
	routePatchModel := new(vpcv1.RoutePatch)
	if d.HasChange("advertise") {
		advertiseVal := d.Get("advertise").(bool)
		routePatchModel.Advertise = &advertiseVal
		hasChange = true

	}
	if d.HasChange(rName) {
		name := d.Get(rName).(string)
		routePatchModel.Name = &name
		hasChange = true
	}
	if d.HasChange("priority") {
		rp := d.Get("priority").(int)
		routePriority := int64(rp)
		routePatchModel.Priority = &routePriority
		hasChange = true
	}

	if d.HasChange(rNextHop) {
		if add, ok := d.GetOk(rNextHop); ok {
			item := add.(string)
			if net.ParseIP(item) == nil {
				routePatchModel.NextHop = &vpcv1.RouteNextHopPatch{
					ID: core.StringPtr(item),
				}
				hasChange = true
			} else {
				routePatchModel.NextHop = &vpcv1.RouteNextHopPatch{
					Address: core.StringPtr(item),
				}
				hasChange = true
			}
		}
	}
	if hasChange {
		routePatchModelAsPatch, patchErr := routePatchModel.AsPatch()
		if patchErr != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("routePatchModel.AsPatch() failed: %s", patchErr.Error()), "ibm_is_vpc_routing_table_route", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateVpcRoutingTableRouteOptions.RoutePatch = routePatchModelAsPatch
		_, _, err = sess.UpdateVPCRoutingTableRouteWithContext(context, updateVpcRoutingTableRouteOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVPCRoutingTableRouteWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table_route", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISVPCRoutingTableRouteRead(context, d, meta)
}

func resourceIBMISVPCRoutingTableRouteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table_route", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idSet := strings.Split(d.Id(), "/")
	deleteVpcRoutingTableRouteOptions := sess.NewDeleteVPCRoutingTableRouteOptions(idSet[0], idSet[1], idSet[2])
	response, err := sess.DeleteVPCRoutingTableRouteWithContext(context, deleteVpcRoutingTableRouteOptions)
	if err != nil && response.StatusCode != 404 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVPCRoutingTableRouteWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table_route", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}

func resourceIBMISVPCRoutingTableRouteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table_route", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	idSet := strings.Split(d.Id(), "/")
	if len(idSet) != 3 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of vpcID/routingTableID/routeID", d.Id())
	}
	getVpcRoutingTableRouteOptions := sess.NewGetVPCRoutingTableRouteOptions(idSet[0], idSet[1], idSet[2])
	_, response, err := sess.GetVPCRoutingTableRoute(getVpcRoutingTableRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCRoutingTableRoute failed: %s", err.Error()), "ibm_is_vpc_routing_table_route", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
