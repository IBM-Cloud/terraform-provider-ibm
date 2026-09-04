// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func DataSourceIBMAppConfigWorkflowConfigs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigWorkflowConfigsRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"sort": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sort the workflow config details based on the specified attribute (name, created_time, updated_time, workflow_id).",
			},
			"search": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Searches for the provided keyword in the Name of the workflow config.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to retrieve. Default is 10.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to skip.",
			},
			"workflow_configs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of workflow configuration objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the workflow configuration.",
						},
						"workflow_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for the workflow configuration.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the workflow is enabled.",
						},
						// Named "workflow_provider" instead of "provider" because "provider" is a
						// reserved meta-argument in Terraform's plugin SDK and cannot be used as a
						// schema field name. It maps to the "provider" field in the IBM App
						// Configuration API response.
						"workflow_provider": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Provider configuration for the workflow. Corresponds to the 'provider' field in the API.",
							Elem:        workflowProviderSchema(),
						},
						"scope": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Scope configuration for the workflow.",
							Elem:        workflowScopeSchema(),
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the workflow was created.",
						},
						"updated_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the workflow was last updated.",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL to access this workflow configuration.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of records.",
			},
			"first": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the first page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"next": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the next list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
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
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAppConfigWorkflowConfigsRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.ListWorkflowConfigsOptions{}
	if _, ok := GetFieldExists(d, "sort"); ok {
		options.SetSort(d.Get("sort").(string))
	}
	if _, ok := GetFieldExists(d, "search"); ok {
		options.SetSearch(d.Get("search").(string))
	}

	var configsList *appconfigurationv1.WorkflowConfigsList
	var offset int64
	var limit int64 = 10
	var isLimit bool
	finalList := []appconfigurationv1.WorkflowConfigResponse{}

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
		result, response, err := appconfigClient.ListWorkflowConfigs(options)
		configsList = result
		if err != nil {
			return flex.FmtErrorf("[ERROR] ListWorkflowConfigs failed %s\n%s", err, response)
		}
		if isLimit {
			offset = 0
		} else {
			offset = dataSourceWorkflowConfigsGetNext(result.Next)
		}
		finalList = append(finalList, result.WorkflowConfigs...)
		if offset == 0 {
			break
		}
	}

	configsList.WorkflowConfigs = finalList

	d.SetId(fmt.Sprintf("%s/workflow_configs", guid))

	if configsList.TotalCount != nil {
		if err = d.Set("total_count", configsList.TotalCount); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting total_count: %s", err)
		}
	}
	if configsList.Limit != nil {
		if err = d.Set("limit", configsList.Limit); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting limit: %s", err)
		}
	}
	if configsList.Offset != nil {
		if err = d.Set("offset", configsList.Offset); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting offset: %s", err)
		}
	}
	if configsList.First != nil {
		if err = d.Set("first", dataSourceFeatureListFlattenPagination(*configsList.First)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting first: %s", err)
		}
	}
	if configsList.Previous != nil {
		if err = d.Set("previous", dataSourceFeatureListFlattenPagination(*configsList.Previous)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting previous: %s", err)
		}
	}
	if configsList.Next != nil {
		if err = d.Set("next", dataSourceFeatureListFlattenPagination(*configsList.Next)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting next: %s", err)
		}
	}
	if configsList.Last != nil {
		if err = d.Set("last", dataSourceFeatureListFlattenPagination(*configsList.Last)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting last: %s", err)
		}
	}

	configs := []map[string]interface{}{}
	for _, item := range configsList.WorkflowConfigs {
		configs = append(configs, dataSourceWorkflowConfigToMap(item))
	}
	if err = d.Set("workflow_configs", configs); err != nil {
		return flex.FmtErrorf("[ERROR] Error setting workflow_configs: %s", err)
	}

	return nil
}

func dataSourceWorkflowConfigToMap(item appconfigurationv1.WorkflowConfigResponse) map[string]interface{} {
	m := map[string]interface{}{}
	if item.Name != nil {
		m["name"] = item.Name
	}
	if item.WorkflowID != nil {
		m["workflow_id"] = item.WorkflowID
	}
	if item.Enabled != nil {
		m["enabled"] = item.Enabled
	}
	if item.Provider != nil {
		m["workflow_provider"] = flattenWorkflowProvider(item.Provider)
	}
	if item.Scope != nil {
		m["scope"] = flattenWorkflowScope(item.Scope)
	}
	if item.CreatedTime != nil {
		m["created_time"] = item.CreatedTime.String()
	}
	if item.UpdatedTime != nil {
		m["updated_time"] = item.UpdatedTime.String()
	}
	if item.Href != nil {
		m["href"] = item.Href
	}
	return m
}

func dataSourceWorkflowConfigsGetNext(next *appconfigurationv1.PaginatedListNext) int64 {
	if next == nil || next.Href == nil {
		return 0
	}
	u, err := url.Parse(*next.Href)
	if err != nil {
		return 0
	}
	q := u.Query()
	var page string
	if q.Get("offset") != "" {
		page = q.Get("offset")
	} else if q.Get("start") != "" {
		page = q.Get("start")
	}
	val, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

// Suppress unused import for reflect used in the list pagination helper pattern.
var _ = reflect.DeepEqual
