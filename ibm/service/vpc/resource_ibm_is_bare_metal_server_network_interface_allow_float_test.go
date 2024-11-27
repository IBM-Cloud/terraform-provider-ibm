// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISBareMetalServerNetworkInterfaceAllowFloat_rip_basic(t *testing.T) {
	var serverNic string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	subnetreservedipname := fmt.Sprintf("tfip-subnet-rip-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatRipConfig(vpcname, subnetname, subnetreservedipname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatExists("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", serverNic),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet_reserved_ip.testacc_rip", "name", subnetreservedipname),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "interface_type", "vlan"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_ip_spoofing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "name", "eth21"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "primary_ip.0.address", func(v string) error {
						if v == "0.0.0.0" {
							return fmt.Errorf("Attribute 'address' %s is not updated", v)
						}
						return nil
					}),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server", func(v string) error {
						if v == "" {
							return fmt.Errorf("Attribute 'floating_bare_metal_server' %s is not populated", v)
						}
						return nil
					}),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServerNetworkInterfaceAllowFloat_basic(t *testing.T) {
	var serverNic string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatExists("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", serverNic),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_ip_spoofing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_interface_to_float", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "enable_infrastructure_nat", "false"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "primary_ip.0.address", func(v string) error {
						if v == "0.0.0.0" {
							return fmt.Errorf("Attribute 'address' %s is not updated", v)
						}
						return nil
					}),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server", func(v string) error {
						if v == "" {
							return fmt.Errorf("Attribute 'floating_bare_metal_server' %s is not populated", v)
						}
						return nil
					}),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServerNetworkInterfaceAllowFloat_sg_update(t *testing.T) {
	var serverNic string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatExists("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", serverNic),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_ip_spoofing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "security_groups.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_interface_to_float", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "enable_infrastructure_nat", "false"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "primary_ip.0.address", func(v string) error {
						if v == "0.0.0.0" {
							return fmt.Errorf("Attribute 'address' %s is not updated", v)
						}
						return nil
					}),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server", func(v string) error {
						if v == "" {
							return fmt.Errorf("Attribute 'floating_bare_metal_server' %s is not populated", v)
						}
						return nil
					}),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatSgUpdateConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatExists("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", serverNic),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "security_groups.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_ip_spoofing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_interface_to_float", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "enable_infrastructure_nat", "false"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "primary_ip.0.address", func(v string) error {
						if v == "0.0.0.0" {
							return fmt.Errorf("Attribute 'address' %s is not updated", v)
						}
						return nil
					}),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server"),
					resource.TestCheckResourceAttrWith("ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "floating_bare_metal_server", func(v string) error {
						if v == "" {
							return fmt.Errorf("Attribute 'floating_bare_metal_server' %s is not populated", v)
						}
						return nil
					}),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatExists(n, ip string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		bmsId, nicId, err := vpc.ParseNICTerraformID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("[ERROR] Error parsing ID : %s", rs.Primary.ID)
		}
		getbmsnicoptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &bmsId,
			ID:                &nicId,
		}
		nicIntf, _, err := sess.GetBareMetalServerNetworkInterface(getbmsnicoptions)

		if err != nil {
			return err
		}
		switch reflect.TypeOf(nicIntf).String() {
		case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
			{
				nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				ip = vpc.MakeTerraformNICID(bmsId, *nic.ID)
			}
		case "*vpcv1.BareMetalServerNetworkInterfaceByVlan":
			{
				nic := nicIntf.(*vpcv1.BareMetalServerNetworkInterfaceByVlan)
				ip = vpc.MakeTerraformNICID(bmsId, *nic.ID)
			}
		}
		return nil
	}
}

func testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				allowed_vlans 	= [101, 102]
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}

		resource ibm_is_bare_metal_server_network_interface_allow_float bms_nic {
			bare_metal_server 	= ibm_is_bare_metal_server.testacc_bms.id
			
			subnet 				= ibm_is_subnet.testacc_subnet.id
			name   				= "eth21"
			vlan 				= 101
			allow_ip_spoofing 	= false
			enable_infrastructure_nat = false
			}

		
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatSgUpdateConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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

		resource "ibm_is_security_group" "testacc_sg1" {
			name = "%s-security-group1"
			vpc  = ibm_is_vpc.testacc_vpc.id
		  }
		  
		resource "ibm_is_security_group" "testacc_sg2" {
			name = "%s-security-group2"
			vpc  = ibm_is_vpc.testacc_vpc.id
		}
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       			= "%s"
			public_key 			= "%s"
		}
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				allowed_vlans 	= [101, 102]
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}

		resource ibm_is_bare_metal_server_network_interface_allow_float bms_nic {
			bare_metal_server 			= ibm_is_bare_metal_server.testacc_bms.id
			
			subnet 						= ibm_is_subnet.testacc_subnet.id
			name   						= "eth21"
			vlan 						= 101
			allow_ip_spoofing 			= false
			enable_infrastructure_nat 	= false
			security_groups				= [ibm_is_security_group.testacc_sg1.id, ibm_is_security_group.testacc_sg2.id]

		}

		
`, vpcname, subnetname, acc.ISZoneName, vpcname, vpcname, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerNetworkInterfaceAllowFloatRipConfig(vpcname, subnetname, subnetreservedipname, sshname, publicKey, name string) string {
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
		resource "ibm_is_subnet_reserved_ip" "testacc_rip" {
			subnet = ibm_is_subnet.testacc_subnet.id
			name = "%s"
		}
	  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       			= "%s"
			public_key 			= "%s"
		}
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				allowed_vlans 	= [101, 102]
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}

		resource ibm_is_bare_metal_server_network_interface_allow_float bms_nic {
			bare_metal_server 	= ibm_is_bare_metal_server.testacc_bms.id
			
			subnet 				= ibm_is_subnet.testacc_subnet.id
			name   				= "eth21"
			vlan 				= 101
			primary_ip {
				reserved_ip = ibm_is_subnet_reserved_ip.testacc_rip.reserved_ip
			}
		}

		
`, vpcname, subnetname, acc.ISZoneName, subnetreservedipname, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
