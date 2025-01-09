// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package sdsaas_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/sdsaas"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/sds-go-sdk/sdsaasv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMSdsVolumeBasic(t *testing.T) {
	var conf sdsaasv1.Volume
	capacity := fmt.Sprintf("%d", acctest.RandIntRange(1, 5))
	name := "terraform-test-1"
	nameUpdate := "terraform-test-name-updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSdsVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfigBasic(capacity, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSdsVolumeExists("ibm_sds_volume.sds_volume_instance", conf),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfigBasic(capacity, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMSdsVolumeAllArgs(t *testing.T) {
	var conf sdsaasv1.Volume
	hostnqnstring := "nqn.2014-06.org:9345"
	capacity := fmt.Sprintf("%d", acctest.RandIntRange(1, 5))
	name := "terraform-test-name-1"
	nameUpdate := "terraform-test-name-updated"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSdsVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfig(hostnqnstring, capacity, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSdsVolumeExists("ibm_sds_volume.sds_volume_instance", conf),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "hostnqnstring", hostnqnstring),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSdsVolumeConfig(hostnqnstring, capacity, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "hostnqnstring", hostnqnstring),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "capacity", capacity),
					resource.TestCheckResourceAttr("ibm_sds_volume.sds_volume_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName: "ibm_sds_volume.sds_volume_instance",
				ImportState:  true,
			},
		},
	})
}

func testAccCheckIBMSdsVolumeConfigBasic(capacity string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_sds_volume" "sds_volume_instance" {
			capacity = %s
			name = "%s"
		}
	`, capacity, name)
}

func testAccCheckIBMSdsVolumeConfig(hostnqnstring string, capacity string, name string) string {
	return fmt.Sprintf(`

		resource "ibm_sds_volume" "sds_volume_instance" {
			hostnqnstring = "%s"
			capacity = %s
			name = "%s"
		}
	`, hostnqnstring, capacity, name)
}

func testAccCheckIBMSdsVolumeExists(n string, obj sdsaasv1.Volume) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sdsaasClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SdsaasV1()
		if err != nil {
			return err
		}

		volumeOptions := &sdsaasv1.VolumeOptions{}

		volumeOptions.SetVolumeID(rs.Primary.ID)

		volume, _, err := sdsaasClient.Volume(volumeOptions)
		if err != nil {
			return err
		}

		obj = *volume
		return nil
	}
}

func testAccCheckIBMSdsVolumeDestroy(s *terraform.State) error {
	sdsaasClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SdsaasV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sds_volume" {
			continue
		}

		volumeOptions := &sdsaasv1.VolumeOptions{}

		volumeOptions.SetVolumeID(rs.Primary.ID)

		// Give time for the volume to fully delete before checking its state
		time.Sleep(5 * time.Second)

		// Try to find the key
		_, response, err := sdsaasClient.Volume(volumeOptions)

		if err == nil {
			return fmt.Errorf("sds_volume still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for sds_volume (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMSdsVolumeHostMappingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["host_id"] = "testString"
		model["host_name"] = "testString"
		model["host_nqn"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(sdsaasv1.HostMapping)
	model.HostID = core.StringPtr("testString")
	model.HostName = core.StringPtr("testString")
	model.HostNqn = core.StringPtr("testString")

	result, err := sdsaas.ResourceIBMSdsVolumeHostMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
