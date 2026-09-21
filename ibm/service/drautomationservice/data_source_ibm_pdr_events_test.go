// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrEventsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrEventsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_events.pdr_events_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_events.pdr_events_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_events.pdr_events_instance", "events.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrEventsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_events" "pdr_events_instance" {
			instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
			time = "2021-01-31T09:44:12Z"
			from_time = "2025-06-19T00:00:00Z"
			to_time = "2025-06-19T23:59:59Z"
			Accept-Language = "en-US"
		}
	`)
}

func TestDataSourceIBMPdrEventsEventToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		eventUserModel := make(map[string]interface{})
		eventUserModel["email"] = "testString"
		eventUserModel["name"] = "testString"
		eventUserModel["user_id"] = "testString"

		model := make(map[string]interface{})
		model["action"] = "testString"
		model["api_source"] = "testString"
		model["event_id"] = "testString"
		model["level"] = "notice"
		model["message"] = "testString"
		model["message_data"] = map[string]interface{}{"anyKey": "anyValue"}
		model["metadata"] = map[string]interface{}{"anyKey": "anyValue"}
		model["resource"] = "testString"
		model["time"] = "2019-01-01T12:00:00.000Z"
		model["timestamp"] = "1715251200"
		model["user"] = []map[string]interface{}{eventUserModel}

		assert.Equal(t, result, model)
	}

	eventUserModel := new(drautomationservicev1.EventUser)
	eventUserModel.Email = core.StringPtr("testString")
	eventUserModel.Name = core.StringPtr("testString")
	eventUserModel.UserID = core.StringPtr("testString")

	model := new(drautomationservicev1.Event)
	model.Action = core.StringPtr("testString")
	model.APISource = core.StringPtr("testString")
	model.EventID = core.StringPtr("testString")
	model.Level = core.StringPtr("notice")
	model.Message = core.StringPtr("testString")
	model.MessageData = map[string]interface{}{"anyKey": "anyValue"}
	model.Metadata = map[string]interface{}{"anyKey": "anyValue"}
	model.Resource = core.StringPtr("testString")
	model.Time = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Timestamp = core.StringPtr("1715251200")
	model.User = eventUserModel

	result, err := drautomationservice.DataSourceIBMPdrEventsEventToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrEventsEventUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["email"] = "testString"
		model["name"] = "testString"
		model["user_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.EventUser)
	model.Email = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.UserID = core.StringPtr("testString")

	result, err := drautomationservice.DataSourceIBMPdrEventsEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
