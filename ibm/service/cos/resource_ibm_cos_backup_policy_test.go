package cos_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCosBackup_Policy_Basic_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_Basic_Valid(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName, guid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_BV_With_Activity_Tracking_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_BV_With_Activity_Tracking_Valid(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName, guid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "activity_tracking_management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_BV_With_Metrics_Monitoring_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_BV_With_Metrics_Monitoring_Valid(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName, guid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "metrics_monitoring_usage_metrics", "true"),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_BV_With_KP_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_BV_With_KP_Valid(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName, guid, acc.KmsKeyCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "backup_vault_name", backupVaultName),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "region", region),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault", "kms_key_crn", acc.KmsKeyCrn),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}
func TestAccIBMCosBackup_Policy_Multiple_Policy_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName1 := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	backupVaultName2 := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))

	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName1 := fmt.Sprintf("backup-policy1%d", acctest.RandIntRange(10, 100))
	policyName2 := fmt.Sprintf("backup-policy2%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_Multiple_Policy_Valid(acc.CosCRN, bucketName, backupVaultName1, backupVaultName2, region, accountID, policyName1, policyName2, guid),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault1", "backup_vault_name", backupVaultName1),
					resource.TestCheckResourceAttr("ibm_cos_backup_vault.backup-vault2", "backup_vault_name", backupVaultName2),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy1", "policy_name", policyName1),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy2", "policy_name", policyName2),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_Non_Versioning_Source_Bucket_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_Non_Versioning_Source_Bucket_Invalid(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName, guid),
				ExpectError: regexp.MustCompile("Source bucket does not have versioning enabled"),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_Multiple_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName1 := fmt.Sprintf("backup-policy1%d", acctest.RandIntRange(10, 100))
	policyName2 := fmt.Sprintf("backup-policy2%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_Multiple_Invalid(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName1, policyName2, guid),
				ExpectError: regexp.MustCompile("Bad Request"),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_Invalid_Backup_Type(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	backupVaultName := fmt.Sprintf("terraform-backup-vault%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_Invalid_Backup_Type(acc.CosCRN, bucketName, backupVaultName, region, accountID, policyName, guid),
				ExpectError: regexp.MustCompile("Error:"),
			},
		},
	})
}

func testAccCheckIBMCosBackup_Policy_Basic_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class          = "standard"
	object_versioning {
		enable  = true
		}
	  
	  }
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name           = "%s"
	service_instance_id  = "%s"
	region  = "%s"
	}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, bucketVaultName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName)
}

func testAccCheckIBMCosBackup_Policy_Multiple_Policy_Valid(instance_id string, bucketName string, bucketVaultName1 string, bucketVaultName2 string, region string, accountId string, policyName1 string, policyName2 string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class         = "standard"
	object_versioning {
		enable  = true
	}
}
resource "ibm_cos_backup_vault" "backup-vault1" {
	backup_vault_name    = "%s"
	service_instance_id  = "%s"
	region  = "%s"
	}

resource "ibm_cos_backup_vault" "backup-vault2" {
	backup_vault_name    = "%s"
	service_instance_id  = "%s"
	region  = "%s"
	}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}
resource "ibm_iam_authorization_policy" "policy2" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy1" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault1.backup_vault_crn
	backup_type = "continuous"
}
	
resource "ibm_cos_backup_policy" "policy2" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault2.backup_vault_crn
	backup_type = "continuous"
	}
	`, bucketName, instance_id, bucketVaultName1, instance_id, region, bucketVaultName2, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName1, accountId, guid, bucketName, accountId, guid, bucketVaultName2, policyName1, policyName2)
}

func testAccCheckIBMCosBackup_Policy_BV_With_Activity_Tracking_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
 }
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name           = "%s"
	service_instance_id  = "%s"
	region  = "%s"
	activity_tracking_management_events = true
	}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, bucketVaultName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName)
}

func testAccCheckIBMCosBackup_Policy_BV_With_Metrics_Monitoring_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
	  
}
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name           = "%s"
	service_instance_id  = "%s"
	region  = "%s"
	metrics_monitoring_usage_metrics = true
}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, bucketVaultName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName)
}

func testAccCheckIBMCosBackup_Policy_BV_With_KP_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, key string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name            = "%s"
	resource_instance_id   = "%s"
	cross_region_location  =  "us"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
}
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name    = "%s"
	service_instance_id  = "%s"
	region               = "%s"
	kms_key_crn          = "%s"
}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, bucketVaultName, instance_id, region, key, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName)
}

func testAccCheckIBMCosBackup_Policy_Non_Versioning_Source_Bucket_Invalid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class          = "standard"  
}
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name           = "%s"
	service_instance_id  = "%s"
	region  = "%s"
}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, bucketVaultName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName)
}

func testAccCheckIBMCosBackup_Policy_Multiple_Invalid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName1 string, policyName2 string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}  
}
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name           = "%s"
	service_instance_id  = "%s"
	region  = "%s"
}
resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy1" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
}
	
resource "ibm_cos_backup_policy" "policy2" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "continuous"
   }
	`, bucketName, instance_id, bucketVaultName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName1, policyName2)
}

func testAccCheckIBMCosBackup_Policy_Invalid_Backup_Type(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "us"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
}
resource "ibm_cos_backup_vault" "backup-vault" {
	backup_vault_name           = "%s"
	service_instance_id  = "%s"
	region  = "%s"
}

resource "ibm_iam_authorization_policy" "policy" {
	roles                  = [
		"Backup Manager", "Writer"
	]
	subject_attributes {
	  name  = "accountId"
	  value = "%s"
	}
	subject_attributes {
	  name  = "serviceName"
	  value = "cloud-object-storage"
	}
	subject_attributes {
	  name  = "serviceInstance"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resource"
	  value = "%s"
	}
	subject_attributes {
	  name  = "resourceType"
	  value = "bucket"
	}
	resource_attributes {
	  name     = "accountId"
	  operator = "stringEquals"
	  value    = "%s"
	}
	resource_attributes {
	  name     = "serviceName"
	  operator = "stringEquals"
	  value    = "cloud-object-storage"
	}
	resource_attributes { 
	  name  =  "serviceInstance"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resource"
	  operator = "stringEquals"
	  value =  "%s"
	}
	resource_attributes { 
	  name  =  "resourceType"
	  operator = "stringEquals"
	  value =  "backup-vault" 
	}
}

resource "ibm_cos_backup_policy" "policy" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
	backup_type = "invalid"
}
		
	`, bucketName, instance_id, bucketVaultName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName)
}
