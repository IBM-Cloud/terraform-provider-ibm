// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVolumeProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_volume_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.VolumeProfileName),
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
}`, acc.VolumeProfileName)
}
