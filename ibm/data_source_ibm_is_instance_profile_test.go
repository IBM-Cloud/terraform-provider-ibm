package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISInstanceProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(instanceProfileName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", instanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "generation"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceProfileDataSourceConfig(instanceProfileName string) string {
	return fmt.Sprintf(`

data "ibm_is_instance_profile" "test1" {
	name = "%s",
}`, instanceProfileName)
}
