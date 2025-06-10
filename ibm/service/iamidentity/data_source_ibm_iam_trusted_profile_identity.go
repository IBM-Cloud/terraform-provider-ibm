// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func DataSourceIBMIamTrustedProfileIdentity() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamTrustedProfileIdentityRead,

		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the trusted profile.",
			},
			"identity_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the identity.",
			},
			"identifier_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier of the identity that can assume the trusted profiles.",
			},
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
	}
}

func dataSourceIBMIamTrustedProfileIdentityRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_trusted_profile_identity", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getProfileIdentityOptions := &iamidentityv1.GetProfileIdentityOptions{}

	getProfileIdentityOptions.SetProfileID(d.Get("profile_id").(string))
	getProfileIdentityOptions.SetIdentityType(d.Get("identity_type").(string))
	getProfileIdentityOptions.SetIdentifierID(d.Get("identifier_id").(string))

	profileIdentityResponse, _, err := iamIdentityClient.GetProfileIdentityWithContext(context, getProfileIdentityOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProfileIdentityWithContext failed: %s", err.Error()), "(Data) ibm_iam_trusted_profile_identity", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIamTrustedProfileIdentityID(d))

	if err = d.Set("iam_id", profileIdentityResponse.IamID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting iam_id: %s", err), "(Data) ibm_iam_trusted_profile_identity", "read", "set-iam_id").GetDiag()
	}

	if err = d.Set("identifier", profileIdentityResponse.Identifier); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting identifier: %s", err), "(Data) ibm_iam_trusted_profile_identity", "read", "set-identifier").GetDiag()
	}

	if err = d.Set("type", profileIdentityResponse.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_iam_trusted_profile_identity", "read", "set-type").GetDiag()
	}

	if !core.IsNil(profileIdentityResponse.Accounts) {
		accounts := []interface{}{}
		for _, accountsItem := range profileIdentityResponse.Accounts {
			accounts = append(accounts, accountsItem)
		}
		if err = d.Set("accounts", accounts); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting accounts: %s", err), "(Data) ibm_iam_trusted_profile_identity", "read", "set-accounts").GetDiag()
		}
	}

	if !core.IsNil(profileIdentityResponse.Description) {
		if err = d.Set("description", profileIdentityResponse.Description); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_iam_trusted_profile_identity", "read", "set-description").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMIamTrustedProfileIdentityID returns a reasonable ID for the list.
func dataSourceIBMIamTrustedProfileIdentityID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
