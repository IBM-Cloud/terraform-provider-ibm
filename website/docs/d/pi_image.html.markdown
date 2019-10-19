---
layout: "ibm"
page_title: "IBM : Image"
sidebar_current: "docs-ibm-datasources-pi-image"
description: |-
  Manages IBM Cloud Infrastructure Images for IBM Power
---

# ibm\_pi_image

Import the details of an existing IBM Cloud Infrastructure image for IBM Power as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_pi_image" "ds_image" {
    name = "7200-03-03"
    powerinstanceid="49fba6c9-23f8-40bc-9899-aca322ee7d5b"

}

```

## Argument Reference

The following arguments are supported:

* `imagename` - (Required, string) The name of the image.
* `powerinstanceid` - (Required, string) The service instance associated with the account


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this image.
* `state` - The state for this image.
* `size` - The size of the image.
* `architecture` - The architecture for this image.
* `operatingsystem` - The os for this image



