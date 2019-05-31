---
layout: "ibm"
page_title: "IBM : Image"
sidebar_current: "docs-ibm-datasources-is-image"
description: |-
  Manages IBM Cloud Infrastructure Images.
---

# ibm\_is_images

Import the details of an existing IBM Cloud Infrastructure image as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_image" "ds_image" {
    name = "centos-7.x-amd64"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the image.
* `visibility` - (Optional, string) The visibility of the image. Accepted values are `public` or `private`.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this image.
* `crn` - The CRN for this image.
* `os` - The name of the operating system.
* `status` - The status of this image.
* `architecture` - The architecture for this image.



