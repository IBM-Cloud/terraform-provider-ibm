// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
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

func TestAccIbmPdrGetEventsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrGetEventsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_events.pdr_get_events_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_events.pdr_get_events_instance", "provision_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_events.pdr_get_events_instance", "event.#"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrGetEventsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_events" "pdr_get_events_instance" {
			provision_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			time = "2025-06-19T23:59:59Z"
			from_time = "2025-06-19T00:00:00Z"
			to_time = "2025-06-19T23:59:59Z"
		}
	`)
}

func TestDataSourceIbmPdrGetEventsEventToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		eventUserModel := make(map[string]interface{})
		eventUserModel["email"] = "abcuser@ibm.com"
		eventUserModel["name"] = "abcuser"
		eventUserModel["user_id"] = "IBMid-695000abc7E"

		model := make(map[string]interface{})
		model["action"] = "create"
		model["api_source"] = "dr-automation-api"
		model["event_id"] = "1cecfe43-43cd-4b1b-86be-30c2d3d2a25f"
		model["level"] = "info"
		model["message"] = "Service Instance created successfully"
		model["message_data"] = map[string]interface{}{"anyKey": "anyValue"}
		model["metadata"] = map[string]interface{}{"anyKey": "anyValue"}
		model["resource"] = "ProvisionID"
		model["time"] = "2025-06-23T07:12:49.840Z"
		model["timestamp"] = "1750662769"
		model["user"] = []map[string]interface{}{eventUserModel}

		assert.Equal(t, result, model)
	}

	eventUserModel := new(drautomationservicev1.EventUser)
	eventUserModel.Email = core.StringPtr("abcuser@ibm.com")
	eventUserModel.Name = core.StringPtr("abcuser")
	eventUserModel.UserID = core.StringPtr("IBMid-695000abc7E")

	model := new(drautomationservicev1.Event)
	model.Action = core.StringPtr("create")
	model.APISource = core.StringPtr("dr-automation-api")
	model.EventID = core.StringPtr("1cecfe43-43cd-4b1b-86be-30c2d3d2a25f")
	model.Level = core.StringPtr("info")
	model.Message = core.StringPtr("Service Instance created successfully")
	model.MessageData = map[string]interface{}{"anyKey": "anyValue"}
	model.Metadata = map[string]interface{}{"anyKey": "anyValue"}
	model.Resource = core.StringPtr("ProvisionID")
	model.Time = CreateMockDateTime("2025-06-23T07:12:49.840Z")
	model.Timestamp = core.StringPtr("1750662769")
	model.User = eventUserModel

	result, err := drautomationservice.DataSourceIbmPdrGetEventsEventToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmPdrGetEventsEventUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["email"] = "abcuser@ibm.com"
		model["name"] = "abcuser"
		model["user_id"] = "IBMid-695000abc7E"

		assert.Equal(t, result, model)
	}

	model := new(drautomationservicev1.EventUser)
	model.Email = core.StringPtr("abcuser@ibm.com")
	model.Name = core.StringPtr("abcuser")
	model.UserID = core.StringPtr("IBMid-695000abc7E")

	result, err := drautomationservice.DataSourceIbmPdrGetEventsEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
