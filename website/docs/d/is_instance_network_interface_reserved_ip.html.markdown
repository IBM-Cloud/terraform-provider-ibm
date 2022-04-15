---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : reserved_ip"
description: |-
  Shows the info for a reserved IP and instance network interface.
---

# ibm\_is_instance_network_interface_reserved_ip

Import the details of an existing Reserved IP in a network interface of an instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_instance_network_interface_reserved_ip" "data_reserved_ip" {
  instance = ibm_is_instance.test_instance.id
  network_interface = ibm_is_instance.test_instance.network_interfaces.0.id
  reserved_ip = ibm_is_instance.test_instance.network_interfaces.0.ips.0.id
}
```

## Argument Reference

The following arguments are supported as inputs/request params:

- `instance` - (Required, string) The id for the instance.
- `network_interface` - (Required, string) The id for the network interface.
- `reserved_ip` - (Required, string) The id for the Reserved IP.


## Attribute Reference

The following attributes are exported as output/response:

- `auto_delete` - (String) The auto_delete boolean for reserved IP
- `created_at` - (String) The creation timestamp for the reserved IP
- `href` - (String) The unique reference for the reserved IP
- `id` - (String) The id for the reserved IP
- `name` - (String) The name for the reserved IP
- `owner` - (String) The owner of the reserved IP
- `reserved_ip` - (String) Same as `id`
- `resource_type` - (String) The type of resource
- `target` - (String) The id for the target for the reserved IP
