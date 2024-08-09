---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_property"
description: |-
  Manages cd_tekton_pipeline_property.
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_property

Provides a resource for cd_tekton_pipeline_property. This allows cd_tekton_pipeline_property to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  name = "prop1"
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  type = "text"
  value = "https://github.com/open-toolchain/hello-tekton.git"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `enum` - (Optional, List) Options for `single_select` property type. Only needed when using `single_select` property type.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
* `locked` - (Optional, Boolean) When true, this property cannot be overridden by a trigger property or at runtime. Attempting to override it will result in run requests being rejected. The default is false.
  * Constraints: The default value is `false`.
* `name` - (Required, Forces new resource, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `path` - (Optional, String) A dot notation path for `integration` type properties only, to select a value from the tool integration. If left blank the full tool integration data will be used.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `type` - (Required, Forces new resource, String) Property type.
  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
* `value` - (Optional, String) Property value. Any string value is valid.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_tekton_pipeline_property.
* `href` - (String) API URL for interacting with the property.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.


## Import

You can import the `ibm_cd_tekton_pipeline_property` resource by using `name`.
The `name` property can be formed from `pipeline_id`, and `property_name` in the following format:

<pre>
&lt;pipeline_id&gt;/&lt;property_name&gt;
</pre>
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The Tekton pipeline ID.
* `property_name`: A string in the format `debug-pipeline`. The property name.

# Syntax
```
$ terraform import ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property <pipeline_id>/<property_name>
```
