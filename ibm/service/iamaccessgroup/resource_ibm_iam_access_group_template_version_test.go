// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
)

var versionAGName string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))

func TestAccIBMIAMAccessGroupTemplateVersion(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupTemplateVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateVersionConfigBasic(name, agName, versionAGName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template_version.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "group.0.name", versionAGName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupTemplateVersionUpdateWithCommit(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupTemplateVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateVersionConfigBasicWithoutCommit(name, agName, versionAGName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template_version.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "group.0.name", versionAGName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateVersionUpdateWithCommit(name, agName, versionAGName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template_version.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "group.0.name", versionAGName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "committed", "true"),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_version.template", "description", "Testing3"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupTemplateVersionConfigBasic(name string, agName string, versionAGName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		description = "Test Terraform Description"
		group {
			name = "%s"
		}
	}
	
	resource "ibm_iam_access_group_template_version" "template" {
	  template_id = ibm_iam_access_group_template.template.template_id
	  name = ibm_iam_access_group_template.template.name
	  description = "Testing4"
		group {
			name = "%s"
			assertions {
				action_controls {
					add    = false
					remove = true
				}
			}
		}
		committed = true
	}
	`, name, agName, versionAGName)
}

func testAccCheckIBMIAMAccessGroupTemplateVersionConfigBasicWithoutCommit(name string, agName string, versionAGName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		description = "Test Terraform Description"
		group {
			name = "%s"
		}
	}
	
	resource "ibm_iam_access_group_template_version" "template" {
	  template_id = ibm_iam_access_group_template.template.template_id
	  name = ibm_iam_access_group_template.template.name
	  description = "Testing4"
		group {
			name = "%s"
			assertions {
				action_controls {
					add    = false
					remove = true
				}
			}
		}
	}
	`, name, agName, versionAGName)
}

func testAccCheckIBMIAMAccessGroupTemplateVersionUpdateWithCommit(name string, agName string, versionAGName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		description = "Test Terraform Description"
		group {
			name = "%s"
		}
	}
	
	resource "ibm_iam_access_group_template_version" "template" {
	  template_id = ibm_iam_access_group_template.template.template_id
	  name = ibm_iam_access_group_template.template.name
	  description = "Testing3"
		group {
			name = "%s"
			assertions {
				action_controls {
					add    = true
					remove = true
				}
			}
		}
		committed = true
	}
	`, name, agName, versionAGName)
}

func testAccCheckIBMIAMAccessGroupTemplateVersionExists(n string, obj iamaccessgroupsv2.TemplateVersionResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamAccessGroupsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
		if err != nil {
			return err
		}

		getTemplateVersionOptions := &iamaccessgroupsv2.GetTemplateVersionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTemplateVersionOptions.SetTemplateID(parts[0])
		getTemplateVersionOptions.SetVersionNum(parts[1])

		templateVersionResponse, _, err := iamAccessGroupsClient.GetTemplateVersion(getTemplateVersionOptions)
		if err != nil {
			return err
		}

		obj = *templateVersionResponse
		return nil
	}
}

func testAccCheckIBMIAMAccessGroupTemplateVersionDestroy(s *terraform.State) error {
	iamAccessGroupsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_template_version" {
			continue
		}

		getTemplateVersionOptions := &iamaccessgroupsv2.GetTemplateVersionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTemplateVersionOptions.SetTemplateID(parts[0])
		getTemplateVersionOptions.SetVersionNum(parts[1])

		// Try to find the key
		_, response, err := iamAccessGroupsClient.GetTemplateVersion(getTemplateVersionOptions)

		if err == nil {
			return fmt.Errorf("iam_access_group_template_version still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_access_group_template_version (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
