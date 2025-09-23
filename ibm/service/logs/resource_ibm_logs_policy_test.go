// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIbmLogsPolicyBasic(t *testing.T) {
	var conf logsv0.Policy
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	priority := "type_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	priorityUpdate := "type_high"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfigBasic(name, priority),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsPolicyExists("ibm_logs_policy.logs_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priority),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfigBasic(nameUpdate, priorityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priorityUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsPolicyAllArgs(t *testing.T) {
	var conf logsv0.Policy
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	priority := "type_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	priorityUpdate := "type_high"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfig(name, description, priority),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsPolicyExists("ibm_logs_policy.logs_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priority),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsPolicyConfig(nameUpdate, descriptionUpdate, priorityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_policy.logs_policy_instance", "priority", priorityUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_policy.logs_policy_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsPolicyConfigBasic(name string, priority string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_policy" "logs_policy_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "Test description"
		priority    = "%s"
		application_rule {
		  name         = "otel-links-test"
		  rule_type_id = "start_with"
		}
		log_rules {
		  severities = ["info"]
		}
	  }
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, priority)
}

func testAccCheckIbmLogsPolicyConfig(name string, description string, priority string) string {
	return fmt.Sprintf(`

	resource "ibm_logs_policy" "logs_policy_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "%s"
		priority    = "%s"
		application_rule {
		  name         = "otel-links-test"
		  rule_type_id = "start_with"
		}
		log_rules {
		  severities = ["info"]
		}
	  }
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, priority)
}

func testAccCheckIbmLogsPolicyExists(n string, obj logsv0.Policy) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPolicyOptions := &logsv0.GetPolicyOptions{}

		getPolicyOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		policyIntf, _, err := logsClient.GetPolicy(getPolicyOptions)
		if err != nil {
			return err
		}

		policy := policyIntf.(*logsv0.Policy)
		obj = *policy
		return nil
	}
}

func testAccCheckIbmLogsPolicyDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_policy" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getPolicyOptions := &logsv0.GetPolicyOptions{}

		getPolicyOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetPolicy(getPolicyOptions)

		if err == nil {
			return fmt.Errorf("logs_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
