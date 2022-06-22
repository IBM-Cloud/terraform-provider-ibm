// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
)

func TestAccIBMTektonPipelineBasic(t *testing.T) {
	var conf cdtektonpipelinev2.TektonPipeline
	rgID := acc.CdResourceGroupID
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTektonPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineConfigBasic(tcName, rgID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMTektonPipelineExists("ibm_cd_tekton_pipeline.tekton_pipeline", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline.tekton_pipeline",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMTektonPipelineConfigBasic(tcName string, rgID string) string {
	return fmt.Sprintf(`
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = "%s"
		}

		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "name"
				type = "tekton"
				ui_pipeline = true
			}
		}

		resource "ibm_cd_tekton_pipeline" "tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			worker {
				id = "public"
			}			
		}
	`, tcName, rgID)
}

func testAccCheckIBMTektonPipelineExists(n string, obj cdtektonpipelinev2.TektonPipeline) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
		if err != nil {
			return err
		}

		getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{}

		getTektonPipelineOptions.SetID(rs.Primary.ID)

		tektonPipeline, _, err := cdTektonPipelineClient.GetTektonPipeline(getTektonPipelineOptions)
		if err != nil {
			return err
		}

		obj = *tektonPipeline
		return nil
	}
}

func testAccCheckIBMTektonPipelineDestroy(s *terraform.State) error {
	cdTektonPipelineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CdTektonPipelineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_tekton_pipeline" {
			continue
		}

		getTektonPipelineOptions := &cdtektonpipelinev2.GetTektonPipelineOptions{}

		getTektonPipelineOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := cdTektonPipelineClient.GetTektonPipeline(getTektonPipelineOptions)

		if err == nil {
			return fmt.Errorf("tekton_pipeline still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for tekton_pipeline (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
