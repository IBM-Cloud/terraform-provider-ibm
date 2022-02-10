// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package continuousdeliverypipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTektonPipelineTriggerDataSourceBasic(t *testing.T) {
	triggerPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerDataSourceConfigBasic(triggerPipelineID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger.tekton_pipeline_trigger", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger.tekton_pipeline_trigger", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger.tekton_pipeline_trigger", "trigger_id"),
				),
			},
		},
	})
}

func testAccCheckIBMTektonPipelineTriggerDataSourceConfigBasic(triggerPipelineID string) string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline_trigger" "tekton_pipeline_trigger" {
			pipeline_id = "%s"
		}

		data "ibm_tekton_pipeline_trigger" "tekton_pipeline_trigger" {
			pipeline_id = ibm_tekton_pipeline_trigger.tekton_pipeline_trigger.pipeline_id
			trigger_id = ibm_tekton_pipeline_trigger.tekton_pipeline_trigger.trigger_id
		}
	`, triggerPipelineID)
}
