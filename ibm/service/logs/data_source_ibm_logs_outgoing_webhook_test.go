// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsOutgoingWebhookDataSourceBasic(t *testing.T) {
	outgoingWebhookType := "ibm_event_notifications"
	outgoingWebhookName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	outgoingWebhookURL := fmt.Sprintf("tf_url_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsOutgoingWebhookDataSourceConfigBasic(outgoingWebhookType, outgoingWebhookName, outgoingWebhookURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "logs_outgoing_webhook_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "name"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "external_id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsOutgoingWebhookDataSourceConfigBasic(outgoingWebhookType string, outgoingWebhookName string, outgoingWebhookURL string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
		instance_id            = "%s"
		region                 = "%s"
		name                   = "%s"
		type                   = "%s"
		ibm_event_notifications {
		  event_notifications_instance_id = "%s"
		  region_id                       = "%s"
		}
	  }

		data "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
			instance_id              = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.instance_id
			region                   = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.region
			logs_outgoing_webhook_id =ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.webhook_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, outgoingWebhookName, outgoingWebhookType, acc.LogsEventNotificationInstanceId, acc.LogsEventNotificationInstanceRegion)
}

func TestDataSourceIbmLogsOutgoingWebhookOutgoingWebhooksV1IbmEventNotificationsConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["event_notifications_instance_id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		model["region_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig)
	model.EventNotificationsInstanceID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
	model.RegionID = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsOutgoingWebhookOutgoingWebhooksV1IbmEventNotificationsConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
