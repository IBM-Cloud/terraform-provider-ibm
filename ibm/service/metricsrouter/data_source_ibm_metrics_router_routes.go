// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func DataSourceIBMMetricsRouterRoutes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMMetricsRouterRoutesRead,

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
										Description: "The target ID List. All the metrics will be sent to all targets listed in the rule. You can include targets from other regions.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"inclusion_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A list of conditions to be satisfied for routing metrics to pre-defined target.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operand": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Part of CRN that can be compared with values.",
												},
												"operator": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support upto 20 values in the array.",
												},
												"value": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The provided values of the in operand to be compared with.",
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

func dataSourceIBMMetricsRouterRoutesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	listRoutesOptions := &metricsrouterv3.ListRoutesOptions{}

	routeList, response, err := metricsRouterClient.ListRoutesWithContext(context, listRoutesOptions)
	if err != nil {
		log.Printf("[DEBUG] ListRoutesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListRoutesWithContext failed %s\n%s", err, response))
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchRoutes []metricsrouterv3.Route
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
			return diag.FromErr(fmt.Errorf("no Routes found with name %s", name))
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMMetricsRouterRoutesID(d))
	}

	routes := []map[string]interface{}{}
	if routeList.Routes != nil {
		for _, modelItem := range routeList.Routes {
			modelMap, err := dataSourceIBMMetricsRouterRoutesRouteToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			routes = append(routes, modelMap)
		}
	}
	if err = d.Set("routes", routes); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting routes %s", err))
	}

	return nil
}

// dataSourceIBMMetricsRouterRoutesID returns a reasonable ID for the list.
func dataSourceIBMMetricsRouterRoutesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIBMMetricsRouterRoutesRouteToMap(model *metricsrouterv3.Route) (map[string]interface{}, error) {
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
		modelMap["version"] = *model.Version
	}
	if model.Rules != nil {
		rules := []map[string]interface{}{}
		for _, rulesItem := range model.Rules {
			rulesItemMap, err := dataSourceIBMMetricsRouterRoutesRuleToMap(&rulesItem)
			if err != nil {
				return modelMap, err
			}
			rules = append(rules, rulesItemMap)
		}
		modelMap["rules"] = rules
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.APIVersion != nil {
		modelMap["api_version"] = *model.APIVersion
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}

func dataSourceIBMMetricsRouterRoutesRuleToMap(model *metricsrouterv3.Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetIds != nil {
		modelMap["target_ids"] = model.TargetIds
	}
	if model.InclusionFilters != nil {
		inclusionFilters := []map[string]interface{}{}
		for _, inclusionFiltersItem := range model.InclusionFilters {
			inclusionFiltersItemMap, err := dataSourceIBMMetricsRouterRoutesInclusionFilterToMap(&inclusionFiltersItem)
			if err != nil {
				return modelMap, err
			}
			inclusionFilters = append(inclusionFilters, inclusionFiltersItemMap)
		}
		modelMap["inclusion_filters"] = inclusionFilters
	}
	return modelMap, nil
}

func dataSourceIBMMetricsRouterRoutesInclusionFilterToMap(model *metricsrouterv3.InclusionFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Operand != nil {
		modelMap["operand"] = *model.Operand
	}
	if model.Operator != nil {
		modelMap["operator"] = *model.Operator
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}
