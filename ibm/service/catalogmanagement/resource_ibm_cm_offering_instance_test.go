// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCmOfferingInstance(t *testing.T) {
	clusterId := os.Getenv("CATMGMT_CLUSTERID")
	clusterRegion := os.Getenv("CATMGMT_CLUSTERREGION")
	planId := fmt.Sprintf("plan_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCmOfferingInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCmOfferingInstanceConfig(clusterId, clusterRegion, planId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cm_offering_instance.cm_offering_instance", "label"),
					testAccCheckIBMCmOfferingInstanceExists("ibm_cm_offering_instance.cm_offering_instance"),
				),
			},
			{
				ResourceName:            "ibm_cm_offering_instance.cm_offering_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"wait_until_successful"},
			},
		},
	})
}
func testAccCheckIBMCmOfferingInstanceConfig(clusterId string, clusterRegion string, planId string) string {
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
			zipurl = "https://raw.githubusercontent.com/operator-framework/community-operators/master/community-operators/flux/0.14.2/manifests/flux.v0.14.2.clusterserviceversion.yaml"
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
			install_plan = "Automatic"
			plan_id = "%s"
		}
		`, clusterId, clusterRegion, planId)
}

func testAccCheckIBMCmOfferingInstanceDestroy(s *terraform.State) error {
	catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
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
			return fmt.Errorf("[ERROR] A offering instance resource (provision instance of a catalog offering). still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for A offering instance resource (provision instance of a catalog offering). (%s) has been destroyed: %s", rs.Primary.ID, err)
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

		catalogManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CatalogManagementV1()
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
