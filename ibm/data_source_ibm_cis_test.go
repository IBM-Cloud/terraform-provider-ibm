package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMCisDataSource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  setupCisConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis.instance", "service", "internet-svcs"),
				),
			},
			resource.TestStep{
				Config:  testAccCheckIBMCisDataSourceConfig(instanceName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "name", instanceName),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "service", "internet-svcs"),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "plan", "standard"),
					resource.TestCheckResourceAttr("data.ibm_cis.testacc_ds_cis", "location", "global"),
				),
			},
		},
	})
}

func setupCisConfig(instanceName string) string {
	return fmt.Sprintf(`

resource "ibm_cis" "instance" {
  name       = "%s"
  plan       = "standard"
  location   = "global"
}`, instanceName)

}

func testAccCheckIBMCisDataSourceConfig(instanceName string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_cis" "instance" {
	name       = "%s"
	plan       = "standard"
	location   = "global"
}

data "ibm_cis" "testacc_ds_cis" {
  name = "${ibm_cis.instance.name}"
}
`, instanceName)

}
