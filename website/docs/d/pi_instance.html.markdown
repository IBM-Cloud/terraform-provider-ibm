---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance"
description: |-
  Manages an instance in the Power Virtual Server cloud.
---

# ibm_pi_instance
Retrieve information about a Power Systems Virtual Server instance. For more information, about Power Virtual Server instance, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_instance" "ds_instance" {
  pi_instance_name     = "terraform-test-instance"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
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
- `pi_instance_name` - (Required, String) The name of the instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `addresses` - (Deprecated, List of objects) - The address associated with this instance.

  Nested scheme for `addresses`:
  - `ip` - (String) The IP address of the instance.
  - `external_ip` - (String) The external IP address of the instance.
  - `macaddress` - (String) The MAC address of the instance.
  - `network_id` - (String) The network ID of the instance.
  - `network_name` - (String) The network name of the instance.
  - `type` - (String) The type of the network.
- `deployment_type` - (String) The custom deployment type.
- `health_status` - (String) The health of the instance.
- `id` - (String) The unique identifier of the instance.
- `license_repository_capacity` - The VTL license repository capacity TB value. Only available with VTL instances.
- `memory` - (Float) The amount of memory that is allocated to the instance.
- `minproc`- (Float) The minimum number of processors that must be allocated to the instance. 
- `maxproc`- (Float) The maximum number of processors that can be allocated to the instance without shutting down or rebooting the `LPAR`.
- `max_virtual_cores` - (Integer) The maximum value that you increase without a reboot.
- `minmem`- (Float) The minimum amount of memory that must be allocated to the instance.
- `maxmem`- (Float) The maximum amount of memory that can be allocated to the instance without shutting down or rebooting the `LPAR`.
- `min_virtual_cores` - (Integer) The minimum cores assigned to an instance.
- `networks` - List of objects - The network associated with this instance.

  Nested scheme for `networks`:
  - `ip` - (String) The IP address of the instance.
  - `external_ip` - (String) The external IP address of the instance.
  - `macaddress` - (String) The MAC address of the instance.
  - `network_id` - (String) The network ID of the instance.
  - `network_name` - (String) The network name of the instance.
  - `type` - (String) The type of the network.
- `placement_group_id`- (String) The ID of the placement group that the instance is a member.
- `processors` - (Float) The number of processors that are allocated to the instance.
- `proctype` - (String) The procurement type of the instance. Supported values are `shared` and `dedicated`.
- `shared_processor_pool`- (String) The name of the shared processor pool for the instance.
- `shared_processor_pool_id` - (String)  The ID of the shared processor pool for the instance.
- `status` - (String) The status of the instance.
- `storage_pool` - (String) The storage Pool where server is deployed.
- `storage_pool_affinity` - (Bool) Indicates if all volumes attached to the server must reside in the same storage pool.
- `storage_type` - (String) The storage type where server is deployed.
- `virtual_cores_assigned` - (Integer) The virtual cores that are assigned to the instance.
- `volumes` - (List of strings) The list of volume IDs that are attached to the instance.

  **Notes** Ibmi software licenses for IBMi virtual server instances --software licenses currently only supports IBMi instances
- `ibm_css` - (Bool) IBMi Cloud Storage Solution.
- `ibmi_dbq` - (Bool) IBMi DBQ.
- `ibmi_pha` - (Bool) IBMi Power High Availability.
- `ibmi_rds` - (Bool) IBMi Rational Dev Studio.
- `ibmi_rds_users` - (Integer) IBMi Rational Dev Studio Number of User Licenses.
