---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_resource_binding"
description: |-
  Manages EndpointGatewayResourceBinding.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_virtual_endpoint_gateway_resource_binding

Create, update, and delete EndpointGatewayResourceBindings with this resource.

## Example Usage

```hcl
resource "ibm_is_virtual_endpoint_gateway_resource_binding" "is_virtual_endpoint_gateway_resource_binding_instance" {
  endpoint_gateway_id = "endpoint_gateway_id"
  name = "my-resource-binding"
  target {
		crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/aa2432b1fa4d4ace891e9b80fc104e34:1a0ec336-f391-4091-a6fb-5e084a4c56f4:bucket:bucket-27200-lwx4cfvcue"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `endpoint_gateway_id` - (Required, Forces new resource, String) The endpoint gateway identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `name` - (Optional, String) The name for this resource binding. The name is unique across all resource bindings for the endpoint gateway.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `target` - (Required, List) The target for this endpoint gateway resource binding.
Nested schema for **target**:
	* `crn` - (Optional, String)
	  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the EndpointGatewayResourceBinding.
* `created_at` - (String) The date and time that the resource binding was created.
* `href` - (String) The URL for this endpoint gateway resource binding.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `endpoint_gateway_resource_binding_id` - (String) The unique identifier for this endpoint gateway resource binding.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	* `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this lifecycle state.
	* `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the resource binding.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `endpoint_gateway_resource_binding`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `service_endpoint` - (String) The fully qualified domain name of the service endpoint for the resource targeted by this resource binding.
  * Constraints: The maximum length is `255` characters. The minimum length is `4` characters. The value must match regular expression `/^((?=[A-Za-z0-9-]{1,63}\\.)[A-Za-z0-9-]*\\.)+[A-Za-z]{2,63}\\.?$/`.
* `type` - (String) The type of resource binding:- `weak`: The binding is not dependent on the existence of the target resource.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `weak`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.


## Import

You can import the `ibm_is_virtual_endpoint_gateway_resource_binding` resource by using `id`.
The `id` property can be formed from `endpoint_gateway_id`, and `endpoint_gateway_resource_binding_id` in the following format:

<pre>
&lt;endpoint_gateway_id&gt;/&lt;endpoint_gateway_resource_binding_id&gt;
</pre>
* `endpoint_gateway_id`: A string. The endpoint gateway identifier.
* `endpoint_gateway_resource_binding_id`: A string in the format `r006-a7ba95b6-a254-47e4-b129-10593df8a373`. The unique identifier for this endpoint gateway resource binding.

# Syntax
<pre>
$ terraform import ibm_is_virtual_endpoint_gateway_resource_binding.is_virtual_endpoint_gateway_resource_binding &lt;endpoint_gateway_id&gt;/&lt;endpoint_gateway_resource_binding_id&gt;
</pre>
