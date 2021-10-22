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

func TestAccIBMPrivateDNSCustomResolver_basic(t *testing.T) {
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
	resource "ibm_dns_custom_resolver" "test" {
		name			= "%s"
		instance_id		= "fca34054-1c73-497d-8304-41bba9b03acb"
		description		= "%s"
		high_availability =  false
		enabled		= true
		locations	{
			subnet_crn	= "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
			enabled		= true
		}
		locations {
			subnet_crn  = "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-202fe4570c83"
			enabled     = false
		}
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
