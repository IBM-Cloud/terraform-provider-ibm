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

func TestAccIBMCmCatalog(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog"),
					resource.TestCheckResourceAttrSet("ibm_cm_catalog.cm_catalog", "label"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_catalog.cm_catalog",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCmCatalogConfig() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_catalog"
			short_description = "testing terraform provider with catalog"
		}
		`)
}

func testAccCheckIBMCmCatalogExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		_, response, err := catalogManagementClient.GetCatalog(getCatalogOptions)
		if err != nil {
			if response.StatusCode == 404 {
				return nil
			}
			return err
		}

		return nil
	}
}

func testAccCheckIBMCmCatalogDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_catalog" {
			continue
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetCatalog(getCatalogOptions)

		if err == nil {
			return fmt.Errorf("cm_catalog still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 403 {
			return fmt.Errorf("Error checking for cm_catalog (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
