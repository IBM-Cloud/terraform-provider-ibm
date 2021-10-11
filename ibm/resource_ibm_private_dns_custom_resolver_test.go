// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPrivateDNSCustomResolver_Basic(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSCustomResolverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverBasic(name, description),
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
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSCustomResolverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverBasic(name, description),
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

func testAccCheckIBMPrivateDNSCustomResolverBasic(name, description string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "rg" {
		is_default=true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name = "test-pdns-custom-resolver-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet" {
		name                    = "test-pdns-cr-subnet"
		vpc                     = ibm_is_vpc.test-pdns-cr-vpc.id
		zone            		= "us-south-1"
		ipv4_cidr_block 		= "10.240.0.0/24"
		resource_group 			= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name = "test-pdns-cr-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name        = "%s"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "%s"
		enabled = true
	}
	resource "ibm_dns_custom_resolver_location" "test" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet.crn
		enabled     = true
	}
	  `, name, description)
}

func testAccCheckIBMPrivateDNSCustomResolverDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_custom_resolver" {
			continue
		}
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDNSClientSession()
		if err != nil {
			return err
		}
		parts := rs.Primary.ID
		partslist := strings.Split(parts, ":")
		customResolverID := partslist[0]
		crn := partslist[1]

		getCustomResolverOptions := pdnsClient.NewDeleteCustomResolverOptions(crn, customResolverID)
		res, err := pdnsClient.DeleteCustomResolver(getCustomResolverOptions)
		if err != nil {
			if res != nil && res.StatusCode == 404 {
				return nil
			}
			return fmt.Errorf("testAccCheckIBMPrivateDNSCustomResolverDestroy: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMPrivateDNSCustomResolverExists(n string, result string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDNSClientSession()
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
