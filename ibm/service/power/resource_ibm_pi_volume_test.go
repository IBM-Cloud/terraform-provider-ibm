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

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIVolumebasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	volumeRes := "ibm_pi_volume.power_volume"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists(volumeRes),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_name", name),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeSizeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists(volumeRes),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_name", name),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_size", "30"),
				),
			},
		},
	})
}

func TestAccIBMPIVolumeDeleteWithTargetCRN(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	volumeRes := "ibm_pi_volume.power_volume"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeDeleteWithTargetCRNConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists(volumeRes),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_name", name),
					resource.TestCheckResourceAttr(volumeRes, "pi_replication_enabled", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeDeleteWithTargetCRNConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id    = "%[2]s"
			pi_volume_name          = "%[1]s"
			pi_volume_pool          = "%[3]s"
			pi_volume_shareable     = true
			pi_volume_size          = 20
			pi_volume_type          = "tier3"
			pi_replication_enabled  = true
			pi_target_crn           = "%[4]s"
		}`, name, acc.Pi_cloud_instance_id, acc.PiStoragePool, acc.Pi_target_crn)
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
		volumeC := instance.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)
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
		client := instance.NewIBMPIVolumeClient(context.Background(), sess, cloudInstanceID)

		_, err = client.Get(volumeID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPIVolumeConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id	= "%[2]s"
			pi_volume_name       	= "%[1]s"
			pi_volume_shareable  	= true
			pi_volume_size       	= 20
			pi_volume_type       	= "tier1"
		}`, name, acc.Pi_cloud_instance_id)
}

func testAccCheckIBMPIVolumeSizeConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id	= "%[2]s"
			pi_volume_name       	= "%[1]s"
			pi_volume_shareable  	= true
			pi_volume_size       	= 30
			pi_volume_type       	= "tier1"
		}`, name, acc.Pi_cloud_instance_id)
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
						"ibm_pi_volume.power_volume", "pi_volume_pool", acc.PiStoragePool),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumePoolConfig(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id	= "%[2]s"
			pi_volume_name       	= "%[1]s"
			pi_volume_pool       	= "%[3]s"
			pi_volume_shareable		= true
			pi_volume_size       	= 20
		}`, name, acc.Pi_cloud_instance_id, acc.PiStoragePool)
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
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id   	= "%[2]s"
			pi_replication_enabled	= %[4]v
			pi_volume_name         	= "%[1]s"
			pi_volume_pool         	= "%[3]s"
			pi_volume_shareable    	= true
			pi_volume_size         	= 20
			pi_volume_type         	= "tier3"
		}`, name, piCloudInstanceId, piStoragePool, replicationEnabled)
}

// TestAccIBMPIVolumeUpdate test the volume update
func TestAccIBMPIVolumeUpdate(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	sType := acc.PiStorageType // tier 3
	sTypeUpdate := "tier1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeUpdateStorageConfig(name, sType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeUpdateStorageConfig(name, sTypeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists("ibm_pi_volume.power_volume"),
					resource.TestCheckResourceAttr(
						"ibm_pi_volume.power_volume", "pi_volume_name", name),
					resource.TestCheckResourceAttrSet("ibm_pi_volume.power_volume", "pi_volume_type"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeUpdateStorageConfig(name, piStorageType string) string {
	return testAccCheckIBMPIVolumeUpdateBasicConfig(name, acc.Pi_cloud_instance_id, acc.PiStoragePool, piStorageType)
}

func testAccCheckIBMPIVolumeUpdateBasicConfig(name, piCloudInstanceId, piStoragePool, piStorageType string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id	= "%[2]s"
			pi_volume_name         	= "%[1]s"
			pi_volume_pool         	= "%[3]s"
			pi_volume_shareable    	= true
			pi_volume_size         	= 20
			pi_volume_type       	= "%[4]v"
		}`, name, piCloudInstanceId, piStoragePool, piStorageType)
}

func TestAccIBMPIVolumeUserTags(t *testing.T) {
	name := fmt.Sprintf("tf-pi-volume-%d", acctest.RandIntRange(10, 100))
	volumeRes := "ibm_pi_volume.power_volume"
	userTagsString := `["env:dev","test_tag"]`
	userTagsStringUpdated := `["env:dev","test_tag","test_tag2"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIVolumeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeUserTagsConfig(name, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists(volumeRes),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_name", name),
					resource.TestCheckResourceAttr(volumeRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPIVolumeUserTagsConfig(name, userTagsStringUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIVolumeExists(volumeRes),
					resource.TestCheckResourceAttr(volumeRes, "pi_volume_name", name),
					resource.TestCheckResourceAttr(volumeRes, "pi_user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(volumeRes, "pi_user_tags.*", "test_tag2"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeUserTagsConfig(name string, userTagsString string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_volume" "power_volume" {
			pi_cloud_instance_id	= "%[2]s"
			pi_volume_name       	= "%[1]s"
			pi_volume_shareable  	= true
			pi_volume_size       	= 20
			pi_volume_type       	= "tier1"
			pi_user_tags            = %[3]s
		}`, name, acc.Pi_cloud_instance_id, userTagsString)
}
