---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Servers"
description: |-
  Manages IBM Cloud Bare Metal Servers.
---

# ibm\_is_bare_metal_servers

Import the details of an existing IBM Cloud vBare Metal Server collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal servers, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

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
data "ibm_is_bare_metal_servers" "example" {
}
```

## Attribute Reference

Review the attribute references that you can access after you retrieve your data source. 

- `servers` - List of all all bare metal servers in the region.
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
      - `href` - (String) The URL for this bare metal server disk.
      - `id` - (String) The unique identifier for this bare metal server disk.
      - `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered. [ **nvme**, **sata** ]
      - `name` - (String) The user-defined name for this disk
      - `resource_type` - (String) The resource type
      - `size` - (Integer) The size of the disk in GB (gigabytes)
  - `href` - (String) The URL for this bare metal server
  - `id` - (String) The unique identifier for this bare metal server
  - `image` - (String) Image used in the bare metal server.
  - `keys` - (String) Image used in the bare metal server.
  - `memory` - (Integer) The amount of memory, truncated to whole gibibytes
  - `name` - (String) The name of the bare metal server.
  - `network_interfaces` - (List) A nested block describing the additional network interface of this instance.
    Nested scheme for `network_interfaces`:
      - `allow_ip_spoofing` - (Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
      - `href` - (String) The href of the network interface.
      - `id` - (String) The id of the network interface.
      - `name` - (String) The name of the network interface.
      - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
        Nested scheme for `primary_ip`:
          - `address` - (String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `security_groups` -  (Array) List of security groups.
      - `subnet` -  (String) ID of the subnet.
  - `primary_network_interface` - (List) A nested block describing the primary network interface of this bare metal server.
    Nested scheme for `primary_network_interface`:
      - `allow_ip_spoofing` - (Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
      - `href` - (String) The href of the network interface.
      - `id` - (String) The id of the network interface.
      - `name` - (String) The name of the network interface.
      - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
        Nested scheme for `primary_ip`:
          - `address` - (String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `security_groups` -  (Array) List of security groups.
      - `subnet` -  (String) ID of the subnet.
  - `profile` - (String) The name for this bare metal server profile
  - `resource_group` - (String) resource group id of the bare metal server.
  - `resource_type` - (String) The type of resource referenced
  - `status` - (String) The status of the bare metal server [ **failed**, **pending**, **restarting**, **running**, **starting**, **stopped**, **stopping** ]
  - `status_reasons` - (List) Array of reasons for the current status (if any).  
    Nested scheme for `status_reasons`:
      - `code` - (String) The status reason code
      - `message` - (String) An explanation of the status reason
  - `tags` - (Array) Tags associated with the instance.
  - `vpc` - (String) The VPC this bare metal server resides in.
  - `zone` - (String) The zone this bare metal server resides in.
