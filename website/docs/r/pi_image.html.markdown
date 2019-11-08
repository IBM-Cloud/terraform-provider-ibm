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

* `id` - The unique identifier of the image.
