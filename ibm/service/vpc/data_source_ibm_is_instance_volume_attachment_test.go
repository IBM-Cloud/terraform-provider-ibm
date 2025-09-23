// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceVolumeAttDataSource_basic(t *testing.T) {

	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instance_volume_attachment.ds_vol_att"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentDataSourceConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						resName, "name"),
					resource.TestCheckResourceAttr(
						resName, "type", "boot"),
					resource.TestCheckResourceAttrSet(
						resName, "href"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceVolumeAttachmentDataSourceConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count = 16
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
	  data "ibm_is_instance_volume_attachment" "ds_vol_att" {
		instance = ibm_is_instance.testacc_instance.id
		name = ibm_is_instance.testacc_instance.volume_attachments.0.name
	  }
	 
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func TestAccIBMISInstanceVolumeAttDataSource_nodevice(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	action1 := "stop"
	action2 := "start"
	resName := "data.ibm_is_instance_volume_attachment.ds_vol_att"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentDataSourceNodeviceConfig(
					vpcname, subnetname, sshname, publicKey, name, action1,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "href"),
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth"),
					resource.TestCheckResourceAttr(resName, "type", "data"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeAttachmentDataSourceNodeviceConfig(
					vpcname, subnetname, sshname, publicKey, name, action2,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "name"),
					resource.TestCheckResourceAttrSet(resName, "href"),
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttrSet(resName, "bandwidth"),
					resource.TestCheckResourceAttrSet(resName, "device"),
					resource.TestCheckResourceAttr(resName, "type", "data"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceVolumeAttachmentDataSourceNodeviceConfig(vpcname, subnetname, sshname, publicKey, name, action string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
  name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name                     = "%s"
  vpc                      = ibm_is_vpc.testacc_vpc.id
  zone                     = "%s"
  total_ipv4_address_count = 16
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
    subnet = ibm_is_subnet.testacc_subnet.id
  }

  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "%s"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]

  // Add second (data) volume so .1 is valid
  volume_prototypes {
    name                             = "%s-data-vol"
    delete_volume_on_instance_delete = true
    volume_name                      = "%s-data-vol"
    volume_capacity                  = 20
    volume_profile                   = "general-purpose"
  }
}
resource "ibm_is_instance_action" "is_instance_action" {
	  depends_on = [ibm_is_instance.testacc_instance]
	  action = "%s"
	  instance = ibm_is_instance.testacc_instance.id
}
data "ibm_is_instance_volume_attachment" "ds_vol_att" {
  depends_on = [ibm_is_instance_action.is_instance_action]
  instance = ibm_is_instance.testacc_instance.id
  name     = ibm_is_instance.testacc_instance.volume_attachments.1.name
}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name,
		acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, name, name, action)
}
