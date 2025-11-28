// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIInstanceNetworks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceNetworksRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_InstanceID: {
				Description:  "The unique identifier or ID of the instance.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_Networks: {
				Computed:    true,
				Description: "List of networks associated with this instance.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_ExternalIP: {
							Computed:    true,
							Description: "The external IP address of the instance.",
							Type:        schema.TypeString,
						},
						Attr_Href: {
							Computed:    true,
							Description: "Link to this PVM instance network.",
							Type:        schema.TypeString,
						},
						Attr_IPAddress: {
							Computed:    true,
							Description: "The IP address of the instance.",
							Type:        schema.TypeString,
						},
						Attr_MacAddress: {
							Computed:    true,
							Description: "The MAC address of the instance.",
							Type:        schema.TypeString,
						},
						Attr_NetworkID: {
							Computed:    true,
							Description: "The network ID of the instance.",
							Type:        schema.TypeString,
						},
						Attr_NetworkInterfaceID: {
							Computed:    true,
							Description: "ID of the network interface.",
							Type:        schema.TypeString,
						},
						Attr_NetworkName: {
							Computed:    true,
							Description: "The network name of the instance.",
							Type:        schema.TypeString,
						},
						Attr_NetworkSecurityGroupIDs: {
							Computed:    true,
							Description: "IDs of the network security groups that the network interface is a member of.",
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						Attr_NetworkSecurityGroupsHref: {
							Computed:    true,
							Description: "Links to the network security groups that the network interface is a member of.",
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						Attr_Type: {
							Computed:    true,
							Description: "The type of the network.",
							Type:        schema.TypeString,
						},
						Attr_Version: {
							Computed:    true,
							Description: "Version of the network information.",
							Type:        schema.TypeFloat,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIInstanceNetworksRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_instance_networks", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	instanceID := d.Get(Arg_InstanceID).(string)

	netC := instance.NewIBMPIInstanceNetworksClient(ctx, sess, cloudInstanceID)
	netsWrap, err := netC.GetAll(instanceID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAll networks failed: %s", err.Error()), "(Data) ibm_pi_instance_networks", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	networks := flattenPvmInstanceNetworksv2(netsWrap.Networks)

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, instanceID))
	d.Set(Attr_Networks, networks)

	return nil
}

func flattenPvmInstanceNetworksv2(InstanceNetworks []*models.PVMInstanceNetwork) (networks []map[string]any) {
	if InstanceNetworks != nil {
		networks = make([]map[string]any, len(InstanceNetworks))
		for i, pvmNet := range InstanceNetworks {
			p := make(map[string]any)
			p[Attr_ExternalIP] = pvmNet.ExternalIP
			p[Attr_IPAddress] = pvmNet.IPAddress
			p[Attr_MacAddress] = pvmNet.MacAddress
			p[Attr_NetworkID] = pvmNet.NetworkID
			p[Attr_NetworkInterfaceID] = pvmNet.NetworkInterfaceID
			p[Attr_NetworkName] = pvmNet.NetworkName
			p[Attr_Type] = pvmNet.Type
			p[Attr_Href] = pvmNet.Href
			p[Attr_Version] = pvmNet.Version
			if len(pvmNet.NetworkSecurityGroupIDs) > 0 {
				p[Attr_NetworkSecurityGroupIDs] = pvmNet.NetworkSecurityGroupIDs
			}
			if len(pvmNet.NetworkSecurityGroupsHref) > 0 {
				p[Attr_NetworkSecurityGroupsHref] = pvmNet.NetworkSecurityGroupsHref
			}
			networks[i] = p
		}
		return networks
	}
	return
}
