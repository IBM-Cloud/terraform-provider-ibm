// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package continuousdeliverypipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTektonPipelinePropertyBasic(t *testing.T) {
	var conf continuousdeliverypipelinev2.Property
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelinePropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelinePropertyConfigBasic(pipelineID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelinePropertyExists("ibm_tekton_pipeline_property.tekton_pipeline_property", conf),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_property.tekton_pipeline_property", "pipeline_id", pipelineID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tekton_pipeline_property.tekton_pipeline_property",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelinePropertyConfigBasic(pipelineID string) string {
	return fmt.Sprintf(`

		resource "ibm_tekton_pipeline_property" "tekton_pipeline_property" {
			pipeline_id = "%s"
		}
	`, pipelineID)
}

func testAccCheckIBMTektonPipelinePropertyExists(n string, obj continuousdeliverypipelinev2.Property) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions := &continuousdeliverypipelinev2.GetTektonPipelinePropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
		getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

		property, _, err := continuousDeliveryPipelineClient.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)
		if err != nil {
			return err
		}

		obj = *property
		return nil
	}
}

func testAccCheckIBMTektonPipelinePropertyDestroy(s *terraform.State) error {
	continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tekton_pipeline_property" {
			continue
		}

		getTektonPipelinePropertyOptions := &continuousdeliverypipelinev2.GetTektonPipelinePropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelinePropertyOptions.SetPipelineID(parts[0])
		getTektonPipelinePropertyOptions.SetPropertyName(parts[1])

		// Try to find the key
		_, response, err := continuousDeliveryPipelineClient.GetTektonPipelineProperty(getTektonPipelinePropertyOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
