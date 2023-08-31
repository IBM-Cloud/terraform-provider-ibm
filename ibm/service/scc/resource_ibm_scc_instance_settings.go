// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func ResourceIbmSccInstanceSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSccInstanceSettingsCreate,
		ReadContext:   resourceIbmSccInstanceSettingsRead,
		UpdateContext: resourceIbmSccInstanceSettingsUpdate,
		DeleteContext: resourceIbmSccInstanceSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"event_notifications": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The Event Notifications settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Event Notifications instance CRN.",
						},
						"updated_on": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The date when the Event Notifications connection was updated.",
						},
						"source_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The connected Security and Compliance Center instance CRN.",
						},
						"source_description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "This source is used for integration with IBM Cloud Security and Compliance Center.",
							Description: "The description of the source of the Event Notifications.",
						},
						"source_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "compliance",
							Description: "The name of the source of the Event Notifications.",
						},
					},
				},
			},
			"object_storage": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The Cloud Object Storage settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The connected Cloud Object Storage instance CRN.",
						},
						"bucket": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The connected Cloud Object Storage bucket name.",
						},
						"bucket_location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The connected Cloud Object Storage bucket location.",
						},
						"bucket_endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The connected Cloud Object Storage bucket endpoint.",
						},
						"updated_on": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The date when the bucket connection was updated.",
						},
					},
				},
			},
		},
	}
}

func ResourceIbmSccInstanceSettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_scc_instance_settings", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSccInstanceSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	adminClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	updateSettingsOptions := &securityandcompliancecenterapiv3.UpdateSettingsOptions{}

	if _, ok := d.GetOk("event_notifications"); ok {
		eventNotificationsModel, err := resourceIbmSccInstanceSettingsMapToEventNotifications(d.Get("event_notifications.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateSettingsOptions.SetEventNotifications(eventNotificationsModel)
	}
	if _, ok := d.GetOk("object_storage"); ok {
		objectStorageModel, err := resourceIbmSccInstanceSettingsMapToObjectStorage(d.Get("object_storage.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateSettingsOptions.SetObjectStorage(objectStorageModel)
	}
	_, response, err := adminClient.UpdateSettingsWithContext(context, updateSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateSettingsWithContext failed %s\n%s", err, response))
	}

	service_url := adminClient.GetServiceURL()
	d.SetId(service_url)

	return resourceIbmSccInstanceSettingsRead(context, d, meta)
}

func resourceIbmSccInstanceSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	adminClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{}

	settings, response, err := adminClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSettingsWithContext failed %s\n%s", err, response))
	}

	if !core.IsNil(settings.EventNotifications) {
		eventNotificationsMap, err := resourceIbmSccInstanceSettingsEventNotificationsToMap(settings.EventNotifications)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("event_notifications", []map[string]interface{}{eventNotificationsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting event_notifications: %s", err))
		}
	}
	if !core.IsNil(settings.ObjectStorage) {
		objectStorageMap, err := resourceIbmSccInstanceSettingsObjectStorageToMap(settings.ObjectStorage)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("object_storage", []map[string]interface{}{objectStorageMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting object_storage: %s", err))
		}
	}
	return nil
}

func resourceIbmSccInstanceSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	adminClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	updateSettingsOptions := &securityandcompliancecenterapiv3.UpdateSettingsOptions{}

	hasChange := false

	if d.HasChange("event_notifications") {
		eventNotifications, err := resourceIbmSccInstanceSettingsMapToEventNotifications(d.Get("event_notifications.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateSettingsOptions.SetEventNotifications(eventNotifications)
		hasChange = true
	}
	if d.HasChange("object_storage") {
		objectStorage, err := resourceIbmSccInstanceSettingsMapToObjectStorage(d.Get("object_storage.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateSettingsOptions.SetObjectStorage(objectStorage)
		hasChange = true
	}

	if hasChange {
		_, response, err := adminClient.UpdateSettingsWithContext(context, updateSettingsOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateSettingsWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateSettingsWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmSccInstanceSettingsRead(context, d, meta)
}

func resourceIbmSccInstanceSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	adminClient, err := meta.(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{}

	_, response, err := adminClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSettingsWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmSccInstanceSettingsMapToEventNotifications(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.EventNotifications, error) {
	model := &securityandcompliancecenterapiv3.EventNotifications{}
	if modelMap["instance_crn"] != nil && modelMap["instance_crn"].(string) != "" {
		model.InstanceCrn = core.StringPtr(modelMap["instance_crn"].(string))
	}
	if modelMap["updated_on"] != nil {
		dateTime, err := core.ParseDateTime(modelMap["updated_on"].(string))
		if err != nil {
			return model, err
		}
		model.UpdatedOn = &dateTime
	}
	if modelMap["source_id"] != nil && modelMap["source_id"].(string) != "" {
		model.SourceID = core.StringPtr(modelMap["source_id"].(string))
	}
	if modelMap["source_description"] != nil && modelMap["source_description"].(string) != "" {
		model.SourceDescription = core.StringPtr(modelMap["source_description"].(string))
	}
	if modelMap["source_name"] != nil && modelMap["source_name"].(string) != "" {
		model.SourceName = core.StringPtr(modelMap["source_name"].(string))
	}
	return model, nil
}

func resourceIbmSccInstanceSettingsMapToObjectStorage(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.ObjectStorage, error) {
	model := &securityandcompliancecenterapiv3.ObjectStorage{}
	if modelMap["instance_crn"] != nil && modelMap["instance_crn"].(string) != "" {
		model.InstanceCrn = core.StringPtr(modelMap["instance_crn"].(string))
	}
	if modelMap["bucket"] != nil && modelMap["bucket"].(string) != "" {
		model.Bucket = core.StringPtr(modelMap["bucket"].(string))
	}
	if modelMap["bucket_location"] != nil && modelMap["bucket_location"].(string) != "" {
		model.BucketLocation = core.StringPtr(modelMap["bucket_location"].(string))
	}
	if modelMap["bucket_endpoint"] != nil && modelMap["bucket_endpoint"].(string) != "" {
		model.BucketEndpoint = core.StringPtr(modelMap["bucket_endpoint"].(string))
	}
	if modelMap["updated_on"] != nil {
		dateTime, err := core.ParseDateTime(modelMap["updated_on"].(string))
		if err != nil {
			return model, err
		}
		model.UpdatedOn = &dateTime
	}
	return model, nil
}

func resourceIbmSccInstanceSettingsEventNotificationsToMap(model *securityandcompliancecenterapiv3.EventNotifications) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.InstanceCrn != nil {
		modelMap["instance_crn"] = model.InstanceCrn
	}
	if model.UpdatedOn != nil {
		modelMap["updated_on"] = model.UpdatedOn.String()
	}
	if model.SourceID != nil {
		modelMap["source_id"] = model.SourceID
	}
	if model.SourceDescription != nil {
		modelMap["source_description"] = model.SourceDescription
	}
	if model.SourceName != nil {
		modelMap["source_name"] = model.SourceName
	}
	return modelMap, nil
}

func resourceIbmSccInstanceSettingsObjectStorageToMap(model *securityandcompliancecenterapiv3.ObjectStorage) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.InstanceCrn != nil {
		modelMap["instance_crn"] = model.InstanceCrn
	}
	if model.Bucket != nil {
		modelMap["bucket"] = model.Bucket
	}
	if model.BucketLocation != nil {
		modelMap["bucket_location"] = model.BucketLocation
	}
	if model.BucketEndpoint != nil {
		modelMap["bucket_endpoint"] = model.BucketEndpoint
	}
	if model.UpdatedOn != nil {
		modelMap["updated_on"] = model.UpdatedOn.String()
	}
	return modelMap, nil
}
