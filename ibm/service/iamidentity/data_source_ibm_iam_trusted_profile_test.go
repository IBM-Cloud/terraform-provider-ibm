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
	"github.com/stretchr/testify/assert"
)

func TestAccIBMIAMTrustedProfileDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamTrustedProfileDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "entity_tag"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_trusted_profile.iam_trusted_profile_instance", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMIamTrustedProfileDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
			profile_id = "%s"
		}
	`, acc.IAMTrustedProfileID)
}

func TestDataSourceIBMIamTrustedProfileEnityHistoryRecordToMap(t *testing.T) {
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

	result, err := iamidentity.DataSourceIBMIamTrustedProfileEnityHistoryRecordToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
