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

func DataSourceIBMIsVPNGatewayConnectionLocalCidrs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNGatewayConnectionLocalCidrsRead,

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

func dataSourceIBMIsVPNGatewayConnectionLocalCidrsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_vpn_gateway_connection_local_cidrs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listVPNGatewayConnectionsLocalCidrsOptions := &vpcv1.ListVPNGatewayConnectionsLocalCIDRsOptions{}

	listVPNGatewayConnectionsLocalCidrsOptions.SetVPNGatewayID(d.Get("vpn_gateway").(string))
	listVPNGatewayConnectionsLocalCidrsOptions.SetID(d.Get("vpn_gateway_connection").(string))

	vpnGatewayConnectionCidRs, _, err := vpcClient.ListVPNGatewayConnectionsLocalCIDRsWithContext(context, listVPNGatewayConnectionsLocalCidrsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNGatewayConnectionsLocalCIDRsWithContext failed %s", err), "(Data) ibm_is_vpn_gateway_connection_local_cidrs", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(dataSourceIBMIsVPNGatewayConnectionLocalCidrsID(d))

	if err = d.Set("cidrs", vpnGatewayConnectionCidRs.CIDRs); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cidrs %s", err), "(Data) ibm_is_vpn_gateway_connection_local_cidrs", "read", "cidrs-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsVPNGatewayConnectionLocalCidrsID returns a reasonable ID for the list.
func dataSourceIBMIsVPNGatewayConnectionLocalCidrsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
