// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
)

func DataSourceIbmLogsE2ms() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsE2msRead,

		Schema: map[string]*schema.Schema{
			"events2metrics": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of event to metrics definitions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "E2M id, required on update requests.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "E2M name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "E2m description.",
						},
						"create_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "E2M create time.",
						},
						"update_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "E2M update time.",
						},
						"permutations": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "represents E2M permutations limit.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"limit": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "e2m permutation limit.",
									},
									"has_exceeded_limit": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "flag to indicate if limit was exceeded.",
									},
								},
							},
						},
						"metric_labels": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "E2M metric labels.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_label": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "metric label target label.",
									},
									"source_field": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "metric label source field.",
									},
								},
							},
						},
						"metric_fields": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "E2M metric fields.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_base_metric_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "target metric field.",
									},
									"source_field": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "source field.",
									},
									"aggregations": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "represents Aggregation type list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "is enabled.",
												},
												"agg_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "aggregation type.",
												},
												"target_metric_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "target metric field.",
												},
												"samples": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "e2m sample type metadata.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sample_type": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "sample type min/max.",
															},
														},
													},
												},
												"histogram": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "e2m aggregate histogram type metadata.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"buckets": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "buckets that describe the e2m.",
																Elem: &schema.Schema{
																	Type: schema.TypeFloat,
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
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "e2m type.",
						},
						"is_internal": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "a flag that represents if the e2m is for internal usage.",
						},
						"spans_query": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "spans query.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lucene": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "lucene query.",
									},
									"applicationname_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "application name filters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"subsystemname_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "subsystem name filters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"action_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "action filters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"service_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "service filters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"logs_query": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "logs query.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"lucene": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "lucene query.",
									},
									"alias": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "alias.",
									},
									"applicationname_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "application name filters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"subsystemname_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "subsystem names filters.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severity_filters": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "severity type filters.",
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

func dataSourceIbmLogsE2msRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_e2ms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient = getClientWithLogsInstanceEndpoint(logsClient, instanceId, region, getLogsInstanceEndpointType(logsClient, d))

	listE2mOptions := &logsv0.ListE2mOptions{}

	event2MetricCollection, _, err := logsClient.ListE2mWithContext(context, listE2mOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListE2mWithContext failed: %s", err.Error()), "(Data) ibm_logs_e2ms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmLogsE2msID(d))

	events2metrics := []map[string]interface{}{}
	if event2MetricCollection.Events2metrics != nil {
		for _, modelItem := range event2MetricCollection.Events2metrics {
			modelMap, err := DataSourceIbmLogsE2msEvent2MetricToMap(modelItem)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_e2ms", "read")
				return tfErr.GetDiag()
			}
			events2metrics = append(events2metrics, modelMap)
		}
	}
	if err = d.Set("events2metrics", events2metrics); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting events2metrics: %s", err), "(Data) ibm_logs_e2ms", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIbmLogsE2msID returns a reasonable ID for the list.
func dataSourceIbmLogsE2msID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmLogsE2msEvent2MetricToMap(model logsv0.Event2MetricIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.Event2MetricApisEvents2metricsV2E2mQuerySpansQuery); ok {
		return DataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQuerySpansQueryToMap(model.(*logsv0.Event2MetricApisEvents2metricsV2E2mQuerySpansQuery))
	} else if _, ok := model.(*logsv0.Event2MetricApisEvents2metricsV2E2mQueryLogsQuery); ok {
		return DataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQueryLogsQueryToMap(model.(*logsv0.Event2MetricApisEvents2metricsV2E2mQueryLogsQuery))
	} else if _, ok := model.(*logsv0.Event2Metric); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.Event2Metric)
		if model.ID != nil {
			modelMap["id"] = model.ID.String()
		}
		modelMap["name"] = *model.Name
		if model.Description != nil {
			modelMap["description"] = *model.Description
		}
		if model.CreateTime != nil {
			modelMap["create_time"] = *model.CreateTime
		}
		if model.UpdateTime != nil {
			modelMap["update_time"] = *model.UpdateTime
		}
		if model.Permutations != nil {
			permutationsMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mPermutationsToMap(model.Permutations)
			if err != nil {
				return modelMap, err
			}
			modelMap["permutations"] = []map[string]interface{}{permutationsMap}
		}
		if model.MetricLabels != nil {
			metricLabels := []map[string]interface{}{}
			for _, metricLabelsItem := range model.MetricLabels {
				metricLabelsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2MetricLabelToMap(&metricLabelsItem)
				if err != nil {
					return modelMap, err
				}
				metricLabels = append(metricLabels, metricLabelsItemMap)
			}
			modelMap["metric_labels"] = metricLabels
		}
		if model.MetricFields != nil {
			metricFields := []map[string]interface{}{}
			for _, metricFieldsItem := range model.MetricFields {
				metricFieldsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2MetricFieldToMap(&metricFieldsItem)
				if err != nil {
					return modelMap, err
				}
				metricFields = append(metricFields, metricFieldsItemMap)
			}
			modelMap["metric_fields"] = metricFields
		}
		modelMap["type"] = *model.Type
		if model.IsInternal != nil {
			modelMap["is_internal"] = *model.IsInternal
		}
		if model.SpansQuery != nil {
			spansQueryMap, err := DataSourceIbmLogsE2msApisSpans2metricsV2SpansQueryToMap(model.SpansQuery)
			if err != nil {
				return modelMap, err
			}
			modelMap["spans_query"] = []map[string]interface{}{spansQueryMap}
		}
		if model.LogsQuery != nil {
			logsQueryMap, err := DataSourceIbmLogsE2msApisLogs2metricsV2LogsQueryToMap(model.LogsQuery)
			if err != nil {
				return modelMap, err
			}
			modelMap["logs_query"] = []map[string]interface{}{logsQueryMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.Event2MetricIntf subtype encountered")
	}
}

func DataSourceIbmLogsE2msApisEvents2metricsV2E2mPermutationsToMap(model *logsv0.ApisEvents2metricsV2E2mPermutations) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Limit != nil {
		modelMap["limit"] = flex.IntValue(model.Limit)
	}
	if model.HasExceededLimit != nil {
		modelMap["has_exceeded_limit"] = *model.HasExceededLimit
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisEvents2metricsV2MetricLabelToMap(model *logsv0.ApisEvents2metricsV2MetricLabel) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetLabel != nil {
		modelMap["target_label"] = *model.TargetLabel
	}
	if model.SourceField != nil {
		modelMap["source_field"] = *model.SourceField
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisEvents2metricsV2MetricFieldToMap(model *logsv0.ApisEvents2metricsV2MetricField) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetBaseMetricName != nil {
		modelMap["target_base_metric_name"] = *model.TargetBaseMetricName
	}
	if model.SourceField != nil {
		modelMap["source_field"] = *model.SourceField
	}
	if model.Aggregations != nil {
		aggregations := []map[string]interface{}{}
		for _, aggregationsItem := range model.Aggregations {
			aggregationsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2AggregationToMap(aggregationsItem)
			if err != nil {
				return modelMap, err
			}
			aggregations = append(aggregations, aggregationsItemMap)
		}
		modelMap["aggregations"] = aggregations
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisEvents2metricsV2AggregationToMap(model logsv0.ApisEvents2metricsV2AggregationIntf) (map[string]interface{}, error) {
	if _, ok := model.(*logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples); ok {
		return DataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataSamplesToMap(model.(*logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples))
	} else if _, ok := model.(*logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram); ok {
		return DataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataHistogramToMap(model.(*logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram))
	} else if _, ok := model.(*logsv0.ApisEvents2metricsV2Aggregation); ok {
		modelMap := make(map[string]interface{})
		model := model.(*logsv0.ApisEvents2metricsV2Aggregation)
		if model.Enabled != nil {
			modelMap["enabled"] = *model.Enabled
		}
		if model.AggType != nil {
			modelMap["agg_type"] = *model.AggType
		}
		if model.TargetMetricName != nil {
			modelMap["target_metric_name"] = *model.TargetMetricName
		}
		if model.Samples != nil {
			samplesMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggSamplesToMap(model.Samples)
			if err != nil {
				return modelMap, err
			}
			modelMap["samples"] = []map[string]interface{}{samplesMap}
		}
		if model.Histogram != nil {
			histogramMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggHistogramToMap(model.Histogram)
			if err != nil {
				return modelMap, err
			}
			modelMap["histogram"] = []map[string]interface{}{histogramMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized logsv0.ApisEvents2metricsV2AggregationIntf subtype encountered")
	}
}

func DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggSamplesToMap(model *logsv0.ApisEvents2metricsV2E2mAggSamples) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.SampleType != nil {
		modelMap["sample_type"] = *model.SampleType
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggHistogramToMap(model *logsv0.ApisEvents2metricsV2E2mAggHistogram) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Buckets != nil {
		var buckets []interface{}
		for _, item := range model.Buckets {
			buckets = append(buckets, float64(item))
		}
		modelMap["buckets"] = buckets
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataSamplesToMap(model *logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.AggType != nil {
		modelMap["agg_type"] = *model.AggType
	}
	if model.TargetMetricName != nil {
		modelMap["target_metric_name"] = *model.TargetMetricName
	}
	if model.Samples != nil {
		samplesMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggSamplesToMap(model.Samples)
		if err != nil {
			return modelMap, err
		}
		modelMap["samples"] = []map[string]interface{}{samplesMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataHistogramToMap(model *logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Enabled != nil {
		modelMap["enabled"] = *model.Enabled
	}
	if model.AggType != nil {
		modelMap["agg_type"] = *model.AggType
	}
	if model.TargetMetricName != nil {
		modelMap["target_metric_name"] = *model.TargetMetricName
	}
	if model.Histogram != nil {
		histogramMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggHistogramToMap(model.Histogram)
		if err != nil {
			return modelMap, err
		}
		modelMap["histogram"] = []map[string]interface{}{histogramMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisSpans2metricsV2SpansQueryToMap(model *logsv0.ApisSpans2metricsV2SpansQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Lucene != nil {
		modelMap["lucene"] = *model.Lucene
	}
	if model.ApplicationnameFilters != nil {
		modelMap["applicationname_filters"] = model.ApplicationnameFilters
	}
	if model.SubsystemnameFilters != nil {
		modelMap["subsystemname_filters"] = model.SubsystemnameFilters
	}
	if model.ActionFilters != nil {
		modelMap["action_filters"] = model.ActionFilters
	}
	if model.ServiceFilters != nil {
		modelMap["service_filters"] = model.ServiceFilters
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msApisLogs2metricsV2LogsQueryToMap(model *logsv0.ApisLogs2metricsV2LogsQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Lucene != nil {
		modelMap["lucene"] = *model.Lucene
	}
	if model.Alias != nil {
		modelMap["alias"] = *model.Alias
	}
	if model.ApplicationnameFilters != nil {
		modelMap["applicationname_filters"] = model.ApplicationnameFilters
	}
	if model.SubsystemnameFilters != nil {
		modelMap["subsystemname_filters"] = model.SubsystemnameFilters
	}
	if model.SeverityFilters != nil {
		modelMap["severity_filters"] = model.SeverityFilters
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQuerySpansQueryToMap(model *logsv0.Event2MetricApisEvents2metricsV2E2mQuerySpansQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.CreateTime != nil {
		modelMap["create_time"] = *model.CreateTime
	}
	if model.UpdateTime != nil {
		modelMap["update_time"] = *model.UpdateTime
	}
	if model.Permutations != nil {
		permutationsMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mPermutationsToMap(model.Permutations)
		if err != nil {
			return modelMap, err
		}
		modelMap["permutations"] = []map[string]interface{}{permutationsMap}
	}
	if model.MetricLabels != nil {
		metricLabels := []map[string]interface{}{}
		for _, metricLabelsItem := range model.MetricLabels {
			metricLabelsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2MetricLabelToMap(&metricLabelsItem)
			if err != nil {
				return modelMap, err
			}
			metricLabels = append(metricLabels, metricLabelsItemMap)
		}
		modelMap["metric_labels"] = metricLabels
	}
	if model.MetricFields != nil {
		metricFields := []map[string]interface{}{}
		for _, metricFieldsItem := range model.MetricFields {
			metricFieldsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2MetricFieldToMap(&metricFieldsItem)
			if err != nil {
				return modelMap, err
			}
			metricFields = append(metricFields, metricFieldsItemMap)
		}
		modelMap["metric_fields"] = metricFields
	}
	modelMap["type"] = *model.Type
	if model.IsInternal != nil {
		modelMap["is_internal"] = *model.IsInternal
	}
	if model.SpansQuery != nil {
		spansQueryMap, err := DataSourceIbmLogsE2msApisSpans2metricsV2SpansQueryToMap(model.SpansQuery)
		if err != nil {
			return modelMap, err
		}
		modelMap["spans_query"] = []map[string]interface{}{spansQueryMap}
	}
	return modelMap, nil
}

func DataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQueryLogsQueryToMap(model *logsv0.Event2MetricApisEvents2metricsV2E2mQueryLogsQuery) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID.String()
	}
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.CreateTime != nil {
		modelMap["create_time"] = *model.CreateTime
	}
	if model.UpdateTime != nil {
		modelMap["update_time"] = *model.UpdateTime
	}
	if model.Permutations != nil {
		permutationsMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2E2mPermutationsToMap(model.Permutations)
		if err != nil {
			return modelMap, err
		}
		modelMap["permutations"] = []map[string]interface{}{permutationsMap}
	}
	if model.MetricLabels != nil {
		metricLabels := []map[string]interface{}{}
		for _, metricLabelsItem := range model.MetricLabels {
			metricLabelsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2MetricLabelToMap(&metricLabelsItem)
			if err != nil {
				return modelMap, err
			}
			metricLabels = append(metricLabels, metricLabelsItemMap)
		}
		modelMap["metric_labels"] = metricLabels
	}
	if model.MetricFields != nil {
		metricFields := []map[string]interface{}{}
		for _, metricFieldsItem := range model.MetricFields {
			metricFieldsItemMap, err := DataSourceIbmLogsE2msApisEvents2metricsV2MetricFieldToMap(&metricFieldsItem)
			if err != nil {
				return modelMap, err
			}
			metricFields = append(metricFields, metricFieldsItemMap)
		}
		modelMap["metric_fields"] = metricFields
	}
	modelMap["type"] = *model.Type
	if model.IsInternal != nil {
		modelMap["is_internal"] = *model.IsInternal
	}
	if model.LogsQuery != nil {
		logsQueryMap, err := DataSourceIbmLogsE2msApisLogs2metricsV2LogsQueryToMap(model.LogsQuery)
		if err != nil {
			return modelMap, err
		}
		modelMap["logs_query"] = []map[string]interface{}{logsQueryMap}
	}
	return modelMap, nil
}
