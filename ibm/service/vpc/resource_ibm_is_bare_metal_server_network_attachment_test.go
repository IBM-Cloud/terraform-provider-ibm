// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsBareMetalServerNetworkAttachmentPci(t *testing.T) {
	var conf vpcv1.BareMetalServerNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBareMetalServerNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigPci(vpcname, subnetname, sshname, publicKey, vniname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBareMetalServerNetworkAttachmentExists("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", conf),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "allowed_vlans.#"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "interface_type"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "allowed_vlans.#", "3"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "resource_type", "bare_metal_server_network_attachment"),
				),
			},
		},
	})
}
func TestAccIBMIsBareMetalServerNetworkAttachmentVlan(t *testing.T) {
	var conf vpcv1.BareMetalServerNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBareMetalServerNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlan(vpcname, subnetname, sshname, publicKey, vniname, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBareMetalServerNetworkAttachmentExists("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", conf),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "vlan"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "interface_type"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "vlan", "100"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "resource_type", "bare_metal_server_network_attachment"),
				),
			},
		},
	})
}

func TestAccIBMIsBareMetalServerNetworkAttachmentVlanPSFM(t *testing.T) {
	var conf vpcv1.BareMetalServerNetworkAttachment
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	psfm1 := "auto"
	psfm2 := "disabled"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsBareMetalServerNetworkAttachmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlanPSFM(vpcname, subnetname, sshname, publicKey, vniname, name, psfm1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBareMetalServerNetworkAttachmentExists("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", conf),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "vlan"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "interface_type"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "vlan", "100"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "resource_type", "bare_metal_server_network_attachment"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlanPSFM(vpcname, subnetname, sshname, publicKey, vniname, name, psfm2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsBareMetalServerNetworkAttachmentExists("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", conf),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "vlan"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "interface_type"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "port_speed"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "vlan", "100"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "type", "secondary"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment", "resource_type", "bare_metal_server_network_attachment"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlan(vpcname, subnetname, sshname, publicKey, vniname, name string) string {
	return testAccCheckIBMISBareMetalServerVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name) + fmt.Sprintf(`
	resource "ibm_is_virtual_network_interface" "testacc_vni1"{
		name 						= "test-vni-na"
		subnet 						= "0717-a68ebc16-d63c-44e7-aa37-4b3791415b1d"
		enable_infrastructure_nat 	= true
		allow_ip_spoofing 			= true
	}	
	resource "ibm_is_bare_metal_server_network_attachment" "is_bare_metal_server_network_attachment" {
			bare_metal_server 		= "0717-0579cb7c-5ca7-4cb2-9792-7dcf06de25ca"
			vlan 					= 100
			virtual_network_interface { 
				id = ibm_is_virtual_network_interface.testacc_vni1.id
			}
	}
	`)
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigVlanPSFM(vpcname, subnetname, sshname, publicKey, vniname, name, psfm string) string {
	return testAccCheckIBMISBareMetalServerVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name) + fmt.Sprintf(`
	
	resource "ibm_is_bare_metal_server_network_attachment" "is_bare_metal_server_network_attachment" {
			bare_metal_server 		= ibm_is_bare_metal_server.testacc_bms.id
			vlan 					= 100
			virtual_network_interface { 
				name 						= "%s"
				subnet 						= "0717-a68ebc16-d63c-44e7-aa37-4b3791415b1d"
				enable_infrastructure_nat 	= true
				allow_ip_spoofing 			= true
				protocol_state_filtering_mode = "%s"
			}
	}
	`, vniname, psfm)
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentConfigPci(vpcname, subnetname, sshname, publicKey, vniname, name string) string {
	return testAccCheckIBMISBareMetalServerVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name) + fmt.Sprintf(`
	resource "ibm_is_virtual_network_interface" "testacc_vni1"{
		name = "test-vni-na"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = true
		allow_ip_spoofing = true
	}	
	resource "ibm_is_bare_metal_server_network_attachment" "is_bare_metal_server_network_attachment" {
			bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
			allowed_vlans = [300, 302, 303]
			virtual_network_interface { 
				id = ibm_is_virtual_network_interface.testacc_vni1.id
			}
	}
	`)
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentExists(n string, obj vpcv1.BareMetalServerNetworkAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getBareMetalServerNetworkAttachmentOptions := &vpcv1.GetBareMetalServerNetworkAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBareMetalServerNetworkAttachmentOptions.SetBareMetalServerID(parts[0])
		getBareMetalServerNetworkAttachmentOptions.SetID(parts[1])

		bareMetalServerNetworkAttachmentIntf, _, err := vpcClient.GetBareMetalServerNetworkAttachment(getBareMetalServerNetworkAttachmentOptions)
		if err != nil {
			return err
		}

		bareMetalServerNetworkAttachment := bareMetalServerNetworkAttachmentIntf.(*vpcv1.BareMetalServerNetworkAttachment)
		obj = *bareMetalServerNetworkAttachment
		return nil
	}
}

func testAccCheckIBMIsBareMetalServerNetworkAttachmentDestroy(s *terraform.State) error {
	vpcClient, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_bare_metal_server_network_attachment" {
			continue
		}

		getBareMetalServerNetworkAttachmentOptions := &vpcv1.GetBareMetalServerNetworkAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBareMetalServerNetworkAttachmentOptions.SetBareMetalServerID(parts[0])
		getBareMetalServerNetworkAttachmentOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := vpcClient.GetBareMetalServerNetworkAttachment(getBareMetalServerNetworkAttachmentOptions)

		if err == nil {
			return fmt.Errorf("is_bare_metal_server_network_attachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for is_bare_metal_server_network_attachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
