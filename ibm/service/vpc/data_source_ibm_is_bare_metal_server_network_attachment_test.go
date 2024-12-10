// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsBareMetalServerNetworkAttachmentDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBareMetalServerNetworkAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, vniname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "bare_metal_server"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "network_attachment"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "bare_metal_server_network_attachment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "interface_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, vniname, name string) string {
	return testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlan(vpcname, subnetname, sshname, publicKey, vniname, name) + fmt.Sprintf(`
		data "ibm_is_bare_metal_server_network_attachment" "is_bare_metal_server_network_attachment" {
			bare_metal_server = ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment.bare_metal_server
			network_attachment = ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment.network_attachment
		}
	`)
}
