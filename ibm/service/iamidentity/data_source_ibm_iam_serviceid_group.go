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

func DataSourceIBMIamServiceidGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamServiceidGroupRead,

		Schema: map[string]*schema.Schema{
			"iam_serviceid_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID of the service ID group.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the service ID group details object. You need to specify this value when updating the service ID group to avoid stale updates.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the account the service ID group belongs to.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Resource Name of the item.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the service ID group. Unique in the account.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the service ID group.",
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

func dataSourceIBMIamServiceidGroupRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_serviceid_group", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getServiceIDGroupOptions := &iamidentityv1.GetServiceIDGroupOptions{}

	getServiceIDGroupOptions.SetID(d.Get("iam_serviceid_group_id").(string))

	serviceIDGroup, _, err := iamIdentityClient.GetServiceIDGroupWithContext(context, getServiceIDGroupOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetServiceIDGroupWithContext failed: %s", err.Error()), "(Data) ibm_iam_serviceid_group", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getServiceIDGroupOptions.ID)

	if !core.IsNil(serviceIDGroup.EntityTag) {
		if err = d.Set("entity_tag", serviceIDGroup.EntityTag); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_tag: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-entity_tag").GetDiag()
		}
	}

	if err = d.Set("account_id", serviceIDGroup.AccountID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_id: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-account_id").GetDiag()
	}

	if err = d.Set("crn", serviceIDGroup.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-crn").GetDiag()
	}

	if err = d.Set("name", serviceIDGroup.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-name").GetDiag()
	}

	if !core.IsNil(serviceIDGroup.Description) {
		if err = d.Set("description", serviceIDGroup.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-description").GetDiag()
		}
	}

	if !core.IsNil(serviceIDGroup.CreatedAt) {
		if err = d.Set("created_at", serviceIDGroup.CreatedAt); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-created_at").GetDiag()
		}
	}

	if err = d.Set("created_by", serviceIDGroup.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_by: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(serviceIDGroup.ModifiedAt) {
		if err = d.Set("modified_at", serviceIDGroup.ModifiedAt); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting modified_at: %s", err), "(Data) ibm_iam_serviceid_group", "read", "set-modified_at").GetDiag()
		}
	}

	return nil
}
