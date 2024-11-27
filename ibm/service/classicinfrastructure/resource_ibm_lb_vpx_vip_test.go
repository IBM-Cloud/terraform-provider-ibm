// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/softlayer/softlayer-go/helpers/network"
)

func TestAccIBMLbVpxVip_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbVpxVipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMLbVpxVipConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					// Test VPX 10.1
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "load_balancing_method", "lc"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "name", "test_load_balancer_vip"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "source_port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "type", "HTTP"),
					// Test VPX 10.5
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip105", "load_balancing_method", "lc"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip105", "name", "test_load_balancer_vip105"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip105", "source_port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip105", "type", "HTTP"),
				),
			},
		},
	})
}

func TestAccIBMLbVpxVipWithTag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLbVpxVipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMLbVpxVipWithTag,
				Check: resource.ComposeTestCheckFunc(
					// Test VPX 10.1
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "load_balancing_method", "lc"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "name", "test_load_balancer_vip"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "source_port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMLbVpxVipWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					// Test VPX 10.1
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "load_balancing_method", "lc"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "name", "test_load_balancer_vip"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "source_port", "80"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "type", "HTTP"),
					resource.TestCheckResourceAttr(
						"ibm_lb_vpx_vip.testacc_vip", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMLbVpxVipDestroy(s *terraform.State) error {
	sess := acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_lb_vpx_vip" {
			continue
		}

		nadcId, _ := strconv.Atoi(rs.Primary.Attributes["nad_controller_id"])
		vipName, _ := rs.Primary.Attributes["name"]

		vip, _ := network.GetNadcLbVipByName(sess, nadcId, vipName)

		if vip != nil {
			return fmt.Errorf("Netscaler VPX VIP still exists")
		}
	}

	return nil
}

var testAccCheckIBMLbVpxVipConfig_basic = `
resource "ibm_lb_vpx" "testacc_foobar_nadc" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "${ibm_lb_vpx.testacc_foobar_nadc.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "${ibm_lb_vpx.testacc_foobar_nadc.vip_pool[0]}"
}

resource "ibm_lb_vpx" "testacc_foobar_nadc105" {
    datacenter = "dal09"
    speed = 10
    version = "10.5"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip105" {
    name = "test_load_balancer_vip105"
    nad_controller_id = "${ibm_lb_vpx.testacc_foobar_nadc105.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
    virtual_ip_address = "${ibm_lb_vpx.testacc_foobar_nadc105.vip_pool[0]}"
}
`
var testAccCheckIBMLbVpxVipWithTag = `
resource "ibm_lb_vpx" "testacc_foobar_nadc" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "${ibm_lb_vpx.testacc_foobar_nadc.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
	virtual_ip_address = "${ibm_lb_vpx.testacc_foobar_nadc.vip_pool[0]}"
	tags = ["one", "two"]
}
`
var testAccCheckIBMLbVpxVipWithUpdatedTag = `
resource "ibm_lb_vpx" "testacc_foobar_nadc" {
    datacenter = "dal09"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "${ibm_lb_vpx.testacc_foobar_nadc.id}"
    load_balancing_method = "lc"
    source_port = 80
    type = "HTTP"
	virtual_ip_address = "${ibm_lb_vpx.testacc_foobar_nadc.vip_pool[0]}"
	tags = ["one", "two", "three"]
}
`
