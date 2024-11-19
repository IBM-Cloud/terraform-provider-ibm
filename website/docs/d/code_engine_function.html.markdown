---
layout: "ibm"
page_title: "IBM : ibm_code_engine_function"
description: |-
  Get information about code_engine_function
subcategory: "Code Engine"
---

# ibm_code_engine_function

Provides a read-only data source to retrieve information about a code_engine_function. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_function" "code_engine_function" {
	name = ibm_code_engine_function.code_engine_function_instance.name
	project_id = ibm_code_engine_function.code_engine_function_instance.project_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Required, Forces new resource, String) The name of your function.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_function.

* `function_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `code_binary` - (Boolean) Specifies whether the code is binary or not. Defaults to false when `code_reference` is set to a data URL. When `code_reference` is set to a code bundle URL, this field is always true.

* `code_main` - (String) Specifies the name of the function that should be invoked.
  * Constraints: The default value is `main`. The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z_][a-zA-Z0-9_]*$/`.

* `code_reference` - (String) Specifies either a reference to a code bundle or the source code itself. To specify the source code, use the data URL scheme and include the source code as base64 encoded. The data URL scheme is defined in [RFC 2397](https://tools.ietf.org/html/rfc2397).
  * Constraints: The maximum length is `1048576` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z0-9][a-z0-9\\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\\-_.\/]+[a-z0-9](:[\\w][\\w.\\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$|data:([-\\w]+\/[-+\\w.]+)?(;?\\w+=[-\\w]+)*;base64,.*/`.

* `code_secret` - (String) The name of the secret that is used to access the specified `code_reference`. The secret is used to authenticate with a non-public endpoint that is specified as`code_reference`.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.

* `created_at` - (String) The timestamp when the resource was created.

* `endpoint` - (String) URL to invoke the function.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `endpoint_internal` - (String) URL to function that is only visible within the project.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `entity_tag` - (String) The version of the function instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.

* `href` - (String) When you provision a new function, a relative URL path is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `managed_domain_mappings` - (String) Optional value controlling which of the system managed domain mappings will be setup for the function. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports function private visibility.
  * Constraints: The default value is `local_public`. Allowable values are: `local`, `local_private`, `local_public`.

* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.

* `resource_type` - (String) The type of the function.
  * Constraints: Allowable values are: `function_v2`.

* `run_env_variables` - (List) References to config maps, secrets or literal values, which are defined by the function owner and are exposed as environment variables in the function.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **run_env_variables**:
	* `key` - (String) The key to reference as environment variable.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-._a-zA-Z0-9]+$/`.
	* `name` - (String) The name of the environment variable.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-._a-zA-Z0-9]+$/`.
	* `prefix` - (String) A prefix that can be added to all keys of a full secret or config map reference.
	  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z_][a-zA-Z0-9_]*$/`.
	* `reference` - (String) The name of the secret or config map.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
	* `type` - (String) Specify the type of the environment variable.
	  * Constraints: The default value is `literal`. Allowable values are: `literal`, `config_map_full_reference`, `secret_full_reference`, `config_map_key_reference`, `secret_key_reference`. The value must match regular expression `/^(literal|config_map_full_reference|secret_full_reference|config_map_key_reference|secret_key_reference)$/`.
	* `value` - (String) The literal value of the environment variable.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-._a-zA-Z0-9]+$/`.

* `runtime` - (String) The managed runtime used to execute the injected code.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]*\\-[0-9]*(\\.[0-9]*)?$/`.

* `scale_concurrency` - (Integer) Number of parallel requests handled by a single instance, supported only by Node.js, default is `1`.
  * Constraints: The default value is `1`. The maximum value is `100`. The minimum value is `1`.

* `scale_cpu_limit` - (String) Optional amount of CPU set for the instance of the function. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo).
  * Constraints: The default value is `1`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.

* `scale_down_delay` - (Integer) Optional amount of time in seconds that delays the scale down behavior for a function.
  * Constraints: The default value is `1`. The maximum value is `600`. The minimum value is `0`.

* `scale_max_execution_time` - (Integer) Timeout in secs after which the function is terminated.
  * Constraints: The default value is `60`. The maximum value is `120`. The minimum value is `1`.

* `scale_memory_limit` - (String) Optional amount of memory set for the instance of the function. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `4G`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.

* `status` - (String) The current status of the function.
  * Constraints: Allowable values are: `offline`, `deploying`, `ready`, `failed`.

* `status_details` - (List) The detailed status of the function.
Nested schema for **status_details**:
	* `reason` - (String) Provides additional information about the status of the function.
	  * Constraints: Allowable values are: `offline`, `deploying_configuring_routes`, `ready_update_in_progress`, `deploying`, `ready_last_update_failed`, `ready`, `unknown_reason`, `no_code_bundle`.

