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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBaasDataSourceConnectorRegistration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBaasDataSourceConnectorRegistrationCreate,
		ReadContext:   resourceIbmBaasDataSourceConnectorRegistrationRead,
		DeleteContext: resourceIbmBaasDataSourceConnectorRegistrationDelete,
		UpdateContext: resourceIbmBaasDataSourceConnectorRegistrationUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"connector_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "Specifies the unique ID of the connector which is to be deleted.",
			},
			"x_ibm_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"cluster_side_ip": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the IP of the connector's NIC facing the cluster.",
			},
			"connection_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the ID of the connection to which this connector belongs.",
			},
			"connector_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the name of the connector. The name of a connector need not be unique within a tenant or across tenants. The name of the connector can be updated as needed.",
			},
			"connector_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies status information for the data-source connector. For example if it's currently connected to the cluster, when it last connected to the cluster successfully, etc.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_connected": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies whether the connector is currently connected to the cluster.",
						},
						"last_connected_timestamp_secs": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the last timestamp in UNIX time (seconds) when the connector had successfully connected to the cluster. This property can be present even if the connector is currently disconnected.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies error message when the connector is unable to connect to the cluster.",
						},
					},
				},
			},
			"software_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the connector's software version.",
			},
			"tenant_side_ip": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the IP of the connector's NIC facing the sources of the tenant to which the connector belongs.",
			},
		},
	}
}

func resourceIbmBaasDataSourceConnectorRegistrationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	registerDataSourceConnectorOptions := &backuprecoveryv1.RegisterDataSourceConnectorOptions{}

	registerDataSourceConnectorOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	registerDataSourceConnectorOptions.SetConnectorID(d.Get("connector_id").(string))

	registerDataSourceConnectorOptions.SetRegistrationToken(d.Get("registration_token").(string))
	if _, ok := d.GetOk("connector_id"); ok {
		registerDataSourceConnectorOptions.SetConnectorID(d.Get("connector_id").(string))
	}

	_, err = backupRecoveryClient.RegisterDataSourceConnectorWithContext(context, registerDataSourceConnectorOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RegisterDataSourceConnectorWithContext failed: %s", err.Error()), "ibm_baas_data_source_connector_registration", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(d.Get("connector_id").(string))

	return resourceIbmBaasDataSourceConnectorRegistrationRead(context, d, meta)
}

func resourceIbmBaasDataSourceConnectorRegistrationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDataSourceConnectorsOptions := &backuprecoveryv1.GetDataSourceConnectorsOptions{}

	getDataSourceConnectorsOptions.SetConnectorIds([]string{d.Id()})
	getDataSourceConnectorsOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))

	dataSourceConnectorList, response, err := backupRecoveryClient.GetDataSourceConnectorsWithContext(context, getDataSourceConnectorsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDataSourceConnectorsWithContext failed: %s", err.Error()), "ibm_baas_data_source_connector_registration", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(dataSourceConnectorList.Connectors[0].ClusterSideIp) {
		if err = d.Set("cluster_side_ip", dataSourceConnectorList.Connectors[0].ClusterSideIp); err != nil {
			err = fmt.Errorf("Error setting cluster_side_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "set-cluster_side_ip").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectorList.Connectors[0].ConnectionID) {
		if err = d.Set("connection_id", dataSourceConnectorList.Connectors[0].ConnectionID); err != nil {
			err = fmt.Errorf("Error setting connection_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "set-connection_id").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectorList.Connectors[0].ConnectorName) {
		if err = d.Set("connector_name", dataSourceConnectorList.Connectors[0].ConnectorName); err != nil {
			err = fmt.Errorf("Error setting connector_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "set-connector_name").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectorList.Connectors[0].ConnectorStatus) {
		connectorStatusMap, err := ResourceIbmBaasDataSourceConnectorRegistrationDataSourceConnectorStatusToMap(dataSourceConnectorList.Connectors[0].ConnectorStatus)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "connector_status-to-map").GetDiag()
		}
		if err = d.Set("connector_status", []map[string]interface{}{connectorStatusMap}); err != nil {
			err = fmt.Errorf("Error setting connector_status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "set-connector_status").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectorList.Connectors[0].SoftwareVersion) {
		if err = d.Set("software_version", dataSourceConnectorList.Connectors[0].SoftwareVersion); err != nil {
			err = fmt.Errorf("Error setting software_version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "set-software_version").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectorList.Connectors[0].TenantSideIp) {
		if err = d.Set("tenant_side_ip", dataSourceConnectorList.Connectors[0].TenantSideIp); err != nil {
			err = fmt.Errorf("Error setting tenant_side_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_data_source_connector_registration", "read", "set-tenant_side_ip").GetDiag()
		}
	}

	return nil
}

func resourceIbmBaasDataSourceConnectorRegistrationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "Delete operation is not supported for this resource. The resource will be removed from the terraform state file but will continue to exist in the backend.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmBaasDataSourceConnectorRegistrationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Resource Update Will Only Affect Terraform State",
		Detail:   "Update operation for this resource not supported and will only affect the terraform statefile. No changes will be made to actual backend resource. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func ResourceIbmBaasDataSourceConnectorRegistrationDataSourceConnectorStatusToMap(model *backuprecoveryv1.DataSourceConnectorStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_connected"] = *model.IsConnected
	if model.LastConnectedTimestampSecs != nil {
		modelMap["last_connected_timestamp_secs"] = flex.IntValue(model.LastConnectedTimestampSecs)
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}
