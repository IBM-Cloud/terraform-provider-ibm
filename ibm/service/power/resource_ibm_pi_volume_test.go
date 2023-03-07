// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIVolumebasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeSizeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_size", "30"),
				),
			},
		},
	})
}
func testAccCheckIBMPIVolumeDestroy(s *terraform.State) error {

	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_volume" {
			continue
		}
		cloudInstanceID, volumeID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		volumeC := st.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)
		volume, err := volumeC.Get(volumeID)
		if err == nil {
			log.Println("volume*****", volume.State)
			return fmt.Errorf("PI Volume still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIVolumeExists(n string) resource.TestCheckFunc {
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
		cloudInstanceID, volumeID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		client := st.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(volumeID)
		if err != nil {
			return err
		}
		return nil

	}
}

func testAccCheckIBMPIVolumeConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume"{
		pi_volume_size       = 20
		pi_volume_name       = "%s"
		pi_volume_type       = "tier1"
		pi_volume_shareable  = true
		pi_cloud_instance_id = "%s"
	  }
	`, name, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumeSizeConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume"{
		pi_volume_size       = 30
		pi_volume_name       = "%s"
		pi_volume_type       = "tier1"
		pi_volume_shareable  = true
		pi_cloud_instance_id = "%s"
	  }
	`, name, acc.Pi_cloud_instance_id)
}

func TestAccIBMPIVolumePool(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumePoolConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_pool", "Tier3-Flash-1"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumePoolConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume"{
		pi_volume_size       = 20
		pi_volume_name       = "%s"
		pi_volume_pool       = "Tier3-Flash-1"
		pi_volume_shareable  = true
		pi_cloud_instance_id = "%s"
	  }
	`, name, acc.Pi_cloud_instance_id)
}

// TestAccIBMPIVolumeGRS test the volume replication feature which is part of global replication service(GRS)
func TestAccIBMPIVolumeGRS(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeGRSConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_replication_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "replication_type", "global"),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeGRSUpdateConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_replication_enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "replication_type", ""),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeGRSConfig(name string) string {
	return testAccCheckIBMPIVolumeGRSBasicConfig(name, acc.Pi_cloud_instance_id, acc.PiStoragePool, true)
}

func testAccCheckIBMPIVolumeGRSUpdateConfig(name string) string {
	return testAccCheckIBMPIVolumeGRSBasicConfig(name, acc.Pi_cloud_instance_id, acc.PiStoragePool, false)
}

func testAccCheckIBMPIVolumeGRSBasicConfig(name, piCloudInstanceId, piStoragePool string, replicationEnabled bool) string {
	return fmt.Sprintf(`
	resource "ibm_pi_volume" "power_volume"{
		pi_volume_size         = 20
		pi_volume_name         = "%[1]s"
		pi_volume_pool         = "%[3]s"
		pi_volume_shareable    = true
		pi_cloud_instance_id   = "%[2]s"
		pi_replication_enabled = %[4]v
	  }
	`, name, piCloudInstanceId, piStoragePool, replicationEnabled)
}
