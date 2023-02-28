// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func DataSourceIbmSmConfigurations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmConfigurationsRead,

		Schema: map[string]*schema.Schema{
			"sort": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sort a collection of secrets by the specified field in ascending order. To sort in descending order use the `-` character. Available values: id | created_at | updated_at | expiration_date | secret_type | name",
			},
			"search": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Obtain a collection of secrets that contain the specified string in one or more of the fields: `id`, `name`, `description`,\n        `labels`, `secret_type`.",
			},
			"groups": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter secrets by groups. You can apply multiple filters by using a comma-separated list of secret group IDs. If you need to filter secrets that are in the default secret group, use the `default` keyword.",
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources in a collection.",
			},
			"configurations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A collection of configuration metadata.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration type.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique name of your configuration.",
						},
						"secret_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier that is associated with the entity that created the secret.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when a resource was created. The date format follows RFC 3339.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date when a resource was recently modified. The date format follows RFC 3339.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmConfigurationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	region := getRegion(secretsManagerClient, d)
	instanceId := d.Get("instance_id").(string)
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d))

	listConfigurationsOptions := &secretsmanagerv2.ListConfigurationsOptions{}
	sort, ok := d.GetOk("sort")
	if ok {
		sortStr := sort.(string)
		listConfigurationsOptions.SetSort(sortStr)
	}
	search, ok := d.GetOk("search")
	if ok {
		searchStr := search.(string)
		listConfigurationsOptions.SetSearch(searchStr)
	}

	var pager *secretsmanagerv2.ConfigurationsPager
	pager, err = secretsManagerClient.NewConfigurationsPager(listConfigurationsOptions)
	if err != nil {
		return diag.FromErr(err)
	}

	allItems, err := pager.GetAll()
	if err != nil {
		log.Printf("[DEBUG] ConfigurationsPager.GetAll() failed %s", err)
		return diag.FromErr(fmt.Errorf("ConfigurationsPager.GetAll() failed %s", err))
	}

	d.SetId(dataSourceIbmSmConfigurationsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := dataSourceIbmSmConfigurationsConfigurationMetadataToMap(modelItem)
		if err != nil {
			return diag.FromErr(err)
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("configurations", mapSlice); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting configurations %s", err))
	}
	if err = d.Set("total_count", len(mapSlice)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting locks_total: %s", err))
	}

	return nil
}

// dataSourceIbmSmConfigurationsID returns a reasonable ID for the list.
func dataSourceIbmSmConfigurationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmSmConfigurationsConfigurationMetadataToMap(model secretsmanagerv2.ConfigurationMetadataIntf) (map[string]interface{}, error) {
	if _, ok := model.(*secretsmanagerv2.IAMCredentialsConfigurationMetadata); ok {
		return dataSourceIbmSmConfigurationsIAMCredentialsConfigurationMetadataToMap(model.(*secretsmanagerv2.IAMCredentialsConfigurationMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PublicCertificateConfigurationCALetsEncryptMetadata); ok {
		return dataSourceIbmSmConfigurationsPublicCertificateConfigurationCALetsEncryptMetadataToMap(model.(*secretsmanagerv2.PublicCertificateConfigurationCALetsEncryptMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PublicCertificateConfigurationDNSCloudInternetServicesMetadata); ok {
		return dataSourceIbmSmConfigurationsPublicCertificateConfigurationDNSCloudInternetServicesMetadataToMap(model.(*secretsmanagerv2.PublicCertificateConfigurationDNSCloudInternetServicesMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PublicCertificateConfigurationDNSClassicInfrastructureMetadata); ok {
		return dataSourceIbmSmConfigurationsPublicCertificateConfigurationDNSClassicInfrastructureMetadataToMap(model.(*secretsmanagerv2.PublicCertificateConfigurationDNSClassicInfrastructureMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PrivateCertificateConfigurationRootCAMetadata); ok {
		return dataSourceIbmSmConfigurationsPrivateCertificateConfigurationRootCAMetadataToMap(model.(*secretsmanagerv2.PrivateCertificateConfigurationRootCAMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PrivateCertificateConfigurationIntermediateCAMetadata); ok {
		return dataSourceIbmSmConfigurationsPrivateCertificateConfigurationIntermediateCAMetadataToMap(model.(*secretsmanagerv2.PrivateCertificateConfigurationIntermediateCAMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PrivateCertificateConfigurationTemplateMetadata); ok {
		return dataSourceIbmSmConfigurationsPrivateCertificateConfigurationTemplateMetadataToMap(model.(*secretsmanagerv2.PrivateCertificateConfigurationTemplateMetadata))
	} else if _, ok := model.(*secretsmanagerv2.ConfigurationMetadata); ok {
		modelMap := make(map[string]interface{})
		model := model.(*secretsmanagerv2.ConfigurationMetadata)
		if model.ConfigType != nil {
			modelMap["config_type"] = *model.ConfigType
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.SecretType != nil {
			modelMap["secret_type"] = *model.SecretType
		}
		if model.CreatedBy != nil {
			modelMap["created_by"] = *model.CreatedBy
		}
		if model.CreatedAt != nil {
			modelMap["created_at"] = model.CreatedAt.String()
		}
		if model.UpdatedAt != nil {
			modelMap["updated_at"] = model.UpdatedAt.String()
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized secretsmanagerv2.ConfigurationMetadataIntf subtype encountered")
	}
}

func dataSourceIbmSmConfigurationsPrivateCertificateConfigurationIntermediateCAMetadataToMap(model *secretsmanagerv2.PrivateCertificateConfigurationIntermediateCAMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmConfigurationsIAMCredentialsConfigurationMetadataToMap(model *secretsmanagerv2.IAMCredentialsConfigurationMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmConfigurationsPrivateCertificateConfigurationRootCAMetadataToMap(model *secretsmanagerv2.PrivateCertificateConfigurationRootCAMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmConfigurationsPublicCertificateConfigurationDNSClassicInfrastructureMetadataToMap(model *secretsmanagerv2.PublicCertificateConfigurationDNSClassicInfrastructureMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmConfigurationsPublicCertificateConfigurationCALetsEncryptMetadataToMap(model *secretsmanagerv2.PublicCertificateConfigurationCALetsEncryptMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmConfigurationsPublicCertificateConfigurationDNSCloudInternetServicesMetadataToMap(model *secretsmanagerv2.PublicCertificateConfigurationDNSCloudInternetServicesMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmConfigurationsPrivateCertificateConfigurationTemplateMetadataToMap(model *secretsmanagerv2.PrivateCertificateConfigurationTemplateMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConfigType != nil {
		modelMap["config_type"] = *model.ConfigType
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.SecretType != nil {
		modelMap["secret_type"] = *model.SecretType
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	return modelMap, nil
}
