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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIbmCodeEnginePersistentDataStore() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCodeEnginePersistentDataStoreCreate,
		ReadContext:   resourceIbmCodeEnginePersistentDataStoreRead,
		DeleteContext: resourceIbmCodeEnginePersistentDataStoreDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_persistent_data_store", "project_id"),
				Description:  "The ID of the project.",
			},
			"data": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bucket_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify the location of the bucket.",
						},
						"bucket_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify the name of the bucket.",
						},
						"secret_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specify the name of the HMAC secret.",
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_persistent_data_store", "name"),
				Description:  "The name of the persistent data store.",
			},
			"storage_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_code_engine_persistent_data_store", "storage_type"),
				Description:  "Specify the storage type of the persistent data store.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the resource was created.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the persistent data store, which is used to achieve optimistic locking.",
			},
			"persistent_data_store_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier of the resource.",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.",
			},
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmCodeEnginePersistentDataStoreValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-z0-9]([\-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([\-a-z0-9]*[a-z0-9])?)*$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "storage_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "object_storage",
			Regexp:                     `^(object_storage)$`,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_code_engine_persistent_data_store", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmCodeEnginePersistentDataStoreCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createPersistentDataStoreOptions := &codeenginev2.CreatePersistentDataStoreOptions{}

	createPersistentDataStoreOptions.SetProjectID(d.Get("project_id").(string))
	createPersistentDataStoreOptions.SetName(d.Get("name").(string))
	createPersistentDataStoreOptions.SetStorageType(d.Get("storage_type").(string))
	if _, ok := d.GetOk("data"); ok {
		dataModel, err := ResourceIbmCodeEnginePersistentDataStoreMapToStorageData(d.Get("data.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "create", "parse-data").GetDiag()
		}
		createPersistentDataStoreOptions.SetData(dataModel)
	}

	persistentDataStore, _, err := codeEngineClient.CreatePersistentDataStoreWithContext(context, createPersistentDataStoreOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePersistentDataStoreWithContext failed: %s", err.Error()), "ibm_code_engine_persistent_data_store", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createPersistentDataStoreOptions.ProjectID, *persistentDataStore.Name))

	return resourceIbmCodeEnginePersistentDataStoreRead(context, d, meta)
}

func resourceIbmCodeEnginePersistentDataStoreRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPersistentDataStoreOptions := &codeenginev2.GetPersistentDataStoreOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "sep-id-parts").GetDiag()
	}

	getPersistentDataStoreOptions.SetProjectID(parts[0])
	getPersistentDataStoreOptions.SetName(parts[1])

	persistentDataStore, response, err := codeEngineClient.GetPersistentDataStoreWithContext(context, getPersistentDataStoreOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPersistentDataStoreWithContext failed: %s", err.Error()), "ibm_code_engine_persistent_data_store", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("project_id", persistentDataStore.ProjectID); err != nil {
		err = fmt.Errorf("Error setting project_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-project_id").GetDiag()
	}
	if !core.IsNil(persistentDataStore.Data) {
		dataMap, err := ResourceIbmCodeEnginePersistentDataStoreStorageDataToMap(persistentDataStore.Data)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "data-to-map").GetDiag()
		}
		if err = d.Set("data", []map[string]interface{}{dataMap}); err != nil {
			err = fmt.Errorf("Error setting data: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-data").GetDiag()
		}
	}
	if err = d.Set("name", persistentDataStore.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-name").GetDiag()
	}
	if err = d.Set("storage_type", persistentDataStore.StorageType); err != nil {
		err = fmt.Errorf("Error setting storage_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-storage_type").GetDiag()
	}
	if !core.IsNil(persistentDataStore.CreatedAt) {
		if err = d.Set("created_at", persistentDataStore.CreatedAt); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-created_at").GetDiag()
		}
	}
	if err = d.Set("entity_tag", persistentDataStore.EntityTag); err != nil {
		err = fmt.Errorf("Error setting entity_tag: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-entity_tag").GetDiag()
	}
	if !core.IsNil(persistentDataStore.ID) {
		if err = d.Set("persistent_data_store_id", persistentDataStore.ID); err != nil {
			err = fmt.Errorf("Error setting persistent_data_store_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-id").GetDiag()
		}
	}
	if !core.IsNil(persistentDataStore.Region) {
		if err = d.Set("region", persistentDataStore.Region); err != nil {
			err = fmt.Errorf("Error setting region: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "read", "set-region").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_code_engine_persistent_data_store", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIbmCodeEnginePersistentDataStoreDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePersistentDataStoreOptions := &codeenginev2.DeletePersistentDataStoreOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_code_engine_persistent_data_store", "delete", "sep-id-parts").GetDiag()
	}

	deletePersistentDataStoreOptions.SetProjectID(parts[0])
	deletePersistentDataStoreOptions.SetName(parts[1])

	_, err = codeEngineClient.DeletePersistentDataStoreWithContext(context, deletePersistentDataStoreOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePersistentDataStoreWithContext failed: %s", err.Error()), "ibm_code_engine_persistent_data_store", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmCodeEnginePersistentDataStoreMapToStorageData(modelMap map[string]interface{}) (codeenginev2.StorageDataIntf, error) {
	model := &codeenginev2.StorageData{}
	if modelMap["bucket_location"] != nil && modelMap["bucket_location"].(string) != "" {
		model.BucketLocation = core.StringPtr(modelMap["bucket_location"].(string))
	}
	if modelMap["bucket_name"] != nil && modelMap["bucket_name"].(string) != "" {
		model.BucketName = core.StringPtr(modelMap["bucket_name"].(string))
	}
	if modelMap["secret_name"] != nil && modelMap["secret_name"].(string) != "" {
		model.SecretName = core.StringPtr(modelMap["secret_name"].(string))
	}
	return model, nil
}

func ResourceIbmCodeEnginePersistentDataStoreMapToStorageDataObjectStorageData(modelMap map[string]interface{}) (*codeenginev2.StorageDataObjectStorageData, error) {
	model := &codeenginev2.StorageDataObjectStorageData{}
	model.BucketLocation = core.StringPtr(modelMap["bucket_location"].(string))
	model.BucketName = core.StringPtr(modelMap["bucket_name"].(string))
	model.SecretName = core.StringPtr(modelMap["secret_name"].(string))
	return model, nil
}

func ResourceIbmCodeEnginePersistentDataStoreStorageDataToMap(model codeenginev2.StorageDataIntf) (map[string]interface{}, error) {
	if _, ok := model.(*codeenginev2.StorageDataObjectStorageData); ok {
		return ResourceIbmCodeEnginePersistentDataStoreStorageDataObjectStorageDataToMap(model.(*codeenginev2.StorageDataObjectStorageData))
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

func ResourceIbmCodeEnginePersistentDataStoreStorageDataObjectStorageDataToMap(model *codeenginev2.StorageDataObjectStorageData) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["bucket_location"] = *model.BucketLocation
	modelMap["bucket_name"] = *model.BucketName
	modelMap["secret_name"] = *model.SecretName
	return modelMap, nil
}
