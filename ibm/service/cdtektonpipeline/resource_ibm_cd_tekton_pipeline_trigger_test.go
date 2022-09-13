// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
)

func TestAccIBMCdTektonPipelineTriggerBasic(t *testing.T) {
	var conf cdtektonpipelinev2.Trigger

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfigBasic(""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "trigger_id"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "type", "manual"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "name", "trigger"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "event_listener", "listener"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineTriggerAllArgs(t *testing.T) {
	var conf cdtektonpipelinev2.Trigger
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	typeVar := "manual"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	eventListener := "listener"
	maxConcurrentRuns := fmt.Sprintf("%d", acctest.RandIntRange(3, 4))
	disabled := "true"
	cron := fmt.Sprintf("*/5 10 10 %d *", acctest.RandIntRange(1, 12))
	timezone := "Europe/London"
	typeVarUpdate := "generic"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	eventListenerUpdate := fmt.Sprintf("tf_event_listener_%d", acctest.RandIntRange(10, 100))
	maxConcurrentRunsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 2))
	disabledUpdate := "false"
	cronUpdate := fmt.Sprintf("*/10 %d 10 10 *", acctest.RandIntRange(1, 23))
	timezoneUpdate := "America/New_York"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfig(pipelineID, typeVar, name, eventListener, maxConcurrentRuns, disabled, cron, timezone),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "max_concurrent_runs", maxConcurrentRuns),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "disabled", disabled),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "cron", cron),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "timezone", timezone),
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", conf),
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", conf),
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "trigger_id"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "type", "manual"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "tags.#"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "type", "timer"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "type", "generic"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "secret.#"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "webhook_url"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfig(pipelineID, typeVarUpdate, nameUpdate, eventListenerUpdate, maxConcurrentRunsUpdate, disabledUpdate, cronUpdate, timezoneUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "max_concurrent_runs", maxConcurrentRunsUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "disabled", disabledUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "cron", cronUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "timezone", timezoneUpdate),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "trigger_id"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "type", "manual"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger", "tags.#"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2", "type", "timer"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "type", "generic"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "secret.#"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3", "webhook_url"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerConfigBasic(pipelineID string) string {
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
	`, tcName, rgID)
}

func testAccCheckIBMCdTektonPipelineTriggerConfig(pipelineID string, typeVar string, name string, eventListener string, maxConcurrentRuns string, disabled string, cron string, timezone string) string {
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
			tags = [ "tag1", "tag2" ]
			max_concurrent_runs = %s
			disabled = %s
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger2" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			type = "timer"
			name = "timer1"
			event_listener = "listener"
			cron = "%s"
			timezone = "%s"
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger3" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			type = "generic"
			name = "generic1"
			event_listener = "listener"
			secret {
				type = "token_matches"
				value = "value"
				source = "header"
				key_name = "key_name"
				algorithm = "md4"
			}
		}
	`, tcName, rgID, name, maxConcurrentRuns, disabled, cron, timezone)
}

func testAccCheckIBMCdTektonPipelineTriggerExists(n string, obj cdtektonpipelinev2.Trigger) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerOptions.SetTriggerID(parts[1])

		triggerIntf, _, err := cdTektonPipelineClient.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)
		if err != nil {
			return err
		}

		trigger := triggerIntf.(*cdtektonpipelinev2.Trigger)
		obj = *trigger
		return nil
	}
}

func testAccCheckIBMCdTektonPipelineTriggerDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline_trigger" {
			continue
		}

		getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerOptions.SetTriggerID(parts[1])

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)

		if err == nil {
			return fmt.Errorf("cd_tekton_pipeline_trigger still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_tekton_pipeline_trigger (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
