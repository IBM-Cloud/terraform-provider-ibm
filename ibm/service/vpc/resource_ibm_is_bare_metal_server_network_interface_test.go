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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISBareMetalServerNetworkInterface_basic(t *testing.T) {
	var server string
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
				Config: testAccCheckIBMISBareMetalServerNetworkInterfaceConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerNetworkInterfaceExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServerNetworkInterface_basic_rip(t *testing.T) {
	var server string
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
				Config: testAccCheckIBMISBareMetalServerNetworkInterfaceRipConfig(vpcname, subnetname, subnetreservedipname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerNetworkInterfaceExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerNetworkInterfaceDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_bare_metal_server_network_interface" {
			continue
		}
		bmsId, nicId, err := vpc.ParseNICTerraformID(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("[ERROR] Error parsing ID : %s", rs.Primary.ID)
		}
		getbmsnicoptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
			BareMetalServerID: &bmsId,
			ID:                &nicId,
		}
		_, _, err = sess.GetBareMetalServerNetworkInterface(getbmsnicoptions)
		if err == nil {
			return fmt.Errorf("Bare metal server network interafce still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISBareMetalServerNetworkInterfaceExists(n, ip string) resource.TestCheckFunc {
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

func testAccCheckIBMISBareMetalServerNetworkInterfaceConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
		resource ibm_is_bare_metal_server_network_interface bms_nic {
			bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
		  
			subnet = ibm_is_subnet.testacc_subnet.id
			name   = "eth2"
			allow_ip_spoofing = true
			allowed_vlans = [101, 102]
		  }
		
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerNetworkInterfaceRipConfig(vpcname, subnetname, subnetreservedipname, sshname, publicKey, name string) string {
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
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
		resource ibm_is_bare_metal_server_network_interface bms_nic {
			bare_metal_server = ibm_is_bare_metal_server.testacc_bms.id
		  
			subnet = ibm_is_subnet.testacc_subnet.id
			name   = "eth2"
			allow_ip_spoofing = true
			allowed_vlans = [101, 102]
			primary_ip {
				reserved_ip = ibm_is_subnet_reserved_ip.testacc_rip.reserved_ip
			}
		  }
		
`, vpcname, subnetname, acc.ISZoneName, subnetreservedipname, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsImage, acc.ISZoneName)
}
