// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
*/

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsStreamsDataSourceBasic(t *testing.T) {
	streamName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	streamDpxlExpression := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamsDataSourceConfigBasic(streamName, streamDpxlExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "streams.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_streams.logs_streams_instance", "streams.0.name", streamName),
					resource.TestCheckResourceAttr("data.ibm_logs_streams.logs_streams_instance", "streams.0.dpxl_expression", streamDpxlExpression),
				),
			},
		},
	})
}

func TestAccIbmLogsStreamsDataSourceAllArgs(t *testing.T) {
	streamName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	streamIsActive := "false"
	streamDpxlExpression := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))
	streamCompressionType := "unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamsDataSourceConfig(streamName, streamIsActive, streamDpxlExpression, streamCompressionType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "streams.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "streams.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_streams.logs_streams_instance", "streams.0.name", streamName),
					resource.TestCheckResourceAttr("data.ibm_logs_streams.logs_streams_instance", "streams.0.is_active", streamIsActive),
					resource.TestCheckResourceAttr("data.ibm_logs_streams.logs_streams_instance", "streams.0.dpxl_expression", streamDpxlExpression),
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "streams.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_streams.logs_streams_instance", "streams.0.updated_at"),
					resource.TestCheckResourceAttr("data.ibm_logs_streams.logs_streams_instance", "streams.0.compression_type", streamCompressionType),
				),
			},
		},
	})
}

func testAccCheckIbmLogsStreamsDataSourceConfigBasic(streamName string, streamDpxlExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_stream" "logs_stream_instance" {
			name = "%s"
			dpxl_expression = "%s"
		}

		data "ibm_logs_streams" "logs_streams_instance" {
			depends_on = [
				ibm_logs_stream.logs_stream_instance
			]
		}
	`, streamName, streamDpxlExpression)
}

func testAccCheckIbmLogsStreamsDataSourceConfig(streamName string, streamIsActive string, streamDpxlExpression string, streamCompressionType string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_stream" "logs_stream_instance" {
			name = "%s"
			is_active = %s
			dpxl_expression = "%s"
			compression_type = "%s"
			ibm_event_streams {
				brokers = "kafka01.example.com:9093"
				topic = "live.screen"
			}
		}

		data "ibm_logs_streams" "logs_streams_instance" {
			depends_on = [
				ibm_logs_stream.logs_stream_instance
			]
		}
	`, streamName, streamIsActive, streamDpxlExpression, streamCompressionType)
}

func TestDataSourceIbmLogsStreamsStreamToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		ibmEventStreamsModel := make(map[string]interface{})
		ibmEventStreamsModel["brokers"] = "kafka01.example.com:9093"
		ibmEventStreamsModel["topic"] = "live.screen"

		model := make(map[string]interface{})
		model["id"] = int(0)
		model["name"] = "Live Screen"
		model["is_active"] = true
		model["dpxl_expression"] = ")DPXL/1:version:1/50:payload:<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"
		model["created_at"] = "2021-01-01T00:00:00.000Z"
		model["updated_at"] = "2021-01-01T00:00:00.000Z"
		model["compression_type"] = "gzip"
		model["ibm_event_streams"] = []map[string]interface{}{ibmEventStreamsModel}

		assert.Equal(t, result, model)
	}

	ibmEventStreamsModel := new(logsv0.IbmEventStreams)
	ibmEventStreamsModel.Brokers = core.StringPtr("kafka01.example.com:9093")
	ibmEventStreamsModel.Topic = core.StringPtr("live.screen")

	model := new(logsv0.Stream)
	model.ID = core.Int64Ptr(int64(0))
	model.Name = core.StringPtr("Live Screen")
	model.IsActive = core.BoolPtr(true)
	model.DpxlExpression = core.StringPtr(")DPXL/1:version:1/50:payload:<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')")
	model.CreatedAt = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.CompressionType = core.StringPtr("gzip")
	model.IbmEventStreams = ibmEventStreamsModel

	result, err := logs.DataSourceIbmLogsStreamsStreamToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsStreamsIbmEventStreamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["brokers"] = "kafka01.example.com:9093"
		model["topic"] = "live.screen"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.IbmEventStreams)
	model.Brokers = core.StringPtr("kafka01.example.com:9093")
	model.Topic = core.StringPtr("live.screen")

	result, err := logs.DataSourceIbmLogsStreamsIbmEventStreamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
