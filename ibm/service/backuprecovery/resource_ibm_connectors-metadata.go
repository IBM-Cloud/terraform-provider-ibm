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

func ResourceIbmConnectorsMetadata() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmConnectorsMetadataCreate,
		ReadContext:   resourceIbmConnectorsMetadataRead,
		DeleteContext: resourceIbmConnectorsMetadataDelete,
		UpdateContext: resourceIbmConnectorsMetadataUpdate,
		Importer:      &schema.ResourceImporter{},

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

func resourceIbmConnectorsMetadataCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connectors-metadata", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateConnectorMetadataOptions := &backuprecoveryv1.UpdateConnectorMetadataOptions{}

	connectorMetadataModel, err := ResourceIbmConnectorsMetadataMapToConnectorMetadata(d.Get("connector_metadata.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connectors-metadata", "create", "parse-connector_metadata").GetDiag()
	}
	updateConnectorMetadataOptions.SetConnectorMetadata(connectorMetadataModel)

	_, _, err = backupRecoveryClient.UpdateConnectorMetadataWithContext(context, updateConnectorMetadataOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateConnectorMetadataWithContext failed: %s", err.Error()), "ibm_connectors-metadata", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmConnectorsMetadataID(d))

	return resourceIbmConnectorsMetadataRead(context, d, meta)
}

func resourceIbmConnectorsMetadataID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmConnectorsMetadataRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connectors-metadata", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getConnectorMetadataOptions := &backuprecoveryv1.GetConnectorMetadataOptions{}

	connectorMetadata, response, err := backupRecoveryClient.GetConnectorMetadataWithContext(context, getConnectorMetadataOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConnectorMetadataWithContext failed: %s", err.Error()), "ibm_connectors-metadata", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(connectorMetadata.ConnectorImageMetadata) {
		connectorImageMetadataMap, err := ResourceIbmConnectorsMetadataConnectorImageMetadataToMap(connectorMetadata.ConnectorImageMetadata)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connectors-metadata", "read", "connector_image_metadata-to-map").GetDiag()
		}
		if err = d.Set("connector_image_metadata", []map[string]interface{}{connectorImageMetadataMap}); err != nil {
			err = fmt.Errorf("Error setting connector_image_metadata: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_connectors-metadata", "read", "set-connector_image_metadata").GetDiag()
		}
	}

	return nil
}

func resourceIbmConnectorsMetadataDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "Delete operation is not supported for this resource. The resource will be removed from the terraform file but will continue to exist in the backend.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmConnectorsMetadataUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update Not Supported",
		Detail:   "Update operation is not supported for this resource. No changes will be applied.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func ResourceIbmConnectorsMetadataMapToConnectorMetadata(modelMap map[string]interface{}) (*backuprecoveryv1.ConnectorMetadata, error) {
	model := &backuprecoveryv1.ConnectorMetadata{}

	if modelMap["connector_image_metadata"] != nil && len(modelMap["connector_image_metadata"].([]interface{})) > 0 {
		ConnectorImageMetadataModel, err := ResourceIbmConnectorsMetadataMapToConnectorImageMetadata(modelMap["connector_image_metadata"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ConnectorImageMetadata = ConnectorImageMetadataModel
	}
	return model, nil
}

func ResourceIbmConnectorsMetadataMapToConnectorImageMetadata(modelMap map[string]interface{}) (*backuprecoveryv1.ConnectorImageMetadata, error) {
	model := &backuprecoveryv1.ConnectorImageMetadata{}
	connectorImageFileList := []backuprecoveryv1.ConnectorImageFile{}
	for _, connectorImageFileListItem := range modelMap["connector_image_file_list"].([]interface{}) {
		connectorImageFileListItemModel, err := ResourceIbmConnectorsMetadataMapToConnectorImageFile(connectorImageFileListItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		connectorImageFileList = append(connectorImageFileList, *connectorImageFileListItemModel)
	}
	model.ConnectorImageFileList = connectorImageFileList
	return model, nil
}

func ResourceIbmConnectorsMetadataMapToConnectorImageFile(modelMap map[string]interface{}) (*backuprecoveryv1.ConnectorImageFile, error) {
	model := &backuprecoveryv1.ConnectorImageFile{}
	model.ImageType = core.StringPtr(modelMap["image_type"].(string))
	model.URL = core.StringPtr(modelMap["url"].(string))
	return model, nil
}

func ResourceIbmConnectorsMetadataConnectorImageMetadataToMap(model *backuprecoveryv1.ConnectorImageMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	connectorImageFileList := []map[string]interface{}{}
	for _, connectorImageFileListItem := range model.ConnectorImageFileList {
		connectorImageFileListItemMap, err := ResourceIbmConnectorsMetadataConnectorImageFileToMap(&connectorImageFileListItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		connectorImageFileList = append(connectorImageFileList, connectorImageFileListItemMap)
	}
	modelMap["connector_image_file_list"] = connectorImageFileList
	return modelMap, nil
}

func ResourceIbmConnectorsMetadataConnectorImageFileToMap(model *backuprecoveryv1.ConnectorImageFile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["image_type"] = *model.ImageType
	modelMap["url"] = *model.URL
	return modelMap, nil
}
