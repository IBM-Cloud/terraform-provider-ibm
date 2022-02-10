// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package continuousdeliverypipeline_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTektonPipelineWorkersDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTektonPipelineWorkersDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_workers.tekton_pipeline_workers", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_workers.tekton_pipeline_workers", "pipeline_id"),
					resource.TestCheckResourceAttrSet("data.ibm_tekton_pipeline_workers.tekton_pipeline_workers", "workers.#"),
				),
			},
		},
	})
}

func testAccCheckIBMTektonPipelineWorkersDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_tekton_pipeline_workers" "tekton_pipeline_workers" {
			pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
		}
	`)
}

