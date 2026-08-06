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

func ResourceIBMIamAPIKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamAPIKeyCreate,
		ReadContext:   resourceIBMIamAPIKeyRead,
		UpdateContext: resourceIBMIamAPIKeyUpdate,
		DeleteContext: resourceIBMIamAPIKeyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"entity_lock": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "false",
				Description: "Indicates if the API key is locked for further write operations. False by default.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key.",
			},
			"expires_at": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Date and time when the API key becomes invalid, ISO 8601 datetime in the format 'yyyy-MM-ddTHH:mm+0000'. **WARNING** An API key will be permanently and irrevocably deleted when both the expires_at and modified_at timestamps are more than ninety (90) days in the past, regardless of the key’s locked status or any other state.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
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
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "You can optionally passthrough the API key value for this API key. If passed, NO validation of that apiKey value is done, i.e. the value can be non-URL safe. If omitted, the API key management will create an URL safe opaque API key value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing in this value.",
			},
			"store_value": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing of API keys for users.",
			},
			"file": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "File where api key is to be stored",
			},
			"apikey_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of this API Key.",
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
		},
	}
}

func resourceIBMIamAPIKeyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createAPIKeyOptions := &iamidentityv1.CreateAPIKeyOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "create", "get-user-details").GetDiag()
	}
	iamID := userDetails.UserID
	accountID := userDetails.UserAccount

	createAPIKeyOptions.SetName(d.Get("name").(string))
	createAPIKeyOptions.SetIamID(iamID)
	createAPIKeyOptions.SetAccountID(accountID)

	if _, ok := d.GetOk("description"); ok {
		createAPIKeyOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("apikey"); ok {
		createAPIKeyOptions.SetApikey(d.Get("apikey").(string))
	}
	if _, ok := d.GetOk("store_value"); ok {
		createAPIKeyOptions.SetStoreValue(d.Get("store_value").(bool))
	}
	if _, ok := d.GetOk("expires_at"); ok {
		createAPIKeyOptions.SetExpiresAt(d.Get("expires_at").(string))
	}
	if _, ok := d.GetOk("entity_lock"); ok {
		createAPIKeyOptions.SetEntityLock(d.Get("entity_lock").(string))
	}

	apiKey, _, err := iamIdentityClient.CreateAPIKeyWithContext(context, createAPIKeyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAPIKeyWithContext failed: %s", err.Error()), "ibm_iam_api_key", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*apiKey.ID)
	d.Set("apikey", *apiKey.Apikey)

	if keyfile, ok := d.GetOk("file"); ok {
		if err := saveToFile(apiKey, keyfile.(string)); err != nil {
			log.Printf("Error writing API Key Details to file: %s", err)
		}
	}

	return resourceIBMIamAPIKeyRead(context, d, meta)
}

func resourceIBMIamAPIKeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

	getAPIKeyOptions.SetID(d.Id())

	apiKey, response, err := iamIdentityClient.GetAPIKeyWithContext(context, getAPIKeyOptions)
	if err != nil || apiKey == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAPIKeyWithContext failed: %s", err.Error()), "ibm_iam_api_key", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", apiKey.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-name").GetDiag()
	}
	if !core.IsNil(apiKey.ExpiresAt) {
		if err = d.Set("expires_at", apiKey.ExpiresAt); err != nil {
			err = fmt.Errorf("Error setting expires_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-expires_at").GetDiag()
		}
	}
	if !core.IsNil(apiKey.Description) {
		if err = d.Set("description", apiKey.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("iam_id", apiKey.IamID); err != nil {
		err = fmt.Errorf("Error setting iam_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-iam_id").GetDiag()
	}
	if !core.IsNil(apiKey.AccountID) {
		if err = d.Set("account_id", apiKey.AccountID); err != nil {
			err = fmt.Errorf("Error setting account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-account_id").GetDiag()
		}
	}
	if err = d.Set("locked", apiKey.Locked); err != nil {
		err = fmt.Errorf("Error setting locked: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-locked").GetDiag()
	}
	if !core.IsNil(apiKey.ID) {
		if err = d.Set("apikey_id", apiKey.ID); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting apikey_id: %s", err))
		}
	}
	if !core.IsNil(apiKey.EntityTag) {
		if err = d.Set("entity_tag", apiKey.EntityTag); err != nil {
			err = fmt.Errorf("Error setting entity_tag: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-entity_tag").GetDiag()
		}
	}
	if err = d.Set("crn", apiKey.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(apiKey.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(apiKey.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-created_at").GetDiag()
		}
	}
	if err = d.Set("created_by", apiKey.CreatedBy); err != nil {
		err = fmt.Errorf("Error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(apiKey.ModifiedAt) {
		if err = d.Set("modified_at", flex.DateTimeToString(apiKey.ModifiedAt)); err != nil {
			err = fmt.Errorf("Error setting modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "read", "set-modified_at").GetDiag()
		}
	}

	return nil
}

func resourceIBMIamAPIKeyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateAPIKeyOptions := &iamidentityv1.UpdateAPIKeyOptions{}

	updateAPIKeyOptions.SetIfMatch("*")
	updateAPIKeyOptions.SetID(d.Id())
	if _, ok := d.GetOk("name"); ok {
		updateAPIKeyOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		updateAPIKeyOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("expires_at"); ok {
		updateAPIKeyOptions.SetExpiresAt(d.Get("expires_at").(string))
	}

	_, _, err = iamIdentityClient.UpdateAPIKeyWithContext(context, updateAPIKeyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateAPIKeyWithContext failed: %s", err.Error()), "ibm_iam_api_key", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamAPIKeyRead(context, d, meta)
}

func resourceIBMIamAPIKeyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_api_key", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteAPIKeyOptions := &iamidentityv1.DeleteAPIKeyOptions{}

	deleteAPIKeyOptions.SetID(d.Id())

	_, err = iamIdentityClient.DeleteAPIKeyWithContext(context, deleteAPIKeyOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteAPIKeyWithContext failed: %s", err.Error()), "ibm_iam_api_key", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
