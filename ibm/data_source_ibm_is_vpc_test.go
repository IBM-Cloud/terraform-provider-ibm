package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVPCDatasource_basic(t *testing.T) {
	var vpc string
	name := fmt.Sprintf("tfc-vpc-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISVPCConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", vpc),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "name", name),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "tags.#", "1"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "cse_source_addresses.#", "3"),
				),
			},
		},
	})
}

func testDSCheckIBMISVPCConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
			tags = ["tag1"]
		}
		data "ibm_is_vpc" "ds_vpc" {
		    name = "${ibm_is_vpc.testacc_vpc.name}"
		}`, name)
}
