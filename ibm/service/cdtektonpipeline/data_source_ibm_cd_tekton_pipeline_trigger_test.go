// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdTektonPipelineTriggerDataSourceBasic(t *testing.T) {
	triggerPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	triggerType := "manual"
	triggerName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	triggerEventListener := "listener"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerDataSourceConfigBasic(triggerPipelineID, triggerType, triggerName, triggerEventListener),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "trigger_id"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineTriggerDataSourceAllArgs(t *testing.T) {
	triggerPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	triggerType := "manual"
	triggerName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	triggerEventListener := fmt.Sprintf("tf_event_listener_%d", acctest.RandIntRange(10, 100))
	triggerMaxConcurrentRuns := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	triggerEnabled := "false"
	triggerCron := fmt.Sprintf("tf_cron_%d", acctest.RandIntRange(10, 100))
	triggerTimezone := fmt.Sprintf("tf_timezone_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerDataSourceConfig(triggerPipelineID, triggerType, triggerName, triggerEventListener, triggerMaxConcurrentRuns, triggerEnabled, triggerCron, triggerTimezone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "trigger_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "event_listener"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "max_concurrent_runs"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "secret.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerDataSourceConfigBasic(triggerPipelineID string, triggerType string, triggerName string, triggerEventListener string) string {
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}
		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-name"
			}
		}
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			next_build_number = 5
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
		resource "ibm_cd_toolchain_tool_githubconsolidated" "definition-repo" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			name = "definition-repo"
			initialization {
				type = "link"
				repo_url = "https://github.com/open-toolchain/hello-tekton.git"
			}
			parameters {}
		}
		resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			source {
				type = "git"
				properties {
					url = "https://github.com/open-toolchain/hello-tekton.git"
					branch = "master"
					path = ".tekton"
				}
			}
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			type = "%s"
			name = "%s"
			event_listener = "%s"
		}

		data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
		}
	`, rgName, tcName, triggerType, triggerName, triggerEventListener)
}

func testAccCheckIBMCdTektonPipelineTriggerDataSourceConfig(triggerPipelineID string, triggerType string, triggerName string, triggerEventListener string, triggerMaxConcurrentRuns string, triggerEnabled string, triggerCron string, triggerTimezone string) string {
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}
		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-name"
			}
		}
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			next_build_number = 5
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
		resource "ibm_cd_toolchain_tool_githubconsolidated" "definition-repo" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			name = "definition-repo"
			initialization {
				type = "link"
				repo_url = "https://github.com/open-toolchain/hello-tekton.git"
			}
			parameters {}
		}
		resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			source {
				type = "git"
				properties {
					url = "https://github.com/open-toolchain/hello-tekton.git"
					branch = "master"
					path = ".tekton"
				}
			}
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			type = "%s"
			name = "%s"
			event_listener = "%s"
			max_concurrent_runs = %s
			enabled = %s
		}

		data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
		}
	`, rgName, tcName, triggerType, triggerName, triggerEventListener, triggerMaxConcurrentRuns, triggerEnabled)
}
