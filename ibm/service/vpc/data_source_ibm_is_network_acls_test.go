// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIsNetworkAclsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsNetworkAclsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acls.is_network_acls", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_network_acls.is_network_acls", "network_acls.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsNetworkAclsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_network_acls" "is_network_acls" {
		}
	`)
}
