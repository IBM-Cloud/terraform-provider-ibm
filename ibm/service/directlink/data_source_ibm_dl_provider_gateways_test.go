// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDLProviderGWsDataSource_basic(t *testing.T) {
	name := "dl_provider_gws"
	resName := "data.ibm_dl_provider_gateways.test_dl_provider_gws"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLProviderGWsDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMDLProviderGWsDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_provider_gateways" "test_%s" {
	   }
	  `, name)
}
