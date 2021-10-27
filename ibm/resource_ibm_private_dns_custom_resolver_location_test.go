// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverLocations_basic(t *testing.T) {
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	subnet_crn := "crn:v1:bluemix:public:is:us-south-3:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0737-0d198509-3221-4162-b2d8-4a9326d3d7ad"
	subnet_crn_new := "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-f17967f2-2bbe-427c-bcf6-22f8c2395285"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "subnet_crn", subnet_crn),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "enabled", "true"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "subnet_crn", subnet_crn_new),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCustomResolverLocations_Import(t *testing.T) {

	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	subnet_crn := "crn:v1:bluemix:public:is:us-south-3:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0737-0d198509-3221-4162-b2d8-4a9326d3d7ad"
	subnet_crn_new := "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-f17967f2-2bbe-427c-bcf6-22f8c2395285"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "subnet_crn", subnet_crn),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "enabled", "true"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test1", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "cr_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test2", "subnet_crn", subnet_crn_new),
				),
			},
			{
				ResourceName:      "ibm_dns_custom_resolver_location.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"enabled",
					"instance_id",
					"resolver_id",
					"subnet_crn",
				},
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new string) string {
	return fmt.Sprintf(`
			resource "ibm_dns_custom_resolver" "test" {
				name        = "%s"
				instance_id = "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
				description = "%s"
				high_availability = false
				enabled = false
			}
			resource "ibm_dns_custom_resolver_location" "test1" {
				instance_id = "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
				resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
				subnet_crn = "%s"
				enabled    = true
				cr_enabled = false
			}
			resource "ibm_dns_custom_resolver_location" "test2" {
				instance_id   = "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
				resolver_id   = ibm_dns_custom_resolver.test.custom_resolver_id
				subnet_crn    = "%s"
				enabled       = false
				cr_enabled    = false 
			  }
			  `, name, description, subnet_crn, subnet_crn_new)
}
