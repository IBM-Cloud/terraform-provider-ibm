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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

const (
	accountSettings            = "ibm_iam_account_settings"
	restrictCreateServiceId    = "restrict_create_service_id"
	restrictCreateApiKey       = "restrict_create_platform_apikey"
	restrictUserListVisibility = "restrict_user_list_visibility"
	mfa                        = "mfa"
)

func ResourceIBMIamAccountSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamAccountSettingsCreate,
		ReadContext:   resourceIBMIamAccountSettingsRead,
		UpdateContext: resourceIBMIamAccountSettingsUpdate,
		DeleteContext: resourceIBMIamAccountSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
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
				Description: "Enrich MFA exemptions with user PI.",
			},
			"restrict_create_service_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator(accountSettings, "restrict_create_service_id"),
				Description:  "Defines whether or not creating a Service Id is access controlled. Valid values:  * RESTRICTED - to apply access control  * NOT_RESTRICTED - to remove access control  * NOT_SET - to 'unset' a previous set value.",
			},
			"restrict_create_platform_apikey": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator(accountSettings, "restrict_create_platform_apikey"),
				Description:  "Defines whether or not creating platform API keys is access controlled. Valid values:  * RESTRICTED - to apply access control  * NOT_RESTRICTED - to remove access control  * NOT_SET - to 'unset' a previous set value.",
			},
			"restrict_user_list_visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator(accountSettings, "restrict_user_list_visibility"),
				Description:  "Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.",
			},
			"restrict_user_domains": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Defines if account invitations are restricted to specified domains. To remove an entry for a realm_id, perform an update (PUT) request with only the realm_id set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"realm_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The realm that the restrictions apply to.",
						},
						"invitation_email_allow_patterns": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"restrict_invitation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted.",
						},
					},
				},
			},
			"allowed_ip_addresses": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defines the IP addresses and subnets from which IAM tokens can be created for the account.",
			},
			"entity_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Version of the account settings.",
			},
			"mfa": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator(accountSettings, "mfa"),
				Description:  "Defines the MFA trait for the account. Valid values:  * NONE - No MFA trait set  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
			},
			"if_match": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "*",
				Description: "Version of the account settings to be updated. Specify the version that you retrieved as entity_tag (ETag header) when reading the account. This value helps identifying parallel usage of this API. Pass * to indicate to update any version available. This might result in stale updates.",
			},
			"user_mfa": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "List of users that are exempted from the MFA requirement of the account.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iam_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The iam_id of the user.",
						},
						"mfa": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Defines the MFA requirement for the user. Valid values:  * NONE - No MFA trait set  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "name of the user account.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "userName of the user.",
						},
						"email": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "email of the user.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
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
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message which summarizes the executed action.",
						},
					},
				},
			},
			"session_expiration_in_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.",
			},
			"session_invalidation_in_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.",
			},
			"max_sessions_per_identity": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.",
			},
			"system_access_token_expiration_in_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.",
			},
			"system_refresh_token_expiration_in_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '2592000'  * NOT_SET - To unset account setting and use service default.",
			},
		},
	}
}

func ResourceIBMIAMAccountSettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)

	restrict_values := "RESTRICTED, NOT_RESTRICTED, NOT_SET"
	user_visibility_restrict_values := "RESTRICTED, NOT_RESTRICTED"
	mfa_values := "NONE, NONE_NO_ROPC, TOTP, TOTP4ALL, LEVEL1, LEVEL2, LEVEL3"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 restrictUserListVisibility,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              user_visibility_restrict_values})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 restrictCreateServiceId,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              restrict_values})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 restrictCreateApiKey,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              restrict_values})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 mfa,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              mfa_values})

	ibmIAMAccountSettingsValidator := validate.ResourceValidator{ResourceName: "ibm_iam_account_settings", Schema: validateSchema}
	return &ibmIAMAccountSettingsValidator
}

func resourceIBMIamAccountSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return diag.FromErr(err)
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

	accountSettingsResponse, response, err := iamIdentityClient.GetAccountSettings(getAccountSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetAccountSettings failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*accountSettingsResponse.AccountID)

	return resourceIBMIamAccountSettingsUpdate(context, d, meta)
}

func resourceIBMIamAccountSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAccountSettingsOptions := &iamidentityv1.GetAccountSettingsOptions{}

	getAccountSettingsOptions.SetAccountID(d.Id())
	getAccountSettingsOptions.SetIncludeHistory(d.Get("include_history").(bool))

	if _, ok := d.GetOk("resolve_user_mfa"); ok {
		getAccountSettingsOptions.SetResolveUserMfa(d.Get("resolve_user_mfa").(bool))
	}

	accountSettingsResponse, response, err := iamIdentityClient.GetAccountSettingsWithContext(context, getAccountSettingsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAccountSettingsWithContext failed: %s", err.Error()), "ibm_iam_account_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if err = d.Set("entity_tag", accountSettingsResponse.EntityTag); err != nil {
		err = fmt.Errorf("Error setting entity_tag: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-entity_tag").GetDiag()
	}
	if !core.IsNil(accountSettingsResponse.History) {
		history := []map[string]interface{}{}
		for _, historyItem := range accountSettingsResponse.History {
			historyItemMap, _ := EnityHistoryRecordToMap(&historyItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "history-to-map").GetDiag()
			}
			history = append(history, historyItemMap)
		}
		if err = d.Set("history", history); err != nil {
			err = fmt.Errorf("Error setting history: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-history").GetDiag()
		}
	}
	if err = d.Set("restrict_create_service_id", accountSettingsResponse.RestrictCreateServiceID); err != nil {
		err = fmt.Errorf("Error setting restrict_create_service_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-restrict_create_service_id").GetDiag()
	}
	if err = d.Set("restrict_create_platform_apikey", accountSettingsResponse.RestrictCreatePlatformApikey); err != nil {
		err = fmt.Errorf("Error setting restrict_create_platform_apikey: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-restrict_create_platform_apikey").GetDiag()
	}
	if err = d.Set("restrict_user_list_visibility", accountSettingsResponse.RestrictUserListVisibility); err != nil {
		err = fmt.Errorf("Error setting restrict_user_list_visibility: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-restrict_user_list_visibility").GetDiag()
	}
	restrictUserDomains := []map[string]interface{}{}
	for _, restrictUserDomainsItem := range accountSettingsResponse.RestrictUserDomains {
		restrictUserDomainsItemMap, err := AccountSettingsUserDomainRestrictionToMap(&restrictUserDomainsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "restrict_user_domains-to-map").GetDiag()
		}
		restrictUserDomains = append(restrictUserDomains, restrictUserDomainsItemMap)
	}
	if err = d.Set("restrict_user_domains", restrictUserDomains); err != nil {
		err = fmt.Errorf("Error setting restrict_user_domains: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-restrict_user_domains").GetDiag()
	}
	if err = d.Set("allowed_ip_addresses", accountSettingsResponse.AllowedIPAddresses); err != nil {
		err = fmt.Errorf("Error setting allowed_ip_addresses: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-allowed_ip_addresses").GetDiag()
	}
	if err = d.Set("mfa", accountSettingsResponse.Mfa); err != nil {
		err = fmt.Errorf("Error setting mfa: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-mfa").GetDiag()
	}
	if err = d.Set("session_expiration_in_seconds", accountSettingsResponse.SessionExpirationInSeconds); err != nil {
		err = fmt.Errorf("Error setting session_expiration_in_seconds: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-session_expiration_in_seconds").GetDiag()
	}
	if err = d.Set("session_invalidation_in_seconds", accountSettingsResponse.SessionInvalidationInSeconds); err != nil {
		err = fmt.Errorf("Error setting session_invalidation_in_seconds: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-session_invalidation_in_seconds").GetDiag()
	}
	if err = d.Set("max_sessions_per_identity", accountSettingsResponse.MaxSessionsPerIdentity); err != nil {
		err = fmt.Errorf("Error setting max_sessions_per_identity: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-max_sessions_per_identity").GetDiag()
	}
	if err = d.Set("system_access_token_expiration_in_seconds", accountSettingsResponse.SystemAccessTokenExpirationInSeconds); err != nil {
		err = fmt.Errorf("Error setting system_access_token_expiration_in_seconds: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-system_access_token_expiration_in_seconds").GetDiag()
	}
	if err = d.Set("system_refresh_token_expiration_in_seconds", accountSettingsResponse.SystemRefreshTokenExpirationInSeconds); err != nil {
		err = fmt.Errorf("Error setting system_refresh_token_expiration_in_seconds: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-system_refresh_token_expiration_in_seconds").GetDiag()
	}
	userMfa := []map[string]interface{}{}
	for _, userMfaItem := range accountSettingsResponse.UserMfa {
		userMfaItemMap, err := AccountSettingsUserMfaResponseToMap(&userMfaItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "user_mfa-to-map").GetDiag()
		}
		userMfa = append(userMfa, userMfaItemMap)
	}
	if err = d.Set("user_mfa", userMfa); err != nil {
		err = fmt.Errorf("Error setting user_mfa: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "read", "set-user_mfa").GetDiag()
	}

	return nil
}

func resourceIBMIamAccountSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateAccountSettingsOptions := &iamidentityv1.UpdateAccountSettingsOptions{}

	updateAccountSettingsOptions.SetAccountID(d.Id())
	updateAccountSettingsOptions.SetIfMatch(d.Get("if_match").(string))

	hasChange := false

	if d.HasChange("allowed_ip_addresses") {
		allowed_ip_addresses_str := d.Get("allowed_ip_addresses").(string)
		updateAccountSettingsOptions.SetAllowedIPAddresses(allowed_ip_addresses_str)
		hasChange = true
	}

	if d.HasChange("restrict_create_service_id") {
		restrict_create_service_id_str := d.Get("restrict_create_service_id").(string)
		updateAccountSettingsOptions.SetRestrictCreateServiceID(restrict_create_service_id_str)
		hasChange = true
	}

	if d.HasChange("restrict_create_platform_apikey") {
		restrict_create_platform_apikey_str := d.Get("restrict_create_platform_apikey").(string)
		updateAccountSettingsOptions.SetRestrictCreatePlatformApikey(restrict_create_platform_apikey_str)
		hasChange = true
	}

	if d.HasChange("restrict_user_list_visibility") {
		restrict_user_list_visibility_str := d.Get("restrict_user_list_visibility").(string)
		updateAccountSettingsOptions.SetRestrictUserListVisibility(restrict_user_list_visibility_str)
		hasChange = true
	}

	if d.HasChange("mfa") {
		mfa_str := d.Get("mfa").(string)
		updateAccountSettingsOptions.SetMfa(mfa_str)
		hasChange = true
	}

	if d.HasChange("user_mfa") {
		if _, ok := d.GetOk("user_mfa"); ok {
			var userMfa []iamidentityv1.UserMfa
			for _, v := range d.Get("user_mfa").([]interface{}) {
				value := v.(map[string]interface{})
				userMfaItem, err := ResourceIBMIamAccountSettingsMapToUserMfa(value)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "update", "parse-user_mfa").GetDiag()
				}
				userMfa = append(userMfa, *userMfaItem)
			}
			updateAccountSettingsOptions.SetUserMfa(userMfa)
		}
	}

	if d.HasChange("restrict_user_domains") {
		if _, ok := d.GetOk("restrict_user_domains"); ok {
			var restrictUserDomains []iamidentityv1.AccountSettingsUserDomainRestriction
			for _, v := range d.Get("restrict_user_domains").([]interface{}) {
				value := v.(map[string]interface{})
				restrictUserDomainsItem, err := ResourceIBMIamAccountSettingsMapToAccountSettingsUserDomainRestriction(value)
				if err != nil {
					return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings", "update", "parse-restrict_user_domains").GetDiag()
				}
				restrictUserDomains = append(restrictUserDomains, *restrictUserDomainsItem)
			}
			updateAccountSettingsOptions.SetRestrictUserDomains(restrictUserDomains)
		}

		hasChange = true
	}

	if d.HasChange("session_expiration_in_seconds") {
		session_expiration_in_seconds_str := d.Get("session_expiration_in_seconds").(string)
		updateAccountSettingsOptions.SetSessionExpirationInSeconds(session_expiration_in_seconds_str)
		hasChange = true
	}

	if d.HasChange("session_invalidation_in_seconds") {
		session_invalidation_in_seconds_str := d.Get("session_invalidation_in_seconds").(string)
		updateAccountSettingsOptions.SetSessionInvalidationInSeconds(session_invalidation_in_seconds_str)
		hasChange = true
	}

	if d.HasChange("max_sessions_per_identity") {
		max_sessions_per_identity_str := d.Get("max_sessions_per_identity").(string)
		updateAccountSettingsOptions.SetMaxSessionsPerIdentity(max_sessions_per_identity_str)
		hasChange = true
	}
	if d.HasChange("system_access_token_expiration_in_seconds") {
		updateAccountSettingsOptions.SetSystemAccessTokenExpirationInSeconds(d.Get("system_access_token_expiration_in_seconds").(string))
		hasChange = true
	}
	if d.HasChange("system_refresh_token_expiration_in_seconds") {
		updateAccountSettingsOptions.SetSystemRefreshTokenExpirationInSeconds(d.Get("system_refresh_token_expiration_in_seconds").(string))
		hasChange = true
	}

	if hasChange {
		_, _, err = iamIdentityClient.UpdateAccountSettingsWithContext(context, updateAccountSettingsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateAccountSettingsWithContext failed: %s", err.Error()), "ibm_iam_account_settings", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIamAccountSettingsRead(context, d, meta)
}

func resourceIBMIamAccountSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// DELETE NOT SUPPORTED
	d.SetId("")

	return nil
}

func ResourceIBMIamAccountSettingsMapToAccountSettingsUserDomainRestriction(modelMap map[string]interface{}) (*iamidentityv1.AccountSettingsUserDomainRestriction, error) {
	model := &iamidentityv1.AccountSettingsUserDomainRestriction{}
	model.RealmID = core.StringPtr(modelMap["realm_id"].(string))
	if modelMap["invitation_email_allow_patterns"] != nil {
		invitationEmailAllowPatterns := []string{}
		for _, invitationEmailAllowPatternsItem := range modelMap["invitation_email_allow_patterns"].([]interface{}) {
			invitationEmailAllowPatterns = append(invitationEmailAllowPatterns, invitationEmailAllowPatternsItem.(string))
		}
		model.InvitationEmailAllowPatterns = invitationEmailAllowPatterns
	}
	if modelMap["restrict_invitation"] != nil {
		model.RestrictInvitation = core.BoolPtr(modelMap["restrict_invitation"].(bool))
	}
	return model, nil
}

func ResourceIBMIamAccountSettingsMapToUserMfa(modelMap map[string]interface{}) (*iamidentityv1.UserMfa, error) {
	model := &iamidentityv1.UserMfa{}
	if modelMap["iam_id"] != nil && modelMap["iam_id"].(string) != "" {
		model.IamID = core.StringPtr(modelMap["iam_id"].(string))
	}
	if modelMap["mfa"] != nil && modelMap["mfa"].(string) != "" {
		model.Mfa = core.StringPtr(modelMap["mfa"].(string))
	}
	return model, nil
}
