---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_image_deprecate"
description: |-
  Manages IBM VPC custom images.
---

# ibm_is_image_deprecate

Provide support to deprecate a VPC image. This allows images to be marked as deprecared, resulting in its `status` becoming `deprecated` and `deprecation_at` being set to the current date and time. For more information, about VPC custom images, see [IBM Cloud Docs: Virtual Private Cloud - IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Timeouts

The `ibm_is_image_deprecate` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

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
resource "ibm_is_image_deprecate" "example" {
  image               = ibm_is_image.example.id
}
```
  ~> **NOTE**
      `operating_system` is required with `href`.

## Argument reference
Review the argument references that you can specify for your resource. 

- `image` - (Required, Forces new resource, String) The id of an image to be deprecated.

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

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_image_deprecate` resource by using `id`.
The `id` property can be formed from `image ID`. For example:

```terraform
import {
  to = ibm_is_image_deprecate.example
  id = "<image_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_image_deprecate.example <image_id>
```