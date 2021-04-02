---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_virtual_endpoint_gateway_ips"
description: |-
  Manages IBM Cloud Infrastructure virtual endpoint gateway ips.
---

# ibm_is_virtual_endpoint_gateway_ips

Import the details of an existing IBM Cloud Infrastructure virtual endpoint gateway ips as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_virtual_endpoint_gateway_ips" "data_test1" {
  gateway = ibm_is_virtual_endpoint_gateway.endpoint_gateway.id
}
```

## Argument Reference

The following arguments are supported:

- `gateway`(Required,string)- Endpoint gateway ID

## Attribute Reference

The following attributes are exported:

- `ips` - Endpoint gateway reserved IP id
- `reserved_ip` - Endpoint gateway IP id
- `name` - Endpoint gateway IP name
- `resource_type` - Endpoint gateway IP resource type
- `created_at` - Endpoint gateway created date and time
- `auto_delete` - Endpoint gateway IP auto delete
- `address` - Endpoint gateway IP address
- `target` - Endpoint gateway detail
  - `id` - The IPs target id
  - `name` - The IPs target name
  - `resource_type` - Endpoint gateway resource type
