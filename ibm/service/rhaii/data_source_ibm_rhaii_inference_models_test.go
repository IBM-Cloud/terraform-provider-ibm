// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMRhaiiInferenceModelsDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-rhaii-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMRhaiiProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRhaiiInferenceModelsDataSourceConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_rhaii_inference_models.models", "models.#"),
					resource.TestCheckResourceAttrSet("data.ibm_rhaii_inference_models.models", "models.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_rhaii_inference_models.models", "models.0.owned_by"),
					resource.TestCheckResourceAttrPair("data.ibm_rhaii_inference_model.model", "identifier", "data.ibm_rhaii_inference_models.models", "models.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_rhaii_inference_model.model", "model_type"),
					resource.TestCheckResourceAttrSet("data.ibm_rhaii_inference_model.model", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMRhaiiInferenceModelsDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_rhaii_project" "project" {
		name = "%s"
	}

	data "ibm_rhaii_inference_models" "models" {
		project_id = ibm_rhaii_project.project.project_id
		location   = ibm_rhaii_project.project.location
	}

	data "ibm_rhaii_inference_model" "model" {
		project_id = ibm_rhaii_project.project.project_id
		location   = ibm_rhaii_project.project.location
		model      = data.ibm_rhaii_inference_models.models.models[0].id
	}
	`, name)
}
