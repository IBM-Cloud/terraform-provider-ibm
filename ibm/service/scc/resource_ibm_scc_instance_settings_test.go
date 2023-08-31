// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccInstanceSettingsBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Settings

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccInstanceSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccInstanceSettingsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccInstanceSettingsExists("ibm_scc_instance_settings.scc_instance_settings", conf),
				),
			},
		},
	})
}

func TestAccIbmSccInstanceSettingsAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Settings
	xCorrelationID := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestID := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	settingsID := fmt.Sprintf("tf_settings_id_%d", acctest.RandIntRange(10, 100))
	xCorrelationIDUpdate := fmt.Sprintf("tf_x_correlation_id_%d", acctest.RandIntRange(10, 100))
	xRequestIDUpdate := fmt.Sprintf("tf_x_request_id_%d", acctest.RandIntRange(10, 100))
	settingsIDUpdate := fmt.Sprintf("tf_settings_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccInstanceSettingsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccInstanceSettingsConfig(xCorrelationID, xRequestID, settingsID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccInstanceSettingsExists("ibm_scc_instance_settings.scc_instance_settings", conf),
					resource.TestCheckResourceAttr("ibm_scc_instance_settings.scc_instance_settings", "x_correlation_id", xCorrelationID),
					resource.TestCheckResourceAttr("ibm_scc_instance_settings.scc_instance_settings", "x_request_id", xRequestID),
					resource.TestCheckResourceAttr("ibm_scc_instance_settings.scc_instance_settings", "settings_id", settingsID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSccInstanceSettingsConfig(xCorrelationIDUpdate, xRequestIDUpdate, settingsIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_instance_settings.scc_instance_settings", "x_correlation_id", xCorrelationIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_instance_settings.scc_instance_settings", "x_request_id", xRequestIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_instance_settings.scc_instance_settings", "settings_id", settingsIDUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_instance_settings.scc_instance_settings",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccInstanceSettingsConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_scc_instance_settings" "scc_instance_settings_instance" {
		}
	`)
}

func testAccCheckIbmSccInstanceSettingsConfig(xCorrelationID string, xRequestID string, settingsID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_instance_settings" "scc_instance_settings_instance" {
			x_correlation_id = "%s"
			x_request_id = "%s"
			event_notifications {
				instance_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/ff88f007f9ff4622aac4fbc0eda36255:7199ae60-a214-4dd8-9bf7-ce571de49d01::"
				updated_on = "2021-01-31T09:44:12Z"
				source_id = "crn:v1:bluemix:public:event-notifications:us-south:a/ff88f007f9ff4622aac4fbc0eda36255:b8b07245-0bbe-4478-b11c-0dce523105fd::"
				source_description = "source_description"
				source_name = "source_name"
			}
			object_storage {
				instance_crn = "instance_crn"
				bucket = "bucket"
				bucket_location = "bucket_location"
				bucket_endpoint = "bucket_endpoint"
				updated_on = "2021-01-31T09:44:12Z"
			}
			settings_id = "%s"
		}
	`, xCorrelationID, xRequestID, settingsID)
}

func testAccCheckIbmSccInstanceSettingsExists(n string, obj securityandcompliancecenterapiv3.Settings) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		adminClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{}

		settings, _, err := adminClient.GetSettings(getSettingsOptions)
		if err != nil {
			return err
		}

		obj = *settings
		return nil
	}
}

func testAccCheckIbmSccInstanceSettingsDestroy(s *terraform.State) error {
	adminClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_instance_settings" {
			continue
		}

		getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{}

		// Try to find the key
		_, response, err := adminClient.GetSettings(getSettingsOptions)

		if err == nil {
			return fmt.Errorf("scc_instance_settings still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_instance_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
