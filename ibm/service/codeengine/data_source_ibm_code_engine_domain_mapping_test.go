// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmCodeEngineDomainMappingDataSourceBasic(t *testing.T) {
	appName := fmt.Sprintf("tf-app-domain-mapping-%d", acctest.RandIntRange(10, 1000))
	secretName := fmt.Sprintf("tf-secret-domain-mapping-%d", acctest.RandIntRange(10, 1000))

	projectID := acc.CeProjectId
	domainMappingName := acc.CeDomainMappingName
	domainMappingTLSKey, _ := os.ReadFile(acc.CeTLSKeyFilePath)
	domainMappingTLSCert, _ := os.ReadFile(acc.CeTLSCertFilePath)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCodeEngine(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCodeEngineDomainMappingDataSourceConfigBasic(projectID, appName, string(domainMappingTLSKey), string(domainMappingTLSCert), secretName, domainMappingName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "cname_target"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "user_managed"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "visibility"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "name", domainMappingName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "resource_type", "domain_mapping_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "component.0.resource_type", "app_v2"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "component.0.name", appName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "tls_secret", secretName),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineDomainMappingDataSourceConfigBasic(projectID string, appName string, tlsKey string, tlsCert string, secretName string, domainMappingName string) string {
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

		variable "tls_secret_data" {
			type = map(string)
  			default = {
   				"tls_key" = <<EOT
%s
EOT
				"tls_cert" = <<EOT
%s
EOT
			}
		}

		resource "ibm_code_engine_secret" "code_engine_secret_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			format = "tls"
			name = "%s"
			data = var.tls_secret_data
		}

	 	resource "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			component {
				name = ibm_code_engine_app.code_engine_app_instance.name
				resource_type = "app_v2"
			}
			name = "%s"
			tls_secret = ibm_code_engine_secret.code_engine_secret_instance.name

			depends_on = [
    			ibm_code_engine_app.code_engine_app_instance
  			]
		}

		data "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
			project_id = ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance.project_id
			name = ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance.name
		}

	`, projectID, appName, tlsKey, tlsCert, secretName, domainMappingName)
}
