// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPINetworkPeerInterfaces() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPINetworkPeerInterfacesRead,

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			// Attributes
			Attr_PeerInterfaces: {
				Computed:    true,
				Description: "List of peer interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_DeviceID: {
							Computed:    true,
							Description: "Device ID of the peer interface.",
							Type:        schema.TypeString,
						},
						Attr_Name: {
							Computed:    true,
							Description: "Peer interface name.",
							Type:        schema.TypeString,
						},
						Attr_PeerInterfaceID: {
							Computed:    true,
							Description: "Peer interface ID.",
							Type:        schema.TypeString,
						},
						Attr_PeerType: {
							Computed:    true,
							Description: "Type of peer interface.",
							Type:        schema.TypeString,
						},
						Attr_PortID: {
							Computed:    true,
							Description: "Port ID of the peer interface.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPINetworkPeerInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)

	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, cloudInstanceID)
	networkdata, err := networkC.GetAllNetworkPeersInterfaces()
	if err != nil {
		return diag.FromErr(err)
	}
	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	peerInterfaces := []map[string]interface{}{}
	if len(networkdata) > 0 {
		for _, peerInterface := range networkdata {
			peerInterfaceMap := dataSourceIBMPINetworkPeerInterfacesPeerInterfaceToMap(peerInterface)
			peerInterfaces = append(peerInterfaces, peerInterfaceMap)
		}
		d.Set(Attr_PeerInterfaces, peerInterfaces)
	}

	return nil
}

func dataSourceIBMPINetworkPeerInterfacesPeerInterfaceToMap(peerInterface *models.PeerInterface) map[string]interface{} {
	peerInterfaceMap := make(map[string]interface{})
	peerInterfaceMap[Attr_DeviceID] = peerInterface.DeviceID
	peerInterfaceMap[Attr_Name] = peerInterface.Name
	peerInterfaceMap[Attr_PeerInterfaceID] = peerInterface.PeerInterfaceID
	peerInterfaceMap[Attr_PeerType] = peerInterface.PeerType
	peerInterfaceMap[Attr_PortID] = peerInterface.PortID
	return peerInterfaceMap
}
