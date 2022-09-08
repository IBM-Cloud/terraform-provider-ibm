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

func TestAccIBMCdTektonPipelineDefinitionDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDefinitionDataSourceConfigBasic(""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "definition_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "scm_source.#"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineDefinitionDataSourceConfigBasic(definitionPipelineID string) string {
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
				ui_pipeline = true
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
		data "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition" {
			pipeline_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition.pipeline_id
			definition_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition.definition_id
		}
	`, tcName, rgID)
}
