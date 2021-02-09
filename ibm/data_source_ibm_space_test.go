/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMSpaceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSpaceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_space.testacc_ds_space", "org", cfOrganization),
					resource.TestCheckResourceAttr("data.ibm_space.testacc_ds_space", "space", cfSpace),
				),
			},
		},
	})
}

func testAccCheckIBMSpaceDataSourceConfig() string {
	return fmt.Sprintf(`
data "ibm_space" "testacc_ds_space" {
    org = "%s"
	space = "%s"
}`, cfOrganization, cfSpace)

}
