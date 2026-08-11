// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package restapi_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIBMRestApiDataBasic reads the resource groups of the test account
// through the generic data source.
func TestAccIBMRestApiDataBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRestApiDataConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_restapi_data.restapi_data_instance", "response_status_code", "200"),
					resource.TestCheckResourceAttr("data.ibm_restapi_data.restapi_data_instance", "is_json", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_restapi_data.restapi_data_instance", "response_body"),
					resource.TestCheckResourceAttrSet("data.ibm_restapi_data.restapi_data_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_restapi_data.restapi_data_instance", "response_headers.%"),
				),
			},
		},
	})
}

// TestAccIBMRestApiDataQueryParams checks that query_params reach the API and
// that a non 2xx status code listed in accept_status_codes is tolerated.
func TestAccIBMRestApiDataQueryParams(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRestApiDataConfigQueryParams(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_restapi_data.restapi_data_filtered", "response_status_code", "200"),
					resource.TestCheckResourceAttr("data.ibm_restapi_data.restapi_data_missing", "response_status_code", "404"),
				),
			},
		},
	})
}

func testAccCheckIBMRestApiDataConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_restapi_data" "restapi_data_instance" {
			url = "%s"
			query_params = {
				account_id = "%s"
			}
		}
	`, resourceGroupsURL, acc.IAMAccountId)
}

func testAccCheckIBMRestApiDataConfigQueryParams() string {
	return fmt.Sprintf(`
		data "ibm_restapi_data" "restapi_data_filtered" {
			url = "%[1]s"
			query_params = {
				account_id = "%[2]s"
				default    = "true"
			}
		}

		data "ibm_restapi_data" "restapi_data_missing" {
			url                 = "%[1]s/00000000000000000000000000000000"
			accept_status_codes = [404]
		}
	`, resourceGroupsURL, acc.IAMAccountId)
}
