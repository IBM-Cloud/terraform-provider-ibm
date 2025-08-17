---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Server"
description: |-
  Manages IBM Cloud Bare Metal Server.
---

# ibm_is_bare_metal_server

Import the details of an existing IBM Cloud Bare Metal Server as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal servers, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage with identifier

```terraform
data "ibm_is_bare_metal_server" "example" {
  identifier        = "9328-9849-9849-9849"
}
```

## Example Usage with name
```terraform
data "ibm_is_bare_metal_server" "example" {
  name        = "example-server"
}
```

## Argument Reference

Review the argument references that you can specify for your data source.

- `identifier` - (Optional, String) The id for this bare metal server.
- `name` - (Optional, String) The name for this bare metal server.

  ~> **NOTE**
    `identifier` and `name` are mutually exclusive.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `access_tags`  - (List) Access management tags associated for the bare metal server.
- `bandwidth` - (Integer) The total bandwidth (in megabits per second) shared across the bare metal server's network interfaces.
- `boot_target` - (String) The unique identifier for this bare metal server disk.
- `cpu` - (List) A nested block describing the CPU configuration of this bare metal server.
  Nested scheme for `cpu`:
    - `architecture` - (String) The architecture of the bare metal server.
    - `core_count` - (Integer) The total number of cores
    - `socket_count` - (Integer) The total number of CPU sockets
    - `threads_per_core` - (Integer) The total number of hardware threads per core
- `created_at` - (Timestamp) The date and time that the bare metal server was created.
- `crn` - (String) The CRN for this bare metal server
- `disks` - (List) The disks for this bare metal server, including any disks that are associated with the boot_target.
  Nested scheme for `disks`:
    - `allowed_use` - (List) The usage constraints to be matched against the requested bare metal server properties to determine compatibility.
    
      Nested schema for `allowed_use`:
      - `api_version` - (String) The API version with which to evaluate the expressions.

      - `bare_metal_server` - (String)The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this disk. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
      
      ~> **NOTE** </br> In addition, the following property is supported: </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled for this bare metal server.   
    - `href` - (String) The URL for this bare metal server disk.
    - `id` - (String) The unique identifier for this bare metal server disk.
    - `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered. [ **nvme**, **sata** ]
    - `name` - (String) The user-defined name for this disk
    - `resource_type` - (String) The resource type
    - `size` - (Integer) The size of the disk in GB (gigabytes) 
- `enable_secure_boot` - (Boolean) Indicates whether secure boot is enabled. If enabled, the image must support secure boot or the server will fail to boot.
- `health_reasons` - (List) The reasons for the current health_state (if any).

    Nested scheme for `health_reasons`:
    - `code` - (String) A snake case string succinctly identifying the reason for this health state.
    - `message` - (String) An explanation of the reason for this health state.
    - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.
- `href` - (String) The URL for this bare metal server
- `id` - (String) The unique identifier for this bare metal server
- `image` - (String) Image used in the bare metal server.
- `keys` - (String) Image used in the bare metal server.
- `memory` - (Integer) The amount of memory, truncated to whole gibibytes
- `name` - (String) The name of the bare metal server.
- `network_attachments` - (List) The network attachments for this bare metal server, including the primary network attachment.
  Nested schema for **network_attachments**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this network attachment.
	- `id` - (String) The unique identifier for this network attachment.
	- `name` - (String)
	- `primary_ip` - (List) The primary IP address of the virtual network interface for the network attachment.
	  Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.

	- `reservation`- (List) The reservation used by this bare metal server. 
    Nested scheme for `reservation`:
    - `crn` - (String) The CRN for this reservation.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

      Nested `deleted` blocks have the following structure: 
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this reservation.
    - `id` - (String) The unique identifier for this reservation.
    - `name` - (string) The name for this reservation. The name is unique across all reservations in the region.
    - `resource_type` - (string) The resource type.
  - `reservation_affinity`- (List) The bare metal server reservation affinity. 

    Nested scheme for `reservation_affinity`:
    - `policy` - (String) The reservation affinity policy to use for this bare metal server.
    - `pool` - (List) The pool of reservations available for use by this bare metal server.

      Nested `pool` blocks have the following structure: 
      - `crn` - (String) The CRN for this reservation.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.

        Nested `deleted` blocks have the following structure:
        - `more_info` - (String) Link to documentation about deleted resources. 
      - `href` - (String) The URL for this reservation.
      - `id` - (String) The unique identifier for this reservation.
      - `name` - (string) The name for this reservation. The name is unique across all reservations in the region.
      - `resource_type` - (string) The resource type.
	- `resource_type` - (String) The resource type.
  - `subnet` - (List) The subnet of the virtual network interface for the network attachment.
	  Nested schema for **subnet**:
		- `crn` - (String) The CRN for this subnet.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this subnet.
		- `id` - (String) The unique identifier for this subnet.
		- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
		- `resource_type` - (String) The resource type.
  - `virtual_network_interface` - (List) The virtual network interface for this bare metal server network attachment.
  Nested schema for **virtual_network_interface**:
    - `crn` - (String) The CRN for this virtual network interface.
    - `href` - (String) The URL for this virtual network interface.
    - `id` - (String) The unique identifier for this virtual network interface.
    - `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
    - `resource_type` - (String) The resource type.
    
- `network_interfaces` - (List) A nested block describing the additional network interface of this instance.
  Nested scheme for `network_interfaces`:
    - `allow_ip_spoofing` - (Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
    - `href` - (String) The href of the network interface.
    - `id` - (String) The id of the network interface.
    - `name` - (String) The name of the network interface.
    - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
        - `address` - (String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
        - `href`- (String) The URL for this reserved IP
        - `reserved_ip`- (String) The unique identifier for this reserved IP
        - `name`- (String) The user-defined or system-provided name for this reserved IP
        - `resource_type`- (String) The resource type.

    - `security_groups` -  (Array) List of security groups.
    - `subnet` -  (String) ID of the subnet.
- `primary_network_attachment` - (List) The primary network attachment.
  Nested schema for **primary_network_attachment**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this network attachment.
	- `id` - (String) The unique identifier for this network attachment.
	- `name` - (String)
	- `primary_ip` - (List) The primary IP address of the virtual network interface for the network attachment.
	  Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List) The subnet of the virtual network interface for the network attachment.
	  Nested schema for **subnet**:
		- `crn` - (String) The CRN for this subnet.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this subnet.
		- `id` - (String) The unique identifier for this subnet.
		- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
		- `resource_type` - (String) The resource type.
  - `virtual_network_interface` - (List) The virtual network interface for this bare metal server network attachment.
  Nested schema for **virtual_network_interface**:
    - `crn` - (String) The CRN for this virtual network interface.
    - `href` - (String) The URL for this virtual network interface.
    - `id` - (String) The unique identifier for this virtual network interface.
    - `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
    - `resource_type` - (String) The resource type.
- `primary_network_interface` - (List) A nested block describing the primary network interface of this bare metal server.
  Nested scheme for `primary_network_interface`:
    - `allow_ip_spoofing` - (Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
    - `href` - (String) The href of the network interface.
    - `id` - (String) The id of the network interface.
    - `name` - (String) The name of the network interface.
    - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
        - `address` - (String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
        - `href`- (String) The URL for this reserved IP
        - `reserved_ip`- (String) The unique identifier for this reserved IP
        - `name`- (String) The user-defined or system-provided name for this reserved IP
        - `resource_type`- (String) The resource type.
    - `security_groups` -  (Array) List of security groups.
    - `subnet` -  (String) ID of the subnet.
- `profile` - (String) The name for this bare metal server profile
- `resource_group` - (String) resource group id of the bare metal server.
- `resource_type` - (String) The type of resource referenced
- `firmware_update_type_available` - (String) The firmware update type available for the bare metal server.
  
  -> **Supported firmware update types** 
    </br>&#x2022; none 
    </br>&#x2022; optional 
    </br>&#x2022; required
- `status` - (String) The status of the bare metal server.

  -> **Supported Status** 
    &#x2022; failed
    </br>&#x2022; pending
    </br>&#x2022; restarting
    </br>&#x2022; running
    </br>&#x2022; starting
    </br>&#x2022; stopped
    </br>&#x2022; stopping
    
- `status_reasons` - (List) Array of reasons for the current status (if any).
  Nested scheme for `status_reasons`:
    - `code` - (String) The status reason code
    - `message` - (String) An explanation of the status reason
    - `more_info` - (String) Link to documentation about this status reason
- `tags` - (Array) Tags associated with the instance.
- `trusted_platform_module` - (List) trusted platform module (TPM) configuration for this bare metal server

    Nested scheme for **trusted_platform_module**:

    - `enabled` - (Boolean) Indicates whether the trusted platform module is enabled. 
    - `mode` - (String) The trusted platform module mode to use. The specified value must be listed in the bare metal server profile's supported_trusted_platform_module_modes. Updating trusted_platform_module mode would require the server to be stopped then started again.
      - Constraints: Allowable values are: `disabled`, `tpm_2`.
    - `supported_modes` - (Array) The trusted platform module (TPM) mode:
      - **disabled: No TPM functionality**
      - **tpm_2: TPM 2.0**
      - The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
- `vpc` - (String) The VPC this bare metal server resides in.
- `zone` - (String) The zone this bare metal server resides in.
