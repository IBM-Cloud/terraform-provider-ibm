// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVPNGatewayConnectionPeerCidrs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNGatewayConnectionPeerCidrsRead,

		Schema: map[string]*schema.Schema{
			"vpn_gateway": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPN gateway identifier.",
			},
			"vpn_gateway_connection": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The VPN gateway connection identifier.",
			},
			"cidrs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The CIDRs for this resource.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIBMIsVPNGatewayConnectionPeerCidrsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpn_gateway_connection_peer_cidrs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	listVPNGatewayConnectionsPeerCidrsOptions := &vpcv1.ListVPNGatewayConnectionsPeerCIDRsOptions{}

	listVPNGatewayConnectionsPeerCidrsOptions.SetVPNGatewayID(d.Get("vpn_gateway").(string))
	listVPNGatewayConnectionsPeerCidrsOptions.SetID(d.Get("vpn_gateway_connection").(string))

	vpnGatewayConnectionCidRs, _, err := vpcClient.ListVPNGatewayConnectionsPeerCIDRsWithContext(context, listVPNGatewayConnectionsPeerCidrsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNGatewayConnectionsPeerCIDRsWithContext failed %s", err), "(Data) ibm_is_vpn_gateway_connection_peer_cidrs", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(dataSourceIBMIsVPNGatewayConnectionPeerCidrsID(d))
	d.Set("cidrs", vpnGatewayConnectionCidRs.CIDRs)
	if err = d.Set("cidrs", vpnGatewayConnectionCidRs.CIDRs); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cidrs %s", err), "(Data) ibm_is_vpn_gateway_connection_peer_cidrs", "read", "cidrs-set").GetDiag()
	}
	return nil
}

// dataSourceIBMIsVPNGatewayConnectionPeerCidrsID returns a reasonable ID for the list.
func dataSourceIBMIsVPNGatewayConnectionPeerCidrsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
