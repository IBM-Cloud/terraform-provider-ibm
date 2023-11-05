---
layout: "ibm"
page_title: "IBM : ibm_code_engine_job_run"
description: |-
  Get information about code_engine_job_run
subcategory: "Code Engine"
---

# ibm_code_engine_job_run

Provides a read-only data source for code_engine_job_run. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_job_run" "code_engine_job_run" {
	project_id = data.ibm_code_engine_project.code_engine_project.project_id
	name       = "my-job-run"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Required, Forces new resource, String) The name of your job run.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the code_engine_job_run.
* `created_at` - (String) The timestamp when the resource was created.

* `href` - (String) When you provision a new job,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `image_reference` - (String) The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z0-9][a-z0-9\\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\\-_.\/]+[a-z0-9](:[\\w][\\w.\\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$/`.

* `image_secret` - (String) The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.

* `job_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `resource_type` - (String) The type of the job.
  * Constraints: Allowable values are: `job_run_v2`.

* `run_arguments` - (List) Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container.
  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `100` items. The minimum length is `0` items.

* `run_as_user` - (Integer) The user ID (UID) to run the job (e.g., 1001).

* `run_commands` - (List) Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container.
  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `100` items. The minimum length is `0` items.

* `run_env_variables` - (List) References to config maps, secrets or literal values, which are exposed as environment variables in the job run.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **run_env_variables**:
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

* `run_mode` - (String) The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed.
  * Constraints: The default value is `task`. Allowable values are: `task`, `daemon`. The minimum length is `0` characters. The value must match regular expression `/^(task|daemon)$/`.

* `run_service_account` - (String) The name of the service account. For built-in service accounts, you can use the shortened names `manager`, `none`, `reader`, and `writer`. This property must not be set on a job run, which references a job template.
  * Constraints: The default value is `default`. Allowable values are: `default`, `manager`, `reader`, `writer`, `none`. The minimum length is `0` characters. The value must match regular expression `/^(manager|reader|writer|none|default)$/`.

* `run_volume_mounts` - (List) Optional mounts of config maps or a secrets.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested scheme for **run_volume_mounts**:
	* `mount_path` - (String) The path that should be mounted.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\/([^\/\\0]+\/?)+$/`.
	* `name` - (String) The name of the mount.
	  * Constraints: The maximum length is `63` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-z]([-a-z0-9]*[a-z0-9])?$/`.
	* `reference` - (String) The name of the referenced secret or config map.
	  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
	* `type` - (String) Specify the type of the volume mount. Allowed types are: 'config_map', 'secret'.
	  * Constraints: The default value is `secret`. Allowable values are: `config_map`, `secret`. The value must match regular expression `/^(config_map|secret)$/`.

* `scale_array_spec` - (String) Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges like `5,12-14,23,27`. Each instance can pick up its array index via environment variable `JOB_INDEX`. The number of unique array indices specified here determines the number of job instances to run.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d)(?:-(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d))?(?:,(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d)(?:-(?:[1-9]\\d\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d\\d|[1-9]\\d\\d\\d\\d|[1-9]\\d\\d\\d|[1-9]\\d\\d|[1-9]?\\d))?)*$/`.

* `scale_cpu_limit` - (String) Optional amount of CPU set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo).
  * Constraints: The default value is `1`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.

* `scale_ephemeral_storage_limit` - (String) Optional amount of ephemeral storage to set for the instance of the job. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `400M`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.

* `scale_max_execution_time` - (Integer) The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`.

* `scale_memory_limit` - (String) Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements).
  * Constraints: The default value is `4G`. The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^([0-9.]+)([eEinumkKMGTPB]*)$/`.

* `scale_retry_limit` - (Integer) The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`.

* `status` - (String) The current status of the job run.
  * Constraints: Allowable values are: `failed`, `completed`, `running`

* `status_details` - (Map) The detailed status of the job run.
  * `start_time` - (String) Time the job run started.
  * `completion_time` - (String) Time the job run completed.
  * `failed` - (Integer) Number of failed job run instances.
  * `pending` - (Integer) Number of pending job run instances.
  * `requested` - (Integer) Number of requested job run instances.
  * `running` - (Integer) Number of running job run instances.
  * `succeeded` - (Integer) Number of succeeded job run instances.
  * `unknown` - (Integer) Number of job run instances with unknown state.