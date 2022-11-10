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

func TestAccIBMCdTektonPipelineTriggerPropertyDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfigBasic("", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "type"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineTriggerPropertyDataSourceAllArgs(t *testing.T) {
	triggerPropertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	triggerPropertyValue := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	triggerPropertyType := "text"
	triggerPropertyPath := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfig("", "", triggerPropertyName, triggerPropertyValue, triggerPropertyType, triggerPropertyPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "type"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfigBasic(triggerPropertyPipelineID string, triggerPropertyTriggerID string) string {
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
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			type = "manual"
			event_listener = "listener"
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			name = "trigger"
		}
		resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
			type = "text"
			name = "trig-prop-1"
			value = "trig-prop-value-1"
		}
		data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property.trigger_id
			property_name = "trig-prop-1"
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfig(triggerPropertyPipelineID string, triggerPropertyTriggerID string, triggerPropertyName string, triggerPropertyValue string, triggerPropertyType string, triggerPropertyPath string) string {
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
		resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			type = "manual"
			event_listener = "listener"
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			name = "trigger"
		}
		resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
			name = "%s"
			value = "%s"
			type = "%s"
		}
		data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property.trigger_id
			property_name = "%s"
		}
	`, rgName, tcName, triggerPropertyName, triggerPropertyValue, triggerPropertyType, triggerPropertyName)
}
