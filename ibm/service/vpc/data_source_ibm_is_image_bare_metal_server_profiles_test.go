// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsImageBareMetalServerProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsImageBareMetalServerProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_image_bare_metal_server_profiles.is_image_bare_metal_server_profiles_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_bare_metal_server_profiles.is_image_bare_metal_server_profiles_instance", "bare_metal_server_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_bare_metal_server_profiles.is_image_bare_metal_server_profiles_instance", "bare_metal_server_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_bare_metal_server_profiles.is_image_bare_metal_server_profiles_instance", "bare_metal_server_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_image_bare_metal_server_profiles.is_image_bare_metal_server_profiles_instance", "bare_metal_server_profiles.0.resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsImageBareMetalServerProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_image_bare_metal_server_profiles" "is_image_bare_metal_server_profiles_instance" {
			identifier = "%s"
		}
	`, acc.IsImage)
}
