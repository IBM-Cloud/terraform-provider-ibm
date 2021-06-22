---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : ibm_image_template"
description: |-
  Get information on a IBM Cloud compute image template resource.
---

# ibm_compute_image_template
Retrieve information of an existing image template as a read-only data source. For more information, about IBM Cloud compute image template, see [about bare metal custom image templates](https://cloud.ibm.com/docs/cloud-infrastructure?topic=bare-metal-getting-started-bm-custom-image-templates).

## Example usage
The following example shows how you can retrieve the ID of an image template and reference this ID in your `ibm_compute_vm_instance` resource. 

```terraform
data "ibm_compute_image_template" "img_tpl" {
    name = "jumpbox"
}

resource "ibm_compute_vm_instance" "vm1" {
    image_id = data.ibm_compute_image_template.img_tpl.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the image template. You can find the name in the [IBM Cloud infrastructure customer portal](https://cloud.ibm.com/classic) by navigating to **Devices > Manage > Images**.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the image template.
