// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryManagerGetReportPreview() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetReportPreviewRead,

		Schema: map[string]*schema.Schema{
			"report_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of the report.",
			},
			"component_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of components ids to be evaluated for the given report. If not specified, then all the components are evaluated.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filters": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies list of global filters that are applicable to given components in the report.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the attribute.",
						},
						"filter_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the type of the filter that needs to be applied.",
						},
						"in_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the in filter that are applied on attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute_data_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the data type of the attribute.",
									},
									"attribute_labels": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the optional label values for the attribute.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"bool_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of boolean values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeBool,
										},
									},
									"int32_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of int32 values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"int64_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of int64 values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"string_filter_values": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of string values to filter results on.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"range_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the filters that are applied on attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lower_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the lower bound value. If specified, all the results which are greater than this value will be returned.",
									},
									"upper_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the upper bound value. If specified, all the results which are lesser than this value will be returned.",
									},
								},
							},
						},
						"systems_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the systems filter. Specifying this will pre filter all the results provided list of system identifier before applying aggregations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"system_ids": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies an array of system identifiers. System identifiers may be of format clusterid:clusterincarnationid or a regionid (applicable only in case of DMaaS).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"system_names": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the optional system names labels.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"time_range_filter_params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the time range filter. Specifying this will pre filter all the results on necessary resources like Protection Runs etc before applying aggregations. Currently, maximum allowed time range is 60 days.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"date_range": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Enum value for specifying the date range for a time filter. Considered only if lowerBound and upperBound are empty.",
									},
									"duration_hours": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the duration preceding the current time for which the data must be fetch i.e fetch data between currentTime and currentTime - durationHours. This filter is only considered if neither upperBound, lowerBound or dateRange is specified.",
									},
									"lower_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the lower bound value. If specified, all the results which are greater than this value will be returned.",
									},
									"upper_bound": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the upper bound value. If specified, all the results which are lesser than this value will be returned.",
									},
								},
							},
						},
					},
				},
			},
			"timezone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies timezone of the user. If nil, defaults to UTC. The time specified should be a location name in the IANA Time Zone database, for example, 'America/Los_Angeles'.",
			},
			"components": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the component params and data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggs": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the aggregation related information that needs to be applied on the attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aggregated_attributes": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of aggregation properties to be applied on the attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"aggregation_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the aggregation type.",
												},
												"attribute": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the attribute name.",
												},
												"label": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the label to be generated for this aggregated attribute. If not specified, then by default label of the column in the output will be combination of aggregation type and attribute.",
												},
											},
										},
									},
									"grouped_attributes": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies list of attributes to be grouped in the aggregation.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"config": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the configuration parameters to customize and format the columns in the report artifacts like excel, pdf etc.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"xlsx_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the configuration parameters to customize a component in excel report.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_config": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies customized configuration for the attributes in the report. If not specified, all the attributes will be sent as-is to the report without any formatting.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"attribute_name": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies the name of the attribute.",
															},
															"custom_label": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a custom label for attribute to appear in the xlsx report. If not specified, default attribute name will be used.",
															},
															"format": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Specifies a custom format for attribute to appear in the xlsx report. If not specified, the attribute value is sent as-is.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"data": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the data returned after evaluating the component.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies description of the Component.",
						},
						"filters": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the filters that are applied on specific report type attributes in order to compose this component.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the attribute.",
									},
									"filter_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the type of the filter that needs to be applied.",
									},
									"in_filter_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the in filter that are applied on attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attribute_data_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies the data type of the attribute.",
												},
												"attribute_labels": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the optional label values for the attribute.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"bool_filter_values": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies list of boolean values to filter results on.",
													Elem: &schema.Schema{
														Type: schema.TypeBool,
													},
												},
												"int32_filter_values": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies list of int32 values to filter results on.",
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"int64_filter_values": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies list of int64 values to filter results on.",
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"string_filter_values": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies list of string values to filter results on.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"range_filter_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the filters that are applied on attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"lower_bound": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the lower bound value. If specified, all the results which are greater than this value will be returned.",
												},
												"upper_bound": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the upper bound value. If specified, all the results which are lesser than this value will be returned.",
												},
											},
										},
									},
									"systems_filter_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the systems filter. Specifying this will pre filter all the results provided list of system identifier before applying aggregations.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"system_ids": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies an array of system identifiers. System identifiers may be of format clusterid:clusterincarnationid or a regionid (applicable only in case of DMaaS).",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"system_names": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Specifies the optional system names labels.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"time_range_filter_params": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the time range filter. Specifying this will pre filter all the results on necessary resources like Protection Runs etc before applying aggregations. Currently, maximum allowed time range is 60 days.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"date_range": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Enum value for specifying the date range for a time filter. Considered only if lowerBound and upperBound are empty.",
												},
												"duration_hours": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the duration preceding the current time for which the data must be fetch i.e fetch data between currentTime and currentTime - durationHours. This filter is only considered if neither upperBound, lowerBound or dateRange is specified.",
												},
												"lower_bound": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the lower bound value. If specified, all the results which are greater than this value will be returned.",
												},
												"upper_bound": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specifies the upper bound value. If specified, all the results which are lesser than this value will be returned.",
												},
											},
										},
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the id of the Component.",
						},
						"limit": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the parameters to limit the resulting dataset.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"from": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the offset to which resulting data will be skipped before applying the size parameter. For example if dataset size is 10 objects, from=2 and size=5, then from 10 objects only 5 objects are returned starting from offset 2 i.e., 2 to 7. If not specified, then none of the objects are skipped.",
									},
									"size": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the number of objects to be returned from the offset specified.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the Component.",
						},
						"report_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the report type on top of which this Component is created from.",
						},
						"sort": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the sorting (ordering) parameters to be applied to the resulting data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the name of the attribute.",
									},
									"desc": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Specifies whether the sorting order should be descending. Default value is false.",
									},
								},
							},
						},
					},
				},
			},
			"generated_timestamp_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the epoch timestamp in UTC in microseconds.",
			},
			"last_refresh_timestamp_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the last refresh timestamp of data used to evaluate the component. If this parameter is not returned, then 'generatedTimestampUsecs' can be used for last refreshed timestamp of the data. This is epoch timestamp in UTC in microseconds.",
			},
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the title of the report.",
			},
		},
	}
}

func dataSourceIbmBackupRecoveryManagerGetReportPreviewRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_report_preview", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getReportPreviewOptions := &backuprecoveryv1.GetReportPreviewOptions{}

	getReportPreviewOptions.SetID(d.Get("report_id").(string))
	if _, ok := d.GetOk("component_ids"); ok {
		var componentIds []string
		for _, v := range d.Get("component_ids").([]interface{}) {
			componentIdsItem := v.(string)
			componentIds = append(componentIds, componentIdsItem)
		}
		getReportPreviewOptions.SetComponentIds(componentIds)
	}
	if _, ok := d.GetOk("filters"); ok {
		var filters []backuprecoveryv1.AttributeFilter
		for _, v := range d.Get("filters").([]interface{}) {
			value := v.(map[string]interface{})
			filtersItem, err := ResourceIbmBackupRecoveryManagerGetReportPreviewMapToAttributeFilter(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_get_report_preview", "read", "parse-filters").GetDiag()
			}
			filters = append(filters, *filtersItem)
		}
		getReportPreviewOptions.SetFilters(filters)
	}
	if _, ok := d.GetOk("timezone"); ok {
		getReportPreviewOptions.SetTimezone(d.Get("timezone").(string))
	}

	reportPreview, _, err := backupRecoveryClient.GetReportPreviewWithContext(context, getReportPreviewOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetReportPreviewWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_report_preview", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*reportPreview.ID)

	if !core.IsNil(reportPreview.Components) {
		components := []map[string]interface{}{}
		for _, componentsItem := range reportPreview.Components {
			componentsItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewComponentToMap(&componentsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_report_preview", "read", "components-to-map").GetDiag()
			}
			components = append(components, componentsItemMap)
		}
		if err = d.Set("components", components); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting components: %s", err), "(Data) ibm_backup_recovery_manager_get_report_preview", "read", "set-components").GetDiag()
		}
	}

	if !core.IsNil(reportPreview.GeneratedTimestampUsecs) {
		if err = d.Set("generated_timestamp_usecs", flex.IntValue(reportPreview.GeneratedTimestampUsecs)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting generated_timestamp_usecs: %s", err), "(Data) ibm_backup_recovery_manager_get_report_preview", "read", "set-generated_timestamp_usecs").GetDiag()
		}
	}

	if !core.IsNil(reportPreview.LastRefreshTimestampUsecs) {
		if err = d.Set("last_refresh_timestamp_usecs", flex.IntValue(reportPreview.LastRefreshTimestampUsecs)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_refresh_timestamp_usecs: %s", err), "(Data) ibm_backup_recovery_manager_get_report_preview", "read", "set-last_refresh_timestamp_usecs").GetDiag()
		}
	}

	if !core.IsNil(reportPreview.Title) {
		if err = d.Set("title", reportPreview.Title); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting title: %s", err), "(Data) ibm_backup_recovery_manager_get_report_preview", "read", "set-title").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewComponentToMap(model *backuprecoveryv1.Component) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Aggs != nil {
		aggsMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeAggregationsToMap(model.Aggs)
		if err != nil {
			return modelMap, err
		}
		modelMap["aggs"] = []map[string]interface{}{aggsMap}
	}
	if model.Config != nil {
		configMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewCustomConfigParamsToMap(model.Config)
		if err != nil {
			return modelMap, err
		}
		modelMap["config"] = []map[string]interface{}{configMap}
	}
	if model.Data != nil {
		modelMap["data"] = model.Data
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.Filters != nil {
		filters := []map[string]interface{}{}
		for _, filtersItem := range model.Filters {
			filtersItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeFilterToMap(&filtersItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			filters = append(filters, filtersItemMap)
		}
		modelMap["filters"] = filters
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Limit != nil {
		limitMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewLimitParamsToMap(model.Limit)
		if err != nil {
			return modelMap, err
		}
		modelMap["limit"] = []map[string]interface{}{limitMap}
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ReportType != nil {
		modelMap["report_type"] = *model.ReportType
	}
	if model.Sort != nil {
		sort := []map[string]interface{}{}
		for _, sortItem := range model.Sort {
			sortItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeSortToMap(&sortItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			sort = append(sort, sortItemMap)
		}
		modelMap["sort"] = sort
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeAggregationsToMap(model *backuprecoveryv1.AttributeAggregations) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AggregatedAttributes != nil {
		aggregatedAttributes := []map[string]interface{}{}
		for _, aggregatedAttributesItem := range model.AggregatedAttributes {
			aggregatedAttributesItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewAggregatedAttributesParamsToMap(&aggregatedAttributesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			aggregatedAttributes = append(aggregatedAttributes, aggregatedAttributesItemMap)
		}
		modelMap["aggregated_attributes"] = aggregatedAttributes
	}
	if model.GroupedAttributes != nil {
		modelMap["grouped_attributes"] = model.GroupedAttributes
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewAggregatedAttributesParamsToMap(model *backuprecoveryv1.AggregatedAttributesParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["aggregation_type"] = *model.AggregationType
	modelMap["attribute"] = *model.Attribute
	if model.Label != nil {
		modelMap["label"] = *model.Label
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewCustomConfigParamsToMap(model *backuprecoveryv1.CustomConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.XlsxParams != nil {
		xlsxParamsMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxCustomConfigParamsToMap(model.XlsxParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["xlsx_params"] = []map[string]interface{}{xlsxParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxCustomConfigParamsToMap(model *backuprecoveryv1.XlsxCustomConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AttributeConfig != nil {
		attributeConfig := []map[string]interface{}{}
		for _, attributeConfigItem := range model.AttributeConfig {
			attributeConfigItemMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxAttributeCustomConfigParamsToMap(&attributeConfigItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			attributeConfig = append(attributeConfig, attributeConfigItemMap)
		}
		modelMap["attribute_config"] = attributeConfig
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewXlsxAttributeCustomConfigParamsToMap(model *backuprecoveryv1.XlsxAttributeCustomConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute_name"] = *model.AttributeName
	if model.CustomLabel != nil {
		modelMap["custom_label"] = *model.CustomLabel
	}
	if model.Format != nil {
		modelMap["format"] = *model.Format
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeFilterToMap(model *backuprecoveryv1.AttributeFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute"] = *model.Attribute
	modelMap["filter_type"] = *model.FilterType
	if model.InFilterParams != nil {
		inFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewInFilterParamsToMap(model.InFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["in_filter_params"] = []map[string]interface{}{inFilterParamsMap}
	}
	if model.RangeFilterParams != nil {
		rangeFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewRangeFilterParamsToMap(model.RangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["range_filter_params"] = []map[string]interface{}{rangeFilterParamsMap}
	}
	if model.SystemsFilterParams != nil {
		systemsFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewSystemsFilterParamsToMap(model.SystemsFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["systems_filter_params"] = []map[string]interface{}{systemsFilterParamsMap}
	}
	if model.TimeRangeFilterParams != nil {
		timeRangeFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetReportPreviewTimeRangeFilterParamsToMap(model.TimeRangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["time_range_filter_params"] = []map[string]interface{}{timeRangeFilterParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewInFilterParamsToMap(model *backuprecoveryv1.InFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute_data_type"] = *model.AttributeDataType
	if model.AttributeLabels != nil {
		modelMap["attribute_labels"] = model.AttributeLabels
	}
	if model.BoolFilterValues != nil {
		modelMap["bool_filter_values"] = model.BoolFilterValues
	}
	if model.Int32FilterValues != nil {
		modelMap["int32_filter_values"] = model.Int32FilterValues
	}
	if model.Int64FilterValues != nil {
		modelMap["int64_filter_values"] = model.Int64FilterValues
	}
	if model.StringFilterValues != nil {
		modelMap["string_filter_values"] = model.StringFilterValues
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewRangeFilterParamsToMap(model *backuprecoveryv1.RangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewSystemsFilterParamsToMap(model *backuprecoveryv1.SystemsFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["system_ids"] = model.SystemIds
	if model.SystemNames != nil {
		modelMap["system_names"] = model.SystemNames
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewTimeRangeFilterParamsToMap(model *backuprecoveryv1.TimeRangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DateRange != nil {
		modelMap["date_range"] = *model.DateRange
	}
	if model.DurationHours != nil {
		modelMap["duration_hours"] = flex.IntValue(model.DurationHours)
	}
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewLimitParamsToMap(model *backuprecoveryv1.LimitParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.From != nil {
		modelMap["from"] = flex.IntValue(model.From)
	}
	modelMap["size"] = flex.IntValue(model.Size)
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetReportPreviewAttributeSortToMap(model *backuprecoveryv1.AttributeSort) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute"] = *model.Attribute
	if model.Desc != nil {
		modelMap["desc"] = *model.Desc
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewMapToAttributeFilter(modelMap map[string]interface{}) (*backuprecoveryv1.AttributeFilter, error) {
	model := &backuprecoveryv1.AttributeFilter{}
	model.Attribute = core.StringPtr(modelMap["attribute"].(string))
	model.FilterType = core.StringPtr(modelMap["filter_type"].(string))
	if modelMap["in_filter_params"] != nil && len(modelMap["in_filter_params"].([]interface{})) > 0 {
		InFilterParamsModel, err := ResourceIbmBackupRecoveryManagerGetReportPreviewMapToInFilterParams(modelMap["in_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.InFilterParams = InFilterParamsModel
	}
	if modelMap["range_filter_params"] != nil && len(modelMap["range_filter_params"].([]interface{})) > 0 {
		RangeFilterParamsModel, err := ResourceIbmBackupRecoveryManagerGetReportPreviewMapToRangeFilterParams(modelMap["range_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.RangeFilterParams = RangeFilterParamsModel
	}
	if modelMap["systems_filter_params"] != nil && len(modelMap["systems_filter_params"].([]interface{})) > 0 {
		SystemsFilterParamsModel, err := ResourceIbmBackupRecoveryManagerGetReportPreviewMapToSystemsFilterParams(modelMap["systems_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.SystemsFilterParams = SystemsFilterParamsModel
	}
	if modelMap["time_range_filter_params"] != nil && len(modelMap["time_range_filter_params"].([]interface{})) > 0 {
		TimeRangeFilterParamsModel, err := ResourceIbmBackupRecoveryManagerGetReportPreviewMapToTimeRangeFilterParams(modelMap["time_range_filter_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TimeRangeFilterParams = TimeRangeFilterParamsModel
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewMapToInFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.InFilterParams, error) {
	model := &backuprecoveryv1.InFilterParams{}
	model.AttributeDataType = core.StringPtr(modelMap["attribute_data_type"].(string))
	if modelMap["attribute_labels"] != nil {
		attributeLabels := []string{}
		for _, attributeLabelsItem := range modelMap["attribute_labels"].([]interface{}) {
			attributeLabels = append(attributeLabels, attributeLabelsItem.(string))
		}
		model.AttributeLabels = attributeLabels
	}
	if modelMap["bool_filter_values"] != nil {
		boolFilterValues := []bool{}
		for _, boolFilterValuesItem := range modelMap["bool_filter_values"].([]interface{}) {
			boolFilterValues = append(boolFilterValues, boolFilterValuesItem.(bool))
		}
		model.BoolFilterValues = boolFilterValues
	}
	if modelMap["int32_filter_values"] != nil {
		int32FilterValues := []int64{}
		for _, int32FilterValuesItem := range modelMap["int32_filter_values"].([]interface{}) {
			int32FilterValues = append(int32FilterValues, int64(int32FilterValuesItem.(int)))
		}
		model.Int32FilterValues = int32FilterValues
	}
	if modelMap["int64_filter_values"] != nil {
		int64FilterValues := []int64{}
		for _, int64FilterValuesItem := range modelMap["int64_filter_values"].([]interface{}) {
			int64FilterValues = append(int64FilterValues, int64(int64FilterValuesItem.(int)))
		}
		model.Int64FilterValues = int64FilterValues
	}
	if modelMap["string_filter_values"] != nil {
		stringFilterValues := []string{}
		for _, stringFilterValuesItem := range modelMap["string_filter_values"].([]interface{}) {
			stringFilterValues = append(stringFilterValues, stringFilterValuesItem.(string))
		}
		model.StringFilterValues = stringFilterValues
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewMapToRangeFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.RangeFilterParams, error) {
	model := &backuprecoveryv1.RangeFilterParams{}
	if modelMap["lower_bound"] != nil {
		model.LowerBound = core.Int64Ptr(int64(modelMap["lower_bound"].(int)))
	}
	if modelMap["upper_bound"] != nil {
		model.UpperBound = core.Int64Ptr(int64(modelMap["upper_bound"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewMapToSystemsFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.SystemsFilterParams, error) {
	model := &backuprecoveryv1.SystemsFilterParams{}
	systemIds := []string{}
	for _, systemIdsItem := range modelMap["system_ids"].([]interface{}) {
		systemIds = append(systemIds, systemIdsItem.(string))
	}
	model.SystemIds = systemIds
	if modelMap["system_names"] != nil {
		systemNames := []string{}
		for _, systemNamesItem := range modelMap["system_names"].([]interface{}) {
			systemNames = append(systemNames, systemNamesItem.(string))
		}
		model.SystemNames = systemNames
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewMapToTimeRangeFilterParams(modelMap map[string]interface{}) (*backuprecoveryv1.TimeRangeFilterParams, error) {
	model := &backuprecoveryv1.TimeRangeFilterParams{}
	if modelMap["date_range"] != nil && modelMap["date_range"].(string) != "" {
		model.DateRange = core.StringPtr(modelMap["date_range"].(string))
	}
	if modelMap["duration_hours"] != nil {
		model.DurationHours = core.Int64Ptr(int64(modelMap["duration_hours"].(int)))
	}
	if modelMap["lower_bound"] != nil {
		model.LowerBound = core.Int64Ptr(int64(modelMap["lower_bound"].(int)))
	}
	if modelMap["upper_bound"] != nil {
		model.UpperBound = core.Int64Ptr(int64(modelMap["upper_bound"].(int)))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewAttributeFilterToMap(model *backuprecoveryv1.AttributeFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute"] = *model.Attribute
	modelMap["filter_type"] = *model.FilterType
	if model.InFilterParams != nil {
		inFilterParamsMap, err := ResourceIbmBackupRecoveryManagerGetReportPreviewInFilterParamsToMap(model.InFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["in_filter_params"] = []map[string]interface{}{inFilterParamsMap}
	}
	if model.RangeFilterParams != nil {
		rangeFilterParamsMap, err := ResourceIbmBackupRecoveryManagerGetReportPreviewRangeFilterParamsToMap(model.RangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["range_filter_params"] = []map[string]interface{}{rangeFilterParamsMap}
	}
	if model.SystemsFilterParams != nil {
		systemsFilterParamsMap, err := ResourceIbmBackupRecoveryManagerGetReportPreviewSystemsFilterParamsToMap(model.SystemsFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["systems_filter_params"] = []map[string]interface{}{systemsFilterParamsMap}
	}
	if model.TimeRangeFilterParams != nil {
		timeRangeFilterParamsMap, err := ResourceIbmBackupRecoveryManagerGetReportPreviewTimeRangeFilterParamsToMap(model.TimeRangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["time_range_filter_params"] = []map[string]interface{}{timeRangeFilterParamsMap}
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewInFilterParamsToMap(model *backuprecoveryv1.InFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute_data_type"] = *model.AttributeDataType
	if model.AttributeLabels != nil {
		modelMap["attribute_labels"] = model.AttributeLabels
	}
	if model.BoolFilterValues != nil {
		modelMap["bool_filter_values"] = model.BoolFilterValues
	}
	if model.Int32FilterValues != nil {
		modelMap["int32_filter_values"] = model.Int32FilterValues
	}
	if model.Int64FilterValues != nil {
		modelMap["int64_filter_values"] = model.Int64FilterValues
	}
	if model.StringFilterValues != nil {
		modelMap["string_filter_values"] = model.StringFilterValues
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewRangeFilterParamsToMap(model *backuprecoveryv1.RangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewSystemsFilterParamsToMap(model *backuprecoveryv1.SystemsFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["system_ids"] = model.SystemIds
	if model.SystemNames != nil {
		modelMap["system_names"] = model.SystemNames
	}
	return modelMap, nil
}

func ResourceIbmBackupRecoveryManagerGetReportPreviewTimeRangeFilterParamsToMap(model *backuprecoveryv1.TimeRangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DateRange != nil {
		modelMap["date_range"] = *model.DateRange
	}
	if model.DurationHours != nil {
		modelMap["duration_hours"] = flex.IntValue(model.DurationHours)
	}
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}
