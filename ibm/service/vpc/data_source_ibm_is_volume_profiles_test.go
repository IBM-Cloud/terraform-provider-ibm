// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVolumeProfilesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_volume_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
				),
			},
		},
	})
}
func TestAccIBMISVolumeProfilesDataSource_Sdp(t *testing.T) {
	resName := "data.ibm_is_volume_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.adjustable_capacity_states.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.adjustable_iops_states.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.boot_capacity.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.capacity.#"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.href"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.iops.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeProfilesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_volume_profiles" "test1" {
      }`)
}
