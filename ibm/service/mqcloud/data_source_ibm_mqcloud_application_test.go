// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
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
