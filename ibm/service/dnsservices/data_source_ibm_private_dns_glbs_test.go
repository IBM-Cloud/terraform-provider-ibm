// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSGlbLoadBalancersDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_glbs.test1"
	riname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	poolname := fmt.Sprintf("tf-poolname-%d", acctest.RandIntRange(100, 200))
	lbname := fmt.Sprintf("tf-lbname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbLoadBalancerdDataSConfig(riname, zonename, poolname, lbname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "dns_glbs.0.name"),
					resource.TestCheckResourceAttrSet(node, "dns_glbs.0.description"),
					resource.TestCheckResourceAttrSet(node, "dns_glbs.0.ttl"),
					resource.TestCheckResourceAttrSet(node, "dns_glbs.0.fallback_pool"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSGlbLoadBalancerdDataSConfig(riname, zonename, poolname, lbname string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	  }

	  resource "ibm_resource_instance" "test_pdns_instance" {
		name              = "%[1]s"
		resource_group_id = data.ibm_resource_group.rg.id
		location          = "global"
		service           = "dns-svcs"
		plan              = "standard-dns"
	  }

	  resource "ibm_dns_zone" "test_pdns_glb_zone" {
		depends_on  = [ibm_resource_instance.test_pdns_instance]
		name        = "%[2]s"
		instance_id = ibm_resource_instance.test_pdns_instance.guid
		description = "testdescription"
		label       = "testlabel-updated"
	  }

	  resource "ibm_dns_glb_pool" "test_pdns_glb_pool" {
		depends_on                = [ibm_dns_zone.test_pdns_glb_zone]
		name                      = "%[3]s"
		instance_id               = ibm_resource_instance.test_pdns_instance.guid
		description               = "new pool update"
		enabled                   = "true"
		healthy_origins_threshold = 1
		origins {
		  name        = "example-1"
		  address     = "www.google.com"
		  enabled     = true
		  description = "origin pool"
		}
	  }

	  resource "ibm_dns_glb" "test_pdns_glb" {
		depends_on    = [ibm_dns_glb_pool.test_pdns_glb_pool]
		name          = "%[4]s"
		instance_id   = ibm_resource_instance.test_pdns_instance.guid
		zone_id       = ibm_dns_zone.test_pdns_glb_zone.zone_id
		description   = "new glb"
		ttl           = 120
		fallback_pool = ibm_dns_glb_pool.test_pdns_glb_pool.pool_id
		default_pools = [ibm_dns_glb_pool.test_pdns_glb_pool.pool_id]
		az_pools {
		  availability_zone = "us-south-1"
		  pools             = [ibm_dns_glb_pool.test_pdns_glb_pool.pool_id]
		}
	  }

	  data "ibm_dns_glbs" "test1" {
		instance_id = ibm_dns_glb.test_pdns_glb.instance_id
		zone_id     = ibm_dns_glb.test_pdns_glb.zone_id
	  }
	  `, riname, zonename, poolname, lbname)

}
