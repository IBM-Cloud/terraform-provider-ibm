---

subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM : Cloud Object Storage Backup Vault"
description: 
  "Manages IBM Cloud Object Storage Backup Vault"
---

# ibm_cos_backup_vault
Creates a backup vault to store bucket backup data.

---

## Example usage


```terraform
data "ibm_resource_group" "cos_group" {
  name = "cos-resource-group"
}

resource "ibm_cos_backup_vault" "backup-vault" {
  backup_vault_name           = "backup_vault_name"
  service_instance_id  = "cos_instance_id to create backup vault"
  region  = "us"
  activity_tracking_management_events = true
  metrics_monitoring_usage_metrics = true
  kms_key_crn = "crn:v1:staging:public:kms:us-south:a/997xxxxxxxxxxxxxxxxxxxxxx54:5xxxxxxxa-fxxb-4xx8-9xx4-f1xxxxxxxxx5:key:af5667d5-dxx5-4xxf-8xxf-exxxxxxxf1d"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 
- `backup_vault_name` - (Required, Forces new resource, String) Name of the backup vault.
- `service_instance_id` - (Required, Forces new resource, String) CRN of the COS instance where the backup vault is to be created.
- `region`- (Required, Forces new resource, String) The location of the COS backup vault.
- `activity_tracking_management_events` - (Optional , Bool) Whether to send notification for the management events for backup vault.
- `metrics_monitoring_usage_metrics` - (Optional , Bool)  Whether usage metrics are collected for this backup vault.
-  `kms_key_crn` - (Optional, Forces new resource, String) Crn of the key protect root key.
  
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `backup_vault_crn` - (String) The CRN of the backup vault.
- `id` - (String) The ID of the backup vault.
