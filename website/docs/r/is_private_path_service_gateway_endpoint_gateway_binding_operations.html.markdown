---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations"
description: |-
  Manages PrivatePathServiceGateway endpoint gateway bindings.
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations

Provides a resource for ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations. This allows permitting or denying endpoint gateway bindings.

## Example Usage. Permit all the pending endpoint gateway bindings

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "review"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.testacc_LB.id
  zonal_affinity = true
  service_endpoints = ["myexamplefqdn"]
}
data "ibm_is_private_path_service_gateway_endpoint_gateway_bindings" "bindings" {
  account = "7f75c7b025e54bc5635f754b2f888665"
  status = "pending"
  private_path_service_gateway = ibm_is_private_path_service_gateway.ppsg.id
}
resource "ibm_is_private_path_service_gateway_endpoint_gateway_binding_operations" "policy" {
  count = length(data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.bindings.endpoint_gateway_bindings)
  access_policy = "permit"
  endpoint_gateway_binding = data.ibm_is_private_path_service_gateway_endpoint_gateway_bindings.bindings.endpoint_gateway_bindings[count.index].id
  private_path_service_gateway = ibm_is_private_path_service_gateway.ppsg.id
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `access_policy` - (Required, String) The access policy for the endpoint gateway binding:- permit: access will be permitted- deny:  access will be denied. Allowable values are: `deny`, `permit`. 
- `private_path_service_gateway` - (Required, Forces new resource, String) The private path service gateway 
identifier.
- `endpoint_gateway_binding` - (Required, Forces new resource, String) ID of the endpoint gateway binding

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `access_policy` - (String) The access policy for the endpoint gateway binding:- permit: access will be permitted- deny:  access will be denied. Allowable values are: `deny`, `permit`. 
- `private_path_service_gateway` - (String) The private path service gateway 
identifier.
- `endpoint_gateway_binding` - (String) ID of the endpoint gateway binding

