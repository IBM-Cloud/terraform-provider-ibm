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

func ResourceIBMISVpnGatewayMemberReplace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVpnGatewayMemberReplaceCreate,
		ReadContext:   resourceIBMISVpnGatewayMemberReplaceRead,
		DeleteContext: resourceIBMISVpnGatewayMemberReplaceDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				d.Set("vpn_gateway_id", d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPN gateway identifier",
			},
			"vpn_gateway_member_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPN gateway member identifier",
			},
			"subnet": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The subnet to replace the VPN gateway member with",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The subnet identifier",
						},
						"crn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The subnet CRN",
						},
						"href": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The subnet href",
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the VPN gateway member replacement",
			},
		},
	}
}

func resourceIBMISVpnGatewayMemberReplaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] VPN Gateway Member Replace create")

	vpnGatewayID := d.Get("vpn_gateway_id").(string)
	vpnGatewayMemberID := d.Get("vpn_gateway_member_id").(string)

	err := vpnGatewayMemberReplaceCreate(ctx, d, meta, vpnGatewayID, vpnGatewayMemberID)
	if err != nil {
		return err
	}

	return resourceIBMISVpnGatewayMemberReplaceRead(ctx, d, meta)
}

func vpnGatewayMemberReplaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}, vpnGatewayID, vpnGatewayMemberID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Build subnet prototype
	subnetPrototype := &vpcv1.SubnetIdentity{}

	if subnetRaw, ok := d.GetOk("subnet"); ok {
		subnetList := subnetRaw.([]interface{})
		if len(subnetList) > 0 {
			subnet := subnetList[0].(map[string]interface{})
			if id, ok := subnet["id"]; ok && id.(string) != "" {
				subnetPrototype.ID = flex.PtrToString(id.(string))
			} else if crn, ok := subnet["crn"]; ok && crn.(string) != "" {
				subnetPrototype.CRN = flex.PtrToString(crn.(string))
			} else if href, ok := subnet["href"]; ok && href.(string) != "" {
				subnetPrototype.Href = flex.PtrToString(href.(string))
			} else {
				err := fmt.Errorf("[ERROR] subnet must provide one of id, crn, or href")
				tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "set-subnet")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
		} else {
			err := fmt.Errorf("[ERROR] subnet must be provided")
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "set-subnet")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	} else {
		err := fmt.Errorf("[ERROR] subnet is required")
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "set-subnet")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Build subnet prototype

	replaceOptions := &vpcv1.ReplaceVPNGatewayMemberPrivateIPOptions{
		VPNGatewayID: &vpnGatewayID,
		ID:           &vpnGatewayMemberID,
		PrivateIP: &vpcv1.VPNGatewayMemberPrivateIPPrototype{
			Subnet: subnetPrototype,
		},
	}

	_, response, err := sess.ReplaceVPNGatewayMemberPrivateIPWithContext(ctx, replaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceVPNGatewayMemberWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Set the resource ID as combination of gateway ID and member ID
	d.SetId(fmt.Sprintf("%s/%s", vpnGatewayID, vpnGatewayMemberID))
	log.Printf("[INFO] VPN Gateway Member Replace initiated for gateway: %s, member: %s", vpnGatewayID, vpnGatewayMemberID)

	// Check if response contains the updated member
	if response != nil && response.StatusCode == 204 {
		log.Printf("[INFO] VPN Gateway Member Replace completed successfully (204 No Content)")
	} else if response != nil && response.StatusCode != 204 {
		err := fmt.Errorf("[ERROR] Unexpected status code: %d", response.StatusCode)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceVPNGatewayMemberWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return nil
}

func resourceIBMISVpnGatewayMemberReplaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Split the composite ID
	id := d.Id()
	if id == "" {
		return nil
	}

	// Split the ID to get gateway and member IDs
	parts, err := flex.IdParts(id)
	if err != nil {
		return diag.FromErr(err)
	}
	vpnGatewayID := parts[0]
	vpnGatewayMemberID := parts[1]

	diag := vpnGatewayMemberReplaceGet(d, meta, vpnGatewayID, vpnGatewayMemberID)
	if diag != nil {
		return diag
	}
	return nil
}

func vpnGatewayMemberReplaceGet(d *schema.ResourceData, meta interface{}, vpnGatewayID, vpnGatewayMemberID string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getOptions := &vpcv1.GetVPNGatewayMemberOptions{
		VPNGatewayID: &vpnGatewayID,
		ID:           &vpnGatewayMemberID,
	}

	_, response, err := sess.GetVPNGatewayMember(getOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		err := fmt.Errorf("[ERROR] Error Getting VPN Gateway Member (%s/%s): %s\n%s", vpnGatewayID, vpnGatewayMemberID, err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpnGatewayMemberReplaceGet failed: %s", err.Error()), "ibm_is_vpn_gateway_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Set the fields from the response
	d.Set("vpn_gateway_id", vpnGatewayID)
	d.Set("vpn_gateway_member_id", vpnGatewayMemberID)

	d.Set("status", "204")

	return nil
}

func resourceIBMISVpnGatewayMemberReplaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] VPN Gateway Member Replace delete - setting ID to empty")
	d.SetId("")
	return nil
}
