# Example for CodeEngineV2

This example illustrates how to use the CodeEngineV2

The following types of resources are supported:

* code_engine_app
* code_engine_binding
* code_engine_build
* code_engine_config_map
* code_engine_domain_mapping
* code_engine_job
* code_engine_project
* code_engine_secret

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CodeEngineV2 resources

code_engine_project resource:

```hcl
resource "code_engine_project" "code_engine_project_instance" {
  name              = var.code_engine_project_name
  resource_group_id = var.code_engine_project_resource_group_id
}
```

code_engine_app resource:

```hcl
resource "ibm_code_engine_app" "code_engine_app_instance" {
  project_id      = var.code_engine_project_id
  image_reference = var.code_engine_app_image_reference
  name            = var.code_engine_app_name
  run_env_variables {
    type  = "literal"
    name  = "name"
    value = "value"
  }
}
```

code_engine_build resource:

```hcl
resource "ibm_code_engine_build" "code_engine_build_instance" {
  project_id    = var.code_engine_project_id
  name          = var.code_engine_build_name
  output_image  = var.code_engine_build_output_image
  output_secret = var.code_engine_build_output_secret
  source_url    = var.code_engine_build_source_url
  strategy_type = var.code_engine_build_strategy_type
}
```

code_engine_config_map resource:

```hcl
resource "code_engine_config_map" "code_engine_config_map_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_config_map_name
  data       = var.code_engine_config_map_data
}
```

code_engine_job resource:

```hcl
resource "ibm_code_engine_job" "code_engine_job_instance" {
  project_id      = var.code_engine_project_id
  image_reference = var.code_engine_job_image_reference
  name            = var.code_engine_job_name

  run_env_variables {
    type  = "literal"
    name  = "name"
    value = "value"
  }
}

```

code_engine_secret resource:

```hcl
resource "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = var.code_engine_project_id
  format     = var.code_engine_secret_format
  name       = var.code_engine_secret_name
  data       = var.code_engine_secret_data
}
```

code_engine_binding resource:

```hcl
resource "ibm_code_engine_binding" "code_engine_secret_instance" {
  project_id = var.code_engine_project_id
  component {
    name          = var.code_engine_binding_component_name
    resource_type = var.code_engine_binding_component_resource_type
  }

  prefix      = var.code_engine_binding_prefix
  secret_name = var.code_engine_binding_secret_name
}
```

code_engine_domain_mapping resource:

```hcl
resource "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_domain_mapping_name
  component {
    name          = var.code_engine_domain_mapping_component_name
    resource_type = var.code_engine_domain_mapping_component_resource_type
  }
  tls_secret = var.code_engine_binding_secret_name
}
```

## CodeEngineV2 Data sources

code_engine_project data source:

```hcl
data "code_engine_project" "code_engine_project_instance" {
  project_id = var.code_engine_project_id
}
```

code_engine_app data source:

```hcl
data "ibm_code_engine_app" "code_engine_app_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_app_name
}
```

code_engine_build data source:

```hcl
data "ibm_code_engine_build" "code_engine_build_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_build_name
}
```

code_engine_config_map data source:

```hcl
data "code_engine_config_map" "code_engine_config_map_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_config_map_name
}
```

code_engine_job data source:

```hcl
data "ibm_code_engine_job" "code_engine_job_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_job_name
}

```

code_engine_secret data source:

```hcl
data "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_secret_name
}
```

code_engine_binding data source:

```hcl
data "ibm_code_engine_binding" "code_engine_binding_instance" {
  project_id = var.code_engine_project_id
  binding_id = var.code_engine_binding_id
}
```

code_engine_domain_mapping data source:

```hcl
data "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_domain_mapping_name
}
```

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| project_id | The ID of the project. | `string` | true |
| image_reference | The name of the image that is used for this app or job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`. | `string` | true |
| name | The name of the resource. | `string` | true |
| image_port | Optional port the app listens on. While the app will always be exposed via port `443` for end users, this port is used to connect to the port that is exposed by the container image. | `number` | false |
| image_secret | Optional name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the app will be created but cannot reach the ready status, until this property is provided, too. | `string` | false |
| managed_domain_mappings | Optional value controlling which of the system managed domain mappings will be setup for the application. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports application private visibility. | `string` | false |
| run_arguments | Optional arguments for the app that are passed to start the container. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. | `list(string)` | false |
| run_as_user | Optional user ID (UID) to run the app/job (e.g., `1001`). | `number` | false |
| run_commands | Optional commands for the app/job that are passed to start the container. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. | `list(string)` | false |
| run_env_variables | Optional references to config maps, secrets or literal values that are exposed as environment variables within the running application. | `list()` | false |
| run_service_account | Optional name of the service account. For built-in service accounts, you can use the shortened names `manager` , `none`, `reader`, and `writer`. | `string` | false |
| run_volume_mounts | Optional mounts of config maps or a secrets. | `list()` | false |
| scale_concurrency | Optional maximum number of requests that can be processed concurrently per instance. | `number` | false |
| scale_concurrency_target | Optional threshold of concurrent requests per instance at which one or more additional instances are created. Use this value to scale up instances based on concurrent number of requests. This option defaults to the value of the `scale_concurrency` option, if not specified. | `number` | false |
| scale_cpu_limit | Optional number of CPU set for the instance of the app/job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). | `string` | false |
| scale_down_delay | Optional amount of time in seconds that delays the scale down behavior for an app instance. | `number` | false |
| scale_ephemeral_storage_limit | Optional amount of ephemeral storage to set for the instance of the app. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_initial_instances | Optional initial number of instances that are created upon app creation or app update. | `number` | false |
| scale_max_instances | Optional maximum number of instances for this app. If you set this value to `0`, this property does not set a upper scaling limit. However, the app scaling is still limited by the project quota for instances. See [Limits and quotas for Code Engine](https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits). | `number` | false |
| scale_memory_limit | Optional amount of memory set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_min_instances | Optional minimum number of instances for this app. If you set this value to `0`, the app will scale down to zero, if not hit by any request for some time. | `number` | false |
| scale_request_timeout | Optional amount of time in seconds that is allowed for a running app to respond to a request. | `number` | false |
| component | A reference to another component. | `` | true |
| prefix | Optional value that is set as prefix in the component that is bound. Will be generated if not provided. | `string` | true |
| secret_name | The service access secret that is binding to a component. | `string` | true |
| output_image | The name of the image. | `string` | true |
| output_secret | The secret that is required to access the image registry. Make sure that the secret is granted with push permissions towards the specified container registry namespace. | `string` | true |
| strategy_type | The strategy to use for building the image. | `string` | true |
| source_context_dir | Option directory in the repository that contains the buildpacks file or the Dockerfile. | `string` | false |
| source_revision | Commit, tag, or branch in the source repository to pull. This field is optional if the `source_type` is `git` and uses the HEAD of default branch if not specified. If the `source_type` value is `local`, this field must be omitted. | `string` | false |
| source_secret | Name of the secret that is used access the repository source. This field is optional if the `source_type` is `git`. Additionally, if the `source_url` points to a repository that requires authentication, the build will be created but cannot access any source code, until this property is provided, too. If the `source_type` value is `local`, this field must be omitted. | `string` | false |
| source_type | Specifies the type of source to determine if your build source is in a repository or based on local source code.* local - For builds from local source code.* git - For builds from git version controlled source code. | `string` | false |
| source_url | The URL of the code repository. This field is required if the `source_type` is `git`. If the `source_type` value is `local`, this field must be omitted. If the repository is publicly available you can provide a 'https' URL like `https://github.com/IBM/CodeEngine`. If the repository requires authentication, you need to provide a 'ssh' URL like `git@github.com:IBM/CodeEngine.git` along with a `source_secret` that points to a secret of format `ssh_auth`. | `string` | false |
| strategy_size | Optional size for the build, which determines the amount of resources used. Build sizes are `small`, `medium`, `large`, `xlarge`. | `string` | false |
| strategy_spec_file | Optional path to the specification file that is used for build strategies for building an image. | `string` | false |
| timeout | The maximum amount of time, in seconds, that can pass before the build must succeed or fail. | `number` | false |
| data | The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters. | `map(string)` | false |
| image_reference | The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`. | `string` | true |
| image_secret | The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template. | `string` | false |
| run_mode | The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed. | `string` | false |
| scale_array_spec | Define a custom set of array indices as comma-separated list containing single values and hyphen-separated ranges like `5,12-14,23,27`. Each instance can pick up its array index via environment variable `JOB_INDEX`. The number of unique array indices specified here determines the number of job instances to run. | `string` | false |
| scale_max_execution_time | The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`. | `number` | false |
| scale_memory_limit | Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_retry_limit | The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`. | `number` | false |
| resource_group_id | Optional ID of the resource group for your project deployment. If this field is not defined, the default resource group of the account will be used. | `string` | false |
| format | Specify the format of the secret. | `string` | true |
| service_access | Properties for Service Access Secret Prototypes. | `` | false |
| tls_secret | The name of the TLS secret that holds the certificate and private key of the domain mapping. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| code_engine_project | code_engine_project object |
| code_engine_app | code_engine_app object |
| code_engine_binding | code_engine_binding object |
| code_engine_build | code_engine_build object |
| code_engine_config_map | code_engine_config_map object |
| code_engine_domain_mapping | code_engine_domain_mapping object |
| code_engine_job | code_engine_job object |
| code_engine_secret | code_engine_secret object |
