---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_ip"
description: |-
  Manages IBM Virtual endpoint gateway IP.
---

# ibm_is_virtual_endpoint_gateway_ip
Create, update, or delete a VPC endpoint gateway IP by using virtual endpoint gateway resource. For more information, about the VPC endpoint gateway, see [about VPC gateways](https://cloud.ibm.com/docs/vpc?topic=vpc-about-vpe).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example creates a Virtual Private Endpoint gateway IP.

```terraform
resource "ibm_is_virtual_endpoint_gateway_ip" "example" {
  gateway     = ibm_is_virtual_endpoint_gateway.example.id
  reserved_ip = ibm_is_subnet_reserved_ip.example.reserved_ip
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `gateway` - (Required, Forces new resource, String) The endpoint gateway ID.
- `reserver_ip` - (Required, Forces new resource, String) The endpoint gateway IP ID.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `address` - (String) The endpoint gateway IP address.
- `auto_delete` - (String) The endpoint gateway IP auto delete.
- `created_at` - (Timestamp) The created date and time of the endpoint gateway IP.
- `id` - (String) The unique identifier of the VPE gateway. The ID is composed of `<gateway_id>/<gateway_ip_id>`.
- `name` - (String) The endpoint gateway IP name.
- `resource_type` - (String) The endpoint gateway IP resource type.
- `target` - (List) The endpoint gateway target details.

  Nested scheme for `target`:
  - `id` - (String) The IPs target ID.
  - `name` - (String) The IPs target name.
  - `resource_type` - (String) The endpoint gateway resource type.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_virtual_endpoint_gateway_ip` resource by using `id`.
The `id` property can be formed from `virtual endpoint gateway ID`, and `gateway IP ID`. For example:

```terraform
import {
  to = ibm_is_virtual_endpoint_gateway_ip.example
  id = "<virtual_endpoint_gateway_id>/<gateway_ip_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_virtual_endpoint_gateway_ip.example <virtual_endpoint_gateway_id>/<gateway_ip_id>
```