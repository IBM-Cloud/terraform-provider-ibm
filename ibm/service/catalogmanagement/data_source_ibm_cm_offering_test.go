// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmOfferingDataSource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCmOfferingDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering_data", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering_data", "label"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingDataSourceConfig() string {
	return `

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_data_offering_catalog"
			short_description = "testing terraform provider with catalog"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "tf_test_offering"
			tags = ["dev_ops", "target_roks", "operator"]
		}
		
		data "ibm_cm_offering" "cm_offering_data" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
		}
		`
}
