// Copyright IBM Corp. 2024 All Rights Reserved.
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

func DataSourceIBMIsVPNGatewayServiceConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNGatewayServiceConnectionRead,

		Schema: map[string]*schema.Schema{
			"vpn_gateway": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"vpn_gateway_name", "vpn_gateway"},
				Description:  "The VPN gateway identifier.",
			},
			"vpn_gateway_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"vpn_gateway_name", "vpn_gateway"},
				Description:  "The VPN gateway name.",
			},
			"vpn_gateway_service_connection": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPN gateway connection identifier.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that this VPN service connection was created.",
			},
			"creator": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for transit gateway resource.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for transit gateway resource.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this VPN gateway service connection",
			},
			"lifecycle_reasons": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current `lifecycle_state` (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
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
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the VPN service connection.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this service connection:- `up`: operating normally- `degraded`: operating with compromised performance- `down`: not operational.",
			},
			"status_reasons": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current VPN service connection status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason. The enumerated values for this property may https://cloud.ibm.com/apidocs/vpc#property-value-expansion in the future.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the reason for this VPN service connection's status.",
						},
						"more_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Link to documentation about this status reason.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsVPNGatewayServiceConnectionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	vpn_gateway_id := d.Get("vpn_gateway").(string)
	vpn_gateway_name := d.Get("vpn_gateway_name").(string)
	vpn_gateway_service_connection := d.Get("vpn_gateway_service_connection").(string)

	var vpnGatewayServiceConn vpcv1.VPNGatewayServiceConnection

	if vpn_gateway_name != "" {
		listvpnGWOptions := vpcClient.NewListVPNGatewaysOptions()

		start := ""
		allrecs := []vpcv1.VPNGatewayIntf{}
		for {
			if start != "" {
				listvpnGWOptions.Start = &start
			}
			availableVPNGateways, detail, err := vpcClient.ListVPNGatewaysWithContext(context, listvpnGWOptions)
			if err != nil || availableVPNGateways == nil {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error reading list of VPN Gateways:%s\n%s", err, detail), "(Data) ibm_is_vpn_gateway_service_connection", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return diag.FromErr(tfErr)
			}
			start = flex.GetNext(availableVPNGateways.Next)
			allrecs = append(allrecs, availableVPNGateways.VPNGateways...)
			if start == "" {
				break
			}
		}
		vpn_gateway_found := false
		for _, vpnGatewayIntfItem := range allrecs {
			if *vpnGatewayIntfItem.(*vpcv1.VPNGateway).Name == vpn_gateway_name {
				vpnGateway := vpnGatewayIntfItem.(*vpcv1.VPNGateway)
				vpn_gateway_id = *vpnGateway.ID
				vpn_gateway_found = true
				break
			}
		}
		if !vpn_gateway_found {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No vpn gateway found with given name %s", vpn_gateway_name), "(Data) ibm_is_vpn_gateway_service_connection", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return diag.FromErr(tfErr)
		}
	}

	getVPNGatewayServiceConnectionOptions := &vpcv1.GetVPNGatewayServiceConnectionOptions{}

	getVPNGatewayServiceConnectionOptions.SetVPNGatewayID(vpn_gateway_id)
	getVPNGatewayServiceConnectionOptions.SetID(vpn_gateway_service_connection)

	vpnGatewayServiceConnection, response, err := vpcClient.GetVPNGatewayServiceConnectionWithContext(context, getVPNGatewayServiceConnectionOptions)
	if err != nil || vpnGatewayServiceConnection == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayServiceConnectionWithContext failed %s\n%s", err, response), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	vpnGatewayServiceConn = *vpnGatewayServiceConnection

	d.SetId(*vpnGatewayServiceConn.ID)

	if err = d.Set("created_at", flex.DateTimeToString(vpnGatewayServiceConn.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	if err := d.Set("creator", resourceVPNGatewayServiceConnectionFlattenCreator(vpnGatewayServiceConn.Creator)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting creator: %s", err), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	if err := d.Set("lifecycle_reasons", resourceVPNGatewayServiceConnectionFlattenLifecycleReasons(vpnGatewayServiceConn.LifecycleReasons)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	if err = d.Set("lifecycle_state", vpnGatewayServiceConn.LifecycleState); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	if err = d.Set("status", vpnGatewayServiceConn.Status); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf(" Error setting status: %s", err), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	if err := d.Set("status_reasons", resourceVPNGatewayServiceConnectionFlattenStateReasons(vpnGatewayServiceConn.StatusReasons)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_vpn_gateway_service_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	return nil
}
