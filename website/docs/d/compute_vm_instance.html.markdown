---
layout: "ibm"
page_title: "IBM: ibm_vm_instance"
sidebar_current: "docs-ibm-datasource-compute-vm_instance"
description: |-
  Get information on a IBM Compute VM Instance resource
---

# ibm\_compute_vm_instance

Import the details of an existing VM instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_compute_vm_instance" "vm_instance" {
  hostname    = "jumpbox"
  domain      = "example.com"
  most_recent = true
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (Required, string) The hostname of the VM instance.
* `domain` - (Required, string) The domain of the VM instance.
* `most_recent` - (Optional, boolean) If there are multiple VM instances, you can set this argument to `true` to import only the most recently created instance.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VM instance.
* `datacenter` - The data center in which the VM instance is deployed.
* `public_interface_id` - The ID of the primary public interface.
* `private_interface_id` - The ID of the primary private interface.
* `cores` - The number of CPU cores.
* `status` - The VSI status.
* `last_known_power_state` - The last known power state of a VM instance, in the event the instance is turned off outside the information management system (IMS) or has gone offline.
* `power_state` - The current power state of a VM instance.
* `ipv4_address` - The public IPv4 address of the VM instance.
* `ip_address_id_private` - The unique identifier for the private IPv4 address assigned to the VM instance.
* `ipv4_address_private` - The private IPv4 address of the VM instance.
* `ip_address_id` - The unique identifier for the public IPv4 address assigned to the VM instance.
* `ipv6_address` - The public IPv6 address of the VM instance provided when `ipv6_enabled` is set to `true`.
* `ipv6_address_id` - The unique identifier for the public IPv6 address assigned to the VM instance provided when `ipv6_enabled` is set to `true`.
* `private_subnet_id` - The unique identifier of the subnet `ipv4_address_private` belongs to.
* `public_ipv6_subnet` - The public IPv6 subnet provided when `ipv6_enabled` is set to `true`.
* `public_ipv6_subnet_id` - The unique identifier of the subnet `ipv6_address` belongs to.
* `public_subnet_id` - The unique identifier of the subnet `ipv4_address` belongs to.
* `secondary_ip_addresses` - The public secondary IPv4 addresses of the VM instance.
* `secondary_ip_count` - Number of secondary public IPv4 addresses.