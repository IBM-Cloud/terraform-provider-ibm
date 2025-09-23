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

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsStreamsDataSourceBasic(t *testing.T) {
	streamName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	streamDpxlExpression := "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
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
	streamDpxlExpression := "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"
	streamCompressionType := "gzip"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
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
			instance_id      = "%s"
			region           = "%s"
			name = "%s"
			dpxl_expression = "%s"
			compression_type = "gzip"
			ibm_event_streams {
				brokers = "kafka01.example.com:9093"
				topic   = "live.screen"
			}
		}

		data "ibm_logs_streams" "logs_streams_instance" {
			instance_id      = "%[1]s"
			region           = "%[2]s"
			depends_on = [
				ibm_logs_stream.logs_stream_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, streamName, streamDpxlExpression)
}

func testAccCheckIbmLogsStreamsDataSourceConfig(streamName string, streamIsActive string, streamDpxlExpression string, streamCompressionType string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_stream" "logs_stream_instance" {
			instance_id      = "%s"
			region           = "%s"
			name             = "%s"
			is_active        = %s
			dpxl_expression  = "%s"
			compression_type = "%s"
			ibm_event_streams {
				brokers = "kafka01.example.com:9093"
				topic   = "live.screen"
			}
		}

		data "ibm_logs_streams" "logs_streams_instance" {
			instance_id      = "%[1]s"
			region           = "%[2]s"
			depends_on = [
				ibm_logs_stream.logs_stream_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, streamName, streamIsActive, streamDpxlExpression, streamCompressionType)
}
