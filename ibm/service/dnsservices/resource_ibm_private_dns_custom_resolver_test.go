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

func TestAccIBMPrivateDNSCustomResolver_basic(t *testing.T) {
	var resultprivatedns string
	vpcname := fmt.Sprintf("cr-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("cr-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSCustomResolverExists("ibm_dns_custom_resolver.test", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "name", name),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "description", description),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCustomResolverImport(t *testing.T) {
	var resultprivatedns string
	vpcname := fmt.Sprintf("cr-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("cr-subnet-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSCustomResolverExists("ibm_dns_custom_resolver.test", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "name", name),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "description", description),
				),
			},
			{
				ResourceName:      "ibm_dns_custom_resolver.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"type"},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCustomResolverBasic(vpcname, subnetname, zone, cidr, name, description string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "%s"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "%s"
		ipv4_cidr_block	= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name		= "%s"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "%s"
		high_availability = false
		enabled 	= true
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled		= true
		}
	}
	`, vpcname, subnetname, zone, cidr, name, description)
}

func testAccCheckIBMPrivateDNSCustomResolverExists(n string, result string) resource.TestCheckFunc {

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
		partslist := strings.Split(parts, ":")
		customResolverID := partslist[0]
		crn := partslist[1]

		getCustomResolverOptions := pdnsClient.NewGetCustomResolverOptions(crn, customResolverID)
		r, res, err := pdnsClient.GetCustomResolver(getCustomResolverOptions)

		if err != nil {
			if res != nil && res.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("testAccCheckIBMPrivateDNSCustomResolverExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}
