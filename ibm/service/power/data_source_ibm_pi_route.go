// Copyright IBM Corp. 2025 All Rights Reserved.
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

func DataSourceIBMPIRoute() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIRouteRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_RouteID: {
				Description:  "Unique ID of the route.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_Action: {
				Computed:    true,
				Description: "The route action.",
				Type:        schema.TypeString,
			},
			Attr_Advertise: {
				Computed:    true,
				Description: "Indicates if the route is advertised.",
				Type:        schema.TypeString,
			},
			Attr_CRN: {
				Computed:    true,
				Description: "The CRN of this resource.",
				Type:        schema.TypeString,
			},
			Attr_Destination: {
				Computed:    true,
				Description: "The route destination.",
				Type:        schema.TypeString,
			},
			Attr_DestinationType: {
				Computed:    true,
				Description: "The destination type.",
				Type:        schema.TypeString,
			},
			Attr_Enabled: {
				Computed:    true,
				Description: "Indicates if the route should be enabled in the fabric.",
				Type:        schema.TypeBool,
			},
			Attr_Name: {
				Computed:    true,
				Description: "Name of the route.",
				Type:        schema.TypeString,
			},
			Attr_NextHop: {
				Computed:    true,
				Description: "The next hop in the route.",
				Type:        schema.TypeString,
			},
			Attr_NextHopType: {
				Computed:    true,
				Description: "The next hop type.",
				Type:        schema.TypeString,
			},
			Attr_State: {
				Computed:    true,
				Description: "The state of the route.",
				Type:        schema.TypeString,
			},
			Attr_UserTags: {
				Computed:    true,
				Description: "List of user tags attached to the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},
		},
	}
}

func dataSourceIBMPIRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	routeID := d.Get(Arg_RouteID).(string)
	client := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)

	route, err := client.Get(routeID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set(Attr_Action, route.Action)
	d.Set(Attr_Advertise, route.Advertise)
	if route.Crn != nil {
		d.Set(Attr_CRN, *route.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, *route.Crn, "", UserTagType)
		if err != nil {
			log.Printf("Error on get of route (%s) user_tags: %s", *route.ID, err)
		}
		d.Set(Attr_UserTags, tags)
	}
	d.Set(Attr_Destination, route.Destination)
	d.Set(Attr_DestinationType, route.DestinationType)
	d.Set(Attr_Enabled, route.Enabled)
	d.SetId(*route.ID)
	d.Set(Attr_Name, route.Name)
	d.Set(Attr_NextHop, route.NextHop)
	d.Set(Attr_NextHopType, route.NextHopType)
	d.Set(Attr_State, route.State)

	return nil
}
