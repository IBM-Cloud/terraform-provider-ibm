// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

var (
	agName string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
)

func TestAccIBMIAMAccessGroupTemplateBasic(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupTemplateVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateConfigBasic(name, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupTemplateBasicWithCommit(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupTemplateVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateConfigBasic(name, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateConfigBasicWithCommit(name, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIAMAccessGroupTemplateBasicWithAssertionAndActionControl(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateConfig(name, description, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
				),
			},
			{
				ResourceName:      "ibm_iam_access_group_template.template",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupTemplateConfigBasic(name string, agName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
			name = "%s"
			group {
				name = "%s"
			}
			committed = true
		}
	`, name, agName)
}

func testAccCheckIBMIAMAccessGroupTemplateConfigBasicWithCommit(name string, agName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
			name = "%s"
			group {
				name = "%s"
			}
			committed = true
		}
	`, name, agName)
}

func testAccCheckIBMIAMAccessGroupTemplateConfig(name string, description string, agName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group_template" "template" {
			name = "%s"
			description = "%s"
			group {
				name = "%s"
				description = "description"
				assertions {
					action_controls {
						add = true
						remove = true
					}
				}
				action_controls {
					access {
						add = true
					}
				}
			}
		}
	`, name, description, agName)
}

func testAccCheckIBMIAMAccessGroupTemplateDestroy(s *terraform.State) error {
	iamAccessGroupsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_template" {
			continue
		}

		getLatestTemplateVersionOptions := &iamaccessgroupsv2.GetLatestTemplateVersionOptions{}

		getLatestTemplateVersionOptions.SetTemplateID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamAccessGroupsClient.GetLatestTemplateVersion(getLatestTemplateVersionOptions)

		if err == nil {
			return fmt.Errorf("iam_access_group_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_access_group_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
