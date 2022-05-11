// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsSshKeysDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsSshKeysDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_ssh_keys.is_ssh_keys", "keys.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsSshKeysDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_ssh_keys" "is_ssh_keys" {
		}
	`)
}
