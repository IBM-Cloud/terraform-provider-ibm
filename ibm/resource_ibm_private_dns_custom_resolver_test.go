// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
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
	subnet_crn := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-6c3a997d-72b2-47f6-8788-6bd95e1bdb03"
	log.Println("In side TestAccIBMPrivateDNSCustomResolver_Basic")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSCustomResolverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverBasic(name, description, subnet_crn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSCustomResolverExists("ibm_dns_custom_resolver.test", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "name", name),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "description", description),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "locations.0.subnet_crn", subnet_crn),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCustomResolverImport(t *testing.T) {
	var resultprivatedns string
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"
	subnet_crn := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-6c3a997d-72b2-47f6-8788-6bd95e1bdb03"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPrivateDNSCustomResolverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverBasic(name, description, subnet_crn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPrivateDNSCustomResolverExists("ibm_dns_custom_resolver.test", resultprivatedns),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "name", name),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "description", description),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver.test", "locations.0.subnet_crn", subnet_crn),
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

func testAccCheckIBMPrivateDNSCustomResolverBasic(name, description, subnet_crn string) string {
	return fmt.Sprintf(`

	resource "ibm_dns_custom_resolver" "test" {
		name        = "%s"
		instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
		description = "%s"
		locations {
			subnet_crn = "%s"
			enabled    = "true"
		}
	}
	  `, name, description, subnet_crn)
}

func testAccCheckIBMPrivateDNSCustomResolverDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_custom_resolver" {
			continue
		}
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDNSClientSessionScoped()
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
		pdnsClient, err := testAccProvider.Meta().(ClientSession).PrivateDNSClientSessionScoped()
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
