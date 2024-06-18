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

func TestAccIBMISInstanceDataSource_basic(t *testing.T) {

	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instance.ds_instance"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDataSourceConfig(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", instanceName),
					resource.TestCheckResourceAttr(
						resName, "tags.#", "1"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						resName, "availability_policy_host_failure"),
					resource.TestCheckResourceAttrSet(
						resName, "lifecycle_state"),
					resource.TestCheckResourceAttr(
						resName, "lifecycle_reasons.#", "0"),
					resource.TestCheckResourceAttrSet(
						resName, "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						resName, "vcpu.0.manufacturer"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceDataSourceWithCatalogOffering(t *testing.T) {

	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instance.ds_instance"
	planCrn := "crn:v1:staging:public:globalcatalog-collection:global::1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:plan:sw.1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.279a3cee-ba7d-42d5-ae88-6a0ebc56fa4a-global"
	versionCrn := "crn:v1:staging:public:globalcatalog-collection:global::1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:4f8466eb-2218-42e3-a755-bf352b559c69-global/6a73aa69-5dd9-4243-a908-3b62f467cbf8-global"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDataSourceConfigWithCatalogOffering(vpcname, subnetname, sshname, instanceName, planCrn, versionCrn),
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
func TestAccIBMISInstanceDataSource_vni(t *testing.T) {

	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfins-vni-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instance.ds_instance"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVniDataSourceConfig(vpcname, subnetname, sshname, vniname, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", instanceName),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						resName, "availability_policy_host_failure"),
					resource.TestCheckResourceAttrSet(
						resName, "lifecycle_state"),
					resource.TestCheckResourceAttr(
						resName, "lifecycle_reasons.#", "0"),
					resource.TestCheckResourceAttrSet(
						resName, "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						resName, "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_attachment.#"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_attachment.0.id"),
					resource.TestCheckResourceAttr(
						resName, "primary_network_attachment.0.name", "test-vni"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_attachment.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_attachment.0.subnet.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceDataSource_PKCS8SSH(t *testing.T) {

	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instance.ds_instance"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDataSourcePKCS8SSHConfig(vpcname, subnetname, sshname, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", instanceName),
					resource.TestCheckResourceAttr(
						resName, "tags.#", "1"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						resName, "availability_policy_host_failure"),
					resource.TestCheckResourceAttrSet(
						resName, "lifecycle_state"),
					resource.TestCheckResourceAttr(
						resName, "lifecycle_reasons.#", "0"),
					resource.TestCheckResourceAttrSet(
						resName, "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						resName, "vcpu.0.manufacturer"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceDataSource_reserved_ip(t *testing.T) {

	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instance.ds_instance"
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDataSourceReservedIpConfig(vpcname, subnetname, sshname, publicKey, instanceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resName, "name", instanceName),
					resource.TestCheckResourceAttr(
						resName, "tags.#", "1"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.primary_ip.0.resource_type"),
					resource.TestCheckResourceAttrSet(
						resName, "primary_network_interface.0.primary_ip.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceDataSourceConfig(vpcname, subnetname, sshname, instanceName string) string {
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
  public_key = file("./test-fixtures/.ssh/id_rsa.pub")
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
  tags = ["tag1"]
}
data "ibm_is_instance" "ds_instance" {
  name        = ibm_is_instance.testacc_instance.name
  private_key = file("./test-fixtures/.ssh/id_rsa")
  passphrase  = ""
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, instanceName, acc.IsWinImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceDataSourceConfigWithCatalogOffering(vpcname, subnetname, sshname, instanceName, planCrn, versionCrn string) string {
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
		public_key = file("./test-fixtures/.ssh/id_rsa.pub")
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
data "ibm_is_instance" "ds_instance" {
  name        = ibm_is_instance.testacc_instance.name
}`, vpcname, subnetname, acc.ISZoneName, sshname, instanceName, acc.InstanceProfileName, acc.ISZoneName, versionCrn, planCrn)
}
func testAccCheckIBMISInstanceVniDataSourceConfig(vpcname, subnetname, sshname, vniname, instanceName string) string {
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
  public_key = file("./test-fixtures/.ssh/id_rsa.pub")
}

resource "ibm_is_virtual_network_interface" "testacc_vni"{
	name = "%s"
	allow_ip_spoofing = true
	subnet = ibm_is_subnet.testacc_subnet.id
} 

resource "ibm_is_instance" "testacc_instance" {
  name    = "%s"
  image   = "%s"
  profile = "%s"
  primary_network_attachment {
	name = "test-vni"
	virtual_network_interface { 
		id = ibm_is_virtual_network_interface.testacc_vni.id
	}
  }
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "%s"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
}
data "ibm_is_instance" "ds_instance" {
  name        = ibm_is_instance.testacc_instance.name
  private_key = file("./test-fixtures/.ssh/id_rsa")
  passphrase  = ""
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, vniname, instanceName, acc.IsWinImage, acc.InstanceProfileName, acc.ISZoneName)
}
func testAccCheckIBMISInstanceDataSourcePKCS8SSHConfig(vpcname, subnetname, sshname, instanceName string) string {
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
		public_key = file("%s")
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
		tags = ["tag1"]
	}
	data "ibm_is_instance" "ds_instance" {
		name        = ibm_is_instance.testacc_instance.name
		private_key = file("%s")
		passphrase  = ""
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, acc.ISPublicSSHKeyFilePath, instanceName, acc.IsWinImage, acc.InstanceProfileName, acc.ISZoneName, acc.ISPrivateSSHKeyFilePath)
}

func testAccCheckIBMISInstanceDataSourceReservedIpConfig(vpcname, subnetname, sshname, publicKey, instanceName string) string {
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
  tags = ["tag1"]
}
data "ibm_is_instance" "ds_instance" {
  name        = ibm_is_instance.testacc_instance.name
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, instanceName, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}
