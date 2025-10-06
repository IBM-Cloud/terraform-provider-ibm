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

func DataSourceIBMIamAccountSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamAccountSettingsRead,

		Schema: map[string]*schema.Schema{
			"include_history": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Defines if the entity history is included in the response.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique ID of the account.",
			},
			"restrict_create_service_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Defines whether or not creating a Service Id is access controlled. Valid values:  * RESTRICTED - to apply access control  * NOT_RESTRICTED - to remove access control  * NOT_SET - to 'unset' a previous set value.",
			},
			"restrict_create_platform_apikey": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Defines whether or not creating platform API keys is access controlled. Valid values:  * RESTRICTED - to apply access control  * NOT_RESTRICTED - to remove access control  * NOT_SET - to 'unset' a previous set value.",
			},
			"allowed_ip_addresses": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Defines the IP addresses and subnets from which IAM tokens can be created for the account.",
			},
			"resolve_user_mfa": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enrich MFA exemptions with user PI.",
			},
			"entity_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the account settings.",
			},
			"mfa": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
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
				Description: "Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '2592000'  * NOT_SET - To unset account setting and use service default.",
			},
		},
	}
}

func dataSourceIBMIamAccountSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_account_settings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAccountSettingsOptions := &iamidentityv1.GetAccountSettingsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	getAccountSettingsOptions.SetAccountID(userDetails.UserAccount)

	if _, ok := d.GetOk("include_history"); ok {
		getAccountSettingsOptions.SetIncludeHistory(d.Get("include_history").(bool))
	}
	if _, ok := d.GetOk("resolve_user_mfa"); ok {
		getAccountSettingsOptions.SetResolveUserMfa(d.Get("resolve_user_mfa").(bool))
	}

	accountSettingsResponse, _, err := iamIdentityClient.GetAccountSettingsWithContext(context, getAccountSettingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAccountSettingsWithContext failed: %s", err.Error()), "(Data) ibm_iam_account_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*getAccountSettingsOptions.AccountID)

	if err = d.Set("account_id", accountSettingsResponse.AccountID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_id: %s", err), "(Data) ibm_iam_account_settings", "read", "set-account_id").GetDiag()
	}

	if err = d.Set("entity_tag", accountSettingsResponse.EntityTag); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_tag: %s", err), "(Data) ibm_iam_account_settings", "read", "set-entity_tag").GetDiag()
	}

	if !core.IsNil(accountSettingsResponse.History) {
		history := []map[string]interface{}{}
		for _, historyItem := range accountSettingsResponse.History {
			historyItemMap, err := EnityHistoryRecordToMap(&historyItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_account_settings", "read", "history-to-map").GetDiag()
			}
			history = append(history, historyItemMap)
		}
		if err = d.Set("history", history); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting history: %s", err), "(Data) ibm_iam_account_settings", "read", "set-history").GetDiag()
		}
	}

	if err = d.Set("restrict_create_service_id", accountSettingsResponse.RestrictCreateServiceID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting restrict_create_service_id: %s", err), "(Data) ibm_iam_account_settings", "read", "set-restrict_create_service_id").GetDiag()
	}

	if err = d.Set("restrict_create_platform_apikey", accountSettingsResponse.RestrictCreatePlatformApikey); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting restrict_create_platform_apikey: %s", err), "(Data) ibm_iam_account_settings", "read", "set-restrict_create_platform_apikey").GetDiag()
	}

	if err = d.Set("restrict_user_list_visibility", accountSettingsResponse.RestrictUserListVisibility); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting restrict_user_list_visibility: %s", err), "(Data) ibm_iam_account_settings", "read", "set-restrict_user_list_visibility").GetDiag()
	}

	restrictUserDomains := []map[string]interface{}{}
	for _, restrictUserDomainsItem := range accountSettingsResponse.RestrictUserDomains {
		restrictUserDomainsItemMap, err := AccountSettingsUserDomainRestrictionToMap(&restrictUserDomainsItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_account_settings", "read", "restrict_user_domains-to-map").GetDiag()
		}
		restrictUserDomains = append(restrictUserDomains, restrictUserDomainsItemMap)
	}

	if err = d.Set("restrict_user_domains", restrictUserDomains); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting restrict_user_domains: %s", err), "(Data) ibm_iam_account_settings", "read", "set-restrict_user_domains").GetDiag()
	}

	if err = d.Set("allowed_ip_addresses", accountSettingsResponse.AllowedIPAddresses); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting allowed_ip_addresses: %s", err), "(Data) ibm_iam_account_settings", "read", "set-allowed_ip_addresses").GetDiag()
	}

	if err = d.Set("mfa", accountSettingsResponse.Mfa); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting mfa: %s", err), "(Data) ibm_iam_account_settings", "read", "set-mfa").GetDiag()
	}

	if err = d.Set("session_expiration_in_seconds", accountSettingsResponse.SessionExpirationInSeconds); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting session_expiration_in_seconds: %s", err), "(Data) ibm_iam_account_settings", "read", "set-session_expiration_in_seconds").GetDiag()
	}

	if err = d.Set("session_invalidation_in_seconds", accountSettingsResponse.SessionInvalidationInSeconds); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting session_invalidation_in_seconds: %s", err), "(Data) ibm_iam_account_settings", "read", "set-session_invalidation_in_seconds").GetDiag()
	}

	if err = d.Set("max_sessions_per_identity", accountSettingsResponse.MaxSessionsPerIdentity); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting max_sessions_per_identity: %s", err), "(Data) ibm_iam_account_settings", "read", "set-max_sessions_per_identity").GetDiag()
	}

	if err = d.Set("system_access_token_expiration_in_seconds", accountSettingsResponse.SystemAccessTokenExpirationInSeconds); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting system_access_token_expiration_in_seconds: %s", err), "(Data) ibm_iam_account_settings", "read", "set-system_access_token_expiration_in_seconds").GetDiag()
	}

	if err = d.Set("system_refresh_token_expiration_in_seconds", accountSettingsResponse.SystemRefreshTokenExpirationInSeconds); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting system_refresh_token_expiration_in_seconds: %s", err), "(Data) ibm_iam_account_settings", "read", "set-system_refresh_token_expiration_in_seconds").GetDiag()
	}

	userMfa := []map[string]interface{}{}
	for _, userMfaItem := range accountSettingsResponse.UserMfa {
		userMfaItemMap, err := AccountSettingsUserMfaResponseToMap(&userMfaItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_account_settings", "read", "user_mfa-to-map").GetDiag()
		}
		userMfa = append(userMfa, userMfaItemMap)
	}
	if err = d.Set("user_mfa", userMfa); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_mfa: %s", err), "(Data) ibm_iam_account_settings", "read", "set-user_mfa").GetDiag()
	}

	return nil
}
