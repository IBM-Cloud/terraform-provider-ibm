// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudUserDataSourceBasic(t *testing.T) {
	t.Parallel()
	userDetailsServiceInstanceGuid := acc.MqcloudInstanceID
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
