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

func dataSourceIBMIsDynamicRouteServerStaticRouteDeletedSchema() *schema.Schema {
	return &schema.Schema{
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
	}
}

func DataSourceIBMIsDynamicRouteServerStaticRoute() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerStaticRouteRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"dynamic_route_server_static_route_id": &schema.Schema{
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
							Description: "The route reference for the applied route. If absent, the route could not be added to the routing table.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deleted": dataSourceIBMIsDynamicRouteServerStaticRouteDeletedSchema(),
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
			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The destination CIDR of the route.",
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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this static route.",
			},
			"next_hops": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The next hop peer groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"deleted": dataSourceIBMIsDynamicRouteServerStaticRouteDeletedSchema(),
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
				Description: "The number of seconds to wait before deleting a route from `routing_tables` when the status of the route's associated peer is `down`.",
			},
			"routing_tables": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routing tables to update.",
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
									"deleted": dataSourceIBMIsDynamicRouteServerStaticRouteDeletedSchema(),
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

	dynamicRouteServerID := d.Get("dynamic_route_server_id").(string)
	staticRouteID := d.Get("dynamic_route_server_static_route_id").(string)

	getOptions := &vpcv1.GetDynamicRouteServerStaticRouteOptions{}
	getOptions.SetDynamicRouteServerID(dynamicRouteServerID)
	getOptions.SetID(staticRouteID)

	staticRoute, _, err := vpcClient.GetDynamicRouteServerStaticRouteWithContext(context, getOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerStaticRouteWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_static_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	modelMap, err := DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(staticRoute)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_static_route", "read", "static-route-to-map").GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", dynamicRouteServerID, staticRouteID))

	for key, value := range modelMap {
		if key == "id" {
			continue
		}
		if err = d.Set(key, value); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "(Data) ibm_is_dynamic_route_server_static_route", "read", "set-"+key).GetDiag()
		}
	}

	return nil
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(model *vpcv1.DynamicRouteServerStaticRoute) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model *vpcv1.DynamicRouteServerStaticRouteAddedRoute) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteAddedRouteToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model *vpcv1.RouteReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteRouteReferenceToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteDeletedToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model *vpcv1.LifecycleReason) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteLifecycleReasonToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model vpcv1.DynamicRouteServerStaticRouteNextHopIntf) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model *vpcv1.DynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteNextHopDynamicRouteServerPeerGroupReferenceToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model *vpcv1.DynamicRouteServerStaticRouteRoutingTable) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteDynamicRouteServerStaticRouteRoutingTableToMap(model)
}

func DataSourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model *vpcv1.RoutingTableReference) (map[string]interface{}, error) {
	return ResourceIBMIsDynamicRouteServerStaticRouteRoutingTableReferenceToMap(model)
}
