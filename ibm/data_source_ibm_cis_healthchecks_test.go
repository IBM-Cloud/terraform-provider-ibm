package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisGLBHealthCheckDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_healthchecks.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisGLBHealthCheckDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_healthchecks.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_healthchecks.0.monitor_id"),
					resource.TestCheckResourceAttrSet(node, "cis_healthchecks.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMCisGLBHealthCheckDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_healthchecks" "test" {
		cis_id     = data.ibm_cis.cis.id
	  }`)
}
