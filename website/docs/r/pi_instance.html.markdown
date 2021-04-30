---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance"
description: |-
  Manages an instance (a.k.a. VM/LPAR) in the Power Virtual Server Cloud.
---

# ibm\_pi_instance

Provides an instance resource. This allows instances to be created or updated in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create an instance:

```hcl
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
## Notes:
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  Example Usage:
  ```hcl
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Timeouts

ibm_pi_instance provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating an instance.
* `delete` - (Default 60 minutes) Used for deleting an instance.

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

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the instance.The id is composed of \<power_instance_id\>/\<instance_id\>.
* `instance_id` - The unique identifier of the instance.
* `status` - The status of the VM.
* `health_status` - The health status of the VM.
* `migratable` - The state of instance migratable.
* `min_processors` - Minimum number of processors that were allocated (for resize)
* `min_memory` - Minimum Memory that was  allocated (for resize)
* `max_processors` - Maximumx number of processors that can be allocated (for resize) without a shutdown/reboot of the lpar
* `max_memory` - Maximum amount of memory that can be allocated (for resize) without a shutdown/reboot of the lpar
* `progress` - The progress of the instance.
* `addresses` - A list of addresses assigned to the VM. Nested `addresses` blocks have the following structure:
	* `ip` - IP of the instance.
  * `macaddress` - The macaddress of the instance.
  * `networkid` - The networkID of the instance.
  * `network_name` - The network name of the instance.
  * `type` - The type of the network
  * `external_ip` - The externalIP address of the instance.
* `pin_policy` - The pin policy of the instance
* `max_virtual_cores` - The maximum number of virtual cores
* `min_virtual_cores` - The minimum number of virtual cores
## Import

ibm_pi_instance can be imported using `power_instance_id` and `instance_id`, eg

```
$ terraform import ibm_pi_instance.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
