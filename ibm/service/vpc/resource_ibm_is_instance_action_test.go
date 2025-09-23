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

func TestAccIBMISInstanceAction_basic(t *testing.T) {
	//var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceActionRebootConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
				),
			},
			{
				Config: testAccCheckIBMISInstanceActionCheckStatusConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "status", "running"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceActionStopConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
				),
			},
			{
				Config: testAccCheckIBMISInstanceActionCheckStatusConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "status", "stopped"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceActionRebootConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "im_images" {
  
	}
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
    	image   = data.ibm_is_images.im_images.images.4.id
    	profile = "bx2d-16x64"
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
	
    resource "ibm_is_instance_action" "testacc_instanceaction" {
	  depends_on = [ibm_is_instance.testacc_instance]
	  action = "reboot"
	  instance = ibm_is_instance.testacc_instance.id
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.ISZoneName)
}

func testAccCheckIBMISInstanceActionStopConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "im_images" {
  
	}
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
    	image   = data.ibm_is_images.im_images.images.4.id
    	profile = "bx2d-16x64"
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
	
    resource "ibm_is_instance_action" "testacc_instanceaction" {
	  depends_on = [ibm_is_instance.testacc_instance]
	  action = "stop"
	  instance = ibm_is_instance.testacc_instance.id
	}
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.ISZoneName)
}

func testAccCheckIBMISInstanceActionCheckStatusConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	data "ibm_is_images" "im_images" {
  
	}
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
    	image   = data.ibm_is_images.im_images.images.4.id
    	profile = "bx2d-16x64"
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
	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.ISZoneName)
}
