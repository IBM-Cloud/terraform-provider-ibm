// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/drautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrGetEventDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrGetEventDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pdr_get_event.pdr_get_event_instance", "instance_id"),
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

func testAccCheckIBMPdrGetEventDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pdr_get_event" "pdr_get_event_instance" {
			instance_id = "crn:v1:staging:public:power-dr-automation:global:a/b68c234e719144b18598ae4a7b80c44c:492fef47-3ebf-4090-b089-e9b4199878b6::"
			event_id = "492fef47-3ebf-4090-b089-e9b4199878b6-1764568221673678123"
		}
	`)
}

func TestDataSourceIBMPdrGetEventEventUserToMap(t *testing.T) {
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

	result, err := drautomationservice.DataSourceIBMPdrGetEventEventUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
