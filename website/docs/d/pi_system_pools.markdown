---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_system_pools"
description: |-
  Manages system pools within a particular data center in the Power Virtual Server cloud.
---

# ibm_pi_system_pools
Retrieve information about list of system pools within a particular data center. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_system_pools" "pools" {
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

- `system_pools` - (List) The available system pools within a particular DataCenter.

  Nested scheme for `system_pools`:
  - `system_pool_name` - (String) The system pool name.
  - `capacity` - (Map) Advertised capacity cores and memory (GB).

    Nested scheme for `capacity`:
    - `cores` - (String) The host available Processor units.
    - `id` - (String) The host identifier.
    - `memory`- (String) The host available RAM memory in GiB.

  - `core_memory_ratio` - (Float) Processor to Memory (GB) Ratio.
  - `max_available` - (Map) Maximum configurable cores and memory (GB) (aggregated from all hosts).

    Nested scheme for `max_available`:
    - `cores` - (String) The host available Processor units.
    - `id` - (String) The host identifier.
    - `memory`- (String) The host available RAM memory in GiB.

  - `max_cores_available` - (Map) Maximum configurable cores available combined with available memory of that host.

    Nested scheme for `max_cores_available`:
    - `cores` - (String) The host available Processor units.
    - `id` - (String) The host identifier.
    - `memory`- (String) The host available RAM memory in GiB.

  - `max_memory_available` - (Map) Maximum configurable memory available combined with available cores of that host.

    Nested scheme for `max_memory_available`:
    - `cores` - (String) The host available Processor units.
    - `id` - (String) The host identifier.
    - `memory`- (String) The host available RAM memory in GiB.

  - `shared_core_ratio` - (Map) The min-max-default allocation percentage of shared core per vCPU.

    Nested scheme for `shared_core_ratio`:
    - `default` - (String) The default value.
    - `max` - (String) The max value.
    - `min`- (String) The min value.

  - `systems` - (List) The DataCenter list of servers and their available resources.

    Nested scheme for `systems`:
    - `cores` - (String) The host available Processor units.
    - `id` - (String) The host identifier.
    - `memory`- (String) The host available RAM memory in GiB.

  - `type` - (String) Type of system hardware.
