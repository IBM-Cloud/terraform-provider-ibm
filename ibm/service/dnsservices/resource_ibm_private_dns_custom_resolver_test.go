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
	name := fmt.Sprintf("testpdnscustomresolver%s", acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
	description := "new test CR - TF"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
		instance_id		= "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
		description		= "%s"
		high_availability =  false
		enabled		= true
		locations	{
			subnet_crn	= "crn:v1:bluemix:public:is:us-south-3:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0737-0d198509-3221-4162-b2d8-4a9326d3d7ad"
			enabled		= false
		}
		locations {
			subnet_crn  = "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-f17967f2-2bbe-427c-bcf6-22f8c2395285"
			enabled     = true
		}
	}
	  `, name, description)
}

func testAccCheckIBMPrivateDNSCustomResolverDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dns_custom_resolver" {
			continue
		}
		pdnsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PrivateDNSClientSession()
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
