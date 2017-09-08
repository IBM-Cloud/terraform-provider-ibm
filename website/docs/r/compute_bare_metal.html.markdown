---
layout: "ibm"
page_title: "IBM: compute_bare_metal"
sidebar_current: "docs-ibm-resource-compute-bare-metal"
description: |-
  Manages IBM Compute bare metal servers.
---

# ibm\_compute_bare_metal

Provides a bare metal resource. This allows bare metal servers to be created, updated, and deleted. This resource supports both monthly bare metal servers and hourly bare metal servers. For more detail on bare metal seves, refer to the [link](https://www.ibm.com/cloud-computing/bluemix/bare-metal-servers)


## Hourly bare metal server
If the `ibm_compute_bare_metal` resource definition has a `fixed_config_preset` attribute, terraform will create an hourly
bare metal server. The following example creates a new hourly bare metal server. Hardware specifications 
are already defined in the `fixed_config_preset` attribute and cannot be modified.

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
}
```

In addition, users can use configure optional attributes such as `user_metadata`, `tags`, and `notes` attributes as follows:

### Example of additional attributes for the hourly bare metal server
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
If the `fixed_config_preset` attribute is not configured, terraform will consider it as a monthly bare metal server resource. It provides 
options to configure process, memory, network, disk, and RAID. Users also can assign VLANs and subnets for the target monthly bare metal server. To configure the monthly bare 
metal server, you need to provide additional attributes such as `package_key_name`, `proecss_key_name`, `disk_key_names`, and `os_key_name`. The following example describes a basic configuration
 of the monthly bare metal server.

### Example of a monthly bare metal server
```hcl
resource "ibm_compute_bare_metal" "monthly_bm1" {
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
}
```

Users can configure additional options. The following example configures target VLANs, subnets, and a RAID controller. `storage_groups` 
configures RAIDs and disk partitioning.
### Example of a monthly bare metal server with additional options
```hcl
resource "ibm_compute_bare_metal" "monthly_bm1" {

# Mandatory attributes
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

# Optional attributes
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
       # Basic partition template for windows
       partition_template_id = 17
    }
}
```

_Please Note_: Monthly bare metal servers does not support `immediate cancellation`. If the monthly bare metal server is deleted by terraform, `anniversary date cancellation` option will be used. 

## Create a bare metal server using quote ID
If users already have a quote id for the bare metal server, they can create a new bare metal server with the quote id. The following example describes a basic configuration for a bare metal server with 
 quote_id.
  
### Example of a quote based ordering
```hcl
# Create a new bare metal
resource "ibm_compute_bare_metal" "quote_test" {
    hostname = "quote-bm-test"
    domain = "example.com"
    quote_id = 2209349 
}
```

Users can use additional options when they create a new bare metal server with `quote_id`. The following example defines target VLANs, subnets, 
 user metadata, and tags additionally. 
 
### Example of a quote based ordering with additional options
```hcl
# Create a new bare metal
resource "ibm_compute_bare_metal" "quote_test" {

# Mandatory attributes
    hostname = "quote-bm-test"
    domain = "example.com"
    quote_id = 2209349

# Optional attributes
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

* `hostname` - (Optional, string) Hostname for the computing instance.
* `domain` - (Required, string) Domain for the computing instance.
* `user_metadata` - (Optional, string) Arbitrary data to be made available to the computing instance.
* `notes` - (Optional,string) Specifies a note to associate with the instance.
* `ssh_key_ids` - (Optional, array) SSH key IDs to install on the computing instance upon provisioning.

    **NOTE:** If you don't know the ID(s) for your SSH keys, [you can reference your SSH keys by their labels](../d/compute_ssh_key.html).
* `post_install_script_uri` - (Optional, string) As defined in the [Bluemix Infrastructure (SoftLayer) API docs](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions).
*  `tags` - (Optional, array of strings) Set tags on this bare metal server. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters will be removed.
* `file_storage_ids` - (Optional) An array of numbers. File storage this computing instance should have access to. File storage need to be in the same data center as the bare metal. If you are using this to authorize access to file storage, then you shouldn't use the `allowed_hardware_ids` argument in the `ibm_storage_file` resource in case `ibm_storage_file` represents the same storage as the one being added to the current bare metal. 
* `block_storage_ids` - (Optional) An array of numbers. Block storage this computing instance should have access to. Block storage need to be in the same data center as the bare metal. If you are using this to authorize access to block storage, then you shouldn't use `allowed_hardware_ids` argument in the `ibm_storage_block` resource in case `ibm_storage_block` represents the same storage as the one being added to the current bare metal. 

**Monthly/Hourly bare metal server attributes**

* `datacenter` - (Required, string) The data center the instance is to be provisioned in.
* `hourly_billing` - (Required, boolean) The billing type for the instance. When set to `true`, the computing instance is billed on hourly usage, otherwise it is billed on a monthly basis. Default value: `true`.
* `image_template_id` - (Optional, integer) The ID of the image template you want to use to provision the computing instance. This is not the global identifier (UUID), but the image template group ID that should point to a valid global identifier. You can get the image template ID from the SoftLayer Customer Portal. In the portal, navigate to **Devices > Manage > Images**, clock the desired image, and take note of the ID number in the browser URL location.

    **NOTE**: Conflicts with `os_reference_code`. If you don't know the ID(s) of your image templates, [you can reference them by name](../d/compute_image_template.html).
* `network_speed` - (Optional, integer) Specifies the connection speed (in Mbps) for the instance's network components. Default value: `100`.
* `private_network_only` - (Optional, boolean) Specifies whether or not the instance only has access to the private network. When set to `true`, a compute instance only has access to the private network. Default value: `false`.

**Hourly bare metal server only attributes**

* `fixed_config_preset` - (Required, string) The configuration preset that the bare metal server will be provisioned with. This governs the type of CPU, number of cores, amount of RAM, and hard drives that the bare metal server will have. [Log in to the Bluemix Infrastructure (SoftLayer) API to see the available presets](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json). Use your API key as the password. Log in and find the key called `fixedConfigurationPresets`. The presets are be identified by the key names.

* `os_reference_code` - (Optional, string) An operating system reference code that provisions the computing instance. [Log in to the Bluemix Infrastructure (SoftLayer) API to see available OS reference codes](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest_Block_Device_Template_Group/getVhdImportSoftwareDescriptions.json?objectMask=referenceCode). Use your API as the password to log in. 

    **NOTE**: Conflicts with `image_template_id`.  

**Monthly / Quote based bare metal server provisioning attributes**

* `public_vlan_id` - (Optional, integer) Public VLAN to be used for the public network interface of the instance. Accepted values can be found [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID number in the URL.
* `private_vlan_id` - (Optional, integer) Private VLAN to be used for the private network interface of the instance. Accepted values can be found [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID number in the URL.
* `public_subnet` - (Optional, string) Public subnet to be used for the public network interface of the instance. Accepted values are primary public networks and can be found [here](https://control.softlayer.com/network/subnets).
* `private_subnet` - (Optional, string) Private subnet to be used for the private network interface of the instance. Accepted values are primary private networks and can be found [here](https://control.softlayer.com/network/subnets).


**Monthly bare metal server only attributes**

* `package_key_name` - (Optional, string). Monthly bare metal server's package key name. This attribute is only used when a new monthly bare metal server is created. You can find available key names in the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}). You need your username and api_key to access to the page. 
* `process_key_name` - (Optional, string). Monthly bare metal server's process key name. This attribute is only used when a new monthly bare metal server is created. Note the package key ID from the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}) and replace **PACKAGE_ID** in the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/PACKAGE_ID/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]) to your package key ID. Select a process key name from available process key names.
* `disk_key_names` - (Optional) An array of internal disk key names. This attribute is only used when a new monthly bare metal server is created. Note the package key ID from the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}) and replace **PACKAGE_ID** in the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/PACKAGE_ID/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]) to your package key ID. Select a disk key name from available disk key names.
* `os_key_name` - (Optional, string) An operating system key name that will be used to provision the computing instance. Note the package key ID from the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_CPU"}}}) and replace **PACKAGE_ID** in the [link](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/PACKAGE_ID/getItems?objectMask=mask[prices[id,categories[id,name,categoryCode],capacityRestrictionType,capacityRestrictionMinimum,capacityRestrictionMaximum,locationGroupId]]) to your package key ID. Select a OS key name from available OS key names.
* `redundant_network` -(Optional). If `redundant_network` is `true`, two physical network interfaces will be provided with a bonding configuration. Default value is `False`
* `unbonded_network` - (Optional). If `unbonded_network` is `true`, two physical network interfaces will be provided.Default value is `False`
* `public_bandwidth` - (Optional, int). Allowed public network traffic(GB) per month. It can be greater than 0 when `private_network_only` is `false` and the server is a monthly based server.
* `memory` - (Optional). An amount of memory(GB) for the server.
* `storage_groups` - (Optional) An array of storage group objects.RAID and partition configuration. Refer to the [link](https://sldn.softlayer.com/blog/hansKristian/Ordering-RAID-through-API) to configure `storage_groups`. Each storage group object has the following sub-attributes:
    * `array_type_id` -(Required, int). It provides RAID type. You can find `array_type_id` from the [link](https://api.softlayer.com/rest/v3/SoftLayer_Configuration_Storage_Group_Array_Type/getAllObjects). 
    * `hard_drives` - (Required) An array of integers. Index of hard drives for RAID configuration. The index starts from 0. For example, if you want to use first two hard drives, you will use the following expression: [0,1]
    * `array_size`-(Optional, int) Target RAID disk size in GB unit. 
    * `partition_template_id` - (Optional) Partition template id for OS disk. The templates are different based on the target OS. Check your OS with the [link](https://api.softlayer.com/rest/v3/SoftLayer_Hardware_Component_Partition_OperatingSystem/getAllObjects ). Note the id of the OS and  check available partition templates using the URL : https://api.softlayer.com/rest/v3/SoftLayer_Hardware_Component_Partition_OperatingSystem/OS_ID/getPartitionTemplates . Replace `OS_ID` to your OS ID from the URL and find your template id.  
    
* `redundant_power_supply` - (Optional) If `redundant_power_supply` is true, an additional power supply will be provided. 
* `tcp_monitoring` - (Optional) If `tcp_monitoring` is `false`, ping monitoring service will be provided. If `tcp_monitoring` is `true`, ping and tcp monitoring service will be provided.

**Quote based provisioning only attributes**

* `quote_id` -(Optional). Create a bare metal server using the quote. If quote_id is defined, the terraform uses specifications in the quote to create a bare metal server.You can find the quote id by navigating on the portal to _Account > Sales > Quotes_.

## Attributes Reference

The following attributes are exported:

* `id` - Identifier of the bare metal server.
* `public_ipv4_address` - Public IPv4 address of the bare metal server.
* `private_ipv4_address` - Private IPv4 address of the bare metal server.
