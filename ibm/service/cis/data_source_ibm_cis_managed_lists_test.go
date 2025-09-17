// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisManagedListsDataSource_Basic(t *testing.T) {
	name := "data.ibm_cis_managed_lists.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisManagedListsDataSource_basic("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "lists.#", "5"),
				),
			},
		},
	})
}
func testAccCheckCisManagedListsDataSource_basic(id string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_managed_lists" "%[1]s" {
		cis_id = data.ibm_cis.cis.id
	  }
`, id)
}
