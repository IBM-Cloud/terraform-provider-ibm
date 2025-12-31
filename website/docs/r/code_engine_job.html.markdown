---
layout: "ibm"
page_title: "IBM : ibm_code_engine_job"
description: |-
  Manages code_engine_job.
subcategory: "Code Engine"
---

# ibm_code_engine_job

Create, update, and delete code_engine_jobs with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_job" "code_engine_job_instance" {
  project_id      = ibm_code_engine_project.code_engine_project_instance.project_id
  name            = "my-job"
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

## Argument Reference

You can specify the following arguments for this resource.

* `image_reference` - (Required, String) The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z0-9][a-z0-9\\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\\-_.\/]+[a-z0-9](:[\\w][\\w.\\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$/`.
* `image_secret` - (Optional, String) The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `name` - (Required, Forces new resource, String) The name of the job.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `run_arguments` - (Optional, List) Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container.
  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `run_as_user` - (Optional, Integer) The user ID (UID) to run the job.
  * Constraints: The default value is `0`.
* `run_commands` - (Optional, List) Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container.
  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `100` items. The minimum length is `0` items.
* `run_compute_resource_token_enabled` - (Optional, Boolean) Optional flag to enable the use of a compute resource token mounted to the container file system.
  * Constraints: The default value is `false`.
* `run_env_variables` - (Optional, List) References to config maps, secrets or literal values, which are defined by the function owner and are exposed as environment variables in the job run.
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
	* `type` - (Required, String) Specify the type of the environment variable.
	  * Constraints: The default value is `literal`. Allowable values are: `literal`, `config_map_full_reference`, `secret_full_reference`, `config_map_key_reference`, `secret_key_reference`. The value must match regular expression `/^(literal|config_map_full_reference|secret_full_reference|config_map_key_reference|secret_key_reference)$/`.
	* `value` - (Optional, String) The literal value of the environment variable.
	  * Constraints: The maximum length is `1048576` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `run_mode` - (Optional, String) The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed.
  * Constraints: The default value is `task`. Allowable values are: `task`, `daemon`. The value must match regular expression `/^(task|daemon)$/`.
* `run_service_account` - (Optional, String) The name of the service account. For built-in service accounts, you can use the shortened names `manager`, `none`, `reader`, and `writer`. This property must not be set on a job run, which references a job template.
  * Constraints: The default value is `default`. Allowable values are: `default`, `manager`, `reader`, `writer`, `none`. The value must match regular expression `/^(manager|reader|writer|none|default)$/`.
* `run_volume_mounts` - (Optional, List) Optional mounts of config maps or secrets.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **run_volume_mounts**:
	* `mount_path` - (Required, String) The path that should be mounted.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\/([^\/\\0]+\/?)+$/`.
	* `name` - (Optional, String) The name of the mount.
	  * Constraints: The maximum length is `63` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
	* `read_only` - (Optional, Boolean) Optional flag for a volume mount of type 'persistent_data_store' to specify whether it is read-only.
    * Constraints: The default value is `true`.
	* `reference` - (Required, String) The name of the referenced secret, config map, or persistent data store.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
	* `sub_path` - (Optional, String) The path mounted at the mount path.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `type` - (Required, String) Specify the type of the volume mount. Allowed types are: 'config_map', 'persistent_data_store', 'secret'.
	  * Constraints: The default value is `secret`. Allowable values are: `config_map`, `persistent_data_store`, `secret`. The value must match regular expression `/^(config_map|persistent_data_store|secret)$/`.
* `scale_array_spec` - (Optional, String) Define a custom set of array indices as a comma-separated list containing single values and hyphen-separated ranges, such as  5,12-14,23,27. Each instance gets its array index value from the environment variable JOB_INDEX. The number of unique array indices that you specify with this parameter determines the number of job instances to run.
  * Constraints: The default value is `0`. The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d)(?:-(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d))?(?:,(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d)(?:-(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d))?)*$/`.
* `scale_cpu_limit` - (Optional, String) Optional amount of CPU set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo).
  * Constraints: The default value is `1`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.
* `scale_ephemeral_storage_limit` - (Optional, String) Optional amount of ephemeral storage to set for the instance of the job. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `400M`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.
* `scale_max_execution_time` - (Optional, Integer) The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`.
  * Constraints: The default value is `7200`. The maximum value is `86400`. The minimum value is `1`.
* `scale_memory_limit` - (Optional, String) Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `4G`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.
* `scale_retry_limit` - (Optional, Integer) The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`.
  * Constraints: The default value is `3`. The maximum value is `5`. The minimum value is `0`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_job.
* `job_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `build` - (String) Reference to a build that is associated with the job.
* `build_run` - (String) Reference to a build run that is associated with the job.
* `computed_env_variables` - (List) References to config maps, secrets or literal values, which are defined and set by Code Engine and are exposed as environment variables in the job run.
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
	  * Constraints: The maximum length is `1048576` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the job instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `href` - (String) When you provision a new job,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/^(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.
* `resource_type` - (String) The type of the job.
  * Constraints: Allowable values are: `job_v2`.

* `etag` - ETag identifier for code_engine_job.

## Import

You can import the `ibm_code_engine_job` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-job`. The name of the job.

# Syntax
<pre>
$ terraform import ibm_code_engine_job.code_engine_job &lt;project_id&gt;/&lt;name&gt;
</pre>
