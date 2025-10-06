// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package iamidentity_test

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/stretchr/testify/assert"
)

func TestEnityHistoryRecordToMap(t *testing.T) {
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

	result, err := iamidentity.EnityHistoryRecordToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestAccountSettingsUserDomainRestrictionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["realm_id"] = "IBMid"
		model["invitation_email_allow_patterns"] = []string{"*.*@company.com"}
		model["restrict_invitation"] = true

		assert.Equal(t, result, model)
	}

	model := new(iamidentityv1.AccountSettingsUserDomainRestriction)
	model.RealmID = core.StringPtr("IBMid")
	model.InvitationEmailAllowPatterns = []string{"*.*@company.com"}
	model.RestrictInvitation = core.BoolPtr(true)

	result, err := iamidentity.AccountSettingsUserDomainRestrictionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
