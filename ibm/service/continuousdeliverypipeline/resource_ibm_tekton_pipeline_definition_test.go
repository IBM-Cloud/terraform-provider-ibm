// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package continuousdeliverypipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"
)

func TestAccIBMTektonPipelineDefinitionBasic(t *testing.T) {
	var conf continuousdeliverypipelinev2.Definition
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelineDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineDefinitionConfigBasic(pipelineID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelineDefinitionExists("ibm_tekton_pipeline_definition.tekton_pipeline_definition", conf),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_definition.tekton_pipeline_definition", "pipeline_id", pipelineID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tekton_pipeline_definition.tekton_pipeline_definition",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelineDefinitionConfigBasic(pipelineID string) string {
	return fmt.Sprintf(`

		resource "ibm_tekton_pipeline_definition" "tekton_pipeline_definition" {
			pipeline_id = "%s"
		}
	`, pipelineID)
}

func testAccCheckIBMTektonPipelineDefinitionExists(n string, obj continuousdeliverypipelinev2.Definition) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineDefinitionOptions := &continuousdeliverypipelinev2.GetTektonPipelineDefinitionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
		getTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])

		definition, _, err := continuousDeliveryPipelineClient.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)
		if err != nil {
			return err
		}

		obj = *definition
		return nil
	}
}

func testAccCheckIBMTektonPipelineDefinitionDestroy(s *terraform.State) error {
	continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tekton_pipeline_definition" {
			continue
		}

		getTektonPipelineDefinitionOptions := &continuousdeliverypipelinev2.GetTektonPipelineDefinitionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineDefinitionOptions.SetPipelineID(parts[0])
		getTektonPipelineDefinitionOptions.SetDefinitionID(parts[1])

		// Try to find the key
		_, response, err := continuousDeliveryPipelineClient.GetTektonPipelineDefinition(getTektonPipelineDefinitionOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline_definition still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline_definition (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
