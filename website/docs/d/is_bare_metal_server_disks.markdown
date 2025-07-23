---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_disks"
description: |-
  Manages IBM Cloud Bare Metal Server Disks.
---

# ibm\_is_bare_metal_server_disks

Import the details of an existing IBM Cloud vBare Metal Server Disk collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal server disks, see [Storage of Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-bare-metal-servers-storage).

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

data "ibm_is_bare_metal_server_disks" "ds_bmserver_disks" {
  bare_metal_server = ibm_is_bare_metal_server.example.id
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `bare_metal_server` - (Required, String) The id for this bare metal server.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `disks` - (List of objects) A list of bare metal server disks. Disk is a block device that is locally attached to the physical server. By default, the listed disks are sorted by their created_at property values, with the newest disk first.

  Nested scheme for `disks`:
  - `allowed_use` - (List) The usage constraints to be matched against the requested bare metal server properties to determine compatibility.
    
    Nested schema for `allowed_use`:
    - `api_version` - (String) The API version with which to evaluate the expressions.
	  
    - `bare_metal_server` - (String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this disk. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
      
      ~> **NOTE** </br> In addition, the following property is supported: </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled for this bare metal server.
    - `href` - (String) The URL for this bare metal server disk.
    - `id` - (String) The unique identifier for this bare metal server disk.
    - `interface_type` - (String) The disk interface used for attaching the disk. Supported values are [ **nvme**, **sata** ].
    - `name` - (String) The user-defined name for this disk.
    - `resource_type` - (String) The resource type.
    - `size` - (String) The size of the disk in GB (gigabytes).      
  - `href` - (String) The URL for this bare metal server disk.
  - `id` - (String) The unique identifier for this bare metal server disk.
  - `interface_type` - (String) The disk interface used for attaching the disk. Supported values are [ **nvme**, **sata** ].
  - `name` - (String) The user-defined name for this disk.
  - `resource_type` - (String) The resource type.
  - `size` - (String) The size of the disk in GB (gigabytes).
