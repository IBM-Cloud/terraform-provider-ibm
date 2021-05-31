---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance"
description: |-
  Manages an instance also known as VM/LPAR in the Power Virtual Server cloud.
---

# ibm_pi_instance
Create or update a [Power Systems Virtual Server instance](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server).

## Example usage
The following example creates a Power Systems Virtual Server instance. 

```terraform
resource "ibm_pi_instance" "test-instance" {
    pi_memory             = "4"
    pi_processors         = "2"
    pi_instance_name      = "test-vm"
    pi_proc_type          = "shared"
    pi_image_id           = "${data.ibm_pi_image.powerimages.id}"
    pi_network_ids        = [data.ibm_pi_public_network.dsnetwork.id]
    pi_key_pair_name      = ibm_pi_key.key.key_id
    pi_sys_type           = "s922"
    pi_cloud_instance_id  = "51e1879c-bcbe-4ee1-a008-49cdba0eaf60"
    pi_pin_policy         = "none"
    pi_health_status      = "WARNING"
}
```

**Note**
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

## Timeouts

The `ibm_pi_instance` provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- **create** - The creation of the instance is considered failed if no response is received for 60 minutes. 
- **delete** - The deletion of the instance is considered failed if no response is received for 60 minutes. 


## Argument Reference

The following arguments are supported:

* `pi_instance_name` - (Required, string) The name of the VM.
* `pi_key_pair_name` - (Required, string) The name of the Power Virtual Server Cloud SSH key to used to login to the VM.
* `pi_image_id` - (Required, string) The name of the image to deploy (e.g., 7200-03-03).
* `pi_processors` - (Required, float) The number of vCPUs to assign to the VM (as visibile within the guest operating system).
* `pi_proc_type` - (Required, string) The type of processor mode in which the VM will run (shared/dedicated/capped).
* `pi_memory` - (Required, float) The amount of memory (GB) to assign to the VM.
* `pi_sys_type` - (Required, string) The type of system on which to create the VM (s922/e880/e980).
* `pi_volume_ids` - (Optional, list(string)) The list of volume IDs to attach to the VM at creation time.
* `pi_network_ids` - (Required, list(string)) The list of network IDs assigned to the VM.
* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account
* `pi_user_data` - (Optional, string) The base64-encoded form of the user data (cloud-init) to pass to the VM at creation time.
* `pi_replicants` - (Optional, float) Specifies the number of VMs to deploy; default is 1.
* `pi_replication_policy` - (Optional, string) Specifies the replication policy (e.g., none).
* `pi_replication_scheme` - (Optional, string) Specifies the replicate scheme (prefix/suffix).
* `pi_pin_policy` - (Optional,string) Specifies the pin policy for the lpar (none/soft/hard) - This is dependent on the cloud instance capabilities.
* `pi_health_status` - (Optional,string) Specifies if terraform should poll for the Health Status to be OK or WARNING.  Default is OK. 
* `pi_virtual_cores_assigned` - (Optional,integer) Specifies the number of virtual cores to be assigned 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `addresses` - Map of strings - A list of addresses that are assigned to the instance.

  Nested scheme for `addresses`:
  - `ip` - (String) The IP address of the instance.
  
  Nested scheme for `ip`:
  - `macaddress` - (String) The MAC address of the instance.
  - `networkid` - (String) The network ID of the instance.
  - `network_name` - (String) The network name of the instance.
  - `type` - (String) The type of network.
  - `externalip` - (String) The external IP address of the instance.
- `id` - (String) The unique identifier of the instance. The ID is composed of `<power_instance_id>/<instance_id>`.
- `instance_id` - (String) The unique identifier of the instance. 
- `pin_policy`  - (String) The pinning policy of the instance.
- `progress` - Float - Specifies the overall progress of the instance deployment process in percentage.
- `status` - (String) The status of the instance.
- `health_status` - (String) The health status of the VM.
- `migratable`- (Bool) Indicates the VM is migrated or not.
- `max_processors`- Integer- The maximum number of processors that can be allocated to the instance with shutting down or rebooting the `LPAR`.
- `max_virtual_cores` - (Integer) The maximum number of virtual cores.
- `min_processors` - Float - The minimum number of processors that the instance can have. 
- `min_memory` - (Integer) The minimum memory that was allocated to the instance.
- `max_memory`- (Integer) The maximum amount of memory that can be allocated to the instance without shut down or reboot the `LPAR`.
- `min_virtual_cores` - (Integer) The minimum number of virtual cores.

## Import

The `ibm_pi_instance` can be imported using `power_instance_id` and `instance_id`.

**Example**

```
$ terraform import ibm_pi_instance.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
