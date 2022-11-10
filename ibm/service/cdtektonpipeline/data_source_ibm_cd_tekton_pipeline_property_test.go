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

func TestAccIBMCdTektonPipelinePropertyDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyDataSourceConfigBasic(""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "type"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelinePropertyDataSourceAllArgs(t *testing.T) {
	propertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyValue := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	propertyType := "text"
	propertyPath := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyDataSourceConfig("", propertyName, propertyValue, propertyType, propertyPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property", "type"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelinePropertyDataSourceConfigBasic(propertyPipelineID string) string {
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
		resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			name = "property1"
			type = "text"
			value = "prop1"
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
		}
		data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			property_name = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property.name
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelinePropertyDataSourceConfig(propertyPipelineID string, propertyName string, propertyValue string, propertyType string, propertyPath string) string {
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
		resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			name = "%s"
			type = "text"
			value = "%s"
			depends_on = [
				ibm_cd_tekton_pipeline.cd_tekton_pipeline
			]
		}
		data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline.pipeline_id
			property_name = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property.name
		}
	`, rgName, tcName, propertyName, propertyValue)
}
