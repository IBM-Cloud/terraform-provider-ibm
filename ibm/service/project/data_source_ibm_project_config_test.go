// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectConfigDataSourceBasic(t *testing.T) {
	projectConfigCanonicalName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectConfigCanonicalLocatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfigBasic(projectConfigCanonicalName, projectConfigCanonicalLocatorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.0.locator_id"),
				),
			},
		},
	})
}

func TestAccIbmProjectConfigDataSourceAllArgs(t *testing.T) {
	projectConfigCanonicalName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectConfigCanonicalDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectConfigCanonicalLocatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfig(projectConfigCanonicalName, projectConfigCanonicalDescription, projectConfigCanonicalLocatorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_config_canonical_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "needs_attention_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "last_approved.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "authorizations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "compliance_profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "input.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "setting.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "output.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.0.locator_id"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectConfigDataSourceConfigBasic(projectConfigCanonicalName string, projectConfigCanonicalLocatorID string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			locator_id = "%s"
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config_instance.project_id
			id = ibm_project_config.project_config_instance.project_config_canonical_id
		}
	`, projectConfigCanonicalName, projectConfigCanonicalLocatorID)
}

func testAccCheckIbmProjectConfigDataSourceConfig(projectConfigCanonicalName string, projectConfigCanonicalDescription string, projectConfigCanonicalLocatorID string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			labels = [ "labels" ]
			description = "%s"
			authorizations {
                method = "API_KEY"
                api_key = "<YOUR_APIKEY_HERE>"
            }
			locator_id = "%s"
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config_instance.project_id
			id = ibm_project_config.project_config_instance.project_config_canonical_id
		}
	`, projectConfigCanonicalName, projectConfigCanonicalDescription, projectConfigCanonicalLocatorID)
}
