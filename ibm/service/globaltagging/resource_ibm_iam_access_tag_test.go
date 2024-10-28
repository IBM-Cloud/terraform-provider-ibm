// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package globaltagging_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	iamAccessTagRegex = "^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$"
)

func TestAccIamAccessTag_Basic(t *testing.T) {
	name := fmt.Sprintf("tf%d:iam-access%d", acctest.RandIntRange(10, 100), acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckIamAccessTagCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIamAccessTagExists("ibm_iam_access_tag.tag"),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "id", name),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "tag_type", "access"),
				),
			},
		},
	})
}
func TestAccIamAccessTag_Usage(t *testing.T) {
	name := fmt.Sprintf("tf%d:iam-access%d", acctest.RandIntRange(10, 100), acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckIamAccessTagCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIamAccessTagExists("ibm_iam_access_tag.tag"),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "id", name),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "tag_type", "access"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIamAccessTagUsage(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIamAccessTagExists("ibm_iam_access_tag.tag"),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "id", name),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_tag.tag", "tag_type", "access"),
					testAccCheckResourceTagExists("ibm_resource_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tags.0", name),
					resource.TestCheckResourceAttr("ibm_resource_tag.tag", "tag_type", "access"),
				),
			},
		},
	})
}

func testAccCheckIamAccessTagExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var tagName string
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamAccessTagRegex, err := regexp.Compile(iamAccessTagRegex)
		if err != nil {
			return err
		}

		if iamAccessTagRegex.MatchString(rs.Primary.ID) {
			tagName = rs.Primary.ID
		}

		gtClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).GlobalTaggingAPIv1()
		if err != nil {
			return fmt.Errorf("Error getting global tagging client settings: %s", err)
		}
		accessTagType := "access"
		listTagsOptions := &globaltaggingv1.ListTagsOptions{
			TagType: &accessTagType,
		}
		taggingResult, _, err := gtClient.ListTags(listTagsOptions)
		if err != nil {
			return err
		}

		var taglist []string
		for _, item := range taggingResult.Items {
			taglist = append(taglist, *item.Name)
		}
		existingAccessTags := flex.NewStringSet(flex.ResourceIBMVPCHash, taglist)
		if !existingAccessTags.Contains(tagName) {
			return fmt.Errorf(
				"Error on get of resource tags (%s) : %s", tagName, err)
		}
		return nil
	}
}

func testAccCheckIamAccessTagCreate(name string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_access_tag" "tag" {
		name = "%s"
	  }
`, name)
}
func testAccCheckIamAccessTagUsage(name string) string {
	resource_group_name := fmt.Sprintf("tf%d-iam-access%d", acctest.RandIntRange(10, 100), acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		name = "Default"
	}
	resource "ibm_iam_access_tag" "tag" {
		name = "%s"
	}
	resource "ibm_cd_toolchain" "cd_toolchain_instance" {
		description = "Terraform test"
		name = "%s-toolchain"
		resource_group_id = data.ibm_resource_group.group.id
	}
	resource "ibm_resource_tag" "tag" {
		resource_id = ibm_cd_toolchain.cd_toolchain_instance.crn
		tags        = [ibm_iam_access_tag.tag.name]
		tag_type	= "access"
	}
`, name, resource_group_name)
}
