---
layout: "ibm"
page_title: "IBM: pi_image"
sidebar_current: "docs-ibm-datasources-pi-image"
description: |-
  Manages an image in the Power Virtual Server Cloud.
---

# ibm\_pi_image

Import the details of an existing IBM Power Virtual Server Cloud image as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_image" "ds_image" {
  pi_image_name        = "7200-03-03"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_image_name` - (Required, string) The name of the image.
* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this image.
* `state` - The state for this image.
* `size` - The size of the image.
* `architecture` - The architecture for this image.
* `operatingsystem` - The operating system for this image.
