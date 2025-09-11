// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIBMPINetworkPeerRouteFilter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPINetworkPeerRouteFilterRead,

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_NetworkPeerID: {
				Description:  "Network peer ID.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_RouteFilterID: {
				Description:  "Route filter ID.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_Action: {
				Computed:    true,
				Description: "Action of the filter `allow` or `deny`.",
				Type:        schema.TypeString,
			},
			Attr_CreationDate: {
				Computed:    true,
				Description: "Time stamp for create route filter.",
				Type:        schema.TypeString,
			},
			Attr_Direction: {
				Computed:    true,
				Description: "Direction of the filter `import` or `export`.",
				Type:        schema.TypeString,
			},
			Attr_Error: {
				Computed:    true,
				Description: "Error description.",
				Type:        schema.TypeString,
			},
			Attr_GE: {
				Computed:    true,
				Description: "The minimum matching length of the prefix-set.",
				Type:        schema.TypeInt,
			},
			Attr_Index: {
				Computed:    true,
				Description: "Priority or order of the filter.",
				Type:        schema.TypeInt,
			},
			Attr_LE: {
				Computed:    true,
				Description: "The maximum matching length of the prefix-set.",
				Type:        schema.TypeInt,
			},
			Attr_Prefix: {
				Computed:    true,
				Description: "IP prefix representing an address and mask length of the prefix-set.",
				Type:        schema.TypeString,
			},
			Attr_RouteFilterID: {
				Computed:    true,
				Description: "Route filter ID.",
				Type:        schema.TypeString,
			},
			Attr_State: {
				Computed:    true,
				Description: "Status of the route filter.",
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceIBMPINetworkPeerRouteFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	networkPeerID := d.Get(Arg_NetworkPeerID).(string)
	routeFilterID := d.Get(Arg_RouteFilterID).(string)
	networkC := instance.NewIBMPINetworkPeerClient(ctx, sess, cloudInstanceID)
	routedata, err := networkC.GetNetworkPeersRouteFilter(networkPeerID, routeFilterID)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*routedata.RouteFilterID)
	d.Set(Attr_Action, routedata.Action)
	d.Set(Attr_CreationDate, routedata.CreationDate)
	d.Set(Attr_Direction, routedata.Direction)
	d.Set(Attr_Error, routedata.Error)
	d.Set(Attr_GE, routedata.GE)
	d.Set(Attr_Index, routedata.Index)
	d.Set(Attr_LE, routedata.LE)
	d.Set(Attr_Prefix, routedata.Prefix)
	d.Set(Attr_RouteFilterID, routedata.RouteFilterID)
	d.Set(Attr_State, routedata.State)
	return nil
}
