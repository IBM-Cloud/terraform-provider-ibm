---
layout: "ibm"
page_title: "IBM : ibm_code_engine_binding"
description: |-
  Manages code_engine_binding.
subcategory: "Code Engine"
---

# ibm_code_engine_binding

Create, update, and delete code_engine_bindings with this resource. A `secret` with format `service_access` is required to create a binding.

## Example Usage

```hcl
resource "ibm_code_engine_binding" "code_engine_binding_instance" {
  component {
		name = "my-app-1"
		resource_type = "app_v2"
  }
  prefix = "MY_COS"
  project_id = "15314cc3-85b4-4338-903f-c28cdee6d005"
  secret_name = "my-service-access"
}
```

## Timeouts

code_engine_binding provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 20 minutes) Used for creating a code_engine_binding.
* `delete` - (Default 20 minutes) Used for deleting a code_engine_binding.

## Argument Reference

You can specify the following arguments for this resource.

* `component` - (Required, Forces new resource, List) A reference to another component.
Nested schema for **component**:
	* `name` - (Required, String) The name of the referenced component.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?$/`.
	* `resource_type` - (Required, String) The type of the referenced resource.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/.+/`.
* `prefix` - (Required, Forces new resource, String) The value that is set as a prefix in the component that is bound.
  * Constraints: The maximum length is `31` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Z]([_A-Z0-9]*[A-Z0-9])*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `secret_name` - (Required, Forces new resource, String) The service access secret that is bound to a component.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_binding.
* `binding_id` - (String) The ID of the binding.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/.+/`.
* `href` - (String) When you provision a new binding,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `resource_type` - (String) The type of the binding.
  * Constraints: Allowable values are: `binding_v2`.
* `status` - (String) The current status of the binding.


## Import

You can import the `ibm_code_engine_binding` resource by using `id`.
The `id` property can be formed from `project_id`, and `binding_id` in the following format:

<pre>
&lt;project_id&gt;/&lt;binding_id&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `binding_id`: A string in the format `a172ced-b5f21bc-71ba50c-1638604`. The ID of the binding.

# Syntax
<pre>
$ terraform import ibm_code_engine_binding.code_engine_binding &lt;project_id&gt;/&lt;binding_id&gt;
</pre>
