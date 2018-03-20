package ibm

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMIPSec_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform-%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIPSecConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "datacenter", ipsecDatacenter),
				),
			},
			{
				Config: testAccCheckIBMIPSECConfig_update(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "datacenter", ipsecDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_one.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_two.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "remote_subnet_id", customersubnetid),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "remote_subnet_id", customersubnetid)),
			},
			{
				Config: testAccCheckIBMIPSECConfig_updatesubnet(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "datacenter", ipsecDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_one.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "phase_two.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ipsec.ipsec", "Customer_Peer_IP", customerpeerip),
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIPSECConfig_InvalidEncryptionProtocol,
				ExpectError: regexp.MustCompile("auth protocol can be DES or 3DES or AES128 or AES192 or AES256"),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidDiffieHellmanGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIPSECConfig_InvalidDiffieHellmanGroup,
				ExpectError: regexp.MustCompile("auth protocol can be MD5 or SHA1 or SHA256"),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidAuthProtocol(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIPSECConfig_InvalidAuthProtocol,
				ExpectError: regexp.MustCompile("auth protocol can be MD5 or SHA1 or SHA256"),
			},
		},
	})
}

func TestAccIBMIPSec_InvalidKeyLife(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMLbaasDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMIPSECConfig_InvalidKeyLife,
				ExpectError: regexp.MustCompile("keylife value can be between 120 and 172800"),
			},
		},
	})
}

func testAccCheckIBMIPSecDestroy(s *terraform.State) error {
	sess := testAccProvider.Meta().(ClientSession).SoftLayerSession()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_ipsec_vpn" {
			continue
		}
		id, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err = services.GetNetworkTunnelModuleContextService(sess).
			Id(id).
			GetObject()

		if err == nil {
			return fmt.Errorf("ipsec vpn (%s) to be destroyed still exists", rs.Primary.ID)
		} else if apiErr, ok := err.(sl.Error); ok && apiErr.Exception != NOT_FOUND {
			return fmt.Errorf("Error waiting for IPSec VPN (%s) to be destroyed: %s", rs.Primary.ID, err)
		}

	}

	return nil
}

func testAccCheckIBMIPSecConfig_basic(name string) string {
	return fmt.Sprintf(`
resource "ibm_ipsec_vpn" "ipsec" {
  datacenter        = "%s"
  
}
`, ipsecDatacenter)
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
			
`, customerpeerip, customersubnetid)
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
			
`, ipsecDatacenter, customerpeerip)
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
