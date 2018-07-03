package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMServicePlanDataSource_basic(t *testing.T) {

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
