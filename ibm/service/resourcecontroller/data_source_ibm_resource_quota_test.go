// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package resourcecontroller_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourceQuotaDataSource_basic(t *testing.T) {

	name := "Trial Quota"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourceQuotaDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_resource_quota.testacc_ds_resource_quota", "name", name),
				),
			},
		},
	})
}

func TestAccIBMResourceQuotaDataSource_invalid_name(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMResourceQuotaDataSourceConfig("abc"),
				ExpectError: regexp.MustCompile(`Error retrieving resource quota`),
			},
		},
	})
}

func testAccCheckIBMResourceQuotaDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
data "ibm_resource_quota" "testacc_ds_resource_quota" {
    name = "%s"
}`, name)

}
