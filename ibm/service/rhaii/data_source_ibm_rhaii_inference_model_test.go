// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMRhaiiInferenceModelDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("tf-rhaii-%s", acctest.RandString(8))
	dataSource := "data.ibm_rhaii_inference_model.model"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		CheckDestroy:             testAccCheckIBMRhaiiProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRhaiiInferenceModelDataSourceConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSource, "project_id", "ibm_rhaii_project.project", "project_id"),
					resource.TestCheckResourceAttrPair(dataSource, "model", "data.ibm_rhaii_inference_models.models", "models.0.id"),
					resource.TestCheckResourceAttrPair(dataSource, "identifier", "data.ibm_rhaii_inference_models.models", "models.0.id"),
					resource.TestMatchResourceAttr(dataSource, "id", regexp.MustCompile(`^[0-9a-f-]{36}/[a-z0-9-]+$`)),
					resource.TestCheckResourceAttrSet(dataSource, "provider_id"),
					resource.TestCheckResourceAttrSet(dataSource, "type"),
					resource.TestCheckResourceAttrSet(dataSource, "model_type"),
					resource.TestCheckResourceAttrSet(dataSource, "status"),
					resource.TestCheckResourceAttrSet(dataSource, "config_json"),
				),
			},
		},
	})
}

func TestAccIBMRhaiiInferenceModelDataSource_notFound(t *testing.T) {
	name := fmt.Sprintf("tf-rhaii-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		CheckDestroy:             testAccCheckIBMRhaiiProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMRhaiiInferenceModelNotFoundConfig(name),
				ExpectError: regexp.MustCompile(`Error reading the inference model`),
			},
		},
	})
}

func testAccCheckIBMRhaiiInferenceModelDataSourceConfig(name string) string {
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

func testAccCheckIBMRhaiiInferenceModelNotFoundConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_rhaii_project" "project" {
		name = "%s"
	}

	data "ibm_rhaii_inference_model" "missing" {
		project_id = ibm_rhaii_project.project.project_id
		model      = "model-that-does-not-exist"
	}
	`, name)
}
