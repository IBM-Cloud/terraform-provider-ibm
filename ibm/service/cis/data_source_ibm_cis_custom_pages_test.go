// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisCustomPagesDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_custom_pages.test_custom_pages"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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
	return testAccCheckIBMCisDNSRecordConfigCisDSBasic("test", acc.CisDomainStatic) +
		`
	data "ibm_cis_custom_pages" "test_custom_pages" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = ibm_cis_dns_record.test.domain_id
	  }`
}
