---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_image"
description: |-
  Manages IBM VPC custom images.
---

# ibm_is_image

Provide an image resource. This allows images to be created, retrieved, and deleted. For more information, about VPC custom images, see [IBM Cloud Docs: Virtual Private Cloud - IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

## Example usage

```terraform
resource "ibm_is_image" "test_is_images" {
 name                   = "test_image"
 href                   = "cos://us-south/buckettesttest/livecd.ubuntu-cpc.azure.vhd"
 operating_system       = "ubuntu-16-04-amd64"
 encrypted_data_key     = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
 encryption_key         = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"
   
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `encrypted_data_key` - (Optional, Forces new resource, String) A base64-encoded, encrypted representation of the key that was used to encrypt the data for this image.
- `encryption_key` - (Optional, Forces new resource, String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
- `href` - (Required, String) The path of an image to be uploaded.
- `name` - (Required, String) The descriptive name used to identify an image.
- `operating_system` - (Required, String) Description of underlying OS of an image.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID for this image.
- `tags` (Optional, Array of Strings) A list of tags that you want to your image. Tags can help you find the image more easily later.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `architecture` - (String) The processor architecture that this image is based on.
- `crn` - (String) The CRN of the image.
- `checksum`-  (String) The `SHA256` checksum of the image.
- `encryption` - (String) The type of encryption used on the image.
- `file` - (String) The file.
- `format` - (String) The format of an image.
- `id` - (String) The unique identifier of the image.
- `resourceGroup` - (String) The resource group to which the image belongs to.
- `status`-String The status of an image such as `corrupt`, or `available`.
- `visibility` - (String) The access scope of an image such as `private` or `public`.


## Import
The `ibm_is_image` resource can be imported by using image ID.

**Example**

```
$ terraform import ibm_is_image.example d7bec597-4726-451f-8a63-e62e6f121c32c
```
