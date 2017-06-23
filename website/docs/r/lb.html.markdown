---
layout: "ibm"
page_title: "IBM : lb"
sidebar_current: "docs-ibm-resource-lb"
description: |-
  Manages IBM Load Balancer.
---

# ibm\_lb

Provides a resource for local load balancers. This allows local load balancers to be created, updated, and deleted.

## Example Usage

```hcl
# Create a local load balancer
resource "ibm_lb" "test_lb_local" {
    connections = 1500
    datacenter = "tok02"
    ha_enabled = false
    dedicated = false       
}
```

## Argument Reference

The following arguments are supported:

* `connections` - (Required, integer) Set the number of connections for the local load balancer.
* `datacenter` - (Required, string) Set the data center for the local load balancer.
* `ha_enabled` - (Required, boolean) Set whether the local load balancer needs to be HA enabled or not.
* `security_certificate_id` - (Optional, integer) Set the ID of the security certificate associated with the local load balancer.
* `dedicated` - (Optional, boolean) Set to `true` if the local load balancer should be dedicated. Default value: `false`.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the local load balancer.
* `ip_address` - The IP Address of the local load balancer.
* `subnet_id` - The ID of the subnet associated with the local load balancer.
* `ssl_enabled` - If the local load balancer provides SSL capability or not.
