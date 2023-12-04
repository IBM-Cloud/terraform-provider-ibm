---
layout: "ibm"
page_title: "IBM : ibm_logs_router_tenant"
description: |-
  Get information about logs_router_tenant
subcategory: "Logs Router"
---

# ibm_logs_router_tenant

Provides a read-only data source to retrieve information about a logs_router_tenant. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_router_tenant" "logs_router_tenant" {
	tenant_id = "f3a466c9-c4db-4eee-95cc-ba82db58e2b5"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `tenant_id` - (Required, Forces new resource, String) The instance ID of the tenant.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[A-F,0-9,-]/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_router_tenant.
* `account_id` - (String) The account ID the tenant belongs to.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-z,A-Z,0-9,-]/`.

* `created_at` - (String) Time stamp the tenant was originally created.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.

* `target_host` - (String) Host name of log-sink.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.

* `target_instance_crn` - (String) Cloud resource name of the log-sink target instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,:,-]/`.

* `target_port` - (Integer) Network port of log sink.

* `target_type` - (String) Type of log-sink.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-]/`.

* `updated_at` - (String) time stamp the tenant was last updated.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.

