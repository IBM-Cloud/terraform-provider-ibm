// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMSpaceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSpaceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_space.testacc_ds_space", "org", acc.CfOrganization),
					resource.TestCheckResourceAttr("data.ibm_space.testacc_ds_space", "space", acc.CfSpace),
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
}`, acc.CfOrganization, acc.CfSpace)

}
