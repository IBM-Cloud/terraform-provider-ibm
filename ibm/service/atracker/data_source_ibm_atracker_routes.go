// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.101.0-62624c1e-20250225-192301
 */

package atracker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMAtrackerRoutes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMAtrackerRoutesRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the route.",
			},
			"routes": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of route resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uuid of the route resource.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the route.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crn of the route resource.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version of the route.",
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The target ID List. All the events will be send to all targets listed in the rule. You can include targets from other regions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"locations": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Logs from these locations will be sent to the targets specified. Locations is a superset of regions including global and *.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp of the route creation time.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp of the route last updated time.",
						},
						"api_version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The API version of the route.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An optional message containing information about the route.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMAtrackerRoutesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_atracker_routes", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listRoutesOptions := &atrackerv2.ListRoutesOptions{}

	routeList, _, err := atrackerClient.ListRoutesWithContext(context, listRoutesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListRoutesWithContext failed: %s", err.Error()), "(Data) ibm_atracker_routes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchRoutes []atrackerv2.Route
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range routeList.Routes {
			if *data.Name == name {
				matchRoutes = append(matchRoutes, data)
			}
		}
	} else {
		matchRoutes = routeList.Routes
	}
	routeList.Routes = matchRoutes

	if suppliedFilter {
		if len(routeList.Routes) == 0 {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("no Routes found with name %s", name), "(Data) ibm_atracker_routes", "read", "no-collection-found").GetDiag()
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMAtrackerRoutesID(d))
	}

	routes := []map[string]interface{}{}
	for _, routesItem := range routeList.Routes {
		routesItemMap, err := DataSourceIBMAtrackerRoutesRouteToMap(&routesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_atracker_routes", "read", "routes-to-map").GetDiag()
		}
		routes = append(routes, routesItemMap)
	}
	if err = d.Set("routes", routes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting routes: %s", err), "(Data) ibm_atracker_routes", "read", "set-routes").GetDiag()
	}

	return nil
}

// dataSourceIBMAtrackerRoutesID returns a reasonable ID for the list.
func dataSourceIBMAtrackerRoutesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMAtrackerRoutesRouteToMap(model *atrackerv2.Route) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Version != nil {
		modelMap["version"] = flex.IntValue(model.Version)
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIBMAtrackerRoutesRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.APIVersion != nil {
		modelMap["api_version"] = flex.IntValue(model.APIVersion)
	} else {
		modelMap["api_version"] = 1
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}

func DataSourceIBMAtrackerRoutesRuleToMap(model *atrackerv2.Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetIds != nil {
		modelMap["target_ids"] = model.TargetIds
	}
	if model.Locations != nil {
		modelMap["locations"] = model.Locations
	}
	return modelMap, nil
}
