---
layout: "ibm"
page_title: "IBM : ibm_cm_preset"
description: |-
  Get information about ibm_cm_preset
subcategory: "Catalog Management"
---

# ibm_cm_preset

Provides a read-only data source for ibm_cm_preset. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_preset" "cm_preset" {
	id = "${ibm_cm_object.my_object.id}@1.0.0"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The ID of the preset.  Format is <catalog_id>-<object_name>@<preset_version>

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `preset` - (String) The map of preset values as a JSON string.
