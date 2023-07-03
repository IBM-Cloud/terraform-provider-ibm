// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISImageDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	imageName := fmt.Sprintf("tfimage-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDataSourceConfig(imageName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", imageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}
func TestAccIBMISImageDataSource_catalog(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	resName1 := "data.ibm_is_image.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISCatalogImageDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
					resource.TestCheckResourceAttr(resName, "catalog_offering.0.managed", "true"),
					resource.TestCheckResourceAttrSet(resName, "catalog_offering.0.version.0.crn"),
					resource.TestCheckResourceAttrSet(resName1, "name"),
					resource.TestCheckResourceAttrSet(resName1, "os"),
					resource.TestCheckResourceAttrSet(resName1, "architecture"),
					resource.TestCheckResourceAttrSet(resName1, "visibility"),
					resource.TestCheckResourceAttrSet(resName1, "status"),
					resource.TestCheckResourceAttr(resName1, "catalog_offering.0.managed", "true"),
					resource.TestCheckResourceAttrSet(resName1, "catalog_offering.0.version.0.crn"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_VisibiltyPublic(t *testing.T) {
	resName := "data.ibm_is_image.test1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDataSourceWithVisibilityPublic("public"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", acc.IsImageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func TestAccIBMISImageDataSource_With_VisibiltyPrivate(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	imageName := fmt.Sprintf("tfimage-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDataSourceWithVisibilityPrivate(imageName, "private"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", imageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
				),
			},
		},
	})
}

func testAccCheckIBMISImageDataSourceConfig(imageName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_image" "isExampleImage" {
		href = "%s"
		name = "%s"
		operating_system = "%s"
	}
	data "ibm_is_image" "test1" {
		name = ibm_is_image.isExampleImage.name
	}`, acc.Image_cos_url, imageName, acc.Image_operating_system)
}
func testAccCheckIBMISCatalogImageDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_images" "test1" {
		catalog_managed = true
	}
	data "ibm_is_image" "test1" {
		name = data.ibm_is_images.test1.images.0.name
	}
	data "ibm_is_image" "test2" {
		identifier = data.ibm_is_images.test1.images.0.id
	}`)
}

func testAccCheckIBMISImageDataSourceWithVisibilityPublic(visibility string) string {
	return fmt.Sprintf(`
	data "ibm_is_image" "test1" {
		name = "%s"
		visibility = "%s"
	}`, acc.IsImageName, visibility)
}

func testAccCheckIBMISImageDataSourceWithVisibilityPrivate(imageName, visibility string) string {
	return fmt.Sprintf(`
	resource "ibm_is_image" "isExampleImage" {
		href = "%s"
		name = "%s"
		operating_system = "%s"
	}
	data "ibm_is_image" "test1" {
		name = ibm_is_image.isExampleImage.name
		visibility = "%s"
	}`, acc.Image_cos_url, imageName, acc.Image_operating_system, visibility)
}
