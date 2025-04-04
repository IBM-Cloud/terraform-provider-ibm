---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_shared_processor_pools"
description: |-
  Manages the shared processor pools in the Power Virtual Server cloud.
---

# ibm_pi_shared_processor_pools

Retrieve information about all shared processor pools. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_shared_processor_pools" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
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

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `shared_processor_pools` - (List) List of all the shared processor pools.

  Nested scheme for `shared_processor_pools`:
  - `allocated_cores` - (Float) The allocated cores in the shared processor pool.
  - `available_cores` - (Integer) The available cores in the shared processor pool.
  - `crn` - (String) The CRN of this resource.
  - `dedicated_host_id` - (String) The dedicated host ID where the shared processor pool resides.
  - `host_id` - (Integer) The host ID where the shared processor pool resides.
  - `name` - (String) The name of the shared processor pool.
  - `reserved_cores` - (Integer) The amount of reserved cores for the shared processor pool.
  - `shared_processor_pool_id` - (String) The shared processor pool's unique ID.
  - `status` - (String) The status of the shared processor pool.
  - `status_detail` - (String) The status details of the shared processor pool.
  - `user_tags` - (List) List of user tags attached to the resource.
