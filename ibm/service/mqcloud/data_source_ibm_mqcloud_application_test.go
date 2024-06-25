// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.90.0-5aad763d-20240506-203857
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/mqcloud"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmMqcloudApplicationDataSourceBasic(t *testing.T) {
	t.Parallel()
	applicationDetailsServiceInstanceGuid := acc.MqcloudInstanceID
	applicationDetailsName := "appdsbasic"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudApplicationDataSourceConfigBasic(applicationDetailsServiceInstanceGuid, applicationDetailsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_application.mqcloud_application_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_application.mqcloud_application_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_application.mqcloud_application_instance", "applications.#"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_application.mqcloud_application_instance", "applications.0.name", applicationDetailsName),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudApplicationDataSourceConfigBasic(applicationDetailsServiceInstanceGuid string, applicationDetailsName string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_application" "mqcloud_application_instance" {
			service_instance_guid = "%s"
			name = "%s"
		}

		data "ibm_mqcloud_application" "mqcloud_application_instance" {
			service_instance_guid = ibm_mqcloud_application.mqcloud_application_instance.service_instance_guid
			name = ibm_mqcloud_application.mqcloud_application_instance.name
		}
	`, applicationDetailsServiceInstanceGuid, applicationDetailsName)
}

func TestDataSourceIbmMqcloudApplicationApplicationDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["create_api_key_uri"] = "testString"
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(mqcloudv1.ApplicationDetails)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.CreateApiKeyURI = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")

	result, err := mqcloud.DataSourceIbmMqcloudApplicationApplicationDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
