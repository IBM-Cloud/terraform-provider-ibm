// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package codeengine_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/codeengine"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmCodeEnginePersistentDataStoreDataSourceBasic(t *testing.T) {
	projectID := acc.CeProjectId
	secretName := fmt.Sprintf("tf-hmac-secret-%d", acctest.RandIntRange(10, 1000))
	cosAccessKeyID := acc.CeCosAccessKeyID
	cosSecretAccessKey := acc.CeCosSecretAccessKey
	pdsName := fmt.Sprintf("tf-pds-%d", acctest.RandIntRange(10, 1000))
	cosBucketName := acc.CeCosBucketName
	cosBucketLocation := acc.CeCosBucketLocation

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEnginePersistentDataStoreDataSourceConfigBasic(projectID, secretName, cosAccessKeyID, cosSecretAccessKey, pdsName, cosBucketName, cosBucketLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "region"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "created_at"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "name", pdsName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "storage_type", "object_storage"),
					resource.TestCheckResourceAttr("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "data.0.secret_name", secretName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "data.0.bucket_name", cosBucketName),
					resource.TestCheckResourceAttr("data.ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance", "data.0.bucket_location", cosBucketLocation),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEnginePersistentDataStoreDataSourceConfigBasic(projectID string, secretName string, cosAccessKeyID string, cosSecretAccessKey string, pdsName string, cosBucketName string, cosBucketLocation string) string {
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

		data "ibm_code_engine_persistent_data_store" "code_engine_persistent_data_store_instance" {
			project_id = ibm_code_engine_secret.my-secret.project_id
			name = ibm_code_engine_persistent_data_store.code_engine_persistent_data_store_instance.name
		}
	`, projectID, secretName, cosAccessKeyID, cosSecretAccessKey, pdsName, cosBucketName, cosBucketLocation)
}

func TestDataSourceIbmCodeEnginePersistentDataStoreStorageDataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["bucket_location"] = "eu-de"
		model["bucket_name"] = "my-bucket"
		model["secret_name"] = "my-secret"

		assert.Equal(t, result, model)
	}

	model := new(codeenginev2.StorageData)
	model.BucketLocation = core.StringPtr("eu-de")
	model.BucketName = core.StringPtr("my-bucket")
	model.SecretName = core.StringPtr("my-secret")

	result, err := codeengine.DataSourceIbmCodeEnginePersistentDataStoreStorageDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmCodeEnginePersistentDataStoreStorageDataObjectStorageDataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["bucket_location"] = "au-syd"
		model["bucket_name"] = "testString"
		model["secret_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(codeenginev2.StorageDataObjectStorageData)
	model.BucketLocation = core.StringPtr("au-syd")
	model.BucketName = core.StringPtr("testString")
	model.SecretName = core.StringPtr("testString")

	result, err := codeengine.DataSourceIbmCodeEnginePersistentDataStoreStorageDataObjectStorageDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
