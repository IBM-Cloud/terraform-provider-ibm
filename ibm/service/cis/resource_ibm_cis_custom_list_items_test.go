// Copyright IBM Corp. 2025. All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisCustomListItems_Basic(t *testing.T) {
	name := "ibm_cis_custom_list_items.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomListItems_basic("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckCisCustomListItems_basic(id string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`

	resource "ibm_cis_custom_list" "custom_list" {
		cis_id = data.ibm_cis.cis.id
		kind = "ip"
    	name = "acc_test_list"
	}

	resource "ibm_cis_custom_list_items" "%[1]s" {
    	cis_id    = data.ibm_cis.cis.id
    	list_id = ibm_cis_custom_list.custom_list.list_id
		items {
        	ip = "1.2.3.6"
    	}
    	items {
        	ip = "2.3.4.5"
    	}
	}
`, id)
}
