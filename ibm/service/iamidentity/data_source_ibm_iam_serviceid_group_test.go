// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamServiceidGroupDataSourceBasic(t *testing.T) {
	serviceIDGroupAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	serviceIDGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamServiceidGroupDataSourceConfigBasic(serviceIDGroupAccountID, serviceIDGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "iam_serviceid_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "created_by"),
				),
			},
		},
	})
}

func TestAccIBMIamServiceidGroupDataSourceAllArgs(t *testing.T) {
	serviceIDGroupAccountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	serviceIDGroupName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	serviceIDGroupDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamServiceidGroupDataSourceConfig(serviceIDGroupAccountID, serviceIDGroupName, serviceIDGroupDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "iam_serviceid_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_serviceid_group.iam_serviceid_group_instance", "modified_at"),
				),
			},
		},
	})
}

func testAccCheckIBMIamServiceidGroupDataSourceConfigBasic(serviceIDGroupAccountID string, serviceIDGroupName string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
			account_id = "%s"
			name = "%s"
		}

		data "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
			iam_serviceid_group_id = "iam_serviceid_group_id"
		}
	`, serviceIDGroupAccountID, serviceIDGroupName)
}

func testAccCheckIBMIamServiceidGroupDataSourceConfig(serviceIDGroupAccountID string, serviceIDGroupName string, serviceIDGroupDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
			account_id = "%s"
			name = "%s"
			description = "%s"
		}

		data "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
			iam_serviceid_group_id = "iam_serviceid_group_id"
		}
	`, serviceIDGroupAccountID, serviceIDGroupName, serviceIDGroupDescription)
}
