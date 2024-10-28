variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "ibmcloud_region" {
  description = "IBM Cloud Region"
  type        = string
}

// Resource arguments for code_engine_project
variable "code_engine_project_name" {
  description = "The name of the project."
  type        = string
  default     = "my-terraform-project"
}

// Resource arguments for code_engine_app
variable "code_engine_app_image_reference" {
  description = "The name of the image that is used for this app. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`."
  type        = string
  default     = "icr.io/codeengine/helloworld"
}
variable "code_engine_app_name" {
  description = "The name of the app. Use a name that is unique within the project."
  type        = string
  default     = "my-app"
}

// Resource arguments for code_engine_build
variable "code_engine_build_name" {
  description = "The name of the build. Use a name that is unique within the project."
  type        = string
  default     = "my-build"
}
variable "code_engine_build_output_image" {
  description = "The name of the image."
  type        = string
  default     = "private.de.icr.io/icr_namespace/image-name"
}
variable "code_engine_build_output_secret" {
  description = "The secret that is required to access the image registry. Make sure that the secret is granted with push permissions towards the specified container registry namespace."
  type        = string
  default     = "ce-auto-icr-private-eu-de"
}
variable "code_engine_build_source_url" {
  description = "The URL of the code repository. This field is required if the `source_type` is `git`. If the `source_type` value is `local`, this field must be omitted. If the repository is publicly available you can provide a 'https' URL like `https://github.com/IBM/CodeEngine`. If the repository requires authentication, you need to provide a 'ssh' URL like `git@github.com:IBM/CodeEngine.git` along with a `source_secret` that points to a secret of format `ssh_auth`."
  type        = string
  default     = "https://github.com/IBM/CodeEngine"
}
variable "code_engine_build_strategy_type" {
  description = "The strategy to use for building the image."
  type        = string
  default     = "dockerfile"
}

// Resource arguments for code_engine_config_map
variable "code_engine_config_map_name" {
  description = "The name of the config map. Use a name that is unique within the project."
  type        = string
  default     = "my-config-map"
}
variable "code_engine_config_map_data" {
  description = "The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters."
  type        = map(string)
  default     = { "key" = "inner" }
}

// Resource arguments for code_engine_secret
variable "code_engine_secret_name" {
  description = "The name of the secret. Use a name that is unique within the project."
  type        = string
  default     = "my-generic-secret"
}
variable "code_engine_secret_format" {
  description = "The format of the secret. Use a name that is unique within the project."
  type        = string
  default     = "generic"
}
variable "code_engine_secret_data" {
  description = "The key-value pair for the secret. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters."
  type        = map(string)
  default     = { "key" = "inner" }
}

// Resource arguments for code_engine_job
variable "code_engine_job_image_reference" {
  description = "The name of the image that is used for this job. The format is `REGISTRY/NAMESPACE/REPOSITORY:TAG` where `REGISTRY` and `TAG` are optional. If `REGISTRY` is not specified, the default is `docker.io`. If `TAG` is not specified, the default is `latest`. If the image reference points to a registry that requires authentication, make sure to also specify the property `image_secret`."
  type        = string
  default     = "icr.io/codeengine/helloworld"
}
variable "code_engine_job_name" {
  description = "The name of the job. Use a name that is unique within the project."
  type        = string
  default     = "my-job"
}

// Resource arguments for code_engine_secret with format service_access
variable "code_engine_secret_service_access_name" {
  description = "The name of the service access secret"
  type        = string
  default     = "my-service-access"
}

variable "code_engine_secret_service_access_resource_key" {
  description = "The ID of a resource key to access a resource instance."
  type        = string
}

variable "code_engine_secret_service_access_service_instance" {
  description = "The ID of a service instance."
  type        = string
}

// Resource arguments for code_engine_binding
variable "code_engine_binding_prefix" {
  description = "The name of the service access secret"
  type        = string
  default     = "MY_PREFIX"
}

// Resource arguments for code_engine_domain_mapping
variable "code_engine_domain_mapping_name" {
  description = "The name of the domain mapping."
  type        = string
}

// Resource arguments for code_engine_function
variable "code_engine_function_name" {
  description = "The name of the function."
  type        = string
  default     = "my-function"
}
variable "code_engine_function_runtime" {
  description = "The runtime of the function."
  type        = string
  default     = "nodejs-20"
}
variable "code_engine_function_code_reference_file_path" {
  description = "The path to a file containing the source code."
  type        = string
}

// Resource arguments for code_engine_secret with format tls
variable "code_engine_secret_tls_name" {
  description = "The name of the tls secret."
  type        = string
  default     = "my-tls-secret"
}
variable "code_engine_secret_tls_key_file_path" {
  description = "The path to the .key file containing the private key of the TLS certificate."
  type        = string
}
variable "code_engine_secret_tls_crt_file_path" {
  description = "The path to the .crt file containing the signed TLS certificate."
  type        = string
}

// Data source arguments for code_engine_project
variable "code_engine_project_id" {
  description = "The ID of the project."
  type        = string
  default     = "15314cc3-85b4-4338-903f-c28cdee6d005"
}
