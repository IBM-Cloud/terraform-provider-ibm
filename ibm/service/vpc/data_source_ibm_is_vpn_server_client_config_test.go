// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVPNServerClientConfigDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIBMIsVPNServerClientConfigDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_server_client_configuration.is_vpn_server_client_configuration", "vpn_server"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNServerClientConfigDataSourceConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn string) string {
	return testAccCheckIBMIsVPNServerConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn) + fmt.Sprintf(`
		data "ibm_is_vpn_server_client_configuration" "is_vpn_server_client_configuration" {
			vpn_server = ibm_is_vpn_server.is_vpn_server.id
			file_path = "vpnServerClinetConfigFile.txt"
		}
	`)
}
