package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISVPCDatasource_basic(t *testing.T) {
	var vpc *models.Vpc
	name1 := fmt.Sprintf("terraformvpcuat_create_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISVPCConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCExists("ibm_is_vpc.testacc_vpc", &vpc),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "name", name1),
					resource.TestCheckResourceAttr(
						"data.ibm_is_vpc.ds_vpc", "tags.#", "1"),
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
