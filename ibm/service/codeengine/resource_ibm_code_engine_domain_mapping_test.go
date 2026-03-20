// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
)

func TestAccIbmCodeEngineDomainMappingBasic(t *testing.T) {
	var conf codeenginev2.DomainMapping

	app1Name := fmt.Sprintf("tf-app-domain-mapping-%d", acctest.RandIntRange(10, 1000))
	app2Name := fmt.Sprintf("tf-app-domain-mapping-%d", acctest.RandIntRange(10, 1000))
	secretName := fmt.Sprintf("tf-secret-domain-mapping-%d", acctest.RandIntRange(10, 1000))

	projectID := acc.CeProjectId
	domainMappingName := acc.CeDomainMappingName
	domainMappingTLSKey, _ := os.ReadFile(acc.CeTLSKeyFilePath)
	domainMappingTLSCert, _ := os.ReadFile(acc.CeTLSCertFilePath)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCodeEngine(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineDomainMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCodeEngineDomainMappingConfigBasic(projectID, app1Name, app2Name, string(domainMappingTLSKey), string(domainMappingTLSCert), secretName, app1Name, domainMappingName, "app1_instance"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineDomainMappingExists("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "status"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "cname_target"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "user_managed"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "visibility"),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "name", domainMappingName),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "resource_type", "domain_mapping_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "component.0.resource_type", "app_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "component.0.name", app1Name),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "tls_secret", secretName),
				),
			},
			{
				Config: testAccCheckIbmCodeEngineDomainMappingConfigBasic(projectID, app1Name, app2Name, string(domainMappingTLSKey), string(domainMappingTLSCert), secretName, app2Name, domainMappingName, "app2_instance"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "href"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "status"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "cname_target"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "user_managed"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "visibility"),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "name", domainMappingName),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "resource_type", "domain_mapping_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "component.0.resource_type", "app_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "component.0.name", app2Name),
					resource.TestCheckResourceAttr("ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance", "tls_secret", secretName),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineDomainMappingConfigBasic(projectID string, app1Name string, app2Name string, tlsKey string, tlsCert string, secretName string, componentRefName string, domainMappingName string, dependsOn string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_app" "code_engine_app1_instance" {
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

		resource "ibm_code_engine_app" "code_engine_app2_instance" {
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
				name = "%s"
				resource_type = "app_v2"
			}
			name = "%s"
			tls_secret = ibm_code_engine_secret.code_engine_secret_instance.name

			depends_on = [
    			ibm_code_engine_app.code_engine_%s
  			]
		}
	`, projectID, app1Name, app2Name, tlsKey, tlsCert, secretName, componentRefName, domainMappingName, dependsOn)
}

func testAccCheckIbmCodeEngineDomainMappingExists(n string, obj codeenginev2.DomainMapping) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getDomainMappingOptions := &codeenginev2.GetDomainMappingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDomainMappingOptions.SetProjectID(parts[0])
		getDomainMappingOptions.SetName(parts[1])

		domainMapping, _, err := codeEngineClient.GetDomainMapping(getDomainMappingOptions)
		if err != nil {
			return err
		}

		obj = *domainMapping
		return nil
	}
}

func testAccCheckIbmCodeEngineDomainMappingDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_domain_mapping" {
			continue
		}

		getDomainMappingOptions := &codeenginev2.GetDomainMappingOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDomainMappingOptions.SetProjectID(parts[0])
		getDomainMappingOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetDomainMapping(getDomainMappingOptions)

		if err == nil {
			return fmt.Errorf("code_engine_domain_mapping still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_domain_mapping (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
