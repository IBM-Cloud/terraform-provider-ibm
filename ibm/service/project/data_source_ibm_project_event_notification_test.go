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
	projectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectResourceGroup := fmt.Sprintf("Default")
	projectLocation := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEventNotificationDataSourceConfigBasic(projectResourceGroup, projectLocation, projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "name"),
				),
			},
		},
	})
}

func TestAccIbmProjectEventNotificationDataSourceAllArgs(t *testing.T) {
	projectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectResourceGroup := fmt.Sprintf("Default")
	projectLocation := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEventNotificationDataSourceConfig(projectName, projectDescription, projectResourceGroup, projectLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project_event_notification.project_event_notification", "metadata.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectEventNotificationDataSourceConfigBasic(projectResourceGroup string, projectLocation string, projectName string) string {
	return fmt.Sprintf(`
		resource "ibm_project_instance" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}

		data "ibm_project_event_notification" "project_event_notification" {
			project_id = ibm_project_instance.project_instance.id
		}
	`, projectResourceGroup, projectLocation, projectName)
}

func testAccCheckIbmProjectEventNotificationDataSourceConfig(projectName string, projectDescription string, projectResourceGroup string, projectLocation string) string {
	return fmt.Sprintf(`
		resource "ibm_project_instance" "project_instance" {
			name = "%s"
			description = "%s"
			configs {
				name = "name"
				labels = [ "labels" ]
				description = "description"
				locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global"
				input {
					name = "name"
				}
				setting {
					name = "name"
					value = "value"
				}
			}
			resource_group = "%s"
			location = "%s"
		}

		data "ibm_project_event_notification" "project_event_notification" {
			project_id = ibm_project_instance.project_instance.id
		}
	`, projectName, projectDescription, projectResourceGroup, projectLocation)
}
