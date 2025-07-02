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

func DataSourceIbmBackupRecoveryManagerGetComponent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryManagerGetComponentRead,

		Schema: map[string]*schema.Schema{
			"component_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the id of the report component.",
			},
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
	}
}

func dataSourceIbmBackupRecoveryManagerGetComponentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	heliosReportingServiceApIsClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_component", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getComponentByIdOptions := &backuprecoveryv1.GetComponentByIdOptions{}

	getComponentByIdOptions.SetID(d.Get("component_id").(string))

	component, _, err := heliosReportingServiceApIsClient.GetComponentByIDWithContext(context, getComponentByIdOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetComponentByIDWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_manager_get_component", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*component.ID)

	if !core.IsNil(component.Aggs) {
		aggs := []map[string]interface{}{}
		aggsMap, err := DataSourceIbmBackupRecoveryManagerGetComponentAttributeAggregationsToMap(component.Aggs)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_component", "read", "aggs-to-map").GetDiag()
		}
		aggs = append(aggs, aggsMap)
		if err = d.Set("aggs", aggs); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting aggs: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-aggs").GetDiag()
		}
	}

	if !core.IsNil(component.Config) {
		config := []map[string]interface{}{}
		configMap, err := DataSourceIbmBackupRecoveryManagerGetComponentCustomConfigParamsToMap(component.Config)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_component", "read", "config-to-map").GetDiag()
		}
		config = append(config, configMap)
		if err = d.Set("config", config); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting config: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-config").GetDiag()
		}
	}

	if !core.IsNil(component.Data) {
		data := []map[string]interface{}{}
		for _, dataItem := range component.Data {
			data = append(data, dataItem)
		}
		if err = d.Set("data", data); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting data: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-data").GetDiag()
		}
	}

	if !core.IsNil(component.Description) {
		if err = d.Set("description", component.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(component.Filters) {
		filters := []map[string]interface{}{}
		for _, filtersItem := range component.Filters {
			filtersItemMap, err := DataSourceIbmBackupRecoveryManagerGetComponentAttributeFilterToMap(&filtersItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_component", "read", "filters-to-map").GetDiag()
			}
			filters = append(filters, filtersItemMap)
		}
		if err = d.Set("filters", filters); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting filters: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-filters").GetDiag()
		}
	}

	if !core.IsNil(component.Limit) {
		limit := []map[string]interface{}{}
		limitMap, err := DataSourceIbmBackupRecoveryManagerGetComponentLimitParamsToMap(component.Limit)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_component", "read", "limit-to-map").GetDiag()
		}
		limit = append(limit, limitMap)
		if err = d.Set("limit", limit); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting limit: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-limit").GetDiag()
		}
	}

	if !core.IsNil(component.Name) {
		if err = d.Set("name", component.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-name").GetDiag()
		}
	}

	if !core.IsNil(component.ReportType) {
		if err = d.Set("report_type", component.ReportType); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting report_type: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-report_type").GetDiag()
		}
	}

	if !core.IsNil(component.Sort) {
		sort := []map[string]interface{}{}
		for _, sortItem := range component.Sort {
			sortItemMap, err := DataSourceIbmBackupRecoveryManagerGetComponentAttributeSortToMap(&sortItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_manager_get_component", "read", "sort-to-map").GetDiag()
			}
			sort = append(sort, sortItemMap)
		}
		if err = d.Set("sort", sort); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting sort: %s", err), "(Data) ibm_backup_recovery_manager_get_component", "read", "set-sort").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentAttributeAggregationsToMap(model *backuprecoveryv1.AttributeAggregations) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AggregatedAttributes != nil {
		aggregatedAttributes := []map[string]interface{}{}
		for _, aggregatedAttributesItem := range model.AggregatedAttributes {
			aggregatedAttributesItemMap, err := DataSourceIbmBackupRecoveryManagerGetComponentAggregatedAttributesParamsToMap(&aggregatedAttributesItem) // #nosec G601
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

func DataSourceIbmBackupRecoveryManagerGetComponentAggregatedAttributesParamsToMap(model *backuprecoveryv1.AggregatedAttributesParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["aggregation_type"] = *model.AggregationType
	modelMap["attribute"] = *model.Attribute
	if model.Label != nil {
		modelMap["label"] = *model.Label
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentCustomConfigParamsToMap(model *backuprecoveryv1.CustomConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.XlsxParams != nil {
		xlsxParamsMap, err := DataSourceIbmBackupRecoveryManagerGetComponentXlsxCustomConfigParamsToMap(model.XlsxParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["xlsx_params"] = []map[string]interface{}{xlsxParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentXlsxCustomConfigParamsToMap(model *backuprecoveryv1.XlsxCustomConfigParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AttributeConfig != nil {
		attributeConfig := []map[string]interface{}{}
		for _, attributeConfigItem := range model.AttributeConfig {
			attributeConfigItemMap, err := DataSourceIbmBackupRecoveryManagerGetComponentXlsxAttributeCustomConfigParamsToMap(&attributeConfigItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			attributeConfig = append(attributeConfig, attributeConfigItemMap)
		}
		modelMap["attribute_config"] = attributeConfig
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentXlsxAttributeCustomConfigParamsToMap(model *backuprecoveryv1.XlsxAttributeCustomConfigParams) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryManagerGetComponentAttributeFilterToMap(model *backuprecoveryv1.AttributeFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute"] = *model.Attribute
	modelMap["filter_type"] = *model.FilterType
	if model.InFilterParams != nil {
		inFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetComponentInFilterParamsToMap(model.InFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["in_filter_params"] = []map[string]interface{}{inFilterParamsMap}
	}
	if model.RangeFilterParams != nil {
		rangeFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetComponentRangeFilterParamsToMap(model.RangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["range_filter_params"] = []map[string]interface{}{rangeFilterParamsMap}
	}
	if model.SystemsFilterParams != nil {
		systemsFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetComponentSystemsFilterParamsToMap(model.SystemsFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["systems_filter_params"] = []map[string]interface{}{systemsFilterParamsMap}
	}
	if model.TimeRangeFilterParams != nil {
		timeRangeFilterParamsMap, err := DataSourceIbmBackupRecoveryManagerGetComponentTimeRangeFilterParamsToMap(model.TimeRangeFilterParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["time_range_filter_params"] = []map[string]interface{}{timeRangeFilterParamsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentInFilterParamsToMap(model *backuprecoveryv1.InFilterParams) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryManagerGetComponentRangeFilterParamsToMap(model *backuprecoveryv1.RangeFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LowerBound != nil {
		modelMap["lower_bound"] = flex.IntValue(model.LowerBound)
	}
	if model.UpperBound != nil {
		modelMap["upper_bound"] = flex.IntValue(model.UpperBound)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentSystemsFilterParamsToMap(model *backuprecoveryv1.SystemsFilterParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["system_ids"] = model.SystemIds
	if model.SystemNames != nil {
		modelMap["system_names"] = model.SystemNames
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentTimeRangeFilterParamsToMap(model *backuprecoveryv1.TimeRangeFilterParams) (map[string]interface{}, error) {
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

func DataSourceIbmBackupRecoveryManagerGetComponentLimitParamsToMap(model *backuprecoveryv1.LimitParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.From != nil {
		modelMap["from"] = flex.IntValue(model.From)
	}
	modelMap["size"] = flex.IntValue(model.Size)
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryManagerGetComponentAttributeSortToMap(model *backuprecoveryv1.AttributeSort) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["attribute"] = *model.Attribute
	if model.Desc != nil {
		modelMap["desc"] = *model.Desc
	}
	return modelMap, nil
}
