// Copyright IBM Corp. 2025 All Rights Reserved.
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

func DataSourceIbmBackupRecoveryConnectorStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryConnectorStatusRead,

		Schema: map[string]*schema.Schema{
			"access_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Token required to authenticate to the connector. Token can be obtained using ibm_backup_recovery_connector_access_token resource",
			},
			"cluster_connection_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the data-source connector-cluster connectivity status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_active": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the connection to the cluster is active.",
						},
						"last_connected_timestamp_msecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies last known connectivity status time in milliseconds.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies possible connectivity error message.",
						},
					},
				},
			},
			"is_certificate_valid": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag to indicate if connector certificate is valid.",
			},
			"registration_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the data-source connector registration status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the message corresponding the registration.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the registration status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryConnectorStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryConnectorClient, err := meta.(conns.ClientSession).BackupRecoveryV1Connector()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if backupRecoveryConnectorClient.GetConnectorURL() == "" {
		tfErr := flex.DiscriminatedTerraformErrorf(nil, "No connector URL specified. Please set the `IBMCLOUD_BACKUP_RECOVERY_CONNECTOR_ENDPOINT` environment variable or specify the endpoint in endpoints.json file.", "ibm_backup_recovery_connector_status", "read", "initialize-client")
		return tfErr.GetDiag()
	}

	accessToken := d.Get("access_token").(string)
	var auth core.Authenticator
	auth = &core.BearerTokenAuthenticator{BearerToken: accessToken}
	backupRecoveryConnectorClient.Service.Options.Authenticator = auth

	getDataSourceConnectorStatusOptions := &backuprecoveryv1.GetDataSourceConnectorStatusOptions{}

	dataSourceConnectorLocalStatus, _, err := backupRecoveryConnectorClient.GetDataSourceConnectorStatusWithContext(context, getDataSourceConnectorStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDataSourceConnectorStatusWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_connector_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryConnectorStatusID(d))

	if !core.IsNil(dataSourceConnectorLocalStatus.ClusterConnectionStatus) {
		clusterConnectionStatus := []map[string]interface{}{}
		clusterConnectionStatusMap, err := DataSourceIbmBackupRecoveryConnectorStatusDataSourceConnectorClusterConnectionStatusToMap(dataSourceConnectorLocalStatus.ClusterConnectionStatus)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_status", "read", "cluster_connection_status-to-map").GetDiag()
		}
		clusterConnectionStatus = append(clusterConnectionStatus, clusterConnectionStatusMap)
		if err = d.Set("cluster_connection_status", clusterConnectionStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_connection_status: %s", err), "(Data) ibm_backup_recovery_connector_status", "read", "set-cluster_connection_status").GetDiag()
		}
	}

	if !core.IsNil(dataSourceConnectorLocalStatus.IsCertificateValid) {
		if err = d.Set("is_certificate_valid", dataSourceConnectorLocalStatus.IsCertificateValid); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_certificate_valid: %s", err), "(Data) ibm_backup_recovery_connector_status", "read", "set-is_certificate_valid").GetDiag()
		}
	}

	if !core.IsNil(dataSourceConnectorLocalStatus.RegistrationStatus) {
		registrationStatus := []map[string]interface{}{}
		registrationStatusMap, err := DataSourceIbmBackupRecoveryConnectorStatusDataSourceConnectorRegistrationStatusToMap(dataSourceConnectorLocalStatus.RegistrationStatus)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_connector_status", "read", "registration_status-to-map").GetDiag()
		}
		registrationStatus = append(registrationStatus, registrationStatusMap)
		if err = d.Set("registration_status", registrationStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting registration_status: %s", err), "(Data) ibm_backup_recovery_connector_status", "read", "set-registration_status").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryConnectorStatusID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryConnectorStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryConnectorStatusDataSourceConnectorClusterConnectionStatusToMap(model *backuprecoveryv1.DataSourceConnectorClusterConnectionStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.IsActive != nil {
		modelMap["is_active"] = *model.IsActive
	}
	if model.LastConnectedTimestampMsecs != nil {
		modelMap["last_connected_timestamp_msecs"] = flex.IntValue(model.LastConnectedTimestampMsecs)
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryConnectorStatusDataSourceConnectorRegistrationStatusToMap(model *backuprecoveryv1.DataSourceConnectorRegistrationStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}
