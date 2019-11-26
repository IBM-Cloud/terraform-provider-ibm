---
layout: "ibm"
page_title: "IBM: compute_images"
sidebar_current: "docs-ibm-resource-compute-images"
description: |-
  Manages IBM Compute Images.
---

# ibm\_compute_images

Provide an image resource. This allows images to be created, retrieved, and deleted.

For additional details, see the [IBM Cloud Docs: Virtual Private Cloud - IBM Cloud Importing and managing custom images](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-images).

## Example Usage

```
resource "ibm_is_image" "test_is_images" {
 name                   = "test_image"
 href                   = "test_image_path"
 operating_system       = "test_os_info"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify an image.
* `href` - (Required, string) The path of an image to be uploaded.
* `operating_system` - (Required, string) Description of underlying OS of an image.
* `tags` - (Optional, array of strings) Tags associated with the image.  

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the image.
* `name` - The name of an image.
* `architecture` - The architecture which image is based on
* `createdAt` - The date and time in which the image was created
* `crn` - The CRN for an image
* `file` - The file
* `format` - The format of an image
* `href` - The URL of an image
* `os` - The operating system which image is based on
* `resourceGroup` - The resource group which image is belonging to
* `status` - The status of an image such as corrupt, available
* `visibility` - The access scope of an image such as private or public

``