// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmDb2AutoscaleDataSourceBasic(t *testing.T) {
	xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-east%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A8e3a219f-65d3-43cd-86da-b231d53732ef%3A%3A"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2AutoscaleDataSourceConfigBasic(xDbProfile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "auto_scaling_allow_plan_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "auto_scaling_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "auto_scaling_max_storage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "auto_scaling_over_time_period"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "auto_scaling_pause_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "auto_scaling_threshold"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "storage_unit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "storage_utilization_percentage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.Db2-v0-test-public", "support_auto_scaling"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2AutoscaleDataSourceConfigBasic(xDbProfile string) string {
	return fmt.Sprintf(`
		 data "ibm_db2_autoscale" "Db2-v0-test-public" {
    deployment_id = "%[1]s"
}
	`, xDbProfile)
}
