// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsStreamBasic(t *testing.T) {
	var conf logsv0.Stream
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dpxlExpression := "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dpxlExpressionUpdate := "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-2')"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsStreamDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamConfigBasic(name, dpxlExpression),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsStreamExists("ibm_logs_stream.logs_stream_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "dpxl_expression", dpxlExpression),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamConfigBasic(nameUpdate, dpxlExpressionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "dpxl_expression", dpxlExpressionUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsStreamAllArgs(t *testing.T) {
	var conf logsv0.Stream
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	isActive := "false"
	dpxlExpression := "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"
	compressionType := "gzip"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	isActiveUpdate := "true"
	dpxlExpressionUpdate := "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-2')"
	compressionTypeUpdate := "gzip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsStreamDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamConfig(name, isActive, dpxlExpression, compressionType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsStreamExists("ibm_logs_stream.logs_stream_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "is_active", isActive),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "dpxl_expression", dpxlExpression),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "compression_type", compressionType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsStreamConfig(nameUpdate, isActiveUpdate, dpxlExpressionUpdate, compressionTypeUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "is_active", isActiveUpdate),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "dpxl_expression", dpxlExpressionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_stream.logs_stream_instance", "compression_type", compressionTypeUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_stream.logs_stream_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsStreamConfigBasic(name string, dpxlExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_stream" "logs_stream_instance" {
			instance_id      = "%s"
			region           = "%s"
			name 			 = "%s"
			dpxl_expression  = "%s"
			compression_type = "gzip"
			ibm_event_streams {
				brokers = "kafka01.example.com:9093"
				topic   = "live.screen"
			}
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, dpxlExpression)
}

func testAccCheckIbmLogsStreamConfig(name string, isActive string, dpxlExpression string, compressionType string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, isActive, dpxlExpression, compressionType)
}

func testAccCheckIbmLogsStreamExists(n string, obj logsv0.Stream) resource.TestCheckFunc {

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
		streamsIDInt, _ := strconv.ParseInt(resourceID[2], 10, 64)

		getEventStreamTargetsOptions := &logsv0.GetEventStreamTargetsOptions{}

		stream, _, err := logsClient.GetEventStreamTargets(getEventStreamTargetsOptions)
		if err != nil {
			return err
		}
		for _, stream := range stream.Streams {
			if stream.ID == &streamsIDInt {
				obj = stream
				return nil
			}
		}
		return nil
	}
}

func testAccCheckIbmLogsStreamDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_stream" {
			continue
		}

		getEventStreamTargetsOptions := &logsv0.GetEventStreamTargetsOptions{}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		streamsIDInt, _ := strconv.ParseInt(resourceID[2], 10, 64)
		// Try to find the key
		streams, _, _ := logsClient.GetEventStreamTargets(getEventStreamTargetsOptions)
		for _, stream := range streams.Streams {
			if stream.ID == &streamsIDInt {
				return fmt.Errorf("logs_streams still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
