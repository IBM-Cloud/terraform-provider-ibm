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

func TestAccIBMPrivateDNSZone_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSZoneBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSZoneExists("ibm_dns_zone.test-pdns-zone-zone", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_zone.test-pdns-zone-zone", "name", name),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSZoneImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnszone%s.com", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	//var resultendpoint apigatewaysdk.V2Endpoint
	//name := fmt.Sprintf("tftest-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSZoneBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSZoneExists("ibm_dns_zone.test-pdns-zone-zone", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_zone.test-pdns-zone-zone", "name", name),
				),
			},
			{
				ResourceName:      "ibm_dns_zone.test-pdns-zone-zone",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"type"},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSZoneBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}
	resource "ibm_resource_instance" "test-pdns-zone-instance" {
		name = "test-pdns-zone-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}
	resource "ibm_dns_zone" "test-pdns-zone-zone" {
		depends_on = ["ibm_resource_instance.test-pdns-zone-instance"]
		name = "%s"
		instance_id = ibm_resource_instance.test-pdns-zone-instance.guid
		description = "testdescription"
		label = "testlabel"
	}
	  `, name)
}

func testAccCheckIBMPrivateDNSZoneDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_zone" {
			continue
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")

		getZoneOptions := pdnsClient.NewGetDnszoneOptions(partslist[0], partslist[1])
		_, res, err := pdnsClient.GetDnszone(getZoneOptions)
		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {

			return flex.FmtErrorf("testAccCheckIBMPrivateDNSZoneDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSZoneExists(n string, result string) resource.TestCheckFunc {

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

		getZoneOptions := pdnsClient.NewGetDnszoneOptions(partslist[0], partslist[1])
		r, res, err := pdnsClient.GetDnszone(getZoneOptions)

		if err != nil &&
			res.StatusCode != 403 &&
			!strings.Contains(err.Error(), "The service instance was disabled, any access is not allowed.") {
			return flex.FmtErrorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}
