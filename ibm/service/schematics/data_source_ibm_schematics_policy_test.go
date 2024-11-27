// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsPolicyDataSourceBasic(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyKind := "agent_assignment_policy"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsPolicyDataSourceConfigBasic(policyName, policyKind),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "policy_id"),
				),
			},
		},
	})
}

func TestAccIbmSchematicsPolicyDataSourceAllArgs(t *testing.T) {
	policyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	policyResourceGroup := "Default"
	policyLocation := "us-south"
	policyKind := "agent_assignment_policy"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsPolicyDataSourceConfig(policyName, policyDescription, policyResourceGroup, policyLocation, policyKind),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "kind"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "target.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "parameter.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_policy.schematics_policy_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsPolicyDataSourceConfigBasic(policyName, policyKind string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_policy" "schematics_policy_instance" {
			name = "%s"
			kind = "%s"
		}

		data "ibm_schematics_policy" "schematics_policy_instance" {
			policy_id = ibm_schematics_policy.schematics_policy_instance.id
		}
	`, policyName, policyKind)
}

func testAccCheckIbmSchematicsPolicyDataSourceConfig(policyName string, policyDescription string, policyResourceGroup string, policyLocation string, policyKind string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_policy" "schematics_policy_instance" {
			name = "%s"
			description = "%s"
			resource_group = "%s"
			tags = ["policy-tag"]
			location = "%s"
			kind = "%s"
			target {
				selector_kind = "ids"
				selector_ids = [ "selector_ids" ]
			}
			parameter {
				agent_assignment_policy_parameter {
					selector_kind = "scoped"
					selector_scope {
						kind = "workspace"
						tags = [ "tags" ]
						resource_groups = [ "Default" ]
						locations = [ "us-south" ]
					}
				}
			}
		}

		data "ibm_schematics_policy" "schematics_policy_instance" {
			policy_id = ibm_schematics_policy.schematics_policy_instance.id
		}
	`, policyName, policyDescription, policyResourceGroup, policyLocation, policyKind)
}
