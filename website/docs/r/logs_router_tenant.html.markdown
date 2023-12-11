---
layout: "ibm"
page_title: "IBM : ibm_logs_router_tenant"
description: |-
  Manages logs_router_tenant.
subcategory: "Logs Router"
---

# ibm_logs_router_tenant

~> **Beta:** This resource is in Beta, and is subject to change.

Create, update, and delete logs_router_tenants with this resource.

## Example Usage

```hcl
resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  target_host = "www.example.com"
  target_instance_crn = "crn:v1:bluemix:public:logdna:eu-de:a/3516b8fa0a174a71899f5affa4f18d78:3517d2ed-9429-af34-ad52-34278391cbc8::"
  target_port = 10
  target_type = "logdna"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `target_host` - (Required, String) Host name of log-sink.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
* `target_instance_crn` - (Required, String) Cloud resource name of the log-sink target instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,:,-]/`.
* `target_port` - (Required, Integer) Network port of log sink.
* `target_type` - (Required, String) Type of log-sink.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-]/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_router_tenant.
* `account_id` - (String) The account ID the tenant belongs to.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-z,A-Z,0-9,-]/`.
* `created_at` - (String) Time stamp the tenant was originally created.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.
* `updated_at` - (String) time stamp the tenant was last updated.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.

## Import

You can import the `ibm_logs_router_tenant` resource by using `id`. Unique ID of the created instance.

# Syntax
<pre>
$ terraform import ibm_logs_router_tenant.logs_router_tenant &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_logs_router_tenant.logs_router_tenant aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa
```
