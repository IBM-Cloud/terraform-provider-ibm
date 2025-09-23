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

func TestAccIBMIsBareMetalServerNetworkAttachmentsDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsBareMetalServerNetworkAttachmentsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, vniname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "bare_metal_server"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.interface_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.port_speed"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.subnet.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_bare_metal_server_network_attachments.is_bare_metal_server_network_attachments", "network_attachments.0.virtual_network_interface.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentsDataSourceConfigBasic(vpcname, subnetname, sshname, publicKey, vniname, name string) string {
	return testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlan(vpcname, subnetname, sshname, publicKey, vniname, name) + fmt.Sprintf(`
		data "ibm_is_bare_metal_server_network_attachments" "is_bare_metal_server_network_attachments" {
			bare_metal_server = "0717-1193c3f7-b23c-4e35-9e65-01f3a8741085"
		}
	`)
}
