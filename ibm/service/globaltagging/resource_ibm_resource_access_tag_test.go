// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package globaltagging_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const (
	accessTagRegex = "^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$"
)

func TestAccResourceAccessTag_Basic(t *testing.T) {
	name := fmt.Sprintf("tf%d:access%d", acctest.RandIntRange(10, 100), acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckResourceAccessTagCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAccessTagExists("ibm_resource_access_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "id", name),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "name", name),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "tag_type", "access"),
				),
			},
		},
	})
}
func TestAccResourceAccessTag_Usage(t *testing.T) {
	name := fmt.Sprintf("tf%d:access%d", acctest.RandIntRange(10, 100), acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshkeyname := fmt.Sprintf("tfssh-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			resource.TestStep{
				Config: testAccCheckResourceAccessTagCreate(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAccessTagExists("ibm_resource_access_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "id", name),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "name", name),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "tag_type", "access"),
				),
			},
			resource.TestStep{
				Config: testAccCheckResourceAccessTagUsage(name, sshkeyname, publicKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAccessTagExists("ibm_resource_access_tag.tag"),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "id", name),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "name", name),
					resource.TestCheckResourceAttr("ibm_resource_access_tag.tag", "tag_type", "access"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key.key", "name", sshkeyname),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key.key", "access_tags.#"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key.key", "access_tags.0", name),
				),
			},
		},
	})
}

func testAccCheckResourceAccessTagExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		var tagName string
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		accessTagRegex, err := regexp.Compile(accessTagRegex)
		if err != nil {
			return err
		}

		if accessTagRegex.MatchString(rs.Primary.ID) {
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

func testAccCheckResourceAccessTagCreate(name string) string {
	return fmt.Sprintf(`
	resource ibm_resource_access_tag tag {
		name = "%s"
	  }
`, name)
}
func testAccCheckResourceAccessTagUsage(name, sshkeyname, publicKey string) string {
	return fmt.Sprintf(`
	resource ibm_resource_access_tag tag {
		name = "%s"
	}
	resource "ibm_is_ssh_key" "key" {
		name = "%s"
		public_key = "%s"
		access_tags = [ibm_resource_access_tag.tag.name]
	}

`, name, sshkeyname, publicKey)
}
