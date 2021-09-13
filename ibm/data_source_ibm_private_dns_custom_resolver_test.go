// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_custom_resolvers.test-cr"
	crname := fmt.Sprintf("tf-pdns-custom-resolver-%d", acctest.RandIntRange(100, 200))
	crdescription := fmt.Sprintf("tf-pdns-custom-resolver-tf-test%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverDataSourceConfig(crname, crdescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.name"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.description"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.enabled"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.locations.0.subnet_crn"),
					resource.TestCheckResourceAttrSet(node, "custom_resolvers.0.locations.0.enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCustomResolverDataSourceConfig(crname, crdescription string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`

	resource "ibm_dns_custom_resolver" "test" {
		name                      = "%s"
		instance_id               = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
		description               = "%s"
		locations {
		  subnet_crn  = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-6c3a997d-72b2-47f6-8788-6bd95e1bdb03"
		  enabled     = true
		}
	}

	data "ibm_dns_custom_resolvers" "test-cr" {
		instance_id = ibm_dns_custom_resolver.test.instance_id
	}`, crname, crdescription)
}
