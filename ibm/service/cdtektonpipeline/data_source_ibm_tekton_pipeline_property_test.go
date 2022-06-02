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

func TestAccIBMTektonPipelinePropertyDataSourceBasic(t *testing.T) {
	propertyPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelinePropertyDataSourceConfigBasic(propertyPipelineID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "type"),
				),
			},
		},
	})
}

func TestAccIBMTektonPipelinePropertyDataSourceAllArgs(t *testing.T) {
	propertyPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	propertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	propertyValue := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	propertyDefault := fmt.Sprintf("tf_default_%d", acctest.RandIntRange(10, 100))
	propertyType := "SECURE"
	propertyPath := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelinePropertyDataSourceConfig(propertyPipelineID, propertyName, propertyValue, propertyDefault, propertyType, propertyPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "enum.#"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "default"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_tekton_pipeline_property.tekton_pipeline_property", "path"),
				),
			},
		},
	})
}

func testAccCheckIBMTektonPipelinePropertyDataSourceConfigBasic(propertyPipelineID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = "%s"
		}

		data "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline_property.tekton_pipeline_property.pipeline_id
			property_name = "debug-pipeline"
		}
	`, propertyPipelineID)
}

func testAccCheckIBMTektonPipelinePropertyDataSourceConfig(propertyPipelineID string, propertyName string, propertyValue string, propertyDefault string, propertyType string, propertyPath string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = "%s"
			name = "%s"
			value = "%s"
			enum = "FIXME"
			default = "%s"
			type = "%s"
			path = "%s"
		}

		data "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = ibm_cd_tekton_pipeline_property.tekton_pipeline_property.pipeline_id
			property_name = "debug-pipeline"
		}
	`, propertyPipelineID, propertyName, propertyValue, propertyDefault, propertyType, propertyPath)
}
