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

func TestAccIBMCdTektonPipelinePropertyDataSourceBasic(t *testing.T) {
	propertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyType := "secure"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyDataSourceConfigBasic("", propertyName, propertyType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "type"),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelinePropertyDataSourceAllArgs(t *testing.T) {
	propertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyValue := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	propertyType := "text"
	propertyLocked := "true"
	propertyPath := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelinePropertyDataSourceConfig("", propertyName, propertyValue, propertyType, propertyLocked, propertyPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance", "locked"),
				),
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelinePropertyDataSourceConfigBasic(propertyPipelineID string, propertyName string, propertyType string) string {
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
		data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			property_name = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance.name
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelinePropertyDataSourceConfig(propertyPipelineID string, propertyName string, propertyValue string, propertyType string, propertyLocked string, propertyPath string) string {
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
		data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
			pipeline_id = ibm_cd_tekton_pipeline.cd_tekton_pipeline_instance.pipeline_id
			property_name = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance.name
		}
	`, rgName, tcName, propertyName, propertyValue, propertyLocked)
}
