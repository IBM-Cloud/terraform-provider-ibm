// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	rtID                         = "routing_table"
	rtVpcID                      = "vpc"
	rtName                       = "name"
	rtRouteDirectLinkIngress     = "route_direct_link_ingress"
	rtRouteInternetIngress       = "route_internet_ingress"
	rtRouteTransitGatewayIngress = "route_transit_gateway_ingress"
	rtRouteVPCZoneIngress        = "route_vpc_zone_ingress"
	rtCreateAt                   = "created_at"
	rtHref                       = "href"
	rtIsDefault                  = "is_default"
	rtResourceType               = "resource_type"
	rtLifecycleState             = "lifecycle_state"
	rtSubnets                    = "subnets"
	rtDestination                = "destination"
	rtAction                     = "action"
	rtNextHop                    = "next_hop"
	rtZone                       = "zone"
	rtOrigin                     = "origin"
	rtResourceGroup              = "resource_group"
	rtResourceGroupHref          = "href"
	rtResourceGroupId            = "id"
	rtResourceGroupName          = "name"
	rtAccessTags                 = "access_tags"
	rtAccessTagType              = "access"
	rtTags                       = "tags"
	rtUserTagType                = "user"
)

func ResourceIBMISVPCRoutingTable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPCRoutingTableCreate,
		ReadContext:   resourceIBMISVPCRoutingTableRead,
		UpdateContext: resourceIBMISVPCRoutingTableUpdate,
		DeleteContext: resourceIBMISVPCRoutingTableDelete,
		Exists:        resourceIBMISVPCRoutingTableExists,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),
		Schema: map[string]*schema.Schema{
			rtVpcID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPC identifier.",
			},
			"accept_routes_from_resource_type": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The filters specifying the resources that may create routes in this routing table, The resource type: vpn_gateway or vpn_server",
			},
			"advertise_routes_to": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         schema.HashString,
				Description: "The ingress sources to advertise routes to. Routes in the table with `advertise` enabled will be advertised to these sources.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			rtRouteDirectLinkIngress: {
				Type:        schema.TypeBool,
				ForceNew:    false,
				Default:     false,
				Optional:    true,
				Description: "If set to true, this routing table will be used to route traffic that originates from Direct Link to this VPC.",
			},
			rtRouteInternetIngress: {
				Type:        schema.TypeBool,
				ForceNew:    false,
				Default:     false,
				Optional:    true,
				Description: "If set to true, this routing table will be used to route traffic that originates from the internet. For this to succeed, the VPC must not already have a routing table with this property set to true.",
			},
			rtRouteTransitGatewayIngress: {
				Type:        schema.TypeBool,
				ForceNew:    false,
				Default:     false,
				Optional:    true,
				Description: "If set to true, this routing table will be used to route traffic that originates from Transit Gateway to this VPC.",
			},
			rtRouteVPCZoneIngress: {
				Type:        schema.TypeBool,
				ForceNew:    false,
				Default:     false,
				Optional:    true,
				Description: "If set to true, this routing table will be used to route traffic that originates from subnets in other zones in this VPC.",
			},
			rtName: {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     false,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_vpc_routing_table", rtName),
				Description:  "The user-defined name for this routing table.",
			},
			rtID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The routing table identifier.",
			},
			rtCrn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The routing table CRN.",
			},
			rtHref: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table Href",
			},
			rtResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table Resource Type",
			},
			rtCreateAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table Created At",
			},
			rtLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Routing table Lifecycle State",
			},
			rtIsDefault: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this is the default routing table for this VPC",
			},
			rtSubnets: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						rtName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet name",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet ID",
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
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_routing_table", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags",
			},

			rtAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_vpc_routing_table", "accesstag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags",
			},
		},
	}
}

func ResourceIBMISVPCRoutingTableValidator() *validate.ResourceValidator {

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

	validateSchema = append(validateSchema, validate.ValidateSchema{
		Identifier:                 "accesstag",
		ValidateFunctionIdentifier: validate.ValidateRegexpLen,
		Type:                       validate.TypeString,
		Optional:                   true,
		Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
		MinValueLength:             1,
		MaxValueLength:             128,
	})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 rtAction,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   false,
			AllowedValues:              actionAllowedValues})

	ibmISVPCRoutingTableValidator := validate.ResourceValidator{ResourceName: "ibm_is_vpc_routing_table", Schema: validateSchema}
	return &ibmISVPCRoutingTableValidator
}

func resourceIBMISVPCRoutingTableCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpcID := d.Get(rtVpcID).(string)
	rtName := d.Get(rtName).(string)
	// acptresourcetype

	createVpcRoutingTableOptions := sess.NewCreateVPCRoutingTableOptions(vpcID)
	createVpcRoutingTableOptions.SetName(rtName)
	if _, ok := d.GetOk(rtRouteDirectLinkIngress); ok {
		routeDirectLinkIngress := d.Get(rtRouteDirectLinkIngress).(bool)
		createVpcRoutingTableOptions.RouteDirectLinkIngress = &routeDirectLinkIngress
	}

	if acceptRoutesFrom, ok := d.GetOk("accept_routes_from_resource_type"); ok {
		var aroutes []vpcv1.ResourceFilter
		acptRoutes := acceptRoutesFrom.(*schema.Set)
		for _, val := range acptRoutes.List() {
			value := val.(string)
			resourceFilter := vpcv1.ResourceFilter{
				ResourceType: &value,
			}
			aroutes = append(aroutes, resourceFilter)
		}
		createVpcRoutingTableOptions.AcceptRoutesFrom = aroutes
	}
	if _, ok := d.GetOk("advertise_routes_to"); ok {
		var advertiseRoutesToList []string
		advertiseRoutesTo := d.Get("advertise_routes_to").(*schema.Set)

		for _, val := range advertiseRoutesTo.List() {
			advertiseRoutesToList = append(advertiseRoutesToList, val.(string))
		}
		createVpcRoutingTableOptions.AdvertiseRoutesTo = advertiseRoutesToList
	}

	if _, ok := d.GetOk(rtRouteInternetIngress); ok {
		rtRouteInternetIngress := d.Get(rtRouteInternetIngress).(bool)
		createVpcRoutingTableOptions.RouteInternetIngress = &rtRouteInternetIngress
	}
	if _, ok := d.GetOk(rtRouteTransitGatewayIngress); ok {
		routeTransitGatewayIngress := d.Get(rtRouteTransitGatewayIngress).(bool)
		createVpcRoutingTableOptions.RouteTransitGatewayIngress = &routeTransitGatewayIngress
	}
	if _, ok := d.GetOk(rtRouteVPCZoneIngress); ok {
		routeVPCZoneIngress := d.Get(rtRouteVPCZoneIngress).(bool)
		createVpcRoutingTableOptions.RouteVPCZoneIngress = &routeVPCZoneIngress
	}
	routeTable, _, err := sess.CreateVPCRoutingTableWithContext(context, createVpcRoutingTableOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVPCRoutingTableWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", vpcID, *routeTable.ID))

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(rtTags); ok || v != "" {
		oldList, newList := d.GetChange(rtTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *routeTable.CRN, "", rtUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of resource routing table (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(rtAccessTags); ok {
		oldList, newList := d.GetChange(rtAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *routeTable.CRN, "", rtAccessTags)
		if err != nil {
			log.Printf(
				"Error on create of resource routing table (%s) access tags: %s", d.Id(), err)
		}
	}

	return resourceIBMISVPCRoutingTableRead(context, d, meta)
}

func resourceIBMISVPCRoutingTableRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idSet := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(idSet[0], idSet[1])
	routeTable, response, err := sess.GetVPCRoutingTableWithContext(context, getVpcRoutingTableOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCRoutingTableWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set(rtVpcID, idSet[0])

	if err = d.Set(rtID, routeTable.ID); err != nil {
		err = fmt.Errorf("Error setting routing_table: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-routing_table").GetDiag()
	}
	if err = d.Set("crn", routeTable.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(routeTable.Name) {
		if err = d.Set("name", routeTable.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-name").GetDiag()
		}
	}
	if err = d.Set("href", routeTable.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-href").GetDiag()
	}
	if err = d.Set("lifecycle_state", routeTable.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(routeTable.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("resource_type", routeTable.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(routeTable.RouteDirectLinkIngress) {
		if err = d.Set("route_direct_link_ingress", routeTable.RouteDirectLinkIngress); err != nil {
			err = fmt.Errorf("Error setting route_direct_link_ingress: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-route_direct_link_ingress").GetDiag()
		}
	}
	if !core.IsNil(routeTable.RouteInternetIngress) {
		if err = d.Set("route_internet_ingress", routeTable.RouteInternetIngress); err != nil {
			err = fmt.Errorf("Error setting route_internet_ingress: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-route_internet_ingress").GetDiag()
		}
	}
	if !core.IsNil(routeTable.RouteTransitGatewayIngress) {
		if err = d.Set("route_transit_gateway_ingress", routeTable.RouteTransitGatewayIngress); err != nil {
			err = fmt.Errorf("Error setting route_transit_gateway_ingress: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-route_transit_gateway_ingress").GetDiag()
		}
	}
	if !core.IsNil(routeTable.RouteVPCZoneIngress) {
		if err = d.Set("route_vpc_zone_ingress", routeTable.RouteVPCZoneIngress); err != nil {
			err = fmt.Errorf("Error setting route_vpc_zone_ingress: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-route_vpc_zone_ingress").GetDiag()
		}
	}
	if err = d.Set("is_default", routeTable.IsDefault); err != nil {
		err = fmt.Errorf("Error setting is_default: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-is_default").GetDiag()
	}
	acceptRoutesFromArray := make([]string, 0)
	advertiseRoutesToArray := make([]string, 0)
	for i := 0; i < len(routeTable.AcceptRoutesFrom); i++ {
		acceptRoutesFromArray = append(acceptRoutesFromArray, string(*(routeTable.AcceptRoutesFrom[i].ResourceType)))
	}
	if err = d.Set("accept_routes_from_resource_type", acceptRoutesFromArray); err != nil {
		err = fmt.Errorf("Error setting accept_routes_from_resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-accept_routes_from_resource_type").GetDiag()
	}

	for i := 0; i < len(routeTable.AdvertiseRoutesTo); i++ {
		advertiseRoutesToArray = append(advertiseRoutesToArray, routeTable.AdvertiseRoutesTo[i])
	}

	if err = d.Set("advertise_routes_to", advertiseRoutesToArray); err != nil {
		err = fmt.Errorf("Error setting advertise_routes_to: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-advertise_routes_to").GetDiag()
	}

	subnets := make([]map[string]interface{}, 0)

	for _, s := range routeTable.Subnets {
		subnet := make(map[string]interface{})
		subnet[ID] = *s.ID
		subnet["name"] = *s.Name
		subnets = append(subnets, subnet)
	}

	if err = d.Set("subnets", subnets); err != nil {
		err = fmt.Errorf("Error setting subnets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "read", "set-subnets").GetDiag()
	}
	resourceGroupList := []map[string]interface{}{}
	if routeTable.ResourceGroup != nil {
		resourceGroupMap := routingTableResourceGroupToMap(*routeTable.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
	}

	if err = d.Set(rtResourceGroup, resourceGroupList); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_group: %s", err), "ibm_is_vpc_routing_table", "read", "set-resource_group").GetDiag()
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *routeTable.CRN, "", rtUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource routing table (%s) tags: %s", d.Id(), err)
	}
	if err = d.Set(rtTags, tags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting tags: %s", err), "ibm_is_vpc_routing_table", "read", "set-tags").GetDiag()
	}
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *routeTable.CRN, "", rtAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of resource routing table (%s) access tags: %s", d.Id(), err)
	}
	if err = d.Set(rtAccessTags, accesstags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting access_tags: %s", err), "ibm_is_vpc_routing_table", "read", "set-access_tags").GetDiag()
	}
	return nil
}

func routingTableResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap[isVolumesResourceGroupHref] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap[isVolumesResourceGroupId] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap[isVolumesResourceGroupName] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func resourceIBMISVPCRoutingTableUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	//Etag
	idSett := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(idSett[0], idSett[1])
	routingTableGet, respGet, err := sess.GetVPCRoutingTableWithContext(context, getVpcRoutingTableOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCRoutingTableWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	eTag := respGet.Headers.Get("ETag")

	if d.HasChange(rtAccessTags) {
		oldList, newList := d.GetChange(rtAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *routingTableGet.CRN, "", rtAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource routing table (%s) access tags: %s", d.Id(), err)
		}
	}

	if d.HasChange(rtTags) {
		oldList, newList := d.GetChange(rtTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *routingTableGet.CRN, "", rtUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource routing table (%s) tags: %s", d.Id(), err)
		}
	}

	idSet := strings.Split(d.Id(), "/")
	updateVpcRoutingTableOptions := new(vpcv1.UpdateVPCRoutingTableOptions)
	updateVpcRoutingTableOptions.VPCID = &idSet[0]
	updateVpcRoutingTableOptions.ID = &idSet[1]
	hasChange := false
	// Construct an instance of the RoutingTablePatch model
	routingTablePatchModel := new(vpcv1.RoutingTablePatch)

	if d.HasChange(rtName) {
		name := d.Get(rtName).(string)
		routingTablePatchModel.Name = core.StringPtr(name)
		hasChange = true
	}
	removeAcceptRoutesFromFilter := false
	if d.HasChange("accept_routes_from_resource_type") {
		var aroutes []vpcv1.ResourceFilter
		acptRoutes := d.Get("accept_routes_from_resource_type").(*schema.Set)
		if len(acptRoutes.List()) == 0 {
			removeAcceptRoutesFromFilter = true
		} else {
			for _, val := range acptRoutes.List() {
				value := val.(string)
				resourceFilter := vpcv1.ResourceFilter{
					ResourceType: &value,
				}
				aroutes = append(aroutes, resourceFilter)
			}
		}
		routingTablePatchModel.AcceptRoutesFrom = aroutes
		hasChange = true
	}
	removeAdvertiseRoutesTo := false
	if d.HasChange("advertise_routes_to") {
		var advertiseRoutesToList []string
		advertiseRoutesTo := d.Get("advertise_routes_to").(*schema.Set)

		if len(advertiseRoutesTo.List()) == 0 {
			removeAdvertiseRoutesTo = true
		} else {
			for _, val := range advertiseRoutesTo.List() {
				advertiseRoutesToList = append(advertiseRoutesToList, val.(string))
			}
		}

		routingTablePatchModel.AdvertiseRoutesTo = advertiseRoutesToList
		hasChange = true
	}
	if d.HasChange(rtRouteDirectLinkIngress) {
		routeDirectLinkIngress := d.Get(rtRouteDirectLinkIngress).(bool)
		routingTablePatchModel.RouteDirectLinkIngress = core.BoolPtr(routeDirectLinkIngress)
		hasChange = true
	}
	if d.HasChange(rtRouteInternetIngress) {
		rtRouteInternetIngress := d.Get(rtRouteInternetIngress).(bool)
		routingTablePatchModel.RouteInternetIngress = core.BoolPtr(rtRouteInternetIngress)
	}
	if d.HasChange(rtRouteTransitGatewayIngress) {
		routeTransitGatewayIngress := d.Get(rtRouteTransitGatewayIngress).(bool)
		routingTablePatchModel.RouteTransitGatewayIngress = core.BoolPtr(routeTransitGatewayIngress)
		hasChange = true
	}
	if d.HasChange(rtRouteVPCZoneIngress) {
		routeVPCZoneIngress := d.Get(rtRouteVPCZoneIngress).(bool)
		routingTablePatchModel.RouteVPCZoneIngress = core.BoolPtr(routeVPCZoneIngress)
		hasChange = true
	}
	if hasChange {
		updateVpcRoutingTableOptions.IfMatch = &eTag
	}

	routingTablePatchModelAsPatch, asPatchErr := routingTablePatchModel.AsPatch()
	if asPatchErr != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("routingTablePatchModel.AsPatch() failed: %s", asPatchErr.Error()), "ibm_is_vpc_routing_table", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if removeAdvertiseRoutesTo {
		routingTablePatchModelAsPatch["advertise_routes_to"] = []string{}
	}
	if removeAcceptRoutesFromFilter {
		routingTablePatchModelAsPatch["accept_routes_from"] = []vpcv1.ResourceFilter{}
	}
	updateVpcRoutingTableOptions.RoutingTablePatch = routingTablePatchModelAsPatch
	_, _, err = sess.UpdateVPCRoutingTableWithContext(context, updateVpcRoutingTableOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVPCRoutingTableWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	return resourceIBMISVPCRoutingTableRead(context, d, meta)
}

func resourceIBMISVPCRoutingTableDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idSet := strings.Split(d.Id(), "/")

	deleteTableOptions := sess.NewDeleteVPCRoutingTableOptions(idSet[0], idSet[1])
	response, err := sess.DeleteVPCRoutingTableWithContext(context, deleteTableOptions)
	if err != nil && response.StatusCode != 404 {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVPCRoutingTableWithContext failed: %s", err.Error()), "ibm_is_vpc_routing_table", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}

func resourceIBMISVPCRoutingTableExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpc_routing_table", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}

	idSet := strings.Split(d.Id(), "/")
	if len(idSet) != 2 {
		return false, fmt.Errorf("[ERROR] Incorrect ID %s: ID should be a combination of vpcID/routingTableID", d.Id())
	}
	getVpcRoutingTableOptions := sess.NewGetVPCRoutingTableOptions(idSet[0], idSet[1])
	_, response, err := sess.GetVPCRoutingTable(getVpcRoutingTableOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return false, nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCRoutingTable failed: %s", err.Error()), "ibm_is_vpc_routing_table", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return false, tfErr
	}
	return true, nil
}
