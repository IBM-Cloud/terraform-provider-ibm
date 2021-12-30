// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"errors"
	"fmt"
	"testing"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPICaptureBasic(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"

	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPICaptureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICaptureConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICaptureExists(captureRes),
					resource.TestCheckResourceAttr(captureRes, "pi_capture_name", name),
				),
			},
		},
	})
}
func TestAccIBMPICaptureWithVolume(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"
	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPICaptureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPICaptureWithVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPICaptureExists(captureRes),
					resource.TestCheckResourceAttr(captureRes, "pi_capture_name", name),
				),
			},
		},
	})
}

func TestAccIBMPICaptureCloudStorage(t *testing.T) {
	captureRes := "ibm_pi_capture.capture_instance"
	name := fmt.Sprintf("tf-pi-capture-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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

func testAccCheckIBMPICaptureExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		cloudInstanceID, captureID, err := splitID(rs.Primary.ID)
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
	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_capture" {
			continue
		}
		cloudInstanceID, captureID, err := splitID(rs.Primary.ID)
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

func testAccCheckIBMPICaptureWithVolumeConfig(name string) string {
	return fmt.Sprintf(`
	data "ibm_pi_volume" "dsvolume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name      = "%[5]s"
	  }

	resource "ibm_pi_capture" "capture_instance" {
		pi_cloud_instance_id="%[1]s"
		pi_capture_name       = "%s"
		pi_instance_name		= "%s"
		pi_capture_destination  = "%s"
		pi_capture_volume_ids = [data.ibm_pi_volume.dsvolume.id]

	}
	`, pi_cloud_instance_id, name, pi_instance_name, pi_capture_destination, pi_volume_name)
}

func testAccCheckIBMPICaptureConfigBasic(name string) string {
	return fmt.Sprintf(`

	resource "ibm_pi_capture" "capture_instance" {
		pi_cloud_instance_id="%[1]s"
		pi_capture_name       = "%s"
		pi_instance_name		= "%s"
		pi_capture_destination  = "%s"

	}
	`, pi_cloud_instance_id, name, pi_instance_name, pi_capture_destination)
}

func testAccCheckIBMPICaptureCloudStorageConfig(name string) string {
	return fmt.Sprintf(`
	data "ibm_pi_volume" "dsvolume" {
		pi_cloud_instance_id = "%[1]s"
		pi_volume_name      = "%[8]s"
	  }

	resource "ibm_pi_capture" "capture_instance" {
		pi_cloud_instance_id="%[1]s"
		pi_capture_name       = "%s"
		pi_instance_name		= "%s"
		pi_capture_destination  = "%s"
		pi_capture_volume_ids = [data.ibm_pi_volume.dsvolume.id]
		pi_capture_cloud_storage_region = "us-east"
		pi_capture_cloud_storage_access_key = "%s"
		pi_capture_cloud_storage_secret_key = "%s"
		pi_capture_storage_image_path = "%s"

	}
	`, pi_cloud_instance_id, name, pi_instance_name, pi_capture_destination, pi_capture_cloud_storage_access_key, pi_capture_cloud_storage_secret_key, pi_capture_storage_image_path, pi_volume_name)
}
