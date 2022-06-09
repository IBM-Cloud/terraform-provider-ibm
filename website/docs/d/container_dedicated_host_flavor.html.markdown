---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_dedicated_host_flavor"
description: |-
  Get information about a dedicated host flavor.
---

# ibm_container_dedicated_host_flavor

Retrieve information about a dedicated host flavor. For more information, about the use of dedicated host flavors, see [Creating a cluster on dedicated host infrastructure](https://cloud.ibm.com/docs/containers?topic=containers-clusters#cluster_dedicated_host_cli).


## Example usage
In the following example, you can retrieve a dedicated host flavor:

```terraform
data "ibm_container_dedicated_host_flavor" "test_dhost_flavor" {
  host_flavor_id = "bx2d.host.152x608"
  zone           = "us-south-1"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 
- `host_flavor_id` - (Required, String) The unique identifier of the dedicated host flavor.
- `zone` - (Required, String) The zone of the dedicated host flavor.
 
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created.
- `flavor_class` - (String) The flavor class of the dedicated host flavor.
- `region` (String) The region of the dedicated host flavor.
- `deprecated` - (String) Describes if the dedicated host flavor is deprecated.
- `max_vcpus` - (String) The maximum available vcpus in the dedicated host flavor.
- `max_memory` - (String) The maximum available memory in the dedicated host flavor.
- `instance_storage` - (List) A nested block describes the instance storage of this dedicated host flavor.

  Nested scheme for `instance_storage`:
  - `count` - (Int) The count of the disks.
  - `size` - (Int) The size of the instance storage.

