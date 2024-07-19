// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIImagebasic(t *testing.T) {

	name := fmt.Sprintf("tf-pi-image-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIImageExists("ibm_pi_image.power_image"),
					resource.TestCheckResourceAttr(
						"ibm_pi_image.power_image", "pi_image_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_image" {
			continue
		}
		cloudInstanceID, imageID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		imageC := st.NewIBMPIImageClient(context.Background(), sess, cloudInstanceID)
		_, err = imageC.Get(imageID)
		if err == nil {
			return fmt.Errorf("PI Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIImageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		cloudInstanceID, imageID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := st.NewIBMPIImageClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(imageID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPIImageConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "power_image" {
		pi_image_name       = "%s"
		pi_image_id         = "IBMi-74-01-001"
		pi_cloud_instance_id = "%s"
	  }
	`, name, acc.Pi_cloud_instance_id)
}

func TestAccIBMPIImageCOSPublicImport(t *testing.T) {
	imageRes := "ibm_pi_image.cos_image"
	name := fmt.Sprintf("tf-pi-image-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageCOSPublicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIImageExists(imageRes),
					resource.TestCheckResourceAttr(imageRes, "pi_image_name", name),
					resource.TestCheckResourceAttrSet(imageRes, "image_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageCOSPublicConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "cos_image" {
		pi_image_name       = "%[1]s"
		pi_cloud_instance_id = "%[2]s"
		pi_image_bucket_name = "%[3]s"
		pi_image_bucket_access = "public"
		pi_image_bucket_region = "us-south"
		pi_image_bucket_file_name = "%[4]s"
		pi_image_storage_type = "tier1"
	}
	`, name, acc.Pi_cloud_instance_id, acc.Pi_image_bucket_name, acc.Pi_image_bucket_file_name)
}

func TestAccIBMPIImageBYOLImport(t *testing.T) {
	imageRes := "ibm_pi_image.cos_image"
	name := fmt.Sprintf("tf-pi-image-byoi-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageBYOLConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIImageExists(imageRes),
					resource.TestCheckResourceAttr(imageRes, "pi_image_name", name),
					resource.TestCheckResourceAttrSet(imageRes, "image_id"),
				),
			},
		},
	})
}
func testAccCheckIBMPIImageBYOLConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "cos_image" {
		pi_cloud_instance_id = "%[2]s"
		pi_image_access_key = "%[5]s"
		pi_image_bucket_access = "private"
		pi_image_bucket_file_name = "%[4]s" 
		pi_image_bucket_name = "%[3]s" 
		pi_image_bucket_region = "us-east"
		pi_image_name       = "%[1]s"
		pi_image_secret_key = "%[6]s"
		pi_image_storage_type = "tier3"
		pi_image_import_details {
			license_type = "byol"
			product = "Hana"
			vendor = "SAP"
		}
	}
	`, name, acc.Pi_cloud_instance_id, acc.Pi_image_bucket_name, acc.Pi_image_bucket_file_name, acc.Pi_image_bucket_access_key, acc.Pi_image_bucket_secret_key)
}
