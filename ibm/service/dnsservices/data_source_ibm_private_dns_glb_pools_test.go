// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPrivateDNSGlbPoolsDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_glb_pools.test1"
	riname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	poolname := fmt.Sprintf("tf-poolname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbPoolsDataSourceConfig(vpcname, riname, zonename, poolname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "dns_glb_pools.0.name"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_pools.0.description"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_pools.0.enabled"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_pools.0.healthy_origins_threshold"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_pools.0.created_on"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSGlbPoolsDataSourceConfig(vpcname, riname, zonename, poolname string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}
    resource "ibm_is_vpc" "test-pdns-vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "%s"
		resource_group = data.ibm_resource_group.rg.id
    }
    resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test-pdns-vpc]
		name = "%s"
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
	resource "ibm_dns_glb_pool" "test-pdns-pool" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "new test pool"
		enabled=true
		healthy_origins_threshold=1
		origins {
				name    = "example-1"
				address = "www.google.com"
				enabled = true
				description="origin pool"
		}
    }

	data "ibm_dns_glb_pools" "test1" {
		instance_id = ibm_dns_glb_pool.test-pdns-pool.instance_id
	}`, vpcname, riname, zonename, poolname)

}
