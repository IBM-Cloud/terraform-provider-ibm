// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsE2mDataSourceBasic(t *testing.T) {
	event2MetricName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mDataSourceConfigBasic(event2MetricName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "logs_e2m_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "type"),
				),
			},
		},
	})
}

func TestAccIbmLogsE2mDataSourceAllArgs(t *testing.T) {
	event2MetricName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	event2MetricDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	event2MetricType := "logs2metrics"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mDataSourceConfig(event2MetricName, event2MetricDescription, event2MetricType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "logs_e2m_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "create_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "update_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "permutations.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "metric_labels.#"), //Todo: @kavya498
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "metric_labels.0.target_label"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "metric_labels.0.source_field"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "metric_fields.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "metric_fields.0.target_base_metric_name"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "metric_fields.0.source_field"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "is_internal"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "spans_query.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_e2m.logs_e2m_instance", "logs_query.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsE2mDataSourceConfigBasic(event2MetricName string) string {
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
	  data "ibm_logs_e2m" "logs_e2m_instance" {
		instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
		region      = ibm_logs_e2m.logs_e2m_instance.region
		logs_e2m_id = ibm_logs_e2m.logs_e2m_instance.e2m_id
	  }
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, event2MetricName)
}

func testAccCheckIbmLogsE2mDataSourceConfig(event2MetricName string, event2MetricDescription string, event2MetricType string) string {
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

	data "ibm_logs_e2m" "logs_e2m_instance" {
		instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
		region      = ibm_logs_e2m.logs_e2m_instance.region
		logs_e2m_id = ibm_logs_e2m.logs_e2m_instance.e2m_id
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, event2MetricName, event2MetricDescription, event2MetricType)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsE2mApisEvents2metricsV2E2mPermutationsToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["limit"] = int(38)
// 		model["has_exceeded_limit"] = true

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2E2mPermutations)
// 	model.Limit = core.Int64Ptr(int64(38))
// 	model.HasExceededLimit = core.BoolPtr(true)

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2E2mPermutationsToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2MetricLabelToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["target_label"] = "testString"
// 		model["source_field"] = "testString"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2MetricLabel)
// 	model.TargetLabel = core.StringPtr("testString")
// 	model.SourceField = core.StringPtr("testString")

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2MetricLabelToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2MetricFieldToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2MetricFieldToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2AggregationToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2AggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2E2mAggSamplesToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["sample_type"] = "unspecified"

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 	model.SampleType = core.StringPtr("unspecified")

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2E2mAggSamplesToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2E2mAggHistogramToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["buckets"] = []interface{}{float64(36.0)}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
// 	model.Buckets = []float32{float32(36.0)}

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2E2mAggHistogramToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataSamplesToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataSamplesToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataHistogramToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataHistogramToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisSpans2metricsV2SpansQueryToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsE2mApisSpans2metricsV2SpansQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

// func TestDataSourceIbmLogsE2mApisLogs2metricsV2LogsQueryToMap(t *testing.T) {
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

// 	result, err := logs.DataSourceIbmLogsE2mApisLogs2metricsV2LogsQueryToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }
