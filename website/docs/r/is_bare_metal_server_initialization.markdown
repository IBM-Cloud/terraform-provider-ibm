---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_initialization"
description: |-
  Replaces the IBM bare metal sever initialization.
---

# ibm\_is_bare_metal_server_initialization

Reinitialize a Bare Metal Server with the existing image, keys and user data. This is a one time action resource, which would reinitialize/reload/replace the OS, keys, user_data on the bare metal server with image and keys (with/without user_data). [Read more about Bare Metal Servers reinitialization](https://cloud.ibm.com/apidocs/vpc/latest#replace-bare-metal-server-initialization). For multiple reload, multiple `ibm_is_bare_metal_server_initialization` resource need to be used. For more information, about managing VPC Bare Metal Server, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

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
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  image             = var.image_id
  keys              = [var.keys]
  user_data         = var.userdata
}

## to avoid changes on the ibm_is_bare_metal_server resource, use lifecycle meta argument ignore_changes
resource "ibm_is_bare_metal_server" "bms" {
  ....
  lifecycle{
    ignore_changes = [ image, keys, user_data ]
  }
}
```

## Default trusted profile Example

In the following example, you can update default trusted profile of a Bare Metal Server disk:

```terraform
resource "ibm_is_bare_metal_server_initialization" "initialization" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id
  image             = var.image_id
  keys              = [var.keys]
  user_data         = var.userdata
  default_trusted_profile {
    auto_link = true
    target {
      id = var.profile_id
    }
  }
}

## to avoid changes on the ibm_is_bare_metal_server resource, use lifecycle meta argument ignore_changes
resource "ibm_is_bare_metal_server" "bms" {
  ....
  lifecycle{
    ignore_changes = [ image, keys, user_data, default_trusted_profile ]
  }
}
```

## Argument Reference

Review the argument references that you can specify for your resource. 


- `bare_metal_server` - (Required, String) Bare metal server identifier. 
- `default_trusted_profile`- (Optional, List) The default trusted profile to be used when initializing the bare metal server.

  Nested scheme for `default_trusted_profile`:
  - `auto_link` - (Boolean) If set to true, the system will create a link to the specified target trusted profile during server creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the server is deleted.
  - `target` - (List) The default IAM trusted profile to use for this bare metal server.
     Nested scheme for `target`: 
     - `id` - (String) The unique identifier for this trusted profile
     - `crn`- (String) The CRN for this trusted profile
- `image` - (Required, String) Image id to use to reinitialize the bare metal server. 
- `keys` - (Required, Array) Keys ids to use to reinitialize the bare metal server. 
- `user_data` - (Optional, String) User data to transfer to the server bare metal server. If unspecified, no user data will be made available. (For reload/reinitialize provide the same user data as at the time of provisioning)


To reinitialize a bare metal server, the server status must be stopped, or have failed a previous reinitialization. For more information, see Managing Bare Metal Servers for VPC.
