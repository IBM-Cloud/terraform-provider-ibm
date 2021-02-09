/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMPrivateDNSResourceRecordsDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_resource_records.test1"
	riname := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(100, 200))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	recname := fmt.Sprintf("tf-recname-%d.com", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSResourceRecordsDataSourceConfig(riname, zonename, vpcname, recname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "dns_resource_records.0.name"),
					resource.TestCheckResourceAttrSet(node, "dns_resource_records.0.rdata"),
					resource.TestCheckResourceAttrSet(node, "dns_resource_records.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSResourceRecordsDataSourceConfig(riname, zonename, vpcname, recname string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "default"
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
		description = "testdescription100"
		label       = "testlabel-updated100"
	  }

	resource "ibm_is_vpc" "test_pdns_vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "%s"
		resource_group = data.ibm_resource_group.rg.id
	}

	resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
		depends_on = [ibm_dns_zone.test-pdns-zone]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		vpc_crn = ibm_is_vpc.test_pdns_vpc.resource_crn
	}
	resource "ibm_dns_resource_record" "test-pdns-resource-record-a" {
		depends_on = [ibm_dns_permitted_network.test-pdns-permitted-network-nw]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "A"
		name = "%s"
		rdata = "5.6.7.8"
	}

    data "ibm_dns_resource_records" "test1" {
		instance_id = ibm_dns_zone.test-pdns-zone.instance_id
		zone_id = 	ibm_dns_resource_record.test-pdns-resource-record-a.zone_id
	}`, riname, zonename, vpcname, recname)
}
