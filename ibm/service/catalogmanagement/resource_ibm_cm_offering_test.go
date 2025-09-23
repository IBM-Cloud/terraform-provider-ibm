// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIBMCmOfferingBasic(t *testing.T) {
	var conf catalogmanagementv1.Offering

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingExists("ibm_cm_offering.cm_offering", conf),
				),
			},
		},
	})
}

func TestAccIBMCmOfferingSimpleArgs(t *testing.T) {
	var conf catalogmanagementv1.Offering
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescription := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))
	labelUpdate := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	shortDescriptionUpdate := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	longDescriptionUpdate := fmt.Sprintf("tf_long_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfig(label, name, shortDescription, longDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmOfferingExists("ibm_cm_offering.cm_offering", conf),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "name", name),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "long_description", longDescription),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingConfig(labelUpdate, name, shortDescriptionUpdate, longDescriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "label", labelUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "short_description", shortDescriptionUpdate),
					resource.TestCheckResourceAttr("ibm_cm_offering.cm_offering", "long_description", longDescriptionUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_catalog_label_1"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
		}
	`)
}

func testAccCheckIBMCmOfferingConfig(label string, name string, shortDescription string, longDescription string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_tf_catalog_label_2"
			kind = "offering"
		}

		resource "ibm_cm_offering" "cm_offering" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			label = "%s"
			name = "%s"
			short_description = "%s"
			long_description = "%s"
		}
	`, label, name, shortDescription, longDescription)
}

func testAccCheckIBMCmOfferingExists(n string, obj catalogmanagementv1.Offering) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		getOfferingOptions.SetCatalogIdentifier(rs.Primary.Attributes["catalog_id"])
		getOfferingOptions.SetOfferingID(rs.Primary.ID)

		offering, _, err := catalogManagementClient.GetOffering(getOfferingOptions)
		if err != nil {
			return err
		}

		obj = *offering
		return nil
	}
}

func testAccCheckIBMCmOfferingCreate(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_offering" {
			continue
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		getOfferingOptions.SetCatalogIdentifier(rs.Primary.Attributes["catalog_id"])
		getOfferingOptions.SetOfferingID(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetOffering(getOfferingOptions)

		if err == nil {
			return fmt.Errorf("cm_offering still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_offering (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMCmOfferingDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_offering" {
			continue
		}

		getOfferingOptions := &catalogmanagementv1.GetOfferingOptions{}

		getOfferingOptions.SetCatalogIdentifier(rs.Primary.Attributes["catalog_id"])
		getOfferingOptions.SetOfferingID(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetOffering(getOfferingOptions)

		if err == nil {
			return fmt.Errorf("cm_offering still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_offering (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
