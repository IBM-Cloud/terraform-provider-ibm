// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package logsrouter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
)

func DataSourceIBMLogsRouterRoutes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMLogsRouterRoutesRead,

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
							Description: "The UUID of the route resource.",
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
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The action if the inclusion_filters matches, default is `send` action.",
									},
									"targets": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The target ID List. Platform logs will be sent to all targets listed in the rule. You can include targets from other regions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The target uuid for a pre-defined platform logs router target.",
												},
												"crn": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The CRN of a pre-defined logs-router target.",
												},
												"name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of a pre-defined logs-router target.",
												},
												"target_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of the target.",
												},
											},
										},
									},
									"inclusion_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A list of conditions to be satisfied for routing platform logs to pre-defined target.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operand": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Part of CRN that can be compared with values. Currently only location is supported.",
												},
												"operator": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support up to 20 values in the array.",
												},
												"values": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The provided string values of the operand to be compared with.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
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
						"managed_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Present when the route is enterprise-managed (`managed_by: enterprise`).",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMLogsRouterRoutesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_router_routes", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listRoutesOptions := &logsrouterv3.ListRoutesOptions{}

	routeCollection, _, err := logsRouterClient.ListRoutesWithContext(context, listRoutesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListRoutesWithContext failed: %s", err.Error()), "(Data) ibm_logs_router_routes", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchRoutes []logsrouterv3.Route
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range routeCollection.Routes {
			if *data.Name == name {
				matchRoutes = append(matchRoutes, data)
			}
		}
	} else {
		matchRoutes = routeCollection.Routes
	}
	routeCollection.Routes = matchRoutes

	if suppliedFilter {
		if len(routeCollection.Routes) == 0 {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("no Routes found with name %s", name), "(Data) ibm_logs_router_routes", "read", "no-collection-found").GetDiag()
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMLogsRouterRoutesID(d))
	}

	routes := []map[string]interface{}{}
	for _, routesItem := range routeCollection.Routes {
		routesItemMap, err := DataSourceIBMLogsRouterRoutesRouteToMap(&routesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_logs_router_routes", "read", "routes-to-map").GetDiag()
		}
		routes = append(routes, routesItemMap)
	}
	if err = d.Set("routes", routes); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting routes: %s", err), "(Data) ibm_logs_router_routes", "read", "set-routes").GetDiag()
	}

	return nil
}

// dataSourceIBMLogsRouterRoutesID returns a reasonable ID for the list.
func dataSourceIBMLogsRouterRoutesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMLogsRouterRoutesRouteToMap(model *logsrouterv3.Route) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["name"] = *model.Name
	modelMap["crn"] = *model.CRN
	rules := []map[string]interface{}{}
	for _, rulesItem := range model.Rules {
		rulesItemMap, err := DataSourceIBMLogsRouterRoutesRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		rules = append(rules, rulesItemMap)
	}
	modelMap["rules"] = rules
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	if model.ManagedBy != nil {
		modelMap["managed_by"] = *model.ManagedBy
	}
	return modelMap, nil
}

func DataSourceIBMLogsRouterRoutesRuleToMap(model *logsrouterv3.Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Action != nil {
		modelMap["action"] = *model.Action
	}
	targets := []map[string]interface{}{}
	for _, targetsItem := range model.Targets {
		targetsItemMap, err := DataSourceIBMLogsRouterRoutesTargetReferenceToMap(&targetsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		targets = append(targets, targetsItemMap)
	}
	modelMap["targets"] = targets
	inclusionFilters := []map[string]interface{}{}
	for _, inclusionFiltersItem := range model.InclusionFilters {
		inclusionFiltersItemMap, err := DataSourceIBMLogsRouterRoutesInclusionFilterToMap(&inclusionFiltersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		inclusionFilters = append(inclusionFilters, inclusionFiltersItemMap)
	}
	modelMap["inclusion_filters"] = inclusionFilters
	return modelMap, nil
}

func DataSourceIBMLogsRouterRoutesTargetReferenceToMap(model *logsrouterv3.TargetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["crn"] = *model.CRN
	modelMap["name"] = *model.Name
	modelMap["target_type"] = *model.TargetType
	return modelMap, nil
}

func DataSourceIBMLogsRouterRoutesInclusionFilterToMap(model *logsrouterv3.InclusionFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["operand"] = *model.Operand
	modelMap["operator"] = *model.Operator
	modelMap["values"] = model.Values
	return modelMap, nil
}
