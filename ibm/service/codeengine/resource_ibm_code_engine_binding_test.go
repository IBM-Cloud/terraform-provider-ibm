// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineBindingBasic(t *testing.T) {
	var conf codeenginev2.Binding
	prefix := fmt.Sprintf("PREFIX_%d", acctest.RandIntRange(10, 100))
	secretName := fmt.Sprintf("tf-secret-service-access-binding-%d", acctest.RandIntRange(10, 1000))
	appName := fmt.Sprintf("tf-app-binding-%d", acctest.RandIntRange(10, 1000))

	projectID := acc.CeProjectId
	resourceKeyId := acc.CeResourceKeyID
	serviceInstanceId := acc.CeServiceInstanceID

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineBindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineBindingConfigBasic(projectID, appName, secretName, resourceKeyId, serviceInstanceId, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineBindingExists("ibm_code_engine_binding.code_engine_binding_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_binding.code_engine_binding_instance", "status"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_binding.code_engine_binding_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_binding.code_engine_binding_instance", "href"),
					resource.TestCheckResourceAttr("ibm_code_engine_binding.code_engine_binding_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_binding.code_engine_binding_instance", "resource_type", "binding_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_binding.code_engine_binding_instance", "component.0.resource_type", "app_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_binding.code_engine_binding_instance", "component.0.name", appName),
					resource.TestCheckResourceAttr("ibm_code_engine_binding.code_engine_binding_instance", "prefix", prefix),
					resource.TestCheckResourceAttr("ibm_code_engine_binding.code_engine_binding_instance", "secret_name", secretName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_binding.code_engine_binding_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmCodeEngineBindingConfigBasic(projectID string, appName string, secretName string, resourceKeyId string, serviceInstanceId string, prefix string) string {
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
	`, projectID, appName, secretName, resourceKeyId, serviceInstanceId, prefix)
}

func testAccCheckIbmCodeEngineBindingExists(n string, obj codeenginev2.Binding) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getBindingOptions := &codeenginev2.GetBindingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBindingOptions.SetProjectID(parts[0])
		getBindingOptions.SetID(parts[1])

		binding, _, err := codeEngineClient.GetBinding(getBindingOptions)
		if err != nil {
			return err
		}

		obj = *binding
		return nil
	}
}

func testAccCheckIbmCodeEngineBindingDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_binding" {
			continue
		}

		getBindingOptions := &codeenginev2.GetBindingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getBindingOptions.SetProjectID(parts[0])
		getBindingOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetBinding(getBindingOptions)

		if err == nil {
			return fmt.Errorf("code_engine_binding_instance still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_binding_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
