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

func TestAccIBMCdTektonPipelineDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "runs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enabled"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineDataSourceAllArgs(t *testing.T) {
	tektonPipelineEnableNotifications := "true"
	tektonPipelineEnablePartialCloning := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineDataSourceConfig(tektonPipelineEnableNotifications, tektonPipelineEnablePartialCloning),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "runs_url"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_notifications"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_partial_cloning"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enabled"),
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
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
		data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineDataSourceConfig(tektonPipelineEnableNotifications string, tektonPipelineEnablePartialCloning string) string {
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
		resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			name = "property1"
			type = "text"
			value = "prop1"
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
		}
		data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
		}
	`, rgName, tcName, tektonPipelineEnableNotifications, tektonPipelineEnablePartialCloning)
}
