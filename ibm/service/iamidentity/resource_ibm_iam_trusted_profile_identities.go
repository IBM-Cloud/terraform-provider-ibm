// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.103.0-e8b84313-20250402-201816
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func ResourceIBMIamTrustedProfileIdentities() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamTrustedProfileIdentitiesCreate,
		ReadContext:   resourceIBMIamTrustedProfileIdentitiesRead,
		UpdateContext: resourceIBMIamTrustedProfileIdentitiesUpdate,
		DeleteContext: resourceIBMIamTrustedProfileIdentitiesDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the trusted profile.",
			},
			"if_match": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity tag of the Identities to be updated. Specify the tag that you retrieved when reading the Profile Identities. This value helps identify parallel usage of this API. Pass * to indicate updating any available version, which may result in stale updates.",
			},
			"identities": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of identities.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iam_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "IAM ID of the identity.",
						},
						"identifier": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Type of the identity.",
						},
						"accounts": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Only valid for the type user. Accounts from which a user can assume the trusted profile.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMIamTrustedProfileIdentitiesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	setProfileIdentitiesOptions := &iamidentityv1.SetProfileIdentitiesOptions{}

	setProfileIdentitiesOptions.SetProfileID(d.Get("profile_id").(string))
	setProfileIdentitiesOptions.SetIfMatch("*")
	if _, ok := d.GetOk("identities"); ok {
		var identities []iamidentityv1.ProfileIdentityRequest
		for _, v := range d.Get("identities").([]interface{}) {
			value := v.(map[string]interface{})
			identitiesItem, err := ResourceIBMIamTrustedProfileIdentitiesMapToProfileIdentityRequest(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "delete", "parse-identities").GetDiag()
			}
			identities = append(identities, *identitiesItem)
		}
		setProfileIdentitiesOptions.SetIdentities(identities)
	}

	_, _, err = iamIdentityClient.SetProfileIdentitiesWithContext(context, setProfileIdentitiesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetProfileIdentitiesWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile_identities", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*setProfileIdentitiesOptions.ProfileID)

	return resourceIBMIamTrustedProfileIdentitiesRead(context, d, meta)
}

func resourceIBMIamTrustedProfileIdentitiesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getProfileIdentitiesOptions := &iamidentityv1.GetProfileIdentitiesOptions{}

	getProfileIdentitiesOptions.SetProfileID(d.Id())

	profileIdentitiesResponse, response, err := iamIdentityClient.GetProfileIdentitiesWithContext(context, getProfileIdentitiesOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProfileIdentitiesWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile_identities", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err := d.Set("profile_id", d.Id()); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting profile_id: %s", err))
	}
	if !core.IsNil(profileIdentitiesResponse.Identities) {
		identities := []map[string]interface{}{}
		for _, identitiesItem := range profileIdentitiesResponse.Identities {
			identitiesItemMap, err := ResourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(&identitiesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "read", "identities-to-map").GetDiag()
			}
			identities = append(identities, identitiesItemMap)
		}
		if err = d.Set("identities", identities); err != nil {
			err = fmt.Errorf("Error setting identities: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "read", "set-identities").GetDiag()
		}

	} else {
		// Always set an empty list if no identities
		if err = d.Set("identities", []map[string]interface{}{}); err != nil {
			err = fmt.Errorf("Error setting identities: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "read", "set-identities").GetDiag()
		}
	}
	if err = d.Set("if_match", response.Headers.Get("ETag")); err != nil {
		return diag.FromErr(flex.FmtErrorf("Error setting etag: %s", err))
	}

	return nil
}

func resourceIBMIamTrustedProfileIdentitiesUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	setProfileIdentitiesOptions := &iamidentityv1.SetProfileIdentitiesOptions{}

	setProfileIdentitiesOptions.SetProfileID(d.Get("profile_id").(string))
	setProfileIdentitiesOptions.SetIfMatch(d.Get("if_match").(string))
	var identities []iamidentityv1.ProfileIdentityRequest
	setProfileIdentitiesOptions.SetIdentities(identities)

	_, _, err = iamIdentityClient.SetProfileIdentitiesWithContext(context, setProfileIdentitiesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetProfileIdentitiesWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile_identities", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", *setProfileIdentitiesOptions.ProfileID))

	return resourceIBMIamTrustedProfileIdentitiesRead(context, d, meta)
}

func resourceIBMIamTrustedProfileIdentitiesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile_identities", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	setProfileIdentitiesOptions := &iamidentityv1.SetProfileIdentitiesOptions{}
	setProfileIdentitiesOptions.SetIfMatch(d.Get("if_match").(string))
	setProfileIdentitiesOptions.SetIdentities([]iamidentityv1.ProfileIdentityRequest{})

	// This is a workaround to delete all identities from the profile
	// as the SetProfileIdentities API does not support deleting identities directly.
	// Instead, we set an empty list of identities to remove all existing identities.
	setProfileIdentitiesOptions.SetProfileID(d.Id())

	_, _, err = iamIdentityClient.SetProfileIdentitiesWithContext(context, setProfileIdentitiesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("SetProfileIdentitiesWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile_identities", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIamTrustedProfileIdentitiesMapToProfileIdentityRequest(modelMap map[string]interface{}) (*iamidentityv1.ProfileIdentityRequest, error) {
	model := &iamidentityv1.ProfileIdentityRequest{}
	model.Identifier = core.StringPtr(modelMap["identifier"].(string))
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["accounts"] != nil {
		accounts := []string{}
		for _, accountsItem := range modelMap["accounts"].([]interface{}) {
			accounts = append(accounts, accountsItem.(string))
		}
		model.Accounts = accounts
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	return model, nil
}

func ResourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(model *iamidentityv1.ProfileIdentityResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["iam_id"] = *model.IamID
	modelMap["identifier"] = *model.Identifier
	modelMap["type"] = *model.Type
	if model.Accounts != nil {
		modelMap["accounts"] = model.Accounts
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	return modelMap, nil
}
