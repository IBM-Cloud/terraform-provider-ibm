---
subcategory: "Object Storage"
layout: "ibm"
page_title: "IBM: ibm_cos_backup_policy"
description: |-
  Get information about IBM Cloud Object Storage Backup Policy.
---

# ibm_cos_backup_policy

Retrieves information of a backup policy on a given source bucket.

## Example usage

```terraform
data "ibm_cos_backup_policy" "policy" {
  bucket_name          = "name of the source bucket"
	policys_id  = "id of the policy"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 
- `bucket_name` - (Required, String) Name of the bucket name.
- `policy_id` - (Required,String) Id of the policy to be retrieved
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `policy_name` - (String) Name of the policy.
- `delete_after_days` - (Int) Number of days after which the data contained in a RecoveryRange will be deleted.
- `backup_type`- (String) CRN of the backuo vault.
- `target_backup_vault_crn` - (String) Type of backup supported.
