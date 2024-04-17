---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Server Profile"
description: |-
  Manages IBM Cloud Bare Metal Server Profile.
---

# ibm\_is_bare_metal_server_profile

Import the details of an existing IBM Cloud Bare Metal Server profile as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal server profile, see [Bare Metal Servers for VPC profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-bare-metal-servers-profile).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
data "ibm_is_bare_metal_server_profile" "ds_bmsprofile" {
  name        = "profile-name"
}

```

## Argument Reference

Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name for this profile .

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `bandwidth` - (List) The total bandwidth (in megabits per second) shared across the network interfaces of a bare metal server with this profile.
  
  Nested scheme for `bandwidth`:
    - `default` - (Integer) The default value for this profile field.
    - `max` - (Integer) The maximum value for this profile field.
    - `min` - (Integer) The minimum value for this profile field.
    - `step` - (Integer) The increment step value for this profile field.
    - `type` - (String) The type for this profile field.
    - `value` - (Integer) The value for this profile field.
    - `values` - (List) The permitted values for this profile field.
- `console_types` - (List) The console type configuration for a bare metal server with this profile.
  
  Nested schema for  `console_types`:
	- `type` - (String) The type for this profile field.
	- `values` - (List) The console types for a bare metal server with this profile.
- `cpu_architecture` - (List) The CPU architecture for a bare metal server with this profile.
  
  Nested scheme for `cpu_architecture`:
    - `type` - (String) The type for this profile field.
    - `value` - (Integer) The value for this profile field.
- `cpu_core_count` - (List) The CPU core count for a bare metal server with this profile.
  
  Nested scheme for `cpu_core_count`:
    - `type` - (String) The type for this profile field.
    - `value` - (Integer) The value for this profile field.
- `cpu_socket_count` - (List) The number of CPU sockets for a bare metal server with this profile.
  
  Nested scheme for `cpu_socket_count`:
    - `type` - (String) The type for this profile field.
    - `value` - (Integer) The value for this profile field.
- `disks` - (List) A nested block describing the collection of the bare metal server profile's disks.
  
  Nested scheme for `disk`:
    - `quantity` - (List) The number of disks of this configuration for a bare metal server with this profile.

      Nested scheme for `quantity`:
        - `type` - (String) The type for this profile field.
        - `value` - (Integer) The value for this profile field.

    - `size` - (List) The size of the disk in GB (gigabytes).

      Nested scheme for `size`:
        - `type` - (String) The type for this profile field.
        - `value` - (Integer) The value for this profile field.
    - `supported_interface_types` - (List) The disk interface used for attaching the disk.
      
      Nested scheme for `supported_interface_types`:

        - `default` - (String) The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
        - `type` - (String) The type for this profile field.
        - `values` - (Array) The supported disk interfaces used for attaching the disk.
- `family` - (String) The product family this bare metal server profile belongs to.
- `href` - (String) The URL for this bare metal server profile.
- `id` - (String) The name of the profile.
- `memory` - (List) The memory (in gibibytes) for a bare metal server with this profile.
  Nested scheme for `memory`:
    - `type` - (String) The type for this profile field.
    - `value` - (String) The value for this profile field.
- `name` - (String) The name of the profile.
- `network_attachment_count` - (List)

  Nested schema for **network_attachment_count**:
    - `max` - (Integer) The maximum value for this profile field.
    - `min` - (Integer) The minimum value for this profile field.
    - `type` - (String) The type for this profile field.

- `network_interface_count` - (List) 
  
  Nested schema for **network_interface_count**:
	- `max` - (Integer) The maximum value for this profile field.
	- `min` - (Integer) The minimum value for this profile field.
	- `type` - (String) The type for this profile field.
- `os_architecture` - (List) The supported OS architecture(s) for a bare metal server with this profile.
  Nested scheme for `os_architecture`:
    - `default` - (String) The default OS architecture for a bare metal server with this profile
    - `type` - (String) The type for this profile field.
    - `values` - (Array) The supported OS architecture(s) for a bare metal server with this profile.
- `resource_type` - (String) The resource type.
- `supported_image_flags` - (Array) An array of flags supported by this bare metal server profile.
- `supported_trusted_platform_module_modes` - (List) An array of supported trusted platform module (TPM) modes for this bare metal server profile.

  Nested scheme for `supported_trusted_platform_module_modes`:
    - `type` - (String) The type for this profile field.
    - `values` - (Array) The supported trusted platform module (TPM) modes.

- `virtual_network_interfaces_supported` - (List) Indicates whether this profile supports virtual network interfaces.

  Nested schema for **virtual_network_interfaces_supported**:
	- `type` - (String) The type for this profile field.
	- `value` - (Boolean) The value for this profile field.

