// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	resourceIBMKmsKMIPAdapterValidProfiles = []string{"native_1.0"}
)

func ResourceIBMKmsKMIPAdapter() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMKmsKMIPAdapterCreate,
		Read:     resourceIBMKmsKMIPAdapterRead,
		Delete:   resourceIBMKmsKMIPAdapterDelete,
		Exists:   resourceIBMKmsKMIPAdapterExists,
		Importer: &schema.ResourceImporter{},
		// Timeouts: &schema.ResourceTimeout{
		// 	Create: schema.DefaultTimeout(10 * time.Minute),
		// 	Update: schema.DefaultTimeout(10 * time.Minute),
		// },
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Key protect Instance GUID",
				ForceNew:         true,
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"profile": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "The profile of the KMIP adapter",
				ValidateFunc: validate.ValidateAllowedStringValues(resourceIBMKmsKMIPAdapterValidProfiles),
			},
			"profile_data": {
				Type:        schema.TypeMap,
				Required:    true,
				ForceNew:    true,
				Description: "The data specific to the KMIP Adapter profile",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The name of the KMIP adapter",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the KMIP adapter",
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
		},
	}
}

func resourceIBMKmsKMIPAdapterProfileToProfileFunc(profile string, profileData map[string]string) kp.CreateKMIPAdapterProfile {
	if profile == resourceIBMKmsKMIPAdapterValidProfiles[0] {
		//native_1.0
		return kp.WithNativeProfile(profileData["crk_id"])
	}
	// Shouldn't reach here, since we check for profile validity before this.
	return nil
}

func resourceIBMKmsKMIPAdapterCreate(d *schema.ResourceData, meta interface{}) error {
	adapterToCreate, instanceID, err := ExtractAndValidateKMIPAdapterDataFromSchema(d)
	if err != nil {
		return err
	}
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	adapter, err := kpAPI.CreateKMIPAdapter(context.Background(),
		resourceIBMKmsKMIPAdapterProfileToProfileFunc(adapterToCreate.Profile, adapterToCreate.ProfileData),
		kp.WithKMIPAdapterName(adapterToCreate.Name),
		kp.WithKMIPAdapterDescription(adapterToCreate.Description),
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Error while creating KMIP adapter: %s", err)
	}
	populateKMIPAdapterSchemaDataFromStruct(d, *adapter)
	return nil
}

func resourceIBMKmsKMIPAdapterRead(d *schema.ResourceData, meta interface{}) error {
	instanceID := d.Get("instance_id").(string)
	adapterID := d.Id()
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	ctx := context.Background()
	adapter, err := kpAPI.GetKMIPAdapter(ctx, adapterID)
	if err != nil {
		return err
	}
	err = populateKMIPAdapterSchemaDataFromStruct(d, *adapter)
	if err != nil {
		return err
	}
	return nil
}

func resourceIBMKmsKMIPAdapterUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMKmsKMIPAdapterRead(d, meta)
}

func resourceIBMKmsKMIPAdapterDelete(d *schema.ResourceData, meta interface{}) error {
	instanceID := d.Get("instance_id").(string)
	adapterID := d.Id()
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	ctx := context.Background()
	objects, err := kpAPI.GetKMIPObjects(ctx, adapterID, nil)
	for _, object := range objects.Objects {
		err = kpAPI.DeleteKMIPObject(ctx, adapterID, object.ID)
		if err != nil {
			return fmt.Errorf("[ERROR] Failed to delete KMIP object associated with adapter (%s): %s",
				adapterID,
				err,
			)
		}
	}

	err = kpAPI.DeleteKMIPAdapter(ctx, adapterID)
	return err
}

func resourceIBMKmsKMIPAdapterExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	instanceID := d.Get("instance_id").(string)
	adapterID := d.Id()
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return false, err
	}
	ctx := context.Background()
	_, err = kpAPI.GetKMIPAdapter(ctx, adapterID)
	if err != nil {
		kpError := err.(*kp.Error)
		if kpError.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func ExtractAndValidateKMIPAdapterDataFromSchema(d *schema.ResourceData) (adapter kp.KMIPAdapter, instanceID string, err error) {
	err = nil
	instanceID = getInstanceIDFromResourceData(d, "instance_id")
	profile, ok := d.Get("profile").(string)
	if !ok {
		err = fmt.Errorf("Error converting profile to string")
		return
	}
	adapter = kp.KMIPAdapter{
		Profile: profile,
	}
	if name, ok := d.GetOk("name"); ok {
		nameStr, ok2 := name.(string)
		if !ok2 {
			err = fmt.Errorf("Error converting name to string")
			return
		}
		adapter.Name = nameStr
	}
	if desc, ok := d.GetOk("description"); ok {
		descStr, ok2 := desc.(string)
		if !ok2 {
			err = fmt.Errorf("Error converting description to string")
			return
		}
		adapter.Description = descStr
	}
	if data, ok := d.GetOk("profile_data"); ok {
		dataMap, ok2 := data.(map[string]interface{})
		if !ok2 {
			err = fmt.Errorf("Error converting profile data to map[string]interface{}")
			return
		}
		profileData := map[string]string{}
		for key := range dataMap {
			if val, ok := dataMap[key].(string); ok {
				profileData[key] = val
			} else {
				err = fmt.Errorf("Error converting value with key {%s} into string", key)
				return
			}
		}
		adapter.ProfileData = profileData
	}
	return
}

func populateKMIPAdapterSchemaDataFromStruct(d *schema.ResourceData, adapter kp.KMIPAdapter) (err error) {
	d.SetId(adapter.ID)
	if err = d.Set("name", adapter.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err = d.Set("description", adapter.Description); err != nil {
		return fmt.Errorf("Error setting description: %s", err)
	}
	if err = d.Set("profile", adapter.Profile); err != nil {
		return fmt.Errorf("Error setting profile: %s", err)
	}
	if err = d.Set("profile_data", adapter.ProfileData); err != nil {
		return fmt.Errorf("Error setting profile_data: %s", err)
	}
	if err = d.Set("created_at", adapter.CreatedAt.String()); err != nil {
		return fmt.Errorf("Error setting created_at: %s", err)
	}
	if err = d.Set("created_by", adapter.CreatedBy); err != nil {
		return fmt.Errorf("Error setting created_by: %s", err)
	}
	if err = d.Set("updated_at", adapter.UpdatedAt.String()); err != nil {
		return fmt.Errorf("Error setting updated_at: %s", err)
	}
	if err = d.Set("updated_by", adapter.UpdatedBy); err != nil {
		return fmt.Errorf("Error setting updated_by: %s", err)
	}
	return nil
}
