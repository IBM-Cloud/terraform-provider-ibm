package cos_test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCosBackup_Policy_Basic_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_Basic_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_BV_With_Activity_Tracking_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_BV_With_Activity_Tracking_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_BV_With_Metrics_Monitoring_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_BV_With_Metrics_Monitoring_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_BV_With_KP_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_BV_With_KP_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
				),
			},
		},
	})
}
func TestAccIBMCosBackup_Policy_Multiple_Policy_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIBMCosBackup_Policy_Multiple_Policy_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, acc.BackupVaultName2, region, accountID, policyName1, policyName2, guid, acc.BackupVaultCrn, acc.BackupVaultCrn2),
				Check: resource.ComposeAggregateTestCheckFunc(
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
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_Non_Versioning_Source_Bucket_Invalid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				ExpectError: regexp.MustCompile("Source bucket does not have versioning enabled"),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_Multiple_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
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
				Config:      testAccCheckIBMCosBackup_Policy_Multiple_Invalid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName1, policyName2, guid, acc.BackupVaultCrn),
				ExpectError: regexp.MustCompile("Bad Request"),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_Invalid_Backup_Type(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_Invalid_Backup_Type(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				ExpectError: regexp.MustCompile("Error:"),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_With_Initial_Retention_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	deleteAfterDays := 1
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_With_Initial_Retention_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn, deleteAfterDays),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "initial_delete_after_days", strconv.Itoa(deleteAfterDays)),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_With_Initial_Retention_Max_Days_Valid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	deleteAfterDays := 36500
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBackup_Policy_With_Initial_Retention_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn, deleteAfterDays),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "policy_name", policyName),
					resource.TestCheckResourceAttr("ibm_cos_backup_policy.policy", "initial_delete_after_days", strconv.Itoa(deleteAfterDays)),
				),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_With_Initial_Retention_Zero_Days_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	deleteAfterDays := 0
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_With_Initial_Retention_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn, deleteAfterDays),
				ExpectError: regexp.MustCompile(`"initial_delete_after_days" must contain a valid int value should be in range\(1, 36500\), got 0`),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_With_Initial_Retention_Less_Than_Zero_Days_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	deleteAfterDays := -1
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_With_Initial_Retention_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn, deleteAfterDays),
				ExpectError: regexp.MustCompile(`"initial_delete_after_days" must contain a valid int value should be in range\(1, 36500\), got -1`),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_With_Initial_Retention_More_Than_36500_Days_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	deleteAfterDays := 36501
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_With_Initial_Retention_Valid(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn, deleteAfterDays),
				ExpectError: regexp.MustCompile(`"initial_delete_after_days" must contain a valid int value should be in range\(1, 36500\), got 36501`),
			},
		},
	})
}

func TestAccIBMCosBackup_Policy_With_Initial_Retention_Invalid(t *testing.T) {
	accountID := acc.IBM_AccountID_REPL
	bucketName := fmt.Sprintf("terraform-backup-source%d", acctest.RandIntRange(10, 100))
	region := "us"
	guid := strings.Split(acc.CosCRN, ":")[7]
	policyName := fmt.Sprintf("backup-policy%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBackup_Policy_Without_Retention(acc.CosCRN, bucketName, acc.BackupVaultName, region, accountID, policyName, guid, acc.BackupVaultCrn),
				ExpectError: regexp.MustCompile("Missing required argument"),
			},
		},
	})
}

func testAccCheckIBMCosBackup_Policy_Basic_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
		}
	  
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_Multiple_Policy_Valid(instance_id string, bucketName string, bucketVaultName1 string, bucketVaultName2 string, region string, accountId string, policyName1 string, policyName2 string, guid string, backupVaultCrn string, backupVaultCrn2 string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class         = "standard"
	object_versioning {
		enable  = true
	}
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
	
resource "ibm_cos_backup_policy" "policy2" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
	}
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName1, accountId, guid, bucketName, accountId, guid, bucketVaultName2, policyName1, backupVaultCrn, policyName2, backupVaultCrn2)
}

func testAccCheckIBMCosBackup_Policy_BV_With_Activity_Tracking_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_BV_With_Metrics_Monitoring_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
	  
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_BV_With_KP_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name            = "%s"
	resource_instance_id   = "%s"
	cross_region_location  =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_Non_Versioning_Source_Bucket_Invalid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"  
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_Multiple_Invalid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName1 string, policyName2 string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}  
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
	
resource "ibm_cos_backup_policy" "policy2" {
	bucket_crn      = ibm_cos_bucket.bucket.crn
	policy_name = "%s"
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
   }
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName1, backupVaultCrn, policyName2, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_Invalid_Backup_Type(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
	}
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
	initial_delete_after_days  = 2
	target_backup_vault_crn = "%s"
	backup_type = "invalid"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_With_Initial_Retention_Valid(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string, deleteAfterDays int) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
		}
	  
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
	initial_delete_after_days       = "%d"
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, deleteAfterDays, backupVaultCrn)
}

func testAccCheckIBMCosBackup_Policy_Without_Retention(instance_id string, bucketName string, bucketVaultName string, region string, accountId string, policyName string, guid string, backupVaultCrn string) string {

	return fmt.Sprintf(`

resource "ibm_cos_bucket" "bucket" {
	bucket_name           = "%s"
	resource_instance_id  = "%s"
	cross_region_location =  "%s"
	storage_class          = "standard"
	object_versioning {
		enable  = true
		}
	  
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
	target_backup_vault_crn = "%s"
	backup_type = "continuous"
}
		
	`, bucketName, instance_id, region, accountId, guid, bucketName, accountId, guid, bucketVaultName, policyName, backupVaultCrn)
}
