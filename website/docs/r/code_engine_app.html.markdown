---
layout: "ibm"
page_title: "IBM : ibm_code_engine_app"
description: |-
  Manages code_engine_app.
subcategory: "Code Engine"
---

# ibm_code_engine_app

Create, update, and delete code_engine_apps with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_app" "code_engine_app_instance" {
  project_id      = ibm_code_engine_project.code_engine_project_instance.project_id
  name            = "my-app"
  image_reference = "icr.io/codeengine/helloworld"

  run_env_variables {
    type  = "literal"
    name  = "name"
    value = "value"
  }

  run_env_variables {
    type      = "secret_full_reference"
    name      = "secret_env_var"
    reference = "secret_name"
  }
}
```

## Timeouts

code_engine_app provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating a code_engine_app.
* `update` - (Default 10 minutes) Used for updating a code_engine_app.

## Argument Reference

You can specify the following arguments for this resource.

* `image_port` - (Optional, Integer) Optional port the app listens on. While the app will always be exposed via port `443` for end users, this port is used to connect to the port that is exposed by the container image.
  * Constraints: The default value is `8080`.
* `image_reference` - (Required, String) The name of the image that is used for this app. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z0-9][a-z0-9\\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\\-_.\/]+[a-z0-9](:[\\w][\\w.\\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$/`.
* `image_secret` - (Optional, String) Optional name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the app will be created but cannot reach the ready status, until this property is provided, too.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `managed_domain_mappings` - (Optional, String) Optional value controlling which of the system managed domain mappings will be setup for the application. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports application private visibility.
  * Constraints: The default value is `local_public`. Allowable values are: `local`, `local_private`, `local_public`.
* `name` - (Required, Forces new resource, String) The name of the app.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
* `probe_liveness` - (Optional, List) Response model for probes.
Nested schema for **probe_liveness**:
	* `failure_threshold` - (Optional, Integer) The number of consecutive, unsuccessful checks for the probe to be considered failed.
	  * Constraints: The default value is `1`. The maximum value is `10`. The minimum value is `1`.
	* `initial_delay` - (Optional, Integer) The amount of time in seconds to wait before the first probe check is performed.
	  * Constraints: The maximum value is `10`. The minimum value is `1`.
	* `interval` - (Optional, Integer) The amount of time in seconds between probe checks.
	  * Constraints: The default value is `10`. The maximum value is `60`. The minimum value is `1`.
	* `path` - (Optional, String) The path of the HTTP request to the resource. A path is only supported for a probe with a `type` of `http`.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$/`.
	* `port` - (Optional, Integer) The port on which to probe the resource.
	  * Constraints: The maximum value is `65535`. The minimum value is `1`.
	* `timeout` - (Optional, Integer) The amount of time in seconds that the probe waits for a response from the application before it times out and fails.
	  * Constraints: The default value is `1`. The maximum value is `3600`. The minimum value is `1`.
	* `type` - (Optional, String) Specifies whether to use HTTP or TCP for the probe checks. The default is TCP.
	  * Constraints: Allowable values are: `tcp`, `http`.
* `probe_readiness` - (Optional, List) Response model for probes.
Nested schema for **probe_readiness**:
	* `failure_threshold` - (Optional, Integer) The number of consecutive, unsuccessful checks for the probe to be considered failed.
	  * Constraints: The default value is `1`. The maximum value is `10`. The minimum value is `1`.
	* `initial_delay` - (Optional, Integer) The amount of time in seconds to wait before the first probe check is performed.
	  * Constraints: The maximum value is `10`. The minimum value is `1`.
	* `interval` - (Optional, Integer) The amount of time in seconds between probe checks.
	  * Constraints: The default value is `10`. The maximum value is `60`. The minimum value is `1`.
	* `path` - (Optional, String) The path of the HTTP request to the resource. A path is only supported for a probe with a `type` of `http`.
	  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^\/(([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})+(\/([a-zA-Z0-9-._~!$&'()*+,;=:@]|%[a-fA-F0-9]{2})*)*)?(\\?([a-zA-Z0-9-._~!$&'()*+,;=:@\/?]|%[a-fA-F0-9]{2})*)?$/`.
	* `port` - (Optional, Integer) The port on which to probe the resource.
	  * Constraints: The maximum value is `65535`. The minimum value is `1`.
	* `timeout` - (Optional, Integer) The amount of time in seconds that the probe waits for a response from the application before it times out and fails.
	  * Constraints: The default value is `1`. The maximum value is `3600`. The minimum value is `1`.
	* `type` - (Optional, String) Specifies whether to use HTTP or TCP for the probe checks. The default is TCP.
	  * Constraints: Allowable values are: `tcp`, `http`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `run_arguments` - (Optional, List) Optional arguments for the app that are passed to start the container. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container.
  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `run_as_user` - (Optional, Integer) Optional user ID (UID) to run the app.
  * Constraints: The default value is `0`.
* `run_commands` - (Optional, List) Optional commands for the app that are passed to start the container. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container.
  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `run_env_variables` - (Optional, List) References to config maps, secrets or literal values, which are exposed as environment variables in the application.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **run_env_variables**:
	* `key` - (Optional, String) The key to reference as environment variable.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-._a-zA-Z0-9]+$/`.
	* `name` - (Optional, String) The name of the environment variable.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-._a-zA-Z0-9]+$/`.
	* `prefix` - (Optional, String) A prefix that can be added to all keys of a full secret or config map reference.
	  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z_][a-zA-Z0-9_]*$/`.
	* `reference` - (Optional, String) The name of the secret or config map.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
	* `type` - (Optional, String) Specify the type of the environment variable.
	  * Constraints: The default value is `literal`. Allowable values are: `literal`, `config_map_full_reference`, `secret_full_reference`, `config_map_key_reference`, `secret_key_reference`. The value must match regular expression `/^(literal|config_map_full_reference|secret_full_reference|config_map_key_reference|secret_key_reference)$/`. When referencing a secret or configmap, the `reference` must be specified. When referencing a secret or configmap key, a `key` must also be specified.
	* `value` - (Optional, String) The literal value of the environment variable.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-._a-zA-Z0-9]+$/`.
* `run_service_account` - (Optional, String) Optional name of the service account. For built-in service accounts, you can use the shortened names `manager` , `none`, `reader`, and `writer`.
  * Constraints: The default value is `default`. Allowable values are: `default`, `manager`, `reader`, `writer`, `none`. The minimum length is `0` characters. The value must match regular expression `/^(manager|reader|writer|none|default)$/`.
* `run_volume_mounts` - (Optional, List) Mounts of config maps or secrets.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **run_volume_mounts**:
	* `mount_path` - (Required, String) The path that should be mounted.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\/([^\/\\0]+\/?)+$/`.
	* `name` - (Required, String) The name of the mount.
	  * Constraints: The maximum length is `63` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
	* `reference` - (Required, String) The name of the referenced secret or config map.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
	* `type` - (Required, String) Specify the type of the volume mount. Allowed types are: 'config_map', 'secret'.
	  * Constraints: The default value is `secret`. Allowable values are: `config_map`, `secret`. The value must match regular expression `/^(config_map|secret)$/`.
* `scale_concurrency` - (Optional, Integer) Optional maximum number of requests that can be processed concurrently per instance.
  * Constraints: The default value is `100`.
* `scale_concurrency_target` - (Optional, Integer) Optional threshold of concurrent requests per instance at which one or more additional instances are created. Use this value to scale up instances based on concurrent number of requests. This option defaults to the value of the `scale_concurrency` option, if not specified.
  * Constraints: The default value is `100`. The maximum value is `1000`. The minimum value is `1`.
* `scale_cpu_limit` - (Optional, String) Optional number of CPU set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo).
  * Constraints: The default value is `1`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.
* `scale_down_delay` - (Optional, Integer) Optional amount of time in seconds that delays the scale-down behavior for an app instance.
  * Constraints: The default value is `0`. The maximum value is `3600`. The minimum value is `0`.
* `scale_ephemeral_storage_limit` - (Optional, String) Optional amount of ephemeral storage to set for the instance of the app. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `400M`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.
* `scale_initial_instances` - (Optional, Integer) Optional initial number of instances that are created upon app creation or app update.
  * Constraints: The default value is `1`.
* `scale_max_instances` - (Optional, Integer) Optional maximum number of instances for this app. If you set this value to `0`, this property does not set a upper scaling limit. However, the app scaling is still limited by the project quota for instances. See [Limits and quotas for Code Engine](https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits).
  * Constraints: The default value is `10`.
* `scale_memory_limit` - (Optional, String) Optional amount of memory set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `4G`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.
* `scale_min_instances` - (Optional, Integer) Optional minimum number of instances for this app. If you set this value to `0`, the app will scale down to zero, if not hit by any request for some time.
  * Constraints: The default value is `0`.
* `scale_request_timeout` - (Optional, Integer) Optional amount of time in seconds that is allowed for a running app to respond to a request.
  * Constraints: The default value is `300`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_app.
* `app_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `build` - (String) Reference to a build that is associated with the application.
* `build_run` - (String) Reference to a build run that is associated with the application.
* `computed_env_variables` - (List) References to config maps, secrets or literal values, which are defined and set by Code Engine and are exposed as environment variables in the application.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **computed_env_variables**:
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
* `created_at` - (String) The timestamp when the resource was created.
* `endpoint` - (String) Optional URL to invoke the app. Depending on visibility,  this is accessible publicly or in the private network only. Empty in case 'managed_domain_mappings' is set to 'local'.
* `endpoint_internal` - (String) The URL to the app that is only visible within the project.
* `entity_tag` - (String) The version of the app instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `href` - (String) When you provision a new app,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.
* `resource_type` - (String) The type of the app.
  * Constraints: Allowable values are: `app_v2`.
* `status` - (String) The current status of the app.
  * Constraints: Allowable values are: `ready`, `deploying`, `failed`, `warning`.
* `status_details` - (List) The detailed status of the application.
Nested schema for **status_details**:
	* `latest_created_revision` - (String) Latest app revision that has been created.
	* `latest_ready_revision` - (String) Latest app revision that reached a ready state.
	* `reason` - (String) Optional information to provide more context in case of a 'failed' or 'warning' status.
	  * Constraints: Allowable values are: `ready`, `deploying`, `waiting_for_resources`, `no_revision_ready`, `ready_but_latest_revision_failed`.
* `etag` - ETag identifier for code_engine_app.

## Import

You can import the `ibm_code_engine_app` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-app`. The name of the app.

# Syntax
<pre>
$ terraform import ibm_code_engine_app.code_engine_app &lt;project_id&gt;/&lt;name&gt;
</pre>
