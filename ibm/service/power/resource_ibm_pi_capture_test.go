// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPICaptureBasic(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"
	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICaptureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICaptureConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICaptureExists(captureRes),
					resource.TestCheckResourceAttr(captureRes, "pi_capture_name", name),
					resource.TestCheckResourceAttrSet(captureRes, "image_id"),
				),
			},
		},
	})
}
func TestAccIBMPICaptureWithVolume(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"
	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPICaptureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICaptureWithVolumeConfig(name, helpers.PIInstanceHealthOk),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICaptureExists(captureRes),
					resource.TestCheckResourceAttr(captureRes, "pi_capture_name", name),
					resource.TestCheckResourceAttrSet(captureRes, "image_id"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIBMPICaptureCloudStorage(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"
	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICaptureCloudStorageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(captureRes, "pi_capture_name", name),
				),
			},
		},
	})
}

func TestAccIBMPICaptureBoth(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"
	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICaptureBothConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(captureRes, "pi_capture_name", name),
					resource.TestCheckResourceAttrSet(captureRes, "image_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPICaptureExists(n string) resource.TestCheckFunc {
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
		parts, err := flex.IdParts(rs.Primary.ID)
		cloudInstanceID := parts[0]
		captureID := parts[1]
		if err != nil {
			return err
		}
		client := st.NewIBMPIImageClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(captureID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPICaptureDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_capture" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		cloudInstanceID := parts[0]
		captureID := parts[1]
		if err != nil {
			return err
		}
		imageClient := st.NewIBMPIImageClient(context.Background(), sess, cloudInstanceID)
		_, err = imageClient.Get(captureID)
		if err == nil {
			return fmt.Errorf("PI Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMPICaptureWithVolumeConfig(name string, healthStatus string) string {
	return testAccCheckIBMPIInstanceConfig(name, healthStatus) + fmt.Sprintf(`
	resource "ibm_pi_capture" "capture_instance" {
		depends_on=[ibm_pi_instance.power_instance]
		pi_cloud_instance_id="%[1]s"
		pi_capture_name  = "%[2]s"
		pi_instance_name = ibm_pi_instance.power_instance.pi_instance_name
		pi_capture_destination = "image-catalog"
		pi_capture_volume_ids = [ibm_pi_volume.power_volume.volume_id]
	}
	`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPICaptureConfigBasic(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_capture" "capture_instance" {
		pi_cloud_instance_id="%[1]s"
		pi_capture_name = "%s"
		pi_instance_name = "%s"
		pi_capture_destination = "image-catalog"
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_instance_name)
}

func testAccCheckIBMPICaptureCloudStorageConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_capture" "capture_instance" {
		pi_cloud_instance_id="%[1]s"
		pi_capture_name  = "%s"
		pi_instance_name = "%s"
		pi_capture_destination = "cloud-storage"
		pi_capture_cloud_storage_region = "us-east"
		pi_capture_cloud_storage_access_key = "%s"
		pi_capture_cloud_storage_secret_key = "%s"
		pi_capture_storage_image_path = "%s"
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_instance_name, acc.Pi_capture_cloud_storage_access_key, acc.Pi_capture_cloud_storage_secret_key, acc.Pi_capture_storage_image_path)
}

func testAccCheckIBMPICaptureBothConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_capture" "capture_instance" {
		pi_cloud_instance_id="%[1]s"
		pi_capture_name = "%s"
		pi_instance_name = "%s"
		pi_capture_destination  = "both"
		pi_capture_cloud_storage_region = "us-east"
		pi_capture_cloud_storage_access_key = "%s"
		pi_capture_cloud_storage_secret_key = "%s"
		pi_capture_storage_image_path = "%s"
	}
	`, acc.Pi_cloud_instance_id, name, acc.Pi_instance_name, acc.Pi_capture_cloud_storage_access_key, acc.Pi_capture_cloud_storage_secret_key, acc.Pi_capture_storage_image_path)
}
