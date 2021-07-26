---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_vm_instance"
description: |-
  Get information on a IBM Cloud compute Virtual Machine instance resource
---

# ibm_compute_vm_instance
Retrieve information of an existing Virtual Machine (VM) instance as a read-only data source. For more information, about computer VM instance, see [enabling auto scale for better capacity and resiliency](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-auto-scale).

## Example usage

```terraform
data "ibm_compute_vm_instance" "vm_instance" {
  hostname    = "jumpbox"
  domain      = "example.com"
  most_recent = true
}
```


## Argument reference
Review the argument references that you can specify for your data source.

- `domain` - (Required, String) The domain of the VM instance.
- `hostname` - (Required, String) The hostname of the VM instance.
- `most_recent` - (Optional, Bool) For multiple VM instances, you can set this argument to **true** to import only the most recently created instance.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `cores`- (Integer) The number of CPU cores.
- `datacenter` - (String) The data center in which the VM instance is deployed.
- `id` - (String) The unique identifier of the VM instance.
- `ipv4_address` - (String) The public IPv4 address of the VM instance.
- `ip_address_id_private` - (String) The unique identifier for the private IPv4 address that is assigned to the VM instance.
- `ipv4_address_private` - (String) The private IPv4 address of the VM instance.
- `ip_address_id` - (String) The unique identifier for the public IPv4 address that is assigned to the VM instance.
- `ipv6_address` - (String) The public IPv6 address of the VM instance provided when `ipv6_enabled` is set to **true**.
- `ipv6_address_id` - (String) The unique identifier for the public IPv6 address assigned to the VM instance provided when `ipv6_enabled` is set to **true**.
- `last_known_power_state` - (String) The last known power state of a VM instance, if the instance is turned off outside the information management system (IMS) is offline.
- `power_state` - (String) The current power state of a VM instance.
- `private_interface_id` - (String) The ID of the primary private interface.
- `private_subnet_id` - (String) The unique identifier of the subnet `ipv4_address_private` belongs to.
- `public_ipv6_subnet` - (String) The public IPv6 subnet provided when `ipv6_enabled` is set to **true**.
- `public_ipv6_subnet_id` - (String) The unique identifier of the subnet `ipv6_address` belongs to.
- `public_subnet_id` - (String) The unique identifier of the subnet `ipv4_address` belongs to.
- `public_interface_id` - (String) The ID of the primary public interface.
- `secondary_ip_addresses` - (String) The public secondary IPv4 addresses of the VM instance.
- `secondary_ip_count`- (Integer) Number of secondary public IPv4 addresses.
- `status` - (String) The VSI status.
