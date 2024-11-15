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

 you must have `writer` or `backuo manager` platform roles on source bucket and backup vault. And from backup vault to target bucket.
 Add depends_on on ibm_iam_authorization_policy.policy in template to make sure.
 The source bucket should have object versioning enabled.

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
  bucket_crn      = ibm_cos_bucket.bucket.crn
  policy_name = "policy_name"
  target_backup_vault_crn = ibm_cos_backup_vault.backup-vault.backup_vault_crn
  backup_type = "continuous"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 
- `bucket_crn` - (Required, Forces new resource, String) CRN of the source bucket.
- `policy_name` - (Required, Forces new resource, String) Name of the policy.
- `backup_type`- (Required, Forces new resource, String) CRN of the backuo vault.
- `target_backup_vault_crn` - (Required ,Forces new resource, String) Type of backup supported.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the backup policy.
