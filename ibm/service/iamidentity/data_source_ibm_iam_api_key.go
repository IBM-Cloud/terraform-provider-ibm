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

func DataSourceIBMIamAPIKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamAPIKeyRead,

		Schema: map[string]*schema.Schema{
			"apikey_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID of the API key.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the API Key details object. You need to specify this value when updating the API key to avoid stale updates.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-5678'.",
			},
			"locked": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The API key cannot be changed if set to true.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set contains a date time string of the creation date in ISO format.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM ID of the user or service which created the API key.",
			},
			"modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set contains a date time string of the last modification date in ISO format.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The optional description of the API key. The 'description' property is only available if a description was provided during a create of an API key.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The iam_id that this API key authenticates.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the account that this API key authenticates for.",
			},
		},
	}
}

func dataSourceIBMIamAPIKeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_api_key", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

	getAPIKeyOptions.SetID(d.Get("apikey_id").(string))

	apiKey, _, err := iamIdentityClient.GetAPIKeyWithContext(context, getAPIKeyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAPIKeyWithContext failed: %s", err.Error()), "(Data) ibm_iam_api_key", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*apiKey.ID)

	if !core.IsNil(apiKey.EntityTag) {
		if err = d.Set("entity_tag", apiKey.EntityTag); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_tag: %s", err), "(Data) ibm_iam_api_key", "read", "set-entity_tag").GetDiag()
		}
	}

	if err = d.Set("crn", apiKey.CRN); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting crn: %s", err), "(Data) ibm_iam_api_key", "read", "set-crn").GetDiag()
	}

	if err = d.Set("locked", apiKey.Locked); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting locked: %s", err), "(Data) ibm_iam_api_key", "read", "set-locked").GetDiag()
	}

	if !core.IsNil(apiKey.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(apiKey.CreatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_iam_api_key", "read", "set-created_at").GetDiag()
		}
	}

	if err = d.Set("created_by", apiKey.CreatedBy); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_by: %s", err), "(Data) ibm_iam_api_key", "read", "set-created_by").GetDiag()
	}

	if !core.IsNil(apiKey.ModifiedAt) {
		if err = d.Set("modified_at", flex.DateTimeToString(apiKey.ModifiedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting modified_at: %s", err), "(Data) ibm_iam_api_key", "read", "set-modified_at").GetDiag()
		}
	}

	if err = d.Set("name", apiKey.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_iam_api_key", "read", "set-name").GetDiag()
	}

	if !core.IsNil(apiKey.ExpiresAt) {
		if err = d.Set("expires_at", apiKey.ExpiresAt); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting expires_at: %s", err), "(Data) ibm_iam_api_key", "read", "set-expires_at").GetDiag()
		}
	}

	if !core.IsNil(apiKey.Description) {
		if err = d.Set("description", apiKey.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_iam_api_key", "read", "set-description").GetDiag()
		}
	}

	if err = d.Set("iam_id", apiKey.IamID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting iam_id: %s", err), "(Data) ibm_iam_api_key", "read", "set-iam_id").GetDiag()
	}

	if err = d.Set("account_id", apiKey.AccountID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_id: %s", err), "(Data) ibm_iam_api_key", "read", "set-account_id").GetDiag()
	}

	return nil
}
