// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMDLProviderPortsDataSource_basic(t *testing.T) {
	name := "dl_provider_ports"
	resName := "data.ibm_dl_provider_ports.test_dl_provider_ports"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
