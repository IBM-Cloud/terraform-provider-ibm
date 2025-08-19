// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.106.0-09823488-20250707-071701
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	db2saas "github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/db2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDb2SaasUsersDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2SaasUsersDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_users.db2_saas_users_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_users.db2_saas_users_instance", "x_deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_users.db2_saas_users_instance", "count"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_saas_users.db2_saas_users_instance", "resources.#"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2SaasUsersDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_db2_saas_users" "db2_saas_users_instance" {
			x-deployment-id = "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
		}
	`)
}

func TestDataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		successGetUserInfoResourcesItemAuthenticationModel := make(map[string]interface{})
		successGetUserInfoResourcesItemAuthenticationModel["method"] = "internal"
		successGetUserInfoResourcesItemAuthenticationModel["policy_id"] = "Default"

		model := make(map[string]interface{})
		model["dv_role"] = "test-role"
		model["metadata"] = map[string]interface{}{"anyKey": "anyValue"}
		model["formated_ibmid"] = "test-formated-ibm-id"
		model["role"] = "bluadmin"
		model["iamid"] = "test-iam-id"
		model["permitted_actions"] = []string{"testString"}
		model["all_clean"] = false
		model["password"] = "nd!@aegr63@989hcRFTcdcs63"
		model["iam"] = false
		model["name"] = "admin"
		model["ibmid"] = "test-ibm-id"
		model["id"] = "admin"
		model["locked"] = "no"
		model["init_error_msg"] = "testString"
		model["email"] = "user@host.org"
		model["authentication"] = []map[string]interface{}{successGetUserInfoResourcesItemAuthenticationModel}

		assert.Equal(t, result, model)
	}

	successGetUserInfoResourcesItemAuthenticationModel := new(db2saasv1.SuccessGetUserInfoResourcesItemAuthentication)
	successGetUserInfoResourcesItemAuthenticationModel.Method = core.StringPtr("internal")
	successGetUserInfoResourcesItemAuthenticationModel.PolicyID = core.StringPtr("Default")

	model := new(db2saasv1.SuccessGetUserInfoResourcesItem)
	model.DvRole = core.StringPtr("test-role")
	model.Metadata = map[string]interface{}{"anyKey": "anyValue"}
	model.FormatedIbmid = core.StringPtr("test-formated-ibm-id")
	model.Role = core.StringPtr("bluadmin")
	model.Iamid = core.StringPtr("test-iam-id")
	model.PermittedActions = []string{"testString"}
	model.AllClean = core.BoolPtr(false)
	model.Password = core.StringPtr("nd!@aegr63@989hcRFTcdcs63")
	model.Iam = core.BoolPtr(false)
	model.Name = core.StringPtr("admin")
	model.Ibmid = core.StringPtr("test-ibm-id")
	model.ID = core.StringPtr("admin")
	model.Locked = core.StringPtr("no")
	model.InitErrorMsg = core.StringPtr("testString")
	model.Email = core.StringPtr("user@host.org")
	model.Authentication = successGetUserInfoResourcesItemAuthenticationModel

	result, err := db2saas.DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemAuthenticationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["method"] = "internal"
		model["policy_id"] = "Default"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessGetUserInfoResourcesItemAuthentication)
	model.Method = core.StringPtr("internal")
	model.PolicyID = core.StringPtr("Default")

	result, err := db2saas.DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemAuthenticationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
