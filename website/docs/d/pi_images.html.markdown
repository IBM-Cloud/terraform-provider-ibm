---
layout: "ibm"
page_title: "IBM : Images"
sidebar_current: "docs-ibm-datasources-pi-images"
description: |-
  Manages IBM Cloud Infrastructure Images for IBM Power
---

# ibm\_pi_images

Import the details of an existing IBM Cloud Infrastructure images as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_pi_images" "ds_images" {
   powerinstanceid="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}

```

The following arguments are supported:

* `powerinstanceid` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `images` - List of all images in the IBM Cloud Infrastructure for IBM Power
  * `name` - The name for this image.
  * `id` - The unique identifier for this image
  * `state` - The state of the operating system.
  * `href` - The href  of this image.
  


