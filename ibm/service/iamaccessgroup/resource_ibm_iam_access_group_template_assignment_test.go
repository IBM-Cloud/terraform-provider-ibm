// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
)

var target string = os.Getenv("IAM_ACCESS_GROUP_ASSIGNMENT_TARGET_ACCOUNT")

func TestAccIBMIAMAccessGroupTemplateAssignmentBasic(t *testing.T) {
	var conf iamaccessgroupsv2.TemplateAssignmentVerboseResponse
	var versionConf iamaccessgroupsv2.TemplateVersionResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupTemplateAssignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateAssignmentConfigBasic(name, agName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateVersionExists("ibm_iam_access_group_template.template", versionConf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template.template", "group.0.name", agName),
				),
			},
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateAssignmentConfigUpdateCommit(name, agName, target),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMAccessGroupTemplateAssignmentExists("ibm_iam_access_group_template_assignment.assignment", conf),
					resource.TestCheckResourceAttr("ibm_iam_access_group_template_assignment.assignment", "target", target),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentConfigBasic(name string, agName string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		group {
			name = "%s"
		}
	}
	`, name, agName)
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentVersionUpdate(name string, agName string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		group {
			name = "%s"
		}
	}

	resource "ibm_iam_access_group_template_version" "template" {
		template_id = ibm_iam_access_group_template.template.template_id
		name = ibm_iam_access_group_template.template.name
		description = "Test Terraform Description"
		  group {
			  name = "Test Terraform Access Group Name"
			  assertions {
				  action_controls {
					  add    = false
					  remove = true
					  update = true
				  }
			  }
		  }
	  }

	`, name, agName)
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentConfigVersionUpdateCommit(name string, agName string, target string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		group {
			name = "%s"
		}
		committed = true
	}

	resource "ibm_iam_access_group_template_version" "template" {
		template_id = ibm_iam_access_group_template.template.template_id
		name = ibm_iam_access_group_template.template.name
		description = "Test Terraform Description"
		  group {
			  name = "Test Terraform Access Group Name"
			  assertions {
				  action_controls {
					  add    = true
					  remove = true
					  update = true
				  }
			  }
		  }
		  committed = true
	  }

	resource ibm_iam_access_group_template_assignment "assignment" {
		template_id = ibm_iam_access_group_template_version.template.template_id
		template_version = ibm_iam_access_group_template_version.template.version
		target_type = "AccountGroup"
		target = "%s"
	}
	`, name, agName, target)
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentUpdate(agName string, target string, name string) string {
	return fmt.Sprintf(`
	resource ibm_iam_access_group_template_assignment "assignment" {
		template_id = ibm_iam_access_group_template.template.template_id
		template_version = ibm_iam_access_group_template.template.version
		target_type = "AccountGroup"
		target = "%s"
	}
	
	resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		group {
			name = "%s"
		}
		committed = true
	}

	resource "ibm_iam_access_group_template_version" "template" {
		template_id = ibm_iam_access_group_template.template.template_id
		name = ibm_iam_access_group_template.template.name
		description = "Test Terraform Description"
		  group {
			  name = "Test Terraform Access Group Name"
			  assertions {
				  action_controls {
					  add    = true
					  remove = true
					  update = true
				  }
			  }
		  }
		  committed = true
	  }
	`, target, name, agName)
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentConfigUpdateCommit(name string, agName string, target string) string {
	return fmt.Sprintf(`
	resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		group {
			name = "%s"
		}
		committed = true
	}

	resource ibm_iam_access_group_template_assignment "assignment" {
		template_id = ibm_iam_access_group_template.template.template_id
		template_version = ibm_iam_access_group_template.template.version
		target_type = "AccountGroup"
		target = "%s"
	}
	`, name, agName, target)
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentExists(n string, obj iamaccessgroupsv2.TemplateAssignmentVerboseResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		iamAccessGroupsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
		if err != nil {
			return err
		}

		getAssignmentOptions := &iamaccessgroupsv2.GetAssignmentOptions{}

		getAssignmentOptions.SetAssignmentID(rs.Primary.ID)

		templateAssignmentResponse, _, err := iamAccessGroupsClient.GetAssignment(getAssignmentOptions)
		if err != nil {
			return err
		}

		obj = *templateAssignmentResponse
		return nil
	}
}

func testAccCheckIBMIAMAccessGroupTemplateAssignmentDestroy(s *terraform.State) error {
	iamAccessGroupsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group_template_assignment" {
			continue
		}

		getAssignmentOptions := &iamaccessgroupsv2.GetAssignmentOptions{}

		getAssignmentOptions.SetAssignmentID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamAccessGroupsClient.GetAssignment(getAssignmentOptions)

		if err == nil {
			return flex.FmtErrorf("iam_access_group_template_assignment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("Error checking for iam_access_group_template_assignment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
