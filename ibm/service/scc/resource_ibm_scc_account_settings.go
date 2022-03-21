// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v3/adminserviceapiv1"
)

func ResourceIBMSccAccountSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSccAccountSettingsCreate,
		ReadContext:   resourceIbmSccAccountSettingsRead,
		UpdateContext: resourceIbmSccAccountSettingsUpdate,
		DeleteContext: resourceIbmSccAccountSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"location_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_scc_account_settings", "location_id"),
				Description:  "The programatic ID of the location that you want to work in.",
				Deprecated:   "use location instead",
			},
			"location": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true, // Made this Required to avoid drift
				ConflictsWith: []string{"location_id"},
				Description: "Location Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location_id": &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							Description:  "The programatic ID of the location that you want to work in.",
							ValidateFunc: validate.InvokeValidator("ibm_scc_account_settings", "location_id"),
						},
					},
				},
			},
			"event_notifications": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true, // Made this Required to avoid drift during testing
				Description: "The Event Notification settings to register.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "",
							Description: "The Cloud Resource Name (CRN) of the Event Notifications instance that you want to connect.",
						},
					},
				},
			},
		},
	}
}

func resourceIbmSccAccountSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Starting resourceIbmSccAccountSettings%s \n", "Create")
	adminServiceApiClient, err := meta.(conns.ClientSession).AdminServiceApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	// Get the available body that you can put from the SDK
	patchAccountSettingsOptions := &adminserviceapiv1.PatchAccountSettingsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the patchSettings to use userAccount tied to the API_KEY
	patchAccountSettingsOptions.SetAccountID(userDetails.UserAccount)

	getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{}
	getSettingsOptions.SetAccountID(userDetails.UserAccount)

	// Check with GetSettings what the current setting is
	accountSettings, response, err := adminServiceApiClient.GetSettingsWithContext(context, getSettingsOptions)

	hasChange := false
	// check from the local tf file is location is defined
	if _, ok := d.GetOk("location"); ok {
		location, err := resourceIbmSccAccountSettingsMapToLocationID(d.Get("location.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		// if GetSettings is different than the terrafrom config file, prepare a PATCH call
		if location.ID != accountSettings.Location.ID {
			patchAccountSettingsOptions.SetLocation(location)
			hasChange = true
		}
	}

	// check from the local tf file if event_notifications is defined
	event_obj := d.Get("event_notifications.0").(map[string]interface{})
	if _, ok := d.GetOk("event_notifications"); ok && event_obj["instance_crn"] != nil {
		eventNotifications, err := resourceIbmSccAccountSettingsMapToNotificationsRegistration(d.Get("event_notifications.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		// if GetSettings is different than the terrafrom config file, prepare a PATCH call
		if eventNotifications.InstanceCrn != event_obj["instance_crn"] {
			patchAccountSettingsOptions.SetEventNotifications(eventNotifications)
			hasChange = true
		}
	}

	// use scc-go-sdk to send the PATCH request if there is a change
	if hasChange {
		_, response, err = adminServiceApiClient.PatchAccountSettingsWithContext(context, patchAccountSettingsOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchAccountSettingsWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchAccountSettingsWithContext failed %s\n%s", err, response))
		}
		// Set the ID of the Terraform object to the current timestamp
		d.SetId(resourceIbmSccAccountSettingID(d))
	}

	return resourceIbmSccAccountSettingsRead(context, d, meta)
}

func resourceIbmSccAccountSettingID(d *schema.ResourceData) string {
	// make a unique ID according to the timestamp
	return time.Now().UTC().String()
}

func resourceIbmSccAccountSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Starting resourceIbmSccAccountSettings%s \n", "Read")
	adminServiceApiClient, err := meta.(conns.ClientSession).AdminServiceApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	// Get the Settings to call GetSettings
	getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	getSettingsOptions.SetAccountID(userDetails.UserAccount)

	// Return back the current Settings according to GetSettings
	accountSettings, response, err := adminServiceApiClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSettingsWithContext failed %s\n%s", err, response))
	}

	if accountSettings.Location != nil {
		locationMap, err := resourceIbmSccAccountSettingsLocationIDToMap(accountSettings.Location)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("location", []map[string]interface{}{locationMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
		}
	}
	if accountSettings.EventNotifications != nil {
		eventNotificationsMap, err := resourceIbmSccAccountSettingsNotificationsRegistrationToMap(accountSettings.EventNotifications)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("event_notifications", []map[string]interface{}{eventNotificationsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting event_notifications during the read: %s", err))
		}
	}

	return nil
}

func resourceIbmSccAccountSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Starting resourceIbmSccAccountSettings%s \n", "Update")
	adminServiceApiClient, err := meta.(conns.ClientSession).AdminServiceApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	// Use the same logic as resourceIbmSccAccountSettingsCreate
	patchAccountSettingsOptions := &adminserviceapiv1.PatchAccountSettingsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}
	patchAccountSettingsOptions.SetAccountID(userDetails.UserAccount)

	// Flag to see if anything has been changed from the Update(terraform apply)
	hasChange := false

	if d.HasChange("location") {
		location, err := resourceIbmSccAccountSettingsMapToLocationID(d.Get("location.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchAccountSettingsOptions.SetLocation(location)
		hasChange = true
	}

	if d.HasChange("event_notifications") {
		eventNotifications, err := resourceIbmSccAccountSettingsMapToNotificationsRegistration(d.Get("event_notifications.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		patchAccountSettingsOptions.SetEventNotifications(eventNotifications)
		hasChange = true
	}

	if hasChange {
		_, response, err := adminServiceApiClient.PatchAccountSettingsWithContext(context, patchAccountSettingsOptions)
		if err != nil {
			log.Printf("[DEBUG] PatchAccountSettingsWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("PatchAccountSettingsWithContext failed %s\n%s", err, response))
		}
		// Change the id of the this object only if anything has changed
		d.SetId(resourceIbmSccAccountSettingID(d))
	}

	return resourceIbmSccAccountSettingsRead(context, d, meta)
}

func resourceIbmSccAccountSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Use GetSettings since there is no API to delete the configuration of the AccountSettings and avoid compiler warnings
	adminServiceApiClient, err := meta.(conns.ClientSession).AdminServiceApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getSettingsOptions := &adminserviceapiv1.GetSettingsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	getSettingsOptions.SetAccountID(userDetails.UserAccount)

	_, response, err := adminServiceApiClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] PatchAccountSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PatchAccountSettingsWithContext failed %s\n%s", err, response))
	}
	// Set the object to a empty string so Terraform deletes the object
	d.SetId("")

	return nil
}

func resourceIbmSccAccountSettingsMapToLocationID(modelMap map[string]interface{}) (*adminserviceapiv1.LocationID, error) {
	model := &adminserviceapiv1.LocationID{}
	model.ID = core.StringPtr(modelMap["location_id"].(string))
	return model, nil
}

func resourceIbmSccAccountSettingsMapToNotificationsRegistration(modelMap map[string]interface{}) (*adminserviceapiv1.NotificationsRegistration, error) {
	model := &adminserviceapiv1.NotificationsRegistration{}
	model.InstanceCrn = core.StringPtr(modelMap["instance_crn"].(string))
	return model, nil
}

func resourceIbmSccAccountSettingsLocationIDToMap(model *adminserviceapiv1.LocationID) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["location_id"] = model.ID
	return modelMap, nil
}

func resourceIbmSccAccountSettingsNotificationsRegistrationToMap(model *adminserviceapiv1.NotificationsRegistration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["instance_crn"] = model.InstanceCrn
	return modelMap, nil
}
