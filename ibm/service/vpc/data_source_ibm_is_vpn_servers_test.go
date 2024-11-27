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

func TestAccIBMIsVPNServersDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsVPNServersDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.client_auto_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.client_auto_delete_timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.hostname"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.vpc.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.vpc.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.vpc.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_servers.is_vpn_servers", "vpn_servers.0.vpc.name"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNServersDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn string) string {
	return testAccCheckIBMIsVPNServerConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn) + fmt.Sprintf(`
		data "ibm_is_vpn_servers" "is_vpn_servers" {
		}
	`)
}
