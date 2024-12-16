// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package globaltagging_test

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceTag_Basic(t *testing.T) {
	name := fmt.Sprintf("tf-cos-%d", acctest.RandIntRange(10, 100))
	var tags []string
	tags = append(tags, "env:dev")
	tags = append(tags, "cpu:4")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckTagDestroy("ibm_resource_tag.tag", "user", tags),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckResourceTagCreate(name, tags),
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

func TestAccResourceTag_FakeCrnExpectingError(t *testing.T) {
	var tags []string
	tags = append(tags, "env:dev")
	tags = append(tags, "cpu:4")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckResourceTagCreateFakeCrnExpectingError(tags),
				ExpectError: regexp.MustCompile("\"is_error\": true"),
			},
		},
	})
}

func TestAccResourceTag_AttachOnExistingResource(t *testing.T) {
	crn := "crn:v1:bluemix:public:toolchain:eu-gb:a/970f5cb4bbc04119ab0a0f399e4b776c:8784b8c3-1c7f-476a-ac30-50ae07e3cce3::"
	var tags []string
	tags = append(tags, "env:dev")
	tags = append(tags, "cpu:4")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckTagOnExistingResourceDestroy(crn, "ibm_resource_tag.tag", "user", tags),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckResourceAttachOnExistingResource(crn, tags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceTagExists("ibm_resource_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.#", "2"),
				),
			},
		},
	})
}

func testAccCheckTagDestroy(n, tagType string, tagNames []string) resource.TestCheckFunc {
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
		results, errorGet := flex.GetTagsUsingResourceCRNFromTaggingApi(acc.TestAccProvider.Meta(), resourceID, "", tagType)
		if errorGet != nil {
			return fmt.Errorf("Error on get of resource tags (%s) : %s", resourceID, errorGet)
		}
		var taglist []string
		for _, v := range results.List() {
			taglist = append(taglist, fmt.Sprint(v))
		}
		existingAccessTags := flex.NewStringSet(flex.ResourceIBMVPCHash, taglist)
		for _, tagName := range tagNames {
			if existingAccessTags.Contains(tagName) {
				return fmt.Errorf("Tag still exists")
			}
		}
		return nil
	}
}

func testAccCheckTagOnExistingResourceDestroy(resourceCrn, n, tagType string, tagNames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var resourceID string
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		crnRegex := "^crn:v1(:[a-zA-Z0-9 \\-\\._~\\*\\+,;=!$&'\\(\\)\\/\\?#\\[\\]@]*){8}$|^[0-9]+$"
		crn, err := regexp.Compile(crnRegex)
		if err != nil {
			return err
		}
		if crn.MatchString(resourceCrn) {
			resourceID = resourceCrn
		} else {
			return fmt.Errorf("CRN not correct: %s", resourceCrn)
		}
		results, errorGet := flex.GetTagsUsingResourceCRNFromTaggingApi(acc.TestAccProvider.Meta(), resourceID, "", tagType)
		if errorGet != nil {
			return fmt.Errorf("Error on get of resource tags (%s) : %s", resourceID, errorGet)
		}
		var taglist []string
		for _, v := range results.List() {
			taglist = append(taglist, fmt.Sprint(v))
		}
		existingAccessTags := flex.NewStringSet(flex.ResourceIBMVPCHash, taglist)
		for _, tagName := range tagNames {
			if existingAccessTags.Contains(tagName) {
				return fmt.Errorf("Tag still exists")
			}
		}
		return nil
	}
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

func testAccCheckResourceTagCreate(name string, tagNames []string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "resource_1" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "lite"
		location          = "global"
	}

	resource "ibm_resource_tag" "tag" {
		resource_id = ibm_resource_instance.resource_1.crn
		tags        = ["%s"]
	}
`, name, strings.Join(tagNames[:], "\",\""))
}

func testAccCheckResourceTagCreateFakeCrnExpectingError(tagNames []string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_tag" "tag" {
		resource_id = "crn:v1:staging:public:cloud-object-storage:global:a/d99e99999dfe99ee999999f99bddd099:ab99d9be-9e9c-99dd-ad99-9bced9999999::"
		tags        = ["%s"]
	}
`, strings.Join(tagNames[:], "\",\""))
}

func testAccCheckResourceAttachOnExistingResource(crn string, tagNames []string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_tag" "tag" {
		resource_id = "%s"
		tags        = ["%s"]
	}
`, crn, strings.Join(tagNames[:], "\",\""))
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
