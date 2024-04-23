// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMKMSKMIPObjectBaseSchema(isForList bool) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"object_id": {
			Type:        schema.TypeString,
			Required:    !isForList,
			Computed:    isForList,
			Description: "The id of the KMIP Object to be fetched",
		},
		"object_type": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The type of KMIP object",
		},
		"object_state": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The state of the KMIP object",
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
		"created_by_cert_id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the certificate that created the object",
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
		"updated_by_cert_id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the certificate that updated the object",
		},
		"destroyed_by": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The unique identifier that is associated with the entity that destroyed the adapter.",
		},
		"destroyed_at": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date when a resource was destroyed. The date format follows RFC 3339.",
		},
		"destroyed_by_cert_id": &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The ID of the certificate that destroyed the object",
		},
	}
}

func DataSourceIBMKmsKMIPObject() *schema.Resource {
	baseMap := dataSourceIBMKMSKMIPObjectBaseSchema(false)

	baseMap["instance_id"] = &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		Description:      "Key protect Instance GUID",
		ForceNew:         true,
		DiffSuppressFunc: suppressKMSInstanceIDDiff,
	}
	baseMap["adapter_id"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		Description:  "The id of the KMIP adapter that contains the cert",
		ForceNew:     true,
		ExactlyOneOf: []string{"adapter_id", "adapter_name"},
	}
	baseMap["adapter_name"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		Description:  "The name of the KMIP adapter that contains the cert",
		ForceNew:     true,
		ExactlyOneOf: []string{"adapter_id", "adapter_name"},
	}

	return &schema.Resource{
		Read:     resourceIBMKmsKMIPClientCertRead,
		Importer: &schema.ResourceImporter{},
		Schema:   baseMap,
	}
}

func dataSourceIBMKmsKMIPObjectRead(d *schema.ResourceData, meta interface{}) error {
	// initialize API
	instanceID := getInstanceIDFromResourceData(d, "instance_id")
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	// get adapterID and certID
	nameOrID, hasID := d.GetOk("adapter_id")
	if !hasID {
		nameOrID = d.Get("adapter_name")
	}
	adapterNameOrID := nameOrID.(string)

	objectID := d.Get("object_id").(string)

	ctx := context.Background()
	adapter, err := kpAPI.GetKMIPAdapter(ctx, adapterNameOrID)
	if err != nil {
		return fmt.Errorf("[ERROR] Error while retriving KMIP adapter to get KMIP object: %s", err)
	}
	if err = d.Set("adapter_id", adapter.ID); err != nil {
		return fmt.Errorf("Error setting adapter_id: %s", err)
	}
	if err = d.Set("adapter_name", adapter.Name); err != nil {
		return fmt.Errorf("Error setting adapter_name: %s", err)
	}

	cert, err := kpAPI.GetKMIPObject(ctx, adapterNameOrID, objectID)
	if err != nil {
		return err
	}
	err = populateKMIPClientCertSchemaDataFromStruct(d, *cert)
	if err != nil {
		return err
	}
	return nil
}

func populateKMIPObjectSchemaDataFromStruct(d *schema.ResourceData, object kp.GetKMIPObject) (err error) {
	if err = d.Set("name", object.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err = d.Set("certificate", object.Certificate); err != nil {
		return fmt.Errorf("Error setting certificate: %s", err)
	}
	if err = d.Set("created_at", object.CreatedAt.String()); err != nil {
		return fmt.Errorf("Error setting created_at: %s", err)
	}
	if err = d.Set("created_by", object.CreatedBy); err != nil {
		return fmt.Errorf("Error setting created_by: %s", err)
	}
	if err = d.Set("created_by_cert_id", object.CreatedByCertID); err != nil {
		return fmt.Errorf("Error setting created_by_cert_id: %s", err)
	}
	if err = d.Set("updated_at", object.UpdatedAt.String()); err != nil {
		return fmt.Errorf("Error setting updated_at: %s", err)
	}
	if err = d.Set("updated_by", object.UpdatedBy); err != nil {
		return fmt.Errorf("Error setting created_by: %s", err)
	}
	if err = d.Set("updated_by_cert_id", object.UpdatedByCertID); err != nil {
		return fmt.Errorf("Error setting updated_by_cert_id: %s", err)
	}
	if err = d.Set("destroyed_at", object.DestroyedAt.String()); err != nil {
		return fmt.Errorf("Error setting destroyed_at: %s", err)
	}
	if err = d.Set("destroyed_by", object.DestroyedBy); err != nil {
		return fmt.Errorf("Error setting destroyed_by: %s", err)
	}
	if err = d.Set("destroyed_by_cert_id", object.DestroyedByCertID); err != nil {
		return fmt.Errorf("Error setting destroyed_by_cert_id: %s", err)
	}

	// d.SetID(object.ID)
	return nil
}
