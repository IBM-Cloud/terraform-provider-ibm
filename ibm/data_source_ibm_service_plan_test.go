// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMServicePlanDataSource_basic(t *testing.T) {
	t.Skip()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMServicePlanDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_service_plan.testacc_ds_service_plan", "service", "cloudantNoSQLDB"),
					resource.TestCheckResourceAttr("data.ibm_service_plan.testacc_ds_service_plan", "plan", "Lite"),
				),
			},
		},
	})
}

func testAccCheckIBMServicePlanDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_service_plan" "testacc_ds_service_plan" {
    service = "cloudantNoSQLDB"
	plan = "Lite"
}`,
	)

}
