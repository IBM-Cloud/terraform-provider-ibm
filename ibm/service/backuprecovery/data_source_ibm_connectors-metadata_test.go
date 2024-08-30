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

func TestAccIbmConnectorsMetadataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConnectorsMetadataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_connectors-metadata.connectors_metadata_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmConnectorsMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_connectors-metadata" "connectors_metadata_instance" {
		}
	`)
}

func TestDataSourceIbmConnectorsMetadataConnectorImageMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		connectorImageFileModel := make(map[string]interface{})
		connectorImageFileModel["image_type"] = "VSI"
		connectorImageFileModel["url"] = "testString"

		model := make(map[string]interface{})
		model["connector_image_file_list"] = []map[string]interface{}{connectorImageFileModel}

		assert.Equal(t, result, model)
	}

	connectorImageFileModel := new(backuprecoveryv1.ConnectorImageFile)
	connectorImageFileModel.ImageType = core.StringPtr("VSI")
	connectorImageFileModel.URL = core.StringPtr("testString")

	model := new(backuprecoveryv1.ConnectorImageMetadata)
	model.ConnectorImageFileList = []backuprecoveryv1.ConnectorImageFile{*connectorImageFileModel}

	result, err := backuprecovery.DataSourceIbmConnectorsMetadataConnectorImageMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmConnectorsMetadataConnectorImageFileToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["image_type"] = "VSI"
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ConnectorImageFile)
	model.ImageType = core.StringPtr("VSI")
	model.URL = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmConnectorsMetadataConnectorImageFileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
