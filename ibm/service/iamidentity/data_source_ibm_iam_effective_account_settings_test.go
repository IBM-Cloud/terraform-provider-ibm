// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.93.0-c40121e6-20240729-182103
 */

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamEffectiveAccountSettingsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamEffectiveAccountSettingsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "effective.#"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_effective_account_settings.iam_effective_account_settings_instance", "account.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamEffectiveAccountSettingsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_effective_account_settings" "iam_effective_account_settings_instance" {
			account_id = "account_id"
			include_history = true
			resolve_user_mfa = true
		}
	`)
}

func TestDataSourceIBMIamEffectiveAccountSettingsResponseContextToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["transaction_id"] = "testString"
		model["operation"] = "testString"
		model["user_agent"] = "testString"
		model["url"] = "testString"
		model["instance_id"] = "testString"
		model["thread_id"] = "testString"
		model["host"] = "testString"
		model["start_time"] = "testString"
		model["end_time"] = "testString"
		model["elapsed_time"] = "testString"
		model["cluster_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.ResponseContext)
	model.TransactionID = core.StringPtr("testString")
	model.Operation = core.StringPtr("testString")
	model.UserAgent = core.StringPtr("testString")
	model.URL = core.StringPtr("testString")
	model.InstanceID = core.StringPtr("testString")
	model.ThreadID = core.StringPtr("testString")
	model.Host = core.StringPtr("testString")
	model.StartTime = core.StringPtr("testString")
	model.EndTime = core.StringPtr("testString")
	model.ElapsedTime = core.StringPtr("testString")
	model.ClusterName = core.StringPtr("testString")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsResponseContextToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsAccountSettingsEffectiveSectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		effectiveAccountSettingsUserMfaModel := make(map[string]interface{})
		effectiveAccountSettingsUserMfaModel["iam_id"] = "testString"
		effectiveAccountSettingsUserMfaModel["mfa"] = "NONE"
		effectiveAccountSettingsUserMfaModel["name"] = "testString"
		effectiveAccountSettingsUserMfaModel["user_name"] = "testString"
		effectiveAccountSettingsUserMfaModel["email"] = "testString"
		effectiveAccountSettingsUserMfaModel["description"] = "testString"

		model := make(map[string]interface{})
		model["restrict_create_service_id"] = "NOT_SET"
		model["restrict_create_platform_apikey"] = "NOT_SET"
		model["allowed_ip_addresses"] = "testString"
		model["mfa"] = "NONE"
		model["user_mfa"] = []map[string]interface{}{effectiveAccountSettingsUserMfaModel}
		model["session_expiration_in_seconds"] = "86400"
		model["session_invalidation_in_seconds"] = "7200"
		model["max_sessions_per_identity"] = "testString"
		model["system_access_token_expiration_in_seconds"] = "3600"
		model["system_refresh_token_expiration_in_seconds"] = "259200"

		assert.Equal(t, result, model)
	}

	effectiveAccountSettingsUserMfaModel := new(iamidentityv1.EffectiveAccountSettingsUserMfa)
	effectiveAccountSettingsUserMfaModel.IamID = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
	effectiveAccountSettingsUserMfaModel.Name = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.UserName = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Email = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Description = core.StringPtr("testString")

	model := new(iamidentityv1.AccountSettingsEffectiveSection)
	model.RestrictCreateServiceID = core.StringPtr("NOT_SET")
	model.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
	model.AllowedIPAddresses = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.UserMfa = []iamidentityv1.EffectiveAccountSettingsUserMfa{*effectiveAccountSettingsUserMfaModel}
	model.SessionExpirationInSeconds = core.StringPtr("86400")
	model.SessionInvalidationInSeconds = core.StringPtr("7200")
	model.MaxSessionsPerIdentity = core.StringPtr("testString")
	model.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
	model.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsAccountSettingsEffectiveSectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsEffectiveAccountSettingsUserMfaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["iam_id"] = "testString"
		model["mfa"] = "NONE"
		model["name"] = "testString"
		model["user_name"] = "testString"
		model["email"] = "testString"
		model["description"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.EffectiveAccountSettingsUserMfa)
	model.IamID = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.Name = core.StringPtr("testString")
	model.UserName = core.StringPtr("testString")
	model.Email = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsEffectiveAccountSettingsUserMfaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsAccountSettingsAccountSectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		effectiveAccountSettingsUserMfaModel := make(map[string]interface{})
		effectiveAccountSettingsUserMfaModel["iam_id"] = "testString"
		effectiveAccountSettingsUserMfaModel["mfa"] = "NONE"
		effectiveAccountSettingsUserMfaModel["name"] = "testString"
		effectiveAccountSettingsUserMfaModel["user_name"] = "testString"
		effectiveAccountSettingsUserMfaModel["email"] = "testString"
		effectiveAccountSettingsUserMfaModel["description"] = "testString"

		enityHistoryRecordModel := make(map[string]interface{})
		enityHistoryRecordModel["timestamp"] = "testString"
		enityHistoryRecordModel["iam_id"] = "testString"
		enityHistoryRecordModel["iam_id_account"] = "testString"
		enityHistoryRecordModel["action"] = "testString"
		enityHistoryRecordModel["params"] = []string{"testString"}
		enityHistoryRecordModel["message"] = "testString"

		model := make(map[string]interface{})
		model["account_id"] = "testString"
		model["restrict_create_service_id"] = "NOT_SET"
		model["restrict_create_platform_apikey"] = "NOT_SET"
		model["allowed_ip_addresses"] = "testString"
		model["mfa"] = "NONE"
		model["user_mfa"] = []map[string]interface{}{effectiveAccountSettingsUserMfaModel}
		model["history"] = []map[string]interface{}{enityHistoryRecordModel}
		model["session_expiration_in_seconds"] = "86400"
		model["session_invalidation_in_seconds"] = "7200"
		model["max_sessions_per_identity"] = "testString"
		model["system_access_token_expiration_in_seconds"] = "3600"
		model["system_refresh_token_expiration_in_seconds"] = "259200"

		assert.Equal(t, result, model)
	}

	effectiveAccountSettingsUserMfaModel := new(iamidentityv1.EffectiveAccountSettingsUserMfa)
	effectiveAccountSettingsUserMfaModel.IamID = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
	effectiveAccountSettingsUserMfaModel.Name = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.UserName = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Email = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Description = core.StringPtr("testString")

	enityHistoryRecordModel := new(iamidentityv1.EnityHistoryRecord)
	enityHistoryRecordModel.Timestamp = core.StringPtr("testString")
	enityHistoryRecordModel.IamID = core.StringPtr("testString")
	enityHistoryRecordModel.IamIDAccount = core.StringPtr("testString")
	enityHistoryRecordModel.Action = core.StringPtr("testString")
	enityHistoryRecordModel.Params = []string{"testString"}
	enityHistoryRecordModel.Message = core.StringPtr("testString")

	model := new(iamidentityv1.AccountSettingsAccountSection)
	model.AccountID = core.StringPtr("testString")
	model.RestrictCreateServiceID = core.StringPtr("NOT_SET")
	model.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
	model.AllowedIPAddresses = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.UserMfa = []iamidentityv1.EffectiveAccountSettingsUserMfa{*effectiveAccountSettingsUserMfaModel}
	model.History = []iamidentityv1.EnityHistoryRecord{*enityHistoryRecordModel}
	model.SessionExpirationInSeconds = core.StringPtr("86400")
	model.SessionInvalidationInSeconds = core.StringPtr("7200")
	model.MaxSessionsPerIdentity = core.StringPtr("testString")
	model.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
	model.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAccountSectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsEnityHistoryRecordToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["timestamp"] = "testString"
		model["iam_id"] = "testString"
		model["iam_id_account"] = "testString"
		model["action"] = "testString"
		model["params"] = []string{"testString"}
		model["message"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.EnityHistoryRecord)
	model.Timestamp = core.StringPtr("testString")
	model.IamID = core.StringPtr("testString")
	model.IamIDAccount = core.StringPtr("testString")
	model.Action = core.StringPtr("testString")
	model.Params = []string{"testString"}
	model.Message = core.StringPtr("testString")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsEnityHistoryRecordToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamEffectiveAccountSettingsAccountSettingsAssignedTemplatesSectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		effectiveAccountSettingsUserMfaModel := make(map[string]interface{})
		effectiveAccountSettingsUserMfaModel["iam_id"] = "testString"
		effectiveAccountSettingsUserMfaModel["mfa"] = "NONE"
		effectiveAccountSettingsUserMfaModel["name"] = "testString"
		effectiveAccountSettingsUserMfaModel["user_name"] = "testString"
		effectiveAccountSettingsUserMfaModel["email"] = "testString"
		effectiveAccountSettingsUserMfaModel["description"] = "testString"

		model := make(map[string]interface{})
		model["template_id"] = "testString"
		model["template_version"] = int(26)
		model["template_name"] = "testString"
		model["restrict_create_service_id"] = "NOT_SET"
		model["restrict_create_platform_apikey"] = "NOT_SET"
		model["allowed_ip_addresses"] = "testString"
		model["mfa"] = "NONE"
		model["user_mfa"] = []map[string]interface{}{effectiveAccountSettingsUserMfaModel}
		model["session_expiration_in_seconds"] = "86400"
		model["session_invalidation_in_seconds"] = "7200"
		model["max_sessions_per_identity"] = "testString"
		model["system_access_token_expiration_in_seconds"] = "3600"
		model["system_refresh_token_expiration_in_seconds"] = "259200"

		assert.Equal(t, result, model)
	}

	effectiveAccountSettingsUserMfaModel := new(iamidentityv1.EffectiveAccountSettingsUserMfa)
	effectiveAccountSettingsUserMfaModel.IamID = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Mfa = core.StringPtr("NONE")
	effectiveAccountSettingsUserMfaModel.Name = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.UserName = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Email = core.StringPtr("testString")
	effectiveAccountSettingsUserMfaModel.Description = core.StringPtr("testString")

	model := new(iamidentityv1.AccountSettingsAssignedTemplatesSection)
	model.TemplateID = core.StringPtr("testString")
	model.TemplateVersion = core.Int64Ptr(int64(26))
	model.TemplateName = core.StringPtr("testString")
	model.RestrictCreateServiceID = core.StringPtr("NOT_SET")
	model.RestrictCreatePlatformApikey = core.StringPtr("NOT_SET")
	model.AllowedIPAddresses = core.StringPtr("testString")
	model.Mfa = core.StringPtr("NONE")
	model.UserMfa = []iamidentityv1.EffectiveAccountSettingsUserMfa{*effectiveAccountSettingsUserMfaModel}
	model.SessionExpirationInSeconds = core.StringPtr("86400")
	model.SessionInvalidationInSeconds = core.StringPtr("7200")
	model.MaxSessionsPerIdentity = core.StringPtr("testString")
	model.SystemAccessTokenExpirationInSeconds = core.StringPtr("3600")
	model.SystemRefreshTokenExpirationInSeconds = core.StringPtr("259200")

	result, err := iamidentity.DataSourceIBMIamEffectiveAccountSettingsAccountSettingsAssignedTemplatesSectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
