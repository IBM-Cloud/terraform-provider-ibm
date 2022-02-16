---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_disk"
description: |-
  Manages IBM Cloud Bare Metal Server Disk.
---

# ibm\_is_bare_metal_server_disk

Import the details of configuration variables used to initialize the bare metal server, such as the image used, SSH keys, and any configured usernames and passwords as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal servers, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

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

data "ibm_is_bare_metal_server_initialization" "ds_bms_ini" {
  bare_metal_server         = ibm_is_bare_metal_server.example.id
}

```

## Argument reference
Review the argument references that you can specify for your data source.

- `bare_metal_server` - (Required, String) The id for this bare metal server.
- `passphrase` - (Optional, String) The passphrase that you used when you created your SSH key. If you did not enter a passphrase when you created the SSH key, do not provide this input parameter.
- `private_key` - (Optional, String) The private key of an SSH key that you want to add to your Bare metal server during creation in PEM format. It is used to decrypt the default password of the Windows administrator for the bare metal server if the image is used of type `windows`.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The unique identifier for this bare metal server.
- `image` - (String) The unique identifier for this image
- `image_name` - (String) The user-defined or system-provided name for this image.
- `keys` - (Array) List of public SSH keys used at initialization.
- `user_accounts` - (List) The size of the disk in GB (gigabytes).
  Nested scheme for `user_accounts`:
    - `encrypted_password` - (String) The password at initialization, encrypted using encryption_key, and returned base64-encoded.
    - `encryption_key` - (String) The CRN for this key.
    - `password` - (String) The password that you can use to access your bare metal server.
    - `resource_type` - (String) The type of resource referenced.
    - `username` - (String) The username for the account created at initialization.
