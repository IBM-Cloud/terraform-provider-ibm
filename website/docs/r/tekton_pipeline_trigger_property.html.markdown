---
layout: "ibm"
page_title: "IBM : ibm_tekton_pipeline_trigger_property"
description: |-
  Manages tekton_pipeline_trigger_property.
subcategory: "Continuous Delivery Pipeline"
---

# ibm_tekton_pipeline_trigger_property

Provides a resource for tekton_pipeline_trigger_property. This allows tekton_pipeline_trigger_property to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property" {
  create_tekton_pipeline_trigger_properties_request = {"name":"key1","value":"https://github.com/IBM/tekton-tutorial.git","type":"TEXT"}
  pipeline_id = "94619026-912b-4d92-8f51-6c74f0692d90"
  trigger_id = "1bb892a1-2e04-4768-a369-b1159eace147"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `create_tekton_pipeline_trigger_properties_request` - (Optional, List) 
Nested scheme for **create_tekton_pipeline_trigger_properties_request**:
	* `name` - (Optional, String) Property name.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
	* `options` - (Optional, Map) Options for SINGLE_SELECT property type.
	* `path` - (Optional, String) property path for INTEGRATION type properties.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `properties` - (Optional, List) Trigger properties list.
	Nested scheme for **properties**:
		* `href` - (Optional, String) General href URL.
		  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `name` - (Required, String) Property name.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-zA-Z_.]{1,234}$/`.
		* `options` - (Optional, Map) Options for SINGLE_SELECT property type.
		* `path` - (Optional, String) property path for INTEGRATION type properties.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
		* `type` - (Required, String) Property type.
		  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`.
		* `value` - (Optional, String) String format property value.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
	* `type` - (Optional, String) Property type.
	  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`.
	* `value` - (Optional, String) String format property value.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
* `pipeline_id` - (Required, Forces new resource, String) The tekton pipeline ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.
* `trigger_id` - (Required, Forces new resource, String) The trigger ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[-0-9a-z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the tekton_pipeline_trigger_property.
* `href` - (Optional, String) General href URL.
  * Constraints: The maximum length is `2048` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `options` - (Optional, Map) Options for SINGLE_SELECT property type.
* `path` - (Optional, String) property path for INTEGRATION type properties.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.
* `type` - (Required, String) Property type.
  * Constraints: Allowable values are: `SECURE`, `TEXT`, `INTEGRATION`, `SINGLE_SELECT`.
* `value` - (Optional, String) String format property value.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/./`.

## Import

You can import the `ibm_tekton_pipeline_trigger_property` resource by using `name`.
The `name` property can be formed from `pipeline_id`, `trigger_id`, and `property_name` in the following format:

```
<pipeline_id>/<trigger_id>/<property_name>
```
* `pipeline_id`: A string in the format `94619026-912b-4d92-8f51-6c74f0692d90`. The tekton pipeline ID.
* `trigger_id`: A string in the format `1bb892a1-2e04-4768-a369-b1159eace147`. The trigger ID.
* `property_name`: A string in the format `debug-pipeline`. The property's name.

# Syntax
```
$ terraform import ibm_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property <pipeline_id>/<trigger_id>/<property_name>
```
