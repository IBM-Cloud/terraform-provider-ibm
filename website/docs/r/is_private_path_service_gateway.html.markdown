---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway"
description: |-
  Manages PrivatePathServiceGateway.
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway

Provides a resource for PrivatePathServiceGateway. This allows PrivatePathServiceGateway to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_is_private_path_service_gateway" "example" {
  default_access_policy = "permit"
  name = "my-example-ppsg"
  load_balancer = ibm_is_lb.testacc_LB.id
  zonal_affinity = true
  service_endpoints = ["myexamplefqdn"]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `default_access_policy` - (Optional, String) The policy to use for bindings from accounts without an explicit account policy. `permit`: access will be permitted. `deny`:  access will be denied. `review`: access will be manually reviewed. Allowable values are: `deny`, `permit`, `review`. 
- `load_balancer` - (Required, String) The ID of the load balancer for this private path service gateway. This load balancer must be in the same VPC as the private path service gateway and must have is_private_path set to true.
- `service_endpoints` - (Required, List of Strings) The fully qualified domain names for this private path service gateway.
- `name` - (Optional, String) The name for this private path service gateway. The name must not be used by another private path service gateway in the VPC. 
- `resource_group` - (Optional, String) ID of the resource group to use.
- `zonal_affinity` - (Optional, String) Indicates whether this private path service gateway has zonal affinity.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.


- `created_at` - (String) The date and time that the private path service gateway was created.
- `crn` - (String) The CRN for this private path service gateway.
- `default_access_policy` - (String) The policy to use for bindings from accounts without an explicit account policy.
- `endpoint_gateway_count` - (Integer) The number of endpoint gateways using this private path service gateway.
- `endpoint_gateway_binding_auto_delete` - (Boolean) Indicates whether endpoint gateway bindings will be automatically deleted after endpoint_gateway_binding_auto_delete_timeout hours have passed. At present, this is always true, but may be modifiable in the future.
- `endpoint_gateway_binding_auto_delete_timeout` - (Integer) If endpoint_gateway_binding_auto_delete is true, the hours after which endpoint gateway bindings will be automatically deleted. If the value is 0, abandoned endpoint gateway bindings will be deleted immediately. At present, this is always set to 0. This value may be modifiable in the future.
- `href` - (String) The URL for this private path service gateway.
- `id` - The unique identifier of the PrivatePathServiceGateway
- `lifecycle_state` - (String) The lifecycle state of the private path service gateway.
- `load_balancer` - (String) The load balancer for this private path service gateway.
- `name` - (String) The name for this private path service gateway. The name is unique across all private path service gateways in the VPC.
- `published` - (Boolean) Indicates the availability of this private path service gateway- `true`: Any account can request access to this private path service gateway.- `false`: Access is restricted to the account that created this private path service gateway.
- `resource_group` - (String) The resource group for this private path service gateway.
- `resource_type` - (String) The resource type.
- `service_endpoints` - (List of strings) The fully qualified domain names for this private path service gateway.
- `vpc` - (String) The VPC this private path service gateway resides in.
- `zonal_affinity` - (Boolean) Indicates whether this private path service gateway has zonal affinity.- `true`:  Traffic to the service from a zone will favor service endpoints in the same zone.- `false`: Traffic to the service from a zone will be load balanced across all zones           in the region the service resides in.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_private_path_service_gateway` resource by using `id`.
The `id` property can be formed using the private_path_service_gateway id. For example:

```terraform
import {
  to = ibm_is_private_path_service_gateway.example
  id = "<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_private_path_service_gateway.example <id>
```