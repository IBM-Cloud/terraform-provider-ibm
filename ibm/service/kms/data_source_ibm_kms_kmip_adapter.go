// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMKMSKmipAdapterBaseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"adapter_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The UUID of the KMIP adapter to be fetched",
		},
		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The name of the KMIP adapter to be fetched",
		},
		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The description of the KMIP adapter to be fetched",
		},
		"profile": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The profile of the KMIP adapter to be fetched",
		},
		"profile_data": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "The data specific to the KMIP Adapter profile",
		},
		"created_by": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier that is associated with the entity that created the adapter.",
		},
		"created_at": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date when a resource was created. The date format follows RFC 3339.",
		},
		"updated_by": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier that is associated with the entity that updated the adapter.",
		},
		"updated_at": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date when a resource was updated. The date format follows RFC 3339.",
		},
	}
}

func DataSourceIBMKMSKmipAdapter() *schema.Resource {
	baseMap := dataSourceIBMKMSKmipAdapterBaseSchema()
	baseMap["instance_id"] = *schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		Description:      "Key protect or hpcs instance GUID",
		DiffSuppressFunc: suppressKMSInstanceIDDiff,
	}
	adapterIDSchema := baseMap["adapter_id"]
	adapterIDSchema.Optional = true
	adapterIDSchema.ExactlyOneOF = []string{"adapter_id", "name"}

	adapterNameSchema := baseMap["name"]
	adapterNameSchema.Optional = true
	adapterNameSchema.ExactlyOneOf = []string{"adapter_id", "name"}

	return &schema.Resource{
		Read:   dataSourceIBMKMSKmipAdapterRead,
		Schema: baseMap,
	}
}

func dataSourceIBMKMSKmipAdapterRead(d *schema.ResourceData, meta interface{}) error {
	// initialize API
	api, err := meta.(conns.ClientSession).KeyProtectAPI()
	if err != nil {
		return err
	}

	instanceID := getInstanceIDFromResourceData(d, "instance_id")
	api.Config.InstanceID = instanceID
	nameOrID, hasID := d.GetOk("adapter_id").(string)
	if !hasID {
		nameOrID = d.Get("name")
	}

	// call GetKMIPAdapter api
	adapter, err := api.GetKMIPAdapter(context.Background(), nameOrID)
	if err != nil {
		return err
	}

	// set computed values
	populateKMIPAdapterSchemaDataFromStruct(d, adapter)
	return nil
}
