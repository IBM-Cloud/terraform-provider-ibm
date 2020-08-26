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
			resource.TestStep{
				Config: testAccCheckIBMCisGLBHealthCheckDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_glb_health_check.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_glb_health_check.0.monitor_id"),
					resource.TestCheckResourceAttrSet(node, "cis_glb_health_check.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMCisGLBHealthCheckDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "web_group" {
		name = "Default"
	}
	data "ibm_cis" "web_instance" {
		name              = "CISTest"
		resource_group_id = data.ibm_resource_group.web_group.id
	}
	data "ibm_cis_glb_health_checks" "test" {
		cis_id     = data.ibm_cis.web_instance.id
	  }`)
}
