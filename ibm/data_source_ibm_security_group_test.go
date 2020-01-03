package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/softlayer/softlayer-go/datatypes"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccIBMSecurityGroupDataSource_basic(t *testing.T) {
	var sg datatypes.Network_SecurityGroup

	name1 := fmt.Sprintf("terraformsguat_create_step_name_%d", acctest.RandInt())
	desc1 := fmt.Sprintf("terraformsguat_create_step_desc_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSecurityGroupDataSourceConfig(name1, desc1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMSecurityGroupExists("ibm_security_group.testacc_security_group", &sg),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_security_group", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_security_group.testacc_security_group", "description", desc1),
					resource.TestCheckResourceAttr(
						"data.ibm_security_group.tfsg", "name", name1),
					resource.TestCheckResourceAttr(
						"data.ibm_security_group.tfsg", "description", desc1),
				),
			},
		},
	})
}

func testAccCheckIBMSecurityGroupDataSourceConfig(name, description string) string {
	return fmt.Sprintf(`
data "ibm_security_group" "tfsg"{
	name = ibm_security_group.testacc_security_group.name
}
resource "ibm_security_group" "testacc_security_group" {
    name = "%s"
    description = "%s"
}`, name, description)

}
