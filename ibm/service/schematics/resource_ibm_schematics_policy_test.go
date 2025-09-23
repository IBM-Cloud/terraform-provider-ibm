// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmSchematicsPolicyBasic(t *testing.T) {
	var conf schematicsv1.Policy
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	kind := "agent_assignment_policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSchematicsPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsPolicyConfigBasic(name, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSchematicsPolicyExists("ibm_schematics_policy.schematics_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "kind", kind),
				),
			},
		},
	})
}

func TestAccIbmSchematicsPolicyAllArgs(t *testing.T) {
	var conf schematicsv1.Policy
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resourceGroup := "Default"
	location := "us-south"
	kind := "agent_assignment_policy"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	kindUpdate := "agent_assignment_policy"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSchematicsPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsPolicyConfig(name, description, resourceGroup, location, kind),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSchematicsPolicyExists("ibm_schematics_policy.schematics_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "kind", kind),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSchematicsPolicyConfig(nameUpdate, descriptionUpdate, resourceGroup, location, kindUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_policy.schematics_policy_instance", "kind", kindUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_policy.schematics_policy_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSchematicsPolicyConfigBasic(name string, kind string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_policy" "schematics_policy_instance" {
			name = "%s"
			kind = "%s"
		}
	`, name, kind)
}

func testAccCheckIbmSchematicsPolicyConfig(name string, description string, resourceGroup string, location string, kind string) string {
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
	`, name, description, resourceGroup, location, kind)
}

func testAccCheckIbmSchematicsPolicyExists(n string, obj schematicsv1.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getPolicyOptions := &schematicsv1.GetPolicyOptions{}

		getPolicyOptions.SetPolicyID(rs.Primary.ID)

		policy, _, err := schematicsClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return err
		}

		obj = *policy
		return nil
	}
}

func testAccCheckIbmSchematicsPolicyDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_policy" {
			continue
		}

		getPolicyOptions := &schematicsv1.GetPolicyOptions{}

		getPolicyOptions.SetPolicyID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetPolicy(getPolicyOptions)

		if err == nil {
			return fmt.Errorf("schematics_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
