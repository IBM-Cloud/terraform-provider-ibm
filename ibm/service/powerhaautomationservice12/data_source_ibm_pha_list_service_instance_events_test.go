// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaListServiceInstanceEventsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaListServiceInstanceEventsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_list_service_instance_events.pha_list_service_instance_events_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_list_service_instance_events.pha_list_service_instance_events_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_list_service_instance_events.pha_list_service_instance_events_instance", "events.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaListServiceInstanceEventsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_list_service_instance_events" "pha_list_service_instance_events_instance" {
			pha_instance_id = "8ce2a099-a463-479a-9a1d-eedc19287a62"
			from_time = "2025-06-19T00:00:00Z"
		}
	`)
}

func TestDataSourceIBMPhaListServiceInstanceEventsPhaEventToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		phaEventUserModel := make(map[string]interface{})
		phaEventUserModel["email"] = "abcuser@ibm.com"
		phaEventUserModel["name"] = "abcuser"
		phaEventUserModel["user_id"] = "IBMid-695000abc7E"

		model := make(map[string]interface{})
		model["action"] = "create"
		model["api_source"] = "pha-automation-api"
		model["event_id"] = "d20d80f7-c87c-4ab7-bea2-2bc4736e7c2e"
		model["level"] = "info"
		model["message"] = "Service instance created successfully"
		model["message_data"] = map[string]interface{}{"anyKey": "anyValue"}
		model["meta_data"] = map[string]interface{}{"key1": "testString"}
		model["resource"] = "pha_instance_id"
		model["time"] = "2025-06-23T07:12:49.840Z"
		model["time_stamp"] = "2026-01-10T08:15:30Z"
		model["user"] = []map[string]interface{}{phaEventUserModel}

		assert.Equal(t, result, model)
	}

	phaEventUserModel := new(powerhaautomationservicev1.PhaEventUser)
	phaEventUserModel.Email = core.StringPtr("abcuser@ibm.com")
	phaEventUserModel.Name = core.StringPtr("abcuser")
	phaEventUserModel.UserID = core.StringPtr("IBMid-695000abc7E")

	model := new(powerhaautomationservicev1.PhaEvent)
	model.Action = core.StringPtr("create")
	model.APISource = core.StringPtr("pha-automation-api")
	model.EventID = core.StringPtr("d20d80f7-c87c-4ab7-bea2-2bc4736e7c2e")
	model.Level = core.StringPtr("info")
	model.Message = core.StringPtr("Service instance created successfully")
	model.MessageData = map[string]interface{}{"anyKey": "anyValue"}
	model.MetaData = map[string]string{"key1": "testString"}
	model.Resource = core.StringPtr("pha_instance_id")
	model.Time = core.StringPtr("2025-06-23T07:12:49.840Z")
	model.TimeStamp = core.StringPtr("2026-01-10T08:15:30Z")
	model.User = phaEventUserModel

	result, err := powerhaautomationservice.DataSourceIBMPhaListServiceInstanceEventsPhaEventToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMPhaListServiceInstanceEventsPhaEventUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["email"] = "abcuser@ibm.com"
		model["name"] = "abcuser"
		model["user_id"] = "IBMid-695000abc7E"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.PhaEventUser)
	model.Email = core.StringPtr("abcuser@ibm.com")
	model.Name = core.StringPtr("abcuser")
	model.UserID = core.StringPtr("IBMid-695000abc7E")

	result, err := powerhaautomationservice.DataSourceIBMPhaListServiceInstanceEventsPhaEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
