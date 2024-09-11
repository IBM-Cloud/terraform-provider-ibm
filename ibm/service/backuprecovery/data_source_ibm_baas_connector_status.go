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

func DataSourceIbmBaasConnectorStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBaasConnectorStatusRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
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

func dataSourceIbmBaasConnectorStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_baas_connector_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if backupRecoveryClient.ConnectorUrl == "" {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", "Connector URL is not set")
		return tfErr.GetDiag()
	}

	backupRecoveryClient.SetServiceURL(backupRecoveryClient.ConnectorUrl)

	getDataSourceConnectorStatusOptions := &backuprecoveryv1.GetDataSourceConnectorStatusOptions{}

	getDataSourceConnectorStatusOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	dataSourceConnectorLocalStatus, _, err := backupRecoveryClient.GetDataSourceConnectorStatusWithContext(context, getDataSourceConnectorStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDataSourceConnectorStatusWithContext failed: %s", err.Error()), "(Data) ibm_baas_connector_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBaasConnectorStatusID(d))

	if !core.IsNil(dataSourceConnectorLocalStatus.ClusterConnectionStatus) {
		clusterConnectionStatus := []map[string]interface{}{}
		clusterConnectionStatusMap, err := DataSourceIbmBaasConnectorStatusDataSourceConnectorClusterConnectionStatusToMap(dataSourceConnectorLocalStatus.ClusterConnectionStatus)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_baas_connector_status", "read", "cluster_connection_status-to-map").GetDiag()
		}
		clusterConnectionStatus = append(clusterConnectionStatus, clusterConnectionStatusMap)
		if err = d.Set("cluster_connection_status", clusterConnectionStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cluster_connection_status: %s", err), "(Data) ibm_baas_connector_status", "read", "set-cluster_connection_status").GetDiag()
		}
	}

	if !core.IsNil(dataSourceConnectorLocalStatus.IsCertificateValid) {
		if err = d.Set("is_certificate_valid", dataSourceConnectorLocalStatus.IsCertificateValid); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting is_certificate_valid: %s", err), "(Data) ibm_baas_connector_status", "read", "set-is_certificate_valid").GetDiag()
		}
	}

	if !core.IsNil(dataSourceConnectorLocalStatus.RegistrationStatus) {
		registrationStatus := []map[string]interface{}{}
		registrationStatusMap, err := DataSourceIbmBaasConnectorStatusDataSourceConnectorRegistrationStatusToMap(dataSourceConnectorLocalStatus.RegistrationStatus)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_baas_connector_status", "read", "registration_status-to-map").GetDiag()
		}
		registrationStatus = append(registrationStatus, registrationStatusMap)
		if err = d.Set("registration_status", registrationStatus); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting registration_status: %s", err), "(Data) ibm_baas_connector_status", "read", "set-registration_status").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBaasConnectorStatusID returns a reasonable ID for the list.
func dataSourceIbmBaasConnectorStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBaasConnectorStatusDataSourceConnectorClusterConnectionStatusToMap(model *backuprecoveryv1.DataSourceConnectorClusterConnectionStatus) (map[string]interface{}, error) {
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

func DataSourceIbmBaasConnectorStatusDataSourceConnectorRegistrationStatusToMap(model *backuprecoveryv1.DataSourceConnectorRegistrationStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}
