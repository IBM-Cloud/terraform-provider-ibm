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

resource "ibm_resource_group" "resourceGroup" {
  name     = "prod"
}

resource "ibm_is_image" "test_is_images" {
 name                   = "test_image"
 href                   = "cos://us-south/buckettesttest/livecd.ubuntu-cpc.azure.vhd"
 operating_system       = "my-image-ubuntu-16-04-amd64"
 visibility = "private"
 resource_group = ibm_resource_group.resourceGroup.id
 encrypted_data_key     = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
 encryption_key         = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"
}

data "ibm_is_images" "ds_images" {
  name = "my-image-ubuntu-16-04-amd64"
}

data "ibm_is_images" "ds_images" {
  resource_group = ibm_resource_group.resourceGroup.id
}

data "ibm_is_images" "ds_images" {
  visibility = "private"
}

```
## Argument Reference

The following arguments are supported:

* `resource_group` - (Optional, string) The id of the resource group.
* `name` - (Optional, string) The name of the image.
* `visibility` - (Optional, string) Visibility of the image.
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


