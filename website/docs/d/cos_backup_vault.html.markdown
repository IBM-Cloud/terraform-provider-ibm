---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM: ibm_cos_backup_vault"
description: |-
  Get information about IBM Cloud Object Storage Backup Vault.
---

# ibm_cos_bucket_object

Retrieves information of a backup vault that store the objects to be backed up from a source bucket whenever a backup policy is applied.

## Example usage

```terraform
data "ibm_cos_backup_vault" "vault" {
  backup_vault_name          = "name of the vault"
	service_instance_id = "instance_id"
	region = "us"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 
- `backup_vault_name` - (Required, Forces new resource, String) Name of the backup vault.
- `service_instance_id` - (Required, Forces new resource, String) The location of the COS bucket.
- `region`- (Required, Forces new resource, String) The location of the COS backup vault.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `activity_tracking_management_events` - (Bool) Whether  notification has been sent for the management events for backup vault.
- `metrics_monitoring_usage_metrics` - (Bool)  Whether usage metrics is collected for this BackupVault.
-  `kms_key_crn` - (String) Crn of the Key protect root key.
  
