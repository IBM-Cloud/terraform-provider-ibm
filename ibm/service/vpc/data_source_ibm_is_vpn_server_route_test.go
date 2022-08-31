// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVPNServerRouteDataSourceBasic(t *testing.T) {
	if acc.ISCertificateCrn == "" {
		fmt.Println("[ERROR] Set the environment variable IS_CERTIFICATE_CRN for testing ibm_is_vpn_server resource")
	}

	if acc.ISClientCaCrn == "" {
		fmt.Println("[ERROR] Set the environment variable IS_CLIENT_CA_CRN for testing ibm_is_vpn_server resource")
	}
	isCertificateCrn := acc.ISCertificateCrn
	isClientCaCrn := acc.ISClientCaCrn
	nameVpc := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	nameSubnet1 := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vpnServerName := fmt.Sprintf("tf-vpnserver-%d", acctest.RandIntRange(10, 100))
	clientIPPool := "10.5.0.0/21"
	clientIdleTimeout := fmt.Sprintf("%d", acctest.RandIntRange(0, 28800))
	enableSplitTunneling := "true"
	destination := "172.16.0.0/16"
	name := fmt.Sprintf("tfname%d", acctest.RandIntRange(10, 100))
	action := "translate"
	port := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocol := "udp"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVPNServerRouteDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, destination, action, name, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "vpn_server"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "destination"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_route.is_vpn_server_route", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNServerRouteDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, destination, action, name, isCertificateCrn, isClientCaCrn string) string {
	return testAccCheckIBMIsVPNServerRouteConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, destination, action, name, isCertificateCrn, isClientCaCrn) + fmt.Sprintf(`
		data "ibm_is_vpn_server_route" "is_vpn_server_route" {
			vpn_server = ibm_is_vpn_server_route.is_vpn_server_route.vpn_server
			identifier = ibm_is_vpn_server_route.is_vpn_server_route.vpn_route
		}
	`)
}
