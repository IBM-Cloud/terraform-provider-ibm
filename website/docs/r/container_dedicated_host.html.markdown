---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_dedicated_host"
description: |-
  Manages dedicated host.
---

# ibm_container_dedicated_host

Provides a resource for managing a dedicated host. A dedicated host can be created or deleted. The only supported update is enabling or disabling the placement on the dedicated host. For more information about dedicated host, see [Creating and managing dedicated hosts on VPC Gen 2 infrastructure](https://cloud.ibm.com/docs/containers?topic=containers-dedicated-hosts).


## Example usage
In the following example, you can create a dedicated host:

```terraform
resource "ibm_container_dedicated_host" "test_dhost" {
  flavor            = "bx2d.host.152x608"
  host_pool_id      = "dh-abcdefgh1234567"
  zone              = "us-south-1"
  placement_enabled = "true"
}
```

## Timeouts

ibm_container_dedicated_host provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* `create` - (Default 40 minutes) Used for creating the dedicated host. Please note that after creating the host, terraform may need to execute some update logic.
* `read`   - (Default 10 minutes) Used for reading the dedicated host.
* `update` - (Default 15 minutes) Used for updating the dedicated host.
* `delete` - (Default 40 minutes) Used for deleting the dedicated host. Please note that before deleting the host, terraform may need to execute some update logic.

## Argument reference
Review the argument references that you can specify for your resource. 

- `flavor` - (Required, Forces new resource, String) The flavor of the dedicated host.
- `host_pool_id`- (Required, Forces new resource, String) The id of the dedicated host pool the dedicated host is associated with.
- `zone` - (Required, Forces new resource, String) The zone of the dedicated host.
- `placement_enabled` - (Optional, Bool) Enables/disables placement on the dedicated host.
 
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `host_id` - (String) The unique identifier of the dedicated host.
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

## Import

The `ibm_container_dedicated_host` can be imported by using the dedicated host pool id and the dedicated host id in the following format: `<dedicated host pool id>/<dedicated host id>`.

**Example**

```
$ terraform import ibm_container_dedicated_host.test_dhost dh-abcdefgh1234567/abcd12-dh-abcdefgh1234567-abcd123-acbd1234
