---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Images"
description: |-
  Manages IBM Cloud Infrastructure Images.
---

# ibm\_is_images

Import the details of an existing IBM Cloud Infrastructure images as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_images" "ds_images" {
}

```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `images` - List of all images in the IBM Cloud Infrastructure.
  * `name` - The name for this image.
  * `id` - The unique identifier for this image.
  * `crn` - The CRN for this image.
  * `checksum` - The SHA256 Checksum for this image
  * `os` - The name of the operating system.
  * `status` - The status of this image.
  * `architecture` - The architecture for this image.
  * `visibility` - The visibility of the image public or private.
  * `encryption` - The type of encryption used on the image.
  * `encryption_key` - The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
  * `source_volume` - The source volume id of the image.


