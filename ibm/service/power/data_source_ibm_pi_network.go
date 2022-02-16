// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	//"fmt"

	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	// Arguments
	PINetworkName  = "pi_network_name"
	PINetworkType  = "pi_network_type"
	PINetworkDNS   = "pi_dns"
	PINetworkCIDR  = "pi_cidr"
	PINetworkJumbo = "pi_network_jumbo"

	// Attributes
	NetworkAvailableIPCount = "available_ip_count"
	NetworkCIDR             = "cidr"
	NetworkDNS              = "dns"
	NetworkGateway          = "gateway"
	NetworkNetworkID        = "networkid"
	NetworkID               = "id"
	NetworkType             = "type"
	NetworkUsedIPCount      = "used_ip_count"
	NetworkUsedIPPercent    = "used_ip_percent"
	NetworkVlanID           = "vlan_id"
	NetworkJumbo            = "jumbo"
	NetworkName             = "name"
)

func DataSourceIBMPINetwork() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPINetworkRead,
		Schema: map[string]*schema.Schema{
			PINetworkName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Network Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed Attributes
			NetworkCIDR: {
				Type:     schema.TypeString,
				Computed: true,
			},

			NetworkType: {
				Type:     schema.TypeString,
				Computed: true,
			},

			NetworkVlanID: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			NetworkGateway: {
				Type:     schema.TypeString,
				Computed: true,
			},
			NetworkAvailableIPCount: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			NetworkUsedIPCount: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			NetworkUsedIPPercent: {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			NetworkDNS: {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			NetworkJumbo: {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPINetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)

	networkC := instance.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)
	networkdata, err := networkC.Get(d.Get(PINetworkName).(string))
	if err != nil || networkdata == nil {
		return diag.FromErr(err)
	}

	d.SetId(*networkdata.NetworkID)
	if networkdata.Cidr != nil {
		d.Set(NetworkCIDR, networkdata.Cidr)
	}
	if networkdata.Type != nil {
		d.Set(NetworkType, networkdata.Type)
	}
	d.Set(NetworkGateway, networkdata.Gateway)
	if networkdata.VlanID != nil {
		d.Set(NetworkVlanID, networkdata.VlanID)
	}
	if networkdata.IPAddressMetrics.Available != nil {
		d.Set(NetworkAvailableIPCount, networkdata.IPAddressMetrics.Available)
	}
	if networkdata.IPAddressMetrics.Used != nil {
		d.Set(NetworkUsedIPCount, networkdata.IPAddressMetrics.Used)
	}
	if networkdata.IPAddressMetrics.Utilization != nil {
		d.Set(NetworkUsedIPPercent, networkdata.IPAddressMetrics.Utilization)
	}
	if len(networkdata.DNSServers) > 0 {
		d.Set(NetworkDNS, networkdata.DNSServers)
	}
	d.Set(NetworkJumbo, networkdata.Jumbo)

	return nil

}
