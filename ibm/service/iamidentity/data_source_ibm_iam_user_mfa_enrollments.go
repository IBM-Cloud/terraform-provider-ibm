// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func DataSourceIBMIamUserMfaEnrollments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamUserMfaEnrollmentsRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the account.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "iam_id of the user. This user must be the member of the account.",
			},
			"effective_mfa_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "currently effective mfa type i.e. id_based_mfa or account_based_mfa.",
			},
			"id_based_mfa": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trait_account_default": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"trait_user_specific": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"trait_effective": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"complies": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The enrollment complies to the effective requirement.",
						},
					},
				},
			},
			"account_based_mfa": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_questions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Describes whether the enrollment type is required.",
									},
									"enrolled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Describes whether the enrollment type is enrolled.",
									},
								},
							},
						},
						"totp": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Describes whether the enrollment type is required.",
									},
									"enrolled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Describes whether the enrollment type is enrolled.",
									},
								},
							},
						},
						"verisign": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"required": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Describes whether the enrollment type is required.",
									},
									"enrolled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Describes whether the enrollment type is enrolled.",
									},
								},
							},
						},
						"complies": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The enrollment complies to the effective requirement.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIamUserMfaEnrollmentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_user_mfa_enrollments", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getMfaStatusOptions := &iamidentityv1.GetMfaStatusOptions{}

	getMfaStatusOptions.SetAccountID(d.Get("account_id").(string))
	getMfaStatusOptions.SetIamID(d.Get("iam_id").(string))

	userMfaEnrollments, response, err := iamIdentityClient.GetMfaStatusWithContext(context, getMfaStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetMfaStatusWithContext failed: %s", err.Error()), "ibm_iam_user_mfa_enrollments", "read")
		log.Printf("[DEBUG]\n%s\n%s", tfErr.GetDebugMessage(), response)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIamUserMfaEnrollmentsID(d))

	if err = d.Set("effective_mfa_type", userMfaEnrollments.EffectiveMfaType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_user_mfa_enrollments", "read", "set-effective_mfa_type").GetDiag()
	}

	idBasedMfa := []map[string]interface{}{}
	if userMfaEnrollments.IDBasedMfa != nil {
		modelMap, err := dataSourceIBMIamUserMfaEnrollmentsIDBasedMfaEnrollmentToMap(userMfaEnrollments.IDBasedMfa)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_user_mfa_enrollments", "read", "mfa-enrollment-to-map").GetDiag()
		}
		idBasedMfa = append(idBasedMfa, modelMap)
	}
	if err = d.Set("id_based_mfa", idBasedMfa); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_user_mfa_enrollments", "read", "set-id_based_mfa").GetDiag()
	}

	accountBasedMfa := []map[string]interface{}{}
	if userMfaEnrollments.AccountBasedMfa != nil {
		modelMap, err := dataSourceIBMIamUserMfaEnrollmentsAccountBasedMfaEnrollmentToMap(userMfaEnrollments.AccountBasedMfa)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_user_mfa_enrollments", "read", "mfa-enrollment-to-map").GetDiag()
		}
		accountBasedMfa = append(accountBasedMfa, modelMap)
	}
	if err = d.Set("account_based_mfa", accountBasedMfa); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_user_mfa_enrollments", "read", "set-account_based_mfa").GetDiag()
	}

	return nil
}

// dataSourceIBMIamUserMfaEnrollmentsID returns a reasonable ID for the list.
func dataSourceIBMIamUserMfaEnrollmentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIBMIamUserMfaEnrollmentsIDBasedMfaEnrollmentToMap(model *iamidentityv1.IDBasedMfaEnrollment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["trait_account_default"] = model.TraitAccountDefault
	if model.TraitUserSpecific != nil {
		modelMap["trait_user_specific"] = model.TraitUserSpecific
	}
	modelMap["trait_effective"] = model.TraitEffective
	modelMap["complies"] = model.Complies
	return modelMap, nil
}

func dataSourceIBMIamUserMfaEnrollmentsAccountBasedMfaEnrollmentToMap(model *iamidentityv1.AccountBasedMfaEnrollment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	securityQuestionsMap, err := dataSourceIBMIamUserMfaEnrollmentsMfaEnrollmentTypeStatusToMap(model.SecurityQuestions)
	if err != nil {
		return modelMap, err
	}
	modelMap["security_questions"] = []map[string]interface{}{securityQuestionsMap}
	totpMap, err := dataSourceIBMIamUserMfaEnrollmentsMfaEnrollmentTypeStatusToMap(model.Totp)
	if err != nil {
		return modelMap, err
	}
	modelMap["totp"] = []map[string]interface{}{totpMap}
	verisignMap, err := dataSourceIBMIamUserMfaEnrollmentsMfaEnrollmentTypeStatusToMap(model.Verisign)
	if err != nil {
		return modelMap, err
	}
	modelMap["verisign"] = []map[string]interface{}{verisignMap}
	modelMap["complies"] = model.Complies
	return modelMap, nil
}

func dataSourceIBMIamUserMfaEnrollmentsMfaEnrollmentTypeStatusToMap(model *iamidentityv1.MfaEnrollmentTypeStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["required"] = model.Required
	modelMap["enrolled"] = model.Enrolled
	return modelMap, nil
}
