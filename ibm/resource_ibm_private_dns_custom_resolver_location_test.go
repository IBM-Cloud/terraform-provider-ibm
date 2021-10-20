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
	subnet_crn := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
	subnet_crn_new := "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-202fe4570c83"
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
	subnet_crn := "crn:v1:bluemix:public:is:us-south-1:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0717-4f53a236-cd7a-4688-9347-066bb5058a5c"
	subnet_crn_new := "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-a248c17c-0872-417e-9dd5-e9927dc64b56"
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
				instance_id = "d515a480-a702-4837-9f40-6c0c285262fd"
				description = "%s"
				high_availability = false
				enabled = false
			}
			resource "ibm_dns_custom_resolver_location" "test1" {
				depends_on  = [ibm_dns_custom_resolver.test]
				instance_id = "d515a480-a702-4837-9f40-6c0c285262fd"
				resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
				subnet_crn = "%s"
				enabled    = true
				cr_enabled = false
			}
			resource "ibm_dns_custom_resolver_location" "test2" {
				depends_on  = [ibm_dns_custom_resolver.test]
				instance_id   = "d515a480-a702-4837-9f40-6c0c285262fd"
				resolver_id   = ibm_dns_custom_resolver.test.custom_resolver_id
				subnet_crn    = "%s"
				enabled       = false
				cr_enabled    = false 
			  }
			  `, name, description, subnet_crn, subnet_crn_new)
}
