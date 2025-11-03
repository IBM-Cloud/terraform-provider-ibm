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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func DataSourceIBMIamTrustedProfileIdentities() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamTrustedProfileIdentitiesRead,

		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the trusted profile.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity tag of the profile identities response.",
			},
			"identities": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of identities.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iam_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IAM ID of the identity.",
						},
						"identifier": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the identity.",
						},
						"accounts": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Only valid for the type user. Accounts from which a user can assume the trusted profile.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIamTrustedProfileIdentitiesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_trusted_profile_identities", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getProfileIdentitiesOptions := &iamidentityv1.GetProfileIdentitiesOptions{}

	getProfileIdentitiesOptions.SetProfileID(d.Get("profile_id").(string))

	profileIdentitiesResponse, _, err := iamIdentityClient.GetProfileIdentitiesWithContext(context, getProfileIdentitiesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProfileIdentitiesWithContext failed: %s", err.Error()), "(Data) ibm_iam_trusted_profile_identities", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getProfileIdentitiesOptions.ProfileID)

	if !core.IsNil(profileIdentitiesResponse.EntityTag) {
		if err = d.Set("entity_tag", profileIdentitiesResponse.EntityTag); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_tag: %s", err), "(Data) ibm_iam_trusted_profile_identities", "read", "set-entity_tag").GetDiag()
		}
	}

	if !core.IsNil(profileIdentitiesResponse.Identities) {
		identities := []map[string]interface{}{}
		for _, identitiesItem := range profileIdentitiesResponse.Identities {
			identitiesItemMap, err := DataSourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(&identitiesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_trusted_profile_identities", "read", "identities-to-map").GetDiag()
			}
			identities = append(identities, identitiesItemMap)
		}
		if err = d.Set("identities", identities); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting identities: %s", err), "(Data) ibm_iam_trusted_profile_identities", "read", "set-identities").GetDiag()
		}
	}

	return nil
}

func DataSourceIBMIamTrustedProfileIdentitiesProfileIdentityResponseToMap(model *iamidentityv1.ProfileIdentityResponse) (map[string]interface{}, error) {
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
