// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/mqcloud"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmMqcloudUserDataSourceBasic(t *testing.T) {
	userDetailsServiceInstanceGuid := acc.MqcloudDeploymentID
	userDetailsName := fmt.Sprintf("tfname%d", acctest.RandIntRange(10, 100))
	userDetailsEmail := fmt.Sprintf("tfemail%d@ibm.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudUserDataSourceConfigBasic(userDetailsServiceInstanceGuid, userDetailsName, userDetailsEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_user.mqcloud_user_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_user.mqcloud_user_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_user.mqcloud_user_instance", "users.#"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_user.mqcloud_user_instance", "users.0.name", userDetailsName),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_user.mqcloud_user_instance", "users.0.email", userDetailsEmail),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudUserDataSourceConfigBasic(userDetailsServiceInstanceGuid string, userDetailsName string, userDetailsEmail string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_user" "mqcloud_user_instance" {
			service_instance_guid = "%s"
			name = "%s"
			email = "%s"
		}

		data "ibm_mqcloud_user" "mqcloud_user_instance" {
			service_instance_guid = ibm_mqcloud_user.mqcloud_user_instance.service_instance_guid
			name = ibm_mqcloud_user.mqcloud_user_instance.name
		}
	`, userDetailsServiceInstanceGuid, userDetailsName, userDetailsEmail)
}

func TestDataSourceIbmMqcloudUserUserDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["email"] = "user@host.org"
		model["iam_service_id"] = "testString"
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(mqcloudv1.UserDetails)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Email = core.StringPtr("user@host.org")
	model.IamServiceID = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")

	result, err := mqcloud.DataSourceIbmMqcloudUserUserDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
