// Copyright IBM Corp. 2026 All Rights Reserved.
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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtektonpipeline"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCdTektonPipelineTriggerBasic(t *testing.T) {
	var conf cdtektonpipelinev2.Trigger
	typeVar := "manual"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	eventListener := "listener"
	typeVarUpdate := "manual"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	eventListenerUpdate := fmt.Sprintf("tf_event_listener_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfigBasic("", typeVar, name, eventListener),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "trigger_id"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "event_listener", eventListener),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfigBasic("", typeVarUpdate, nameUpdate, eventListenerUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "event_listener", eventListenerUpdate),
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
	enabled := "false"
	cron := fmt.Sprintf("*/5 10 10 %d *", acctest.RandIntRange(1, 12))
	timezone := "Europe/London"
	filter := "test"
	favorite := "false"
	enableEventsFromForks := "false"
	disableDraftEvents := "false"
	limitWaitingRuns := "false"
	typeVarUpdate := "generic"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	eventListenerUpdate := fmt.Sprintf("tf_event_listener_%d", acctest.RandIntRange(10, 100))
	maxConcurrentRunsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 2))
	enabledUpdate := "true"
	cronUpdate := fmt.Sprintf("*/10 %d 10 10 *", acctest.RandIntRange(1, 23))
	timezoneUpdate := "America/New_York"
	filterUpdate := "true"
	favoriteUpdate := "true"
	enableEventsFromForksUpdate := "true"
	disableDraftEventsUpdate := "true"
	limitWaitingRunsUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfig(pipelineID, typeVar, name, eventListener, maxConcurrentRuns, enabled, favorite, limitWaitingRuns, enableEventsFromForks, disableDraftEvents, filter, cron, timezone),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", conf),
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", conf),
					testAccCheckIBMCdTektonPipelineTriggerExists("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "limit_waiting_runs", limitWaitingRuns),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "max_concurrent_runs", maxConcurrentRuns),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "cron", cron),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "timezone", timezone),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "favorite", favorite),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "trigger_id"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "type", "manual"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "tags.#"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "type", "timer"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "type", "generic"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "secret.#"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "webhook_url"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerConfig(pipelineID, typeVarUpdate, nameUpdate, eventListenerUpdate, maxConcurrentRunsUpdate, enabledUpdate, favoriteUpdate, limitWaitingRunsUpdate, enableEventsFromForksUpdate, disableDraftEventsUpdate, filterUpdate, cronUpdate, timezoneUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "limit_waiting_runs", limitWaitingRunsUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "max_concurrent_runs", maxConcurrentRunsUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "cron", cronUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "timezone", timezoneUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "favorite", favoriteUpdate),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "trigger_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "trigger_id"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "type", "manual"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "tags.#"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger2_instance", "type", "timer"),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "type", "generic"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "secret.#"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger3_instance", "webhook_url"),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pipeline_id", "enable_events_from_forks", "disable_draft_events"},
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerConfigBasic(pipelineID string, typeVar string, name string, eventListener string) string {
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
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
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
		resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			source {
				type = "git"
				properties {
					url = "https://github.com/open-toolchain/hello-tekton.git"
					branch = "master"
					path = ".tekton"
				}
			}
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance
			]
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance
			]
			type = "%s"
			name = "%s"
			event_listener = "%s"
		}
	`, rgName, tcName, typeVar, name, eventListener)
}

func testAccCheckIBMCdTektonPipelineTriggerConfig(pipelineID string, typeVar string, name string, eventListener string, maxConcurrentRuns string, enabled string, favorite string, limitWaitingRuns string, enableEventsFromForks string, disableDraftEvents string, filter string, cron string, timezone string) string {
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
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
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
		resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			source {
				type = "git"
				properties {
					url = "https://github.com/open-toolchain/hello-tekton.git"
					branch = "master"
					path = ".tekton"
				}
			}
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance
			]
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance
			]
			type = "manual"
			event_listener = "listener"
			name = "%s"
			tags = [ "tag1", "tag2" ]
			limit_waiting_runs = %s
			max_concurrent_runs = %s
			enabled = %s
			favorite = %s
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger2_instance" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance
			]
			type = "timer"
			name = "timer1"
			event_listener = "listener"
			cron = "%s"
			timezone = "%s"
		}
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger3_instance" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance
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
	`, rgName, tcName, name, limitWaitingRuns, maxConcurrentRuns, enabled, favorite, cron, timezone)
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

func TestResourceIBMCdTektonPipelineTriggerWorkerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["type"] = "testString"
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Worker)
	model.Name = core.StringPtr("testString")
	model.Type = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerWorkerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerTriggerSourceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		triggerSourcePropertiesModel := make(map[string]interface{})
		triggerSourcePropertiesModel["url"] = "testString"
		triggerSourcePropertiesModel["branch"] = "testString"
		triggerSourcePropertiesModel["pattern"] = "testString"
		triggerSourcePropertiesModel["blind_connection"] = true
		triggerSourcePropertiesModel["hook_id"] = "testString"
		triggerSourcePropertiesModel["tool"] = []map[string]interface{}{toolModel}

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["properties"] = []map[string]interface{}{triggerSourcePropertiesModel}

		assert.Equal(t, result, model)
	}

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	triggerSourcePropertiesModel := new(cdtektonpipelinev2.TriggerSourceProperties)
	triggerSourcePropertiesModel.URL = core.StringPtr("testString")
	triggerSourcePropertiesModel.Branch = core.StringPtr("testString")
	triggerSourcePropertiesModel.Pattern = core.StringPtr("testString")
	triggerSourcePropertiesModel.BlindConnection = core.BoolPtr(true)
	triggerSourcePropertiesModel.HookID = core.StringPtr("testString")
	triggerSourcePropertiesModel.Tool = toolModel

	model := new(cdtektonpipelinev2.TriggerSource)
	model.Type = core.StringPtr("testString")
	model.Properties = triggerSourcePropertiesModel

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerTriggerSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerTriggerSourcePropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		model := make(map[string]interface{})
		model["url"] = "testString"
		model["branch"] = "testString"
		model["pattern"] = "testString"
		model["blind_connection"] = true
		model["hook_id"] = "testString"
		model["tool"] = []map[string]interface{}{toolModel}

		assert.Equal(t, result, model)
	}

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	model := new(cdtektonpipelinev2.TriggerSourceProperties)
	model.URL = core.StringPtr("testString")
	model.Branch = core.StringPtr("testString")
	model.Pattern = core.StringPtr("testString")
	model.BlindConnection = core.BoolPtr(true)
	model.HookID = core.StringPtr("testString")
	model.Tool = toolModel

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerTriggerSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerToolToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Tool)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerToolToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerGenericSecretToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["type"] = "token_matches"
		model["value"] = "testString"
		model["source"] = "header"
		model["key_name"] = "testString"
		model["algorithm"] = "md4"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.GenericSecret)
	model.Type = core.StringPtr("token_matches")
	model.Value = core.StringPtr("testString")
	model.Source = core.StringPtr("header")
	model.KeyName = core.StringPtr("testString")
	model.Algorithm = core.StringPtr("md4")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerGenericSecretToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerTriggerPropertyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["value"] = "testString"
		model["href"] = "testString"
		model["enum"] = []string{"testString"}
		model["type"] = "secure"
		model["path"] = "testString"
		model["locked"] = true

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.TriggerProperty)
	model.Name = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")
	model.Enum = []string{"testString"}
	model.Type = core.StringPtr("secure")
	model.Path = core.StringPtr("testString")
	model.Locked = core.BoolPtr(true)

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerTriggerPropertyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerMapToWorkerIdentity(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.WorkerIdentity) {
		model := new(cdtektonpipelinev2.WorkerIdentity)
		model.ID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerMapToWorkerIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerMapToGenericSecret(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.GenericSecret) {
		model := new(cdtektonpipelinev2.GenericSecret)
		model.Type = core.StringPtr("token_matches")
		model.Value = core.StringPtr("testString")
		model.Source = core.StringPtr("header")
		model.KeyName = core.StringPtr("testString")
		model.Algorithm = core.StringPtr("md4")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["type"] = "token_matches"
	model["value"] = "testString"
	model["source"] = "header"
	model["key_name"] = "testString"
	model["algorithm"] = "md4"

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerMapToGenericSecret(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerMapToTriggerSourcePrototype(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.TriggerSourcePrototype) {
		triggerSourcePropertiesPrototypeModel := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
		triggerSourcePropertiesPrototypeModel.URL = core.StringPtr("testString")
		triggerSourcePropertiesPrototypeModel.Branch = core.StringPtr("testString")
		triggerSourcePropertiesPrototypeModel.Pattern = core.StringPtr("testString")

		model := new(cdtektonpipelinev2.TriggerSourcePrototype)
		model.Type = core.StringPtr("testString")
		model.Properties = triggerSourcePropertiesPrototypeModel

		assert.Equal(t, result, model)
	}

	triggerSourcePropertiesPrototypeModel := make(map[string]interface{})
	triggerSourcePropertiesPrototypeModel["url"] = "testString"
	triggerSourcePropertiesPrototypeModel["branch"] = "testString"
	triggerSourcePropertiesPrototypeModel["pattern"] = "testString"

	model := make(map[string]interface{})
	model["type"] = "testString"
	model["properties"] = []interface{}{triggerSourcePropertiesPrototypeModel}

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerMapToTriggerSourcePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerMapToTriggerSourcePropertiesPrototype(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.TriggerSourcePropertiesPrototype) {
		model := new(cdtektonpipelinev2.TriggerSourcePropertiesPrototype)
		model.URL = core.StringPtr("testString")
		model.Branch = core.StringPtr("testString")
		model.Pattern = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["url"] = "testString"
	model["branch"] = "testString"
	model["pattern"] = "testString"

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerMapToTriggerSourcePropertiesPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
