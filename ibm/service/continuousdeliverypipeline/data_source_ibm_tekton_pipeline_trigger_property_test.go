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

func TestAccIBMTektonPipelineTriggerPropertyDataSourceBasic(t *testing.T) {
	triggerPropertyPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	triggerPropertyTriggerID := fmt.Sprintf("tf_trigger_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerPropertyDataSourceConfigBasic(triggerPropertyPipelineID, triggerPropertyTriggerID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "trigger_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "type"),
				),
			},
		},
	})
}

func TestAccIBMTektonPipelineTriggerPropertyDataSourceAllArgs(t *testing.T) {
	triggerPropertyPipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	triggerPropertyTriggerID := fmt.Sprintf("tf_trigger_id_%d", acctest.RandIntRange(10, 100))
	triggerPropertyName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	triggerPropertyValue := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	triggerPropertyDefault := fmt.Sprintf("tf_default_%d", acctest.RandIntRange(10, 100))
	triggerPropertyType := "SECURE"
	triggerPropertyPath := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerPropertyDataSourceConfig(triggerPropertyPipelineID, triggerPropertyTriggerID, triggerPropertyName, triggerPropertyValue, triggerPropertyDefault, triggerPropertyType, triggerPropertyPath),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "trigger_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "property_name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "enum.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "default"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "path"),
				),
			},
		},
	})
}

func testAccCheckIBMTektonPipelineTriggerPropertyDataSourceConfigBasic(triggerPropertyPipelineID string, triggerPropertyTriggerID string) string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
			pipeline_id = "%s"
			trigger_id = "%s"
		}

		data "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
			pipeline_id = ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property.pipeline_id
			trigger_id = ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property.trigger_id
			property_name = "debug-pipeline"
		}
	`, triggerPropertyPipelineID, triggerPropertyTriggerID)
}

func testAccCheckIBMTektonPipelineTriggerPropertyDataSourceConfig(triggerPropertyPipelineID string, triggerPropertyTriggerID string, triggerPropertyName string, triggerPropertyValue string, triggerPropertyDefault string, triggerPropertyType string, triggerPropertyPath string) string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
			pipeline_id = "%s"
			trigger_id = "%s"
			name = "%s"
			value = "%s"
			enum = "FIXME"
			default = "%s"
			type = "%s"
			path = "%s"
		}

		data "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
			pipeline_id = ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property.pipeline_id
			trigger_id = ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property.trigger_id
			property_name = "debug-pipeline"
		}
	`, triggerPropertyPipelineID, triggerPropertyTriggerID, triggerPropertyName, triggerPropertyValue, triggerPropertyDefault, triggerPropertyType, triggerPropertyPath)
}
