// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
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
	. "github.com/Mavrickk3/terraform-provider-ibm/ibm/unittest"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIamTrustedProfilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance", "account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profiles.iam_trusted_profiles_instance", "profiles.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profiles" "iam_trusted_profiles_instance" {
			account_id = "%s"
			include_history = true
		}
	`, acc.AccountId)
}

func TestDataSourceIBMIamTrustedProfilesTrustedProfileToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		enityHistoryRecordModel := make(map[string]interface{})
		enityHistoryRecordModel["timestamp"] = "testString"
		enityHistoryRecordModel["iam_id"] = "testString"
		enityHistoryRecordModel["iam_id_account"] = "testString"
		enityHistoryRecordModel["action"] = "testString"
		enityHistoryRecordModel["params"] = []string{"testString"}
		enityHistoryRecordModel["message"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["entity_tag"] = "testString"
		model["crn"] = "testString"
		model["name"] = "testString"
		model["description"] = "testString"
		model["created_at"] = "2019-01-01T12:00:00.000Z"
		model["modified_at"] = "2019-01-01T12:00:00.000Z"
		model["iam_id"] = "testString"
		model["account_id"] = "testString"
		model["template_id"] = "testString"
		model["assignment_id"] = "testString"
		model["ims_account_id"] = int(26)
		model["ims_user_id"] = int(26)
		model["history"] = []map[string]interface{}{enityHistoryRecordModel}

		assert.Equal(t, result, model)
	}

	enityHistoryRecordModel := new(iamidentityv1.EnityHistoryRecord)
	enityHistoryRecordModel.Timestamp = core.StringPtr("testString")
	enityHistoryRecordModel.IamID = core.StringPtr("testString")
	enityHistoryRecordModel.IamIDAccount = core.StringPtr("testString")
	enityHistoryRecordModel.Action = core.StringPtr("testString")
	enityHistoryRecordModel.Params = []string{"testString"}
	enityHistoryRecordModel.Message = core.StringPtr("testString")

	model := new(iamidentityv1.TrustedProfile)
	model.ID = core.StringPtr("testString")
	model.EntityTag = core.StringPtr("testString")
	model.CRN = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.CreatedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.ModifiedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.IamID = core.StringPtr("testString")
	model.AccountID = core.StringPtr("testString")
	model.TemplateID = core.StringPtr("testString")
	model.AssignmentID = core.StringPtr("testString")
	model.ImsAccountID = core.Int64Ptr(int64(26))
	model.ImsUserID = core.Int64Ptr(int64(26))
	model.History = []iamidentityv1.EnityHistoryRecord{*enityHistoryRecordModel}

	result, err := iamidentity.DataSourceIBMIamTrustedProfilesTrustedProfileToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMIamTrustedProfilesEnityHistoryRecordToMap(t *testing.T) {
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

	result, err := iamidentity.DataSourceIBMIamTrustedProfilesEnityHistoryRecordToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
