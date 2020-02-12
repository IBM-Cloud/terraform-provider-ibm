---
layout: "ibm"
page_title: "IBM: pi_tenant"
sidebar_current: "docs-ibm-datasources-pi-tenant"
description: |-
  Manages a tenant in the IBM Power Virtual Server Cloud.
---

# ibm\_pi_tenant

Import the details of an existing IBM Power Virtual Server Cloud tenant as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_tenant" "ds_tenant" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this tenant.
* `creation_date` - The date on which the tenant was created.
* `enabled` - Indicates whether the tenant is enabled.
* `tenant_name` - The name of the tenant.
* `cloudinstances` - Lists the regions and instance IDs this tenant owns.
  * `cloud_instance_id` - The unique identifier of the cloud instance.
  * `region` - The region of the cloud instance.
