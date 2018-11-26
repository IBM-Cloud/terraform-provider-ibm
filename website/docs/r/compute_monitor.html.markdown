---
layout: "ibm"
page_title: "IBM: compute_monitor"
sidebar_current: "docs-ibm-resource-compute-monitor"
description: |-
  Manages IBM Compute monitor resources.
---


# ibm\_compute_monitor

Provides a monitoring instance resource. This allows monitoring instances to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Monitor_Version1_Query_Host).

## Example Usage

In the following example, you can create a monitor:

```hcl
resource "ibm_compute_monitor" "test_monitor" {
    guest_id = ${ibm_compute_vm_instance.test_server.id}
    ip_address = ${ibm_compute_vm_instance.test_server.id.ipv4_address}
    query_type_id = 1
    response_action_id = 1
    wait_cycles = 5
    notified_users = [460547]
}
```

## Argument Reference

The following arguments are supported:

* `guest_id` - (Required, integer) The ID of the virtual guest you want to monitor.
* `ip_address` - (Optional, strings) The IP address you want to monitor.
* `query_type_id` - (Required, integer) The ID of the query type.
* `response_action_id` - (Required, integer) The ID of the response action to take if the monitor fails. Accepted values are `1` or `2`.
* `wait_cycles` - (Optional, integer) The number of five-minute cycles to wait before the response action is taken.
* `notified_users` - (Optional, array of integers) The list of user IDs that will be notified.
* `tags` - (Optional, array of strings) Tags associated with the monitoring instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the monitor.
* `notified_users` - The list of user IDs that will be notified.
