// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsStreamBasic(t *testing.T) {
	var conf logsv0.Stream
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dpxlExpression := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	dpxlExpressionUpdate := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
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
	dpxlExpression := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))
	compressionType := "unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	isActiveUpdate := "true"
	dpxlExpressionUpdate := fmt.Sprintf("tf_dpxl_expression_%d", acctest.RandIntRange(10, 100))
	compressionTypeUpdate := "gzip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
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
				ResourceName:      "ibm_logs_stream.logs_stream",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsStreamConfigBasic(name string, dpxlExpression string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_stream" "logs_stream_instance" {
			name = "%s"
			dpxl_expression = "%s"
		}
	`, name, dpxlExpression)
}

func testAccCheckIbmLogsStreamConfig(name string, isActive string, dpxlExpression string, compressionType string) string {
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
	`, name, isActive, dpxlExpression, compressionType)
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

		getEventStreamTargetsOptions := &logsv0.GetEventStreamTargetsOptions{}

		getEventStreamTargetsOptions.SetID(rs.Primary.ID)

		stream, _, err := logsClient.GetEventStreamTargets(getEventStreamTargetsOptions)
		if err != nil {
			return err
		}

		obj = *stream
		return nil
	}
}

func testAccCheckIbmLogsStreamDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_stream" {
			continue
		}

		getEventStreamTargetsOptions := &logsv0.GetEventStreamTargetsOptions{}

		getEventStreamTargetsOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := logsClient.GetEventStreamTargets(getEventStreamTargetsOptions)

		if err == nil {
			return fmt.Errorf("logs_stream still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_stream (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsStreamIbmEventStreamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["brokers"] = "kafka01.example.com:9093"
		model["topic"] = "live.screen"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.IbmEventStreams)
	model.Brokers = core.StringPtr("kafka01.example.com:9093")
	model.Topic = core.StringPtr("live.screen")

	result, err := logs.ResourceIbmLogsStreamIbmEventStreamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsStreamMapToIbmEventStreams(t *testing.T) {
	checkResult := func(result *logsv0.IbmEventStreams) {
		model := new(logsv0.IbmEventStreams)
		model.Brokers = core.StringPtr("kafka01.example.com:9093")
		model.Topic = core.StringPtr("live.screen")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["brokers"] = "kafka01.example.com:9093"
	model["topic"] = "live.screen"

	result, err := logs.ResourceIbmLogsStreamMapToIbmEventStreams(model)
	assert.Nil(t, err)
	checkResult(result)
}
