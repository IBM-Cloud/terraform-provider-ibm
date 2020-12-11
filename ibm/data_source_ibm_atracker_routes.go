/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func dataSourceIBMAtrackerRoutes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMAtrackerRoutesRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of this route.",
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
							Description: "The uuid of this route resource.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of this route.",
						},
						"instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uuid of ATracker services in this region.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crn of this route type resource.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version of this route.",
						},
						"receive_global_events": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not all global events should be forwarded to this region.",
						},
						"rules": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The routing rules that will be evaluated in their order of the array.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The target ID List. Only one target id is supported. For regional route, the id must be V4 uuid of a target in the same region. For global route, it will be region-code and target-id separated by colon.",
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
		},
	}
}

func dataSourceIBMAtrackerRoutesRead(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	listRoutesOptions := &atrackerv1.ListRoutesOptions{}

	routeList, response, err := atrackerClient.ListRoutes(listRoutesOptions)
	if err != nil {
		log.Printf("[DEBUG] ListRoutes failed %s\n%s", err, response)
		return err
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchRoutes []atrackerv1.Route
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

	if len(routeList.Routes) == 0 {
		return fmt.Errorf("no Routes found with name %s\nIf not specified, please specify more filters", name)
	}

	if suppliedFilter {
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMAtrackerRoutesID(d))
	}

	if routeList.Routes != nil {
		err = d.Set("routes", dataSourceRouteListFlattenRoutes(routeList.Routes))
		if err != nil {
			log.Printf("Error flattening routes list %s", err)
		}
	}

	return nil
}

// dataSourceIBMAtrackerRoutesID returns a reasonable ID for the list.
func dataSourceIBMAtrackerRoutesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceRouteListFlattenRoutes(result []atrackerv1.Route) (routes []map[string]interface{}) {
	for _, routesItem := range result {
		routes = append(routes, dataSourceRouteListRoutesToMap(routesItem))
	}

	return routes
}

func dataSourceRouteListRoutesToMap(routesItem atrackerv1.Route) (routesMap map[string]interface{}) {
	routesMap = map[string]interface{}{}

	if routesItem.ID != nil {
		routesMap["id"] = routesItem.ID
	}
	if routesItem.Name != nil {
		routesMap["name"] = routesItem.Name
	}
	if routesItem.InstanceID != nil {
		routesMap["instance_id"] = routesItem.InstanceID
	}
	if routesItem.CRN != nil {
		routesMap["crn"] = routesItem.CRN
	}
	if routesItem.Version != nil {
		routesMap["version"] = routesItem.Version
	}
	if routesItem.ReceiveGlobalEvents != nil {
		routesMap["receive_global_events"] = routesItem.ReceiveGlobalEvents
	}
	if routesItem.Rules != nil {
		rulesList := []map[string]interface{}{}
		for _, rulesItem := range routesItem.Rules {
			rulesList = append(rulesList, dataSourceRouteListRoutesRulesToMap(rulesItem))
		}
		routesMap["rules"] = rulesList
	}

	return routesMap
}

func dataSourceRouteListRoutesRulesToMap(rulesItem atrackerv1.Rule) (rulesMap map[string]interface{}) {
	rulesMap = map[string]interface{}{}

	if rulesItem.TargetIds != nil {
		rulesMap["target_ids"] = rulesItem.TargetIds
	}

	return rulesMap
}
