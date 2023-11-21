// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsShareProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsShareProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.0.capacity.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "profiles.0.iops.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_share_profiles.is_share_profiles", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIbmIsShareProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_share_profiles" "is_share_profiles" {
		}
	`)
}
