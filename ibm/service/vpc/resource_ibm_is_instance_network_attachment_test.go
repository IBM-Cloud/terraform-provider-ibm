// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsInstanceNetworkAttachmentBasic(t *testing.T) {
	var conf vpcv1.InstanceNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-vsi-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tvni-subnet-%d", acctest.RandIntRange(10, 100))
	naname := fmt.Sprintf("tvni-na-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentConfigBasic(vpcname, subnetname, sshname, publicKey, vniname, name, naname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "resource_type", "instance_network_attachment"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "instance"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "network_attachment"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.#"),
				),
			},
		},
	})
}

func TestAccIBMIsInstanceNetworkAttachmentAllArgs(t *testing.T) {
	var conf vpcv1.InstanceNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-vsi-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tvni-subnet-%d", acctest.RandIntRange(10, 100))
	naname := fmt.Sprintf("tvni-na-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentConfig(vpcname, subnetname, sshname, publicKey, vniname, name, naname),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "resource_type", "instance_network_attachment"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "instance"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "network_attachment"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.allow_ip_spoofing", "true"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.enable_infrastructure_nat"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.name"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.name", vniname+"-inline"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.0.address", "10.240.64.12"),
				),
			},
		},
	})
}
func TestAccIBMIsInstanceNetworkAttachmentAllArgsUpdate(t *testing.T) {
	var conf vpcv1.InstanceNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-vsi-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	vninameupdated := fmt.Sprintf("tf-vni-upd-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tvni-subnet-%d", acctest.RandIntRange(10, 100))
	naname := fmt.Sprintf("tvni-na-%d", acctest.RandIntRange(10, 100))
	nanameupdated := fmt.Sprintf("tvni-na-upd-%d", acctest.RandIntRange(10, 100))
	protocolStateFilteringMode := "auto"
	protocolStateFilteringUpdated := "enabled"
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentUpdateConfig(vpcname, subnetname, sshname, publicKey, vniname, name, naname, protocolStateFilteringMode),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "resource_type", "instance_network_attachment"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "instance"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "network_attachment"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "name"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "name", naname),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.allow_ip_spoofing", "true"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.enable_infrastructure_nat"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.name"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.name", vniname+"-inline"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.0.address"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.protocol_state_filtering_mode", protocolStateFilteringMode),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentUpdateConfig(vpcname, subnetname, sshname, publicKey, vninameupdated, name, nanameupdated, protocolStateFilteringUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "resource_type", "instance_network_attachment"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "instance"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "network_attachment"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "name"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "name", nanameupdated),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.allow_ip_spoofing", "true"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.enable_infrastructure_nat"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.name"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.name", vninameupdated+"-inline"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.primary_ip.0.address"),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.protocol_state_filtering_mode", protocolStateFilteringUpdated),
				),
			},
		},
	})
}

func testAccCheckIBMIsInstanceNetworkAttachmentConfigBasic(vpcname, subnetname, sshname, publicKey, vniname, name, naname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	}
  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       			= "%s"
		public_key 			= "%s"
	}
	resource "ibm_is_virtual_network_interface" "testacc_vni" {
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = true
		allow_ip_spoofing = %t
	}
	resource "ibm_is_virtual_network_interface" "testacc_vni2" {
		name = "%s2"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = true
		allow_ip_spoofing = %t
	}
	resource "ibm_is_instance" "testacc_vsi" {
		profile 			= "%s"
		name 				= "%s"
		image 				= "%s"
		zone 				= "%s"
		keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
		primary_network_attachment {
		         name = "vni-2"
		         virtual_network_interface {
					id = ibm_is_virtual_network_interface.testacc_vni.id
				 }
		}
		vpc 				= ibm_is_vpc.testacc_vpc.id
	}

	resource "ibm_is_instance_network_attachment" "is_instance_network_attachment" {
		instance = ibm_is_instance.testacc_vsi.id
		name = "%s"
		virtual_network_interface {
			id = ibm_is_virtual_network_interface.testacc_vni2.id
		}
	}
	`, vpcname, subnetname, acc.ISZoneName2, sshname, publicKey, vniname, true, vniname, true, acc.InstanceProfileName, name, acc.IsImage, acc.ISZoneName2, naname)
}

func testAccCheckIBMIsInstanceNetworkAttachmentConfig(vpcname, subnetname, sshname, publicKey, vniname, name, naname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	}
  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       			= "%s"
		public_key 			= "%s"
	}
	resource "ibm_is_virtual_network_interface" "testacc_vni" {
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = true
		allow_ip_spoofing = %t
	}

	resource "ibm_is_instance" "testacc_vsi" {
		profile 			= "%s"
		name 				= "%s"
		image 				= "%s"
		zone 				= "%s"
		keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
		primary_network_attachment {
		         name = "vni-2"
		         virtual_network_interface {
					id = ibm_is_virtual_network_interface.testacc_vni.id
				 }
		}
		vpc 				= ibm_is_vpc.testacc_vpc.id
	}

	resource "ibm_is_instance_network_attachment" "is_instance_network_attachment" {
		instance = ibm_is_instance.testacc_vsi.id
		name = "%s"
		virtual_network_interface {
			name = "%s-inline"
			allow_ip_spoofing = %t
			enable_infrastructure_nat = true
			primary_ip {
				auto_delete = true
				address = "10.240.64.12"
			}
			subnet = ibm_is_subnet.testacc_subnet.id
		}
	}
	`, vpcname, subnetname, acc.ISZoneName2, sshname, publicKey, vniname, true, acc.InstanceProfileName, name, acc.IsImage, acc.ISZoneName2, naname, vniname, true)
}
func testAccCheckIBMIsInstanceNetworkAttachmentUpdateConfig(vpcname, subnetname, sshname, publicKey, vniname, name, naname, psftMode string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            			= "%s"
		vpc             			= ibm_is_vpc.testacc_vpc.id
		zone            			= "%s"
		total_ipv4_address_count 	= 16
	}
  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       			= "%s"
		public_key 			= "%s"
	}
	resource "ibm_is_virtual_network_interface" "testacc_vni" {
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = true
		allow_ip_spoofing = %t
	}

	resource "ibm_is_instance" "testacc_vsi" {
		profile 			= "%s"
		name 				= "%s"
		image 				= "%s"
		zone 				= "%s"
		keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
		primary_network_attachment {
			name = "vni-2"
			virtual_network_interface {
				id = ibm_is_virtual_network_interface.testacc_vni.id
			}
		}
		vpc 				= ibm_is_vpc.testacc_vpc.id
	}

	resource "ibm_is_instance_network_attachment" "is_instance_network_attachment" {
		instance = ibm_is_instance.testacc_vsi.id
		name = "%s"
		virtual_network_interface {
			auto_delete = true
			name 		= "%s-inline"
			allow_ip_spoofing = %t
			enable_infrastructure_nat = true
			primary_ip {
				auto_delete = true
				address = cidrhost(ibm_is_subnet.testacc_subnet.ipv4_cidr_block, 11)
			}
			protocol_state_filtering_mode = "%s"
			subnet = ibm_is_subnet.testacc_subnet.id
		}
	}
	`, vpcname, subnetname, acc.ISZoneName2, sshname, publicKey, vniname, true, acc.InstanceProfileName, name, acc.IsImage, acc.ISZoneName2, naname, vniname, true, psftMode)
}

func testAccCheckIBMIsInstanceNetworkAttachmentExists(n string, obj vpcv1.InstanceNetworkAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getInstanceNetworkAttachmentOptions := &vpcv1.GetInstanceNetworkAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceNetworkAttachmentOptions.SetInstanceID(parts[0])
		getInstanceNetworkAttachmentOptions.SetID(parts[1])

		instanceByNetworkAttachment, _, err := vpcClient.GetInstanceNetworkAttachment(getInstanceNetworkAttachmentOptions)
		if err != nil {
			return err
		}

		obj = *instanceByNetworkAttachment
		return nil
	}
}

func testAccCheckIBMIsInstanceNetworkAttachmentDestroy(s *terraform.State) error {
	vpcClient, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_network_attachment" {
			continue
		}

		getInstanceNetworkAttachmentOptions := &vpcv1.GetInstanceNetworkAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getInstanceNetworkAttachmentOptions.SetInstanceID(parts[0])
		getInstanceNetworkAttachmentOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetInstanceNetworkAttachment(getInstanceNetworkAttachmentOptions)

		if err == nil {
			return fmt.Errorf("InstanceByNetworkAttachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for InstanceByNetworkAttachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

// TestAccIBMIsInstanceNetworkAttachmentVniResourceGroupChange tests attachment with external VNI's resource group change
func TestAccIBMIsInstanceNetworkAttachmentVniResourceGroupChange(t *testing.T) {
	var conf vpcv1.InstanceNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-vsi-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tvni-subnet-%d", acctest.RandIntRange(10, 100))
	naname := fmt.Sprintf("tvni-na-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	rg1 := acc.IsResourceGroupID
	rg2 := acc.IsResourceGroupIDUpdate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			// Initial setup with resource group 1
			{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentVniResourceGroup(vpcname, subnetname, sshname, publicKey, vniname, name, naname, rg1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "name", naname),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni2", "resource_group", rg1),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
				),
			},
			// Change to resource group 2, should trigger force-new for VNI and recreate attachment
			{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentVniResourceGroup(vpcname, subnetname, sshname, publicKey, vniname, name, naname, rg2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "name", naname),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni2", "resource_group", rg2),
					resource.TestCheckResourceAttrSet("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.#"),
				),
			},
		},
	})
}

// TestAccIBMIsInstanceNetworkAttachmentInlineVniResourceGroupChange tests attachment with inline VNI resource group change
func TestAccIBMIsInstanceNetworkAttachmentInlineVniResourceGroupChange(t *testing.T) {
	var conf vpcv1.InstanceNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-vsi-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tvni-subnet-%d", acctest.RandIntRange(10, 100))
	naname := fmt.Sprintf("tvni-na-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	rg1 := acc.IsResourceGroupID
	rg2 := acc.IsResourceGroupIDUpdate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsInstanceNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			// Initial setup with resource group 1
			{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentInlineVniResourceGroup(vpcname, subnetname, sshname, publicKey, vniname, name, naname, rg1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "name", naname),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.resource_group", rg1),
				),
			},
			// Change to resource group 2, should trigger force-new
			{
				Config: testAccCheckIBMIsInstanceNetworkAttachmentInlineVniResourceGroup(vpcname, subnetname, sshname, publicKey, vniname, name, naname, rg2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsInstanceNetworkAttachmentExists("ibm_is_instance_network_attachment.is_instance_network_attachment", conf),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "name", naname),
					resource.TestCheckResourceAttr("ibm_is_instance_network_attachment.is_instance_network_attachment", "virtual_network_interface.0.resource_group", rg2),
				),
			},
		},
	})
}

// Config generators

func testAccCheckIBMIsInstanceNetworkAttachmentVniResourceGroup(vpcname, subnetname, sshname, publicKey, vniname, name, naname, resourceGroup string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
  
resource "ibm_is_subnet" "testacc_subnet" {
	name            			= "%s"
	vpc             			= ibm_is_vpc.testacc_vpc.id
	zone            			= "%s"
	total_ipv4_address_count 	= 16
}
  
resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       			= "%s"
	public_key 			= "%s"
}

resource "ibm_is_virtual_network_interface" "testacc_vni" {
	name = "%s"
	subnet = ibm_is_subnet.testacc_subnet.id
	auto_delete = false
}

resource "ibm_is_virtual_network_interface" "testacc_vni2" {
	name = "%s2"
	subnet = ibm_is_subnet.testacc_subnet.id
	resource_group = "%s"
	auto_delete = false
}

resource "ibm_is_instance" "testacc_vsi" {
	profile 			= "%s"
	name 				= "%s"
	image 				= "%s"
	zone 				= "%s"
	keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
	primary_network_attachment {
		 name = "vni-primary"
		 virtual_network_interface {
			id = ibm_is_virtual_network_interface.testacc_vni.id
		 }
	}
	vpc 				= ibm_is_vpc.testacc_vpc.id
}

resource "ibm_is_instance_network_attachment" "is_instance_network_attachment" {
	instance = ibm_is_instance.testacc_vsi.id
	name = "%s"
	virtual_network_interface {
		id = ibm_is_virtual_network_interface.testacc_vni2.id
	}
}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, vniname, vniname, resourceGroup, acc.InstanceProfileName, name, acc.IsImage, acc.ISZoneName, naname)
}

func testAccCheckIBMIsInstanceNetworkAttachmentInlineVniResourceGroup(vpcname, subnetname, sshname, publicKey, vniname, name, naname, resourceGroup string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}
  
resource "ibm_is_subnet" "testacc_subnet" {
	name            			= "%s"
	vpc             			= ibm_is_vpc.testacc_vpc.id
	zone            			= "%s"
	total_ipv4_address_count 	= 16
}
  
resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       			= "%s"
	public_key 			= "%s"
}

resource "ibm_is_virtual_network_interface" "testacc_vni" {
	name = "%s"
	subnet = ibm_is_subnet.testacc_subnet.id
	enable_infrastructure_nat = true
	auto_delete = false
}

resource "ibm_is_instance" "testacc_vsi" {
	profile 			= "%s"
	name 				= "%s"
	image 				= "%s"
	zone 				= "%s"
	keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
	primary_network_attachment {
		 name = "vni-primary"
		 virtual_network_interface {
			id = ibm_is_virtual_network_interface.testacc_vni.id
		 }
	}
	vpc 				= ibm_is_vpc.testacc_vpc.id
}

resource "ibm_is_instance_network_attachment" "is_instance_network_attachment" {
	instance = ibm_is_instance.testacc_vsi.id
	name = "%s"
	virtual_network_interface {
		subnet = ibm_is_subnet.testacc_subnet.id
		resource_group = "%s"
		name = "inline-vni-na"
		auto_delete = true
	}
}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, vniname, acc.InstanceProfileName, name, acc.IsImage, acc.ISZoneName, naname, resourceGroup)
}
