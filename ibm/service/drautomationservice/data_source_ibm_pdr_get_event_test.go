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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIbmPdrGetEventDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmPdrGetEventDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "provision_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "event_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "action"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "level"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "message"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "resource"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "time"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "timestamp"),
				),
			},
		},
	})
}

func testAccCheckIbmPdrGetEventDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_event" "pdr_get_event_instance" {
			provision_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
			event_id = "00116b2a-9326-4024-839e-fb5364b76898"
			Accept-Language = "Accept-Language"
			If-None-Match = "If-None-Match"
		}
	`)
}

func TestDataSourceIbmPdrGetEventEventUserToMap(t *testing.T) {
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

	result, err := drautomationservice.DataSourceIbmPdrGetEventEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
