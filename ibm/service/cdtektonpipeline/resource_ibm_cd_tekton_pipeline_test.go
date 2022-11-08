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

func TestAccIBMCdTektonPipelineBasic(t *testing.T) {
	var conf cdtektonpipelinev2.TektonPipeline

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineExists("ibm_cd_tekton_pipeline.cd_tekton_pipeline", conf),
				),
			},
		},
	})
}

func TestAccIBMCdTektonPipelineAllArgs(t *testing.T) {
	var conf cdtektonpipelinev2.TektonPipeline
	enableNotifications := "true"
	enablePartialCloning := "true"
	enableNotificationsUpdate := "false"
	enablePartialCloningUpdate := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCdTektonPipelineDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineConfig(enableNotifications, enablePartialCloning),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdTektonPipelineExists("ibm_cd_tekton_pipeline.cd_tekton_pipeline", conf),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_notifications", enableNotifications),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_partial_cloning", enablePartialCloning),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMCdTektonPipelineConfig(enableNotificationsUpdate, enablePartialCloningUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_notifications", enableNotificationsUpdate),
					resource.TestCheckResourceAttr("ibm_cd_tekton_pipeline.cd_tekton_pipeline", "enable_partial_cloning", enablePartialCloningUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cd_tekton_pipeline.cd_tekton_pipeline",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMCdTektonPipelineConfigBasic() string {
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}
		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-name"
			}
		}
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
	`, rgName, tcName)
}

func testAccCheckIBMCdTektonPipelineConfig(enableNotifications string, enablePartialCloning string) string {
	rgName := acc.CdResourceGroupName
	tcName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}
		resource "ibm_cd_toolchain_tool_pipeline" "ibm_cd_toolchain_tool_pipeline" {
			toolchain_id = ibm_cd_toolchain.cd_toolchain.id
			parameters {
				name = "pipeline-name"
			}
		}
		resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline" {
			pipeline_id = ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline.tool_id
			enable_notifications = %s
			enable_partial_cloning = %s
			worker {
				id = "public"
			}
			depends_on = [
				ibm_cd_toolchain_tool_pipeline.ibm_cd_toolchain_tool_pipeline
			]
		}
	`, rgName, tcName, enableNotifications, enablePartialCloning)
}

func testAccCheckIBMCdTektonPipelineExists(n string, obj cdtektonpipelinev2.TektonPipeline) resource.TestCheckFunc {

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

func testAccCheckIBMCdTektonPipelineDestroy(s *terraform.State) error {
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
			return fmt.Errorf("cd_tekton_pipeline still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for cd_tekton_pipeline (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
