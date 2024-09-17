# Examples for Code Engine

These examples illustrate how to use the resources and data sources associated with Code Engine.

The following resources are supported:
* ibm_code_engine_app
* ibm_code_engine_binding
* ibm_code_engine_build
* ibm_code_engine_config_map
* ibm_code_engine_domain_mapping
* ibm_code_engine_function
* ibm_code_engine_job
* ibm_code_engine_project
* ibm_code_engine_secret

The following data sources are supported:
* ibm_code_engine_app
* ibm_code_engine_binding
* ibm_code_engine_build
* ibm_code_engine_config_map
* ibm_code_engine_domain_mapping
* ibm_code_engine_function
* ibm_code_engine_job
* ibm_code_engine_project
* ibm_code_engine_secret

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Code Engine resources

### Resource: ibm_code_engine_project

```hcl
resource "ibm_code_engine_project" "code_engine_project_instance" {
  name              = var.code_engine_project_name
  resource_group_id = var.code_engine_project_resource_group_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of the project. | `string` | true |
| resource_group_id | The ID of the resource group. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| account_id | An alphanumeric value identifying the account ID. |
| created_at | The timestamp when the project was created. |
| crn | The CRN of the project. |
| href | When you provision a new resource, a URL is created identifying the location of the instance. |
| region | The region for your project deployment. Possible values: `au-syd`, `br-sao`, `ca-tor`, `eu-de`, `eu-es`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`. |
| resource_type | The type of the project. |
| status | The current state of the project. For example, when the project is created and is ready for use, the status of the project is `active`. |

### Resource: ibm_code_engine_app

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

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| image_port | Optional port the app listens on. While the app will always be exposed via port `443` for end users, this port is used to connect to the port that is exposed by the container image. | `number` | false |
| image_reference | The name of the image that is used for this app. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`. | `string` | true |
| image_secret | Optional name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the app will be created but cannot reach the ready status, until this property is provided, too. | `string` | false |
| managed_domain_mappings | Optional value controlling which of the system managed domain mappings will be setup for the application. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports application private visibility. | `string` | false |
| name | The name of the app. | `string` | true |
| probe_liveness | Response model for probes. | `` | false |
| probe_readiness | Response model for probes. | `` | false |
| run_arguments | Optional arguments for the app that are passed to start the container. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. | `list(string)` | false |
| run_as_user | Optional user ID (UID) to run the app. | `number` | false |
| run_commands | Optional commands for the app that are passed to start the container. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. | `list(string)` | false |
| run_env_variables | References to config maps, secrets or literal values, which are exposed as environment variables in the application. | `list()` | false |
| run_service_account | Optional name of the service account. For built-in service accounts, you can use the shortened names `manager` , `none`, `reader`, and `writer`. | `string` | false |
| run_volume_mounts | Mounts of config maps or secrets. | `list()` | false |
| scale_concurrency | Optional maximum number of requests that can be processed concurrently per instance. | `number` | false |
| scale_concurrency_target | Optional threshold of concurrent requests per instance at which one or more additional instances are created. Use this value to scale up instances based on concurrent number of requests. This option defaults to the value of the `scale_concurrency` option, if not specified. | `number` | false |
| scale_cpu_limit | Optional number of CPU set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). | `string` | false |
| scale_down_delay | Optional amount of time in seconds that delays the scale-down behavior for an app instance. | `number` | false |
| scale_ephemeral_storage_limit | Optional amount of ephemeral storage to set for the instance of the app. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_initial_instances | Optional initial number of instances that are created upon app creation or app update. | `number` | false |
| scale_max_instances | Optional maximum number of instances for this app. If you set this value to `0`, this property does not set a upper scaling limit. However, the app scaling is still limited by the project quota for instances. See [Limits and quotas for Code Engine](https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits). | `number` | false |
| scale_memory_limit | Optional amount of memory set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_min_instances | Optional minimum number of instances for this app. If you set this value to `0`, the app will scale down to zero, if not hit by any request for some time. | `number` | false |
| scale_request_timeout | Optional amount of time in seconds that is allowed for a running app to respond to a request. | `number` | false |

#### Outputs

| Name | Description |
|------|-------------|
| build | Reference to a build that is associated with the application. |
| build_run | Reference to a build run that is associated with the application. |
| created_at | The timestamp when the resource was created. |
| endpoint | Optional URL to invoke the app. Depending on visibility,  this is accessible publicly or in the private network only. Empty in case 'managed_domain_mappings' is set to 'local'. |
| endpoint_internal | The URL to the app that is only visible within the project. |
| entity_tag | The version of the app instance, which is used to achieve optimistic locking. |
| href | When you provision a new app,  a URL is created identifying the location of the instance. |
| app_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the app. |
| status | The current status of the app. |
| status_details | The detailed status of the application. |

### Resource: ibm_code_engine_binding

```hcl
resource "ibm_code_engine_binding" "code_engine_binding_instance" {
  project_id  = var.code_engine_project_id
  prefix      = var.code_engine_binding_prefix
  secret_name = "my-service-access-secret"
  component {
    name          = var.code_engine_binding_component_name
    resource_type = var.code_engine_binding_component_resource_type
  }
}

```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| component | A reference to another component. | `` | true |
| prefix | The value that is set as a prefix in the component that is bound. | `string` | true |
| secret_name | The service access secret that is bound to a component. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | When you provision a new binding,  a URL is created identifying the location of the instance. |
| resource_type | The type of the binding. |
| status | The current status of the binding. |
| code_engine_binding_id | The ID of the binding. |

### Resource: ibm_code_engine_build

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

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of the build. | `string` | true |
| output_image | The name of the image. | `string` | true |
| output_secret | The secret that is required to access the image registry. Make sure that the secret is granted with push permissions towards the specified container registry namespace. | `string` | true |
| source_context_dir | Optional directory in the repository that contains the buildpacks file or the Dockerfile. | `string` | false |
| source_revision | Commit, tag, or branch in the source repository to pull. This field is optional if the `source_type` is `git` and uses the HEAD of default branch if not specified. If the `source_type` value is `local`, this field must be omitted. | `string` | false |
| source_secret | Name of the secret that is used access the repository source. This field is optional if the `source_type` is `git`. Additionally, if the `source_url` points to a repository that requires authentication, the build will be created but cannot access any source code, until this property is provided, too. If the `source_type` value is `local`, this field must be omitted. | `string` | false |
| source_type | Specifies the type of source to determine if your build source is in a repository or based on local source code.* local - For builds from local source code.* git - For builds from git version controlled source code. | `string` | false |
| source_url | The URL of the code repository. This field is required if the `source_type` is `git`. If the `source_type` value is `local`, this field must be omitted. If the repository is publicly available you can provide a 'https' URL like `https://github.com/IBM/CodeEngine`. If the repository requires authentication, you need to provide a 'ssh' URL like `git@github.com:IBM/CodeEngine.git` along with a `source_secret` that points to a secret of format `ssh_auth`. | `string` | false |
| strategy_size | Optional size for the build, which determines the amount of resources used. Build sizes are `small`, `medium`, `large`, `xlarge`, `xxlarge`. | `string` | false |
| strategy_spec_file | Optional path to the specification file that is used for build strategies for building an image. | `string` | false |
| strategy_type | The strategy to use for building the image. | `string` | true |
| timeout | The maximum amount of time, in seconds, that can pass before the build must succeed or fail. | `number` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the build instance, which is used to achieve optimistic locking. |
| href | When you provision a new build,  a URL is created identifying the location of the instance. |
| build_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the build. |
| status | The current status of the build. |
| status_details | The detailed status of the build. |

### Resource: ibm_code_engine_config_map

```hcl
resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
  project_id = var.code_engine_project_id
  name       = var.code_engine_config_map_name
  data       = var.code_engine_config_map_data
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| data | The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. | `map(string)` | false |
| name | The name of the config map. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the config map instance, which is used to achieve optimistic locking. |
| href | When you provision a new config map,  a URL is created identifying the location of the instance. |
| config_map_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the config map. |

### Resource: ibm_code_engine_domain_mapping

```hcl
resource "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
  project_id = var.code_engine_domain_mapping_project_id
  name       = var.code_engine_domain_mapping_name
  tls_secret = var.code_engine_domain_mapping_tls_secret
  component {
    name          = var.code_engine_domain_mapping_component_name
    resource_type = var.code_engine_domain_mapping_component_resource_type
  }
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| component | A reference to another component. | `` | true |
| name | The name of the domain mapping. | `string` | true |
| tls_secret | The name of the TLS secret that includes the certificate and private key of this domain mapping. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| cname_target | The value of the CNAME record that must be configured in the DNS settings of the domain, to route traffic properly to the target Code Engine region. |
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the domain mapping instance, which is used to achieve optimistic locking. |
| href | When you provision a new domain mapping, a URL is created identifying the location of the instance. |
| domain_mapping_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the Code Engine resource. |
| status | The current status of the domain mapping. |
| status_details | The detailed status of the domain mapping. |
| user_managed | Specifies whether the domain mapping is managed by the user or by Code Engine. |
| visibility | Specifies whether the domain mapping is reachable through the public internet, or private IBM network, or only through other components within the same Code Engine project. |

### Resource: ibm_code_engine_function

```hcl
resource "ibm_code_engine_function" "code_engine_function_instance" {
  project_id = var.code_engine_function_project_id
  name = var.code_engine_function_name
  runtime = var.code_engine_function_runtime
  code_reference = var.code_engine_function_code_reference
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| code_binary | Specifies whether the code is binary or not. Defaults to false when `code_reference` is set to a data URL. When `code_reference` is set to a code bundle URL, this field is always true. | `bool` | false |
| code_main | Specifies the name of the function that should be invoked. | `string` | false |
| code_reference | Specifies either a reference to a code bundle or the source code itself. To specify the source code, use the data URL scheme and include the source code as base64 encoded. The data URL scheme is defined in [RFC 2397](https://tools.ietf.org/html/rfc2397). | `string` | true |
| code_secret | The name of the secret that is used to access the specified `code_reference`. The secret is used to authenticate with a non-public endpoint that is specified as`code_reference`. | `string` | false |
| managed_domain_mappings | Optional value controlling which of the system managed domain mappings will be setup for the function. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports function private visibility. | `string` | false |
| name | The name of the function. | `string` | true |
| run_env_variables | References to config maps, secrets or literal values, which are defined by the function owner and are exposed as environment variables in the function. | `list()` | false |
| runtime | The managed runtime used to execute the injected code. | `string` | true |
| scale_concurrency | Number of parallel requests handled by a single instance, supported only by Node.js, default is `1`. | `number` | false |
| scale_cpu_limit | Optional amount of CPU set for the instance of the function. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). | `string` | false |
| scale_down_delay | Optional amount of time in seconds that delays the scale down behavior for a function. | `number` | false |
| scale_max_execution_time | Timeout in secs after which the function is terminated. | `number` | false |
| scale_memory_limit | Optional amount of memory set for the instance of the function. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| endpoint | URL to invoke the function. |
| endpoint_internal | URL to function that is only visible within the project. |
| entity_tag | The version of the function instance, which is used to achieve optimistic locking. |
| href | When you provision a new function, a relative URL path is created identifying the location of the instance. |
| function_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the function. |
| status | The current status of the function. |
| status_details | The detailed status of the function. |

### Resource: ibm_code_engine_job

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

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| image_reference | The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`. | `string` | true |
| image_secret | The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template. | `string` | false |
| name | The name of the job. | `string` | true |
| run_arguments | Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. | `list(string)` | false |
| run_as_user | The user ID (UID) to run the job. | `number` | false |
| run_commands | Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. | `list(string)` | false |
| run_env_variables | References to config maps, secrets or literal values, which are exposed as environment variables in the job run. | `list()` | false |
| run_mode | The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed. | `string` | false |
| run_service_account | The name of the service account. For built-in service accounts, you can use the shortened names `manager`, `none`, `reader`, and `writer`. This property must not be set on a job run, which references a job template. | `string` | false |
| run_volume_mounts | Optional mounts of config maps or secrets. | `list()` | false |
| scale_array_spec | Define a custom set of array indices as a comma-separated list containing single values and hyphen-separated ranges, such as  5,12-14,23,27. Each instance gets its array index value from the environment variable JOB_INDEX. The number of unique array indices that you specify with this parameter determines the number of job instances to run. | `string` | false |
| scale_cpu_limit | Optional amount of CPU set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). | `string` | false |
| scale_ephemeral_storage_limit | Optional amount of ephemeral storage to set for the instance of the job. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_max_execution_time | The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`. | `number` | false |
| scale_memory_limit | Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). | `string` | false |
| scale_retry_limit | The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`. | `number` | false |

#### Outputs

| Name | Description |
|------|-------------|
| build | Reference to a build that is associated with the job. |
| build_run | Reference to a build run that is associated with the job. |
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the job instance, which is used to achieve optimistic locking. |
| href | When you provision a new job,  a URL is created identifying the location of the instance. |
| job_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the job. |

### Resource: ibm_code_engine_secret

```hcl
resource "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = var.code_engine_project_id
  format     = var.code_engine_secret_format
  name       = var.code_engine_secret_name
  data       = var.code_engine_secret_data
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| data | Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters. | `map(string)` | false |
| format | Specify the format of the secret. | `string` | true |
| name | The name of the secret. | `string` | true |
| service_access | Properties for Service Access Secrets. | `` | false |
| service_operator | Properties for the IBM Cloud Operator Secret. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the secret instance, which is used to achieve optimistic locking. |
| href | When you provision a new secret,  a URL is created identifying the location of the instance. |
| secret_id | The identifier of the resource. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the secret. |

## Code Engine data sources

### Data source: ibm_code_engine_project

```hcl
data "ibm_code_engine_project" "code_engine_project_instance" {
  project_id = var.data_code_engine_project_code_engine_project_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| account_id | An alphanumeric value identifying the account ID. |
| created_at | The timestamp when the project was created. |
| crn | The CRN of the project. |
| href | When you provision a new resource, a URL is created identifying the location of the instance. |
| name | The name of the project. |
| region | The region for your project deployment. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_group_id | The ID of the resource group. |
| resource_type | The type of the project. |
| status | The current state of the project. For example, when the project is created and is ready for use, the status of the project is active. |

### Data source: ibm_code_engine_app

```hcl
data "ibm_code_engine_app" "code_engine_app_instance" {
  project_id = var.data_code_engine_app_project_id
  name       = var.data_code_engine_app_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your application. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| build | Reference to a build that is associated with the application. |
| build_run | Reference to a build run that is associated with the application. |
| created_at | The timestamp when the resource was created. |
| endpoint | Optional URL to invoke the app. Depending on visibility,  this is accessible publicly or in the private network only. Empty in case 'managed_domain_mappings' is set to 'local'. |
| endpoint_internal | The URL to the app that is only visible within the project. |
| entity_tag | The version of the app instance, which is used to achieve optimistic locking. |
| href | When you provision a new app,  a URL is created identifying the location of the instance. |
| image_port | Optional port the app listens on. While the app will always be exposed via port `443` for end users, this port is used to connect to the port that is exposed by the container image. |
| image_reference | The name of the image that is used for this app. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`. |
| image_secret | Optional name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the app will be created but cannot reach the ready status, until this property is provided, too. |
| managed_domain_mappings | Optional value controlling which of the system managed domain mappings will be setup for the application. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports application private visibility. |
| probe_liveness | Response model for probes. |
| probe_readiness | Response model for probes. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the app. |
| run_arguments | Optional arguments for the app that are passed to start the container. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. |
| run_as_user | Optional user ID (UID) to run the app. |
| run_commands | Optional commands for the app that are passed to start the container. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. |
| run_env_variables | References to config maps, secrets or literal values, which are exposed as environment variables in the application. |
| run_service_account | Optional name of the service account. For built-in service accounts, you can use the shortened names `manager` , `none`, `reader`, and `writer`. |
| run_volume_mounts | Mounts of config maps or secrets. |
| scale_concurrency | Optional maximum number of requests that can be processed concurrently per instance. |
| scale_concurrency_target | Optional threshold of concurrent requests per instance at which one or more additional instances are created. Use this value to scale up instances based on concurrent number of requests. This option defaults to the value of the `scale_concurrency` option, if not specified. |
| scale_cpu_limit | Optional number of CPU set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). |
| scale_down_delay | Optional amount of time in seconds that delays the scale-down behavior for an app instance. |
| scale_ephemeral_storage_limit | Optional amount of ephemeral storage to set for the instance of the app. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). |
| scale_initial_instances | Optional initial number of instances that are created upon app creation or app update. |
| scale_max_instances | Optional maximum number of instances for this app. If you set this value to `0`, this property does not set a upper scaling limit. However, the app scaling is still limited by the project quota for instances. See [Limits and quotas for Code Engine](https://cloud.ibm.com/docs/codeengine?topic=codeengine-limits). |
| scale_memory_limit | Optional amount of memory set for the instance of the app. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). |
| scale_min_instances | Optional minimum number of instances for this app. If you set this value to `0`, the app will scale down to zero, if not hit by any request for some time. |
| scale_request_timeout | Optional amount of time in seconds that is allowed for a running app to respond to a request. |
| status | The current status of the app. |
| status_details | The detailed status of the application. |

### Data source: ibm_code_engine_binding

```hcl
data "ibm_code_engine_binding" "code_engine_binding_instance" {
  project_id             = var.data_code_engine_binding_project_id
  code_engine_binding_id = var.data_code_engine_binding_code_engine_binding_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| code_engine_binding_id | The id of your binding. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| component | A reference to another component. |
| href | When you provision a new binding,  a URL is created identifying the location of the instance. |
| prefix | The value that is set as a prefix in the component that is bound. |
| resource_type | The type of the binding. |
| secret_name | The service access secret that is bound to a component. |
| status | The current status of the binding. |

### Data source: ibm_code_engine_build

```hcl
data "ibm_code_engine_build" "code_engine_build_instance" {
  project_id = var.data_code_engine_build_project_id
  name       = var.data_code_engine_build_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your build. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the build instance, which is used to achieve optimistic locking. |
| href | When you provision a new build,  a URL is created identifying the location of the instance. |
| output_image | The name of the image. |
| output_secret | The secret that is required to access the image registry. Make sure that the secret is granted with push permissions towards the specified container registry namespace. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the build. |
| source_context_dir | Optional directory in the repository that contains the buildpacks file or the Dockerfile. |
| source_revision | Commit, tag, or branch in the source repository to pull. This field is optional if the `source_type` is `git` and uses the HEAD of default branch if not specified. If the `source_type` value is `local`, this field must be omitted. |
| source_secret | Name of the secret that is used access the repository source. This field is optional if the `source_type` is `git`. Additionally, if the `source_url` points to a repository that requires authentication, the build will be created but cannot access any source code, until this property is provided, too. If the `source_type` value is `local`, this field must be omitted. |
| source_type | Specifies the type of source to determine if your build source is in a repository or based on local source code.* local - For builds from local source code.* git - For builds from git version controlled source code. |
| source_url | The URL of the code repository. This field is required if the `source_type` is `git`. If the `source_type` value is `local`, this field must be omitted. If the repository is publicly available you can provide a 'https' URL like `https://github.com/IBM/CodeEngine`. If the repository requires authentication, you need to provide a 'ssh' URL like `git@github.com:IBM/CodeEngine.git` along with a `source_secret` that points to a secret of format `ssh_auth`. |
| status | The current status of the build. |
| status_details | The detailed status of the build. |
| strategy_size | Optional size for the build, which determines the amount of resources used. Build sizes are `small`, `medium`, `large`, `xlarge`, `xxlarge`. |
| strategy_spec_file | Optional path to the specification file that is used for build strategies for building an image. |
| strategy_type | The strategy to use for building the image. |
| timeout | The maximum amount of time, in seconds, that can pass before the build must succeed or fail. |

### Data source: ibm_code_engine_config_map

```hcl
data "ibm_code_engine_config_map" "code_engine_config_map_instance" {
  project_id = var.data_code_engine_config_map_project_id
  name       = var.data_code_engine_config_map_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your configmap. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| data | The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. |
| entity_tag | The version of the config map instance, which is used to achieve optimistic locking. |
| href | When you provision a new config map,  a URL is created identifying the location of the instance. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the config map. |

### Data source: ibm_code_engine_domain_mapping

```hcl
data "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
  project_id = var.data_code_engine_domain_mapping_project_id
  name       = var.data_code_engine_domain_mapping_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your domain mapping. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| cname_target | The value of the CNAME record that must be configured in the DNS settings of the domain, to route traffic properly to the target Code Engine region. |
| component | A reference to another component. |
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the domain mapping instance, which is used to achieve optimistic locking. |
| href | When you provision a new domain mapping, a URL is created identifying the location of the instance. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the Code Engine resource. |
| status | The current status of the domain mapping. |
| status_details | The detailed status of the domain mapping. |
| tls_secret | The name of the TLS secret that includes the certificate and private key of this domain mapping. |
| user_managed | Specifies whether the domain mapping is managed by the user or by Code Engine. |
| visibility | Specifies whether the domain mapping is reachable through the public internet, or private IBM network, or only through other components within the same Code Engine project. |

### Data source: ibm_code_engine_function

```hcl
data "ibm_code_engine_function" "code_engine_function_instance" {
  project_id = var.data_code_engine_function_project_id
  name       = var.data_code_engine_function_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your function. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| code_binary | Specifies whether the code is binary or not. Defaults to false when `code_reference` is set to a data URL. When `code_reference` is set to a code bundle URL, this field is always true. |
| code_main | Specifies the name of the function that should be invoked. |
| code_reference | Specifies either a reference to a code bundle or the source code itself. To specify the source code, use the data URL scheme and include the source code as base64 encoded. The data URL scheme is defined in [RFC 2397](https://tools.ietf.org/html/rfc2397). |
| code_secret | The name of the secret that is used to access the specified `code_reference`. The secret is used to authenticate with a non-public endpoint that is specified as`code_reference`. |
| created_at | The timestamp when the resource was created. |
| endpoint | URL to invoke the function. |
| endpoint_internal | URL to function that is only visible within the project. |
| entity_tag | The version of the function instance, which is used to achieve optimistic locking. |
| href | When you provision a new function, a relative URL path is created identifying the location of the instance. |
| managed_domain_mappings | Optional value controlling which of the system managed domain mappings will be setup for the function. Valid values are 'local_public', 'local_private' and 'local'. Visibility can only be 'local_private' if the project supports function private visibility. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the function. |
| run_env_variables | References to config maps, secrets or literal values, which are defined by the function owner and are exposed as environment variables in the function. |
| runtime | The managed runtime used to execute the injected code. |
| scale_concurrency | Number of parallel requests handled by a single instance, supported only by Node.js, default is `1`. |
| scale_cpu_limit | Optional amount of CPU set for the instance of the function. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). |
| scale_down_delay | Optional amount of time in seconds that delays the scale down behavior for a function. |
| scale_max_execution_time | Timeout in secs after which the function is terminated. |
| scale_memory_limit | Optional amount of memory set for the instance of the function. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). |
| status | The current status of the function. |
| status_details | The detailed status of the function. |

### Data source: ibm_code_engine_job

```hcl
data "ibm_code_engine_job" "code_engine_job_instance" {
  project_id = var.data_code_engine_job_project_id
  name       = var.data_code_engine_job_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your job. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| build | Reference to a build that is associated with the job. |
| build_run | Reference to a build run that is associated with the job. |
| created_at | The timestamp when the resource was created. |
| entity_tag | The version of the job instance, which is used to achieve optimistic locking. |
| href | When you provision a new job,  a URL is created identifying the location of the instance. |
| image_reference | The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`. |
| image_secret | The name of the image registry access secret. The image registry access secret is used to authenticate with a private registry when you download the container image. If the image reference points to a registry that requires authentication, the job / job runs will be created but submitted job runs will fail, until this property is provided, too. This property must not be set on a job run, which references a job template. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the job. |
| run_arguments | Set arguments for the job that are passed to start job run containers. If not specified an empty string array will be applied and the arguments specified by the container image, will be used to start the container. |
| run_as_user | The user ID (UID) to run the job. |
| run_commands | Set commands for the job that are passed to start job run containers. If not specified an empty string array will be applied and the command specified by the container image, will be used to start the container. |
| run_env_variables | References to config maps, secrets or literal values, which are exposed as environment variables in the job run. |
| run_mode | The mode for runs of the job. Valid values are `task` and `daemon`. In `task` mode, the `max_execution_time` and `retry_limit` properties apply. In `daemon` mode, since there is no timeout and failed instances are restarted indefinitely, the `max_execution_time` and `retry_limit` properties are not allowed. |
| run_service_account | The name of the service account. For built-in service accounts, you can use the shortened names `manager`, `none`, `reader`, and `writer`. This property must not be set on a job run, which references a job template. |
| run_volume_mounts | Optional mounts of config maps or secrets. |
| scale_array_spec | Define a custom set of array indices as a comma-separated list containing single values and hyphen-separated ranges, such as  5,12-14,23,27. Each instance gets its array index value from the environment variable JOB_INDEX. The number of unique array indices that you specify with this parameter determines the number of job instances to run. |
| scale_cpu_limit | Optional amount of CPU set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). |
| scale_ephemeral_storage_limit | Optional amount of ephemeral storage to set for the instance of the job. The amount specified as ephemeral storage, must not exceed the amount of `scale_memory_limit`. The units for specifying ephemeral storage are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). |
| scale_max_execution_time | The maximum execution time in seconds for runs of the job. This property can only be specified if `run_mode` is `task`. |
| scale_memory_limit | Optional amount of memory set for the instance of the job. For valid values see [Supported memory and CPU combinations](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo). The units for specifying memory are Megabyte (M) or Gigabyte (G), whereas G and M are the shorthand expressions for GB and MB. For more information see [Units of measurement](https://cloud.ibm.com/docs/codeengine?topic=codeengine-mem-cpu-combo#unit-measurements). |
| scale_retry_limit | The number of times to rerun an instance of the job before the job is marked as failed. This property can only be specified if `run_mode` is `task`. |

### Data source: ibm_code_engine_secret

```hcl
data "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = var.data_code_engine_secret_project_id
  name       = var.data_code_engine_secret_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The ID of the project. | `string` | true |
| name | The name of your secret. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The timestamp when the resource was created. |
| data | Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters. |
| entity_tag | The version of the secret instance, which is used to achieve optimistic locking. |
| format | Specify the format of the secret. |
| href | When you provision a new secret,  a URL is created identifying the location of the instance. |
| region | The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'. |
| resource_type | The type of the secret. |
| service_access | Properties for Service Access Secrets. |
| service_operator | Properties for the IBM Cloud Operator Secret. |
