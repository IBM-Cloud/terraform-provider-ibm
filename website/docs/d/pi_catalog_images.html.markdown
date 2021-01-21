---
layout: "ibm"
page_title: "IBM: pi_catalog_images"
sidebar_current: "docs-ibm-datasources-pi-catalog-images"
description: |-
  List all images available for copying into cloud instances
---

# ibm\_pi_catalog_images

List all images available for copying into cloud instances

## Example Usage

```hcl
data "ibm_pi_catalog_images" "ds_images" {
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

* `images` - List of all images in the IBM Power Virtual Server Cloud.
  * `image_id` - The unique identifier for this image
  * `name` - The name for this image.
  * `state` - The state of the operating system.
  * `description` - description of the image.
  * `storage_type` - The storage type for this image
  * `href` - The href  of this image.
  * `creation_date` - Creation Date of the Image.
  * `last_update_date` - The Last Updated Date of the image.
  * `image_type` - The Type of the Format.
  * `disk_format` - Disk format.
  * `operating_system` - Operating System
  * `hypervisor_type` - Hypervisor Type
  * `architecture` - Architechture
  * `endianness` - Endianness
  * `href` - Address of the Image.
