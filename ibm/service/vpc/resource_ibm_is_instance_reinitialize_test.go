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

func TestAccIBMISInstanceReinitialize_basic(t *testing.T) {
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
				Config: testAccCheckIBMISInstanceReinitializeByImageConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "image", acc.IsImage),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "id"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceReinitializeByImageUpdateConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_reinitialize.test_reinit", "instance_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_reinitialize.test_reinit", "status"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance.testacc_instance_data", "image", acc.IsImage2),
				),
			},
		},
	})
}

// Test configuration for reinitialization by image
func testAccCheckIBMISInstanceReinitializeByImageConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	

	
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}
	
	
	`, sshname, publicKey, name, acc.IsImage, acc.ISZoneName)
}
func testAccCheckIBMISInstanceReinitializeByImageUpdateConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
		profile = "bx2d-2x8"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		lifecycle {
			ignore_changes = [image]
		}
	}
	
	resource "ibm_is_instance_reinitialize" "test_reinit" {
		instance_id = ibm_is_instance.testacc_instance.id
		image       = "%s"
		keys        = [ibm_is_ssh_key.testacc_sshkey.id]
		user_data   = "#!/bin/bash\necho reinit-by-image > /tmp/reinit.log"
	}
	
	data "ibm_is_instance" "testacc_instance_data" {
		depends_on  = [ibm_is_instance_reinitialize.test_reinit]
		name    	= ibm_is_instance.testacc_instance.name
	}

	`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.ISZoneName, acc.IsImage2)
}
