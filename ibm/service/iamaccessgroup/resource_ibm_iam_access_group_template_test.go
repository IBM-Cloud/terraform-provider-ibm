// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
)

var (
	accountID string = acc.IAMAccountId
	agName    string = fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
)

func TestAccIBMIamAccessGroupTemplateBasic(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamAccessGroupTemplateVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamAccessGroupTemplateConfigBasic(name, accountID, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
				),
			},
		},
	})
}

func TestAccIBMIamAccessGroupTemplateBasicWithCommit(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamAccessGroupTemplateVersionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamAccessGroupTemplateConfigBasic(name, accountID, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
				),
			},
			{
				Config: testAccCheckIBMIamAccessGroupTemplateConfigBasicWithCommit(name, accountID, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "committed", "true"),
				),
			},
		},
	})
}

func TestAccIBMIamAccessGroupTemplateBasicWithAssertionAndActionControl(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamAccessGroupTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamAccessGroupTemplateConfig(name, description, accountID, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "account_id", accountID),
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

func testAccCheckIBMIamAccessGroupTemplateConfigBasic(name string, accountID string, agName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
			name = "%s"
			account_id = "%s"
			group {
				name = "%s"
			}
		}
	`, name, accountID, agName)
}

func testAccCheckIBMIamAccessGroupTemplateConfigBasicWithCommit(name string, accountID string, agName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_access_group_template" "template" {
			name = "%s"
			account_id = "%s"
			group {
				name = "%s"
			}
			committed = true
		}
	`, name, accountID, agName)
}

func testAccCheckIBMIamAccessGroupTemplateConfig(name string, description string, accountID string, agName string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_access_group_template" "template" {
			name = "%s"
			description = "%s"
			account_id = "%s"
			group {
				name = "%s"
				description = "description"
				assertions {
					action_controls {
						add = true
						remove = true
						update = true
					}
				}
				action_controls {
					access {
						add = true
					}
				}
			}
		}
	`, name, description, accountID, agName)
}

func testAccCheckIBMIamAccessGroupTemplateDestroy(s *terraform.State) error {
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
