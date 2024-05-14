// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsOutgoingWebhookBasic(t *testing.T) {
	var conf logsv0.OutgoingWebhook
	typeVar := "ibm_event_notifications"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "ibm_event_notifications"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsOutgoingWebhookDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsOutgoingWebhookConfigBasic(typeVar, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsOutgoingWebhookExists("ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsOutgoingWebhookConfigBasic(typeVarUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsOutgoingWebhookConfigBasic(typeVar string, name string) string {
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, typeVar, acc.LogsEventNotificationInstanceId, acc.LogsEventNotificationInstanceRegion)
}

func testAccCheckIbmLogsOutgoingWebhookExists(n string, obj logsv0.OutgoingWebhook) resource.TestCheckFunc {

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
		getOutgoingWebhookOptions := &logsv0.GetOutgoingWebhookOptions{}

		getOutgoingWebhookOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		outgoingWebhookIntf, _, err := logsClient.GetOutgoingWebhook(getOutgoingWebhookOptions)
		if err != nil {
			return err
		}

		outgoingWebhook := outgoingWebhookIntf.(*logsv0.OutgoingWebhook)
		obj = *outgoingWebhook
		return nil
	}
}

func testAccCheckIbmLogsOutgoingWebhookDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_outgoing_webhook" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		getOutgoingWebhookOptions := &logsv0.GetOutgoingWebhookOptions{}

		getOutgoingWebhookOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetOutgoingWebhook(getOutgoingWebhookOptions)

		if err == nil {
			return fmt.Errorf("logs_outgoing_webhook still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_outgoing_webhook (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmLogsOutgoingWebhookOutgoingWebhooksV1IbmEventNotificationsConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["event_notifications_instance_id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		model["region_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig)
	model.EventNotificationsInstanceID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
	model.RegionID = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsOutgoingWebhookOutgoingWebhooksV1IbmEventNotificationsConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsOutgoingWebhookMapToOutgoingWebhooksV1IbmEventNotificationsConfig(t *testing.T) {
	checkResult := func(result *logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig) {
		model := new(logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig)
		model.EventNotificationsInstanceID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
		model.RegionID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["event_notifications_instance_id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
	model["region_id"] = "testString"

	result, err := logs.ResourceIbmLogsOutgoingWebhookMapToOutgoingWebhooksV1IbmEventNotificationsConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsOutgoingWebhookMapToOutgoingWebhookPrototype(t *testing.T) {
	checkResult := func(result logsv0.OutgoingWebhookPrototypeIntf) {
		outgoingWebhooksV1IbmEventNotificationsConfigModel := new(logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig)
		outgoingWebhooksV1IbmEventNotificationsConfigModel.EventNotificationsInstanceID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
		outgoingWebhooksV1IbmEventNotificationsConfigModel.RegionID = core.StringPtr("testString")

		model := new(logsv0.OutgoingWebhookPrototype)
		model.Type = core.StringPtr("ibm_event_notifications")
		model.Name = core.StringPtr("testString")
		model.URL = core.StringPtr("testString")
		model.IbmEventNotifications = outgoingWebhooksV1IbmEventNotificationsConfigModel

		assert.Equal(t, result, model)
	}

	outgoingWebhooksV1IbmEventNotificationsConfigModel := make(map[string]interface{})
	outgoingWebhooksV1IbmEventNotificationsConfigModel["event_notifications_instance_id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
	outgoingWebhooksV1IbmEventNotificationsConfigModel["region_id"] = "testString"

	model := make(map[string]interface{})
	model["type"] = "ibm_event_notifications"
	model["name"] = "testString"
	model["url"] = "testString"
	model["ibm_event_notifications"] = []interface{}{outgoingWebhooksV1IbmEventNotificationsConfigModel}

	result, err := logs.ResourceIbmLogsOutgoingWebhookMapToOutgoingWebhookPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsOutgoingWebhookMapToOutgoingWebhookPrototypeOutgoingWebhooksV1OutgoingWebhookInputDataConfigIbmEventNotifications(t *testing.T) {
	checkResult := func(result *logsv0.OutgoingWebhookPrototypeOutgoingWebhooksV1OutgoingWebhookInputDataConfigIbmEventNotifications) {
		outgoingWebhooksV1IbmEventNotificationsConfigModel := new(logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig)
		outgoingWebhooksV1IbmEventNotificationsConfigModel.EventNotificationsInstanceID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
		outgoingWebhooksV1IbmEventNotificationsConfigModel.RegionID = core.StringPtr("testString")

		model := new(logsv0.OutgoingWebhookPrototypeOutgoingWebhooksV1OutgoingWebhookInputDataConfigIbmEventNotifications)
		model.Type = core.StringPtr("ibm_event_notifications")
		model.Name = core.StringPtr("testString")
		model.URL = core.StringPtr("testString")
		model.IbmEventNotifications = outgoingWebhooksV1IbmEventNotificationsConfigModel

		assert.Equal(t, result, model)
	}

	outgoingWebhooksV1IbmEventNotificationsConfigModel := make(map[string]interface{})
	outgoingWebhooksV1IbmEventNotificationsConfigModel["event_notifications_instance_id"] = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
	outgoingWebhooksV1IbmEventNotificationsConfigModel["region_id"] = "testString"

	model := make(map[string]interface{})
	model["type"] = "ibm_event_notifications"
	model["name"] = "testString"
	model["url"] = "testString"
	model["ibm_event_notifications"] = []interface{}{outgoingWebhooksV1IbmEventNotificationsConfigModel}

	result, err := logs.ResourceIbmLogsOutgoingWebhookMapToOutgoingWebhookPrototypeOutgoingWebhooksV1OutgoingWebhookInputDataConfigIbmEventNotifications(model)
	assert.Nil(t, err)
	checkResult(result)
}
