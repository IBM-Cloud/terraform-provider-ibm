// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsInstanceNetworkAttachmentsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR`)
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, vniname, userData1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "instance"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.virtual_network_interface.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachments.is_instance_network_attachments", "network_attachments.0.virtual_network_interface.0.id"),
				),
			},
		},
	})
}

func testAccCheckIBMIsInstanceNetworkAttachmentsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, vniname, userData1 string) string {
	return testAccCheckIBMISInstanceVniConfig(vpcname, subnetname, sshname, publicKey, name, vniname, userData1) + fmt.Sprintf(`
		data "ibm_is_instance_network_attachments" "is_instance_network_attachments" {
			instance = ibm_is_instance.testacc_instance.id
		}
	`)
}
