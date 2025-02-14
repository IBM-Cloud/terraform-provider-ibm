// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfigurationevaluation_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppConfigEvaluateFeatureFlagDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppConfigEvaluateFeatureFlagDataSourceBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_evaluate_feature_flag.evaluate_feature_flag", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppConfigEvaluateFeatureFlagDataSourceBasic() string {
	return fmt.Sprintf(`
		data "ibm_app_config_evaluate_feature_flag" "evaluate_feature_flag" {
			guid = "36401ffc-6280-459a-ba98-456aba10d0c7"
			environment_id = "dev"
			collection_id = "car-rentals"
			feature_id = "weekend-discount"
			entity_id = "john_doe"
			entity_attributes = {
				"city" : "Bangalore",
				"radius" : 60
			}
		}
	`)
}
