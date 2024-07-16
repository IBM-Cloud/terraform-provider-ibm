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

func TestAccIbmCodeEngineSecretGeneric(t *testing.T) {
	var conf codeenginev2.Secret
	format := "generic"
	name := fmt.Sprintf("tf-secret-generic-%d", acctest.RandIntRange(10, 1000))
	data := `{ "key" = "value" }`
	nameUpdate := fmt.Sprintf("tf-secret-generic-update-%d", acctest.RandIntRange(10, 1000))
	dataUpdate := `{
		"key1" = "value1"
		"key2" = "value2"
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, name, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineSecretExists("ibm_code_engine_secret.code_engine_secret_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.key", "value"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_generic_v2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, nameUpdate, dataUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.key1", "value1"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.key2", "value2"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_generic_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_secret.code_engine_secret_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmCodeEngineSecretBasicAuth(t *testing.T) {
	var conf codeenginev2.Secret
	format := "basic_auth"
	name := fmt.Sprintf("tf-secret-basic-auth-%d", acctest.RandIntRange(10, 1000))
	data := `{
		"username" = "name"
		"password" = "password"
	}`
	nameUpdate := fmt.Sprintf("tf-secret-basic-auth-update-%d", acctest.RandIntRange(10, 1000))
	dataUpdate := `{
		"username" = "name_update"
		"password" = "password_update"
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, name, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineSecretExists("ibm_code_engine_secret.code_engine_secret_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.username", "name"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.password", "password"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_basic_auth_v2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, nameUpdate, dataUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.username", "name_update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.password", "password_update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_basic_auth_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_secret.code_engine_secret_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmCodeEngineSecretRegistry(t *testing.T) {
	var conf codeenginev2.Secret
	format := "registry"
	name := fmt.Sprintf("tf-secret-registry-%d", acctest.RandIntRange(10, 1000))
	data := `{
		"username" = "foo.bar"
    "password" = "foouser"
    "server"   = "foopass"
    "email"    = "foo@mail.com"
	}`
	nameUpdate := fmt.Sprintf("tf-secret-registry-update-%d", acctest.RandIntRange(10, 1000))
	dataUpdate := `{
		"username" = "foo.update"
    "password" = "foouser_update"
    "server"   = "foopass_update"
    "email"    = "foo@mail.co.uk"
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, name, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineSecretExists("ibm_code_engine_secret.code_engine_secret_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.username", "foo.bar"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.password", "foouser"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.server", "foopass"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.email", "foo@mail.com"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_registry_v2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, nameUpdate, dataUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.username", "foo.update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.password", "foouser_update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.server", "foopass_update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.email", "foo@mail.co.uk"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_registry_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_secret.code_engine_secret_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmCodeEngineSecretSSHAuth(t *testing.T) {
	var conf codeenginev2.Secret
	format := "ssh_auth"
	name := fmt.Sprintf("tf-secret-ssh-auth-%d", acctest.RandIntRange(10, 1000))
	data := `{
		"ssh_key"     = "ssh-key",
    "known_hosts" = "knownhosts",
	}`
	nameUpdate := fmt.Sprintf("tf-secret-ssh-auth-update-%d", acctest.RandIntRange(10, 1000))
	dataUpdate := `{
		"ssh_key"     = "ssh-key-update",
    "known_hosts" = "knownhosts-update",
	}`

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, name, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineSecretExists("ibm_code_engine_secret.code_engine_secret_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.ssh_key", "ssh-key"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.known_hosts", "knownhosts"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_auth_ssh_v2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretConfig(projectID, format, nameUpdate, dataUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.ssh_key", "ssh-key-update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "data.known_hosts", "knownhosts-update"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_auth_ssh_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_secret.code_engine_secret_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmCodeEngineSecretTls(t *testing.T) {
	var conf codeenginev2.Secret
	format := "tls"
	name := fmt.Sprintf("tf-secret-tls-%d", acctest.RandIntRange(10, 1000))
	nameUpdate := fmt.Sprintf("tf-secret-tls-update-%d", acctest.RandIntRange(10, 1000))
	tlsKey, _ := os.ReadFile(acc.CeTLSKeyFilePath)
	tlsCert, _ := os.ReadFile(acc.CeTLSCertFilePath)

	projectID := acc.CeProjectId

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCodeEngine(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretTLSConfig(projectID, string(tlsKey), string(tlsCert), format, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineSecretExists("ibm_code_engine_secret.code_engine_secret_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", name),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "data.tls_cert"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "data.tls_key"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_tls_v2"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineSecretTLSConfig(projectID, string(tlsKey), string(tlsCert), format, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", nameUpdate),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "data.tls_cert"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "data.tls_key"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_tls_v2"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_secret.code_engine_secret_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIbmCodeEngineSecretServiceAccess(t *testing.T) {
	var conf codeenginev2.Secret
	format := "service_access"
	name := fmt.Sprintf("tf-secret-service-access-%d", acctest.RandIntRange(10, 1000))

	projectID := acc.CeProjectId
	resourceKeyId := acc.CeResourceKeyID
	serviceInstanceId := acc.CeServiceInstanceID

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineServiceAccessSecretConfig(projectID, format, name, resourceKeyId, serviceInstanceId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineSecretExists("ibm_code_engine_secret.code_engine_secret_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "format", format),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "resource_type", "secret_service_access_v2"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "service_access.0.service_instance.0.id", serviceInstanceId),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "service_access.0.service_instance.0.type"),
					resource.TestCheckResourceAttr("ibm_code_engine_secret.code_engine_secret_instance", "service_access.0.resource_key.0.id", resourceKeyId),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "service_access.0.resource_key.0.name"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_secret.code_engine_secret_instance", "service_access.0.role.0.name"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_secret.code_engine_secret_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmCodeEngineSecretConfig(projectID string, format string, name string, data string) string {
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
	`, projectID, format, name, data)
}

func testAccCheckIbmCodeEngineSecretTLSConfig(projectID string, tlsKey string, tlsCert string, secretFormat string, secretName string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
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
			format = "%s"
			name = "%s"
			data = var.tls_secret_data
		}
	`, projectID, tlsKey, tlsCert, secretFormat, secretName)
}

func testAccCheckIbmCodeEngineServiceAccessSecretConfig(projectID string, format string, name string, resourceKeyId string, serviceInstanceId string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_secret" "code_engine_secret_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			format = "%s"
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
	`, projectID, format, name, resourceKeyId, serviceInstanceId)
}

func testAccCheckIbmCodeEngineSecretExists(n string, obj codeenginev2.Secret) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getSecretOptions := &codeenginev2.GetSecretOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSecretOptions.SetProjectID(parts[0])
		getSecretOptions.SetName(parts[1])

		secret, _, err := codeEngineClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		obj = *secret
		return nil
	}
}

func testAccCheckIbmCodeEngineSecretDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_secret" {
			continue
		}

		getSecretOptions := &codeenginev2.GetSecretOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSecretOptions.SetProjectID(parts[0])
		getSecretOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("code_engine_secret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_secret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
