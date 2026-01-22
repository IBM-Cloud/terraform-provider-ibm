// Copyright IBM Corp. 2025 All Rights Reserved.
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

func TestAccIbmCodeEnginePersistentDataStoreBasic(t *testing.T) {
	var conf codeenginev2.PersistentDataStore
	projectID := acc.CeProjectId
	secretName := fmt.Sprintf("tf-hmac-secret-%d", acctest.RandIntRange(10, 1000))
	cosAccessKeyID := acc.CeCosAccessKeyID
	cosSecretAccessKey := acc.CeCosSecretAccessKey
	pdsName := fmt.Sprintf("tf-pds-%d", acctest.RandIntRange(10, 1000))
	cosBucketName := acc.CeCosBucketName
	cosBucketLocation := acc.CeCosBucketLocation
	appName := fmt.Sprintf("tf-app-%d", acctest.RandIntRange(10, 1000))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEnginePersistentDataStoreDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEnginePersistentDataStoreConfigBasic(projectID, secretName, cosAccessKeyID, cosSecretAccessKey, pdsName, cosBucketName, cosBucketLocation, appName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEnginePersistentDataStoreExists("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", conf),
					resource.TestCheckResourceAttrSet("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "region"),
					resource.TestCheckResourceAttr("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "name", pdsName),
					resource.TestCheckResourceAttr("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "storage_type", "object_storage"),
					resource.TestCheckResourceAttr("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "data.0.secret_name", secretName),
					resource.TestCheckResourceAttr("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "data.0.bucket_name", cosBucketName),
					resource.TestCheckResourceAttr("ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "data.0.bucket_location", cosBucketLocation),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.#", "1"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.0.mount_path", "/foo"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.0.type", "persistent_data_store"),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.0.reference", pdsName),
					resource.TestCheckResourceAttr("ibm_code_engine_app.code_engine_app_instance", "run_volume_mounts.0.read_only", "true"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmCodeEnginePersistentDataStoreConfigBasic(projectID string, secretName string, cosAccessKeyID string, cosSecretAccessKey string, pdsName string, cosBucketName string, cosBucketLocation string, appName string) string {
	return fmt.Sprintf(`
		resource "ibm_code_engine_secret" "my-secret" {
			project_id = "%s"
			name       = "%s"
			format     = "hmac_auth"
			data = {
				"access_key_id"     = "%s"
				"secret_access_key" = "%s"
			}
		}

		resource "ibm_code_engine_persistent_data_store" "code_engine_persistent_data_store_instance" {
			project_id = ibm_code_engine_secret.my-secret.project_id
			name = "%s"
			storage_type = "object_storage"
			data {
				bucket_name     = "%s"
				bucket_location = "%s"
				secret_name     = ibm_code_engine_secret.my-secret.name
			}
		}

		resource "ibm_code_engine_app" "code_engine_app_instance" {
			project_id = ibm_code_engine_secret.my-secret.project_id
			image_reference = "icr.io/codeengine/helloworld"
			name = "%s"

			run_volume_mounts {
				mount_path = "/foo"
				type       = "persistent_data_store"
				reference  = ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance.name
				read_only  = true
			}

			lifecycle {
				ignore_changes = [
					probe_liveness,
					probe_readiness
				]
			}
		}
	`, projectID, secretName, cosAccessKeyID, cosSecretAccessKey, pdsName, cosBucketName, cosBucketLocation, appName)
}

func testAccCheckIbmCodeEnginePersistentDataStoreExists(n string, obj codeenginev2.PersistentDataStore) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getPersistentDataStoreOptions := &codeenginev2.GetPersistentDataStoreOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPersistentDataStoreOptions.SetProjectID(parts[0])
		getPersistentDataStoreOptions.SetName(parts[1])

		persistentDataStore, _, err := codeEngineClient.GetPersistentDataStore(getPersistentDataStoreOptions)
		if err != nil {
			return err
		}

		obj = *persistentDataStore
		return nil
	}
}

func testAccCheckIbmCodeEnginePersistentDataStoreDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_persistent_data_store" {
			continue
		}

		getPersistentDataStoreOptions := &codeenginev2.GetPersistentDataStoreOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPersistentDataStoreOptions.SetProjectID(parts[0])
		getPersistentDataStoreOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetPersistentDataStore(getPersistentDataStoreOptions)

		if err == nil {
			return fmt.Errorf("code_engine_persistent_data_store still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_persistent_data_store (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
