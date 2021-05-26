// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_instance_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISInstanceProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", instanceProfileName),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "bandwidth.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "memory.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "architecture"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "port_speed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_architecture.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_profile.test1", "vcpu_count.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceProfileDataSourceConfig() string {
	return fmt.Sprintf(`

data "ibm_is_instance_profile" "test1" {
	name = "%s"
}`, instanceProfileName)
}
