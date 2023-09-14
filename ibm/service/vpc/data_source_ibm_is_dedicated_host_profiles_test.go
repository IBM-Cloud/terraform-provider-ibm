// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmIsDedicatedHostProfilesDataSourceBasic(t *testing.T) {

	resName := "data.ibm_is_dedicated_host_profiles.dhprofiles"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsDedicatedHostProfilesDataSourceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "profiles.0.name"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.class"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.family"),
					resource.TestCheckResourceAttrSet(resName, "profiles.0.vcpu_manufacturer.#"),
				),
			},
		},
	})
}

func testAccCheckIbmIsDedicatedHostProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
	  
	  data "ibm_is_dedicated_host_profiles" "dhprofiles" {
	  }
	  `)
}
