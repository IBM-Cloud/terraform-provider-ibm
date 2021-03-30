// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmVersion(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmVersionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmVersionConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmVersionExists("ibm_cm_version.cm_version"),
					resource.TestCheckResourceAttrSet("ibm_cm_version.cm_version", "repo_url"),
				),
			},
		},
	})
}

func testAccCheckIBMCmVersionConfig() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_version_catalog"
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
		`)
}

func testAccCheckIBMCmVersionExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

		getVersionOptions.SetVersionLocID(rs.Primary.ID)

		_, _, err = catalogManagementClient.GetVersion(getVersionOptions)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMCmVersionDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_version" {
			continue
		}

		getVersionOptions := &catalogmanagementv1.GetVersionOptions{}

		getVersionOptions.SetVersionLocID(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetVersion(getVersionOptions)

		if err == nil {
			return fmt.Errorf("cm_version still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 403 {
			return fmt.Errorf("Error checking for cm_version (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
