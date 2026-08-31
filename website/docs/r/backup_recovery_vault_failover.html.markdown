---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_vault_failover"
description: |-
  Manages backup_recovery_vault_failover.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_vault_failover

Create, update, and delete backup_recovery_vault_failovers with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_vault_failover" "backup_recovery_vault_failover_instance" {
  cloud_type = "ibm"
  failover_request_params {
		vault_id = 1
  }
  x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `cloud_type` - (Required, Forces new resource, String) Specifies the type of the Backup and Recovery instance. Currently, only 'ibm' is supported.
  * Constraints: Allowable values are: `ibm`.
* `failover_request_params` - (Optional, Forces new resource, List) Specifies the parameters specific to the Backup and Recovery instance. viz the vault.
Nested schema for **failover_request_params**:
	* `vault_id` - (Required, Integer) Specifies the unique id of the IBM Cloud Backup and Recovery instance for which the failover is to be initiated.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the unique id of the tenant.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.
* `uid` - (String) Specifies the unique id of the failover.


## Import
Not Supported

