// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMKMSKmipAdapters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMKMSKmipAdaptersList,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Key protect or hpcs instance GUID",
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
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
			"adapters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A collection of KMIP adapters",
				Elem: &schema.Resource{
					Schema: dataSourceIBMKMSKmipAdapterBaseSchema(),
				},
			},
		},
	}
}

func dataSourceIBMKMSKmipAdaptersList(d *schema.ResourceData, meta interface{}) error {
	// initialize API
	api, err := meta.(conns.ClientSession).KeyProtectAPI()
	if err != nil {
		return err
	}

	instanceID := getInstanceIDFromResourceData(d, "instance_id")
	api.Config.InstanceID = instanceID

	// call GetKMIPAdapters api
	opts := &kp.ListKmipAdaptersOptions{}
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

	adapters, err := api.GetKMIPAdapters(context.Background(), opts)
	if err != nil {
		return fmt.Errorf("[ERROR] Error listing KMIP adapters: %s", err)
	}

	adaptersList := adapters.Adapters

	// set computed values
	mySlice := make([]map[string]interface{}, 0, len(adaptersList))
	for _, adapter := range adaptersList {
		adapterMap := dataSourceIBMKMSKmipAdapterToMap(adapter)
		mySlice = append(mySlice, adapterMap)
	}
	d.Set("adapters", mySlice)
	d.Set("instance_id", instanceID)
	d.SetId(instanceID)
	return nil
}

func dataSourceIBMKMSKmipAdapterToMap(model kp.KMIPAdapter) map[string]interface{} {
	modelMap := make(map[string]interface{})
	modelMap["adapter_id"] = model.ID
	modelMap["name"] = model.Name
	modelMap["description"] = model.Description
	modelMap["profile"] = model.Profile
	modelMap["profile_data"] = model.ProfileData
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["created_by"] = model.CreatedBy
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["updated_by"] = model.CreatedBy
	return modelMap
}
