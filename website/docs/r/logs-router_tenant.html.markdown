---
layout: "ibm"
page_title: "IBM : ibm_logs_router_tenant"
description: |-
  Manages logs_router_tenant.
subcategory: "IBM Cloud Logs Routing"
---

# ibm_logs_router_tenant

Create, update, and delete logs_router_tenants with this resource.

## Example Usage

```hcl
resource "ibm_resource_instance" "logs_instance" {
  name     = "logs-instance"
  service  = "logs"
  plan     = "standard"
  location = "eu-de"
  parameters = {
    retention_period        = "14"
    logs_bucket_crn         = "crn:v1:bluemix:public:cloud-object-storage:global:a/4448261269a14562b839e0a3019ed980:f8b3176e-af8e-4e14-a2f9-7f82634e7f0b:bucket:logs-bucket"
    logs_bucket_endpoint    = "s3.direct.eu-de.cloud-object-storage.appdomain.cloud"
    metrics_bucket_crn      = "crn:v1:bluemix:public:cloud-object-storage:global:a/4448261269a14562b839e0a3019ed980:f8b3176e-af8e-4e14-a2f9-7f82634e7f0b:bucket:metrics-bucket"
    metrics_bucket_endpoint = "s3.direct.eu-de.cloud-object-storage.appdomain.cloud"
  }
}

resource "ibm_logs_router_tenant" "logs_router_tenant_instance" {
  name = "cloud-logs-router-tenant"
  region = "eu-de"
  targets {
		log_sink_crn = ibm_resource_instance.logs_instance.target_crn
		name = "my-cloud-logs-target"
		parameters {
			host = ibm_resource_instance.logs_instance.extensions.external_ingress_private
			port = 443
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, String) The name for this tenant. The name is regionally unique across all tenants in the account.
  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
* `region` - (Required, Forces new resource, String) The region to create the tenant.
  * Constraints: The value must match one of the available regions. For a list of regions, see the available [IBM Cloud Logs Router Endpoints](https://cloud.ibm.com/docs/logs-router?topic=logs-router-locations).
* `targets` - (Required, List) List of targets.
  * Constraints: The maximum length is `2` items. The minimum length is `1` item.
Nested schema for **targets**:
	* `log_sink_crn` - (Required, String) Cloud resource name of the log-sink target instance.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,:,-]/`.
	* `name` - (Optional, String) The name for this tenant target. The name is unique across all targets for this tenant.
	  * Constraints: The maximum length is `35` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
	* `parameters` - (Required, List) List of properties returned from a successful list operation for a log-sink of type IBM Log Analysis (logdna).
	Nested schema for **parameters**:
		* `host` - (Required, String) Host name of the log-sink.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,-,.]/`.
		* `port` - (Required, Integer) Network port of the log-sink.
		  * Constraints: The maximum value is `65535`. The minimum value is `1`.
		* `access_credential` - (Optional, String) Secret to connect to the Mezmo log-sink. This is not required for log-sink of type Cloud Logs.


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_router_tenant.
* `created_at` - (String) Time stamp the tenant was originally created.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.
* `crn` - (String) Cloud resource name of the tenant.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/[a-z,A-Z,0-9,:,-]/`.
* `etag` - (String) Resource version identifier.
  * Constraints: The maximum length is `66` characters. The minimum length is `66` characters. The value must match regular expression `/(?:W\/)?"(?:[ !#-\\x7E\\x80-\\xFF]*|  [  ]|\\.)*"/`.
* `updated_at` - (String) Time stamp the tenant was last updated.
  * Constraints: The maximum length is `36` characters. The minimum length is `1` character. The value must match regular expression `/[0-9,:,.,-,T,Z]/`.

* `etag` - ETag identifier for logs_router_tenant.

* `target.0.etag` - ETag identifier for logs_router_tenant target.

* `target.0.id` -  The unique identifier of the logs_router_tenant target

## Import

You can import the `ibm_logs_router_tenant` resource by using `id`, the unique `id` of the tenant, and `region` where the tenant exists in the format `id/region`.

# Syntax
<pre>
$ terraform import ibm_logs_router_tenant.logs_router_tenant &lt;id/region&gt;
</pre>

# Example
```
$ terraform import ibm_logs_router_tenant.logs_router_tenant 8717db99-2cfb-4ba6-a033-89c994c2e9f0/us-east
```
