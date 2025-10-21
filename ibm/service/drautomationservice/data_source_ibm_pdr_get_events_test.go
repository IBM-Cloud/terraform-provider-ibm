// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMPdrGetEventsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetEventsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_events.pdr_get_events_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_events.pdr_get_events_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_events.pdr_get_events_instance", "event.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPdrGetEventsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_events" "pdr_get_events_instance" {
			instance_id = "xxxx2ec4-xxxx-4f84-xxxx-c2aa834dd4ed"
			time = "2025-06-19T23:59:59Z"
			from_time = "2025-06-19T00:00:00Z"
			to_time = "2025-06-19T23:59:59Z"
			Accept-Language = "it"
		}
	`)
}

func TestDataSourceIBMPdrGetEventsEventToMap(t *testing.T) {
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

	result, err := drautomationservice.DataSourceIBMPdrGetEventsEventToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPdrGetEventsEventUserToMap(t *testing.T) {
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

	result, err := drautomationservice.DataSourceIBMPdrGetEventsEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
