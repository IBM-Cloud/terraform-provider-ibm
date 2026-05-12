---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_resource_bindings"
description: |-
  Get information about EndpointGatewayResourceBindingCollection
subcategory: "Virtual Private Cloud API"
---

# ibm_is_virtual_endpoint_gateway_resource_bindings

Provides a read-only data source to retrieve information about an EndpointGatewayResourceBindingCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_virtual_endpoint_gateway_resource_bindings" "is_virtual_endpoint_gateway_resource_bindings" {
	endpoint_gateway_id = ibm_is_virtual_endpoint_gateway_resource_binding.is_virtual_endpoint_gateway_resource_binding_instance.endpoint_gateway_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `endpoint_gateway_id` - (Required, Forces new resource, String) The endpoint gateway identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the EndpointGatewayResourceBindingCollection.
* `resource_bindings` - (List) A page of resource bindings for the endpoint gateway.
Nested schema for **resource_bindings**:
	* `created_at` - (String) The date and time that the resource binding was created.
	* `href` - (String) The URL for this endpoint gateway resource binding.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this endpoint gateway resource binding.
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
	* `name` - (String) The name for this resource binding. The name is unique across all resource bindings for the endpoint gateway.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `endpoint_gateway_resource_binding`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `service_endpoint` - (String) The fully qualified domain name of the service endpoint for the resource targeted by this resource binding.
	  * Constraints: The maximum length is `255` characters. The minimum length is `4` characters. The value must match regular expression `/^((?=[A-Za-z0-9-]{1,63}\\.)[A-Za-z0-9-]*\\.)+[A-Za-z]{2,63}\\.?$/`.
	* `target` - (List) The target for this endpoint gateway resource binding.
	Nested schema for **target**:
		* `crn` - (String)
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
	* `type` - (String) The type of resource binding:- `weak`: The binding is not dependent on the existence of the target resource.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `weak`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

