---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_images"
description: |-
  List all of the images for the respective cloud instance that are imported from catalog by the user.
---

# ibm_pi_images

Retrieve a list of supported images that you can use in your Power Systems Virtual Server instance. The image represents the version of the operation system that is installed in your Power Systems Virtual Server instance. For more information, about power instance images, see [capturing and exporting a virtual machine (VM)](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-capturing-exporting-vm).

## Example Usage

The following example retrieves all images for a cloud instance ID.

```terraform
data "ibm_pi_images" "ds_images" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`
  
Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `image_info` - (List) List of all supported images.

  Nested scheme for `image_info`:
  - `crn` - (String) The CRN of this resource.
  - `href` - (String) The hyper link of an image.
  - `id` - (String) The unique identifier of an image.
  - `image_type` - (String) The identifier of this image type.
  - `name`-  (String) The name of an image.
  - `source_checksum` - (String) Checksum of the image.
  - `state` - (String) The state of an image.
  - `storage_pool` - (String) Storage pool where image resides.
  - `storage_type` - (String) The storage type of an image.
  - `user_tags` - (List) List of user tags attached to the resource.
