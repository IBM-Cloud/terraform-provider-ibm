// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIAMRoleDataSourceAction_basic(t *testing.T) {
	serviceName := "kms"
	kmsManagerAction := "kms.instancepolicies.read"
	countActions := "1"
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMRoleActionConfig(name, displayName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_role_actions.test", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "actions.#", countActions),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.customrole", "actions.0", kmsManagerAction),
				),
			},
		},
	})
}

func TestAccIBMIAMRoleDataSourceAction_withServiceSpecificRoleActions(t *testing.T) {
	serviceName := "cloud-object-storage"
	countActionsContentReaderAndObjectWriter := "22"
	contentReaderActionInCos := "cloud-object-storage.bucket.get"
	name := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	displayName := fmt.Sprintf("Terraform%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMCustomServiceRoleActionsConfig(name, displayName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_role_actions.example", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_iam_custom_role.read_write", "service", serviceName),
					resource.TestCheckResourceAttr(
						"ibm_iam_custom_role.read_write", "actions.#", countActionsContentReaderAndObjectWriter),
					resource.TestCheckResourceAttr(
						"ibm_iam_custom_role.read_write", "actions.0", contentReaderActionInCos),
				),
			},
		},
	})
}

func testAccCheckIBMIAMRoleActionConfig(name, displayName, serviceName string) string {
	return fmt.Sprintf(`

data "ibm_iam_role_actions" "test" {
  service = "%s"
}

resource "ibm_iam_custom_role" "customrole" {
    name         = "%s"
    display_name = "%s"
    description  = "Custom Role for test scenario2"
    service = "kms"
    actions      = [data.ibm_iam_role_actions.test.manager.18]
}
`, serviceName, name, displayName)
}

func testAccCheckIBMIAMCustomServiceRoleActionsConfig(name, displayName, serviceName string) string {
	return fmt.Sprintf(`

data "ibm_iam_role_actions" "example" {
  service = "%s"
}

resource "ibm_iam_custom_role" "read_write" {
  name = "%s"
  display_name = "%s"
  service = "%s"
  actions = concat(
             split(",", data.ibm_iam_role_actions.example.actions["Content Reader"]),
             split(",", data.ibm_iam_role_actions.example.actions["Object Writer"])
  )
}
`, serviceName, name, displayName, serviceName)
}
