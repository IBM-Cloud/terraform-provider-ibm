// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMDLProviderPortsDataSource_basic(t *testing.T) {
	name := "dl_provider_ports"
	resName := "data.ibm_dl_provider_ports.test_dl_provider_ports"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLProviderPortsDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "ports.0.port_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDLProviderPortsDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	   data "ibm_dl_provider_ports" "test_%s" {
	   }
	  `, name)
}
