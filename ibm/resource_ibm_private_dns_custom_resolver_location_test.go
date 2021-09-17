// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCRLocations_Basic(t *testing.T) {
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	subnet_crn := "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-202fe4570c83"
	subnet_crn_new := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test", "subnet_crn", subnet_crn_new),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test", "enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCRLocations_Import(t *testing.T) {
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR Locations - TF"
	subnet_crn := "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-202fe4570c83"
	subnet_crn_new := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(name, description, subnet_crn, subnet_crn_new),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test", "subnet_crn", subnet_crn_new),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_location.test", "enabled", "false"),
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
			instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
			description = "%s"
			enabled = true
		}
		resource "ibm_dns_custom_resolver_location" "test" {
			depends_on  = [ibm_dns_custom_resolver.test]
			instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
			resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
			subnet_crn = "%s"
			enabled    = true
		}
		resource "ibm_dns_custom_resolver_location" "test" {
			depends_on  = [ibm_dns_custom_resolver.test]
			instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
			resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
			subnet_crn = "%s"
			enabled    = true
		}
	  	`, name, description, subnet_crn, subnet_crn_new)
}
