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

func TestAccIBMCdTektonPipelineTriggerDataSourceBasic(t *testing.T) {
	triggerPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerDataSourceConfigBasic(triggerPipelineID),
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
	triggerDisabled := "true"
	triggerCron := fmt.Sprintf("tf_cron_%d", acctest.RandIntRange(10, 100))
	triggerTimezone := fmt.Sprintf("tf_timezone_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerDataSourceConfig(triggerPipelineID, triggerType, triggerName, triggerEventListener, triggerMaxConcurrentRuns, triggerDisabled, triggerCron, triggerTimezone),
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
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "scm_source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "events.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "secret.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerDataSourceConfigBasic(triggerPipelineID string) string {
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
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
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			scm_source {
				url = "https://github.com/open-toolchain/hello-tekton.git"
				branch = "master"
				path = ".tekton"
			}
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
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
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			type = "manual"
			event_listener = "listener"
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			name = "trigger"
		}

		data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
		}
	`, tcName, rgID)
}

func testAccCheckIBMCdTektonPipelineTriggerDataSourceConfig(triggerPipelineID string, triggerType string, triggerName string, triggerEventListener string, triggerMaxConcurrentRuns string, triggerDisabled string, triggerCron string, triggerTimezone string) string {
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
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
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			scm_source {
				url = "https://github.com/open-toolchain/hello-tekton.git"
				branch = "master"
				path = ".tekton"
			}
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
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
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			type = "manual"
			event_listener = "listener"
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			name = "%s"
			max_concurrent_runs = %s
			disabled = %s
		}

		data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
		}
	`, tcName, rgID, triggerName, triggerMaxConcurrentRuns, triggerDisabled)
}
