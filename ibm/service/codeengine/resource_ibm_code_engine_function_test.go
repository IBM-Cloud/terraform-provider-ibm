// Copyright IBM Corp. 2024 All Rights Reserved.
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

func TestAccIbmCodeEngineFunctionBasic(t *testing.T) {
	var conf codeenginev2.Function
	functionName := fmt.Sprintf("%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	functionRuntime := "nodejs-20"
	functionCodeReference := "data:text/plain;base64,foo"

	projectID := acc.CeProjectId

	functionCodeReferenceUpdate := "data:text/plain;base64,bar"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineFunctionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineFunctionConfigBasic(projectID, functionCodeReference, functionName, functionRuntime),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineFunctionExists("ibm_code_engine_function.code_engine_function_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "name", functionName),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "runtime", functionRuntime),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_binary", "false"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_reference", functionCodeReference),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "managed_domain_mappings", "local_public"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_concurrency", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_down_delay", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_max_execution_time", "60"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_memory_limit", "4G"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineFunctionConfigBasic(projectID, functionCodeReferenceUpdate, functionName, functionRuntime),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "name", functionName),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "runtime", functionRuntime),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_binary", "false"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_reference", functionCodeReferenceUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "managed_domain_mappings", "local_public"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_concurrency", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_cpu_limit", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_down_delay", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_max_execution_time", "60"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_memory_limit", "4G"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineFunctionExtended(t *testing.T) {
	var conf codeenginev2.Function

	functionName := fmt.Sprintf("%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlpha))
	functionRuntime := "nodejs-20"
	functionCodeReference := "data:text/plain;base64,foo"
	functionManagedDomainMappings := "local_public"
	functionScaleCpuLimit := "1"
	functionScaleDownDelay := "1"
	functionScaleMaxExecutionTime := "60"
	functionScaleMemoryLimit := "4G"

	projectID := acc.CeProjectId

	functionCodeReferenceUpdate := "data:text/plain;base64,bar"
	functionManagedDomainMappingsUpdate := "local_private"
	functionScaleCpuLimitUpdate := "0.5"
	functionScaleDownDelayUpdate := "20"
	functionScaleMaxExecutionTimeUpdate := "30"
	functionScaleMemoryLimitUpdate := "2G"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineFunctionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineFunctionConfig(projectID, functionCodeReference, functionManagedDomainMappings, functionName, functionRuntime, functionScaleCpuLimit, functionScaleDownDelay, functionScaleMaxExecutionTime, functionScaleMemoryLimit),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineFunctionExists("ibm_code_engine_function.code_engine_function_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "name", functionName),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "runtime", functionRuntime),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_binary", "false"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_reference", functionCodeReference),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "managed_domain_mappings", functionManagedDomainMappings),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_concurrency", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_cpu_limit", functionScaleCpuLimit),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_down_delay", functionScaleDownDelay),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_max_execution_time", functionScaleMaxExecutionTime),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_memory_limit", functionScaleMemoryLimit),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineFunctionConfig(projectID, functionCodeReferenceUpdate, functionManagedDomainMappingsUpdate, functionName, functionRuntime, functionScaleCpuLimitUpdate, functionScaleDownDelayUpdate, functionScaleMaxExecutionTimeUpdate, functionScaleMemoryLimitUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "name", functionName),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "runtime", functionRuntime),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_binary", "false"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "code_reference", functionCodeReferenceUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "managed_domain_mappings", functionManagedDomainMappingsUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_concurrency", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_cpu_limit", functionScaleCpuLimitUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_down_delay", functionScaleDownDelayUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_max_execution_time", functionScaleMaxExecutionTimeUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_function.code_engine_function_instance", "scale_memory_limit", functionScaleMemoryLimitUpdate),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineFunctionConfigBasic(projectID string, codeReference string, name string, runtime string) string {
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
	`, projectID, codeReference, name, runtime)
}

func testAccCheckIbmCodeEngineFunctionConfig(projectID string, codeReference string, managedDomainMappings string, name string, runtime string, scaleCpuLimit string, scaleDownDelay string, scaleMaxExecutionTime string, scaleMemoryLimit string) string {
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
	`, projectID, codeReference, managedDomainMappings, name, runtime, scaleCpuLimit, scaleDownDelay, scaleMaxExecutionTime, scaleMemoryLimit)
}

func testAccCheckIbmCodeEngineFunctionExists(n string, obj codeenginev2.Function) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getFunctionOptions := &codeenginev2.GetFunctionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getFunctionOptions.SetProjectID(parts[0])
		getFunctionOptions.SetName(parts[1])

		function, _, err := codeEngineClient.GetFunction(getFunctionOptions)
		if err != nil {
			return err
		}

		obj = *function
		return nil
	}
}

func testAccCheckIbmCodeEngineFunctionDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_function" {
			continue
		}

		getFunctionOptions := &codeenginev2.GetFunctionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getFunctionOptions.SetProjectID(parts[0])
		getFunctionOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetFunction(getFunctionOptions)

		if err == nil {
			return fmt.Errorf("code_engine_function still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_function (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
