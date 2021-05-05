---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : floating_ip"
description: |-
  Fechtes floating ip information.
---

# ibm\_floating\_ip

Import the details of vpc floating ip on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

    data "ibm_is_floating_ip" "test" {
        name   = "test-fp"
    }

```
## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the floating ip.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:
* `id` - An alphanumeric value identifying the floating ip.	
* `address` - Floating IP address that was created.
* `status` - Provisioning status of the floating IP address. 
* `tags` - Tags associate with VPC.
* `target` -  ID of the network interface use to allocate the IP address.
* `zone` -   Name of the zone where to create the floating IP address. 