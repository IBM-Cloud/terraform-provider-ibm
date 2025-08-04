// Copyright IBM Corp. 2021 All Rights Reserved.
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

func TestAccIBMIsVolumesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVolumesDataSourceConfigFilterByZone(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
			{
				Config: testAccCheckIBMIsVolumesDataSourceConfigFilterByName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
				),
			},
		},
	})
}
func TestAccIBMIsVolumesDataSourceSdpBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesDataSourceSdpConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.0.adjustable_iops_states.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.0.adjustable_capacity_states.#"),
					resource.TestCheckResourceAttr("data.ibm_is_volumes.is_volumes", "volumes.0.profile.0.name", "sdp"),
				),
			},
		},
	})
}
func TestAccIBMIsVolumesFromSnapshotDataSourceBasic(t *testing.T) {
	resName := "data.ibm_is_volumes.is_volumes"
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
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesFromSnapshotDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.active"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.attachment_state"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.bandwidth"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.busy"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.created_at"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.resource_group.0.id"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.profile.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.architecture"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.display_name"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.dedicated_host_only"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.family"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.name"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.vendor"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.operating_system.0.version"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.allowed_use.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.allowed_use.0.api_version"),
				),
			},
		},
	})
}

func TestAccIBMIsVolumesFromSnapshotDataSourceWithCatalogOffering(t *testing.T) {
	resName := "data.ibm_is_volumes.is_volumes"
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
			resource.TestStep{
				Config: testAccCheckIBMIsVolumesWithCatalogOffering(vpcname, subnetname, sshname, publicKey, name, planCrn, versionCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_volumes.is_volumes", "volumes.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.catalog_offering.#"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.catalog_offering.0.plan_crn"),
					resource.TestCheckResourceAttrSet(
						resName, "volumes.0.catalog_offering.0.version_crn"),
				),
			},
		},
	})
}
func testAccCheckIBMIsVolumesFromSnapshotDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, volname, name, name1 string) string {
	return testAccCheckIBMISVolumeConfigSnapshot(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			volume_name = ibm_is_volume.storage.name
			attachment_state = "unattached"
			encryption = "provider_managed"
			operating_system_family = "not:null"
			operating_system_architecture = "not:null"
		}
	`)
}
func testAccCheckIBMIsVolumesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
		}
	`)
}
func testAccCheckIBMIsVolumesDataSourceSdpConfig() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
		}
	`)
}

func testAccCheckIBMIsVolumesDataSourceConfigFilterByZone() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			zone_name = "us-south-1"
		}
	`)
}

func testAccCheckIBMIsVolumesDataSourceConfigFilterByName() string {
	return fmt.Sprintf(`
		data "ibm_is_volumes" "is_volumes" {
			volume_name = "worrier-mailable-timpani-scowling"
		}
	`)
}

func testAccCheckIBMIsVolumesWithCatalogOffering(vpcname, subnetname, sshname, publicKey, name, planCrn, versionCrn string) string {
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
	  data "ibm_is_volumes" "is_volumes" {
		volume_name = ibm_is_instance.testacc_instance.boot_volume.0.name
	}`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.InstanceProfileName, acc.ISZoneName, versionCrn, planCrn)
}
