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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMResourceQuotaDataSource_basic(t *testing.T) {

	name := "Trial Quota"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
