// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISSubnetRoutingTableAttachment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISSubnetRoutingTableAttachmentCreate,
		ReadContext:   resourceIBMISSubnetRoutingTableAttachmentRead,
		UpdateContext: resourceIBMISSubnetRoutingTableAttachmentUpdate,
		DeleteContext: resourceIBMISSubnetRoutingTableAttachmentDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isSubnetID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet identifier",
			},

			isRoutingTableID: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{isRoutingTableID, isRoutingTableCrn},
				Description:  "The unique identifier of routing table",
			},

			isRoutingTableCrn: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{isRoutingTableID, isRoutingTableCrn},
				Description:  "The crn of routing table",
			},

			rtRouteDirectLinkIngress: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, this routing table will be used to route traffic that originates from Direct Link to this VPC.",
			},

			rtIsDefault: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this is the default routing table for this VPC",
			},
			rtLifecycleState: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "he lifecycle state of the routing table [ deleting, failed, pending, stable, suspended, updating, waiting ]",
			},

			isRoutingTableName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the routing table",
			},
			isRoutingTableResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type",
			},

			rtRouteTransitGatewayIngress: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, this routing table will be used to route traffic that originates from Transit Gateway to this VPC.",
			},

			rtRouteVPCZoneIngress: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If true, this routing table will be used to route traffic that originates from subnets in other zones in this VPC.",
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

			rtRoutes: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						rtName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "route name",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "route ID",
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
		},
	}
}

func resourceIBMISSubnetRoutingTableAttachmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	subnet := d.Get(isSubnetID).(string)
	routingTableID := d.Get(isRoutingTableID).(string)
	routingTableCrn := d.Get(isRoutingTableCrn).(string)
	replaceSubnetRoutingTableOptionsModel := new(vpcv1.ReplaceSubnetRoutingTableOptions)
	replaceSubnetRoutingTableOptionsModel.ID = &subnet

	if routingTableID != "" {
		// Construct an instance of the RoutingTableIdentityByID model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = &routingTableID

		// Construct an instance of the ReplaceSubnetRoutingTableOptions model
		replaceSubnetRoutingTableOptionsModel.RoutingTableIdentity = routingTableIdentityModel
	}

	if routingTableCrn != "" {
		// Construct an instance of the RoutingTableIdentityByID model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByCRN)
		routingTableIdentityModel.CRN = &routingTableCrn

		// Construct an instance of the ReplaceSubnetRoutingTableOptions model
		replaceSubnetRoutingTableOptionsModel.RoutingTableIdentity = routingTableIdentityModel
	}

	resultRT, _, err := sess.ReplaceSubnetRoutingTableWithContext(context, replaceSubnetRoutingTableOptionsModel)

	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSubnetReservedIPWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(subnet)
	log.Printf("[INFO] Routing Table : %s", *resultRT.ID)
	log.Printf("[INFO] Subnet ID : %s", subnet)

	return resourceIBMISSubnetRoutingTableAttachmentRead(context, d, meta)
}

func resourceIBMISSubnetRoutingTableAttachmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSubnetRoutingTableOptionsModel := &vpcv1.GetSubnetRoutingTableOptions{
		ID: &id,
	}
	subRT, response, err := sess.GetSubnetRoutingTableWithContext(context, getSubnetRoutingTableOptionsModel)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetRoutingTableWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set(isRoutingTableName, *subRT.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-name").GetDiag()
	}
	if err = d.Set(isSubnetID, id); err != nil {
		err = fmt.Errorf("Error setting subnet: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-subnet").GetDiag()
	}
	if err = d.Set(isRoutingTableID, *subRT.ID); err != nil {
		err = fmt.Errorf("Error setting routing_table: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-routing_table").GetDiag()
	}
	if err = d.Set(isRoutingTableCrn, *subRT.CRN); err != nil {
		err = fmt.Errorf("Error setting routing_table_crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-routing_table_crn").GetDiag()
	}
	if err = d.Set(isRoutingTableResourceType, *subRT.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set(rtRouteDirectLinkIngress, *subRT.RouteDirectLinkIngress); err != nil {
		err = fmt.Errorf("Error setting route_direct_link_ingress: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-route_direct_link_ingress").GetDiag()
	}
	if err = d.Set(rtIsDefault, *subRT.IsDefault); err != nil {
		err = fmt.Errorf("Error setting is_default: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-is_default").GetDiag()
	}
	if err = d.Set(rtLifecycleState, *subRT.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set(isRoutingTableResourceType, *subRT.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set(rtRouteTransitGatewayIngress, *subRT.RouteTransitGatewayIngress); err != nil {
		err = fmt.Errorf("Error setting route_transit_gateway_ingress: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-route_transit_gateway_ingress").GetDiag()
	}
	if err = d.Set(rtRouteVPCZoneIngress, *subRT.RouteVPCZoneIngress); err != nil {
		err = fmt.Errorf("Error setting route_vpc_zone_ingress: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-route_vpc_zone_ingress").GetDiag()
	}
	subnets := make([]map[string]interface{}, 0)

	for _, s := range subRT.Subnets {
		subnet := make(map[string]interface{})
		subnet[ID] = *s.ID
		subnet["name"] = *s.Name
		subnets = append(subnets, subnet)
	}
	if err = d.Set(rtSubnets, subnets); err != nil {
		err = fmt.Errorf("Error setting subnets: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-subnets").GetDiag()
	}
	routes := make([]map[string]interface{}, 0)
	for _, s := range subRT.Routes {
		route := make(map[string]interface{})
		route[ID] = *s.ID
		route["name"] = *s.Name
		routes = append(routes, route)
	}
	if err = d.Set(rtRoutes, routes); err != nil {
		err = fmt.Errorf("Error setting routes: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-routes").GetDiag()
	}
	resourceGroupList := []map[string]interface{}{}
	if subRT.ResourceGroup != nil {
		resourceGroupMap := routingTableResourceGroupToMap(*subRT.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
	}
	if err = d.Set(rtResourceGroup, resourceGroupList); err != nil {
		err = fmt.Errorf("Error setting resource_group: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "read", "set-resource_group").GetDiag()
	}
	return nil
}

func resourceIBMISSubnetRoutingTableAttachmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange(isRoutingTableID) {
		subnet := d.Get(isSubnetID).(string)
		routingTable := d.Get(isRoutingTableID).(string)

		// Construct an instance of the RoutingTableIdentityByID model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = &routingTable

		// Construct an instance of the ReplaceSubnetRoutingTableOptions model
		replaceSubnetRoutingTableOptionsModel := new(vpcv1.ReplaceSubnetRoutingTableOptions)
		replaceSubnetRoutingTableOptionsModel.ID = &subnet
		replaceSubnetRoutingTableOptionsModel.RoutingTableIdentity = routingTableIdentityModel
		resultRT, _, err := sess.ReplaceSubnetRoutingTableWithContext(context, replaceSubnetRoutingTableOptionsModel)

		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceSubnetRoutingTableWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[INFO] Updated subnet %s with Routing Table : %s", subnet, *resultRT.ID)

		d.SetId(subnet)
		return resourceIBMISSubnetRoutingTableAttachmentRead(context, d, meta)
	}

	if d.HasChange(isRoutingTableCrn) {
		subnet := d.Get(isSubnetID).(string)
		routingTableCrn := d.Get(isRoutingTableCrn).(string)

		// Construct an instance of the RoutingTableIdentityByID model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByCRN)
		routingTableIdentityModel.CRN = &routingTableCrn

		// Construct an instance of the ReplaceSubnetRoutingTableOptions model
		replaceSubnetRoutingTableOptionsModel := new(vpcv1.ReplaceSubnetRoutingTableOptions)
		replaceSubnetRoutingTableOptionsModel.ID = &subnet
		replaceSubnetRoutingTableOptionsModel.RoutingTableIdentity = routingTableIdentityModel
		resultRT, _, err := sess.ReplaceSubnetRoutingTableWithContext(context, replaceSubnetRoutingTableOptionsModel)

		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceSubnetRoutingTableWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[INFO] Updated subnet %s with Routing Table Crn : %s", subnet, *resultRT.CRN)

		d.SetId(subnet)
		return resourceIBMISSubnetRoutingTableAttachmentRead(context, d, meta)
	}

	return resourceIBMISSubnetRoutingTableAttachmentRead(context, d, meta)
}

func resourceIBMISSubnetRoutingTableAttachmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_subnet_routing_table_attachment", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Set the subnet with VPC default routing table
	getSubnetOptions := &vpcv1.GetSubnetOptions{
		ID: &id,
	}
	subnet, response, err := sess.GetSubnetWithContext(context, getSubnetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubnetWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	// Fetch VPC
	vpcID := *subnet.VPC.ID

	getvpcOptions := &vpcv1.GetVPCOptions{
		ID: &vpcID,
	}
	vpc, response, err := sess.GetVPCWithContext(context, getvpcOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPCWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Fetch default routing table
	if vpc.DefaultRoutingTable != nil {
		log.Printf("[DEBUG] vpc default routing table is not null :%s", *vpc.DefaultRoutingTable.ID)
		// Construct an instance of the RoutingTableIdentityByID model
		routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
		routingTableIdentityModel.ID = vpc.DefaultRoutingTable.ID

		// Construct an instance of the ReplaceSubnetRoutingTableOptions model
		replaceSubnetRoutingTableOptionsModel := new(vpcv1.ReplaceSubnetRoutingTableOptions)
		replaceSubnetRoutingTableOptionsModel.ID = &id
		replaceSubnetRoutingTableOptionsModel.RoutingTableIdentity = routingTableIdentityModel
		resultRT, _, err := sess.ReplaceSubnetRoutingTableWithContext(context, replaceSubnetRoutingTableOptionsModel)

		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceSubnetRoutingTableWithContext failed: %s", err.Error()), "ibm_is_subnet_routing_table_attachment", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		log.Printf("[INFO] Updated subnet %s with VPC default Routing Table : %s", id, *resultRT.ID)
	} else {
		log.Printf("[DEBUG] vpc default routing table is  null")
	}

	d.SetId("")
	return nil
}
