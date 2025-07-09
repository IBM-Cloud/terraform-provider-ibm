// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVPNServerRoute() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNServerRouteRead,

		Schema: map[string]*schema.Schema{
			"vpn_server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPN server identifier.",
			},

			"identifier": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The unique identifier for this VPN server route",
			},

			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The unique user-defined name for this VPN server route",
			},

			"action": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The action to perform with a packet matching the VPN route:- `translate`: translate the source IP address to one of the private IP addresses of the VPN server.- `deliver`: deliver the packet into the VPC.- `drop`: drop the packetThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the VPN route on which the unexpected property value was encountered.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the VPN route was created.",
			},
			"destination": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The destination for this VPN route in the VPN server. If an incoming packet does not match any destination, it will be dropped.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this VPN route.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health of this resource.- `ok`: Healthy- `degraded`: Suffering from compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.",
			},
			"health_reasons": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this health state.",
						},

						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this health state.",
						},

						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this health state.",
						},
					},
				},
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the VPN route.",
			},
			"lifecycle_reasons": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current lifecycle_state (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the reason for this lifecycle state.",
						},

						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this lifecycle state.",
						},

						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about the reason for this lifecycle state.",
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
		},
	}
}

func dataSourceIBMIsVPNServerRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpn_server_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var vpnServerRoute *vpcv1.VPNServerRoute

	if v, ok := d.GetOk("identifier"); ok {

		getVPNServerRouteOptions := &vpcv1.GetVPNServerRouteOptions{}

		getVPNServerRouteOptions.SetVPNServerID(d.Get("vpn_server").(string))
		getVPNServerRouteOptions.SetID(v.(string))

		vpnServerRouteInfo, _, err := vpcClient.GetVPNServerRouteWithContext(context, getVPNServerRouteOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNServerRouteWithContext failed: %s", err.Error()), "(Data) ibm_is_vpn_server_route", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		vpnServerRoute = vpnServerRouteInfo
	} else if v, ok := d.GetOk("name"); ok {
		name := v.(string)

		start := ""
		allrecs := []vpcv1.VPNServerRoute{}

		for {
			listVPNServerRoutesOptions := &vpcv1.ListVPNServerRoutesOptions{}
			listVPNServerRoutesOptions.SetVPNServerID(d.Get("vpn_server").(string))

			if start != "" {
				listVPNServerRoutesOptions.Start = &start
			}
			vpnServerRouteCollection, _, err := vpcClient.ListVPNServerRoutesWithContext(context, listVPNServerRoutesOptions)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNServerRoutesWithContext failed: %s", err.Error()), "(Data) ibm_is_vpn_server_route", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			start = flex.GetNext(vpnServerRouteCollection.Next)
			allrecs = append(allrecs, vpnServerRouteCollection.Routes...)
			if start == "" {
				break
			}
		}

		for _, vpnServerRouteInfo := range allrecs {
			if *vpnServerRouteInfo.Name == name {
				vpnServerRoute = &vpnServerRouteInfo
				break
			}
		}
		if vpnServerRoute == nil {
			log.Printf("[DEBUG] No vpnServer route found with name %s", name)
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("No vpn server route found with name: %s", err), "(Data) ibm_is_vpn_server_route", "read", "not-found").GetDiag()

		}
	}

	d.SetId(fmt.Sprintf("%s/%s", d.Get("vpn_server").(string), *vpnServerRoute.ID))

	if err = d.Set("vpn_server", d.Get("vpn_server").(string)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vpn-server: %s", err), "(Data) ibm_is_vpn_server_route", "read", "vpn-server").GetDiag()
	}

	if err = d.Set("identifier", *vpnServerRoute.ID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting identifier: %s", err), "(Data) ibm_is_vpn_server_route", "read", "identifier").GetDiag()
	}
	if err = d.Set("action", vpnServerRoute.Action); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-action").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(vpnServerRoute.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("destination", vpnServerRoute.Destination); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting destination: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-destination").GetDiag()
	}
	if err = d.Set("href", vpnServerRoute.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-href").GetDiag()
	}
	if err = d.Set("health_state", vpnServerRoute.HealthState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_state: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-health_state").GetDiag()
	}
	if vpnServerRoute.HealthReasons != nil {
		if err := d.Set("health_reasons", resourceVPNServerRouteFlattenHealthReasons(vpnServerRoute.HealthReasons)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting health_reasons: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-health_reasons").GetDiag()
		}
	}
	if err = d.Set("lifecycle_state", vpnServerRoute.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-lifecycle_state").GetDiag()
	}
	if vpnServerRoute.LifecycleReasons != nil {
		if err := d.Set("lifecycle_reasons", resourceVPNServerRouteFlattenLifecycleReasons(vpnServerRoute.LifecycleReasons)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-lifecycle_reasons").GetDiag()
		}
	}
	if err = d.Set("name", vpnServerRoute.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-name").GetDiag()
	}
	if err = d.Set("resource_type", vpnServerRoute.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_vpn_server_route", "read", "set-resource_type").GetDiag()
	}

	return nil
}
