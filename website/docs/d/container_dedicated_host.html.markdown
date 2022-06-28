---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_dedicated_host"
description: |-
  Get information about a dedicated host.
---

# ibm_container_dedicated_host

Retrieve information about a dedicated host. For more information about dedicated hosts, see [Creating and managing dedicated hosts on VPC Gen 2 infrastructure](https://cloud.ibm.com/docs/containers?topic=containers-dedicated-hosts).


## Example usage
In the following example, you can retrieve a dedicated host:

```terraform
data "ibm_container_dedicated_host" "test_dhost" {
  host_id      = "abcd12-dh-abcdefgh1234567-abcd123-acbd1234"
  host_pool_id = "dh-abcdefgh1234567"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 
- `host_id` - (Required, String) The unique identifier of the dedicated host.
- `host_pool_id` - (Required, String) The unique identifier of the dedicated host pool the dedicated host is associated with.
 
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created.
- `flavor` - (String) The name of the dedicated host.
- `placement_enabled`- (Bool) Describes if the placement on the dedicated host is enabled
- `life_cycle` - (List) A nested block describes the lifecycle state of the dedicated host.

  Nested scheme for `life_cycle`:
  - `actual_state` - (String) The actual state of the dedicated host.
  - `desired_state` - (String) The desired state of the dedicated host.
  - `message` - (String) Information message about the dedicated host's lifecycle.
  - `message_date` - (String) The date of the information message.
  - `message_details` - (String) Additional details of the information message.
  - `message_details_date` - (String) The date of the additional details.
- `resources` - (List) A nested block describes the resources of the dedicated host.

  Nested scheme for `resources`:
  - `capacity` - (List) A nested block describes the capacity of the dedicated host.
    Nested scheme for `capacity`:
    - `memory_bytes` - (Int) Memory capacity of the dedicated host.
    - `vcpu` - (Int) VCPU capacity of the dedicated host.
  - `consumed` - (List) A nested block describes the consumed resources of the dedicated host.
    Nested scheme for `capacity`:
    - `memory_bytes` - (Int) Consumed memory capacity of the dedicated host.
    - `vcpu` - (Int) Consumed VCPU capacity of the dedicated host.
- `workers` - (List) A nested block describes the workers associated with this dedicated host.

  Nested scheme for `workers`:
  - `cluster_id` - (String) The ID of the cluster the worker is associated with.
  - `flavor` - (String) The flavor of the worker.
  - `worker_id` - (String) The ID of the worker.
  - `worker_pool_id` -  (String) The ID of the worker pool the worker is associated with.
- `zone` - (String) The zone of the dedicated host.
