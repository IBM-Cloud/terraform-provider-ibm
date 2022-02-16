// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	//"fmt"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	// Argumnets
	PINetworkPortName         = "pi_network_name"
	PINetworkPortDescription  = "pi_network_port_description"
	PINetworkPortIP           = "pi_network_port_ipaddress"
	PINetworkPortInstanceName = "pi_instance_name"
	PINetworkPortID           = "port_id"

	// Attributes
	NetworkPorts           = "network_ports"
	NetworkPortIP          = "ipaddress"
	NetworkPortMAC         = "macaddress"
	NetworkPortID          = "portid"
	NetworkPortStatus      = "status"
	NetworkPortHref        = "href"
	NetworkPortDescription = "description"
	NetworkPortPublicIP    = "public_ip"
	NetworkPortInstance    = "pvminstance"
)

func DataSourceIBMPINetworkPort() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceIBMPINetworkPortsRead,
		Schema: map[string]*schema.Schema{
			PINetworkPortName: {
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
			NetworkPorts: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						NetworkPortIP: {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						NetworkPortMAC: {
							Type:     schema.TypeString,
							Computed: true,
						},
						NetworkPortID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						NetworkPortStatus: {
							Type:     schema.TypeString,
							Computed: true,
						},
						NetworkPortHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
						NetworkPortDescription: {
							Type:     schema.TypeString,
							Required: true,
						},
						NetworkPortPublicIP: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPINetworkPortsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(PICloudInstanceID).(string)
	networkportC := instance.NewIBMPINetworkClient(ctx, sess, cloudInstanceID)
	networkportdata, err := networkportC.GetAllPorts(d.Get(PINetworkPortName).(string))
	if err != nil {
		return diag.FromErr(err)
	}
	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	d.Set(NetworkPorts, flattenNetworkPorts(networkportdata.Ports))

	return nil

}

func flattenNetworkPorts(networkPorts []*models.NetworkPort) interface{} {
	result := make([]map[string]interface{}, 0, len(networkPorts))
	log.Printf("the number of ports is %d", len(networkPorts))
	for _, i := range networkPorts {
		l := map[string]interface{}{
			NetworkPortID:          *i.PortID,
			NetworkPortStatus:      *i.Status,
			NetworkPortHref:        i.Href,
			NetworkPortIP:          *i.IPAddress,
			NetworkPortMAC:         *i.MacAddress,
			NetworkPortDescription: *i.Description,
			NetworkPortPublicIP:    i.ExternalIP,
		}

		result = append(result, l)
	}
	return result
}
