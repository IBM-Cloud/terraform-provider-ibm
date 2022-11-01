---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_property"
description: |-
  Get information about cd_tekton_pipeline_property
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_property

Provides a read-only data source for cd_tekton_pipeline_property. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property" {
	pipeline_id = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property.pipeline_id
	property_name = "debug-pipeline"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `property_name` - (Required, Forces new resource, String) The property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_tekton_pipeline_property.
* `enum` - (List) Options for `single_select` property type. Only needed when using `single_select` property type.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.

* `name` - (Forces new resource, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.

* `path` - (String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.

* `type` - (String) Property type.
  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.

* `value` - (String) Property value. Any string value is valid.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

