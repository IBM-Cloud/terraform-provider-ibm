// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
* IBM OpenAPI Terraform Generator Version: 3.95.2-120e65bc-20240924-152329
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudQueueManagerStatusDataSourceBasic(t *testing.T) {
	t.Parallel()
	service_instance_guid := acc.MqcloudDeploymentID
	queue_manager_id := acc.MqcloudQueueManagerID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudQueueManagerStatusDataSourceConfigBasic(service_instance_guid, queue_manager_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager_status.mqcloud_queue_manager_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager_status.mqcloud_queue_manager_status_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager_status.mqcloud_queue_manager_status_instance", "queue_manager_id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager_status.mqcloud_queue_manager_status_instance", "status"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudQueueManagerStatusDataSourceConfigBasic(service_instance_guid string, queue_manager_id string) string {
	return fmt.Sprintf(`
		data "ibm_mqcloud_queue_manager_status" "mqcloud_queue_manager_status_instance" {
			service_instance_guid = "%s"
			queue_manager_id = "%s"
		}
	`, service_instance_guid, queue_manager_id)
}
