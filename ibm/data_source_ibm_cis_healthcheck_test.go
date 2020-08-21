package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCISGLBHealthCheckDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_glb_health_checks.test"
	rname := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMNetworkCISGLBHealthCheckDataSourceConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_glb_health_check.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_glb_health_check.0.monitor_id"),
					resource.TestCheckResourceAttrSet(node, "cis_glb_health_check.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMNetworkCISGLBHealthCheckDataSourceConfig(crn string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_cis_glb_health_checks" "test" {
		crn     = %s
	  }`, crn)
}
