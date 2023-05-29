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

func TestAccIbmCodeEngineSecretDataSourceGeneric(t *testing.T) {
	secretFormat := "generic"
	secretName := fmt.Sprintf("tf-data-secret-generic-%d", acctest.RandIntRange(10, 1000))
	secretData := `{ "key" = "value" }`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretDataSourceConfigBasic(projectID, secretFormat, secretName, secretData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "format", secretFormat),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "name", secretName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.key", "value"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_generic_v2"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineSecretDataSourceBasicAuth(t *testing.T) {
	secretFormat := "basic_auth"
	secretName := fmt.Sprintf("tf-data-secret-basic-auth-%d", acctest.RandIntRange(10, 1000))
	secretData := `{
		"username" = "name"
		"password" = "password"
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretDataSourceConfigBasic(projectID, secretFormat, secretName, secretData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "format", secretFormat),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "name", secretName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.username", "name"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.password", "password"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_basic_auth_v2"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineSecretDataSourceRegistry(t *testing.T) {
	secretFormat := "registry"
	secretName := fmt.Sprintf("tf-data-secret-registry-%d", acctest.RandIntRange(10, 1000))
	secretData := `{
		"username" = "foo.bar"
    "password" = "foouser"
    "server"   = "foopass"
    "email"    = "foo@mail.com"
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretDataSourceConfigBasic(projectID, secretFormat, secretName, secretData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "format", secretFormat),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "name", secretName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.username", "foo.bar"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.password", "foouser"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.server", "foopass"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.email", "foo@mail.com"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_registry_v2"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineSecretDataSourceSSHAuth(t *testing.T) {
	secretFormat := "ssh_auth"
	secretName := fmt.Sprintf("tf-data-secret-ssh-auth-%d", acctest.RandIntRange(10, 1000))
	secretData := `{
		"ssh_key"     = "ssh-key",
    "known_hosts" = "knownhosts",
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretDataSourceConfigBasic(projectID, secretFormat, secretName, secretData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "format", secretFormat),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "name", secretName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.ssh_key", "ssh-key"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.known_hosts", "knownhosts"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_auth_ssh_v2"),
				),
			},
		},
	})
}

func TestAccIbmCodeEngineSecretDataSourceTls(t *testing.T) {
	secretFormat := "tls"
	secretName := fmt.Sprintf("tf-data-secret-tls-%d", acctest.RandIntRange(10, 1000))
	secretData := `{
		"tls_cert" = "---BEGIN CERTIFICATE--- ---END CERTIFICATE---",
    "tls_key"  = "---BEGIN PRIVATE KEY--- ---END PRIVATE KEY---",
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretDataSourceConfigBasic(projectID, secretFormat, secretName, secretData),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "format", secretFormat),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "name", secretName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.tls_cert", "---BEGIN CERTIFICATE--- ---END CERTIFICATE---"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "data.tls_key", "---BEGIN PRIVATE KEY--- ---END PRIVATE KEY---"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_tls_v2"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineSecretDataSourceConfigBasic(projectID string, secretFormat string, secretName string, secretData string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_secret" "code_engine_secret_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			format = "%s"
			name = "%s"
			data = %s
		}

		data "ibm_code_engine_secret" "code_engine_secret_instance" {
			project_id = ibm_code_engine_secret.code_engine_secret_instance.project_id
			name = ibm_code_engine_secret.code_engine_secret_instance.name
		}
	`, projectID, secretFormat, secretName, secretData)
}
