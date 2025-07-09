// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPNGatewayConnection_basic(t *testing.T) {
	var VPNGatewayConnection string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(10, 100))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(10, 100))
	updname2 := fmt.Sprintf("tfvpngc-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayConnectionUpdate(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, updname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", updname2),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "interval"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "status"),
				),
			},
		},
	})
}

func TestAccIBMISVPNGatewayConnection_route(t *testing.T) {
	var VPNGatewayConnection string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	updname2 := fmt.Sprintf("tfvpngc-updatename-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionRouteConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayConnectionRouteUpdate(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, updname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", updname2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
				),
			},
		},
	})
}
func TestAccIBMISVPNGatewayConnection_admin_state(t *testing.T) {
	var VPNGatewayConnection string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	name3 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	name4 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	adminStateUp1 := true
	adminStateUp1Update := false
	adminStateUp2 := true
	adminStateUp2Update := false
	updname2 := fmt.Sprintf("tfvpngc-updatename-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionAdminStateConfig(vpcname1, subnetname1, vpnname1, name1, name2, vpcname2, subnetname2, vpnname2, name3, name4, adminStateUp1, adminStateUp2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "name", name3),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "admin_state_up", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "name", name4),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "admin_state_up", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "status"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayConnectionAdminStateConfig(vpcname1, subnetname1, vpnname1, name1, updname2, vpcname2, subnetname2, vpnname2, name3, name4, adminStateUp1Update, adminStateUp2Update),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", updname2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up", "false"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "name", name3),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "admin_state_up", "true"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "name", name4),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "admin_state_up"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "admin_state_up", "false"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "status"),
				),
			},
		},
	})
}

// distribute_traffic
func TestAccIBMISVPNGatewayConnection_routeDistributeTraffic(t *testing.T) {
	var VPNGatewayConnection string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	updname2 := fmt.Sprintf("tfvpngc-updatename-%d", acctest.RandIntRange(100, 200))
	dt := true
	dt2 := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionRouteDistributeTrafficConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2, dt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "distribute_traffic"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "distribute_traffic", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "distribute_traffic", fmt.Sprintf("%t", dt)),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayConnectionRouteDistributeTrafficUpdate(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, updname2, dt2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", updname2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "distribute_traffic"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "distribute_traffic", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "distribute_traffic", fmt.Sprintf("%t", dt2)),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayConnectionRouteDistributeTrafficConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string, distributeTraffic bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		distribute_traffic = %t
	}
	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name2, distributeTraffic)

}

func testAccCheckIBMISVPNGatewayConnectionRouteDistributeTrafficUpdate(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string, distributeTraffic bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		distribute_traffic = %t
	}
	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name2, distributeTraffic)

}

func TestAccIBMISVPNGatewayConnection_multiple(t *testing.T) {
	var VPNGatewayConnection string
	var VPNGatewayConnection2 string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionMultipleConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "status"),
				),
			},
		},
	})
}
func TestAccIBMISVPNGatewayConnection_advanced(t *testing.T) {
	var VPNGatewayConnection string
	var VPNGatewayConnection2 string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	subnetname3 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	subnetname4 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))

	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	name3 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	name4 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionAdvanceConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2, subnetname3, subnetname4, name3, name4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_VPNGateway1", "name", vpnname1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_VPNGateway2", "name", vpnname2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "name"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection3", "name", name3),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection4", "name", name4),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "status"),
				),
			},
		},
	})
}
func TestAccIBMISVPNGatewayConnection_breakingchange(t *testing.T) {
	var VPNGatewayConnection string
	var VPNGatewayConnection2 string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionBreakingChangeConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "status"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", name2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "action"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "admin_state_up"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "authentication_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "establish_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "href"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "interval"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "preshared_key"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "resource_type", "vpn_gateway_connection"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayConnectionDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_gateway_connection" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		gID := parts[0]
		gConnID := parts[1]

		getvpngcoptions := &vpcv1.GetVPNGatewayConnectionOptions{
			VPNGatewayID: &gID,
			ID:           &gConnID,
		}
		_, _, err1 := sess.GetVPNGatewayConnection(getvpngcoptions)

		if err1 == nil {
			return fmt.Errorf("VPNGatewayConnection still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayConnectionExists(n, vpngcID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		gID := parts[0]
		gConnID := parts[1]

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvpngcoptions := &vpcv1.GetVPNGatewayConnectionOptions{
			VPNGatewayID: &gID,
			ID:           &gConnID,
		}
		vpnGatewayConnectionIntf, res, err := sess.GetVPNGatewayConnection(getvpngcoptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting VPN Gateway connection: %s\n%s", err, res)
		}

		if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionPolicyMode)
			vpngcID = *vpnGatewayConnection.ID
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteMode)
			vpngcID = *vpnGatewayConnection.ID
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnectionRouteModeVPNGatewayConnectionStaticRouteMode)
			vpngcID = *vpnGatewayConnection.ID
		} else if _, ok := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection); ok {
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection)
			vpngcID = *vpnGatewayConnection.ID
		} else {
			return fmt.Errorf("[ERROR] Unrecognized vpcv1.vpnGatewayConnectionIntf subtype encountered")
		}
		return nil
	}
}

func testAccCheckIBMISVPNGatewayConnectionConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]

	}

	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]

	}

	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name2)

}

func testAccCheckIBMISVPNGatewayConnectionUpdate(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]

	}

	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]

	}

	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name2)

}

func testAccCheckIBMISVPNGatewayConnectionRouteConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name2)

}
func testAccCheckIBMISVPNGatewayConnectionAdminStateConfig(vpc1, subnet1, vpnname1, name1, name2, vpc2, subnet2, vpnname2, name3, name4 string, adminStateUp1, adminStateUp2 bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name 			= "%s"
		vpc 			= "${ibm_is_vpc.testacc_vpc1.id}"
		zone 			= "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name 	= "%s"
		subnet 	= "${ibm_is_subnet.testacc_subnet1.id}"
		mode 	= "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name 			= "%s"
		vpn_gateway 	= "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address 	= "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key 	= "VPNDemoPassword"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name 			= "%s"
		admin_state_up	= %t
		vpn_gateway 	= "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address 	= "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key 	= "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name 				= "%s"
		vpc 				= "${ibm_is_vpc.testacc_vpc2.id}"
		zone 				= "%s"
		ipv4_cidr_block 	= "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name 	= "%s"
		subnet 	= "${ibm_is_subnet.testacc_subnet2.id}"
		mode 	= "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection3" {
		name 			= "%s"
		vpn_gateway 	= "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address 	= "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key 	= "VPNDemoPassword"
		local_cidrs   	= ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
		peer_cidrs    	= ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection4" {
		name 			= "%s"
		admin_state_up	= %t
		vpn_gateway 	= "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address 	= "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key 	= "VPNDemoPassword"
		local_cidrs   	= ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
		peer_cidrs    	= ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
	}
	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, name2, adminStateUp1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name3, name4, adminStateUp2)

}
func testAccCheckIBMISVPNGatewayConnectionMultipleConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc1.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet1.id
		mode   = "policy"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name          	= "%s"
		vpn_gateway   	= ibm_is_vpn_gateway.testacc_VPNGateway1.id
		peer_cidrs		= [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
		peer_address  	= cidrhost(ibm_is_subnet.testacc_subnet1.ipv4_cidr_block, 14)
		local_cidrs 	= [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
		preshared_key 	= "VPNDemoPassword"
	  }
	  resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc2.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet2.id
		mode   = "route"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name          = "%s"
		vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway2.id
		peer_address  = cidrhost(ibm_is_subnet.testacc_subnet2.ipv4_cidr_block, 15)
		preshared_key = "VPNDemoPassword"
	  }
	`, vpc1, subnet1, acc.ISZoneName, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, vpnname2, name2)

}
func testAccCheckIBMISVPNGatewayConnectionBreakingChangeConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc1.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet1.id
		mode   = "policy"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name          	= "%s"
		vpn_gateway   	= ibm_is_vpn_gateway.testacc_VPNGateway1.id
		peer_cidrs		= [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
		peer_address  	= cidrhost(ibm_is_subnet.testacc_subnet1.ipv4_cidr_block, 14)
		local_cidrs 	= [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
		preshared_key 	= "VPNDemoPassword"
	  }
	  resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc2.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet2.id
		mode   = "route"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name          = "%s"
		vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway2.id
		peer_address  = cidrhost(ibm_is_subnet.testacc_subnet2.ipv4_cidr_block, 15)
		preshared_key = "VPNDemoPassword"
	  }
	`, vpc1, subnet1, acc.ISZoneName, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, vpnname2, name2)

}
func testAccCheckIBMISVPNGatewayConnectionAdvanceConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2, subnet3, subnet4, name3, name4 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc1.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_subnet" "testacc_subnet3" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc1.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet1.id
		mode   = "policy"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name          	= "%s"
		vpn_gateway   	= ibm_is_vpn_gateway.testacc_VPNGateway1.id
		peer_cidrs		= [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
		peer_address  	= cidrhost(ibm_is_subnet.testacc_subnet1.ipv4_cidr_block, 14)
		local_cidrs 	= [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
		preshared_key 	= "VPNDemoPassword"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection3" {
		name          	= "%s"
		vpn_gateway   	= ibm_is_vpn_gateway.testacc_VPNGateway1.id
		peer {
			cidrs   = [ibm_is_subnet.testacc_subnet3.ipv4_cidr_block]
			address = cidrhost(ibm_is_subnet.testacc_subnet3.ipv4_cidr_block, 14)
		}
		local {
			cidrs = [ibm_is_subnet.testacc_subnet3.ipv4_cidr_block]
		}
		preshared_key 	= "VPNDemoPassword"
	  }

	  resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	  }
	  resource "ibm_is_subnet" "testacc_subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc2.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_subnet" "testacc_subnet4" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc2.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	  }
	  resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet2.id
		mode   = "route"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name          = "%s"
		vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway2.id
		peer_address  = cidrhost(ibm_is_subnet.testacc_subnet2.ipv4_cidr_block, 15)
		preshared_key = "VPNDemoPassword"
	  }
	  resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection4" {
		name          = "%s"
		vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway2.id
		peer {
				address  = cidrhost(ibm_is_subnet.testacc_subnet4.ipv4_cidr_block, 15)
			}	
		preshared_key = "VPNDemoPassword"
	  }
	`, vpc1, subnet1, acc.ISZoneName, subnet3, acc.ISZoneName, vpnname1, name1, name3, vpc2, subnet2, acc.ISZoneName, subnet4, acc.ISZoneName, vpnname2, name2, name4)

}

func testAccCheckIBMISVPNGatewayConnectionRouteUpdate(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	`, vpc1, subnet1, acc.ISZoneName, acc.ISCIDR, vpnname1, name1, vpc2, subnet2, acc.ISZoneName, acc.ISCIDR, vpnname2, name2)

}

func TestAccIBMISVPNGatewayConnection_ike_ipsec_null_patch(t *testing.T) {
	var VPNGatewayConnection string
	vpcname := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(10, 100))
	noNullPass := ""
	nullPass := "null"
	ikepolicyname := fmt.Sprintf("tfvpngc-ike-%d", acctest.RandIntRange(10, 100))
	ipsecpolicyname := fmt.Sprintf("tfvpngc-ipsec-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConnectionNullPatchConfig(vpcname, subnetname, vpnname, ikepolicyname, ipsecpolicyname, name, noNullPass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "gateway_connection"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", vpcname),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet1", "name", subnetname),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_VPNGateway1", "name", vpnname),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.testacc_ike", "name", ikepolicyname),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.testacc_ipsec", "name", ipsecpolicyname),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "ike_policy"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "ipsec_policy"),
				),
			},
			{
				Config: testAccCheckIBMISVPNGatewayConnectionNullPatchConfig(vpcname, subnetname, vpnname, ikepolicyname, ipsecpolicyname, name, nullPass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpc.testacc_vpc1", "name", vpcname),
					resource.TestCheckResourceAttr(
						"ibm_is_subnet.testacc_subnet1", "name", subnetname),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_VPNGateway1", "name", vpnname),
					resource.TestCheckResourceAttr(
						"ibm_is_ike_policy.testacc_ike", "name", ikepolicyname),
					resource.TestCheckResourceAttr(
						"ibm_is_ipsec_policy.testacc_ipsec", "name", ipsecpolicyname),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "ike_policy", ""),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "ipsec_policy", ""),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayConnectionNullPatchConfig(vpc, subnet, vpnname, ikepolicyname, ipsecpolicyname, name, noNullPass string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		timeouts {
			create = "18m"
			delete = "18m"
		}
	}
	resource "ibm_is_ike_policy" "testacc_ike" {
		name                     = "%s"
		authentication_algorithm = "md5"
		encryption_algorithm     = "triple_des"
		dh_group                 = 2
		ike_version              = 1
	}
	resource "ibm_is_ipsec_policy" "testacc_ipsec" {
		name                     = "%s"
		authentication_algorithm = "md5"
		encryption_algorithm     = "triple_des"
		pfs                      = "disabled"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name 				= "%s"
		vpn_gateway 		= "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		preshared_key 		= "VPNDemoPassword"
		peer_address 		= ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address != "0.0.0.0" ? ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address : ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address2
		ike_policy 			= "%s" == "null" ? "" : ibm_is_ike_policy.testacc_ike.id
		ipsec_policy  		= "%s" == "null" ? "" : ibm_is_ipsec_policy.testacc_ipsec.id
	}

	`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpnname, ikepolicyname, ipsecpolicyname, name, noNullPass, noNullPass)

}

func TestAccIBMISVPNGatewayConnection_CIDRUpdates(t *testing.T) {
	var VPNGatewayConnection string
	vpcname := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name := fmt.Sprintf("tfvpngc-conn-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			// Initial configuration
			{
				Config: testAccCheckIBMISVPNGatewayConnectionCIDRConfig(vpcname, subnetname1, subnetname2, vpnname, name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "peer.0.cidrs.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "local.0.cidrs.#", "1"),
				),
			},
			// Add additional CIDRs
			{
				Config: testAccCheckIBMISVPNGatewayConnectionCIDRConfig(vpcname, subnetname1, subnetname2, vpnname, name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "peer.0.cidrs.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "local.0.cidrs.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayConnectionCIDRConfig(vpc, subnet1, subnet2, vpnname, name string, additionalCIDRs bool) string {
	base := fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet1" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	}
	
	resource "ibm_is_subnet" "testacc_subnet2" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	}
	
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet1.id
		mode   = "policy"
	}
	`, vpc, subnet1, acc.ISZoneName, subnet2, acc.ISZoneName, vpnname)

	if !additionalCIDRs {
		return base + fmt.Sprintf(`
		resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection" {
			name          = "%s"
			vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway.id
			peer {
				cidrs    = [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
				address  = cidrhost(ibm_is_subnet.testacc_subnet1.ipv4_cidr_block, 14)
			}
			local {
				cidrs    = [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
			}
			preshared_key = "VPNDemoPassword"
		}
		`, name)
	}

	return base + fmt.Sprintf(`
		resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection" {
			name          = "%s"
			vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway.id
			peer {
				cidrs    = [
					ibm_is_subnet.testacc_subnet1.ipv4_cidr_block,
					ibm_is_subnet.testacc_subnet2.ipv4_cidr_block
				]
				address  = cidrhost(ibm_is_subnet.testacc_subnet1.ipv4_cidr_block, 14)
			}
			local {
				cidrs    = [
					ibm_is_subnet.testacc_subnet1.ipv4_cidr_block,
					ibm_is_subnet.testacc_subnet2.ipv4_cidr_block
				]
			}
			preshared_key = "VPNDemoPassword"
		}
		`, name)
}
