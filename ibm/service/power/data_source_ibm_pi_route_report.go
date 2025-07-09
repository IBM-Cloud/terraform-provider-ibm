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

func DataSourceIBMPIRouteReport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPIRouteReportRead,
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
				Description: "A report of routes in a workspace.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						Attr_Destination: {
							Computed:    true,
							Description: "The destination CIDR.",
							Type:        schema.TypeString,
						},
						Attr_Type: {
							Computed:    true,
							Description: "The route type.",
							Type:        schema.TypeString,
						},
					},
				},
				Type: schema.TypeList,
			},
		},
	}
}

func dataSourceIBMPIRouteReportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	client := instance.NewIBMPIRouteClient(ctx, sess, cloudInstanceID)

	routeReports, err := client.GetRouteReport()
	if err != nil {
		return diag.FromErr(err)
	}
	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set(Attr_Routes, flattenRouteReports(routeReports.Routes))

	return nil
}

func flattenRouteReports(routes []*models.RouteReportRoute) []map[string]interface{} {
	result := make([]map[string]interface{}, len(routes))
	for _, r := range routes {
		routeReport := map[string]interface{}{
			Attr_Destination: r.Destination,
			Attr_Type:        r.Type,
		}

		result = append(result, routeReport)
	}

	return result
}
