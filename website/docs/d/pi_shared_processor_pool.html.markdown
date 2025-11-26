---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_shared_processor_pool"
description: |-
  Manages a shared processor pool in the Power Virtual Server cloud.
---

# ibm_pi_shared_processor_pool

Retrieve information about a shared processor pool. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_shared_processor_pool" "ds_pool" {
  pi_shared_processor_pool_id   = "my-spp"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`
  
Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_shared_processor_pool_id` - (Required, String) The ID of the shared processor pool.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `allocated_cores` - (Float) The allocated cores in the shared processor pool.
- `available_cores` - (Integer) The available cores in the shared processor pool.
- `creation_date` - (String) Date of shared processor pool creation.
- `crn` - (String) The CRN of this resource.
- `dedicated_host_id` - (String) The dedicated host ID where the shared processor pool resides.
- `host_id` - (Integer) The host ID where the shared processor pool resides.
- `id` - (String) The shared processor pool's unique ID.
- `instances` - (List) List of server instances deployed in the shared processor pool.

  Nested scheme for `instances`:
  - `availability_zone` - (String) Availability zone for the server instances.
  - `cpus` - (Integer) The amount of cpus for the server instance.
  - `id` - (String) The server instance ID.
  - `memory` - (Integer) The amount of memory for the server instance.
  - `name` - (String) The server instance name.
  - `status` - (String) Status of the instance.
  - `uncapped` - (Bool) Identifies if uncapped or not.
  - `vcpus` - (Float) The amout of vcpus for the server instance.
- `name` - (String) The name of the shared processor pool.
- `reserved_cores` - (Integer) The amount of reserved cores for the shared processor pool.
- `status` - (String) The status of the shared processor pool.
- `status_detail` - (String) The status details of the shared processor pool.
- `user_tags` - (List) List of user tags attached to the resource.
