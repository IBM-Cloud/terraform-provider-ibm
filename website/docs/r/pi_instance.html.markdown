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

The `ibm_pi_instance` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - The creation of the instance is considered failed if no response is received for 60 minutes. 
- **delete** - The deletion of the instance is considered failed if no response is received for 60 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_health_status` - (Optional, String) Specifies if Terraform should poll for the health status to be `OK` or `WARNING`. The default value is `OK`.
- `pi_image_id` - (Required, String) The ID of the image that you want to use for your Power Systems Virtual Server instance. The image determines the operating system that is installed in your instance. To list available images, run the `ibmcloud pi images` command.
- `pi_instance_name` - (Required, String) The name of the Power Systems Virtual Server instance. 
- `pi_key_pair_name` - (Required, String) The name of the SSH key that you want to use to access your Power Systems Virtual Server instance. The SSH key must be uploaded to IBM Cloud.
- `pi_memory` - (Required, Float) The amount of memory that you want to assign to your instance in gigabytes.
- `pi_migratable`- (Optional, Bool) Indicates the VM is migrated or not.
- `pi_network_ids` - (Required, String) The list of network IDs that you want to assign to the instance. 
- `pi_pin_policy` - (Optional, String) Select the pinning policy for your Power Systems Virtual Server instance. Supported values are `soft`, `hard`, and `none`.    **Note** You can choose to soft pin (`soft`) or hard pin (`hard`) a virtual server to the physical host where it runs. When you soft pin an instance for high availability, the instance automatically migrates back to the original host once the host is back to its operating state. If the instance has a licensing restriction with the host, the hard pin option restricts the movement of the instance during remote restart, automated remote restart, DRO, and live partition migration. The default pinning policy is `none`. 
- `pi_processors` - (Required, Float) The number of vCPUs to assign to the VM as visible within the guest Operating System. 
- `pi_proc_type` - (Required, String) The type of processor mode in which the VM will run with `shared` or `dedicated`.
- `pi_replicants` - (Optional, Float) The number of instances that you want to provision with the same configuration. If this parameter is not set,  `1` is used by default.
- `pi_replication_policy` - (Optional, String) The replication policy that you want to use. If this parameter is not set, `none` is used by default. 
- `pi_replication_scheme` - (Optional, String) The replication scheme that you want to set, either `prefix` or `suffix`.
- `pi_storage_type` - (Optional, String) - Storage type for server deployment. Only valid when you deploy one of the IBM supplied stock images. Storage type for a custom image (an imported image or an image that is created from a VM capture) defaults to the storage type the image was created in
- `pi_sys_type` - (Required, String) The type of system on which to create the VM (s922/e880/any). 
- `pi_user_data` - (Optional, String) The base64 encoded form of the user data `cloud-init` to pass to the instance during creation. 
- `pi_virtual_cores_assigned`  - (Optional, Integer) Specify the number of virtual cores to be assigned.
- `pi_volume_ids` - (Required, String) The list of volume IDs that you want to attach to the instance during creation.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `addresses` - Map of strings - A list of addresses that are assigned to the instance.

  Nested scheme for `addresses`:
  - `ip` - (String) The IP address of the instance.
  - `macaddress` - (String) The MAC address of the instance.
  - `networkid` - (String) The network ID of the instance.
  - `network_name` - (String) The network name of the instance.
  - `type` - (String) The type of network.
  - `external_ip` - (String) The external IP address of the instance.
- `health_status` - (String) The health status of the VM.
- `id` - (String) The unique identifier of the instance. The ID is composed of `<power_instance_id>/<instance_id>`.
- `instance_id` - (String) The unique identifier of the instance. 
- `max_processors`- Integer- The maximum number of processors that can be allocated to the instance with shutting down or rebooting the `LPAR`.
- `max_virtual_cores` - (Integer) The maximum number of virtual cores.
- `migratable` - (Bool) Indicates the VM is migrated or not.This attribute is Deprecated use `pi_migratable` instead
- `min_processors` - Float - The minimum number of processors that the instance can have. 
- `min_memory` - (Integer) The minimum memory that was allocated to the instance.
- `max_memory`- (Integer) The maximum amount of memory that can be allocated to the instance without shut down or reboot the `LPAR`.
- `min_virtual_cores` - (Integer) The minimum number of virtual cores.
- `status` - (String) The status of the instance.
- `pin_policy`  - (String) The pinning policy of the instance.
- `progress` - (Float) - Specifies the overall progress of the instance deployment process in percentage.

## Import

The `ibm_pi_instance` can be imported using `power_instance_id` and `instance_id`.

**Example**

```
$ terraform import ibm_pi_instance.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770b112ebb
```
