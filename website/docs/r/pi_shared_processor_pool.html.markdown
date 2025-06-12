---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_shared_processor_pool"
description: |-
  Manages a shared processor pool in the Power Virtual Server cloud.
---

# ibm_pi_shared_processor_pool

Create, update or delete a shared processor pool.

## Example Usage

The following example enables you to create a shared processor pool with a group 2 reserved cores on a s922 host group:

```terraform
resource "ibm_pi_shared_processor_pool" "testacc_shared_processor_pool" {
  pi_cloud_instance_id                    = "<value of the cloud_instance_id>"
  pi_shared_processor_pool_host_group     = "s922"
  pi_shared_processor_pool_name           = "my_spp"
  pi_shared_processor_pool_reserved_cores = "2"
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

## Timeouts

ibm_pi_shared_processor_pool provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating a shared processor pool.
- **delete** - (Default 60 minutes) Used for deleting a shared processor pool.
- **update** - (Default 60 minutes) Used for updating a shared processor pool.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_host_id` - (Optional, String) The host id of a host in a host group (only available for dedicated hosts).
- `pi_shared_processor_pool_host_group` - (Required, String) Host group of the shared processor pool. Valid values are 's922', 'e980' and 's1022'.
- `pi_shared_processor_pool_name` - (Required, String) The name of the shared processor pool.
- `pi_shared_processor_pool_placement_group_id` - (Deprecated, Optional, String) The ID of the placement group the shared processor pool is created in. Please use pi_shared_processor_pool_placement_groups instead.
- `pi_shared_processor_pool_placement_groups` - (Optional, List) The list of shared processor pool placement groups that the shared processor pool is in.
- `pi_shared_processor_pool_reserved_cores` - (Required, Integer) The amount of reserved cores for the shared processor pool.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `allocated_cores` - (Float) The allocated cores in the shared processor pool.
- `available_cores` - (Integer) The available cores in the shared processor pool.
- `crn` - (String) The CRN of this resource.
- `dedicated_host_id` - (String) The dedicated host ID where the shared processor pool resides.
- `host_id` - (Integer) The host ID where the shared processor pool resides.
- `instances` - (List of Map) The list of server instances that are deployed in the shared processor pool.
  
  Nested scheme for `instances`:
  - `availability_zone` - (String) Availability zone for the server instances.
  - `cpus` - (Integer) The amount of cpus for the server instance.
  - `id` - (String) The server instance ID.
  - `memory` - (Integer) The amount of memory for the server instance.
  - `name` - (String) The server instance name.
  - `status` - (String) Status of the instance.
  - `uncapped` - (Bool) Identifies if uncapped or not.
  - `vcpus` - (Float) The amout of vcpus for the server instance.
- `shared_processor_pool_id` - (String) The shared processor pool's unique ID.
- `status` - (String) The status of the shared processor pool.
- `status_detail` - (String) The status details of the shared processor pool.

## Import

The `ibm_pi_shared_processor_pool` resource can be imported by using `pi_cloud_instance_id` and `shared_processor_pool_id`.

### Example

```bash
terraform import ibm_pi_shared_processor_pool.example d7bec597-4726-451f-8a63-e62e6f19c32c/b17a2b7f-77ab-491c-811e-495f8d4c8947
```
