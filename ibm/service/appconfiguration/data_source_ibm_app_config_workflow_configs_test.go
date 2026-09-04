// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIbmAppConfigWorkflowConfigsDataSource tests the data source for listing
// workflow configurations. Requires an enterprise App Configuration instance
// with the workflow feature enabled. Set APP_CONFIG_WORKFLOW_INSTANCE_GUID to
// enable this test; without it the test is skipped.
func TestAccIbmAppConfigWorkflowConfigsDataSource(t *testing.T) {
	if os.Getenv("APP_CONFIG_WORKFLOW_INSTANCE_GUID") == "" {
		t.Skip("APP_CONFIG_WORKFLOW_INSTANCE_GUID not set — skipping workflow test (requires enterprise plan with workflow enabled)")
	}

	instanceName := fmt.Sprintf("tf_app_config_test_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_workflow_%d", acctest.RandIntRange(10, 100))
	workflowID := fmt.Sprintf("tf-wf-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAppConfigWorkflowConfigsDataSourceConfig(instanceName, name, workflowID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "workflow_configs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "workflow_configs.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "workflow_configs.0.enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "workflow_configs.0.href"),
					// workflow_provider is the renamed schema field (API field is "provider",
					// but "provider" is reserved in Terraform's plugin SDK).
					resource.TestCheckResourceAttrSet("data.ibm_app_config_workflow_configs.test_ds", "workflow_configs.0.workflow_provider.#"),
				),
			},
		},
	})
}

func testAccCheckIbmAppConfigWorkflowConfigsDataSourceConfig(instanceName, name, workflowID string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "app_config_wfs_ds_test" {
			name     = "%s"
			location = "us-south"
			service  = "apprapp"
			plan     = "enterprise"
		}

		resource "ibm_app_config_workflow_config" "wfs_ds_resource" {
			guid        = ibm_resource_instance.app_config_wfs_ds_test.guid
			name        = "%s"
			workflow_id = "%s"
			enabled     = true

			workflow_provider {
				type = "SERVICENOW_EXTERNAL"
				metadata {
					workflow_url        = "https://example.service-now.com"
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
		}

		data "ibm_app_config_workflow_configs" "test_ds" {
			guid = ibm_resource_instance.app_config_wfs_ds_test.guid
			depends_on = [ibm_app_config_workflow_config.wfs_ds_resource]
		}`, instanceName, name, workflowID)
}
