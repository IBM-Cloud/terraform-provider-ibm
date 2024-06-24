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
)

func TestAccIbmMqcloudQueueManagerOptionsDataSourceBasic(t *testing.T) {
	t.Parallel()
	service_instance_guid := acc.MqcloudInstanceID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudQueueManagerOptionsDataSourceConfigBasic(service_instance_guid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager_options.mqcloud_queue_manager_options_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager_options.mqcloud_queue_manager_options_instance", "service_instance_guid"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudQueueManagerOptionsDataSourceConfigBasic(service_instance_guid string) string {
	return fmt.Sprintf(`
		data "ibm_mqcloud_queue_manager_options" "mqcloud_queue_manager_options_instance" {
			service_instance_guid = "%s"
		}
	`, service_instance_guid)
}
