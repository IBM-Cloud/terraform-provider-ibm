// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdTektonPipelineDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "runs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enabled"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineDataSourceAllArgs(t *testing.T) {
	tektonPipelineEnableSlackNotifications := "true"
	tektonPipelineEnablePartialCloning := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDataSourceConfig(tektonPipelineEnableSlackNotifications, tektonPipelineEnablePartialCloning),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "definitions.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.0.path"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.event_listener"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.max_concurrent_runs"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.cron"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.timezone"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.0.webhook_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "runs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_slack_notifications"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_partial_cloning"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineDataSourceConfigBasic() string {
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}
		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-name"
				type = "tekton"
			}
		}
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
		data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
		}
	`, tcName, rgID)
}

func testAccCheckIBMCdTektonPipelineDataSourceConfig(tektonPipelineEnableSlackNotifications string, tektonPipelineEnablePartialCloning string) string {
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}
		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-name"
				type = "tekton"
			}
		}
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			enable_slack_notifications = %s
			enable_partial_cloning = %s
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
		data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
		}
	`, tcName, rgID, tektonPipelineEnableSlackNotifications, tektonPipelineEnablePartialCloning)
}
