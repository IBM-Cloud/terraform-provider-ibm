---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : reserved_ip"
description: |-
  Shows the info for a reserved IP and Subnet.
---

# ibm\_is_subnet_reserved_ip

Import the details of an existing Reserved IP in a Subnet as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_subnet_reserved_ip" "data_reserved_ip" {
  subnet = ibm_is_subnet.test_subnet.id
  reserved_ip = ibm_is_subnet_reserved_ip.resource_res_ip.reserved_ip
}
```

## Argument Reference

The following arguments are supported as inputs/request params:

* `subnet` - (Required, string) The id for the Subnet.
* `reserved_ip` - (Required, string) The id for the Reserved IP.


## Attribute Reference

The following attributes are exported as output/response:

* `auto_delete` - The auto_delete boolean for reserved IP
* `created_at` - The creation timestamp for the reserved IP
* `href` - The unique reference for the reserved IP
* `id` - The id for the reserved IP
* `name` - The name for the reserved IP
* `owner` - The owner of the reserved IP
* `reserved_ip` - Same as `id`
* `resource_type` - The type of resource
* `subnet` - The id for the subnet for the reserved IP
