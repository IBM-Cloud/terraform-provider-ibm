// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccCheckIBMPICatalogImagesDataSourceBasicConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_catalog_images" "power_catalog_images_basic" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPICatalogImagesDataSourceSAPConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_catalog_images" "power_catalog_images_sap" {
			pi_cloud_instance_id = "%s"
			sap = "true"
		}`, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPICatalogImagesDataSourceVTLConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_catalog_images" "power_catalog_images_vtl" {
			pi_cloud_instance_id = "%s"
			vtl = "true"
		}`, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPICatalogImagesDataSourceSAP_And_VTLConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_catalog_images" "power_catalog_images_sap_and_vtl" {
			pi_cloud_instance_id = "%s"
			sap = "true"
			vtl = "true"
		}`, acc.Pi_cloud_instance_id)
}

func TestAccIBMPICatalogImagesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICatalogImagesDataSourceBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_catalog_images.power_catalog_images_basic", "id"),
				),
			},
		},
	})
}

func TestAccIBMPICatalogImagesDataSourceSAP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICatalogImagesDataSourceSAPConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_catalog_images.power_catalog_images_sap", "id"),
				),
			},
		},
	})
}

func TestAccIBMPICatalogImagesDataSourceVTL(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICatalogImagesDataSourceVTLConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_catalog_images.power_catalog_images_vtl", "id"),
				),
			},
		},
	})
}

func TestAccIBMPICatalogImagesDataSourceSAPAndVTL(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICatalogImagesDataSourceSAP_And_VTLConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_catalog_images.power_catalog_images_sap_and_vtl", "id"),
				),
			},
		},
	})
}
