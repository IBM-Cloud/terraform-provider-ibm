// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/continuous-delivery-go-sdk/v2/cdtektonpipelinev2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMCdTektonPipelinePropertyBasic(t *testing.T) {
	var conf cdtektonpipelinev2.Property
	name := "property1"
	typeVar := "text"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelinePropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyConfigBasic("", name, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelinePropertyExists("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "type", typeVar),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelinePropertyAllArgs(t *testing.T) {
	var conf cdtektonpipelinev2.Property
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	value := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	typeVar := "text"
	locked := "true"
	path := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))
	valueUpdate := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	lockedUpdate := "false"
	pathUpdate := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelinePropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyConfig("", name, value, typeVar, locked, path),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelinePropertyExists("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "value", value),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "locked", locked),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyConfig("", name, valueUpdate, typeVar, lockedUpdate, pathUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "value", valueUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "locked", lockedUpdate),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pipeline_id"},
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelinePropertyConfigBasic(pipelineID string, name string, typeVar string) string {
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
		resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			name = "property1"
			type = "text"
			value = "prop1"
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance
			]
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelinePropertyConfig(pipelineID string, name string, value string, typeVar string, locked string, path string) string {
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
		resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			name = "%s"
			type = "text"
			value = "%s"
			locked = "%s"
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance
			]
		}
	`, rgName, tcName, name, value, locked)
}

func testAccCheckIBMCdTektonPipelinePropertyExists(n string, obj cdtektonpipelinev2.Property) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions := &cdtektonpipelinev2.GetTektonPipelinePropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
		getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

		property, _, err := cdTektonPipelineClient.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)
		if err != nil {
			return err
		}

		obj = *property
		return nil
	}
}

func testAccCheckIBMCdTektonPipelinePropertyDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline_property" {
			continue
		}

		getTektonPipelinePropertyOptions := &cdtektonpipelinev2.GetTektonPipelinePropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
		getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)

		if err == nil {
			return fmt.Errorf("cd_tekton_pipeline_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_tekton_pipeline_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
