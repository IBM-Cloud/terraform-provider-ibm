// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
)

func TestAccIbmSccInstanceSettingsBasic(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Settings

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccInstanceSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccInstanceSettingsConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccInstanceSettingsExists("ibm_scc_instance_settings.scc_instance_settings_instance", conf),
				),
			},
		},
	})
}

func TestAccIbmSccInstanceSettingsAllArgs(t *testing.T) {
	var conf securityandcompliancecenterapiv3.Settings

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckScc(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSccInstanceSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccInstanceSettingsConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccInstanceSettingsExists("ibm_scc_instance_settings.scc_instance_settings_instance", conf),
				),
			},
			{
				Config: testAccCheckIbmSccInstanceSettingsConfig(acc.SccInstanceID, acc.SccEventNotificationsCRN, acc.SccObjectStorageCRN, acc.SccObjectStorageBucket),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSccInstanceSettingsExists("ibm_scc_instance_settings.scc_instance_settings_instance", conf),
				),
			},
			{
				ResourceName:      "ibm_scc_instance_settings.scc_instance_settings_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSccInstanceSettingsConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_instance_settings" "scc_instance_settings_instance" {
			instance_id = "%s"
			event_notifications { }
			object_storage { }
		}
	`, instanceID)
}

func testAccCheckIbmSccInstanceSettingsConfig(instanceID, enInstanceCRN, objStorInstanceCRN, objStorBucket string) string {
	return fmt.Sprintf(`
		resource "ibm_scc_instance_settings" "scc_instance_settings_instance" {
			instance_id = "%s"
			event_notifications {
				instance_crn = "%s"
				source_name = "scc compliance"
			}
			object_storage {
				instance_crn = "%s"
				bucket = "%s"
			}
		}
	`, instanceID, enInstanceCRN, objStorInstanceCRN, objStorBucket)
}

func testAccCheckIbmSccInstanceSettingsExists(n string, obj securityandcompliancecenterapiv3.Settings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
		}

		adminClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecurityAndComplianceCenterV3()
		if err != nil {
			return err
		}

		getSettingsOptions := &securityandcompliancecenterapiv3.GetSettingsOptions{}
		instanceID := acc.SccInstanceID
		getSettingsOptions.SetInstanceID(instanceID)

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
		instanceID := acc.SccInstanceID
		getSettingsOptions.SetInstanceID(instanceID)

		// Deleting a instance_settings_resource doesn't delete the entity
		_, response, err := adminClient.GetSettings(getSettingsOptions)
		if response.StatusCode != 200 {
			return flex.FmtErrorf("Error checking for scc_instance_settings (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
