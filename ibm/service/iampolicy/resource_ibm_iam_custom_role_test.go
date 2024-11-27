// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMIAMCustomRole_Basic(t *testing.T) {
	var conf iampolicymanagementv1.CustomRole
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	updateDisplayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMCustomRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMCustomRoleBasic(name, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMCustomRoleExists("ibm_iam_custom_role.customrole", conf),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "display_name", displayName),
				),
			},
			{
				Config: testAccCheckIBMIAMCustomRoleUpdateWithSameName(name, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMCustomRoleExists("ibm_iam_custom_role.customrole", conf),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "description", "role for test scenario1"),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "display_name", displayName),
				),
			},
			{
				Config: testAccCheckIBMIAMCustomRoleUpdate(name, updateDisplayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "display_name", updateDisplayName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "description", "role for test scenario2"),
				),
			},
		},
	})
}
func testAccCheckIBMIAMAccessGroupDestroy(s *terraform.State) error {
	accClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_access_group" {
			continue
		}

		agID := rs.Primary.ID

		// Try to find the key
		getAccessGroupOptions := &iamaccessgroupsv2.GetAccessGroupOptions{
			AccessGroupID: &agID,
		}
		_, detailResponse, _ := accClient.GetAccessGroup(getAccessGroupOptions)

		if err == nil {
			return fmt.Errorf("Access group still exists: %s", rs.Primary.ID)
		} else if detailResponse.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error waiting for access group (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
func TestAccIBMIAMCustomRole_import(t *testing.T) {
	var conf iampolicymanagementv1.CustomRole
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_iam_custom_role.customrole"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIAMAccessGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMCustomRoleMultipleAction(name, displayName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIAMCustomRoleExists(resourceName, conf),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "description", "Custom Role for test scenario2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIAMCustomRoleDestroy(s *terraform.State) error {
	roleClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_custom_role" {
			continue
		}

		roleID := rs.Primary.ID

		getRoleOptions := &iampolicymanagementv1.GetRoleOptions{
			RoleID: &roleID,
		}

		// Try to find the role
		_, response, err := roleClient.GetRole(getRoleOptions)
		if err == nil {
			return fmt.Errorf("Custom Role still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error waiting for Custom Role (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMIAMCustomRoleExists(n string, obj iampolicymanagementv1.CustomRole) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		roleClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return err
		}

		getRoleOptions := &iampolicymanagementv1.GetRoleOptions{
			RoleID: &rs.Primary.ID,
		}

		customrole, _, err := roleClient.GetRole(getRoleOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error retrieving Custom Role %s err: %s", rs.Primary.ID, err)
		}

		obj = *customrole
		return nil
	}
}

func testAccCheckIBMIAMCustomRoleBasic(name, displayName string) string {
	return fmt.Sprintf(`
		
	resource "ibm_iam_custom_role" "customrole" {
		name         = "%s"
		display_name = "%s"
		description  = "role for test scenario1"
		service = "kms"
		actions      = ["kms.secrets.rotate"]
	  }
	`, name, displayName)
}

func testAccCheckIBMIAMCustomRoleUpdateWithSameName(name, displayName string) string {
	return fmt.Sprintf(`
		
	resource "ibm_iam_custom_role" "customrole" {
		name         = "%s"
		display_name = "%s"
		description  = "role for test scenario1"
		service = "kms"
		actions      = ["kms.secrets.rotate"]
	  }
	`, name, displayName)
}

func testAccCheckIBMIAMCustomRoleUpdate(name, updateName string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_custom_role" "customrole" {
		name         = "%s"
		display_name = "%s"
		description  = "role for test scenario2"
		service = "kms"
		actions      = ["kms.secrets.rotate"]
	  }
	`, name, updateName)
}

func testAccCheckIBMIAMCustomRoleMultipleAction(name, displayName string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_custom_role" "customrole" {
		name         = "%s"
		display_name = "%s"
		description  = "Custom Role for test scenario2"
		service = "kms"
		actions      = ["kms.registrations.merge","kms.registrations.write"]
	  }
	`, name, displayName)
}
