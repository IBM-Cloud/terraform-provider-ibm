// Copyright IBM Corp. 2025. All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisCustomList_Basic(t *testing.T) {
	name := "ibm_cis_custom_list.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisCustomList_basic("test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "kind", "ip"),
					resource.TestCheckResourceAttr(name, "name", "acc_test_list"),
				),
			},
		},
	})
}

func testAccCheckCisCustomList_basic(id string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`

	resource "ibm_cis_custom_list" "%[1]s" {
		cis_id = data.ibm_cis.cis.id
		kind = "ip"
    	name = "acc_test_list"
	  }
`, id)
}
