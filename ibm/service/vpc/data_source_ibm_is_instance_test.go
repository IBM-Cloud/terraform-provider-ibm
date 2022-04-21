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
						resName, "availability_policy.#"),
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
  public_key = file("../../test-fixtures/.ssh/id_rsa.pub")
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
  private_key = file("../../test-fixtures/.ssh/id_rsa")
  passphrase  = ""
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, instanceName, acc.IsWinImage, acc.InstanceProfileName, acc.ISZoneName)
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
