// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

const NOT_FOUND = "SoftLayer_Exception_Network_LBaaS_ObjectNotFound"

func TestAccIBMIPSec_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIPSecConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "datacenter", acc.IpsecDatacenter),
				),
			},
			{
				Config: testAccCheckIBMIPSECConfig_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "datacenter", acc.IpsecDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_one.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_two.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "remote_subnet_id", acc.Customersubnetid),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "remote_subnet_id", acc.Customersubnetid)),
			},
			{
				Config: testAccCheckIBMIPSECConfig_updatesubnet(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "datacenter", acc.IpsecDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_one.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_two.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "Customer_Peer_IP", acc.Customerpeerip),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "remote_subnet.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "remote_subnet.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidEncryptionProtocol(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIPSECConfig_InvalidEncryptionProtocol,
				ExpectError: regexp.MustCompile("auth protocol can be DES or 3DES or AES128 or AES192 or AES256"),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidDiffieHellmanGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIPSECConfig_InvalidDiffieHellmanGroup,
				ExpectError: regexp.MustCompile("auth protocol can be MD5 or SHA1 or SHA256"),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidAuthProtocol(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIPSECConfig_InvalidAuthProtocol,
				ExpectError: regexp.MustCompile("auth protocol can be MD5 or SHA1 or SHA256"),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidKeyLife(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMIPSECConfig_InvalidKeyLife,
				ExpectError: regexp.MustCompile("keylife value can be between 120 and 172800"),
			},
		},
	})
}

func testAccCheckIBMIPSecDestroy(s *terraform.State) error {
	sess := acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_ipsec_vpn" {
			continue
		}
		id, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := services.GetNetworkTunnelModuleContextService(sess).
			Id(id).
			GetObject()

		if err == nil {
			return fmt.Errorf("ipsec vpn (%s) to be destroyed still exists", rs.Primary.ID)
		} else if apiErr, ok := err.(sl.Error); ok && apiErr.Exception != NOT_FOUND {
			return fmt.Errorf("[ERROR] Error waiting for IPSec VPN (%s) to be destroyed: %s", rs.Primary.ID, err)
		}

	}

	return nil
}

func testAccCheckIBMIPSecConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "ibm_ipsec_vpn" "ipsec" {
  datacenter        = "%s"
  
}
`, acc.IpsecDatacenter)
}

func testAccCheckIBMIPSECConfig_update(name string) string {
	return fmt.Sprintf(`
		resource "ibm_ipsec_vpn" "ipsec" {
			datacenter = "tok02"
			Customer_Peer_IP = %s
			phase_one = [{Encryption="3DES",Diffie-Hellman-Group=2,Keylife=131}]
			phase_two = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=2,Keylife=133}]
			remote_subnet_id = %s
			}
			
`, acc.Customerpeerip, acc.Customersubnetid)
}

func testAccCheckIBMIPSECConfig_updatesubnet(name string) string {
	return fmt.Sprintf(`
		resource "ibm_ipsec_vpn" "ipsec" {
			datacenter = %s
			Customer_Peer_IP = %s
			phase_one = [{Encryption="3DES",Diffie-Hellman-Group=2,Keylife=131}]
			phase_two = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=2,Keylife=133}]
			remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
			}
			
`, acc.IpsecDatacenter, acc.Customerpeerip)
}

const testAccCheckIBMIPSECConfig_InvalidEncryptionProtocol = `resource "ibm_ipsec_vpn" "ipsec" {
	datacenter = "tok02"
	Customer_Peer_IP = "192.168.32.2"
	phase_one = [{Authentication="SHA1",Encryption="5DES",Diffie-Hellman-Group=2,Keylife=131}]
	phase_two = [{Authentication="SHA1",Encryption="5DES",Diffie-Hellman-Group=2,Keylife=133}]
	remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
	}`

const testAccCheckIBMIPSECConfig_InvalidDiffieHellmanGroup = `resource "ibm_ipsec_vpn" "ipsec" {
	datacenter = "tok02"
	Customer_Peer_IP = "192.168.32.2"
	phase_one = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=131}]
	phase_two = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=133}]
	remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
	}`

const testAccCheckIBMIPSECConfig_InvalidAuthProtocol = `resource "ibm_ipsec_vpn" "ipsec" {
	datacenter = "tok02"
	Customer_Peer_IP = "192.168.32.2"
	phase_one = [{Authentication="SABC",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=131}]
	phase_two = [{Authentication="SABC",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=133}]
	remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
	}`

const testAccCheckIBMIPSECConfig_InvalidKeyLife = `resource "ibm_ipsec_vpn" "ipsec" {
	datacenter = "tok02"
	Customer_Peer_IP = "192.168.32.2"
	phase_one = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=12,Keylife=2}]
	phase_two = [{Authentication="SHA1",Encryption="3DES",Diffie-Hellman-Group=2,Keylife=2}]
	remote_subnet = [{Remote_ip_adress = "10.0.0.0",Remote_IP_CIDR = 22}]
	}`
