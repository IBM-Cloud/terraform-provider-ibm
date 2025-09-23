// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package appconfiguration

import (
	"net/url"
	"reflect"
	"strconv"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func DataSourceIBMAppConfigIntegrations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigIntegrationsRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of records returned in the current response.",
			},
			"next": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the next list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"first": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the first page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"previous": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the previous list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"last": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the last page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"integrations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of integrations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"integration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration ID",
						},
						"integration_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Integration Type",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAppConfigIntegrationsRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("getAppConfigClient failed %s", err)
	}

	options := &appconfigurationv1.ListIntegrationsOptions{}

	var integrationsList *appconfigurationv1.IntegrationList
	var offset int64
	var limit int64 = 10
	var isLimit bool
	finalList := []appconfigurationv1.Integration{}

	if _, ok := GetFieldExists(d, "limit"); ok {
		isLimit = true
		limit = int64(d.Get("limit").(int))
	}
	options.SetLimit(limit)

	if _, ok := GetFieldExists(d, "offset"); ok {
		offset = int64(d.Get("offset").(int))
	}

	for {
		options.Offset = &offset
		result, response, err := appconfigClient.ListIntegrations(options)
		integrationsList = result
		if err != nil {
			return flex.FmtErrorf("List Integrations failed %s\n%s", err, response)
		}
		if isLimit {
			offset = 0
		} else {
			offset = dataSourceIntegrationsListGetNext(result.Next)
		}
		finalList = append(finalList, result.Integrations...)
		if offset == 0 {
			break
		}
	}

	integrationsList.Integrations = finalList

	d.SetId(guid)

	if integrationsList.Integrations != nil {
		err = d.Set("integrations", dataSourceIntegrationListFlattenintegrations(integrationsList.Integrations))
		if err != nil {
			return flex.FmtErrorf("[ERROR] Error setting integrations %s", err)
		}
	}
	if integrationsList.Limit != nil {
		if err = d.Set("limit", integrationsList.Limit); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting limit: %s", err)
		}
	}
	if integrationsList.Offset != nil {
		if err = d.Set("offset", integrationsList.Offset); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting offset: %s", err)
		}
	}
	if integrationsList.TotalCount != nil {
		if err = d.Set("total_count", integrationsList.TotalCount); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting total_count: %s", err)
		}
	}

	return nil
}

func dataSourceIntegrationListFlattenintegrations(integrations []appconfigurationv1.Integration) (integrationsMap []map[string]any) {
	for _, integration := range integrations {
		integrationMap := map[string]any{}
		integrationMap["integration_id"] = *integration.IntegrationID
		integrationMap["integration_type"] = *integration.IntegrationType
		integrationsMap = append(integrationsMap, integrationMap)
	}
	return integrationsMap
}

func dataSourceIntegrationsListGetNext(next any) int64 {
	if reflect.ValueOf(next).IsNil() {
		return 0
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("href").Elem().String())
	if err != nil {
		return 0
	}

	q := u.Query()
	var page string

	if q.Get("start") != "" {
		page = q.Get("start")
	} else if q.Get("offset") != "" {
		page = q.Get("offset")
	}

	convertedVal, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return convertedVal
}
