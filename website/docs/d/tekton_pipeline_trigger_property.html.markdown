---
layout: "ibm"
page_title: "IBM : ibm_tekton_pipeline_trigger_property"
description: |-
  Get information about tekton_pipeline_trigger_property
subcategory: "Continuous Delivery Pipeline"
---

# ibm_tekton_pipeline_trigger_property

Provides a read-only data source for tekton_pipeline_trigger_property. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
	pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
	property_name = "debug-pipeline"
	trigger_id = "1bb892a1-2e04-4768-a369-b1159eace147"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `property_name` - (Required, Forces new resource, String) The property's name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
* `trigger_id` - (Required, Forces new resource, String) The trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the tekton_pipeline_trigger_property.
* `name` - (Required, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.

* `options` - (Optional, Map) Options for SINGLE_SELECT property type.

* `path` - (Optional, String) property path for INTEGRATION type properties.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.

* `type` - (Required, String) Property type.
  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`.

* `value` - (Optional, String) String format property value.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.

