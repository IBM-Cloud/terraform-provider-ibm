// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVolumeDatasource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	zone := "us-south-1"
	resName := "data.ibm_is_volume.testacc_dsvol"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceConfig(name, zone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name),
					resource.TestCheckResourceAttr(
						resName, "zone", zone),
					resource.TestCheckResourceAttrSet(
						resName, "active"),
					resource.TestCheckResourceAttrSet(
						resName, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "busy"),
					resource.TestCheckResourceAttrSet(
						resName, "created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resName, "profile"),
				),
			},
		},
	})
}
func TestAccIBMISVolumeDatasourceIdnetifier_basic(t *testing.T) {
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	zone := acc.ISZoneName
	resName := "data.ibm_is_volume.testacc_dsvol"
	resNameId := "data.ibm_is_volume.testacc_dsvolidentifier"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceWithIdentifierConfig(name, zone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name),
					resource.TestCheckResourceAttr(
						resName, "zone", zone),
					resource.TestCheckResourceAttrSet(
						resName, "active"),
					resource.TestCheckResourceAttrSet(
						resName, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "busy"),
					resource.TestCheckResourceAttrSet(
						resName, "created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resName, "profile"),
					resource.TestCheckResourceAttr(
						resNameId, "name", name),
					resource.TestCheckResourceAttr(
						resNameId, "zone", zone),
					resource.TestCheckResourceAttrSet(
						resNameId, "active"),
					resource.TestCheckResourceAttrSet(
						resNameId, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resNameId, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resNameId, "busy"),
					resource.TestCheckResourceAttrSet(
						resNameId, "created_at"),
					resource.TestCheckResourceAttrSet(
						resNameId, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resNameId, "profile"),
				),
			},
		},
	})
}
func TestAccIBMISVolumeDatasource_Sdp(t *testing.T) {
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	zone := "eu-gb-1"
	resName := "data.ibm_is_volume.testacc_dsvol"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceSdpConfig(name, zone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", name),
					resource.TestCheckResourceAttr(
						resName, "zone", zone),
					resource.TestCheckResourceAttrSet(
						resName, "active"),
					resource.TestCheckResourceAttrSet(
						resName, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "busy"),
					resource.TestCheckResourceAttrSet(
						resName, "created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resName, "profile"),
					resource.TestCheckResourceAttrSet(
						resName, "adjustable_capacity_states.#"),
					resource.TestCheckResourceAttrSet(
						resName, "adjustable_iops_states.#"),
					resource.TestCheckResourceAttr(
						resName, "profile", "sdp"),
				),
			},
		},
	})
}
func TestAccIBMISVolumeDatasource_from_snapshot(t *testing.T) {

	resName := "data.ibm_is_volume.testacc_dsvol"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceFromSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "active"),
					resource.TestCheckResourceAttrSet(
						resName, "attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "busy"),
					resource.TestCheckResourceAttrSet(
						resName, "created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "resource_group"),
					resource.TestCheckResourceAttrSet(
						resName, "profile"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.#"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.architecture"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.display_name"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.dedicated_host_only"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.family"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.name"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.vendor"),
					resource.TestCheckResourceAttrSet(
						resName, "operating_system.0.version"),
				),
			},
		},
	})
}
func TestAccIBMISVolumeDatasourceWithCatalogOffering(t *testing.T) {

	resName := "data.ibm_is_volume.testacc_dsvol"
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	planCrn := "crn:v1:staging:public:globalcatalog-collection:global::1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:plan:sw.1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.279a3cee-ba7d-42d5-ae88-6a0ebc56fa4a-global"
	versionCrn := "crn:v1:staging:public:globalcatalog-collection:global::1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:4f8466eb-2218-42e3-a755-bf352b559c69-global/6a73aa69-5dd9-4243-a908-3b62f467cbf8-global"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVolumeDataSourceWithCatalogOffering(vpcname, subnetname, sshname, publicKey, name, planCrn, versionCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "catalog_offering.#"),
					resource.TestCheckResourceAttrSet(
						resName, "catalog_offering.0.plan_crn"),
					resource.TestCheckResourceAttrSet(
						resName, "catalog_offering.0.version_crn"),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeDataSourceFromSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1 string) string {
	return testAccCheckIBMISVolumeConfigSnapshot(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(`
	
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.storage.name
	}`)
}
func testAccCheckIBMISVolumeDataSourceConfig(name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_volume" "testacc_volume"{
		name = "%s"
		profile = "10iops-tier"
		zone = "%s"
	}
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.testacc_volume.name
	}`, name, zone)
}
func testAccCheckIBMISVolumeDataSourceWithIdentifierConfig(name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_volume" "testacc_volume"{
		name = "%s"
		profile = "10iops-tier"
		zone = "%s"
	}
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.testacc_volume.name
	}
	data "ibm_is_volume" "testacc_dsvolidentifier" {
		identifier = ibm_is_volume.testacc_volume.id
	}
		
	`, name, zone)
}
func testAccCheckIBMISVolumeDataSourceSdpConfig(name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_volume" "testacc_volume"{
		name 		= "%s"
		profile 	= "sdp"
		zone 		= "%s"
	}
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_volume.testacc_volume.name
	}`, name, zone)
}

func testAccCheckIBMISVolumeDataSourceWithCatalogOffering(vpcname, subnetname, sshname, publicKey, name, planCrn, versionCrn string) string {
	return fmt.Sprintf(`
	
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name           				= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  } 
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		boot_volume {
		  auto_delete_volume = false  
		}
		catalog_offering {
		  version_crn = "%s"
		  plan_crn    = "%s"
		}
	  }
	
	data "ibm_is_volume" "testacc_dsvol" {
		name = ibm_is_instance.testacc_instance.boot_volume.0.name
	}`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.InstanceProfileName, acc.ISZoneName, versionCrn, planCrn)
}
