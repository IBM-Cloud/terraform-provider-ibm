/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMIAMUserProfileDataSource_Basic(t *testing.T) {
	t.Skip()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMUserProfileDataSourceConfig(),
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

func testAccCheckIBMIAMUserProfileDataSourceConfig() string {
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
