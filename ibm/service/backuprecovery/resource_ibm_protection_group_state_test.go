// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmProtectionGroupStateBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	action := "kPause"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProtectionGroupStateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupStateConfigBasic(action),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProtectionGroupStateExists("ibm_protection_group_state.protection_group_state_instance", conf),
					resource.TestCheckResourceAttr("ibm_protection_group_state.protection_group_state_instance", "action", action),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_protection_group_state.protection_group_state",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProtectionGroupStateConfigBasic(action string) string {
	return fmt.Sprintf(`
		resource "ibm_protection_group_state" "protection_group_state_instance" {
			action = "%s"
			ids = "FIXME"
		}
	`, action)
}

func testAccCheckIbmProtectionGroupStateExists(n string, obj backuprecoveryv1.ProtectionGroupResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)

		updateProtectionGroupsStateRequest, _, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)
		if err != nil {
			return err
		}

		obj = *updateProtectionGroupsStateRequest
		return nil
	}
}

func testAccCheckIbmProtectionGroupStateDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_protection_group_state" {
			continue
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)

		if err == nil {
			return fmt.Errorf("Update state of Protection Groups. still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Update state of Protection Groups. (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
