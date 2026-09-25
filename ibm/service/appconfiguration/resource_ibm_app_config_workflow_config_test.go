// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TestAccIbmAppConfigWorkflowConfigBasic tests CREATE, UPDATE, and DELETE of a
// workflow configuration. Requires an App Configuration enterprise plan instance
// with the workflow feature enabled. Set APP_CONFIG_WORKFLOW_INSTANCE_GUID to
// enable this test; without it the test is skipped.
func TestAccIbmAppConfigWorkflowConfigBasic(t *testing.T) {
	if os.Getenv("APP_CONFIG_WORKFLOW_INSTANCE_GUID") == "" {
		t.Skip("APP_CONFIG_WORKFLOW_INSTANCE_GUID not set — skipping workflow test (requires enterprise plan with workflow enabled)")
	}

	var conf appconfigurationv1.WorkflowConfigResponse

	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_workflow_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_workflow_updated_%d", acctest.RandIntRange(10, 100))
	workflowID := fmt.Sprintf("tf-wf-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmAppConfigWorkflowConfigDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: CREATE
				Config: testAccCheckIbmAppConfigWorkflowConfigBasic(instanceName, name, workflowID, "https://example.service-now.com", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmAppConfigWorkflowConfigExists("ibm_app_config_workflow_config.test", conf),
					resource.TestCheckResourceAttrSet("ibm_app_config_workflow_config.test", "id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_workflow_config.test", "workflow_config_id"),
					resource.TestCheckResourceAttrSet("ibm_app_config_workflow_config.test", "href"),
					resource.TestCheckResourceAttrSet("ibm_app_config_workflow_config.test", "created_time"),
					resource.TestCheckResourceAttrSet("ibm_app_config_workflow_config.test", "updated_time"),
					resource.TestCheckResourceAttr("ibm_app_config_workflow_config.test", "name", name),
					resource.TestCheckResourceAttr("ibm_app_config_workflow_config.test", "enabled", "true"),
				),
			},
			{
				// Step 2: UPDATE — change name and disable
				Config: testAccCheckIbmAppConfigWorkflowConfigBasic(instanceName, nameUpdate, workflowID, "https://example.service-now.com", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_app_config_workflow_config.test", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_app_config_workflow_config.test", "enabled", "false"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigWorkflowConfigBasic(instanceName, name, workflowID, workflowURL string, enabled bool) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_wf_test" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "enterprise"
		}

		resource "ibm_app_config_workflow_config" "test" {
			guid        = ibm_resource_instance.app_config_wf_test.guid
			name        = "%s"
			workflow_id = "%s"
			enabled     = %t

			# Named "workflow_provider" instead of "provider" because "provider" is
			# a reserved meta-argument in Terraform's plugin SDK.
			workflow_provider {
				type = "SERVICENOW_EXTERNAL"
				metadata {
					workflow_url        = "%s"
					approval_group_name = "TF Test Group"
					approval_expiration = 7
				}
			}

			scope {
				environments {
					environment_id = "dev"
					resources {
						environment {
							enable = true
						}
					}
				}
			}
		}`, instanceName, name, workflowID, enabled, workflowURL)
}

func testAccCheckIbmAppConfigWorkflowConfigExists(n string, obj appconfigurationv1.WorkflowConfigResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}
		// ID format: guid/workflow_config_id
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}
		if len(parts) < 2 {
			return flex.FmtErrorf("invalid ID format, expected guid/workflow_config_id")
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}

		options := &appconfigurationv1.GetWorkflowConfigOptions{}
		options.SetWorkflowConfigID(parts[1])

		result, _, err := appconfigClient.GetWorkflowConfig(options)
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}

		obj = *result
		return nil
	}
}

func testAccCheckIbmAppConfigWorkflowConfigDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_app_config_workflow_config" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}
		if len(parts) < 2 {
			continue
		}

		appconfigClient, err := getAppConfigClient(acc.TestAccProvider.Meta(), parts[0])
		if err != nil {
			return flex.FmtErrorf("%s", err)
		}

		options := &appconfigurationv1.GetWorkflowConfigOptions{}
		options.SetWorkflowConfigID(parts[1])

		_, response, err := appconfigClient.GetWorkflowConfig(options)
		if err == nil {
			return flex.FmtErrorf("WorkflowConfig still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return flex.FmtErrorf("[ERROR] Error checking for WorkflowConfig (%s) destroyed: %s", rs.Primary.ID, err)
		}
	}
	return nil
}
