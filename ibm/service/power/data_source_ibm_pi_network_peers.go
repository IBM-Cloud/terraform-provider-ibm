// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPINetworkPeers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPINetworkPeersRead,

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_NetworkPeers: {
				Computed:    true,
				Description: "List of network peers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_CreationDate: {
							Computed:    true,
							Description: "Time stamp for create network peer.",
							Type:        schema.TypeString,
						},
						Attr_CustomerASN: {
							Computed:    true,
							Description: "ASN number at customer network side.",
							Type:        schema.TypeInt,
						},
						Attr_CustomerCIDR: {
							Computed:    true,
							Description: "IP address used for configuring customer network interface with network subnet mask.",
							Type:        schema.TypeString,
						},
						Attr_DefaultExportRouteFilter: {
							Computed:    true,
							Description: "Default action for export route filter.",
							Type:        schema.TypeString,
						},
						Attr_DefaultImportRouteFilter: {
							Computed:    true,
							Description: "Default action for import route filter.",
							Type:        schema.TypeString,
						},
						Attr_Description: {
							Computed:    true,
							Description: "[Depracated] Description of the network peer.",
							Type:        schema.TypeString,
						},
						Attr_Error: {
							Computed:    true,
							Description: "Error description.",
							Type:        schema.TypeString,
						},
						Attr_ExportRouteFilters: routeFilterSchema("List of export route filters."),
						Attr_IBMASN: {
							Computed:    true,
							Description: "ASN number at IBM PowerVS side.",
							Type:        schema.TypeInt,
						},
						Attr_IBMCIDR: {
							Computed:    true,
							Description: "IP address used for configuring IBM network interface with network subnet mask.",
							Type:        schema.TypeString,
						},
						Attr_ID: {
							Computed:    true,
							Description: "ID of the network peer.",
							Type:        schema.TypeString,
						},
						Attr_ImportRouteFilters: routeFilterSchema("List of import route filters."),
						Attr_Name: {
							Computed:    true,
							Description: "User defined name.",
							Type:        schema.TypeString,
						},
						Attr_PeerInterfaceID: {
							Computed:    true,
							Description: "Peer interface id.",
							Type:        schema.TypeString,
						},
						Attr_State: {
							Computed:    true,
							Description: "Status of the network peer.",
							Type:        schema.TypeString,
						},
						Attr_Type: {
							Computed:    true,
							Description: "Type of the peer network.",
							Type:        schema.TypeString,
						},
						Attr_UpdatedDate: {
							Computed:    true,
							Description: "Time stamp for update network peer.",
							Type:        schema.TypeString,
						},
						Attr_VLAN: {
							Computed:    true,
							Description: "A vlan configured at the customer network.",
							Type:        schema.TypeInt,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPINetworkPeersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)

	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, cloudInstanceID)
	networkdata, err := networkC.GetNetworkPeers()
	if err != nil {
		return diag.FromErr(err)
	}
	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	networkPeers := []map[string]interface{}{}
	if networkdata.NetworkPeers != nil {
		for _, np := range networkdata.NetworkPeers {
			npMap := dataSourceIBMPINetworkPeersNetworkPeerToMap(np)

			networkPeers = append(networkPeers, npMap)
		}
	}
	d.Set(Attr_NetworkPeers, networkPeers)

	return nil
}

func dataSourceIBMPINetworkPeersNetworkPeerToMap(np *models.NetworkPeer) map[string]interface{} {
	npMap := make(map[string]interface{})
	npMap[Attr_CreationDate] = np.CreationDate
	npMap[Attr_CustomerASN] = np.CustomerASN
	npMap[Attr_CustomerCIDR] = np.CustomerCidr
	npMap[Attr_DefaultExportRouteFilter] = np.DefaultExportRouteFilter
	npMap[Attr_DefaultImportRouteFilter] = np.DefaultImportRouteFilter
	npMap[Attr_Error] = np.Error
	exportRouteFilters := []map[string]interface{}{}
	if len(np.ExportRouteFilters) > 0 {
		for _, erf := range np.ExportRouteFilters {
			exportRouteFilter := dataSourceIBMPINetworkPeerRouteFilterToMap(erf)
			exportRouteFilters = append(exportRouteFilters, exportRouteFilter)
		}
	}
	npMap[Attr_ExportRouteFilters] = exportRouteFilters
	npMap[Attr_IBMASN] = np.IbmASN
	npMap[Attr_IBMCIDR] = np.IbmCidr
	npMap[Attr_ID] = np.ID
	importRouteFilters := []map[string]interface{}{}
	if len(np.ImportRouteFilters) > 0 {
		for _, irf := range np.ImportRouteFilters {
			importRouteFilter := dataSourceIBMPINetworkPeerRouteFilterToMap(irf)
			importRouteFilters = append(importRouteFilters, importRouteFilter)
		}
	}
	if np.ID != nil {
		npMap[Attr_ID] = np.ID
	}
	npMap[Attr_ImportRouteFilters] = importRouteFilters
	if np.Name != nil {
		npMap[Attr_Name] = np.Name
	}
	npMap[Attr_PeerInterfaceID] = np.PeerInterfaceID
	if np.Type != nil {
		npMap[Attr_Type] = np.Type
	}
	npMap[Attr_VLAN] = np.Vlan
	return npMap
}

func dataSourceIBMPINetworkPeerRouteFilterToMap(rf *models.RouteFilter) map[string]interface{} {
	rfMap := make(map[string]interface{})

	if rf.Action != nil {
		rfMap[Attr_Action] = rf.Action
	}
	if rf.CreationDate != nil {
		rfMap[Attr_CreationDate] = rf.CreationDate
	}
	if rf.Direction != nil {
		rfMap[Attr_Direction] = rf.Direction
	}
	if rf.Error != nil {
		rfMap[Attr_Error] = rf.Error
	}
	if rf.GE != nil {
		rfMap[Attr_GE] = rf.GE
	}
	if rf.Index != nil {
		rfMap[Attr_Index] = rf.Index
	}
	if rf.LE != nil {
		rfMap[Attr_LE] = rf.LE
	}
	if rf.Prefix != nil {
		rfMap[Attr_Prefix] = rf.Prefix
	}
	if rf.RouteFilterID != nil {
		rfMap[Attr_RouteFilterID] = rf.RouteFilterID
	}
	if rf.State != nil {
		rfMap[Attr_State] = rf.State
	}
	return rfMap
}
