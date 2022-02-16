// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"
	"net"
	"strconv"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Attributes and Arguments defined in data_source_ibm_pi_instance.go
func DataSourceIBMPIInstanceIP() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPIInstancesIPRead,
		Schema: map[string]*schema.Schema{
			PIInstanceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Server Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},
			PICloudInstanceID: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			PIInstanceNetworkName: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},

			// Computed attributes
			InstanceNetworkIP: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceIpOctet: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceNetworkMAC: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceNetworkID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceNetworkType: {
				Type:     schema.TypeString,
				Computed: true,
			},
			InstanceExternalIP: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPIInstancesIPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	networkName := d.Get(PIInstanceNetworkName).(string)
	powerC := instance.NewIBMPIInstanceClient(ctx, sess, cloudInstanceID)

	powervmdata, err := powerC.Get(d.Get(PIInstanceName).(string))
	if err != nil {
		return diag.FromErr(err)
	}

	for _, address := range powervmdata.Addresses {
		if address.NetworkName == networkName {
			log.Printf("Printing the ip %s", address.IP)
			d.SetId(address.NetworkID)
			d.Set(InstanceNetworkIP, address.IP)
			d.Set(InstanceNetworkID, address.NetworkID)
			d.Set(InstanceNetworkMAC, address.MacAddress)
			d.Set(InstanceExternalIP, address.ExternalIP)
			d.Set(InstanceNetworkType, address.Type)

			IPObject := net.ParseIP(address.IP).To4()

			d.Set(InstanceIpOctet, strconv.Itoa(int(IPObject[3])))

			return nil
		}
	}

	return diag.Errorf("failed to find instance ip that belongs to the given network")
}
