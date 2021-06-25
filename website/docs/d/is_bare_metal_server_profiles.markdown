---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Server Profiles"
description: |-
  Manages IBM Cloud Bare Metal Server Profiles.
---

# ibm\_is_bare_metal_server_profiles

Import the details of an existing IBM Cloud virtual server instances as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_bare_metal_server_profiles" "ds_bmsprofiles" {
}

```

## Attribute Reference

The following attributes are exported:

* `profiles` - List of all bare metal server profiles in the IBM Cloud Infrastructure.
  * `name` - The name of the profile.
  * `bandwidth` - The total bandwidth (in megabits per second) shared across the network interfaces of a bare metal server with this profile.
  Nested `bandwidth` blocks have the following profile:
    * `type` - The type for this profile field.
    * `value` - The value for this profile field.
  * `cpu_architecture` - The CPU architecture for a bare metal server with this profile.
  Nested `cpu_architecture` blocks have the following profile:
    * `type` - The type for this profile field.
    * `value` - The value for this profile field.
  * `cpu_core_count` - The CPU core count for a bare metal server with this profile.
  Nested `cpu_core_count` blocks have the following profile:
    * `type` - The type for this profile field.
    * `value` - The value for this profile field.
  * `cpu_socket_count` - The number of CPU sockets for a bare metal server with this profile.
  Nested `cpu_socket_count` blocks have the following profile:
    * `type` - The type for this profile field.
    * `value` - The value for this profile field.
  * `family` - The product family this bare metal server profile belongs to.
  * `href` - The URL for this bare metal server profile.
  * `memory` - The memory (in gibibytes) for a bare metal server with this profile.
  Nested `memory` blocks have the following profile:
    * `type` - The type for this profile field.
    * `value` - The value for this profile field.
  * `os_architecture` - The supported OS architecture(s) for a bare metal server with this profile.
  Nested `bandwidth` blocks have the following profile:
    * `default` - The default OS architecture for a bare metal server with this profile
    * `type` - The type for this profile field.
    * `values` - The supported OS architecture(s) for a bare metal server with this profile.
  * `resource_type` - The resource type.
  * `supported_image_flags` - An array of flags supported by this bare metal server profile.
  * `supported_trusted_platform_module_modes` - An array of supported trusted platform module (TPM) modes for this bare metal server profile.
  Nested `supported_trusted_platform_module_modes` blocks have the following profile:
    * `type` - The type for this profile field.
    * `values` - The supported trusted platform module (TPM) modes.
  * `disks` - A nested block describing the collection of the bare metal server profile's disks.
  Nested `disk` blocks have the following profile:
    * `quantity` - The number of disks of this configuration for a bare metal server with this profile.
    Nested `quantity` blocks have the following profile:
      * `type` - The type for this profile field.
      * `value` - The value for this profile field.
    * `size` - The size of the disk in GB (gigabytes).
      Nested `size` blocks have the following profile:
      * `type` - The type for this profile field.
      * `value` - The value for this profile field.
    * `supported_interface_types` - The disk interface used for attaching the disk.
      Nested `supported_interface_types` blocks have the following profile:
      * `default` - The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
      * `type` - The type for this profile field.
      * `values` - The supported disk interfaces used for attaching the disk.