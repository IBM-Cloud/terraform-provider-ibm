// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"log"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceIBMPIRoutes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIRoutesRead,
		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},

			// Attributes
			Attr_Routes: {
				Computed:    true,
				Description: "List of routes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_RouteID: {
							Description: "Unique ID of the route.",
							Required:    true,
							Type:        schema.TypeString,
						},
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
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPIRoutesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)

	routes, err := client.GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(Attr_Routes, flattenRoutes(routes.Routes, meta))

	return nil
}

func flattenRoutes(routes []*models.Route, meta interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, len(routes))
	for _, r := range routes {
		route := map[string]interface{}{
			Attr_RouteID:         r.ID,
			Attr_Action:          r.Action,
			Attr_Advertise:       r.Advertise,
			Attr_Destination:     r.Destination,
			Attr_DestinationType: r.DestinationType,
			Attr_Enabled:         r.Enabled,
			Attr_Name:            r.Name,
			Attr_NextHop:         r.NextHop,
			Attr_NextHopType:     r.NextHopType,
			Attr_State:           r.State,
		}

		if r.Crn != nil {
			route[Attr_CRN] = r.Crn
			tags, err := flex.GetGlobalTagsUsingCRN(meta, *r.Crn, "", UserTagType)
			if err != nil {
				log.Printf("Error on get of pi route (%s) user_tags: %s", *r.ID, err)
			}
			route[Attr_UserTags] = tags
		}
		result = append(result, route)
	}

	return result
}
