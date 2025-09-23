// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIBMIsVPNServerBasic(t *testing.T) {
	var vpnserver string
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
	name := fmt.Sprintf("tf-name%d", acctest.RandIntRange(10, 100))
	port := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocol := "udp"

	clientIPPoolUpdate := "10.6.0.0/21"
	clientIdleTimeoutUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 28800))
	enableSplitTunnelingUpdate := "false"
	nameUpdate := fmt.Sprintf("tf-name%d", acctest.RandIntRange(10, 100))
	portUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocolUpdate := "tcp"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVPNServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNServerConfigBasic(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, name, port, protocol, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVPNServerExists("ibm_is_vpn_server.is_vpn_server", vpnserver),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "certificate_crn"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_authentication.0.method"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_auto_delete"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_auto_delete_timeout"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_dns_server_ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "hostname"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "health_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "vpn_server"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.address"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "security_groups.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "subnets.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "vpc.#"),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_ip_pool", clientIPPool),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_idle_timeout", clientIdleTimeout),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "enable_split_tunneling", enableSplitTunneling),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "name", name),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "port", port),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "protocol", protocol),
				),
			},
			{
				Config: testAccCheckIBMIsVPNServerConfigBasic(nameVpc, nameSubnet1, clientIPPoolUpdate, clientIdleTimeoutUpdate, enableSplitTunnelingUpdate, nameUpdate, portUpdate, protocolUpdate, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_ip_pool", clientIPPoolUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_idle_timeout", clientIdleTimeoutUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "enable_split_tunneling", enableSplitTunnelingUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "port", portUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "protocol", protocolUpdate),
				),
			},
			{
				ResourceName:      "ibm_is_vpn_server.is_vpn_server",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccIBMIsVPNServerBasicTags(t *testing.T) {
	var vpnserver string
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
	name := fmt.Sprintf("tf-name%d", acctest.RandIntRange(10, 100))
	port := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocol := "udp"

	clientIPPoolUpdate := "10.6.0.0/21"
	clientIdleTimeoutUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 28800))
	enableSplitTunnelingUpdate := "false"
	nameUpdate := fmt.Sprintf("tf-name%d", acctest.RandIntRange(10, 100))
	portUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 65535))
	protocolUpdate := "tcp"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVPNServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNServerConfigBasicTags(nameVpc, nameSubnet1, clientIPPool, clientIdleTimeout, enableSplitTunneling, name, port, protocol, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVPNServerExists("ibm_is_vpn_server.is_vpn_server", vpnserver),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "certificate_crn"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_authentication.0.method"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_auto_delete"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_auto_delete_timeout"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "client_dns_server_ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "hostname"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "health_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "vpn_server"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.address"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.id"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "private_ips.0.resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "security_groups.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "resource_type"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "subnets.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "vpc.#"),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_ip_pool", clientIPPool),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_idle_timeout", clientIdleTimeout),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "enable_split_tunneling", enableSplitTunneling),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "name", name),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "port", port),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "protocol", protocol),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "tags.#"),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMIsVPNServerConfigBasicTags(nameVpc, nameSubnet1, clientIPPoolUpdate, clientIdleTimeoutUpdate, enableSplitTunnelingUpdate, nameUpdate, portUpdate, protocolUpdate, isCertificateCrn, isClientCaCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_ip_pool", clientIPPoolUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "client_idle_timeout", clientIdleTimeoutUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "enable_split_tunneling", enableSplitTunnelingUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "port", portUpdate),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "protocol", protocolUpdate),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_server.is_vpn_server", "tags.#"),
					resource.TestCheckResourceAttr("ibm_is_vpn_server.is_vpn_server", "tags.#", "2"),
				),
			},
			{
				ResourceName:      "ibm_is_vpn_server.is_vpn_server",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIsVPNServerConfigBasic(nameVpc string, nameSubnet1 string, clientIPPool string, clientIdleTimeout string, enableSplitTunneling string, vpnServerName string, port string, protocol string, isCertificateCrn string, isClientCaCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s" 
		}
		
		resource "ibm_is_subnet" "testacc_subnet-1" {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
			zone = "us-south-1"
			ipv4_cidr_block = "10.240.0.0/24"
		}
		
		resource "ibm_is_vpn_server" "is_vpn_server" {
			certificate_crn = "%s"
			client_authentication {
				method = "certificate"
				client_ca_crn = "%s"
			}
			client_ip_pool = "%s"
			subnets = [ibm_is_subnet.testacc_subnet-1.id]
			client_dns_server_ips = ["192.168.3.4"]
			client_idle_timeout = %s
			enable_split_tunneling = %s
			name = "%s"
			port = %s
			protocol = "%s"
		}
	`, nameVpc, nameSubnet1, isCertificateCrn, isClientCaCrn, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol)
}
func testAccCheckIBMIsVPNServerConfigBasicTags(nameVpc string, nameSubnet1 string, clientIPPool string, clientIdleTimeout string, enableSplitTunneling string, vpnServerName string, port string, protocol string, isCertificateCrn string, isClientCaCrn string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s" 
		}
		
		resource "ibm_is_subnet" "testacc_subnet-1" {
			name = "%s"
			vpc = ibm_is_vpc.testacc_vpc.id
			zone = "us-south-1"
			ipv4_cidr_block = "10.240.0.0/24"
		}
		
		resource "ibm_is_vpn_server" "is_vpn_server" {
			certificate_crn = "%s"
			client_authentication {
				method = "certificate"
				client_ca_crn = "%s"
			}
			client_ip_pool = "%s"
			subnets = [ibm_is_subnet.testacc_subnet-1.id]
			client_dns_server_ips = ["192.168.3.4"]
			client_idle_timeout = %s
			enable_split_tunneling = %s
			name = "%s"
			port = %s
			protocol = "%s"
			tags = [ "test:tags", "test:tags2" ]
		}
	`, nameVpc, nameSubnet1, isCertificateCrn, isClientCaCrn, clientIPPool, clientIdleTimeout, enableSplitTunneling, vpnServerName, port, protocol)
}

func testAccCheckIBMIsVPNServerExists(n string, obj string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getVPNServerOptions := &vpcv1.GetVPNServerOptions{}

		getVPNServerOptions.SetID(rs.Primary.ID)

		vpnServer, _, err := sess.GetVPNServer(getVPNServerOptions)
		if err != nil {
			return err
		}

		obj = *vpnServer.ID
		return nil
	}
}

func testAccCheckIBMIsVPNServerDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_server" {
			continue
		}

		getVPNServerOptions := &vpcv1.GetVPNServerOptions{}

		getVPNServerOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := sess.GetVPNServer(getVPNServerOptions)

		if err == nil {
			return fmt.Errorf("VPNServer still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VPNServer (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
