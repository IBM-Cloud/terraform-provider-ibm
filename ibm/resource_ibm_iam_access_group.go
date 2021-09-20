// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMIAMAccessGroup() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMAccessGroupCreate,
		Read:     resourceIBMIAMAccessGroupRead,
		Update:   resourceIBMIAMAccessGroupUpdate,
		Delete:   resourceIBMIAMAccessGroupDelete,
		Exists:   resourceIBMIAMAccessGroupExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the access group",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the access group",
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMIAMAccessGroupCreate(d *schema.ResourceData, meta interface{}) error {
	iamAccessGroupsClient, err := meta.(ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	creatAccessGroupOptions := iamAccessGroupsClient.NewCreateAccessGroupOptions(userDetails.userAccount, name)
	if des, ok := d.GetOk("description"); ok {
		description := des.(string)
		creatAccessGroupOptions.Description = &description
	}
	agrp, detailedResponse, err := iamAccessGroupsClient.CreateAccessGroup(creatAccessGroupOptions)
	if err != nil || agrp == nil {
		return fmt.Errorf("Error creating access group: %s. API Response: %s", err, detailedResponse)
	}

	d.SetId(*agrp.ID)

	return resourceIBMIAMAccessGroupRead(d, meta)
}

func resourceIBMIAMAccessGroupRead(d *schema.ResourceData, meta interface{}) error {
	iamAccessGroupsClient, err := meta.(ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	agrpID := d.Id()
	getAccessGroupOptions := iamAccessGroupsClient.NewGetAccessGroupOptions(agrpID)
	agrp, detailedResponse, err := iamAccessGroupsClient.GetAccessGroup(getAccessGroupOptions)
	if err != nil || agrp == nil {
		return fmt.Errorf("Error retrieving access group: %s. API Response: %s", err, detailedResponse)
	}
	version := detailedResponse.GetHeaders().Get("etag")
	d.Set("name", agrp.Name)
	d.Set("description", agrp.Description)
	d.Set("version", version)

	return nil
}

func resourceIBMIAMAccessGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	iamAccessGroupsClient, err := meta.(ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}
	agrpID := d.Id()

	hasChange := false
	version := d.Get("version").(string)
	updateAccessGroupOptions := iamAccessGroupsClient.NewUpdateAccessGroupOptions(agrpID, version)

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateAccessGroupOptions.Name = &name
		hasChange = true
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateAccessGroupOptions.Description = &description
		hasChange = true
	}

	if hasChange {
		agrp, detailedResponse, err := iamAccessGroupsClient.UpdateAccessGroup(updateAccessGroupOptions)
		if err != nil || agrp == nil {
			return fmt.Errorf("Error updating access group: %s. API Response: %s", err, detailedResponse)
		}
	}

	return resourceIBMIAMAccessGroupRead(d, meta)

}

func resourceIBMIAMAccessGroupDelete(d *schema.ResourceData, meta interface{}) error {
	iamAccessGroupsClient, err := meta.(ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return err
	}

	agID := d.Id()
	force := true
	deleteAccessGroupOptions := iamAccessGroupsClient.NewDeleteAccessGroupOptions(agID)
	deleteAccessGroupOptions.SetForce(force)
	detailedResponse, err := iamAccessGroupsClient.DeleteAccessGroup(deleteAccessGroupOptions)
	if err != nil {
		return fmt.Errorf("Error deleting access group: %s, API Response: %s", err, detailedResponse)
	}

	d.SetId("")

	return nil
}

func resourceIBMIAMAccessGroupExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iamAccessGroupsClient, err := meta.(ClientSession).IAMAccessGroupsV2()
	if err != nil {
		return false, err
	}
	agID := d.Id()
	getAccessGroupOptions := iamAccessGroupsClient.NewGetAccessGroupOptions(agID)
	agrp, detailedResponse, err := iamAccessGroupsClient.GetAccessGroup(getAccessGroupOptions)
	if err != nil || agrp == nil {
		if detailedResponse != nil && detailedResponse.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with the API: %s, API Response: %s", err, detailedResponse)
	}

	return *agrp.ID == agID, nil
}
