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

func DataSourceIBMIsVPNGatewayAdvertisedCidrs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVPNGatewayAdvertisedCidrsRead,

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
			"advertised_cidrs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The additional CIDRs advertised through any enabled routing protocol (for example, BGP). The routing protocol will advertise routes with these CIDRs and VPC prefixes as route destinations.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIBMIsVPNGatewayAdvertisedCidrsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "(Data) ibm_is_vpn_gateway_advertised_cidrs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	vpn_gateway_id := d.Get("vpn_gateway").(string)
	vpn_gateway_name := d.Get("vpn_gateway_name").(string)

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
				if err != nil {
					tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error reading list of VPN Gateways:%s\n%s", err, detail), "(Data) ibm_is_vpn_gateway_advertised_cidrs", "read")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					return diag.FromErr(tfErr)
				}
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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("No vpn gateway found with given name %s", vpn_gateway_name), "(Data) ibm_is_vpn_gateway_advertised_cidrs", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return diag.FromErr(tfErr)
		}
	}

	listVPNGatewayAdvertisedCIDRsOptions := &vpcv1.ListVPNGatewayAdvertisedCIDRsOptions{}

	listVPNGatewayAdvertisedCIDRsOptions.SetVPNGatewayID(vpn_gateway_id)

	vpnGatewayAdvertisedCidRs, response, err := vpcClient.ListVPNGatewayAdvertisedCIDRsWithContext(context, listVPNGatewayAdvertisedCIDRsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVPNGatewayAdvertisedCIDRsWithContext failed %s\n%s", err, response), "(Data) ibm_is_vpn_gateway_advertised_cidrs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	d.SetId(time.Now().UTC().String())
	d.Set("advertised_cidrs", vpnGatewayAdvertisedCidRs.AdvertisedCIDRs)

	return nil
}
