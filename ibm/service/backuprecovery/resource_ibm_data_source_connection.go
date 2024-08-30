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

func ResourceIbmDataSourceConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmDataSourceConnectionCreate,
		ReadContext:   resourceIbmDataSourceConnectionRead,
		UpdateContext: resourceIbmDataSourceConnectionUpdate,
		DeleteContext: resourceIbmDataSourceConnectionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"tenant_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Id of the tenant accessing the cluster.",
			},
			"connection_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the name of the connection. For a given tenant, different connections can't have the same name. However, two (or more) different tenants can each have a connection with the same name.",
			},
			"connector_ids": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the IDs of the connectors in this connection.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"network_settings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the common network settings for the connectors associated with this connection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_fqdn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the FQDN for the cluster as visible to the connectors in this connection.",
						},
						"dns": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the DNS servers to be used by the connectors in this connection.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"network_gateway": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the network gateway to be used by the connectors in this connection.",
						},
						"ntp": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies the NTP server to be used by the connectors in this connection.",
						},
					},
				},
			},
			"registration_token": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies a token that can be used to register a connector against this connection.",
			},
		},
	}
}

func resourceIbmDataSourceConnectionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDataSourceConnectionOptions := &backuprecoveryv1.CreateDataSourceConnectionOptions{}

	createDataSourceConnectionOptions.SetConnectionName(d.Get("connection_name").(string))
	if _, ok := d.GetOk("tenant_id"); ok {
		createDataSourceConnectionOptions.SetTenantID(int64(d.Get("tenant_id").(int)))
	}

	dataSourceConnection, _, err := backupRecoveryClient.CreateDataSourceConnectionWithContext(context, createDataSourceConnectionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDataSourceConnectionWithContext failed: %s", err.Error()), "ibm_data_source_connection", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*dataSourceConnection.ConnectionID)
	d.Set("registration_token", dataSourceConnection.RegistrationToken)

	return resourceIbmDataSourceConnectionRead(context, d, meta)
}

func resourceIbmDataSourceConnectionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDataSourceConnectionsOptions := &backuprecoveryv1.GetDataSourceConnectionsOptions{}
	getDataSourceConnectionsOptions.ConnectionIds = []string{d.Id()}
	if _, ok := d.GetOk("tenant_id"); ok {
		getDataSourceConnectionsOptions.SetTenantID(d.Get("tenant_id").(string))
	}

	dataSourceConnectionList, response, err := backupRecoveryClient.GetDataSourceConnectionsWithContext(context, getDataSourceConnectionsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDataSourceConnectionsWithContext failed: %s", err.Error()), "ibm_data_source_connection", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("connection_name", dataSourceConnectionList.Connections[0].ConnectionName); err != nil {
		err = fmt.Errorf("Error setting connection_name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "read", "set-connection_name").GetDiag()
	}
	if !core.IsNil(dataSourceConnectionList.Connections[0].ConnectorIds) {
		if err = d.Set("connector_ids", dataSourceConnectionList.Connections[0].ConnectorIds); err != nil {
			err = fmt.Errorf("Error setting connector_ids: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "read", "set-connector_ids").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectionList.Connections[0].NetworkSettings) {
		networkSettingsMap, err := ResourceIbmDataSourceConnectionNetworkSettingsToMap(dataSourceConnectionList.Connections[0].NetworkSettings)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "read", "network_settings-to-map").GetDiag()
		}
		if err = d.Set("network_settings", []map[string]interface{}{networkSettingsMap}); err != nil {
			err = fmt.Errorf("Error setting network_settings: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "read", "set-network_settings").GetDiag()
		}
	}
	if !core.IsNil(dataSourceConnectionList.Connections[0].RegistrationToken) {
		if err = d.Set("registration_token", dataSourceConnectionList.Connections[0].RegistrationToken); err != nil {
			err = fmt.Errorf("Error setting registration_token: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "read", "set-registration_token").GetDiag()
		}
	}

	return nil
}

func resourceIbmDataSourceConnectionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	patchDataSourceConnectionOptions := &backuprecoveryv1.PatchDataSourceConnectionOptions{}

	patchDataSourceConnectionOptions.SetConnectionID(d.Id())

	hasChange := false

	if d.HasChange("connection_name") {
		patchDataSourceConnectionOptions.SetConnectionName(d.Get("connection_name").(string))
		hasChange = true
	}

	if hasChange {
		_, _, err = backupRecoveryClient.PatchDataSourceConnectionWithContext(context, patchDataSourceConnectionOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PatchDataSourceConnectionWithContext failed: %s", err.Error()), "ibm_data_source_connection", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmDataSourceConnectionRead(context, d, meta)
}

func resourceIbmDataSourceConnectionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_data_source_connection", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDataSourceConnectionOptions := &backuprecoveryv1.DeleteDataSourceConnectionOptions{}

	deleteDataSourceConnectionOptions.SetConnectionID(d.Id())

	_, err = backupRecoveryClient.DeleteDataSourceConnectionWithContext(context, deleteDataSourceConnectionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDataSourceConnectionWithContext failed: %s", err.Error()), "ibm_data_source_connection", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}

func ResourceIbmDataSourceConnectionNetworkSettingsToMap(model *backuprecoveryv1.NetworkSettings) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ClusterFqdn != nil {
		modelMap["cluster_fqdn"] = *model.ClusterFqdn
	}
	if model.Dns != nil {
		modelMap["dns"] = model.Dns
	}
	if model.NetworkGateway != nil {
		modelMap["network_gateway"] = *model.NetworkGateway
	}
	if model.Ntp != nil {
		modelMap["ntp"] = *model.Ntp
	}
	return modelMap, nil
}
