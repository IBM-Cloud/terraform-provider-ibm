---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_image"
description: |-
  Manages IBM VPC Custom Images.
---

# ibm\_is_image

Provide an image resource. This allows images to be created, retrieved, and deleted.

For additional details, see the [IBM Cloud Docs: Virtual Private Cloud - IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

## Example Usage

```
resource "ibm_is_image" "test_is_images" {
 name                   = "test_image"
 href                   = "cos://us-south/buckettesttest/livecd.ubuntu-cpc.azure.vhd"
 operating_system       = "ubuntu-16-04-amd64"
 encrypted_data_key     = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
 encryption_key         = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"
   
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify an image.
* `href` - (Required, string) The path(SQL URL of COS Bucket Object) of an image to be uploaded.
* `operating_system` - (Required, string) Description of underlying OS of an image.
* `resource_group` - (Optional, Forces new resource, string) The resource group ID for this image.
* `encrypted_data_key` - (Optional, Forces new resource, string) A base64-encoded, encrypted representation of the key that was used to encrypt the data for this image.
* `encryption_key` - (Optional, Forces new resource, string) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
* `tags` - (Optional, array of strings) Tags associated with the image.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the image.
* `architecture` - The architecture which image is based on
* `crn` - The CRN for an image
* `checksum` - The SHA256 checksum of this image
* `file` - The file
* `format` - The format of an image
* `resourceGroup` - The resource group which image is belonging to
* `status` - The status of an image such as corrupt, available
* `visibility` - The access scope of an image such as private or public
* `encryption` - The type of encryption used on the image

``

## Import

ibm_is_image can be imported using Image ID, eg

```
$ terraform import ibm_is_image.example d7bec597-4726-451f-8a63-e62e6f19c32c
```