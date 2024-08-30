// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmConnectorsMetadataBasic(t *testing.T) {
	var conf backuprecoveryv1.ConnectorMetadata

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmConnectorsMetadataDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConnectorsMetadataConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmConnectorsMetadataExists("ibm_connectors-metadata.connectors_metadata_instance", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_connectors-metadata.connectors_metadata",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmConnectorsMetadataConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_connectors-metadata" "connectors_metadata_instance" {
		}
	`)
}

func testAccCheckIbmConnectorsMetadataExists(n string, obj backuprecoveryv1.ConnectorMetadata) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getConnectorMetadataOptions := &backuprecoveryv1.GetConnectorMetadataOptions{}

		connectorMetadata, _, err := backupRecoveryClient.GetConnectorMetadata(getConnectorMetadataOptions)
		if err != nil {
			return err
		}

		obj = *connectorMetadata
		return nil
	}
}

func testAccCheckIbmConnectorsMetadataDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_connectors-metadata" {
			continue
		}

		getConnectorMetadataOptions := &backuprecoveryv1.GetConnectorMetadataOptions{}

		// Try to find the key
		_, response, err := backupRecoveryClient.GetConnectorMetadata(getConnectorMetadataOptions)

		if err == nil {
			return fmt.Errorf("connectors-metadata still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for connectors-metadata (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmConnectorsMetadataConnectorImageMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmConnectorsMetadataConnectorImageMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConnectorsMetadataConnectorImageFileToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["image_type"] = "VSI"
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ConnectorImageFile)
	model.ImageType = core.StringPtr("VSI")
	model.URL = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmConnectorsMetadataConnectorImageFileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConnectorsMetadataMapToConnectorMetadata(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ConnectorMetadata) {
		connectorImageFileModel := new(backuprecoveryv1.ConnectorImageFile)
		connectorImageFileModel.ImageType = core.StringPtr("VSI")
		connectorImageFileModel.URL = core.StringPtr("testString")

		connectorImageMetadataModel := new(backuprecoveryv1.ConnectorImageMetadata)
		connectorImageMetadataModel.ConnectorImageFileList = []backuprecoveryv1.ConnectorImageFile{*connectorImageFileModel}

		model := new(backuprecoveryv1.ConnectorMetadata)
		model.ConnectorImageMetadata = connectorImageMetadataModel

		assert.Equal(t, result, model)
	}

	connectorImageFileModel := make(map[string]interface{})
	connectorImageFileModel["image_type"] = "VSI"
	connectorImageFileModel["url"] = "testString"

	connectorImageMetadataModel := make(map[string]interface{})
	connectorImageMetadataModel["connector_image_file_list"] = []interface{}{connectorImageFileModel}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["connector_image_metadata"] = []interface{}{connectorImageMetadataModel}

	result, err := backuprecovery.ResourceIbmConnectorsMetadataMapToConnectorMetadata(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConnectorsMetadataMapToConnectorImageMetadata(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ConnectorImageMetadata) {
		connectorImageFileModel := new(backuprecoveryv1.ConnectorImageFile)
		connectorImageFileModel.ImageType = core.StringPtr("VSI")
		connectorImageFileModel.URL = core.StringPtr("testString")

		model := new(backuprecoveryv1.ConnectorImageMetadata)
		model.ConnectorImageFileList = []backuprecoveryv1.ConnectorImageFile{*connectorImageFileModel}

		assert.Equal(t, result, model)
	}

	connectorImageFileModel := make(map[string]interface{})
	connectorImageFileModel["image_type"] = "VSI"
	connectorImageFileModel["url"] = "testString"

	model := make(map[string]interface{})
	model["connector_image_file_list"] = []interface{}{connectorImageFileModel}

	result, err := backuprecovery.ResourceIbmConnectorsMetadataMapToConnectorImageMetadata(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmConnectorsMetadataMapToConnectorImageFile(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ConnectorImageFile) {
		model := new(backuprecoveryv1.ConnectorImageFile)
		model.ImageType = core.StringPtr("VSI")
		model.URL = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["image_type"] = "VSI"
	model["url"] = "testString"

	result, err := backuprecovery.ResourceIbmConnectorsMetadataMapToConnectorImageFile(model)
	assert.Nil(t, err)
	checkResult(result)
}
