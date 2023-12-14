// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudQueueManagerDataSourceBasic(t *testing.T) {
	t.Parallel()
	queueManagerDetailsServiceInstanceGuid := acc.MqcloudInstanceID
	queueManagerDetailsName := "queue_manager_ds_basic"
	queueManagerDetailsLocation := "ibmcloud_eu_de"
	queueManagerDetailsSize := "small"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudQueueManagerDataSourceConfigBasic(queueManagerDetailsServiceInstanceGuid, queueManagerDetailsName, queueManagerDetailsLocation, queueManagerDetailsSize),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.#"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.name", queueManagerDetailsName),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.location", queueManagerDetailsLocation),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.size", queueManagerDetailsSize),
				),
			},
		},
	})
}

func TestAccIbmMqcloudQueueManagerDataSourceAllArgs(t *testing.T) {
	t.Parallel()
	queueManagerDetailsServiceInstanceGuid := acc.MqcloudInstanceID
	queueManagerDetailsName := "queue_manager_ds_allargs"
	queueManagerDetailsDisplayName := "queue_manager_ds_allargs"
	queueManagerDetailsLocation := "ibmcloud_eu_de"
	queueManagerDetailsSize := "small"
	queueManagerDetailsVersion := "9.3.3_3"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudQueueManagerDataSourceConfig(queueManagerDetailsServiceInstanceGuid, queueManagerDetailsName, queueManagerDetailsDisplayName, queueManagerDetailsLocation, queueManagerDetailsSize, queueManagerDetailsVersion),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.#"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.id"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.name", queueManagerDetailsName),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.display_name", queueManagerDetailsDisplayName),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.location", queueManagerDetailsLocation),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.size", queueManagerDetailsSize),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.status_uri"),
					resource.TestCheckResourceAttr("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.version", queueManagerDetailsVersion),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.web_console_url"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.rest_api_endpoint_url"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.administrator_api_endpoint_url"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.connection_info_uri"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.date_created"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.upgrade_available"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.available_upgrade_versions_uri"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "queue_managers.0.href"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudQueueManagerDataSourceConfigBasic(queueManagerDetailsServiceInstanceGuid string, queueManagerDetailsName string, queueManagerDetailsLocation string, queueManagerDetailsSize string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
			service_instance_guid = "%s"
			name = "%s"
			location = "%s"
			size = "%s"
		}

		data "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
			service_instance_guid = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance.service_instance_guid
			name = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance.name
		}
	`, queueManagerDetailsServiceInstanceGuid, queueManagerDetailsName, queueManagerDetailsLocation, queueManagerDetailsSize)
}

func testAccCheckIbmMqcloudQueueManagerDataSourceConfig(queueManagerDetailsServiceInstanceGuid string, queueManagerDetailsName string, queueManagerDetailsDisplayName string, queueManagerDetailsLocation string, queueManagerDetailsSize string, queueManagerDetailsVersion string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
			service_instance_guid = "%s"
			name = "%s"
			display_name = "%s"
			location = "%s"
			size = "%s"
			version = "%s"
		}

		data "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
			service_instance_guid = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance.service_instance_guid
			name = ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance.name
		}
	`, queueManagerDetailsServiceInstanceGuid, queueManagerDetailsName, queueManagerDetailsDisplayName, queueManagerDetailsLocation, queueManagerDetailsSize, queueManagerDetailsVersion)
}
