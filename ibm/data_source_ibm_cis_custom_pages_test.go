// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisCustomPagesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_custom_pages.test_custom_pages"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisCustomPagesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_custom_pages.0.page_id"),
					resource.TestCheckResourceAttrSet(node, "cis_custom_pages.0.state"),
				),
			},
		},
	})
}

func testAccCheckIBMCisCustomPagesDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckIBMCisDNSRecordConfigCisDSBasic("test", cisDomainStatic) +
		fmt.Sprintf(`
	data "ibm_cis_custom_pages" "test_custom_pages" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = ibm_cis_dns_record.test.domain_id
	  }`)
}
