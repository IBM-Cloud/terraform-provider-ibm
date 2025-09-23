// Copyright IBM Corp. 2025 All Rights Reserved.
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

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	"github.com/IBM/go-sdk-core/v5/core"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/cdtektonpipeline"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMCdTektonPipelineDefinitionDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDefinitionDataSourceConfigBasic(""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "definition_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance", "source.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineDefinitionDataSourceConfigBasic(definitionPipelineID string) string {
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
		data "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
			pipeline_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance.pipeline_id
			definition_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance.definition_id
		}
	`, rgName, tcName)
}

func TestDataSourceIBMCdTektonPipelineDefinitionDefinitionSourceToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinitionDefinitionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineDefinitionDefinitionSourcePropertiesToMap(t *testing.T) {
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

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinitionDefinitionSourcePropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMCdTektonPipelineDefinitionToolToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(cdtektonpipelinev2.Tool)
	model.ID = core.StringPtr("testString")

	result, err := cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinitionToolToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
