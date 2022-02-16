// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISBMSProfileDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_server_profile.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISBMSProfileDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
				),
			},
		},
	})
}

func testAccCheckIBMISBMSProfileDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
		data "ibm_is_bare_metal_server_profiles" "testbmsps" {
		}

		data "ibm_is_bare_metal_server_profile" "test1" {
			name = data.ibm_is_bare_metal_server_profiles.testbmsps.profiles.0.name
		}`)
}
