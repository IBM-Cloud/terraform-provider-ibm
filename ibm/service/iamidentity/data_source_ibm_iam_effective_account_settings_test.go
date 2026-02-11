// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamEffectiveAccountSettingsDataSourceBasic(t *testing.T) {
	accountId := acc.IAMAccountId
	includeHistory := "false"
	resolveUserMfa := "true"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acc.TestAccPreCheck(t)
			acc.TestAccPreCheckCbr(t)
		},
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIamEffectiveAccountSettingsDataSourceConfigBasic(accountId, includeHistory, resolveUserMfa),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "include_history"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "resolve_user_mfa"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "effective.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "account.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamEffectiveAccountSettingsDataSourceConfigBasic(accountId string, includeHistory string, resolveUserMfa string) string {
	return fmt.Sprintf(`
		data "ibm_iam_effective_account_settings" "iam_effective_account_settings_instance" {
			account_id = "%s"
			include_history = %s
			resolve_user_mfa = %s
		}
	`, accountId, includeHistory, resolveUserMfa)
}

func TestDataSourceIBMIamEffectiveAccountSettingsAccountSettingsEffectiveSectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		accountSettingsUserMfaResponseModel := make(map[string]interface{})
		accountSettingsUserMfaResponseModel["iam_id"] = "testString"
		accountSettingsUserMfaResponseModel["mfa"] = "NONE"
		accountSettingsUserMfaResponseModel["name"] = "testString"
		accountSettingsUserMfaResponseModel["user_name"] = "testString"
		accountSettingsUserMfaResponseModel["email"] = "testString"
		accountSettingsUserMfaResponseModel["description"] = "testString"

		model := make(map[string]interface{})
		model["restrict_create_service_id"] = "NOT_SET"
		model["restrict_create_platform_apikey"] = "NOT_SET"
		model["restrict_user_list_visibility"] = "NOT_RESTRICTED"
		model["allowed_ip_addresses"] = "testString"
		model["mfa"] = "NONE"
		model["user_mfa"] = []map[string]interface{}{accountSettingsUserMfaResponseModel}
		model["session_expiration_in_seconds"] = "86400"
		model["session_invalidation_in_seconds"] = "7200"
		model["max_sessions_per_identity"] = "testString"
		model["system_access_token_expiration_in_seconds"] = "3600"
		model["system_refresh_token_expiration_in_seconds"] = "259200"

		assert.Equal(t, result, model)
	}

	accountSettingsUserMfaResponseModel := new(iamidentityv1.AccountSettingsUserMfaResponse)
	accountSettingsUserMfaResponseModel.IamID = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Mfa = core.StringPtr("NONE")
	accountSettingsUserMfaResponseModel.Name = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.UserName = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Email = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Description = core.StringPtr("testString")

	model := new(iamidentityv1.AccountSettingsEffectiveSection)
	model.RestrictCreateServiceID = core.StringPtr("NOT_SET")
	model.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
	model.RestrictUserListVisibility = core.StringPtr("NOT_RESTRICTED")
	model.AllowedIPAddresses = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.UserMfa = []iamidentityv1.AccountSettingsUserMfaResponse{*accountSettingsUserMfaResponseModel}
	model.SessionExpirationInSeconds = core.StringPtr("86400")
	model.SessionInvalidationInSeconds = core.StringPtr("7200")
	model.MaxSessionsPerIdentity = core.StringPtr("testString")
	model.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
	model.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsAccountSettingsEffectiveSectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsAccountSettingsResponseToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		enityHistoryRecordModel := make(map[string]interface{})
		enityHistoryRecordModel["timestamp"] = "testString"
		enityHistoryRecordModel["iam_id"] = "testString"
		enityHistoryRecordModel["iam_id_account"] = "testString"
		enityHistoryRecordModel["action"] = "testString"
		enityHistoryRecordModel["params"] = []string{"testString"}
		enityHistoryRecordModel["message"] = "testString"

		accountSettingsUserDomainRestrictionModel := make(map[string]interface{})
		accountSettingsUserDomainRestrictionModel["realm_id"] = "IBMid"
		accountSettingsUserDomainRestrictionModel["invitation_email_allow_patterns"] = []string{"*.*@company.com"}
		accountSettingsUserDomainRestrictionModel["restrict_invitation"] = true

		accountSettingsUserMfaResponseModel := make(map[string]interface{})
		accountSettingsUserMfaResponseModel["iam_id"] = "testString"
		accountSettingsUserMfaResponseModel["mfa"] = "NONE"
		accountSettingsUserMfaResponseModel["name"] = "testString"
		accountSettingsUserMfaResponseModel["user_name"] = "testString"
		accountSettingsUserMfaResponseModel["email"] = "testString"
		accountSettingsUserMfaResponseModel["description"] = "testString"

		model := make(map[string]interface{})
		model["entity_tag"] = "testString"
		model["history"] = []map[string]interface{}{enityHistoryRecordModel}
		model["restrict_create_service_id"] = "NOT_SET"
		model["restrict_create_platform_apikey"] = "NOT_SET"
		model["restrict_user_list_visibility"] = "NOT_RESTRICTED"
		model["restrict_user_domains"] = []map[string]interface{}{accountSettingsUserDomainRestrictionModel}
		model["allowed_ip_addresses"] = "testString"
		model["mfa"] = "NONE"
		model["session_expiration_in_seconds"] = "86400"
		model["session_invalidation_in_seconds"] = "7200"
		model["max_sessions_per_identity"] = "testString"
		model["system_access_token_expiration_in_seconds"] = "3600"
		model["system_refresh_token_expiration_in_seconds"] = "259200"
		model["user_mfa"] = []map[string]interface{}{accountSettingsUserMfaResponseModel}

		assert.Equal(t, result, model)
	}

	enityHistoryRecordModel := new(iamidentityv1.EnityHistoryRecord)
	enityHistoryRecordModel.Timestamp = core.StringPtr("testString")
	enityHistoryRecordModel.IamID = core.StringPtr("testString")
	enityHistoryRecordModel.IamIDAccount = core.StringPtr("testString")
	enityHistoryRecordModel.Action = core.StringPtr("testString")
	enityHistoryRecordModel.Params = []string{"testString"}
	enityHistoryRecordModel.Message = core.StringPtr("testString")

	accountSettingsUserDomainRestrictionModel := new(iamidentityv1.AccountSettingsUserDomainRestriction)
	accountSettingsUserDomainRestrictionModel.RealmID = core.StringPtr("IBMid")
	accountSettingsUserDomainRestrictionModel.InvitationEmailAllowPatterns = []string{"*.*@company.com"}
	accountSettingsUserDomainRestrictionModel.RestrictInvitation = core.BoolPtr(true)

	accountSettingsUserMfaResponseModel := new(iamidentityv1.AccountSettingsUserMfaResponse)
	accountSettingsUserMfaResponseModel.IamID = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Mfa = core.StringPtr("NONE")
	accountSettingsUserMfaResponseModel.Name = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.UserName = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Email = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Description = core.StringPtr("testString")

	model := new(iamidentityv1.AccountSettingsResponse)
	model.AccountID = core.StringPtr("testString")
	model.EntityTag = core.StringPtr("testString")
	model.History = []iamidentityv1.EnityHistoryRecord{*enityHistoryRecordModel}
	model.RestrictCreateServiceID = core.StringPtr("NOT_SET")
	model.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
	model.RestrictUserListVisibility = core.StringPtr("NOT_RESTRICTED")
	model.RestrictUserDomains = []iamidentityv1.AccountSettingsUserDomainRestriction{*accountSettingsUserDomainRestrictionModel}
	model.AllowedIPAddresses = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.SessionExpirationInSeconds = core.StringPtr("86400")
	model.SessionInvalidationInSeconds = core.StringPtr("7200")
	model.MaxSessionsPerIdentity = core.StringPtr("testString")
	model.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
	model.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
	model.UserMfa = []iamidentityv1.AccountSettingsUserMfaResponse{*accountSettingsUserMfaResponseModel}

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAccountSectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsAccountSettingsAssignedTemplatesSectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		accountSettingsUserDomainRestrictionModel := make(map[string]interface{})
		accountSettingsUserDomainRestrictionModel["realm_id"] = "IBMid"
		accountSettingsUserDomainRestrictionModel["invitation_email_allow_patterns"] = []string{"*.*@company.com"}
		accountSettingsUserDomainRestrictionModel["restrict_invitation"] = true

		accountSettingsUserMfaResponseModel := make(map[string]interface{})
		accountSettingsUserMfaResponseModel["iam_id"] = "testString"
		accountSettingsUserMfaResponseModel["mfa"] = "NONE"
		accountSettingsUserMfaResponseModel["name"] = "testString"
		accountSettingsUserMfaResponseModel["user_name"] = "testString"
		accountSettingsUserMfaResponseModel["email"] = "testString"
		accountSettingsUserMfaResponseModel["description"] = "testString"

		model := make(map[string]interface{})
		model["template_id"] = "testString"
		model["template_version"] = int(26)
		model["template_name"] = "testString"
		model["restrict_create_service_id"] = "NOT_SET"
		model["restrict_create_platform_apikey"] = "NOT_SET"
		model["restrict_user_list_visibility"] = "NOT_RESTRICTED"
		//model["restrict_user_domains"] = []map[string]interface{}{accountSettingsUserDomainRestrictionModel}
		model["restrict_user_domains"] = []map[string]interface{}{
			{
				"account_sufficient": true,
				"restrictions": []map[string]interface{}{
					{
						"realm_id":                        "IBMid",
						"invitation_email_allow_patterns": []string{"*.*@company.com"},
						"restrict_invitation":             true,
					},
				},
			},
		}
		model["allowed_ip_addresses"] = "testString"
		model["mfa"] = "NONE"
		model["session_expiration_in_seconds"] = "86400"
		model["session_invalidation_in_seconds"] = "7200"
		model["max_sessions_per_identity"] = "testString"
		model["system_access_token_expiration_in_seconds"] = "3600"
		model["system_refresh_token_expiration_in_seconds"] = "259200"
		model["user_mfa"] = []map[string]interface{}{accountSettingsUserMfaResponseModel}

		assert.Equal(t, result, model)
	}

	accountSettingsUserDomainRestrictionModel := new(iamidentityv1.AccountSettingsUserDomainRestriction)
	accountSettingsUserDomainRestrictionModel.RealmID = core.StringPtr("IBMid")
	accountSettingsUserDomainRestrictionModel.InvitationEmailAllowPatterns = []string{"*.*@company.com"}
	accountSettingsUserDomainRestrictionModel.RestrictInvitation = core.BoolPtr(true)

	accountSettingsUserMfaResponseModel := new(iamidentityv1.AccountSettingsUserMfaResponse)
	accountSettingsUserMfaResponseModel.IamID = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Mfa = core.StringPtr("NONE")
	accountSettingsUserMfaResponseModel.Name = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.UserName = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Email = core.StringPtr("testString")
	accountSettingsUserMfaResponseModel.Description = core.StringPtr("testString")

	model := new(iamidentityv1.AccountSettingsAssignedTemplatesSection)
	model.TemplateID = core.StringPtr("testString")
	model.TemplateVersion = core.Int64Ptr(int64(26))
	model.TemplateName = core.StringPtr("testString")
	model.RestrictCreateServiceID = core.StringPtr("NOT_SET")
	model.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
	model.RestrictUserListVisibility = core.StringPtr("NOT_RESTRICTED")
	restrictUserDomains := &iamidentityv1.AssignedTemplatesAccountSettingsRestrictUserDomains{
		AccountSufficient: core.BoolPtr(true),
		Restrictions: []iamidentityv1.AccountSettingsUserDomainRestriction{
			*accountSettingsUserDomainRestrictionModel,
		},
	}
	model.RestrictUserDomains = restrictUserDomains
	model.AllowedIPAddresses = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.SessionExpirationInSeconds = core.StringPtr("86400")
	model.SessionInvalidationInSeconds = core.StringPtr("7200")
	model.MaxSessionsPerIdentity = core.StringPtr("testString")
	model.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
	model.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")
	model.UserMfa = []iamidentityv1.AccountSettingsUserMfaResponse{*accountSettingsUserMfaResponseModel}

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAssignedTemplatesSectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
