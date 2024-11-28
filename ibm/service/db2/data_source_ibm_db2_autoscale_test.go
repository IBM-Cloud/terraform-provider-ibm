package db2

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIBMDb2AutoscaleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
		},
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckDb2AutoscaleDataSourceBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "auto_scaling_allow_plan_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "auto_scaling_enable"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "auto_scaling_max_storage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "auto_scaling_over_time_period"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "auto_scaling_pause_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "auto_scaling_threshold"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "storage_unit"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "storage_utilization_percentage"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_autoscale.instance_autoscale", "support_auto_scaling"),
				),
			},
		},
	})
}

func testAccCheckDb2AutoscaleDataSourceBasic() string {
	return fmt.Sprintf(`
data "ibm_db2_autoscale" "test" {
auto_scaling_allow_plan_limit = "%[1]t"
auto_scaling_enabled = "%[1]f"
auto_scaling_max_storage = "%[1]d"
auto_scaling_over_time_period = "%[1]d"
auto_scaling_pause_limit = "%[1]d"
auto_scaling_threshold = "%[1]d"
storage_unit = "%[1]s"
storage_utilization_percentage = "%[1]d"
support_auto_scaling = "%[1]t"
}
`)
}
