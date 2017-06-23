---
layout: "ibm"
page_title: "IBM : ibm_image_template"
sidebar_current: "docs-ibm-datasource-compute-image-template"
description: |-
  Get information on a IBM Compute Image Template resource
---

# ibm\_compute_image_template

Import the details of an existing image template as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_compute_image_template" "img_tpl" {
    name = "jumpbox"
}
```

The following example shows how you can use this data source to reference the image template ID in the `ibm_compute_vm_instance` resource, since the numeric IDs are often unknown.

```hcl
resource "ibm_compute_vm_instance" "vm1" {
    ...
    image_id = "${data.ibm_compute_image_template.img_tpl.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the image template as it was defined in Bluemix Infrastructure (SoftLayer). The names can be found in the [SoftLayer Customer Portal](https://control.softlayer.com), by navigating to **Devices > Manage > Images**.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the image template.
