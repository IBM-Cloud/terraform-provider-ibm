// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCmOfferingDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_id"),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingDataSourceSimpleArgs(t *testing.T) {
	offeringLabel := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	offeringName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	offeringOfferingIconURL := fmt.Sprintf("tf_offering_icon_url_%d", acctest.RandIntRange(10, 100))
	offeringShortDescription := fmt.Sprintf("tf_offering_icon_url_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingDataSourceConfig(offeringLabel, offeringName, offeringOfferingIconURL, offeringShortDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "rev"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "offering_icon_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "updated"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "short_description"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "catalog_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering.cm_offering", "support.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_basic_catalog_label_for_offering_data_source"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
		}

		data "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_offering.cm_offering.catalog_id
			offering_id = ibm_cm_offering.cm_offering.id
		}
	`)
}

func testAccCheckIBMCmOfferingDataSourceConfig(offeringLabel string, offeringName string, offeringOfferingIconURL string, offeringShortDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "catalog_%s"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "%s"
			name = "%s"
			offering_icon_url = "%s"
			short_description = "%s"
		}

		data "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_offering.cm_offering.catalog_id
			offering_id = ibm_cm_offering.cm_offering.id
		}
	`, offeringLabel, offeringLabel, offeringName, offeringOfferingIconURL, offeringShortDescription)
}
