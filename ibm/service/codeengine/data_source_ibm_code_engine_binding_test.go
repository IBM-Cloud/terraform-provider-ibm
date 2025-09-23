// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineBindingDataSourceBasic(t *testing.T) {
	prefix := fmt.Sprintf("PREFIX_%d", acctest.RandIntRange(10, 100))
	secretName := fmt.Sprintf("tf-secret-service-access-binding-%d", acctest.RandIntRange(10, 1000))
	appName := fmt.Sprintf("tf-app-binding-%d", acctest.RandIntRange(10, 1000))

	projectID := acc.CeProjectId
	resourceKeyId := acc.CeResourceKeyID
	serviceInstanceId := acc.CeServiceInstanceID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineBindingDataSourceConfigBasic(projectID, appName, secretName, resourceKeyId, serviceInstanceId, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_binding.code_engine_binding_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_binding.code_engine_binding_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_binding.code_engine_binding_instance", "href"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "resource_type", "binding_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "component.0.resource_type", "app_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "component.0.name", appName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "prefix", prefix),
					resource.TestCheckResourceAttr("data.ibm_code_engine_binding.code_engine_binding_instance", "secret_name", secretName),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineBindingDataSourceConfigBasic(projectID string, appName string, secretName string, resourceKeyId string, serviceInstanceId string, prefix string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			image_reference = "icr.io/codeengine/helloworld"
			name = "%s"

			lifecycle {
				ignore_changes = [
					probe_liveness,
					probe_readiness
				]
			}
		}

		resource "ibm_code_engine_secret" "code_engine_secret_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			format = "service_access"
			name = "%s"
			service_access {
				resource_key {
					id = "%s"
				}
				service_instance {
					id = "%s"
				}
			}
			lifecycle {
				ignore_changes = [
					data, service_access
				]
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
			binding_id = ibm_code_engine_binding.code_engine_binding_instance.binding_id
		}
	`, projectID, appName, secretName, resourceKeyId, serviceInstanceId, prefix)
}
