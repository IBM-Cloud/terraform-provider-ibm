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

func TestAccIBMTektonPipelineDefinitionDataSourceBasic(t *testing.T) {
	definitionPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineDefinitionDataSourceConfigBasic(definitionPipelineID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_definition.tekton_pipeline_definition", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_definition.tekton_pipeline_definition", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_definition.tekton_pipeline_definition", "definition_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_definition.tekton_pipeline_definition", "scm_source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_definition.tekton_pipeline_definition", "service_instance_id"),
				),
			},
		},
	})
}

func testAccCheckIBMTektonPipelineDefinitionDataSourceConfigBasic(definitionPipelineID string) string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline_definition" "tekton_pipeline_definition" {
			pipeline_id = "%s"
		}

		data "ibm_tekton_pipeline_definition" "tekton_pipeline_definition" {
			pipeline_id = ibm_tekton_pipeline_definition.tekton_pipeline_definition.pipeline_id
			definition_id = ibm_tekton_pipeline_definition.tekton_pipeline_definition.definition_id
		}
	`, definitionPipelineID)
}
