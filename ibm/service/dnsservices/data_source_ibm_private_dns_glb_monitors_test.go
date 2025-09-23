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

func TestAccIBMPrivateDNSGlbMonitorsDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_glb_monitors.test1"
	riname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	moniname := fmt.Sprintf("tf-monitorname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSGlbMonitordDataSConfig(vpcname, riname, zonename, moniname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "dns_glb_monitors.0.name"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_monitors.0.interval"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_monitors.0.type"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_monitors.0.retries"),
					resource.TestCheckResourceAttrSet(node, "dns_glb_monitors.0.timeout"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSGlbMonitordDataSConfig(vpcname, riname, zonename, moniname string) string {
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

	resource "ibm_dns_glb_monitor" "test-pdns-monitor" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "test monitor description"
		interval=63
		retries=3
		timeout=8
		port=8080
		type="HTTP"
		path="/"
		method="GET"

	}

	data "ibm_dns_glb_monitors" "test1" {
		instance_id = ibm_dns_glb_monitor.test-pdns-monitor.instance_id
	}`, vpcname, riname, zonename, moniname)

}
