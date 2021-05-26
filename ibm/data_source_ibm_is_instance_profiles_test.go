// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceProfilesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_profiles.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISInstanceProfilesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profiles.test1", "profiles.0.vcpu_count.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceProfilesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_instance_profiles" "test1" {
      }`)
}
