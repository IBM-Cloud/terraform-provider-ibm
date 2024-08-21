---
layout: "ibm"
page_title: "IBM : ibm_code_engine_binding"
description: |-
  Get information about code_engine_binding
subcategory: "Code Engine"
---

# ibm_code_engine_binding

Provides a read-only data source to retrieve information about a code_engine_binding. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_binding" "code_engine_binding" {
	binding_id = "a172ced-b5f21bc-71ba50c-1638604"
	project_id = ibm_code_engine_binding.code_engine_binding_instance.project_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `binding_id` - (Required, Forces new resource, String) The id of your binding.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/.+/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_binding.

* `component` - (List) A reference to another component.
Nested schema for **component**:
	* `name` - (String) The name of the referenced component.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?$/`.
	* `resource_type` - (String) The type of the referenced resource.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/.+/`.

* `href` - (String) When you provision a new binding,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `prefix` - (String) The value that is set as a prefix in the component that is bound.
  * Constraints: The maximum length is `31` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Z]([_A-Z0-9]*[A-Z0-9])*$/`.

* `resource_type` - (String) The type of the binding.
  * Constraints: Allowable values are: `binding_v2`.

* `secret_name` - (String) The service access secret that is bound to a component.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.

* `status` - (String) The current status of the binding.

