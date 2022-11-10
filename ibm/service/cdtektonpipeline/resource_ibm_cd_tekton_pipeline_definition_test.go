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
					testAccCheckIBMCdTektonPipelineDefinitionExists("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", conf),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "definition_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "pipeline_id"),
					resource.TestCheckResourceAttrSet("ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition", "source.#"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition",
				ImportState:       true,
				ImportStateVerify: true,
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
