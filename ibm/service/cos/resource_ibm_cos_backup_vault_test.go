package cos_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCosBackup_Vault_Create_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_Valid(acc.CosCRN, backupVaultName, region),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Vault_Create_With_Activity_Tracking_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	activityTrackingManagementEvents := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_With_Activity_Tracking_Valid(acc.CosCRN, backupVaultName, region, activityTrackingManagementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "activity_tracking_management_events", "true"),
				),
			},
		},
	})
}
func TestAccIBMCosBackup_Vault_Create_With_Metrics_Monitoring_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	usageMetrics := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_With_Metrics_Monitoring_Valid(acc.CosCRN, backupVaultName, region, usageMetrics),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "metrics_monitoring_usage_metrics", "true"),
				),
			},
		},
	})
}
func TestAccIBMCosBackup_Vault_Create_With_KP_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_With_KP_Valid(acc.CosCRN, backupVaultName, region, acc.KmsKeyCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "kms_key_crn", acc.KmsKeyCrn),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Vault_Create_With_AT_Enabled_MM_Disabled_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	activityTrackingManagementEvents := true
	usageMetrics := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_With_AT_Enabled_MM_Disabled_Valid(acc.CosCRN, backupVaultName, region, activityTrackingManagementEvents, usageMetrics),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "activity_tracking_management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "metrics_monitoring_usage_metrics", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Vault_Create_With_AT_Disabled_MM_Enabled_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	activityTrackingManagementEvents := false
	usageMetrics := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_With_AT_Disabled_MM_Enabled_Valid(acc.CosCRN, backupVaultName, region, activityTrackingManagementEvents, usageMetrics),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "activity_tracking_management_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "metrics_monitoring_usage_metrics", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Vault_Create_Multiple_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	backupVaultName2 := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_Multiple_Valid(acc.CosCRN, backupVaultName, backupVaultName2, region),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault2", "backup_vault_name", backupVaultName2),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault2", "region", region),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Vault_Create_Multiple_Vault_With_Same_Name_Invalid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Vault_Create_Multiple_Invalid(acc.CosCRN, backupVaultName, region),
				ExpectError: regexp.MustCompile("Error: Failed to create the backup vault"),
			},
		},
	})
}

func TestAccIBMCosBackup_Vault_Create_All_Configurations_Valid(t *testing.T) {
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	activityTrackingManagementEvents := true
	usageMetrics := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Vault_Create_With_All_Configurations_Valid(acc.CosCRN, backupVaultName, region, activityTrackingManagementEvents, usageMetrics, acc.KmsKeyCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "activity_tracking_management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "metrics_monitoring_usage_metrics", "true"),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "kms_key_crn", acc.KmsKeyCrn),
				),
			},
		},
	})
}

// tests

func testAccCheckIBMCosBackup_Vault_Create_Valid(instance_id string, bucketVaultName string, region string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
	}
		
	`, bucketVaultName, instance_id, region)
}

func testAccCheckIBMCosBackup_Vault_Create_With_Activity_Tracking_Valid(instance_id string, bucketVaultName string, region string, managementEvents bool) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
		activity_tracking_management_events = "%t"
	}
		
	`, bucketVaultName, instance_id, region, managementEvents)
}

func testAccCheckIBMCosBackup_Vault_Create_With_Metrics_Monitoring_Valid(instance_id string, bucketVaultName string, region string, usageMetrics bool) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
		metrics_monitoring_usage_metrics = "%t"
	}
		
	`, bucketVaultName, instance_id, region, usageMetrics)
}

func testAccCheckIBMCosBackup_Vault_Create_With_KP_Valid(instance_id string, bucketVaultName string, region string, kmsKeyCrn string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
		kms_key_crn = "%s"
	}
		
	`, bucketVaultName, instance_id, region, kmsKeyCrn)
}

func testAccCheckIBMCosBackup_Vault_Create_With_AT_Enabled_MM_Disabled_Valid(instance_id string, bucketVaultName string, region string, managementEvents bool, usageMetrics bool) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
		activity_tracking_management_events = "%t"
		metrics_monitoring_usage_metrics = "%t"
	}
		
	`, bucketVaultName, instance_id, region, managementEvents, usageMetrics)
}

func testAccCheckIBMCosBackup_Vault_Create_With_AT_Disabled_MM_Enabled_Valid(instance_id string, bucketVaultName string, region string, managementEvents bool, usageMetrics bool) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
		activity_tracking_management_events = "%t"
		metrics_monitoring_usage_metrics = "%t"
	}
		
	`, bucketVaultName, instance_id, region, managementEvents, usageMetrics)
}

func testAccCheckIBMCosBackup_Vault_Create_Multiple_Valid(instance_id string, bucketVaultName string, bucketVaultName2 string, region string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
	}

	resource "ibm_cos_backup_vault" "backup-vault2" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
	}
		
	`, bucketVaultName, instance_id, region, bucketVaultName2, instance_id, region)
}

func testAccCheckIBMCosBackup_Vault_Create_Multiple_Invalid(instance_id string, bucketVaultName string, region string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
	}

	resource "ibm_cos_backup_vault" "backup-vault2" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
	}
		
	`, bucketVaultName, instance_id, region, bucketVaultName, instance_id, region)
}

func testAccCheckIBMCosBackup_Vault_Create_With_All_Configurations_Valid(instance_id string, bucketVaultName string, region string, managementEvents bool, usageMetrics bool, key string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_backup_vault" "backup-vault" {
		backup_vault_name           = "%s"
		service_instance_id  = "%s"
		region  = "%s"
		activity_tracking_management_events = "%t"
		metrics_monitoring_usage_metrics = "%t"
		kms_key_crn = "%s"
	}
		
	`, bucketVaultName, instance_id, region, managementEvents, usageMetrics, key)
}
