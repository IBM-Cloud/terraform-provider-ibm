---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance"
description: |-
  Manages an instance also known as VM/LPAR in the Power Virtual Server cloud.
---

# ibm_pi_instance

Create, delete or update a [Power Systems Virtual Server instance](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server).

## Example Usage

The following example creates a Power Systems Virtual Server instance.

```terraform
resource "ibm_pi_instance" "test-instance" {
    pi_memory             = "4"
    pi_processors         = "2"
    pi_instance_name      = "test-vm"
    pi_proc_type          = "shared"
    pi_image_id           = "${data.ibm_pi_image.powerimages.id}"
    pi_key_pair_name      = ibm_pi_key.key.key_id
    pi_sys_type           = "s922"
    pi_cloud_instance_id  = "51e1879c-bcbe-4ee1-a008-49cdba0eaf60"
    pi_pin_policy         = "none"
    pi_health_status      = "WARNING"
    pi_network {
      network_id = data.ibm_pi_public_network.dsnetwork.id
    }
}
```

~> **WARNING:** Updating a ibm_pi_instance resource with `pi_replicants` set does not update replicant vms!

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

The `ibm_pi_instance` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 120 minutes) Used for creating an instance.
- **update** - (Default 60 minutes) Used for updating an instance.
- **delete** - (Default 60 minutes) Used for deleting an instance.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_affinity_instance` - (Optional, String) PVM Instance (ID or Name) to base storage affinity policy against; required if requesting `affinity` and `pi_affinity_volume` is not provided.
- `pi_affinity_policy` - (Optional, String) Affinity policy for pvm instance being created; ignored if `pi_storage_pool` provided; for policy affinity requires one of `pi_affinity_instance` or `pi_affinity_volume` to be specified; for policy anti-affinity requires one of `pi_anti_affinity_instances` or `pi_anti_affinity_volumes` to be specified; Allowable values: `affinity`, `anti-affinity`
- `pi_affinity_volume`- (Optional, String) Volume (ID or Name) to base storage affinity policy against; required if requesting `affinity` and `pi_affinity_instance` is not provided.
- `pi_anti_affinity_instances` - (Optional, String) List of pvmInstances to base storage anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_volumes` is not provided.
- `pi_anti_affinity_volumes`- (Optional, String) List of volumes to base storage anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_instances` is not provided.
- `pi_boot_volume_replication_enabled` - (Optional, Boolean) Indicates if the boot volume should be replication enabled or not.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_deployment_target` - (Optional, List) The deployment of a dedicated host. Max items: 1.
  
  Nested scheme for `pi_deployment_target` :
  - `id` - (Required, String) The uuid of the host group or host.
  - `type` - (Required, String) The deployment target type. Supported values are `host` and `hostGroup`.

- `pi_deployment_type` - (Optional, String) Custom deployment type; Allowable value: `EPIC` or `VMNoStorage`.
- `pi_health_status` - (Optional, String) Specifies if Terraform should poll for the health status to be `OK` or `WARNING`. The default value is `OK`.

**Note**: IBM i software licenses for IBM i virtual server instances -- only for IBM i instances. Default to `false` and `0` if no values provided

- `pi_ibmi_css` - (Optional, Boolean) IBM i Cloud Storage Solution.
- `pi_ibmi_pha` - (Optional, Boolean) IBM i Power High Availability.
- `pi_ibmi_rds_users` - (Optional, Integer) IBM i Rational Dev Studio Number of User Licenses.
- `pi_image_id` - (Required, String) The ID of the image that you want to use for your Power Systems Virtual Server instance. The image determines the operating system that is installed in your instance. To list available images, run the `ibmcloud pi images` command.
  - **Note**:
        - Only images belonging to your project can be used image for deploying a Power Systems Virtual Server instance. To import an images to your project, see [ibm_pi_image](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_image).
        - If using `pi_deployment_type = VMNoStorage` then use the following images for the respective OS you intend to create the instance: `AIX-EMPTY`, `IBMI-EMPTY`, `SLES-EMPTY`, `RHEL-EMPTY`.
- `pi_instance_name` - (Required, String) The name of the Power Systems Virtual Server instance.
- `pi_key_pair_name` - (Optional, String) The name of the SSH key that you want to use to access your Power Systems Virtual Server instance. The SSH key must be uploaded to IBM Cloud.
- `pi_license_repository_capacity` - (Optional, Integer) The VTL license repository capacity TB value. Only use with VTL instances. `pi_memory >= 16 + (2 * pi_license_repository_capacity)`.
- `pi_memory` - (Optional, Float) The amount of memory that you want to assign to your instance in GB.
  - Required when not creating SAP instances. Conflicts with `pi_sap_profile_id`.
- `pi_network` - (Required, List of Map) List of one or more networks to attach to the instance.

  The `pi_network` block supports:
  - `ip_address` - (Optional, String) The ip address to be used of this network.
  - `network_id` - (Required, String) The network ID to assign to the instance.
  - `network_security_group_ids` - (Optional, List) The Network security groups that the network interface is a member of. There is a limit of 1 network security group in the array. If not specified, default network security group is used.
- `pi_pin_policy` - (Optional, String) Select the pinning policy for your Power Systems Virtual Server instance. Supported values are `soft`, `hard`, and `none`.    **Note** You can choose to soft pin (`soft`) or hard pin (`hard`) a virtual server to the physical host where it runs. When you soft pin an instance for high availability, the instance automatically migrates back to the original host once the host is back to its operating state. If the instance has a licensing restriction with the host, the hard pin option restricts the movement of the instance during remote restart, automated remote restart, DRO, and live partition migration. The default pinning policy is `none`.
- `pi_placement_group_id` - (Optional, String) The ID of the placement group that the instance is in or empty quotes `""` to indicate it is not in a placement group. The meta-argument `count` and a `pi_replicants` cannot be used when specifying a placement group ID. Instances provisioning in the same placement group must be provisioned one at a time; however, to provision multiple instances on the same host or different hosts then use `pi_replicants` and `pi_replication_policy` instead of `pi_placement_group_id`.
- `pi_processors` - (Optional, Float) The number of vCPUs to assign to the VM as visible within the guest Operating System.
  - Required when not creating SAP instances. Conflicts with `pi_sap_profile_id`.
- `pi_proc_type` - (Optional, String) The type of processor mode in which the VM will run with `shared`, `capped` or `dedicated`.
  - Required when not creating SAP instances. Conflicts with `pi_sap_profile_id`.
- `pi_replicants` - (Optional, Integer) The number of instances that you want to provision with the same configuration. If this parameter is not set, `1` is used by default.
- `pi_replication_policy` - (Optional, String) The replication policy that you want to use, either `affinity`, `anti-affinity` or `none`. If this parameter is not set, `none` is used by default.
- `pi_replication_scheme` - (Optional, String) The replication scheme that you want to set, either `prefix` or `suffix`.
- `pi_replication_sites` - (Optional, List) Indicates the replication sites of the boot volume.
- `pi_retain_virtual_serial_number` - (Optional, Boolean) Indicates whether attached virtual serial number will be reserved when serial assigned to instance is changed, removed, or instance is deleted. If using `ibm_pi_virtual_serial_number` resource, will unassign and unreserved virtual serial number attached to instance if set to false. Default value is `false`.
- `pi_sap_profile_id` - (Optional, String) SAP Profile ID for the amount of cores and memory.
  - Required only when creating SAP instances.
- `pi_sap_deployment_type` - (Optional, String) Custom SAP deployment type information (For Internal Use Only).
- `pi_shared_processor_pool` - (Optional, String) The shared processor pool for instance deployment. Conflicts with `pi_sap_profile_id`.
- `pi_storage_pool` - (Optional, String) Storage Pool for server deployment; if provided then `pi_affinity_policy` will be ignored; Only valid when you deploy one of the IBM supplied stock images. Storage pool for a custom image (an imported image or an image that is created from a VM capture) defaults to the storage pool the image was created in.
- `pi_storage_pool_affinity` - (Optional, Boolean) Indicates if all volumes attached to the server must reside in the same storage pool. The default value is `true`. To attach data volumes from a different storage pool (mixed storage) set to `false` and use `pi_volume_attach` resource. Once set to `false`, cannot be set back to `true` unless all volumes attached reside in the same storage type and pool.
- `pi_storage_type` - (Optional, String) - Storage type for server deployment; If storage type is not provided the storage type will default to `tier3`. To get a list of available storage types, please use the [ibm_pi_storage_types_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_types_capacity) data source.
- `pi_storage_connection` - (Optional, String) - Storage Connectivity Group (SCG) for server deployment. Supported values are `vSCSI`, `maxVolumeSupport`.
- `pi_sys_type` - (Optional, String) The type of system on which to create the VM (e880/e980/e1080/s922/s1022).
  - Supported SAP system types are (e880/e980/e1080).
- `pi_user_data` - (Optional, String) The user data `cloud-init` to pass to the instance during creation. It can be a base64 encoded or an unencoded string. If it is an unencoded string, the provider will encode it before it passing it down.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.
- `pi_virtual_cores_assigned`  - (Optional, Integer) Specify the number of virtual cores to be assigned.
- `pi_virtual_optical_device` - (Optional, String) Virtual Machine's Cloud Initialization Virtual Optical Device.
- `pi_virtual_serial_number` - (Optional, List)  Virtual Serial Number information. If using `ibm_pi_virtual_serial_number` resource to manage a virtual serial number assigned to this instance, it is strongly recommended to ignore changes in this argument using the `ignore_changes` meta-argument in the `lifecycle`.
  
  Nested scheme for `pi_virtual_serial_number`:
  - `description` - (Optional, String) Description of virtual serial number.
  - `serial` - (Required, String) Provide an existing reserved Virtual Serial Number or specify 'auto-assign' for auto generated Virtual Serial Number. Updates to this will shutdown then restart power VM instance.
  - `software_tier` - (Optional, String) Software tier for virtual serial number. Allowed values are: ["P05", "P10", "P20", "P30"]. Updates to this will shutdown then restart power VM instance.

      ~> **Note** When set to "auto-assign", changes to `serial` outside of terraform will not be detected. In addition, if a new generated virtual serial number is needed, the old serial must be removed before a new one is generated.
- `pi_volume_ids` - (Optional, List of String) The list of volume IDs that you want to attach to the instance during creation.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of this resource.
- `dedicated_host_id` - (String) The dedicated host ID where the shared processor pool resides.
- `fault` - (Map) Fault information, if any.
  
   Nested scheme for `fault`:
      - `code` - (String) The fault status of the server.
      - `created` - (String) The date and time the fault occurred.
      - `details` - (String) The fault details of the server.
      - `message` -  (String) The fault message of the server.

- `health_status` - (String) The health status of the VM.
- `ibmi_rds` - (Boolean) IBM i Rational Dev Studio.
- `id` - (String) The unique identifier of the instance. The ID is composed of `<pi_cloud_instance_id>/<instance_id_1>/.../<instance_id_n>`.
- `instance_id` - (String) The unique identifier of the instance.
- `max_processors`- (Float) The maximum number of processors that can be allocated to the instance with shutting down or rebooting the `LPAR`.
- `max_virtual_cores` - (Integer) The maximum number of virtual cores.
- `min_processors` - (Float) The minimum number of processors that the instance can have.
- `min_memory` - (Float) The minimum memory that was allocated to the instance.
- `max_memory`- (Float) The maximum amount of memory that can be allocated to the instance without shut down or reboot the `LPAR`.
- `min_virtual_cores` - (Integer) The minimum number of virtual cores.
- `pin_policy`  - (String) The pinning policy of the instance.
- `pi_network` - (List of Map) - A list of networks that are assigned to the instance.
  Nested scheme for `pi_network`:
  - `external_ip` - (String) The external IP address of the network.
  - `ip_address` - (String) The IP address of the network.
  - `mac_address` - (String) The MAC address of the network.
  - `network_id` - (String) The ID of the network.
  - `network_interface_id` - (String) ID of the network interface.
  - `network_name` - (String) The name of the network.
  - `network_security_group_ids` - (List) IDs of the network necurity groups that the network interface is a member of.
  - `network_security_groups_href` - (List) Links to the network security groups that the network interface is a member of.
  - `type` - (String) The type of network.
- `progress` - (Float) - Specifies the overall progress of the instance deployment process in percentage.
- `shared_processor_pool_id` - (String)  The ID of the shared processor pool for the instance.
- `status` - (String) The status of the instance.

## Import

The `ibm_pi_instance` can be imported using `pi_cloud_instance_id` and `instance_id`.

### Example

```bash
terraform import ibm_pi_instance.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770b112ebb
```
