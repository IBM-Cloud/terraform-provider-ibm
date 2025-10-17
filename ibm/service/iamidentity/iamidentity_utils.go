package iamidentity

import "github.com/IBM/platform-services-go-sdk/iamidentityv1"

func EnityHistoryRecordToMap(model *iamidentityv1.EnityHistoryRecord) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	if model.Timestamp != nil {
		modelMap["timestamp"] = *model.Timestamp
	}
	if model.IamID != nil {
		modelMap["iam_id"] = *model.IamID
	}
	if model.IamIDAccount != nil {
		modelMap["iam_id_account"] = *model.IamIDAccount
	}
	if model.Action != nil {
		modelMap["action"] = *model.Action
	}
	if model.Params != nil {
		modelMap["params"] = model.Params
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}

	return modelMap, nil
}

func AccountSettingsUserMfaResponseToMap(model *iamidentityv1.AccountSettingsUserMfaResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	modelMap["iam_id"] = *model.IamID
	modelMap["mfa"] = *model.Mfa

	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.UserName != nil {
		modelMap["user_name"] = *model.UserName
	}
	if model.Email != nil {
		modelMap["email"] = *model.Email
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}

	return modelMap, nil
}

func AccountSettingsUserDomainRestrictionToMap(model *iamidentityv1.AccountSettingsUserDomainRestriction) (map[string]interface{}, error) {

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
