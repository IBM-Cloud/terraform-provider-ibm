---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_image_obsolete"
description: |-
  Manages IBM VPC custom images.
---

# ibm_is_image_obsolete

Provide support to obsolete a VPC image. This resource obsoletes an image, resulting in its status becoming `obsolete` and `obsolescence_at` being set to the current date and time. For more information, about VPC custom images, see [IBM Cloud Docs: Virtual Private Cloud - IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Timeouts

The `ibm_is_image_obsolete` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- **create** - (Default 10 minutes) Used for creating image.
- **update** - (Default 10 minutes) Used for updating image.
- **delete** - (Default 10 minutes) Used for deleting image.



## Example usage (using href and operating_system)

```terraform
resource "ibm_is_image" "example" {
  name               = "example-image"
  href               = "cos://us-south/buckettesttest/livecd.ubuntu-cpc.azure.vhd"
  operating_system   = "ubuntu-16-04-amd64"
  encrypted_data_key = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
  encryption_key     = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"
}
resource "ibm_is_image_obsolete" "example" {
  image               = ibm_is_image.example.id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `image` - (Required, Forces new resource, String) The id of an image to be made obsolete.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `architecture` - (String) The processor architecture that this image is based on.
- `created_at` - (String) The date and time that the image was created
- `crn` - (String) The CRN of the image.
- `checksum`-  (String) The `SHA256` checksum of the image.
- `deprecation_at` - (String) The deprecation date and time (UTC) for this image. If absent, no deprecation date and time has been set.
- `encryption` - (String) The type of encryption used on the image.
- `file` - (String) The file.
- `format` - (String) The format of an image.
- `id` - (String) The unique identifier of the image.
- `obsolescence_at` - (String) The obsolescence date and time (UTC) for this image. If absent, no obsolescence date and time has been set.
- `resourceGroup` - (String) The resource group to which the image belongs to.
- `status`- (String) The status of an image such as `corrupt`, or `available`.
- `visibility` - (String) The access scope of an image such as `private` or `public`.


## Import
The `ibm_is_image_obsolete` resource can be imported by using image ID.

**Example**

```
$ terraform import ibm_is_image_obsolete.example d7bec597-4726-451f-8a63-e62e6f121c32c
```
