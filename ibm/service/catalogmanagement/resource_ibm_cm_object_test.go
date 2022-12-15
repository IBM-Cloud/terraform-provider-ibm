// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func TestAccIBMCmObjectSimpleArgs(t *testing.T) {
	var conf catalogmanagementv1.CatalogObject
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	parentID := fmt.Sprintf("tf_parent_id_%d", acctest.RandIntRange(10, 100))
	label := fmt.Sprintf("tf_label_%d", acctest.RandIntRange(10, 100))
	shortDescription := fmt.Sprintf("tf_short_description_%d", acctest.RandIntRange(10, 100))
	kind := "preset_configuration"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmObjectConfig(name, parentID, label, shortDescription, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCmObjectExists("ibm_cm_object.cm_object", conf),
					resource.TestCheckResourceAttr("ibm_cm_object.cm_object", "name", name),
					resource.TestCheckResourceAttr("ibm_cm_object.cm_object", "parent_id", parentID),
					resource.TestCheckResourceAttr("ibm_cm_object.cm_object", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_object.cm_object", "short_description", shortDescription),
					resource.TestCheckResourceAttr("ibm_cm_object.cm_object", "kind", kind),
				),
			},
		},
	})
}

func testAccCheckIBMCmObjectConfig(name string, parentID string, label string, shortDescription string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "test_preset_catalog_tf_test"
			kind = "preset_configuration"
		}

		resource "ibm_cm_object" "cm_object" {
			catalog_id = ibm_cm_catalog.cm_catalog.id
			name = "%s"
			parent_id = "%s"
			label = "%s"
			short_description = "%s"
			kind = "%s"
		}
	`, name, parentID, label, shortDescription, kind)
}

func testAccCheckIBMCmObjectExists(n string, obj catalogmanagementv1.CatalogObject) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getObjectOptions := &catalogmanagementv1.GetObjectOptions{}

		getObjectOptions.SetCatalogIdentifier(rs.Primary.Attributes["catalog_id"])
		getObjectOptions.SetObjectIdentifier(rs.Primary.ID)

		catalogObject, _, err := catalogManagementClient.GetObject(getObjectOptions)
		if err != nil {
			return err
		}

		obj = *catalogObject
		return nil
	}
}

func testAccCheckIBMCmObjectDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_object" {
			continue
		}

		getObjectOptions := &catalogmanagementv1.GetObjectOptions{}

		getObjectOptions.SetCatalogIdentifier(rs.Primary.Attributes["catalog_id"])
		getObjectOptions.SetObjectIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetObject(getObjectOptions)

		if err == nil {
			return fmt.Errorf("cm_object still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cm_object (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
