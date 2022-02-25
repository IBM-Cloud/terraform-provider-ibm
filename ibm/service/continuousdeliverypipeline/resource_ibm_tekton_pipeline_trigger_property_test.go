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

func TestAccIBMTektonPipelineTriggerPropertyBasic(t *testing.T) {
	var conf continuousdeliverypipelinev2.TriggerProperty
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	triggerID := fmt.Sprintf("tf_trigger_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelineTriggerPropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerPropertyConfigBasic(pipelineID, triggerID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelineTriggerPropertyExists("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", conf),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "pipeline_id", pipelineID),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "trigger_id", triggerID),
				),
			},
		},
	})
}

func TestAccIBMTektonPipelineTriggerPropertyAllArgs(t *testing.T) {
	var conf continuousdeliverypipelinev2.TriggerProperty
	pipelineID := fmt.Sprintf("tf_pipeline_id_%d", acctest.RandIntRange(10, 100))
	triggerID := fmt.Sprintf("tf_trigger_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	value := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	defaultVar := fmt.Sprintf("tf_default_%d", acctest.RandIntRange(10, 100))
	typeVar := "SECURE"
	path := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	valueUpdate := fmt.Sprintf("tf_value_%d", acctest.RandIntRange(10, 100))
	defaultVarUpdate := fmt.Sprintf("tf_default_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "SINGLE_SELECT"
	pathUpdate := fmt.Sprintf("tf_path_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelineTriggerPropertyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerPropertyConfig(pipelineID, triggerID, name, value, defaultVar, typeVar, path),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelineTriggerPropertyExists("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", conf),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "pipeline_id", pipelineID),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "trigger_id", triggerID),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "name", name),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "value", value),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "default", defaultVar),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "path", path),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineTriggerPropertyConfig(pipelineID, triggerID, nameUpdate, valueUpdate, defaultVarUpdate, typeVarUpdate, pathUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "pipeline_id", pipelineID),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "trigger_id", triggerID),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "value", valueUpdate),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "default", defaultVarUpdate),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property", "path", pathUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelineTriggerPropertyConfigBasic(pipelineID string, triggerID string) string {
	return fmt.Sprintf(`

		resource "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
			pipeline_id = "%s"
			trigger_id = "%s"
		}
	`, pipelineID, triggerID)
}

func testAccCheckIBMTektonPipelineTriggerPropertyConfig(pipelineID string, triggerID string, name string, value string, defaultVar string, typeVar string, path string) string {
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
	`, pipelineID, triggerID, name, value, defaultVar, typeVar, path)
}

func testAccCheckIBMTektonPipelineTriggerPropertyExists(n string, obj continuousdeliverypipelinev2.TriggerProperty) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineTriggerPropertyOptions := &continuousdeliverypipelinev2.GetTektonPipelineTriggerPropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
		getTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

		triggerProperty, _, err := continuousDeliveryPipelineClient.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)
		if err != nil {
			return err
		}

		obj = *triggerProperty
		return nil
	}
}

func testAccCheckIBMTektonPipelineTriggerPropertyDestroy(s *terraform.State) error {
	continuousDeliveryPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContinuousDeliveryPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tekton_pipeline_trigger_property" {
			continue
		}

		getTektonPipelineTriggerPropertyOptions := &continuousdeliverypipelinev2.GetTektonPipelineTriggerPropertyOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTektonPipelineTriggerPropertyOptions.SetPipelineID(parts[0])
		getTektonPipelineTriggerPropertyOptions.SetTriggerID(parts[1])
		getTektonPipelineTriggerPropertyOptions.SetPropertyName(parts[2])

		// Try to find the key
		_, response, err := continuousDeliveryPipelineClient.GetTektonPipelineTriggerProperty(getTektonPipelineTriggerPropertyOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline_trigger_property still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline_trigger_property (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
