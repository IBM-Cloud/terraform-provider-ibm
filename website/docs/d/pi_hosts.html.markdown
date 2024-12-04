---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : ibm_pi_hosts"
description: |-
    Get information about hosts in Power Virtual Server.
---

# ibm_pi_hosts

Provides a read-only data source to retrieve information about hosts. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_hosts" "hosts" {
    pi_cloud_instance_id  = "<value of the cloud_instance_id>"
    
}
```

## Notes

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

You can specify the following arguments for this data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
  
## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `hosts` - (List) List of hosts.
  
    Nested schema for `hosts`:
  - `capacity` - (List) Capacities of the host.
  
     Nested schema for `capacity`:
        - `available_core` - (Float) Number of cores currently available.
        - `available_memory` - (Float) Amount of memory currently available (in GB).
        - `reserved_core` - (Float) Number of cores reserved for system use.
        - `reserved_memory` - (Float) Amount of memory reserved for system use (in GB).
        - `total_core` - (Float) Total number of cores of the host.
        - `total_memory` - (Float) Total amount of memory of the host (in GB).
        - `used_core` - (Float) Number of cores in use on the host.
        - `used_memory` - (Float) Amount of memory used on the host (in GB).

  - `display_name` - (String) Name of the host.
  - `host_group` - (Map)  Information about the owning host group.

       Nested schema for `host_group`:
        - `access` - (String) Whether the host group is a primary or secondary host group.
        - `href` - (String) Link to the host group resource.
        - `name` - (String) Name of the host group.
  - `host_id` - (String)  ID of the host.
  - `host_reference` - (Integer) Current physical ID of the host.
  - `id` - (String) The unique identifier of the pi_hosts.
  - `state` - (String) State of the host `up` or `down`.
  - `status` - (String) Status of the host `enabled` or `disabled`.
  - `sys_type` - (String) System type.
