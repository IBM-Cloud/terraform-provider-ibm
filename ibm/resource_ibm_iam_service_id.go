// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMIAMServiceID() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMServiceIDCreate,
		Read:     resourceIBMIAMServiceIDRead,
		Update:   resourceIBMIAMServiceIDUpdate,
		Delete:   resourceIBMIAMServiceIDDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the serviceID",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the serviceID",
			},

			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "version of the serviceID",
			},

			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "crn of the serviceID",
			},

			"iam_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the serviceID",
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"unique_instance_crns": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceIBMIAMServiceIDCreate(d *schema.ResourceData, meta interface{}) error {
	iamIdentityClient, err := meta.(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}

	name := d.Get("name").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()

	createServiceIDOptions := iamidentityv1.CreateServiceIDOptions{
		Name:      &name,
		AccountID: &userDetails.userAccount,
	}

	if d, ok := d.GetOk("description"); ok {
		des := d.(string)
		createServiceIDOptions.Description = &des
	}
	if a, ok := d.GetOk("unique_instance_crns"); ok {
		crns := expandStringList(a.([]interface{}))
		createServiceIDOptions.UniqueInstanceCrns = crns
	}

	serviceID, resp, err := iamIdentityClient.CreateServiceID(&createServiceIDOptions)
	if err != nil || serviceID == nil {
		return fmt.Errorf("Error creating serviceID: %s, %s", err, resp)
	}
	d.SetId(*serviceID.ID)

	return resourceIBMIAMServiceIDRead(d, meta)
}

func resourceIBMIAMServiceIDRead(d *schema.ResourceData, meta interface{}) error {
	iamIdentityClient, err := meta.(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	serviceIDUUID := d.Id()
	getServiceIDOptions := iamidentityv1.GetServiceIDOptions{
		ID: &serviceIDUUID,
	}
	serviceID, resp, err := iamIdentityClient.GetServiceID(&getServiceIDOptions)
	if err != nil {
		if resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving serviceID: %s %s", err, resp)
	}
	if serviceID.Name != nil {
		d.Set("name", *serviceID.Name)
	}
	if serviceID.Description != nil {
		d.Set("description", *serviceID.Description)
	}
	if serviceID.CRN != nil {
		d.Set("crn", *serviceID.CRN)
	}
	if serviceID.EntityTag != nil {
		d.Set("version", serviceID.EntityTag)
	}
	if serviceID.IamID != nil {
		d.Set("iam_id", serviceID.IamID)
	}
	if serviceID.Locked != nil {
		d.Set("locked", serviceID.Locked)
	}
	if serviceID.UniqueInstanceCrns != nil && len(serviceID.UniqueInstanceCrns) > 0 {
		d.Set("unique_instance_crns", flattenStringList(serviceID.UniqueInstanceCrns))
	}
	return nil
}

func resourceIBMIAMServiceIDUpdate(d *schema.ResourceData, meta interface{}) error {

	iamIdentityClient, err := meta.(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	serviceIDUUID := d.Id()

	hasChange := false
	ifMatch := "*"
	updateServiceIDOptions := iamidentityv1.UpdateServiceIDOptions{
		ID:      &serviceIDUUID,
		IfMatch: &ifMatch,
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateServiceIDOptions.Name = &name
		hasChange = true
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updateServiceIDOptions.Description = &description
		hasChange = true
	}
	if d.HasChange("unique_instance_crns") {
		u := d.Get("unique_instance_crns").([]interface{})
		updateServiceIDOptions.UniqueInstanceCrns = expandStringList(u)
		hasChange = true
	}

	if hasChange {
		_, resp, err := iamIdentityClient.UpdateServiceID(&updateServiceIDOptions)
		if err != nil {
			return fmt.Errorf("Error updating serviceID: %s, %s", err, resp)
		}
	}

	return resourceIBMIAMServiceIDRead(d, meta)

}

func resourceIBMIAMServiceIDDelete(d *schema.ResourceData, meta interface{}) error {
	iamIdentityClient, err := meta.(ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}

	serviceIDUUID := d.Id()
	deleteServiceIDOptions := iamidentityv1.DeleteServiceIDOptions{
		ID: &serviceIDUUID,
	}
	resp, err := iamIdentityClient.DeleteServiceID(&deleteServiceIDOptions)
	if err != nil {
		return fmt.Errorf("Error deleting serviceID: %s %s", err, resp)
	}

	d.SetId("")

	return nil
}
