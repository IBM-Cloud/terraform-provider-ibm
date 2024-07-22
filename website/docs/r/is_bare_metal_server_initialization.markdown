---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_initialization"
description: |-
  Reloads the IBM bare metal sever operating system.
---

# ibm\_is_bare_metal_server_disk

Reload a Bare Metal Server with the existing image, keys and user data. This is a one time action resource, which would reload the OS on the bare metal server with image and keys (with/without user_data). For multiple reload, multiple `ibm_is_bare_metal_server_initialization` resource need to be used. For more information, about managing VPC Bare Metal Server, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

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
resource "ibm_is_bare_metal_server_initialization" "initialization" {
  bare_metal_server   = ibm_is_bare_metal_server.bms.id
  image               = var.image_id
  keys                = [ var.keys ]
  user_data           = var.userdata
}
```

## Argument Reference

Review the argument references that you can specify for your resource. 


- `bare_metal_server` - (Required, String) Bare metal server identifier. 
- `image` - (Required, String) Image id to use to reinitialize the bare metal server. 
- `keys` - (Required, Array) Keys ids to use to reinitialize the bare metal server. 
- `user_data` - (Optional, String) User data to transfer to the server bare metal server. (For reload provide the same user data as at the time of provisioning)
