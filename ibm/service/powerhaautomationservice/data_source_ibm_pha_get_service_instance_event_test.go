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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaGetServiceInstanceEventDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaGetServiceInstanceEventDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "event_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "action"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "message"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "resource"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_get_service_instance_event.pha_get_service_instance_event_instance", "time_stamp"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaGetServiceInstanceEventDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_get_service_instance_event" "pha_get_service_instance_event_instance" {
			pha_instance_id = "8ce2a099-a463-479a-9a1d-eedc19287a62"
			event_id = "8ce2a099-a463-479a-9a1d-eedc19287a62-1772012788623714351"
		}
	`)
}


func TestDataSourceIBMPhaGetServiceInstanceEventPhaEventUserToMap(t *testing.T) {
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

	result, err := powerhaautomationservice.DataSourceIBMPhaGetServiceInstanceEventPhaEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
