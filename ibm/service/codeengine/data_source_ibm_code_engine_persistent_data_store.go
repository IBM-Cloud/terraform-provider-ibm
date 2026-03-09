// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package codeengine

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIbmCodeEnginePersistentDataStore() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmCodeEnginePersistentDataStoreRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the project.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of your persistent data store.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the resource was created.",
			},
			"data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket_location": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the location of the bucket.",
						},
						"bucket_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the name of the bucket.",
						},
						"secret_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the name of the HMAC secret.",
						},
					},
				},
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the persistent data store, which is used to achieve optimistic locking.",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.",
			},
			"storage_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specify the storage type of the persistent data store.",
			},
		},
	}
}

func dataSourceIbmCodeEnginePersistentDataStoreRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_code_engine_persistent_data_store", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPersistentDataStoreOptions := &codeenginev2.GetPersistentDataStoreOptions{}

	getPersistentDataStoreOptions.SetProjectID(d.Get("project_id").(string))
	getPersistentDataStoreOptions.SetName(d.Get("name").(string))

	persistentDataStore, _, err := codeEngineClient.GetPersistentDataStoreWithContext(context, getPersistentDataStoreOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPersistentDataStoreWithContext failed: %s", err.Error()), "(Data) ibm_code_engine_persistent_data_store", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getPersistentDataStoreOptions.ProjectID, *getPersistentDataStoreOptions.Name))

	if !core.IsNil(persistentDataStore.CreatedAt) {
		if err = d.Set("created_at", persistentDataStore.CreatedAt); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_code_engine_persistent_data_store", "read", "set-created_at").GetDiag()
		}
	}

	data := []map[string]interface{}{}
	dataMap, err := DataSourceIbmCodeEnginePersistentDataStoreStorageDataToMap(persistentDataStore.Data)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_code_engine_persistent_data_store", "read", "data-to-map").GetDiag()
	}
	data = append(data, dataMap)
	if err = d.Set("data", data); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting data: %s", err), "(Data) ibm_code_engine_persistent_data_store", "read", "set-data").GetDiag()
	}

	if err = d.Set("entity_tag", persistentDataStore.EntityTag); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_tag: %s", err), "(Data) ibm_code_engine_persistent_data_store", "read", "set-entity_tag").GetDiag()
	}

	if !core.IsNil(persistentDataStore.Region) {
		if err = d.Set("region", persistentDataStore.Region); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting region: %s", err), "(Data) ibm_code_engine_persistent_data_store", "read", "set-region").GetDiag()
		}
	}

	if err = d.Set("storage_type", persistentDataStore.StorageType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_type: %s", err), "(Data) ibm_code_engine_persistent_data_store", "read", "set-storage_type").GetDiag()
	}

	return nil
}

func DataSourceIbmCodeEnginePersistentDataStoreStorageDataToMap(model codeenginev2.StorageDataIntf) (map[string]interface{}, error) {
	if _, ok := model.(*codeenginev2.StorageDataObjectStorageData); ok {
		return DataSourceIbmCodeEnginePersistentDataStoreStorageDataObjectStorageDataToMap(model.(*codeenginev2.StorageDataObjectStorageData))
	} else if _, ok := model.(*codeenginev2.StorageData); ok {
		modelMap := make(map[string]interface{})
		model := model.(*codeenginev2.StorageData)
		if model.BucketLocation != nil {
			modelMap["bucket_location"] = *model.BucketLocation
		}
		if model.BucketName != nil {
			modelMap["bucket_name"] = *model.BucketName
		}
		if model.SecretName != nil {
			modelMap["secret_name"] = *model.SecretName
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized codeenginev2.StorageDataIntf subtype encountered")
	}
}

func DataSourceIbmCodeEnginePersistentDataStoreStorageDataObjectStorageDataToMap(model *codeenginev2.StorageDataObjectStorageData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["bucket_location"] = *model.BucketLocation
	modelMap["bucket_name"] = *model.BucketName
	modelMap["secret_name"] = *model.SecretName
	return modelMap, nil
}
