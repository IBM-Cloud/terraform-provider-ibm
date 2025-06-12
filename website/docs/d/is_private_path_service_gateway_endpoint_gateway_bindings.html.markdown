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
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name = "example-subnet"
  vpc = ibm_is_vpc.example.id
  zone = "us-south-2"
  ipv4_cidr_block = "10.240.0.0/24"
}
resource "ibm_is_lb" "example" {
  name = "example-lb"
  subnets = [ibm_is_subnet.example.id]
}
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "review"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.example.id
  zonal_affinity = true
  service_endpoints = ["example-fqdn"]
}
data "ibm_is_private_path_service_gateway_endpoint_gateway_bindings" "example" {
	status = "pending"
	private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `private_path_service_gateway` - (Required, String) The private path service gateway identifier.
- `status` - (Optional, String) Status of the binding
- `account` - (Optional, String) ID of the account to filter

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the PrivatePathServiceGatewayEndpointGatewayBindingCollection.
- `endpoint_gateway_bindings` - (List) Collection of endpoint gateway bindings.
	Nested scheme for **endpoint_gateway_bindings**:
	- `account` - (List) The account that created the endpoint gateway.
		Nested scheme for **account**:
		- `id` - (String)
		- `resource_type` - (String) The resource type.
	- `created_at` - (String) The date and time that the endpoint gateway binding was created.
	- `expiration_at` - (String) The expiration date and time for the endpoint gateway binding.
	- `href` - (String) The URL for this endpoint gateway binding.
	- `id` - (String) The unique identifier for this endpoint gateway binding.
	- `lifecycle_state` - (String) The lifecycle state of the endpoint gateway binding.
	- `resource_type` - (String) The resource type.
	- `status` - (String) The status of the endpoint gateway binding- `denied`: endpoint gateway binding was denied- `expired`: endpoint gateway binding has expired- `pending`: endpoint gateway binding is awaiting review- `permitted`: endpoint gateway binding was permittedThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	- `updated_at` - (String) The date and time that the endpoint gateway binding was updated.
