// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmCatalogBasic(t *testing.T) {
	var conf catalogmanagementv1.Catalog

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog", conf),
				),
			},
		},
	})
}

func TestAccIBMCmCatalogSimpleArgs(t *testing.T) {
	var conf catalogmanagementv1.Catalog
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmCatalogDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(label, shortDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmCatalogExists("ibm_cm_catalog.cm_catalog", conf),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "short_description", shortDescription),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmCatalogConfig(label, shortDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_catalog.cm_catalog", "short_description", shortDescription),
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

func testAccCheckIBMCmCatalogConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "basic-catalog-label-test"
			kind = "offering"
		}
	`)
}

func testAccCheckIBMCmCatalogConfig(label string, shortDescription string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "%s"
			kind = "offering"
			short_description = "%s"
		}
	`, label, shortDescription)
}

func testAccCheckIBMCmCatalogExists(n string, obj catalogmanagementv1.Catalog) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getCatalogOptions := &catalogmanagementv1.GetCatalogOptions{}

		getCatalogOptions.SetCatalogIdentifier(rs.Primary.ID)

		catalog, _, err := catalogManagementClient.GetCatalog(getCatalogOptions)
		if err != nil {
			return err
		}

		obj = *catalog
		return nil
	}
}

func testAccCheckIBMCmCatalogDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
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
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_catalog (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
