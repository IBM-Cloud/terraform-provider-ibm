// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmIsShareProfileDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareProfileDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "capacity.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "capacity.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "iops.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profile.is_share_profile", "iops.0.type"),
				),
			},
		},
	})
}

func testAccCheckIbmIsShareProfileDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_share_profile" "is_share_profile" {
			name = "%s"
		}
	`, acc.ShareProfileName)
}
