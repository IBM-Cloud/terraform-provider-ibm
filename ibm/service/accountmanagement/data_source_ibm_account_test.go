// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

package accountmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/accountmanagement"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmAccountDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmAccountDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "owner"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "owner_userid"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "owner_iamid"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "linked_softlayer_account"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "team_directory_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_account_info.account_instance", "traits.#"),
				),
			},
		},
	})
}

func testAccCheckIbmAccountDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_account_info" "account_instance" {
			account_id = "4904bf0fc13042878784054a1fd0f320"
		}
	`)
}

func TestDataSourceIbmAccountAccountResponseTraitsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["eu_supported"] = true
		model["poc"] = true
		model["hippa"] = true

		assert.Equal(t, result, model)
	}

	model := new(accountmanagementv4.AccountResponseTraits)
	model.EuSupported = core.BoolPtr(true)
	model.Poc = core.BoolPtr(true)
	model.Hippa = core.BoolPtr(true)

	result, err := accountmanagement.DataSourceIbmAccountAccountResponseTraitsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
