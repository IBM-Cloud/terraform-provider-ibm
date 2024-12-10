// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMTransitGatewaysLocationDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_location.test_tg_location"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewaysLocationDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "billing_location"),
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "type"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewaysLocationDataSourceConfig() string {
	// status filter defaults to empty
	return `
	data "ibm_tg_location" "test_tg_location" {
		name = "us-south"
		}   `
}
