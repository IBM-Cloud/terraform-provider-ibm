---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_revoke_account"
description: |-
  Manages PrivatePathServiceGateway revoke account.
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_revoke_account

Provides a resource for ibm_is_private_path_service_gateway_revoke_account. This revokes the access to provided account.

## Example Usage.
```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "permit"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.testacc_LB.id
  zonal_affinity = true
  service_endpoints = ["myexamplefqdn"]
}
 resource "ibm_is_private_path_service_gateway_revoke_account" "example" {
  account = "7f75c7b025e54bc5635f754b2f888665"
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `account` - (Required, String) Account ID to revoke.
- `private_path_service_gateway` - (Required, Forces new resource, String) The private path service gateway 
identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `account` - (Required, String) Account ID to revoke.
- `private_path_service_gateway` - (String) The private path service gateway 
identifier.

