/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPrivateDNSGlbPoolsDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_glb_pools.test1"
	riname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	poolname := fmt.Sprintf("tf-poolname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
		name = "default"
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
