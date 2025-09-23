// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

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
					resource.TestCheckResourceAttrSet(resName, "allowed_use.#"),
					resource.TestCheckResourceAttrSet(resName, "allowed_use.0.api_version"),
					resource.TestCheckResourceAttrSet(resName, "allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(resName, "allowed_use.0.instance"),
				),
			},
		},
	})
}
func TestAccIBMISImageDataSource_id404(t *testing.T) {
	imageId := "8843-5fr454ft-f6-4565-9555-5f889f5f3f7777"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMISImageDataSource404Config(imageId),
				ExpectError: regexp.MustCompile("Error fetching image with id"),
			},
		},
	})
}
func TestAccIBMISImageDataSource_All(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	imageName := fmt.Sprintf("tfimage-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDataSourceAllConfig(imageName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.name"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.dedicated_host_only"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.display_name"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.family"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.href"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.vendor"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.version"),
					resource.TestCheckResourceAttrSet(resName, "operating_system.0.architecture"),
					resource.TestCheckResourceAttrSet(resName, "status"),
					resource.TestCheckResourceAttrSet(resName, "resource_group.0.id"),
					resource.TestCheckResourceAttrSet(resName, "resource_group.0.name"),
				),
			},
		},
	})
}
func TestAccIBMISImageDataSource_ilc(t *testing.T) {
	resName := "data.ibm_is_image.test1"
	imageName := fmt.Sprintf("tfimage-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDataSourceConfigIlc(imageName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resName, "name", imageName),
					resource.TestCheckResourceAttrSet(resName, "os"),
					resource.TestCheckResourceAttrSet(resName, "architecture"),
					resource.TestCheckResourceAttrSet(resName, "visibility"),
					resource.TestCheckResourceAttrSet(resName, "status"),
					resource.TestCheckResourceAttrSet(resName, "created_at"),
					resource.TestCheckResourceAttrSet(resName, "deprecation_at"),
					resource.TestCheckResourceAttrSet(resName, "obsolescence_at"),
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

func TestAccIBMISImageDataSourceRemoteAccountId(t *testing.T) {
	resName := "data.ibm_is_image.example"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISImageDataSourceWithRemoteAccountId(),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttrSet(resName, "image.0.remote.#"),
					// resource.TestCheckResourceAttrSet(resName, "image.0.remote.0.account.#"),
					resource.TestCheckResourceAttrSet(resName, "remote.0.account.0.id"),
					resource.TestCheckResourceAttrSet(resName, "remote.0.account.0.resource_type"),
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

func testAccCheckIBMISImageDataSource404Config(imageId string) string {
	return fmt.Sprintf(`
	data "ibm_is_image" "test1" {
		identifier = "%s"
	}`, imageId)
}

func testAccCheckIBMISImageDataSourceAllConfig(imageName string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "test1" {
		status = "available"
	}
	data "ibm_is_image" "test1" {
		name = data.ibm_is_images.test1.images.0.name
	}`)
}

func testAccCheckIBMISImageDataSourceConfigIlc(imageName string) string {
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

func testAccCheckIBMISCatalogImageDataSourceRemoteAccountId() string {
	return fmt.Sprintf(`
	data "ibm_is_images" "test1" {
		catalog_managed = true
	}`)
}

func testAccCheckIBMISImageDataSourceWithRemoteAccountId() string {
	return fmt.Sprintf(`
		data "ibm_is_image" "example" {
  		name = "ibm-ubuntu-18-04-1-minimal-amd64-1"
	}`)
}
