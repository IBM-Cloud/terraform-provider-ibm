// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"log"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Arguments
	PIDhcpID              = "pi_dhcp_id"
	PIDhcpCloudConnection = "pi_cloud_connection_id"

	// Attributes
	DhcpID          = "dhcp_id"
	DhcpStatus      = "status"
	DhcpNetwork     = "network"
	DhcpLeases      = "leases"
	DhcpInstanceIP  = "instance_ip"
	DhcpInstanceMAC = "instance_mac"
	DhcpServers     = "servers"
)

func DataSourceIBMPIDhcp() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIDhcpRead,
		Schema: map[string]*schema.Schema{
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			PIDhcpID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the DHCP Server",
			},
			// Computed Attributes
			DhcpStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the DHCP Server",
			},
			DhcpNetwork: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DHCP Server private network",
			},
			DhcpLeases: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of DHCP Server PVM Instance leases",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DhcpInstanceIP: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP of the PVM Instance",
						},
						DhcpInstanceMAC: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The MAC Address of the PVM Instance",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPIDhcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	dhcpID := d.Get(PIDhcpID).(string)

	client := st.NewIBMPIDhcpClient(ctx, sess, cloudInstanceID)
	dhcpServer, err := client.Get(dhcpID)
	if err != nil {
		log.Printf("[DEBUG] get DHCP failed %v", err)
		return diag.FromErr(err)
	}

	d.SetId(*dhcpServer.ID)
	d.Set(DhcpStatus, *dhcpServer.Status)
	dhcpNetwork := dhcpServer.Network
	if dhcpNetwork != nil {
		d.Set(DhcpNetwork, *dhcpNetwork.ID)
	}
	dhcpLeases := dhcpServer.Leases
	if dhcpLeases != nil {
		leaseList := make([]map[string]string, len(dhcpLeases))
		for i, lease := range dhcpLeases {
			leaseList[i] = map[string]string{
				DhcpInstanceIP:  *lease.InstanceIP,
				DhcpInstanceMAC: *lease.InstanceMacAddress,
			}
		}
		d.Set(DhcpLeases, leaseList)
	}

	return nil
}
