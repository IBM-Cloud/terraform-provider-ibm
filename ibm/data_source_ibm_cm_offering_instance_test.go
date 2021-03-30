// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCmOfferingInstanceDataSource(t *testing.T) {
	clusterId := os.Getenv("CATMGMT_CLUSTERID")
	clusterRegion := os.Getenv("CATMGMT_CLUSTERREGION")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCmOfferingInstanceDataSourceConfig(clusterId, clusterRegion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance_data", "label"),
					resource.TestCheckResourceAttrSet("data.ibm_cm_offering_instance.cm_offering_instance_data", "url"),
				),
			},
		},
	})
}

func testAccCheckIBMCmOfferingInstanceDataSourceConfig(clusterId string, clusterRegion string) string {
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
			zipurl = "https://raw.githubusercontent.com/operator-framework/community-operators/master/community-operators/flux/0.9.1/manifests/flux.v0.9.1.clusterserviceversion.yaml"
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
		}

		data "ibm_cm_offering_instance" "cm_offering_instance_data" {
			instance_identifier = ibm_cm_offering_instance.cm_offering_instance.id
		}
		  
		`, clusterId, clusterRegion)
}
