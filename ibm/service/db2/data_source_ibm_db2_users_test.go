// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/db2"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDb2UsersDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2UsersDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "x_deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "dv_role"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "metadata.%"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "formated_ibmid"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "role"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "iamid"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "permitted_actions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "all_clean"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "password"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "iam"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "ibmid"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "locked"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "init_error_msg"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "email"),
					resource.TestCheckResourceAttrSet("data.ibm_db2_users.db2_users_instance", "authentication.#"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2UsersDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_db2_users" "db2_users_instance" {
			x-deployment-id = "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
		}
	`)
}

func TestDataSourceIbmDb2UsersSuccessGetUserByIDAuthenticationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["method"] = "testString"
		model["policy_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessGetUserByIDAuthentication)
	model.Method = core.StringPtr("testString")
	model.PolicyID = core.StringPtr("testString")

	result, err := db2.DataSourceIbmDb2UsersSuccessGetUserByIDAuthenticationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
