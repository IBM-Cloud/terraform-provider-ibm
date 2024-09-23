---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: compute_bare_metal"
description: |-
  Manages IBM Cloud compute bare metal servers.
---

# ibm_compute_bare_metal
Create, update, and delete a bare metal resource. This resource supports both monthly bare metal servers and hourly bare metal servers. For more detail on bare metal severs, refer to the [IBM Cloud bare metal server page](https://www.ibm.com/cloud/bare-metal-servers).

## Example usage

### Example of an hourly bare metal server
When the `ibm_compute_bare_metal` resource definition has a `fixed_config_preset` attribute, Terraform creates an hourly bare metal server. Hardware specifications are predefined in the `fixed_config_preset` attribute and cannot be modified. The following example shows you how to create a new hourly bare metal server.

```terraform
resource "ibm_compute_bare_metal" "hourly-bm1" {
  hostname             = "hourly-bm1"
  domain               = "example.com"
  os_reference_code    = "UBUNTU_16_64"
  datacenter           = "dal01"
  network_speed        = 100   # Optional
  hourly_billing       = true  # Optional
  private_network_only = false # Optional
  fixed_config_preset  = "S1270_8GB_2X1TBSATA_NORAID"

  user_metadata = "{\"value\":\"newvalue\"}" # Optional
  tags = [
    "collectd",
    "mesos-master",
  ]
  notes = "note test"
}
```

### Example of a monthly bare metal server
When the `fixed_config_preset` attribute is not configured, Terraform creates a monthly bare metal server resource. The monthly bare metal server resource provides options to configure process, memory, network, disk, and RAID. You can also can assign VLANs and subnets for the target monthly bare metal server. To configure the monthly bare metal server, you must provide additional attributes such as `package_key_name`, `proecss_key_name`, `disk_key_names`, and `os_key_name`. The following example shows you how to create a new monthly bare metal server.

```terraform
resource "ibm_compute_bare_metal" "monthly_bm1" {
  # Mandatory fields
  package_key_name = "DUAL_E52600_V4_12_DRIVES"
  process_key_name = "INTEL_INTEL_XEON_E52620_V4_2_10"
  memory           = 64
  os_key_name      = "OS_WINDOWS_2012_R2_FULL_DC_64_BIT_2"
  hostname         = "cust-bm"
  domain           = "ms.com"
  datacenter       = "wdc04"
  network_speed    = 100
  public_bandwidth = 500
  disk_key_names   = ["HARD_DRIVE_800GB_SSD", "HARD_DRIVE_800GB_SSD", "HARD_DRIVE_800GB_SSD"]
  hourly_billing   = false

  # Optional fields
  private_network_only = false
  unbonded_network     = true
  user_metadata        = "{\"value\":\"newvalue\"}"
  public_vlan_id       = 12345678
  private_vlan_id      = 87654321
  public_subnet        = "50.97.46.160/28"
  private_subnet       = "10.56.109.128/26"
  tags = [
    "collectd",
    "mesos-master",
  ]
  redundant_power_supply = true
  storage_groups {
    # RAID 5
    array_type_id = 3

    # Use three disks
    hard_drives = [0, 1, 2]
    array_size  = 1600

    # Basic partition template for Windows
    partition_template_id = 17
  }
}

```

**Note**

Monthly bare metal servers do not support `immediate cancellation`. When Terraform deletes the monthly bare metal server, the `anniversary date cancellation` option is used.

### Example to create a bare metal server using quote ID based ordering
If you already have a quote ID for the bare metal server, you can create a new bare metal server with the quote ID. The following example shows you how to create a new bare metal server with a quote ID.


```terraform
resource "ibm_compute_bare_metal" "quote_test" {
  # Mandatory fields
  hostname = "quote-bm-test"
  domain   = "example.com"
  quote_id = 2209349

  # Optional fields
  user_metadata   = "{\"value\":\"newvalue\"}"
  public_vlan_id  = 12345678
  private_vlan_id = 87654321
  public_subnet   = "50.97.46.160/28"
  private_subnet  = "10.56.109.128/26"
  tags = [
    "collectd",
    "mesos-master",
  ]
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

### Argument reference for all the bare metal server types 

- `block_storage_ids`- (Optional, Array of Integers) Block storage to which this computing instance has access. Block storage must be in the same data center as the Bare Metal server. If you use this argument to authorize, access to block storage, do not use the `allowed_hardware_ids` argument in the `ibm_storage_file` resource in order to prevent the same storage be added twice.
- `domain` - (Required, Forces new resource, String) The domain for the computing instance.
- `file_storage_ids`- (Optional, Array of Integers) File storage to which this computing instance has access. File storage must be in the same data center as the Bare Metal server. If you use this argument to authorize, access to file storage, do not use the `allowed_hardware_ids` argument in the `ibm_storage_file` resource in order to prevent the same storage be added twice.
- `hostname` - (Optional, Forces new resource, String) The host name for the compute instance.
- `notes` - (Optional, String) Notes to associate with the instance.
- `post_install_script_uri` - (Optional, Forces new resource, String) The URI of the script to be downloaded and executed after installation is complete.
- `ssh_key_ids`- (Optional, Forces new resources, Array of Integers) The SSH key IDs to install on the compute instance when the instance is provisioned. **Note** If you don't know the IDs for your SSH keys, you can reference your SSH keys by their labels.
- `tags` (Optional, Array of Strings) Tags associated with this Bare Metal server. Permitted characters include A-Z, 0-9, whitespace, `_` (underscore), `- ` (hyphen), `.` (period), and `:` (colon). All other characters are removed.
- `user_metadata` - (Optional, Forces new resource, String) Arbitrary data to be made available to the compute instance.

### Arguments reference common to hourly and monthly server

- `datacenter` - (Required, String) The data center in which you want to provision the instance.
- `extended_hardware_testing` - (Optional, Bool) Enable the extended hardware testing while the Bare Metal server. The default value is **false**. **Note** Enabling the `extended_hardware_testing` cause considerable delays in the deployment.
- `gpu_key_name` - (Optional, String) The key name for the primary Graphics Processing Unit (GPU). For example, `GPU_NVIDIA_GRID_K2`. Locate your package ID. See `package_key_name` attribute. To fetch the `PACKAGE_ID`, you need to access [Package ID](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={%22type%22:{%22keyName%22:{%22operation%22:%22BARE_METAL_CPU%22}}}) URL to view the `ID`. Once you have the ID, for example provide `PACKAGE_ID` as `142`, fetch its details by accessing the URL `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/142/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]`. Select a `GPU` key name from the resulting available `GPU` key names where category code is `gpu0`. **Note** To view the package ID. Log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json) that uses your API key as the password. For more information, about creating classic infrastructure keys and locating your VPN username and password, refer [Managing classic infrastructure API keys](https://cloud.ibm.com/docs/account?topic=account-classic_keys).
- `gpu_secondary_key_name` - (Optional, String) The key name for the secondary Graphics Processing Unit (GPU). For example, `GPU_NVIDIA_GRID_K2`. Key names can be fetched in the similar way as `gpu_key_name` and  category code is `gpu1`.
- `hourly_billing`- (Required, Bool) The billing type for the instance. When set to **true**, the compute instance is billed on hourly usage. Otherwise, the instance is billed monthly. The default value is **true**.
- `image_template_id` - (Optional, Integer) The image template ID you want to use to provision the computing instance. This is not the global identifier (UUID), but the image template group ID that should point to a valid global identifier. To retrieve the image template ID from the IBM Cloud infrastructure customer portal, navigate to **Devices > Manage > Images**, click the image that you want, and note the ID number in the resulting URL.       **Note** Conflicts with `os_reference_code`.
- `ipv6_enabled` - (Optional, Bool) The primary public IPv6 address. The default value is **false**.
- `ipv6_static_enabled` - (Optional, Bool) The public static IPv6 address block of `/64`. The default value is **false**.
- `network_speed` - (Optional, Integer) The connection speed, expressed in Mbps,  for the instance's network components. The default value is `100`.
- `private_network_only` - (Optional, Bool) Specifies whether the instance has only access to the private network. When the value is **true**, a compute instance has only access to the private network. The default value is **false**.
- `secondary_ip_count` - (Optional, Integer) Specifies secondary public IPv4 addresses. Accepted values are `4` and `8`.
- `redundant_power_supply` - (Optional, Bool) When the value is **true**, more power supply is provided.
- `redundant_network` - (Optional, Bool) When the value is **true**, two physical network interfaces are provided with a bonding configuration. The default value is **false**.
- `unbonded_network` - (Optional, Bool) When the value is **true**, two physical network interfaces are provided without a bonding configuration. The default value is **false**.

### Argument reference for hourly bare metal servers

- `fixed_config_preset` - (Required, String) The configuration preset with which you want to provision the Bare Metal server. This preset governs the type of CPU, number of cores, amount of RAM, and number of hard disks that the Bare Metal server has. To see the available presets, log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json) that uses your API key as the password. For more information, about creating classic infrastructure keys and locating your VPN username and password, refer [Managing classic infrastructure API keys](https://cloud.ibm.com/docs/account?topic=account-classic_keys).  You can search for `fixedConfigurationPresets` to view the presets key names. **Note** The `fixedConfigurationPresets` key names are displayed in the `JSON` or `txt` format.
- `os_reference_code` - (Optional, String) An operating system reference code that provisions the computing instance. To see available OS reference codes, log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest_Block_Device_Template_Group/getVhdImportSoftwareDescriptions.json?objectMask=referenceCode), that uses your API key as the password.  **Note** Conflicts with `image_template_id`.
- `software_guard_extensions` - (Optional, Bool) The Software Guard Extensions product is added to a compatible server package, selecting `Intel SGX-enabled BIOS` and `hardware`. The default value is **false**.

### Argument reference for monthly bare metal servers

- `backend_network_component` - (Optional, List of Objects)Configurations for backend network components.

  Nested scheme for `backend_network_component`:
  - `vlan_id`- (Required, Integer) The private vlan id.
  - `subent_id` - (Required, Integer) The private subnet id.
- `disk_key_names` (Optional, Array of Strings) The internal key names for the monthly Bare Metal server's disk. Use this argument when you create a new monthly Bare Metal server. To get disk key names, first find the package key name in the [IBM Cloud API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). Then, replace <PACKAGE_NAME> with your package key name in the following [URL](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/PACKAGE_NAME/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]). Select disk key names from the resulting available disk key names.
- `memory` - (Optional, Integer) The amount of memory, which is specified in gigabytes, for the server.
- `os_key_name` - (Optional, String) The operating system key name that you want to use to provision the computing instance. To get disk key names, first find the package key name in the [IBM Cloud API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). Then, replace <PACKAGE_NAME> with your package key name in the following [URL](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/<PACKAGE_NAME>/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]). Select an OS key name from the resulting available OS key names.
- `public_vlan_id` - (Optional, Integer) The public VLAN to be used for the public network interface of the instance. You can find accepted values in the [VLAN networks](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want and notes the ID in the resulting URL.
- `private_vlan_id` - (Optional, Integer) The private VLAN to be used for the private network interface of the instance. You can find accepted values in the [VLAN networks](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want and notes the ID in the resulting URL.
- `public_subnet` - (Optional, String) The public subnet to be used for the public network interface of the instance. Accepted values are primary public networks. You can find accepted values in the [subnets Docs](https://cloud.ibm.com/classic/network/subnets).
- `private_subnet` - (Optional, String) The private subnet to be used for the private network interface of the instance. Accepted values are primary private networks. You can find accepted values in the [subnets Docs](https://cloud.ibm.com/classic/network/subnets).
- `package_key_name` - (Optional, String) The key name for the monthly Bare Metal server's package.Use this argument when you create a new monthly Bare Metal server. You can find available package key names in the IBM Cloud API URL `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}`, that uses your API key as the password.
- `process_key_name` - (Optional, String) The key name for the monthly Bare Metal server's process. Use this argument when you create a new monthly Bare Metal server. To get a process key name, first find the package key name in the [IBM Cloud API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). To fetch the `PACKAGE_ID`, you need to access [Package ID](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={%22type%22:{%22keyName%22:{%22operation%22:%22BARE_METAL_CPU%22}}}) URL to view the `ID`. Once you have the ID, for example provide `PACKAGE_ID` as `142`. Then, replace <PACKAGE_ID> with your package key name in the following URL `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/142/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]`. Select a process key name from the resulting available process key names. **Note** To view the package ID. log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json) that uses your API key as the password. For more information, about creating classic infrastructure keys and locating your VPN username and password, refer [Managing classic infrastructure API keys](https://cloud.ibm.com/docs/account?topic=account-classic_keys).
- `public_bandwidth` - (Optional, Integer) The amount of public network traffic, which is specified in gigabytes, allowed per month. The value can be greater than 0 when `private_network_only` is set to **false** and the server is a monthly based server.
- `restricted_network` - (Optional, Bool) The non-datacenter restricted port speed. The default value is **false**.
- `storage_groups` - (Optional, List of Objects)Configurations for RAID and partition.

  Nested scheme for `storage-groups`:
  - `array_type_id`- (Required, Integer) The RAID type. You can retrieve the value from the [IBM Cloud API](https://api.softlayer.com/rest/v3/SoftLayer_Configuration_Storage_Group_Array_Type/getAllObjects).
  - `array_size` - (Optional, Integer) The target RAID disk size, specific in gigabytes.
  - `hard_drives`-Array of integers-Required-The index of hard disks for RAID configuration. The index starts at 0. For example, the array [0,1] is an index of two hard disks.
  - `hard_drive_category_codes`-Array of strings allowing the category codes to be specified instead of just a disk index 
  - `partition_template_id` - (Optional, String) The partition template ID for the OS disk. Templates are different based on the target OS. To get the partition template ID, first find the OS ID in the [IBM Cloud API](https://api.softlayer.com/rest/v3/SoftLayer_Hardware_Component_Partition_OperatingSystem/getAllObjects). Then, replace <OS_ID> with your OS ID in the following URL `https://api.softlayer.com/rest/v3/SoftLayer_Hardware_Component_Partition_OperatingSystem/<OS_ID>/getPartitionTemplates`. Select your template ID in resulting available partition template IDs.
- `software_guard_extensions` - (Optional, Bool) The Software Guard Extensions product is added to a compatible server package, selecting Intel SGX-enabled BIOS and hardware. The default value is **false**.
- `tcp_monitoring` - (Optional, Bool)  When the value is **false**, a ping monitoring service is provided. When the value is **true**, a ping monitoring service and a TCP monitoring service are provided.#### Arguments for quote-based Bare Metal servers-

### Argument reference for quote-based bare metal servers

- `public_vlan_id` - (Optional, Integer) The public VLAN to be used for the public network interface of the instance. You can find accepted values in the [VLAN network](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want and notes the ID in the resulting URL.
- `private_vlan_id` - (Optional, Integer) The private VLAN to be used for the private network interface of the instance. You can find accepted values in the [VLAN network](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want and notes the ID in the resulting URL.
- `public_subnet` - (Optional, String) The public subnet to be used for the public network interface of the instance. Accepted values are primary public networks. You can find accepted values in the [subnets Docs](https://cloud.ibm.com/classic/network/subnets).
- `private_subnet` - (Optional, String) The private subnet to be used for the private network interface of the instance. Accepted values are primary private networks. You can find accepted values in the [subnets Docs](https://cloud.ibm.com/classic/network/subnets).
- `quote_id` - (Optional, String) When you define `quote_id`,  Terraform uses specifications in the quote to create a Bare Metal server. You can find the quote ID in the [IBM Cloud infrastructure customer portal](https://cloud.ibm.com/classic) by navigating to **Account > Sales > Quotes**.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `global_identifier` - (String) The unique global identifier of the Bare Metal server.
- `id` - (String) The unique identifier of the Bare Metal server.
- `ipv6_address` - (String) The public IPv6 address of the Bare Metal server instance provided when `ipv6_enabled` is set to **true**.
- `ipv6_address_id` - (String) The unique identifier for the public IPv6 address of the Bare Metal server.
- `public_ipv4_address` - (String) The public IPv4 address of the Bare Metal server.
- `public_ipv4_address_id` - (String) The unique identifier for the public IPv4 address of the Bare Metal server.
- `private_ipv4_address` - (String) The private IPv4 address of the Bare Metal server.
- `private_ipv4_address_id` - (String) The unique identifier for the private IPv4 address of the Bare Metal server.
- `secondary_ip_addresses` - (String) The public secondary IPv4 addresses of the Bare Metal server instance when `secondary_ip_count` is set to non zero value.

## Import

The `ibm_compute_bare_metal` resource can be imported by using Bare Metal server ID.

**Example**

```
$ terraform import ibm_compute_bare_metal.server <server_id>
```
