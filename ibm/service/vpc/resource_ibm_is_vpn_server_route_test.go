// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVPNServerRouteBasic(t *testing.T) {
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
	nameUpdate := fmt.Sprintf("tfname%d", acctest.RandIntRange(10, 100))
	action := "translate"
	port := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocol := "udp"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVPNServerRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVPNServerRouteConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, destination, action, name, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpn_server_route.is_vpn_server_route", "destination", destination),
					resource.TestCheckResourceAttr("ibm_is_vpn_server_route.is_vpn_server_route", "action", action),
					resource.TestCheckResourceAttr("ibm_is_vpn_server_route.is_vpn_server_route", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server_route.is_vpn_server_route", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server_route.is_vpn_server_route", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server_route.is_vpn_server_route", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server_route.is_vpn_server_route", "health_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server_route.is_vpn_server_route", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server_route.is_vpn_server_route", "vpn_route"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVPNServerRouteConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, destination, action, nameUpdate, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpn_server_route.is_vpn_server_route", "destination", destination),
					resource.TestCheckResourceAttr("ibm_is_vpn_server_route.is_vpn_server_route", "action", action),
					resource.TestCheckResourceAttr("ibm_is_vpn_server_route.is_vpn_server_route", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_vpn_server_route.is_vpn_server_route",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVPNServerRouteConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol string, destination string, action string, name, isCertificateCrn, isClientCaCrn string) string {
	return testAccCheckIBMIsVPNServerConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol, isCertificateCrn, isClientCaCrn) + fmt.Sprintf(`
		resource "ibm_is_vpn_server_route" "is_vpn_server_route" {
			vpn_server = ibm_is_vpn_server.is_vpn_server.id
			destination = "%s"
			action = "%s"
			name = "%s"
		}
	`, destination, action, name)
}

func testAccCheckIBMIsVPNServerRouteExists(n string, obj vpcv1.VPNServerRoute) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()

		getVPNServerRouteOptions := &vpcv1.GetVPNServerRouteOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVPNServerRouteOptions.SetVPNServerID(parts[0])
		getVPNServerRouteOptions.SetID(parts[1])

		vpnServerRoute, _, err := sess.GetVPNServerRoute(getVPNServerRouteOptions)
		if err != nil {
			return err
		}

		obj = *vpnServerRoute
		return nil
	}
}

func testAccCheckIBMIsVPNServerRouteDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_server_route" {
			continue
		}

		getVPNServerRouteOptions := &vpcv1.GetVPNServerRouteOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getVPNServerRouteOptions.SetVPNServerID(parts[0])
		getVPNServerRouteOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := sess.GetVPNServerRoute(getVPNServerRouteOptions)

		if err == nil {
			return fmt.Errorf("VPNServerRoute still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VPNServerRoute (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
