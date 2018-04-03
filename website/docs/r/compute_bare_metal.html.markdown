---
layout: "ibm"
page_title: "IBM: compute_bare_metal"
sidebar_current: "docs-ibm-resource-compute-bare-metal"
description: |-
  Manages IBM Compute bare metal servers.
---

# ibm\_compute_bare_metal

Provides a bare metal resource. This allows bare metal servers to be created, updated, and deleted. This resource supports both monthly bare metal servers and hourly bare metal servers. For more detail on bare metal severs, refer to the [IBM Cloud bare metal server page](https://www.ibm.com/cloud-computing/bluemix/bare-metal-servers).


## Hourly bare metal server
When the `ibm_compute_bare_metal` resource definition has a `fixed_config_preset` attribute, Terraform creates an hourly bare metal server. Hardware specifications are predefined in the `fixed_config_preset` attribute and cannot be modified. The following example shows you how to create a new hourly bare metal server.

### Example of an hourly bare metal server
```hcl
resource "ibm_compute_bare_metal" "hourly-bm1" {
    hostname = "hourly-bm1"
    domain = "example.com"
    os_reference_code = "UBUNTU_16_64"
    datacenter = "dal01"
    network_speed = 100 # Optional
    hourly_billing = true # Optional
    private_network_only = false # Optional
    fixed_config_preset = "S1270_8GB_2X1TBSATA_NORAID"

    user_metadata = "{\"value\":\"newvalue\"}" # Optional
    tags = [
      "collectd",
      "mesos-master"
    ]
    notes = "note test"
}
```

## Monthly bare metal server
When the `fixed_config_preset` attribute is not configured, Terraform creates a monthly bare metal server resource. The monthly bare metal server resource provides options to configure process, memory, network, disk, and RAID. You can also can assign VLANs and subnets for the target monthly bare metal server. To configure the monthly bare metal server, you must provide additional attributes such as `package_key_name`, `proecss_key_name`, `disk_key_names`, and `os_key_name`. The following example shows you how to create a new monthly bare metal server.

### Example of a monthly bare metal server
```hcl
resource "ibm_compute_bare_metal" "monthly_bm1" {

# Mandatory fields
    package_key_name = "DUAL_E52600_V4_12_DRIVES"
    process_key_name = "INTEL_INTEL_XEON_E52620_V4_2_10"
    memory = 64
    os_key_name = "OS_WINDOWS_2012_R2_FULL_DC_64_BIT_2"
    hostname = "cust-bm"
    domain = "ms.com"
    datacenter = "wdc04"
    network_speed = 100
    public_bandwidth = 500
    disk_key_names = [ "HARD_DRIVE_800GB_SSD", "HARD_DRIVE_800GB_SSD", "HARD_DRIVE_800GB_SSD" ]
    hourly_billing = false

# Optional fields
    private_network_only = false
    unbonded_network = true
    user_metadata = "{\"value\":\"newvalue\"}"
    public_vlan_id = 12345678
    private_vlan_id = 87654321
    public_subnet = "50.97.46.160/28"
    private_subnet = "10.56.109.128/26"
    tags = [
      "collectd",
      "mesos-master"
    ]
    redundant_power_supply = true
    storage_groups = {
       # RAID 5
       array_type_id = 3
       # Use three disks
       hard_drives = [ 0, 1, 2]
       array_size = 1600
       # Basic partition template for Windows
       partition_template_id = 17
    }
}
```

**Note**: Monthly bare metal servers do not support `immediate cancellation`. When Terraform deletes the monthly bare metal server, the `anniversary date cancellation` option is used.

## Create a bare metal server using quote ID
If you already have a quote ID for the bare metal server, you can create a new bare metal server with the quote ID. The following example shows you how to create a new bare metal server with a quote ID.

### Example of a quote based ordering
```hcl
resource "ibm_compute_bare_metal" "quote_test" {

# Mandatory fields
    hostname = "quote-bm-test"
    domain = "example.com"
    quote_id = 2209349

# Optional fields
    user_metadata = "{\"value\":\"newvalue\"}"
    public_vlan_id = 12345678
    private_vlan_id = 87654321
    public_subnet = "50.97.46.160/28"
    private_subnet = "10.56.109.128/26"
    tags = [
      "collectd",
      "mesos-master"
    ]  
}
```

## Argument Reference

The following arguments are supported:

### Arguments for all bare metal server types

* `hostname` - (Optional, string) The hostname for the computing instance.
* `domain` - (Required, string) The domain for the computing instance.
* `user_metadata` - (Optional, string) Arbitrary data to be made available to the computing instance.
* `notes` - (Optional, string) Notes to associate with the instance.
* `ssh_key_ids` - (Optional, array of numbers) The SSH key IDs to install on the computing instance when the instance is provisioned.
    **NOTE:** If you don't know the ID(s) for your SSH keys, you can [reference your SSH keys by their labels](../d/compute_ssh_key.html).
* `post_install_script_uri` - (Optional, string) The URI of the script to be downloaded and executed after installation is complete.
* `tags` - (Optional, array of strings) Tags associated with this bare metal server. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters will be removed.
* `file_storage_ids` - (Optional, array of numbers) File storage to which this computing instance should have access. File storage must be in the same data center as the bare metal server. If you use this argument to authorize access to file storage, do not use the `allowed_hardware_ids` argument in the `ibm_storage_file` resource in order to prevent the same storage being added twice.
* `block_storage_ids` - (Optional, array of numbers) Block storage to which this computing instance should have access. Block storage must be in the same data center as the bare metal server. If you use this argument to authorize access to block storage, do not use the `allowed_hardware_ids` argument in the `ibm_storage_file` resource in order to prevent the same storage being added twice.

### Arguments common to hourly and monthly server

* `datacenter` - (Required, string) The datacenter in which you want to provision the instance.
* `gpu_key_name` - (Optional, string) The key name for the primary Graphics Processing Unit (GPU). For example - `GPU_NVIDIA_GRID_K2`.
Locate your package ID. See `package_key_name` attribute. Once you have the ID fetch its [details](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/<PACKAGE_ID>/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]). Select a gpu key name from the resulting available gpu key names where category code is `gpu0`.
* `gpu_secondary_key_name` - (Optional, string) The key name for the secondary Graphics Processing Unit (GPU). For example - `GPU_NVIDIA_GRID_K2`. Key names can be fetched in the similar way as `gpu_key_name` and  category code is `gpu1`.
* `hourly_billing` - (Required, boolean) The billing type for the instance. When set to `true`, the computing instance is billed on hourly usage. Otherwise the instance is billed on a monthly basis. The default value is `true`.
* `redundant_power_supply` - (Optional, boolean) When the value is `true`, an additional power supply is provided.
* `redundant_network` - (Optional, boolean) When the value is `true`, two physical network interfaces are provided with a bonding configuration. The default value is `false`.
* `unbonded_network` - (Optional, boolean) When the value is `true`, two physical network interfaces are provided without a bonding configuration. The default value is `false`.
* `network_speed` - (Optional, integer) The connection speed, expressed in Mbps,  for the instance's network components. The default value is `100`.
* `private_network_only` - (Optional, boolean) Specifies whether the instance only has access to the private network. When the value is `true`, a compute instance only has access to the private network. The default value is `false`.
* `ipv6_enabled` - (Optional, boolean) The primary public IPv6 address. The default value is `false`.
* `ipv6_static_enabled` - (Optional, boolean) The public static IPv6 address block of `/64`. The default value is `false`.
* `secondary_ip_count` - (Optional, integer) Specifies secondary public IPv4 addresses. Accepted values are `4` and `8`.
* `image_template_id` - (Optional, integer) The image template ID you want to use to provision the computing instance. This is not the global identifier (UUID), but the image template group ID that should point to a valid global identifier. To retrieve the image template ID from the IBM Cloud infrastructure customer portal, navigate to **Devices > Manage > Images**, click the desired image, and note the ID number in the resulting URL.
    **NOTE**: Conflicts with `os_reference_code`. If you don't know the ID(s) of your image templates, you can [reference them by name](../d/compute_image_template.html).

### Arguments for hourly bare metal servers

* `fixed_config_preset` - (Required, string) The configuration preset with which you want to provision the bare metal server. This preset governs the type of CPU, number of cores, amount of RAM, and number of hard drives that the bare metal server has. To see the available presets, log in to the [IBM Cloud Infrastructure (SoftLayer) API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json) using your API key as the password. Find the key called `fixedConfigurationPresets`. The presets are identified by the key names.
* `os_reference_code` - (Optional, string) An operating system reference code that provisions the computing instance. To see available OS reference codes, log in to the [IBM Cloud Infrastructure (SoftLayer) API](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest_Block_Device_Template_Group/getVhdImportSoftwareDescriptions.json?objectMask=referenceCode), using your API key as the password.
    **NOTE**: Conflicts with `image_template_id`.  

### Arguments for monthly bare metal servers

* `public_vlan_id` - (Optional, integer) The public VLAN to be used for the public network interface of the instance. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID in the resulting URL.
* `private_vlan_id` - (Optional, integer) The private VLAN to be used for the private network interface of the instance. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID in the resulting URL.
* `public_subnet` - (Optional, string) The public subnet to be used for the public network interface of the instance. Accepted values are primary public networks. You can find accepted values in the [subnets docs](https://control.softlayer.com/network/subnets).
* `private_subnet` - (Optional, string) The private subnet to be used for the private network interface of the instance. Accepted values are primary private networks. You can find accepted values in the [subnets docs](https://control.softlayer.com/network/subnets).
* `package_key_name` - (Optional, string) The key name for the monthly bare metal server's package. Only use this argument when you create a new monthly bare metal server. You can find available package key names in the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}), using your API key as the password.
* `process_key_name` - (Optional, string). The key name for the monthly bare metal server's process. Only use this argument when you create a new monthly bare metal server. To get a process key name, first find the package key name in the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). Then replace <PACKAGE_NAME> with your package key name in the following URL: `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/<PACKAGE_NAME>/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]`. Select a process key name from the resulting available process key names.
* `disk_key_names` - (Optional, array of strings) The internal key names for the monthly bare metal server's disk. Only use this argument when you create a new monthly bare metal server. To get disk key names, first find the package key name in the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). Then replace <PACKAGE_NAME> with your package key name in the following URL: `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/PACKAGE_NAME/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]`. Select disk key names from the resulting available disk key names.
* `os_key_name` - (Optional, string) The operating system key name that you want to use to provision the computing instance. To get disk key names, first find the package key name in the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). Then replace <PACKAGE_NAME> with your package key name in the following URL: `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/<PACKAGE_NAME>/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]`. Select an OS key name from the resulting available OS key names.
* `public_bandwidth` - (Optional, integer) The amount of public network traffic, specified in gigabytes, allowed per month. The value can be greater than 0 when `private_network_only` is set to `false` and the server is a monthly-based server.
* `memory` - (Optional, integer) The amount of memory, specified in gigabytes, for the server.
* `storage_groups` - (Optional, array of storage group objects) Configurations for RAID and partition. For more information on configuring `storage_groups`, refer to [Ordering RAID Through API](https://sldn.softlayer.com/blog/hansKristian/Ordering-RAID-through-API). Each storage group object has the following sub-arguments:
    * `array_type_id` -(Required, integer) The RAID type. You can retrieve the value from the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Configuration_Storage_Group_Array_Type/getAllObjects).
    * `hard_drives` - (Required, array of integers) The index of hard drives for RAID configuration. The index starts at 0. For example, the array [0,1] is an index of two hard drives.
    * `array_size` - (Optional, integer) The target RAID disk size, specific in gigabytes.
    * `partition_template_id` - (Optional, string) The partition template ID for the OS disk. Templates are different based on the target OS. To get the partition template ID, first find the OS ID in the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware_Component_Partition_OperatingSystem/getAllObjects). Then replace <OS_ID> with your OS ID in the following URL: `https://api.softlayer.com/rest/v3/SoftLayer_Hardware_Component_Partition_OperatingSystem/<OS_ID>/getPartitionTemplates`. Select you template ID in resulting available parition template IDs.  
* `restricted_network` - (Optional, boolean) The non-datacenter restricted port speed. The default value is `false`.
* `tcp_monitoring` - (Optional) When the value is `false`, a ping monitoring service is provided. When the value is `true`, a ping monitoring service and a TCP monitoring service are provided.

### Arguments for quote-based bare metal servers

* `public_vlan_id` - (Optional, integer) The public VLAN to be used for the public network interface of the instance. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID in the resulting URL.
* `private_vlan_id` - (Optional, integer) The private VLAN to be used for the private network interface of the instance. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID in the resulting URL.
* `public_subnet` - (Optional, string) The public subnet to be used for the public network interface of the instance. Accepted values are primary public networks. You can find accepted values in the [subnets docs](https://control.softlayer.com/network/subnets).
* `private_subnet` - (Optional, string) The private subnet to be used for the private network interface of the instance. Accepted values are primary private networks. You can find accepted values in the [subnets docs](https://control.softlayer.com/network/subnets).
* `quote_id` - (Optional, string) When you define `quote_id`, Terraform uses specifications in the quote to create a bare metal server. You can find the quote ID in the [IBM Cloud infrastructure customer portal](https://control.softlayer.com) by navigating to **Account > Sales > Quotes**.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the bare metal server.
* `public_ipv4_address` - The public IPv4 address of the bare metal server.
* `private_ipv4_address` - The private IPv4 address of the bare metal server.
* `ipv6_address` - The public IPv6 address of the bare metal server instance provided when `ipv6_enabled` is set to `true`.
* `secondary_ip_addresses` - The public secondary IPv4 addresses of the bare metal server instance when `secondary_ip_count` is set to non-zero value.
* `global_identifier` - The unique global identifier of the bare metal server.

