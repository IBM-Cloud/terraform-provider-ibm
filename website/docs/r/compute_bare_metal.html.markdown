---
layout: "ibm"
page_title: "IBM: compute_bare_metal"
sidebar_current: "docs-ibm-resource-compute-bare-metal"
description: |-
  Manages IBM Compute bare metal servers.
---

# ibm\_compute_bare_metal

Provides a bare metal resource. This allows bare metal servers to be created, updated, and deleted.

```hcl
# Create a bare metal server
resource "ibm_compute_bare_metal" "twc_terraform_sample" {
    hostname = "twc-terraform-sample-name"
    domain = "bar.example.com"
    os_reference_code = "UBUNTU_16_64"
    datacenter = "dal01"
    network_speed = 100 # Optional
    hourly_billing = true # Optional
    private_network_only = false # Optional
    user_metadata = "{\"value\":\"newvalue\"}" # Optional
    public_vlan_id = 12345678 # Optional
    private_vlan_id = 87654321 # Optional
    public_subnet = "50.97.46.160/28" # Optional
    private_subnet = "10.56.109.128/26" # Optional
    fixed_config_preset = "S1270_8GB_2X1TBSATA_NORAID"
    image_template_id = 12345 # Optional
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
* `datacenter` - (Required, string) The data center the instance is to be provisioned in.
* `fixed_config_preset` - (Required, string) The configuration preset that the bare metal server will be provisioned with. This governs the type of CPU, number of cores, amount of RAM, and hard drives that the bare metal server will have. [Log in to the Bluemix Infrastructure (SoftLayer) API to see the available presets](https://api.softlayer.com/rest/v3/SoftLayer_Hardware/getCreateObjectOptions.json). Use your API key as the password. Log in and find the key called `fixedConfigurationPresets`. The presets are be identified by the key names.
* `hourly_billing` - (Required, boolean) The billing type for the instance. When set to `true`, the computing instance is billed on hourly usage, otherwise it is billed on a monthly basis. Default value: `true`.
* `os_reference_code` - (Optional, string) An operating system reference code that provisions the computing instance. [Log in to the Bluemix Infrastructure (SoftLayer) API to see available OS reference codes](https://api.softlayer.com/rest/v3/SoftLayer_Virtual_Guest_Block_Device_Template_Group/getVhdImportSoftwareDescriptions.json?objectMask=referenceCode). Use your API as the password to log in. 

    **NOTE**: Conflicts with `image_template_id`.  

* `image_template_id` - (Optional, integer) The ID of the image template you want to use to provision the computing instance. This is not the global identifier (UUID), but the image template group ID that should point to a valid global identifier. You can get the image template ID from the SoftLayer Customer Portal. In the portal, navigate to **Devices > Manage > Images**, clock the desired image, and take note of the ID number in the browser URL location.

    **NOTE**: Conflicts with `os_reference_code`. If you don't know the ID(s) of your image templates, [you can reference them by name](../d/compute_image_template.html).

* `network_speed` - (Optional, integer) Specifies the connection speed (in Mbps) for the instance's network components. Default value: `100`.
* `private_network_only` - (Optional, boolean) Specifies whether or not the instance only has access to the private network. When set to `true`, a compute instance only has access to the private network. Default value: `false`.
* `public_vlan_id` - (Optional, integer) Public VLAN to be used for the public network interface of the instance. Accepted values can be found [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID number in the URL.
* `private_vlan_id` - (Optional, integer) Private VLAN to be used for the private network interface of the instance. Accepted values can be found [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID number in the URL.
* `public_subnet` - (Optional, string) Public subnet to be used for the public network interface of the instance. Accepted values are primary public networks and can be found [here](https://control.softlayer.com/network/subnets).
* `private_subnet` - (Optional, string) Private subnet to be used for the private network interface of the instance. Accepted values are primary private networks and can be found [here](https://control.softlayer.com/network/subnets).
* `user_metadata` - (Optional, string) Arbitrary data to be made available to the computing instance.
* `ssh_key_ids` - (Optional, array) SSH key IDs to install on the computing instance upon provisioning.

    **NOTE:** If you don't know the ID(s) for your SSH keys, [you can reference your SSH keys by their labels](../d/compute_ssh_key.html).

* `file_storage_ids` - (Optional) An array of numbers. File storage this computing instance should have access to. File storage need to be in the same data center as the bare metal. If you are using this to authorize access to file storage, then you shouldn't use the `allowed_hardware_ids` argument in the `ibm_storage_file` resource in case `ibm_storage_file` represents the same storage as the one being added to the current bare metal. 
* `block_storage_ids` - (Optional) An array of numbers. Block storage this computing instance should have access to. Block storage need to be in the same data center as the bare metal. If you are using this to authorize access to block storage, then you shouldn't use `allowed_hardware_ids` argument in the `ibm_storage_block` resource in case `ibm_storage_block` represents the same storage as the one being added to the current bare metal. 

* `post_install_script_uri` - (Optional, string) As defined in the [Bluemix Infrastructure (SoftLayer) API docs](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions).
*   `tags` - (Optional, array of strings) Set tags on this bare metal server. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters will be removed.

## Attributes Reference

The following attributes are exported:

* `id` - Identifier of the bare metal server.
* `public_ipv4_address` - Public IPv4 address of the bare metal server.
* `private_ipv4_address` - Private IPv4 address of the bare metal server.
