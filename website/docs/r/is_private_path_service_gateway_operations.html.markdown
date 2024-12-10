---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_operations"
description: |-
  Manages PrivatePathServiceGateway publish and unpublish.
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_operations

Provides a resource for ibm_is_private_path_service_gateway_operations. This allows publishing or unpublishing the PPSG.

## Example Usage. Publish a PPSG.

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "permit"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.testacc_LB.id
  zonal_affinity = true
  service_endpoints = ["myexamplefqdn"]
}
resource "ibm_is_private_path_service_gateway_operations" "publish" {
  published = true
  private_path_service_gateway = ibm_is_private_path_service_gateway.ppsg.id
}
```
## Example Usage. Unpublish a PPSG.

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "permit"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.testacc_LB.id
  zonal_affinity = true
  service_endpoints = ["myexamplefqdn"]
}
resource "ibm_is_private_path_service_gateway_operations" "publish" {
  published = false
  private_path_service_gateway = ibm_is_private_path_service_gateway.ppsg.id
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `published` - (Required, Boolean) Boolean to specify whether to publish or unpublish the PPSG.
- `private_path_service_gateway` - (Required, Forces new resource, String) The private path service gateway 
identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `published` - (Boolean) Boolean to specify whether to publish or unpublish the PPSG.
- `private_path_service_gateway` - (String) The private path service gateway 
identifier.

