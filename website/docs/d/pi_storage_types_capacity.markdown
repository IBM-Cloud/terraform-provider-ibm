---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_storage_types_capacity"
description: |-
  Manages storages capacity for all available storage types in the Power Virtual Server cloud.
---

# ibm_pi_storage_types_capacity
Retrieve information about storages capacity for all available storage types in a region. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_storage_types_capacity" "types" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

**Notes**

* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument reference
Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `maximum_storage_allocation` - (Map) Maximum storage allocation.

  Nested scheme for `maximum_storage_allocation`:
  - `max_allocation_size` - (Integer) Maximum allocation storage size (GB).
  - `storage_pool` - (String) The storage pool.
  - `storage_type`- (String) The storage type.

- `storage_types_capacity` - (List) List of storage types capacity.

  Nested scheme for `storage_types_capacity`:
  - `maximum_storage_allocation` - (Map) Maximum storage allocation.

    Nested scheme for `maximum_storage_allocation`:
    - `max_allocation_size` - (Integer) Maximum allocation storage size (GB).
    - `storage_pool` - (String) The storage pool.
    - `storage_type`- (String) The storage type.
 
  - `storage_pools_capacity` - (List) List of storage pools capacity.

    Nested scheme for `storage_pools_capacity`:
    - `max_allocation_size` - (Integer) Maximum allocation storage size (GB).
    - `pool_name` - (String) The pool name.
    - `storage_type` - (String) Storage type of the storage pool.
    - `total_capacity` - (Integer) Total pool capacity (GB).

  - `storage_type` - (String) The storage type.
