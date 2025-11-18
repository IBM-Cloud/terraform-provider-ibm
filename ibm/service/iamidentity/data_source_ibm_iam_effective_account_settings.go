// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
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

func DataSourceIBMIamEffectiveAccountSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamEffectiveAccountSettingsRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID of the account.",
			},
			"include_history": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Defines if the entity history is included in the response.",
			},
			"resolve_user_mfa": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enrich MFA exemptions with user information.",
			},
			"effective": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"restrict_create_service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_create_platform_apikey": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_user_list_visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.",
						},
						"allowed_ip_addresses": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the IP addresses and subnets from which IAM tokens can be created for the account.",
						},
						"mfa": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"user_mfa": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of users that are exempted from the MFA requirement of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iam_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The iam_id of the user.",
									},
									"mfa": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "name of the user account.",
									},
									"user_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "userName of the user.",
									},
									"email": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "email of the user.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "optional description.",
									},
								},
							},
						},
						"session_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.",
						},
						"session_invalidation_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.",
						},
						"max_sessions_per_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.",
						},
						"system_access_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.",
						},
						"system_refresh_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.",
						},
					},
				},
			},
			"account": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Input body parameters for the Account Settings REST request.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entity_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version of the account settings.",
						},
						"history": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "History of the Account Settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"timestamp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Timestamp when the action was triggered.",
									},
									"iam_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IAM ID of the identity which triggered the action.",
									},
									"iam_id_account": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Account of the identity which triggered the action.",
									},
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action of the history entry.",
									},
									"params": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Params of the history entry.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Message which summarizes the executed action.",
									},
								},
							},
						},
						"restrict_create_service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_create_platform_apikey": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_user_list_visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.",
						},
						"restrict_user_domains": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Defines if account invitations are restricted to specified domains. To remove an entry for a realm_id, perform an update (PUT) request with only the realm_id set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"realm_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The realm that the restrictions apply to.",
									},
									"invitation_email_allow_patterns": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"restrict_invitation": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted.",
									},
								},
							},
						},
						"allowed_ip_addresses": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the IP addresses and subnets from which IAM tokens can be created for the account.",
						},
						"mfa": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"session_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.",
						},
						"session_invalidation_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.",
						},
						"max_sessions_per_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.",
						},
						"system_access_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.",
						},
						"system_refresh_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.",
						},
						"user_mfa": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of users that are exempted from the MFA requirement of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iam_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The iam_id of the user.",
									},
									"mfa": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "name of the user account.",
									},
									"user_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "userName of the user.",
									},
									"email": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "email of the user.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "optional description.",
									},
								},
							},
						},
					},
				},
			},
			"assigned_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "assigned template section.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template Id.",
						},
						"template_version": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Template version.",
						},
						"template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template name.",
						},
						"restrict_create_service_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"restrict_create_platform_apikey": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.",
						},
						"allowed_ip_addresses": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the IP addresses and subnets from which IAM tokens can be created for the account.",
						},
						"mfa": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"session_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.",
						},
						"session_invalidation_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.",
						},
						"max_sessions_per_identity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.",
						},
						"system_access_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.",
						},
						"system_refresh_token_expiration_in_seconds": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.",
						},
						"user_mfa": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of users that are exempted from the MFA requirement of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"iam_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The iam_id of the user.",
									},
									"mfa": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "name of the user account.",
									},
									"user_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "userName of the user.",
									},
									"email": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "email of the user.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "optional description.",
									},
								},
							},
						},
						"restrict_user_list_visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.",
						},
						"restrict_user_domains": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_sufficient": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restrictions": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Defines if account invitations are restricted to specified domains. To remove an entry for a realm_id, perform an update (PUT) request with only the realm_id set.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"realm_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The realm that the restrictions apply to.",
												},
												"invitation_email_allow_patterns": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"restrict_invitation": &schema.Schema{
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIamEffectiveAccountSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_effective_account_settings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEffectiveAccountSettingsOptions := &iamidentityv1.GetEffectiveAccountSettingsOptions{}

	getEffectiveAccountSettingsOptions.SetAccountID(d.Get("account_id").(string))
	if _, ok := d.GetOk("include_history"); ok {
		getEffectiveAccountSettingsOptions.SetIncludeHistory(d.Get("include_history").(bool))
	}
	if _, ok := d.GetOk("resolve_user_mfa"); ok {
		getEffectiveAccountSettingsOptions.SetResolveUserMfa(d.Get("resolve_user_mfa").(bool))
	}

	effectiveAccountSettingsResponse, _, err := iamIdentityClient.GetEffectiveAccountSettingsWithContext(context, getEffectiveAccountSettingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEffectiveAccountSettingsWithContext failed: %s", err.Error()), "(Data) ibm_iam_effective_account_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getEffectiveAccountSettingsOptions.AccountID)

	var effective []map[string]interface{}
	effectiveMap, err := DataSourceIBMIamEffectiveAccountSettingsAccountSettingsEffectiveSectionToMap(effectiveAccountSettingsResponse.Effective)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_effective_account_settings", "read", "effective-to-map").GetDiag()
	}
	effective = append(effective, effectiveMap)
	if err = d.Set("effective", effective); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting effective: %s", err), "(Data) ibm_iam_effective_account_settings", "read", "set-effective").GetDiag()
	}

	var account []map[string]interface{}
	accountMap, err := DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAccountSectionToMap(effectiveAccountSettingsResponse.Account)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_effective_account_settings", "read", "account-to-map").GetDiag()
	}
	account = append(account, accountMap)
	if err = d.Set("account", account); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account: %s", err), "(Data) ibm_iam_effective_account_settings", "read", "set-account").GetDiag()
	}

	if !core.IsNil(effectiveAccountSettingsResponse.AssignedTemplates) {
		var assignedTemplates []map[string]interface{}
		for _, assignedTemplatesItem := range effectiveAccountSettingsResponse.AssignedTemplates {
			assignedTemplatesItemMap, err := DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAssignedTemplatesSectionToMap(&assignedTemplatesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_effective_account_settings", "read", "assigned_templates-to-map").GetDiag()
			}
			assignedTemplates = append(assignedTemplates, assignedTemplatesItemMap)
		}
		if err = d.Set("assigned_templates", assignedTemplates); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting assigned_templates: %s", err), "(Data) ibm_iam_effective_account_settings", "read", "set-assigned_templates").GetDiag()
		}
	}

	return nil
}

func DataSourceIBMIamEffectiveAccountSettingsAccountSettingsEffectiveSectionToMap(model *iamidentityv1.AccountSettingsEffectiveSection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RestrictCreateServiceID != nil {
		modelMap["restrict_create_service_id"] = *model.RestrictCreateServiceID
	}
	if model.RestrictCreatePlatformApikey != nil {
		modelMap["restrict_create_platform_apikey"] = *model.RestrictCreatePlatformApikey
	}
	if model.RestrictUserListVisibility != nil {
		modelMap["restrict_user_list_visibility"] = *model.RestrictUserListVisibility
	}
	if model.AllowedIPAddresses != nil {
		modelMap["allowed_ip_addresses"] = *model.AllowedIPAddresses
	}
	if model.Mfa != nil {
		modelMap["mfa"] = *model.Mfa
	}
	if model.UserMfa != nil {
		var userMfa []map[string]interface{}
		for _, userMfaItem := range model.UserMfa {
			userMfaItemMap, err := AccountSettingsUserMfaResponseToMap(&userMfaItem)
			if err != nil {
				return modelMap, err
			}
			userMfa = append(userMfa, userMfaItemMap)
		}
		modelMap["user_mfa"] = userMfa
	}
	if model.SessionExpirationInSeconds != nil {
		modelMap["session_expiration_in_seconds"] = *model.SessionExpirationInSeconds
	}
	if model.SessionInvalidationInSeconds != nil {
		modelMap["session_invalidation_in_seconds"] = *model.SessionInvalidationInSeconds
	}
	if model.MaxSessionsPerIdentity != nil {
		modelMap["max_sessions_per_identity"] = *model.MaxSessionsPerIdentity
	}
	if model.SystemAccessTokenExpirationInSeconds != nil {
		modelMap["system_access_token_expiration_in_seconds"] = *model.SystemAccessTokenExpirationInSeconds
	}
	if model.SystemRefreshTokenExpirationInSeconds != nil {
		modelMap["system_refresh_token_expiration_in_seconds"] = *model.SystemRefreshTokenExpirationInSeconds
	}
	return modelMap, nil
}

func DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAccountSectionToMap(model *iamidentityv1.AccountSettingsResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	//modelMap["account_id"] = "testString"
	modelMap["entity_tag"] = *model.EntityTag
	if model.History != nil {
		var history []map[string]interface{}
		for _, historyItem := range model.History {
			historyItemMap, err := EnityHistoryRecordToMap(&historyItem)
			if err != nil {
				return modelMap, err
			}
			history = append(history, historyItemMap)
		}
		modelMap["history"] = history
	}
	modelMap["restrict_create_service_id"] = *model.RestrictCreateServiceID
	modelMap["restrict_create_platform_apikey"] = *model.RestrictCreatePlatformApikey
	modelMap["restrict_user_list_visibility"] = *model.RestrictUserListVisibility

	var restrictUserDomains []map[string]interface{}
	for _, restrictUserDomainsItem := range model.RestrictUserDomains {
		restrictUserDomainsItemMap, err := AccountSettingsUserDomainRestrictionToMap(&restrictUserDomainsItem)
		if err != nil {
			return modelMap, err
		}
		restrictUserDomains = append(restrictUserDomains, restrictUserDomainsItemMap)
	}
	modelMap["restrict_user_domains"] = restrictUserDomains
	modelMap["allowed_ip_addresses"] = *model.AllowedIPAddresses
	modelMap["mfa"] = *model.Mfa
	modelMap["session_expiration_in_seconds"] = *model.SessionExpirationInSeconds
	modelMap["session_invalidation_in_seconds"] = *model.SessionInvalidationInSeconds
	modelMap["max_sessions_per_identity"] = *model.MaxSessionsPerIdentity
	modelMap["system_access_token_expiration_in_seconds"] = *model.SystemAccessTokenExpirationInSeconds
	modelMap["system_refresh_token_expiration_in_seconds"] = *model.SystemRefreshTokenExpirationInSeconds

	var userMfa []map[string]interface{}
	for _, userMfaItem := range model.UserMfa {
		userMfaItemMap, err := AccountSettingsUserMfaResponseToMap(&userMfaItem)
		if err != nil {
			return modelMap, err
		}
		userMfa = append(userMfa, userMfaItemMap)
	}
	modelMap["user_mfa"] = userMfa

	return modelMap, nil
}

func DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAssignedTemplatesSectionToMap(model *iamidentityv1.AccountSettingsAssignedTemplatesSection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["template_id"] = *model.TemplateID
	modelMap["template_version"] = flex.IntValue(model.TemplateVersion)
	modelMap["template_name"] = *model.TemplateName
	if model.RestrictCreateServiceID != nil {
		modelMap["restrict_create_service_id"] = *model.RestrictCreateServiceID
	}
	if model.RestrictCreatePlatformApikey != nil {
		modelMap["restrict_create_platform_apikey"] = *model.RestrictCreatePlatformApikey
	}
	if model.RestrictUserListVisibility != nil {
		modelMap["restrict_user_list_visibility"] = *model.RestrictUserListVisibility
	}
	if model.RestrictUserDomains != nil {
		restrictUserDomainsMap, err := DataSourceIBMEffectiveAccountSettingsAssignedTemplatesAccountSettingsRestrictUserDomainsToMap(model.RestrictUserDomains)
		if err != nil {
			return modelMap, err
		}
		modelMap["restrict_user_domains"] = []map[string]interface{}{restrictUserDomainsMap}
	}
	if model.AllowedIPAddresses != nil {
		modelMap["allowed_ip_addresses"] = *model.AllowedIPAddresses
	}
	if model.Mfa != nil {
		modelMap["mfa"] = *model.Mfa
	}
	if model.SessionExpirationInSeconds != nil {
		modelMap["session_expiration_in_seconds"] = *model.SessionExpirationInSeconds
	}
	if model.SessionInvalidationInSeconds != nil {
		modelMap["session_invalidation_in_seconds"] = *model.SessionInvalidationInSeconds
	}
	if model.MaxSessionsPerIdentity != nil {
		modelMap["max_sessions_per_identity"] = *model.MaxSessionsPerIdentity
	}
	if model.SystemAccessTokenExpirationInSeconds != nil {
		modelMap["system_access_token_expiration_in_seconds"] = *model.SystemAccessTokenExpirationInSeconds
	}
	if model.SystemRefreshTokenExpirationInSeconds != nil {
		modelMap["system_refresh_token_expiration_in_seconds"] = *model.SystemRefreshTokenExpirationInSeconds
	}
	if model.UserMfa != nil {
		var userMfa []map[string]interface{}
		for _, userMfaItem := range model.UserMfa {
			userMfaItemMap, err := AccountSettingsUserMfaResponseToMap(&userMfaItem)
			if err != nil {
				return modelMap, err
			}
			userMfa = append(userMfa, userMfaItemMap)
		}
		modelMap["user_mfa"] = userMfa
	}
	return modelMap, nil
}

func DataSourceIBMEffectiveAccountSettingsAssignedTemplatesAccountSettingsRestrictUserDomainsToMap(model *iamidentityv1.AssignedTemplatesAccountSettingsRestrictUserDomains) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountSufficient != nil {
		modelMap["account_sufficient"] = *model.AccountSufficient
	}
	if model.Restrictions != nil {
		restrictions := []map[string]interface{}{}
		for _, restrictionsItem := range model.Restrictions {
			restrictionsItemMap, err := DataSourceIBMEffectiveAccountSettingsAccountSettingsUserDomainRestrictionToMap(&restrictionsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			restrictions = append(restrictions, restrictionsItemMap)
		}
		modelMap["restrictions"] = restrictions
	}
	return modelMap, nil
}

func DataSourceIBMEffectiveAccountSettingsAccountSettingsUserDomainRestrictionToMap(model *iamidentityv1.AccountSettingsUserDomainRestriction) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["realm_id"] = *model.RealmID
	if model.InvitationEmailAllowPatterns != nil {
		modelMap["invitation_email_allow_patterns"] = model.InvitationEmailAllowPatterns
	}
	if model.RestrictInvitation != nil {
		modelMap["restrict_invitation"] = *model.RestrictInvitation
	}
	return modelMap, nil
}
