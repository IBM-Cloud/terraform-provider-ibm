// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsDedicatedHostProfileDataSourceBasic(t *testing.T) {

	resName := "data.ibm_is_dedicated_host_profile.dhprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostProfileDataSourceConfigBasic(acc.DedicatedHostProfileName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.DedicatedHostProfileName),
					resource.TestCheckResourceAttrSet(resName, "class"),
					resource.TestCheckResourceAttrSet(resName, "family"),
					resource.TestCheckResourceAttrSet(resName, "vcpu_manufacturer.#"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostProfileDataSourceConfigBasic(profile string) string {
	return fmt.Sprintf(`
	 
	 data "ibm_is_dedicated_host_profile" "dhprofile" {
		 name = "%s"
	 }
	 `, profile)
}
