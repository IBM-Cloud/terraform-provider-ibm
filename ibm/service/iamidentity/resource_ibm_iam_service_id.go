// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func ResourceIBMIamServiceID() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamServiceIDCreate,
		ReadContext:   resourceIBMIamServiceIDRead,
		UpdateContext: resourceIBMIamServiceIDUpdate,
		DeleteContext: resourceIBMIamServiceIDDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the Service Id. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the Service Id.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The optional description of the Service Id. The 'description' property is only available if a description was provided during a create of a Service Id.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud wide identifier for identities of this service ID.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the account.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::serviceid:1234-5678-9012'.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the ServiceID object.",
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"locked": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The service ID cannot be changed if set to true.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set contains a date time string of the creation date in ISO format.",
			},
			"modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set contains a date time string of the last modification date in ISO format.",
			},
		},
	}
}

func resourceIBMIamServiceIDCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_service_id", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	accountID := userDetails.UserAccount

	createServiceIDOptions := &iamidentityv1.CreateServiceIDOptions{}
	createServiceIDOptions.SetAccountID(accountID)
	createServiceIDOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		createServiceIDOptions.SetDescription(d.Get("description").(string))
	}

	serviceID, _, err := iamIdentityClient.CreateServiceIDWithContext(context, createServiceIDOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateServiceIDWithContext failed: %s", err.Error()), "ibm_iam_service_id", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if serviceID == nil {
		err = fmt.Errorf("Create ServiceID failed: %s", d.Get("name").(string))
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_service_id", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*serviceID.ID)

	return resourceIBMIamServiceIDRead(context, d, meta)
}

func resourceIBMIamServiceIDRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getServiceIDOptions := &iamidentityv1.GetServiceIDOptions{}

	getServiceIDOptions.SetID(d.Id())

	serviceID, response, err := iamIdentityClient.GetServiceIDWithContext(context, getServiceIDOptions)
	if err != nil || serviceID == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetServiceIDWithContext failed: %s", err.Error()), "ibm_iam_service_id", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("account_id", serviceID.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-account_id").GetDiag()
	}
	if err = d.Set("name", serviceID.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-name").GetDiag()
	}
	if !core.IsNil(serviceID.Description) {
		if err = d.Set("description", serviceID.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("iam_id", serviceID.IamID); err != nil {
		err = fmt.Errorf("Error setting iam_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-iam_id").GetDiag()
	}
	if err = d.Set("crn", serviceID.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-crn").GetDiag()
	}
	if err = d.Set("version", serviceID.EntityTag); err != nil {
		err = fmt.Errorf("Error setting version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-version").GetDiag()
	}
	if err = d.Set("locked", serviceID.Locked); err != nil {
		err = fmt.Errorf("Error setting locked: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-locked").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(serviceID.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("modified_at", flex.DateTimeToString(serviceID.ModifiedAt)); err != nil {
		err = fmt.Errorf("Error setting modified_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-modified_at").GetDiag()
	}
	return nil
}

func resourceIBMIamServiceIDUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateServiceIDOptions := &iamidentityv1.UpdateServiceIDOptions{}

	updateServiceIDOptions.SetID(d.Id())
	updateServiceIDOptions.SetIfMatch("*")
	if _, ok := d.GetOk("name"); ok {
		updateServiceIDOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		updateServiceIDOptions.SetDescription(d.Get("description").(string))
	}

	_, _, err = iamIdentityClient.UpdateServiceIDWithContext(context, updateServiceIDOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateServiceIDWithContext failed: %s", err.Error()), "ibm_iam_service_id", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamServiceIDRead(context, d, meta)
}

func resourceIBMIamServiceIDDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteServiceIDOptions := &iamidentityv1.DeleteServiceIDOptions{}

	deleteServiceIDOptions.SetID(d.Id())

	_, err = iamIdentityClient.DeleteServiceIDWithContext(context, deleteServiceIDOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteServiceID failed: %s", err.Error()), "ibm_iam_service_id", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
