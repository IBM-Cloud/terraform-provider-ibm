// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsInstanceNetworkAttachmentDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsInstanceNetworkAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, vniname, userData1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "instance"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "network_attachment"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.id"),
				),
			},
		},
	})
}

func testAccCheckIBMIsInstanceNetworkAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, name, vniname, userData1 string) string {
	return testAccCheckIBMISInstanceVniConfig(vpcname, subnetname, sshname, publicKey, name, vniname, userData1) + fmt.Sprintf(`
		data "ibm_is_instance_network_attachment" "is_instance_network_attachment" {
			instance 			= ibm_is_instance.testacc_instance.id
			network_attachment 	= ibm_is_instance.testacc_instance.primary_network_attachment.0.id
		}
	`)
}
