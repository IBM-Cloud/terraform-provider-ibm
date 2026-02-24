// Copyright IBM Corp. 2017, 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstanceBootVolumeManager_basic(t *testing.T) {
	var volumeID string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "boot_volume"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "bandwidth"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "capacity"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "encryption_type"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "health_state"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "iops"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "profile"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "zone"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "status"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "resource_group.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "volume_attachments.#"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_name_update(t *testing.T) {
	var volumeID string
	name1 := fmt.Sprintf("tfbootvol-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfbootvol-updated-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name", name1),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name", name2),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_capacity_update(t *testing.T) {
	var volumeID string
	capacity1 := 120
	capacity2 := 180

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "capacity"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithCapacity(capacity1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "capacity", fmt.Sprintf("%d", capacity1)),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithCapacity(capacity2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "capacity", fmt.Sprintf("%d", capacity2)),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_profile_update(t *testing.T) {
	var volumeID string
	profile := "10iops-tier"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "profile"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithProfile(profile),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "profile", profile),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_iops_update(t *testing.T) {
	var volumeID string
	iops1 := 600
	iops2 := 900

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithIOPS("custom", iops1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "iops", fmt.Sprintf("%d", iops1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "profile", "custom"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithIOPS("custom", iops2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "iops", fmt.Sprintf("%d", iops2)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "profile", "custom"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_tags(t *testing.T) {
	var volumeID string
	tag1 := "env:test"
	tag2 := "team:dev"
	tag3 := "boot:managed"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithTags(tag1, tag2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithTags(tag1, tag2, tag3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "tags.#", "3"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_access_tags(t *testing.T) {
	var volumeID string
	accessTag := "project:test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithAccessTags(accessTag),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "access_tags.#", "1"),
				),
			},
			{
				ResourceName:            "ibm_is_instance_boot_volume_manager.test_boot_volume",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"delete_volume", "delete_all_snapshots"},
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_complete_update(t *testing.T) {
	var volumeID string
	name := fmt.Sprintf("tfbootvol-%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("tfbootvol-updated-%d", acctest.RandIntRange(10, 100))
	profile := "10iops-tier"
	capacity := 200
	tag1 := "env:prod"
	tag2 := "boot:managed"
	accessTag := "project:production"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerCompleteConfig(name, profile, capacity, tag1, tag2, accessTag),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "profile", profile),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "capacity", fmt.Sprintf("%d", capacity)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "tags.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "access_tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerCompleteConfig(updatedName, profile, capacity+50, tag1, tag2, accessTag),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name", updatedName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "capacity", fmt.Sprintf("%d", capacity+50)),
				),
			},
		},
	})
}

func TestAccIBMISInstanceBootVolumeManager_delete_volume(t *testing.T) {
	var volumeID string
	vpcName := fmt.Sprintf("tfvpc-%d", acctest.RandIntRange(10, 100))
	subnetName := fmt.Sprintf("tfsubnet-%d", acctest.RandIntRange(10, 100))
	sshName := fmt.Sprintf("tfssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfinstance-%d", acctest.RandIntRange(10, 100))
	bootVolumeName := fmt.Sprintf("tfbootvol-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceBootVolumeManagerDestroyWithDeletion,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeManagerConfigWithDeletion(vpcName, subnetName, sshName, publicKey, instanceName, bootVolumeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceBootVolumeManagerExists("ibm_is_instance_boot_volume_manager.test_boot_volume", volumeID),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "name", bootVolumeName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "delete_volume", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_boot_volume_manager.test_boot_volume", "delete_all_snapshots", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceBootVolumeManagerDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_boot_volume_manager" {
			continue
		}

		getVolumeOptions := &vpcv1.GetVolumeOptions{
			ID: &rs.Primary.ID,
		}
		_, response, err := sess.GetVolume(getVolumeOptions)

		if err == nil {
			// Volume still exists, which is expected if delete_volume was false
			continue
		}
		if response != nil && response.StatusCode != 404 {
			return fmt.Errorf("Error checking for boot volume (%s) deletion: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceBootVolumeManagerDestroyWithDeletion(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_boot_volume_manager" {
			continue
		}

		getVolumeOptions := &vpcv1.GetVolumeOptions{
			ID: &rs.Primary.ID,
		}
		_, response, err := sess.GetVolume(getVolumeOptions)

		if err == nil {
			return fmt.Errorf("Boot volume still exists: %s", rs.Primary.ID)
		}
		if response != nil && response.StatusCode != 404 {
			return fmt.Errorf("Error checking for boot volume (%s) deletion: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceBootVolumeManagerExists(n, volumeID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getVolumeOptions := &vpcv1.GetVolumeOptions{
			ID: &rs.Primary.ID,
		}
		foundVolume, _, err := sess.GetVolume(getVolumeOptions)
		if err != nil {
			return err
		}
		volumeID = *foundVolume.ID
		return nil
	}
}

func testAccCheckIBMISInstanceBootVolumeManagerConfig() string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
}`, acc.VSIUnattachedBootVolumeID)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithName(name string) string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	name        = "%s"
}`, acc.VSIUnattachedBootVolumeID, name)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithCapacity(capacity int) string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	capacity    = %d
}`, acc.VSIUnattachedBootVolumeID, capacity)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithProfile(profile string) string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	profile     = "%s"
}`, acc.VSIUnattachedBootVolumeID, profile)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithIOPS(profile string, iops int) string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	profile     = "%s"
	iops        = %d
}`, acc.VSIUnattachedBootVolumeID, profile, iops)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithTags(tags ...string) string {
	tagList := `["` + strings.Join(tags, `", "`) + `"]`
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	tags        = %s
}`, acc.VSIUnattachedBootVolumeID, tagList)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithAccessTags(accessTag string) string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	access_tags = ["%s"]
}`, acc.VSIUnattachedBootVolumeID, accessTag)
}

func testAccCheckIBMISInstanceBootVolumeManagerCompleteConfig(name, profile string, capacity int, tag1, tag2, accessTag string) string {
	return fmt.Sprintf(`
resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume = "%s"
	name        = "%s"
	profile     = "%s"
	capacity    = %d
	tags        = ["%s", "%s"]
	access_tags = ["%s"]
}`, acc.VSIUnattachedBootVolumeID, name, profile, capacity, tag1, tag2, accessTag)
}

func testAccCheckIBMISInstanceBootVolumeManagerConfigWithDeletion(vpcName, subnetName, sshName, publicKey, instanceName, bootVolumeName string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "test_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "test_subnet" {
	name                     = "%s"
	vpc                      = ibm_is_vpc.test_vpc.id
	zone                     = "%s"
	total_ipv4_address_count = 16
}

resource "ibm_is_ssh_key" "test_ssh_key" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_volume" "test_boot_volume" {
	name     = "%s"
	profile  = "general-purpose"
	zone     = "%s"
	capacity = 100
}

resource "ibm_is_instance" "test_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
	
	boot_volume {
		volume_id           = ibm_is_instance_boot_volume_manager.test_boot_volume.id
		auto_delete_volume  = false
	}
	
	primary_network_interface {
		subnet = ibm_is_subnet.test_subnet.id
	}
	
	vpc  = ibm_is_vpc.test_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.test_ssh_key.id]
}

resource "ibm_is_instance_boot_volume_manager" "test_boot_volume" {
	boot_volume           = ibm_is_volume.test_boot_volume.id
	name                  = "%s"
	delete_volume         = true
	delete_all_snapshots  = true
}`, vpcName, subnetName, acc.ISZoneName, sshName, publicKey, bootVolumeName, acc.ISZoneName, instanceName, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, bootVolumeName)
}
