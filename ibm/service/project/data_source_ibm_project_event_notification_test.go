// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectEventNotificationDataSourceBasic(t *testing.T) {
	projectResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	projectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEventNotificationDataSourceConfigBasic(projectResourceGroup, projectLocation, projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "name"),
				),
			},
		},
	})
}

func TestAccIbmProjectEventNotificationDataSourceAllArgs(t *testing.T) {
	projectResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	projectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEventNotificationDataSourceConfig(projectResourceGroup, projectLocation, projectName, projectDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "exclude_configs"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "configs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "configs.0.id"),
					resource.TestCheckResourceAttr("data.ibm_project_event_notification.project_event_notification", "configs.0.name", projectName),
					resource.TestCheckResourceAttr("data.ibm_project_event_notification.project_event_notification", "configs.0.description", projectDescription),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "configs.0.locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "configs.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "metadata.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectEventNotificationDataSourceConfigBasic(projectResourceGroup string, projectLocation string, projectName string) string {
	return fmt.Sprintf(`
		resource "ibm_project_instance" "project_instance_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}

		data "ibm_project_event_notification" "project_event_notification_instance" {
			id = ibm_project_config.project_config_instance.projectConfig_id
			exclude_configs = true
		}
	`, projectResourceGroup, projectLocation, projectName)
}

func testAccCheckIbmProjectEventNotificationDataSourceConfig(projectResourceGroup string, projectLocation string, projectName string, projectDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_project_instance" "project_instance_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
			description = "%s"
			configs {
				id = "id"
				name = "name"
				labels = [ "labels" ]
				description = "description"
				locator_id = "locator_id"
				input {
					name = "name"
					value = "anything as a string"
				}
				setting {
					name = "name"
					value = "value"
				}
			}
		}

		data "ibm_project_event_notification" "project_event_notification_instance" {
			id = ibm_project_config.project_config_instance.projectConfig_id
			exclude_configs = true
		}
	`, projectResourceGroup, projectLocation, projectName, projectDescription)
}
