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

func TestAccIBMPrivateDNSZonesDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_zones.test1"
	riname := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSZonesDataSourceConfig(riname, zonename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "dns_zones.0.name"),
					resource.TestCheckResourceAttrSet(node, "dns_zones.0.description"),
					resource.TestCheckResourceAttrSet(node, "dns_zones.0.label"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSZonesDataSourceConfig(riname, zonename string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}

	resource "ibm_resource_instance" "test-pdns-instance" {
		name = "%s"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}

	resource "ibm_dns_zone" "test-pdns-zone" {
		name        = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription6"
		label       = "testlabel-updated6"
	  }

    data "ibm_dns_zones" "test1" {
		instance_id = ibm_dns_zone.test-pdns-zone.instance_id
	}`, riname, zonename)
}
