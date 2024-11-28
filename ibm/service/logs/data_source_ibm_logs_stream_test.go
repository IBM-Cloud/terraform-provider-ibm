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
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsStreamDataSourceBasic(t *testing.T) {
	streamName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	streamDpxlExpression := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamDataSourceConfigBasic(streamName, streamDpxlExpression),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "dpxl_expression"),
				),
			},
		},
	})
}

func TestAccIbmLogsStreamDataSourceAllArgs(t *testing.T) {
	streamName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	streamIsActive := "false"
	streamDpxlExpression := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))
	streamCompressionType := "unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamDataSourceConfig(streamName, streamIsActive, streamDpxlExpression, streamCompressionType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "is_active"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "dpxl_expression"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "compression_type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_stream.logs_stream_instance", "ibm_event_streams.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsStreamDataSourceConfigBasic(streamName string, streamDpxlExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_stream" "logs_stream_instance" {
			name = "%s"
			dpxl_expression = "%s"
		}

		data "ibm_logs_stream" "logs_stream_instance" {
			depends_on = [
				ibm_logs_stream.logs_stream_instance
			]
		}
	`, streamName, streamDpxlExpression)
}

func testAccCheckIbmLogsStreamDataSourceConfig(streamName string, streamIsActive string, streamDpxlExpression string, streamCompressionType string) string {
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

		data "ibm_logs_stream" "logs_stream_instance" {
			depends_on = [
				ibm_logs_stream.logs_stream_instance
			]
		}
	`, streamName, streamIsActive, streamDpxlExpression, streamCompressionType)
}

func TestDataSourceIbmLogsStreamIbmEventStreamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["brokers"] = "kafka01.example.com:9093"
		model["topic"] = "live.screen"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.IbmEventStreams)
	model.Brokers = core.StringPtr("kafka01.example.com:9093")
	model.Topic = core.StringPtr("live.screen")

	result, err := logs.DataSourceIbmLogsStreamIbmEventStreamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
