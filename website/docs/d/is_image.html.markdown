---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Image"
description: |-
  Manages IBM Cloud infrastructure image.
---

# ibm_is_image
Retrieve information of an existing IBM Cloud Infrastructure image as a read-only data source. For more information, about VPC custom images, see [IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

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

data "ibm_is_image" "example" {
  name = "centos-7.x-amd64"
}
```
```terraform
data "ibm_is_image" "example" {
  identifier = ibm_is_image.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `identifier` - (Optional, String) The id of the image.

    ~> **Note:** `name` and `identifier` are mutually exclusive.

- `name` - (Optional, String) The name of the image.

    ~> **Note:** `name` and `identifier` are mutually exclusive.

- `visibility` - (Optional, String) The visibility of the image. Accepted values are `public` or `private`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `access_tags`  - (List) Access management tags associated for image.
- `architecture` - (String) The architecture of the image.
- `catalog_offering` - (List) The catalog offering for this image.

  Nested scheme for **catalog_offering**:
  - `managed` - (Bool) Indicates whether this image is managed as part of a catalog offering. If an image is managed, accounts in the same enterprise with access to that catalog can specify the image's catalog offering version CRN to provision virtual server instances using the image.
  - `version` - (List) The catalog offering version associated with this image. If absent, this image is not associated with a cloud catalog offering.
  
      Nested scheme for **version**:
        - `crn` - (String) The CRN for this version of a catalog offering
- `created_at` - (String) The date and time that the image was created
- `checksum`-  (String) The `SHA256` checksum of the image.
- `crn` - (String) The CRN for this image.
- `deprecation_at` - (String) The deprecation date and time (UTC) for this image. If absent, no deprecation date and time has been set.
- `encryption` - (String) The type of encryption used of the image.
- `encryption_key`-  (String) The CRN of the Key Protect or Hyper Protect Crypto Service root key for this resource.
- `id` - (String) The unique identifier of the image.
- `obsolescence_at` - (String) The obsolescence date and time (UTC) for this image. If absent, no obsolescence date and time has been set.
- `os` - (String) The name of the operating system.
- `operating_system` - (List) The operating system details. 
    
  Nested scheme for `operating_system`:
  - `architecture` - (String) The operating system architecture.
  - `dedicated_host_only` - (Bool) Images with this operating system can only be used on dedicated hosts or dedicated host groups.
  - `display_name` - (String) A unique, display-friendly name for the operating system.
  - `family` - (String) The software family for this operating system.
  - `href` - (String) The URL for this operating system.
  - `name` - (String) The globally unique name for this operating system.
  - `vendor` - (String) The vendor of the operating system.
  - `version` - (String) The major release version of this operating system.
- `resource_group` - (List) The resource group object, for this image.
  
  Nested scheme for `resource_group`:
  - `href` - (String) The URL for this resource group.
  - `id` - (String) The unique identifier for this resource group.
  - `name` - (String) The user-defined name for this resource group.
- `status` - (String) The status of this image.
- `status_reasons` - (List) The reasons for the current status (if any).

    Nested scheme for `status_reasons`:
  - `code` - (String) The status reason code
  - `message` - (String) An explanation of the status reason
  - `more_info` - (String) Link to documentation about this status reason

- `source_volume` - The source volume id of the image.
