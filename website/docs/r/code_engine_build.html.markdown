---
layout: "ibm"
page_title: "IBM : ibm_code_engine_build"
description: |-
  Manages code_engine_build.
subcategory: "Code Engine"
---

# ibm_code_engine_build

Create, update, and delete code_engine_builds with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_build" "code_engine_build_instance" {
  project_id    = ibm_code_engine_project.code_engine_project_instance.project_id
  name          = "my-build"
  output_image  = "private.de.icr.io/icr_namespace/image-name"
  output_secret = "ce-auto-icr-private-eu-de"
  source_url    = "https://github.com/IBM/CodeEngine"
  strategy_type = "dockerfile"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, Forces new resource, String) The name of the build.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?$/`.
* `output_image` - (Required, String) The name of the image.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z0-9][a-z0-9\\-_.]+[a-z0-9][\/])?([a-z0-9][a-z0-9\\-_]+[a-z0-9][\/])?[a-z0-9][a-z0-9\\-_.\/]+[a-z0-9](:[\\w][\\w.\\-]{0,127})?(@sha256:[a-fA-F0-9]{64})?$/`.
* `output_secret` - (Required, String) The secret that is required to access the image registry. Make sure that the secret is granted with push permissions towards the specified container registry namespace.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `source_context_dir` - (Optional, String) Optional directory in the repository that contains the buildpacks file or the Dockerfile.
  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^(.*)+$/`.
* `source_revision` - (Optional, String) Commit, tag, or branch in the source repository to pull. This field is optional if the `source_type` is `git` and uses the HEAD of default branch if not specified. If the `source_type` value is `local`, this field must be omitted.
  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^[\\S]*$/`.
* `source_secret` - (Optional, String) Name of the secret that is used access the repository source. This field is optional if the `source_type` is `git`. Additionally, if the `source_url` points to a repository that requires authentication, the build will be created but cannot access any source code, until this property is provided, too. If the `source_type` value is `local`, this field must be omitted.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `source_type` - (Optional, String) Specifies the type of source to determine if your build source is in a repository or based on local source code.* local - For builds from local source code.* git - For builds from git version controlled source code.
  * Constraints: The default value is `git`. Allowable values are: `local`, `git`.
* `source_url` - (Optional, String) The URL of the code repository. This field is required if the `source_type` is `git`. If the `source_type` value is `local`, this field must be omitted. If the repository is publicly available you can provide a 'https' URL like `https://github.com/IBM/CodeEngine`. If the repository requires authentication, you need to provide a 'ssh' URL like `git@github.com:IBM/CodeEngine.git` along with a `source_secret` that points to a secret of format `ssh_auth`.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^((https:\/\/[a-z0-9]([\\-.]?[a-z0-9])+(:\\d{1,5})?)|((ssh:\/\/)?git@[a-z0-9]([\\-.]{0,1}[a-z0-9])+(:[a-zA-Z0-9\/][\\w\\-.]*)?))(\/([\\w\\-.]|%20)+)*$/`.
* `strategy_size` - (Optional, String) Optional size for the build, which determines the amount of resources used. Build sizes are `small`, `medium`, `large`, `xlarge`, `xxlarge`.
  * Constraints: The default value is `medium`. The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/[\\S]*/`.
* `strategy_spec_file` - (Optional, String) Optional path to the specification file that is used for build strategies for building an image.
  * Constraints: The default value is `Dockerfile`. The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[\\S]*$/`.
* `strategy_type` - (Required, String) The strategy to use for building the image.
  * Constraints: The default value is `dockerfile`. The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/[\\S]*/`.
* `timeout` - (Optional, Integer) The maximum amount of time, in seconds, that can pass before the build must succeed or fail.
  * Constraints: The default value is `600`. The maximum value is `3600`. The minimum value is `1`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_build.
* `build_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the build instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `href` - (String) When you provision a new build,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.
* `resource_type` - (String) The type of the build.
  * Constraints: Allowable values are: `build_v2`.
* `status` - (String) The current status of the build.
  * Constraints: Allowable values are: `ready`, `failed`.
* `status_details` - (List) The detailed status of the build.
Nested schema for **status_details**:
	* `reason` - (String) Optional information to provide more context in case of a 'failed' or 'warning' status.
	  * Constraints: Allowable values are: `registered`, `strategy_not_found`, `cluster_build_strategy_not_found`, `set_owner_reference_failed`, `spec_source_secret_not_found`, `spec_output_secret_ref_not_found`, `spec_runtime_secret_ref_not_found`, `multiple_secret_ref_not_found`, `runtime_paths_can_not_be_empty`, `remote_repository_unreachable`, `failed`.

* `etag` - ETag identifier for code_engine_build.

## Import

You can import the `ibm_code_engine_build` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-build`. The name of the build.

# Syntax
<pre>
$ terraform import ibm_code_engine_build.code_engine_build &lt;project_id&gt;/&lt;name&gt;
</pre>
