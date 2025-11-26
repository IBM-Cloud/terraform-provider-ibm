// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsOutgoingWebhooksDataSourceBasic(t *testing.T) {
	outgoingWebhookType := "ibm_event_notifications"
	outgoingWebhookName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsOutgoingWebhooksDataSourceConfigBasic(outgoingWebhookType, outgoingWebhookName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_outgoing_webhooks.logs_outgoing_webhooks_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsOutgoingWebhooksDataSourceConfigBasic(outgoingWebhookType string, outgoingWebhookName string) string {
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
		data "ibm_logs_outgoing_webhooks" "logs_outgoing_webhooks_instance" {
			instance_id              = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.instance_id
			region                   = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.region
			type					 = ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.type
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, outgoingWebhookName, outgoingWebhookType, acc.LogsEventNotificationInstanceId, acc.LogsEventNotificationInstanceRegion)
}

func TestDataSourceIbmLogsOutgoingWebhooksOutgoingWebhookSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		model["name"] = "testString"
		model["url"] = "testString"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["updated_at"] = "2019-01-01T12:00:00.000Z"
		model["external_id"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.OutgoingWebhookSummary)
	model.ID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
	model.Name = core.StringPtr("testString")
	model.URL = core.StringPtr("testString")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.ExternalID = core.Int64Ptr(int64(0))

	result, err := logs.DataSourceIbmLogsOutgoingWebhooksOutgoingWebhookSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
