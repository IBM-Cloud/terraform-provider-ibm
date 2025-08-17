---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_disk"
description: |-
  Manages IBM bare metal sever disk name.
---

# ibm\_is_bare_metal_server_disk

Rename a Bare Metal Server for disk. For more information, about managing VPC Bare Metal Server, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

In the following example, you can update name of a Bare Metal Server disk:

```terraform
resource "ibm_is_bare_metal_server_disk" "disk" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  disk              = ibm_is_bare_metal_server.bms.disks.0.id
  name              = "name1"
}
```

## Argument Reference

Review the argument references that you can specify for your resource. 


- `bare_metal_server` - (Required, String) Bare metal server identifier. 
- `disk` - (Required, String) The unique identifier for the disk to be renamed on the  Bare metal server.
- `name` - (Optional, String) The name for the disk.


## Attribute reference
- `allowed_use` - (List) The usage constraints to be matched against the requested bare metal server properties to determine compatibility.
    
    Nested schema for `allowed_use`:
    - `api_version` - (String) The API version with which to evaluate the expressions. If specified, the value must be between 2019-01-01 and today's date (in UTC). If unspecified, the version query parameter value will be used.
	  
    - `bare_metal_server` - (String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this disk..The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
    ~> **NOTE** </br> In addition, the following property is supported: </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled for this bare metal server.
- `href` - (String) The URL for this bare metal server disk.
- `id` - (String) The unique identifier for this bare metal server disk.
- `interface_type` - (String) The disk interface used for attaching the disk. Supported values are [ **nvme**, **sata** ].
- `name` - (String) The user-defined name for this disk.
- `resource_type` - (String) The resource type.
- `size` - (String) The size of the disk in GB (gigabytes).