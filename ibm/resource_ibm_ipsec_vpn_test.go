package ibm

import (
	"fmt"
	"regexp"
	"strconv"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
	"testing"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
						"ibm_lbaas.lbaas", "phase_two.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMLbaasConfig_updateHTTPS(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ipsec_vpn.lbaas", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "description", "updated desc-used for terraform uat"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "datacenter", lbaasDatacenter),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "subnets.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "protocols.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_lbaas.lbaas", "server_instances.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMIPSecDestroy(s *terraform.State) error {
	sess := testAccProvider.Meta().(ClientSession).SoftLayerSession()
	service := services.GetNetworkTunnelModuleContextService(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_ipsec_vpn" {
			continue
		}
		id,_ := strconv.Atoi(rs.Primary.ID)

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
			
`, customerpeerip,customersubnetid)
}
