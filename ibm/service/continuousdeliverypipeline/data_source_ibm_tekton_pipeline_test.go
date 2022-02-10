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

func TestAccIBMTektonPipelineDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "html_url"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "enabled"),
				),
			},
		},
	})
}

func TestAccIBMTektonPipelineDataSourceAllArgs(t *testing.T) {
	tektonPipelineIntegrationInstanceID := fmt.Sprintf("tf_integration_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineDataSourceConfig(tektonPipelineIntegrationInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "toolchain.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "definitions.0.service_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "definitions.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.0.options"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "env_properties.0.path"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "created"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "pipeline_definition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.event_listener"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.service_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.cron"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "triggers.0.timezone"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.created"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.trigger_name"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.trigger_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.timezone"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.sub"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.event_listener"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.cron"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.disabled"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "next_timers.0.expiration"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "worker.#"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "html_url"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "build_number"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "enabled"),
				),
			},
		},
	})
}

func testAccCheckIBMTektonPipelineDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline" "tekton_pipeline" {
		}

		data "ibm_tekton_pipeline" "tekton_pipeline" {
			id = "94619026-912b-4d92-8f51-6c74f0692d90"
		}
	`)
}

func testAccCheckIBMTektonPipelineDataSourceConfig(tektonPipelineIntegrationInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline" "tekton_pipeline" {
			integration_instance_id = "%s"
			worker {
				id = "id"
			}
		}

		data "ibm_tekton_pipeline" "tekton_pipeline" {
			id = "94619026-912b-4d92-8f51-6c74f0692d90"
		}
	`, tektonPipelineIntegrationInstanceID)
}
