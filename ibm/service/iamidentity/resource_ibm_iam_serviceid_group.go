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

func ResourceIBMIamServiceidGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamServiceidGroupCreate,
		ReadContext:   resourceIBMIamServiceidGroupRead,
		UpdateContext: resourceIBMIamServiceidGroupUpdate,
		DeleteContext: resourceIBMIamServiceidGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the account the service ID group belongs to.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the service ID group. Unique in the account.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the service ID group.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the service ID group details object. You need to specify this value when updating the service ID group to avoid stale updates.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Resource Name of the item.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the service ID group was created.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM ID of the user or service which created the Service Id group.",
			},
			"modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of when the service ID group was modified.",
			},
		},
	}
}

func resourceIBMIamServiceidGroupCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createServiceIDGroupOptions := &iamidentityv1.CreateServiceIDGroupOptions{}

	createServiceIDGroupOptions.SetAccountID(d.Get("account_id").(string))
	createServiceIDGroupOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		createServiceIDGroupOptions.SetDescription(d.Get("description").(string))
	}

	serviceIDGroup, _, err := iamIdentityClient.CreateServiceIDGroupWithContext(context, createServiceIDGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateServiceIDGroupWithContext failed: %s", err.Error()), "ibm_iam_serviceid_group", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*serviceIDGroup.ID)

	return resourceIBMIamServiceidGroupRead(context, d, meta)
}

func resourceIBMIamServiceidGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getServiceIDGroupOptions := &iamidentityv1.GetServiceIDGroupOptions{}

	getServiceIDGroupOptions.SetID(d.Id())

	serviceIDGroup, response, err := iamIdentityClient.GetServiceIDGroupWithContext(context, getServiceIDGroupOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetServiceIDGroupWithContext failed: %s", err.Error()), "ibm_iam_serviceid_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("account_id", serviceIDGroup.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-account_id").GetDiag()
	}
	if err = d.Set("name", serviceIDGroup.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-name").GetDiag()
	}
	if !core.IsNil(serviceIDGroup.Description) {
		if err = d.Set("description", serviceIDGroup.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(serviceIDGroup.EntityTag) {
		if err = d.Set("entity_tag", serviceIDGroup.EntityTag); err != nil {
			err = fmt.Errorf("Error setting entity_tag: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-entity_tag").GetDiag()
		}
	}
	if err = d.Set("crn", serviceIDGroup.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(serviceIDGroup.CreatedAt) {
		if err = d.Set("created_at", serviceIDGroup.CreatedAt); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-created_at").GetDiag()
		}
	}
	if err = d.Set("created_by", serviceIDGroup.CreatedBy); err != nil {
		err = fmt.Errorf("Error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(serviceIDGroup.ModifiedAt) {
		if err = d.Set("modified_at", serviceIDGroup.ModifiedAt); err != nil {
			err = fmt.Errorf("Error setting modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "read", "set-modified_at").GetDiag()
		}
	}

	return nil
}

func resourceIBMIamServiceidGroupUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateServiceIDGroupOptions := &iamidentityv1.UpdateServiceIDGroupOptions{}

	updateServiceIDGroupOptions.SetID(d.Id())
	updateServiceIDGroupOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		updateServiceIDGroupOptions.SetDescription(d.Get("description").(string))
	}

	_, _, err = iamIdentityClient.UpdateServiceIDGroupWithContext(context, updateServiceIDGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateServiceIDGroupWithContext failed: %s", err.Error()), "ibm_iam_serviceid_group", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamServiceidGroupRead(context, d, meta)
}

func resourceIBMIamServiceidGroupDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_serviceid_group", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteServiceIDGroupOptions := &iamidentityv1.DeleteServiceIDGroupOptions{}

	deleteServiceIDGroupOptions.SetID(d.Id())

	_, err = iamIdentityClient.DeleteServiceIDGroupWithContext(context, deleteServiceIDGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteServiceIDGroupWithContext failed: %s", err.Error()), "ibm_iam_serviceid_group", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
