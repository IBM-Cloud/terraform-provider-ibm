---
layout: "ibm"
page_title: "IBM: Tenant"
sidebar_current: "docs-ibm-datasources-pi-tenant"
description: |-
  Manages a tenant in the IBM Power Virtual Server Cloud.
---

# ibm\_pi_network

Import the details of an existing IBM Power Virtual Server Cloud tenant as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_tenant" "ds_tenant" {
    powerinstanceid = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `powerinstanceid` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `tenantid` - The unique identifier for this tenant.
* `creationdate` - The date on which the tenant was created.
* `enabled` - Indicates whether the tenant is enabled.
* `tenantname` - The name of the tenant.
* `cloudinstances` - Lists the regions and instance IDs this tenant owns.
