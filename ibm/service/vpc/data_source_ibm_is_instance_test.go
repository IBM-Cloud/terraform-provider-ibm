// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
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
