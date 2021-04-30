// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCmOfferingInstance(t *testing.T) {
	clusterId := os.Getenv("CATMGMT_CLUSTERID")
	clusterRegion := os.Getenv("CATMGMT_CLUSTERREGION")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceConfig(clusterId, clusterRegion),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cm_offering_instance.cm_offering_instance", "label"),
					testAccCheckIBMCmOfferingInstanceExists("ibm_cm_offering_instance.cm_offering_instance"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cm_offering_instance.cm_offering_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func testAccCheckIBMCmOfferingInstanceConfig(clusterId string, clusterRegion string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_instance_catalog"
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

		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			label = "tf_test_offering_instance_label"
			catalog_id = ibm_cm_catalog.cm_catalog.id
			offering_id = ibm_cm_offering.cm_offering.id
			kind_format = "operator"
			version = ibm_cm_version.cm_version.version
			cluster_id = "%s"
			cluster_region = "%s"
			cluster_namespaces = ["tf-cm-test"]
			cluster_all_namespaces = false
		}
		`, clusterId, clusterRegion)
}

func testAccCheckIBMCmOfferingInstanceDestroy(s *terraform.State) error {
	catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cm_offering_instance" {
			continue
		}

		getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

		getOfferingInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		// Try to find the key
		_, response, err := catalogManagementClient.GetOfferingInstance(getOfferingInstanceOptions)

		if err == nil {
			return fmt.Errorf("A offering instance resource (provision instance of a catalog offering). still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for A offering instance resource (provision instance of a catalog offering). (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMCmOfferingInstanceExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		catalogManagementClient, err := testAccProvider.Meta().(ClientSession).CatalogManagementV1()
		if err != nil {
			return err
		}

		getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

		getOfferingInstanceOptions.SetInstanceIdentifier(rs.Primary.ID)

		_, _, err = catalogManagementClient.GetOfferingInstance(getOfferingInstanceOptions)
		if err != nil {
			return err
		}

		return nil
	}
}
