package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMIAMUserSettingsDataSource_Basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserSettingsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_user_profile.user_profile", "allowed_ip_addresses.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_iam_user_profile.user_profile", "state", "ACTIVE"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_profile.user_profile", "firstname"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_user_profile.user_profile", "phonenumber"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMUserSettingsDataSourceConfig() string {
	return fmt.Sprintf(`

	resource "ibm_iam_user_settings" "user_settings" {
		iam_id = "%s"
		allowed_ip_addresses = ["192.168.0.0","192.168.0.1"]
	  }
	  
	  data "ibm_iam_user_profile" "user_profile" {
		iam_id = ibm_iam_user_settings.user_settings.iam_id
	  }
	  
`, IAMUser)

}
