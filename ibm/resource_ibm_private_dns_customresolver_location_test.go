// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCRLocations_Basic(t *testing.T) {
	subnet_crn := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-03d54d71-b438-4d20-b943-76d3d2a1a590"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(subnet_crn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_cr_locations.test", "subnet_crn", subnet_crn),
					resource.TestCheckResourceAttr("ibm_dns_cr_locations.test", "enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMPrivateDNSCRLocations_Import(t *testing.T) {
	subnet_crn := "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-03d54d71-b438-4d20-b943-76d3d2a1a590"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCRLocationsBasic(subnet_crn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_cr_locations.test", "subnet_crn", subnet_crn),
					resource.TestCheckResourceAttr("ibm_dns_cr_locations.test", "enabled", "false"),
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

func testAccCheckIBMPrivateDNSCRLocationsBasic(subnet_crn string) string {
	return fmt.Sprintf(`
	resource "ibm_dns_cr_locations" "test" {
		instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
		resolver_id = "095604bd-265a-4cad-9e23-fc3d2fb7f5dc"
		subnet_crn = "%s"
		enabled    = false
	  }
	  `, subnet_crn)
}
