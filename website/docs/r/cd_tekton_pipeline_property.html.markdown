---
layout: "ibm"
page_title: "IBM : ibm_cd_tekton_pipeline_property"
description: |-
  Manages cd_tekton_pipeline_property.
subcategory: "Continuous Delivery"
---

# ibm_cd_tekton_pipeline_property

Create, update, and delete cd_tekton_pipeline_propertys with this resource.

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

You can specify the following arguments for this resource.

* `enum` - (Optional, List) Options for `single_select` property type. Only needed when using `single_select` property type.
  * Constraints: The list items must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`. The maximum length is `256` items. The minimum length is `0` items.
* `locked` - (Optional, Boolean) When true, this property cannot be overridden by a trigger property or at runtime. Attempting to override it will result in run requests being rejected. The default is false.
* `name` - (Required, Forces new resource, String) Property name.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,253}$/`.
* `path` - (Optional, String) A dot notation path for `integration` type properties only, that selects a value from the tool integration. If left blank the full tool integration data will be used.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^[-0-9a-zA-Z_.]*$/`.
* `pipeline_id` - (Required, Forces new resource, String) The Tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `type` - (Required, Forces new resource, String) Property type.
  * Constraints: Allowable values are: `secure`, `text`, `integration`, `single_select`, `appconfig`.
* `value` - (Optional, String) Property value. Any string value is valid.
  * Constraints: The maximum length is `4096` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_tekton_pipeline_property.
* `href` - (String) API URL for interacting with the property.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.


## Import

You can import the `ibm_cd_tekton_pipeline_property` resource by using `name`.
The `name` property can be formed from `pipeline_id`, and `name` in the following format:

<pre>
&lt;pipeline_id&gt;/&lt;name&gt;
</pre>
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The Tekton pipeline ID.
* `name`: A string in the format `prop1`. Property name.

# Syntax
<pre>
$ terraform import ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property &lt;pipeline_id&gt;/&lt;name&gt;
</pre>
