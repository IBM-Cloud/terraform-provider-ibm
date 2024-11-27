// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMOrgDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOrgDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org.testacc_ds_org", "org", acc.CfOrganization),
				),
			},
		},
	})
}

func testAccCheckIBMOrgDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}`, acc.CfOrganization)

}
