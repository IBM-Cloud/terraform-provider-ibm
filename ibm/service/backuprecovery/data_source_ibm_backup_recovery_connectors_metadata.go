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
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryConnectorsMetadata() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorsMetadataRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
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
			"k8s_connector_info_list": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "k8sConnectorInfoList specifies information about supported kubernetes environments where Data-Source Connectors can be deployed. Also, specifies the helm chart location (OCI URL) for each supported Kubernetes environment and instructions for installing it.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"helm_chart_oci_ref": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Represents the structured components of an OCI (Open Container Initiative) artifact reference. A full reference string can be constructed from these parts. See Also: https://github.com/opencontainers/distribution-spec/blob/main/spec.md.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"digest": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The immutable, content-addressable digest of the artifact's manifest. If only digest is set, the artifact is fetched by its immutable reference. If both tag and digest are set, the application should verify that the tag resolves to the given digest before proceeding. This should include the algorithm prefix.",
									},
									"namespace": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The namespace or organization within the registry. For public registries like Docker Hub, this can be 'library' for official images or a user's account name. May be optional for certain registry configurations.",
									},
									"registry_host": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The address of the OCI-compliant container registry. This can be a hostname or an IP address, and may optionally include a port number.",
									},
									"repository": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the repository that holds the artifact.",
									},
									"tag": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The mutable tag for the artifact.",
									},
								},
							},
						},
						"helm_install_cmd": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the Helm install command for this type of k8s connector.",
						},
						"k8s_platform_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Enum representing the different supported Kubernetes platform types.",
						},
						"ugrade_doc_url": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL for upgrade documentation for this type of k8s connector.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryConnectorsMetadataRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connectors_metadata", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	if instanceId != "" && region != "" {
		bmxsession, err := meta.(conns.ClientSession).BluemixSession()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupRecoveryClient = getClientWithInstanceEndpoint(backupRecoveryClient, bmxsession, instanceId, region, endpointType)
	}

	getConnectorMetadataOptions := &backuprecoveryv1.GetConnectorMetadataOptions{}

	getConnectorMetadataOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	connectorMetadata, _, err := backupRecoveryClient.GetConnectorMetadataWithContext(context, getConnectorMetadataOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetConnectorMetadataWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connectors_metadata", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorsMetadataID(d))

	if !core.IsNil(connectorMetadata.ConnectorImageMetadata) {
		connectorImageMetadata := []map[string]interface{}{}
		connectorImageMetadataMap, err := DataSourceIbmBackupRecoveryConnectorsMetadataConnectorImageMetadataToMap(connectorMetadata.ConnectorImageMetadata)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connectors_metadata", "read", "connector_image_metadata-to-map").GetDiag()
		}
		connectorImageMetadata = append(connectorImageMetadata, connectorImageMetadataMap)
		if err = d.Set("connector_image_metadata", connectorImageMetadata); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting connector_image_metadata: %s", err), "(Data) ibm_backup_recovery_connectors_metadata", "read", "set-connector_image_metadata").GetDiag()
		}
	}
	if !core.IsNil(connectorMetadata.K8sConnectorInfoList) {
		k8sConnectorInfoList := []map[string]interface{}{}
		for _, k8sConnectorInfoListItem := range connectorMetadata.K8sConnectorInfoList {
			k8sConnectorInfoListItemMap, err := DataSourceIbmBackupRecoveryConnectorsMetadataK8sConnectorInfoToMap(&k8sConnectorInfoListItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connectors_metadata", "read", "k8s_connector_info_list-to-map").GetDiag()
			}
			k8sConnectorInfoList = append(k8sConnectorInfoList, k8sConnectorInfoListItemMap)
		}
		if err = d.Set("k8s_connector_info_list", k8sConnectorInfoList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting k8s_connector_info_list: %s", err), "(Data) ibm_backup_recovery_connectors_metadata", "read", "set-k8s_connector_info_list").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorsMetadataID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorsMetadataID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryConnectorsMetadataConnectorImageMetadataToMap(model *backuprecoveryv1.ConnectorImageMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	connectorImageFileList := []map[string]interface{}{}
	for _, connectorImageFileListItem := range model.ConnectorImageFileList {
		connectorImageFileListItemMap, err := DataSourceIbmBackupRecoveryConnectorsMetadataConnectorImageFileToMap(&connectorImageFileListItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		connectorImageFileList = append(connectorImageFileList, connectorImageFileListItemMap)
	}
	modelMap["connector_image_file_list"] = connectorImageFileList
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorsMetadataConnectorImageFileToMap(model *backuprecoveryv1.ConnectorImageFile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["image_type"] = *model.ImageType
	modelMap["url"] = *model.URL
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorsMetadataK8sConnectorInfoToMap(model *backuprecoveryv1.K8sConnectorInfo) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.HelmChartOciRef != nil {
		helmChartOciRefMap, err := DataSourceIbmBackupRecoveryConnectorsMetadataOciArtifactReferenceToMap(model.HelmChartOciRef)
		if err != nil {
			return modelMap, err
		}
		modelMap["helm_chart_oci_ref"] = []map[string]interface{}{helmChartOciRefMap}
	}
	if model.HelmInstallCmd != nil {
		modelMap["helm_install_cmd"] = *model.HelmInstallCmd
	}
	modelMap["k8s_platform_type"] = *model.K8sPlatformType
	if model.UgradeDocURL != nil {
		modelMap["ugrade_doc_url"] = *model.UgradeDocURL
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorsMetadataOciArtifactReferenceToMap(model *backuprecoveryv1.OciArtifactReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["digest"] = *model.Digest
	if model.Namespace != nil {
		modelMap["namespace"] = *model.Namespace
	}
	modelMap["registry_host"] = *model.RegistryHost
	modelMap["repository"] = *model.Repository
	modelMap["tag"] = *model.Tag
	return modelMap, nil
}
