// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineFunctionDataSourceBasic(t *testing.T) {
	functionName := fmt.Sprintf("%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	functionRuntime := "nodejs-20"
	functionCodeReference := "data:text/plain;base64,foo"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineFunctionDataSourceConfigBasic(projectID, functionCodeReference, functionName, functionRuntime),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_function.code_engine_function_instance", "function_id"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "resource_type", "function_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "name", functionName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "code_binary", "false"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "code_reference", functionCodeReference),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "managed_domain_mappings", "local_public"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "runtime", functionRuntime),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_concurrency", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_down_delay", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_max_execution_time", "60"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_memory_limit", "4G"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineFunctionDataSourceExtended(t *testing.T) {
	functionName := fmt.Sprintf("%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	functionRuntime := "nodejs-20"
	functionCodeReference := "data:text/plain;base64,foo"
	functionManagedDomainMappings := "local_private"
	functionScaleCpuLimit := "0.5"
	functionScaleDownDelay := "20"
	functionScaleMaxExecutionTime := "30"
	functionScaleMemoryLimit := "2G"

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineFunctionDataSourceConfig(projectID, functionCodeReference, functionManagedDomainMappings, functionName, functionRuntime, functionScaleCpuLimit, functionScaleDownDelay, functionScaleMaxExecutionTime, functionScaleMemoryLimit),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "resource_type", "function_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "name", functionName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "code_binary", "false"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "code_reference", functionCodeReference),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "managed_domain_mappings", functionManagedDomainMappings),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "runtime", functionRuntime),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_concurrency", "1"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_cpu_limit", functionScaleCpuLimit),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_down_delay", functionScaleDownDelay),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_max_execution_time", functionScaleMaxExecutionTime),
					resource.TestCheckResourceAttr("data.ibm_code_engine_function.code_engine_function_instance", "scale_memory_limit", functionScaleMemoryLimit),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineFunctionDataSourceConfigBasic(projectID string, functionCodeReference string, functionName string, functionRuntime string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_function" "code_engine_function_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			code_reference = "%s"
			name = "%s"
			runtime = "%s"

			lifecycle {
				ignore_changes = [
					run_env_variables
				]
			}
		}

		data "ibm_code_engine_function" "code_engine_function_instance" {
			project_id = ibm_code_engine_function.code_engine_function_instance.project_id
			name = ibm_code_engine_function.code_engine_function_instance.name
		}
	`, projectID, functionCodeReference, functionName, functionRuntime)
}

func testAccCheckIbmCodeEngineFunctionDataSourceConfig(projectID string, functionCodeReference string, functionManagedDomainMappings string, functionName string, functionRuntime string, functionScaleCpuLimit string, functionScaleDownDelay string, functionScaleMaxExecutionTime string, functionScaleMemoryLimit string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_function" "code_engine_function_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			code_reference = "%s"
			managed_domain_mappings = "%s"
			name = "%s"
			runtime = "%s"
			scale_cpu_limit = "%s"
			scale_down_delay = %s
			scale_max_execution_time = %s
			scale_memory_limit = "%s"

			run_env_variables {
				type  = "literal"
				name  = "name"
				value = "value"
			}

			lifecycle {
				ignore_changes = [
					run_env_variables
				]
			}
		}

		data "ibm_code_engine_function" "code_engine_function_instance" {
			project_id = ibm_code_engine_function.code_engine_function_instance.project_id
			name = ibm_code_engine_function.code_engine_function_instance.name
		}
	`, projectID, functionCodeReference, functionManagedDomainMappings, functionName, functionRuntime, functionScaleCpuLimit, functionScaleDownDelay, functionScaleMaxExecutionTime, functionScaleMemoryLimit)
}
