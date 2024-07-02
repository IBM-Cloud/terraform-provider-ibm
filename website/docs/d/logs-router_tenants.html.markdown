---
layout: "ibm"
page_title: "IBM : ibm_logs-router_tenants"
description: |-
  Get information about logs-router_tenants
subcategory: "IBM Cloud Logs Routing"
---

# ibm_logs-router_tenants

Provides a read-only data source to retrieve information about logs-router_tenants. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs-router_tenants" "logs_router_tenants" {
	ibm_api_version = ibm_logs-router_tenant.logs_router_tenant_instance.ibm_api_version
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `ibm_api_version` - (Required, String) Requests the version of the API as of a date in the format YYYY-MM-DD. Any date up to the current date can be provided. Specify the current date to request the latest version.
  * Constraints: The maximum length is `10` characters. The minimum length is `10` characters. The value must match regular expression `/^[0-9]{4}-[0-9]{2}-[0-9]{2}$/`.
* `name` - (Optional, String) Optional: The name of a tenant.
  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/[A-F,0-9,-]/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs-router_tenants.
* `tenants` - (List) List of tenants in the account.
  * Constraints: The maximum length is `1` item. The minimum length is `0` items.
Nested schema for **tenants**:
	* `created_at` - (String) Time stamp the tenant was originally created.
	  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.
	* `crn` - (String) Cloud resource name of the tenant.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,:,-]/`.
	* `etag` - (String) Resource version identifier.
	  * Constraints: The maximum length is `66` characters. The minimum length is `66` characters. The value must match regular expression `/(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"/`.
	* `id` - (String) Unique ID of the tenant.
	* `name` - (String) The name for this tenant. The name is regionally unique across all tenants in the account.
	  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
	* `targets` - (List) List of targets.
	  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
	Nested schema for **targets**:
		* `created_at` - (String) Time stamp the target was originally created.
		  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.
		* `etag` - (String) Resource version identifier.
		  * Constraints: The maximum length is `66` characters. The minimum length is `66` characters. The value must match regular expression `/(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"/`.
		* `id` - (String) Unique ID of the target.
		* `log_sink_crn` - (String) Cloud resource name of the log-sink target instance.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,:,-]/`.
		* `name` - (String) The name for this tenant target. The name is unique across all targets for this tenant.
		  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
		* `parameters` - (List) List of properties returned from a successful list operation for a log-sink of type IBM Log Analysis (logdna).
		Nested schema for **parameters**:
			* `host` - (String) Host name of the log-sink.
			  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
			* `port` - (Integer) Network port of the log-sink.
			  * Constraints: The maximum value is `65535`. The minimum value is `1`.
		* `type` - (String) Type of log-sink. Identical to the <code>service-name</code> segment of <code>log_sink_crn</code>.
		  * Constraints: Allowable values are: `logdna`.
		* `updated_at` - (String) Time stamp the target was last updated.
		  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.
	* `updated_at` - (String) Time stamp the tenant was last updated.
	  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.

