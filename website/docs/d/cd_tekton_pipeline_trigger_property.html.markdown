---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_trigger_property"
description: |-
  Get information about cd_tekton_pipeline_trigger_property
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_trigger_property

Provides a read-only data source to retrieve information about a cd_tekton_pipeline_trigger_property. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property" {
	pipeline_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.pipeline_id
	property_name = "debug-pipeline"
	trigger_id = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.trigger_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `property_name` - (Required, Forces new resource, String) The property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `trigger_id` - (Required, Forces new resource, String) The trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_trigger_property.
* `enum` - (List) Options for `single_select` property type. Only needed for `single_select` property type.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
* `href` - (String) API URL for interacting with the trigger property.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `locked` - (Boolean) When true, this property cannot be overridden at runtime. Attempting to override it will result in run requests being rejected. The default is false.
* `name` - (Forces new resource, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `path` - (String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
* `type` - (Forces new resource, String) Property type.
  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
* `value` - (String) Property value. Any string value is valid.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

