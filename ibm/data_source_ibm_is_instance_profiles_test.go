package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISInstanceProfilesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.generation"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceProfilesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_instance_profiles" "test1" {
      }`)
}
