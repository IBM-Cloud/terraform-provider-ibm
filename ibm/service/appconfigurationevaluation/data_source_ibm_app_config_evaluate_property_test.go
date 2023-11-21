// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfigurationevaluation_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppConfigEvaluatePropertyDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppConfigEvaluatePropertyDataSourceBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_evaluate_property.evaluate_property", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppConfigEvaluatePropertyDataSourceBasic() string {
	return fmt.Sprintf(`
		data "ibm_app_config_evaluate_property" "evaluate_property" {
			guid = "36401ffc-6280-459a-ba98-456aba10d0c7"
			environment_id = "dev"
			collection_id = "car-rentals"
			property_id = "users-location"
			entity_id = "john_doe"
			entity_attributes = {
				"city" : "Bangalore",
				"radius" : 60
			}
		}
	`)
}
