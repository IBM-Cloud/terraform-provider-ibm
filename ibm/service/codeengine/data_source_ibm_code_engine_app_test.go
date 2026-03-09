// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineAppDataSourceBasic(t *testing.T) {
	appName := fmt.Sprintf("tf-data-app-basic-%d", acctest.RandIntRange(10, 1000))
	appImageReference := "icr.io/codeengine/helloworld"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAppDataSourceConfigBasic(projectID, appImageReference, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_app.code_engine_app_instance", "app_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "resource_type", "app_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "image_reference", appImageReference),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "name", appName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "image_port", "8080"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "managed_domain_mappings", "local_public"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "run_service_account", "default"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_concurrency", "100"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_ephemeral_storage_limit", "400M"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_initial_instances", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_max_instances", "10"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_memory_limit", "4G"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_min_instances", "0"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", "300"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "run_compute_resource_token_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "computed_env_variables.#", "6"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineAppDataSourceExtended(t *testing.T) {
	appImageReference := "icr.io/codeengine/helloworld"
	appName := fmt.Sprintf("tf-data-app-extended-%d", acctest.RandIntRange(10, 1000))
	appImagePort := "8080"
	appManagedDomainMappings := "local_public"
	appRunAsUser := "10"
	appRunServiceAccount := "default"
	appScaleConcurrency := fmt.Sprintf("%d", acctest.RandIntRange(50, 100))
	appScaleConcurrencyTarget := fmt.Sprintf("%d", acctest.RandIntRange(20, 50))
	appScaleCPULimit := "0.5"
	appScaleEphemeralStorageLimit := "500M"
	appScaleInitialInstances := "2"
	appScaleMaxInstances := "2"
	appScaleMemoryLimit := "1G"
	appScaleMinInstances := "1"
	appScaleRequestTimeout := fmt.Sprintf("%d", acctest.RandIntRange(10, 30))

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAppDataSourceConfig(projectID, appImageReference, appName, appImagePort, appManagedDomainMappings, appRunAsUser, appRunServiceAccount, appScaleConcurrency, appScaleConcurrencyTarget, appScaleCPULimit, appScaleEphemeralStorageLimit, appScaleInitialInstances, appScaleMaxInstances, appScaleMemoryLimit, appScaleMinInstances, appScaleRequestTimeout),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "image_reference", appImageReference),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "name", appName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "image_port", appImagePort),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "managed_domain_mappings", appManagedDomainMappings),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "run_as_user", appRunAsUser),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "run_service_account", appRunServiceAccount),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_concurrency", appScaleConcurrency),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_concurrency_target", appScaleConcurrencyTarget),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_cpu_limit", appScaleCPULimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_ephemeral_storage_limit", appScaleEphemeralStorageLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_initial_instances", appScaleInitialInstances),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_max_instances", appScaleMaxInstances),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_memory_limit", appScaleMemoryLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_min_instances", appScaleMinInstances),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", appScaleRequestTimeout),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "run_compute_resource_token_enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "computed_env_variables.#", "6"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_app.code_engine_app_instance", "run_env_variables.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineAppDataSourceConfigBasic(projectID string, appImageReference string, appName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "%s"
			name = "%s"

			lifecycle {
				ignore_changes = [
					probe_liveness,
					probe_readiness
				]
			}
		}

		data "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = ibm_code_engine_app.code_engine_app_instance.project_id
			name = ibm_code_engine_app.code_engine_app_instance.name
		}
	`, projectID, appImageReference, appName)
}

func testAccCheckIbmCodeEngineAppDataSourceConfig(projectID string, appImageReference string, appName string, appImagePort string, appManagedDomainMappings string, appRunAsUser string, appRunServiceAccount string, appScaleConcurrency string, appScaleConcurrencyTarget string, appScaleCPULimit string, appScaleEphemeralStorageLimit string, appScaleInitialInstances string, appScaleMaxInstances string, appScaleMemoryLimit string, appScaleMinInstances string, appScaleRequestTimeout string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "%s"
			name = "%s"
			image_port = %s
			managed_domain_mappings = "%s"
			run_as_user = %s
			run_service_account = "%s"
			scale_concurrency = %s
			scale_concurrency_target = %s
			scale_cpu_limit = "%s"
			scale_ephemeral_storage_limit = "%s"
			scale_initial_instances = %s
			scale_max_instances = %s
			scale_memory_limit = "%s"
			scale_min_instances = %s
			scale_request_timeout = %s
			run_compute_resource_token_enabled = true

			run_env_variables {
				type  = "literal"
				name  = "name"
				value = "value"
			}

			lifecycle {
				ignore_changes = [
					probe_liveness,
					probe_readiness
				]
			}
		}

		data "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = ibm_code_engine_app.code_engine_app_instance.project_id
			name = ibm_code_engine_app.code_engine_app_instance.name
		}
	`, projectID, appImageReference, appName, appImagePort, appManagedDomainMappings, appRunAsUser, appRunServiceAccount, appScaleConcurrency, appScaleConcurrencyTarget, appScaleCPULimit, appScaleEphemeralStorageLimit, appScaleInitialInstances, appScaleMaxInstances, appScaleMemoryLimit, appScaleMinInstances, appScaleRequestTimeout)
}
