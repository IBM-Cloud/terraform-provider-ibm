---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_images"
description: |-
  List all of the images for the respective cloud instance that are imported  from catalog by the user
---

# ibm\_pi_images

Import the details of existing IBM Power Virtual Server Cloud images as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_images" "ds_images" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```
## Notes:
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  Example Usage:
  ```hcl
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
## Argument Reference
The following arguments are supported:

* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `image_info` - List of all images in the IBM Power Virtual Server Cloud.
  * `name` - The name for this image.
  * `id` - The unique identifier for this image
  * `state` - The state of the operating system.
  * `href` - The href  of this image.
  * `storage_type` - The storage type for this image
