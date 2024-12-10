// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmDb2SaasAutoscaleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2SaasAutoscaleDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "x_db_profile"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "auto_scaling_allow_plan_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "auto_scaling_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "auto_scaling_max_storage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "auto_scaling_over_time_period"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "auto_scaling_pause_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "auto_scaling_threshold"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "storage_unit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "storage_utilization_percentage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_autoscale.db2_saas_autoscale_instance", "support_auto_scaling"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2SaasAutoscaleDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		 data "ibm_db2_saas_autoscale" "db2_autoscale" {
    x_db_profile = "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
}
	`)
}
