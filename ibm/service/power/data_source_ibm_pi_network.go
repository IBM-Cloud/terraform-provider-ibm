// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPINetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPINetworkRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_NetworkID: {
				AtLeastOneOf:  []string{Arg_NetworkID, Arg_NetworkName},
				ConflictsWith: []string{Arg_NetworkName},
				Description:   "The network ID.",
				Optional:      true,
				Type:          schema.TypeString,
				ValidateFunc:  validation.NoZeroValues,
			},
			Arg_NetworkName: {
				AtLeastOneOf:  []string{Arg_NetworkID, Arg_NetworkName},
				ConflictsWith: []string{Arg_NetworkID},
				Deprecated:    "The pi_network_name field is deprecated. Please use pi_network_id instead",
				Description:   "The unique identifier or name of a network.",
				Optional:      true,
				Type:          schema.TypeString,
				ValidateFunc:  validation.NoZeroValues,
			},

			// Attributes
			Attr_Advertise: {
				Computed:    true,
				Description: "Indicates if the network is advertised.",
				Type:        schema.TypeString,
			},
			Attr_ARPBroadcast: {
				Computed:    true,
				Description: "Indicates if ARP Broadcast is enabled.",
				Type:        schema.TypeString,
			},
			Attr_AvailableIPCount: {
				Computed:    true,
				Description: "The total number of IP addresses that you have in your network.",
				Type:        schema.TypeFloat,
			},
			Attr_CIDR: {
				Computed:    true,
				Description: "The CIDR of the network.",
				Type:        schema.TypeString,
			},
			Attr_CRN: {
				Computed:    true,
				Description: "The CRN of this resource.",
				Type:        schema.TypeString,
			},
			Attr_DNS: {
				Computed:    true,
				Description: "The DNS Servers for the network.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeSet,
			},
			Attr_Gateway: {
				Computed:    true,
				Description: "The network gateway that is attached to your network.",
				Type:        schema.TypeString,
			},
			Attr_MTU: {
				Computed:    true,
				Description: "Maximum Transmission Unit option of the network.",
				Type:        schema.TypeInt,
			},
			Attr_Name: {
				Computed:    true,
				Description: "The name of the network.",
				Type:        schema.TypeString,
			},
			Attr_NetworkAddressTranslation: {
				Computed:    true,
				Description: "Contains the network address translation details (for on prem locations only).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_SourceIP: {
							Computed:    true,
							Description: "source IP address.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
			Attr_PeerID: {
				Computed:    true,
				Deprecated:  "This field is deprecated",
				Description: "Network peer ID (for on prem locations only).",
				Type:        schema.TypeString,
			},
			Attr_Type: {
				Computed:    true,
				Description: "The type of network.",
				Type:        schema.TypeString,
			},
			Attr_UsedIPCount: {
				Computed:    true,
				Description: "The number of used IP addresses.",
				Type:        schema.TypeFloat,
			},
			Attr_UsedIPPercent: {
				Computed:    true,
				Description: "The percentage of IP addresses used.",
				Type:        schema.TypeFloat,
			},
			Attr_UserTags: {
				Computed:    true,
				Description: "List of user tags attached to the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
			Attr_VLanID: {
				Computed:    true,
				Description: "The VLAN ID that the network is connected to.",
				Type:        schema.TypeInt,
			},
		},
	}
}

func dataSourceIBMPINetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	var networkID string
	if v, ok := d.GetOk(Arg_NetworkID); ok {
		networkID = v.(string)
	} else if v, ok := d.GetOk(Arg_NetworkName); ok {
		networkID = v.(string)
	}

	networkC := instance.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)
	networkdata, err := networkC.Get(networkID)
	if err != nil || networkdata == nil {
		return diag.FromErr(err)
	}

	d.SetId(*networkdata.NetworkID)
	d.Set(Attr_Advertise, networkdata.Advertise)
	d.Set(Attr_ARPBroadcast, networkdata.ArpBroadcast)
	if networkdata.IPAddressMetrics.Available != nil {
		d.Set(Attr_AvailableIPCount, networkdata.IPAddressMetrics.Available)
	}
	if networkdata.Cidr != nil {
		d.Set(Attr_CIDR, networkdata.Cidr)
	}
	if networkdata.Crn != "" {
		d.Set(Attr_CRN, networkdata.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(networkdata.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of pi network (%s) user_tags: %s", *networkdata.NetworkID, err)
		}
		d.Set(Attr_UserTags, tags)
	}
	if len(networkdata.DNSServers) > 0 {
		d.Set(Attr_DNS, networkdata.DNSServers)
	}
	d.Set(Attr_Gateway, networkdata.Gateway)
	d.Set(Attr_MTU, networkdata.Mtu)
	if networkdata.Name != nil {
		d.Set(Attr_Name, networkdata.Name)
	}
	networkAddressTranslation := []map[string]interface{}{}
	if networkdata.NetworkAddressTranslation != nil {
		natMap := networkAddressTranslationToMap(networkdata.NetworkAddressTranslation)
		networkAddressTranslation = append(networkAddressTranslation, natMap)
	}
	d.Set(Attr_NetworkAddressTranslation, networkAddressTranslation)
	d.Set(Attr_PeerID, networkdata.PeerID)
	if networkdata.Type != nil {
		d.Set(Attr_Type, networkdata.Type)
	}
	if networkdata.IPAddressMetrics.Used != nil {
		d.Set(Attr_UsedIPCount, networkdata.IPAddressMetrics.Used)
	}
	if networkdata.IPAddressMetrics.Utilization != nil {
		d.Set(Attr_UsedIPPercent, networkdata.IPAddressMetrics.Utilization)
	}
	if networkdata.VlanID != nil {
		d.Set(Attr_VLanID, networkdata.VlanID)
	}

	return nil
}
