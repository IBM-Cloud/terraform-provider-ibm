// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasRestorePointsBasic(t *testing.T) {
	var conf backuprecoveryv1.GetRestorePointsInTimeRangeResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	endTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	startTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasRestorePointsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasRestorePointsConfigBasic(xIbmTenantID, endTimeUsecs, startTimeUsecs),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRestorePointsExists("ibm_baas_restore_points.baas_restore_points_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "end_time_usecs", endTimeUsecs),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "start_time_usecs", startTimeUsecs),
				),
			},
		},
	})
}

func TestAccIbmBaasRestorePointsAllArgs(t *testing.T) {
	var conf backuprecoveryv1.GetRestorePointsInTimeRangeResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	endTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	environment := "kVMware"
	sourceID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	startTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasRestorePointsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasRestorePointsConfig(xIbmTenantID, endTimeUsecs, environment, sourceID, startTimeUsecs),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRestorePointsExists("ibm_baas_restore_points.baas_restore_points_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "end_time_usecs", endTimeUsecs),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "environment", environment),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "source_id", sourceID),
					resource.TestCheckResourceAttr("ibm_baas_restore_points.baas_restore_points_instance", "start_time_usecs", startTimeUsecs),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_restore_points.baas_restore_points",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasRestorePointsConfigBasic(xIbmTenantID string, endTimeUsecs string, startTimeUsecs string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_restore_points" "baas_restore_points_instance" {
			x_ibm_tenant_id = "%s"
			end_time_usecs = %s
			protection_group_ids = "FIXME"
			start_time_usecs = %s
		}
	`, xIbmTenantID, endTimeUsecs, startTimeUsecs)
}

func testAccCheckIbmBaasRestorePointsConfig(xIbmTenantID string, endTimeUsecs string, environment string, sourceID string, startTimeUsecs string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_restore_points" "baas_restore_points_instance" {
			x_ibm_tenant_id = "%s"
			end_time_usecs = %s
			environment = "%s"
			protection_group_ids = "FIXME"
			source_id = %s
			start_time_usecs = %s
		}
	`, xIbmTenantID, endTimeUsecs, environment, sourceID, startTimeUsecs)
}

func testAccCheckIbmBaasRestorePointsExists(n string, obj backuprecoveryv1.GetRestorePointsInTimeRangeResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getRestorePointsInTimeRangeOptions := &backuprecoveryv1.GetRestorePointsInTimeRangeOptions{}

		getRestorePointsInTimeRangeParams, _, err := backupRecoveryClient.GetRestorePointsInTimeRange(getRestorePointsInTimeRangeOptions)
		if err != nil {
			return err
		}

		obj = *getRestorePointsInTimeRangeParams
		return nil
	}
}

func testAccCheckIbmBaasRestorePointsDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_restore_points" {
			continue
		}

		getRestorePointsInTimeRangeOptions := &backuprecoveryv1.GetRestorePointsInTimeRangeOptions{}

		// Try to find the key
		_, response, err := backupRecoveryClient.GetRestorePointsInTimeRange(getRestorePointsInTimeRangeOptions)

		if err == nil {
			return fmt.Errorf("baas_restore_points still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_restore_points (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
