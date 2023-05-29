---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Images"
description: |-
  Manages IBM Cloud infrastructure images.
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

data "ibm_is_images" "ds_images" {
  visibility = "public"
}

```
## Argument reference

Review the argument references that you can specify for your data source. 

* `catalog_managed` - (Optional, bool) Lists only those images which are managed as part of a catalog offering.
* `resource_group` - (Optional, string) The id of the resource group.
* `name` - (Optional, string) The name of the image.
* `visibility` - (Optional, string) Visibility of the image.
* `status` - (Optional, string) Status of the image.

## Attribute reference
You can access the following attribute references after your data source is created. 

- `images` - (List) List of all images in the IBM Cloud Infrastructure.

  Nested scheme for `images`:
  - `access_tags`  - (List) Access management tags associated for image.
  - `architecture` - (String) The architecture for this image.
  - `crn` - (String) The CRN for this image.
  - `catalog_offering` - (List) The catalog offering for this image.

      Nested scheme for **catalog_offering**:
      - `managed` - (Bool) Indicates whether this image is managed as part of a catalog offering. If an image is managed, accounts in the same enterprise with access to that catalog can specify the image's catalog offering version CRN to provision virtual server instances using the image.
      - `version` - (List) The catalog offering version associated with this image. If absent, this image is not associated with a cloud catalog offering.
      
          Nested scheme for **version**:
            - `crn` - (String) The CRN for this version of a catalog offering
  - `checksum` - (String) TThe SHA256 checksum for this image.
  - `encryption` - (String) The type of encryption used on the image.
  - `encryption_key` - (String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
  - `id` - (String) The unique identifier for this image.
  - `name` - (String) The name for this image.
  - `os` - (String) The name of the Operating System.
  - `status` - (String) The status of this image.
  - `visibility` - (String) The visibility of the image public or private.
  - `source_volume` - The source volume id of the image.

