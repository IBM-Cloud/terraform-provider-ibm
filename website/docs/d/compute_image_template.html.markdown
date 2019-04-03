---
layout: "ibm"
page_title: "IBM : ibm_image_template"
sidebar_current: "docs-ibm-datasource-compute-image-template"
description: |-
  Get information on a IBM Compute Image Template resource
---

# ibm\_compute_image_template

Import the details of an existing image template as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_compute_image_template" "img_tpl" {
    name = "jumpbox"
}
```

The following example shows how you can use this data source to reference the image template ID in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```hcl
resource "ibm_compute_vm_instance" "vm1" {
    ...
    image_id = "${data.ibm_compute_image_template.img_tpl.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the image template as it was defined in IBM Cloud Infrastructure (SoftLayer). You can find the name in the [IBM Cloud infrastructure customer portal](https://control.softlayer.com) by navigating to **Devices > Manage > Images**.
* `most_recent` - (Optional, boolean) Ask the provider for the latest version of the image template by selecting the one with the highest identifier with the requested name.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the image template.
