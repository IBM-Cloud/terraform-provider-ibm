// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineAppBasic(t *testing.T) {
	var conf codeenginev2.App
	name := fmt.Sprintf("tf-app-basic-%d", acctest.RandIntRange(10, 1000))
	imageReference := "icr.io/codeengine/helloworld"

	nameUpdate := fmt.Sprintf("tf-app-basic-update-%d", acctest.RandIntRange(10, 1000))
	imageReferenceUpdate := "icr.io/codeengine/hello"

	projectID := acc.CeProjectId

	envVars := `run_env_variables {
			type  = "literal"
			name  = "name"
			value = "value"
		}`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineAppDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAppConfigBasic(projectID, imageReference, name, envVars),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineAppExists("ibm_code_engine_app.code_engine_app_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_app.code_engine_app_instance", "app_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "resource_type", "app_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_reference", imageReference),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_port", "8080"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "managed_domain_mappings", "local_public"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_service_account", "default"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_concurrency", "100"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_ephemeral_storage_limit", "400M"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_initial_instances", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_max_instances", "10"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_memory_limit", "4G"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_min_instances", "0"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", "300"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "computed_env_variables.#", "6"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_env_variables.#", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.#", "0"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAppConfigBasic(projectID, imageReferenceUpdate, nameUpdate, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_app.code_engine_app_instance", "project_id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_app.code_engine_app_instance", "app_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "resource_type", "app_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_reference", imageReferenceUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_port", "8080"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "managed_domain_mappings", "local_public"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_service_account", "default"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_concurrency", "100"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_ephemeral_storage_limit", "400M"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_initial_instances", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_max_instances", "10"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_memory_limit", "4G"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_min_instances", "0"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", "300"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "computed_env_variables.#", "6"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_env_variables.#", "0"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.#", "0"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineAppExtended(t *testing.T) {
	var conf codeenginev2.App
	name := fmt.Sprintf("tf-app-extended-%d", acctest.RandIntRange(10, 1000))
	imageReference := "icr.io/codeengine/helloworld"
	imagePort := "8080"
	managedDomainMappings := "local_public"
	runAsUser := "10"
	runServiceAccount := "default"
	scaleConcurrency := fmt.Sprintf("%d", acctest.RandIntRange(50, 100))
	scaleConcurrencyTarget := fmt.Sprintf("%d", acctest.RandIntRange(20, 50))
	scaleCpuLimit := "0.5"
	scaleEphemeralStorageLimit := "500M"
	scaleInitialInstances := "2"
	scaleMaxInstances := "2"
	scaleMemoryLimit := "1G"
	scaleMinInstances := "1"
	scaleRequestTimeout := fmt.Sprintf("%d", acctest.RandIntRange(10, 30))
	configMapName := "my-config-map"
	configMapData := `{ "key" = "inner" }`

	envVars := `run_env_variables {
			type  = "literal"
			name  = "key1"
			value = "value1"
		}

		run_env_variables {
			type  = "literal"
			name  = "key2"
			value = "value2"
		}`

	volumeMounts := `
		run_volume_mounts {
			mount_path = "/mount"
			name       = "mymount"
			reference  = ibm_code_engine_config_map.code_engine_config_map_instance.name
			type       = "config_map"
		}`

	nameUpdate := fmt.Sprintf("tf-app-extended-update-%d", acctest.RandIntRange(10, 1000))
	imageReferenceUpdate := "icr.io/codeengine/hello"
	imagePortUpdate := "8080"
	managedDomainMappingsUpdate := "local"
	runServiceAccountUpdate := "default"
	scaleConcurrencyUpdate := fmt.Sprintf("%d", acctest.RandIntRange(50, 100))
	scaleConcurrencyTargetUpdate := fmt.Sprintf("%d", acctest.RandIntRange(20, 50))
	scaleCpuLimitUpdate := "1"
	scaleEphemeralStorageLimitUpdate := "1G"
	scaleInitialInstancesUpdate := "1"
	scaleMaxInstancesUpdate := "2"
	scaleMemoryLimitUpdate := "2G"
	scaleMinInstancesUpdate := "0"
	scaleRequestTimeoutUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 30))

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineAppDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAppConfig(projectID, configMapName, configMapData, imageReference, name, imagePort, managedDomainMappings, runAsUser, runServiceAccount, scaleConcurrency, scaleConcurrencyTarget, scaleCpuLimit, scaleEphemeralStorageLimit, scaleInitialInstances, scaleMaxInstances, scaleMemoryLimit, scaleMinInstances, scaleRequestTimeout, "", volumeMounts),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineAppExists("ibm_code_engine_app.code_engine_app_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_reference", imageReference),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_port", imagePort),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "managed_domain_mappings", managedDomainMappings),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_as_user", runAsUser),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_service_account", runServiceAccount),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_concurrency", scaleConcurrency),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_concurrency_target", scaleConcurrencyTarget),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_cpu_limit", scaleCpuLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_ephemeral_storage_limit", scaleEphemeralStorageLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_initial_instances", scaleInitialInstances),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_max_instances", scaleMaxInstances),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_memory_limit", scaleMemoryLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_min_instances", scaleMinInstances),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", scaleRequestTimeout),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "computed_env_variables.#", "6"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_env_variables.#", "2"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAppConfig(projectID, configMapName, configMapData, imageReferenceUpdate, nameUpdate, imagePortUpdate, managedDomainMappingsUpdate, runAsUser, runServiceAccountUpdate, scaleConcurrencyUpdate, scaleConcurrencyTargetUpdate, scaleCpuLimitUpdate, scaleEphemeralStorageLimitUpdate, scaleInitialInstancesUpdate, scaleMaxInstancesUpdate, scaleMemoryLimitUpdate, scaleMinInstancesUpdate, scaleRequestTimeoutUpdate, envVars, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_reference", imageReferenceUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "image_port", imagePortUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "managed_domain_mappings", managedDomainMappingsUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_as_user", runAsUser),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_service_account", runServiceAccountUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_concurrency", scaleConcurrencyUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_concurrency_target", scaleConcurrencyTargetUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_cpu_limit", scaleCpuLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_ephemeral_storage_limit", scaleEphemeralStorageLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_initial_instances", scaleInitialInstancesUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_max_instances", scaleMaxInstancesUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_memory_limit", scaleMemoryLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_min_instances", scaleMinInstancesUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", scaleRequestTimeoutUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "scale_request_timeout", scaleRequestTimeoutUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "computed_env_variables.#", "6"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_env_variables.#", "4"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.#", "0"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_app.code_engine_app_instance",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccCheckIbmCodeEngineAppConfigBasic(projectID string, imageReference string, name string, envVars string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "%s"
			name = "%s"

			%s

			lifecycle {
				ignore_changes = [
					probe_liveness,
					probe_readiness
				]
			}
		}
	`, projectID, imageReference, name, envVars)
}

func testAccCheckIbmCodeEngineAppConfig(projectID string, configMapName string, configMapData string, imageReference string, name string, imagePort string, managedDomainMappings string, runAsUser string, runServiceAccount string, scaleConcurrency string, scaleConcurrencyTarget string, scaleCpuLimit string, scaleEphemeralStorageLimit string, scaleInitialInstances string, scaleMaxInstances string, scaleMemoryLimit string, scaleMinInstances string, scaleRequestTimeout string, envVars string, volumeMounts string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name       = "%s"
			data       = %s
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
			run_env_variables {
				reference = ibm_code_engine_config_map.code_engine_config_map_instance.name
				type      = "config_map_full_reference"
				prefix    = "PREFIX_"
			}

			run_env_variables {
				type  = "literal"
				name  = "name"
				value = "value"
			}

			%s

			%s

			lifecycle {
				ignore_changes = [
					probe_liveness,
					probe_readiness
				]
			}
		}
	`, projectID, configMapName, configMapData, imageReference, name, imagePort, managedDomainMappings, runAsUser, runServiceAccount, scaleConcurrency, scaleConcurrencyTarget, scaleCpuLimit, scaleEphemeralStorageLimit, scaleInitialInstances, scaleMaxInstances, scaleMemoryLimit, scaleMinInstances, scaleRequestTimeout, envVars, volumeMounts)
}

func testAccCheckIbmCodeEngineAppExists(n string, obj codeenginev2.App) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getAppOptions := &codeenginev2.GetAppOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAppOptions.SetProjectID(parts[0])
		getAppOptions.SetName(parts[1])

		app, _, err := codeEngineClient.GetApp(getAppOptions)
		if err != nil {
			return err
		}

		obj = *app
		return nil
	}
}

func testAccCheckIbmCodeEngineAppDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_app" {
			continue
		}

		getAppOptions := &codeenginev2.GetAppOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAppOptions.SetProjectID(parts[0])
		getAppOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetApp(getAppOptions)

		if err == nil {
			return fmt.Errorf("code_engine_app still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_app (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
