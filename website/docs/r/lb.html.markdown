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

In the following example, you can create a local load balancer:

```hcl
resource "ibm_lb" "test_lb_local" {
    connections = 1500
    datacenter  = "tok02"
    ha_enabled  = false
    dedicated   = false       
}
```

## Argument Reference

The following arguments are supported:

* `connections` - (Required, integer) The number of connections for the local load balancer. Only incremental upgrade is supported . For downgrade, please open the softlayer support ticket.
* `datacenter` - (Required, string) The data center for the local load balancer.
* `ha_enabled` - (Required, boolean) Specifies whether the local load balancer must be HA-enabled.
* `security_certificate_id` - (Optional, integer) The ID of the security certificate associated with the local load balancer.
* `dedicated` - (Optional, boolean) Specifies whether the local load balancer must be dedicated. The default value is `false`.
* `tags` - (Optional, array of strings) Tags associated with the local load balancer instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the local load balancer.
* `ip_address` - The IP address of the local load balancer.
* `subnet_id` - The unique identifier of the subnet associated with the local load balancer.
* `ssl_enabled` - The status of whether the local load balancer provides SSL capability.
