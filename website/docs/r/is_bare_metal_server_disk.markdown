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

