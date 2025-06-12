// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPrivateDNSResourceRecord_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnsresourcerecord%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSResourceRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSResourceRecordBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSResourceRecordExists("ibm_dns_resource_record.test-pdns-resource-record-txt", &resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_resource_record.test-pdns-resource-record-txt", "type", "TXT"),
				),
			},
			{
				Config: testAccCheckIBMPrivateDNSResourceRecordUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSResourceRecordExists("ibm_dns_resource_record.test-pdns-resource-record-txt", &resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_resource_record.test-pdns-resource-record-txt", "type", "TXT"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSResourceRecordImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSResourceRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSResourceRecordBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSResourceRecordExists("ibm_dns_resource_record.test-pdns-resource-record-txt", &resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_resource_record.test-pdns-resource-record-txt", "type", "TXT"),
				),
			},
			{
				ResourceName:      "ibm_dns_resource_record.test-pdns-resource-record-txt",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"preference",
					"priority",
					"rdata",
					"weight",
				},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSResourceRecordBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}

	resource "ibm_is_vpc" "test_pdns_vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "test-pdns-record-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}

	resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test_pdns_vpc]
		name = "test-pdns-record-instance"
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
		name = "testA"
		rdata = "1.2.3.4"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-aaaa" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-a]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "AAAA"
		name = "testAAAA"
		rdata = "2001:0db8:0012:0001:3c5e:7354:0000:5db5"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-cname" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-aaaa]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "CNAME"
		name = "testCNAME"
		rdata = "%s"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-ptr" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-cname]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "PTR"
		name = "1.2.3.4"
		rdata = "testA.%s"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-mx" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-ptr]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "MX"
		name = "testMX"
		rdata = "mailserver.%s"
		preference = 10
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-srv" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-mx]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "SRV"
		name = "testSRV"
		rdata = "tester.com"
		priority = 100
		weight = 100
		port = 8000
		service = "_sip"
		protocol = "udp"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-txt" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-srv]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "TXT"
		name = "testTXT"
		rdata = "textinformation"
	}
	  `, name, name, name, name)
}

func testAccCheckIBMPrivateDNSResourceRecordUpdate(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}

	resource "ibm_is_vpc" "test_pdns_vpc" {
		depends_on = [data.ibm_resource_group.rg]
		name = "test-pdns-record-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}

	resource "ibm_resource_instance" "test-pdns-instance" {
		depends_on = [ibm_is_vpc.test_pdns_vpc]
		name = "test-pdns-record-instance"
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
		name = "testAupdate"
		rdata = "1.2.3.5"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-aaaa" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-a]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "AAAA"
		name = "testAAAAUpdate"
		rdata = "2001:0db8:0012:0001:3c5e:7354:0000:5db5"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-cname" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-aaaa]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "CNAME"
		name = "testCNAMEUpdate"
		rdata = "%s"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-ptr" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-cname]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "PTR"
		name = "1.2.3.7"
		rdata = "testAupdate.%s"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-mx" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-ptr]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "MX"
		name = "testMXUpdate"
		rdata = "mailserverupdate.%s"
		preference = 10
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-srv" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-mx]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "SRV"
		name = "testSRVUpdate"
		rdata = "testerupdate.com"
		priority = 100
		weight = 100
		port = 8000
		service = "_sip"
		protocol = "udp"
	}

	resource "ibm_dns_resource_record" "test-pdns-resource-record-txt" {
		depends_on = [ibm_dns_resource_record.test-pdns-resource-record-srv]
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		zone_id = ibm_dns_zone.test-pdns-zone.zone_id
		type = "TXT"
		name = "testTXTUpdate"
		rdata = "textinformation"
	}
	  `, name, name, name, name)
}

func testAccCheckIBMPrivateDNSResourceRecordDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_resource_record" {
			continue
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")
		getResourceRecordOptions := pdnsClient.NewGetResourceRecordOptions(partslist[0], partslist[1], partslist[2])
		_, res, err := pdnsClient.GetResourceRecord(getResourceRecordOptions)
		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return flex.FmtErrorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSResourceRecordExists(n string, result *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")
		getResourceRecordOptions := pdnsClient.NewGetResourceRecordOptions(partslist[0], partslist[1], partslist[2])
		_, res, err := pdnsClient.GetResourceRecord(getResourceRecordOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return flex.FmtErrorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		*result = parts
		return nil
	}
}
