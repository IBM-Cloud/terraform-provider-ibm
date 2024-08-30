// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmConnectorsMetadata() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmConnectorsMetadataRead,

		Schema: map[string]*schema.Schema{
			"connector_image_metadata": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies information about the connector images for various platforms.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connector_image_file_list": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies info about connector images for the supported platforms.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"image_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the platform on which the image can be deployed.",
									},
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the URL to access the file.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmConnectorsMetadataRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_connectors-metadata", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getConnectorMetadataOptions := &backuprecoveryv1.GetConnectorMetadataOptions{}

	connectorMetadata, _, err := backupRecoveryClient.GetConnectorMetadataWithContext(context, getConnectorMetadataOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConnectorMetadataWithContext failed: %s", err.Error()), "(Data) ibm_connectors-metadata", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmConnectorsMetadataID(d))

	if !core.IsNil(connectorMetadata.ConnectorImageMetadata) {
		connectorImageMetadata := []map[string]interface{}{}
		connectorImageMetadataMap, err := DataSourceIbmConnectorsMetadataConnectorImageMetadataToMap(connectorMetadata.ConnectorImageMetadata)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_connectors-metadata", "read", "connector_image_metadata-to-map").GetDiag()
		}
		connectorImageMetadata = append(connectorImageMetadata, connectorImageMetadataMap)
		if err = d.Set("connector_image_metadata", connectorImageMetadata); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connector_image_metadata: %s", err), "(Data) ibm_connectors-metadata", "read", "set-connector_image_metadata").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmConnectorsMetadataID returns a reasonable ID for the list.
func dataSourceIbmConnectorsMetadataID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmConnectorsMetadataConnectorImageMetadataToMap(model *backuprecoveryv1.ConnectorImageMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	connectorImageFileList := []map[string]interface{}{}
	for _, connectorImageFileListItem := range model.ConnectorImageFileList {
		connectorImageFileListItemMap, err := DataSourceIbmConnectorsMetadataConnectorImageFileToMap(&connectorImageFileListItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		connectorImageFileList = append(connectorImageFileList, connectorImageFileListItemMap)
	}
	modelMap["connector_image_file_list"] = connectorImageFileList
	return modelMap, nil
}

func DataSourceIbmConnectorsMetadataConnectorImageFileToMap(model *backuprecoveryv1.ConnectorImageFile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["image_type"] = *model.ImageType
	modelMap["url"] = *model.URL
	return modelMap, nil
}
