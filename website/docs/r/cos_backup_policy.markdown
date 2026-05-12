---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Backup Policy"
description: 
  "Manages IBM Cloud Object Storage Backup Policy"
---

# ibm_cos_backup_policy
Creates a backup policy on a given source bucket.


**Note:**

 The source bucket must have object versioning enabled.

 **Note:**

 Backup policies require a service-to-service IAM policy granting sync permissions from the source to the target bucket. Adding a `depends_on` for `ibm_iam_authorization_policy.policy` ensures that this policy is in place.

---

## Example usage


```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

resource "ibm_cos_bucket" "backup-source-bucket" {
  bucket_name           = "bucket-name"
  resource_instance_id  = "cos_instance_id"
  cross_region_location =  "us"
  storage_class          = "standard"
  object_versioning {
    enable  = true
  }

}

resource "ibm_cos_backup_vault" "backup-vault" {
  backup_vault_name           = "backup_vault_name"
  service_instance_id  = "cos_instance_id to create backup vault"
  region  = "us"
  activity_tracking_management_events = true
  metrics_monitoring_usage_metrics = true
  kms_key_crn = "crn:v1:staging:public:kms:us-south:a/997xxxxxxxxxxxxxxxxxxxxxx54:5xxxxxxxa-fxxb-4xx8-9xx4-f1xxxxxxxxx5:key:af5667d5-dxx5-4xxf-8xxf-exxxxxxxf1d"
}

resource "ibm_iam_authorization_policy" "policy" {
		roles                  = [
			"Backup Manager", "Writer"
		]
		subject_attributes {
		  name  = "accountId"
		  value = "account_id of the cos account"
		}
		subject_attributes {
		  name  = "serviceName"
		  value = "cloud-object-storage"
		}
		subject_attributes {
		  name  = "serviceInstance"
		  value = "exxxxx34-xxxx-xxxx-xxxx-d6xxxxxxxx9"
		}
		subject_attributes {
		  name  = "resource"
		  value = "source-bucket-name"
		}
		subject_attributes {
		  name  = "resourceType"
		  value = "bucket"
		}
		resource_attributes {
		  name     = "accountId"
		  operator = "stringEquals"
		  value    = "account id of the cos account of backup vault"
		}
		resource_attributes {
		  name     = "serviceName"
		  operator = "stringEquals"
		  value    = "cloud-object-storage"
		}
		resource_attributes { 
		  name  =  "serviceInstance"
		  operator = "stringEquals"
		  value =  "exxxxx34-xxxx-xxxx-xxxx-d6xxxxxxxx9"
		}
		resource_attributes { 
		  name  =  "resource"
		  operator = "stringEquals"
		  value =  "backup-vault-name"
		}
		resource_attributes { 
		  name  =  "resourceType"
		  operator = "stringEquals"
		  value =  "backup-vault" 
		}
	}

resource "ibm_cos_backup_policy" "policy" {
  depends_on      = [ibm_iam_authorization_policy.policy]
  bucket_crn      = ibm_cos_bucket.bucket.crn
  initial_delete_after_days = 2
  policy_name = "policy_name"
  target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
  backup_type = "continuous"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 
- `bucket_crn` - (Required, Forces new resource, String) CRN of the source bucket.
- `initial_delete_after_days` - (Required, String) Number of days after which the data contained within the RecoveryRange will be deleted.

  **Note:**
Once set the value of `initial_delete_after_days` cannot be updated.

- `policy_name` - (Required, Forces new resource, String) Name of the policy.
- `backup_type`- (Required, Forces new resource, String) Backup type. Currently only  `continuous` is supported.
- `target_backup_vault_crn` - (Required ,Forces new resource, String) CRN of the target backup vault.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the backup policy.
