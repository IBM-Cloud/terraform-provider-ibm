---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway_account_policy"
description: |-
  Manages PrivatePathServiceGatewayAccountPolicy.
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway_account_policy

Provides a resource for PrivatePathServiceGatewayAccountPolicy. This allows PrivatePathServiceGatewayAccountPolicy to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "deny"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.testacc_LB.id
  zonal_affinity = true
  service_endpoints = ["myexamplefqdn"]
}
resource "ibm_is_private_path_service_gateway_account_policy" "example" {
  access_policy = "deny"
  account = "fee82deba12e4c0fb69c3b09d1f12345"
  private_path_service_gateway = ibm_is_private_path_service_gateway.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `access_policy` - (Required, String) The access policy for the account:- permit: access will be permitted- deny:  access will be denied- review: access will be manually reviewed. Allowable values are: `deny`, `permit`, `review`. 
- `account` - (Required, Forces new resource, String) The ID of the account for this access policy.
- `private_path_service_gateway` - (Required, Forces new resource, String) The private path service gateway identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the PrivatePathServiceGatewayAccountPolicy. The ID is composed of `<private_path_service_gateway_id>/<account_policy_id>`.
- `created_at` - (String) The date and time that the account policy was created.
- `href` - (String) The URL for this account policy.
- `account_policy` - (String) The unique identifier for this account policy.
- `resource_type` - (String) The resource type.
- `updated_at` - (String) The date and time that the account policy was updated.


## Import

You can import the `ibm_is_private_path_service_gateway_account_policy` resource by using `id` in the following format `<private_path_service_gateway_id>/<id>`

- `private_path_service_gateway`: A string. The private path service gateway identifier.
- `id`: A string. The account policy identifier.

# Syntax
```
$ terraform import ibm_is_private_path_service_gateway_account_policy.example 0767-fa41aecb-4f21-423d-8082-630bfba1e1d9/d7bec597-4726-451f-8a63-e62e6f19c32c
```
