// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmOfferingInstanceDataSource(t *testing.T) {
	clusterId := os.Getenv("CATMGMT_CLUSTERID")
	clusterRegion := os.Getenv("CATMGMT_CLUSTERREGION")
	resourceGroupID := os.Getenv("CATMGMT_RGID")
	planId := fmt.Sprintf("plan_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCmOfferingInstanceDataSourceConfig(clusterId, clusterRegion, resourceGroupID, planId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance_data", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance_data", "url"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance_data", "resource_group_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingInstanceDataSourceConfig(clusterId string, clusterRegion string, resourceGroupID string, planId string) string {
	return fmt.Sprintf(`

		resource "ibm_cm_catalog" "cm_catalog" {
			label = "tf_test_data_instance_catalog"
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
			cluster_namespaces = ["tf-cm-data-test"]
			cluster_all_namespaces = false
			resource_group_id = "%s"
			install_plan = "Automatic"
			plan_id = "%s"
		}

		data "ibm_cm_offering_instance" "cm_offering_instance_data" {
			instance_identifier = ibm_cm_offering_instance.cm_offering_instance.id
		}
		  
		`, clusterId, clusterRegion, resourceGroupID, planId)
}
