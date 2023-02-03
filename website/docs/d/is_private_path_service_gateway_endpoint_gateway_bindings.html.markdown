---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_endpoint_gateway_bindings"
description: |-
  Get information about PrivatePathServiceGatewayEndpointGatewayBindingCollection
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_endpoint_gateway_bindings

Provides a read-only data source for PrivatePathServiceGatewayEndpointGatewayBindingCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_private_path_service_gateway_endpoint_gateway_bindings" "is_private_path_service_gateway_endpoint_gateway_bindings" {
	private_path_service_gateway_id = "private_path_service_gateway_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `private_path_service_gateway_id` - (Required, Forces new resource, String) The private path service gateway identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the PrivatePathServiceGatewayEndpointGatewayBindingCollection.
* `endpoint_gateway_bindings` - (List) Collection of endpoint gateway bindings.
Nested scheme for **endpoint_gateway_bindings**:
	* `account` - (List) The account that created the endpoint gateway.
	Nested scheme for **account**:
		* `id` - (String)
		  * Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
		* `resource_type` - (String) The resource type.
		  * Constraints: Allowable values are: `account`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `created_at` - (String) The date and time that the endpoint gateway binding was created.
	* `expiration_at` - (String) The expiration date and time for the endpoint gateway binding.
	* `href` - (String) The URL for this endpoint gateway binding.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this endpoint gateway binding.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `lifecycle_state` - (String) The lifecycle state of the endpoint gateway binding.
	  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `private_path_service_gateway_endpoint_gateway_binding`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `status` - (String) The status of the endpoint gateway binding- `denied`: endpoint gateway binding was denied- `expired`: endpoint gateway binding has expired- `pending`: endpoint gateway binding is awaiting review- `permitted`: endpoint gateway binding was permittedThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  * Constraints: Allowable values are: `denied`, `expired`, `pending`, `permitted`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `updated_at` - (String) The date and time that the endpoint gateway binding was updated.

* `first` - (List) A link to the first page of resources.
Nested scheme for **first**:
	* `href` - (String) The URL for a page of resources.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `limit` - (Integer) The maximum number of resources that can be returned by the request.
  * Constraints: The maximum value is `100`. The minimum value is `1`.

* `next` - (List) A link to the next page of resources. This property is present for all pagesexcept the last page.
Nested scheme for **next**:
	* `href` - (String) The URL for a page of resources.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `total_count` - (Integer) The total number of resources across all pages.
  * Constraints: The minimum value is `0`.

