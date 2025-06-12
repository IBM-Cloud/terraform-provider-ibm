// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVirtualNetworkInterfaceFloatingIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVirtualNetworkInterfaceFloatingIPRead,

		Schema: map[string]*schema.Schema{
			"virtual_network_interface": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The virtual network interface identifier",
			},
			"floating_ip": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The floating IP identifier",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this floating IP. The name is unique across all floating IPs in the region.",
			},

			"deleted": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Link to documentation about deleted resources",
						},
					},
				},
			},
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The globally unique IP address.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this floating IP.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this floating IP.",
			},
		},
	}
}

func dataSourceIBMIsVirtualNetworkInterfaceFloatingIPRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_floating_ip", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	vniId := d.Get("virtual_network_interface").(string)
	fipId := d.Get("floating_ip").(string)

	getNetworkInterfaceFloatingIPOptions := &vpcv1.GetNetworkInterfaceFloatingIPOptions{}
	getNetworkInterfaceFloatingIPOptions.SetVirtualNetworkInterfaceID(vniId)
	getNetworkInterfaceFloatingIPOptions.SetID(fipId)

	floatingIP, _, err := sess.GetNetworkInterfaceFloatingIPWithContext(context, getNetworkInterfaceFloatingIPOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetNetworkInterfaceFloatingIPWithContext failed %s", err), "(Data) ibm_is_virtual_network_interface_floating_ip", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*floatingIP.ID)
	diagErr := dataIBMIsVirtualNetworkInterfaceFloatingIPGet(d, floatingIP)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func dataIBMIsVirtualNetworkInterfaceFloatingIPGet(d *schema.ResourceData, floatingIP *vpcv1.FloatingIPReference) diag.Diagnostics {
	if !core.IsNil(floatingIP.Name) {
		if err := d.Set("name", floatingIP.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_floating_ip", "read", "set-name").GetDiag()
		}
	}
	if err := d.Set("address", floatingIP.Address); err != nil {
		err = fmt.Errorf("Error setting address: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_floating_ip", "read", "set-address").GetDiag()
	}

	if err := d.Set("crn", floatingIP.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_floating_ip", "read", "set-crn").GetDiag()
	}
	if err := d.Set("href", floatingIP.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_floating_ip", "read", "set-href").GetDiag()
	}
	deleted := make(map[string]interface{})

	if floatingIP.Deleted != nil && floatingIP.Deleted.MoreInfo != nil {
		deleted["more_info"] = floatingIP.Deleted
	}

	if err := d.Set("deleted", []map[string]interface{}{deleted}); err != nil {
		err = fmt.Errorf("Error setting deleted: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_network_interface_floating_ip", "read", "set-deleted").GetDiag()
	}

	return nil
}
