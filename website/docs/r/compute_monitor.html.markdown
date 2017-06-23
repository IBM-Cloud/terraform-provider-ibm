---
layout: "ibm"
page_title: "IBM: compute_monitor"
sidebar_current: "docs-ibm-resource-compute-monitor"
description: |-
  Manages IBM Compute monitor resources.
---


# ibm\_compute_monitor

Provides a resource to create, update, and delete a monitoring instance.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Monitor_Version1_Query_Host).

## Example Usage

```hcl
# Create a monitor
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

* `guest_id` - (Required, integer) The ID of the virtual guest to be monitored.
* `ip_address` - (Optional, strings) The IP address to be monitored.
* `query_type_id` - (Required, integer) The ID of the query type.
* `response_action_id` - (Required, integer) The ID of the response action to take if the monitor fails. Accepted values are `1` or `2`.
* `wait_cycles` - (Optional, integer) The number of five-minute cycles to wait before the response action is taken.
* `notified_users` - (Optional, array of integers) The list of user IDs to be notified.

## Attributes Reference

The following attributes are exported:

* `id` - ID of the monitor.
* `notified_users` - The list of user IDs to be notified.
