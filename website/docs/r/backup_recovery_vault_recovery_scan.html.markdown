---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_vault_recovery_scan"
description: |-
  Manages backup_recovery_vault_recovery_scan.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_vault_recovery_scan

Create, update, and delete backup_recovery_vault_recovery_scans with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_vault_recovery_scan" "backup_recovery_vault_recovery_scan_instance" {
  cloud_type = "ibm"
  recovery_scan_request_params {
		vault_id = 1
  }
  x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cloud_type` - (Required, Forces new resource, String) Specifies the cloud type where the vault is registered for recovery scan. Currently, only 'ibm' is supported.
  * Constraints: Allowable values are: `ibm`.
* `recovery_scan_request_params` - (Optional, Forces new resource, List) Specifies the parameters specific to the Backup and Recovery instance. which is the vault.
Nested schema for **recovery_scan_request_params**:
	* `vault_id` - (Required, Integer) Specifies the unique id of the IBM Cloud Backup and Recovery instance for which the recovery scan is to be initiated.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the unique id of the tenant.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.
* `uid` - (String) Specifies the unique id of the recovery scan.


## Import
Not Supported