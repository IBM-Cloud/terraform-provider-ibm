/* IBM Confidential
 *  Object Code Only Source Materials
 *  5747-SM3
 *  (c) Copyright IBM Corp. 2017,2021
 *
 *  The source code for this program is not published or otherwise divested
 *  of its trade secrets, irrespective of what has been deposited with the
 *  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCmOfferingInstance(t *testing.T) {
	//	var conf catalogmanagementv1.OfferingInstance
	label := os.Getenv("CATMGMT_LABEL")
	catalogID := os.Getenv("CATMGMT_CATALOGID")
	offeringID := os.Getenv("CATMGMT_OFFERINGID")
	kindFormat := "operator"
	version := os.Getenv("CATMGMT_VERSION")
	clusterID := os.Getenv("CATMGMT_CLUSTERID")
	region := os.Getenv("CATMGMT_CLUSTERREGION")
	clusterNamespace := os.Getenv("CATMGMT_NAMESPACE")
	clusterAllNamespaces := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceConfig(label, catalogID, offeringID, kindFormat, version, clusterID, region, clusterNamespace, clusterAllNamespaces),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "label", label),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "catalog_id", catalogID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "offering_id", offeringID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "kind_format", kindFormat),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_id", clusterID),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_region", region),
					resource.TestCheckResourceAttr("ibm_cm_offering_instance.cm_offering_instance", "cluster_all_namespaces", clusterAllNamespaces),
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

func testAccCheckIBMCmOfferingInstanceConfig(label string, catalogID string, offeringID string, kindFormat string, version string, clusterID string, region string, clusterNamespace string, clusterAllNamespaces string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_offering_instance" "cm_offering_instance" {
			label = "%s"
			catalog_id = "%s"
			offering_id = "%s"
			kind_format = "%s"
			version = "%s"
			cluster_id = "%s"
			cluster_region = "%s"
			cluster_namespaces = ["%s"]
			cluster_all_namespaces = %s
		}
	`, label, catalogID, offeringID, kindFormat, version, clusterID, region, clusterNamespace, clusterAllNamespaces)
}

func testAccCheckIBMCmOfferingInstanceExists(n string, obj catalogmanagementv1.OfferingInstance) resource.TestCheckFunc {

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

		offeringInstance, _, err := catalogManagementClient.GetOfferingInstance(getOfferingInstanceOptions)
		if err != nil {
			return err
		}

		obj = *offeringInstance
		return nil
	}
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
