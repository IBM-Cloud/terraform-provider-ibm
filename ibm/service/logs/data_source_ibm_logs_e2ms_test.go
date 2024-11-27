// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	// . "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
)

func TestAccIbmLogsE2msDataSourceBasic(t *testing.T) {
	event2MetricName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2msDataSourceConfigBasic(event2MetricName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsE2msDataSourceAllArgs(t *testing.T) {
	event2MetricName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	event2MetricDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	event2MetricType := "logs2metrics"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2msDataSourceConfig(event2MetricName, event2MetricDescription, event2MetricType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.name", event2MetricName),
					resource.TestCheckResourceAttr("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.description", event2MetricDescription),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.create_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.update_time"),
					resource.TestCheckResourceAttr("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.type", event2MetricType),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2ms.logs_e2ms_instance", "events2metrics.0.is_internal"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsE2msDataSourceConfigBasic(event2MetricName string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_e2m" "logs_e2m_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "test description"
		logs_query {
		  applicationname_filters = []
		  severity_filters = [
			"debug", "error"
		  ]
		  subsystemname_filters = []
		}
		type = "logs2metrics"
	  }
	data "ibm_logs_e2ms" "logs_e2ms_instance" {
		instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
		region      = ibm_logs_e2m.logs_e2m_instance.region
		depends_on = [
			ibm_logs_e2m.logs_e2m_instance
		]
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, event2MetricName)
}

func testAccCheckIbmLogsE2msDataSourceConfig(event2MetricName string, event2MetricDescription string, event2MetricType string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_e2m" "logs_e2m_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "%s"
			logs_query {
			applicationname_filters = []
			severity_filters = [
				"debug", "error"
			]
			subsystemname_filters = []
			}
			type = "%s"
		}

		data "ibm_logs_e2ms" "logs_e2ms_instance" {
			instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
			region      = ibm_logs_e2m.logs_e2m_instance.region
			depends_on  = [
				ibm_logs_e2m.logs_e2m_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, event2MetricName, event2MetricDescription, event2MetricType)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsE2msEvent2MetricToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mPermutationsModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mPermutationsModel["limit"] = int(1)
// 		apisEvents2metricsV2E2mPermutationsModel["has_exceeded_limit"] = true

// 		apisEvents2metricsV2MetricLabelModel := make(map[string]interface{})
// 		apisEvents2metricsV2MetricLabelModel["target_label"] = "testString"
// 		apisEvents2metricsV2MetricLabelModel["source_field"] = "testString"

// 		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 		apisEvents2metricsV2AggregationModel := make(map[string]interface{})
// 		apisEvents2metricsV2AggregationModel["enabled"] = true
// 		apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
// 		apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
// 		apisEvents2metricsV2AggregationModel["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

// 		apisEvents2metricsV2MetricFieldModel := make(map[string]interface{})
// 		apisEvents2metricsV2MetricFieldModel["target_base_metric_name"] = "testString"
// 		apisEvents2metricsV2MetricFieldModel["source_field"] = "testString"
// 		apisEvents2metricsV2MetricFieldModel["aggregations"] = []map[string]interface{}{apisEvents2metricsV2AggregationModel}

// 		apisLogs2metricsV2LogsQueryModel := make(map[string]interface{})
// 		apisLogs2metricsV2LogsQueryModel["lucene"] = "logs"
// 		apisLogs2metricsV2LogsQueryModel["alias"] = "testString"
// 		apisLogs2metricsV2LogsQueryModel["applicationname_filters"] = []string{"testString"}
// 		apisLogs2metricsV2LogsQueryModel["subsystemname_filters"] = []string{"testString"}
// 		apisLogs2metricsV2LogsQueryModel["severity_filters"] = []string{"unspecified"}

// 		model := make(map[string]interface{})
// 		model["id"] = "d6a3658e-78d2-47d0-9b81-b2c551f01b09"
// 		model["name"] = "Service_catalog_latency"
// 		model["description"] = "avg and max the latency of catalog service"
// 		model["create_time"] = "2022-06-30T12:30:00Z'"
// 		model["update_time"] = "2022-06-30T12:30:00Z'"
// 		model["permutations"] = []map[string]interface{}{apisEvents2metricsV2E2mPermutationsModel}
// 		model["metric_labels"] = []map[string]interface{}{apisEvents2metricsV2MetricLabelModel}
// 		model["metric_fields"] = []map[string]interface{}{apisEvents2metricsV2MetricFieldModel}
// 		model["type"] = "unspecified"
// 		model["is_internal"] = true
// 		model["spans_query"] = []map[string]interface{}{apisSpans2metricsV2SpansQueryModel}
// 		model["logs_query"] = []map[string]interface{}{apisLogs2metricsV2LogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mPermutationsModel := new(logsv0.ApisEvents2metricsV2E2mPermutations)
// 	apisEvents2metricsV2E2mPermutationsModel.Limit = core.Int64Ptr(int64(1))
// 	apisEvents2metricsV2E2mPermutationsModel.HasExceededLimit = core.BoolPtr(true)

// 	apisEvents2metricsV2MetricLabelModel := new(logsv0.ApisEvents2metricsV2MetricLabel)
// 	apisEvents2metricsV2MetricLabelModel.TargetLabel = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricLabelModel.SourceField = core.StringPtr("testString")

// 	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 	apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
// 	apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
// 	apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
// 	apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

// 	apisEvents2metricsV2MetricFieldModel := new(logsv0.ApisEvents2metricsV2MetricField)
// 	apisEvents2metricsV2MetricFieldModel.TargetBaseMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricFieldModel.SourceField = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricFieldModel.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

// 	apisLogs2metricsV2LogsQueryModel := new(logsv0.ApisLogs2metricsV2LogsQuery)
// 	apisLogs2metricsV2LogsQueryModel.Lucene = core.StringPtr("logs")
// 	apisLogs2metricsV2LogsQueryModel.Alias = core.StringPtr("testString")
// 	apisLogs2metricsV2LogsQueryModel.ApplicationnameFilters = []string{"testString"}
// 	apisLogs2metricsV2LogsQueryModel.SubsystemnameFilters = []string{"testString"}
// 	apisLogs2metricsV2LogsQueryModel.SeverityFilters = []string{"unspecified"}

// 	model := new(logsv0.Event2Metric)
// 	model.ID = CreateMockUUID("d6a3658e-78d2-47d0-9b81-b2c551f01b09")
// 	model.Name = core.StringPtr("Service_catalog_latency")
// 	model.Description = core.StringPtr("avg and max the latency of catalog service")
// 	model.CreateTime = core.StringPtr("2022-06-30T12:30:00Z'")
// 	model.UpdateTime = core.StringPtr("2022-06-30T12:30:00Z'")
// 	model.Permutations = apisEvents2metricsV2E2mPermutationsModel
// 	model.MetricLabels = []logsv0.ApisEvents2metricsV2MetricLabel{*apisEvents2metricsV2MetricLabelModel}
// 	model.MetricFields = []logsv0.ApisEvents2metricsV2MetricField{*apisEvents2metricsV2MetricFieldModel}
// 	model.Type = core.StringPtr("unspecified")
// 	model.IsInternal = core.BoolPtr(true)
// 	model.SpansQuery = apisSpans2metricsV2SpansQueryModel
// 	model.LogsQuery = apisLogs2metricsV2LogsQueryModel

// 	result, err := logs.DataSourceIbmLogsE2msEvent2MetricToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2E2mPermutationsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["limit"] = int(38)
// 		model["has_exceeded_limit"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2E2mPermutations)
// 	model.Limit = core.Int64Ptr(int64(38))
// 	model.HasExceededLimit = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2E2mPermutationsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2MetricLabelToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["target_label"] = "testString"
// 		model["source_field"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2MetricLabel)
// 	model.TargetLabel = core.StringPtr("testString")
// 	model.SourceField = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2MetricLabelToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2MetricFieldToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 		apisEvents2metricsV2AggregationModel := make(map[string]interface{})
// 		apisEvents2metricsV2AggregationModel["enabled"] = true
// 		apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
// 		apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
// 		apisEvents2metricsV2AggregationModel["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

// 		model := make(map[string]interface{})
// 		model["target_base_metric_name"] = "testString"
// 		model["source_field"] = "testString"
// 		model["aggregations"] = []map[string]interface{}{apisEvents2metricsV2AggregationModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 	apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
// 	apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
// 	apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
// 	apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

// 	model := new(logsv0.ApisEvents2metricsV2MetricField)
// 	model.TargetBaseMetricName = core.StringPtr("testString")
// 	model.SourceField = core.StringPtr("testString")
// 	model.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2MetricFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2AggregationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["enabled"] = true
// 		model["agg_type"] = "unspecified"
// 		model["target_metric_name"] = "testString"
// 		model["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}
// 		model["histogram"] = []map[string]interface{}{apisEvents2metricsV2E2mAggHistogramModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisEvents2metricsV2Aggregation)
// 	model.Enabled = core.BoolPtr(true)
// 	model.AggType = core.StringPtr("unspecified")
// 	model.TargetMetricName = core.StringPtr("testString")
// 	model.Samples = apisEvents2metricsV2E2mAggSamplesModel
// 	model.Histogram = apisEvents2metricsV2E2mAggHistogramModel

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2AggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2E2mAggSamplesToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["sample_type"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	model.SampleType = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggSamplesToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2E2mAggHistogramToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["buckets"] = []interface{}{float64(36.0)}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
// 	model.Buckets = []float32{float32(36.0)}

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2E2mAggHistogramToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataSamplesToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 		model := make(map[string]interface{})
// 		model["enabled"] = true
// 		model["agg_type"] = "unspecified"
// 		model["target_metric_name"] = "testString"
// 		model["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 	model := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
// 	model.Enabled = core.BoolPtr(true)
// 	model.AggType = core.StringPtr("unspecified")
// 	model.TargetMetricName = core.StringPtr("testString")
// 	model.Samples = apisEvents2metricsV2E2mAggSamplesModel

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataSamplesToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataHistogramToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mAggHistogramModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggHistogramModel["buckets"] = []interface{}{float64(36.0)}

// 		model := make(map[string]interface{})
// 		model["enabled"] = true
// 		model["agg_type"] = "unspecified"
// 		model["target_metric_name"] = "testString"
// 		model["histogram"] = []map[string]interface{}{apisEvents2metricsV2E2mAggHistogramModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mAggHistogramModel := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
// 	apisEvents2metricsV2E2mAggHistogramModel.Buckets = []float32{float32(36.0)}

// 	model := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram)
// 	model.Enabled = core.BoolPtr(true)
// 	model.AggType = core.StringPtr("unspecified")
// 	model.TargetMetricName = core.StringPtr("testString")
// 	model.Histogram = apisEvents2metricsV2E2mAggHistogramModel

// 	result, err := logs.DataSourceIbmLogsE2msApisEvents2metricsV2AggregationAggMetadataHistogramToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisSpans2metricsV2SpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["lucene"] = "testString"
// 		model["applicationname_filters"] = []string{"testString"}
// 		model["subsystemname_filters"] = []string{"testString"}
// 		model["action_filters"] = []string{"testString"}
// 		model["service_filters"] = []string{"testString"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisSpans2metricsV2SpansQuery)
// 	model.Lucene = core.StringPtr("testString")
// 	model.ApplicationnameFilters = []string{"testString"}
// 	model.SubsystemnameFilters = []string{"testString"}
// 	model.ActionFilters = []string{"testString"}
// 	model.ServiceFilters = []string{"testString"}

// 	result, err := logs.DataSourceIbmLogsE2msApisSpans2metricsV2SpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msApisLogs2metricsV2LogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["lucene"] = "testString"
// 		model["alias"] = "testString"
// 		model["applicationname_filters"] = []string{"testString"}
// 		model["subsystemname_filters"] = []string{"testString"}
// 		model["severity_filters"] = []string{"unspecified"}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisLogs2metricsV2LogsQuery)
// 	model.Lucene = core.StringPtr("testString")
// 	model.Alias = core.StringPtr("testString")
// 	model.ApplicationnameFilters = []string{"testString"}
// 	model.SubsystemnameFilters = []string{"testString"}
// 	model.SeverityFilters = []string{"unspecified"}

// 	result, err := logs.DataSourceIbmLogsE2msApisLogs2metricsV2LogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQuerySpansQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mPermutationsModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mPermutationsModel["limit"] = int(38)
// 		apisEvents2metricsV2E2mPermutationsModel["has_exceeded_limit"] = true

// 		apisEvents2metricsV2MetricLabelModel := make(map[string]interface{})
// 		apisEvents2metricsV2MetricLabelModel["target_label"] = "testString"
// 		apisEvents2metricsV2MetricLabelModel["source_field"] = "testString"

// 		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 		apisEvents2metricsV2AggregationModel := make(map[string]interface{})
// 		apisEvents2metricsV2AggregationModel["enabled"] = true
// 		apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
// 		apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
// 		apisEvents2metricsV2AggregationModel["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

// 		apisEvents2metricsV2MetricFieldModel := make(map[string]interface{})
// 		apisEvents2metricsV2MetricFieldModel["target_base_metric_name"] = "testString"
// 		apisEvents2metricsV2MetricFieldModel["source_field"] = "testString"
// 		apisEvents2metricsV2MetricFieldModel["aggregations"] = []map[string]interface{}{apisEvents2metricsV2AggregationModel}

// 		apisSpans2metricsV2SpansQueryModel := make(map[string]interface{})
// 		apisSpans2metricsV2SpansQueryModel["lucene"] = "testString"
// 		apisSpans2metricsV2SpansQueryModel["applicationname_filters"] = []string{"testString"}
// 		apisSpans2metricsV2SpansQueryModel["subsystemname_filters"] = []string{"testString"}
// 		apisSpans2metricsV2SpansQueryModel["action_filters"] = []string{"testString"}
// 		apisSpans2metricsV2SpansQueryModel["service_filters"] = []string{"testString"}

// 		model := make(map[string]interface{})
// 		model["id"] = "d6a3658e-78d2-47d0-9b81-b2c551f01b09"
// 		model["name"] = "Service_catalog_latency"
// 		model["description"] = "avg and max the latency of catalog service"
// 		model["create_time"] = "2022-06-30T12:30:00Z'"
// 		model["update_time"] = "2022-06-30T12:30:00Z'"
// 		model["permutations"] = []map[string]interface{}{apisEvents2metricsV2E2mPermutationsModel}
// 		model["metric_labels"] = []map[string]interface{}{apisEvents2metricsV2MetricLabelModel}
// 		model["metric_fields"] = []map[string]interface{}{apisEvents2metricsV2MetricFieldModel}
// 		model["type"] = "unspecified"
// 		model["is_internal"] = true
// 		model["spans_query"] = []map[string]interface{}{apisSpans2metricsV2SpansQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mPermutationsModel := new(logsv0.ApisEvents2metricsV2E2mPermutations)
// 	apisEvents2metricsV2E2mPermutationsModel.Limit = core.Int64Ptr(int64(38))
// 	apisEvents2metricsV2E2mPermutationsModel.HasExceededLimit = core.BoolPtr(true)

// 	apisEvents2metricsV2MetricLabelModel := new(logsv0.ApisEvents2metricsV2MetricLabel)
// 	apisEvents2metricsV2MetricLabelModel.TargetLabel = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricLabelModel.SourceField = core.StringPtr("testString")

// 	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 	apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
// 	apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
// 	apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
// 	apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

// 	apisEvents2metricsV2MetricFieldModel := new(logsv0.ApisEvents2metricsV2MetricField)
// 	apisEvents2metricsV2MetricFieldModel.TargetBaseMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricFieldModel.SourceField = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricFieldModel.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

// 	apisSpans2metricsV2SpansQueryModel := new(logsv0.ApisSpans2metricsV2SpansQuery)
// 	apisSpans2metricsV2SpansQueryModel.Lucene = core.StringPtr("testString")
// 	apisSpans2metricsV2SpansQueryModel.ApplicationnameFilters = []string{"testString"}
// 	apisSpans2metricsV2SpansQueryModel.SubsystemnameFilters = []string{"testString"}
// 	apisSpans2metricsV2SpansQueryModel.ActionFilters = []string{"testString"}
// 	apisSpans2metricsV2SpansQueryModel.ServiceFilters = []string{"testString"}

// 	model := new(logsv0.Event2MetricApisEvents2metricsV2E2mQuerySpansQuery)
// 	model.ID = CreateMockUUID("d6a3658e-78d2-47d0-9b81-b2c551f01b09")
// 	model.Name = core.StringPtr("Service_catalog_latency")
// 	model.Description = core.StringPtr("avg and max the latency of catalog service")
// 	model.CreateTime = core.StringPtr("2022-06-30T12:30:00Z'")
// 	model.UpdateTime = core.StringPtr("2022-06-30T12:30:00Z'")
// 	model.Permutations = apisEvents2metricsV2E2mPermutationsModel
// 	model.MetricLabels = []logsv0.ApisEvents2metricsV2MetricLabel{*apisEvents2metricsV2MetricLabelModel}
// 	model.MetricFields = []logsv0.ApisEvents2metricsV2MetricField{*apisEvents2metricsV2MetricFieldModel}
// 	model.Type = core.StringPtr("unspecified")
// 	model.IsInternal = core.BoolPtr(true)
// 	model.SpansQuery = apisSpans2metricsV2SpansQueryModel

// 	result, err := logs.DataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQuerySpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQueryLogsQueryToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		apisEvents2metricsV2E2mPermutationsModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mPermutationsModel["limit"] = int(38)
// 		apisEvents2metricsV2E2mPermutationsModel["has_exceeded_limit"] = true

// 		apisEvents2metricsV2MetricLabelModel := make(map[string]interface{})
// 		apisEvents2metricsV2MetricLabelModel["target_label"] = "testString"
// 		apisEvents2metricsV2MetricLabelModel["source_field"] = "testString"

// 		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 		apisEvents2metricsV2AggregationModel := make(map[string]interface{})
// 		apisEvents2metricsV2AggregationModel["enabled"] = true
// 		apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
// 		apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
// 		apisEvents2metricsV2AggregationModel["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

// 		apisEvents2metricsV2MetricFieldModel := make(map[string]interface{})
// 		apisEvents2metricsV2MetricFieldModel["target_base_metric_name"] = "testString"
// 		apisEvents2metricsV2MetricFieldModel["source_field"] = "testString"
// 		apisEvents2metricsV2MetricFieldModel["aggregations"] = []map[string]interface{}{apisEvents2metricsV2AggregationModel}

// 		apisLogs2metricsV2LogsQueryModel := make(map[string]interface{})
// 		apisLogs2metricsV2LogsQueryModel["lucene"] = "testString"
// 		apisLogs2metricsV2LogsQueryModel["alias"] = "testString"
// 		apisLogs2metricsV2LogsQueryModel["applicationname_filters"] = []string{"testString"}
// 		apisLogs2metricsV2LogsQueryModel["subsystemname_filters"] = []string{"testString"}
// 		apisLogs2metricsV2LogsQueryModel["severity_filters"] = []string{"unspecified"}

// 		model := make(map[string]interface{})
// 		model["id"] = "d6a3658e-78d2-47d0-9b81-b2c551f01b09"
// 		model["name"] = "Service_catalog_latency"
// 		model["description"] = "avg and max the latency of catalog service"
// 		model["create_time"] = "2022-06-30T12:30:00Z'"
// 		model["update_time"] = "2022-06-30T12:30:00Z'"
// 		model["permutations"] = []map[string]interface{}{apisEvents2metricsV2E2mPermutationsModel}
// 		model["metric_labels"] = []map[string]interface{}{apisEvents2metricsV2MetricLabelModel}
// 		model["metric_fields"] = []map[string]interface{}{apisEvents2metricsV2MetricFieldModel}
// 		model["type"] = "unspecified"
// 		model["is_internal"] = true
// 		model["logs_query"] = []map[string]interface{}{apisLogs2metricsV2LogsQueryModel}

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mPermutationsModel := new(logsv0.ApisEvents2metricsV2E2mPermutations)
// 	apisEvents2metricsV2E2mPermutationsModel.Limit = core.Int64Ptr(int64(38))
// 	apisEvents2metricsV2E2mPermutationsModel.HasExceededLimit = core.BoolPtr(true)

// 	apisEvents2metricsV2MetricLabelModel := new(logsv0.ApisEvents2metricsV2MetricLabel)
// 	apisEvents2metricsV2MetricLabelModel.TargetLabel = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricLabelModel.SourceField = core.StringPtr("testString")

// 	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 	apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
// 	apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
// 	apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
// 	apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

// 	apisEvents2metricsV2MetricFieldModel := new(logsv0.ApisEvents2metricsV2MetricField)
// 	apisEvents2metricsV2MetricFieldModel.TargetBaseMetricName = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricFieldModel.SourceField = core.StringPtr("testString")
// 	apisEvents2metricsV2MetricFieldModel.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

// 	apisLogs2metricsV2LogsQueryModel := new(logsv0.ApisLogs2metricsV2LogsQuery)
// 	apisLogs2metricsV2LogsQueryModel.Lucene = core.StringPtr("testString")
// 	apisLogs2metricsV2LogsQueryModel.Alias = core.StringPtr("testString")
// 	apisLogs2metricsV2LogsQueryModel.ApplicationnameFilters = []string{"testString"}
// 	apisLogs2metricsV2LogsQueryModel.SubsystemnameFilters = []string{"testString"}
// 	apisLogs2metricsV2LogsQueryModel.SeverityFilters = []string{"unspecified"}

// 	model := new(logsv0.Event2MetricApisEvents2metricsV2E2mQueryLogsQuery)
// 	model.ID = CreateMockUUID("d6a3658e-78d2-47d0-9b81-b2c551f01b09")
// 	model.Name = core.StringPtr("Service_catalog_latency")
// 	model.Description = core.StringPtr("avg and max the latency of catalog service")
// 	model.CreateTime = core.StringPtr("2022-06-30T12:30:00Z'")
// 	model.UpdateTime = core.StringPtr("2022-06-30T12:30:00Z'")
// 	model.Permutations = apisEvents2metricsV2E2mPermutationsModel
// 	model.MetricLabels = []logsv0.ApisEvents2metricsV2MetricLabel{*apisEvents2metricsV2MetricLabelModel}
// 	model.MetricFields = []logsv0.ApisEvents2metricsV2MetricField{*apisEvents2metricsV2MetricFieldModel}
// 	model.Type = core.StringPtr("unspecified")
// 	model.IsInternal = core.BoolPtr(true)
// 	model.LogsQuery = apisLogs2metricsV2LogsQueryModel

// 	result, err := logs.DataSourceIbmLogsE2msEvent2MetricApisEvents2metricsV2E2mQueryLogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }
