// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudfoundry_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMOrgQuotaDataSource_basic(t *testing.T) {

	name := "qIBM"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMOrgQuotaDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org_quota.testacc_ds_org_quota", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMOrgQuotaDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
data "ibm_org_quota" "testacc_ds_org_quota" {
    name = "%s"
}`, name)

}
