// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.103.0-e8b84313-20250402-201816
 */

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtektonpipeline"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
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
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "trigger_id"),
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
	triggerFavorite := "true"
	triggerLimitWaitingRuns := "true"
	triggerEnableEventsFromForks := "true"
	triggerFilter := fmt.Sprintf("tf_filter_%d", acctest.RandIntRange(10, 100))
	triggerCron := fmt.Sprintf("tf_cron_%d", acctest.RandIntRange(10, 100))
	triggerTimezone := fmt.Sprintf("tf_timezone_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerDataSourceConfig(triggerPipelineID, triggerType, triggerName, triggerEventListener, triggerMaxConcurrentRuns, triggerEnabled, triggerFavorite, triggerLimitWaitingRuns, triggerEnableEventsFromForks, triggerFilter, triggerCron, triggerTimezone),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "trigger_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "event_listener"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "max_concurrent_runs"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "favorite"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance", "limit_waiting_runs"),
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

		data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
		}
	`, rgName, tcName, triggerType, triggerName, triggerEventListener)
}

func testAccCheckIBMCdTektonPipelineTriggerDataSourceConfig(triggerPipelineID string, triggerType string, triggerName string, triggerEventListener string, triggerMaxConcurrentRuns string, triggerEnabled string, triggerFavorite string, triggerLimitWaitingRuns string, triggerEnableEventsFromForks string, triggerFilter string, triggerCron string, triggerTimezone string) string {
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
			limit_waiting_runs = %s
			max_concurrent_runs = %s
			enabled = %s
			favorite = %s
		}

		data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
		}
	`, rgName, tcName, triggerType, triggerName, triggerEventListener, triggerLimitWaitingRuns, triggerMaxConcurrentRuns, triggerEnabled, triggerFavorite)
}

func TestDataSourceIBMCdTektonPipelineTriggerTriggerPropertyToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerTriggerPropertyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerWorkerToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerWorkerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerTriggerSourceToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerTriggerSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerTriggerSourcePropertiesToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerTriggerSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerToolToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Tool)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerToolToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerGenericSecretToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerGenericSecretToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
