// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.4-9b56d441-20260612-210048
 */

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmSmInstanceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmInstanceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "instance_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "plan"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "vault_cluster.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "endpoints.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_instance.sm_instance_instance", "encryption.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmInstanceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_sm_instance" "sm_instance_instance" {
			instance_id = "60b40daa-1fd3-4f35-a994-2409cc0f270c"
		}
	`)
}

func TestDataSourceIbmSmInstanceVaultDedicatedClusterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "healthy"
		model["version"] = "1.21.2+ent.hsm"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.VaultDedicatedCluster)
	model.Status = core.StringPtr("healthy")
	model.Version = core.StringPtr("1.21.2+ent.hsm")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmInstanceVaultDedicatedClusterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmSmInstanceVaultDedicatedInstanceEndpointsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vaultDedicatedEndpointsDataModel := make(map[string]interface{})
		vaultDedicatedEndpointsDataModel["vault_api"] = "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud"
		vaultDedicatedEndpointsDataModel["vault_ui"] = "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"

		model := make(map[string]interface{})
		model["public"] = []map[string]interface{}{vaultDedicatedEndpointsDataModel}
		model["private"] = []map[string]interface{}{vaultDedicatedEndpointsDataModel}

		assert.Equal(t, result, model)
	}

	vaultDedicatedEndpointsDataModel := new(secretsmanagerinstancemanagementv2.VaultDedicatedEndpointsData)
	vaultDedicatedEndpointsDataModel.VaultApi = core.StringPtr("https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud")
	vaultDedicatedEndpointsDataModel.VaultUi = core.StringPtr("https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui")

	model := new(secretsmanagerinstancemanagementv2.VaultDedicatedInstanceEndpoints)
	model.Public = vaultDedicatedEndpointsDataModel
	model.Private = vaultDedicatedEndpointsDataModel

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmInstanceVaultDedicatedInstanceEndpointsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmSmInstanceVaultDedicatedEndpointsDataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["vault_api"] = "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud"
		model["vault_ui"] = "https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.VaultDedicatedEndpointsData)
	model.VaultApi = core.StringPtr("https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud")
	model.VaultUi = core.StringPtr("https://f85f512b-e21b-4a9a-ac45-7bbc2f5cew2e.us-south.secrets-manager.appdomain.cloud/ui")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmInstanceVaultDedicatedEndpointsDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmSmInstanceVaultDedicatedInstanceEncryptionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["mode"] = "service_managed"
		model["provider"] = "key_protect"
		model["key_crn"] = "crn:v1:bluemix:public:kms:us-south:a/791f5fb10986423e97aa8512f18b7e65:31639268-42e8-4420-9872-590a6ee20506:key:b4af8f76-e6ea-4dc5-89cc-5f1b9bb207cc"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.VaultDedicatedInstanceEncryption)
	model.Mode = core.StringPtr("service_managed")
	model.Provider = core.StringPtr("key_protect")
	model.KeyCrn = core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/791f5fb10986423e97aa8512f18b7e65:31639268-42e8-4420-9872-590a6ee20506:key:b4af8f76-e6ea-4dc5-89cc-5f1b9bb207cc")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmInstanceVaultDedicatedInstanceEncryptionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
