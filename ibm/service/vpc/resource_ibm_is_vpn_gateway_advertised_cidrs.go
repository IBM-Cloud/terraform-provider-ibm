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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isVPNGatewayAdvertisedCidr           = "cidr"
	isVPNGatewayAdvertisedCidrVPNGateway = "vpn_gateway"
	isVPNGatewayAdvertisedCidrDeleting   = "deleting"
	isVPNGatewayAdvertisedCidrDeleted    = "done"
)

func ResourceIBMISVPNGatewayAdvertisedCidr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVPNGatewayAdvertisedCidrCreate,
		ReadContext:   resourceIBMISVPNGatewayAdvertisedCidrRead,
		DeleteContext: resourceIBMISVPNGatewayAdvertisedCidrDelete,

		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			isVPNGatewayAdvertisedCidrVPNGateway: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPN gateway identifier",
			},
			isVPNGatewayAdvertisedCidr: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The IP address range in CIDR block notation.",
			},
		},
	}
}

func resourceIBMISVPNGatewayAdvertisedCidrCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Adding advertised cidr to vpn gateway")

	gatewayID := d.Get(isVPNGatewayAdvertisedCidrVPNGateway).(string)
	cidr := d.Get(isVPNGatewayAdvertisedCidr).(string)

	options := &vpcv1.AddVPNGatewayAdvertisedCIDROptions{
		VPNGatewayID: &gatewayID,
		CIDR:         &cidr,
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "ibm_is_vpn_gateway_advertised_cidr", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	response, err := sess.AddVPNGatewayAdvertisedCIDR(options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error adding advertised cidr to VPN Gateway err %s\n%s", err, response), "ibm_is_vpn_gateway_advertised_cidr", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	d.SetId(fmt.Sprintf("%s/%s", gatewayID, cidr))

	return resourceIBMISVPNGatewayAdvertisedCidrRead(context, d, meta)
}

func resourceIBMISVPNGatewayAdvertisedCidrRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "ibm_is_vpn_gateway_advertised_cidr", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_advertised_cidr", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	gID := parts[0]
	gAdvertisedCidr := parts[1] + "/" + parts[2]

	checkVPNGatewayAdvertisedCIDROptions := &vpcv1.CheckVPNGatewayAdvertisedCIDROptions{
		VPNGatewayID: &gID,
		CIDR:         &gAdvertisedCidr,
	}
	response, err := sess.CheckVPNGatewayAdvertisedCIDR(checkVPNGatewayAdvertisedCIDROptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error getting advertised cidr : %s\n%s", err, response), "ibm_is_vpn_gateway_advertised_cidr", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	return nil
}

func resourceIBMISVPNGatewayAdvertisedCidrDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_advertised_cidr", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	gID := parts[0]
	cidr := parts[1] + "/" + parts[2]

	deleteErr := vpngwAdvertisedCidrDelete(d, meta, gID, cidr)
	if deleteErr != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error deleting advertised cidr err %s", err), "ibm_is_vpn_gateway_advertised_cidr", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	return nil
}

func vpngwAdvertisedCidrDelete(d *schema.ResourceData, meta interface{}, gID, gCidr string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "ibm_is_vpn_gateway_advertised_cidr", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}
	checkVPNGatewayAdvertisedCIDROptions := &vpcv1.CheckVPNGatewayAdvertisedCIDROptions{
		VPNGatewayID: &gID,
		CIDR:         &gCidr,
	}
	response, err := sess.CheckVPNGatewayAdvertisedCIDR(checkVPNGatewayAdvertisedCIDROptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error getting Vpn Gateway advertised cidr(%s): %s\n%s", gCidr, err, response), "ibm_is_vpn_gateway_advertised_cidr", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	removeVPNGatewayAdvertisedCIDROptions := &vpcv1.RemoveVPNGatewayAdvertisedCIDROptions{
		VPNGatewayID: &gID,
		CIDR:         &gCidr,
	}
	response, err = sess.RemoveVPNGatewayAdvertisedCIDR(removeVPNGatewayAdvertisedCIDROptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error removing advertised cidr from Vpn Gateway: %s\n%s", err, response), "ibm_is_vpn_gateway_advertised_cidr", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	_, err = isWaitForVPNGatewayAdvertisedCIDRDeleted(sess, gID, gCidr, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error checking for Vpn Gateway advertised cidr (%s) is deleted: %s", gCidr, err), "ibm_is_vpn_gateway_advertised_cidr", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return diag.FromErr(tfErr)
	}

	d.SetId("")
	return nil
}

func isWaitForVPNGatewayAdvertisedCIDRDeleted(vpnGatewayAdverisedCidr *vpcv1.VpcV1, gID, gCidr string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPNGatewayAdvertisedCIDR (%s) to be deleted.", gCidr)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPNGatewayAdvertisedCidrDeleting},
		Target:     []string{"", isVPNGatewayAdvertisedCidrDeleted},
		Refresh:    isVPNGatewayAdvertisedCIDRDeleteRefreshFunc(vpnGatewayAdverisedCidr, gID, gCidr),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPNGatewayAdvertisedCIDRDeleteRefreshFunc(vpnGatewayAdverisedCidr *vpcv1.VpcV1, gID, gCidr string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		checkVPNGatewayAdvertisedCIDROptions := &vpcv1.CheckVPNGatewayAdvertisedCIDROptions{
			VPNGatewayID: &gID,
			CIDR:         &gCidr,
		}
		response, err := vpnGatewayAdverisedCidr.CheckVPNGatewayAdvertisedCIDR(checkVPNGatewayAdvertisedCIDROptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return "", isVPNGatewayConnectionDeleted, nil
			}
			return "", "", fmt.Errorf("[ERROR] The VPNGateway Advertised CIDR %s failed to delete: %s\n%s", gCidr, err, response)
		}
		return nil, isVPNGatewayConnectionDeleting, nil
	}
}
