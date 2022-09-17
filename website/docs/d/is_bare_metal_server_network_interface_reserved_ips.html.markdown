---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : reserved_ips"
description: |-
  Lists all the info in reserved IP for Bare Metal Server network interface.
---

# ibm\_is_bare_metal_server_network_interface_reserved_ips

Import the details of all the Reserved IPs in a network interface of an bare metal server as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_bare_metal_server_network_interface_reserved_ips" "data_reserved_ips" {
  bare_metal_server = ibm_is_bare_metal_server.test_bare_metal_server.id
  network_interface = ibm_is_bare_metal_server.test_bare_metal_server.network_interfaces.0.id
}
```

## Argument Reference

The following arguments are supported as inputs/request params:

- `bare_metal_server` - (Required, string) The id for the bare metal server.
- `network_interface` - (Required, string) The id for the network interface.


## Attribute Reference

The following attributes are exported as output/response:

- `id` - The id for the all the reserved ID (current timestamp)
- `reserved_ips` - The collection of all the reserved IPs in the network interface
   - `address` - (String) The IP bound for the reserved IP
   - `auto_delete` - (Bool) If reserved ip shall be deleted automatically
   - `created_at` - (String) The date and time that the reserved IP was created
   - `href` - (String) The URL for this reserved IP
   - `reserved_ip` - (String) The unique identifier for this reserved IP
   - `name` - (String) The user-defined or system-provided name for this reserved IP
   - `owner` -(String)  The owner of a reserved IP, defining whether it is managed by the user or the provider
   - `resource_type` - (String) The resource type
   - `target` - (String) The id for the target for the reserved IP.

- `total_count` - The number of reserved IP in the network interface of the bare metal server
