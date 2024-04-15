// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsE2mBasic(t *testing.T) {
	var conf logsv0.Event2Metric
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsE2mDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsE2mExists("ibm_logs_e2m.logs_e2m_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsE2mAllArgs(t *testing.T) {
	var conf logsv0.Event2Metric
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	typeVar := "logs2metrics"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "logs2metrics"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsE2mDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfig(name, description, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsE2mExists("ibm_logs_e2m.logs_e2m_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsE2mConfig(nameUpdate, descriptionUpdate, typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_e2m.logs_e2m_instance", "type", typeVarUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_e2m.logs_e2m_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsE2mConfigBasic(name string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsE2mConfig(name string, description string, typeVar string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, typeVar)
}

func testAccCheckIbmLogsE2mExists(n string, obj logsv0.Event2Metric) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getE2mOptions := &logsv0.GetE2mOptions{}

		getE2mOptions.SetID(resourceID[2])

		event2MetricIntf, _, err := logsClient.GetE2m(getE2mOptions)
		if err != nil {
			return err
		}

		event2Metric := event2MetricIntf.(*logsv0.Event2Metric)
		obj = *event2Metric
		return nil
	}
}

func testAccCheckIbmLogsE2mDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_e2m" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getE2mOptions := &logsv0.GetE2mOptions{}

		getE2mOptions.SetID(resourceID[2])

		// Try to find the key
		_, response, err := logsClient.GetE2m(getE2mOptions)

		if err == nil {
			return fmt.Errorf("logs_e2m still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_e2m (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsE2mApisEvents2metricsV2MetricLabelToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_label"] = "testString"
		model["source_field"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisEvents2metricsV2MetricLabel)
	model.TargetLabel = core.StringPtr("testString")
	model.SourceField = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2MetricLabelToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisEvents2metricsV2MetricFieldToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

		apisEvents2metricsV2AggregationModel := make(map[string]interface{})
		apisEvents2metricsV2AggregationModel["enabled"] = true
		apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
		apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
		apisEvents2metricsV2AggregationModel["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

		model := make(map[string]interface{})
		model["target_base_metric_name"] = "testString"
		model["source_field"] = "testString"
		model["aggregations"] = []map[string]interface{}{apisEvents2metricsV2AggregationModel}

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

	apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
	apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
	apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
	apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
	apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

	model := new(logsv0.ApisEvents2metricsV2MetricField)
	model.TargetBaseMetricName = core.StringPtr("testString")
	model.SourceField = core.StringPtr("testString")
	model.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2MetricFieldToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsE2mApisEvents2metricsV2AggregationToMap(t *testing.T) {
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

// 	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2AggregationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsE2mApisEvents2metricsV2E2mAggSamplesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["sample_type"] = "unspecified"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
	model.SampleType = core.StringPtr("unspecified")

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2E2mAggSamplesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisEvents2metricsV2E2mAggHistogramToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["buckets"] = []interface{}{float64(36.0)}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
	model.Buckets = []float32{float32(36.0)}

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2E2mAggHistogramToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataSamplesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
		apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

		model := make(map[string]interface{})
		model["enabled"] = true
		model["agg_type"] = "unspecified"
		model["target_metric_name"] = "testString"
		model["samples"] = []map[string]interface{}{apisEvents2metricsV2E2mAggSamplesModel}

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
	apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

	model := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
	model.Enabled = core.BoolPtr(true)
	model.AggType = core.StringPtr("unspecified")
	model.TargetMetricName = core.StringPtr("testString")
	model.Samples = apisEvents2metricsV2E2mAggSamplesModel

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataSamplesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataHistogramToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisEvents2metricsV2E2mAggHistogramModel := make(map[string]interface{})
		apisEvents2metricsV2E2mAggHistogramModel["buckets"] = []interface{}{float64(36.0)}

		model := make(map[string]interface{})
		model["enabled"] = true
		model["agg_type"] = "unspecified"
		model["target_metric_name"] = "testString"
		model["histogram"] = []map[string]interface{}{apisEvents2metricsV2E2mAggHistogramModel}

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2E2mAggHistogramModel := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
	apisEvents2metricsV2E2mAggHistogramModel.Buckets = []float32{float32(36.0)}

	model := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram)
	model.Enabled = core.BoolPtr(true)
	model.AggType = core.StringPtr("unspecified")
	model.TargetMetricName = core.StringPtr("testString")
	model.Histogram = apisEvents2metricsV2E2mAggHistogramModel

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2AggregationAggMetadataHistogramToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisSpans2metricsV2SpansQueryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["lucene"] = "testString"
		model["applicationname_filters"] = []string{"testString"}
		model["subsystemname_filters"] = []string{"testString"}
		model["action_filters"] = []string{"testString"}
		model["service_filters"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisSpans2metricsV2SpansQuery)
	model.Lucene = core.StringPtr("testString")
	model.ApplicationnameFilters = []string{"testString"}
	model.SubsystemnameFilters = []string{"testString"}
	model.ActionFilters = []string{"testString"}
	model.ServiceFilters = []string{"testString"}

	result, err := logs.ResourceIbmLogsE2mApisSpans2metricsV2SpansQueryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisLogs2metricsV2LogsQueryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["lucene"] = "testString"
		model["alias"] = "testString"
		model["applicationname_filters"] = []string{"testString"}
		model["subsystemname_filters"] = []string{"testString"}
		model["severity_filters"] = []string{"unspecified"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisLogs2metricsV2LogsQuery)
	model.Lucene = core.StringPtr("testString")
	model.Alias = core.StringPtr("testString")
	model.ApplicationnameFilters = []string{"testString"}
	model.SubsystemnameFilters = []string{"testString"}
	model.SeverityFilters = []string{"unspecified"}

	result, err := logs.ResourceIbmLogsE2mApisLogs2metricsV2LogsQueryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mApisEvents2metricsV2E2mPermutationsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["limit"] = int(38)
		model["has_exceeded_limit"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisEvents2metricsV2E2mPermutations)
	model.Limit = core.Int64Ptr(int64(38))
	model.HasExceededLimit = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsE2mApisEvents2metricsV2E2mPermutationsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisEvents2metricsV2MetricLabel(t *testing.T) {
	checkResult := func(result *logsv0.ApisEvents2metricsV2MetricLabel) {
		model := new(logsv0.ApisEvents2metricsV2MetricLabel)
		model.TargetLabel = core.StringPtr("testString")
		model.SourceField = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["target_label"] = "testString"
	model["source_field"] = "testString"

	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2MetricLabel(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisEvents2metricsV2MetricField(t *testing.T) {
	checkResult := func(result *logsv0.ApisEvents2metricsV2MetricField) {
		apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
		apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

		apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
		apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
		apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
		apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
		apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

		model := new(logsv0.ApisEvents2metricsV2MetricField)
		model.TargetBaseMetricName = core.StringPtr("testString")
		model.SourceField = core.StringPtr("testString")
		model.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
	apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

	apisEvents2metricsV2AggregationModel := make(map[string]interface{})
	apisEvents2metricsV2AggregationModel["enabled"] = true
	apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
	apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
	apisEvents2metricsV2AggregationModel["samples"] = []interface{}{apisEvents2metricsV2E2mAggSamplesModel}

	model := make(map[string]interface{})
	model["target_base_metric_name"] = "testString"
	model["source_field"] = "testString"
	model["aggregations"] = []interface{}{apisEvents2metricsV2AggregationModel}

	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2MetricField(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsE2mMapToApisEvents2metricsV2Aggregation(t *testing.T) {
// 	checkResult := func(result logsv0.ApisEvents2metricsV2AggregationIntf) {
// 		apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 		apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 		model := new(logsv0.ApisEvents2metricsV2Aggregation)
// 		model.Enabled = core.BoolPtr(true)
// 		model.AggType = core.StringPtr("unspecified")
// 		model.TargetMetricName = core.StringPtr("testString")
// 		model.Samples = apisEvents2metricsV2E2mAggSamplesModel
// 		model.Histogram = apisEvents2metricsV2E2mAggHistogramModel

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 	apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 	model := make(map[string]interface{})
// 	model["enabled"] = true
// 	model["agg_type"] = "unspecified"
// 	model["target_metric_name"] = "testString"
// 	model["samples"] = []interface{}{apisEvents2metricsV2E2mAggSamplesModel}
// 	model["histogram"] = []interface{}{apisEvents2metricsV2E2mAggHistogramModel}

// 	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2Aggregation(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsE2mMapToApisEvents2metricsV2E2mAggSamples(t *testing.T) {
	checkResult := func(result *logsv0.ApisEvents2metricsV2E2mAggSamples) {
		model := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
		model.SampleType = core.StringPtr("unspecified")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["sample_type"] = "unspecified"

	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2E2mAggSamples(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisEvents2metricsV2E2mAggHistogram(t *testing.T) {
	checkResult := func(result *logsv0.ApisEvents2metricsV2E2mAggHistogram) {
		model := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
		model.Buckets = []float32{float32(36.0)}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["buckets"] = []interface{}{float64(36.0)}

	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2E2mAggHistogram(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisEvents2metricsV2AggregationAggMetadataSamples(t *testing.T) {
	checkResult := func(result *logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples) {
		apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
		apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

		model := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
		model.Enabled = core.BoolPtr(true)
		model.AggType = core.StringPtr("unspecified")
		model.TargetMetricName = core.StringPtr("testString")
		model.Samples = apisEvents2metricsV2E2mAggSamplesModel

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
	apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

	model := make(map[string]interface{})
	model["enabled"] = true
	model["agg_type"] = "unspecified"
	model["target_metric_name"] = "testString"
	model["samples"] = []interface{}{apisEvents2metricsV2E2mAggSamplesModel}

	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2AggregationAggMetadataSamples(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisEvents2metricsV2AggregationAggMetadataHistogram(t *testing.T) {
	checkResult := func(result *logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram) {
		apisEvents2metricsV2E2mAggHistogramModel := new(logsv0.ApisEvents2metricsV2E2mAggHistogram)
		apisEvents2metricsV2E2mAggHistogramModel.Buckets = []float32{float32(36.0)}

		model := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataHistogram)
		model.Enabled = core.BoolPtr(true)
		model.AggType = core.StringPtr("unspecified")
		model.TargetMetricName = core.StringPtr("testString")
		model.Histogram = apisEvents2metricsV2E2mAggHistogramModel

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2E2mAggHistogramModel := make(map[string]interface{})
	apisEvents2metricsV2E2mAggHistogramModel["buckets"] = []interface{}{float64(36.0)}

	model := make(map[string]interface{})
	model["enabled"] = true
	model["agg_type"] = "unspecified"
	model["target_metric_name"] = "testString"
	model["histogram"] = []interface{}{apisEvents2metricsV2E2mAggHistogramModel}

	result, err := logs.ResourceIbmLogsE2mMapToApisEvents2metricsV2AggregationAggMetadataHistogram(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisSpans2metricsV2SpansQuery(t *testing.T) {
	checkResult := func(result *logsv0.ApisSpans2metricsV2SpansQuery) {
		model := new(logsv0.ApisSpans2metricsV2SpansQuery)
		model.Lucene = core.StringPtr("testString")
		model.ApplicationnameFilters = []string{"testString"}
		model.SubsystemnameFilters = []string{"testString"}
		model.ActionFilters = []string{"testString"}
		model.ServiceFilters = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["lucene"] = "testString"
	model["applicationname_filters"] = []interface{}{"testString"}
	model["subsystemname_filters"] = []interface{}{"testString"}
	model["action_filters"] = []interface{}{"testString"}
	model["service_filters"] = []interface{}{"testString"}

	result, err := logs.ResourceIbmLogsE2mMapToApisSpans2metricsV2SpansQuery(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToApisLogs2metricsV2LogsQuery(t *testing.T) {
	checkResult := func(result *logsv0.ApisLogs2metricsV2LogsQuery) {
		model := new(logsv0.ApisLogs2metricsV2LogsQuery)
		model.Lucene = core.StringPtr("testString")
		model.Alias = core.StringPtr("testString")
		model.ApplicationnameFilters = []string{"testString"}
		model.SubsystemnameFilters = []string{"testString"}
		model.SeverityFilters = []string{"unspecified"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["lucene"] = "testString"
	model["alias"] = "testString"
	model["applicationname_filters"] = []interface{}{"testString"}
	model["subsystemname_filters"] = []interface{}{"testString"}
	model["severity_filters"] = []interface{}{"unspecified"}

	result, err := logs.ResourceIbmLogsE2mMapToApisLogs2metricsV2LogsQuery(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsE2mMapToEvent2MetricPrototype(t *testing.T) {
// 	checkResult := func(result logsv0.Event2MetricPrototypeIntf) {
// 		apisEvents2metricsV2MetricLabelModel := new(logsv0.ApisEvents2metricsV2MetricLabel)
// 		apisEvents2metricsV2MetricLabelModel.TargetLabel = core.StringPtr("testString")
// 		apisEvents2metricsV2MetricLabelModel.SourceField = core.StringPtr("testString")

// 		apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
// 		apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

// 		apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
// 		apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
// 		apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
// 		apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
// 		apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

// 		apisEvents2metricsV2MetricFieldModel := new(logsv0.ApisEvents2metricsV2MetricField)
// 		apisEvents2metricsV2MetricFieldModel.TargetBaseMetricName = core.StringPtr("testString")
// 		apisEvents2metricsV2MetricFieldModel.SourceField = core.StringPtr("testString")
// 		apisEvents2metricsV2MetricFieldModel.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

// 		apisLogs2metricsV2LogsQueryModel := new(logsv0.ApisLogs2metricsV2LogsQuery)
// 		apisLogs2metricsV2LogsQueryModel.Lucene = core.StringPtr("logs")
// 		apisLogs2metricsV2LogsQueryModel.Alias = core.StringPtr("testString")
// 		apisLogs2metricsV2LogsQueryModel.ApplicationnameFilters = []string{"testString"}
// 		apisLogs2metricsV2LogsQueryModel.SubsystemnameFilters = []string{"testString"}
// 		apisLogs2metricsV2LogsQueryModel.SeverityFilters = []string{"unspecified"}

// 		model := new(logsv0.Event2MetricPrototype)
// 		model.Name = core.StringPtr("Service catalog latency")
// 		model.Description = core.StringPtr("avg and max the latency of catalog service")
// 		model.PermutationsLimit = core.Int64Ptr(int64(38))
// 		model.MetricLabels = []logsv0.ApisEvents2metricsV2MetricLabel{*apisEvents2metricsV2MetricLabelModel}
// 		model.MetricFields = []logsv0.ApisEvents2metricsV2MetricField{*apisEvents2metricsV2MetricFieldModel}
// 		model.Type = core.StringPtr("unspecified")
// 		model.SpansQuery = apisSpans2metricsV2SpansQueryModel
// 		model.LogsQuery = apisLogs2metricsV2LogsQueryModel

// 		assert.Equal(t, result, model)
// 	}

// 	apisEvents2metricsV2MetricLabelModel := make(map[string]interface{})
// 	apisEvents2metricsV2MetricLabelModel["target_label"] = "testString"
// 	apisEvents2metricsV2MetricLabelModel["source_field"] = "testString"

// 	apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
// 	apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

// 	apisEvents2metricsV2AggregationModel := make(map[string]interface{})
// 	apisEvents2metricsV2AggregationModel["enabled"] = true
// 	apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
// 	apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
// 	apisEvents2metricsV2AggregationModel["samples"] = []interface{}{apisEvents2metricsV2E2mAggSamplesModel}

// 	apisEvents2metricsV2MetricFieldModel := make(map[string]interface{})
// 	apisEvents2metricsV2MetricFieldModel["target_base_metric_name"] = "testString"
// 	apisEvents2metricsV2MetricFieldModel["source_field"] = "testString"
// 	apisEvents2metricsV2MetricFieldModel["aggregations"] = []interface{}{apisEvents2metricsV2AggregationModel}

// 	apisLogs2metricsV2LogsQueryModel := make(map[string]interface{})
// 	apisLogs2metricsV2LogsQueryModel["lucene"] = "logs"
// 	apisLogs2metricsV2LogsQueryModel["alias"] = "testString"
// 	apisLogs2metricsV2LogsQueryModel["applicationname_filters"] = []interface{}{"testString"}
// 	apisLogs2metricsV2LogsQueryModel["subsystemname_filters"] = []interface{}{"testString"}
// 	apisLogs2metricsV2LogsQueryModel["severity_filters"] = []interface{}{"unspecified"}

// 	model := make(map[string]interface{})
// 	model["name"] = "Service catalog latency"
// 	model["description"] = "avg and max the latency of catalog service"
// 	model["permutations_limit"] = int(38)
// 	model["metric_labels"] = []interface{}{apisEvents2metricsV2MetricLabelModel}
// 	model["metric_fields"] = []interface{}{apisEvents2metricsV2MetricFieldModel}
// 	model["type"] = "unspecified"
// 	model["spans_query"] = []interface{}{apisSpans2metricsV2SpansQueryModel}
// 	model["logs_query"] = []interface{}{apisLogs2metricsV2LogsQueryModel}

// 	result, err := logs.ResourceIbmLogsE2mMapToEvent2MetricPrototype(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsE2mMapToEvent2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQuerySpansQuery(t *testing.T) {
	checkResult := func(result *logsv0.Event2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQuerySpansQuery) {
		apisEvents2metricsV2MetricLabelModel := new(logsv0.ApisEvents2metricsV2MetricLabel)
		apisEvents2metricsV2MetricLabelModel.TargetLabel = core.StringPtr("testString")
		apisEvents2metricsV2MetricLabelModel.SourceField = core.StringPtr("testString")

		apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
		apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

		apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
		apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
		apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
		apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
		apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

		apisEvents2metricsV2MetricFieldModel := new(logsv0.ApisEvents2metricsV2MetricField)
		apisEvents2metricsV2MetricFieldModel.TargetBaseMetricName = core.StringPtr("testString")
		apisEvents2metricsV2MetricFieldModel.SourceField = core.StringPtr("testString")
		apisEvents2metricsV2MetricFieldModel.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

		apisSpans2metricsV2SpansQueryModel := new(logsv0.ApisSpans2metricsV2SpansQuery)
		apisSpans2metricsV2SpansQueryModel.Lucene = core.StringPtr("testString")
		apisSpans2metricsV2SpansQueryModel.ApplicationnameFilters = []string{"testString"}
		apisSpans2metricsV2SpansQueryModel.SubsystemnameFilters = []string{"testString"}
		apisSpans2metricsV2SpansQueryModel.ActionFilters = []string{"testString"}
		apisSpans2metricsV2SpansQueryModel.ServiceFilters = []string{"testString"}

		model := new(logsv0.Event2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQuerySpansQuery)
		model.Name = core.StringPtr("Service catalog latency")
		model.Description = core.StringPtr("avg and max the latency of catalog service")
		model.PermutationsLimit = core.Int64Ptr(int64(38))
		model.MetricLabels = []logsv0.ApisEvents2metricsV2MetricLabel{*apisEvents2metricsV2MetricLabelModel}
		model.MetricFields = []logsv0.ApisEvents2metricsV2MetricField{*apisEvents2metricsV2MetricFieldModel}
		model.Type = core.StringPtr("unspecified")
		model.SpansQuery = apisSpans2metricsV2SpansQueryModel

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2MetricLabelModel := make(map[string]interface{})
	apisEvents2metricsV2MetricLabelModel["target_label"] = "testString"
	apisEvents2metricsV2MetricLabelModel["source_field"] = "testString"

	apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
	apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

	apisEvents2metricsV2AggregationModel := make(map[string]interface{})
	apisEvents2metricsV2AggregationModel["enabled"] = true
	apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
	apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
	apisEvents2metricsV2AggregationModel["samples"] = []interface{}{apisEvents2metricsV2E2mAggSamplesModel}

	apisEvents2metricsV2MetricFieldModel := make(map[string]interface{})
	apisEvents2metricsV2MetricFieldModel["target_base_metric_name"] = "testString"
	apisEvents2metricsV2MetricFieldModel["source_field"] = "testString"
	apisEvents2metricsV2MetricFieldModel["aggregations"] = []interface{}{apisEvents2metricsV2AggregationModel}

	apisSpans2metricsV2SpansQueryModel := make(map[string]interface{})
	apisSpans2metricsV2SpansQueryModel["lucene"] = "testString"
	apisSpans2metricsV2SpansQueryModel["applicationname_filters"] = []interface{}{"testString"}
	apisSpans2metricsV2SpansQueryModel["subsystemname_filters"] = []interface{}{"testString"}
	apisSpans2metricsV2SpansQueryModel["action_filters"] = []interface{}{"testString"}
	apisSpans2metricsV2SpansQueryModel["service_filters"] = []interface{}{"testString"}

	model := make(map[string]interface{})
	model["name"] = "Service catalog latency"
	model["description"] = "avg and max the latency of catalog service"
	model["permutations_limit"] = int(38)
	model["metric_labels"] = []interface{}{apisEvents2metricsV2MetricLabelModel}
	model["metric_fields"] = []interface{}{apisEvents2metricsV2MetricFieldModel}
	model["type"] = "unspecified"
	model["spans_query"] = []interface{}{apisSpans2metricsV2SpansQueryModel}

	result, err := logs.ResourceIbmLogsE2mMapToEvent2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQuerySpansQuery(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsE2mMapToEvent2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQueryLogsQuery(t *testing.T) {
	checkResult := func(result *logsv0.Event2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQueryLogsQuery) {
		apisEvents2metricsV2MetricLabelModel := new(logsv0.ApisEvents2metricsV2MetricLabel)
		apisEvents2metricsV2MetricLabelModel.TargetLabel = core.StringPtr("testString")
		apisEvents2metricsV2MetricLabelModel.SourceField = core.StringPtr("testString")

		apisEvents2metricsV2E2mAggSamplesModel := new(logsv0.ApisEvents2metricsV2E2mAggSamples)
		apisEvents2metricsV2E2mAggSamplesModel.SampleType = core.StringPtr("unspecified")

		apisEvents2metricsV2AggregationModel := new(logsv0.ApisEvents2metricsV2AggregationAggMetadataSamples)
		apisEvents2metricsV2AggregationModel.Enabled = core.BoolPtr(true)
		apisEvents2metricsV2AggregationModel.AggType = core.StringPtr("unspecified")
		apisEvents2metricsV2AggregationModel.TargetMetricName = core.StringPtr("testString")
		apisEvents2metricsV2AggregationModel.Samples = apisEvents2metricsV2E2mAggSamplesModel

		apisEvents2metricsV2MetricFieldModel := new(logsv0.ApisEvents2metricsV2MetricField)
		apisEvents2metricsV2MetricFieldModel.TargetBaseMetricName = core.StringPtr("testString")
		apisEvents2metricsV2MetricFieldModel.SourceField = core.StringPtr("testString")
		apisEvents2metricsV2MetricFieldModel.Aggregations = []logsv0.ApisEvents2metricsV2AggregationIntf{apisEvents2metricsV2AggregationModel}

		apisLogs2metricsV2LogsQueryModel := new(logsv0.ApisLogs2metricsV2LogsQuery)
		apisLogs2metricsV2LogsQueryModel.Lucene = core.StringPtr("testString")
		apisLogs2metricsV2LogsQueryModel.Alias = core.StringPtr("testString")
		apisLogs2metricsV2LogsQueryModel.ApplicationnameFilters = []string{"testString"}
		apisLogs2metricsV2LogsQueryModel.SubsystemnameFilters = []string{"testString"}
		apisLogs2metricsV2LogsQueryModel.SeverityFilters = []string{"unspecified"}

		model := new(logsv0.Event2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQueryLogsQuery)
		model.Name = core.StringPtr("Service catalog latency")
		model.Description = core.StringPtr("avg and max the latency of catalog service")
		model.PermutationsLimit = core.Int64Ptr(int64(38))
		model.MetricLabels = []logsv0.ApisEvents2metricsV2MetricLabel{*apisEvents2metricsV2MetricLabelModel}
		model.MetricFields = []logsv0.ApisEvents2metricsV2MetricField{*apisEvents2metricsV2MetricFieldModel}
		model.Type = core.StringPtr("unspecified")
		model.LogsQuery = apisLogs2metricsV2LogsQueryModel

		assert.Equal(t, result, model)
	}

	apisEvents2metricsV2MetricLabelModel := make(map[string]interface{})
	apisEvents2metricsV2MetricLabelModel["target_label"] = "testString"
	apisEvents2metricsV2MetricLabelModel["source_field"] = "testString"

	apisEvents2metricsV2E2mAggSamplesModel := make(map[string]interface{})
	apisEvents2metricsV2E2mAggSamplesModel["sample_type"] = "unspecified"

	apisEvents2metricsV2AggregationModel := make(map[string]interface{})
	apisEvents2metricsV2AggregationModel["enabled"] = true
	apisEvents2metricsV2AggregationModel["agg_type"] = "unspecified"
	apisEvents2metricsV2AggregationModel["target_metric_name"] = "testString"
	apisEvents2metricsV2AggregationModel["samples"] = []interface{}{apisEvents2metricsV2E2mAggSamplesModel}

	apisEvents2metricsV2MetricFieldModel := make(map[string]interface{})
	apisEvents2metricsV2MetricFieldModel["target_base_metric_name"] = "testString"
	apisEvents2metricsV2MetricFieldModel["source_field"] = "testString"
	apisEvents2metricsV2MetricFieldModel["aggregations"] = []interface{}{apisEvents2metricsV2AggregationModel}

	apisLogs2metricsV2LogsQueryModel := make(map[string]interface{})
	apisLogs2metricsV2LogsQueryModel["lucene"] = "testString"
	apisLogs2metricsV2LogsQueryModel["alias"] = "testString"
	apisLogs2metricsV2LogsQueryModel["applicationname_filters"] = []interface{}{"testString"}
	apisLogs2metricsV2LogsQueryModel["subsystemname_filters"] = []interface{}{"testString"}
	apisLogs2metricsV2LogsQueryModel["severity_filters"] = []interface{}{"unspecified"}

	model := make(map[string]interface{})
	model["name"] = "Service catalog latency"
	model["description"] = "avg and max the latency of catalog service"
	model["permutations_limit"] = int(38)
	model["metric_labels"] = []interface{}{apisEvents2metricsV2MetricLabelModel}
	model["metric_fields"] = []interface{}{apisEvents2metricsV2MetricFieldModel}
	model["type"] = "unspecified"
	model["logs_query"] = []interface{}{apisLogs2metricsV2LogsQueryModel}

	result, err := logs.ResourceIbmLogsE2mMapToEvent2MetricPrototypeApisEvents2metricsV2E2mCreateParamsQueryLogsQuery(model)
	assert.Nil(t, err)
	checkResult(result)
}
