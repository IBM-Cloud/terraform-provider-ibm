---
layout: "ibm"
page_title: "IBM: pi_image"
sidebar_current: "docs-ibm-resource-pi-image"
description: |-
  Manages IBM Image in the Power Virtual Server Cloud.
---

# ibm\_pi_image

Provides a image resource. This allows image to be created, updated, and cancelled in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create a image:

```hcl
resource "ibm_pi_image" "testacc_image  "{
  pi_image_name       = "7200-03-02"
  pi_image_id         = [ "image id obtained from the datasource"]
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
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
## Timeouts

ibm_pi_image provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for Creating image.
* `delete` - (Default 60 minutes) Used for Deleting image.

## Argument Reference

The following arguments are supported:

* `pi_image_name` - (Required, string) The name for this image.
* `pi_image_id` - (Required, string) The image id for this image.
* `pi_cloud_instance_id` - (Required, string) The cloud_instance_id for this account.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the image.The id is composed of \<power_instance_id\>/\<image_id\>.
* `image_id` - The unique identifier of the image.

## Import

ibm_pi_image can be imported using `power_instance_id` and `image_id`, eg

```
$ terraform import ibm_pi_image.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
