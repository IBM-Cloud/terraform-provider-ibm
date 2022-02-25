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
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/continuousdeliverypipelinev2"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTektonPipelineBasic(t *testing.T) {
	var conf continuousdeliverypipelinev2.TektonPipeline

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelineExists("ibm_tekton_pipeline.tekton_pipeline", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tekton_pipeline.tekton_pipeline",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelineConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_tekton_pipeline" "tekton_pipeline" {
		}
	`, )
}

func testAccCheckIBMTektonPipelineExists(n string, obj continuousdeliverypipelinev2.TektonPipeline) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineOptions := &continuousdeliverypipelinev2.GetTektonPipelineOptions{}

		getTektonPipelineOptions.SetID(rs.Primary.ID)

		tektonPipeline, _, err := continuousDeliveryPipelineClient.GetTektonPipeline(getTektonPipelineOptions)
		if err != nil {
			return err
		}

		obj = *tektonPipeline
		return nil
	}
}

func testAccCheckIBMTektonPipelineDestroy(s *terraform.State) error {
	continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tekton_pipeline" {
			continue
		}

		getTektonPipelineOptions := &continuousdeliverypipelinev2.GetTektonPipelineOptions{}

		getTektonPipelineOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := continuousDeliveryPipelineClient.GetTektonPipeline(getTektonPipelineOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
