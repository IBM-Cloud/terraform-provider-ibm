// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmVersionDataSource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version_data", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_version.cm_version_data", "repo_url"),
				),
			},
		},
	})
}

func testAccCheckIBMCmVersionDataSourceConfig() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_data_version_catalog"
			short_description = "testing terraform provider with catalog"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "tf_test_offering"
			tags = ["dev_ops", "target_roks", "operator"]
		}

		resource "ibm_cm_version" "cm_version" {
			catalog_identifier = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			zipurl = "https://raw.githubusercontent.com/operator-framework/community-operators/master/community-operators/cockroachdb/5.0.3/manifests/cockroachdb.clusterserviceversion.yaml"
		}

		data "ibm_cm_version" "cm_version_data" {
			version_loc_id = ibm_cm_version.cm_version.id
		}
		`)
}
