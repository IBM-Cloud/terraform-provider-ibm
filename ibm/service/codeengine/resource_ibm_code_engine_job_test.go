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

func TestAccIbmCodeEngineJobBasic(t *testing.T) {
	var conf codeenginev2.Job
	name := fmt.Sprintf("tf-job-basic-%d", acctest.RandIntRange(10, 1000))
	imageReference := "icr.io/codeengine/helloworld"

	nameUpdate := fmt.Sprintf("tf-job-basic-update-%d", acctest.RandIntRange(10, 1000))
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
		CheckDestroy: testAccCheckIbmCodeEngineJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobConfigBasic(projectID, name, imageReference, envVars),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineJobExists("ibm_code_engine_job.code_engine_job_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_job.code_engine_job_instance", "job_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "resource_type", "job_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "image_reference", imageReference),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_service_account", "default"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_ephemeral_storage_limit", "400M"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_max_execution_time", "7200"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_memory_limit", "4G"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_retry_limit", "3"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "computed_env_variables.#", "3"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_env_variables.#", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_volume_mounts.#", "0"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobConfigBasic(projectID, nameUpdate, imageReferenceUpdate, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_job.code_engine_job_instance", "job_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "resource_type", "job_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "image_reference", imageReferenceUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_service_account", "default"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_ephemeral_storage_limit", "400M"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_max_execution_time", "7200"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_memory_limit", "4G"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_retry_limit", "3"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "computed_env_variables.#", "3"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_env_variables.#", "0"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_volume_mounts.#", "0"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineJobExtended(t *testing.T) {
	var conf codeenginev2.Job
	name := fmt.Sprintf("tf-job-extended-%d", acctest.RandIntRange(10, 1000))
	imageReference := "icr.io/codeengine/helloworld"
	runAsUser := "1001"
	runMode := "task"
	runServiceAccount := "default"
	scaleCpuLimit := "0.5"
	scaleEphemeralStorageLimit := "500M"
	scaleMaxExecutionTime := "3600"
	scaleMemoryLimit := "1G"
	scaleRetryLimit := "2"
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

	nameUpdate := fmt.Sprintf("tf-job-extended-update-%d", acctest.RandIntRange(10, 1000))
	imageReferenceUpdate := "icr.io/codeengine/hello"
	runAsUserUpdate := "0"
	runModeUpdate := "task"
	runServiceAccountUpdate := "none"
	scaleCpuLimitUpdate := "1"
	scaleEphemeralStorageLimitUpdate := "1G"
	scaleMaxExecutionTimeUpdate := "7200"
	scaleMemoryLimitUpdate := "2G"
	scaleRetryLimitUpdate := "3"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobConfig(projectID, configMapName, configMapData, name, imageReference, runAsUser, runMode, runServiceAccount, scaleCpuLimit, scaleEphemeralStorageLimit, scaleMaxExecutionTime, scaleMemoryLimit, scaleRetryLimit, "", volumeMounts),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineJobExists("ibm_code_engine_job.code_engine_job_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_job.code_engine_job_instance", "job_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "image_reference", imageReference),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_as_user", runAsUser),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_mode", runMode),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_service_account", runServiceAccount),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_cpu_limit", scaleCpuLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_ephemeral_storage_limit", scaleEphemeralStorageLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_max_execution_time", scaleMaxExecutionTime),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_memory_limit", scaleMemoryLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_retry_limit", scaleRetryLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "computed_env_variables.#", "3"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_env_variables.#", "2"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_volume_mounts.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineJobConfig(projectID, configMapName, configMapData, nameUpdate, imageReferenceUpdate, runAsUserUpdate, runModeUpdate, runServiceAccountUpdate, scaleCpuLimitUpdate, scaleEphemeralStorageLimitUpdate, scaleMaxExecutionTimeUpdate, scaleMemoryLimitUpdate, scaleRetryLimitUpdate, envVars, ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_job.code_engine_job_instance", "job_id"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "image_reference", imageReferenceUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_as_user", runAsUserUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_mode", runModeUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_service_account", runServiceAccountUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_cpu_limit", scaleCpuLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_ephemeral_storage_limit", scaleEphemeralStorageLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_max_execution_time", scaleMaxExecutionTimeUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_memory_limit", scaleMemoryLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "scale_retry_limit", scaleRetryLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "computed_env_variables.#", "3"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_env_variables.#", "4"),
					resource.TestCheckResourceAttr("ibm_code_engine_job.code_engine_job_instance", "run_volume_mounts.#", "0"),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_code_engine_job.code_engine_job_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"run_as_user"},
			},
		},
	})
}

func testAccCheckIbmCodeEngineJobConfigBasic(projectID string, name string, imageReference string, envVars string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_job" "code_engine_job_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name = "%s"
			image_reference = "%s"

			%s
		}
	`, projectID, name, imageReference, envVars)
}

func testAccCheckIbmCodeEngineJobConfig(projectID string, configMapName string, configMapData string, name string, imageReference string, runAsUser string, runMode string, runServiceAccount string, scaleCpuLimit string, scaleEphemeralStorageLimit string, scaleMaxExecutionTime string, scaleMemoryLimit string, scaleRetryLimit string, envVars string, volumeMounts string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name       = "%s"
			data       = %s
		}

		resource "ibm_code_engine_job" "code_engine_job_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			name = "%s"
			image_reference = "%s"
			run_as_user = %s
			run_mode = "%s"
			run_service_account = "%s"
			scale_cpu_limit = "%s"
			scale_ephemeral_storage_limit = "%s"
			scale_max_execution_time = %s
			scale_memory_limit = "%s"
			scale_retry_limit = %s

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
		}
	
	`, projectID, configMapName, configMapData, name, imageReference, runAsUser, runMode, runServiceAccount, scaleCpuLimit, scaleEphemeralStorageLimit, scaleMaxExecutionTime, scaleMemoryLimit, scaleRetryLimit, envVars, volumeMounts)
}

func testAccCheckIbmCodeEngineJobExists(n string, obj codeenginev2.Job) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getJobOptions := &codeenginev2.GetJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getJobOptions.SetProjectID(parts[0])
		getJobOptions.SetName(parts[1])

		job, _, err := codeEngineClient.GetJob(getJobOptions)
		if err != nil {
			return err
		}

		obj = *job
		return nil
	}
}

func testAccCheckIbmCodeEngineJobDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_job" {
			continue
		}

		getJobOptions := &codeenginev2.GetJobOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getJobOptions.SetProjectID(parts[0])
		getJobOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetJob(getJobOptions)

		if err == nil {
			return fmt.Errorf("code_engine_job still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_job (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
