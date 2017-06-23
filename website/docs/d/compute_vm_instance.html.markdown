---
layout: "ibm"
page_title: "IBM: ibm_vm_instance"
sidebar_current: "docs-ibm-datasource-compute-vm_instance"
description: |-
  Get information on a IBM Compute VM Instance resource
---

# ibm\_compute_vm_instance

Import the details of an existing VM instance as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.

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

* `hostname` - (Required) The hostname of the VM instance.
* `domain` - (Required) The domain of the VM instance.
* `most_recent` - (Optional) `True` or `False`. If `true` and multiple entries are found, the most recently created VM instance is used. If `false`, an error is returned.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the vm instance.
* `datacenter` - Data center in which the VM instance is deployed.
* `cores` - Number of CPU cores.
* `status` - The VSI status.
* `last_known_power_state` - The last known power state of a VM instance, in the event the instance is turned off outside of IMS or has gone offline.
* `power_state` - The current power state of a VM instance.
