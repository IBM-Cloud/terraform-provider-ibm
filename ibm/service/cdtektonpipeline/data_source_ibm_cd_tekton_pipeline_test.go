// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.103.0-e8b84313-20250402-201816
 */

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtektonpipeline"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCdTektonPipelineDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "runs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_notifications"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_partial_cloning"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enabled"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineDataSourceAllArgs(t *testing.T) {
	tektonPipelineNextBuildNumber := fmt.Sprintf("%d", acctest.RandIntRange(1, 99999999999999))
	tektonPipelineEnableNotifications := "true"
	tektonPipelineEnablePartialCloning := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDataSourceConfig(tektonPipelineNextBuildNumber, tektonPipelineEnableNotifications, tektonPipelineEnablePartialCloning),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "runs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "next_build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_notifications"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_partial_cloning"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineDataSourceConfigBasic() string {
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
		data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineDataSourceConfig(tektonPipelineNextBuildNumber string, tektonPipelineEnableNotifications string, tektonPipelineEnablePartialCloning string) string {
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
			next_build_number = %s
			enable_notifications = %s
			enable_partial_cloning = %s
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
		resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			name = "property1"
			type = "text"
			value = "prop1"
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance
			]
		}
		data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
		}
	`, rgName, tcName, tektonPipelineNextBuildNumber, tektonPipelineEnableNotifications, tektonPipelineEnablePartialCloning)
}

func TestDataSourceIBMCdTektonPipelineResourceGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.ResourceGroupReference)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineToolchainReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["crn"] = "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.ToolchainReference)
	model.ID = core.StringPtr("testString")
	model.CRN = core.StringPtr("crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineToolchainReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		definitionSourcePropertiesModel := make(map[string]interface{})
		definitionSourcePropertiesModel["url"] = "testString"
		definitionSourcePropertiesModel["branch"] = "testString"
		definitionSourcePropertiesModel["tag"] = "testString"
		definitionSourcePropertiesModel["path"] = "testString"
		definitionSourcePropertiesModel["tool"] = []map[string]interface{}{toolModel}

		definitionSourceModel := make(map[string]interface{})
		definitionSourceModel["type"] = "testString"
		definitionSourceModel["properties"] = []map[string]interface{}{definitionSourcePropertiesModel}

		model := make(map[string]interface{})
		model["source"] = []map[string]interface{}{definitionSourceModel}
		model["href"] = "testString"
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
	definitionSourcePropertiesModel.URL = core.StringPtr("testString")
	definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
	definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
	definitionSourcePropertiesModel.Path = core.StringPtr("testString")
	definitionSourcePropertiesModel.Tool = toolModel

	definitionSourceModel := new(cdtektonpipelinev2.DefinitionSource)
	definitionSourceModel.Type = core.StringPtr("testString")
	definitionSourceModel.Properties = definitionSourcePropertiesModel

	model := new(cdtektonpipelinev2.Definition)
	model.Source = definitionSourceModel
	model.Href = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineDefinitionSourceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		definitionSourcePropertiesModel := make(map[string]interface{})
		definitionSourcePropertiesModel["url"] = "testString"
		definitionSourcePropertiesModel["branch"] = "testString"
		definitionSourcePropertiesModel["tag"] = "testString"
		definitionSourcePropertiesModel["path"] = "testString"
		definitionSourcePropertiesModel["tool"] = []map[string]interface{}{toolModel}

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["properties"] = []map[string]interface{}{definitionSourcePropertiesModel}

		assert.Equal(t, result, model)
	}

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	definitionSourcePropertiesModel := new(cdtektonpipelinev2.DefinitionSourceProperties)
	definitionSourcePropertiesModel.URL = core.StringPtr("testString")
	definitionSourcePropertiesModel.Branch = core.StringPtr("testString")
	definitionSourcePropertiesModel.Tag = core.StringPtr("testString")
	definitionSourcePropertiesModel.Path = core.StringPtr("testString")
	definitionSourcePropertiesModel.Tool = toolModel

	model := new(cdtektonpipelinev2.DefinitionSource)
	model.Type = core.StringPtr("testString")
	model.Properties = definitionSourcePropertiesModel

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinitionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineDefinitionSourcePropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		model := make(map[string]interface{})
		model["url"] = "testString"
		model["branch"] = "testString"
		model["tag"] = "testString"
		model["path"] = "testString"
		model["tool"] = []map[string]interface{}{toolModel}

		assert.Equal(t, result, model)
	}

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	model := new(cdtektonpipelinev2.DefinitionSourceProperties)
	model.URL = core.StringPtr("testString")
	model.Branch = core.StringPtr("testString")
	model.Tag = core.StringPtr("testString")
	model.Path = core.StringPtr("testString")
	model.Tool = toolModel

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinitionSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineToolToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Tool)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineToolToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelinePropertyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["value"] = "testString"
		model["href"] = "testString"
		model["enum"] = []string{"testString"}
		model["type"] = "secure"
		model["locked"] = true
		model["path"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Property)
	model.Name = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")
	model.Enum = []string{"testString"}
	model.Type = core.StringPtr("secure")
	model.Locked = core.BoolPtr(true)
	model.Path = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelinePropertyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		triggerPropertyModel := make(map[string]interface{})
		triggerPropertyModel["name"] = "testString"
		triggerPropertyModel["value"] = "testString"
		triggerPropertyModel["href"] = "testString"
		triggerPropertyModel["enum"] = []string{"testString"}
		triggerPropertyModel["type"] = "secure"
		triggerPropertyModel["path"] = "testString"
		triggerPropertyModel["locked"] = true

		workerModel := make(map[string]interface{})
		workerModel["name"] = "testString"
		workerModel["type"] = "testString"
		workerModel["id"] = "testString"

		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		triggerSourcePropertiesModel := make(map[string]interface{})
		triggerSourcePropertiesModel["url"] = "testString"
		triggerSourcePropertiesModel["branch"] = "testString"
		triggerSourcePropertiesModel["pattern"] = "testString"
		triggerSourcePropertiesModel["blind_connection"] = true
		triggerSourcePropertiesModel["hook_id"] = "testString"
		triggerSourcePropertiesModel["tool"] = []map[string]interface{}{toolModel}

		triggerSourceModel := make(map[string]interface{})
		triggerSourceModel["type"] = "testString"
		triggerSourceModel["properties"] = []map[string]interface{}{triggerSourcePropertiesModel}

		genericSecretModel := make(map[string]interface{})
		genericSecretModel["type"] = "token_matches"
		genericSecretModel["value"] = "testString"
		genericSecretModel["source"] = "header"
		genericSecretModel["key_name"] = "testString"
		genericSecretModel["algorithm"] = "md4"

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["name"] = "start-deploy"
		model["href"] = "testString"
		model["event_listener"] = "testString"
		model["id"] = "testString"
		model["properties"] = []map[string]interface{}{triggerPropertyModel}
		model["tags"] = []string{"testString"}
		model["worker"] = []map[string]interface{}{workerModel}
		model["max_concurrent_runs"] = int(4)
		model["enabled"] = true
		model["favorite"] = false
		model["limit_waiting_runs"] = false
		model["enable_events_from_forks"] = false
		model["disable_draft_events"] = false
		model["source"] = []map[string]interface{}{triggerSourceModel}
		model["events"] = []string{"push", "pull_request"}
		model["filter"] = "header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'"
		model["cron"] = "testString"
		model["timezone"] = "America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC"
		model["secret"] = []map[string]interface{}{genericSecretModel}
		model["webhook_url"] = "testString"

		assert.Equal(t, result, model)
	}

	triggerPropertyModel := new(cdtektonpipelinev2.TriggerProperty)
	triggerPropertyModel.Name = core.StringPtr("testString")
	triggerPropertyModel.Value = core.StringPtr("testString")
	triggerPropertyModel.Href = core.StringPtr("testString")
	triggerPropertyModel.Enum = []string{"testString"}
	triggerPropertyModel.Type = core.StringPtr("secure")
	triggerPropertyModel.Path = core.StringPtr("testString")
	triggerPropertyModel.Locked = core.BoolPtr(true)

	workerModel := new(cdtektonpipelinev2.Worker)
	workerModel.Name = core.StringPtr("testString")
	workerModel.Type = core.StringPtr("testString")
	workerModel.ID = core.StringPtr("testString")

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	triggerSourcePropertiesModel := new(cdtektonpipelinev2.TriggerSourceProperties)
	triggerSourcePropertiesModel.URL = core.StringPtr("testString")
	triggerSourcePropertiesModel.Branch = core.StringPtr("testString")
	triggerSourcePropertiesModel.Pattern = core.StringPtr("testString")
	triggerSourcePropertiesModel.BlindConnection = core.BoolPtr(true)
	triggerSourcePropertiesModel.HookID = core.StringPtr("testString")
	triggerSourcePropertiesModel.Tool = toolModel

	triggerSourceModel := new(cdtektonpipelinev2.TriggerSource)
	triggerSourceModel.Type = core.StringPtr("testString")
	triggerSourceModel.Properties = triggerSourcePropertiesModel

	genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
	genericSecretModel.Type = core.StringPtr("token_matches")
	genericSecretModel.Value = core.StringPtr("testString")
	genericSecretModel.Source = core.StringPtr("header")
	genericSecretModel.KeyName = core.StringPtr("testString")
	genericSecretModel.Algorithm = core.StringPtr("md4")

	model := new(cdtektonpipelinev2.Trigger)
	model.Type = core.StringPtr("testString")
	model.Name = core.StringPtr("start-deploy")
	model.Href = core.StringPtr("testString")
	model.EventListener = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.Properties = []cdtektonpipelinev2.TriggerProperty{*triggerPropertyModel}
	model.Tags = []string{"testString"}
	model.Worker = workerModel
	model.MaxConcurrentRuns = core.Int64Ptr(int64(4))
	model.Enabled = core.BoolPtr(true)
	model.Favorite = core.BoolPtr(false)
	model.LimitWaitingRuns = core.BoolPtr(false)
	model.EnableEventsFromForks = core.BoolPtr(false)
	model.DisableDraftEvents = core.BoolPtr(false)
	model.Source = triggerSourceModel
	model.Events = []string{"push", "pull_request"}
	model.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
	model.Cron = core.StringPtr("testString")
	model.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
	model.Secret = genericSecretModel
	model.WebhookURL = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerPropertyToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerPropertyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineWorkerToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineWorkerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerSourceToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerSourcePropertiesToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineGenericSecretToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineGenericSecretToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerManualTriggerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		triggerPropertyModel := make(map[string]interface{})
		triggerPropertyModel["name"] = "testString"
		triggerPropertyModel["value"] = "testString"
		triggerPropertyModel["href"] = "testString"
		triggerPropertyModel["enum"] = []string{"testString"}
		triggerPropertyModel["type"] = "secure"
		triggerPropertyModel["path"] = "testString"
		triggerPropertyModel["locked"] = true

		workerModel := make(map[string]interface{})
		workerModel["name"] = "testString"
		workerModel["type"] = "testString"
		workerModel["id"] = "testString"

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["name"] = "start-deploy"
		model["href"] = "testString"
		model["event_listener"] = "testString"
		model["id"] = "testString"
		model["properties"] = []map[string]interface{}{triggerPropertyModel}
		model["tags"] = []string{"testString"}
		model["worker"] = []map[string]interface{}{workerModel}
		model["max_concurrent_runs"] = int(4)
		model["enabled"] = true
		model["favorite"] = false
		model["limit_waiting_runs"] = false

		assert.Equal(t, result, model)
	}

	triggerPropertyModel := new(cdtektonpipelinev2.TriggerProperty)
	triggerPropertyModel.Name = core.StringPtr("testString")
	triggerPropertyModel.Value = core.StringPtr("testString")
	triggerPropertyModel.Href = core.StringPtr("testString")
	triggerPropertyModel.Enum = []string{"testString"}
	triggerPropertyModel.Type = core.StringPtr("secure")
	triggerPropertyModel.Path = core.StringPtr("testString")
	triggerPropertyModel.Locked = core.BoolPtr(true)

	workerModel := new(cdtektonpipelinev2.Worker)
	workerModel.Name = core.StringPtr("testString")
	workerModel.Type = core.StringPtr("testString")
	workerModel.ID = core.StringPtr("testString")

	model := new(cdtektonpipelinev2.TriggerManualTrigger)
	model.Type = core.StringPtr("testString")
	model.Name = core.StringPtr("start-deploy")
	model.Href = core.StringPtr("testString")
	model.EventListener = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.Properties = []cdtektonpipelinev2.TriggerProperty{*triggerPropertyModel}
	model.Tags = []string{"testString"}
	model.Worker = workerModel
	model.MaxConcurrentRuns = core.Int64Ptr(int64(4))
	model.Enabled = core.BoolPtr(true)
	model.Favorite = core.BoolPtr(false)
	model.LimitWaitingRuns = core.BoolPtr(false)

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerManualTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerScmTriggerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		triggerPropertyModel := make(map[string]interface{})
		triggerPropertyModel["name"] = "testString"
		triggerPropertyModel["value"] = "testString"
		triggerPropertyModel["href"] = "testString"
		triggerPropertyModel["enum"] = []string{"testString"}
		triggerPropertyModel["type"] = "secure"
		triggerPropertyModel["path"] = "testString"
		triggerPropertyModel["locked"] = true

		workerModel := make(map[string]interface{})
		workerModel["name"] = "testString"
		workerModel["type"] = "testString"
		workerModel["id"] = "testString"

		toolModel := make(map[string]interface{})
		toolModel["id"] = "testString"

		triggerSourcePropertiesModel := make(map[string]interface{})
		triggerSourcePropertiesModel["url"] = "testString"
		triggerSourcePropertiesModel["branch"] = "testString"
		triggerSourcePropertiesModel["pattern"] = "testString"
		triggerSourcePropertiesModel["blind_connection"] = true
		triggerSourcePropertiesModel["hook_id"] = "testString"
		triggerSourcePropertiesModel["tool"] = []map[string]interface{}{toolModel}

		triggerSourceModel := make(map[string]interface{})
		triggerSourceModel["type"] = "testString"
		triggerSourceModel["properties"] = []map[string]interface{}{triggerSourcePropertiesModel}

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["name"] = "start-deploy"
		model["href"] = "testString"
		model["event_listener"] = "testString"
		model["id"] = "testString"
		model["properties"] = []map[string]interface{}{triggerPropertyModel}
		model["tags"] = []string{"testString"}
		model["worker"] = []map[string]interface{}{workerModel}
		model["max_concurrent_runs"] = int(4)
		model["enabled"] = true
		model["favorite"] = false
		model["limit_waiting_runs"] = false
		model["enable_events_from_forks"] = false
		model["disable_draft_events"] = false
		model["source"] = []map[string]interface{}{triggerSourceModel}
		model["events"] = []string{"push", "pull_request"}
		model["filter"] = "header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'"

		assert.Equal(t, result, model)
	}

	triggerPropertyModel := new(cdtektonpipelinev2.TriggerProperty)
	triggerPropertyModel.Name = core.StringPtr("testString")
	triggerPropertyModel.Value = core.StringPtr("testString")
	triggerPropertyModel.Href = core.StringPtr("testString")
	triggerPropertyModel.Enum = []string{"testString"}
	triggerPropertyModel.Type = core.StringPtr("secure")
	triggerPropertyModel.Path = core.StringPtr("testString")
	triggerPropertyModel.Locked = core.BoolPtr(true)

	workerModel := new(cdtektonpipelinev2.Worker)
	workerModel.Name = core.StringPtr("testString")
	workerModel.Type = core.StringPtr("testString")
	workerModel.ID = core.StringPtr("testString")

	toolModel := new(cdtektonpipelinev2.Tool)
	toolModel.ID = core.StringPtr("testString")

	triggerSourcePropertiesModel := new(cdtektonpipelinev2.TriggerSourceProperties)
	triggerSourcePropertiesModel.URL = core.StringPtr("testString")
	triggerSourcePropertiesModel.Branch = core.StringPtr("testString")
	triggerSourcePropertiesModel.Pattern = core.StringPtr("testString")
	triggerSourcePropertiesModel.BlindConnection = core.BoolPtr(true)
	triggerSourcePropertiesModel.HookID = core.StringPtr("testString")
	triggerSourcePropertiesModel.Tool = toolModel

	triggerSourceModel := new(cdtektonpipelinev2.TriggerSource)
	triggerSourceModel.Type = core.StringPtr("testString")
	triggerSourceModel.Properties = triggerSourcePropertiesModel

	model := new(cdtektonpipelinev2.TriggerScmTrigger)
	model.Type = core.StringPtr("testString")
	model.Name = core.StringPtr("start-deploy")
	model.Href = core.StringPtr("testString")
	model.EventListener = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.Properties = []cdtektonpipelinev2.TriggerProperty{*triggerPropertyModel}
	model.Tags = []string{"testString"}
	model.Worker = workerModel
	model.MaxConcurrentRuns = core.Int64Ptr(int64(4))
	model.Enabled = core.BoolPtr(true)
	model.Favorite = core.BoolPtr(false)
	model.LimitWaitingRuns = core.BoolPtr(false)
	model.EnableEventsFromForks = core.BoolPtr(false)
	model.DisableDraftEvents = core.BoolPtr(false)
	model.Source = triggerSourceModel
	model.Events = []string{"push", "pull_request"}
	model.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerScmTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerTimerTriggerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		triggerPropertyModel := make(map[string]interface{})
		triggerPropertyModel["name"] = "testString"
		triggerPropertyModel["value"] = "testString"
		triggerPropertyModel["href"] = "testString"
		triggerPropertyModel["enum"] = []string{"testString"}
		triggerPropertyModel["type"] = "secure"
		triggerPropertyModel["path"] = "testString"
		triggerPropertyModel["locked"] = true

		workerModel := make(map[string]interface{})
		workerModel["name"] = "testString"
		workerModel["type"] = "testString"
		workerModel["id"] = "testString"

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["name"] = "start-deploy"
		model["href"] = "testString"
		model["event_listener"] = "testString"
		model["id"] = "testString"
		model["properties"] = []map[string]interface{}{triggerPropertyModel}
		model["tags"] = []string{"testString"}
		model["worker"] = []map[string]interface{}{workerModel}
		model["max_concurrent_runs"] = int(4)
		model["enabled"] = true
		model["favorite"] = false
		model["limit_waiting_runs"] = false
		model["cron"] = "testString"
		model["timezone"] = "America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC"

		assert.Equal(t, result, model)
	}

	triggerPropertyModel := new(cdtektonpipelinev2.TriggerProperty)
	triggerPropertyModel.Name = core.StringPtr("testString")
	triggerPropertyModel.Value = core.StringPtr("testString")
	triggerPropertyModel.Href = core.StringPtr("testString")
	triggerPropertyModel.Enum = []string{"testString"}
	triggerPropertyModel.Type = core.StringPtr("secure")
	triggerPropertyModel.Path = core.StringPtr("testString")
	triggerPropertyModel.Locked = core.BoolPtr(true)

	workerModel := new(cdtektonpipelinev2.Worker)
	workerModel.Name = core.StringPtr("testString")
	workerModel.Type = core.StringPtr("testString")
	workerModel.ID = core.StringPtr("testString")

	model := new(cdtektonpipelinev2.TriggerTimerTrigger)
	model.Type = core.StringPtr("testString")
	model.Name = core.StringPtr("start-deploy")
	model.Href = core.StringPtr("testString")
	model.EventListener = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.Properties = []cdtektonpipelinev2.TriggerProperty{*triggerPropertyModel}
	model.Tags = []string{"testString"}
	model.Worker = workerModel
	model.MaxConcurrentRuns = core.Int64Ptr(int64(4))
	model.Enabled = core.BoolPtr(true)
	model.Favorite = core.BoolPtr(false)
	model.LimitWaitingRuns = core.BoolPtr(false)
	model.Cron = core.StringPtr("testString")
	model.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerTimerTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineTriggerGenericTriggerToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		triggerPropertyModel := make(map[string]interface{})
		triggerPropertyModel["name"] = "testString"
		triggerPropertyModel["value"] = "testString"
		triggerPropertyModel["href"] = "testString"
		triggerPropertyModel["enum"] = []string{"testString"}
		triggerPropertyModel["type"] = "secure"
		triggerPropertyModel["path"] = "testString"
		triggerPropertyModel["locked"] = true

		workerModel := make(map[string]interface{})
		workerModel["name"] = "testString"
		workerModel["type"] = "testString"
		workerModel["id"] = "testString"

		genericSecretModel := make(map[string]interface{})
		genericSecretModel["type"] = "token_matches"
		genericSecretModel["value"] = "testString"
		genericSecretModel["source"] = "header"
		genericSecretModel["key_name"] = "testString"
		genericSecretModel["algorithm"] = "md4"

		model := make(map[string]interface{})
		model["type"] = "testString"
		model["name"] = "start-deploy"
		model["href"] = "testString"
		model["event_listener"] = "testString"
		model["id"] = "testString"
		model["properties"] = []map[string]interface{}{triggerPropertyModel}
		model["tags"] = []string{"testString"}
		model["worker"] = []map[string]interface{}{workerModel}
		model["max_concurrent_runs"] = int(4)
		model["enabled"] = true
		model["favorite"] = false
		model["limit_waiting_runs"] = false
		model["secret"] = []map[string]interface{}{genericSecretModel}
		model["webhook_url"] = "testString"
		model["filter"] = "event.type == 'message' && event.text.contains('urgent')"

		assert.Equal(t, result, model)
	}

	triggerPropertyModel := new(cdtektonpipelinev2.TriggerProperty)
	triggerPropertyModel.Name = core.StringPtr("testString")
	triggerPropertyModel.Value = core.StringPtr("testString")
	triggerPropertyModel.Href = core.StringPtr("testString")
	triggerPropertyModel.Enum = []string{"testString"}
	triggerPropertyModel.Type = core.StringPtr("secure")
	triggerPropertyModel.Path = core.StringPtr("testString")
	triggerPropertyModel.Locked = core.BoolPtr(true)

	workerModel := new(cdtektonpipelinev2.Worker)
	workerModel.Name = core.StringPtr("testString")
	workerModel.Type = core.StringPtr("testString")
	workerModel.ID = core.StringPtr("testString")

	genericSecretModel := new(cdtektonpipelinev2.GenericSecret)
	genericSecretModel.Type = core.StringPtr("token_matches")
	genericSecretModel.Value = core.StringPtr("testString")
	genericSecretModel.Source = core.StringPtr("header")
	genericSecretModel.KeyName = core.StringPtr("testString")
	genericSecretModel.Algorithm = core.StringPtr("md4")

	model := new(cdtektonpipelinev2.TriggerGenericTrigger)
	model.Type = core.StringPtr("testString")
	model.Name = core.StringPtr("start-deploy")
	model.Href = core.StringPtr("testString")
	model.EventListener = core.StringPtr("testString")
	model.ID = core.StringPtr("testString")
	model.Properties = []cdtektonpipelinev2.TriggerProperty{*triggerPropertyModel}
	model.Tags = []string{"testString"}
	model.Worker = workerModel
	model.MaxConcurrentRuns = core.Int64Ptr(int64(4))
	model.Enabled = core.BoolPtr(true)
	model.Favorite = core.BoolPtr(false)
	model.LimitWaitingRuns = core.BoolPtr(false)
	model.Secret = genericSecretModel
	model.WebhookURL = core.StringPtr("testString")
	model.Filter = core.StringPtr("event.type == 'message' && event.text.contains('urgent')")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerGenericTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
