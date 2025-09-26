// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmBackupRecoveryManagerGetResourcesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetResourcesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_resources.backup_recovery_manager_get_resources_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_resources.backup_recovery_manager_get_resources_instance", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetResourcesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_resources" "backup_recovery_manager_get_resources_instance" {
			resourceType = "Policies"
		}
	`)
}

func TestDataSourceIbmBackupRecoveryManagerGetResourcesExternalTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["system_id"] = "testString"
		model["system_name"] = "testString"
		model["target_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalTarget)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.SystemID = core.StringPtr("testString")
	model.SystemName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetResourcesExternalTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetResourcesMessageCodeMappingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["message_code"] = "testString"
		model["message_guid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MessageCodeMapping)
	model.MessageCode = core.StringPtr("testString")
	model.MessageGuid = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetResourcesMessageCodeMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetResourcesPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["is_global_policy"] = true
		model["name"] = "testString"
		model["system_id"] = "testString"
		model["system_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Policy)
	model.ID = core.StringPtr("testString")
	model.IsGlobalPolicy = core.BoolPtr(true)
	model.Name = core.StringPtr("testString")
	model.SystemID = core.StringPtr("testString")
	model.SystemName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetResourcesPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetResourcesProtectionGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["system_id"] = "testString"
		model["system_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ProtectionGroup)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.SystemID = core.StringPtr("testString")
	model.SystemName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetResourcesProtectionGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetResourcesRegisteredSourceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environments"] = []string{"testString"}
		model["name"] = "testString"
		model["uuid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RegisteredSource)
	model.Environments = []string{"testString"}
	model.Name = core.StringPtr("testString")
	model.UUID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetResourcesRegisteredSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryManagerGetResourcesTenantToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Tenant)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryManagerGetResourcesTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
