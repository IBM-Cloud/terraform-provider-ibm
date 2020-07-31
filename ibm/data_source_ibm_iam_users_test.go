package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMIAMUsersDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUsersDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_users.users_profiles", "users.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUsersDataSourceConfig() string {
	return fmt.Sprintf(`

	data "ibm_iam_users" "users_profiles"{
  
	} 
`)

}
