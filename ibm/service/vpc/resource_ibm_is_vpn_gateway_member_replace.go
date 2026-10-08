// Copyright IBM Corp. 2026 All Rights Reserved.
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
	isVPNGatewayMemberLifecycleStable = "stable"
	isVPNGatewayMemberLifecycleFailed = "failed"
)

func ResourceIBMISVpnGatewayMemberReplace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISVpnGatewayMemberReplaceCreate,
		ReadContext:   resourceIBMISVpnGatewayMemberReplaceRead,
		DeleteContext: resourceIBMISVpnGatewayMemberReplaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceIBMISVpnGatewayMemberReplaceImport,
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
				Description: "The VPN gateway identifier.",
			},
			"vpn_gateway_member_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The VPN gateway member identifier.",
			},
			"subnet": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				MaxItems:    1,
				Description: "The subnet to move the VPN gateway member to. A reserved IP is allocated from this subnet.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The subnet identifier.",
						},
						"crn": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The subnet CRN.",
						},
						"href": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							ForceNew:    true,
							Description: "The subnet URL.",
						},
					},
				},
			},
			"role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The high availability role assigned to the VPN gateway member.",
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the VPN gateway member.",
			},
			"health_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The health state of the VPN gateway member.",
			},
			"private_ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The private IP address of the VPN gateway member in the new subnet.",
			},
			"public_ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The public IP address of the VPN gateway member.",
			},
		},
	}
}

func resourceIBMISVpnGatewayMemberReplaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	vpnGatewayID := d.Get("vpn_gateway_id").(string)
	vpnGatewayMemberID := d.Get("vpn_gateway_member_id").(string)

	subnetList := d.Get("subnet").([]interface{})
	if len(subnetList) == 0 || subnetList[0] == nil {
		err := fmt.Errorf("subnet must specify one of id, crn or href")
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "parse-subnet").GetDiag()
	}
	subnetIdentity, err := vpnGatewaySubnetIdentityFromMap(subnetList[0].(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "create", "parse-subnet").GetDiag()
	}

	replaceOptions := sess.NewReplaceVPNGatewayMemberOptions(vpnGatewayID, vpnGatewayMemberID,
		&vpcv1.VPNGatewayMemberPrivateIPPrototypeReservedIPPrototypeVPNGatewayContext{
			Subnet: subnetIdentity,
		})

	member, _, err := sess.ReplaceVPNGatewayMemberWithContext(ctx, replaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceVPNGatewayMemberWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_member_replace", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// The replace recreates the member. Track it by the ID the API returns so
	// Read does not lose it if the service hands back a new member ID.
	memberID := vpnGatewayMemberID
	if member != nil && member.ID != nil {
		memberID = *member.ID
	}
	d.SetId(fmt.Sprintf("%s/%s", vpnGatewayID, memberID))
	log.Printf("[INFO] VPN gateway member replace started for gateway %s, member %s", vpnGatewayID, memberID)

	if _, err = isWaitForVPNGatewayMemberStable(ctx, sess, vpnGatewayID, memberID, d.Timeout(schema.TimeoutCreate)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForVPNGatewayMemberStable failed: %s", err.Error()), "ibm_is_vpn_gateway_member_replace", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMISVpnGatewayMemberReplaceRead(ctx, d, meta)
}

func isWaitForVPNGatewayMemberStable(ctx context.Context, sess *vpcv1.VpcV1, vpnGatewayID, memberID string, timeout time.Duration) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"pending", "updating", "waiting"},
		Target:  []string{isVPNGatewayMemberLifecycleStable},
		Refresh: func() (interface{}, string, error) {
			member, _, err := sess.GetVPNGatewayMemberWithContext(ctx, &vpcv1.GetVPNGatewayMemberOptions{
				VPNGatewayID: &vpnGatewayID,
				ID:           &memberID,
			})
			if err != nil {
				return nil, "", fmt.Errorf("error getting VPN gateway member (%s/%s): %s", vpnGatewayID, memberID, err)
			}
			if member.LifecycleState == nil {
				return member, "pending", nil
			}
			if *member.LifecycleState == isVPNGatewayMemberLifecycleFailed {
				return member, *member.LifecycleState, fmt.Errorf("VPN gateway member (%s/%s) went into failed state", vpnGatewayID, memberID)
			}
			return member, *member.LifecycleState, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForStateContext(ctx)
}

func resourceIBMISVpnGatewayMemberReplaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	parts, err := flex.IdParts(d.Id())
	if err != nil || len(parts) != 2 {
		if err == nil {
			err = fmt.Errorf("unexpected ID format (%s), expected vpn_gateway_id/vpn_gateway_member_id", d.Id())
		}
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "read", "sep-id-parts").GetDiag()
	}
	vpnGatewayID, memberID := parts[0], parts[1]

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	member, response, err := sess.GetVPNGatewayMemberWithContext(ctx, &vpcv1.GetVPNGatewayMemberOptions{
		VPNGatewayID: &vpnGatewayID,
		ID:           &memberID,
	})
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVPNGatewayMemberWithContext failed: %s", err.Error()), "ibm_is_vpn_gateway_member_replace", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("vpn_gateway_id", vpnGatewayID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "read", "set-vpn_gateway_id").GetDiag()
	}
	// Keep the configured member ID; only fill it in on import.
	if d.Get("vpn_gateway_member_id").(string) == "" {
		if err = d.Set("vpn_gateway_member_id", memberID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "read", "set-vpn_gateway_member_id").GetDiag()
		}
	}
	if member.Role != nil {
		d.Set("role", *member.Role)
	}
	if member.LifecycleState != nil {
		d.Set("lifecycle_state", *member.LifecycleState)
	}
	if member.HealthState != nil {
		d.Set("health_state", *member.HealthState)
	}
	if member.PublicIP != nil && member.PublicIP.Address != nil {
		d.Set("public_ip_address", *member.PublicIP.Address)
	}
	if member.PrivateIP != nil {
		if member.PrivateIP.Address != nil {
			d.Set("private_ip_address", *member.PrivateIP.Address)
		}
		if member.PrivateIP.Subnet != nil {
			subnet := map[string]interface{}{}
			if member.PrivateIP.Subnet.ID != nil {
				subnet["id"] = *member.PrivateIP.Subnet.ID
			}
			if member.PrivateIP.Subnet.CRN != nil {
				subnet["crn"] = *member.PrivateIP.Subnet.CRN
			}
			if member.PrivateIP.Subnet.Href != nil {
				subnet["href"] = *member.PrivateIP.Subnet.Href
			}
			if err = d.Set("subnet", []map[string]interface{}{subnet}); err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_vpn_gateway_member_replace", "read", "set-subnet").GetDiag()
			}
		}
	}
	return nil
}

func resourceIBMISVpnGatewayMemberReplaceImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil, err
	}
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("unexpected import ID (%s), expected vpn_gateway_id/vpn_gateway_member_id", d.Id())
	}
	d.Set("vpn_gateway_id", parts[0])
	d.Set("vpn_gateway_member_id", parts[1])
	return []*schema.ResourceData{d}, nil
}

// resourceIBMISVpnGatewayMemberReplaceDelete only removes the resource from
// state. There is no API to undo a member replace, and the member itself is
// owned by the VPN gateway.
func resourceIBMISVpnGatewayMemberReplaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
