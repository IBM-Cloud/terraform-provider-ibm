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
)

func TestAccIBMCdTektonPipelineTriggerPropertyDataSourceBasic(t *testing.T) {
	triggerPropertyName := "trig-prop-1"
	triggerPropertyType := "text"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfigBasic("", "", triggerPropertyName, triggerPropertyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "type"),
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
	triggerPropertyLocked := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfig("", "", triggerPropertyName, triggerPropertyValue, triggerPropertyType, triggerPropertyPath, triggerPropertyLocked),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance", "locked"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfigBasic(triggerPropertyPipelineID string, triggerPropertyTriggerID string, triggerPropertyName string, triggerPropertyType string) string {
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
			name = "trigger"
			type = "manual"
			event_listener = "listener"
		}
		resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
			type = "text"
			name = "trig-prop-1"
			value = "trig-prop-value-1"
		}
		data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.trigger_id
			property_name = "trig-prop-1"
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyDataSourceConfig(triggerPropertyPipelineID string, triggerPropertyTriggerID string, triggerPropertyName string, triggerPropertyValue string, triggerPropertyType string, triggerPropertyPath string, triggerPropertyLocked string) string {
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
			name = "trigger"
			type = "manual"
			event_listener = "listener"
		}
		resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
			name = "%s"
			type = "%s"
			value = "%s"
			locked = "%s"
		}
		data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.trigger_id
			property_name = "%s"
		}
	`, rgName, tcName, triggerPropertyName, triggerPropertyType, triggerPropertyValue, triggerPropertyLocked, triggerPropertyName)
}
