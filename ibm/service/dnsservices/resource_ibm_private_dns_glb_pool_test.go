// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMPrivateDNSGlbPool_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnspn%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMGlbPoolBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMGlbPoolExists("ibm_dns_glb_pool.test-pdns-pool-nw", resultprivatedns),

					resource.TestCheckResourceAttr("ibm_dns_glb_pool.test-pdns-pool-nw", "name", "testpool"),
					resource.TestCheckResourceAttr("ibm_dns_glb_pool.test-pdns-pool-nw", "healthy_origins_threshold", "1"), // default value
				),
			},
			{
				Config: testAccCheckIBMGlbPoolUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMGlbPoolExists("ibm_dns_glb_pool.test-pdns-pool-nw", resultprivatedns),

					resource.TestCheckResourceAttr("ibm_dns_glb_pool.test-pdns-pool-nw", "name", "testpoolUpdate"),
					resource.TestCheckResourceAttr("ibm_dns_glb_pool.test-pdns-pool-nw", "healthy_origins_threshold", "1"), // default value
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSGlbPoolImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbPoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMGlbPoolBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMGlbPoolExists("ibm_dns_glb_pool.test-pdns-pool-nw", resultprivatedns),
				),
			},
			{
				ResourceName:      "ibm_dns_glb_pool.test-pdns-pool-nw",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMGlbPoolBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}
	resource "ibm_is_vpc" "test-pdns-glb-pool-vpc" {
		name = "test-pdns-glb-pool-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-glb-subnet" {
		name                     = "test-pdns-glb-subnet"
		vpc                      = ibm_is_vpc.test-pdns-glb-pool-vpc.id
		zone            = "us-south-1"
		ipv4_cidr_block = "10.240.25.0/24"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-glb-pool-instance" {
		name = "test-pdns-glb-pool-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}
	resource "ibm_dns_zone" "test-pdns-glb-pool-zone" {
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-glb-pool-instance.guid
		description = "testdescription"
		label = "testlabel"
	}
	resource "ibm_dns_glb_monitor" "test-pdns-glb-monitor" {
		depends_on = [ibm_dns_zone.test-pdns-glb-pool-zone]
		name = "test-pdns-glb-monitor"
		instance_id = ibm_resource_instance.test-pdns-glb-pool-instance.guid
		description = "Monitor description"
		interval=60
		retries=3
		timeout=8
		port=8080
		type="HTTPS"
    }
	resource "ibm_dns_glb_pool" "test-pdns-pool-nw" {
		depends_on = [ibm_dns_zone.test-pdns-glb-pool-zone]
		name = "testpool"
		instance_id = ibm_resource_instance.test-pdns-glb-pool-instance.guid
		description = "New test pool"
		enabled=true
		healthy_origins_threshold=1
		origins {
				name    = "example-1"
				address = "www.google.com"
				enabled = true
				description="origin pool"
		}
		monitor=ibm_dns_glb_monitor.test-pdns-glb-monitor.monitor_id
		notification_channel="https://mywebsite.com/dns/webhook"
		healthcheck_region="us-south"
		healthcheck_subnets=[ibm_is_subnet.test-pdns-glb-subnet.resource_crn]
    }
	  `, name)
}

func testAccCheckIBMGlbPoolUpdate(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}
	resource "ibm_is_vpc" "test-pdns-glb-pool-vpc" {
		name = "test-pdns-glb-pool-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-glb-subnet" {
		name                     = "test-pdns-glb-pool-subnet"
		vpc                      = ibm_is_vpc.test-pdns-glb-pool-vpc.id
		zone            = "us-south-1"
		ipv4_cidr_block = "10.240.25.0/24"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-glb-pool-instance" {
		name = "test-pdns-glb-pool-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}
	resource "ibm_dns_zone" "test-pdns-glb-pool-zone" {
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-glb-pool-instance.guid
		description = "testdescription"
		label = "testlabel"
	}
	resource "ibm_dns_glb_monitor" "test-pdns-glb-monitor" {
		depends_on = [ibm_dns_zone.test-pdns-glb-pool-zone]
		name = "test-pdns-glb-monitor"
		instance_id = ibm_resource_instance.test-pdns-glb-pool-instance.guid
		description = "Monitor description"
		interval=60
		retries=3
		timeout=8
		port=8080
		type="HTTPS"
    }
	resource "ibm_dns_glb_pool" "test-pdns-pool-nw" {
		depends_on = [ibm_dns_zone.test-pdns-glb-pool-zone]
		name = "testpoolUpdate"
		instance_id = ibm_resource_instance.test-pdns-glb-pool-instance.guid
		description = "Update test pool"
		enabled=true
		healthy_origins_threshold=1
		origins {
				name    = "example-1"
				address = "www.google.com"
				enabled = true
				description="origin pool"
		}
		monitor=ibm_dns_glb_monitor.test-pdns-glb-monitor.monitor_id
		notification_channel="https://mywebsite.com/dns/webhook"
		healthcheck_region="us-south"
		healthcheck_subnets=[ibm_is_subnet.test-pdns-glb-subnet.resource_crn]
    }
	  `, name)
}

func testAccCheckIBMPrivateDNSGlbPoolDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_glb_pool" {
			continue
		}

		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getGlbPoolOptions := pdnsClient.NewGetPoolOptions(partslist[0], partslist[1])
		_, res, err := pdnsClient.GetPool(getGlbPoolOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return fmt.Errorf("testAccCheckIBMPrivateDNSGlbPoolDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMGlbPoolExists(n string, result string) resource.TestCheckFunc {

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

		getGlbPoolOptions := pdnsClient.NewGetPoolOptions(partslist[0], partslist[1])
		r, res, err := pdnsClient.GetPool(getGlbPoolOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return fmt.Errorf("testAccCheckIBMGlbPoolExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}
