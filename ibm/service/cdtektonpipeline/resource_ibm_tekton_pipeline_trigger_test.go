// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/org-ids/tekton-pipeline-go-sdk/cdtektonpipelinev2"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTektonPipelineTriggerBasic(t *testing.T) {
	var conf cdtektonpipelinev2.Trigger
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelineTriggerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerConfigBasic(pipelineID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelineTriggerExists("ibm_tekton_pipeline_trigger.tekton_pipeline_trigger", conf),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger.tekton_pipeline_trigger", "pipeline_id", pipelineID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tekton_pipeline_trigger.tekton_pipeline_trigger",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelineTriggerConfigBasic(pipelineID string) string {
	return fmt.Sprintf(`

		resource "ibm_tekton_pipeline_trigger" "tekton_pipeline_trigger" {
			pipeline_id = "%s"
		}
	`, pipelineID)
}

func testAccCheckIBMTektonPipelineTriggerExists(n string, obj cdtektonpipelinev2.Trigger) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerOptions.SetTriggerID(parts[1])

		_, response, err := cdTektonPipelineClient.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)
		if err != nil {
			return err
		}

		if response.StatusCode != 200 {
			return fmt.Errorf("Error checking for tekton_pipeline_trigger (%s) has been fetched: %s", rs.Primary.ID, err)
		}
		return nil
	}
}

func testAccCheckIBMTektonPipelineTriggerDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tekton_pipeline_trigger" {
			continue
		}

		getTektonPipelineTriggerOptions := &cdtektonpipelinev2.GetTektonPipelineTriggerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerOptions.SetTriggerID(parts[1])

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipelineTrigger(getTektonPipelineTriggerOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline_trigger still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline_trigger (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
