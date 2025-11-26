// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/db2"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDb2BackupDataSourceBasic(t *testing.T) {
	xDbProfile := "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3Af5c9359c-9a66-4087-9eda-0024c2603c92%3A%3A"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2BackupDataSourceConfigBasic(xDbProfile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_backup.Db2-44-test-both", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_backup.Db2-44-test-both", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_backup.Db2-44-test-both", "backups.#"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2BackupDataSourceConfigBasic(xDbProfile string) string {
	return fmt.Sprintf(`
		data "ibm_db2_backup" "Db2-44-test-both" {
			deployment_id = "%[1]s"
		}
	`, xDbProfile)
}

func TestDataSourceIbmDb2BackupBackupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:84792aeb-2a9c-4dee-bfad-2e529f16945d::"
		model["type"] = "on_demand"
		model["status"] = "running"
		model["created_at"] = "2025-01-27T08:43:46.000Z"
		model["size"] = 4000000000
		model["duration"] = 131

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.Backup)
	model.ID = core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:84792aeb-2a9c-4dee-bfad-2e529f16945d::")
	model.Type = core.StringPtr("on_demand")
	model.Status = core.StringPtr("running")
	model.CreatedAt = core.StringPtr("2025-01-27T08:43:46.000Z")
	model.Size = core.Int64Ptr(4000000000)
	model.Duration = core.Int64Ptr(131)

	result, err := db2.DataSourceIbmDb2BackupBackupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
