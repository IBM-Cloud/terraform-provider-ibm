// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISImagesDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_images.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImagesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "images.0.name"),
					resource.TestCheckResourceAttrSet(resName, "images.0.status"),
					resource.TestCheckResourceAttrSet(resName, "images.0.architecture"),
				),
			},
		},
	})
}
func TestAccIBMISImagesDataSource_catalog(t *testing.T) {
	resName := "data.ibm_is_images.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISCatalogImagesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "images.0.name"),
					resource.TestCheckResourceAttr(resName, "images.0.catalog_offering.0.managed", "true"),
					resource.TestCheckResourceAttrSet(resName, "images.0.status"),
					resource.TestCheckResourceAttrSet(resName, "images.0.architecture"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_FilterVisibilty(t *testing.T) {
	resName := "data.ibm_is_images.test1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImagesDataSourceWithVisibilityPublic("public"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "images.0.name"),
					resource.TestCheckResourceAttrSet(resName, "images.0.status"),
					resource.TestCheckResourceAttrSet(resName, "images.0.architecture"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_FilterStatus(t *testing.T) {
	resName := "data.ibm_is_images.test1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImagesDataSourceWithStatusPublic("available"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "images.0.name"),
					resource.TestCheckResourceAttrSet(resName, "images.0.status"),
					resource.TestCheckResourceAttrSet(resName, "images.0.architecture"),
				),
			},
		},
	})
}

func testAccCheckIBMISImagesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_images" "test1" {
      }`)
}
func testAccCheckIBMISCatalogImagesDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_is_images" "test1" {
		catalog_managed = true
      }`)
}

func testAccCheckIBMISImagesDataSourceWithVisibilityPublic(visibility string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "test1" {
		visibility = "%s"
	}
	`, visibility)
}

func testAccCheckIBMISImagesDataSourceWithStatusPublic(status string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "test1" {
		status = "%s"
	}
	`, status)
}
