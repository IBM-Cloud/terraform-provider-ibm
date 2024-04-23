// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMKmsKMIPObjects() *schema.Resource {
	return &schema.Resource{
		Read:     resourceIBMKmsKMIPClientCertRead,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Key protect Instance GUID",
				ForceNew:         true,
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"adapter_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "The id of the KMIP adapter that contains the cert",
				ForceNew:     true,
				ExactlyOneOf: []string{"adapter_id", "adapter_name"},
			},
			"adapter_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "The name of the KMIP adapter that contains the cert",
				ForceNew:     true,
				ExactlyOneOf: []string{"adapter_id", "adapter_name"},
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Limit of how many adapters to be fetched",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Offset of adapters to be fetched",
			},
			"show_total_count": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Flag to return the count of how many adapters there are in total",
			},
			"object_state_filter": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The objects contained in the specified adapter",
				Elem: &schema.Resource{
					Schema: dataSourceIBMKMSKMIPObjectBaseSchema(true),
				},
			},
		},
	}
}

func dataSourceIBMKmsKMIPObjectList(d *schema.ResourceData, meta interface{}) error {
	// initialize API
	instanceID := getInstanceIDFromResourceData(d, "instance_id")
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	// get adapterID
	nameOrID, hasID := d.GetOk("adapter_id")
	if !hasID {
		nameOrID = d.Get("adapter_name")
	}
	adapterNameOrID := nameOrID.(string)

	opts := &kp.ListKmipObjectsOptions{}
	if limit, ok := d.GetOk("limit"); ok {
		limitVal := uint32(limit.(int))
		opts.Limit = &limitVal
	}
	if offset, ok := d.GetOk("offset"); ok {
		offsetVal := uint32(offset.(int))
		opts.Offset = &offsetVal
	}
	if showTotalCount, ok := d.GetOk("show_total_count"); ok {
		boolVal := showTotalCount.(bool)
		opts.TotalCount = &boolVal
	}
	if stateFilter, ok := d.GetOk("object_state_filter"); ok {
		arrayVal, ok2 := stateFilter.([]int32)
		// TODO: might have to convert each int into int32 manually
		if !ok2 {
			return fmt.Errorf("Error converting object_state_filter into []int32")
		}
		opts.ObjectStateFilter = &arrayVal
	}

	ctx := context.Background()
	adapter, err := kpAPI.GetKMIPAdapter(ctx, adapterNameOrID)
	if err != nil {
		return fmt.Errorf("[ERROR] Error while retriving KMIP adapter to list certificates: %s", err)
	}
	if err = d.Set("adapter_id", adapter.ID); err != nil {
		return fmt.Errorf("Error setting adapter_id: %s", err)
	}
	if err = d.Set("adapter_name", adapter.Name); err != nil {
		return fmt.Errorf("Error setting adapter_name: %s", err)
	}
	objs, err := kpAPI.GetKMIPObjects(ctx, adapterNameOrID, opts)
	objsList := objs.Objects
	// set computed values
	mySlice := make([]map[string]interface{}, 0, len(objsList))
	for _, obj := range objsList {
		objMap := dataSourceIBMKMSKmipObjectToMap(obj)
		mySlice = append(mySlice, objMap)
	}
	d.Set("objects", mySlice)
	return nil
}

func dataSourceIBMKMSKmipObjectToMap(model kp.KMIPObject) map[string]interface{} {
	modelMap := make(map[string]interface{})
	modelMap["object_id"] = model.ID
	modelMap["object_type"] = model.KMIPObjectType
	modelMap["object_state"] = model.ObjectState
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["created_by"] = model.CreatedBy
	modelMap["created_by_cert_id"] = model.CreatedByCertID
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["updated_by"] = model.UpdatedBy
	modelMap["updated_by_cert_id"] = model.UpdatedByCertID
	modelMap["destroyed_at"] = model.DestroyedAt.String()
	modelMap["destroyed_by"] = model.DestroyedBy
	modelMap["destroyed_by_cert_id"] = model.DestroyedByCertID
	return modelMap
}
