// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.4-9b56d441-20260612-210048
 */

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func DataSourceIbmSmInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmInstanceRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The service instance ID.",
			},
			"instance_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The instance CRN identifier.",
			},
			"plan": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Instance plan name.",
			},
			"vault_cluster": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Vault cluster information for Vault Dedicated instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vault cluster status. Possible values:- sealed: The Vault cluster is sealed and requires unsealing to access secrets- not_initialized: The Vault cluster has not been initialized yet- healthy: The Vault cluster is operational and ready to serve requests.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vault cluster version.",
						},
					},
				},
			},
			"endpoints": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Instance endpoints for Vault Dedicated instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"public": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Endpoint URLs for accessing the Vault Dedicated instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vault_api": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Vault API endpoint URL.",
									},
									"vault_ui": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Vault UI endpoint URL.",
									},
								},
							},
						},
						"private": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Endpoint URLs for accessing the Vault Dedicated instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vault_api": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Vault API endpoint URL.",
									},
									"vault_ui": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Vault UI endpoint URL.",
									},
								},
							},
						},
					},
				},
			},
			"encryption": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Vault encryption configuration for Vault Dedicated instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vault encryption mode.",
						},
						"provider": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vault encryption provider (only present for customer_managed mode). Valid value - 'key_protect'.",
						},
						"key_crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vault encryption key CRN (only present for customer_managed mode).",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmInstanceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := meta.(conns.ClientSession).SecretsManagerInstanceManagementV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_instance", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceOptions := &secretsmanagerinstancemanagementv2.GetInstanceOptions{}

	getInstanceOptions.SetInstanceID(d.Get("instance_id").(string))

	instance, _, err := secretsManagerInstanceManagementClient.GetInstanceWithContext(context, getInstanceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceWithContext failed: %s", err.Error()), "(Data) ibm_sm_instance", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmSmInstanceID(d))

	if err = d.Set("instance_crn", instance.InstanceCrn); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting instance_crn: %s", err), "(Data) ibm_sm_instance", "read", "set-instance_crn").GetDiag()
	}

	if err = d.Set("plan", instance.Plan); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting plan: %s", err), "(Data) ibm_sm_instance", "read", "set-plan").GetDiag()
	}

	vaultCluster := []map[string]interface{}{}
	vaultClusterMap, err := DataSourceIbmSmInstanceVaultDedicatedClusterToMap(instance.VaultCluster)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_instance", "read", "vault_cluster-to-map").GetDiag()
	}
	vaultCluster = append(vaultCluster, vaultClusterMap)
	if err = d.Set("vault_cluster", vaultCluster); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting vault_cluster: %s", err), "(Data) ibm_sm_instance", "read", "set-vault_cluster").GetDiag()
	}

	endpoints := []map[string]interface{}{}
	endpointsMap, err := DataSourceIbmSmInstanceVaultDedicatedInstanceEndpointsToMap(instance.Endpoints)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_instance", "read", "endpoints-to-map").GetDiag()
	}
	endpoints = append(endpoints, endpointsMap)
	if err = d.Set("endpoints", endpoints); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting endpoints: %s", err), "(Data) ibm_sm_instance", "read", "set-endpoints").GetDiag()
	}

	encryption := []map[string]interface{}{}
	encryptionMap, err := DataSourceIbmSmInstanceVaultDedicatedInstanceEncryptionToMap(instance.Encryption)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_sm_instance", "read", "encryption-to-map").GetDiag()
	}
	encryption = append(encryption, encryptionMap)
	if err = d.Set("encryption", encryption); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting encryption: %s", err), "(Data) ibm_sm_instance", "read", "set-encryption").GetDiag()
	}

	return nil
}

// dataSourceIbmSmInstanceID returns a reasonable ID for the list.
func dataSourceIbmSmInstanceID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmSmInstanceVaultDedicatedClusterToMap(model *secretsmanagerinstancemanagementv2.VaultDedicatedCluster) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = *model.Status
	modelMap["version"] = *model.Version
	return modelMap, nil
}

func DataSourceIbmSmInstanceVaultDedicatedInstanceEndpointsToMap(model *secretsmanagerinstancemanagementv2.VaultDedicatedInstanceEndpoints) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Public != nil {
		publicMap, err := DataSourceIbmSmInstanceVaultDedicatedEndpointsDataToMap(model.Public)
		if err != nil {
			return modelMap, err
		}
		modelMap["public"] = []map[string]interface{}{publicMap}
	}
	privateMap, err := DataSourceIbmSmInstanceVaultDedicatedEndpointsDataToMap(model.Private)
	if err != nil {
		return modelMap, err
	}
	modelMap["private"] = []map[string]interface{}{privateMap}
	return modelMap, nil
}

func DataSourceIbmSmInstanceVaultDedicatedEndpointsDataToMap(model *secretsmanagerinstancemanagementv2.VaultDedicatedEndpointsData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["vault_api"] = *model.VaultApi
	modelMap["vault_ui"] = *model.VaultUi
	return modelMap, nil
}

func DataSourceIbmSmInstanceVaultDedicatedInstanceEncryptionToMap(model *secretsmanagerinstancemanagementv2.VaultDedicatedInstanceEncryption) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mode"] = *model.Mode
	if model.Provider != nil {
		modelMap["provider"] = *model.Provider
	}
	if model.KeyCrn != nil {
		modelMap["key_crn"] = *model.KeyCrn
	}
	return modelMap, nil
}
