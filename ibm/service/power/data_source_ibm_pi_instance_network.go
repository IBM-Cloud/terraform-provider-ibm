// Copyright IBM Corp.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"

	instance "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIInstanceNetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstanceNetworkRead,
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
			Arg_NetworkID: {
				Description:  "The network ID on the instance.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_ExternalIP: {
				Computed:    true,
				Description: "The external IP address of the instance.",
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
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeSet,
			},
			Attr_NetworkSecurityGroupsHref: {
				Computed:    true,
				Description: "Links to the network security groups that the network interface is a member of.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeList,
			},
			Attr_Type: {
				Computed:    true,
				Description: "The type of the network.",
				Type:        schema.TypeString,
			},
			Attr_Href: {
				Computed:    true,
				Description: "Link to this PVM instance network.",
				Type:        schema.TypeString,
			},
			Attr_Version: {
				Computed:    true,
				Description: "Version of the network information.",
				Type:        schema.TypeFloat,
			},
		},
	}
}

func dataSourceIBMPIInstanceNetworkRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("IBMPISession failed: %s", err.Error()), "(Data) ibm_pi_instance_network", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	instanceID := d.Get(Arg_InstanceID).(string)
	networkID := d.Get(Arg_NetworkID).(string)

	netC := instance.NewIBMPIInstanceNetworksClient(ctx, sess, cloudInstanceID)
	netsWrap, err := netC.Get(instanceID, networkID)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Get network failed: %s", err.Error()), "(Data) ibm_pi_instance_network", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	flatList := flattenPvmInstanceNetworksv2(netsWrap.Networks)
	m := flatList[0]

	d.SetId(fmt.Sprintf("%s/%s/%s", cloudInstanceID, instanceID, networkID))
	d.Set(Attr_ExternalIP, m[Attr_ExternalIP])
	d.Set(Attr_Href, m[Attr_Href])
	d.Set(Attr_IPAddress, m[Attr_IPAddress])
	d.Set(Attr_MacAddress, m[Attr_MacAddress])
	d.Set(Attr_NetworkID, m[Attr_NetworkID])
	d.Set(Attr_NetworkInterfaceID, m[Attr_NetworkInterfaceID])
	d.Set(Attr_NetworkName, m[Attr_NetworkName])
	d.Set(Attr_NetworkSecurityGroupIDs, m[Attr_NetworkSecurityGroupIDs])
	d.Set(Attr_NetworkSecurityGroupsHref, m[Attr_NetworkSecurityGroupsHref])
	d.Set(Attr_Type, m[Attr_Type])
	d.Set(Attr_Version, m[Attr_Version])

	return nil
}
