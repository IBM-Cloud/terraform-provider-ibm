// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package globaltagging_test

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceTag_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckResourceTagCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceTagExists("ibm_resource_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.#", "2"),
				),
			},
			{
				ResourceName:      "ibm_resource_tag.tag",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"replace"},
			},
		},
	})
}
func TestAccResourceTag_Wait(t *testing.T) {
	name := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckResourceTagWaitCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceTagExists("ibm_resource_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.#", "3"),
				),
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
		crnRegex := "^crn:v1(:[a-zA-Z0-9 \\-\\._~\\*\\+,;=!$&'\\(\\)\\/\\?#\\[\\]@]*){8}$|^[0-9]+$"
		crn, err := regexp.Compile(crnRegex)
		if err != nil {
			return err
		}

		if crn.MatchString(rs.Primary.ID) {
			resourceID = rs.Primary.ID
		} else {
			parts, err := flex.VmIdParts(rs.Primary.ID)
			if err != nil {
				return err
			}
			resourceID = parts[0]
		}
		_, err = flex.GetGlobalTagsUsingCRN(acc.TestAccProvider.Meta(), resourceID, "", "")
		if err != nil {
			log.Printf(
				"Error on get of resource tags (%s) : %s", resourceID, err)
		}
		return nil
	}
}

func testAccCheckResourceTagWaitCreate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "resource_1" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "lite"
		location          = "global"
	}

	resource "ibm_resource_tag" "tag" {
		resource_id = ibm_resource_instance.resource_1.crn
		tags        = ["env:dev", "cpu:4", "user:8"]
	}
`, name)
}
func testAccCheckResourceTagCreate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "resource_1" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "lite"
		location          = "global"
	}

	resource "ibm_resource_tag" "tag" {
		resource_id = ibm_resource_instance.resource_1.crn
		tags        = ["env:dev", "cpu:4"]
	}
`, name)
}

func TestAccResourceTag_replace_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckResourceTagCreate_replace(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceTagExists("ibm_resource_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.#", "1"),
				),
			},
			{
				ResourceName:      "ibm_resource_tag.tag",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"replace"},
			},
		},
	})
}

func testAccCheckResourceTagCreate_replace(name string) string {
	return fmt.Sprintf(`
        resource "ibm_resource_instance" "resource_1" {
          name              = "%s"
          service           = "cloud-object-storage"
          plan              = "lite"
          location          = "global"
        }

        resource "ibm_resource_tag" "tag" {
            resource_id = resource.ibm_resource_instance.resource_1.crn
            tags        = ["test:test"]
            replace     = true
        }
    `, name)
}
