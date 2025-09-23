// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMUserProfileDataSource_Basic(t *testing.T) {
	t.Skip()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
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
	  
`, acc.IAMUser)

}
