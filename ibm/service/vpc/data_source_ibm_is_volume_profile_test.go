// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

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
func TestAccIBMISVolumeProfileDataSource_sdpbasic(t *testing.T) {
	resName := "data.ibm_is_volume_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeProfileDataSourceSdpConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", "sdp"),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "adjustable_capacity_states.#"),
					resource.TestCheckResourceAttrSet(resName, "adjustable_capacity_states.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "adjustable_iops_states.#"),
					resource.TestCheckResourceAttrSet(resName, "adjustable_iops_states.0.values.#"),
					resource.TestCheckResourceAttrSet(resName, "boot_capacity.#"),
					resource.TestCheckResourceAttrSet(resName, "capacity.#"),
					resource.TestCheckResourceAttrSet(resName, "iops.#"),
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
func testAccCheckIBMISVolumeProfileDataSourceSdpConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_volume_profile" "test1" {
		name = "%s"
	}`, "sdp")
}
