// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstance_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISInstance_profile(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigWithProfile(vpcname, subnetname, sshname, publicKey, name, instanceProfileName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", instanceProfileName),
				),
			},

			{
				Config: testAccCheckIBMISInstanceConfigWithProfile(vpcname, subnetname, sshname, publicKey, name, instanceProfileNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", instanceProfileNameUpdate),
				),
			},
		},
	})
}

func TestAccIBMISInstance_basicwithipv4(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	ipv4address := "10.240.0.6"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigwithipv4(vpcname, subnetname, sshname, publicKey, name, ipv4address),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.primary_ipv4_address", ipv4address),
				),
			},
		},
	})
}

func TestAccIBMISInstance_VolumeAutoDelete(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAutoDelete(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "auto_delete_volume", "true"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Volume(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolume(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeUpdate(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolume(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Disk(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDisk(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "disks.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "disks.0.name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "disks.0.size"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Placement(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	dhostname := fmt.Sprintf("tf-dhost-%d", acctest.RandIntRange(10, 100))
	dhostgrpname := fmt.Sprintf("tf-dhostgrp-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstancePlacement(vpcname, subnetname, sshname, publicKey, volname, name, dhostname, dhostgrpname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_dedicated_host.dhost", "instances.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_dedicated_host.dhost", "instances.0.name", name),
				),
			},
		},
	})
}

func TestAccIBMISInstanceSnapshotRestore_basic(t *testing.T) {
	var instance, instanceRestore string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	snapshot := fmt.Sprintf("tf-snapshot-%d", acctest.RandIntRange(10, 100))
	vsiRestore := fmt.Sprintf("tf-instancerestore-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceSnapshotRestoreConfig(vpcname, subnetname, sshname, publicKey, name, snapshot, vsiRestore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_restore", instanceRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "zone", ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "name", vsiRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "boot_volume.0.name", "boot-restore"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceDestroy(s *terraform.State) error {

	instanceC, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance" {
			continue
		}
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := instanceC.GetInstance(getinsOptions)

		if err == nil {
			return fmt.Errorf("instance still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISInstanceExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		instanceC, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &rs.Primary.ID,
		}
		foundins, _, err := instanceC.GetInstance(getinsOptions)
		if err != nil {
			return err
		}
		instance = *foundins.ID
		return nil
	}
}

func testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, name, isImage, instanceProfileName, ISZoneName)
}
func testAccCheckIBMISInstanceSnapshotRestoreConfig(vpcname, subnetname, sshname, publicKey, name, snapshot, insRestore string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
	  }

	  resource "ibm_is_instance" "testacc_instance_restore" {
		name    = "%s"
		profile = "%s"
		boot_volume {
			name     = "boot-restore"
			snapshot = ibm_is_snapshot.testacc_snapshot.id
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  `, vpcname, subnetname, ISZoneName, sshname, publicKey, name, isImage, instanceProfileName, ISZoneName, snapshot, insRestore, instanceProfileName, ISZoneName)
}

func testAccCheckIBMISInstanceConfigWithProfile(vpcname, subnetname, sshname, publicKey, name, isInstanceProfileName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, name, isImage, isInstanceProfileName, ISZoneName)
}

func testAccCheckIBMISInstanceConfigwithipv4(vpcname, subnetname, sshname, publicKey, name, ipv4address string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		  primary_ipv4_address = "%s"
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, name, isImage, instanceProfileName, ipv4address, ISZoneName)
}

func testAccCheckIBMISInstanceVolume(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
		# capacity= 200
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		volumes = [ibm_is_volume.storage.id]
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, volName, ISZoneName, name, isImage, instanceProfileName, ISZoneName)
}

func testAccCheckIBMISInstanceDisk(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
		# capacity= 200
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		volumes = [ibm_is_volume.storage.id]
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, volName, ISZoneName, name, isImage, instanceDiskProfileName, ISZoneName)
}

func testAccCheckIBMISInstancePlacement(vpcname, subnetname, sshname, publicKey, volName, name, dhostname, dhostgrpname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		dedicated_host_group = "%s"
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		
	  }
	  data "ibm_is_dedicated_host" "dhost"{
		  host_group = ibm_is_instance.testacc_instance.dedicated_host_group
		  name 		 = "%s"
	  } 
	 
	  `, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, volName, ISZoneName, name, isImage, instanceProfileName, dedicatedHostGroupID, ISZoneName, dedicatedHostName)
}

func testAccCheckIBMISInstanceVolumeAutoDelete(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     			= ibm_is_vpc.testacc_vpc.id
		zone    			= "%s"
		keys    			= [ibm_is_ssh_key.testacc_sshkey.id]
		volumes 			= [ibm_is_volume.storage.id]
		auto_delete_volume 	= true
	  }`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, volName, ISZoneName, name, isImage, instanceProfileName, ISZoneName)
}

func testAccCheckIBMISInstanceVolumeUpdate(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }	  
	
`, vpcname, subnetname, ISZoneName, ISCIDR, sshname, publicKey, volName, ISZoneName, name, isImage, instanceProfileName, ISZoneName)
}
