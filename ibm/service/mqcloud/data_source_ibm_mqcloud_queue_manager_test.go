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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/mqcloud"
	. "github.com/Mavrickk3/terraform-provider-ibm/ibm/unittest"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmMqcloudQueueManagerDataSourceBasic(t *testing.T) {
	queueManagerDetailsServiceInstanceGuid := acc.MqcloudDeploymentID
	queueManagerDetailsName := fmt.Sprintf("tf_queue_manager_ds_basic%d", acctest.RandIntRange(10, 100))
	queueManagerDetailsLocation := acc.MqCloudQueueManagerLocation
	queueManagerDetailsSize := "xsmall"

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
	queueManagerDetailsServiceInstanceGuid := acc.MqcloudDeploymentID
	queueManagerDetailsName := fmt.Sprintf("tf_queue_manager_ds_allargs%d", acctest.RandIntRange(10, 100))
	queueManagerDetailsDisplayName := queueManagerDetailsName
	queueManagerDetailsLocation := acc.MqCloudQueueManagerLocation
	queueManagerDetailsSize := "xsmall"
	queueManagerDetailsVersion := acc.MqCloudQueueManagerVersionUpdate

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

func TestDataSourceIbmMqcloudQueueManagerQueueManagerDetailsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["display_name"] = "testString"
		model["location"] = "reserved-eu-de-cluster-f884"
		model["size"] = "small"
		model["status_uri"] = "testString"
		model["version"] = "9.3.2_2"
		model["web_console_url"] = "testString"
		model["rest_api_endpoint_url"] = "testString"
		model["administrator_api_endpoint_url"] = "testString"
		model["connection_info_uri"] = "testString"
		model["date_created"] = "2020-01-13T15:39:35.000Z"
		model["upgrade_available"] = true
		model["available_upgrade_versions_uri"] = "testString"
		model["href"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(mqcloudv1.QueueManagerDetails)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.DisplayName = core.StringPtr("testString")
	model.Location = core.StringPtr("reserved-eu-de-cluster-f884")
	model.Size = core.StringPtr("small")
	model.StatusURI = core.StringPtr("testString")
	model.Version = core.StringPtr("9.3.2_2")
	model.WebConsoleURL = core.StringPtr("testString")
	model.RestApiEndpointURL = core.StringPtr("testString")
	model.AdministratorApiEndpointURL = core.StringPtr("testString")
	model.ConnectionInfoURI = core.StringPtr("testString")
	model.DateCreated = CreateMockDateTime("2020-01-13T15:39:35.000Z")
	model.UpgradeAvailable = core.BoolPtr(true)
	model.AvailableUpgradeVersionsURI = core.StringPtr("testString")
	model.Href = core.StringPtr("testString")

	result, err := mqcloud.DataSourceIbmMqcloudQueueManagerQueueManagerDetailsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
