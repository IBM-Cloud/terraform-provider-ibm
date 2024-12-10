// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIsVPNServerDataSourceBasic(t *testing.T) {
	if acc.ISCertificateCrn == "" {
		fmt.Println("[ERROR] Set the environment variable IS_CERTIFICATE_CRN for testing ibm_is_vpn_server resource")
	}

	if acc.ISClientCaCrn == "" {
		fmt.Println("[ERROR] Set the environment variable IS_CLIENT_CA_CRN for testing ibm_is_vpn_server resource")
	}
	isCertificateCrn := acc.ISCertificateCrn
	isClientCaCrn := acc.ISClientCaCrn
	clientIPPool := "10.5.0.0/21"
	clientIdleTimeout := fmt.Sprintf("%d", acctest.RandIntRange(0, 28800))
	enableSplitTunneling := "true"
	nameVpc := fmt.Sprintf("test-vpc-tf-%d", acctest.RandIntRange(10, 100))
	nameSubnet1 := fmt.Sprintf("test-subnet1-tf-%d", acctest.RandIntRange(10, 100))
	vpnServerName := fmt.Sprintf("tfname%d", acctest.RandIntRange(10, 100))
	port := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocol := "udp"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVPNServerDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "certificate.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "client_authentication.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "client_auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "client_auto_delete_timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "client_dns_server_ips.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "client_idle_timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "client_ip_pool"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "enable_split_tunneling"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "hostname"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "port"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "private_ips.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "protocol"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "security_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "subnets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "vpc.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "vpc.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "vpc.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server.is_vpn_server", "vpc.name"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNServerDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn string) string {
	return testAccCheckIBMIsVPNServerConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn) + fmt.Sprintf(`
		data "ibm_is_vpn_server" "is_vpn_server" {
			identifier = ibm_is_vpn_server.is_vpn_server.id
		}
	`)
}
