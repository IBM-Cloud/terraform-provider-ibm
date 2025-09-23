// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPrivateDNSGlbMonitor_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnspn%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbMonitorBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbMonitorExists("ibm_dns_glb_monitor.test-pdns-monitor", resultprivatedns),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr("ibm_dns_glb_monitor.test-pdns-monitor", "name", "test-pdns-glb-monitor"),
					resource.TestCheckResourceAttr("ibm_dns_glb_monitor.test-pdns-monitor", "description", "Monitordescription"),
					resource.TestCheckResourceAttr("ibm_dns_glb_monitor.test-pdns-monitor", "port", "80"), // default value
				),
			},
			{
				Config: testAccCheckIBMPrivateDNSGlbUpdateMonitorBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbMonitorExists("ibm_dns_glb_monitor.test-pdns-monitor", resultprivatedns),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr("ibm_dns_glb_monitor.test-pdns-monitor", "name", "test-pdns-glb-monitor-update"),
					resource.TestCheckResourceAttr("ibm_dns_glb_monitor.test-pdns-monitor", "description", "UpdatedMonitordescription"),
					resource.TestCheckResourceAttr("ibm_dns_glb_monitor.test-pdns-monitor", "port", "80"), // default value
				),
			},
		},
	})

}

func TestAccIBMPrivateDNSGlbMonitorImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbMonitorBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbMonitorExists("ibm_dns_glb_monitor.test-pdns-monitor", resultprivatedns),
				),
			},
			{
				ResourceName:      "ibm_dns_glb_monitor.test-pdns-monitor",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})

}

func testAccCheckIBMPrivateDNSGlbMonitorBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
    }

    resource "ibm_is_vpc" "test-pdns-vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "test-pdns-glb-monitor-vpc"
		resource_group = data.ibm_resource_group.rg.id
    }

    resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test-pdns-vpc]
		name = "test-pdns-glb-monitor-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
    }

    resource "ibm_dns_zone" "test-pdns-zone" {
		depends_on = [ibm_resource_instance.test-pdns-instance]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription"
		label = "testlabel-updated"
    }

	resource "ibm_dns_glb_monitor" "test-pdns-monitor" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "test-pdns-glb-monitor"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "Monitordescription"
		interval=63
		retries=3
		timeout=8
		type="HTTP"
		expected_codes= "200"
		path="/health"
		method="GET"
		expected_body="alive"
		headers{
			name="headerName1"
			value=["example1","abc1"]
		}
		headers{
			name="headerName2"
			value=["example2","abc2"]
		}
		headers{
			name="headerName3"
			value=["example3","ab3c"]
		}
    }
	  `, name)

}

func testAccCheckIBMPrivateDNSGlbUpdateMonitorBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
    }

    resource "ibm_is_vpc" "test-pdns-vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "test-pdns-glb-monitor-vpc"
		resource_group = data.ibm_resource_group.rg.id
    }

    resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test-pdns-vpc]
		name = "test-pdns-glb-monitor-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
    }

    resource "ibm_dns_zone" "test-pdns-zone" {
		depends_on = [ibm_resource_instance.test-pdns-instance]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription"
		label = "testlabel-updated"
    }

	resource "ibm_dns_glb_monitor" "test-pdns-monitor" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "test-pdns-glb-monitor-update"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "UpdatedMonitordescription"
		interval=63
		retries=3
		timeout=8
		type="HTTP"
		expected_codes= "200"
		path="/health"
		method="GET"
		expected_body="alive"
		headers{
			name="headerName2"
			value=["example2","abc2"]
		}
		headers{
			name="headerName4"
			value=["example4","abc4"]
		}
		headers{
			name="headerName3"
			value=["example3","ab3c"]
		}
    }
	  `, name)

}

func testAccCheckIBMPrivateDNSGlbMonitorDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_glb_monitor" {
			continue
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getMonitorOptions := pdnsClient.NewGetMonitorOptions(partslist[0], partslist[1])
		_, res, err := pdnsClient.GetMonitor(getMonitorOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSGlbMonitorExists(n string, result string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getMonitorOptions := pdnsClient.NewGetMonitorOptions(partslist[0], partslist[1])
		r, res, err := pdnsClient.GetMonitor(getMonitorOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
		result = *r.ID
		return nil
	}
}
