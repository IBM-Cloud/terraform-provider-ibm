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

func TestAccIbmCodeEngineBindingDataSourceBasic(t *testing.T) {
	prefix := fmt.Sprintf("PREFIX_%d", acctest.RandIntRange(10, 100))

	projectID := acc.CeProjectId
	resourceKeyId := acc.CeResourceKeyID
	serviceInstanceId := acc.CeServiceInstanceID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineBindingDataSourceConfigBasic(projectID, resourceKeyId, serviceInstanceId, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_binding.code_engine_binding_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_binding.code_engine_binding_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_binding.code_engine_binding_instance", "href"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "resource_type", "binding_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "component.resource_type", "app_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "component.name", "app_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "prefix", prefix),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "secret_name", "my-service-access"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineBindingDataSourceConfigBasic(projectID string, resourceKeyId string, serviceInstanceId string, prefix string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "icr.io/codeengine/helloworld"
			name = "my-app"

			lifecycle {
				ignore_changes = [
					run_env_variables
				]
			}
		}

		resource "ibm_code_engine_secret" "code_engine_secret_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			format = "service_access"
			name = "my-service-access"
			service_access {
				resource_key {
					id = "%s"
				}
				service_instance {
					id = "%s"
				}
			}
		}

		resource "ibm_code_engine_binding" "code_engine_binding_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			component {
				name = ibm_code_engine_app.code_engine_app_instance.name
				resource_type = ibm_code_engine_app.code_engine_app_instance.resource_type
			}
			prefix = "%s"
			secret_name = ibm_code_engine_secret.code_engine_secret_instance.name
		}

		data "ibm_code_engine_binding" "code_engine_binding_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			id = ibm_code_engine_binding.code_engine_binding_instance.binding_id
		}
	`, projectID, resourceKeyId, serviceInstanceId, prefix)
}
