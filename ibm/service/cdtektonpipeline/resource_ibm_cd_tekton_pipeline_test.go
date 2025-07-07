// Copyright IBM Corp. 2025 All Rights Reserved.
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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtektonpipeline"
	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCdTektonPipelineBasic(t *testing.T) {
	var conf cdtektonpipelinev2.TektonPipeline

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineExists("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", conf),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineAllArgs(t *testing.T) {
	var conf cdtektonpipelinev2.TektonPipeline
	nextBuildNumber := "5"
	enableNotifications := "true"
	enablePartialCloning := "true"
	nextBuildNumberUpdate := fmt.Sprintf("%d", acctest.RandIntRange(1, 99999999999999))
	enableNotificationsUpdate := "false"
	enablePartialCloningUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineConfig(nextBuildNumber, enableNotifications, enablePartialCloning),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineExists("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "next_build_number", nextBuildNumber),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_notifications", enableNotifications),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_partial_cloning", enablePartialCloning),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineConfig(nextBuildNumberUpdate, enableNotificationsUpdate, enablePartialCloningUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "next_build_number", nextBuildNumberUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_notifications", enableNotificationsUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance", "enable_partial_cloning", enablePartialCloningUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineConfigBasic() string {
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
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineConfig(nextBuildNumber string, enableNotifications string, enablePartialCloning string) string {
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
	`, rgName, tcName, nextBuildNumber, enableNotifications, enablePartialCloning)
}

func testAccCheckIBMCdTektonPipelineExists(n string, obj cdtektonpipelinev2.TektonPipeline) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{}

		getTektonPipelineOptions.SetID(rs.Primary.ID)

		tektonPipeline, _, err := cdTektonPipelineClient.GetTektonPipeline(getTektonPipelineOptions)
		if err != nil {
			return err
		}

		obj = *tektonPipeline
		return nil
	}
}

func testAccCheckIBMCdTektonPipelineDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline" {
			continue
		}

		getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{}

		getTektonPipelineOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipeline(getTektonPipelineOptions)

		if err == nil {
			return fmt.Errorf("cd_tekton_pipeline still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_tekton_pipeline (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMCdTektonPipelineWorkerToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineWorkerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineResourceGroupReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.ResourceGroupReference)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineResourceGroupReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineToolchainReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["crn"] = "crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.ToolchainReference)
	model.ID = core.StringPtr("testString")
	model.CRN = core.StringPtr("crn:v1:staging:public:toolchain:us-south:a/0ba224679d6c697f9baee5e14ade83ac:bf5fa00f-ddef-4298-b87b-aa8b6da0e1a6::")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineToolchainReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionSourceToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionSourcePropertiesToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineToolToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Tool)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineToolToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelinePropertyToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelinePropertyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerToMap(t *testing.T) {
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
		model["enable_events_from_forks"] = false
		model["source"] = []map[string]interface{}{triggerSourceModel}
		model["events"] = []string{"push", "pull_request"}
		model["filter"] = "header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'"
		model["limit_waiting_runs"] = false
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
	model.EnableEventsFromForks = core.BoolPtr(false)
	model.Source = triggerSourceModel
	model.Events = []string{"push", "pull_request"}
	model.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
	model.LimitWaitingRuns = core.BoolPtr(false)
	model.Cron = core.StringPtr("testString")
	model.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")
	model.Secret = genericSecretModel
	model.WebhookURL = core.StringPtr("testString")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerPropertyToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerPropertyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerSourceToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerSourcePropertiesToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineGenericSecretToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineGenericSecretToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerManualTriggerToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerManualTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerScmTriggerToMap(t *testing.T) {
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
		model["enable_events_from_forks"] = false
		model["source"] = []map[string]interface{}{triggerSourceModel}
		model["events"] = []string{"push", "pull_request"}
		model["filter"] = "header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'"
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
	model.EnableEventsFromForks = core.BoolPtr(false)
	model.Source = triggerSourceModel
	model.Events = []string{"push", "pull_request"}
	model.Filter = core.StringPtr("header['x-github-event'] == 'push' && body.ref == 'refs/heads/main'")
	model.LimitWaitingRuns = core.BoolPtr(false)

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerScmTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerTimerTriggerToMap(t *testing.T) {
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
	model.Cron = core.StringPtr("testString")
	model.Timezone = core.StringPtr("America/Los_Angeles, CET, Europe/London, GMT, US/Eastern, or UTC")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerTimerTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineTriggerGenericTriggerToMap(t *testing.T) {
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
	model.Secret = genericSecretModel
	model.WebhookURL = core.StringPtr("testString")
	model.Filter = core.StringPtr("event.type == 'message' && event.text.contains('urgent')")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerGenericTriggerToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineMapToWorkerIdentity(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.WorkerIdentity) {
		model := new(cdtektonpipelinev2.WorkerIdentity)
		model.ID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineMapToWorkerIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}
