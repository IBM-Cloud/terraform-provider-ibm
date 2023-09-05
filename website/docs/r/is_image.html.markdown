---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_image"
description: |-
  Manages IBM VPC custom images.
---

# ibm_is_image

Provide an image resource. This allows images to be created, retrieved, and deleted. For more information, about VPC custom images, see [IBM Cloud Docs: Virtual Private Cloud - IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Timeouts

The `ibm_is_image` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

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
```
  ~> **NOTE**
      `operating_system` is required with `href`.

## Example usage (using volume)      
```terraform
resource "ibm_is_image" "example" {
  name = "example-image"

  //optional field, either of (href, operating_system) or source_volume is required

  source_volume      = "xxxx-xxxx-xxxxxxx"
  encrypted_data_key = "eJxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx0="
  encryption_key     = "crn:v1:bluemix:public:kms:us-south:a/6xxxxxxxxxxxxxxx:xxxxxxx-xxxx-xxxx-xxxxxxx:key:dxxxxxx-fxxx-4xxx-9xxx-7xxxxxxxx"

  //increase timeouts as per volume size
  timeouts {
    create = "45m"
  }

}
```
## Example usage (lifecycle)      
```terraform
resource "ibm_is_image" "example" {
  name               = "example-image"
  href               = "cos://us-south/buckettesttest/livecd.ubuntu-cpc.azure.vhd"
  operating_system   = "ubuntu-16-04-amd64"
  deprecation_at     = "2023-09-28T15:10:00.000Z"
  obsolescence_at    = "2023-11-28T15:10:00.000Z"
}
```
  ~> **NOTE**
      `obsolescence_at` must be later than `deprecation_at` (if `deprecation_at` is set).



## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the image

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `deprecate` - (Bool) This flag deprecates an image, resulting in its status becoming deprecated and deprecation_at being set to the current date and time. The image must:

    - be an existing image and have a status of available
    - have catalog_offering.managed set to false
    - not have deprecation_at set

A system-provided image is not allowed to be deprecated.
- `deprecation_at` - (String) The deprecation date and time (UTC) for this image. If absent, no deprecation date and time has been set.
  
  ~> **NOTE**
      Specify "null" to remove an existing deprecation date and time. If the image status is currently deprecated, it will become available.
  	string($date-time)

    - This cannot be set if the image has a status of `failed` or `deleting`, or if `catalog_offering`.`managed` is true.
    - The date and time must not be in the past, and must be earlier than `obsolescence_at` (if `obsolescence_at` is set). Additionally, if the image status is currently deprecated, the value cannot  be changed (but may be removed).
    - If the deprecation date and time is reached while the image has a status of pending, the image's     status will transition to deprecated upon its successful creation (or obsolete if the obsolescence     date and time was also reached).

- `encrypted_data_key` - (Optional, Forces new resource, String) A base64-encoded, encrypted representation of the key that was used to encrypt the data for this image.
- `encryption_key` - (Optional, Forces new resource, String) The CRN of the Key Protect Root Key or Hyper Protect Crypto Service Root Key for this resource.
- `href` - (Optional, String) The path of an image to be uploaded. The Cloud Object Store (COS) location of the image file.

  ~> **NOTE**
      either `href` or `source_volume` is required
- `name` - (Required, String) The descriptive name used to identify an image.
- `obsolete` - (Optional, Bool) This flag obsoletes an image, resulting in its status becoming obsolete and obsolescence_at being set to the current date and time. The image must:

    - be an existing image and have a status of available or deprecated
    - have catalog_offering.managed set to false
    - not have deprecation_at set in the future
    - not have obsolescence_at set
    - A system-provided image is not allowed to be obsolescence.

- `obsolescence_at` - (Optional, String) The obsolescence date and time (UTC) for this image. If absent, no obsolescence date and time has been set.
  
  ~> **NOTE**
      Specify "null" to remove an existing obsolescence date and time. If the image status is currently obsolete, it will become deprecated if deprecation_at is in the past. Otherwise, it will become available.

    - This cannot be set if the image has a status of failed or deleting, or if `catalog_offering`.`managed` is true.
    - The date and time must not be in the past, and must be later than `deprecation_at` (if `deprecation_at` is set). Additionally, if the image status is currently obsolete, the value cannot  be changed (but may be removed).
    - If the obsolescence date and time is reached while the image has a status of pending, the image's status will transition to obsolete upon its successful creation.
- `operating_system` - (Required, String) Description of underlying OS of an image.

  ~> **NOTE**
      `operating_system` is required with `href`
- `resource_group` - (Optional, Forces new resource, String) The resource group ID for this image.
- `source_volume` - (Optional, string) The volume id of the volume from which to create the image.

  ~> **NOTE**
      either `source_volume` or `href` is required.

  The specified volume must:
    - Originate from an image, which will be used to populate this image's operating system information.(boot type volumes)
    - Not be active or busy.
    - During image creation, the specified volume may briefly become busy.
    - Creating image from volume requires instance to which volume is attached to be in stopped status, running instance will be stopped on using this option.
    - increase the default timeout as per the volume size.
- `tags` (Optional, Array of Strings) A list of tags that you want to your image. Tags can help you find the image more easily later.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `architecture` - (String) The processor architecture that this image is based on.
- `created_at` - (String) The date and time that the image was created
- `crn` - (String) The CRN of the image.
- `checksum`-  (String) The `SHA256` checksum of the image.
- `encryption` - (String) The type of encryption used on the image.
- `file` - (String) The file.
- `format` - (String) The format of an image.
- `id` - (String) The unique identifier of the image.
- `resourceGroup` - (String) The resource group to which the image belongs to.
- `status`- (String) The status of an image such as `corrupt`, or `available`.
- `visibility` - (String) The access scope of an image such as `private` or `public`.


## Import
The `ibm_is_image` resource can be imported by using image ID.

**Example**

```
$ terraform import ibm_is_image.example d7bec597-4726-451f-8a63-e62e6f121c32c
```
