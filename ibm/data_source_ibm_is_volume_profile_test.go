// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVolumeProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_volume_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVolumeProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", volumeProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeProfileDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_volume_profile" "test1" {
	name = "%s"
}`, volumeProfileName)
}
