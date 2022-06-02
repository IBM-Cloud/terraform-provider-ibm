// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtektonpipeline_test

import (
	"fmt"
	"testing"

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
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline.tekton_pipeline", "properties.#"),
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

func testAccCheckIBMTektonPipelineDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_tekton_pipeline" "tekton_pipeline" {
		}

		data "ibm_tekton_pipeline" "tekton_pipeline" {
			id = "94619026-912b-4d92-8f51-6c74f0692d90"
		}
	`)
}
