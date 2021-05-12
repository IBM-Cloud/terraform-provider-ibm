// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceTag_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-satellitelocation-%d", acctest.RandIntRange(10, 100))
	managed_from := "wdc04"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckResourceTagCreate(name, managed_from),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceTagExists("ibm_resource_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.#", "2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_resource_tag.tag",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckResourceTagExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var resourceID string
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		crn, err := regexp.Compile(crnRegex)
		if err != nil {
			return err
		}

		if crn.MatchString(rs.Primary.ID) {
			resourceID = rs.Primary.ID
		} else {
			parts, err := vmIdParts(rs.Primary.ID)
			if err != nil {
				return err
			}
			resourceID = parts[0]
		}
		_, err = GetGlobalTagsUsingCRN(testAccProvider.Meta(), resourceID, "", "")
		if err != nil {
			log.Printf(
				"Error on get of resource tags (%s) : %s", resourceID, err)
		}
		return nil
	}
}

func testAccCheckResourceTagCreate(name, managed_from string) string {
	return fmt.Sprintf(`

	resource "ibm_satellite_location" "location" {
		location      = "%s"
		managed_from  = "%s"
		description	  = "satellite service"	
		zones		  = ["us-east-1", "us-east-2", "us-east-3"]
	}

	data "ibm_satellite_location" "test_location" {
		location  = ibm_satellite_location.location.location
	}

	resource "ibm_resource_tag" "tag" {
		resource_id = data.ibm_satellite_location.test_location.crn
		tags        = ["env:dev", "cpu:4"]
	}
`, name, managed_from)
}
