---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: compute_monitor"
description: |-
  Manages IBM Cloud compute monitor resources.
---


# ibm_compute_monitor
Create, update, or delete a monitor for your virtual server instance. With monitors, you can verify the health of your virtual server instance by sending ping requests to the instance and checking the responsiveness of your instance. For more information, about compute monitor resource, see [viewing and managing monitors](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-virtual-server-instances).

**Note**

For more information, see [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Monitor_Version1_Query_Host).

## Example usage

In the following example, you can create a monitor:

```terraform
resource "ibm_compute_monitor" "test_monitor" {
    guest_id = ibm_compute_vm_instance.test_server.id
    ip_address = ibm_compute_vm_instance.test_server.id.ipv4_address
    query_type_id = 1
    response_action_id = 1
    wait_cycles = 5
    notified_users = [460547]
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `guest_id` - (Required, Forces new resource, Integer) The ID of the virtual guest that you want to monitor.
- `ip_address`- (Optional, String) The IP address that you want to monitor.
- `notified_users`- (Optional, Array of Integers) The list of user IDs that is notified.
- `query_type_id` - (Required, Integer) The ID of the query type.
- `response_action_id`- (Required, Integer) The ID of the response action to take if the monitor fails. Accepted values are `1` or `2`.
- `tags`- (Optional, Array of Integers) Tags associated with the monitoring instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `wait_cycles` - (Optional, Integer)The number of 5-minute cycles to wait before the response action is taken.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the monitor.
- `notified_users`- (String) The list of user IDs that is notified.
