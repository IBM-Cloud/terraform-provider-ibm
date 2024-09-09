// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasConnectorStatusDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasConnectorStatusDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_connector_status.baas_connector_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_connector_status.baas_connector_status_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasConnectorStatusDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_connector_status" "baas_connector_status_instance" {
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
		}
	`)
}

func TestDataSourceIbmBaasConnectorStatusDataSourceConnectorClusterConnectionStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_active"] = true
		model["last_connected_timestamp_msecs"] = int(26)
		model["message"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DataSourceConnectorClusterConnectionStatus)
	model.IsActive = core.BoolPtr(true)
	model.LastConnectedTimestampMsecs = core.Int64Ptr(int64(26))
	model.Message = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasConnectorStatusDataSourceConnectorClusterConnectionStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasConnectorStatusDataSourceConnectorRegistrationStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["message"] = "testString"
		model["status"] = "NotDone"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DataSourceConnectorRegistrationStatus)
	model.Message = core.StringPtr("testString")
	model.Status = core.StringPtr("NotDone")

	result, err := backuprecovery.DataSourceIbmBaasConnectorStatusDataSourceConnectorRegistrationStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
