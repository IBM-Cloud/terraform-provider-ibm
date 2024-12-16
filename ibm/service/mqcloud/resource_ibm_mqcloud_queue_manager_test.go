// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
)

func TestAccIbmMqcloudQueueManagerBasic(t *testing.T) {
	t.Parallel()
	var conf mqcloudv1.QueueManagerDetails
	serviceInstanceGuid := acc.MqcloudDeploymentID
	name := fmt.Sprintf("tf_queue_manager_basic%d", acctest.RandIntRange(10, 100))
	location := acc.MqCloudQueueManagerLocation
	size := "xsmall"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudQueueManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudQueueManagerConfigBasic(serviceInstanceGuid, name, location, size),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudQueueManagerExists("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "size", size),
				),
			},
		},
	})
}

func TestAccIbmMqcloudQueueManagerAllArgs(t *testing.T) {
	t.Parallel()
	var conf mqcloudv1.QueueManagerDetails
	serviceInstanceGuid := acc.MqcloudDeploymentID
	name := fmt.Sprintf("tf_queue_manager_allargs%d", acctest.RandIntRange(10, 100))
	displayName := name
	location := acc.MqCloudQueueManagerLocation
	size := "xsmall"
	version := acc.MqCloudQueueManagerVersion
	versionUpdate := acc.MqCloudQueueManagerVersionUpdate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckMqcloud(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmMqcloudQueueManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudQueueManagerConfig(serviceInstanceGuid, name, displayName, location, size, version),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmMqcloudQueueManagerExists("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", conf),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "size", size),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "version", version),
				),
			},
			{
				Config: testAccCheckIbmMqcloudQueueManagerConfig(serviceInstanceGuid, name, displayName, location, size, versionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "service_instance_guid", serviceInstanceGuid),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "display_name", displayName),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "location", location),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "size", size),
					resource.TestCheckResourceAttr("ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance", "version", versionUpdate),
				),
			},
			{
				ResourceName:      "ibm_mqcloud_queue_manager.mqcloud_queue_manager_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmMqcloudQueueManagerConfigBasic(serviceInstanceGuid string, name string, location string, size string) string {
	return fmt.Sprintf(`
		resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
			service_instance_guid = "%s"
			name = "%s"
			location = "%s"
			size = "%s"
		}
	`, serviceInstanceGuid, name, location, size)
}

func testAccCheckIbmMqcloudQueueManagerConfig(serviceInstanceGuid string, name string, displayName string, location string, size string, version string) string {
	return fmt.Sprintf(`

		resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
			service_instance_guid = "%s"
			name = "%s"
			display_name = "%s"
			location = "%s"
			size = "%s"
			version = "%s"
		}
	`, serviceInstanceGuid, name, displayName, location, size, version)
}

func testAccCheckIbmMqcloudQueueManagerExists(n string, obj mqcloudv1.QueueManagerDetails) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
		if err != nil {
			return err
		}

		getQueueManagerOptions := &mqcloudv1.GetQueueManagerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getQueueManagerOptions.SetServiceInstanceGuid(parts[0])
		getQueueManagerOptions.SetQueueManagerID(parts[1])

		queueManagerDetails, _, err := mqcloudClient.GetQueueManager(getQueueManagerOptions)
		if err != nil {
			return err
		}

		obj = *queueManagerDetails
		return nil
	}
}

func testAccCheckIbmMqcloudQueueManagerDestroy(s *terraform.State) error {
	mqcloudClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MqcloudV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_mqcloud_queue_manager" {
			continue
		}

		getQueueManagerOptions := &mqcloudv1.GetQueueManagerOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getQueueManagerOptions.SetServiceInstanceGuid(parts[0])
		getQueueManagerOptions.SetQueueManagerID(parts[1])

		// Try to find the key
		_, response, err := mqcloudClient.GetQueueManager(getQueueManagerOptions)

		if err == nil {
			return fmt.Errorf("mqcloud_queue_manager_instance still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for mqcloud_queue_manager_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
