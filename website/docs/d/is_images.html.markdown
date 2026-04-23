---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Images"
description: |-
  Lists IBM Cloud VPC infrastructure images.
---

# ibm_is_images
Retrieve information of an existing IBM Cloud Infrastructure images as a read-only data source. For more information, about IBM Cloud infrastructure images, see [Images](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_images" "ds_images" {
}
```

```terraform
data "ibm_is_images" "ds_images" {
  visibility = "public"
}
```

```terraform
data "ibm_is_images" "remote_images" {
  remote_account_id = "provider"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `catalog_managed` - (Optional, Bool) Lists only those images which are managed as part of a catalog offering.
- `name` - (Optional, String) The name of the image.
- `resource_group` - (Optional, String) The id of the resource group.
- `status` - (Optional, String) Status of the image. Accepted values: **available**, **deleting**, **deprecated**, **failed**, **obsolete**, **pending**, **tentative**, **unusable**
- `user_data_format` - (Optional, Set of String) Filters the collection to images with a user_data_format property matching one of the specified values.

    ~> **Note:** Supported values are:
    - `cloud_init`: user_data will be interpreted according to the cloud-init standard.
    - `esxi_kickstart`: user_data will be interpreted as a VMware ESXi installation script.
    - `ipxe`: user_data will be interpreted as a single URL to an iPXE script or as the text of an iPXE script.

- `visibility` - (Optional, String) Visibility of the image. Accepted values: **private**, **public**
- `remote_account_id` - (Optional, String) Filters the collection to images with a remote account id matching the specified value. Accepted values are `provider`, `user`, or a valid account ID.

## Attribute reference

You can access the following attribute references after your data source is created.

- `images` - (List) List of all images in the IBM Cloud Infrastructure.

  Nested scheme for `images`:
  - `access_tags` - (List) Access management tags associated for image.
  - `allowed_use` - (List) The usage constraints to match against the requested instance or bare metal server properties to determine compatibility.

    Nested schema for `allowed_use`:
    - `api_version` - (String) The API version with which to evaluate the expressions.
    - `bare_metal_server` - (String) The expression that must be satisfied by the properties of a bare metal server provisioned using this image. If unspecified, the expression will be set to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros.

    ~> **NOTE** In addition, the following property is supported, corresponding to `BareMetalServer` properties:
      - `enable_secure_boot` - (Boolean) Indicates whether secure boot is enabled.

    - `instance` - (String) The expression that must be satisfied by the properties of a virtual server instance provisioned using this image. If unspecified, the expression will be set to true. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros.

    ~> **NOTE** In addition, the following variables are supported, corresponding to `Instance` properties:
      - `gpu.count` - (Integer) The number of GPUs.
      - `gpu.manufacturer` - (String) The GPU manufacturer.
      - `gpu.memory` - (Integer) The overall amount of GPU memory in GiB (gibibytes).
      - `gpu.model` - (String) The GPU model.
      - `enable_secure_boot` - (Boolean) Indicates whether secure boot is enabled.

  - `architecture` - (String) The architecture for this image.
  - `catalog_offering` - (List) The catalog offering for this image.

    Nested scheme for `catalog_offering`:
    - `managed` - (Bool) Indicates whether this image is managed as part of a catalog offering. If an image is managed, accounts in the same enterprise with access to that catalog can specify the image's catalog offering version CRN to provision virtual server instances using the image.
    - `version` - (List) The catalog offering version associated with this image. If absent, this image is not associated with a cloud catalog offering.

      Nested scheme for `version`:
      - `crn` - (String) The CRN for this version of a catalog offering.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.

        Nested scheme for `deleted`:
        - `more_info` - (String) Link to documentation about deleted resources.

  - `checksum` - (String) The SHA256 checksum for this image.
  - `crn` - (String) The CRN for this image.
  - `encryption` - (String) The type of encryption used on the image.
  - `encryption_key` - (String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
  - `id` - (String) The unique identifier for this image.
  - `name` - (String) The name for this image.
  - `os` - (String) The name of the Operating System.
  - `operating_system` - (List) The operating system details.

    Nested scheme for `operating_system`:
    - `allow_user_image_creation` - (Bool) Users may create new images with this operating system.
    - `architecture` - (String) The operating system architecture.
    - `dedicated_host_only` - (Bool) Images with this operating system can only be used on dedicated hosts or dedicated host groups.
    - `display_name` - (String) A unique, display-friendly name for the operating system.
    - `family` - (String) The software family for this operating system.
    - `href` - (String) The URL for this operating system.
    - `name` - (String) The globally unique name for this operating system.
    - `user_data_format` - (String) The user data format for this image.

    ~> **Note:** Supported values are:
      - `cloud_init`: user_data will be interpreted according to the cloud-init standard.
      - `esxi_kickstart`: user_data will be interpreted as a VMware ESXi installation script.
      - `ipxe`: user_data will be interpreted as a single URL to an iPXE script or as the text of an iPXE script.

    - `vendor` - (String) The vendor of the operating system.
    - `version` - (String) The major release version of this operating system.

  - `remote` - (List) If present, this property indicates that the resource associated with this reference is remote and therefore may not be directly retrievable.

    Nested schema for `remote`:
    - `account` - (List) Indicates that the referenced resource is remote to this account, and identifies the owning account.

      Nested schema for `account`:
      - `id` - (String) The unique identifier for this account.
      - `resource_type` - (String) The resource type.

  - `resource_group` - (List) The resource group object, for this image.

    Nested scheme for `resource_group`:
    - `href` - (String) The URL for this resource group.
    - `id` - (String) The unique identifier for this resource group.
    - `name` - (String) The user-defined name for this resource group.

  - `source_volume` - (String) The source volume id of the image.
  - `status` - (String) The status of this image.
  - `status_reasons` - (List) The reasons for the current status (if any).

    Nested scheme for `status_reasons`:
    - `code` - (String) The status reason code.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) Link to documentation about this status reason.

  - `user_data_format` - (String) The user data format for this image.
  - `visibility` - (String) The visibility of the image public or private.
  - `zones` - (List) The zones in which this image is available for use. If the image has a status of `available` or `deprecated`, this will include all zones in the region.If the image has a status of `partially_available`, this will include one or more zones in the region. If the image has a status of `failed`, `obsolete`, `pending`, or `unusable`, this will be empty.
      Nested schema for **zones**:
      - `href` - (String) The URL for this zone.
      - `name` - (String) The globally unique name for this zone.