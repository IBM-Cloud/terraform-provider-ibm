---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_source_refresh"
description: |-
  Manages backup_recovery_protection_source_refresh.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_protection_source_refresh

Refresh protection source with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_protection_source_refresh" "backup_recovery_protection_source_refresh_instance" {
  x_ibm_tenant_id = "tenantId"
  backup_recovery_protection_source_id = "id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the unique id of the tenant.
* `backup_recovery_protection_source_id` - (Required, Forces new resource,Integer) protection source Id.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.
* `backup_recovery_protection_source_id` - (Integer) protection source Id.


## Import
Import is not supported