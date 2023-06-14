// Copyright IBM Corp. 2023 All Rights Reserved.
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

func TestAccIBMCdTektonPipelineTriggerPropertyBasic(t *testing.T) {
	var conf cdtektonpipelinev2.TriggerProperty
	name := "trig-prop-1"
	typeVar := "text"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineTriggerPropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyConfigBasic("", "", name, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineTriggerPropertyExists("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "type", typeVar),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineTriggerPropertyAllArgs(t *testing.T) {
	var conf cdtektonpipelinev2.TriggerProperty
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVar := "text"
	value := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	path := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))
	valueUpdate := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	pathUpdate := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineTriggerPropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyConfig("", "", name, typeVar, value, path),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineTriggerPropertyExists("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "value", value),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineTriggerPropertyConfig("", "", name, typeVar, valueUpdate, pathUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property", "value", valueUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyConfigBasic(pipelineID string, triggerID string, name string, typeVar string) string {
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
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			name = "trigger"
			type = "manual"
			event_listener = "listener"
		}
		resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
			type = "text"
			name = "trig-prop-1"
			value = "trig-prop-value-1"
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyConfig(pipelineID string, triggerID string, name string, typeVar string, value string, path string) string {
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
			depends_on = [
				ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition
			]
			name = "trigger"
			type = "manual"
			event_listener = "listener"
		}
		resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger.trigger_id
			name = "%s"
			type = "%s"
			value = "%s"
		}
	`, rgName, tcName, name, typeVar, value)
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyExists(n string, obj cdtektonpipelinev2.TriggerProperty) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
		getTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

		triggerProperty, _, err := cdTektonPipelineClient.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)
		if err != nil {
			return err
		}

		obj = *triggerProperty
		return nil
	}
}

func testAccCheckIBMCdTektonPipelineTriggerPropertyDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline_trigger_property" {
			continue
		}

		getTektonPipelineTriggerPropertyOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerPropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
		getTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)

		if err == nil {
			return fmt.Errorf("cd_tekton_pipeline_trigger_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_tekton_pipeline_trigger_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
