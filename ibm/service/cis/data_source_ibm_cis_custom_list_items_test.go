// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisCustomListItemsDataSource_Basic(t *testing.T) {
	name := "data.ibm_cis_custom_list_items.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomListItemsDataSource_basic("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "items.#"),
				),
			},
		},
	})
}
func testAccCheckCisCustomListItemsDataSource_basic(id string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data ibm_cis_custom_lists custom_lists {
    	cis_id = data.ibm_cis.cis.id
	}

	data "ibm_cis_custom_list_items" "%[1]s" {
		cis_id = data.ibm_cis.cis.id
		list_id = data.ibm_cis_custom_lists.custom_lists.lists[0].list_id
	  }
`, id)
}
