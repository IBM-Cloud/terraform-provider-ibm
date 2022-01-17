// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPrivateDNSGlbLoadBalancer_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnspn%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	newName := fmt.Sprintf("Test-load-balancer.%s", name)
	updateName := fmt.Sprintf("Update-load-balancer.%s", name)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbLoadBalancerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbLoadBalancerExists("ibm_dns_glb.test-pdns-lb", &resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_glb.test-pdns-lb", "name", newName),
					resource.TestCheckResourceAttr("ibm_dns_glb.test-pdns-lb", "description", "new lb"),
				),
			},
			{
				Config: testAccCheckIBMPrivateDNSGlbUpdateLoadBalancerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbLoadBalancerExists("ibm_dns_glb.test-pdns-lb", &resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_glb.test-pdns-lb", "name", updateName),
					resource.TestCheckResourceAttr("ibm_dns_glb.test-pdns-lb", "description", "update lb"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSGlboadBalancerImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSGlbMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbLoadBalancerBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSGlbLoadBalancerExists("ibm_dns_glb.test-pdns-lb", &resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_glb.test-pdns-lb", "ttl", "120"),
				),
			},
			{
				ResourceName:      "ibm_dns_glb.test-pdns-lb",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPrivateDNSGlbLoadBalancerBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	  }

	  resource "ibm_resource_instance" "test-pdns-instance" {
		name              = "test-pdns-glb-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location          = "global"
		service           = "dns-svcs"
		plan              = "standard-dns"
	  }

	  resource "ibm_dns_zone" "test-pdns-zone" {
		depends_on  = [ibm_resource_instance.test-pdns-instance]
		name        = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription"
		label       = "testlabel-updated"
	  }

	  resource "ibm_dns_glb_pool" "test-pdns-pool" {
		depends_on                = [ibm_dns_zone.test-pdns-zone]
		name                      = "testpool"
		instance_id               = ibm_resource_instance.test-pdns-instance.guid
		description               = "test pool"
		enabled                   = true
		healthy_origins_threshold = 1
		origins {
		  name        = "example-1"
		  address     = "www.google.com"
		  enabled     = true
		  description = "origin pool"
		}
	  }

	  resource "ibm_dns_glb" "test-pdns-lb" {
		depends_on    = [ibm_dns_zone.test-pdns-zone]
		name          = "Test-load-balancer"
		instance_id   = ibm_resource_instance.test-pdns-instance.guid
		zone_id       = ibm_dns_zone.test-pdns-zone.zone_id
		description   = "new lb"
		ttl           = 120
		fallback_pool = ibm_dns_glb_pool.test-pdns-pool.pool_id
		default_pools = [ibm_dns_glb_pool.test-pdns-pool.pool_id]
		az_pools {
		  availability_zone = "us-south-1"
		  pools             = [ibm_dns_glb_pool.test-pdns-pool.pool_id]
		}
	  }
	  `, name)

}

func testAccCheckIBMPrivateDNSGlbUpdateLoadBalancerBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	  }

	  resource "ibm_resource_instance" "test-pdns-instance" {
		name              = "test-pdns-glb-monitor-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location          = "global"
		service           = "dns-svcs"
		plan              = "standard-dns"
	  }

	  resource "ibm_dns_zone" "test-pdns-zone" {
		depends_on  = [ibm_resource_instance.test-pdns-instance]
		name        = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription"
		label       = "testlabel-updated"
	  }

	  resource "ibm_dns_glb_pool" "test-pdns-pool" {
		depends_on                = [ibm_dns_zone.test-pdns-zone]
		name                      = "testpool"
		instance_id               = ibm_resource_instance.test-pdns-instance.guid
		description               = "test pool"
		enabled                   = true
		healthy_origins_threshold = 1
		origins {
		  name        = "example-1"
		  address     = "www.google.com"
		  enabled     = true
		  description = "origin pool"
		}
	  }
	  resource "ibm_dns_glb" "test-pdns-lb" {
		depends_on    = [ibm_dns_zone.test-pdns-zone]
		name          = "Update-load-balancer"
		instance_id   = ibm_resource_instance.test-pdns-instance.guid
		zone_id       = ibm_dns_zone.test-pdns-zone.zone_id
		description   = "update lb"
		ttl           = 120
		fallback_pool = ibm_dns_glb_pool.test-pdns-pool.pool_id
		default_pools = [ibm_dns_glb_pool.test-pdns-pool.pool_id]
		az_pools {
		  availability_zone = "us-south-1"
		  pools             = [ibm_dns_glb_pool.test-pdns-pool.pool_id]
		}
	  }
	  `, name)

}

func testAccCheckIBMPrivateDNSGlbDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_glb" {
			continue
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getLbOptions := pdnsClient.NewGetLoadBalancerOptions(partslist[0], partslist[1], partslist[2])
		_, res, err := pdnsClient.GetLoadBalancer(getLbOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSGlbLoadBalancerExists(n string, result *string) resource.TestCheckFunc {
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

		getLbOptions := pdnsClient.NewGetLoadBalancerOptions(partslist[0], partslist[1], partslist[2])
		r, res, err := pdnsClient.GetLoadBalancer(getLbOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
		*result = *r.ID
		return nil
	}
}
