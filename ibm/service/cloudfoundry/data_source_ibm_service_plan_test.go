// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMServicePlanDataSource_basic(t *testing.T) {
	t.Skip()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
