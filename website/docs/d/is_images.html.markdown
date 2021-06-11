---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Images"
description: |-
  Manages IBM Cloud infrastructure images.
---

# ibm_is_images

Retrieve information of an existing IBM Cloud Infrastructure images as a read-only data source. For more information, about IBM Cloud infrastructure images, see [Images](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images).

## Example usage

```terraform
data "ibm_is_images" "ds_images" {
}
```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `images` - (List) List of all images in the IBM Cloud Infrastructure.
  - `architecture` - (String) The architecture for this image.
  - `crn` - (String) The CRN for this image.
  - `checksum` - (String) TThe SHA256 checksum for this image.
  - `encryption` - (String) The type of encryption used on the image.
  - `encryption_key` - (String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
  - `id` - (String) The unique identifier for this image.
  - `name` - (String) The name for this image.
  - `os` - (String) The name of the Operating System.
  - `status` - (String) The status of this image.
  - `visibility` - (String) The visibility of the image public or private.


