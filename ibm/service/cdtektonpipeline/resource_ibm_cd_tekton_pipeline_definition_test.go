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

func TestAccIBMCdTektonPipelineDefinitionBasic(t *testing.T) {
	var conf cdtektonpipelinev2.Definition

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDefinitionConfigBasic(""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineDefinitionExists("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "definition_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "source.#"),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pipeline_id"},
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineDefinitionConfigBasic(pipelineID string) string {
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
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineDefinitionExists(n string, obj cdtektonpipelinev2.Definition) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineDefinitionOptions := &cdtektonpipelinev2.GetTektonPipelineDefinitionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
		getTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])

		definition, _, err := cdTektonPipelineClient.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)
		if err != nil {
			return err
		}

		obj = *definition
		return nil
	}
}

func testAccCheckIBMCdTektonPipelineDefinitionDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline_definition" {
			continue
		}

		getTektonPipelineDefinitionOptions := &cdtektonpipelinev2.GetTektonPipelineDefinitionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
		getTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)

		if err == nil {
			return fmt.Errorf("cd_tekton_pipeline_definition still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_tekton_pipeline_definition (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMCdTektonPipelineDefinitionDefinitionSourceToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionDefinitionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionDefinitionSourcePropertiesToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionDefinitionSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionToolToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Tool)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionToolToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionMapToDefinitionSource(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.DefinitionSource) {
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

		assert.Equal(t, result, model)
	}

	toolModel := make(map[string]interface{})
	toolModel["id"] = "testString"

	definitionSourcePropertiesModel := make(map[string]interface{})
	definitionSourcePropertiesModel["url"] = "testString"
	definitionSourcePropertiesModel["branch"] = "testString"
	definitionSourcePropertiesModel["tag"] = "testString"
	definitionSourcePropertiesModel["path"] = "testString"
	definitionSourcePropertiesModel["tool"] = []interface{}{toolModel}

	model := make(map[string]interface{})
	model["type"] = "testString"
	model["properties"] = []interface{}{definitionSourcePropertiesModel}

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionMapToDefinitionSource(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionMapToDefinitionSourceProperties(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.DefinitionSourceProperties) {
		toolModel := new(cdtektonpipelinev2.Tool)
		toolModel.ID = core.StringPtr("testString")

		model := new(cdtektonpipelinev2.DefinitionSourceProperties)
		model.URL = core.StringPtr("testString")
		model.Branch = core.StringPtr("testString")
		model.Tag = core.StringPtr("testString")
		model.Path = core.StringPtr("testString")
		model.Tool = toolModel

		assert.Equal(t, result, model)
	}

	toolModel := make(map[string]interface{})
	toolModel["id"] = "testString"

	model := make(map[string]interface{})
	model["url"] = "testString"
	model["branch"] = "testString"
	model["tag"] = "testString"
	model["path"] = "testString"
	model["tool"] = []interface{}{toolModel}

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionMapToDefinitionSourceProperties(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMCdTektonPipelineDefinitionMapToTool(t *testing.T) {
	checkResult := func(result *cdtektonpipelinev2.Tool) {
		model := new(cdtektonpipelinev2.Tool)
		model.ID = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"

	result, err := cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionMapToTool(model)
	assert.Nil(t, err)
	checkResult(result)
}
